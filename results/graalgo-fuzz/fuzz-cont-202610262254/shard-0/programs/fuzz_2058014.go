package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int64 }

func main() {
    v1 := []string{}
    var v2 float32 = (float32(3.14159) * (float32((1e10 - (-2.25 + 0.0))) / 1.5))
    v3 := []int{}
    v4 := []string{(("foo bar" + (("hello" + "\n\t") + "hello")) + "\n\t"), "", "", "\n\t"}
    var v5 uint = 255
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    if true {
        v2 -= ((-2.25 * v2) * float32(0.0))
    } else {
        var v6 int64 = -9223372036854775808
        _ = v6
    }
    var v7 int = 32767
    if true {
        v7 = v7
    }
    fmt.Println(v7, v2, v2)
    v3 = append(v3, v7)
    var v8 string = ("foo bar" + ("\n\t" + "foo bar"))
    fmt.Println(v7, v7, v5)
    for fi9 := 0; fi9 < 17; fi9++ {
        fmt.Println(v5, fi9)
        _ = fi9
    }
    var v10 uint64 = 1
    v4 = append(v4, "")
    fmt.Println(v8, v7, v2)
    v11 := GenS1{F0: 0, F1: v10, F2: -1}
    if true {
        switch v8 {
        case "hello":
            v8 = v8
        case "a":
            switch (v8 + (v8 + (("a" + (v8 + "hello")) + "\n\t"))) {
            case "\n\t":
                v5 += v5
            case "a":
                for fi12 := 0; fi12 < 13; fi12++ {
                    switch v8 {
                    case "":
                        var v13 uint8 = 19
                        _ = v13
                    }
                    v7 -= (fi12 >> (((v5 << (((v5 - v5)) & 63))) & 63))
                    _ = fi12
                }
            case "":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                var v14 uint = v5
                _ = v14
            }
        case "\n\t":
            v10 -= (v10 & v10)
        default:
            var v15 int = v7
            _ = v15
        }
        var v16 float64 = 0.1
        _ = v16
    } else {
        var v17 int8 = -8
        _ = len(v3)
        var v18 int32 = 42
        _ = v17
        _ = v18
    }
    v7 += v7
    var v19 int16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
