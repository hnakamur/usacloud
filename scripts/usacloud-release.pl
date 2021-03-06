#!/usr/bin/env perl

=head DESCRIPTION

Release Engineering scripts for [myProject].
(copied from https://github.com/mackerelio/mackerel-agent/blob/master/_tools/releng)

=head SYNOPSIS

    % myproject-release
        --help
        --task=<task>            # 'create-pullrequest' or 'github-release'
        --next-version=<version> # specify next version. if not specified, a
                                 # value derived from the branch name is used.
        --package-name=<name>    # target package name. if not specified,
                                 # the repository name is used.
        --verbose
        --dry-run

=head DEPENDENCY

`git` command and `hub` (or `gh`) command and `github-release` command are required.

=cut

use 5.014;
use strict;
use warnings;
use utf8;
use Carp;

use IPC::Cmd qw/run/;
use JSON::PP;
use ExtUtils::MakeMaker qw/prompt/;
use Time::Piece;
use POSIX qw(setlocale LC_TIME);
use version;
use Getopt::Long;
use File::Basename;

chomp(my $REPO_NAME = `git rev-parse --show-toplevel`);
$REPO_NAME =~ s|^.+/([^/]+/[^/]+)$|$1|;
$REPO_NAME =~ m!/([^/]+)$!;
my $REPO = $1;
$REPO_NAME =~ m!^([^/]+)!;
my $REPO_USER = $1;
my $dry_run;
my $verbose;
my $next_version;
my $package_name = $REPO;

sub DEBUG() { $ENV{RELEASE_DEBUG} || $verbose }

sub command {
    say('+ '. join ' ', @_) if DEBUG;
    my @args = (@_ , ("> /dev/null","2>&1"));
    !system("@args") or croak $!;
}
sub command_with_exit_code {
    say('+ '. join ' ', @_) if DEBUG;
    my @args = (@_ , ("> /dev/null","2>&1"));
    system("@args");
}
sub _git {
    if($dry_run){
        return
    }
    state $com = do {
        chomp(my $c = `which git`);
        die "git command is required\n" unless $c;
        $c;
    };
}
sub git {
    if($dry_run){
        return
    }
    unshift  @_, _git; goto \&command
}

sub git_with_exit_code {
    if($dry_run){
        return
    }
    unshift  @_, _git; goto \&command_with_exit_code
}


sub _hub {
    if($dry_run){
        return
    }
    state $com = do {
        chomp(my $c = `which hub`);
        unless ($c) {
            chomp($c = `which gh`);
        }
        die "hub or gh command is required\n" unless $c;
        $c;
    };
}
sub hub {
    if($dry_run){
      return
    }
    unshift @_, _hub; goto \&command;
}

sub hub_with_exit_code {
    if($dry_run){
      return
    }
    unshift @_, _hub; goto \&command_with_exit_code;
}

sub _github_release {
    if($dry_run){
        return
    }
    state $com = do {
        chomp(my $c = `which github-release`);
        die "github-release command is required\n" unless $c;
        $c;
    };
}
sub github_release {
    if($dry_run){
        return
    }
    unshift  @_, _github_release; goto \&command
}

sub github_release_with_exit_code {
    if($dry_run){
        return
    }
    unshift  @_, _github_release; goto \&command_with_exit_code
}

sub http_get {
    my $url = shift;
    my ($ok, $err, undef, $stdout, $stderr) = run(command => [qw{curl -sf}, $url]);
    return {
        success => $ok,
        content => join('', @$stdout),
    };
}

# logger. steal from minilla
use Term::ANSIColor qw(colored);
use constant { LOG_DEBUG => 1, LOG_INFO  => 2, LOG_WARN  => 3, LOG_ERROR => 4 };

my $Colors = {
    LOG_DEBUG,   => 'green',
    LOG_WARN,    => 'yellow',
    LOG_INFO,    => 'cyan',
    LOG_ERROR,   => 'red',
};

sub _printf {
    my $type = pop;
    return if $type == LOG_DEBUG && !DEBUG;
    my ($temp, @args) = @_;
    my $msg = sprintf($temp, map { defined($_) ? $_ : '-' } @args);
    $msg = colored $msg, $Colors->{$type} if defined $type;
    my $fh = $type && $type >= LOG_WARN ? *STDERR : *STDOUT;
    print $fh $msg;
}

sub infof  {_printf(@_, LOG_INFO)}
sub warnf  {_printf(@_, LOG_WARN)}
sub debugf {_printf(@_, LOG_DEBUG)}
sub errorf {
    my(@msg) = @_;
    _printf(@msg, LOG_ERROR);

    my $fmt = shift @msg;
    die sprintf($fmt, @msg);
}

# file utils
sub slurp {
    my $file = shift;
    local $/;
    open my $fh, '<:encoding(UTF-8)', $file or die $!;
    <$fh>
}
sub spew {
    my ($file, $data) = @_;
    open my $fh, '>:encoding(UTF-8)', $file or die $!;
    $data .= "\n" if $data !~ /\n\z/ms;
    print $fh $data;
}
sub replace {
    my ($file, $code) = @_;
    if(!-f $file){ errorf "File does not exist: %s\n", $file; }
    my $content = $code->(slurp($file), $file);
    spew($file, $content);
}

# scope_guard
package __g {
    sub new {
        my ($class, $code) = @_;
        bless $code, $class;
    }
    sub DESTROY {
        my $self = shift;
        $self->();
    }
}
sub scope_guard(&) {
    my $code = shift;
    __g->new($code);
}

###
sub last_release {
    my @out = `git tag`;

    my ($tag) =
        sort { version->parse($b) <=> version->parse($a) }
        map {if(/^\d.\d$/){ $_ . ".0" } else { $_ } }
        map {if(/^v([0-9]+(?:\.[0-9]+)+)$/){ $1 } else { () }}
        map {chomp; $_} @out;
    $tag;
}

sub current_branch {
  chomp(my $branch =  `git symbolic-ref --short HEAD`);
  return $branch;
}

sub next_version {
    return $next_version if $next_version;
    if(current_branch() =~ /^bump-version-([\d\.]+)$/){
        return $next_version = $1;
    } elsif(current_branch() =~ /^master$/){
        my $next = `git log --merges --oneline`;
        if($next =~ /^.+Merge pull request #[0-9]+ from .+\/bump-version-([0-9\.]+)$/m){
            return $next_version = $1;
        }
    }
    return "";
}

sub diff {
  git_with_exit_code qw/diff --exit-code/
}

sub merged_prs {
    my $current_tag = shift;
    my @pull_nums = sort {$a <=> $b} map {
        if(m/Merge pull request #([0-9]+) /){ $1 } else { () }
    } `git log v$current_tag... --merges --oneline`;
    say "Pull Requests: " . join(", ", @pull_nums);
    my @releases;
    for my $pull_num (@pull_nums) {
        my $url = sprintf "https://api.github.com/repos/%s/pulls/%d", $REPO_NAME, $pull_num;
        my $res = http_get($url);
        unless ($res->{success}) {
            warnf "request to $url failed\n";
            next; #exit;
        }
        my $data = eval { decode_json $res->{content} };
        if ($@) {
            warnf "parse json failed. url: $url\n";
            next;
        }

        push @releases, {
            num   => $pull_num,
            title => $data->{title},
            user  => $data->{user}{login},
            url   => $data->{html_url},
        } if $data->{title} !~ /\[nit\]/i;
    }
    @releases;
}

sub find_latest_release_pullrequest {
    my $url = sprintf "https://api.github.com/repos/%s/pulls?state=close&sort=updated&direction=desc", $REPO_NAME;
    my $res = http_get($url);
    unless ($res->{success}) {
        warnf "request to $url failed\n";
        return; #exit;
    }
    my $data = eval { decode_json $res->{content} };
    if ($@) {
        warnf "parse json failed. url: $url\n";
        return;
    }
    for my $pull (@$data){
        if($pull->{head}->{ref} =~ /bump-version-[0-9\.]+/){
            return $pull;
        }
    }
}

sub parse_version {
    my $ver = shift;
    my ($major, $minor, $patch) = $ver =~ /^([0-9]+)\.([0-9]+)\.([0-9]+)$/;
    ($major, $minor, $patch)
}

sub update_changelog {
    my ($next_version, @releases) = @_;

    my $email = 'sacloud.users@gmail.com';
    my $name = 'sacloud-bot';

    my $old_locale = setlocale(LC_TIME);
    setlocale(LC_TIME, "C");
    my $g = scope_guard {
        setlocale(LC_TIME, $old_locale);
    };

    my $now = localtime;
    my $ret = 0;

    replace 'package/deb/debian/changelog' => sub {
        my ($content, $filename) = @_;

        my $update = "$package_name ($next_version-1) stable; urgency=low\n\n";
        for my $rel (@releases) {
            $update .= sprintf "  * %s (by %s)\n    <%s>\n", $rel->{title}, $rel->{user}, $rel->{url};
        }
        $update .= sprintf "\n -- %s <%s>  %s\n\n", $name, $email, $now->strftime("%a, %d %b %Y %H:%M:%S %z");

        # '35' means a rough length to cut timestamp at the end of update.
        if(index($content, substr($update, 0, length($update) - 35)) == -1){
            infof "update '$filename'.\n";
            $ret = 1;
            $update . $content;
        } else {
            infof "skip to update '$filename'.\n";
            $content;
        }
    };

    my $rpm_changelog_replacer = sub {
        my ($content, $filename) = @_;

        my $update = sprintf "* %s <%s> - %s-1\n", $now->strftime('%a %b %d %Y'), $email, $next_version;
        for my $rel (@releases) {
            $update .= sprintf "- %s (by %s)\n", $rel->{title}, $rel->{user};
        }
        if (index($content, $update) == -1) {
            infof "update '%s'\n", $filename;
            $ret = 1;
            $content =~ s/%changelog/%changelog\n$update/;
        } else {
            infof "skip to update '%s'\n", $filename;
        }
        $content;
    };

    my $rpm_spec = sprintf "package/rpm/%s.spec", $package_name;
    replace $rpm_spec => $rpm_changelog_replacer;

    replace 'CHANGELOG.md' => sub {
        my ($content, $filename) = @_;

        my $update = sprintf "\n\n## %s (%s)\n\n", $next_version, $now->strftime('%Y-%m-%d');
        for my $rel (@releases) {
            $update .= sprintf "* %s #%d (%s)\n", $rel->{title}, $rel->{num}, $rel->{user};
        }
        if(index($content, $update) == -1){
            infof "update '$filename'\n";
            $ret = 1;
            $content =~ s/\A# Changelog/# Changelog$update/;
        } else {
            infof "skip to update '$filename'\n";
        }
        $content;
    };
    $ret;
}

sub build_pull_request_body {
    my ($next_version, @releases) = @_;
    my $body = "Release version $next_version\n\n";
    for my $rel (@releases) {
        $body .= sprintf "- %s #%s\n", $rel->{title}, $rel->{num};
    }
    $body;
}

sub upload_to_github_release {
    my $tag = shift || "v" . next_version();
    my $name = shift || $REPO ."-" . next_version();
    my $erase_if_exist = shift || 0;
    my $is_staging = $name eq "staging";
    my $cont = 0;
    do {
        $cont = 0;
        my $ret = github_release_with_exit_code qw/release --user/, $REPO_USER,
            qw/--repo/, $REPO,
            qw/--tag/, $tag,
            qw/--name/, $name,
            qw/--description/, "pre-release",
            qw/--pre-release/;

        if($ret != 0){
            if($erase_if_exist){
                $ret = github_release_with_exit_code qw/delete --user/, $REPO_USER,
                    qw/--repo/, $REPO,
                    qw/--tag/, $tag;
                if($ret != 0){
                    infof "The release of this version has already existed at GitHub Release, and cannot be erased.\n";
                    return;
                }
                $cont = 1;
            } else {
                infof "The release of this version has already existed at GitHub Release, so skip the process.\n";
                return;
            }
        }
    } while($cont);

    # upload files
    my $latest_release_pullrequest = find_latest_release_pullrequest();
    my $description = $latest_release_pullrequest->{body};
    if ($is_staging){
        infof "Delete staging tag\n";
        git qw/push/, "https://$ENV{GITHUB_TOKEN}\@github.com/$REPO_NAME.git" , qw/:staging/;
        $description = "This release includes latest commits with unreleased features";
    }
    my @filepaths = glob("bin/*.zip");
    infof "uploading following files:\n" . join("\n", @filepaths). "\n";
    for my $path (@filepaths){
        my $filename = basename($path);
        github_release qw/upload --user/, $REPO_USER,
            qw/--repo/, $REPO,
            qw/--tag/, $tag,
            qw/--name/, $filename,
            qw/--file/, $path;
    }

    my @args = (qw/edit --user/, $REPO_USER,
        qw/--repo/, $REPO,
        qw/--tag/, $tag,
        qw/--description/, "'$description'");
    push @args, "--pre-release" if $erase_if_exist;
    github_release @args;

}

sub create_pull_request {
    my($current_branch) = @_;
    if(!$current_branch){
        $current_branch = current_branch()
    }

    git qw/checkout/, $current_branch;

    my $current_version = last_release;
    my $next_version    = next_version;
    if(!$next_version) {
        warnf("current branch should be 'bump-version-x.y.z'.\n");
        return;
    }

    # return if last commit is updating specs and changelogs
    my $last_commitlog = `git log --oneline master..HEAD | cut -d ' ' -f2-`;
    if($last_commitlog =~ /update changelogs/){
        infof "skip to update changelogs because the last commit is 'update changelogs'.\n";
        return;
    }

    infof "Update changelogs, between %s and %s.\n", $current_version, $next_version;
    my @releases = merged_prs $current_version;
    my $ret = 0;
    if(@releases){
        $ret = update_changelog($next_version, @releases);
    } else {
        infof "skip to update changelogs because no merged pull request is found after the last release.\n"
    }

    infof "Build GitHub Pages.\n";
    system "make", "build-docs";

    infof "Update AUTHORS.\n";
    system "scripts/generate-authors.sh";

    if($ret || git_with_exit_code qw/diff --exit-code/){
        git qw/config --global push.default matching/;
        git qw/config user.email/, 'sacloud.users@gmail.com';
        git qw/config user.name/, 'sacloud-bot';
        git qw/commit -am/, "'update changelogs'";
        git qw/push -u/, "https://$ENV{GITHUB_TOKEN}\@github.com/$REPO_NAME.git";
        infof "Difference is pushed to GitHub.\n";
    }

    #git qw/diff/, qw/--word-diff/, "master..".current_branch();
    my $pr_body = build_pull_request_body($next_version, @releases);
    say "Pull Request\n--------";
    say $pr_body;

    infof "Push changes\n";
    git qw/pull/;
    $ret = hub_with_exit_code qw/pull-request -m/, "'$pr_body'", qq/-b/, qq/$REPO_NAME:master/;
    if(!defined($ret) || $ret == 0){
        infof "Releasing pull request is created. Review and merge it. You can update changelogs and commit more in this branch before merging.\n";
    } else {
        # do push and pull-request again
        git qw/push -u/, "https://$ENV{GITHUB_TOKEN}\@github.com/$REPO_NAME.git";
        $ret = hub_with_exit_code qw/pull-request -m/, "'$pr_body'", qq/-b/, qq/$REPO_NAME:master/;
        if(!defined($ret) || $ret == 0){
            infof "Releasing pull request is created. Review and merge it. You can update changelogs and commit more in this branch before merging.\n";
        } else {
            infof "Pull Request already seems to exist.\n"
        }
    }
}

### main process
if (!$ENV{HARNESS_ACTIVE}) {
    main();
} else {
    # When called via `prove`, tests will run.
    run_tests();
}

sub help {
  print <<EOS;
$0
  --help
  --task=<task>            # 'create-pullrequest'
                           # or 'upload-to-github-release'
                           # or 'upload-master-to-github-release'
  --next-version=<version> # specify next version. if not specified, a
                           # value derived from the branch name is used.
  --package-name=<name>    # target package name. if not specified,
                           # the repository name is used.
  --verbose
  --dry-run
EOS

}

sub main {
    my $help;
    my $task;
    my $current_branch;
    GetOptions ("help" => \$help,
        "task=s"   => \$task,  # create-pullrequest, upload-to-github-release
        "verbose"  => \$verbose,
        "dry-run" => \$dry_run,
        "current-branch=s" => \$current_branch,
        "next-version=s" => \$next_version,
        "package-name=s" => \$package_name,
        );

    if($help){
        help();
        exit 0;
    }
    # check command
    _git;_hub;

    if($task && $task eq "create-pullrequest"){
        infof "Create a Pull Request for version up.\n";
        create_pull_request($current_branch);
        return;
    } elsif($task && $task eq "upload-to-github-release"){
        infof "Upload files to GitHub Release.\n";
        upload_to_github_release();
        return;
    } elsif($task && $task eq "update-changelog"){
        infof "Updating changelog.\n";
        update_changelog(next_version);
        return;
    } elsif($task && $task eq "upload-master-to-github-release"){
        infof "Upload files on master branch to GitHub Release.\n";
        upload_to_github_release("staging", "staging", 1);
        return;
    }
    help();
}

sub run_tests {
    require Test::More;
    Test::More->import;

    my $version = '0.1.2';
    my ($major, $minor, $patch) = parse_version($version);
    is($major, 0);
    is($minor, 1);
    is($patch, 2);

    done_testing();
}