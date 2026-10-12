package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 bool }
type GenS2 struct { F0 uint64; F1 uint16 }
func genFn0(p0 string, p1 int16) int8 {
    if true {
        for fi4 := 0; fi4 < 5; fi4++ {
            var v5 uint = 42
            fmt.Println(p0, v5)
            if ((-1 > -51) || (!true)) {
                fmt.Println(v5, fi4)
            }
            _ = fi4
            _ = v5
        }
        v6 := []int{0, -9223372036854775808}
        _ = v6
    } else {
        v7 := GenS1{F0: (0 <= 1), F1: 2147483647, F2: true}
        var v8 uint8 = 128
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 int64 = -196
            var v11 bool = (!false)
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v7
        _ = v8
    }
    var v12 uint = 4294967295
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            v12 += (v12 & v12)
            v15 := []int{fi14, (fi13 & (fi14 - fi14))}
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    v16 := GenS2{F0: 4294967295, F1: 826}
    _ = p0
    _ = p1
    _ = v12
    _ = v16
    return 82
}

func main() {
    var v1 uint32 = 65535
    var v2 uint64 = 4294967295
    var v3 uint8 = 0
    if false {
        v17 := []int{2147483647, -863, 32768}
        v2 += 184
        fmt.Println(v2, v1)
        _ = v17
    }
    fmt.Println(v3, v3, v3)
    v18 := []int{1}
    if len(v18) > 0 {
        v18[0] = 42
    }
    var v19 float64 = (-2.25 + 3.14159)
    var v20 int8 = 42
    v2 -= (v2 % ((v2) | 1))
    switch "" {
    case "\n\t":
        var v21 float32 = 0.1
        _ = v21
    case "a":
        var v22 int = -2147483648
        _ = v22
    }
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        var v24 bool = false
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        for fi26 := 0; fi26 < 10; fi26++ {
            var v27 uint8 = v3
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    v1 += (^(v1))
    _ = len(v18)
    switch -2147483648 {
    case 1:
        v19 = 3.14159
    case 32767:
        var v28 bool = true
        _ = v28
    case -956:
        var v29 uint32 = 419
        _ = v29
    }
    var v30 uint8 = (^(v3))
    v18 = append(v18, 484)
    if len(v18) > 0 {
        v18[0] = 2147483647
    }
    v31 := GenS2{F0: (v2 / ((646) | 1)), F1: uint16(v3)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
}
