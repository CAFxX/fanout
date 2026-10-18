package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 int32; F2 int64 }
func genFn0(p0 uint32) uint32 {
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint16 = 347
        switch fi7 {
        case 562:
            var v9 int16 = 32767
            _ = v9
        default:
            v8 -= v8
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    return p0
}
func genFn1(p0 uint8, p1 int64) uint {
    v10 := GenS1{F0: -424, F1: false}
    p1 = 774
    switch "hello" {
    case "foo bar":
        var v11 int = -9223372036854775808
        _ = v11
    case "\n\t":
        for fi12 := 0; fi12 < 6; fi12++ {
            for fi13 := 0; fi13 < 10; fi13++ {
                if true {
                    p1 = (p1 % ((0) | 1))
                }
                fmt.Println(p1)
                _ = fi13
            }
            _ = fi12
        }
    }
    switch "foo bar" {
    case "foo bar":
        if false {
            for fi14 := 0; fi14 < 16; fi14++ {
                v15 := GenS1{F0: (fi14 * fi14), F1: (false || (190 >= (p0 * (p0 / (((p0 - 255)) | 1)))))}
                var v16 float32 = (3.14159 + (((((0.1 + float32(1.5)) - (float32(-2.25) * (float32((((1.5 - 0.1) / 1.0) / 0.1)) / ((((-2.25 + (3.14159 * 3.14159)) + 0.1) / 1.0) + 0.1)))) + 1.5) + (1e10 * 0.1)) * 0.0))
                var v17 int = fi14
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
        }
    case "\n\t":
        fmt.Println(p0, p1)
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = v10
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 109
    var v3 int64 = 488
    var v4 int = 42
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]uint32)
    if (!((v4 - v4) < v4)) {
        var v18 int32 = 1
        _ = v18
    } else {
        for fi19 := 0; fi19 < 3; fi19++ {
            v20 := GenS2{F0: 1}
            _ = fi19
            _ = v20
        }
    }
    v3 = v3
    v3 = v3
    v21 := GenS3{F0: -33, F1: -12, F2: (-(v3))}
    v4 = (v4 | v4)
    var v22 string = ("" + "\n\t")
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 float32 = 1e10
        switch 32768 {
        case 9223372036854775807:
            for fi25 := 0; fi25 < 2; fi25++ {
                var v26 int = v4
                _ = fi25
                _ = v26
            }
        default:
            v27 := []int{(fi23 + (^(fi23))), -1, fi23, (fi23 % ((v4) | 1))}
            _ = v27
        }
        v2 -= (v2 % ((v2) | 1))
        _ = fi23
        _ = v24
    }
    if false {
        v28 := GenS1{F0: (v4 & v4), F1: ((-932 >= (v3 << ((65535) & 63))) || (v2 > (v2 & (v2 % ((40) | 1)))))}
        var v29 int8 = -1
        _ = v28
        _ = v29
    }
    var v30 bool = ("hello" <= "foo bar")
    v31 := GenS1{F0: v4, F1: (108 >= 76)}
    v5[2147483648] = v4
    var v32 int32 = 193
    if ("hello" <= (v22 + "")) {
        var v33 int8 = 44
        fmt.Println(v22, v32, v32)
        _ = v33
    }
    v34 := []int{v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
