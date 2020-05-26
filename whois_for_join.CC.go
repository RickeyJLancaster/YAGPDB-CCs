{{define "whois"}}
    {{$roles := .Member.Roles}}{{$pos := 0}}{{$final := 0}}{{range .Guild.Roles}}{{if in $roles .ID}}{{if and (lt $pos .Position) (.Color)}}{{$pos = .Position}}{{$final = .Color}}{{end}}{{end}}{{end}}
    {{$var_roles := ""}}{{$counter := 0}}{{range $roles}}{{$counter = add 1 $counter}}{{$var_roles = (print $var_roles "<@&" . "> ")}}{{end}}
    {{$un := ""}}{{$un_counter := 0}}{{range pastUsernames .User.ID 0}}{{$un_counter = add 1 $un_counter}}{{if (le $un_counter 8)}}{{$un = (print $un "\n" (.Time.Format "Jan,06") ": " .Name)}}{{end}}{{end}}{{$un := print "```" $un "```"}}
    {{$nn_ := ""}}{{$nn_counter := 0}}{{range pastNicknames .User.ID 0}}{{$nn_counter = add 1 $nn_counter}}{{if (le $nn_counter 8)}}{{$nn_ = (print $nn_ "\n" (.Time.Format "Jan,06") ": " .Name)}}{{end}}{{end}}{{$nn := print "```" $nn_ "```"}}

    {{$embed := sdict
        "author" (sdict "name" "USER JOINED")
        "title" (print "User: " .User.String "\nID: " .User.ID)
        "description" (print "\n**Sticky Roles Added:** " $var_roles)
        "thumbnail" (sdict "url" (print "https://cdn.discordapp.com/avatars/" .User.ID "/" (toString .User.Avatar) ".png"))
        "color" $final
        "timestamp" currentTime
        "fields" (cslice
            (sdict "name" "Account Age" "value" currentUserAgeHuman)
            (sdict "name" "Last Usernames:" "value" $un "inline" false)
        )
    }}
    {{if ne $nn_ ""}}
        {{$embed.Set "fields"  (cslice
            (sdict "name" "Account Age" "value" currentUserAgeHuman)
            (sdict "name" "Last 8 Usernames:" "value" $un "inline" false)
            (sdict "name" "Last 8 Nicknames:" "value" $nn "inline" false)
        )}}
    {{end}}
    {{sendMessage nil (cembed $embed)}}
{{end}}
{{sendTemplate nil "whois"}}
