package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int16; F1 int32; F2 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 string }

func main() {
    var v1 int64 = -779
    var v2 int16 = 1
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]bool)
    v5 := []int{128, 32767, 687}
    switch 9223372036854775807 {
    case 32768:
        for fi6 := 0; fi6 < 19; fi6++ {
            _ = len(v5)
            fmt.Println(v1, v2, v2)
            _ = fi6
        }
    }
    v7 := GenS1{F0: ((("hello" + "foo bar") + "a") + "a"), F1: -1}
    for fi8 := 0; fi8 < 1; fi8++ {
        v3[42] = 150
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 uint64 = 0
            v11 := v4[uint32(v10)]
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := GenS2{F0: v2, F1: 1, F2: v2}
        _ = fi8
        _ = v12
    }
    if false {
        v1 = v1
        if true {
            v2 = (v2 % ((32767) | 1))
            v2 = 723
        } else {
            for fi13 := 0; fi13 < 2; fi13++ {
                delete(v3, 1)
                _ = fi13
            }
            var v14 float32 = (float32(1e10) / 1e10)
            _ = v14
        }
        v2 += v2
    } else {
        var v15 uint16 = 2
        v16 := GenS2{F0: (-(v2)), F1: -309, F2: v2}
        for fi17 := 0; fi17 < 14; fi17++ {
            fmt.Println(fi17, fi17)
            v2 -= (v2 % ((42) | 1))
            _ = fi17
        }
        _ = v15
        _ = v16
    }
    v2 = (-(v2))
    v18 := GenS3{F0: 1}
    switch 42 {
    case 9223372036854775807:
        var v19 int16 = 488
        _ = v19
    case 128:
        v20 := GenS2{F0: (v2 >> ((1) & 63)), F1: -639, F2: (v2 % ((v2) | 1))}
        _ = v20
    }
    v1 = v1
    if true {
        if false {
            var v21 float32 = (1e10 / ((-2.25 + (3.14159 + 1e10)) + float32(-2.25)))
            v22 := GenS4{F0: ""}
            _ = v21
            _ = v22
        }
    }
    if false {
        v1 = (v1 / (((^(v1))) | 1))
        switch 42 {
        case 128:
            var v23 bool = true
            _ = v23
        case -939:
            var v24 float64 = 3.14159
            _ = v24
        default:
            v1 = (v1 * v1)
        }
    } else {
        for fi25 := 0; fi25 < 10; fi25++ {
            v26 := GenS5{F0: (v2 / (((v2 << ((1) & 63))) | 1)), F1: "hello"}
            v27 := v4[2]
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 uint16 = 284
                switch ("a" + "a") {
                case "":
                    _ = len(v4)
                case "a":
                    v30 := GenS5{F0: v2, F1: ""}
                    _ = v30
                }
                fmt.Println(v27)
                _ = fi28
                _ = v29
            }
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    var v31 float32 = 3.14159
    for fi32 := 0; fi32 < 15; fi32++ {
        v2 = v2
        _ = fi32
    }
    var v33 int8 = 127
    if len(v5) > 0 {
        v5[0] = 32768
    }
    var v34 string = ((("" + "") + ("" + ("foo bar" + "\n\t"))) + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v7
}
