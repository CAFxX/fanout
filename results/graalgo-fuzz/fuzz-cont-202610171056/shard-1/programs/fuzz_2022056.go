package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
func genFn0() int8 {
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: 104, F1: 100}
        switch (^(fi6)) {
        case -2147483648:
            v8 := []int{0, (fi6 / (((fi6 | fi6)) | 1)), fi6}
            _ = v8
        case -1:
            v9 := []int{(fi6 % ((fi6) | 1)), (fi6 >> ((2) & 63)), fi6}
            _ = v9
        case 4294967295:
            var v10 int8 = 73
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    v11 := GenS1{F0: 128, F1: -105}
    var v12 uint32 = 42
    _ = v11
    _ = v12
    return 53
}

func main() {
    var v1 string = ("a" + ("hello" + "hello"))
    v2 := []int{-324}
    var v3 uint = 384
    var v4 float64 = 3.14159
    var v5 int32 = -521
    switch v1 {
    case "\n\t":
        v3 += (v3 << (((^(v3))) & 63))
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v13 string = (v1 + "foo bar")
    if true {
        var v14 int16 = 1000
        _ = v14
    } else {
        for fi15 := 0; fi15 < 17; fi15++ {
            v5 = (v5 & v5)
            var v16 uint16 = 156
            _ = fi15
            _ = v16
        }
        if false {
            v1 = ("hello" + v13)
        }
    }
    if (42 != 1) {
        var v17 int16 = -1
        fmt.Println(v3, v13, v3)
        _ = v17
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v4 = (((float64((-2.25 * 1.5)) + 3.14159) / (v4 / 1e10)) - ((-2.25 + ((float64(1e10) / v4) / 0.1)) * float64(1e10)))
    }
    fmt.Println(v5, v3)
    var v18 int8 = -1
    v19 := GenS1{F0: 2, F1: 127}
    for fi20 := 0; fi20 < 10; fi20++ {
        var v21 int64 = 9223372036854775807
        _ = fi20
        _ = v21
    }
    v5 -= (v5 - (v5 * v5))
    v18 -= (^(v18))
    v22 := GenS1{F0: 42, F1: 0}
    var v23 float64 = 3.14159
    var v24 uint32 = 203
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
