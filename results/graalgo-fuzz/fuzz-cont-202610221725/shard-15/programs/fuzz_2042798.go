package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint; F1 bool }
func genFn0(p0 uint8) bool {
    v6 := GenS1{F0: -1, F1: 198, F2: 32767}
    for fi7 := 0; fi7 < 6; fi7++ {
        switch (("" + "a") + ("foo bar" + "a")) {
        case "hello":
            v8 := GenS2{F0: 4294967295, F1: true}
            _ = v8
        default:
            if false {
                switch ("a" + "hello") {
                case "":
                    switch fi7 {
                    case 32768:
                        var v9 string = "a"
                        _ = v9
                    default:
                        p0 = 2
                    }
                }
                var v10 uint8 = 255
                _ = v10
            }
        }
        _ = fi7
    }
    p0 = (p0 ^ p0)
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    _ = v6
    return true
}
func genFn1() uint64 {
    if false {
        var v11 int8 = 100
        _ = v11
    }
    if (!true) {
        var v12 string = "hello"
        _ = v12
    } else {
        var v13 uint = 255
        switch "foo bar" {
        case "\n\t":
            v13 = (v13 ^ v13)
        case "hello":
            switch 127 {
            case 128:
                v13 -= v13
            case 220:
                var v14 float64 = float64(1.5)
                _ = v14
            default:
                v13 = v13
            }
        case "foo bar":
            var v15 int64 = 42
            _ = v15
        default:
            v13 -= v13
        }
        _ = v13
    }
    var v16 int64 = 0
    fmt.Println(v16)
    _ = v16
    return 2
}

func main() {
    var v1 uint = 998
    var v2 uint = 983
    var v3 uint32 = 42
    v4 := []uint32{v3, v3, 2147483647}
    var v5 int = 32768
    switch "a" {
    case "\n\t":
        var v17 int16 = -543
        _ = v17
    default:
        v3 -= (v3 / (((^(v3))) | 1))
    }
    if len(v4) > 0 {
        v4[0] = 1
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        for fi19 := 0; fi19 < 10; fi19++ {
            if len(v4) > 0 {
                v4[0] = 2
            }
            v20 := GenS2{F0: 211, F1: ((0.1 * 0.1) <= (1e10 / (-2.25 * (float32((float64(1.5) / 3.14159)) - 0.0))))}
            _ = fi19
            _ = v20
        }
        var v21 int32 = -578
        v4 = append(v4, v3)
        _ = fi18
        _ = v21
    }
    var v22 uint16 = 32767
    v23 := GenS2{F0: (^(v2)), F1: true}
    var v24 string = "foo bar"
    var v25 uint8 = 127
    var v26 int = (v5 * v5)
    v1 -= 260
    v26 = v5
    v27 := GenS2{F0: v1, F1: ((("foo bar" + v24) + v24) < (v24 + "a"))}
    var v28 uint = uint(v25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
