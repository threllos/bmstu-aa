# Лабораторная работа №3

## Задание

Реализовать:
- Сортировку вставками (*src/array/sort_insert.go*);
- Сортировку расчёской (*src/array/sort_comb.pdf*);
- Блинную сортировку (*src/array/sort_pancake.pdf*).

Написать отчет в tex (_docs/report.pdf_).

Unit-tests: *src/array/pkg_test.go*
Benchmark: *src/benchmark/benchmark.go*

## Программа
> _использовать в папке src_

**Пользовательский ввод:**
`make`
_или_
`make run`
_или_
`make build && ./app.exe`

**Ввод из файла:**
`make build && ./app.exe < test_data`

**Рандомный ввод:**
`make build && ./app.exe -r` — случайное заполнение
`make build && ./app.exe -ra` — заполнение по возрастанию
`make build && ./app.exe -rd` — заполнение по убыванию

**Изменить количество повторений:**
_по умолчанию 1000_
`make build && ./app.exe -repeats=5000`

**Запустить тестирование:**
`make test`

**Очистить папку:**
`make clean`

## Отчёт
> _использовать в папке docs/tex_

**Создать отчёт:**
`./tex-create.sh`

**Очистить папку:**
`./tex-clear.sh`