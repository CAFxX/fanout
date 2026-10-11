package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int8; F1 bool; F2 string }
func genFn0() int {
    var v5 uint64 = 42
    var v6 int64 = 0
    v7 := GenS1{F0: 381, F1: (v5 - v5)}
    v8 := []int{128}
    if len(v8) > 0 {
        v8[0] = -311
    }
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 bool = true
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    v4 := []uint32{955, 450}
    var v9 uint64 = 625
    v1 = (v1 || (4294967295 == 104))
    v9 = (v9 / (((^(v9))) | 1))
    v1 = v1
    var v10 int32 = -210
    v1 = v1
    var v11 int16 = 524
    if ((32767 > 255) && v1) {
        v11 += (v11 & (v11 >> ((65535) & 63)))
        switch "a" {
        case "\n\t":
            var v12 int = -459
            _ = v12
        case "":
            var v13 int32 = (v10 << ((345) & 63))
            _ = v13
        case "a":
            switch ("foo bar" + "hello") {
            case "":
                v10 += v10
            default:
                switch 9223372036854775807 {
                case 550:
                    fmt.Println(v10, v9)
                }
            }
        }
    }
    var v14 bool = v1
    if v1 {
        fmt.Println(v11, v11, v9)
    } else {
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        if v14 {
            v15 := GenS1{F0: 255, F1: v9}
            _ = v15
        }
    }
    v16 := GenS1{F0: 906, F1: 856}
    switch "foo bar" {
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "":
        v14 = (v10 > v10)
    case "\n\t":
        if (2147483648 == 1) {
            fmt.Println(v14)
            switch 1 {
            case 711:
                if (true && false) {
                    var v17 uint64 = v9
                    v18 := GenS2{F0: -1, F1: v1, F2: "foo bar"}
                    _ = v17
                    _ = v18
                }
            case 9223372036854775807:
                v11 = v11
            default:
                v14 = (-9223372036854775808 != -9223372036854775808)
            }
        } else {
            var v19 bool = false
            var v20 int = 1
            _ = v19
            _ = v20
        }
    }
    v1 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
