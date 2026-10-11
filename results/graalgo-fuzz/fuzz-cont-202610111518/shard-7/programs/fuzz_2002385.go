package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int16; F2 uint64 }
type GenS4 struct { F0 int64; F1 int32; F2 uint }
func genFn0() int8 {
    var v6 uint = 2
    v7 := []int{42}
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 int32) uint {
    var v8 uint8 = 42
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 10; fi10++ {
            if false {
                var v11 string = ((("foo bar" + "foo bar") + "a") + "\n\t")
                var v12 int = (fi10 + fi9)
                _ = v11
                _ = v12
            }
            _ = fi10
        }
        var v13 int = (fi9 | (fi9 + fi9))
        _ = fi9
        _ = v13
    }
    var v14 int8 = -1
    v14 -= (v14 & (v14 % ((v14) | 1)))
    _ = p0
    _ = v14
    _ = v8
    return 115
}

func main() {
    var v1 int32 = 42
    var v2 float64 = 0.1
    v3 := []uint32{}
    v4 := []uint32{4294967295}
    var v5 uint64 = 18446744073709551615
    var v15 float32 = (-2.25 + (1e10 + 1e10))
    if (-769 == 32767) {
        if (480 <= 185) {
            v16 := GenS1{F0: 42, F1: -87, F2: 100}
            var v17 string = ("foo bar" + ("" + ("" + "")))
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 uint64 = 2
            _ = fi18
            _ = v19
        }
        var v20 float32 = 3.14159
        _ = v20
    }
    v21 := GenS1{F0: 1, F1: 0, F2: -128}
    v1 = 1
    if (2 > 159) {
        v22 := GenS1{F0: 944, F1: 100, F2: 42}
        _ = len(v3)
        _ = v22
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "" {
    case "hello":
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 uint32 = 2147483648
            var v25 float64 = ((v2 * 1.5) - (v2 * (v2 - (0.0 - v2))))
            v26 := GenS2{F0: -247}
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
    case "a":
        v27 := GenS3{F0: -122, F1: 899, F2: 600}
        _ = v27
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        if ((v15 - (0.1 + float32(v2))) < v15) {
            fmt.Println(v2)
            v5 = v5
            v2 = (1.5 * float64(v15))
        }
    }
    if false {
        v28 := GenS4{F0: 1, F1: -2147483648, F2: 256}
        v29 := GenS1{F0: 258, F1: genFn0(), F2: 39}
        _ = v28
        _ = v29
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 int32 = (v1 / (((v1 << ((42) & 63))) | 1))
            _ = fi30
            _ = v31
        }
        for fi32 := 0; fi32 < 12; fi32++ {
            switch "\n\t" {
            case "hello":
                var v33 float32 = v15
                _ = v33
            case "\n\t":
                v2 += (v2 / float64((float32(v2) + (3.14159 - (v15 * ((v15 / v15) / 1e10))))))
            default:
                if (true && (42 == 32767)) {
                    v34 := GenS3{F0: -645, F1: 32767, F2: 4294967295}
                    var v35 float64 = 0.1
                    _ = v34
                    _ = v35
                }
            }
            _ = fi32
        }
    }
    v36 := GenS4{F0: 0, F1: v1, F2: 4294967295}
    switch 32767 {
    case 32768:
        v37 := GenS1{F0: 42, F1: genFn0(), F2: 1}
        _ = v37
    case 128:
        for fi38 := 0; fi38 < 8; fi38++ {
            v1 = (^(v1))
            _ = fi38
        }
    default:
        for fi39 := 0; fi39 < 13; fi39++ {
            v4 = append(v4, 1)
            var v40 uint32 = 42
            _ = fi39
            _ = v40
        }
    }
    v5 += v5
    v15 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
