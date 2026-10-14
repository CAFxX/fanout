package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int64; F2 int8 }
type GenS3 struct { F0 int16; F1 bool }
func genFn0(p0 int16, p1 uint32) int16 {
    p1 -= (p1 | 2147483648)
    switch "hello" {
    case "\n\t":
        v5 := GenS1{F0: ("foo bar" + ("a" + (("foo bar" + "a") + ("\n\t" + "a"))))}
        _ = v5
    case "":
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    return p0
}
func genFn1() int16 {
    v6 := []int{-1, 32768, 2147483647}
    var v7 int = 0
    _ = v6
    _ = v7
    return genFn0(-233, 300)
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int32 = 544
    var v3 uint32 = 601
    v4 := make(map[int]bool)
    v1 += v1
    v8 := GenS1{F0: "hello"}
    v2 -= -181
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 string = ("" + (("\n\t" + "") + ""))
        _ = fi9
        _ = v10
    }
    v11 := v4[897]
    var v12 int = 558
    switch 670 {
    case -2147483648:
        delete(v4, -285)
    case 4294967295:
        if (v11 && v11) {
            _ = len(v4)
            v13 := GenS2{F0: (-(v1)), F1: v1, F2: 85}
            _ = v13
        } else {
            var v14 int = (v12 ^ 593)
            v15 := GenS1{F0: "hello"}
            v16 := GenS3{F0: 32767, F1: (2 < 255)}
            _ = v14
            _ = v15
            _ = v16
        }
    }
    v17 := GenS2{F0: (v1 | int64(v3)), F1: (v1 & v1), F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
