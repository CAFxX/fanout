package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int16 }

func main() {
    var v1 string = ("" + "\n\t")
    v2 := []string{v1}
    var v3 bool = (9223372036854775807 <= 452)
    var v4 uint = 2
    var v5 int8 = -60
    v6 := []uint32{199, 42, 0}
    v3 = (2 != 89)
    v4 += (v4 % (((v4 / ((42) | 1))) | 1))
    fmt.Println(v4, v4, v1)
    v4 = (^(v4))
    var v7 float64 = float64(3.14159)
    fmt.Println(v4, v4)
    if (v1 != "hello") {
        switch "hello" {
        case "hello":
            _ = len(v6)
        }
    } else {
        switch -2147483648 {
        case -9223372036854775808:
            _ = len(v2)
        case 42:
            if (-1 >= 42) {
                v8 := GenS1{F0: 1, F1: (419 == 481)}
                v9 := GenS2{F0: 1}
                var v10 int = 1
                _ = v10
                _ = v8
                _ = v9
            }
        case 2147483647:
            fmt.Println(v3, v4, v1)
        default:
            var v11 uint = (^(v4))
            _ = v11
        }
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 uint16 = 240
            var v14 int32 = 0
            fmt.Println(v3)
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    v15 := GenS1{F0: 495, F1: false}
    var v16 uint32 = 2147483647
    fmt.Println(v3, v16, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
