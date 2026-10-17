package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 uint32; F2 uint64 }
func genFn0(p0 bool, p1 uint64) uint {
    var v4 int = 736
    v5 := []int{(v4 & 9223372036854775807), v4}
    v6 := GenS1{F0: 922, F1: (p1 / (((p1 >> ((0) & 63))) | 1)), F2: 502}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int8, p1 int8, p2 bool) string {
    for fi7 := 0; fi7 < 15; fi7++ {
        if (p2 || (2 >= 1)) {
            v8 := GenS2{F0: 1, F1: (42 == 1)}
            _ = v8
        }
        _ = fi7
    }
    var v9 int64 = 9223372036854775807
    var v10 uint = 4294967295
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(v10)
        if (187 == 2) {
            v9 += v9
        }
        fmt.Println(p0)
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return ((("\n\t" + ("" + "hello")) + "a") + "")
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    var v3 string = ("\n\t" + "")
    var v12 uint64 = 1
    if (1 > 42) {
        v12 = v12
    }
    v13 := []int{128, 42, 32767, 837}
    var v14 int = -9223372036854775808
    switch (^(v14)) {
    case 559:
        var v15 uint8 = 2
        _ = v15
    case -1:
        if (32767 <= 1) {
            var v16 int16 = 290
            _ = v16
        } else {
            var v17 int8 = 112
            _ = v17
        }
    case 2147483647:
        var v18 int32 = -2147483648
        _ = v18
    }
    if (!(!false)) {
        v19 := v1[v3]
        v20 := GenS1{F0: 2, F1: v12, F2: 1}
        _ = v19
        _ = v20
    }
    var v21 int8 = 127
    if ((1 != 1) && (((!false) && true) && true)) {
        delete(v1, v3)
        v22 := GenS3{F0: 1, F1: -446}
        _ = v22
    } else {
        fmt.Println(v3, v3)
        var v23 int64 = 9223372036854775807
        _ = v23
    }
    var v24 bool = (v21 > (v21 % (((v21 << ((18446744073709551615) & 63))) | 1)))
    if (!(v12 <= (^(v12)))) {
        v12 += v12
        v12 -= v12
        v14 = (v14 % ((v14) | 1))
    } else {
        v24 = false
    }
    v25 := GenS4{F0: 0}
    if true {
        var v26 int16 = -557
        _ = v26
    } else {
        v12 -= 65
    }
    if true {
        fmt.Println(v14, v3)
        v13 = append(v13, (-(v14)))
    }
    v27 := GenS5{F0: v24, F1: 65535, F2: (v12 % ((491) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v27
    _ = v3
}
