package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int64 = -1
    var v2 int8 = 42
    var v3 uint = 87
    v4 := []string{"foo bar", "foo bar"}
    var v5 int16 = 398
    var v6 uint64 = 18446744073709551615
    var v7 uint = (^(v3))
    v2 += v2
    v4 = append(v4, "foo bar")
    var v8 int32 = 599
    var v9 float32 = -2.25
    switch 2147483647 {
    case 9223372036854775807:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 766:
        v10 := GenS1{F0: -1}
        _ = v10
    }
    _ = len(v4)
    for fi11 := 0; fi11 < 20; fi11++ {
        v9 = 0.1
        v2 = (-(v2))
        _ = fi11
    }
    v1 = (v1 - (v1 >> ((42) & 63)))
    var v12 int16 = 42
    fmt.Println(v7)
    v13 := GenS1{F0: v12}
    _ = len(v4)
    if true {
        v3 = (v7 ^ 0)
        if (!false) {
            v2 = 0
            v14 := GenS1{F0: 42}
            switch -915 {
            case -1:
                v2 += v2
            default:
                v7 = (v7 & (v3 | v7))
            }
            _ = v14
        }
    }
    var v15 int = 127
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS1{F0: 1000}
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
