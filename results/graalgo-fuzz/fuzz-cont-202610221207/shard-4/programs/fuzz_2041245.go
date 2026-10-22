package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 int32 = -587
    var v3 string = "\n\t"
    v4 := make(map[int]int)
    var v5 int64 = 994
    v6 := make(map[string]uint32)
    var v7 uint16 = 2
    v8 := v6[v3]
    fmt.Println(v8, v2)
    _ = len(v1)
    v5 -= (v5 ^ v5)
    var v9 float64 = 0.0
    _ = len(v6)
    switch 128 {
    case 42:
        if ((v3 + (((((v3 + (v3 + (v3 + v3))) + "a") + v3) + (v3 + (v3 + v3))) + ((v3 + (("hello" + (v3 + v3)) + (v3 + (v3 + v3)))) + ("\n\t" + v3)))) <= ("hello" + "foo bar")) {
            v10 := GenS1{F0: v2, F1: (v7 * 1), F2: (v8 / ((4294967295) | 1))}
            _ = v10
        } else {
            v7 -= (v7 / (((v7 / ((42) | 1))) | 1))
        }
    case 486:
        var v11 uint = 65535
        _ = v11
    case -772:
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 uint64 = 18446744073709551615
            _ = fi12
            _ = v13
        }
    }
    switch (("\n\t" + v3) + ("a" + v3)) {
    case "":
        v7 += 0
    }
    fmt.Println(v8, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
