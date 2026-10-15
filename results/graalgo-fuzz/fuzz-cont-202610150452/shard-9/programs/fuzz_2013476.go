package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int16 }
type GenS2 struct { F0 int }
func genFn0(p0 uint16) int32 {
    var v5 int64 = 9223372036854775807
    var v6 uint = 255
    fmt.Println(v6, p0)
    v5 = int64(v6)
    p0 += (p0 + p0)
    _ = p0
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]int)
    var v3 string = "hello"
    var v4 int16 = 1000
    switch -9223372036854775808 {
    case -2147483648:
        if (2 == 2) {
            v7 := GenS1{F0: v4, F1: (v3 + v3), F2: (v4 >> ((2) & 63))}
            _ = v7
        } else {
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 uint = 1
                v3 = v3
                var v10 uint8 = 163
                _ = fi8
                _ = v10
                _ = v9
            }
        }
    case 513:
        v3 = (v3 + v3)
    }
    fmt.Println(v4)
    if false {
        if true {
            fmt.Println(v4, v4)
            var v11 uint32 = 65535
            v11 -= 944
            _ = v11
        } else {
            if true {
                var v12 int16 = (v4 + (^(v4)))
                var v13 float32 = 1.5
                _ = v12
                _ = v13
            }
        }
        fmt.Println(v4)
        var v14 int8 = 127
        _ = v14
    }
    v15 := []int{-1, 32767, -2147483648}
    v16 := GenS2{F0: 128}
    v4 = (v4 | (v4 - (v4 ^ v4)))
    for fi17 := 0; fi17 < 5; fi17++ {
        v4 += 32767
        switch v3 {
        case "a":
            var v18 uint64 = 500
            _ = v18
        case "hello":
            if false {
                v4 = (^(v4))
                fmt.Println(fi17)
            } else {
                var v19 float64 = -2.25
                v20 := v2[2]
                _ = v19
                _ = v20
            }
        default:
            var v21 int = 0
            _ = v21
        }
        _ = fi17
    }
    var v22 uint64 = 742
    var v23 int32 = 42
    var v24 uint = 42
    fmt.Println(v22, v24)
    v3 = v3
    v25 := GenS1{F0: (-(v4)), F1: v3, F2: (^(v4))}
    if true {
        v4 = v4
        var v26 string = ("a" + "")
        if (!((((0.1 - 3.14159) + ((0.0 / (((float32(((((float64(-2.25) + 0.0) * (((0.1 / -2.25) - 1.5) * 1e10)) + (1.5 * (1e10 / 1e10))) * 1.5)) / 1.5) / float32(0.1)) / -2.25)) / 1e10)) <= 1.5) || false)) {
            delete(v1, 1)
        }
        _ = v26
    }
    var v27 bool = false
    for fi28 := 0; fi28 < 18; fi28++ {
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = fi28
    }
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    var v29 int32 = v23
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v29
    _ = v3
    _ = v4
}
