package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 string; F2 int8 }
type GenS5 struct { F0 bool; F1 int8 }

func main() {
    var v1 uint8 = 225
    var v2 uint8 = 224
    var v3 int32 = 1
    v4 := []int{245, 4294967295, 0, -2147483648}
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            v2 -= v2
            v3 = v3
            v3 -= 2147483647
            _ = fi6
        }
        v2 += (v1 / ((2) | 1))
        _ = fi5
    }
    var v7 int8 = -46
    v8 := GenS1{F0: (^(v2)), F1: ("" + "")}
    fmt.Println(v2)
    v7 = 1
    v9 := GenS1{F0: 96, F1: "foo bar"}
    v10 := GenS2{F0: true}
    v2 = v2
    if true {
        var v11 string = ("a" + "hello")
        v4 = append(v4, -9223372036854775808)
        fmt.Println(v11, v7)
        _ = v11
    } else {
        v12 := GenS3{F0: 4294967295}
        v13 := GenS4{F0: 707, F1: ("hello" + "\n\t"), F2: -52}
        _ = v12
        _ = v13
    }
    fmt.Println(v3)
    v3 -= (^(v3))
    fmt.Println(v3)
    v14 := GenS5{F0: (((true || (true && true)) || true) || true), F1: 100}
    var v15 uint16 = 0
    if (((1.5 / (3.14159 + 1.5)) > ((1e10 + (float32(1e10) * 3.14159)) - 1e10)) || (((("foo bar" + "a") + (("a" + (("hello" + ("hello" + "hello")) + (("a" + "hello") + "\n\t"))) + "a")) + "\n\t") > (("" + "a") + ("" + "")))) {
        var v16 int16 = 42
        _ = v16
    }
    var v17 bool = true
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
