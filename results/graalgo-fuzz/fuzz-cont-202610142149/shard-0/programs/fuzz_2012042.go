package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 int16; F2 string }
func genFn0(p0 uint64, p1 int32, p2 string) string {
    var v7 uint64 = (p0 ^ 18446744073709551615)
    var v8 uint32 = 2147483648
    if false {
        v7 += (^(v7))
        v9 := GenS1{F0: 127}
        var v10 uint = 18446744073709551615
        _ = v10
        _ = v9
    } else {
        v11 := []int{}
        fmt.Println(v8, v7, v7)
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        v13 := GenS2{F0: 32767}
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ("\n\t" + p2)
}

func main() {
    var v1 int = 1
    var v2 int32 = 2147483647
    var v3 uint64 = 18446744073709551615
    var v4 uint64 = 4294967295
    var v5 uint8 = 222
    v6 := []int{-708, -542, (^(v1))}
    for fi14 := 0; fi14 < 14; fi14++ {
        fmt.Println(v4, v5)
        v15 := GenS2{F0: 1000}
        v16 := GenS1{F0: v5}
        _ = fi14
        _ = v15
        _ = v16
    }
    if false {
        _ = len(v6)
        v3 = v4
        v4 = (v3 * v4)
    }
    v17 := GenS3{F0: 722, F1: -32768, F2: "foo bar"}
    v18 := GenS3{F0: -32768, F1: 789, F2: ""}
    var v19 bool = (!true)
    v20 := GenS2{F0: -32768}
    v2 = 1
    v1 += (v1 % ((v1) | 1))
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 bool = false
        var v23 uint16 = 2
        _ = len(v6)
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 int8 = 42
        var v26 uint16 = 255
        v26 = 255
        _ = fi24
        _ = v25
        _ = v26
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v19 = false
    v3 += v3
    v27 := GenS2{F0: 42}
    var v28 float32 = (float32(3.14159) - 0.1)
    v19 = (-128 <= 127)
    var v29 bool = ((v1 << ((1) & 63)) > 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
