package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint }
func genFn0() uint16 {
    var v7 uint16 = 154
    v7 -= v7
    v7 = (v7 + v7)
    _ = v7
    return v7
}

func main() {
    var v1 uint = 255
    var v2 bool = ((("a" + "foo bar") + "hello") <= "foo bar")
    var v3 uint32 = 2147483647
    v4 := make(map[int]bool)
    var v5 uint16 = 718
    var v6 int32 = 42
    switch 9223372036854775807 {
    case 32768:
        for fi8 := 0; fi8 < 15; fi8++ {
            fmt.Println(v6, v3)
            var v9 int16 = -32768
            _ = len(v4)
            _ = fi8
            _ = v9
        }
    default:
        var v10 int = 32767
        _ = v10
    }
    v3 = (v3 ^ 0)
    switch 180 {
    case -9223372036854775808:
        v6 += v6
    }
    var v11 float32 = float32(1.5)
    v2 = (-1 >= 703)
    if (!false) {
        v12 := []int{-2147483648, -427, 336}
        _ = v12
    }
    v3 -= v3
    v13 := GenS1{F0: 93, F1: (v6 - (v6 | (^(v6)))), F2: (^(v1))}
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        for fi14 := 0; fi14 < 3; fi14++ {
            v15 := GenS1{F0: 2, F1: v6, F2: v1}
            _ = fi14
            _ = v15
        }
    default:
        for fi16 := 0; fi16 < 14; fi16++ {
            if v2 {
                var v17 uint32 = 2147483647
                fmt.Println(v1, fi16, fi16)
                _ = v17
            }
            var v18 uint8 = 1
            switch "foo bar" {
            case "hello":
                var v19 uint = (v1 & v1)
                _ = v19
            case "foo bar":
                v20 := GenS1{F0: v18, F1: v6, F2: v1}
                _ = v20
            case "\n\t":
                v21 := []int{(fi16 << (((v1 - v1)) & 63)), int(v18)}
                _ = v21
            }
            _ = fi16
            _ = v18
        }
    }
    fmt.Println(v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
