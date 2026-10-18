package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 bool }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "\n\t"
    v3 := []string{"\n\t", (((v2 + ("hello" + v2)) + v2) + ("\n\t" + (((v2 + (((v2 + v2) + "hello") + ("a" + (((v2 + ((("" + "a") + (v2 + ("\n\t" + "\n\t"))) + v2)) + (v2 + "\n\t")) + "foo bar")))) + v2) + "hello"))), v2, v2}
    var v4 string = "foo bar"
    var v5 uint8 = 255
    v6 := make(map[int]string)
    fmt.Println(v4, v5)
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            if true {
                v9 := v6[128]
                _ = v9
            }
            _ = fi8
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            if true {
                var v11 uint32 = uint32(v5)
                var v12 bool = ((fi7 >= fi10) || (!true))
                switch fi7 {
                case -2147483648:
                    v5 -= v5
                default:
                    switch v4 {
                    case "":
                        v2 = v4
                        var v13 uint16 = 275
                        _ = v13
                    }
                }
                _ = v11
                _ = v12
            } else {
                var v14 uint16 = 255
                if false {
                    var v15 uint16 = 32767
                    v16 := GenS1{F0: 0, F1: 9223372036854775807, F2: true}
                    var v17 uint = 362
                    _ = v15
                    _ = v16
                    _ = v17
                }
                if len(v3) > 0 {
                    v3[0] = ""
                }
                _ = v14
            }
            switch v2 {
            case "hello":
                var v18 int16 = -939
                _ = v18
            default:
                _ = len(v3)
            }
            v19 := GenS2{F0: -2147483648}
            _ = fi10
            _ = v19
        }
        var v20 int64 = 9223372036854775807
        _ = fi7
        _ = v20
    }
    v21 := GenS2{F0: -348}
    if (false && true) {
        var v22 float64 = float64(((1.5 - 1.5) + (-2.25 - 3.14159)))
        fmt.Println(v4)
        _ = v22
    } else {
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 float32 = 0.0
            _ = fi23
            _ = v24
        }
        if true {
            fmt.Println(v5)
            var v25 int64 = -581
            _ = v25
        } else {
            if len(v3) > 0 {
                v3[0] = v4
            }
        }
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        fmt.Println(v4, v5, v2)
        _ = fi26
    }
    v4 = v4
    v27 := GenS1{F0: -128, F1: -961, F2: true}
    v2 = "foo bar"
    for fi28 := 0; fi28 < 1; fi28++ {
        v1[("hello" + v4)] = 796
        var v29 float32 = float32(0.0)
        if false {
            v2 = v2
            _ = len(v3)
            var v30 int64 = -9223372036854775808
            _ = v30
        }
        _ = fi28
        _ = v29
    }
    var v31 float32 = (1.5 - (-2.25 - 1e10))
    var v32 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
