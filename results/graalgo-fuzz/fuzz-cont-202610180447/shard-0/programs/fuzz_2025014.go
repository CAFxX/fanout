package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint32 }
type GenS2 struct { F0 int }

func main() {
    var v1 int64 = 1
    var v2 uint16 = 0
    v3 := []uint32{42, 65535}
    v4 := make(map[int]string)
    v5 := GenS1{F0: -489, F1: 0, F2: 129}
    v6 := GenS1{F0: 159, F1: -128, F2: 42}
    for fi7 := 0; fi7 < 14; fi7++ {
        v2 -= 0
        fmt.Println(v2, v1, v1)
        _ = fi7
    }
    var v8 uint32 = 4294967295
    v9 := GenS1{F0: -931, F1: 0, F2: 452}
    v2 += v2
    switch 137 {
    case 0:
        var v10 uint8 = 2
        _ = v10
    case 1:
        switch "" {
        case "foo bar":
            v11 := GenS2{F0: -1}
            _ = v11
        case "":
            if (false || (false && ((0.0 * float64(1e10)) == 1e10))) {
                v12 := GenS1{F0: 1, F1: 100, F2: (^(v8))}
                v2 -= v2
                _ = v12
            }
        }
    case 9223372036854775807:
        var v13 uint8 = 42
        _ = v13
    }
    _ = len(v4)
    if (true && (!(!true))) {
        if (!true) {
            var v14 bool = true
            for fi15 := 0; fi15 < 19; fi15++ {
                for fi16 := 0; fi16 < 19; fi16++ {
                    fmt.Println(v14, v1)
                    fmt.Println(v1, v1)
                    fmt.Println(v14)
                    _ = fi16
                }
                fmt.Println(fi15, v2, v8)
                v8 = 104
                _ = fi15
            }
            fmt.Println(v8, v14)
            _ = v14
        }
        switch int(v1) {
        case -592:
            var v17 uint = 4294967295
            _ = v17
        }
        v18 := GenS2{F0: -2147483648}
        _ = v18
    }
    var v19 bool = true
    fmt.Println(v1, v1)
    fmt.Println(v2, v19, v1)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
