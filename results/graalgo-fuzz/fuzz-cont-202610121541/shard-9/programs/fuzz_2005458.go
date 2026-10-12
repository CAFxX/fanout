package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int32 }
func genFn0() int8 {
    v4 := []int{127}
    if len(v4) > 0 {
        v4[0] = -924
    }
    v5 := GenS1{F0: 127, F1: 255, F2: 2147483647}
    var v6 int = -912
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(fi7, v6)
        for fi8 := 0; fi8 < 11; fi8++ {
            v6 = (-(fi7))
            switch ("a" + "\n\t") {
            case "hello":
                v9 := GenS1{F0: 143, F1: 42, F2: 2147483647}
                _ = len(v4)
                _ = v9
            case "\n\t":
                if len(v4) > 0 {
                    v4[0] = fi8
                }
                v6 += (v6 / (((v6 ^ (-(fi7)))) | 1))
            case "foo bar":
                var v10 int8 = 100
                _ = v10
            default:
                var v11 uint32 = 380
                _ = v11
            }
            switch "hello" {
            case "a":
                switch "\n\t" {
                case "":
                    if ((65535 != 4294967295) || (!((fi8 * fi8) >= fi7))) {
                        v6 += fi7
                    } else {
                        v12 := GenS1{F0: 255, F1: 574, F2: 0}
                        switch "a" {
                        case "a":
                            var v13 uint = 4294967295
                            _ = v13
                        }
                        fmt.Println(fi7, fi8, v6)
                        _ = v12
                    }
                default:
                    v6 -= (^(fi8))
                }
            case "\n\t":
                v4 = append(v4, v6)
            case "":
                v6 -= fi8
            }
            v6 = -792
            var v14 uint = 2
            _ = fi8
            _ = v14
        }
        _ = fi7
    }
    _ = v4
    _ = v5
    _ = v6
    return 127
}

func main() {
    var v1 int8 = -1
    var v2 int16 = -32768
    v3 := make(map[string]bool)
    v15 := GenS1{F0: 134, F1: 255, F2: 2147483647}
    fmt.Println(v2, v1)
    var v16 string = "hello"
    var v17 uint8 = 149
    if true {
        var v18 uint64 = 0
        _ = v18
    }
    if (true && ((!(2 <= 1)) || false)) {
        var v19 float64 = -2.25
        _ = v19
    }
    v2 = (v2 >> ((228) & 63))
    for fi20 := 0; fi20 < 8; fi20++ {
        fmt.Println(fi20, v17, v17)
        v2 = v2
        var v21 uint32 = 845
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        v17 = 252
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
