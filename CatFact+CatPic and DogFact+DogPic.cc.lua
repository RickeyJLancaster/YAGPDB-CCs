{{/*  ->    DogFact+DogImage and CatFact+CatImage Commands      <-    */}}
{{/*  ->     To avoid duplication of the response disable       <-    */}}
{{/*  ->       the original CatFact and DogFact commands        <-    */}}
{{/*  ->     on Core > Command Settings > Command Override      <-    */}}

{{/*  ->     Trigger Type: Regex                                <-    */}}
{{/*  ->     Trigger: ^-(dog|doggo|dogfact|cat|kitten|catfact)  <-    */}}


{{if (reFind "(dog|doggo|dogfact)" .Cmd)}}
    {{$title := "Dog Fact:"}}
    {{$fact := (exec "dogfact")}}
    {{$x := randInt 1000}}
    {{$link := joinStr "" "https://placedog.net/500/280/id="  (urlescape (toString $x))}}
    {{$footer_img:= (joinStr "" "https://cdn.discordapp.com/avatars/" (toString .User.ID) "/" .User.Avatar ".png")}}
    {{$footer_txt := (joinStr " " "Doggo requested by" .User.Username)}}
    {{sendMessage nil (cembed "title" $title "description" $fact "image" (sdict "url" $link) "footer" (sdict "text" $footer_txt "icon_url" $footer_img))}}
{{else if (reFind "(cat|kitten|catfact)" .Cmd)}}
    {{$title := "Cat Fact:"}}
    {{$fact := (exec "catfact")}}
    {{$h := randInt 300 500}}
    {{$w := randInt 200 600}}
    {{$x := randInt 0 16}}
    {{$link := (joinStr "" "http://placekitten.com/" $h "/" $w "?image=" (urlescape (toString $x)))}}
    {{$footer_img:= (joinStr "" "https://cdn.discordapp.com/avatars/" (toString .User.ID) "/" .User.Avatar ".png")}}
    {{$footer_txt := (joinStr " " "Kitten requested by" .User.Username)}}
    {{sendMessage nil (cembed "title" $title "description" $fact "image" (sdict "url" $link) "footer" (sdict "text" $footer_txt "icon_url" $footer_img))}}
{{else}}{{end}}
