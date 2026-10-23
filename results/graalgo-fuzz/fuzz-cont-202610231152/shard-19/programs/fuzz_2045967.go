package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int64; F2 int64 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int16; F1 bool; F2 bool }
func genFn0(p0 int16, p1 int8, p2 uint32) int16 {
    v7 := []int{-439, 205, 32768, 418}
    p2 = (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -454
}
func genFn1(p0 bool) uint8 {
    v8 := GenS1{F0: 42}
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := []int{(^(fi9)), fi9, fi9, (fi9 ^ (fi9 | fi9))}
        v11 := GenS2{F0: 169, F1: 1, F2: 42}
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS3{F0: 74, F1: false}
    v13 := GenS4{F0: -836, F1: p0, F2: false}
    p0 = false
    _ = p0
    _ = v12
    _ = v13
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{677}
    var v2 uint64 = 247
    var v3 uint8 = 127
    v4 := make(map[string]bool)
    v5 := make(map[string]bool)
    v6 := []string{"a", (("hello" + ("foo bar" + ("foo bar" + "a"))) + "\n\t"), "foo bar"}
    var v14 uint32 = 2147483648
    switch 4294967295 {
    case 128:
        switch 32768 {
        case 1:
            switch "" {
            case "\n\t":
                v15 := GenS1{F0: 42}
                _ = v15
            case "foo bar":
                v16 := GenS4{F0: genFn0(-1, -10, v14), F1: false, F2: true}
                _ = v16
            case "hello":
                _ = len(v6)
            default:
                var v17 uint16 = 459
                _ = v17
            }
        case 965:
            var v18 int32 = int32(v2)
            _ = v18
        case -2147483648:
            v3 = v3
        default:
            v2 -= v2
        }
    case -1:
        v19 := v4[((("foo bar" + (("\n\t" + (("\n\t" + "hello") + "\n\t")) + "")) + "\n\t") + "a")]
        _ = v19
    }
    var v20 int32 = 42
    var v21 int16 = -32768
    v22 := GenS2{F0: (v14 & (v14 * v14)), F1: 362, F2: -9223372036854775808}
    for fi23 := 0; fi23 < 12; fi23++ {
        for fi24 := 0; fi24 < 3; fi24++ {
            v6 = append(v6, ("foo bar" + "hello"))
            var v25 string = ("" + ("" + "\n\t"))
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    if false {
        _ = len(v5)
        switch "\n\t" {
        case "a":
            v20 -= v20
        case "":
            if true {
                var v26 int64 = 666
                if len(v1) > 0 {
                    v1[0] = (v14 * 2147483648)
                }
                _ = v26
            } else {
                v14 = 346
            }
        default:
            v4["hello"] = true
        }
    } else {
        v2 -= (v2 + 431)
        v27 := GenS4{F0: v21, F1: false, F2: (v20 >= v20)}
        _ = v27
    }
    var v28 uint32 = (v14 + v14)
    var v29 float64 = (float64((((-2.25 * -2.25) + (1e10 + 1e10)) * 0.1)) * 0.0)
    delete(v5, ("\n\t" + ""))
    v30 := GenS2{F0: 2147483647, F1: 1, F2: 0}
    var v31 int8 = -1
    switch -2147483648 {
    case -898:
        if true {
            var v32 int32 = v20
            var v33 int32 = v32
            _ = v32
            _ = v33
        }
    case -1:
        var v34 int = 9223372036854775807
        _ = v34
    case -2147483648:
        v29 += 0.0
    default:
        for fi35 := 0; fi35 < 18; fi35++ {
            v28 -= v14
            var v36 int16 = (v21 * (v21 / ((v21) | 1)))
            _ = fi35
            _ = v36
        }
    }
    _ = len(v1)
    var v37 int = 128
    switch (^(v37)) {
    case 277:
        var v38 int64 = -1
        _ = v38
    case -693:
        v14 = v14
    default:
        var v39 int64 = 1
        _ = v39
    }
    v3 = (v3 % ((v3) | 1))
    switch "hello" {
    case "foo bar":
        v40 := GenS4{F0: v21, F1: (255 < 255), F2: false}
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
