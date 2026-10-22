package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 bool; F1 int8; F2 uint }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 32767
    v3 := make(map[int]bool)
    v2 += 42
    v4 := GenS1{F0: 255, F1: false}
    for fi5 := 0; fi5 < 8; fi5++ {
        switch "" {
        case "hello":
            fmt.Println(fi5)
        case "foo bar":
            v2 -= v2
        case "a":
            var v6 int = (fi5 << ((4294967295) & 63))
            _ = v6
        default:
            fmt.Println(v2, v2)
        }
        _ = fi5
    }
    var v7 uint64 = 42
    v2 = 952
    var v8 bool = false
    v2 = v2
    _ = len(v3)
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 int = (fi9 >> ((620) & 63))
        var v11 int32 = -1
        for fi12 := 0; fi12 < 10; fi12++ {
            v8 = (!false)
            _ = fi12
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v8 = ((235 != 1) && (-32768 >= 1))
    v13 := GenS2{F0: (42 == -706), F1: 80, F2: 255}
    v7 += (v7 * (v7 - (v7 & v7)))
    switch (((("a" + "") + "foo bar") + ("a" + ("\n\t" + "\n\t"))) + ((((("a" + "") + "a") + "") + "\n\t") + (("a" + (("foo bar" + (("\n\t" + ("hello" + "a")) + "")) + ("hello" + "hello"))) + "foo bar"))) {
    case "\n\t":
        delete(v3, 9223372036854775807)
    case "":
        if v8 {
            v7 += v7
        } else {
            var v14 float32 = 0.0
            var v15 int = 0
            _ = v14
            _ = v15
        }
    }
    switch -2147483648 {
    case 2147483647:
        v2 -= (v2 / ((749) | 1))
    case 128:
        v16 := v1[2147483647]
        _ = v16
    case 1:
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 bool = (0 > -726)
            var v19 uint32 = 775
            var v20 float32 = ((3.14159 - 0.0) - 0.1)
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
    }
    v21 := GenS2{F0: v8, F1: 42, F2: 42}
    if (false || (false || v8)) {
        fmt.Println(v7, v7, v7)
        var v22 int = 9223372036854775807
        _ = v22
    } else {
        var v23 uint8 = 112
        var v24 float32 = 1e10
        var v25 uint32 = 435
        _ = v23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        if v8 {
            v3[fi26] = v8
            v7 -= 758
            v27 := []int{(fi26 * fi26), fi26, (fi26 | fi26)}
            _ = v27
        }
        _ = fi26
    }
    fmt.Println(v7, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
