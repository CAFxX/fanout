package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int8; F1 uint32; F2 bool }
type GenS4 struct { F0 int64; F1 uint8; F2 string }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint32 }
func genFn0() uint8 {
    var v5 uint32 = 8
    v5 = (v5 * (^(v5)))
    v6 := GenS1{F0: "foo bar"}
    if true {
        fmt.Println(v5)
        v7 := []int{}
        if (394 >= 65535) {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v5)
            case "\n\t":
                v8 := GenS1{F0: ("a" + "a")}
                _ = v8
            case "a":
                for fi9 := 0; fi9 < 4; fi9++ {
                    var v10 float64 = (1e10 - -2.25)
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    _ = fi9
                    _ = v10
                }
            default:
                var v11 float32 = 1.5
                _ = v11
            }
        } else {
            var v12 int8 = -73
            if (814 <= 0) {
                for fi13 := 0; fi13 < 15; fi13++ {
                    v14 := GenS2{F0: 42, F1: (!((v12 << ((220) & 63)) == (v12 * 42)))}
                    v15 := GenS2{F0: 0, F1: false}
                    fmt.Println(v5)
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
            fmt.Println(v5, v12)
            _ = v12
        }
        _ = v7
    } else {
        switch 127 {
        case 9223372036854775807:
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 float32 = ((((0.1 * 0.1) + -2.25) - 3.14159) / 0.1)
                v18 := []int{fi16}
                v5 = 31
                _ = fi16
                _ = v17
                _ = v18
            }
        case -9223372036854775808:
            for fi19 := 0; fi19 < 20; fi19++ {
                fmt.Println(fi19, v5)
                _ = fi19
            }
        case -1:
            switch 128 {
            case 410:
                var v20 int16 = 1
                _ = v20
            case 32767:
                if ("" != "foo bar") {
                    fmt.Println(v5)
                    v5 = (v5 * (^(v5)))
                } else {
                    var v21 int16 = 537
                    v21 -= -1
                    var v22 int32 = -1
                    _ = v21
                    _ = v22
                }
            case -363:
                v23 := GenS3{F0: 1, F1: 47, F2: (!false)}
                _ = v23
            }
        }
    }
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[string]bool)
    v3 := []int{}
    v4 := make(map[int]bool)
    for fi24 := 0; fi24 < 1; fi24++ {
        switch fi24 {
        case -186:
            v25 := GenS4{F0: 653, F1: 128, F2: "foo bar"}
            _ = v25
        default:
            var v26 int = fi24
            _ = v26
        }
        _ = fi24
    }
    var v27 int32 = -475
    var v28 uint8 = 255
    var v29 int16 = 0
    v30 := v4[-402]
    if ((!(!true)) && (112 <= 42)) {
        v31 := GenS5{F0: (v1 / ((42) | 1)), F1: uint64(v28), F2: 0}
        _ = v31
    }
    fmt.Println(v27, v29)
    v27 -= v27
    var v32 int16 = 1
    var v33 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
}
