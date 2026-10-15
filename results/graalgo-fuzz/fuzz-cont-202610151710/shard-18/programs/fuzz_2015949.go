package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int32; F1 int64 }
func genFn0(p0 bool, p1 uint32, p2 uint16) uint8 {
    v6 := GenS1{F0: uint8(p1)}
    v7 := []int{2147483647, -894}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 199
}
func genFn1(p0 bool, p1 uint32) int8 {
    var v8 string = "a"
    p0 = ((p0 && ((-2.25 > (float32(((0.0 + -2.25) / 1e10)) - 3.14159)) && p0)) || false)
    v9 := GenS1{F0: genFn0(p0, (p1 / ((704) | 1)), 171)}
    p0 = (0.1 <= ((3.14159 / 0.1) - 1.5))
    v10 := []int{-255, 32767, 128, 42}
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int16 = 1
    var v2 float32 = ((0.0 / 1e10) * ((0.0 + (1.5 * 0.1)) + (1e10 + ((-2.25 - 3.14159) * (0.0 + 0.1)))))
    var v3 uint32 = 2
    var v4 uint = 865
    var v5 uint16 = uint16(v4)
    if false {
        v5 -= 42
    }
    var v11 uint32 = 354
    if true {
        var v12 bool = true
        if (9223372036854775807 < 42) {
            if true {
                var v13 float64 = -2.25
                _ = v13
            } else {
                var v14 int64 = -940
                var v15 uint32 = (v11 % (((v3 & v11)) | 1))
                _ = v14
                _ = v15
            }
            v16 := GenS2{F0: 1, F1: 0}
            v17 := []int{-9223372036854775808}
            _ = v16
            _ = v17
        } else {
            fmt.Println(v11, v3, v3)
            v18 := GenS1{F0: uint8(v11)}
            _ = v18
        }
        _ = v12
    }
    if false {
        var v19 uint32 = v11
        _ = v19
    }
    v1 = v1
    var v20 uint8 = 0
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 uint8 = (v20 << ((v4) & 63))
        var v23 uint32 = 2147483647
        v4 += 18446744073709551615
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := []int{0}
    v3 = (^(v3))
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    v24 = append(v24, 958)
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    v1 = (-(v1))
    if false {
        if len(v24) > 0 {
            v24[0] = 4294967295
        }
        if false {
            v24 = append(v24, -1)
            switch ("a" + ("a" + "foo bar")) {
            case "foo bar":
                var v25 int8 = 100
                _ = v25
            default:
                for fi26 := 0; fi26 < 2; fi26++ {
                    v27 := GenS3{F0: 0, F1: 1}
                    var v28 string = "foo bar"
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
        }
        var v29 int8 = genFn1(true, (v11 + (v11 / (((v3 ^ (v3 | (v3 & v11)))) | 1))))
        _ = v29
    }
    switch 4294967295 {
    case 32767:
        v20 -= (^(v20))
    case -2147483648:
        for fi30 := 0; fi30 < 11; fi30++ {
            v31 := GenS3{F0: -1, F1: int64(v1)}
            var v32 int8 = 8
            v24 = append(v24, fi30)
            _ = fi30
            _ = v31
            _ = v32
        }
    default:
        var v33 uint64 = 1
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
