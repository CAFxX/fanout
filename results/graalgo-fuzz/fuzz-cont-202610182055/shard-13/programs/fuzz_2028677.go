package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int, p1 uint32) uint {
    if true {
        p0 = p0
    }
    v5 := GenS1{F0: 65535, F1: -628, F2: 65535}
    var v6 uint16 = 32767
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 96
}
func genFn1(p0 uint16, p1 int64, p2 uint8) int8 {
    fmt.Println(p1, p1)
    var v7 string = ""
    var v8 float32 = 1e10
    v9 := []int{32768, 42, 38}
    if len(v9) > 0 {
        v9[0] = 2147483647
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 float64 = (((-2.25 + float64(-2.25)) - -2.25) / ((1.5 - (3.14159 * ((1e10 * float64(1e10)) / 1.0))) + 3.14159))
    var v2 string = "\n\t"
    var v3 int8 = -51
    v4 := []int{}
    var v10 int64 = 42
    var v11 uint = 42
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v12 int8 = (^(v3))
    if (!((true || false) && (false || true))) {
        switch 4294967295 {
        case 553:
            v2 = ("hello" + (v2 + v2))
        }
    } else {
        var v13 int8 = v3
        _ = v13
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v3 = 127
        var v15 int8 = (v12 / ((v3) | 1))
        switch (v2 + v2) {
        case "a":
            v2 = v2
        case "hello":
            v2 = v2
        }
        _ = fi14
        _ = v15
    }
    var v16 bool = (!true)
    for fi17 := 0; fi17 < 19; fi17++ {
        switch fi17 {
        case 642:
            var v18 int = fi17
            _ = v18
        default:
            v19 := GenS1{F0: uint16(v11), F1: -2147483648, F2: 65535}
            _ = v19
        }
        for fi20 := 0; fi20 < 10; fi20++ {
            if len(v4) > 0 {
                v4[0] = (^(fi20))
            }
            _ = fi20
        }
        _ = fi17
    }
    if v16 {
        if v16 {
            v4 = append(v4, 633)
            var v21 float32 = 1e10
            _ = v21
        }
        fmt.Println(v3, v11, v12)
    } else {
        v16 = (!true)
        v22 := GenS1{F0: 42, F1: 0, F2: 42}
        _ = v22
    }
    if false {
        fmt.Println(v1, v1, v11)
    } else {
        if v16 {
            var v23 int64 = 1
            var v24 float64 = (float64((0.0 - 0.0)) / float64((1e10 / 1.0)))
            v25 := GenS2{F0: 2147483648}
            _ = v23
            _ = v24
            _ = v25
        } else {
            fmt.Println(v16, v1)
        }
    }
    v4 = append(v4, int(v10))
    v26 := GenS1{F0: 1, F1: 2147483647, F2: 65535}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
