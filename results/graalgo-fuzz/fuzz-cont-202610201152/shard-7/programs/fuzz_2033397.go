package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int }
func genFn0(p0 bool) int {
    var v7 float64 = 1e10
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint32 = 1
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 string = "a"
        if (0.1 > 0.1) {
            v12 := []int{fi10}
            v11 = v11
            for fi13 := 0; fi13 < 16; fi13++ {
                switch (v11 + v11) {
                case "\n\t":
                    _ = len(v12)
                default:
                    fmt.Println(v11, fi10)
                }
                var v14 float64 = v7
                v7 = -2.25
                _ = fi13
                _ = v14
            }
            _ = v12
        } else {
            v11 = v11
            v15 := []int{fi10, 1, (fi10 * (fi10 ^ 4294967295))}
            if len(v15) > 0 {
                v15[0] = fi10
            }
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    p0 = (!false)
    _ = p0
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 0
    var v2 int = -1
    v3 := make(map[int]uint32)
    v4 := []int{}
    var v5 int = v2
    v6 := make(map[string]uint32)
    var v16 int64 = 178
    if len(v4) > 0 {
        v4[0] = int(v16)
    }
    delete(v6, ("foo bar" + ("hello" + "foo bar")))
    v2 = v5
    if (!(((v2 >> ((555) & 63)) >= v2) && (-1 >= 1000))) {
        v17 := GenS1{F0: -128, F1: 255, F2: (v2 >> ((42) & 63))}
        v5 = (-(v5))
        _ = v17
    } else {
        var v18 uint16 = (v1 / (((v1 ^ v1)) | 1))
        v19 := GenS1{F0: 45, F1: 4294967295, F2: (v5 % ((v2) | 1))}
        fmt.Println(v1, v2)
        _ = v18
        _ = v19
    }
    v5 -= (v5 * v5)
    if (false && false) {
        v20 := GenS1{F0: -1, F1: 255, F2: genFn0((3.14159 >= (float32(3.14159) + -2.25)))}
        fmt.Println(v2, v5, v2)
        _ = v20
    }
    v5 = (v5 / ((649) | 1))
    v5 += v5
    v5 = v2
    if false {
        v1 -= v1
    } else {
        v2 = 9223372036854775807
        var v21 uint8 = 42
        _ = v21
    }
    if true {
        var v22 float64 = -2.25
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
