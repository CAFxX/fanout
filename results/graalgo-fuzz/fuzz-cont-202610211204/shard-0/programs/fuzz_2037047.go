package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int; F1 uint }

func main() {
    var v1 uint32 = 791
    var v2 uint32 = v1
    var v3 uint16 = 2
    var v4 uint16 = (v3 & v3)
    v5 := make(map[uint32]bool)
    var v6 float32 = 3.14159
    var v7 uint = 255
    v8 := GenS1{F0: 0, F1: (v1 % (((v1 % ((2147483648) | 1))) | 1))}
    v7 = v7
    fmt.Println(v7)
    switch 392 {
    case 1:
        for fi9 := 0; fi9 < 10; fi9++ {
            if true {
                _ = len(v5)
                v10 := GenS2{F0: 9223372036854775807, F1: -710}
                _ = v10
            } else {
                v6 += v6
            }
            var v11 string = "foo bar"
            var v12 int32 = 62
            _ = fi9
            _ = v11
            _ = v12
        }
    case 9223372036854775807:
        var v13 int16 = 10
        _ = v13
    case -1:
        var v14 float32 = (v6 - ((v6 * 1.5) + 0.0))
        _ = v14
    default:
        v15 := GenS3{F0: 42, F1: ((("hello" + ("\n\t" + (("\n\t" + ("foo bar" + "a")) + ((("\n\t" + "\n\t") + "\n\t") + "foo bar")))) + ("a" + "")) + ("hello" + "a"))}
        _ = v15
    }
    fmt.Println(v7, v4, v7)
    v16 := []int{-796, -298, -392, -486}
    v1 = (v2 & v2)
    v6 = (float32(0.0) - v6)
    if true {
        v17 := GenS2{F0: 0, F1: 32767}
        var v18 bool = (-1 == -6)
        v7 -= (^(v7))
        _ = v17
        _ = v18
    } else {
        if false {
            v19 := GenS1{F0: 42, F1: v1}
            var v20 float32 = ((v6 + v6) - ((0.1 - 0.1) * (v6 / (v6 + (v6 - 1e10)))))
            _ = v19
            _ = v20
        }
        v2 = 2
    }
    v21 := GenS2{F0: 9223372036854775807, F1: -32768}
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 uint32 = 0
        fmt.Println(v23, v7)
        switch (fi22 ^ 2147483647) {
        case 42:
            v6 = 3.14159
        default:
            fmt.Println(fi22, v1)
        }
        _ = fi22
        _ = v23
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    v24 := GenS4{F0: 127, F1: 876}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
