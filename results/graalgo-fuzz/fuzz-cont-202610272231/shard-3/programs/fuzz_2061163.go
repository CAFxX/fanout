package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int64 }
type GenS2 struct { F0 uint32; F1 int16; F2 int32 }
type GenS3 struct { F0 bool }

func main() {
    var v1 bool = (!(!true))
    var v2 int64 = 0
    var v3 uint16 = 337
    v4 := []uint32{2147483648, 776}
    v5 := make(map[string]int)
    var v6 string = "\n\t"
    v7 := GenS1{F0: 88, F1: (v2 % ((v2) | 1)), F2: v2}
    var v8 int64 = (-(v2))
    v2 += (v2 * -140)
    var v9 uint16 = v3
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(v3)
        _ = fi10
    }
    switch 175 {
    case 4294967295:
        var v11 bool = (3.14159 < 0.0)
        _ = v11
    default:
        switch 32767 {
        case 9223372036854775807:
            if ((-(v2)) != v2) {
                var v12 uint32 = 627
                _ = v12
            }
        case 0:
            if (!(898 <= 2147483648)) {
                var v13 float64 = float64((3.14159 + float32((0.0 / 1.0))))
                if (18446744073709551615 != 2) {
                    for fi14 := 0; fi14 < 9; fi14++ {
                        v15 := GenS1{F0: 2, F1: v8, F2: v2}
                        var v16 uint64 = 0
                        var v17 uint32 = 42
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    v18 := GenS2{F0: 751, F1: -1, F2: 42}
                    _ = v18
                }
                fmt.Println(v2)
                _ = v13
            } else {
                fmt.Println(v1, v3, v3)
            }
        }
    }
    var v19 int8 = 127
    if v1 {
        v9 += v3
        v19 -= 34
    } else {
        var v20 uint16 = v9
        switch int(v19) {
        case -1:
            v21 := GenS3{F0: (v8 >= 42)}
            _ = v21
        case 32767:
            var v22 bool = (127 >= 128)
            _ = v22
        }
        _ = v20
    }
    var v23 int64 = (^(v2))
    v6 = (v6 + ("hello" + v6))
    var v24 int8 = (v19 ^ (-(v19)))
    v8 = (v2 - v2)
    if v1 {
        v8 += (-(v23))
        var v25 string = "foo bar"
        fmt.Println(v23, v8)
        _ = v25
    } else {
        v4 = append(v4, 2147483648)
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        for fi27 := 0; fi27 < 8; fi27++ {
            fmt.Println(v9, v3, v8)
            _ = fi27
        }
        for fi28 := 0; fi28 < 15; fi28++ {
            fmt.Println(fi28, v1, v23)
            _ = fi28
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
