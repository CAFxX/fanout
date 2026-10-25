package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint; F1 int16; F2 bool }
type GenS3 struct { F0 int8 }
func genFn0(p0 int16) uint64 {
    v6 := []int{}
    p0 -= (p0 | (p0 | (p0 << ((1) & 63))))
    var v7 uint = 0
    for fi8 := 0; fi8 < 9; fi8++ {
        v7 -= v7
        var v9 bool = ((p0 - (p0 + (p0 + p0))) > (^(p0)))
        _ = fi8
        _ = v9
    }
    v6 = append(v6, 170)
    _ = p0
    _ = v6
    _ = v7
    return 56
}

func main() {
    var v1 int8 = 1
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    var v4 uint64 = 0
    v5 := []string{("" + "foo bar")}
    fmt.Println(v4, v4)
    switch -9223372036854775808 {
    case -104:
        v10 := GenS1{F0: -1, F1: "\n\t"}
        _ = v10
    default:
        var v11 int8 = 0
        _ = v11
    }
    v12 := GenS2{F0: 18446744073709551615, F1: 79, F2: (false && true)}
    _ = len(v5)
    switch (("" + (("\n\t" + ("\n\t" + (("a" + (("a" + "a") + "a")) + "a"))) + (("hello" + "a") + ("hello" + ("" + "a"))))) + ("hello" + "")) {
    case "a":
        for fi13 := 0; fi13 < 12; fi13++ {
            v4 += 4294967295
            var v14 float64 = float64(1.5)
            var v15 float32 = 3.14159
            _ = fi13
            _ = v14
            _ = v15
        }
    case "foo bar":
        delete(v3, ("foo bar" + "hello"))
    }
    var v16 int = -361
    v1 -= (v1 / ((v1) | 1))
    var v17 uint32 = 190
    v1 = v1
    v18 := GenS1{F0: -9223372036854775808, F1: (("hello" + "a") + "foo bar")}
    var v19 bool = true
    v20 := GenS3{F0: v1}
    v16 += (^(v16))
    var v21 int = 1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
