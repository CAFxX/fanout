package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 float32 = float32((0.1 - 1.5))
    var v2 int8 = 29
    v3 := make(map[uint32]bool)
    var v4 int8 = (-(v2))
    var v5 int = -9223372036854775808
    v1 += (v1 * 3.14159)
    var v6 uint16 = 883
    var v7 uint8 = 46
    if ((float64(v1) - 3.14159) == ((0.0 + 0.1) * 0.0)) {
        fmt.Println(v6, v4)
        v8 := GenS1{F0: (v5 & (^(v5))), F1: 42}
        _ = v8
    }
    if false {
        v9 := GenS1{F0: -2147483648, F1: 0}
        v10 := []int{0}
        _ = len(v3)
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: v5, F1: 170}
    v12 := GenS2{F0: 65535}
    switch v5 {
    case -562:
        if (1 != 0) {
            if (("\n\t" + "a") < "") {
                var v13 bool = false
                _ = v13
            } else {
                v14 := []int{(-(v5)), (v5 ^ v5), (v5 ^ v5)}
                for fi15 := 0; fi15 < 5; fi15++ {
                    switch "hello" {
                    case "":
                        v6 = v6
                    }
                    _ = len(v14)
                    fmt.Println(v6, v4, fi15)
                    _ = fi15
                }
                switch v5 {
                case 4294967295:
                    fmt.Println(v6, v1)
                case 924:
                    var v16 int32 = -959
                    _ = v16
                case 32767:
                    var v17 int64 = 1
                    _ = v17
                default:
                    _ = len(v14)
                }
                _ = v14
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
