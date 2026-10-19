package main

import (
    "fmt"
)

func genFn0() int16 {
    if false {
        var v6 uint64 = 0
        fmt.Println(v6)
        _ = v6
    }
    var v7 uint = 65535
    if false {
        var v8 bool = (!(42 <= 962))
        v9 := []int{-1, -2147483648, 128, 59}
        _ = v8
        _ = v9
    } else {
        fmt.Println(v7)
        var v10 int16 = 526
        _ = v10
    }
    _ = v7
    return -32768
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = (42 >= 1)
    var v3 uint32 = 65535
    var v4 int16 = -381
    var v5 float64 = 1e10
    var v11 int = -10
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := v1[571]
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        v4 = v4
        v1[v11] = 32767
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 int = v11
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    var v17 int64 = -1
    var v18 uint32 = 2147483648
    var v19 int8 = -128
    v20 := []int{}
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
