{{$prob := 10}} {{/* --> Chance (1/100) of triggering the command     */}}

{{$rand := randInt 0 100}}
{{$nospace := (reReplace "([ ])" .StrippedMsg "-")}}
{{$length := (len $nospace)}}
{{$prefix := "mao-"}}

{{if not .Message.Attachments}}
    {{if not (reFind `(<@!?\d+>|www.|https)` .StrippedMsg)}}
        {{if lt $rand $prob}}
            {{if lt $length 96}}
                {{editChannelName nil (joinStr "" $prefix $nospace)}}
            {{else}}
            {{$crop := (slice $nospace 0 96)}}
                {{editChannelName nil (joinStr "" $prefix $crop)}}
            {{end}}
        {{end}}
    {{end}}
{{end}}
