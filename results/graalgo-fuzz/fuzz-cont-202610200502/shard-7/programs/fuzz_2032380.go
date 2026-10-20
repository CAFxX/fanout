package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int16 }
func genFn0(p0 int64, p1 int) int32 {
    if true {
        var v5 uint16 = 65535
        _ = v5
    }
    if true {
        var v6 uint64 = 4294967295
        _ = v6
    } else {
        for fi7 := 0; fi7 < 4; fi7++ {
            switch "\n\t" {
            case "foo bar":
                v8 := GenS1{F0: 42}
                _ = v8
            default:
                for fi9 := 0; fi9 < 18; fi9++ {
                    var v10 int64 = -1
                    var v11 int = 128
                    var v12 int = fi7
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
            }
            var v13 string = "hello"
            for fi14 := 0; fi14 < 17; fi14++ {
                p0 -= (p0 + int64(fi14))
                var v15 int64 = p0
                _ = fi14
                _ = v15
            }
            _ = fi7
            _ = v13
        }
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]uint32)
    v3 := []uint32{2147483648, 796, 65535}
    var v4 uint = 110
    var v16 uint8 = (v1 % ((2) | 1))
    var v17 int32 = genFn0(-1, 763)
    switch "hello" {
    case "foo bar":
        v4 = (^(v4))
    default:
        v1 += (v16 - (v1 / ((v16) | 1)))
    }
    var v18 int8 = -1
    if (1.5 <= float64(0.1)) {
        fmt.Println(v4)
        for fi19 := 0; fi19 < 6; fi19++ {
            v4 += v4
            v17 -= genFn0(0, -2147483648)
            fmt.Println(v4, fi19)
            _ = fi19
        }
        v18 = v18
    }
    fmt.Println(v1)
    if (true && (!false)) {
        switch -45 {
        case -631:
            v20 := v2[((("a" + ("\n\t" + "hello")) + "\n\t") + "\n\t")]
            _ = v20
        case 901:
            var v21 int = -159
            _ = v21
        default:
            v3 = append(v3, 2147483648)
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v4 -= 18446744073709551615
    }
    switch 32768 {
    case 128:
        var v22 int8 = 95
        _ = v22
    case -128:
        _ = len(v2)
    }
    v23 := GenS2{F0: 349, F1: 514, F2: -715}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
