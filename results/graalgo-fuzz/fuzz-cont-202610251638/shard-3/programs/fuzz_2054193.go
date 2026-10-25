package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS3 struct { F0 int; F1 int; F2 uint16 }

func main() {
    v1 := []string{"a", "foo bar", ("" + "hello"), ""}
    var v2 float32 = -2.25
    var v3 uint16 = 255
    v4 := make(map[string]int)
    var v5 uint32 = 2147483648
    fmt.Println(v2)
    var v6 int8 = 1
    v6 = (v6 ^ v6)
    var v7 uint32 = v5
    var v8 float32 = v2
    v5 -= uint32(v3)
    if false {
        var v9 uint32 = (v5 - v5)
        if false {
            var v10 string = ((("foo bar" + "foo bar") + ("hello" + "\n\t")) + "")
            _ = v10
        }
        _ = v9
    }
    var v11 float64 = ((0.1 + (0.0 / 1.5)) - 1e10)
    v8 += float32(float64(((float32(((v11 * float64(v8)) + v11)) / (v2 / v2)) - v8)))
    var v12 float32 = ((v2 - ((v2 / 1e10) * -2.25)) / 0.1)
    _ = len(v4)
    v13 := GenS1{F0: ("hello" + "hello")}
    v1 = append(v1, "\n\t")
    if (v3 < 32767) {
        switch 286 {
        case -558:
            fmt.Println(v2, v7, v2)
        default:
            if (!false) {
                for fi14 := 0; fi14 < 7; fi14++ {
                    var v15 int8 = 42
                    var v16 float64 = (v11 + (float64(v8) - 0.0))
                    v17 := GenS1{F0: "\n\t"}
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
            } else {
                var v18 uint32 = (v5 + (v5 | 885))
                var v19 uint16 = v3
                if ((v5 <= 4294967295) || true) {
                    var v20 uint32 = 2147483648
                    _ = v20
                } else {
                    var v21 float64 = float64(0.0)
                    for fi22 := 0; fi22 < 16; fi22++ {
                        var v23 int16 = 1
                        var v24 uint16 = 721
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    _ = v21
                }
                _ = v18
                _ = v19
            }
        }
        v8 = 1e10
    }
    v25 := GenS2{F0: 127, F1: (^(v5)), F2: -795}
    var v26 int16 = 0
    v27 := GenS3{F0: 32768, F1: -961, F2: (v3 % (((v3 - 363)) | 1))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
