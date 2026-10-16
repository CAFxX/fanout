package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int }
func genFn0(p0 int32, p1 uint64) uint {
    v5 := []int{2147483647}
    p0 -= p0
    _ = p0
    _ = p1
    _ = v5
    return 427
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{}
    var v3 uint8 = 255
    v4 := make(map[uint32]int)
    var v6 int32 = 900
    switch -9223372036854775808 {
    case 32767:
        if (4294967295 >= 42) {
            v7 := GenS1{F0: "foo bar", F1: true, F2: 0}
            for fi8 := 0; fi8 < 11; fi8++ {
                if (2 <= (v3 * 1)) {
                    fmt.Println(fi8, fi8, v3)
                    delete(v4, 2147483648)
                    fmt.Println(v3, v3, v3)
                } else {
                    var v9 uint32 = 65535
                    var v10 int64 = -1
                    _ = v10
                    _ = v9
                }
                _ = fi8
            }
            _ = v7
        } else {
            switch "foo bar" {
            case "hello":
                if false {
                    var v11 int8 = 100
                    _ = v11
                }
            case "":
                fmt.Println(v6, v6)
            default:
                switch 32767 {
                case 1:
                    fmt.Println(v3, v3)
                case 86:
                    v6 = v6
                case 2147483647:
                    var v12 int64 = -848
                    _ = v12
                }
            }
        }
    }
    delete(v1, 127)
    v2 = append(v2, 656)
    v13 := GenS1{F0: (("\n\t" + "hello") + "hello"), F1: false, F2: 32767}
    switch 583 {
    case 32767:
        v6 += (^(v6))
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 uint16 = 653
        fmt.Println(v6, v3)
        delete(v1, fi14)
        _ = fi14
        _ = v15
    }
    v6 -= (v6 >> ((18446744073709551615) & 63))
    var v16 uint16 = 65535
    v6 += (v6 + 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
