package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int32; F1 uint64; F2 bool }
func genFn0() int {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint32 = 633
        v8 = (^(v8))
        _ = fi7
        _ = v8
    }
    var v9 int32 = -2147483648
    v10 := []int{}
    v9 += (v9 % ((v9) | 1))
    if len(v10) > 0 {
        v10[0] = -2147483648
    }
    _ = v10
    _ = v9
    return 13
}
func genFn1(p0 int8) string {
    v11 := GenS1{F0: -2147483648, F1: false, F2: 1}
    var v12 uint32 = 4294967295
    if false {
        v13 := []int{-1, 9223372036854775807, -376, 42}
        v14 := GenS2{F0: 931}
        _ = v13
        _ = v14
    }
    if ("\n\t" < "") {
        v12 = v12
    }
    _ = p0
    _ = v11
    _ = v12
    return "foo bar"
}

func main() {
    var v1 uint64 = 2
    var v2 int = 4294967295
    v3 := make(map[uint32]int)
    var v4 bool = (true && true)
    var v5 uint32 = 412
    var v6 int = (-(v2))
    switch (v6 >> ((4294967295) & 63)) {
    case 42:
        v4 = false
    case 9223372036854775807:
        var v15 int16 = 1
        _ = v15
    }
    v3[49] = genFn0()
    switch "hello" {
    case "hello":
        v16 := GenS3{F0: 1, F1: 1}
        _ = v16
    case "foo bar":
        var v17 uint = 65535
        _ = v17
    }
    switch v6 {
    case 32767:
        v18 := GenS4{F0: 879, F1: 0, F2: true}
        _ = v18
    }
    v2 = (v2 + v2)
    switch ("a" + "") {
    case "a":
        var v19 int32 = 1
        _ = v19
    }
    var v20 int8 = 109
    v21 := []int{}
    v20 -= (v20 >> ((4294967295) & 63))
    var v22 float64 = ((0.1 - -2.25) + (-2.25 * (0.0 * 1e10)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
