package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int8 }
type GenS2 struct { F0 uint; F1 uint64; F2 uint32 }
func genFn0() string {
    var v4 uint16 = 954
    v4 -= 42
    _ = v4
    return (("foo bar" + (("\n\t" + ((("hello" + (("" + "hello") + "hello")) + "foo bar") + (("\n\t" + ("a" + "hello")) + ""))) + "a")) + "")
}
func genFn1() uint8 {
    var v5 int16 = 172
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 float64 = (0.1 - float64(1e10))
        v8 := []int{(fi6 / ((0) | 1)), (fi6 << ((42) & 63)), fi6, fi6}
        v7 -= (v7 + float64((0.0 / 0.1)))
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = v5
    return 42
}

func main() {
    var v1 int8 = -73
    var v2 uint32 = 195
    var v3 int16 = -32768
    v9 := GenS1{F0: -2147483648, F1: 9223372036854775807, F2: (v1 >> ((255) & 63))}
    if (42 < v3) {
        var v10 int8 = v1
        var v11 uint = 344
        _ = v10
        _ = v11
    }
    var v12 uint = 18446744073709551615
    v12 -= (v12 * (v12 ^ 42))
    v1 = v1
    v2 = v2
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 uint16 = 65535
        v15 := GenS1{F0: -2147483648, F1: -1, F2: -52}
        var v16 int32 = -798
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    v12 += v12
    switch 992 {
    case 4294967295:
        v17 := []int{32768, int(v2)}
        _ = v17
    case 0:
        var v18 int16 = v3
        _ = v18
    case 912:
        v19 := GenS2{F0: 0, F1: 1, F2: 65535}
        _ = v19
    default:
        v20 := []int{1, 32768, -9223372036854775808}
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
