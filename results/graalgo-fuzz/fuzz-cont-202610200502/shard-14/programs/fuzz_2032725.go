package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int32 }

func main() {
    var v1 int8 = 3
    var v2 uint32 = 679
    var v3 int16 = 1000
    v4 := GenS1{F0: 0, F1: v1, F2: 1}
    var v5 int8 = (^(v1))
    v6 := []int{-415, -9223372036854775808, 378}
    v6 = append(v6, 42)
    v7 := GenS1{F0: 1, F1: (v5 % ((100) | 1)), F2: -2147483648}
    var v8 int8 = (^(v1))
    var v9 string = "foo bar"
    v10 := GenS1{F0: 255, F1: v5, F2: 0}
    v1 = (v5 / ((v5) | 1))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v8 = v5
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS1{F0: 128, F1: 99, F2: -1}
        var v13 uint = 1
        v6 = append(v6, (^(fi11)))
        _ = fi11
        _ = v12
        _ = v13
    }
    if true {
        for fi14 := 0; fi14 < 7; fi14++ {
            _ = len(v6)
            _ = fi14
        }
    } else {
        if (9223372036854775807 <= 32768) {
            for fi15 := 0; fi15 < 13; fi15++ {
                v3 = -1
                _ = fi15
            }
            fmt.Println(v8, v2, v5)
            var v16 int = 42
            _ = v16
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
