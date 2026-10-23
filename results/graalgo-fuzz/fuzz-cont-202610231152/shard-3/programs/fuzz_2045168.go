package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int8; F1 bool; F2 int64 }
func genFn0() uint8 {
    var v7 uint = 255
    v7 = (v7 * v7)
    v8 := []int{}
    if true {
        var v9 uint64 = 743
        v10 := GenS1{F0: 9223372036854775807, F1: 105}
        for fi11 := 0; fi11 < 11; fi11++ {
            v9 = (v9 % (((v9 << (((v7 - (v7 - (^(v7))))) & 63))) | 1))
            _ = fi11
        }
        _ = v10
        _ = v9
    } else {
        var v12 bool = true
        switch ("hello" + "") {
        case "foo bar":
            for fi13 := 0; fi13 < 8; fi13++ {
                if len(v8) > 0 {
                    v8[0] = fi13
                }
                v7 += v7
                _ = fi13
            }
        case "hello":
            if v12 {
                var v14 float32 = (float32(0.0) / (0.1 * 3.14159))
                if (1 == 9223372036854775807) {
                    fmt.Println(v7, v12)
                }
                _ = v14
            }
        }
        v15 := GenS1{F0: 1, F1: 28}
        _ = v12
        _ = v15
    }
    _ = v7
    _ = v8
    return 128
}

func main() {
    v1 := []int{42, 326, -2147483648}
    var v2 string = "foo bar"
    v3 := []uint32{1, 4294967295}
    var v4 int8 = 100
    var v5 string = "foo bar"
    var v6 int32 = -2147483648
    v16 := GenS1{F0: -1, F1: 0}
    var v17 bool = (0 == 32767)
    fmt.Println(v17)
    if v17 {
        var v18 float32 = (((1.5 + 3.14159) - 0.1) - -2.25)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v18
    }
    v17 = v17
    v19 := GenS1{F0: 32767, F1: (v4 & -1)}
    var v20 uint8 = 42
    v21 := GenS2{F0: (v4 % ((120) | 1)), F1: (32767 >= -1), F2: 1}
    v22 := GenS2{F0: (^(v4)), F1: ((v4 & v4) == 127), F2: 337}
    v4 = -128
    v4 = v4
    for fi23 := 0; fi23 < 8; fi23++ {
        v2 = ((((v2 + v5) + "foo bar") + v2) + v5)
        v24 := GenS2{F0: (v4 << ((0) & 63)), F1: v17, F2: -1}
        v5 = ((v5 + v2) + ("hello" + v5))
        _ = fi23
        _ = v24
    }
    v1 = append(v1, -9223372036854775808)
    v17 = false
    v2 = "foo bar"
    fmt.Println(v2, v20, v17)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
