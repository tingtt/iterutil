# iterutil

## Usage

```sh
go get github.com/tingtt/iterutil
```

```go
import "github.com/tingtt/iterutil"
```

- Only
  ```go
  for item := range iterutil.Only(slices.Values([]string{"a", "b", "c"}), "a") {
    fmt.Println(item)
  }
  // will print:
  // a
  ```
- OnlyFunc
- OnlyKey
- OnlyKeyFunc
- OnlyValue
- OnlyValueFunc
- Filter
- FilterFunc
  ```go
  filterLargerThanOrEqual := func(i int) func(v int) bool {
    return func(v int) bool {
      return v > i
    }
  }
  for item := range iterutil.FilterFunc(slices.Values([]int{1, 2, 3, 4, 5}), filterLargerThanOrEqual(3)) {
    fmt.Println(item)
  }
  // will print:
  // 3
  // 4
  // 5
  ```
- FilterKey
- FilterKeyFunc
- FilterValue
- FilterValueFunc
