package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 int = 128
    var v5 float32 = ((-2.25 * (float32(((-2.25 * 0.0) / 1.5)) - (1.5 * (0.1 * (1e10 * 0.1))))) - 0.1)
    v6 := []string{"", "\n\t", "hello"}
    fmt.Println(v5, v4)
    delete(v2, v4)
    v5 = v5
    _ = len(v6)
    v7 := GenS1{F0: 1, F1: 400, F2: 437}
    v8 := GenS1{F0: 155, F1: 2, F2: 2}
    if ((v4 ^ (v4 >> ((0) & 63))) == 42) {
        v9 := v2[248]
        _ = v9
    } else {
        switch -2147483648 {
        case -465:
            if ((!true) && true) {
                var v10 string = ""
                if true {
                    v5 += float32((float64(v5) * float64(v5)))
                    var v11 uint64 = 2
                    _ = v11
                }
                _ = v10
            }
        }
        var v12 float64 = 3.14159
        v13 := GenS2{F0: 9223372036854775807}
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 255, F1: 923, F2: 255}
    if (false || (1 > 290)) {
        if true {
            _ = len(v6)
        } else {
            var v15 int16 = 1000
            v16 := GenS1{F0: 130, F1: 1, F2: 712}
            _ = v15
            _ = v16
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
