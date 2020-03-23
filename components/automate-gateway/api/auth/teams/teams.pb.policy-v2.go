// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/auth/teams/teams.proto

package teams

import (
	request "github.com/chef/automate/components/automate-gateway/api/auth/teams/request"
	policyv2 "github.com/chef/automate/components/automate-gateway/authz/policy_v2"
)

func init() {
	policyv2.MapMethodTo("/chef.automate.api.teams.Teams/GetVersion", "system:service:version", "system:serviceVersion:get", "GET", "/auth/teams/version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.teams.Teams/GetTeams", "iam:teams", "iam:teams:list", "GET", "/auth/teams", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.teams.Teams/GetTeam", "iam:teams:{id}", "iam:teams:get", "GET", "/auth/teams/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetTeamReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policyv2.MapMethodTo("/chef.automate.api.teams.Teams/CreateTeam", "iam:teams", "iam:teams:create", "POST", "/auth/teams", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateTeamReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "name":
					return m.Name
				case "description":
					return m.Description
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.teams.Teams/UpdateTeam", "iam:teams:{id}", "iam:teams:update", "PUT", "/auth/teams/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateTeamReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "description":
					return m.Description
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.teams.Teams/DeleteTeam", "iam:teams:{id}", "iam:teams:delete", "DELETE", "/auth/teams/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteTeamReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policyv2.MapMethodTo("/chef.automate.api.teams.Teams/GetUsers", "iam:teams:{id}:users", "iam:teamUsers:list", "GET", "/auth/teams/{id}/users", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetUsersReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policyv2.MapMethodTo("/chef.automate.api.teams.Teams/AddUsers", "iam:teams:{id}", "iam:teamUsers:create", "POST", "/auth/teams/{id}/users", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.AddUsersReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policyv2.MapMethodTo("/chef.automate.api.teams.Teams/RemoveUsers", "iam:teams:{id}", "iam:teamUsers:delete", "PUT", "/auth/teams/{id}/users", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.RemoveUsersReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policyv2.MapMethodTo("/chef.automate.api.teams.Teams/GetTeamsForUser", "iam:users:{id}:teams", "iam:userTeams:get", "GET", "/auth/users/{id}/teams", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetTeamsForUserReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
}