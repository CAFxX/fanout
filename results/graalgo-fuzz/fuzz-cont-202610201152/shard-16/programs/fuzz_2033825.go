package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 string; F2 uint8 }
type GenS3 struct { F0 int8 }

func main() {
    v1 := []string{("a" + ("foo bar" + "hello")), "foo bar", ((("foo bar" + "") + ("foo bar" + "foo bar")) + "foo bar"), "foo bar"}
    var v2 uint64 = 0
    v3 := make(map[string]int)
    var v4 int16 = -1
    v5 := []int{32767, 736}
    var v6 int8 = -98
    v3["\n\t"] = 32768
    v7 := GenS1{F0: v2}
    var v8 int = 2147483647
    for fi9 := 0; fi9 < 10; fi9++ {
        delete(v3, ("a" + ""))
        v2 = (v2 ^ (v2 / ((1) | 1)))
        v6 -= 1
        _ = fi9
    }
    var v10 bool = false
    var v11 string = "\n\t"
    var v12 float32 = -2.25
    v13 := GenS2{F0: v4, F1: v11, F2: 91}
    switch ((((v11 + (v11 + (("\n\t" + v11) + v11))) + "hello") + (v11 + v11)) + ("foo bar" + (v11 + v11))) {
    case "a":
        if len(v5) > 0 {
            v5[0] = 32767
        }
    case "\n\t":
        var v14 int16 = (v4 * v4)
        _ = v14
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if (1 < 68) {
        v8 += (^(v8))
    } else {
        fmt.Println(v11)
        fmt.Println(v2, v12, v2)
    }
    if v10 {
        if v10 {
            v15 := GenS3{F0: 113}
            _ = v15
        } else {
            var v16 uint = 0
            v17 := GenS3{F0: v6}
            fmt.Println(v16, v12, v11)
            _ = v16
            _ = v17
        }
        if (1 > (v4 >> ((739) & 63))) {
            var v18 float64 = (1.5 - 1e10)
            var v19 int32 = 661
            var v20 float64 = v18
            _ = v18
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 18; fi21++ {
            v22 := GenS2{F0: int16(v2), F1: "", F2: uint8(v2)}
            _ = fi21
            _ = v22
        }
    }
    fmt.Println(v2, v8, v12)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
