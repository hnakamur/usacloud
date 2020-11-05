// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-params'; DO NOT EDIT

package params

import (
	"io"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/config"
	"github.com/sacloud/usacloud/pkg/define"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/sacloud/usacloud/pkg/validation"
)

// ListAutoBackupParam is input parameters for the sacloud API
type ListAutoBackupParam struct {
	Name             []string
	Id               []types.ID
	Tags             []string
	From             int
	Max              int
	Sort             []string
	Parameters       string
	ParameterFile    string
	GenerateSkeleton bool
	OutputType       string
	Column           []string
	Quiet            bool
	Format           string
	FormatFile       string
	Query            string
	QueryFile        string

	config *config.Config
	input  Input
}

// NewListAutoBackupParam return new ListAutoBackupParam
func NewListAutoBackupParam() *ListAutoBackupParam {
	return &ListAutoBackupParam{}
}

// Initialize init ListAutoBackupParam
func (p *ListAutoBackupParam) Initialize(in Input, args []string, config *config.Config) error {
	p.input = in
	p.config = config
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListAutoBackupParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *ListAutoBackupParam) FillValueToSkeleton() {
	if util.IsEmpty(p.Name) {
		p.Name = []string{""}
	}
	if util.IsEmpty(p.Id) {
		p.Id = []types.ID{}
	}
	if util.IsEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if util.IsEmpty(p.From) {
		p.From = 0
	}
	if util.IsEmpty(p.Max) {
		p.Max = 0
	}
	if util.IsEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if util.IsEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if util.IsEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if util.IsEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if util.IsEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if util.IsEmpty(p.Column) {
		p.Column = []string{""}
	}
	if util.IsEmpty(p.Quiet) {
		p.Quiet = false
	}
	if util.IsEmpty(p.Format) {
		p.Format = ""
	}
	if util.IsEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if util.IsEmpty(p.Query) {
		p.Query = ""
	}
	if util.IsEmpty(p.QueryFile) {
		p.QueryFile = ""
	}

}

func (p *ListAutoBackupParam) validate() error {
	var errors []error

	{
		errs := validation.ConflictsWith("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["AutoBackup"].Commands["list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validation.ConflictsWith("--id", p.Id, map[string]interface{}{

			"--name": p.Name,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["AutoBackup"].Commands["list"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validateParameterOptions(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(config.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := cli.ValidateOutputOption(p, p.config.DefaultOutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return util.FlattenErrors(errors)
}

func (p *ListAutoBackupParam) ResourceDef() *schema.Resource {
	return define.Resources["AutoBackup"]
}

func (p *ListAutoBackupParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListAutoBackupParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListAutoBackupParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListAutoBackupParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListAutoBackupParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListAutoBackupParam) SetName(v []string) {
	p.Name = v
}

func (p *ListAutoBackupParam) GetName() []string {
	return p.Name
}
func (p *ListAutoBackupParam) SetId(v []types.ID) {
	p.Id = v
}

func (p *ListAutoBackupParam) GetId() []types.ID {
	return p.Id
}
func (p *ListAutoBackupParam) SetTags(v []string) {
	p.Tags = v
}

func (p *ListAutoBackupParam) GetTags() []string {
	return p.Tags
}
func (p *ListAutoBackupParam) SetFrom(v int) {
	p.From = v
}

func (p *ListAutoBackupParam) GetFrom() int {
	return p.From
}
func (p *ListAutoBackupParam) SetMax(v int) {
	p.Max = v
}

func (p *ListAutoBackupParam) GetMax() int {
	return p.Max
}
func (p *ListAutoBackupParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListAutoBackupParam) GetSort() []string {
	return p.Sort
}
func (p *ListAutoBackupParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ListAutoBackupParam) GetParameters() string {
	return p.Parameters
}
func (p *ListAutoBackupParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ListAutoBackupParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ListAutoBackupParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListAutoBackupParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListAutoBackupParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListAutoBackupParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListAutoBackupParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListAutoBackupParam) GetColumn() []string {
	return p.Column
}
func (p *ListAutoBackupParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListAutoBackupParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListAutoBackupParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListAutoBackupParam) GetFormat() string {
	return p.Format
}
func (p *ListAutoBackupParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListAutoBackupParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListAutoBackupParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListAutoBackupParam) GetQuery() string {
	return p.Query
}
func (p *ListAutoBackupParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListAutoBackupParam) GetQueryFile() string {
	return p.QueryFile
}

// Changed 指定の項目に入力があった場合にtrueを返す
func (p *ListAutoBackupParam) Changed(name string) bool {
	return p.input.Changed(name)
}

// CreateAutoBackupParam is input parameters for the sacloud API
type CreateAutoBackupParam struct {
	DiskId           types.ID
	Weekdays         []string
	Generation       int
	Name             string
	Description      string
	Tags             []string
	IconId           types.ID
	Assumeyes        bool
	Parameters       string
	ParameterFile    string
	GenerateSkeleton bool
	OutputType       string
	Column           []string
	Quiet            bool
	Format           string
	FormatFile       string
	Query            string
	QueryFile        string

	config *config.Config
	input  Input
}

// NewCreateAutoBackupParam return new CreateAutoBackupParam
func NewCreateAutoBackupParam() *CreateAutoBackupParam {
	return &CreateAutoBackupParam{
		Weekdays:   []string{"all"},
		Generation: 1,
	}
}

// Initialize init CreateAutoBackupParam
func (p *CreateAutoBackupParam) Initialize(in Input, args []string, config *config.Config) error {
	p.input = in
	p.config = config
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *CreateAutoBackupParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *CreateAutoBackupParam) FillValueToSkeleton() {
	if util.IsEmpty(p.DiskId) {
		p.DiskId = types.ID(0)
	}
	if util.IsEmpty(p.Weekdays) {
		p.Weekdays = []string{""}
	}
	if util.IsEmpty(p.Generation) {
		p.Generation = 0
	}
	if util.IsEmpty(p.Name) {
		p.Name = ""
	}
	if util.IsEmpty(p.Description) {
		p.Description = ""
	}
	if util.IsEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if util.IsEmpty(p.IconId) {
		p.IconId = types.ID(0)
	}
	if util.IsEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if util.IsEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if util.IsEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if util.IsEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if util.IsEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if util.IsEmpty(p.Column) {
		p.Column = []string{""}
	}
	if util.IsEmpty(p.Quiet) {
		p.Quiet = false
	}
	if util.IsEmpty(p.Format) {
		p.Format = ""
	}
	if util.IsEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if util.IsEmpty(p.Query) {
		p.Query = ""
	}
	if util.IsEmpty(p.QueryFile) {
		p.QueryFile = ""
	}

}

func (p *CreateAutoBackupParam) validate() error {
	var errors []error

	{
		validator := cli.ValidateRequired
		errs := validator("--disk-id", p.DiskId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["create"].Params["disk-id"].ValidateFunc
		errs := validator("--disk-id", p.DiskId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := cli.ValidateRequired
		errs := validator("--weekdays", p.Weekdays)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["create"].Params["weekdays"].ValidateFunc
		errs := validator("--weekdays", p.Weekdays)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := cli.ValidateRequired
		errs := validator("--generation", p.Generation)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["create"].Params["generation"].ValidateFunc
		errs := validator("--generation", p.Generation)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := cli.ValidateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["AutoBackup"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["AutoBackup"].Commands["create"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["AutoBackup"].Commands["create"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validateParameterOptions(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(config.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := cli.ValidateOutputOption(p, p.config.DefaultOutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return util.FlattenErrors(errors)
}

func (p *CreateAutoBackupParam) ResourceDef() *schema.Resource {
	return define.Resources["AutoBackup"]
}

func (p *CreateAutoBackupParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["create"]
}

func (p *CreateAutoBackupParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CreateAutoBackupParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CreateAutoBackupParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CreateAutoBackupParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *CreateAutoBackupParam) SetDiskId(v types.ID) {
	p.DiskId = v
}

func (p *CreateAutoBackupParam) GetDiskId() types.ID {
	return p.DiskId
}
func (p *CreateAutoBackupParam) SetWeekdays(v []string) {
	p.Weekdays = v
}

func (p *CreateAutoBackupParam) GetWeekdays() []string {
	return p.Weekdays
}
func (p *CreateAutoBackupParam) SetGeneration(v int) {
	p.Generation = v
}

func (p *CreateAutoBackupParam) GetGeneration() int {
	return p.Generation
}
func (p *CreateAutoBackupParam) SetName(v string) {
	p.Name = v
}

func (p *CreateAutoBackupParam) GetName() string {
	return p.Name
}
func (p *CreateAutoBackupParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateAutoBackupParam) GetDescription() string {
	return p.Description
}
func (p *CreateAutoBackupParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateAutoBackupParam) GetTags() []string {
	return p.Tags
}
func (p *CreateAutoBackupParam) SetIconId(v types.ID) {
	p.IconId = v
}

func (p *CreateAutoBackupParam) GetIconId() types.ID {
	return p.IconId
}
func (p *CreateAutoBackupParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *CreateAutoBackupParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *CreateAutoBackupParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *CreateAutoBackupParam) GetParameters() string {
	return p.Parameters
}
func (p *CreateAutoBackupParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *CreateAutoBackupParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *CreateAutoBackupParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *CreateAutoBackupParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *CreateAutoBackupParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CreateAutoBackupParam) GetOutputType() string {
	return p.OutputType
}
func (p *CreateAutoBackupParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CreateAutoBackupParam) GetColumn() []string {
	return p.Column
}
func (p *CreateAutoBackupParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CreateAutoBackupParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CreateAutoBackupParam) SetFormat(v string) {
	p.Format = v
}

func (p *CreateAutoBackupParam) GetFormat() string {
	return p.Format
}
func (p *CreateAutoBackupParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *CreateAutoBackupParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *CreateAutoBackupParam) SetQuery(v string) {
	p.Query = v
}

func (p *CreateAutoBackupParam) GetQuery() string {
	return p.Query
}
func (p *CreateAutoBackupParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *CreateAutoBackupParam) GetQueryFile() string {
	return p.QueryFile
}

// Changed 指定の項目に入力があった場合にtrueを返す
func (p *CreateAutoBackupParam) Changed(name string) bool {
	return p.input.Changed(name)
}

// ReadAutoBackupParam is input parameters for the sacloud API
type ReadAutoBackupParam struct {
	Selector         []string
	Parameters       string
	ParameterFile    string
	GenerateSkeleton bool
	OutputType       string
	Column           []string
	Quiet            bool
	Format           string
	FormatFile       string
	Query            string
	QueryFile        string
	Id               types.ID

	config *config.Config
	input  Input
}

// NewReadAutoBackupParam return new ReadAutoBackupParam
func NewReadAutoBackupParam() *ReadAutoBackupParam {
	return &ReadAutoBackupParam{}
}

// WithID returns new *ReadAutoBackupParam with id
func (p *ReadAutoBackupParam) WithID(id types.ID) *ReadAutoBackupParam {
	return &ReadAutoBackupParam{
		Selector:         p.Selector,
		Parameters:       p.Parameters,
		ParameterFile:    p.ParameterFile,
		GenerateSkeleton: p.GenerateSkeleton,
		OutputType:       p.OutputType,
		Column:           p.Column,
		Quiet:            p.Quiet,
		Format:           p.Format,
		FormatFile:       p.FormatFile,
		Query:            p.Query,
		QueryFile:        p.QueryFile,
		Id:               id,
	}
}

// Initialize init ReadAutoBackupParam
func (p *ReadAutoBackupParam) Initialize(in Input, args []string, config *config.Config) error {
	p.input = in
	p.config = config
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ReadAutoBackupParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *ReadAutoBackupParam) FillValueToSkeleton() {
	if util.IsEmpty(p.Selector) {
		p.Selector = []string{""}
	}
	if util.IsEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if util.IsEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if util.IsEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if util.IsEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if util.IsEmpty(p.Column) {
		p.Column = []string{""}
	}
	if util.IsEmpty(p.Quiet) {
		p.Quiet = false
	}
	if util.IsEmpty(p.Format) {
		p.Format = ""
	}
	if util.IsEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if util.IsEmpty(p.Query) {
		p.Query = ""
	}
	if util.IsEmpty(p.QueryFile) {
		p.QueryFile = ""
	}
	if util.IsEmpty(p.Id) {
		p.Id = types.ID(0)
	}

}

func (p *ReadAutoBackupParam) validate() error {
	var errors []error

	{
		validator := cli.ValidateSakuraID
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validateParameterOptions(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(config.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := cli.ValidateOutputOption(p, p.config.DefaultOutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return util.FlattenErrors(errors)
}

func (p *ReadAutoBackupParam) ResourceDef() *schema.Resource {
	return define.Resources["AutoBackup"]
}

func (p *ReadAutoBackupParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadAutoBackupParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadAutoBackupParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadAutoBackupParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadAutoBackupParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ReadAutoBackupParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *ReadAutoBackupParam) GetSelector() []string {
	return p.Selector
}
func (p *ReadAutoBackupParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ReadAutoBackupParam) GetParameters() string {
	return p.Parameters
}
func (p *ReadAutoBackupParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ReadAutoBackupParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ReadAutoBackupParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadAutoBackupParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadAutoBackupParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadAutoBackupParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadAutoBackupParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadAutoBackupParam) GetColumn() []string {
	return p.Column
}
func (p *ReadAutoBackupParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadAutoBackupParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadAutoBackupParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadAutoBackupParam) GetFormat() string {
	return p.Format
}
func (p *ReadAutoBackupParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadAutoBackupParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadAutoBackupParam) SetQuery(v string) {
	p.Query = v
}

func (p *ReadAutoBackupParam) GetQuery() string {
	return p.Query
}
func (p *ReadAutoBackupParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ReadAutoBackupParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *ReadAutoBackupParam) SetId(v types.ID) {
	p.Id = v
}

func (p *ReadAutoBackupParam) GetId() types.ID {
	return p.Id
}

// Changed 指定の項目に入力があった場合にtrueを返す
func (p *ReadAutoBackupParam) Changed(name string) bool {
	return p.input.Changed(name)
}

// UpdateAutoBackupParam is input parameters for the sacloud API
type UpdateAutoBackupParam struct {
	Weekdays         []string
	Generation       int
	Selector         []string
	Name             string
	Description      string
	Tags             []string
	IconId           types.ID
	Assumeyes        bool
	Parameters       string
	ParameterFile    string
	GenerateSkeleton bool
	OutputType       string
	Column           []string
	Quiet            bool
	Format           string
	FormatFile       string
	Query            string
	QueryFile        string
	Id               types.ID

	config *config.Config
	input  Input
}

// NewUpdateAutoBackupParam return new UpdateAutoBackupParam
func NewUpdateAutoBackupParam() *UpdateAutoBackupParam {
	return &UpdateAutoBackupParam{}
}

// WithID returns new *UpdateAutoBackupParam with id
func (p *UpdateAutoBackupParam) WithID(id types.ID) *UpdateAutoBackupParam {
	return &UpdateAutoBackupParam{
		Weekdays:         p.Weekdays,
		Generation:       p.Generation,
		Selector:         p.Selector,
		Name:             p.Name,
		Description:      p.Description,
		Tags:             p.Tags,
		IconId:           p.IconId,
		Assumeyes:        p.Assumeyes,
		Parameters:       p.Parameters,
		ParameterFile:    p.ParameterFile,
		GenerateSkeleton: p.GenerateSkeleton,
		OutputType:       p.OutputType,
		Column:           p.Column,
		Quiet:            p.Quiet,
		Format:           p.Format,
		FormatFile:       p.FormatFile,
		Query:            p.Query,
		QueryFile:        p.QueryFile,
		Id:               id,
	}
}

// Initialize init UpdateAutoBackupParam
func (p *UpdateAutoBackupParam) Initialize(in Input, args []string, config *config.Config) error {
	p.input = in
	p.config = config
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *UpdateAutoBackupParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *UpdateAutoBackupParam) FillValueToSkeleton() {
	if util.IsEmpty(p.Weekdays) {
		p.Weekdays = []string{""}
	}
	if util.IsEmpty(p.Generation) {
		p.Generation = 0
	}
	if util.IsEmpty(p.Selector) {
		p.Selector = []string{""}
	}
	if util.IsEmpty(p.Name) {
		p.Name = ""
	}
	if util.IsEmpty(p.Description) {
		p.Description = ""
	}
	if util.IsEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if util.IsEmpty(p.IconId) {
		p.IconId = types.ID(0)
	}
	if util.IsEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if util.IsEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if util.IsEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if util.IsEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if util.IsEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if util.IsEmpty(p.Column) {
		p.Column = []string{""}
	}
	if util.IsEmpty(p.Quiet) {
		p.Quiet = false
	}
	if util.IsEmpty(p.Format) {
		p.Format = ""
	}
	if util.IsEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if util.IsEmpty(p.Query) {
		p.Query = ""
	}
	if util.IsEmpty(p.QueryFile) {
		p.QueryFile = ""
	}
	if util.IsEmpty(p.Id) {
		p.Id = types.ID(0)
	}

}

func (p *UpdateAutoBackupParam) validate() error {
	var errors []error

	{
		validator := define.Resources["AutoBackup"].Commands["update"].Params["weekdays"].ValidateFunc
		errs := validator("--weekdays", p.Weekdays)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["AutoBackup"].Commands["update"].Params["generation"].ValidateFunc
		errs := validator("--generation", p.Generation)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["AutoBackup"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["AutoBackup"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["AutoBackup"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["AutoBackup"].Commands["update"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := cli.ValidateSakuraID
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validateParameterOptions(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(config.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := cli.ValidateOutputOption(p, p.config.DefaultOutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return util.FlattenErrors(errors)
}

func (p *UpdateAutoBackupParam) ResourceDef() *schema.Resource {
	return define.Resources["AutoBackup"]
}

func (p *UpdateAutoBackupParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["update"]
}

func (p *UpdateAutoBackupParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *UpdateAutoBackupParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *UpdateAutoBackupParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *UpdateAutoBackupParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *UpdateAutoBackupParam) SetWeekdays(v []string) {
	p.Weekdays = v
}

func (p *UpdateAutoBackupParam) GetWeekdays() []string {
	return p.Weekdays
}
func (p *UpdateAutoBackupParam) SetGeneration(v int) {
	p.Generation = v
}

func (p *UpdateAutoBackupParam) GetGeneration() int {
	return p.Generation
}
func (p *UpdateAutoBackupParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *UpdateAutoBackupParam) GetSelector() []string {
	return p.Selector
}
func (p *UpdateAutoBackupParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateAutoBackupParam) GetName() string {
	return p.Name
}
func (p *UpdateAutoBackupParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateAutoBackupParam) GetDescription() string {
	return p.Description
}
func (p *UpdateAutoBackupParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateAutoBackupParam) GetTags() []string {
	return p.Tags
}
func (p *UpdateAutoBackupParam) SetIconId(v types.ID) {
	p.IconId = v
}

func (p *UpdateAutoBackupParam) GetIconId() types.ID {
	return p.IconId
}
func (p *UpdateAutoBackupParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *UpdateAutoBackupParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *UpdateAutoBackupParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *UpdateAutoBackupParam) GetParameters() string {
	return p.Parameters
}
func (p *UpdateAutoBackupParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *UpdateAutoBackupParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *UpdateAutoBackupParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *UpdateAutoBackupParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *UpdateAutoBackupParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *UpdateAutoBackupParam) GetOutputType() string {
	return p.OutputType
}
func (p *UpdateAutoBackupParam) SetColumn(v []string) {
	p.Column = v
}

func (p *UpdateAutoBackupParam) GetColumn() []string {
	return p.Column
}
func (p *UpdateAutoBackupParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *UpdateAutoBackupParam) GetQuiet() bool {
	return p.Quiet
}
func (p *UpdateAutoBackupParam) SetFormat(v string) {
	p.Format = v
}

func (p *UpdateAutoBackupParam) GetFormat() string {
	return p.Format
}
func (p *UpdateAutoBackupParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *UpdateAutoBackupParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *UpdateAutoBackupParam) SetQuery(v string) {
	p.Query = v
}

func (p *UpdateAutoBackupParam) GetQuery() string {
	return p.Query
}
func (p *UpdateAutoBackupParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *UpdateAutoBackupParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *UpdateAutoBackupParam) SetId(v types.ID) {
	p.Id = v
}

func (p *UpdateAutoBackupParam) GetId() types.ID {
	return p.Id
}

// Changed 指定の項目に入力があった場合にtrueを返す
func (p *UpdateAutoBackupParam) Changed(name string) bool {
	return p.input.Changed(name)
}

// DeleteAutoBackupParam is input parameters for the sacloud API
type DeleteAutoBackupParam struct {
	Selector         []string
	Assumeyes        bool
	Parameters       string
	ParameterFile    string
	GenerateSkeleton bool
	OutputType       string
	Column           []string
	Quiet            bool
	Format           string
	FormatFile       string
	Query            string
	QueryFile        string
	Id               types.ID

	config *config.Config
	input  Input
}

// NewDeleteAutoBackupParam return new DeleteAutoBackupParam
func NewDeleteAutoBackupParam() *DeleteAutoBackupParam {
	return &DeleteAutoBackupParam{}
}

// WithID returns new *DeleteAutoBackupParam with id
func (p *DeleteAutoBackupParam) WithID(id types.ID) *DeleteAutoBackupParam {
	return &DeleteAutoBackupParam{
		Selector:         p.Selector,
		Assumeyes:        p.Assumeyes,
		Parameters:       p.Parameters,
		ParameterFile:    p.ParameterFile,
		GenerateSkeleton: p.GenerateSkeleton,
		OutputType:       p.OutputType,
		Column:           p.Column,
		Quiet:            p.Quiet,
		Format:           p.Format,
		FormatFile:       p.FormatFile,
		Query:            p.Query,
		QueryFile:        p.QueryFile,
		Id:               id,
	}
}

// Initialize init DeleteAutoBackupParam
func (p *DeleteAutoBackupParam) Initialize(in Input, args []string, config *config.Config) error {
	p.input = in
	p.config = config
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *DeleteAutoBackupParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *DeleteAutoBackupParam) FillValueToSkeleton() {
	if util.IsEmpty(p.Selector) {
		p.Selector = []string{""}
	}
	if util.IsEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if util.IsEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if util.IsEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if util.IsEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if util.IsEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if util.IsEmpty(p.Column) {
		p.Column = []string{""}
	}
	if util.IsEmpty(p.Quiet) {
		p.Quiet = false
	}
	if util.IsEmpty(p.Format) {
		p.Format = ""
	}
	if util.IsEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if util.IsEmpty(p.Query) {
		p.Query = ""
	}
	if util.IsEmpty(p.QueryFile) {
		p.QueryFile = ""
	}
	if util.IsEmpty(p.Id) {
		p.Id = types.ID(0)
	}

}

func (p *DeleteAutoBackupParam) validate() error {
	var errors []error

	{
		validator := cli.ValidateSakuraID
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validateParameterOptions(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(config.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := cli.ValidateOutputOption(p, p.config.DefaultOutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return util.FlattenErrors(errors)
}

func (p *DeleteAutoBackupParam) ResourceDef() *schema.Resource {
	return define.Resources["AutoBackup"]
}

func (p *DeleteAutoBackupParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["delete"]
}

func (p *DeleteAutoBackupParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *DeleteAutoBackupParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *DeleteAutoBackupParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *DeleteAutoBackupParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *DeleteAutoBackupParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *DeleteAutoBackupParam) GetSelector() []string {
	return p.Selector
}
func (p *DeleteAutoBackupParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeleteAutoBackupParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeleteAutoBackupParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *DeleteAutoBackupParam) GetParameters() string {
	return p.Parameters
}
func (p *DeleteAutoBackupParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *DeleteAutoBackupParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *DeleteAutoBackupParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *DeleteAutoBackupParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *DeleteAutoBackupParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *DeleteAutoBackupParam) GetOutputType() string {
	return p.OutputType
}
func (p *DeleteAutoBackupParam) SetColumn(v []string) {
	p.Column = v
}

func (p *DeleteAutoBackupParam) GetColumn() []string {
	return p.Column
}
func (p *DeleteAutoBackupParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *DeleteAutoBackupParam) GetQuiet() bool {
	return p.Quiet
}
func (p *DeleteAutoBackupParam) SetFormat(v string) {
	p.Format = v
}

func (p *DeleteAutoBackupParam) GetFormat() string {
	return p.Format
}
func (p *DeleteAutoBackupParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *DeleteAutoBackupParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *DeleteAutoBackupParam) SetQuery(v string) {
	p.Query = v
}

func (p *DeleteAutoBackupParam) GetQuery() string {
	return p.Query
}
func (p *DeleteAutoBackupParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *DeleteAutoBackupParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *DeleteAutoBackupParam) SetId(v types.ID) {
	p.Id = v
}

func (p *DeleteAutoBackupParam) GetId() types.ID {
	return p.Id
}

// Changed 指定の項目に入力があった場合にtrueを返す
func (p *DeleteAutoBackupParam) Changed(name string) bool {
	return p.input.Changed(name)
}