// CHECK IF USER HAS A ROLE OF THE LIST WITH .Member.Roles
{{$roles := (cslice RoleID1 RoleID2 RoleID3)}}
{{range .Member.Roles}}
    {{- if in $roles .}}
        do stuff
    {{end -}}
{{end}}

// CHECK IF USER HAS A ROLE OF THE LIST WITH targetHasRoleID
{{$roles := (cslice RoleID1 RoleID2 RoleID3)}}
{{range $roles}}
    {{if targetHasRoleID $.User.ID .}}
        do stuff
    {{end}}
{{end}}

