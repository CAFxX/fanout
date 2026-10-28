package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 string) uint {
    if ((255 >= 804) || true) {
        for fi7 := 0; fi7 < 16; fi7++ {
            p0 = "hello"
            fmt.Println(p0)
            v8 := GenS1{F0: 2, F1: 209}
            _ = fi7
            _ = v8
        }
        p0 = "foo bar"
        fmt.Println(p0)
    }
    v9 := []int{}
    v10 := GenS1{F0: 127, F1: 255}
    fmt.Println(p0)
    _ = p0
    _ = v10
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    v3 := []int{-658, 32767}
    var v4 uint32 = 0
    v5 := []string{((("a" + "a") + "\n\t") + "foo bar"), "foo bar"}
    v6 := make(map[int]uint32)
    _ = len(v3)
    v11 := GenS1{F0: 2, F1: 0}
    if false {
        v4 = (v4 - 1)
        var v12 string = "hello"
        _ = v12
    }
    v4 += (v4 + v4)
    for fi13 := 0; fi13 < 7; fi13++ {
        v4 += (^(v4))
        switch "a" {
        case "a":
            var v14 float64 = 0.0
            _ = v14
        case "\n\t":
            switch "\n\t" {
            case "a":
                switch ("\n\t" + "hello") {
                case "\n\t":
                    _ = len(v1)
                default:
                    var v15 string = ("\n\t" + "")
                    _ = v15
                }
            case "hello":
                fmt.Println(v4)
            case "foo bar":
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 uint16 = 0
                    _ = fi16
                    _ = v17
                }
            default:
                if true {
                    v4 += (v4 >> ((2) & 63))
                    var v18 int32 = -1
                    _ = len(v2)
                    _ = v18
                } else {
                    var v19 int8 = -128
                    var v20 float64 = float64(1e10)
                    fmt.Println(v20, v19, v19)
                    _ = v19
                    _ = v20
                }
            }
        case "hello":
            if ((2 <= 128) || false) {
                v21 := GenS1{F0: 144, F1: 1}
                _ = v21
            } else {
                if (1 < -152) {
                    fmt.Println(v4, fi13)
                }
                v4 = 991
                var v22 uint16 = 255
                _ = v22
            }
        default:
            v4 = (v4 * v4)
        }
        _ = fi13
    }
    v4 += (^(v4))
    v4 = (v4 * v4)
    v4 = v4
    switch 42 {
    case 32768:
        v4 += v4
    case 4294967295:
        v23 := GenS2{F0: 128, F1: -128}
        _ = v23
    case 699:
        var v24 float32 = (float32((1e10 * (1.5 * -2.25))) / (1.5 / ((0.1 + 1.5) * (float32(1.5) + float32((((float64(float32(-2.25)) / (3.14159 / 1e10)) + 3.14159) - -2.25))))))
        _ = v24
    default:
        if true {
            v4 += v4
            var v25 uint64 = 50
            v26 := GenS3{F0: (v25 ^ v25)}
            _ = v25
            _ = v26
        }
    }
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
