package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint }
func genFn0(p0 int8) int {
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(fi6)
        var v7 int8 = p0
        _ = fi6
        _ = v7
    }
    var v8 string = ("hello" + "a")
    _ = p0
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint16 = 78
    v2 := []int{128, 4294967295, 32767, -118}
    var v3 string = "\n\t"
    var v4 uint64 = 18446744073709551615
    var v5 float32 = -2.25
    var v9 uint = 4294967295
    if (((-2.25 + -2.25) < 1e10) && ((v1 * 346) <= 139)) {
        fmt.Println(v9, v3, v3)
        v1 -= (^(v1))
        v10 := GenS1{F0: (^(v4))}
        _ = v10
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v2 = append(v2, 421)
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 string = ("a" + "hello")
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: (v4 << ((uint(v4)) & 63))}
    var v14 uint = 255
    var v15 uint64 = (v4 % (((v4 - v4)) | 1))
    var v16 uint8 = 128
    switch 128 {
    case 128:
        v4 = (v15 / ((v15) | 1))
    case 42:
        switch v3 {
        case "foo bar":
            for fi17 := 0; fi17 < 17; fi17++ {
                v3 = v3
                _ = fi17
            }
        case "\n\t":
            var v18 int = -596
            _ = v18
        case "hello":
            if (false || (1 <= 42)) {
                if false {
                    v19 := GenS1{F0: (v15 ^ (^(v4)))}
                    var v20 int16 = -32768
                    for fi21 := 0; fi21 < 14; fi21++ {
                        var v22 uint16 = uint16(v14)
                        _ = fi21
                        _ = v22
                    }
                    _ = v19
                    _ = v20
                }
                var v23 int = 78
                v15 += v15
                _ = v23
            } else {
                fmt.Println(v16, v16)
                v2 = append(v2, int(v15))
                v24 := GenS2{F0: 252, F1: (v14 << (((v14 ^ v14)) & 63))}
                _ = v24
            }
        }
    case 32768:
        var v25 int64 = 0
        _ = v25
    }
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
