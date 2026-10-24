package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int8 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 bool; F1 int8; F2 int32 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int32, p1 uint16) bool {
    switch -472 {
    case 543:
        fmt.Println(p1, p1)
    case 0:
        var v5 int32 = p0
        _ = v5
    }
    p1 += 65535
    v6 := GenS1{F0: 34, F1: 1, F2: -128}
    _ = p0
    _ = p1
    _ = v6
    return ((-2.25 <= (((3.14159 - 1e10) - float32(float64((0.0 * (3.14159 + 0.1))))) / 1.0)) && (!true))
}
func genFn1(p0 uint16) uint {
    if (20 < uint8(p0)) {
        switch 684 {
        case 42:
            p0 = (p0 * p0)
        default:
            v7 := []int{128, 0, -2147483648}
            _ = v7
        }
        v8 := GenS2{F0: 226, F1: -646}
        _ = v8
    } else {
        var v9 string = "\n\t"
        var v10 float32 = 1.5
        _ = v10
        _ = v9
    }
    if true {
        v11 := GenS1{F0: 1, F1: -1, F2: 1}
        p0 += (^(p0))
        _ = v11
    }
    p0 = (p0 << ((2) & 63))
    var v12 float32 = ((3.14159 + 1e10) / 0.1)
    p0 -= 42
    _ = p0
    _ = v12
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 255
    v2 := []uint32{65535, 258, 4294967295}
    var v3 int16 = -32768
    v4 := make(map[string]uint32)
    v13 := GenS3{F0: true, F1: 42, F2: -1}
    v1 -= v1
    v3 = (v3 * v3)
    v14 := GenS4{F0: 1, F1: v1, F2: 65535}
    v15 := GenS4{F0: -2147483648, F1: 232, F2: 2}
    var v16 int16 = 1
    v17 := GenS4{F0: -1, F1: v1, F2: 501}
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 int64 = 42
        for fi20 := 0; fi20 < 10; fi20++ {
            var v21 int64 = (v19 * v19)
            _ = fi20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    v22 := GenS1{F0: 107, F1: (v3 | (-(v3))), F2: 0}
    switch -452 {
    case 42:
        v23 := GenS4{F0: -1, F1: v1, F2: 2}
        _ = v23
    case 442:
        switch 1 {
        case 696:
            for fi24 := 0; fi24 < 16; fi24++ {
                v3 = (v3 + v3)
                _ = fi24
            }
        }
    default:
        var v25 uint = 0
        _ = v25
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        v1 = v1
        switch "" {
        case "":
            v27 := GenS2{F0: 1, F1: -9223372036854775808}
            _ = v27
        case "hello":
            v28 := GenS3{F0: (290 < 4294967295), F1: 127, F2: -31}
            _ = v28
        case "\n\t":
            if false {
                v29 := GenS1{F0: 2, F1: (v16 - v3), F2: 1}
                _ = v29
            } else {
                v1 -= v1
                for fi30 := 0; fi30 < 19; fi30++ {
                    v31 := v4[("hello" + ("" + ("" + "a")))]
                    var v32 int8 = 42
                    _ = len(v4)
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
        }
        _ = fi26
    }
    v33 := GenS1{F0: 0, F1: (v3 << ((756) & 63)), F2: 100}
    v34 := GenS3{F0: false, F1: -128, F2: -427}
    fmt.Println(v3, v16, v3)
    v16 = -1
    v35 := GenS1{F0: 4294967295, F1: v3, F2: -55}
    for fi36 := 0; fi36 < 15; fi36++ {
        v16 = v16
        v37 := GenS5{F0: 42}
        _ = fi36
        _ = v37
    }
    var v38 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v4
}
