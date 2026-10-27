package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int }
type GenS2 struct { F0 uint; F1 uint8 }

func main() {
    var v1 int = -1
    v2 := []int{v1, (-(v1)), v1}
    v3 := []uint32{0, 65535, 2147483647, 2147483648}
    v4 := []string{""}
    var v5 int8 = 127
    var v6 uint64 = 1
    v7 := GenS1{F0: true, F1: 0, F2: v1}
    v4 = append(v4, "\n\t")
    var v8 int8 = (v5 | v5)
    v4 = append(v4, "\n\t")
    var v9 float64 = 1.5
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            v5 += (-(v5))
            v12 := GenS2{F0: 0, F1: 1}
            v9 -= 0.0
            var v13 float64 = float64((0.0 + -2.25))
            v14 := GenS2{F0: 2, F1: 124}
            v8 -= v5
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = len(v3)
        v6 += (v6 + v6)
        _ = fi10
    }
    v15 := GenS2{F0: 2, F1: 42}
    var v16 int8 = v8
    v16 += v8
    if (!(int8(v1) > v5)) {
        var v17 uint32 = 4294967295
        v4 = append(v4, ("\n\t" + "hello"))
        if true {
            switch "a" {
            case "a":
                var v18 int32 = 2147483647
                _ = v18
            case "foo bar":
                var v19 int32 = 452
                _ = v19
            }
        } else {
            _ = len(v2)
        }
        _ = v17
    } else {
        var v20 int = 4294967295
        _ = v20
    }
    _ = len(v2)
    if false {
        v3 = append(v3, 42)
        var v21 int = (v1 - (-(v1)))
        v22 := GenS1{F0: false, F1: 4294967295, F2: v21}
        _ = v21
        _ = v22
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v23 int64 = -9223372036854775808
    var v24 float64 = 1e10
    v8 = (v16 / (((^(v5))) | 1))
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
