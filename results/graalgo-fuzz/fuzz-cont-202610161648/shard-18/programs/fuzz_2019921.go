package main

import (
    "fmt"
)

func genFn0() int32 {
    if false {
        var v6 int64 = 0
        _ = v6
    } else {
        var v7 string = (("foo bar" + "hello") + "hello")
        _ = v7
    }
    var v8 float32 = 1e10
    _ = v8
    return -27
}

func main() {
    var v1 int = 682
    var v2 int64 = 42
    v3 := make(map[uint32]int)
    var v4 int64 = (v2 | -996)
    v5 := make(map[uint32]uint32)
    v2 = v2
    var v9 int32 = genFn0()
    if (42 != 109) {
        switch (v1 + v1) {
        case -9223372036854775808:
            var v10 int16 = -70
            _ = v10
        case 0:
            var v11 uint = 65535
            _ = v11
        default:
            fmt.Println(v4, v1)
        }
        var v12 uint = 1
        _ = v12
    }
    var v13 float64 = 1.5
    for fi14 := 0; fi14 < 11; fi14++ {
        fmt.Println(v13)
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 uint64 = 292
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    var v17 int32 = v9
    fmt.Println(v2, v2, v4)
    var v18 uint64 = 1
    v1 = (v1 & -421)
    var v19 int8 = 100
    switch v1 {
    case 4294967295:
        v17 = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
