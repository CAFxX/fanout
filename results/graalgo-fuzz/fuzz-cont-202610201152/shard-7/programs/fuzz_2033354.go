package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint16; F1 string; F2 int32 }
type GenS4 struct { F0 uint8 }
func genFn0() int32 {
    if (false && ((599 <= 18446744073709551615) && true)) {
        var v5 int8 = -1
        _ = v5
    } else {
        var v6 int32 = 2147483647
        fmt.Println(v6)
        _ = v6
    }
    var v7 uint8 = 42
    switch "\n\t" {
    case "a":
        v8 := []int{4294967295}
        _ = v8
    default:
        var v9 int32 = -814
        _ = v9
    }
    var v10 int = 1
    v11 := GenS1{F0: 118, F1: 1}
    _ = v10
    _ = v11
    _ = v7
    return -588
}
func genFn1(p0 bool, p1 uint32, p2 uint8) uint16 {
    var v12 uint8 = (p2 + p2)
    var v13 bool = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 201
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    var v3 uint16 = 65535
    var v4 int8 = 19
    _ = len(v1)
    v4 -= v4
    for fi14 := 0; fi14 < 9; fi14++ {
        v4 = v4
        var v15 uint8 = 42
        var v16 float32 = float32(float64((1.5 * float32(float64(3.14159)))))
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        _ = len(v2)
        _ = fi17
    }
    v4 = (v4 % ((v4) | 1))
    for fi18 := 0; fi18 < 14; fi18++ {
        v19 := []int{fi18, (fi18 | -9223372036854775808), (fi18 >> ((18446744073709551615) & 63)), fi18}
        _ = fi18
        _ = v19
    }
    if (true || true) {
        switch 2147483647 {
        case 128:
            v20 := []int{}
            _ = v20
        default:
            var v21 bool = true
            _ = v21
        }
    }
    fmt.Println(v4, v3)
    v22 := GenS1{F0: (v4 % ((v4) | 1)), F1: -1}
    var v23 uint32 = 895
    v24 := GenS2{F0: v3, F1: 0}
    switch 127 {
    case 2147483647:
        delete(v1, ("" + ("\n\t" + (("foo bar" + "\n\t") + "foo bar"))))
    }
    v25 := GenS3{F0: v3, F1: "\n\t", F2: -275}
    v26 := GenS4{F0: 42}
    var v27 int = -633
    _ = len(v2)
    var v28 int16 = -1
    v4 += v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
