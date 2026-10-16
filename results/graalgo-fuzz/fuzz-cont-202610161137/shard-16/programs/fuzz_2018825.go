package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 937
    var v3 uint = (v2 ^ (v2 % (((v2 / ((v2) | 1))) | 1)))
    v4 := []int{-384, 32768, 42, 32767}
    var v5 float64 = 0.0
    v3 = (v2 / (((^(v3))) | 1))
    switch 874 {
    case 2147483647:
        if len(v4) > 0 {
            v4[0] = 32767
        }
    case 42:
        v3 = (v2 | 42)
    case 159:
        var v6 uint = (v2 / ((v2) | 1))
        _ = v6
    }
    v5 -= (((1.5 * v5) / 3.14159) * 1e10)
    v7 := GenS1{F0: 42, F1: 168}
    switch "foo bar" {
    case "hello":
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 int = fi8
            if (!(false && (4294967295 > 18446744073709551615))) {
                v10 := GenS2{F0: 724, F1: 1}
                if true {
                    fmt.Println(v2, v2)
                }
                _ = v10
            } else {
                v5 += float64(1e10)
                fmt.Println(fi8, v9)
            }
            _ = fi8
            _ = v9
        }
    case "\n\t":
        v3 = v2
    case "":
        if (255 > 42) {
            fmt.Println(v5, v3)
        } else {
            var v11 float32 = 1e10
            v12 := GenS3{F0: 1000}
            for fi13 := 0; fi13 < 5; fi13++ {
                switch ("foo bar" + ((("a" + "a") + ("a" + "a")) + "foo bar")) {
                case "\n\t":
                    v14 := GenS1{F0: 175, F1: 127}
                    _ = v14
                case "a":
                    var v15 uint16 = 374
                    _ = v15
                }
                v16 := GenS2{F0: 42, F1: 271}
                _ = fi13
                _ = v16
            }
            _ = v11
            _ = v12
        }
    default:
        fmt.Println(v3, v3)
    }
    v5 += ((0.0 - 3.14159) - -2.25)
    v3 += v3
    v2 += (v2 - v3)
    v17 := GenS1{F0: 876, F1: 128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
