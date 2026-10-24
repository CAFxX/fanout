package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    var v2 int = -1
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    v5 := make(map[uint32]uint32)
    v6 := []uint32{}
    if false {
        var v7 uint16 = 1
        _ = v7
    }
    if false {
        v2 = 1
        var v8 int64 = 9223372036854775807
        _ = v8
    } else {
        var v9 uint = 985
        _ = v9
    }
    fmt.Println(v2)
    _ = len(v6)
    var v10 string = "a"
    for fi11 := 0; fi11 < 4; fi11++ {
        if true {
            v10 = ("foo bar" + v10)
        }
        var v12 uint32 = 1
        _ = fi11
        _ = v12
    }
    v10 = v10
    switch v2 {
    case 552:
        fmt.Println(v10, v10)
    case 4294967295:
        v10 = v10
    case 128:
        if (65535 == 0) {
            var v13 float32 = 0.1
            var v14 float64 = ((-2.25 / 3.14159) + 0.1)
            v10 = (v10 + ("foo bar" + v10))
            _ = v13
            _ = v14
        }
    }
    switch (v10 + v10) {
    case "":
        var v15 uint32 = 0
        _ = v15
    case "\n\t":
        for fi16 := 0; fi16 < 9; fi16++ {
            v10 = ("" + v10)
            v2 -= 1
            _ = fi16
        }
    default:
        _ = len(v4)
    }
    fmt.Println(v2, v2)
    fmt.Println(v2, v2)
    v10 = "a"
    _ = len(v6)
    if true {
        v10 = v10
    }
    v2 = (v2 ^ (-(v2)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
