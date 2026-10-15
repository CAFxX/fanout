package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
func genFn0(p0 int, p1 int, p2 uint16) int8 {
    fmt.Println(p0, p2)
    fmt.Println(p0, p0, p2)
    p2 = (p2 % ((p2) | 1))
    v6 := []int{128, -97, p0}
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(p1, p0)
        p0 = p1
        fmt.Println(p0, p0, p0)
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -128
}

func main() {
    v1 := []string{"\n\t", ("a" + (("\n\t" + ("foo bar" + "foo bar")) + ""))}
    var v2 int32 = 0
    v3 := make(map[string]bool)
    var v4 uint64 = 632
    v5 := make(map[string]uint32)
    var v8 int16 = -1
    if false {
        switch (((("hello" + "\n\t") + "foo bar") + "\n\t") + ("a" + "\n\t")) {
        case "hello":
            for fi9 := 0; fi9 < 4; fi9++ {
                fmt.Println(v4)
                v10 := GenS1{F0: 100, F1: 0}
                _ = fi9
                _ = v10
            }
        case "\n\t":
            switch 2147483647 {
            case 42:
                for fi11 := 0; fi11 < 9; fi11++ {
                    v4 += (v4 & (v4 / ((v4) | 1)))
                    _ = fi11
                }
            case -2147483648:
                v2 = v2
            default:
                var v12 float32 = (3.14159 + 1e10)
                _ = v12
            }
        case "foo bar":
            v4 = v4
        default:
            var v13 int32 = -1
            _ = v13
        }
    }
    v2 -= 2147483647
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 uint32 = 2
        _ = fi14
        _ = v15
    }
    switch 2147483647 {
    case 127:
        var v16 uint = 18446744073709551615
        _ = v16
    case -9223372036854775808:
        v2 += (v2 * -1)
    case -981:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        var v17 int = 42
        _ = v17
    }
    _ = len(v5)
    v4 = v4
    v2 -= 0
    v18 := v5["foo bar"]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
