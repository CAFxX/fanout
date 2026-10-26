package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint32 }
func genFn0(p0 uint16) uint32 {
    v4 := GenS1{F0: "foo bar", F1: 1000, F2: 293}
    p0 -= 32767
    p0 = (p0 * 1)
    _ = p0
    _ = v4
    return 402
}

func main() {
    v1 := []int{32768, 747}
    v2 := []uint32{495, 1, 2147483647}
    var v3 int32 = -2147483648
    fmt.Println(v3)
    var v5 int64 = 1
    if true {
        if true {
            fmt.Println(v5)
            v5 += (v5 % ((0) | 1))
        }
    }
    if (0.0 > -2.25) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v6 int32 = v3
        fmt.Println(v3)
        _ = v6
    } else {
        v7 := GenS1{F0: "foo bar", F1: -840, F2: genFn0(2)}
        _ = v7
    }
    var v8 int = 1
    switch v8 {
    case 32767:
        var v9 string = "foo bar"
        _ = v9
    case 127:
        v3 += -1
    case 4294967295:
        if (-1 == v8) {
            v8 += v8
            _ = len(v1)
        }
    default:
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := GenS1{F0: "a", F1: -601, F2: 2147483648}
            var v12 uint32 = 2147483648
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    v5 += 42
    v1 = append(v1, (v8 << ((0) & 63)))
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            v8 = fi14
            v3 += 324
            v8 -= -9223372036854775808
            _ = fi14
        }
        v5 += -9223372036854775808
        _ = fi13
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 int64 = v5
        _ = len(v2)
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        switch "a" {
        case "a":
            var v18 uint16 = 65535
            _ = v18
        default:
            for fi19 := 0; fi19 < 12; fi19++ {
                if (false && (v3 == (v3 / (((v3 - v3)) | 1)))) {
                    fmt.Println(v3, fi19)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v20 float64 = (-2.25 - ((3.14159 - 1e10) + -2.25))
                    if len(v1) > 0 {
                        v1[0] = fi19
                    }
                    var v21 float32 = 1.5
                    v21 = float32(v20)
                    _ = v20
                    _ = v21
                } else {
                    var v22 float32 = (((float32(3.14159) + (0.1 * float32(3.14159))) + -2.25) + 3.14159)
                    _ = v22
                }
                _ = fi19
            }
        }
        var v23 float32 = 1e10
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi17
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v8
}
