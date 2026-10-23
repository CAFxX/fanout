package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint64 }
func genFn0(p0 string, p1 int32) int16 {
    fmt.Println(p1, p1)
    if false {
        v7 := []int{1}
        _ = v7
    }
    var v8 uint = 664
    p0 = "foo bar"
    _ = p0
    _ = p1
    _ = v8
    return 32767
}

func main() {
    var v1 int = -9223372036854775808
    var v2 uint32 = 791
    v3 := []uint32{v2, v2, (v2 % ((v2) | 1))}
    var v4 bool = true
    v5 := []int{}
    var v6 string = "a"
    var v9 int8 = -128
    var v10 uint32 = (v2 - 565)
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 uint64 = 18446744073709551615
        if len(v3) > 0 {
            v3[0] = 42
        }
        _ = fi11
        _ = v12
    }
    v6 = "a"
    if len(v5) > 0 {
        v5[0] = 32768
    }
    var v13 uint = 4294967295
    for fi14 := 0; fi14 < 18; fi14++ {
        switch -802 {
        case -512:
            for fi15 := 0; fi15 < 13; fi15++ {
                v16 := GenS1{F0: 0, F1: 1, F2: 18446744073709551615}
                _ = fi15
                _ = v16
            }
        default:
            var v17 uint = 65535
            _ = v17
        }
        _ = fi14
    }
    var v18 int16 = 0
    v19 := GenS1{F0: v18, F1: 1, F2: 0}
    if v4 {
        v20 := GenS1{F0: genFn0(v6, 2147483647), F1: 2, F2: uint64(v13)}
        v4 = v4
        for fi21 := 0; fi21 < 9; fi21++ {
            v22 := GenS1{F0: 42, F1: 128, F2: 4294967295}
            fmt.Println(v2, v2, v18)
            v3 = append(v3, uint32(v13))
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    var v23 int8 = (^(v9))
    v18 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
