package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int64 }
func genFn0(p0 int16, p1 int32, p2 bool) int64 {
    var v5 int = -60
    v6 := GenS1{F0: 90}
    if (("foo bar" + "hello") > "foo bar") {
        v7 := GenS1{F0: 9223372036854775807}
        if (!p2) {
            fmt.Println(v5, v5, p2)
        } else {
            var v8 int8 = 0
            switch "a" {
            case "a":
                var v9 uint64 = 995
                _ = v9
            case "":
                p0 -= (-(p0))
            }
            var v10 uint16 = 1
            _ = v10
            _ = v8
        }
        v11 := GenS2{F0: p0, F1: 1}
        _ = v11
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -1
}

func main() {
    v1 := []uint32{42, 2147483648}
    v2 := make(map[uint32]bool)
    var v3 int16 = 433
    var v4 float32 = (1.5 + -2.25)
    fmt.Println(v4, v4)
    fmt.Println(v3, v4)
    if (295 <= 42) {
        if true {
            v3 += v3
        }
        v2[1] = (!((v3 - 481) <= v3))
    }
    v12 := GenS1{F0: 69}
    v13 := GenS2{F0: v3, F1: -9223372036854775808}
    v3 += v3
    _ = len(v2)
    var v14 uint = 18446744073709551615
    fmt.Println(v4, v4)
    v15 := GenS1{F0: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
