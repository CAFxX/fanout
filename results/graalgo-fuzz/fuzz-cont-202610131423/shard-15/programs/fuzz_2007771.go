package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
func genFn0() int32 {
    var v7 int16 = -309
    if (!((true || (2147483648 > 2147483647)) || true)) {
        v7 = (-(v7))
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{fi8, (^(fi8)), (fi8 / (((fi8 % ((-119) | 1))) | 1))}
        v10 := GenS1{F0: (v7 - int16(fi8)), F1: (2 <= 1)}
        var v11 uint = 4294967295
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v7 -= (^(v7))
    _ = v7
    return 1
}
func genFn1(p0 int16, p1 int8, p2 int32) uint {
    p0 = 0
    v12 := []int{-647, 4294967295, -1, -606}
    v12 = append(v12, 32768)
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 int64 = 9223372036854775807
        var v15 float32 = 1e10
        v16 := GenS1{F0: p0, F1: (653 != 1)}
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 65535
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    var v3 uint32 = 166
    var v4 float64 = 1e10
    var v5 uint16 = 2
    var v6 uint = 813
    v3 = (^(v3))
    v4 = ((1e10 / (v4 * (v4 + 0.0))) - v4)
    for fi17 := 0; fi17 < 13; fi17++ {
        fmt.Println(v5)
        v4 = 1e10
        _ = fi17
    }
    var v18 int8 = 0
    var v19 uint32 = 664
    if ((v18 / (((v18 * int8(v3))) | 1)) != v18) {
        v20 := []int{}
        for fi21 := 0; fi21 < 9; fi21++ {
            v3 += v3
            v22 := GenS1{F0: 0, F1: ((0.0 / (1e10 + 3.14159)) > 1e10)}
            _ = fi21
            _ = v22
        }
        fmt.Println(v19, v5)
        _ = v20
    }
    v19 = (v3 - (^(v19)))
    fmt.Println(v6)
    v23 := GenS1{F0: -32768, F1: (float64(-2.25) != (v4 * (v4 + 0.0)))}
    switch 2147483647 {
    case -456:
        v4 = 0.1
    case 50:
        v24 := []int{int(v6)}
        _ = v24
    case 42:
        var v25 string = ("" + ("\n\t" + "\n\t"))
        _ = v25
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        v27 := GenS1{F0: -1, F1: false}
        _ = fi26
        _ = v27
    }
    var v28 int32 = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
