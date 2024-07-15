## Перестановка максимального и минимального числа

### Описание
Написать программу, которая будет получать на ввод последовательность целых неотрицательных чисел, менять местами первое вхождение максимального и минимального числа в этой последовательности и выводить изменённую последовательность. Входная последовательность может быть любого размера, окончанием последовательности является число `-1` и оно не является частью последовательности.

#### Дополнительные требования
1. Ввод - `stdin`, вывод - `stdout`.
2. Проверка корректности ввода. В случае некорректности ввода выводить "n/a".
3. При выводе результата, после последнего числа не должно быть пробела.
4. В конце вывода не должно быть переноса строки.

### Примеры входных и выходных данных

| Входные данные       | Выходные данные   |
|----------------------|-------------------|
| `1 2 3 4 5 -1`       | `5 2 3 4 1`       |
| `-1`                 |                   |
| `1 -1`               | `1`               |
| `5 9 9 1 5 9 9 5 -1` | `5 1 9 9 5 9 9 5` |


### Решения
Сначала попробуй самостоятельно решить :)

<details>
<summary>C</summary>

```c
#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

int main() {
  int capacity = 10;
  int *numbers = (int *)malloc(capacity * sizeof(int));
  if (numbers == NULL) {
    printf("n/a");
    return 1;
  }

  int error = 0;
  int num = 0;
  int count = 0;
  int max_i = -1, min_i = -1;
  int max_val = INT_MIN, min_val = INT_MAX;
  while (scanf("%d", &num) && (num != -1) && (error == 0)) {
    if (count >= capacity) {
      capacity *= 2;
      numbers = (int *)realloc(numbers, capacity * sizeof(int));
      if (numbers == NULL) {
        printf("n/a");
        error = 1;
        continue;
      }
    }

    if (num > max_val) {
      max_val = num;
      max_i = count;
    }
    if (num < min_val) {
      min_val = num;
      min_i = count;
    }

    numbers[count++] = num;
  }

  if (error == 0) {
    if (max_i != -1 && min_i != -1) {
      int temp = numbers[max_i];
      numbers[max_i] = numbers[min_i];
      numbers[min_i] = temp;
    }

    for (int i = 0; i < count; i++) {
      if (i != count - 1) printf("%d ", numbers[i]);
      else printf("%d", numbers[i]);
    }
  }

  free(numbers);

  return error;
}
```

</details>
