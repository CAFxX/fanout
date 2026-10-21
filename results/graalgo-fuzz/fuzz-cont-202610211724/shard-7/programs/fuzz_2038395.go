package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }
type GenS2 struct { F0 int16; F1 uint }
func genFn0(p0 int16, p1 int) uint {
    var v4 string = (((("" + "\n\t") + "a") + "\n\t") + "")
    var v5 int = p1
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: ((fi6 + 92) >= (p1 / ((428) | 1))), F1: 540, F2: (("" + v4) + v4)}
        v8 := []int{}
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(p0, p1)
    p0 = 0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 863
}
func genFn1() uint8 {
    v9 := GenS2{F0: -303, F1: 255}
    var v10 int32 = -356
    v10 = -1
    fmt.Println(v10)
    if ((3.14159 + 1e10) == 0.1) {
        var v11 uint8 = 128
        v10 = (v10 << ((42) & 63))
        _ = v11
    }
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = (float64((((3.14159 - (1.5 - (0.1 / 1.0))) + float32((3.14159 - 1e10))) * float32((1.5 * -2.25)))) - (1e10 * (float64(0.0) / (3.14159 - (1.5 + 1e10)))))
    v3 := make(map[string]bool)
    var v12 uint = 0
    v13 := []int{42}
    v13 = append(v13, -2147483648)
    switch 626 {
    case -836:
        v2 = 0.1
    case 127:
        v14 := v1[0]
        _ = v14
    case 656:
        v2 = float64(0.0)
    default:
        v15 := GenS2{F0: -1, F1: v12}
        _ = v15
    }
    fmt.Println(v12, v12)
    for fi16 := 0; fi16 < 15; fi16++ {
        fmt.Println(fi16, v2, v2)
        for fi17 := 0; fi17 < 14; fi17++ {
            v2 = v2
            _ = fi17
        }
        _ = fi16
    }
    var v18 uint32 = 210
    v12 += (v12 % ((v12) | 1))
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 uint16 = 32767
        var v21 float32 = (0.0 + -2.25)
        var v22 uint16 = v20
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    if false {
        v12 -= 0
        v23 := v3["\n\t"]
        v24 := GenS2{F0: 42, F1: (v12 % (((v12 / ((36) | 1))) | 1))}
        _ = v23
        _ = v24
    }
    v18 = v18
    var v25 bool = true
    var v26 int8 = -128
    var v27 float32 = -2.25
    v26 += v26
    var v28 float32 = 0.1
    if (((("foo bar" + "foo bar") + "hello") + "\n\t") != ("a" + "hello")) {
        var v29 int = -9223372036854775808
        var v30 int8 = v26
        var v31 uint64 = 4294967295
        _ = v29
        _ = v30
        _ = v31
    } else {
        var v32 uint32 = 958
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
