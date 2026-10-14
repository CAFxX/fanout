package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int32, p1 int64, p2 uint32) uint8 {
    switch 127 {
    case 666:
        fmt.Println(p1, p1)
    case 128:
        for fi5 := 0; fi5 < 4; fi5++ {
            p1 -= (-(p1))
            _ = fi5
        }
    case -434:
        var v6 int8 = 0
        _ = v6
    }
    v7 := GenS1{F0: 0, F1: p2}
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 float32 = 1.5
        var v10 string = "foo bar"
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 128
}
func genFn1(p0 uint32, p1 uint, p2 bool) int16 {
    var v11 uint8 = 71
    v12 := []int{-395, 4294967295, 32768, 330}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 32767
}

func main() {
    var v1 float64 = (1.5 * 0.1)
    v2 := []uint32{}
    var v3 uint16 = 1
    var v4 uint = 0
    v13 := GenS1{F0: 1, F1: 2147483647}
    fmt.Println(v4, v4, v3)
    v3 = (v3 << (((v4 << (((v4 ^ 0)) & 63))) & 63))
    var v14 string = (("a" + (("hello" + (("hello" + "\n\t") + ("foo bar" + "a"))) + ("foo bar" + ""))) + "foo bar")
    var v15 uint16 = v3
    v14 = (("\n\t" + v14) + "a")
    var v16 bool = false
    if v16 {
        var v17 uint64 = 0
        var v18 uint64 = (v17 | (^(v17)))
        _ = v17
        _ = v18
    }
    var v19 int8 = 110
    _ = len(v2)
    v20 := GenS2{F0: 808}
    v2 = append(v2, 1)
    for fi21 := 0; fi21 < 13; fi21++ {
        fmt.Println(fi21, v16)
        if len(v2) > 0 {
            v2[0] = 42
        }
        v19 += -4
        _ = fi21
    }
    if (42 != 42) {
        v22 := GenS2{F0: 1}
        fmt.Println(v1, v14)
        v3 = 65535
        _ = v22
    }
    if (!v16) {
        var v23 uint16 = (v15 + (^(v15)))
        v2 = append(v2, 4294967295)
        _ = v23
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    v15 -= (v15 + 65535)
    var v24 int8 = v19
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
}
