package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 float64 = (3.14159 * ((3.14159 - 0.0) * 0.1))
    v2 := []uint32{2147483648, 4294967295}
    var v3 uint16 = 32767
    v1 -= v1
    var v4 bool = (2 >= 0)
    v1 = 0.1
    v5 := GenS1{F0: "\n\t", F1: 18446744073709551615, F2: (v3 | (v3 | 32767))}
    if len(v2) > 0 {
        v2[0] = 0
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS2{F0: 0}
        v4 = false
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: "foo bar", F1: 2, F2: 2}
    var v9 string = ((("a" + ("\n\t" + "foo bar")) + "") + "\n\t")
    var v10 string = "hello"
    var v11 uint32 = 4294967295
    switch "hello" {
    case "":
        var v12 float32 = (1.5 * (float32(-2.25) + 1e10))
        _ = v12
    case "foo bar":
        _ = len(v2)
    case "a":
        v11 = 20
    }
    v13 := GenS1{F0: v9, F1: 42, F2: (v3 - v3)}
    var v14 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
