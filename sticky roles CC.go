//  CC:
//  Trigger: Regex `.*`

{{if not (dbGet .User.ID "cooldown_stickyroles")}}
    {{dbSetExpire .User.ID "cooldown_stickyroles" "cooldown" 7200}}
    {{$roles := .Member.Roles}}
    {{$data := (dbGet 69420 "stickyroles").Value}}
    {{if $data}}
        {{$data_converted := (sdict $data)}}
        {{if ($data_converted.Get (toString .User.ID))}}
            {{$retrieved_roles:= ($data_converted.Get (toString .User.ID))}}
            {{$isdifferent := false}}
            {{if eq (len $roles) (len $retrieved_roles)}}
                {{range $index, $role := $roles}}
                    {{if not (in $retrieved_roles $role)}}
                        {{$isdifferent = true}}
                    {{end}}
                {{end}}
                {{if $isdifferent}}
                    {{$data_converted.Set (toString .User.ID) $roles}}
                    {{dbSet 69420 "stickyroles" $data_converted}}
                {{end}}
            {{else}}
                {{$data_converted.Set (toString .User.ID) $roles}}
                {{dbSet 69420 "stickyroles" $data_converted}}
            {{end}}
        {{else}}
            {{$data_converted.Set (toString .User.ID) $roles}}
            {{dbSet 69420 "stickyroles" $data_converted}}
        {{end}}
    {{else}}
        {{dbSet 69420 "stickyroles" (sdict (toString .User.ID) $roles)}}
    {{end}}
{{end}}

// Kick DM and Ban DM
{{dbSetExpire .User.ID "ban_kick" 1 3600}}

// Leave Message:
{{$ban_kick := (dbGet .User.ID "ban_kick")}}
{{if $ban_kick}}
    {{$data := (dbGet 69420 "stickyroles").Value}}
    {{$data_converted := sdict $data}}
    {{$userdata := $data_converted.Get (toString .User.ID)}}
    {{if $userdata}}
        {{dbSet 69420 "stickyroles" ($data_converted.Del (toString .User.ID))}}
    {{end}}
{{end}}

// Join Message:
{{$data := (dbGet 69420 "stickyroles").Value}}
{{$data_converted := sdict $data}}
{{$userdata := $data_converted.Get (toString .User.ID)}}
{{sleep 1}}
{{if $userdata}}
    {{range $userdata}}
        {{addRoleID .}}
    {{end}}
{{end}}
