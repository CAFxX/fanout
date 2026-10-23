package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int; F2 uint8 }
type GenS3 struct { F0 int64; F1 int8; F2 bool }
func genFn0(p0 int16, p1 uint32, p2 uint) bool {
    fmt.Println(p0, p1, p2)
    if false {
        v4 := GenS1{F0: 18446744073709551615, F1: true, F2: 1}
        _ = v4
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return true
}

func main() {
    v1 := []string{"hello"}
    var v2 uint16 = 32767
    var v3 uint8 = 1
    var v5 uint8 = (v3 - v3)
    if (3.14159 == 1.5) {
        if ((((v2 != (v2 - 42)) && true) && false) || false) {
            var v6 uint16 = uint16(v3)
            _ = v6
        } else {
            v2 += v2
            v3 -= (v5 | 127)
            if ((float64(0.1) * 1e10) >= float64(((0.1 * float32(((0.1 - (-2.25 - 3.14159)) + float64((((1.5 / 1.0) * ((float32(1e10) + ((1e10 - -2.25) * (float32(float64(((1e10 * float32(1.5)) + 1.5))) - 0.0))) / (1e10 + float32((3.14159 / 1.0))))) / (3.14159 * 1e10)))))) / -2.25))) {
                v5 = v5
                var v7 int16 = -275
                _ = v7
            }
        }
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        v3 = (^(v3))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi8
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v9 bool = true
    fmt.Println(v9, v3, v2)
    if v9 {
        v10 := GenS2{F0: 2, F1: -780, F2: (v3 + uint8(v2))}
        switch -445 {
        case 447:
            fmt.Println(v9, v3)
        }
        _ = v10
    }
    v5 = v3
    v9 = (42 != 1)
    if v9 {
        _ = len(v1)
        var v11 float32 = 0.0
        v1 = append(v1, ("" + "foo bar"))
        _ = v11
    } else {
        switch -9223372036854775808 {
        case 127:
            fmt.Println(v2, v3, v2)
        case -370:
            v3 -= (v3 - (^(v3)))
        }
        for fi12 := 0; fi12 < 12; fi12++ {
            v3 -= 127
            _ = fi12
        }
        fmt.Println(v9, v2)
    }
    v13 := GenS3{F0: -844, F1: 127, F2: ((0 == 32768) && v9)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v5
    _ = v9
}
