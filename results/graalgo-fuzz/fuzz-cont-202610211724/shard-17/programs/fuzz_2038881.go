package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 bool; F2 uint64 }
type GenS3 struct { F0 int16; F1 int64; F2 bool }
type GenS4 struct { F0 bool; F1 uint16 }
func genFn0(p0 int32, p1 int32) int32 {
    v5 := []int{32768}
    v6 := GenS1{F0: ("" + ("foo bar" + ("\n\t" + "\n\t")))}
    if len(v5) > 0 {
        v5[0] = -1
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 float64 = (1.5 * 0.0)
        if (v8 >= float64(0.1)) {
            v9 := GenS1{F0: "foo bar"}
            v8 = v8
            _ = v9
        }
        if (!false) {
            p0 = 927
        }
        _ = fi7
        _ = v8
    }
    v10 := GenS1{F0: ("\n\t" + ("a" + ("foo bar" + ("" + "a"))))}
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    _ = v6
    return (p0 & (p1 - (p1 * 42)))
}

func main() {
    var v1 uint64 = 1
    var v2 float64 = (3.14159 - ((3.14159 - 0.0) - (3.14159 + (-2.25 * 1.5))))
    v3 := make(map[uint32]uint32)
    var v4 bool = (9223372036854775807 >= -846)
    var v11 bool = v4
    switch 553 {
    case -832:
        switch "hello" {
        case "\n\t":
            v12 := GenS1{F0: "foo bar"}
            _ = v12
        default:
            v2 = (1.5 + ((1.5 * v2) / (float64(0.0) / (v2 - v2))))
        }
    case -2147483648:
        var v13 uint8 = 128
        _ = v13
    }
    v14 := []int{439, -813, -668, -568}
    fmt.Println(v2)
    var v15 int32 = -412
    var v16 uint32 = 2
    for fi17 := 0; fi17 < 20; fi17++ {
        if v11 {
            v14 = append(v14, (fi17 + fi17))
            v4 = v11
            switch fi17 {
            case 128:
                var v18 float64 = v2
                _ = v18
            default:
                var v19 int64 = 42
                _ = v19
            }
        }
        _ = fi17
    }
    switch 4294967295 {
    case -208:
        v16 = (^(v16))
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        var v21 bool = v4
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        if true {
            fmt.Println(v4, v1, v1)
            _ = len(v14)
        }
        if (!false) {
            var v23 int = (fi22 - int(v16))
            _ = v23
        } else {
            v24 := GenS2{F0: 127, F1: true, F2: 105}
            fmt.Println(v4, v4)
            var v25 float64 = 1e10
            _ = v24
            _ = v25
        }
        var v26 int8 = 127
        _ = fi22
        _ = v26
    }
    v27 := GenS2{F0: 42, F1: true, F2: v1}
    v28 := GenS3{F0: 1, F1: 9223372036854775807, F2: false}
    v29 := GenS4{F0: false, F1: 2}
    var v30 int8 = 0
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
