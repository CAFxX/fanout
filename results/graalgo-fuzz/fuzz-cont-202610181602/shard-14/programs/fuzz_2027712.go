package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 string }

func main() {
    var v1 bool = (42 == 100)
    var v2 int = 2147483647
    var v3 uint16 = 255
    var v4 string = "\n\t"
    var v5 float64 = 0.0
    v6 := GenS1{F0: v3, F1: v3}
    v7 := GenS2{F0: "foo bar"}
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int64 = 335
        _ = fi8
        _ = v9
    }
    v10 := []int{v2}
    fmt.Println(v4)
    v4 = "a"
    var v11 bool = (v1 || (v4 == v4))
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(v1, v5, fi12)
        v3 += v3
        if len(v10) > 0 {
            v10[0] = -9223372036854775808
        }
        _ = fi12
    }
    fmt.Println(v4)
    for fi13 := 0; fi13 < 14; fi13++ {
        if v11 {
            var v14 int32 = 1
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            fmt.Println(fi13, fi13, fi13)
            _ = v14
        } else {
            _ = len(v10)
        }
        for fi15 := 0; fi15 < 6; fi15++ {
            switch (fi13 | (^(fi13))) {
            case -908:
                var v16 int16 = -1
                _ = v16
            case 127:
                fmt.Println(v1, v5, v2)
            case -979:
                fmt.Println(fi13)
            default:
                var v17 uint = 0
                _ = v17
            }
            var v18 uint8 = 0
            _ = fi15
            _ = v18
        }
        _ = fi13
    }
    fmt.Println(v4)
    if len(v10) > 0 {
        v10[0] = -1
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        v2 = fi19
        var v20 uint8 = 127
        var v21 float64 = v5
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS2{F0: v4}
    v2 = (v2 % (((v2 * v2)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
