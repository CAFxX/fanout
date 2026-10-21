package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }

func main() {
    v1 := []int{4294967295, -2147483648}
    var v2 uint16 = 2
    var v3 float64 = 0.1
    v4 := make(map[int]uint32)
    var v5 float64 = v3
    var v6 uint = 1
    switch -949 {
    case 42:
        v7 := GenS1{F0: (v6 ^ v6), F1: 9223372036854775807}
        _ = v7
    case -582:
        fmt.Println(v5, v6, v5)
    case -2147483648:
        var v8 uint64 = 42
        _ = v8
    }
    fmt.Println(v3)
    var v9 uint64 = 2
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 int64 = -9223372036854775808
        v3 -= v3
        fmt.Println(fi10)
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        v3 = v3
        _ = fi12
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        var v13 int32 = 0
        _ = v13
    case "hello":
        v14 := GenS1{F0: v6, F1: -9223372036854775808}
        _ = v14
    default:
        v6 = v6
    }
    v6 = 65535
    var v15 int8 = 0
    v16 := GenS1{F0: 4294967295, F1: -9223372036854775808}
    v17 := GenS1{F0: 593, F1: 42}
    if false {
        for fi18 := 0; fi18 < 3; fi18++ {
            v19 := GenS1{F0: 437, F1: -1}
            _ = fi18
            _ = v19
        }
        var v20 float64 = (-2.25 + (0.1 + (v3 + v5)))
        if len(v1) > 0 {
            v1[0] = -70
        }
        _ = v20
    } else {
        v3 = 0.1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
