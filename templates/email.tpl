<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Product List</title>
</head>
<body>
    {{range $key, $value := .}}
    <h1>Product List for {{ $key }}</h1>


    {{range $result := $value}}
    <ul>
        <li>
            <a href="{{ $result.ProductURL }}"><img src="{{ $result.ImageURL }}"></a>
            <strong>{{ $result.Name }} </strong><br>
            Price: {{ $result.Price }}<br>
        </li>
    </ul>
    {{end}}
    {{end}}
</body>
</html>
`
