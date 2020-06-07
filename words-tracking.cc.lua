{{/*
        Words Tracking Command. Track messages from all the server and send in logs channel.
                   To add terms replace the examples in (word1|word2|word3).
                  The (?i) at the beggining make the tracking case insensitive
                            Trigger Type: Regex      |      Trigger: (?)
*/}}


{{$logchannel := Channel-ID-Here}}
{{$wordslist := `(?i)(word1|word2|word3|word4)`}}

{{if reFind $wordslist .StrippedMsg}}
    {{$channel := .Channel.ID}}
    {{$msg := .Message}}
    {{$msglink := (joinStr "" "https://discordapp.com/channels/" .Server.ID "/" $channel "/" $msg.ID)}}
    {{$roles := .Member.Roles}}{{$pos := 0}}{{$final := 0}}{{range .Guild.Roles}}{{if in $roles .ID}}{{if and (lt $pos .Position) (.Color)}}{{$pos = .Position}}{{$final = .Color}}{{end}}{{end}}{{end}}
    {{$trigger := reFind $wordslist .StrippedMsg}}

    {{$embedRaw := sdict
        "color" $final
        "author" (sdict "name"  $msg.Author.Username "icon_url" ($msg.Author.AvatarURL "64")) 
        "timestamp" $msg.Timestamp}}

    {{if $msg.Attachments}}
        {{$file := ((index $msg.Attachments 0).URL)}}
        {{$format := " "}}
        {{if (reFind "(.jpg|.jpeg|.png|.gif|.tif|.tiff|.gifv)$" $file) }}
            {{$embedRaw.Set "description" (joinStr "" "**[Message Link](" $msglink ")  to <#" $channel ">\nTrigger:** " $trigger "\n" $msg.Content)}}
            {{$embedRaw.Set "image" (sdict "url" (index $msg.Attachments 0).URL) }}
        {{else}}
            {{$attachement := (joinStr "" "⚠️ **The attachment cannot be displayed. [Click here](" $file ") to open it.**")}}
            {{$embedRaw.Set "description" (joinStr "" "**[Message Link](" $msglink ")  to <#" $channel ">\nTrigger:** " $trigger "\n\n" $attachement)}}
        {{end}}
    {{else}}
            {{$embedRaw.Set "description" (joinStr "" "**[Message Link](" $msglink ")  to <#" $channel ">\nTrigger:** " $trigger "\n" $msg.Content)}}
    {{end}}

    {{sendMessage $logchannel (cembed $embedRaw)}}
{{end}}
