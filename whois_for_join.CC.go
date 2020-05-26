{{define "whois"}}
    {{$roles := .Member.Roles}}{{$pos := 0}}{{$final := 0}}{{range .Guild.Roles}}{{if in $roles .ID}}{{if and (lt $pos .Position) (.Color)}}{{$pos = .Position}}{{$final = .Color}}{{end}}{{end}}{{end}}
    {{$var_roles := ""}}{{$counter := 0}}{{range $roles}}{{$counter = add 1 $counter}}{{$var_roles = (print $var_roles "<@&" . "> ")}}{{end}}
    {{$avatar_check := .User.Avatar}}{{$userAvatar := "https://i.imgur.com/joMi99X.png"}}{{if $avatar_check}}{{$userAvatar = (print "https://cdn.discordapp.com/avatars/" .User.ID "/" (toString .User.Avatar) ".png")}}{{end}}
    {{$un := ""}}{{$un_bool := false}}{{$un_counter := 0}}{{range pastUsernames .User.ID 0}}{{$un_counter = add 1 $un_counter}}{{if (le $un_counter 8)}}{{$un = (print $un "\n" (.Time.Format "Jan,06") ": " .Name)}}{{$un_bool = true}}{{end}}{{end}}
    {{$nn := ""}}{{$nn_bool := false}}{{$nn_counter := 0}}{{range pastNicknames .User.ID 0}}{{$nn_counter = add 1 $nn_counter}}{{if (le $nn_counter 8)}}{{$nn = (print $nn "\n" (.Time.Format "Jan,06") ": " .Name)}}{{$nn_bool = true}}{{end}}{{end}}
    {{$fields := cslice (sdict "name" "Account Age:" "value" currentUserAgeHuman "inline" true)}}
    {{if $un_bool}}{{$fields = $fields.Append (sdict "name" "Last 8 Usernames:" "value" (print "```" $un "```") "inline" false)}}{{end}}
    {{if $nn_bool}}{{$fields = $fields.Append (sdict "name" "Last 8 Nicknames:" "value" (print "```" $nn "```") "inline" false)}}{{end}}
    {{$embed := sdict
        "author" (sdict "name" "USER JOINED")
        "title" (print "User: " .User.String "\nID: " .User.ID)
        "description" (print "\n**Sticky Roles Added:** " $var_roles)
        "thumbnail" (sdict "url" $userAvatar)
        "color" $final
        "timestamp" currentTime
        "fields" $fields}}
    {{sendMessage nil (cembed $embed)}}{{end}}
{{sendTemplate nil "whois"}}
