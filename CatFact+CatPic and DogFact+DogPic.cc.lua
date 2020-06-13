{{/*
DogFact+DogImage and CatFact+CatImage Commands
To avoid duplication of the response disable the original CatFact and DogFact commands on Core > Command Settings > Command Override
> Trigger Type: Regex
> Trigger: ^-(dog|doggo|dogfact|cat|kitten|catfact)
*/}}
{{$stuff:=""}}
{{$link:=""}}
{{if (reFind "(dog|doggo|dogfact)" .Cmd)}}
    {{$stuff = cslice "Doggo" "Dog Fact:" "dogfact"}}
    {{$link = joinStr "" "https://placedog.net/500/280/id="  (urlescape (toString (randInt 1000)))}}
{{else if (reFind "(cat|kitten|catfact)" .Cmd)}}
    {{$stuff = cslice "Kitten" "Cat Fact:" "catfact"}}
    {{$link = (print "http://placekitten.com/" (randInt 300 500) "/" (randInt 200 600) "?image=" (urlescape (toString (randInt 0 16))))}}
{{end}}
{{with $stuff}}
    {{sendMessage nil (cembed "title" (index . 1) "description" (exec (index . 2)) "image" (sdict "url" $link) "footer" (sdict "text" (print (index . 0) " requested by " $.User.Username) "icon_url" ($.User.AvatarURL "256")) "color" (randInt 111111 999999))}}
{{end}}
