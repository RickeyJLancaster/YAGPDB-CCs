{{$Final:=0}}
{{$Pos:=0}}
{{range .Guild.Roles}}
    {{- if in $.Member.Roles .ID -}}
        {{- if and (lt $Pos .Position) (.Color) -}}
            {{- $Pos = .Position -}}
            {{- $Final = .Color -}}
        {{- end -}}
    {{- end -}}
{{end}}
{{$Roles := ""}}
{{$Counter := 0}}
{{range .Member.Roles}}
    {{- $Counter = add 1 $Counter -}}
    {{- $Roles = (print $Roles "<@&" . "> ") -}}
{{end}}
{{if $Roles}}
    {{$Roles = print "\n**Roles:** " $Roles}}
{{end}}
{{$fields := cslice (sdict "name" "Account Age:" "value" currentUserAgeHuman "inline" true)}}
{{$Un :=""}}
{{$User := pastUsernames .User.ID 0}}
{{if gt (len $User) 8}}
    {{$User = slice $User 0 8}}
{{end}}
{{range $User}}
    {{- $Un = (print $Un "\n" (.Time.Format "Jan,06") ": " .Name) -}}
{{end}}
{{if $Un}}
    {{$fields = $fields.Append (sdict "name" "Last 8 Usernames:" "value" (print "```" $Un "```") "inline" false)}}
{{end}}
{{$Nn :=""}}
{{$Nick := pastNicknames .User.ID 0}}
{{if gt (len $Nick) 8}}
    {{$Nick = slice $Nick 0 8}}
{{end}}
{{range $Nick}}
    {{- $Nn = (print $Nn "\n" (.Time.Format "Jan,06") ": " .Name) -}}
{{end}}
{{if $Nn}}
    {{$fields = $fields.Append (sdict "name" "Last 8 Nicknames:" "value" (print "```" $Nn "```") "inline" false)}}
{{end}}
{{$Avatar:="https://i.imgur.com/joMi99X.png"}}
{{if .User.Avatar}}
    {{$Avatar = .User.AvatarURL "256"}}
{{end}}
{{$embed := sdict
    "author" (sdict "name" "USER JOINED")
    "title" (print "User: " .User.String "\nID: " .User.ID)
    "description" $Roles
    "thumbnail" (sdict "url" $Avatar)
    "color" $Final
    "timestamp" currentTime
    "fields" $fields}}
{{sendMessage nil (cembed $embed)}}
