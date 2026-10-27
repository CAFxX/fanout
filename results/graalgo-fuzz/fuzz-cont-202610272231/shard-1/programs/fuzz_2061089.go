package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint16 = 2
    var v2 int8 = 0
    var v3 uint = 65535
    var v4 int = 891
    var v5 bool = ((v4 % ((v4) | 1)) != (v4 * -1))
    var v6 uint32 = 2147483647
    var v7 string = ("" + "a")
    v8 := []int{-1, (^(v4)), v4, v4}
    switch (-(v4)) {
    case -9223372036854775808:
        v9 := GenS1{F0: v6}
        _ = v9
    case 573:
        var v10 float64 = 0.0
        _ = v10
    case -610:
        var v11 bool = v5
        _ = v11
    }
    var v12 uint = (v3 / ((v3) | 1))
    v8 = append(v8, (v4 << (((v12 - (v3 >> (((^(v3))) & 63)))) & 63)))
    if len(v8) > 0 {
        v8[0] = (v4 & (^(v4)))
    }
    if (((638 == 0) || v5) || (!v5)) {
        v13 := GenS2{F0: v1}
        _ = v13
    } else {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    var v14 uint16 = 568
    if true {
        for fi15 := 0; fi15 < 15; fi15++ {
            fmt.Println(fi15, v1, fi15)
            var v16 uint64 = 2
            v4 += fi15
            _ = fi15
            _ = v16
        }
    }
    var v17 float32 = (-2.25 * ((-2.25 / 1e10) - (1.5 - 0.1)))
    fmt.Println(v12, v5)
    _ = len(v8)
    for fi18 := 0; fi18 < 8; fi18++ {
        v2 -= -128
        var v19 int32 = 61
        _ = fi18
        _ = v19
    }
    var v20 float64 = 0.0
    for fi21 := 0; fi21 < 17; fi21++ {
        if ((3.14159 * ((0.0 * v20) * v20)) <= v20) {
            if true {
                if (v5 || ((v6 * (v6 ^ (v6 >> ((1) & 63)))) != v6)) {
                    _ = len(v8)
                } else {
                    v3 -= (v3 + v12)
                    var v22 uint = v12
                    v23 := GenS3{F0: 881}
                    _ = v22
                    _ = v23
                }
                _ = len(v8)
            }
        } else {
            fmt.Println(v5)
            v14 -= uint16(v3)
        }
        if (!((v14 << ((255) & 63)) >= (v1 % (((v1 | 910)) | 1)))) {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            v7 = "foo bar"
        }
        _ = fi21
    }
    var v24 int16 = -457
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
