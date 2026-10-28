package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint16, p1 int8, p2 uint) int8 {
    var v5 string = ("hello" + "")
    var v6 bool = (0 >= 42)
    var v7 int64 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return p1
}

func main() {
    var v1 int32 = 1
    var v2 string = "foo bar"
    var v3 float64 = (0.1 - ((0.1 * 1e10) / (-2.25 + 1.5)))
    var v4 int = 4294967295
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: -413}
        v2 = "a"
        _ = fi8
        _ = v9
    }
    v1 = -1
    fmt.Println(v2)
    v1 = (^(v1))
    v10 := GenS1{F0: 0}
    v4 -= (v4 * v4)
    v4 += 1
    v1 -= v1
    v3 -= 0.0
    for fi11 := 0; fi11 < 10; fi11++ {
        fmt.Println(v2, v1, v1)
        _ = fi11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
