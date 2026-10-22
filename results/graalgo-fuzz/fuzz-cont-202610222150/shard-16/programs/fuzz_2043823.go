package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint32; F1 int }
func genFn0(p0 uint32, p1 int32) string {
    p1 -= p1
    if ((p0 % (((p0 >> ((65535) & 63))) | 1)) < (p0 % ((302) | 1))) {
        p1 -= (^(p1))
        v6 := []int{-9223372036854775808, 471, 42, 1}
        _ = v6
    } else {
        var v7 uint64 = 4294967295
        var v8 int = 360
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    return ("" + ((("a" + "") + "") + ("\n\t" + ("a" + ""))))
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]bool)
    var v3 string = ("foo bar" + ("\n\t" + "a"))
    var v4 uint16 = 2
    var v5 bool = (!(127 < uint8(v4)))
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: v1}
        _ = fi9
        _ = v10
    }
    v3 = (("a" + (v3 + v3)) + v3)
    v11 := []int{1, 16}
    var v12 uint64 = (v1 ^ (v1 / (((v1 - (^(v1)))) | 1)))
    switch -1 {
    case -226:
        for fi13 := 0; fi13 < 3; fi13++ {
            for fi14 := 0; fi14 < 8; fi14++ {
                v12 = (v1 * v1)
                v15 := GenS2{F0: (^(v4)), F1: v3}
                _ = fi14
                _ = v15
            }
            switch 127 {
            case 0:
                var v16 int = fi13
                _ = v16
            }
            _ = fi13
        }
    case 127:
        var v17 uint8 = 60
        _ = v17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        if true {
            var v19 string = (v3 + v3)
            var v20 uint64 = v1
            fmt.Println(fi18)
            _ = v19
            _ = v20
        } else {
            fmt.Println(v12, v12)
        }
        var v21 float32 = 0.1
        _ = fi18
        _ = v21
    }
    var v22 string = v3
    var v23 uint32 = 562
    for fi24 := 0; fi24 < 14; fi24++ {
        v22 = v3
        fmt.Println(v3)
        _ = fi24
    }
    v5 = (true && v5)
    v3 = (v22 + v22)
    v25 := GenS3{F0: (v23 + (^(v23))), F1: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
