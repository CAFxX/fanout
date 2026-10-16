package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint8; F1 int64; F2 int8 }
type GenS5 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS6 struct { F0 uint32; F1 uint32; F2 uint16 }
func genFn0(p0 uint64, p1 uint8) string {
    p0 = (^(p0))
    var v5 int = -309
    fmt.Println(v5, p0, p1)
    v6 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return "\n\t"
}
func genFn1(p0 int, p1 uint) int64 {
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int32 = 854
        if false {
            v9 := GenS2{F0: fi7, F1: 2}
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    if true {
        p1 += (^(p1))
        p0 = -2147483648
        var v10 uint8 = 255
        _ = v10
    }
    if (((-(p0)) > p0) || (true || (!true))) {
        p0 = -2147483648
    } else {
        for fi11 := 0; fi11 < 12; fi11++ {
            v12 := []int{}
            v13 := GenS3{F0: 9223372036854775807, F1: 128}
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 uint16 = 32767
            var v16 int16 = 1
            for fi17 := 0; fi17 < 16; fi17++ {
                fmt.Println(v16, p1)
                _ = fi17
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        var v18 int32 = 102
        _ = v18
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    var v4 float64 = 1e10
    v19 := GenS1{F0: 65535}
    v1 = (v1 + v1)
    v4 = v4
    if true {
        v4 = 1.5
    }
    v1 = -1
    var v20 float32 = 3.14159
    v21 := GenS4{F0: 128, F1: genFn1(128, 399), F2: 0}
    v22 := v2[""]
    switch "foo bar" {
    case "hello":
        v23 := v2["foo bar"]
        _ = v23
    case "foo bar":
        switch 9223372036854775807 {
        case 826:
            for fi24 := 0; fi24 < 12; fi24++ {
                v25 := GenS2{F0: fi24, F1: 69}
                _ = fi24
                _ = v25
            }
        case 128:
            for fi26 := 0; fi26 < 13; fi26++ {
                v27 := GenS5{F0: 1, F1: -120, F2: 2}
                var v28 string = ("\n\t" + ((("" + ("foo bar" + ("a" + ("" + ("" + "a"))))) + ("\n\t" + "\n\t")) + ""))
                _ = fi26
                _ = v27
                _ = v28
            }
        default:
            v29 := GenS1{F0: (v22 ^ (v22 - (v22 - (v22 - (v22 + 2)))))}
            _ = v29
        }
    case "\n\t":
        v20 -= v20
    default:
        v30 := GenS1{F0: (v22 / ((v22) | 1))}
        _ = v30
    }
    v31 := GenS6{F0: (v22 ^ 2), F1: (v22 + 348), F2: 0}
    for fi32 := 0; fi32 < 17; fi32++ {
        fmt.Println(fi32)
        _ = fi32
    }
    if (uint64(v22) <= 389) {
        var v33 float64 = ((float64(-2.25) + (0.0 + ((0.1 - 1.5) + ((v4 * v4) - 1.5)))) / float64(((v20 - ((1e10 * ((3.14159 / (v20 - 1.5)) - 0.1)) / (v20 + v20))) - ((v20 * v20) * float32(1e10)))))
        var v34 bool = (false || (32767 > -32768))
        _ = v33
        _ = v34
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v4
}
