package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 bool; F2 int32 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 uint32; F1 bool }
func genFn0(p0 int32, p1 uint64, p2 string) uint8 {
    if (!(!false)) {
        p1 += 1
        v7 := GenS1{F0: 194, F1: -458}
        var v8 int8 = 127
        _ = v7
        _ = v8
    }
    fmt.Println(p2, p1, p1)
    for fi9 := 0; fi9 < 16; fi9++ {
        if false {
            var v10 int32 = p0
            _ = v10
        }
        if (65535 >= 145) {
            var v11 int64 = 9223372036854775807
            v11 += (v11 - v11)
            _ = v11
        }
        _ = fi9
    }
    var v12 int32 = (^(p0))
    v13 := GenS1{F0: 200, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 111
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 4294967295
    v3 := []int{128, 1, -243}
    var v4 uint64 = 0
    var v5 uint64 = (^(v4))
    v6 := make(map[uint32]bool)
    fmt.Println(v5, v2, v2)
    v14 := GenS1{F0: (v2 / (((v2 >> ((18446744073709551615) & 63))) | 1)), F1: -1}
    v15 := v1["\n\t"]
    v16 := GenS1{F0: 562, F1: 142}
    if (-1 >= -128) {
        var v17 int = v15
        fmt.Println(v17, v5, v2)
        _ = v17
    } else {
        var v18 float64 = 0.0
        _ = v18
    }
    fmt.Println(v15)
    var v19 float32 = ((0.0 + 1e10) - 0.0)
    fmt.Println(v2)
    fmt.Println(v2, v5, v4)
    v5 = v4
    for fi20 := 0; fi20 < 7; fi20++ {
        if (!(false && (-367 <= -974))) {
            v21 := GenS2{F0: -9223372036854775808}
            var v22 float64 = 0.0
            _ = v21
            _ = v22
        }
        _ = fi20
    }
    v5 += v4
    if ((true && false) || true) {
        delete(v6, (v2 * (v2 + 226)))
        for fi23 := 0; fi23 < 6; fi23++ {
            v24 := GenS2{F0: -501}
            v25 := GenS3{F0: v15, F1: (283 <= 1), F2: 737}
            _ = fi23
            _ = v24
            _ = v25
        }
    } else {
        var v26 string = "\n\t"
        _ = v26
    }
    v27 := GenS4{F0: 65535, F1: -19}
    switch -2147483648 {
    case 42:
        for fi28 := 0; fi28 < 19; fi28++ {
            var v29 uint = 597
            v29 += v29
            _ = fi28
            _ = v29
        }
    case 32:
        var v30 int64 = 42
        _ = v30
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v31 := v6[(v2 * v2)]
    if (-2147483648 < 728) {
        v19 = 1e10
        switch v15 {
        case 4294967295:
            var v32 uint16 = 787
            _ = v32
        case 128:
            var v33 float32 = v19
            _ = v33
        case 32767:
            v34 := GenS1{F0: v2, F1: 42}
            _ = v34
        }
        var v35 int8 = 91
        _ = v35
    } else {
        for fi36 := 0; fi36 < 19; fi36++ {
            _ = len(v6)
            v5 += 565
            v37 := GenS5{F0: v2, F1: v31}
            _ = fi36
            _ = v37
        }
        var v38 uint32 = v2
        if len(v3) > 0 {
            v3[0] = v15
        }
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
