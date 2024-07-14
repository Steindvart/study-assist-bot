## Шифр цезаря

### Описание
Написать программу, которая будет применять шифр Цезаря к содержимому указанного файла и выводить преобразованное сообщение. Первый ввод строка — путь к файлу с содержимым. Второй ввод целое число — алфавитный сдвиг. Преобразование должно происходить только для букв латинского алфавита в верхнем и нижнем регистрах, остальные символы должны выводиться без изменений. Сдвиг должен быть зацикленным на самого себя, то есть `'a'` при сдвиге `-1` должна стать `z`, а `z` при сдвиге `1` должна стать `a`.

#### Дополнительные требования
1. Проверка корректности ввода. В случае некорректности ввода выводить "n/a".
2. Ввод - `stdin`, вывод - `stdout`.
3. В конце вывода не должно быть переноса строки.

### Предметная область
#### Преобразование символа

$$x = character + (shift \bmod alphabet)$$

- $\bmod$ - взятие остатка от деления.
- $alphabet$ - размер алфавита.

### Примеры входных и выходных данных

| Входные данные        | Содержимое файла               | Выходные данные                |
|-----------------------|--------------------------------|--------------------------------|
| `1.txt 1`             | `abcdefABCDEF`                 | `bcdefgBCDEFG`                 |
| `1.txt 3`             | `abcdefABCDEF`                 | `defghiDEFGHI`                 |
| `1.txt -1`            | `abcdefABCDEF`                 | `zabcdeZABCDE`                 |
| `2.txt 1`             | `a-b1c3d*e)1fhghn fDFDbK01234` | `b-c1d3e*f)1gihio gEGEcL01234` |
| `2.txt 55`            | `a-b1c3d*e)1fhghn fDFDbK01234` | `d-e1f3g*h)1ikjkq iGIGeN01234` |
| `3.txt 10`            | `z-y-x a-b-c Z-Y-X A-B-C`      | `j-i-h k-l-m J-I-H K-L-M`      |
| `3.txt 1`             | `z-y-x a-b-c Z-Y-X A-B-C`      | `a-z-y b-c-d A-Z-Y B-C-D`      |
| `3.txt -1`            | `z-y-x a-b-c Z-Y-X A-B-C`      | `y-x-w z-a-b Y-X-W Z-A-B`      |

Эта таблица демонстрирует различные входные значения, содержимое файлов и соответствующие результаты программы.

### Решения
Сначала попробуй самостоятельно решить :)

<details>
<summary>C</summary>

```c
#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

void caesarCipherOutput(FILE* content, int shift, int alphabet_size) {
  char ch = 0;
  while ((ch = getc(content)) != EOF) {
    char x = ch;
    if (isalpha(ch)) {
      char base = islower(ch) ? 'a' : 'A';
        x = (ch - base + shift) % alphabet_size;
        if (x < 0) {
          x += alphabet_size;
        }
        x += base;
      }

    printf("%c", x);
  }
}

int main() {
  char filepath[256];
  if (scanf("%255s", filepath) != 1) {
    printf("n/a");
    return 1;
  }

  int shift = 0;
  if (scanf("%d", &shift) != 1) {
    printf("n/a");
    return 1;
  }

  FILE *file = fopen(filepath, "r");
  if (file == NULL) {
    printf("n/a");
    return 1;
  }

  const int alphabet_size = ('z' - 'a') + 1;
  caesarCipherOutput(file, shift, alphabet_size);
  fclose(file);

  return 0;
}
```

</details>
