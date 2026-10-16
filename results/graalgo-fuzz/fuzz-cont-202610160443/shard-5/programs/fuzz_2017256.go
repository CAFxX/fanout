package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int64; F1 int32 }
func genFn0() uint8 {
    for fi5 := 0; fi5 < 5; fi5++ {
        if (0.0 == 1e10) {
            var v6 int16 = 1000
            v7 := GenS1{F0: v6, F1: 1}
            _ = v6
            _ = v7
        }
        _ = fi5
    }
    var v8 int = 42
    var v9 bool = false
    v8 += v8
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 string = ("a" + ((((("a" + "") + "a") + ((("\n\t" + "hello") + "\n\t") + "hello")) + "a") + ""))
    var v2 int8 = -1
    var v3 int = 9223372036854775807
    var v4 float64 = (1e10 + 3.14159)
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int8 = -39
        var v12 float32 = (0.1 - 0.0)
        v2 = (v11 / ((v2) | 1))
        _ = fi10
        _ = v11
        _ = v12
    }
    v2 = 70
    switch v1 {
    case "\n\t":
        v4 = v4
    case "foo bar":
        v1 = "foo bar"
    case "hello":
        v13 := []int{v3, -1, 32768}
        _ = v13
    default:
        var v14 float32 = 0.1
        _ = v14
    }
    v15 := []int{}
    for fi16 := 0; fi16 < 6; fi16++ {
        fmt.Println(v4, v3)
        fmt.Println(v3, v1, v4)
        _ = fi16
    }
    v15 = append(v15, 2147483647)
    switch v3 {
    case 2147483647:
        fmt.Println(v2)
    case 32767:
        v1 = (("a" + ((v1 + "foo bar") + v1)) + "a")
    }
    switch "" {
    case "":
        if true {
            var v17 float64 = (-2.25 + (float64(3.14159) * float64(1.5)))
            v18 := GenS2{F0: -1, F1: -2147483648}
            var v19 float32 = 1.5
            _ = v17
            _ = v18
            _ = v19
        } else {
            var v20 int64 = 0
            if len(v15) > 0 {
                v15[0] = 32767
            }
            _ = v20
        }
    case "a":
        for fi21 := 0; fi21 < 2; fi21++ {
            fmt.Println(v1, v3, v1)
            _ = fi21
        }
    case "foo bar":
        v2 += (v2 & (v2 & (v2 - (v2 | v2))))
    default:
        v22 := GenS1{F0: 32767, F1: 462}
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
