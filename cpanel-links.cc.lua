{{$title := "<:yagpdb:505114640032858114> Control Panel Menu"}}
{{$id := .Guild.ID}}
{{$avatar := (sdict "url" "https://cdn.discordapp.com/avatars/204255221017214977/2fa57b425415134d4f8b279174131ad6.png?size=1024")}}
{{$home := (joinStr "" "**[Home](https://yagpdb.xyz/manage/" $id "/home)**")}}
{{$stats := (joinStr "" "**[Stats](https://yagpdb.xyz/manage/" $id "/stats)**")}}
{{$discovery := (joinStr "" "**[Server Discovery](https://yagpdb.xyz/manage/" $id "/serverdiscovery)**")}}
{{$core := (joinStr "" "**[Core](https://yagpdb.xyz/manage/" $id "/core)**" "\n ‎ ‎ ‎ [Control Panel Logs](https://yagpdb.xyz/manage/" $id "/cplogs)" "\n ‎ ‎ ‎ [Command Settings](https://yagpdb.xyz/manage/" $id "/commands/settings)" "\n ‎ ‎ ‎ [Custom Commands](https://yagpdb.xyz/manage/" $id "/customcommands)")}}
{{$feeds := (joinStr "" "**Notif. & Feeds**" "\n ‎ ‎ ‎ [General](https://yagpdb.xyz/manage/" $id "/notifications/general)" "\n ‎ ‎ ‎ [Reddit](https://yagpdb.xyz/manage/" $id "/reddit)" "\n ‎ ‎ ‎ [Streaming](https://yagpdb.xyz/manage/" $id "/streaming)" "\n ‎ ‎ ‎ [YouTube](https://yagpdb.xyz/manage/" $id "/youtube)" "\n ‎ ‎ ‎ [Twitter](https://yagpdb.xyz/manage/" $id "/twitter)")}}
{{$tools := (joinStr "" "**Tools & Utilities**" "\n ‎ ‎ ‎ [Moderation](https://yagpdb.xyz/manage/" $id "/moderation)" "\n ‎ ‎ ‎ [Automod](https://yagpdb.xyz/manage/" $id "/automod_legacy)" "\n ‎ ‎ ‎ [Automod v2](https://yagpdb.xyz/manage/" $id "/automod)" "\n ‎ ‎ ‎ [Logging](https://yagpdb.xyz/manage/" $id "/logging)" "\n ‎ ‎ ‎ [Autorole](https://yagpdb.xyz/manage/" $id "/autorole)" "\n ‎ ‎ ‎ [Role Commands](https://yagpdb.xyz/manage/" $id "/rolecommands)" "\n ‎ ‎ ‎ [System Tickets](https://yagpdb.xyz/manage/" $id "/tickets/settings/)" "\n ‎ ‎ ‎ [Verification](https://yagpdb.xyz/manage/" $id "/verification/)")}}
{{$fun := (joinStr "" "**Fun**" "\n ‎ ‎ ‎ [Reputation](https://yagpdb.xyz/manage/" $id "/reputation/)" "\n ‎ ‎ ‎ [Soundboard](https://yagpdb.xyz/manage/" $id "/soundboard/)")}}

{{if (reFind "(home)" .StrippedMsg)}}
    {{$embed := cembed
        "title" $title
        "description" $home
        "thumbnail" $avatar}}
    {{sendMessage nil $embed}}
{{else if (reFind "(stats)" .StrippedMsg)}}
    {{$embed := cembed
        "title" $title
        "description" $stats
        "thumbnail" $avatar}}
    {{sendMessage nil $embed}}
{{else if (reFind "(discovery)" .StrippedMsg)}}
    {{$embed := cembed
        "title" $title
        "description" $discovery
        "thumbnail" $avatar}}
    {{sendMessage nil $embed}}
{{else if (reFind "(core|logs|panel logs|command settings|cc|custom command.)" .StrippedMsg)}}
    {{$embed := cembed
        "title" $title
        "description" $core
        "thumbnail" $avatar}}
    {{sendMessage nil $embed}}
{{else if (reFind "(notification|notif|feeds|general|reddit|streaming|stream|youtube|twitter)" .StrippedMsg)}}
    {{$embed := cembed
        "title" $title
        "description" $feeds
        "thumbnail" $avatar}}
    {{sendMessage nil $embed}}
{{else if (reFind "(tools|moderation|automod|logging|autorole|role commands|ticket|tickets|verifiation)" .StrippedMsg)}}
    {{$embed := cembed
        "title" $title
        "description" $tools
        "thumbnail" $avatar}}
    {{sendMessage nil $embed}}
{{else if (reFind "(fun|reputation|sb|soundboard)" .StrippedMsg)}}
    {{$embed := cembed
        "title" $title
        "description" $fun
        "thumbnail" $avatar}}
    {{sendMessage nil $embed}}
{{else}}
    {{$embed := cembed
        "title" $title
        "description" (joinStr "" $home "\n" $stats "\n" $discovery "\n\n" $core "\n\n" $feeds "\n\n" $tools "\n\n" $fun)
        "thumbnail" $avatar}}
    {{sendMessage nil $embed}}
{{end}}
