package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int8, p1 uint, p2 int32) uint {
    if false {
        if true {
            v4 := []int{}
            var v5 int64 = -303
            _ = v4
            _ = v5
        }
        var v6 uint16 = 1
        _ = v6
    }
    switch 42 {
    case 128:
        fmt.Println(p2, p0)
    case -1:
        p1 = (p1 << (((p1 << ((0) & 63))) & 63))
    case 318:
        v7 := []int{32767}
        _ = v7
    default:
        p0 -= (p0 * (-(p0)))
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 % ((1) | 1))
}
func genFn1(p0 string) string {
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 uint32 = 1
        _ = fi8
        _ = v9
    }
    var v10 int = 42
    var v11 int64 = 519
    _ = p0
    _ = v10
    _ = v11
    return "foo bar"
}

func main() {
    var v1 string = ""
    v2 := make(map[string]int)
    var v3 uint8 = 1
    v12 := GenS1{F0: 42}
    v13 := []int{1, 254, 0, 0}
    var v14 int32 = 42
    var v15 int = 4294967295
    if ((!true) && ((0 > 42) && true)) {
        v13 = append(v13, -75)
        var v16 int = (^(v15))
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        _ = v16
    }
    var v17 float32 = (float32(0.0) + (1.5 * -2.25))
    v13 = append(v13, (v15 * v15))
    v1 = (v1 + "\n\t")
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 uint = genFn0(-128, 513, v14)
        fmt.Println(v1, v14)
        v20 := GenS1{F0: -32768}
        _ = fi18
        _ = v19
        _ = v20
    }
    v15 = (^(v15))
    var v21 int = 32767
    var v22 uint32 = 4294967295
    v17 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
