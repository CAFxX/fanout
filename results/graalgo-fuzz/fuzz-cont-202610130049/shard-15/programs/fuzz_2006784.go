package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint16; F2 int }
func genFn0(p0 uint64, p1 uint16, p2 uint16) string {
    v5 := GenS1{F0: 42, F1: false, F2: (p0 & 4294967295)}
    if (p0 > (p0 | p0)) {
        var v6 int = 4294967295
        p2 += (p2 & 460)
        _ = v6
    } else {
        for fi7 := 0; fi7 < 6; fi7++ {
            fmt.Println(p0)
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 uint8 = 127
                v10 := []int{fi8, (fi8 << ((4294967295) & 63))}
                var v11 uint32 = 768
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            _ = fi7
        }
        v12 := GenS2{F0: false, F1: (^(p1)), F2: int(p1)}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "a"
}
func genFn1(p0 int64, p1 int) int32 {
    p1 = -9223372036854775808
    fmt.Println(p0, p1)
    fmt.Println(p0, p0)
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 uint64 = 4294967295
        p1 += 1
        v15 := GenS2{F0: (!true), F1: 1, F2: (p1 >> ((0) & 63))}
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := []int{42, p1}
    _ = p0
    _ = p1
    _ = v16
    return -1
}

func main() {
    var v1 float32 = float32((-2.25 * 1e10))
    v2 := make(map[string]uint32)
    var v3 float32 = v1
    var v4 uint8 = 88
    v17 := v2[""]
    var v18 int = -2147483648
    fmt.Println(v3, v17)
    switch "hello" {
    case "":
        fmt.Println(v3)
    case "foo bar":
        v4 = v4
    case "\n\t":
        v17 += v17
    }
    var v19 uint64 = 945
    switch (v18 >> ((255) & 63)) {
    case -717:
        v18 = 9223372036854775807
    }
    var v20 int = v18
    for fi21 := 0; fi21 < 4; fi21++ {
        fmt.Println(v17, v20)
        if true {
            var v22 bool = false
            _ = v22
        } else {
            v17 = (v17 % ((1) | 1))
            v4 += 1
        }
        _ = fi21
    }
    fmt.Println(v19, v19, v3)
    v23 := []int{v20}
    var v24 int = (^(v20))
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
