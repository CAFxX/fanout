package main

import (
    "fmt"
)

func genFn0(p0 uint32) uint16 {
    for fi5 := 0; fi5 < 16; fi5++ {
        p0 = 622
        var v6 float32 = ((((1.5 - 3.14159) * -2.25) + (3.14159 * 0.0)) / (0.1 - (0.0 / 1.0)))
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(p0)
        _ = fi7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        switch "\n\t" {
        case "foo bar":
            v9 := []int{fi8}
            _ = v9
        default:
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 int8 = 42
                var v12 uint = 153
                var v13 int16 = 42
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
        }
        p0 = p0
        var v14 int32 = -1
        _ = fi8
        _ = v14
    }
    _ = p0
    return 823
}
func genFn1(p0 int8) int8 {
    p0 -= 1
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 float32 = 3.14159
        v17 := []int{fi15}
        v16 -= (-2.25 * -2.25)
        _ = fi15
        _ = v16
        _ = v17
    }
    p0 = -91
    _ = p0
    return p0
}

func main() {
    var v1 uint16 = 42
    var v2 float64 = -2.25
    var v3 bool = (2 > 51)
    v4 := []uint32{2147483647, 4294967295, 288, 291}
    var v18 bool = v3
    _ = len(v4)
    _ = len(v4)
    v1 -= (v1 % ((32767) | 1))
    v1 -= genFn0(2147483648)
    var v19 uint32 = 2147483648
    for fi20 := 0; fi20 < 17; fi20++ {
        v4 = append(v4, (v19 & v19))
        if len(v4) > 0 {
            v4[0] = v19
        }
        _ = fi20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        v1 -= v1
        v19 += v19
        if v18 {
            v1 = v1
        }
        _ = fi21
    }
    if v3 {
        var v22 bool = v18
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
