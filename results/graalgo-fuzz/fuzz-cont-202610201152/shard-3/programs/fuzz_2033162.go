package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int16; F2 uint64 }
func genFn0(p0 int16, p1 uint16, p2 int64) bool {
    v6 := []int{2147483647, 32767, -9223372036854775808}
    if len(v6) > 0 {
        v6[0] = 32767
    }
    switch "a" {
    case "a":
        if true {
            if len(v6) > 0 {
                v6[0] = 9223372036854775807
            }
            v6 = append(v6, 4294967295)
        }
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        v6 = append(v6, -646)
        _ = fi7
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return false
}
func genFn1() int64 {
    var v8 uint32 = 65535
    var v9 uint16 = 255
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS1{F0: 65535}
        var v12 uint32 = v8
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 uint8 = 159
    var v2 float32 = 3.14159
    var v3 uint64 = 4294967295
    v4 := []int{127, -9223372036854775808, -9223372036854775808}
    v5 := make(map[uint32]bool)
    if false {
        v13 := GenS1{F0: 255}
        for fi14 := 0; fi14 < 18; fi14++ {
            v1 += (v1 * (v1 % ((v1) | 1)))
            var v15 uint64 = v3
            v2 += v2
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        var v16 int32 = -2147483648
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v16
    }
    switch "\n\t" {
    case "":
        fmt.Println(v3, v2, v2)
    default:
        v3 = 97
    }
    if true {
        v2 += float32(float64(v2))
        v1 -= v1
        v2 += v2
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := v5[2147483647]
        _ = fi17
        _ = v18
    }
    v2 = 1.5
    if ((-2.25 > 0.0) || (0 <= genFn1())) {
        var v19 int64 = 564
        var v20 uint = 4294967295
        _ = v19
        _ = v20
    }
    fmt.Println(v1, v1, v1)
    v21 := GenS1{F0: 1}
    v22 := GenS2{F0: (!(0 <= -1)), F1: -32768, F2: v3}
    var v23 float32 = (3.14159 - (v2 * ((3.14159 * 1.5) / v2)))
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
