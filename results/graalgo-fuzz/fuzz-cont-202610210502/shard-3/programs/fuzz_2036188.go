package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint }
func genFn0(p0 int, p1 uint16) int16 {
    var v4 uint64 = 18446744073709551615
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 15; fi6++ {
            fmt.Println(p0, fi5)
            _ = fi6
        }
        switch p0 {
        case -547:
            switch ("hello" + "a") {
            case "\n\t":
                p0 = fi5
            case "a":
                if ((v4 & 517) == 0) {
                    fmt.Println(fi5, p1, v4)
                } else {
                    if (0 < 9223372036854775807) {
                        var v7 string = "hello"
                        v8 := GenS1{F0: uint(v4), F1: 9223372036854775807, F2: 0}
                        var v9 int8 = 1
                        _ = v7
                        _ = v8
                        _ = v9
                    }
                }
            case "hello":
                v10 := []int{p0, 128, p0, (p0 & p0)}
                _ = v10
            default:
                p0 = (p0 << ((4294967295) & 63))
            }
        case -150:
            var v11 uint64 = 42
            _ = v11
        case 128:
            var v12 uint16 = (^(p1))
            _ = v12
        default:
            v13 := []int{p0}
            _ = v13
        }
        _ = fi5
    }
    var v14 float64 = 1.5
    switch "" {
    case "a":
        var v15 uint16 = (p1 & (p1 >> ((1) & 63)))
        _ = v15
    default:
        var v16 int8 = 100
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v14
    _ = v4
    return -59
}

func main() {
    var v1 bool = (true && (302 == 548))
    v2 := make(map[uint32]int)
    var v3 float64 = 0.0
    v3 = -2.25
    v17 := GenS1{F0: 109, F1: -613, F2: 807}
    v3 = (1e10 + v3)
    var v18 uint64 = 719
    for fi19 := 0; fi19 < 10; fi19++ {
        v20 := GenS1{F0: 1, F1: 42, F2: 314}
        _ = fi19
        _ = v20
    }
    var v21 string = "hello"
    delete(v2, 920)
    fmt.Println(v21)
    fmt.Println(v21, v18, v1)
    fmt.Println(v18)
    v1 = v1
    switch "\n\t" {
    case "":
        v1 = true
    case "\n\t":
        switch (v21 + (v21 + (("a" + v21) + v21))) {
        case "a":
            v22 := []int{503, 32767}
            _ = v22
        case "":
            var v23 int16 = -32768
            _ = v23
        default:
            _ = len(v2)
        }
    case "foo bar":
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 int16 = -1
            var v26 int64 = -730
            if false {
                var v27 float64 = (v3 * (1.5 * 1e10))
                _ = v27
            }
            _ = fi24
            _ = v25
            _ = v26
        }
    default:
        if (!("" != "a")) {
            var v28 string = "\n\t"
            _ = v28
        }
    }
    v29 := []int{-2147483648, -9223372036854775808, 42}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v29
    _ = v3
}
