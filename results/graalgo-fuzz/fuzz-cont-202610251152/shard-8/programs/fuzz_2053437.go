package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint }
func genFn0() uint {
    switch "hello" {
    case "":
        var v6 uint16 = 255
        _ = v6
    default:
        var v7 uint8 = 255
        _ = v7
    }
    var v8 int64 = 0
    v9 := GenS1{F0: 32767}
    var v10 uint = 0
    _ = v10
    _ = v8
    _ = v9
    return 541
}

func main() {
    v1 := []string{"a", "\n\t", "foo bar", "foo bar"}
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    v4 := make(map[string]uint32)
    var v5 int16 = 42
    v5 = (v5 % ((v5) | 1))
    v5 = v5
    v5 = 0
    fmt.Println(v5)
    for fi11 := 0; fi11 < 14; fi11++ {
        v1 = append(v1, ("a" + ""))
        var v12 float32 = (1.5 / 3.14159)
        v13 := GenS2{F0: 65535, F1: true, F2: 2}
        _ = fi11
        _ = v12
        _ = v13
    }
    v5 = -561
    var v14 float64 = (0.0 - 0.1)
    var v15 bool = (0.0 != 3.14159)
    var v16 uint16 = 32767
    v14 -= v14
    switch ("foo bar" + "a") {
    case "\n\t":
        switch "" {
        case "\n\t":
            _ = len(v4)
        case "":
            v15 = v15
        case "foo bar":
            fmt.Println(v14, v14, v5)
        }
    case "a":
        var v17 uint8 = 148
        _ = v17
    case "hello":
        v18 := GenS2{F0: (v16 - (v16 >> ((4294967295) & 63))), F1: (true && ((float64((3.14159 - ((float32((0.0 / 1.0)) / 1.5) - float32(1e10)))) * v14) <= v14)), F2: genFn0()}
        _ = v18
    default:
        var v19 uint16 = v16
        _ = v19
    }
    var v20 float64 = float64(0.0)
    if true {
        switch ("a" + "a") {
        case "a":
            var v21 float32 = 0.1
            _ = v21
        case "\n\t":
            v16 -= (v16 ^ 255)
        case "foo bar":
            v15 = v15
        }
        if v15 {
            if ((v15 && (255 != 42)) && v15) {
                switch 0 {
                case 378:
                    fmt.Println(v14, v15, v20)
                }
            }
            v14 += v20
            var v22 string = (("a" + "\n\t") + (("foo bar" + "") + ("" + ("" + ("a" + "a")))))
            _ = v22
        }
        fmt.Println(v14, v15, v14)
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        if (-657 < 0) {
            v20 += ((-2.25 * -2.25) * v20)
        }
        _ = fi23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
