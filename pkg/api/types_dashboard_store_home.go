package api
type DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetTimefields struct {
	Field string `json:"field,omitempty"`
	Desc string `json:"desc,omitempty"`
}
type DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetEvalIncoming struct{
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetEval struct {
	Incoming DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetEvalIncoming `json:"incoming,omitempty"`
}

type DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetFieldval struct{
	Field string `json:"field,omitempty"`
	Desc string `json:"desc,omitempty"`
}

type DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetfiltersOpt struct {
	Value string `json:"value,omitempty"`
}

type DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetfilters struct {
	Type string `json:"type,omitempty"`
	Desc string `json:"desc,omitempty"`
	Options []DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetfiltersOpt `json:"options,omitempty"`
}

type DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSet struct{
	Path string `json:"path,omitempty"`
	Query string `json:"query,omitempty"`
	Method string `json:"method,omitempty"`
	Limit int `json:"limit,omitempty"`
	Total bool `json:"total"`
	Eval DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetEval `json:"eval,omitempty"`
	Timefields []DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetTimefields 	`json:"timefields,omitempty"`
	Fieldsvalues []DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetFieldval `json:"fieldvalues,omitempty"`
	FIlters []DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetfilters `json:"filters,omitempty"`
}

type DashboardStoreHomeDataRowsColWidgetsCfgPanelDS struct{
	Name string `json:"name,omitempty"`
    Type string `json:"type,omitempty"`
    Settings DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSet `json:"settings,omitempty"`
}

type DashboardStoreHomeDataRowsColumnsWidgetsCfgPanelVal struct {

}

type DashboardStoreHomeDataRowsColWidgetsCfgPanelFilters struct{
	Type string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type DashboardStoreHomeDataRowsColWidgetsCfgPanelTimefield struct {
	Field string `json:"field,omitempty"`
	Desc string `json:"desc,omitempty"`
}

type DashboardStoreHomeDataRowsColumnsWidgetsCfgPanelFilterval struct{
	Value string `json:"value,omitempty"`
}

type DashboardStoreHomeDataRowsColWidgetsCfgPanelFilter struct {
	Type string `json:"type,omitempty"`
	Desc string `json:"desc,omitempty"`
	Options []DashboardStoreHomeDataRowsColWidgetsCfgPanelFilterOpt `json:"options,omitempty"`
}

type DashboardStoreHomeDataRowsColWidgetsCfgPanelFilterOpt struct{
	Value string `json:"value,omitempty"`
}

type DashboardStoreHomeDataRowsColumnsWidgetsConfigPanel struct{
	Datasource DashboardStoreHomeDataRowsColWidgetsCfgPanelDS `json:"datasource,omitempty"` 
	Values []DashboardStoreHomeDataRowsColumnsWidgetsCfgPanelVal `json:"values"`
	Filters []DashboardStoreHomeDataRowsColWidgetsCfgPanelFilters `json:"filters,omitempty"`
	Timefield DashboardStoreHomeDataRowsColWidgetsCfgPanelTimefield `json:"timefield,omitempty"`
	Limit int `json:"limit,omitempty"`
	Total bool `json:"total"`
	Filter DashboardStoreHomeDataRowsColWidgetsCfgPanelFilter `json:"filter,omitempty"`
	Filtervalue DashboardStoreHomeDataRowsColumnsWidgetsCfgPanelFilterval `json:"filtervalue,omitempty"`
}

type DashboardStoreHomeDataRowsColumnsWidgetsConfigQuery struct {

}

type DashboardStoreHomeDataRowsColumnsWidgetsConfigChartYaxis struct {
	Title string `json:"title,omitempty"`
}

type DashboardStoreHomeDataRowsColumnsWidgetsConfigChartXaxis struct {
	Title string `json:"title"`
}

type DashboardStoreHomeDataRowsColumnsWidgetsConfigChartlegend struct {
	Align string `json:"align,omitempty"`
	Layout string `json:"layout,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
}

type DashboardStoreHomeDataRowsColumnsWidgetsConfigChartUpdate struct {

}

type DashboardStoreHomeDataRowsColumnsWidgetsConfigChartSize struct {
	Height int `json:"height,omitempty"`
}

type DashboardStoreHomeDataRowsColumnsWidgetsConfigChartLib struct{
	Id int `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}

type DashboardStoreHomeDataRowsColumnsWidgetsConfigChartType struct {
	Id int `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}

type DashboardStoreHomeDataRowsColumnsWidgetsConfigChart struct {
	Type DashboardStoreHomeDataRowsColumnsWidgetsConfigChartType `json:"type"`
	Library DashboardStoreHomeDataRowsColumnsWidgetsConfigChartLib `json:"library"`
	Size DashboardStoreHomeDataRowsColumnsWidgetsConfigChartSize `json:"size"`
	Update []DashboardStoreHomeDataRowsColumnsWidgetsConfigChartUpdate `json:"update"`
	Legend DashboardStoreHomeDataRowsColumnsWidgetsConfigChartlegend `json:"legend"`
	Xaxis DashboardStoreHomeDataRowsColumnsWidgetsConfigChartXaxis `json:"xaxis"`
	Yaxis DashboardStoreHomeDataRowsColumnsWidgetsConfigChartYaxis `json:"yaxis"`
}

type DashboardStoreHomeDataRowsColumnsWidgetsConfigFields struct {
	Name string `json:"name"`
	Selection string `json:"selection"`
}
type DashboardStoreHomeDataRowsColumnsWidgetsConfig struct{
	TimePattern string `json:"timePattern,omitempty"`
	DatePattern string `json:"datePattern,omitempty"`
	Location string `json:"location,omitempty"`
	Showseconds bool `json:"showseconds,omitempty"`
	Fields []DashboardStoreHomeDataRowsColumnsWidgetsConfigFields `json:"fields,omitempty"` 
	Searchbutton bool `json:"searchbutton,omitempty"`
	Panel *DashboardStoreHomeDataRowsColumnsWidgetsConfigPanel `json:"panel,omitempty"`
	Query string `json:"query,omitempty"`
	Path string `json:"path,omitempty"`
	Debugresp string `json:"debugresp,omitempty"`
	Chart *DashboardStoreHomeDataRowsColumnsWidgetsConfigChart `json:"chart,omitempty"`
}

type DashboardStoreHomeDataRowsColumnsWidgets struct{
	Type string `json:"type"`
	Config DashboardStoreHomeDataRowsColumnsWidgetsConfig `json:"config"`
	Title string `json:"title"`
	Wid int `json:"wid"`
	TitleTemplateUrl string `json:"titleTemplateUrl"`
}

type DashboardStoreHomeDataRowsColumns struct{
	StyleClass string `json:"styleClass"`
	Widgets []DashboardStoreHomeDataRowsColumnsWidgets `json:"widgets"`
	Cid int `json:"cid"`
}

type DashboardStoreHomeDataRowsObj struct{
	Columns []DashboardStoreHomeDataRowsColumns `json:"columns"`
}

type DashboardStoreHomeData struct{
  Rows	[]DashboardStoreHomeDataRowsObj `json:"rows"`
  Structure string `json:"structure"`
  Title string `json:"title"`
  Protect bool `json:"protect"`
  SelectedItem string  `json:"selectedItem"`
  Alias string `json:"alias"`
  Weight string `json:"weight"`
  TitleTemplateUrl string `json:"titleTemplateUrl"`
}

//DashboardHome /api/v1/dashboard/store/home
type DashboardStoreHome struct {
	Status int `json:"status"`
	Data DashboardStoreHomeData `json:"data"`
}