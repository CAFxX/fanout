package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint; F1 uint8; F2 string }
type GenS3 struct { F0 string; F1 int8 }
func genFn0(p0 string) int {
    var v5 int8 = 8
    var v6 string = "foo bar"
    _ = p0
    _ = v5
    _ = v6
    return 150
}

func main() {
    v1 := []uint32{247}
    var v2 string = ""
    v3 := []string{}
    var v4 int16 = -1
    for fi7 := 0; fi7 < 15; fi7++ {
        if len(v3) > 0 {
            v3[0] = ((v2 + ("a" + (v2 + v2))) + v2)
        }
        v3 = append(v3, (((((v2 + v2) + "hello") + ((v2 + ((v2 + v2) + "foo bar")) + v2)) + "") + "foo bar"))
        _ = fi7
    }
    switch v2 {
    case "\n\t":
        fmt.Println(v2, v2)
    case "hello":
        v4 = v4
    default:
        if true {
            for fi8 := 0; fi8 < 18; fi8++ {
                for fi9 := 0; fi9 < 7; fi9++ {
                    if false {
                        v3 = append(v3, "\n\t")
                        v10 := GenS1{F0: 558, F1: 18446744073709551615}
                        _ = v10
                    }
                    var v11 int8 = -2
                    _ = fi9
                    _ = v11
                }
                v2 = "a"
                _ = fi8
            }
            var v12 int32 = 42
            for fi13 := 0; fi13 < 17; fi13++ {
                v2 = ((v2 + "\n\t") + ("" + v2))
                _ = fi13
            }
            _ = v12
        }
    }
    var v14 uint = 547
    _ = len(v3)
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 uint8 = 159
        fmt.Println(v16, v14, fi15)
        _ = fi15
        _ = v16
    }
    if ((false && true) || true) {
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        switch v2 {
        case "a":
            if true {
                v2 = "\n\t"
            } else {
                var v17 float32 = (((0.1 * 0.0) * (0.1 - -2.25)) - 3.14159)
                _ = v17
            }
        case "foo bar":
            var v18 uint16 = 255
            _ = v18
        case "":
            var v19 bool = (610 <= 42)
            _ = v19
        default:
            fmt.Println(v14, v2)
        }
    }
    v20 := GenS2{F0: v14, F1: 0, F2: v2}
    fmt.Println(v2)
    var v21 string = ""
    for fi22 := 0; fi22 < 15; fi22++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi22
    }
    v3 = append(v3, (v21 + (v2 + v21)))
    if true {
        fmt.Println(v2)
        v21 = ""
    }
    v23 := GenS1{F0: 0, F1: 42}
    if true {
        v4 -= v4
        var v24 uint8 = 255
        if true {
            v25 := GenS2{F0: v14, F1: 0, F2: v21}
            _ = v25
        }
        _ = v24
    }
    var v26 float64 = ((float64(((1e10 - 1e10) + 0.1)) * 0.1) - -2.25)
    v27 := GenS3{F0: (v2 + v21), F1: -20}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
