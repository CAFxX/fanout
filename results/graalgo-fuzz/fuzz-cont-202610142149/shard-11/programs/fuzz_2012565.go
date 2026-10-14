package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int16 {
    if false {
        var v6 int8 = 1
        _ = v6
    } else {
        var v7 uint8 = 95
        var v8 float32 = -2.25
        if true {
            if (42 >= 0) {
                var v9 uint8 = (^(v7))
                _ = v9
            }
        }
        _ = v7
        _ = v8
    }
    var v10 uint64 = 388
    _ = v10
    return -1
}

func main() {
    var v1 int16 = -1
    var v2 float64 = 3.14159
    var v3 float32 = (0.1 + 0.1)
    var v4 uint = 65535
    var v5 uint32 = 4294967295
    var v11 uint32 = v5
    var v12 int = -9223372036854775808
    if false {
        switch (v12 + 128) {
        case 2147483647:
            var v13 int16 = (-(v1))
            _ = v13
        }
        v14 := GenS1{F0: 2}
        _ = v14
    }
    v15 := []int{127, (-(v12))}
    for fi16 := 0; fi16 < 7; fi16++ {
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = fi16
    }
    fmt.Println(v4, v12, v12)
    if (!true) {
        for fi17 := 0; fi17 < 1; fi17++ {
            if len(v15) > 0 {
                v15[0] = 42
            }
            var v18 int8 = 0
            fmt.Println(v4)
            _ = fi17
            _ = v18
        }
        if false {
            var v19 uint64 = 2
            var v20 int16 = v1
            var v21 uint32 = v11
            _ = v19
            _ = v20
            _ = v21
        } else {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            v22 := GenS1{F0: 4294967295}
            _ = v22
        }
        var v23 float32 = (0.0 + (3.14159 - ((1e10 / 1.5) - v3)))
        _ = v23
    }
    v2 -= 1.5
    fmt.Println(v5, v1, v12)
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 bool = (float64(v3) != v2)
        _ = fi24
        _ = v25
    }
    fmt.Println(v12)
    if false {
        var v26 int = (-(v12))
        _ = v26
    } else {
        switch 0 {
        case -9223372036854775808:
            _ = len(v15)
        default:
            v1 = genFn0()
        }
        var v27 int32 = -2147483648
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
