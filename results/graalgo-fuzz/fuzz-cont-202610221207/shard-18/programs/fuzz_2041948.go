package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 1
    var v3 float32 = 1.5
    var v4 int32 = -356
    v5 := GenS1{F0: 727}
    v6 := GenS2{F0: -470, F1: 580}
    if (("a" + "a") == "a") {
        fmt.Println(v3)
    }
    if true {
        delete(v1, 1)
        delete(v1, 2)
        if true {
            for fi7 := 0; fi7 < 4; fi7++ {
                var v8 uint8 = 127
                v2 = (v2 >> ((255) & 63))
                _ = fi7
                _ = v8
            }
            if (float64(v3) >= 0.0) {
                if ((v4 >> ((2) & 63)) < 1) {
                    fmt.Println(v3, v3, v2)
                    var v9 float64 = (float64((0.0 * ((float32(float64(-2.25)) / (1.5 + (v3 * v3))) + 0.1))) / (0.1 * float64(v3)))
                    _ = v9
                } else {
                    v4 += 1
                    var v10 uint = 18446744073709551615
                    if false {
                        v11 := GenS2{F0: 0, F1: 197}
                        _ = v11
                    }
                    _ = v10
                }
                fmt.Println(v4, v3)
                var v12 int8 = -1
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 6; fi13++ {
                if true {
                    var v14 string = ("a" + "foo bar")
                    v15 := GenS2{F0: 0, F1: 2}
                    _ = v14
                    _ = v15
                } else {
                    var v16 bool = (0 >= 2)
                    _ = v16
                }
                var v17 uint = 699
                _ = fi13
                _ = v17
            }
            var v18 int32 = (v4 << ((758) & 63))
            _ = len(v1)
            _ = v18
        }
    } else {
        v4 -= 42
        switch (("\n\t" + "") + "hello") {
        case "\n\t":
            v19 := GenS2{F0: 1, F1: 0}
            _ = v19
        case "":
            if (false || false) {
                var v20 int8 = -128
                v21 := GenS1{F0: 0}
                _ = v20
                _ = v21
            } else {
                var v22 int64 = 42
                _ = v22
            }
        }
        fmt.Println(v4, v4, v3)
    }
    v2 -= v2
    var v23 int32 = v4
    for fi24 := 0; fi24 < 8; fi24++ {
        v3 = (-2.25 * ((0.1 * v3) - float32((((((((-2.25 - 1e10) + (3.14159 / (-2.25 * 1e10))) - ((float64(-2.25) - ((0.1 + 0.1) * 1.5)) * 1.5)) + 1e10) - 0.1) / 1.5) * (-2.25 - ((0.0 * 0.1) - float64(v3)))))))
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 uint16 = 255
            _ = fi25
            _ = v26
        }
        switch "" {
        case "hello":
            if false {
                fmt.Println(v2)
                delete(v1, 65535)
                _ = len(v1)
            } else {
                if (0 < 0) {
                    delete(v1, 826)
                    v27 := GenS2{F0: 9223372036854775807, F1: 739}
                    v2 = v2
                    _ = v27
                }
            }
        case "\n\t":
            var v28 uint8 = 1
            _ = v28
        case "a":
            v4 = v23
        default:
            v3 -= 0.0
        }
        _ = fi24
    }
    v2 -= (^(v2))
    var v29 string = "foo bar"
    fmt.Println(v3, v29, v23)
    delete(v1, 1)
    v30 := GenS2{F0: -1, F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
