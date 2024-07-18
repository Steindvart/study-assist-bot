## Поворот матрицы на 90 градусов

### Описание
Написать программу осуществляющую переворот матрицы на 90 градусов по часовой стрелке. Размерность матрицы задаётся на стандартный поток ввода в виде двух чисел, после чего следует `M*N` целых чисел описывающих построчно матрицу. Результат работы вывести на стандартный поток вывода в виде таблицы.

#### Дополнительные требования
1. Ввод - `stdin`, вывод - `stdout`.
2. Проверка корректности ввода. В случае некорректности ввода выводить "n/a".
3. В конце вывода не должно быть переноса строки.

### Примеры входных и выходных данных

| Входные данные       | Выходные данные   |
|----------------------|-------------------|
| `2 3`<br>`1 2 3`<br>`4 5 6`    | `4 1`<br>`5 2`<br>`6 3`   |
| `3 3`<br>`1 2 3`<br>`4 5 6`<br>`7 8 9` | `7 4 1`<br>`8 5 2`<br>`9 6 3` |
| `2 2`<br>`1 2`<br>`3 4`       | `3 1`<br>`4 2`    |
| `3 2`<br>`1 2`<br>`3 4`<br>`5 6`    | `5 3 1`<br>`6 4 2` |
| `1 4`<br>`1 2 3 4`       | `1`<br>`2`<br>`3`<br>`4`  |


### Решения
Сначала попробуй самостоятельно решить :)

<details>
<summary>C</summary>

```c
#include <stdio.h>
#include <stdlib.h>

int** alloc_matrix(int M, int N) {
    int **matrix = (int **)malloc(M * sizeof(int *));
    for (int i = 0; i < M; i++) {
        matrix[i] = (int *)malloc(N * sizeof(int));
    }

    return matrix;
}

void free_matrix(int** matrix, int M) {
    for (int i = 0; i < M; i++) {
        free(matrix[i]);
    }
    free(matrix);
}

int main() {
  int M, N;

  if (scanf("%d %d", &M, &N) != 2) {
    printf("n/a");
    return 1;
  }

  int **matrix = alloc_matrix(M, N);
  for (int i = 0; i < M; i++) {
    for (int j = 0; j < N; j++) {
      if (scanf("%d", &matrix[i][j]) != 1) {
        printf("n/a");
        free_matrix(matrix, M);
        return 1;
      }
    }
  }

  int **rotated = alloc_matrix(N, M);
  for (int i = 0; i < M; i++) {
    for (int j = 0; j < N; j++) {
      rotated[j][M - i - 1] = matrix[i][j];
    }
  }

  for (int i = 0; i < N; i++) {
    for (int j = 0; j < M; j++) {
      printf("%d", rotated[i][j]);
      if (j < M - 1) {
        printf(" ");
      }
    }
    if (i < N - 1) {
      printf("\n");
    }
  }

  free_matrix(matrix, M);
  free_matrix(rotated, N);

  return 0;
}
```

</details>
