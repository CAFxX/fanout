package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint32 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]string)
    var v3 bool = false
    var v4 bool = (!false)
    var v5 uint8 = 0
    var v6 float64 = float64(1e10)
    fmt.Println(v5, v6)
    if false {
        v5 -= v5
        var v7 float64 = 1e10
        _ = v7
    } else {
        for fi8 := 0; fi8 < 19; fi8++ {
            for fi9 := 0; fi9 < 20; fi9++ {
                v6 -= (v6 * 3.14159)
                var v10 int = fi9
                fmt.Println(v10, v6)
                _ = fi9
                _ = v10
            }
            var v11 bool = v4
            _ = fi8
            _ = v11
        }
    }
    v6 -= 0.1
    v3 = v4
    if v3 {
        var v12 string = "hello"
        v13 := GenS1{F0: 255, F1: -1, F2: 2147483648}
        v14 := []int{-9223372036854775808, 939}
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 int = 42
    switch 1 {
    case -2147483648:
        v6 = -2.25
    case 9223372036854775807:
        v3 = (((0 > 281) || v4) && v3)
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        for fi17 := 0; fi17 < 4; fi17++ {
            delete(v1, fi16)
            _ = fi17
        }
        var v18 uint64 = 4294967295
        _ = fi16
        _ = v18
    }
    var v19 float32 = 0.1
    v5 -= v5
    switch "" {
    case "":
        fmt.Println(v15, v15, v15)
    case "\n\t":
        if (-427 != 533) {
            var v20 uint32 = 42
            v21 := GenS1{F0: 4294967295, F1: 32767, F2: 1}
            _ = v20
            _ = v21
        }
    case "a":
        v15 = -142
    }
    v4 = (false && ((v5 - (^(v5))) >= 2))
    v22 := []int{(v15 - v15), v15, v15, v15}
    switch (v15 + (v15 % ((-9223372036854775808) | 1))) {
    case -1:
        if (2 != 0) {
            var v23 bool = (!true)
            _ = v23
        }
    case 580:
        var v24 int16 = -864
        _ = v24
    case 375:
        for fi25 := 0; fi25 < 10; fi25++ {
            switch "" {
            case "a":
                v26 := v1[v15]
                _ = v26
            case "hello":
                v27 := v1[877]
                _ = v27
            case "\n\t":
                var v28 bool = v3
                _ = v28
            }
            _ = fi25
        }
    default:
        var v29 float32 = -2.25
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
