


// CC 1:
//  Trigger: Regex `.*`
## TODO: ADD COOLDOWN
{{$roles := .Member.Roles}}
{{$data := (dbGet 69420 "stickyroles").Value}}
{{$data_converted := sdict $data}}

{{if ($data_converted.Get (toString .User.ID))}}
    {{$retrieved_roles:= ($data_converted.Get (toString .User.ID))}}
    {{$isequal := false}}
    {{if eq (len $roles) (len $retrieved_roles)}}
        {{range $inx, $role := $roles}}{{if eq (index $retrieved_roles $inx) $role}}{{$isequal = true}}{{end}}{{end}}
        {{if not $isequal}}
            {{dbSet 69420 "stickyroles" (sdict (toString .User.ID) $roles)}}
        {{end}}
    {{else}}
        {{dbSet 69420 "stickyroles" (sdict (toString .User.ID) $roles)}}
    {{end}}
{{else if not ($data_converted.Get (toString .User.ID))}}
    {{dbSet 69420 "stickyroles" (sdict (toString .User.ID) $roles)}}
{{end}}


// Kick DM and Ban DM
{{dbSetExpire .User.ID "ban_kick" 1 86400}}


// Leave Message:
{{$ban_kick := (dbGet .User.ID "ban_kick")}}
{{if $ban_kick}}
    {{$data := (dbGet 69420 "stickyroles").Value}}
    {{$data_converted := sdict $data}}
    {{if ($data_converted.Get (toString .User.ID))}}
        {{dbSet 69420 "stickyroles" ($data_converted.Del (toString .User.ID))}}
    {{end}}
{{end}}


// Join Message:
{{$data := dbGet 69420 "stickyroles"}}
{{$data_converted := sdict $data}}

{{if ($data_converted.Get .User.ID)}}
    {{$roles := ($data_converted.Get .User.ID)}}
    {{range $roles}}
        {{addRoleID .}}
    {{end}}
{{end}}



