package store

import (
	"encoding/json"

	restful "github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"github.com/sipcapture/heplify-server/pkg/api"
	"github.com/sipcapture/heplify-server/pkg/api/apierr"
)

func GetStoreHome(request *restful.Request, response *restful.Response) {

	w := response.ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	// encoded := request.Request.Header.Get("Authorization")
	statusCode := 200
	var output []byte
	var err error

	defer func() {
		w.WriteHeader(statusCode)
		w.Write(output)
	}()

	// _, err := auth.CheckToken(encoded)
	// if err != nil {
	// 	glog.Errorln("Unauth request ", err)
	// 	newErr := apierr.NewUnauthorized("invalid token")
	// 	output = api.EncodeError(newErr)
	// 	statusCode = 401
	// 	return
	// }
	home := api.DashboardStoreHome{
			Status: 200,
			Data: api.DashboardStoreHomeData {
				Rows: []api.DashboardStoreHomeDataRowsObj{
					{
						Columns: []api.DashboardStoreHomeDataRowsColumns{
							{
								StyleClass: "col-md-4",
                Widgets: []api.DashboardStoreHomeDataRowsColumnsWidgets{
                  {
                    Type: "clock",
                    Config: api.DashboardStoreHomeDataRowsColumnsWidgetsConfig{
                      TimePattern: "HH:mm:ss",
                      DatePattern: "YYYY-MM-DD",
                      Location: "America/Los_Angeles,America/New_York, Europe/Amsterdam",
                      Showseconds: false,
                    },
                    Title: "Clock",
                    Wid: 12,
                    TitleTemplateUrl: `../src/templates/widget-title.html`,
                  },
                  {
                    Type: "quicksearch",
                    Config: api.DashboardStoreHomeDataRowsColumnsWidgetsConfig{
                      Fields: []api.DashboardStoreHomeDataRowsColumnsWidgetsConfigFields{
                        {Name: "from_user",Selection: "From",},
                        {Name: "to_user",Selection: "To",},
                        {Name: "callid",Selection: "Call-ID",},
                        {Name: "transaction",Selection: "Transaction",},
                      },
                      Searchbutton: true,
                    },
                    Title: "Quick Search",
                    Wid: 3,
                    TitleTemplateUrl: `../src/templates/widget-title.html`,
                  },
                },
                Cid: 13,
							},
							{
                StyleClass: "col-md-8",
                Widgets: []api.DashboardStoreHomeDataRowsColumnsWidgets{
                  {
                    Type: "sipcaptureChart",
                    Config: api.DashboardStoreHomeDataRowsColumnsWidgetsConfig{
                      Panel: &api.DashboardStoreHomeDataRowsColumnsWidgetsConfigPanel{
                        Datasource: api.DashboardStoreHomeDataRowsColWidgetsCfgPanelDS{
                          Name: "Method",
                          Type: "JSON",
                          Settings: api.DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSet{
                            Path: `statistic\/method`,
                            Query:`{\n   \"timestamp\": {\n          \"from\": \"@from_ts\",\n          \"to\":  \"@to_ts\"\n   },\n  \"param\": {\n        \"filter\": [ \n             \"@filters\"\n       ],\n       \"limit\": \"@limit\",\n       \"total\": \"@total\"\n   }\n}`,
                            Method: "GET",
                            Limit: 200,
                            Total: false,
                            Eval: api.DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetEval{
                              Incoming: api.DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetEvalIncoming{
                                Name: "test incoming",
                                Value: "var object = @incoming; return object",
                              },
                            },
                            Timefields: []api.DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetTimefields{
                              {
                                Field: "from_ts",
                                Desc: "From Timestamp",
                              },
                              {
                                Field: "to_ts",
                                Desc: "To Timestamp",
                              },
                            },
                            Fieldsvalues: []api.DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetFieldval{
                              {
                                Field: "total",
                                Desc: "All Packets",
                              },
                            },
                            FIlters: []api.DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetfilters{
                              {
                                Type: "method",
                                Desc: "SIP Method",
                                Options: []api.DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetfiltersOpt{
                                  {Value: "!ALL"},
                                  {Value: "TOTAL"},
                                  {Value: "INVITE"},
                                  {Value: "UPDATE"},
                                  {Value: "REGISTER"},
                                  {Value: "CANCEL"},
                                  {Value: "BYE"},
                                  {Value: "OPTIONS"},
                                  {Value: "300"},
                                  {Value: "401"},
                                  {Value: "407"},
                                  {Value: "200"},
                                },
                              },
                              {
                                Type: "cseq",
                                Desc: "SIP Cseq",
                                Options: []api.DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetfiltersOpt{
                                    {Value: "INVITE"},
                                    {Value: "UPDATE"},
                                    {Value: "REGISTER"},
                                    {Value: "CANCEL"},
                                    {Value: "BYE"},
                                    {Value: "OPTIONS"},
                                },
                              },
                              {
                                Type: "auth",
                                Desc: "SIP Auth",
                                Options: []api.DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetfiltersOpt{
                                    {Value: "true"},
                                },
                              },
                              {
                                Type: "totag",
                                Desc: "SIP To Tag",
                                Options: []api.DashboardStoreHomeDataRowsColWidgetsCfgPanelDSSetfiltersOpt{
                                    {Value: "true"},
                                },
                              },
                            },
                          },
                        },
                        Values: []api.DashboardStoreHomeDataRowsColumnsWidgetsCfgPanelVal{
                        },
                        Filters: []api.DashboardStoreHomeDataRowsColWidgetsCfgPanelFilters{
                          {
                            Type: "method",
                            Value: "!TOTAL;ALL",
                          },
                          {
                            Type: "auth",
                            Value: "true",
                          },
                        },
                        Timefield: api.DashboardStoreHomeDataRowsColWidgetsCfgPanelTimefield{
                          Field: "to_ts",
                          Desc: "To Timestamp",
                        },
                        Limit: 1000,
                        Total: false,
                        Filter: api.DashboardStoreHomeDataRowsColWidgetsCfgPanelFilter{
                          Type: "auth",
                          Desc: "SIP Auth",
                          Options: []api.DashboardStoreHomeDataRowsColWidgetsCfgPanelFilterOpt{
                            {Value: "true"},
                          },
                        },
                        Filtervalue: api.DashboardStoreHomeDataRowsColumnsWidgetsCfgPanelFilterval{
                          Value: "true",
                        },
                      },
                      Query: `{\n   \"timestamp\": {\n          \"from\": \"@from_ts\",\n          \"to\":  \"@to_ts\"\n   },\n  \"param\": {\n        \"filter\": [ \n             \"@filters\"\n       ],\n       \"limit\": \"@limit\",\n       \"total\": \"@total\"\n   }\n}`,
                      Path: `statistic/method`,
                      Debugresp: "",
                      Chart: &api.DashboardStoreHomeDataRowsColumnsWidgetsConfigChart{
                        Type: api.DashboardStoreHomeDataRowsColumnsWidgetsConfigChartType{
                          Id: 2,
                          Label: "Line",
                          Value: "line",
                        },
                        Library: api.DashboardStoreHomeDataRowsColumnsWidgetsConfigChartLib{
                          Id: 3,
                          Label: "D3JS",
                          Value: "d3",
                        },
                        Size: api.DashboardStoreHomeDataRowsColumnsWidgetsConfigChartSize{
                          Height: 260,
                        },
                        Update: []api.DashboardStoreHomeDataRowsColumnsWidgetsConfigChartUpdate{
                          
                        },
                        Legend: api.DashboardStoreHomeDataRowsColumnsWidgetsConfigChartlegend{
                          Align: "center",
                          Layout: "horizontal",
                          Enabled: true,
                        },
                        Xaxis: api.DashboardStoreHomeDataRowsColumnsWidgetsConfigChartXaxis{
                          Title: "",
                        },
                        Yaxis: api.DashboardStoreHomeDataRowsColumnsWidgetsConfigChartYaxis{
                          Title: "Packets",
                        },
                      },
                    },
                    Title: "SIPCapture Charts",
                    Wid: 15,
                    TitleTemplateUrl: `../src/templates/widget-title.html`,
                  },
                },
                Cid: 14,
							},
						},
					},
				},
				Structure: "4-8",
				Title: "Home",
				Protect: true,
				SelectedItem: "fa-home",
				Alias: "home",
				Weight: "0",
				TitleTemplateUrl: `../src/templates/dashboard-title.html`,
			},
    }
    output, err = json.MarshalIndent(home, "  ", "  ")  
//  output, err = json.Marshal(home)

	if err != nil {
		glog.Errorln("store failure  ", err)
		newErr := apierr.NewInternalError(err.Error())
		output = api.EncodeError(newErr)
		statusCode = 500
		return
	}

	return
}
