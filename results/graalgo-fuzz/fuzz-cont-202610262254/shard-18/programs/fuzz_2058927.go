package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int64 }
type GenS2 struct { F0 int32; F1 bool; F2 uint16 }
type GenS3 struct { F0 int64; F1 bool }
func genFn0() int16 {
    var v7 uint16 = 255
    var v8 bool = true
    v9 := []int{127, -2147483648, 4294967295}
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 bool = ((42 < 2) && (!(false || (42 >= -9223372036854775808))))
    var v2 float64 = 3.14159
    var v3 string = "\n\t"
    v4 := []int{32768, 42}
    v5 := make(map[int]int)
    v6 := make(map[int]bool)
    switch 232 {
    case 32767:
        var v10 float64 = v2
        _ = v10
    default:
        for fi11 := 0; fi11 < 6; fi11++ {
            v2 += 3.14159
            _ = fi11
        }
    }
    var v12 float64 = v2
    var v13 float64 = ((1e10 * (float64(-2.25) * v12)) * (((0.0 + (((v12 * 3.14159) - (v2 + v2)) / (3.14159 + v2))) * 0.0) * 0.1))
    var v14 float32 = -2.25
    if v1 {
        v3 = v3
        var v15 bool = false
        if v1 {
            v3 = "hello"
            v16 := GenS1{F0: 123, F1: 2, F2: -1}
            _ = v16
        }
        _ = v15
    }
    v12 += float64(v14)
    var v17 int = 0
    delete(v6, (v17 | v17))
    var v18 int16 = 993
    v19 := GenS2{F0: -291, F1: (-1 < 90), F2: 255}
    for fi20 := 0; fi20 < 16; fi20++ {
        v14 += 0.1
        _ = fi20
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        switch v3 {
        case "":
            v22 := GenS1{F0: -128, F1: 1, F2: 220}
            _ = v22
        default:
            var v23 uint32 = 1
            _ = v23
        }
        v17 += fi21
        if len(v4) > 0 {
            v4[0] = v17
        }
        _ = fi21
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        switch v17 {
        case -499:
            switch "hello" {
            case "":
                v13 = v12
            case "\n\t":
                if false {
                    if len(v4) > 0 {
                        v4[0] = (-(v17))
                    }
                    v25 := v5[(^(fi24))]
                    v3 = v3
                    _ = v25
                }
            default:
                _ = len(v5)
            }
        default:
            v26 := GenS3{F0: -731, F1: false}
            _ = v26
        }
        switch v3 {
        case "a":
            var v27 int = (fi24 % (((v17 ^ (v17 & 0))) | 1))
            _ = v27
        }
        _ = fi24
    }
    v28 := GenS2{F0: 658, F1: (!false), F2: 0}
    if (false || true) {
        var v29 bool = v1
        if v29 {
            if (v29 && (!((v14 * v14) > (((0.0 * -2.25) + 1e10) + (v14 - ((v14 * 1.5) * -2.25)))))) {
                if ((v29 && (994 == 1)) || v1) {
                    var v30 bool = (!(((0.1 / v14) / 1.0) < v14))
                    var v31 uint8 = 127
                    _ = v30
                    _ = v31
                }
            }
            v32 := v6[v17]
            v32 = v29
            _ = v32
        } else {
            v4 = append(v4, (-(v17)))
        }
        _ = v29
    } else {
        fmt.Println(v2, v17, v13)
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        v12 -= 1e10
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
