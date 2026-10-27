package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int16 }

func main() {
    v1 := []int{2147483647, 2147483647}
    var v2 uint32 = 42
    var v3 uint64 = 4294967295
    var v4 uint16 = 2
    if (false && (true || true)) {
        v3 = v3
        var v5 string = ("" + (("\n\t" + ("\n\t" + "a")) + ((("foo bar" + ("a" + "")) + (("\n\t" + "") + ("foo bar" + "\n\t"))) + "")))
        _ = v5
    } else {
        v4 -= (^(v4))
        v6 := GenS1{F0: (v4 >> ((724) & 63)), F1: -388, F2: 32767}
        _ = v6
    }
    var v7 uint = 255
    var v8 int8 = 120
    if (0.1 != 0.0) {
        if len(v1) > 0 {
            v1[0] = 32768
        }
    }
    var v9 bool = true
    fmt.Println(v8, v8, v4)
    if (true && true) {
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 float64 = -2.25
            v12 := GenS1{F0: (v4 * v4), F1: 0, F2: int16(v4)}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        if len(v1) > 0 {
            v1[0] = -876
        }
        v1 = append(v1, int(v3))
    }
    fmt.Println(v3, v3, v9)
    v2 = v2
    v9 = ((^(v2)) < 42)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v8 += (-(v8))
    v7 = v7
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
