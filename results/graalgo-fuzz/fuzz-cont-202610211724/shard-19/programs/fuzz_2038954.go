package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int16 }
func genFn0(p0 int8, p1 bool, p2 uint16) string {
    var v6 uint8 = 42
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{(fi7 & -309), 9223372036854775807, (fi7 / ((2147483647) | 1))}
        _ = fi7
        _ = v8
    }
    var v9 int16 = 1
    var v10 uint32 = 2
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v6
    _ = v9
    return "a"
}

func main() {
    v1 := []uint32{4294967295, 2, 4294967295, 42}
    v2 := make(map[uint32]bool)
    v3 := []int{}
    var v4 int8 = -128
    var v5 uint32 = 65535
    v5 -= 0
    if false {
        if (0 == -229) {
            var v12 uint = 720
            _ = v12
        }
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(v4, v5)
            v5 -= v5
        }
    }
    if len(v3) > 0 {
        v3[0] = 535
    }
    _ = len(v3)
    v13 := v2[(v5 & v5)]
    switch ("hello" + "foo bar") {
    case "hello":
        if (!(!(-681 < 132))) {
            v14 := GenS1{F0: (-2147483648 > 147), F1: 42, F2: 32767}
            _ = v14
        } else {
            for fi15 := 0; fi15 < 11; fi15++ {
                if ((0.1 >= ((((float64((float32(3.14159) / 1.0)) + (float64(3.14159) / float64(3.14159))) / ((1e10 / 1e10) + float64(1e10))) * -2.25) / 1e10)) && ((v13 || v13) || false)) {
                    var v16 int64 = 380
                    _ = v16
                }
                v17 := GenS1{F0: (v5 <= v5), F1: 32767, F2: 1000}
                v4 = v4
                _ = fi15
                _ = v17
            }
            fmt.Println(v13, v5)
            var v18 uint32 = v5
            _ = v18
        }
    case "\n\t":
        if v13 {
            v4 = v4
        }
    case "foo bar":
        var v19 int32 = 42
        _ = v19
    default:
        var v20 int8 = v4
        _ = v20
    }
    v13 = v13
    for fi21 := 0; fi21 < 2; fi21++ {
        var v22 int8 = v4
        _ = fi21
        _ = v22
    }
    if ((false && ((v5 < (v5 | (v5 - 15))) || (v4 <= v4))) && (1e10 != ((1e10 - ((-2.25 * (1.5 * (float64(0.1) * (((0.0 * float64(0.0)) * -2.25) / 1.0)))) * 0.1)) - -2.25))) {
        fmt.Println(v5)
        for fi23 := 0; fi23 < 12; fi23++ {
            v24 := GenS1{F0: v13, F1: 32767, F2: -1}
            v25 := v2[(v5 + (^(v5)))]
            _ = fi23
            _ = v24
            _ = v25
        }
        fmt.Println(v5, v4)
    }
    v4 = v4
    var v26 int8 = (v4 * (^(v4)))
    v26 += (^(v26))
    var v27 int = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
