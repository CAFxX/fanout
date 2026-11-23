package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
func genFn0() int32 {
    v6 := []int{127, 32767}
    if false {
        var v7 float32 = (1.5 / (1.5 / (1.5 / 0.1)))
        if len(v6) > 0 {
            v6[0] = 42
        }
        _ = v7
    } else {
        switch 756 {
        case 127:
            var v8 int32 = 612
            _ = v8
        case -2147483648:
            v9 := GenS1{F0: 32767, F1: -2147483648}
            _ = v9
        }
    }
    var v10 uint = 112
    v11 := GenS1{F0: 42, F1: 0}
    var v12 uint16 = 42
    _ = v10
    _ = v11
    _ = v12
    _ = v6
    return 908
}
func genFn1() uint16 {
    switch "a" {
    case "":
        var v13 uint16 = 42
        _ = v13
    }
    switch -9223372036854775808 {
    case 42:
        var v14 string = "foo bar"
        _ = v14
    case 923:
        for fi15 := 0; fi15 < 15; fi15++ {
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 int16 = 32767
                v17 = v17
                fmt.Println(fi16, v17)
                fmt.Println(fi16, fi15)
                _ = fi16
                _ = v17
            }
            var v18 uint8 = 1
            var v19 int8 = int8(fi15)
            _ = fi15
            _ = v18
            _ = v19
        }
    }
    return 1
}

func main() {
    var v1 float32 = (-2.25 / 0.1)
    var v2 uint64 = 42
    var v3 int8 = 127
    var v4 int32 = 0
    var v5 uint16 = 253
    var v20 string = ("\n\t" + "")
    var v21 int64 = -9223372036854775808
    v22 := GenS1{F0: -32768, F1: v4}
    v23 := GenS1{F0: 42, F1: (v4 | (v4 ^ (v4 + v4)))}
    fmt.Println(v2, v3, v5)
    for fi24 := 0; fi24 < 18; fi24++ {
        if true {
            switch v20 {
            case "":
                if false {
                    v3 -= v3
                    var v25 uint16 = v5
                    v20 = ("a" + v20)
                    _ = v25
                }
            default:
                fmt.Println(v3, v21)
            }
            fmt.Println(v20, v20)
            if false {
                var v26 int64 = v21
                _ = v26
            } else {
                v21 = (-(v21))
            }
        }
        _ = fi24
    }
    v27 := []int{-497, 153, -790, 657}
    var v28 uint64 = v2
    v27 = append(v27, -357)
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    _ = len(v27)
    v4 += (v4 % ((genFn0()) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
