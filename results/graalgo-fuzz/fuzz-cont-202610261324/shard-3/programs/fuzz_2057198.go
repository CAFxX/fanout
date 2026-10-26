package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int64; F1 int64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -1
    var v3 uint8 = 0
    var v4 float32 = (0.1 / 3.14159)
    v5 := GenS1{F0: -2147483648}
    switch -9223372036854775808 {
    case 128:
        switch 127 {
        case 0:
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                v3 += v3
            case "a":
                for fi6 := 0; fi6 < 9; fi6++ {
                    v3 = v3
                    switch -9223372036854775808 {
                    case 42:
                        v4 -= ((v4 + 1.5) - (float32(float64(((v4 / v4) * 0.1))) * v4))
                    }
                    v7 := []int{(fi6 << ((42) & 63))}
                    _ = fi6
                    _ = v7
                }
            default:
                v8 := GenS2{F0: 322, F1: (306 <= 9223372036854775807)}
                _ = v8
            }
        case 2147483647:
            v9 := GenS3{F0: (v2 << ((98) & 63)), F1: v2}
            _ = v9
        default:
            switch 4294967295 {
            case -9223372036854775808:
                var v10 uint64 = 4294967295
                _ = v10
            case 127:
                if true {
                    var v11 string = ("foo bar" + "\n\t")
                    _ = v11
                } else {
                    v12 := []int{127, -2147483648, 1}
                    fmt.Println(v3, v3, v2)
                    _ = v12
                }
            default:
                v13 := []int{-841, 1}
                _ = v13
            }
        }
    case -1:
        var v14 int8 = -127
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        delete(v1, "\n\t")
        _ = fi15
    }
    var v16 int64 = -1
    var v17 float32 = (v4 * (v4 * v4))
    v18 := []int{-1}
    if true {
        _ = len(v18)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
