package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 string }
func genFn0(p0 int32, p1 uint64, p2 string) uint16 {
    var v5 uint64 = 42
    if (!true) {
        v6 := []int{9223372036854775807, 0}
        v5 = (^(p1))
        _ = v6
    } else {
        var v7 int64 = -1
        _ = v7
    }
    v8 := GenS1{F0: 42, F1: 1000, F2: p2}
    switch 9223372036854775807 {
    case 364:
        p1 = (p1 | 1)
    default:
        var v9 bool = true
        _ = v9
    }
    v10 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    _ = v8
    return 0
}

func main() {
    var v1 int16 = 628
    var v2 int64 = 886
    var v3 int8 = 1
    v4 := make(map[uint32]string)
    v11 := []int{-168}
    fmt.Println(v2, v2)
    var v12 uint16 = 987
    switch 128 {
    case 127:
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    case 32768:
        for fi13 := 0; fi13 < 14; fi13++ {
            switch (fi13 % ((fi13) | 1)) {
            case 128:
                if false {
                    for fi14 := 0; fi14 < 10; fi14++ {
                        var v15 bool = (v1 >= -32768)
                        switch ("" + "") {
                        case "a":
                            fmt.Println(fi14, v12, v12)
                        case "\n\t":
                            v16 := v4[694]
                            _ = v16
                        case "foo bar":
                            v3 = (v3 >> ((4294967295) & 63))
                        }
                        var v17 uint64 = 101
                        _ = fi14
                        _ = v15
                        _ = v17
                    }
                }
            case -9223372036854775808:
                var v18 float64 = 0.0
                _ = v18
            case 2147483647:
                delete(v4, 0)
            default:
                var v19 uint = 2
                _ = v19
            }
            _ = fi13
        }
    case 128:
        var v20 bool = ("a" >= (("hello" + "foo bar") + ("" + "a")))
        _ = v20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        fmt.Println(v1, v2)
        fmt.Println(fi21, v1, v3)
        v22 := GenS1{F0: v1, F1: (-(v1)), F2: "hello"}
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: 1, F1: v1, F2: "\n\t"}
    if ((!false) || true) {
        var v24 int16 = (v1 + -1)
        _ = v24
    } else {
        var v25 float64 = 1e10
        _ = v25
    }
    var v26 int = 0
    var v27 uint8 = 253
    var v28 int = v26
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
