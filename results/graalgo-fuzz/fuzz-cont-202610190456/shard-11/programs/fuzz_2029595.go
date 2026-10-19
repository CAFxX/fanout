package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 string }

func main() {
    var v1 float32 = -2.25
    var v2 int64 = 42
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    v5 := make(map[string]int)
    var v6 uint8 = 105
    var v7 float32 = 0.1
    fmt.Println(v2, v1)
    switch 1 {
    case 0:
        var v8 int16 = -629
        _ = v8
    case 82:
        v9 := GenS1{F0: (true && false), F1: -2147483648}
        _ = v9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(fi10, v1)
        v1 = ((v1 - 1.5) + 3.14159)
        _ = fi10
    }
    var v11 uint = 65535
    fmt.Println(v7)
    if ((false || (v11 == (^(v11)))) && (v6 >= v6)) {
        delete(v5, "")
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        if ((true && true) || ((65535 < uint16(v11)) || (v7 == float32((float64(3.14159) + (1.5 + (1.5 + 1.5))))))) {
            fmt.Println(v1, v7)
            switch (fi12 | (fi12 % ((fi12) | 1))) {
            case 1:
                v13 := GenS2{F0: "\n\t"}
                _ = v13
            }
        }
        var v14 float64 = (((float64(v1) + (float64(1e10) * 1e10)) + float64((v1 * -2.25))) * 1.5)
        var v15 int16 = 1000
        _ = fi12
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: "hello"}
    var v17 uint32 = 2147483648
    v11 = v11
    v2 -= (v2 - -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
