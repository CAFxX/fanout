package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint64; F1 int16 }
func genFn0() uint32 {
    var v4 int = 1
    var v5 int64 = 42
    switch (("foo bar" + (("foo bar" + ("a" + "\n\t")) + (((("hello" + "foo bar") + "") + "hello") + "foo bar"))) + "foo bar") {
    case "foo bar":
        var v6 int32 = -1
        _ = v6
    case "a":
        v5 = (v5 & v5)
    }
    v7 := []int{v4}
    v8 := GenS1{F0: -32768, F1: (0 != 18446744073709551615)}
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    return 164
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{32767, 266}
    var v3 float32 = float32(0.0)
    var v9 int32 = 128
    for fi10 := 0; fi10 < 19; fi10++ {
        v3 += (((v3 + float32(((0.1 / 1.0) / 1.0))) - v3) + (float32((float64(((v3 + ((((v3 * v3) + v3) - 3.14159) - (float32(((-2.25 / 1.5) + 1e10)) - (0.0 * (3.14159 + 1.5))))) - ((-2.25 - ((v3 * v3) / float32(1e10))) + float32((0.0 + (0.0 * 3.14159)))))) / 1.0)) * v3))
        _ = fi10
    }
    var v11 float32 = v3
    for fi12 := 0; fi12 < 12; fi12++ {
        if (true || (!true)) {
            v2 = append(v2, fi12)
        }
        _ = fi12
    }
    v13 := GenS1{F0: 329, F1: (313 < 0)}
    var v14 float32 = (((v11 + 1.5) + (v11 / (v3 / (v11 * (v11 - v3))))) / v3)
    fmt.Println(v9)
    v15 := GenS2{F0: 42, F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v9
}
