package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint16 = 0
    var v2 int8 = 127
    var v3 int64 = -9223372036854775808
    v4 := make(map[string]int)
    v5 := make(map[string]int)
    var v6 uint8 = 127
    if true {
        var v7 uint16 = (^(v1))
        fmt.Println(v2)
        _ = v7
    }
    fmt.Println(v2, v6, v6)
    for fi8 := 0; fi8 < 3; fi8++ {
        v6 = v6
        _ = fi8
    }
    v9 := GenS1{F0: 0}
    fmt.Println(v3, v2)
    v10 := GenS1{F0: 4294967295}
    v1 = 0
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := []int{fi11, fi11, (fi11 | fi11), fi11}
        var v13 float64 = (0.1 - 0.1)
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v1, v2)
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v1)
    default:
        var v14 int16 = -223
        _ = v14
    }
    delete(v4, ("" + "a"))
    v5["foo bar"] = 450
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
