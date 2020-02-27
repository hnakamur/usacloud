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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListProductDiskParam is input parameters for the sacloud API
type ListProductDiskParam struct {
	Name              []string     `json:"name"`
	Id                []sacloud.ID `json:"id"`
	From              int          `json:"from"`
	Max               int          `json:"max"`
	Sort              []string     `json:"sort"`
	ParamTemplate     string       `json:"param-template"`
	ParamTemplateFile string       `json:"param-template-file"`
	GenerateSkeleton  bool         `json:"generate-skeleton"`
	OutputType        string       `json:"output-type"`
	Column            []string     `json:"column"`
	Quiet             bool         `json:"quiet"`
	Format            string       `json:"format"`
	FormatFile        string       `json:"format-file"`
	Query             string       `json:"query"`
	QueryFile         string       `json:"query-file"`
}

// NewListProductDiskParam return new ListProductDiskParam
func NewListProductDiskParam() *ListProductDiskParam {
	return &ListProductDiskParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ListProductDiskParam) FillValueToSkeleton() {
	if isEmpty(p.Name) {
		p.Name = []string{""}
	}
	if isEmpty(p.Id) {
		p.Id = []sacloud.ID{}
	}
	if isEmpty(p.From) {
		p.From = 0
	}
	if isEmpty(p.Max) {
		p.Max = 0
	}
	if isEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.QueryFile) {
		p.QueryFile = ""
	}

}

// Validate checks current values in model
func (p *ListProductDiskParam) Validate() []error {
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
		validator := define.Resources["ProductDisk"].Commands["list"].Params["id"].ValidateFunc
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

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ListProductDiskParam) GetResourceDef() *schema.Resource {
	return define.Resources["ProductDisk"]
}

func (p *ListProductDiskParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListProductDiskParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListProductDiskParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListProductDiskParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListProductDiskParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListProductDiskParam) SetName(v []string) {
	p.Name = v
}

func (p *ListProductDiskParam) GetName() []string {
	return p.Name
}
func (p *ListProductDiskParam) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListProductDiskParam) GetId() []sacloud.ID {
	return p.Id
}
func (p *ListProductDiskParam) SetFrom(v int) {
	p.From = v
}

func (p *ListProductDiskParam) GetFrom() int {
	return p.From
}
func (p *ListProductDiskParam) SetMax(v int) {
	p.Max = v
}

func (p *ListProductDiskParam) GetMax() int {
	return p.Max
}
func (p *ListProductDiskParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListProductDiskParam) GetSort() []string {
	return p.Sort
}
func (p *ListProductDiskParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListProductDiskParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListProductDiskParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListProductDiskParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListProductDiskParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListProductDiskParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListProductDiskParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListProductDiskParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListProductDiskParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListProductDiskParam) GetColumn() []string {
	return p.Column
}
func (p *ListProductDiskParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListProductDiskParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListProductDiskParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListProductDiskParam) GetFormat() string {
	return p.Format
}
func (p *ListProductDiskParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListProductDiskParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListProductDiskParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListProductDiskParam) GetQuery() string {
	return p.Query
}
func (p *ListProductDiskParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListProductDiskParam) GetQueryFile() string {
	return p.QueryFile
}

// ReadProductDiskParam is input parameters for the sacloud API
type ReadProductDiskParam struct {
	Assumeyes         bool       `json:"assumeyes"`
	ParamTemplate     string     `json:"param-template"`
	ParamTemplateFile string     `json:"param-template-file"`
	GenerateSkeleton  bool       `json:"generate-skeleton"`
	OutputType        string     `json:"output-type"`
	Column            []string   `json:"column"`
	Quiet             bool       `json:"quiet"`
	Format            string     `json:"format"`
	FormatFile        string     `json:"format-file"`
	Query             string     `json:"query"`
	QueryFile         string     `json:"query-file"`
	Id                sacloud.ID `json:"id"`
}

// NewReadProductDiskParam return new ReadProductDiskParam
func NewReadProductDiskParam() *ReadProductDiskParam {
	return &ReadProductDiskParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ReadProductDiskParam) FillValueToSkeleton() {
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.QueryFile) {
		p.QueryFile = ""
	}
	if isEmpty(p.Id) {
		p.Id = sacloud.ID(0)
	}

}

// Validate checks current values in model
func (p *ReadProductDiskParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ProductDisk"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadProductDiskParam) GetResourceDef() *schema.Resource {
	return define.Resources["ProductDisk"]
}

func (p *ReadProductDiskParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["read"]
}

func (p *ReadProductDiskParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ReadProductDiskParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ReadProductDiskParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ReadProductDiskParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ReadProductDiskParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *ReadProductDiskParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *ReadProductDiskParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadProductDiskParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadProductDiskParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadProductDiskParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadProductDiskParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadProductDiskParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadProductDiskParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadProductDiskParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadProductDiskParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadProductDiskParam) GetColumn() []string {
	return p.Column
}
func (p *ReadProductDiskParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadProductDiskParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadProductDiskParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadProductDiskParam) GetFormat() string {
	return p.Format
}
func (p *ReadProductDiskParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadProductDiskParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadProductDiskParam) SetQuery(v string) {
	p.Query = v
}

func (p *ReadProductDiskParam) GetQuery() string {
	return p.Query
}
func (p *ReadProductDiskParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ReadProductDiskParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *ReadProductDiskParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *ReadProductDiskParam) GetId() sacloud.ID {
	return p.Id
}
