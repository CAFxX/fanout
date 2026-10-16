package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int8 }
type GenS2 struct { F0 string; F1 uint; F2 uint32 }

func main() {
    var v1 int8 = 19
    var v2 uint32 = 2
    v3 := make(map[int]bool)
    var v4 float32 = 3.14159
    v5 := make(map[string]string)
    var v6 uint8 = 161
    var v7 float32 = v4
    switch -885 {
    case 599:
        var v8 int64 = 0
        _ = v8
    }
    v2 += v2
    var v9 uint16 = 306
    for fi10 := 0; fi10 < 11; fi10++ {
        if true {
            v11 := v3[9223372036854775807]
            _ = v11
        } else {
            if (!true) {
                v9 += 42
                v4 = v4
            }
            v12 := GenS1{F0: 42, F1: (!true), F2: (v1 + 127)}
            _ = v12
        }
        v7 = 1e10
        _ = fi10
    }
    fmt.Println(v9, v9, v6)
    v13 := []int{}
    for fi14 := 0; fi14 < 2; fi14++ {
        v9 = v9
        _ = fi14
    }
    v9 = 42
    v13 = append(v13, -677)
    v15 := GenS1{F0: 18446744073709551615, F1: false, F2: (v1 * v1)}
    var v16 uint32 = (^(v2))
    v1 = 42
    var v17 int = -2147483648
    var v18 string = "a"
    fmt.Println(v16, v17, v7)
    switch (v17 << ((4294967295) & 63)) {
    case 137:
        v9 += v9
    case 127:
        v19 := GenS2{F0: "\n\t", F1: 825, F2: 1}
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
