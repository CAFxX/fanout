package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint16) int {
    p0 = (^(p0))
    var v5 float32 = ((1e10 - (3.14159 + (float32((3.14159 + 0.0)) * (1.5 - ((1.5 / 0.1) - (0.0 * (((3.14159 + 1.5) * float32(0.0)) - (float32((float64((1e10 - 3.14159)) + 0.1)) / float32((1e10 / 0.1)))))))))) - (0.0 * (3.14159 + 3.14159)))
    v6 := GenS1{F0: false}
    _ = p0
    _ = v5
    _ = v6
    return -1
}
func genFn1() uint16 {
    var v7 uint16 = 249
    v8 := []int{128, -1}
    v8 = append(v8, 2147483647)
    fmt.Println(v7)
    _ = v7
    _ = v8
    return 32767
}

func main() {
    v1 := []int{-369, 4294967295}
    var v2 int = 9223372036854775807
    var v3 uint32 = 2147483647
    var v4 float32 = 1.5
    if len(v1) > 0 {
        v1[0] = (v2 % (((v2 / ((165) | 1))) | 1))
    }
    if true {
        v9 := GenS2{F0: v3, F1: 42}
        v3 += v3
        _ = v9
    }
    _ = len(v1)
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 uint64 = 995
            var v13 string = "a"
            v1 = append(v1, (fi10 & fi10))
            v4 = v4
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    switch (("\n\t" + "\n\t") + "a") {
    case "foo bar":
        var v14 int64 = -9223372036854775808
        _ = v14
    }
    var v15 uint8 = 53
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 int8 = 127
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        v3 = 2147483648
        _ = fi18
    }
    switch "hello" {
    case "foo bar":
        _ = len(v1)
    case "\n\t":
        if ("a" >= "hello") {
            var v19 uint16 = 568
            _ = v19
        }
    case "a":
        v15 -= (v15 >> ((0) & 63))
    }
    v3 -= 2147483647
    switch "" {
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = genFn0(543)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
