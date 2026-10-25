package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int32; F1 int }
func genFn0() uint64 {
    var v6 uint32 = 2147483647
    fmt.Println(v6)
    v7 := []int{-601}
    switch -2147483648 {
    case 822:
        fmt.Println(v6)
    }
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int, p1 int8) uint32 {
    p0 = (p0 % ((p0) | 1))
    var v8 uint = 65535
    v9 := GenS1{F0: v8, F1: uint8(v8), F2: 0}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 int8 = 0
    var v2 int16 = 1000
    var v3 uint32 = 2147483647
    var v4 uint = 493
    var v5 uint16 = 758
    v5 += 65535
    fmt.Println(v2, v3, v1)
    fmt.Println(v2)
    fmt.Println(v5, v4)
    v10 := []int{-423}
    if len(v10) > 0 {
        v10[0] = -135
    }
    switch (("foo bar" + "a") + "\n\t") {
    case "a":
        if len(v10) > 0 {
            v10[0] = 978
        }
    default:
        for fi11 := 0; fi11 < 11; fi11++ {
            if len(v10) > 0 {
                v10[0] = (fi11 | (fi11 * fi11))
            }
            switch "hello" {
            case "\n\t":
                var v12 int64 = 42
                _ = v12
            case "":
                v3 = v3
            case "a":
                var v13 uint64 = 2
                _ = v13
            default:
                switch "hello" {
                case "hello":
                    v14 := GenS2{F0: 42, F1: (fi11 >> ((v4) & 63))}
                    _ = v14
                case "\n\t":
                    fmt.Println(v4, v1)
                default:
                    v15 := GenS2{F0: 2147483647, F1: fi11}
                    _ = v15
                }
            }
            _ = fi11
        }
    }
    var v16 uint32 = (v3 + v3)
    fmt.Println(v16, v4, v5)
    v1 -= v1
    if len(v10) > 0 {
        v10[0] = 32768
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        v5 += v5
        fmt.Println(fi17, v16, v3)
        var v18 int8 = -1
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: -467, F1: 4294967295}
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
