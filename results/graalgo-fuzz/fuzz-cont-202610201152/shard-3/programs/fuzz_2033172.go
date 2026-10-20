package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
func genFn0() string {
    switch 140 {
    case 42:
        var v7 int8 = -128
        _ = v7
    default:
        var v8 int16 = 0
        _ = v8
    }
    var v9 uint32 = 644
    v10 := GenS1{F0: -1}
    for fi11 := 0; fi11 < 6; fi11++ {
        switch "\n\t" {
        case "foo bar":
            v9 -= (v9 - (v9 % ((v9) | 1)))
        case "hello":
            v12 := []int{}
            _ = v12
        }
        var v13 uint64 = 2
        v14 := []int{(fi11 ^ 42), (fi11 / ((fi11) | 1)), 4294967295}
        _ = fi11
        _ = v13
        _ = v14
    }
    _ = v10
    _ = v9
    return "a"
}
func genFn1() uint32 {
    var v15 int64 = -1
    var v16 bool = (v15 == v15)
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 string = "\n\t"
        v19 := []int{-2147483648}
        v20 := GenS2{F0: 18446744073709551615}
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    var v21 uint8 = 253
    _ = v15
    _ = v16
    _ = v21
    return 4294967295
}

func main() {
    v1 := []string{}
    var v2 int64 = -608
    var v3 uint64 = 64
    v4 := []int{9223372036854775807, -1}
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]string)
    var v22 float32 = (3.14159 * float32(float64(-2.25)))
    var v23 int16 = -1
    v23 += -32768
    v2 += v2
    var v24 int64 = v2
    var v25 int64 = v2
    var v26 int8 = 127
    fmt.Println(v25)
    v27 := GenS2{F0: 18446744073709551615}
    var v28 int = -9223372036854775808
    for fi29 := 0; fi29 < 9; fi29++ {
        for fi30 := 0; fi30 < 1; fi30++ {
            fmt.Println(v28, fi29)
            if (128 >= 11) {
                var v31 float64 = (3.14159 * (0.0 * -2.25))
                v2 += (v2 + (v25 + (v25 - (v2 + 42))))
                _ = v31
            }
            _ = fi30
        }
        _ = fi29
    }
    for fi32 := 0; fi32 < 9; fi32++ {
        v24 += v24
        _ = fi32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
