package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 9223372036854775807
    var v3 uint16 = 255
    v4 := make(map[int]uint32)
    var v5 uint16 = 0
    var v6 uint8 = 42
    if (false || false) {
        if (!((((0.1 + float32(0.1)) - (1e10 - 3.14159)) - -2.25) > ((0.1 + float32((0.0 - 0.0))) - 1e10))) {
            for fi7 := 0; fi7 < 9; fi7++ {
                v2 = v2
                v3 = v3
                for fi8 := 0; fi8 < 8; fi8++ {
                    var v9 uint16 = (v3 | (^(v5)))
                    if (-1 < 127) {
                        fmt.Println(v5, v9)
                        v2 = (v2 & (^(v2)))
                        v9 += (v5 << ((2) & 63))
                        v10 := GenS1{F0: "foo bar"}
                        v11 := GenS2{F0: 18446744073709551615}
                        _ = v10
                        _ = v11
                    }
                    switch fi8 {
                    case 1:
                        var v12 uint = 832
                        _ = v12
                    case 42:
                        v13 := GenS2{F0: 18446744073709551615}
                        _ = v13
                    case 32768:
                        v6 = 2
                    default:
                        fmt.Println(fi7)
                    }
                    _ = fi8
                    _ = v9
                }
                _ = fi7
            }
        }
    }
    var v14 float32 = (0.0 * 0.1)
    v15 := []int{-9223372036854775808, 4294967295, -1}
    var v16 int16 = -1
    v17 := GenS1{F0: "hello"}
    fmt.Println(v6)
    v6 -= (v6 << ((1) & 63))
    fmt.Println(v16, v14, v14)
    v14 = v14
    if (-2147483648 <= 901) {
        switch "\n\t" {
        case "a":
            v3 = v3
        case "":
            if (v6 != 1) {
                if false {
                    fmt.Println(v6, v2, v6)
                }
                var v18 int16 = v16
                _ = v18
            }
        }
        fmt.Println(v5, v14)
        v16 = int16(v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
