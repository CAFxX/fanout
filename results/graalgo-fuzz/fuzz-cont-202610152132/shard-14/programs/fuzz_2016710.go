package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint16; F2 int }
type GenS4 struct { F0 uint8; F1 bool; F2 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 0.0
    var v3 int8 = -8
    v4 := make(map[string]bool)
    v5 := GenS1{F0: -1, F1: ((("\n\t" + "\n\t") + "foo bar") + "")}
    v6 := GenS2{F0: ""}
    switch 885 {
    case 9223372036854775807:
        if (true || (true || false)) {
            for fi7 := 0; fi7 < 18; fi7++ {
                v2 -= (v2 * v2)
                var v8 int = fi7
                v9 := GenS2{F0: ""}
                _ = fi7
                _ = v8
                _ = v9
            }
            if (true && false) {
                if true {
                    var v10 int = 1
                    v11 := GenS3{F0: 9223372036854775807, F1: 421, F2: int(v3)}
                    _ = v10
                    _ = v11
                }
                v12 := []int{}
                _ = v12
            } else {
                fmt.Println(v3, v3)
            }
        }
    case -989:
        var v13 int64 = -758
        _ = v13
    default:
        if (!(!true)) {
            v1[4294967295] = 524
        }
    }
    var v14 uint8 = 127
    for fi15 := 0; fi15 < 15; fi15++ {
        v14 -= v14
        var v16 uint16 = 682
        _ = fi15
        _ = v16
    }
    v17 := []int{-1, 9223372036854775807, 1}
    v18 := GenS2{F0: "foo bar"}
    var v19 uint8 = (^(v14))
    fmt.Println(v14)
    switch "foo bar" {
    case "":
        for fi20 := 0; fi20 < 4; fi20++ {
            if (731 == (fi20 << ((42) & 63))) {
                v19 -= (v19 << ((786) & 63))
                v21 := GenS4{F0: (v19 + (^(v19))), F1: (0 <= 42), F2: v14}
                _ = v21
            }
            _ = fi20
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
