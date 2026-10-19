package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint; F1 int; F2 uint }

func main() {
    var v1 int16 = -1
    var v2 float32 = 1e10
    var v3 int8 = 98
    v4 := make(map[uint32]uint32)
    var v5 uint64 = 672
    var v6 int16 = (v1 << ((42) & 63))
    v2 -= v2
    var v7 bool = (1 >= 1)
    for fi8 := 0; fi8 < 16; fi8++ {
        if (!v7) {
            v5 -= (v5 >> ((255) & 63))
            v9 := []int{(fi8 << ((1) & 63)), fi8, (fi8 + fi8), (-(fi8))}
            v2 = 0.0
            _ = v9
        } else {
            v5 = (v5 >> ((0) & 63))
            var v10 int8 = (^(v3))
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := GenS1{F0: v10}
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        var v13 bool = v7
        _ = fi8
        _ = v13
    }
    fmt.Println(v3, v5, v7)
    for fi14 := 0; fi14 < 13; fi14++ {
        v3 -= (v3 | v3)
        switch "foo bar" {
        case "foo bar":
            if v7 {
                v2 = v2
            }
        case "\n\t":
            for fi15 := 0; fi15 < 13; fi15++ {
                switch (fi15 | 9223372036854775807) {
                case 113:
                    v16 := GenS1{F0: (v3 * v3)}
                    _ = v16
                default:
                    fmt.Println(v3, v5)
                }
                _ = fi15
            }
        default:
            v17 := GenS2{F0: "a", F1: -58}
            _ = v17
        }
        _ = fi14
    }
    v18 := GenS3{F0: 42, F1: 2147483647, F2: 18446744073709551615}
    v6 -= (v1 * (v1 | v6))
    var v19 int64 = -176
    fmt.Println(v2, v19, v19)
    if (v2 != (0.1 + ((v2 - (v2 + 1.5)) / (v2 * v2)))) {
        v1 = v1
        var v20 float32 = v2
        switch -295 {
        case 363:
            for fi21 := 0; fi21 < 16; fi21++ {
                var v22 uint8 = 0
                v2 += ((((v2 + ((-2.25 * 1.5) - (3.14159 - 3.14159))) * ((((0.0 * v2) / v2) - float32(3.14159)) * v20)) - (v20 * v2)) / v2)
                _ = fi21
                _ = v22
            }
        case 32768:
            switch 32768 {
            case 32767:
                v6 = v6
            case 63:
                var v23 int16 = (v1 * v1)
                _ = v23
            }
        default:
            if (!(!((0 > 2) && ((1e10 - (0.1 + 1.5)) == 1e10)))) {
                if (v7 && (0 <= 2147483647)) {
                    if (402 > -882) {
                        fmt.Println(v2, v5, v2)
                    }
                }
                for fi24 := 0; fi24 < 20; fi24++ {
                    fmt.Println(v19, v2, v5)
                    _ = fi24
                }
                v6 -= v6
            }
        }
        _ = v20
    } else {
        var v25 uint32 = 942
        switch "\n\t" {
        case "":
            v25 = 2147483648
        case "a":
            v5 += (v5 % ((v5) | 1))
        }
        v26 := GenS1{F0: -128}
        _ = v25
        _ = v26
    }
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
