package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint32 }
func genFn0(p0 uint64, p1 uint) bool {
    v4 := GenS1{F0: "\n\t"}
    if (!true) {
        if (32767 == 872) {
            p0 += (^(p0))
        }
        fmt.Println(p1, p1)
        var v5 uint32 = 2147483647
        _ = v5
    } else {
        var v6 int32 = 0
        p1 -= (p1 << ((p1) & 63))
        _ = v6
    }
    v7 := []int{-876, -2147483648, 9223372036854775807, 936}
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(p0, fi8, fi8)
        switch ("\n\t" + ("" + "foo bar")) {
        case "":
            switch (fi8 ^ -722) {
            case 951:
                p0 = (p0 + p0)
            case 9223372036854775807:
                v9 := GenS2{F0: 42}
                _ = v9
            }
        case "hello":
            _ = len(v7)
        default:
            for fi10 := 0; fi10 < 17; fi10++ {
                v7 = append(v7, 826)
                _ = fi10
            }
        }
        v7 = append(v7, fi8)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return true
}
func genFn1(p0 int, p1 uint16) int8 {
    v11 := GenS3{F0: 4294967295, F1: 0, F2: 2}
    var v12 uint16 = (^(p1))
    p0 += 2147483647
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 127
}

func main() {
    var v1 string = "a"
    var v2 float64 = ((((-2.25 + 0.1) + 0.0) - 1.5) * (3.14159 * -2.25))
    v3 := []uint32{2147483647, 2, 2147483648, 753}
    switch 411 {
    case -260:
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 int32 = 330
            _ = fi13
            _ = v14
        }
    case 32768:
        var v15 uint64 = 441
        _ = v15
    default:
        v1 = "hello"
    }
    var v16 int8 = -1
    if ((v2 / v2) > ((1e10 - 3.14159) - (-2.25 / v2))) {
        var v17 int = -9223372036854775808
        _ = v17
    }
    if false {
        v18 := GenS2{F0: 508}
        v16 = v16
        v2 = v2
        _ = v18
    }
    var v19 int32 = 293
    var v20 uint8 = 0
    switch 0 {
    case 0:
        switch v1 {
        case "\n\t":
            v21 := GenS2{F0: 512}
            _ = v21
        }
    case 9223372036854775807:
        v19 -= (v19 * (-(v19)))
    case 128:
        v2 += 0.1
    default:
        var v22 uint64 = 42
        _ = v22
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 int16 = -1
        fmt.Println(v1, v1)
        var v25 float64 = v2
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 int32 = (v19 << ((65535) & 63))
            switch v1 {
            case "a":
                if false {
                    switch v1 {
                    case "\n\t":
                        v20 -= v20
                    case "hello":
                        fmt.Println(v1, v19)
                    default:
                        var v29 bool = ((1e10 < ((1e10 + 0.0) + (1.5 * (3.14159 * 1.5)))) || (v20 == (v20 >> ((uint(v20)) & 63))))
                        _ = v29
                    }
                    var v30 int = fi26
                    _ = v30
                } else {
                    v31 := GenS3{F0: 2147483648, F1: 32767, F2: 2147483648}
                    var v32 uint16 = 42
                    v33 := GenS3{F0: 42, F1: -1, F2: 781}
                    _ = v31
                    _ = v32
                    _ = v33
                }
            default:
                _ = len(v3)
            }
            _ = fi27
            _ = v28
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v20 += v20
        _ = fi26
    }
    if ((((3.14159 - (3.14159 - -2.25)) * (1e10 - 1e10)) / (1.5 / float32((1e10 - v2)))) <= -2.25) {
        if len(v3) > 0 {
            v3[0] = 2
        }
        var v34 uint = 18446744073709551615
        var v35 uint32 = uint32(v20)
        _ = v34
        _ = v35
    } else {
        v3 = append(v3, 1)
        var v36 uint8 = (v20 & (v20 * (^(v20))))
        _ = v36
    }
    for fi37 := 0; fi37 < 11; fi37++ {
        for fi38 := 0; fi38 < 9; fi38++ {
            var v39 float64 = v2
            _ = fi38
            _ = v39
        }
        v2 = (float64(3.14159) / ((float64(3.14159) + (float64((float32(-2.25) + 1e10)) / v2)) - float64((1e10 + float32((1.5 - v2))))))
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
