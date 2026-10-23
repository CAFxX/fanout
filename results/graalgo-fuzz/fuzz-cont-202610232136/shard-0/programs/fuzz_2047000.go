package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int32 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int16 = 42
    var v2 float64 = 1e10
    var v3 uint16 = 616
    v4 := make(map[uint32]string)
    var v5 uint16 = 897
    var v6 uint8 = 128
    switch 32768 {
    case 127:
        if ("foo bar" < "a") {
            var v7 int64 = 42
            fmt.Println(v7)
            var v8 uint16 = (v5 / (((v5 ^ (v3 + 2))) | 1))
            _ = v7
            _ = v8
        } else {
            v9 := []int{32767, 0, -2147483648, -2147483648}
            _ = v9
        }
    default:
        var v10 string = "\n\t"
        _ = v10
    }
    switch (("\n\t" + (("a" + ("\n\t" + (("\n\t" + "") + ""))) + ((("hello" + "a") + "\n\t") + "hello"))) + "hello") {
    case "\n\t":
        var v11 float64 = v2
        _ = v11
    default:
        v6 = 231
    }
    if true {
        v6 -= 128
        v5 = (v5 * (v5 - v5))
    } else {
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := GenS1{F0: int16(v6), F1: 160, F2: -1}
            _ = fi12
            _ = v13
        }
        var v14 string = "foo bar"
        _ = v14
    }
    var v15 string = (("foo bar" + "a") + "foo bar")
    v6 = (v6 * v6)
    v2 = v2
    v2 = -2.25
    if (4294967295 == 42) {
        v16 := GenS2{F0: -128}
        var v17 float64 = (v2 - v2)
        delete(v4, 761)
        _ = v16
        _ = v17
    }
    v18 := []int{-254, -9223372036854775808}
    var v19 int16 = v1
    if true {
        fmt.Println(v3, v5, v6)
        v20 := GenS2{F0: 127}
        _ = v20
    }
    var v21 int8 = 16
    _ = len(v4)
    fmt.Println(v15, v3)
    delete(v4, 302)
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    _ = len(v18)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
