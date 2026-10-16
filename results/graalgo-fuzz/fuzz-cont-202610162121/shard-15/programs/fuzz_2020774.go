package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint64; F1 string; F2 int }
func genFn0(p0 int32) uint64 {
    v7 := GenS1{F0: -32768, F1: 2147483647}
    v8 := GenS1{F0: 41, F1: 32768}
    if true {
        var v9 bool = ((1e10 / (1e10 * (0.1 - -2.25))) > 0.0)
        switch "" {
        case "":
            v9 = (0 > 0)
        }
        _ = v9
    }
    p0 -= (p0 ^ 2147483647)
    p0 -= p0
    _ = p0
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1() int {
    var v10 uint16 = 32767
    v11 := []int{32767, 2147483647, 0, 32767}
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := GenS2{F0: 42, F1: "foo bar", F2: int(v10)}
        var v14 bool = (0 > 2147483647)
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = v10
    _ = v11
    return 270
}

func main() {
    var v1 int16 = -32768
    var v2 int64 = -218
    var v3 bool = ((127 == 127) || false)
    v4 := make(map[string]int)
    var v5 int16 = v1
    var v6 uint64 = 2
    v15 := GenS2{F0: (^(v6)), F1: "a", F2: 32767}
    fmt.Println(v2)
    var v16 int32 = 685
    var v17 string = ""
    var v18 int = 951
    if true {
        var v19 bool = (float32((-2.25 - 3.14159)) >= 1e10)
        fmt.Println(v5, v6)
        v1 -= (v5 + (^(v1)))
        _ = v19
    }
    var v20 uint16 = 0
    switch (v17 + ("" + "foo bar")) {
    case "\n\t":
        v18 = (-(v18))
    default:
        for fi21 := 0; fi21 < 7; fi21++ {
            v22 := []int{(v18 >> ((1) & 63))}
            var v23 string = (v17 + v17)
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    var v24 int = (-(v18))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
