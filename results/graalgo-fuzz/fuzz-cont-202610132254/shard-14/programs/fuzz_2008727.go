package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint16 }
func genFn0(p0 int64, p1 bool) int32 {
    p1 = p1
    var v4 uint64 = 18446744073709551615
    p0 = (p0 - p0)
    fmt.Println(p1, p1)
    var v5 uint16 = 255
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -435
}
func genFn1(p0 int, p1 string) uint64 {
    p0 -= (-(p0))
    v6 := GenS1{F0: 18446744073709551615, F1: true, F2: 255}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v6
    return 517
}

func main() {
    var v1 int = 857
    v2 := []uint32{2147483648, 2}
    var v3 int8 = 1
    if len(v2) > 0 {
        v2[0] = 1
    }
    v2 = append(v2, 65535)
    fmt.Println(v1, v3)
    v7 := GenS1{F0: 18446744073709551615, F1: false, F2: 65535}
    var v8 int = -775
    v3 = v3
    var v9 float64 = 0.1
    if (!true) {
        v1 += (v1 + 42)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v2 = append(v2, 0)
    var v10 bool = ((((("hello" + "") + ((("foo bar" + "\n\t") + "") + "a")) + ("foo bar" + "hello")) + "hello") >= ("foo bar" + ""))
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 4; fi12++ {
            v8 = 1
            switch "a" {
            case "foo bar":
                v8 = v8
            case "a":
                v9 = 0.1
            default:
                v2 = append(v2, 42)
            }
            _ = fi12
        }
        _ = fi11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
