package main

import (
    "fmt"
)

func genFn0() uint32 {
    var v7 int32 = 265
    v8 := []int{-9223372036854775808, 42, 319, 32768}
    for fi9 := 0; fi9 < 11; fi9++ {
        v7 -= v7
        var v10 uint64 = 4294967295
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        if (0 >= 4294967295) {
            fmt.Println(v7, fi11)
            var v12 uint = 366
            var v13 int16 = 183
            _ = v12
            _ = v13
        } else {
            v7 = (v7 - -529)
        }
        _ = fi11
    }
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = -2.25
    var v3 uint16 = 491
    var v4 float64 = v2
    var v5 string = "foo bar"
    v6 := make(map[string]string)
    v4 = v2
    var v14 uint32 = 0
    v4 -= v2
    if (true && ((3.14159 >= v4) || (!(746 == 84)))) {
        if true {
            var v15 int8 = int8(v3)
            _ = v15
        }
        v3 = v3
    } else {
        switch 2147483647 {
        case -9223372036854775808:
            var v16 uint64 = 2
            _ = v16
        }
    }
    _ = len(v1)
    switch v5 {
    case "\n\t":
        var v17 string = v5
        _ = v17
    default:
        var v18 int32 = -2147483648
        _ = v18
    }
    v4 = 0.0
    v14 -= (^(v14))
    var v19 float64 = (3.14159 / v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
