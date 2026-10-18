package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int32 }
type GenS2 struct { F0 int64; F1 uint }
func genFn0() uint16 {
    var v6 uint64 = 4294967295
    fmt.Println(v6)
    v6 += (v6 << ((255) & 63))
    var v7 int = 32767
    _ = v6
    _ = v7
    return 341
}
func genFn1() int8 {
    var v8 uint = 18446744073709551615
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 string = (("\n\t" + "a") + (("" + "\n\t") + "foo bar"))
        v11 := []int{42, -1}
        _ = fi9
        _ = v10
        _ = v11
    }
    v8 = (v8 % ((v8) | 1))
    var v12 int64 = 1
    var v13 uint32 = 2147483647
    _ = v12
    _ = v13
    _ = v8
    return 100
}

func main() {
    v1 := []int{2147483647, 32767, 0, -9223372036854775808}
    var v2 int64 = -1
    var v3 int64 = v2
    var v4 bool = (42 <= 824)
    var v5 uint8 = 42
    v14 := GenS1{F0: (v3 & v2), F1: 1, F2: -70}
    switch "a" {
    case "a":
        var v15 string = "hello"
        _ = v15
    }
    switch -814 {
    case -9223372036854775808:
        v16 := GenS1{F0: 970, F1: 32767, F2: 0}
        _ = v16
    case 2147483647:
        var v17 uint = 0
        _ = v17
    case 4294967295:
        _ = len(v1)
    }
    if (!(v4 || (!(!true)))) {
        var v18 uint16 = 772
        _ = v18
    }
    v5 = (v5 - (v5 % ((v5) | 1)))
    v19 := GenS1{F0: v3, F1: -899, F2: 42}
    v4 = (v4 || false)
    v2 += (v3 % ((9223372036854775807) | 1))
    if v4 {
        var v20 int8 = -76
        v3 -= (^(v3))
        _ = v20
    }
    v5 = 141
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    v21 := GenS1{F0: v3, F1: 261, F2: -2147483648}
    v22 := GenS2{F0: v2, F1: 349}
    var v23 int32 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
