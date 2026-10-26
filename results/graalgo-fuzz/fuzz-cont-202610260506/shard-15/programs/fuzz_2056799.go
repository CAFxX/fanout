package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint16 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint; F1 uint; F2 uint32 }

func main() {
    var v1 int16 = -32768
    var v2 uint32 = 0
    v3 := []string{"\n\t", "", "foo bar", ("\n\t" + ("" + ("hello" + ("\n\t" + "foo bar"))))}
    v4 := make(map[int]uint32)
    v5 := []uint32{}
    var v6 uint16 = 1
    var v7 uint = 42
    v7 -= v7
    var v8 uint16 = (v6 << (((^(v7))) & 63))
    v9 := GenS1{F0: (-9223372036854775808 < int(v2)), F1: 94, F2: v6}
    var v10 int32 = 2147483647
    fmt.Println(v7, v7, v2)
    fmt.Println(v2)
    v6 = 709
    if len(v5) > 0 {
        v5[0] = v2
    }
    v11 := GenS1{F0: (((!(!false)) && false) || ((v7 / (((v7 | (^(v7)))) | 1)) < uint(v2))), F1: -79, F2: (v6 * 2)}
    var v12 int = 9223372036854775807
    var v13 uint32 = (v2 * uint32(v6))
    switch (^(v12)) {
    case 127:
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS1{F0: (false || true), F1: 127, F2: v6}
            _ = len(v4)
            _ = fi14
            _ = v15
        }
    default:
        var v16 int32 = v10
        _ = v16
    }
    v13 = (v2 & (v2 ^ v13))
    var v17 int16 = (v1 >> ((v7) & 63))
    if ((-41 < -116) || true) {
        switch int(v1) {
        case 9223372036854775807:
            var v18 int = (v12 << (((^(v7))) & 63))
            _ = v18
        case 127:
            v19 := GenS2{F0: "a", F1: ("hello" <= "a")}
            _ = v19
        default:
            var v20 uint64 = 7
            _ = v20
        }
    } else {
        v21 := GenS3{F0: (v7 * (^(v7))), F1: v7, F2: (v13 & (v13 & 65535))}
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
