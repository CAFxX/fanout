package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
func genFn0() uint {
    var v5 uint8 = 2
    var v6 float64 = 3.14159
    if false {
        for fi7 := 0; fi7 < 14; fi7++ {
            v5 = v5
            fmt.Println(fi7, v6)
            _ = fi7
        }
        v8 := GenS1{F0: 465, F1: 126}
        fmt.Println(v6, v6)
        _ = v8
    }
    _ = v5
    _ = v6
    return 808
}

func main() {
    var v1 uint8 = 42
    var v2 float64 = 1e10
    var v3 uint = 1
    var v4 int32 = 42
    if (!(!false)) {
        v9 := GenS1{F0: 2, F1: v4}
        v10 := GenS1{F0: 858, F1: -2147483648}
        _ = v10
        _ = v9
    }
    v11 := []int{-1, int(v4), -2147483648, 233}
    var v12 string = "a"
    v13 := GenS1{F0: (v3 + (v3 / ((1) | 1))), F1: v4}
    var v14 bool = true
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 uint64 = 21
        if v14 {
            v1 = (^(v1))
        }
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        if (v14 && (!(-1 > -128))) {
            var v18 bool = (v2 == 3.14159)
            var v19 bool = true
            _ = v18
            _ = v19
        } else {
            if ((v4 - (v4 ^ (v4 % ((352) | 1)))) > v4) {
                v3 += 2
            }
        }
        var v20 uint = genFn0()
        v21 := GenS1{F0: (v3 << ((genFn0()) & 63)), F1: v4}
        _ = fi17
        _ = v20
        _ = v21
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
