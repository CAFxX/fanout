package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int8; F2 int32 }

func main() {
    v1 := []uint32{2147483647}
    var v2 int64 = 42
    v3 := []int{-1, -2147483648}
    var v4 float32 = 3.14159
    if true {
        v3 = append(v3, 0)
        fmt.Println(v2)
        v5 := GenS1{F0: "a"}
        _ = v5
    } else {
        var v6 uint8 = 128
        _ = v6
    }
    switch ("foo bar" + ("" + "hello")) {
    case "foo bar":
        var v7 uint = 42
        _ = v7
    case "hello":
        v8 := GenS2{F0: ("\n\t" + ""), F1: 100, F2: 2147483647}
        _ = v8
    }
    fmt.Println(v4, v2)
    if true {
        var v9 uint8 = 44
        var v10 int = -1
        v11 := GenS2{F0: "\n\t", F1: -128, F2: 1}
        _ = v10
        _ = v11
        _ = v9
    }
    if len(v3) > 0 {
        v3[0] = -2147483648
    }
    if ("" <= "\n\t") {
        v4 = v4
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(v4)
            for fi13 := 0; fi13 < 3; fi13++ {
                if ("hello" <= "\n\t") {
                    switch -9223372036854775808 {
                    case 32768:
                        switch ("\n\t" + "a") {
                        case "":
                            v2 = (^(v2))
                        default:
                            var v14 uint32 = 442
                            _ = v14
                        }
                    case -9223372036854775808:
                        var v15 uint8 = 25
                        _ = v15
                    default:
                        v1 = append(v1, 1)
                    }
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi13
            }
            _ = fi12
        }
        fmt.Println(v4)
    } else {
        var v16 float64 = 3.14159
        _ = v16
    }
    v2 = v2
    if len(v1) > 0 {
        v1[0] = 829
    }
    var v17 float64 = float64(float32(float64(1.5)))
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 int64 = (v2 ^ (v2 - v2))
        var v20 int64 = (v19 % ((9223372036854775807) | 1))
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: "a"}
    v4 -= v4
    v17 -= v17
    var v22 float32 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
