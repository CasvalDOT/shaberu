# Shaberu

This small tool is born to quicky convert a i18n dictionary of a programming language into a CSV. This will help a programmer to reduce the time of annoing hand export


The process is simple!! Just take an input file path, provide an output path with a supported extension and automatically convert from a source to another.

---

## Supported sources
The source supported:

- `json`
- `csv`
- `php` dictionaries like `Laravel`

---

## Formatters

You can also prettify the document. Currently the supported formatters are:

- `squiz` for format `php` documents using `phpcbf`
- `ps` for format `json` documents using `prettier-standard`

You can add additional formatters editing the `~/.config/shaberu/formatters.json`. The structure of a formatter is:

```json
{
    ...

    "my-formatter" : {
        "binary": "my-beautifier",
        "args": ["--flag=1"]
    }
}
```

You can now use the formatter using the flag `-f` like:

```bash
./shaberu -i input.json -o output.csv -f my-formatter
```


---

## Example
Using the following command:

```bash
./shaberu -i /path/to/input.php -o /path/to/output.csv
```

Where `input.php` contains:
```php
<?php
    return [
        'title' => 'A title',
        'section_1' => [
            'content' => 'Lorem ipsum',
            'footer' => 'A footer'
        ]
    ];
```

This will return a `csv` like:

|   |   |
|---|---|
title|A title
section_1.content|Lorem ipsum
section_1.footer|A footer

Now if you need to translate the `csv` into `json`
you can perform:

```bash
./shaberu -i /path/to/output.csv -o /path/to/another_output.json -f ps
```
This will use `Prettier standard` as formatter and return:

```json
{
    "title": "A title",
    "section_1": {
        "content": "Lorem ipsum",
        "footer": "A footer"
    } 
}
```

