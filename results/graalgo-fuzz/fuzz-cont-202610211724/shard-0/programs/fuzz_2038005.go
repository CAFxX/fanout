package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int; F1 bool; F2 uint8 }
func genFn0() uint16 {
    var v6 int64 = 9223372036854775807
    v7 := GenS1{F0: 255, F1: 42, F2: 49}
    v6 = 42
    v8 := GenS2{F0: 1, F1: false, F2: 249}
    _ = v6
    _ = v7
    _ = v8
    return 2
}
func genFn1() uint8 {
    var v9 float64 = (-2.25 / 1e10)
    var v10 uint16 = 801
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 int = 4294967295
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS1{F0: 88, F1: 169, F2: 0}
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    var v15 float32 = -2.25
    _ = v10
    _ = v15
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{1, 4294967295, 2147483647, 4294967295}
    v2 := []string{"\n\t", "foo bar", ("\n\t" + ""), (("" + "foo bar") + "a")}
    var v3 uint32 = 4294967295
    var v4 uint8 = 96
    v5 := []string{"\n\t", "a", "foo bar"}
    var v16 int32 = -2147483648
    v16 += 765
    var v17 float64 = float64(1.5)
    v3 = v3
    if (!(4294967295 <= 255)) {
        var v18 uint = 1
        v18 -= uint(v3)
        v17 += v17
        _ = v18
    } else {
        v4 = v4
        v3 += (^(v3))
        v19 := GenS2{F0: 4294967295, F1: false, F2: 127}
        _ = v19
    }
    v3 = 70
    v4 -= v4
    var v20 int16 = 42
    var v21 uint = 2
    v17 -= 3.14159
    v17 = ((0.1 * (float64(float32(v17)) + 1e10)) * v17)
    var v22 int32 = (v16 | 0)
    v23 := GenS2{F0: 2147483647, F1: (v4 >= (v4 << (((v21 << (((v21 * (^(v21)))) & 63))) & 63))), F2: v4}
    var v24 bool = ((42 > 100) || (!(v21 > 2)))
    var v25 uint8 = (v4 - v4)
    for fi26 := 0; fi26 < 18; fi26++ {
        v27 := GenS1{F0: 2, F1: uint64(v3), F2: (^(v25))}
        var v28 string = "hello"
        v22 += v22
        _ = fi26
        _ = v27
        _ = v28
    }
    v24 = ((v24 && ((-1 == int64(v22)) && false)) && v24)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
