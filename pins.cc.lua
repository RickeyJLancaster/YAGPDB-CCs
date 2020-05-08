{{/* Trigger: Reaction Added Only */}}

{{$emoji := "📌"}} {{/* the emoji it will react to */}}

{{if eq (toString .Reaction.Emoji.Name) $emoji}}
    {{$channel := .Channel.ID}}
    {{$message := .ReactionMessage.ID}}
    {{$msglink := (joinStr "" "https://discordapp.com/channels/" .Server.ID "/" $channel "/" $message)}}
    {{$msg := .ReactionMessage}}
    {{$roles := .Member.Roles}}{{$pos := 0}}{{$final := 0}}{{range .Guild.Roles}}{{if in $roles .ID}}{{if and (lt $pos .Position) (.Color)}}{{$pos = .Position}}{{$final = .Color}}{{end}}{{end}}{{end}}


    {{$embedRaw := sdict
        "title" (joinStr " " "📌   Message Pin from" .Guild.Name "   📌")
        "color" $final
        "author" (sdict "name"  $msg.Author.Username "icon_url" ($msg.Author.AvatarURL "64")) 
        "timestamp" $msg.Timestamp}}

    {{if $msg.Attachments}}
        {{$file := ((index $msg.Attachments 0).URL)}}
        {{$format := " "}}
        {{if (reFind "(.jpg|.jpeg|.png|.gif|.tif|.tiff|.gifv)$" $file) }}
            {{$embedRaw.Set "description" (joinStr "" "**[Message Link](" $msglink ")  to <#" $channel ">**\n\n" $msg.Content)}}
            {{$embedRaw.Set "image" (sdict "url" (index $msg.Attachments 0).URL) }}
        {{else}}
            {{$attachement := (joinStr "" "⚠️ **The attachment cannot be displayed. [Click here](" $file ") to open it.**")}}
            {{$embedRaw.Set "description" (joinStr "" "**[Message Link](" $msglink ")  to <#" $channel ">**\n\n" $msg.Content "\n\n" $attachement)}}
        {{end}}
    {{else}}
            {{$embedRaw.Set "description" (joinStr "" "**[Message Link](" $msglink ")  to <#" $channel ">**\n\n" $msg.Content)}}
    {{end}}

    {{sendDM (cembed $embedRaw)}}
    {{sleep 5}}
    {{deleteMessageReaction $channel $message .User.ID .Reaction.Emoji.Name}}
{{end}}
