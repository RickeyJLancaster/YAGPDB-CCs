// CC:
//  Trigger: Regex `.*`
{{if not (dbGet .User.ID "cooldown_stickyroles")}}
    {{dbSetExpire .User.ID "cooldown_stickyroles" "cooldown" 7200}}
    {{$roles := .Member.Roles}}
    {{$data := (dbGet 69420 "stickyroles").Value}}

    {{if $data}}
        {{$data_converted := (sdict $data)}}
        {{if ($data_converted.Get (toString .User.ID))}}
            {{$retrieved_roles:= ($data_converted.Get (toString .User.ID))}}
            {{$isdifferent := true}}
            {{if eq (len $roles) (len $retrieved_roles)}}
                {{range $inx, $role := $roles}}{{if eq (index $retrieved_roles $inx) $role}}{{$isdifferent = false}}{{end}}{{end}}
                {{if $isdifferent}}
                    {{dbSet 69420 "stickyroles" ($data_converted.Set (toString .User.ID) $roles)}}
                {{end}}
            {{else}}
                {{dbSet 69420 "stickyroles" ($data_converted.Set (toString .User.ID) $roles)}}
            {{end}}
        {{else if not ($data_converted.Get (toString .User.ID))}}
            {{dbSet 69420 "stickyroles" ($data_converted.Set (toString .User.ID) $roles)}}
        {{end}}
    {{else if not $data}}
        {{dbSet 69420 "stickyroles" (sdict (toString .User.ID) $roles)}}
    {{end}}
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
