package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
func genFn0() uint {
    if true {
        var v4 uint32 = 946
        switch "hello" {
        case "hello":
            v4 -= 2
        }
        var v5 uint16 = 911
        _ = v4
        _ = v5
    }
    switch 32768 {
    case 9223372036854775807:
        var v6 int = 967
        _ = v6
    default:
        if true {
            var v7 string = "hello"
            _ = v7
        } else {
            var v8 bool = (2 >= 127)
            _ = v8
        }
    }
    return 0
}

func main() {
    var v1 float64 = 1.5
    var v2 uint = 18446744073709551615
    var v3 uint64 = 68
    fmt.Println(v3, v1)
    switch "a" {
    case "foo bar":
        v3 = (v3 | v3)
    default:
        if true {
            v9 := GenS1{F0: v3, F1: v3}
            fmt.Println(v1, v3)
            v10 := []int{9223372036854775807, 957, 180, 89}
            _ = v10
            _ = v9
        }
    }
    var v11 int32 = -2147483648
    var v12 int32 = v11
    v3 -= v3
    v13 := []int{1}
    for fi14 := 0; fi14 < 8; fi14++ {
        _ = len(v13)
        var v15 int16 = -1
        _ = fi14
        _ = v15
    }
    switch "a" {
    case "hello":
        var v16 int = 0
        _ = v16
    case "a":
        v3 -= (^(v3))
    }
    v2 = genFn0()
    if false {
        v13 = append(v13, 2147483647)
        v17 := GenS1{F0: v3, F1: v3}
        v18 := GenS1{F0: 831, F1: v3}
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 3; fi19++ {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = fi19
        }
    }
    var v20 int64 = -9223372036854775808
    v11 -= (v12 >> ((genFn0()) & 63))
    if len(v13) > 0 {
        v13[0] = 697
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
}
