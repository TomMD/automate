// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: api/external/compliance/reporting/reporting.proto

package reporting

import policy "github.com/chef/automate/components/automate-gateway/api/authz/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListReports", "compliance:reporting:reports", "search", "POST", "/compliance/reporting/reports", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListReportIds", "compliance:reporting:report-ids", "search", "POST", "/compliance/reporting/report-ids", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListControlItems", "compliance:reporting:controls", "search", "POST", "/compliance/reporting/controls", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*ControlItemRequest); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "text":
					return m.Text
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ReadReport", "compliance:reporting:reports:{id}", "read", "POST", "/compliance/reporting/reports/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListSuggestions", "compliance:reporting:suggestions", "search", "POST", "/compliance/reporting/suggestions", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*SuggestionRequest); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "type":
					return m.Type
				case "text":
					return m.Text
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListProfiles", "compliance:reporting:profiles", "search", "POST", "/compliance/reporting/profiles", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ReadNode", "compliance:reporting:nodes:{id}", "read", "GET", "/compliance/reporting/nodes/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListNodes", "compliance:reporting:nodes", "search", "POST", "/compliance/reporting/nodes/search", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/GetVersion", "compliance:reporting:version", "read", "GET", "/compliance/reporting/version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/LicenseUsageNodes", "compliance:reporting:licenseusage", "list", "", "", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}