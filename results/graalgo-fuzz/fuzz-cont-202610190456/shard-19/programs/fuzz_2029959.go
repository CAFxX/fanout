package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint64 }

func main() {
    var v1 uint8 = 140
    var v2 float32 = -2.25
    var v3 float64 = 0.0
    var v4 float32 = 0.0
    var v5 int32 = -354
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{}
        for fi8 := 0; fi8 < 4; fi8++ {
            fmt.Println(fi8, v2, v4)
            fmt.Println(v5, v5)
            v5 = 2147483647
            v2 -= ((v2 - 0.1) / v2)
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    v9 := []int{}
    v10 := GenS1{F0: 42, F1: -14, F2: 2}
    var v11 int32 = (v5 ^ (v5 * -2147483648))
    for fi12 := 0; fi12 < 5; fi12++ {
        if true {
            fmt.Println(v5)
        }
        _ = fi12
    }
    var v13 int64 = -9223372036854775808
    var v14 int = 9223372036854775807
    v15 := GenS1{F0: 70, F1: 778, F2: 784}
    if (1 >= 42) {
        v4 = v4
        v16 := GenS1{F0: 100, F1: 0, F2: 1}
        _ = v16
    } else {
        if false {
            fmt.Println(v1, v14, v1)
        } else {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            fmt.Println(v11, v13)
        }
        switch v14 {
        case 32768:
            var v17 uint64 = 4294967295
            _ = v17
        case -9223372036854775808:
            switch (v14 / ((v14) | 1)) {
            case 127:
                for fi18 := 0; fi18 < 19; fi18++ {
                    v5 -= (v5 - (v5 >> ((uint(v1)) & 63)))
                    fmt.Println(v3, v1)
                    var v19 string = ("a" + "a")
                    _ = fi18
                    _ = v19
                }
            case 172:
                v5 += v5
            }
        case -1:
            var v20 uint = 2
            _ = v20
        }
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
