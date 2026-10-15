package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int16, p1 int32, p2 uint64) int16 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 int64 = -906
        _ = fi5
        _ = v6
    }
    var v7 float32 = 0.0
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p0 & 32767)
}

func main() {
    var v1 uint = 776
    var v2 int64 = 0
    var v3 uint = 113
    var v4 uint = v3
    switch 1 {
    case 225:
        fmt.Println(v1)
    default:
        v1 -= (v4 >> (((^(v3))) & 63))
    }
    if true {
        v3 += 610
    }
    v4 = 42
    v8 := []int{}
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS1{F0: "foo bar"}
        _ = fi9
        _ = v10
    }
    var v11 uint64 = 18446744073709551615
    fmt.Println(v11, v1)
    v8 = append(v8, -191)
    v4 += (^(v1))
    var v12 uint16 = 1
    var v13 string = ("hello" + "hello")
    switch "hello" {
    case "a":
        var v14 string = (v13 + ((v13 + v13) + v13))
        _ = v14
    case "hello":
        v12 = (v12 * v12)
    }
    var v15 int16 = 1000
    v3 = (^(v1))
    v1 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
