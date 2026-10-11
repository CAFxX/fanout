package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 bool }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint32; F1 int64; F2 int8 }
func genFn0(p0 uint) int16 {
    p0 += p0
    p0 = p0
    p0 = 0
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int8 = 100
        v9 := GenS1{F0: 124, F1: 58, F2: false}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    return 1000
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 0.1
    var v3 string = "a"
    v4 := []string{}
    var v5 uint32 = 518
    v6 := make(map[string]bool)
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 string = ((v3 + "hello") + v3)
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 int16 = genFn0(402)
        _ = fi12
        _ = v13
    }
    if true {
        switch v3 {
        case "\n\t":
            v14 := GenS1{F0: 1000, F1: 127, F2: (0 == -882)}
            _ = v14
        case "":
            v2 = v2
        case "a":
            v15 := GenS2{F0: false, F1: v3}
            _ = v15
        }
        v16 := GenS1{F0: 390, F1: 39, F2: (1.5 >= (3.14159 / ((float32((((v2 - ((float64(1.5) * (0.1 * v2)) - (v2 / float64(0.1)))) - float64(float32(0.1))) + v2)) / 1.0) - (1e10 + 1.5))))}
        v1[616] = v5
        _ = v16
    } else {
        if (!true) {
            var v17 uint = 0
            _ = len(v4)
            v18 := v6[v3]
            _ = v17
            _ = v18
        }
        v19 := GenS3{F0: 604, F1: 9223372036854775807, F2: 125}
        _ = v19
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            v5 = v5
            switch v3 {
            case "foo bar":
                var v22 int16 = 0
                _ = v22
            default:
                v3 = (v3 + "foo bar")
            }
            _ = fi21
        }
        v2 -= v2
        _ = fi20
    }
    fmt.Println(v2)
    var v23 uint64 = 0
    v2 += (v2 + float64(0.0))
    if false {
        if false {
            if (32767 < 32767) {
                v3 = v3
            }
            v3 = (v3 + ("" + v3))
        } else {
            v1[(v5 + v5)] = v5
        }
    } else {
        fmt.Println(v5)
        switch 128 {
        case 32768:
            v2 += 1.5
        default:
            v24 := GenS3{F0: v5, F1: 815, F2: -1}
            _ = v24
        }
        var v25 uint = 479
        _ = v25
    }
    v2 = float64((0.0 - 0.0))
    for fi26 := 0; fi26 < 9; fi26++ {
        v3 = ((("hello" + "") + v3) + (v3 + v3))
        fmt.Println(v5, v3)
        v4 = append(v4, ((("\n\t" + v3) + v3) + ""))
        _ = fi26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
