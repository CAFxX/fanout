package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int }
type GenS2 struct { F0 string }

func main() {
    var v1 int8 = 0
    v2 := []uint32{0}
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]bool)
    v6 := GenS1{F0: v1, F1: 2, F2: 74}
    if (("\n\t" + (("" + "foo bar") + "foo bar")) < "\n\t") {
        v1 = v1
    } else {
        v1 += v1
    }
    v1 = v1
    for fi7 := 0; fi7 < 14; fi7++ {
        if (true && (!false)) {
            for fi8 := 0; fi8 < 19; fi8++ {
                switch ("foo bar" + "\n\t") {
                case "":
                    v9 := GenS2{F0: ((("a" + "hello") + ("foo bar" + "a")) + "a")}
                    _ = v9
                case "a":
                    v1 += -106
                default:
                    fmt.Println(fi7, fi8, fi8)
                }
                _ = fi8
            }
            if false {
                v10 := GenS1{F0: (v1 * (-(v1))), F1: (fi7 ^ fi7), F2: (fi7 + fi7)}
                var v11 int32 = 42
                _ = v10
                _ = v11
            } else {
                v5[65535] = true
                for fi12 := 0; fi12 < 4; fi12++ {
                    var v13 uint16 = 8
                    _ = fi12
                    _ = v13
                }
            }
            v1 = -128
        } else {
            v1 = v1
        }
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 uint = 306
            _ = fi14
            _ = v15
        }
        _ = fi7
    }
    v1 = (v1 * v1)
    switch "foo bar" {
    case "":
        v16 := v3[4294967295]
        _ = v16
    case "\n\t":
        _ = len(v2)
    }
    v17 := GenS2{F0: ("" + "")}
    var v18 uint64 = 2
    fmt.Println(v1, v18)
    v1 += (^(v1))
    for fi19 := 0; fi19 < 18; fi19++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v2 = append(v2, 610)
        _ = fi19
    }
    if false {
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    } else {
        v20 := v4[65535]
        _ = v20
    }
    v18 -= (^(v18))
    v1 += (v1 << ((1) & 63))
    var v21 uint32 = 626
    var v22 float64 = (-2.25 * (((float64(0.0) - -2.25) + 1e10) * float64(((float32(0.0) - (1.5 * 3.14159)) + -2.25))))
    var v23 float64 = ((v22 + ((3.14159 / ((3.14159 - 1.5) / -2.25)) - v22)) + v22)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
