package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int32 }
func genFn0(p0 uint32, p1 bool) int {
    if p1 {
        var v4 float64 = 1.5
        _ = v4
    } else {
        p0 = (p0 << ((4294967295) & 63))
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return -2147483648
}
func genFn1() uint64 {
    if (true && false) {
        var v5 uint32 = 65535
        var v6 uint16 = 42
        _ = v5
        _ = v6
    }
    if true {
        v7 := []int{genFn0(2147483648, false), -2147483648}
        _ = v7
    }
    return 1
}

func main() {
    var v1 string = ((("hello" + "foo bar") + "hello") + "foo bar")
    var v2 uint = 255
    v3 := make(map[int]uint32)
    fmt.Println(v2, v2)
    switch genFn0(42, true) {
    case 127:
        var v8 int64 = 140
        _ = v8
    }
    v9 := []int{-1, 42, 42, 0}
    v2 = (^(v2))
    if ((v2 % ((v2) | 1)) != v2) {
        fmt.Println(v1, v1)
        switch -263 {
        case 42:
            fmt.Println(v2, v1)
        }
        _ = len(v3)
    }
    var v10 int8 = 100
    var v11 float32 = 0.1
    v10 = v10
    fmt.Println(v11, v2)
    if len(v9) > 0 {
        v9[0] = int(v2)
    }
    if (!(5 == 42)) {
        fmt.Println(v1)
        for fi12 := 0; fi12 < 19; fi12++ {
            v2 -= v2
            for fi13 := 0; fi13 < 15; fi13++ {
                var v14 bool = true
                _ = fi13
                _ = v14
            }
            v15 := GenS1{F0: v1, F1: int64(v10), F2: 42}
            _ = fi12
            _ = v15
        }
        v3[71] = 65535
    } else {
        fmt.Println(v1, v11, v2)
    }
    if (2 < 2147483648) {
        var v16 int32 = 470
        _ = v16
    } else {
        var v17 int16 = 0
        _ = v17
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := GenS1{F0: ("a" + v1), F1: -9223372036854775808, F2: -446}
        v10 = v10
        _ = fi18
        _ = v19
    }
    switch "" {
    case "foo bar":
        switch 2147483647 {
        case -68:
            v10 = v10
        case 32767:
            switch v1 {
            case "\n\t":
                v2 += 255
            case "a":
                fmt.Println(v2, v11)
            case "":
                var v20 int = 2147483647
                _ = v20
            }
        case -2147483648:
            v10 = v10
        default:
            v3[4294967295] = 2147483647
        }
    case "\n\t":
        v3[1] = 0
    case "":
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 uint = v2
            _ = fi21
            _ = v22
        }
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
