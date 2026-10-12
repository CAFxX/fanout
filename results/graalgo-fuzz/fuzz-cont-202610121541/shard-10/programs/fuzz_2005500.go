package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int = 22
    var v2 int64 = 229
    var v3 string = "\n\t"
    v4 := GenS1{F0: -9223372036854775808}
    var v5 float64 = 3.14159
    if ((^(v1)) != -508) {
        if false {
            v5 -= -2.25
        }
        var v6 uint16 = 42
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS2{F0: (v6 & v6)}
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    var v9 int32 = 42
    if (0 == 111) {
        fmt.Println(v1)
        v3 = ((v3 + "hello") + v3)
    } else {
        var v10 int64 = v2
        v3 = v3
        _ = v10
    }
    v11 := []int{v1, (^(v1)), -2147483648, 2147483647}
    _ = len(v11)
    v12 := GenS2{F0: 1}
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
