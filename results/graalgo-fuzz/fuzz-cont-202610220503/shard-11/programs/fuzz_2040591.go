package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int) int16 {
    v6 := []int{p0, p0, (-(p0)), 32768}
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 float32 = 0.1
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: ""}
    _ = p0
    _ = v6
    _ = v9
    return -32768
}

func main() {
    var v1 bool = false
    var v2 int16 = 32767
    var v3 float32 = (((float32(((-2.25 / 1.0) / (0.1 - 1e10))) - 3.14159) + 0.1) / (((float32(-2.25) * (3.14159 - 1e10)) + ((1.5 / (3.14159 / -2.25)) / (1e10 * -2.25))) - ((((3.14159 / 1e10) / 1e10) * 3.14159) * 1e10)))
    var v4 int16 = v2
    var v5 int16 = (v4 % ((v2) | 1))
    if v1 {
        var v10 int64 = -9223372036854775808
        v11 := GenS1{F0: ("hello" + "foo bar")}
        fmt.Println(v1)
        _ = v10
        _ = v11
    }
    var v12 uint = 52
    v13 := GenS1{F0: "\n\t"}
    v5 = (^(v2))
    v14 := GenS1{F0: "\n\t"}
    v15 := []int{127}
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 int64 = 743
        v2 = v5
        _ = fi16
        _ = v17
    }
    var v18 int8 = 100
    fmt.Println(v18, v3)
    v4 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
