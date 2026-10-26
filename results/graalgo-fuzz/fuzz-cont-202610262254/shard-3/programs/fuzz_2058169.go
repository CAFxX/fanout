package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint; F1 int8; F2 uint }

func main() {
    var v1 int = 128
    var v2 int16 = 42
    var v3 int16 = (v2 + v2)
    var v4 int = (^(v1))
    v5 := GenS1{F0: 0, F1: v1}
    var v6 string = "\n\t"
    var v7 uint64 = 283
    v8 := GenS2{F0: 4294967295, F1: -83, F2: 18446744073709551615}
    v3 += (v3 & 32767)
    v9 := []int{v4}
    if (!(!true)) {
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 string = ("" + ((v6 + v6) + "hello"))
            switch ((v11 + (v11 + v6)) + "a") {
            case "foo bar":
                var v12 uint16 = 1
                _ = v12
            case "hello":
                _ = len(v9)
            case "\n\t":
                v6 = "a"
            default:
                fmt.Println(v11, v1, v4)
            }
            _ = fi10
            _ = v11
        }
        var v13 uint16 = 42
        var v14 int32 = -410
        _ = v13
        _ = v14
    }
    switch "hello" {
    case "\n\t":
        v15 := GenS2{F0: 4294967295, F1: 1, F2: 255}
        _ = v15
    case "hello":
        for fi16 := 0; fi16 < 1; fi16++ {
            v3 -= (-(v3))
            for fi17 := 0; fi17 < 2; fi17++ {
                if true {
                    v7 = (v7 * v7)
                    var v18 int8 = 106
                    v1 = int(v18)
                    _ = v18
                }
                v9 = append(v9, v1)
                _ = fi17
            }
            var v19 float32 = 3.14159
            _ = fi16
            _ = v19
        }
    default:
        v20 := GenS2{F0: 2, F1: -128, F2: 255}
        _ = v20
    }
    v21 := GenS2{F0: 999, F1: 74, F2: 1}
    v22 := GenS2{F0: 42, F1: 127, F2: 2}
    fmt.Println(v7, v3, v6)
    v3 += (v2 * v3)
    v23 := GenS1{F0: 1, F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
