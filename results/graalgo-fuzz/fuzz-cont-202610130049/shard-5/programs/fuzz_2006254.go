package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 613
    var v3 int16 = -717
    v4 := make(map[string]uint32)
    var v5 int8 = 0
    for fi6 := 0; fi6 < 18; fi6++ {
        switch "a" {
        case "":
            v2 = v2
        case "hello":
            var v7 int16 = (v3 / ((v3) | 1))
            _ = v7
        case "\n\t":
            fmt.Println(v3, v5, v5)
        }
        var v8 string = "a"
        var v9 uint32 = 778
        _ = fi6
        _ = v8
        _ = v9
    }
    v3 = int16(v5)
    fmt.Println(v5)
    var v10 int16 = v3
    var v11 uint = 466
    v3 = (^(v3))
    v2 = v2
    v2 -= 32768
    if (true && (!(!((v2 / (((v2 / ((v2) | 1))) | 1)) >= (-(v2)))))) {
        v2 += v2
        for fi12 := 0; fi12 < 2; fi12++ {
            delete(v1, 629)
            var v13 uint32 = 2
            var v14 uint32 = v13
            _ = fi12
            _ = v13
            _ = v14
        }
        fmt.Println(v2, v2, v2)
    } else {
        if (true && (true || (128 != 53))) {
            fmt.Println(v5, v5, v5)
        } else {
            var v15 int = -9223372036854775808
            _ = v15
        }
        v3 = (v3 - (v10 & 1000))
    }
    switch "foo bar" {
    case "foo bar":
        if (!(false && false)) {
            for fi16 := 0; fi16 < 1; fi16++ {
                for fi17 := 0; fi17 < 10; fi17++ {
                    v2 = v2
                    _ = fi17
                }
                var v18 uint32 = 476
                v19 := GenS1{F0: uint8(v18), F1: 18446744073709551615}
                _ = fi16
                _ = v18
                _ = v19
            }
        } else {
            var v20 float32 = ((0.0 + 1e10) * (0.1 / (0.1 - (((3.14159 / (3.14159 * -2.25)) - 3.14159) + (1e10 / 3.14159)))))
            v21 := GenS2{F0: 0}
            _ = v20
            _ = v21
        }
    case "a":
        if (0 != 458) {
            switch v2 {
            case 2147483647:
                v10 -= 42
            case -9223372036854775808:
                v22 := []int{(v2 << (((v11 << (((^(v11))) & 63))) & 63)), (v2 * v2), -26}
                _ = v22
            case 26:
                v5 += 100
            }
        }
    default:
        v10 = 42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
