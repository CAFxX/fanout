package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int16; F2 uint32 }

func main() {
    v1 := []string{}
    var v2 int8 = 100
    v3 := make(map[uint32]int)
    v4 := []uint32{42}
    v5 := make(map[int]uint32)
    var v6 string = "hello"
    fmt.Println(v2, v2)
    var v7 int64 = 0
    _ = len(v3)
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int64 = -316
        _ = fi8
        _ = v9
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v2)
    }
    var v10 uint32 = 42
    v2 += (v2 - v2)
    v7 = (v7 & 778)
    if ((1 > 233) && (true || true)) {
        v2 = -121
        v7 += (v7 | 1)
    } else {
        v11 := GenS1{F0: 414}
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        v6 = (((v6 + v6) + (v6 + "")) + ((v6 + (v6 + ("a" + "\n\t"))) + "\n\t"))
        if true {
            var v13 float32 = 0.1
            v13 += (-2.25 * 1.5)
            _ = v13
        } else {
            v2 = (v2 / ((127) | 1))
            v6 = ((v6 + v6) + (((v6 + (v6 + v6)) + ((("a" + v6) + "a") + ("" + v6))) + (((v6 + "") + v6) + (v6 + "foo bar"))))
            var v14 int8 = (v2 % ((110) | 1))
            _ = v14
        }
        _ = fi12
    }
    delete(v3, 4294967295)
    v15 := GenS2{F0: (v7 >> ((65535) & 63)), F1: -1, F2: v10}
    v6 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
