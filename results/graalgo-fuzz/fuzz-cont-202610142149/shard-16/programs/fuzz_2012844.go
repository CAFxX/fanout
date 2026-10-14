package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 bool; F1 uint }

func main() {
    v1 := []int{127}
    var v2 int8 = -128
    v3 := make(map[string]uint32)
    var v4 float32 = 1e10
    v5 := []string{"foo bar", "foo bar"}
    var v6 uint8 = 42
    switch "foo bar" {
    case "":
        if true {
            v4 = v4
            v4 = (1e10 - (0.1 - v4))
            if (("hello" + "foo bar") != ("foo bar" + (("a" + "hello") + (("hello" + ("" + "\n\t")) + "a")))) {
                v2 -= (v2 << ((4294967295) & 63))
                v7 := v3[("\n\t" + "\n\t")]
                _ = v7
            }
        } else {
            v4 -= 0.0
            v1 = append(v1, -646)
        }
    }
    var v8 bool = (-419 <= 42)
    switch "\n\t" {
    case "":
        v6 = (v6 * v6)
    }
    if v8 {
        var v9 string = (("a" + ("hello" + (("\n\t" + "foo bar") + "a"))) + "hello")
        fmt.Println(v9, v8, v8)
        _ = v9
    } else {
        var v10 uint = 796
        var v11 int16 = 42
        v12 := GenS1{F0: 0, F1: -700}
        _ = v10
        _ = v11
        _ = v12
    }
    fmt.Println(v2, v2, v4)
    if true {
        if (true && v8) {
            fmt.Println(v6, v2, v4)
            switch ((("hello" + "hello") + "\n\t") + "") {
            case "a":
                v13 := GenS2{F0: 42, F1: 9223372036854775807}
                _ = v13
            case "foo bar":
                v14 := GenS1{F0: -335, F1: -315}
                _ = v14
            case "":
                v15 := GenS3{F0: v8, F1: 255}
                _ = v15
            }
        }
        fmt.Println(v2, v2)
    } else {
        var v16 uint8 = (v6 + (v6 & 128))
        var v17 int = 9223372036854775807
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v16
        _ = v17
    }
    if (0 <= 255) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v18 := GenS3{F0: (true || (v6 >= (v6 | v6))), F1: 2}
    v4 += v4
    for fi19 := 0; fi19 < 10; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            v21 := GenS1{F0: 76, F1: 1}
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    var v22 uint16 = 954
    v1 = append(v1, -1)
    v4 = (float32((1e10 * 1e10)) + 1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
