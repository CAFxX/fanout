package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int64 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]bool)
    v3 := []int{325, -672, -708, 0}
    var v4 string = ("hello" + "foo bar")
    v1 += v1
    v4 = "foo bar"
    if (v4 != ((v4 + "foo bar") + (((v4 + ("a" + (("foo bar" + (v4 + (v4 + v4))) + v4))) + "") + ("" + v4)))) {
        var v5 int = 819
        v4 = (((("" + v4) + v4) + "") + "")
        for fi6 := 0; fi6 < 16; fi6++ {
            v5 -= 127
            for fi7 := 0; fi7 < 3; fi7++ {
                v8 := GenS1{F0: 65535, F1: v4, F2: -531}
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
        _ = v5
    }
    if (2147483647 >= -2147483648) {
        if ((!(1.5 >= -2.25)) && true) {
            var v9 int8 = 42
            _ = v9
        }
    }
    v2[v4] = (((0.0 + -2.25) - float64(0.0)) != 0.0)
    if (-9223372036854775808 <= 128) {
        var v10 uint16 = 1
        _ = v10
    }
    v4 = "a"
    v4 = (v4 + ("a" + v4))
    v3 = append(v3, 4294967295)
    v1 += (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
