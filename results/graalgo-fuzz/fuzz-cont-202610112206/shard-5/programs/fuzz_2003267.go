package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int64 }
type GenS2 struct { F0 uint; F1 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    var v3 int16 = 690
    var v4 uint8 = 1
    var v5 int = -943
    fmt.Println(v4, v5, v5)
    var v6 uint = 2
    v7 := []int{v5, v5, (v5 / ((-2147483648) | 1))}
    v5 = int(v6)
    v5 = v5
    v5 -= (v5 * (v5 % ((4294967295) | 1)))
    for fi8 := 0; fi8 < 13; fi8++ {
        if false {
            var v9 bool = true
            _ = v9
        }
        v6 = 1
        v10 := GenS1{F0: v3, F1: v6, F2: -5}
        _ = fi8
        _ = v10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int32 = 2147483647
        for fi13 := 0; fi13 < 16; fi13++ {
            v6 = v6
            var v14 int8 = 1
            var v15 uint64 = 49
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    if true {
        v16 := GenS1{F0: v3, F1: (v6 ^ (v6 & (v6 % ((v6) | 1)))), F2: -769}
        _ = v16
    } else {
        var v17 float32 = -2.25
        fmt.Println(v4)
        _ = v17
    }
    v18 := GenS2{F0: v6, F1: v3}
    var v19 uint64 = 2
    v20 := GenS2{F0: v6, F1: v3}
    switch (v5 / ((-9223372036854775808) | 1)) {
    case -766:
        var v21 string = ("foo bar" + "")
        _ = v21
    }
    var v22 bool = false
    var v23 uint32 = 2
    delete(v2, 32767)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
