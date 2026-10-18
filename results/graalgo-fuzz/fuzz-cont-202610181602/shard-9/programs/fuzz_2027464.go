package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int64 }

func main() {
    var v1 uint8 = 0
    var v2 uint64 = 795
    var v3 int64 = int64(v2)
    v4 := make(map[string]uint32)
    var v5 uint16 = 778
    v6 := make(map[int]bool)
    if (-1 == 0) {
        v7 := GenS1{F0: v1, F1: 42, F2: (v3 * (v3 | -731))}
        _ = v7
    } else {
        delete(v4, "a")
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        v5 -= v5
        v9 := []int{fi8, fi8}
        if false {
            v10 := GenS1{F0: (v1 & (^(v1))), F1: 18446744073709551615, F2: (v3 >> ((18446744073709551615) & 63))}
            _ = v10
        } else {
            switch 948 {
            case 127:
                fmt.Println(fi8, v5, v3)
            case -2147483648:
                v3 -= v3
            default:
                if (true && (2 == v2)) {
                    v11 := GenS1{F0: (v1 % (((v1 | (v1 << ((0) & 63)))) | 1)), F1: 65535, F2: (v3 >> ((4294967295) & 63))}
                    _ = v11
                }
            }
            var v12 int8 = 127
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    v3 = (-(v3))
    if (4294967295 >= -911) {
        v3 = v3
    } else {
        switch "a" {
        case "\n\t":
            for fi13 := 0; fi13 < 17; fi13++ {
                if ((false && (v3 < v3)) || ((1.5 != (-2.25 - (-2.25 / -2.25))) && false)) {
                    v2 = v2
                }
                _ = fi13
            }
        case "a":
            v14 := GenS1{F0: v1, F1: 65535, F2: v3}
            _ = v14
        case "":
            switch "\n\t" {
            case "":
                switch 32768 {
                case 128:
                    if (2 <= v5) {
                        fmt.Println(v3)
                    }
                case 1:
                    var v15 string = ""
                    _ = v15
                }
            default:
                fmt.Println(v2, v5, v3)
            }
        }
        v1 = (v1 % ((0) | 1))
        v2 -= (v2 ^ (v2 | (v2 + 0)))
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 int64 = (v3 + (v3 / ((v3) | 1)))
        _ = fi16
        _ = v17
    }
    fmt.Println(v2)
    v18 := GenS1{F0: 154, F1: 447, F2: (v3 - 42)}
    v2 -= (v2 + (v2 / ((v2) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
