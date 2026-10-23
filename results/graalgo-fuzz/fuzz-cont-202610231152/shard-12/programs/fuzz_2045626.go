package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int32 = -1
    var v2 int16 = -370
    var v3 uint32 = 42
    var v4 uint8 = 255
    var v5 uint32 = 2147483647
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := []int{508, (fi6 % ((fi6) | 1)), 2147483647}
        _ = fi6
        _ = v7
    }
    v4 += (^(v4))
    if (42 == 1) {
        v8 := []int{32767, 32767}
        fmt.Println(v2, v5)
        _ = v8
    } else {
        var v9 uint = 547
        fmt.Println(v3, v2, v3)
        _ = v9
    }
    var v10 int64 = int64(v2)
    v11 := GenS1{F0: (v10 * (-(v10)))}
    if false {
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(v1, v1)
            _ = fi12
        }
        v3 -= v3
        var v13 string = "foo bar"
        _ = v13
    }
    switch 4294967295 {
    case 2147483647:
        v3 += 942
    case 385:
        var v14 uint = 42
        _ = v14
    }
    v15 := []int{2147483647, 780}
    for fi16 := 0; fi16 < 9; fi16++ {
        v10 = -9223372036854775808
        v10 = v10
        _ = fi16
    }
    var v17 uint64 = 4294967295
    switch 128 {
    case 9223372036854775807:
        var v18 int32 = -1
        _ = v18
    default:
        if true {
            v3 = uint32(v4)
        } else {
            v10 = v10
            v5 -= v3
        }
    }
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    fmt.Println(v10, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
