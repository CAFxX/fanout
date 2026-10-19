package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int32; F2 uint16 }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{42, 1, 2147483648}
    v3 := make(map[string]uint32)
    var v4 bool = true
    fmt.Println(v4)
    for fi5 := 0; fi5 < 2; fi5++ {
        if (!false) {
            for fi6 := 0; fi6 < 16; fi6++ {
                if (!(-1 < 42)) {
                    if (((9223372036854775807 > -38) || v4) && false) {
                        var v7 uint8 = 128
                        var v8 bool = (!true)
                        _ = v7
                        _ = v8
                    }
                    fmt.Println(fi5)
                    v9 := GenS1{F0: -85, F1: (fi6 << ((856) & 63)), F2: "\n\t"}
                    _ = v9
                } else {
                    v4 = (4294967295 != 572)
                    var v10 uint16 = 0
                    var v11 float32 = (((1e10 + ((1e10 * (0.1 * ((0.0 * 0.0) / 1.5))) * ((1.5 / 1.5) + 3.14159))) / ((0.0 + 0.0) - 3.14159)) - 1.5)
                    _ = v10
                    _ = v11
                }
                v4 = ((!(!false)) && v4)
                _ = len(v2)
                var v12 int32 = 2147483647
                _ = fi6
                _ = v12
            }
            v4 = (!false)
        }
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        _ = fi5
    }
    v4 = v4
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 float64 = float64(((1e10 / -2.25) + 0.1))
        v15 := GenS1{F0: 42, F1: fi13, F2: "foo bar"}
        v4 = false
        _ = fi13
        _ = v14
        _ = v15
    }
    if v4 {
        if (!true) {
            fmt.Println(v4)
        }
    } else {
        var v16 int64 = 1
        _ = v16
    }
    if v4 {
        for fi17 := 0; fi17 < 5; fi17++ {
            switch ("\n\t" + "foo bar") {
            case "a":
                v4 = (-2.25 != ((0.0 - float32(0.0)) - 3.14159))
            case "foo bar":
                fmt.Println(v4, fi17)
            case "\n\t":
                switch "" {
                case "hello":
                    v18 := GenS2{F0: (^(fi17))}
                    _ = v18
                case "a":
                    v19 := GenS3{F0: 1, F1: 2147483647, F2: 505}
                    _ = v19
                default:
                    fmt.Println(v4)
                }
            default:
                v4 = (!v4)
            }
            _ = fi17
        }
    } else {
        v4 = (!v4)
        v4 = (v4 && true)
        var v20 int64 = 239
        _ = v20
    }
    var v21 int = 0
    v4 = (((0.1 / float64(float32((-2.25 * -2.25)))) - 0.0) > 0.0)
    var v22 int64 = 1
    v4 = (281 <= -977)
    v22 = 531
    v23 := GenS3{F0: 42, F1: -1, F2: 255}
    v22 -= v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
