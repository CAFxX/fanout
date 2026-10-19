package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int8; F1 uint32 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]string)
    var v3 uint16 = 0
    var v4 int32 = 42
    v5 := GenS1{F0: -853, F1: v4}
    v6 := GenS2{F0: 127, F1: 2147483648}
    if (true && false) {
        v4 = v4
    } else {
        v7 := GenS1{F0: 42, F1: -1}
        _ = v7
    }
    fmt.Println(v4, v1, v3)
    switch "a" {
    case "\n\t":
        fmt.Println(v4)
    case "":
        if false {
            switch "foo bar" {
            case "":
                for fi8 := 0; fi8 < 2; fi8++ {
                    var v9 int32 = 760
                    var v10 uint16 = (v3 & v3)
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            }
            var v11 int8 = 127
            fmt.Println(v1)
            _ = v11
        }
    }
    v12 := v2[2147483648]
    switch -101 {
    case 32768:
        var v13 uint16 = v3
        _ = v13
    case 128:
        v12 = v12
    default:
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS1{F0: -1, F1: v4}
            switch (v12 + v12) {
            case "a":
                _ = len(v2)
            case "\n\t":
                var v16 uint32 = 2147483648
                _ = v16
            case "hello":
                for fi17 := 0; fi17 < 3; fi17++ {
                    var v18 uint32 = 0
                    v4 = (^(v4))
                    v19 := []int{128, 228, (fi17 + (fi17 & 9223372036854775807)), (fi14 % ((fi17) | 1))}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            default:
                if true {
                    v20 := GenS2{F0: 1, F1: 0}
                    var v21 int64 = -9223372036854775808
                    _ = v20
                    _ = v21
                } else {
                    v22 := v2[0]
                    v22 = (v22 + "")
                    _ = v22
                }
            }
            v12 = (v12 + v12)
            _ = fi14
            _ = v15
        }
    }
    fmt.Println(v12, v4, v12)
    v4 = 604
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
