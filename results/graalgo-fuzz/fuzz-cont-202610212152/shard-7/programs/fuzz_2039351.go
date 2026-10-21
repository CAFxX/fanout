package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint }
func genFn0(p0 uint64, p1 int16, p2 uint) uint {
    var v5 uint32 = 0
    v6 := GenS1{F0: false, F1: true, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p2
}
func genFn1(p0 uint16, p1 uint16) int {
    switch 0 {
    case 42:
        var v7 float32 = (3.14159 - 1e10)
        _ = v7
    case -982:
        var v8 int = 127
        _ = v8
    }
    if true {
        v9 := GenS1{F0: (-313 < 42), F1: (1 < 1), F2: 65535}
        var v10 uint8 = 129
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 float64 = (1e10 / 1e10)
            _ = fi11
            _ = v12
        }
        _ = v10
        _ = v9
    }
    var v13 bool = (!true)
    _ = p0
    _ = p1
    _ = v13
    return 748
}

func main() {
    v1 := []uint32{2147483648, 65535, 1, 2147483648}
    var v2 uint8 = 128
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]int)
    v2 = 255
    v2 = 82
    if true {
        if true {
            if true {
                v2 += v2
                for fi14 := 0; fi14 < 20; fi14++ {
                    var v15 uint16 = 32767
                    v2 -= v2
                    _ = fi14
                    _ = v15
                }
            } else {
                v2 = (v2 * (^(v2)))
            }
            switch "" {
            case "":
                fmt.Println(v2)
            }
            var v16 float32 = ((1e10 + float32(-2.25)) * 1e10)
            _ = v16
        }
        _ = len(v1)
        fmt.Println(v2)
    }
    var v17 uint = 65535
    for fi18 := 0; fi18 < 19; fi18++ {
        if (-324 != -1) {
            v2 += v2
            v17 = v17
        }
        _ = fi18
    }
    var v19 float64 = -2.25
    v3[533] = genFn1(1, 1)
    if false {
        var v20 uint = 1
        var v21 bool = false
        _ = v20
        _ = v21
    } else {
        var v22 int8 = 25
        _ = v22
    }
    v2 += (v2 << ((65535) & 63))
    var v23 float64 = (v19 * (((0.0 - float64(((0.1 + ((1.5 / 1.0) / ((3.14159 + 3.14159) - float32((v19 / ((v19 + float64(0.0)) + v19)))))) * 0.0))) + (v19 - ((1e10 + ((((1e10 - 0.1) - v19) + v19) + (float64(3.14159) * v19))) * v19))) + v19))
    fmt.Println(v19, v17, v2)
    var v24 int64 = -9223372036854775808
    var v25 string = "a"
    var v26 int16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
