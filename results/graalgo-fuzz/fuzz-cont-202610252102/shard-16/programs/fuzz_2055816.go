package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 string, p1 uint16, p2 uint32) int32 {
    var v5 int = 42
    p1 = (p1 << ((171) & 63))
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -1
}

func main() {
    var v1 int32 = -2147483648
    var v2 int64 = 42
    var v3 uint64 = 381
    var v4 bool = true
    v4 = true
    v6 := []int{}
    var v7 float32 = (0.1 * ((3.14159 + 3.14159) / float32(((float64(0.0) - ((1e10 - 0.0) + (1e10 * (3.14159 + 0.1)))) * (-2.25 / (float64(((0.1 / 1.0) * (1e10 + float32(((((((-2.25 / ((((3.14159 / (-2.25 + (1.5 / 3.14159))) * 0.1) * float64(((-2.25 + -2.25) / float32(float64(0.1))))) + (0.0 - (0.0 - 0.0)))) + (-2.25 / (float64(0.0) - ((0.1 / 1.0) / 3.14159)))) * float64(0.1)) / 1e10) / 0.1) / 3.14159))))) * (((float64(0.0) / 1.0) + 0.1) - 1e10)))))))
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS1{F0: 18446744073709551615}
            var v11 uint = 255
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 uint = 0
        switch (fi8 - fi8) {
        case 9223372036854775807:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        _ = fi8
        _ = v12
    }
    var v13 bool = ((v1 + (v1 / ((v1) | 1))) < (v1 * (v1 + (v1 + 242))))
    switch 457 {
    case -1:
        var v14 int64 = v2
        _ = v14
    case 681:
        if (v13 || false) {
            _ = len(v6)
            v3 -= (v3 % ((v3) | 1))
            v3 -= 18446744073709551615
        } else {
            for fi15 := 0; fi15 < 16; fi15++ {
                v16 := GenS1{F0: 42}
                v17 := GenS1{F0: 18446744073709551615}
                v6 = append(v6, 2147483647)
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    default:
        v6 = append(v6, 32767)
    }
    var v18 uint = 216
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 uint = 2
        switch "foo bar" {
        case "foo bar":
            v3 = uint64(v18)
        }
        _ = fi19
        _ = v20
    }
    if v13 {
        v7 += 0.1
        for fi21 := 0; fi21 < 11; fi21++ {
            var v22 int16 = 42
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 12; fi23++ {
            if len(v6) > 0 {
                v6[0] = 127
            }
            v24 := GenS1{F0: v18}
            var v25 uint32 = 42
            _ = fi23
            _ = v24
            _ = v25
        }
    } else {
        fmt.Println(v2)
        v13 = ((v3 >= (v3 ^ 2)) || (598 > 0))
        var v26 int16 = 144
        _ = v26
    }
    var v27 int32 = 2147483647
    v28 := GenS1{F0: v18}
    for fi29 := 0; fi29 < 17; fi29++ {
        v1 += 0
        v13 = v4
        v30 := GenS1{F0: v18}
        _ = fi29
        _ = v30
    }
    fmt.Println(v18, v3, v18)
    if (v4 || false) {
        var v31 uint64 = (v3 - 1)
        for fi32 := 0; fi32 < 20; fi32++ {
            var v33 int8 = 127
            _ = fi32
            _ = v33
        }
        _ = v31
    }
    v3 += (v3 - 42)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
