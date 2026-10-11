package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int; F1 int16; F2 bool }
func genFn0() int32 {
    var v4 string = ("foo bar" + "foo bar")
    var v5 int8 = 0
    fmt.Println(v5)
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 4294967295
    var v3 int16 = 32767
    v6 := []int{-2147483648, 0, 0, -1}
    var v7 bool = true
    v8 := GenS1{F0: v2, F1: 569, F2: 2}
    v9 := GenS1{F0: (^(v2)), F1: 1, F2: 255}
    if (0 != -623) {
        v7 = (18446744073709551615 <= 65535)
    }
    var v10 float32 = (3.14159 / ((1.5 * 0.1) / 1.0))
    if v7 {
        v2 = v2
        switch ("" + ("hello" + "hello")) {
        case "hello":
            var v11 string = ("\n\t" + "foo bar")
            _ = v11
        default:
            if v7 {
                fmt.Println(v2)
                _ = len(v6)
            }
        }
        v12 := GenS2{F0: -518, F1: v3, F2: v7}
        _ = v12
    } else {
        switch (("" + "hello") + "hello") {
        case "foo bar":
            v7 = v7
        case "":
            var v13 uint16 = 32767
            _ = v13
        default:
            var v14 int8 = int8(v3)
            _ = v14
        }
        v15 := GenS2{F0: 32767, F1: (v3 ^ v3), F2: (v7 && (474 > 2))}
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 float32 = v10
        var v18 uint16 = 65535
        _ = fi16
        _ = v17
        _ = v18
    }
    v10 -= 0.0
    v7 = false
    var v19 int64 = 256
    for fi20 := 0; fi20 < 13; fi20++ {
        fmt.Println(v19, v7, v10)
        _ = fi20
    }
    var v21 int16 = (v3 * v3)
    v1["a"] = v7
    var v22 uint32 = 0
    var v23 int8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
