package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 string; F2 int8 }
type GenS4 struct { F0 uint64; F1 int16; F2 uint64 }
func genFn0(p0 uint32, p1 int, p2 bool) uint {
    p0 = p0
    v6 := GenS1{F0: (^(p1))}
    var v7 int = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 316
}

func main() {
    var v1 int64 = -1
    var v2 int8 = -128
    v3 := []int{128}
    var v4 uint16 = 65535
    var v5 uint16 = v4
    var v8 uint = 65535
    switch 4294967295 {
    case 1:
        switch 396 {
        case 32767:
            v5 += v4
        default:
            var v9 uint16 = v5
            _ = v9
        }
    case 32768:
        if (float64(0.1) != (-2.25 * float64(1.5))) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    case -153:
        v10 := GenS2{F0: (v1 * v1)}
        _ = v10
    default:
        var v11 uint16 = (v5 >> (((v8 % ((genFn0(4294967295, 128, true)) | 1))) & 63))
        _ = v11
    }
    if ((2 < 2) && false) {
        switch -679 {
        case 979:
            fmt.Println(v8, v5, v1)
        case 127:
            v12 := GenS3{F0: 1, F1: "foo bar", F2: v2}
            _ = v12
        case 32767:
            switch 4294967295 {
            case 1:
                v13 := GenS2{F0: (v1 * v1)}
                _ = v13
            default:
                var v14 int32 = -2147483648
                _ = v14
            }
        }
        v4 -= (^(v4))
    }
    v15 := GenS4{F0: 511, F1: 32767, F2: 1}
    v3 = append(v3, 128)
    if (false || (42 < -2147483648)) {
        v16 := GenS1{F0: -644}
        v5 = (v5 + (^(v5)))
        _ = v16
    }
    fmt.Println(v2)
    v17 := GenS3{F0: -764, F1: ("foo bar" + "a"), F2: (^(v2))}
    var v18 uint32 = 2
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(v8, v1, v5)
        _ = fi19
    }
    v18 = (v18 << ((genFn0(v18, 0, false)) & 63))
    switch -320 {
    case 4294967295:
        v8 += 133
    default:
        fmt.Println(v2)
    }
    fmt.Println(v8, v1, v1)
    if ((441 <= 0) && false) {
        var v20 bool = false
        v1 = (-(v1))
        if v20 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v8 -= (^(v8))
            v8 += (v8 / (((^(v8))) | 1))
            for fi21 := 0; fi21 < 9; fi21++ {
                v3 = append(v3, -123)
                v8 = uint(v5)
                _ = fi21
            }
        }
        _ = v20
    } else {
        fmt.Println(v18, v1, v1)
        if true {
            for fi22 := 0; fi22 < 6; fi22++ {
                switch "\n\t" {
                case "a":
                    v4 = v5
                case "foo bar":
                    var v23 uint16 = v4
                    _ = v23
                case "hello":
                    _ = len(v3)
                }
                _ = fi22
            }
            var v24 int8 = v2
            _ = v24
        }
    }
    var v25 float64 = -2.25
    v3 = append(v3, -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
