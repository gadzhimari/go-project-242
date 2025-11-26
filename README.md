### Tests and linter status:

[![Actions Status](https://github.com/gadzhimari/go-project-242/actions/workflows/main.yml/badge.svg)](https://github.com/gadzhimari/go-project-242/actions)

# hexlet-path-size

Утилита для подсчёта размера файла или директории. Поддерживает рекурсивный подсчёт, человеко-читаемый вывод и учёт скрытых файлов.

## Демо

[![asciicast](https://asciinema.org/a/UJeVZTsfG0oSruBGvj92DZ2Mg.svg)](https://asciinema.org/a/UJeVZTsfG0oSruBGvj92DZ2Mg)

## Установка

```bash
git clone https://github.com/gadzhimari/go-project-242.git
cd hexlet-path-size
make build
```

## Флаги

| Флаг          | Алиас | Описание                                              |
| ------------- | ----- | ----------------------------------------------------- |
| `--recursive` | `-r`  | Считать размер директорий рекурсивно                  |
| `--human`     | `-H`  | Вывод в человеко-читаемом формате (KB, MB, GB и т.д.) |
| `--all`       | `-a`  | Учитывать скрытые файлы и директории                  |
