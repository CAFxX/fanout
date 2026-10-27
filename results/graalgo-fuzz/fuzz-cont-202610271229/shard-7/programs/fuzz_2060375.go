package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint }
func genFn0() int {
    v5 := GenS1{F0: -32768, F1: 18446744073709551615}
    switch "\n\t" {
    case "foo bar":
        var v6 float64 = ((1.5 / 1.5) * 3.14159)
        _ = v6
    case "a":
        var v7 uint32 = 42
        _ = v7
    case "hello":
        if (!(("" + (("foo bar" + ("foo bar" + ("foo bar" + "foo bar"))) + ("\n\t" + ("a" + "hello")))) > "hello")) {
            var v8 int = -807
            _ = v8
        }
    default:
        var v9 int64 = 1
        _ = v9
    }
    switch 87 {
    case 429:
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 float32 = float32(1e10)
            _ = fi10
            _ = v11
        }
    }
    var v12 uint16 = 1
    v13 := []int{128, -85}
    _ = v12
    _ = v13
    _ = v5
    return 2147483647
}

func main() {
    var v1 uint16 = 2
    var v2 float32 = 1e10
    v3 := []string{"", ""}
    var v4 uint64 = 4294967295
    var v14 int8 = -1
    var v15 uint64 = 280
    for fi16 := 0; fi16 < 11; fi16++ {
        v4 -= (^(v4))
        _ = fi16
    }
    var v17 int16 = -1
    v2 = v2
    var v18 bool = false
    _ = len(v3)
    var v19 int64 = 0
    v20 := GenS2{F0: (v14 > 100), F1: 164}
    if false {
        if (2 < 4294967295) {
            var v21 uint16 = (v1 | v1)
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            v22 := GenS2{F0: (-1 >= v19), F1: 23}
            _ = v21
            _ = v22
        } else {
            for fi23 := 0; fi23 < 7; fi23++ {
                var v24 int16 = v17
                _ = fi23
                _ = v24
            }
        }
    }
    var v25 uint64 = 2
    v4 = (^(v25))
    if false {
        fmt.Println(v15, v25)
        _ = len(v3)
        fmt.Println(v15, v25)
    } else {
        var v26 uint32 = 42
        _ = v26
    }
    v15 = (^(v15))
    var v27 int8 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v27
    _ = v3
    _ = v4
}
