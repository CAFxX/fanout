package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }

func main() {
    v1 := make(map[string]string)
    var v2 string = "foo bar"
    var v3 int64 = 42
    v4 := make(map[string]uint32)
    var v5 uint32 = 2
    v2 = v2
    v6 := []int{128, -9223372036854775808, -235}
    v5 -= (v5 % (((v5 % ((0) | 1))) | 1))
    switch 32768 {
    case 182:
        for fi7 := 0; fi7 < 15; fi7++ {
            delete(v4, v2)
            for fi8 := 0; fi8 < 19; fi8++ {
                switch (fi7 - fi8) {
                case -91:
                    v3 = (v3 + (v3 | v3))
                default:
                    v6 = append(v6, (fi7 % ((128) | 1)))
                }
                v2 = ""
                _ = fi8
            }
            v2 = v2
            _ = fi7
        }
    case -205:
        var v9 int32 = 481
        _ = v9
    case -247:
        v3 = (v3 ^ v3)
    default:
        for fi10 := 0; fi10 < 9; fi10++ {
            v5 = v5
            v11 := GenS1{F0: (0.1 != float64((0.1 / (float32(1.5) - -2.25)))), F1: 18446744073709551615}
            v12 := GenS1{F0: true, F1: 0}
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    v5 += v5
    v2 = ((((v2 + ((v2 + (v2 + v2)) + v2)) + "a") + (v2 + ("hello" + v2))) + v2)
    var v13 uint8 = 128
    var v14 int32 = -2147483648
    switch int(v13) {
    case -1:
        var v15 int16 = 32767
        _ = v15
    case 42:
        if false {
            _ = len(v4)
        }
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        fmt.Println(v5)
        _ = fi16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
