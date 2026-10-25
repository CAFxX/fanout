package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 string }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint64 = 498
        v7 = v7
        _ = fi6
        _ = v7
    }
    v8 := []int{}
    if true {
        switch -2147483648 {
        case -205:
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := GenS1{F0: 1}
                _ = fi9
                _ = v10
            }
        case 134:
            var v11 uint = 4294967295
            _ = v11
        case 26:
            var v12 int8 = -128
            _ = v12
        }
        var v13 int8 = -1
        _ = v13
    }
    var v14 int32 = 501
    _ = v14
    _ = v8
    return 678
}
func genFn1(p0 int8, p1 uint16, p2 int64) int8 {
    p2 = 0
    fmt.Println(p2, p1)
    var v15 uint = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return p0
}

func main() {
    var v1 int8 = 100
    var v2 uint64 = 18446744073709551615
    var v3 int32 = 0
    var v4 uint64 = v2
    var v5 int32 = v3
    fmt.Println(v1, v1)
    v4 -= genFn0()
    var v16 int16 = -1
    v16 = (v16 | (v16 * -371))
    v17 := []int{128, 4294967295, 4294967295, 32768}
    var v18 int16 = (v16 >> ((18446744073709551615) & 63))
    var v19 int64 = -9223372036854775808
    for fi20 := 0; fi20 < 10; fi20++ {
        v3 = v3
        if ((255 <= 127) && (!(false || (65535 != 18446744073709551615)))) {
            for fi21 := 0; fi21 < 3; fi21++ {
                v22 := GenS1{F0: 2147483648}
                v5 -= (v3 + (v5 - (v5 - (v3 - v5))))
                _ = fi21
                _ = v22
            }
        } else {
            switch fi20 {
            case -1:
                var v23 float32 = (-2.25 / 3.14159)
                _ = v23
            case 128:
                if true {
                    switch (fi20 - (^(fi20))) {
                    case -9223372036854775808:
                        v17 = append(v17, -1)
                    case 876:
                        v19 += v19
                    default:
                        _ = len(v17)
                    }
                    var v24 int = (fi20 >> ((618) & 63))
                    v18 -= (v16 << ((630) & 63))
                    _ = v24
                }
            case 32767:
                if len(v17) > 0 {
                    v17[0] = (-(fi20))
                }
            default:
                v25 := GenS1{F0: 69}
                _ = v25
            }
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            v16 += 1
        }
        _ = fi20
    }
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    fmt.Println(v5, v18)
    v5 += v3
    v26 := GenS2{F0: (v16 - -781), F1: 9223372036854775807}
    v19 -= v19
    var v27 uint8 = 42
    var v28 int8 = genFn1(127, 42, v19)
    var v29 int32 = (v3 % (((v3 + (v5 & (v5 & v3)))) | 1))
    v30 := GenS3{F0: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
