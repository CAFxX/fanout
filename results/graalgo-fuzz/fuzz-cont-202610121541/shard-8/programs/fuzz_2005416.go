package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint64 }
func genFn0(p0 int64) string {
    switch -2147483648 {
    case 32767:
        p0 += (p0 | (-(p0)))
    case 2147483647:
        if (true || ((18446744073709551615 > 472) || true)) {
            var v4 float32 = float32(-2.25)
            fmt.Println(v4, v4)
            switch 32767 {
            case 42:
                fmt.Println(p0, v4)
            default:
                var v5 int = 42
                _ = v5
            }
            _ = v4
        } else {
            for fi6 := 0; fi6 < 5; fi6++ {
                fmt.Println(fi6, fi6)
                _ = fi6
            }
            p0 = p0
            var v7 float32 = -2.25
            _ = v7
        }
    case 0:
        p0 = p0
    default:
        v8 := []int{}
        _ = v8
    }
    v9 := []int{-1, 939, 42, -1}
    v10 := GenS1{F0: (p0 - p0)}
    p0 -= (p0 ^ p0)
    v11 := GenS1{F0: (p0 ^ 0)}
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return ""
}

func main() {
    var v1 uint16 = 32767
    var v2 uint16 = (v1 % ((v1) | 1))
    var v3 int = -9223372036854775808
    if true {
        v12 := GenS1{F0: 414}
        var v13 uint64 = 18446744073709551615
        _ = v12
        _ = v13
    }
    var v14 uint64 = 981
    switch "a" {
    case "hello":
        if true {
            var v15 int16 = 1
            v2 += v2
            v16 := GenS2{F0: 0, F1: (v1 + 0), F2: 809}
            _ = v15
            _ = v16
        }
    case "":
        if true {
            fmt.Println(v1)
            v14 = uint64(v2)
        } else {
            var v17 int8 = 127
            for fi18 := 0; fi18 < 2; fi18++ {
                var v19 int32 = 42
                _ = fi18
                _ = v19
            }
            _ = v17
        }
    case "a":
        v20 := GenS1{F0: -51}
        _ = v20
    }
    v14 = v14
    v21 := []int{42}
    v1 = v2
    v1 -= v2
    if len(v21) > 0 {
        v21[0] = (^(v3))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
}
