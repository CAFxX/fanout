package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
func genFn0(p0 int8, p1 uint32) uint8 {
    if true {
        if false {
            p1 -= (p1 & (p1 << ((18446744073709551615) & 63)))
            var v7 int32 = 42
            var v8 int16 = 32767
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 9; fi9++ {
                fmt.Println(p1, p1)
                var v10 int64 = -9223372036854775808
                p0 = p0
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: 2, F1: 2147483647}
            _ = v11
        }
    }
    p0 += (p0 | (p0 ^ 1))
    v12 := []int{-667, 0, 0}
    _ = len(v12)
    _ = p0
    _ = p1
    _ = v12
    return 255
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float64 = (-2.25 * 0.1)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]bool)
    var v6 uint64 = (v1 & 595)
    v13 := []int{int(v6), -483, 128}
    if true {
        fmt.Println(v2)
        var v14 int16 = -32768
        var v15 int64 = 42
        _ = v14
        _ = v15
    }
    var v16 int = -1
    v2 -= 0.0
    if false {
        switch "\n\t" {
        case "hello":
            var v17 float64 = v2
            _ = v17
        default:
            for fi18 := 0; fi18 < 19; fi18++ {
                _ = len(v13)
                for fi19 := 0; fi19 < 9; fi19++ {
                    var v20 int8 = -111
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
        }
        var v21 int8 = 40
        var v22 int8 = v21
        _ = v21
        _ = v22
    }
    fmt.Println(v2)
    fmt.Println(v6, v2, v1)
    fmt.Println(v1, v6)
    var v23 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
