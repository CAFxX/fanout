package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 uint64; F2 int8 }

func main() {
    var v1 uint32 = 848
    v2 := make(map[string]bool)
    v3 := []int{32768, 290, 1}
    var v4 uint = 2
    var v5 uint16 = 2
    if len(v3) > 0 {
        v3[0] = 9223372036854775807
    }
    if (((-2.25 == (((1e10 * 0.1) + ((0.0 * (0.0 - float64((float32(-2.25) - 0.0)))) / (((0.1 * 1.5) - ((0.1 / 1.0) / float64((-2.25 / 1.5)))) / (float64((0.1 + -2.25)) - -2.25)))) / 0.1)) && true) || false) {
        v5 = v5
        v4 = (v4 >> ((v4) & 63))
        var v6 int64 = 9223372036854775807
        _ = v6
    } else {
        fmt.Println(v5, v1, v5)
        _ = len(v3)
        v7 := GenS1{F0: 377}
        _ = v7
    }
    var v8 uint = (v4 & (v4 + v4))
    var v9 bool = (42 >= -351)
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS2{F0: 128}
        var v12 bool = (3.14159 == 1.5)
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v9, v5)
    var v13 bool = (v9 && (false || (!((1.5 * ((0.1 / 3.14159) - (1e10 + (0.1 + 1.5)))) == 1e10))))
    if true {
        switch 4294967295 {
        case -694:
            var v14 float32 = (1.5 / -2.25)
            _ = v14
        default:
            _ = len(v3)
        }
    } else {
        v4 -= 285
    }
    switch 398 {
    case 9223372036854775807:
        v9 = (!false)
    case 42:
        v15 := GenS3{F0: v5, F1: (v5 << (((v8 << ((255) & 63))) & 63)), F2: -463}
        _ = v15
    default:
        delete(v2, "\n\t")
    }
    var v16 int = int(v5)
    v17 := GenS4{F0: -653}
    v2[("\n\t" + (("hello" + "hello") + "hello"))] = true
    switch 128 {
    case 127:
        var v18 int64 = 0
        _ = v18
    case -162:
        for fi19 := 0; fi19 < 9; fi19++ {
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            v13 = true
            _ = fi19
        }
    }
    if ((!v9) || (9223372036854775807 != -188)) {
        var v20 uint32 = (v1 & (v1 ^ (v1 & 485)))
        _ = v20
    }
    fmt.Println(v16, v16)
    for fi21 := 0; fi21 < 20; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            _ = len(v3)
            _ = fi22
        }
        var v23 uint64 = 42
        v24 := GenS5{F0: v1, F1: uint64(v1), F2: 1}
        _ = fi21
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        var v26 int = 9223372036854775807
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
