package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint32; F2 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int8, p1 int16, p2 uint16) uint32 {
    p0 = (p0 / (((p0 ^ 76)) | 1))
    v7 := GenS1{F0: uint8(p2), F1: 255, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2147483647
}
func genFn1(p0 uint, p1 bool, p2 uint8) int32 {
    if p1 {
        for fi8 := 0; fi8 < 20; fi8++ {
            switch fi8 {
            case -883:
                var v9 uint = p0
                _ = v9
            case 53:
                p2 -= (^(p2))
            case 1:
                v10 := GenS1{F0: (p2 ^ 128), F1: 0, F2: 2147483647}
                _ = v10
            default:
                v11 := []int{128, fi8, -376}
                _ = v11
            }
            _ = fi8
        }
        fmt.Println(p0, p2)
        p1 = ((0.0 <= (float64((1.5 - float32(1.5))) / 1.5)) || p1)
    } else {
        v12 := GenS2{F0: 4294967295}
        var v13 uint64 = 0
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 int32 = 239
            p0 -= p0
            _ = fi14
            _ = v15
        }
        _ = v12
        _ = v13
    }
    var v16 float64 = ((-2.25 + -2.25) - ((float64(1.5) + 3.14159) + 0.1))
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return -292
}

func main() {
    v1 := []uint32{2147483647, 716}
    v2 := make(map[int]int)
    var v3 uint32 = 65535
    var v4 float32 = (((0.1 / 3.14159) / (0.0 + 1.5)) * float32(1.5))
    v5 := make(map[uint32]bool)
    var v6 uint64 = 18446744073709551615
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        v17 := GenS1{F0: 216, F1: 42, F2: -832}
        if (-91 == 100) {
            v6 += v6
            v18 := GenS1{F0: 255, F1: 18446744073709551615, F2: genFn1(4294967295, (42 <= -128), 255)}
            fmt.Println(v4)
            _ = v18
        }
        _ = v17
    }
    switch 42 {
    case -1:
        var v19 float32 = (1.5 - v4)
        _ = v19
    case 0:
        switch "hello" {
        case "hello":
            switch ((("a" + "hello") + "\n\t") + "a") {
            case "a":
                var v20 uint32 = v3
                _ = v20
            case "\n\t":
                v4 += float32(3.14159)
            case "":
                v4 = v4
            }
        default:
            var v21 int32 = 0
            _ = v21
        }
    case -9223372036854775808:
        v3 -= genFn0(1, 32767, 694)
    default:
        if true {
            var v22 int = 2147483647
            _ = v22
        } else {
            var v23 uint8 = 202
            switch 128 {
            case 127:
                v24 := GenS3{F0: 42, F1: (v3 + (v3 % (((^(v3))) | 1))), F2: 4294967295}
                _ = v24
            }
            _ = v23
        }
    }
    v6 = (v6 + v6)
    _ = len(v5)
    var v25 uint = 80
    for fi26 := 0; fi26 < 5; fi26++ {
        v4 = 0.1
        var v27 int16 = 3
        v28 := GenS2{F0: v6}
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := GenS2{F0: v6}
    for fi30 := 0; fi30 < 1; fi30++ {
        if false {
            fmt.Println(v3, fi30, v6)
            if ((true && (false && ((v6 > v6) && (("a" >= "") || true)))) || (1 <= 1000)) {
                for fi31 := 0; fi31 < 12; fi31++ {
                    fmt.Println(v3, fi31, fi30)
                    _ = fi31
                }
            } else {
                if (!false) {
                    v32 := GenS4{F0: 65535}
                    _ = v32
                } else {
                    v25 = 18446744073709551615
                }
            }
        } else {
            v33 := GenS5{F0: 164, F1: 65535}
            v3 -= v3
            _ = v33
        }
        _ = fi30
    }
    var v34 bool = (-227 > -1)
    var v35 float64 = ((-2.25 - 1.5) - (1e10 - ((3.14159 + 3.14159) - float64(v4))))
    var v36 bool = v34
    if false {
        delete(v2, 127)
        var v37 int16 = 1
        _ = v37
    }
    v1 = append(v1, 2147483648)
    v36 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
