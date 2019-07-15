package rbac.authz

# Policy decisions are made using the following input
#   Not part of the policy but good as documentation
# input = {
#     "subject": "bob",
#     "action": "read",
#     "resource": "server123"
# }

# user-role assignments
# user_roles = data.bundle.user.role

# role-permissions assignments
# role_permissions = data.bundle.role.permission
# logic that implements RBAC.
bundle_name = sprintf("%v", [input.service])

default allow = false
allow {
    # lookup the list of roles for the user
    roles := data[bundle_name].user.role[input.subject]
    # for each role in that list
    r := roles[_]
    # lookup the permissions list for role r
    permissions := data[bundle_name].role.permission[r]
    # for each permission
    p := permissions[_]
    # check if the permission granted to r matches the user's request
    p == {"action": input.action, "object": input.resource}
}
