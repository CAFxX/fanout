package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint64; F1 uint8 }
func genFn0() int8 {
    v6 := GenS1{F0: ("foo bar" + "foo bar"), F1: 63}
    var v7 uint32 = 4294967295
    var v8 uint = 255
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{227, 0, 2147483647}
    var v2 uint32 = 2147483648
    var v3 uint8 = 140
    var v4 float64 = -2.25
    v5 := make(map[string]string)
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := GenS1{F0: "foo bar", F1: 2}
        for fi11 := 0; fi11 < 12; fi11++ {
            fmt.Println(v2, fi9)
            _ = fi11
        }
        v12 := GenS1{F0: "a", F1: 668}
        _ = fi9
        _ = v10
        _ = v12
    }
    v3 += uint8(v2)
    v13 := GenS2{F0: 18446744073709551615, F1: v3}
    v4 -= v4
    v2 -= (v2 << ((4294967295) & 63))
    fmt.Println(v4, v3, v3)
    var v14 string = "hello"
    v3 += uint8(v2)
    switch v14 {
    case "a":
        if false {
            v4 += ((float64((0.0 / float32((float64(1e10) - float64((1e10 * 1.5)))))) - float64(1.5)) * 0.1)
            var v15 int = 1
            _ = v15
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
