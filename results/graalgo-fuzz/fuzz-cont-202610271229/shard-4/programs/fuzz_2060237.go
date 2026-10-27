package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 string, p1 uint8, p2 bool) uint {
    if p2 {
        var v4 float64 = (0.0 - float64(float32(0.1)))
        v5 := []int{127, 42, 32767, 9223372036854775807}
        _ = v4
        _ = v5
    }
    v6 := []int{}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 119
}
func genFn1() uint8 {
    switch 194 {
    case -1:
        switch -1 {
        case 32767:
            var v7 uint16 = 65535
            _ = v7
        case -451:
            var v8 string = ("a" + "\n\t")
            _ = v8
        case -1:
            var v9 float64 = float64(3.14159)
            _ = v9
        }
    case 9223372036854775807:
        v10 := GenS1{F0: 128, F1: (("hello" + ((("\n\t" + "foo bar") + "foo bar") + "\n\t")) + "a"), F2: true}
        _ = v10
    default:
        v11 := GenS2{F0: 468}
        _ = v11
    }
    v12 := GenS3{F0: 42, F1: 1}
    _ = v12
    return 42
}

func main() {
    v1 := []int{9223372036854775807, 42}
    var v2 int = 42
    v3 := []uint32{1, 1}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2)
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            v15 := GenS4{F0: 967}
            fmt.Println(fi13)
            _ = fi14
            _ = v15
        }
        var v16 uint64 = 18446744073709551615
        v2 = -1
        _ = fi13
        _ = v16
    }
    if (2 == 674) {
        v2 = (-(v2))
        v2 = -2147483648
    } else {
        if false {
            _ = len(v3)
        }
        var v17 string = ""
        var v18 uint16 = 45
        _ = v17
        _ = v18
    }
    var v19 uint8 = 197
    if (!true) {
        v19 = genFn1()
        var v20 float32 = ((1.5 - 0.0) * float32(-2.25))
        switch "foo bar" {
        case "":
            var v21 int16 = 1000
            _ = v21
        default:
            fmt.Println(v2)
        }
        _ = v20
    }
    if false {
        _ = len(v3)
        v19 = (v19 & (v19 - genFn1()))
        var v22 bool = false
        _ = v22
    } else {
        var v23 float32 = 0.1
        _ = v23
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        for fi25 := 0; fi25 < 5; fi25++ {
            if ((-1 <= -1) || (542 >= -2147483648)) {
                switch (fi25 | fi24) {
                case 496:
                    v2 -= fi24
                default:
                    _ = len(v3)
                }
            }
            var v26 int16 = 802
            v26 -= (v26 % (((v26 + (v26 >> ((255) & 63)))) | 1))
            _ = fi25
            _ = v26
        }
        var v27 string = ("" + "foo bar")
        switch v2 {
        case 0:
            for fi28 := 0; fi28 < 10; fi28++ {
                _ = len(v3)
                v29 := GenS5{F0: 225, F1: -77, F2: v19}
                if false {
                    v19 -= (v19 | 65)
                    var v30 int16 = 1000
                    v31 := GenS6{F0: 65535}
                    var v32 float64 = (1e10 * 1e10)
                    _ = v30
                    _ = v31
                    _ = v32
                } else {
                    fmt.Println(v27, v19)
                    v19 = v19
                }
                _ = fi28
                _ = v29
            }
        }
        _ = fi24
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
}
