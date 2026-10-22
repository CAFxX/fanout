package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint64 }

func main() {
    var v1 uint = 2
    v2 := []int{357, 42}
    var v3 int64 = 9223372036854775807
    var v4 uint16 = 2
    var v5 int = 9223372036854775807
    if false {
        if (1.5 <= 1e10) {
            var v6 string = "a"
            _ = v6
        }
    } else {
        v3 += (^(v3))
        fmt.Println(v3)
    }
    _ = len(v2)
    var v7 int32 = 42
    if true {
        v2 = append(v2, v5)
        var v8 uint64 = 704
        _ = v8
    }
    var v9 int32 = (v7 ^ 42)
    fmt.Println(v1, v7, v7)
    v3 -= (v3 - (v3 * v3))
    if true {
        fmt.Println(v9, v1, v4)
    }
    fmt.Println(v1, v7)
    var v10 float32 = 1e10
    v7 = -2147483648
    switch v5 {
    case -1:
        v11 := GenS1{F0: (32767 != -1), F1: (v1 >> (((^(v1))) & 63)), F2: 4294967295}
        _ = v11
    case 9223372036854775807:
        var v12 int32 = v7
        _ = v12
    default:
        switch "foo bar" {
        case "a":
            var v13 uint16 = (v4 - v4)
            _ = v13
        default:
            fmt.Println(v1, v7, v9)
        }
    }
    v3 = v3
    v14 := GenS1{F0: true, F1: 664, F2: 2}
    v2 = append(v2, 4294967295)
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 int32 = v7
        switch (v5 << ((v1) & 63)) {
        case -9223372036854775808:
            if (4294967295 <= 65535) {
                v16 += -371
                v17 := GenS1{F0: true, F1: (v1 + v1), F2: 471}
                _ = v17
            }
        case 4294967295:
            fmt.Println(v16, v10)
        default:
            v18 := GenS1{F0: false, F1: (v1 / (((v1 << ((9) & 63))) | 1)), F2: 0}
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    var v19 int8 = 8
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
