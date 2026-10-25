package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint32 }
func genFn0() uint32 {
    var v5 uint8 = 199
    var v6 uint32 = 146
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int8 = int8(v5)
        var v9 int8 = v8
        v10 := []int{}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    _ = v5
    _ = v6
    return (v6 % (((v6 * v6)) | 1))
}

func main() {
    var v1 int8 = -10
    var v2 int64 = -447
    var v3 int = 1
    v4 := make(map[int]uint32)
    fmt.Println(v3, v2, v2)
    switch 32767 {
    case -220:
        var v11 float64 = 1.5
        _ = v11
    }
    if (v1 < (^(v1))) {
        v1 = v1
        v4[88] = 2
    } else {
        switch int(v1) {
        case 9223372036854775807:
            var v12 int16 = -1
            _ = v12
        case 42:
            fmt.Println(v3)
        case 4294967295:
            v3 += v3
        default:
            for fi13 := 0; fi13 < 8; fi13++ {
                var v14 float64 = 3.14159
                var v15 bool = ((!true) || true)
                v16 := []int{(fi13 | v3)}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
        }
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        switch "foo bar" {
        case "foo bar":
            v2 -= v2
        case "a":
            v18 := GenS1{F0: -1, F1: true, F2: 4294967295}
            _ = v18
        case "\n\t":
            if (true && (146 <= 1)) {
                var v19 uint32 = 3
                fmt.Println(v19, fi17)
                _ = v19
            } else {
                v2 += (v2 * (v2 & int64(v3)))
                v20 := GenS2{F0: 32767, F1: 42}
                _ = v20
            }
        default:
            fmt.Println(v1)
        }
        if true {
            v1 += int8(fi17)
        } else {
            v3 = v3
            v3 = v3
        }
        fmt.Println(fi17, v2, v1)
        _ = fi17
    }
    switch "" {
    case "a":
        var v21 uint = 2
        _ = v21
    case "foo bar":
        var v22 bool = true
        _ = v22
    case "\n\t":
        delete(v4, v3)
    default:
        for fi23 := 0; fi23 < 20; fi23++ {
            v2 = 0
            v24 := GenS2{F0: 0, F1: 202}
            v1 = -126
            _ = fi23
            _ = v24
        }
    }
    var v25 int16 = 42
    for fi26 := 0; fi26 < 1; fi26++ {
        fmt.Println(v1)
        for fi27 := 0; fi27 < 15; fi27++ {
            v2 = 21
            fmt.Println(v1)
            var v28 int8 = (v1 >> ((2) & 63))
            _ = fi27
            _ = v28
        }
        fmt.Println(v2, v3)
        _ = fi26
    }
    var v29 bool = (1 < 468)
    v4[(v3 % ((-492) | 1))] = 42
    var v30 int8 = 21
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
