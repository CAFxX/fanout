package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 uint = 255
    v2 := make(map[int]bool)
    var v3 int32 = -2147483648
    v4 := make(map[string]int)
    var v5 int16 = 1
    v6 := []string{}
    fmt.Println(v1)
    var v7 int8 = 19
    if (false && true) {
        v3 = (v3 & (v3 ^ v3))
        _ = len(v4)
    }
    v1 = 18446744073709551615
    fmt.Println(v1, v5)
    delete(v2, 42)
    if (242 != 513) {
        v3 = (v3 % (((v3 | v3)) | 1))
        for fi8 := 0; fi8 < 17; fi8++ {
            v5 = (v5 + v5)
            switch (-(fi8)) {
            case 32767:
                v9 := GenS1{F0: 128, F1: 255}
                _ = v9
            }
            var v10 float32 = (0.0 + float32(float64(0.1)))
            _ = fi8
            _ = v10
        }
        var v11 uint16 = 255
        _ = v11
    }
    fmt.Println(v5)
    v12 := GenS1{F0: 42, F1: 42}
    var v13 int8 = -95
    var v14 float32 = float32(1e10)
    if false {
        var v15 float64 = float64(v14)
        switch "" {
        case "foo bar":
            for fi16 := 0; fi16 < 11; fi16++ {
                if ("foo bar" > "\n\t") {
                    var v17 string = "hello"
                    for fi18 := 0; fi18 < 11; fi18++ {
                        var v19 int16 = (-(v5))
                        fmt.Println(v1, v1)
                        v13 = (v7 * v7)
                        var v20 uint64 = 0
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    _ = v17
                } else {
                    v21 := v2[-756]
                    var v22 uint16 = 0
                    var v23 uint8 = 0
                    _ = v21
                    _ = v22
                    _ = v23
                }
                _ = fi16
            }
        case "hello":
            if len(v6) > 0 {
                v6[0] = ("\n\t" + "foo bar")
            }
        default:
            for fi24 := 0; fi24 < 11; fi24++ {
                v6 = append(v6, "foo bar")
                _ = len(v6)
                v14 += -2.25
                _ = fi24
            }
        }
        _ = v15
    } else {
        fmt.Println(v1, v7, v5)
        v14 = float32((1e10 + -2.25))
    }
    v14 -= 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
