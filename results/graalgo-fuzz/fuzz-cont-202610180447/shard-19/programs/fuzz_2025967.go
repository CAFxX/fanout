package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint64 }
type GenS2 struct { F0 int }
func genFn0(p0 uint8) uint64 {
    v7 := []int{-1}
    p0 = p0
    var v8 float32 = (0.1 / ((float32(0.1) - 0.0) - float32(-2.25)))
    v9 := GenS1{F0: 117, F1: 9223372036854775807, F2: 0}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1() bool {
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 float64 = 0.0
        fmt.Println(v11)
        v12 := []int{}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 string = "foo bar"
    _ = v13
    return ((!(!true)) || true)
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    var v3 uint16 = 2
    v4 := make(map[int]int)
    var v5 uint16 = (v3 & 2)
    var v6 uint = 413
    if ("\n\t" == ("" + ("foo bar" + "foo bar"))) {
        v5 = v3
        v2[-2147483648] = (-676 > 1)
    } else {
        v14 := []int{}
        _ = v14
    }
    v3 += v5
    var v15 uint32 = 1
    var v16 float64 = 1.5
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 int8 = 15
        _ = fi17
        _ = v18
    }
    var v19 int32 = -2147483648
    v19 += (v19 ^ -1)
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 bool = (!true)
        var v22 int = (^(fi20))
        v15 = v15
        _ = fi20
        _ = v21
        _ = v22
    }
    if (float32(0.1) <= 0.0) {
        var v23 uint8 = 0
        var v24 uint8 = (v23 ^ uint8(v3))
        fmt.Println(v6, v23)
        _ = v23
        _ = v24
    }
    if ("" <= "") {
        v19 -= 513
    }
    v25 := GenS2{F0: 222}
    v3 = v5
    v6 -= (v6 + (v6 >> (((v6 / ((v6) | 1))) & 63)))
    fmt.Println(v6, v3)
    var v26 float32 = (float32(1e10) / -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
