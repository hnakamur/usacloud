// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListAutoBackupParam is input parameters for the sacloud API
type ListAutoBackupParam struct {
	Name []string
	Id   []int64
	From int
	Max  int
	Sort []string
}

// NewListAutoBackupParam return new ListAutoBackupParam
func NewListAutoBackupParam() *ListAutoBackupParam {
	return &ListAutoBackupParam{}
}

// Validate checks current values in model
func (p *ListAutoBackupParam) Validate() []error {
	errors := []error{}
	{
		errs := validateConflicts("--name", p.Name, map[string]interface{}{

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
		errs := validateConflicts("--id", p.Id, map[string]interface{}{

			"--name": p.Name,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ListAutoBackupParam) getResourceDef() *schema.Resource {
	return define.Resources["AutoBackup"]
}

func (p *ListAutoBackupParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListAutoBackupParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListAutoBackupParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListAutoBackupParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListAutoBackupParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListAutoBackupParam) SetName(v []string) {
	p.Name = v
}

func (p *ListAutoBackupParam) GetName() []string {
	return p.Name
}
func (p *ListAutoBackupParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListAutoBackupParam) GetId() []int64 {
	return p.Id
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

// CreateAutoBackupParam is input parameters for the sacloud API
type CreateAutoBackupParam struct {
	Name        string
	Description string
	IconId      int64
	Tags        []string
	Generation  int
	StartHour   int
	Weekdays    []string
	DiskId      int64
}

// NewCreateAutoBackupParam return new CreateAutoBackupParam
func NewCreateAutoBackupParam() *CreateAutoBackupParam {
	return &CreateAutoBackupParam{

		Generation: 1,

		Weekdays: []string{"all"},
	}
}

// Validate checks current values in model
func (p *CreateAutoBackupParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
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
		validator := define.Resources["AutoBackup"].Commands["create"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
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
		validator := validateRequired
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
		validator := validateRequired
		errs := validator("--start-hour", p.StartHour)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["create"].Params["start-hour"].ValidateFunc
		errs := validator("--start-hour", p.StartHour)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
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
		validator := validateRequired
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

	return errors
}

func (p *CreateAutoBackupParam) getResourceDef() *schema.Resource {
	return define.Resources["AutoBackup"]
}

func (p *CreateAutoBackupParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["create"]
}

func (p *CreateAutoBackupParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *CreateAutoBackupParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *CreateAutoBackupParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *CreateAutoBackupParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
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
func (p *CreateAutoBackupParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *CreateAutoBackupParam) GetIconId() int64 {
	return p.IconId
}
func (p *CreateAutoBackupParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateAutoBackupParam) GetTags() []string {
	return p.Tags
}
func (p *CreateAutoBackupParam) SetGeneration(v int) {
	p.Generation = v
}

func (p *CreateAutoBackupParam) GetGeneration() int {
	return p.Generation
}
func (p *CreateAutoBackupParam) SetStartHour(v int) {
	p.StartHour = v
}

func (p *CreateAutoBackupParam) GetStartHour() int {
	return p.StartHour
}
func (p *CreateAutoBackupParam) SetWeekdays(v []string) {
	p.Weekdays = v
}

func (p *CreateAutoBackupParam) GetWeekdays() []string {
	return p.Weekdays
}
func (p *CreateAutoBackupParam) SetDiskId(v int64) {
	p.DiskId = v
}

func (p *CreateAutoBackupParam) GetDiskId() int64 {
	return p.DiskId
}

// ReadAutoBackupParam is input parameters for the sacloud API
type ReadAutoBackupParam struct {
	Id int64
}

// NewReadAutoBackupParam return new ReadAutoBackupParam
func NewReadAutoBackupParam() *ReadAutoBackupParam {
	return &ReadAutoBackupParam{}
}

// Validate checks current values in model
func (p *ReadAutoBackupParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadAutoBackupParam) getResourceDef() *schema.Resource {
	return define.Resources["AutoBackup"]
}

func (p *ReadAutoBackupParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadAutoBackupParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadAutoBackupParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadAutoBackupParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadAutoBackupParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadAutoBackupParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadAutoBackupParam) GetId() int64 {
	return p.Id
}

// UpdateAutoBackupParam is input parameters for the sacloud API
type UpdateAutoBackupParam struct {
	Tags        []string
	StartHour   int
	Weekdays    []string
	Id          int64
	Name        string
	Description string
	IconId      int64
	Generation  int
}

// NewUpdateAutoBackupParam return new UpdateAutoBackupParam
func NewUpdateAutoBackupParam() *UpdateAutoBackupParam {
	return &UpdateAutoBackupParam{}
}

// Validate checks current values in model
func (p *UpdateAutoBackupParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["AutoBackup"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["update"].Params["start-hour"].ValidateFunc
		errs := validator("--start-hour", p.StartHour)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["update"].Params["weekdays"].ValidateFunc
		errs := validator("--weekdays", p.Weekdays)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
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
		validator := define.Resources["AutoBackup"].Commands["update"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
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

	return errors
}

func (p *UpdateAutoBackupParam) getResourceDef() *schema.Resource {
	return define.Resources["AutoBackup"]
}

func (p *UpdateAutoBackupParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["update"]
}

func (p *UpdateAutoBackupParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UpdateAutoBackupParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UpdateAutoBackupParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UpdateAutoBackupParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UpdateAutoBackupParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateAutoBackupParam) GetTags() []string {
	return p.Tags
}
func (p *UpdateAutoBackupParam) SetStartHour(v int) {
	p.StartHour = v
}

func (p *UpdateAutoBackupParam) GetStartHour() int {
	return p.StartHour
}
func (p *UpdateAutoBackupParam) SetWeekdays(v []string) {
	p.Weekdays = v
}

func (p *UpdateAutoBackupParam) GetWeekdays() []string {
	return p.Weekdays
}
func (p *UpdateAutoBackupParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateAutoBackupParam) GetId() int64 {
	return p.Id
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
func (p *UpdateAutoBackupParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *UpdateAutoBackupParam) GetIconId() int64 {
	return p.IconId
}
func (p *UpdateAutoBackupParam) SetGeneration(v int) {
	p.Generation = v
}

func (p *UpdateAutoBackupParam) GetGeneration() int {
	return p.Generation
}

// DeleteAutoBackupParam is input parameters for the sacloud API
type DeleteAutoBackupParam struct {
	Id int64
}

// NewDeleteAutoBackupParam return new DeleteAutoBackupParam
func NewDeleteAutoBackupParam() *DeleteAutoBackupParam {
	return &DeleteAutoBackupParam{}
}

// Validate checks current values in model
func (p *DeleteAutoBackupParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteAutoBackupParam) getResourceDef() *schema.Resource {
	return define.Resources["AutoBackup"]
}

func (p *DeleteAutoBackupParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["delete"]
}

func (p *DeleteAutoBackupParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DeleteAutoBackupParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DeleteAutoBackupParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DeleteAutoBackupParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DeleteAutoBackupParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteAutoBackupParam) GetId() int64 {
	return p.Id
}
