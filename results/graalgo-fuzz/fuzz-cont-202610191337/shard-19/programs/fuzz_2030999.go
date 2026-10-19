package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint64; F1 string; F2 uint16 }

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    var v4 int8 = 81
    v5 := GenS1{F0: -201, F1: 42}
    switch "" {
    case "":
        v4 -= v4
    case "a":
        v6 := GenS1{F0: 1, F1: 2}
        _ = v6
    default:
        var v7 uint64 = 31
        _ = v7
    }
    v1 = v1
    v1 = (v1 ^ (v1 | (v1 / ((931) | 1))))
    v1 = (v1 >> ((18446744073709551615) & 63))
    for fi8 := 0; fi8 < 3; fi8++ {
        switch (fi8 % (((v1 & 1)) | 1)) {
        case -2147483648:
            var v9 uint = 65535
            _ = v9
        case -1:
            fmt.Println(v1)
        case -355:
            switch "\n\t" {
            case "hello":
                switch (v1 / (((-(fi8))) | 1)) {
                case 9223372036854775807:
                    var v10 uint16 = 23
                    _ = v10
                default:
                    var v11 string = "hello"
                    _ = v11
                }
            default:
                v4 = (^(v4))
            }
        }
        if (346 <= 636) {
            v4 = v4
            delete(v3, fi8)
            if (1e10 == 3.14159) {
                fmt.Println(v1)
            }
        } else {
            v12 := GenS2{F0: 2, F1: ("hello" + "hello"), F2: 974}
            _ = v12
        }
        v13 := []int{(v1 % ((v1) | 1)), v1, (v1 >> ((1) & 63)), (v1 / ((fi8) | 1))}
        _ = fi8
        _ = v13
    }
    fmt.Println(v1, v1)
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := []int{1}
        v1 += v1
        switch "" {
        case "hello":
            v1 = v1
        case "foo bar":
            var v16 uint8 = 255
            _ = v16
        case "":
            delete(v3, (fi14 | fi14))
        }
        _ = fi14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
