package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int8 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 int16; F1 uint }
func genFn0(p0 int8) uint {
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 605, F1: 2147483647, F2: 53}
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    return 2
}
func genFn1() uint {
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := GenS2{F0: 2, F1: -1}
        _ = fi9
        _ = v10
    }
    var v11 uint8 = 42
    v11 = 127
    var v12 uint16 = 1
    _ = v11
    _ = v12
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    var v3 float64 = 3.14159
    v4 := make(map[int]bool)
    var v5 int = 2147483647
    var v6 int8 = 42
    switch 810 {
    case -9223372036854775808:
        v13 := GenS3{F0: 42, F1: 255, F2: 102}
        _ = v13
    case -833:
        v5 = (-(v5))
    default:
        if ((100 != v6) && (0 > 0)) {
            v14 := GenS4{F0: (1 < -1), F1: -736}
            v15 := GenS5{F0: 42, F1: 32767}
            _ = v14
            _ = v15
        } else {
            fmt.Println(v5, v5, v6)
            fmt.Println(v3, v5)
            delete(v1, -1)
        }
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        v5 = fi16
        _ = fi16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v6 = (v6 % (((-(v6))) | 1))
        _ = fi17
    }
    v18 := GenS6{F0: 255, F1: 296}
    v19 := GenS6{F0: 18446744073709551615, F1: 4294967295}
    fmt.Println(v5)
    var v20 bool = (237 != 1)
    _ = len(v1)
    switch v5 {
    case 2147483647:
        v20 = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        fmt.Println(v5, v6)
        _ = fi21
    }
    v22 := GenS7{F0: 308, F1: 226}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
