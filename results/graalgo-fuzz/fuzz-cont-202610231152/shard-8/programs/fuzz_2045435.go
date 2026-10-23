package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint16; F2 int32 }
func genFn0() int16 {
    var v7 bool = ((!true) || (((74 >= -128) && true) || (4294967295 < 0)))
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: 255}
        switch fi8 {
        case -2147483648:
            var v10 string = "a"
            _ = v10
        case 0:
            if (fi8 > (fi8 * fi8)) {
                v7 = v7
                v11 := GenS1{F0: 65}
                _ = v11
            } else {
                var v12 int8 = -1
                for fi13 := 0; fi13 < 14; fi13++ {
                    v14 := []int{fi8}
                    v12 = (v12 ^ (v12 * v12))
                    v15 := GenS2{F0: (!((v7 && (-9223372036854775808 > 42)) || v7)), F1: 666, F2: 1}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = v12
            }
        case 4294967295:
            var v16 float64 = float64(3.14159)
            _ = v16
        }
        if (!(!v7)) {
            var v17 int16 = -32768
            _ = v17
        } else {
            var v18 uint16 = 0
            v19 := []int{fi8}
            _ = v18
            _ = v19
        }
        _ = fi8
        _ = v9
    }
    v7 = v7
    switch "foo bar" {
    case "foo bar":
        v20 := GenS1{F0: 50}
        _ = v20
    }
    v7 = (((((0.0 / 1e10) * 0.0) / (-2.25 - 1e10)) + 1.5) > (1.5 + 1e10))
    _ = v7
    return -32768
}
func genFn1() uint8 {
    var v21 int16 = -449
    if true {
        var v22 bool = true
        _ = v22
    } else {
        if (127 > 125) {
            v21 += v21
            v23 := GenS1{F0: 8}
            _ = v23
        } else {
            v24 := GenS3{F0: 4294967295, F1: 0, F2: 42}
            v25 := []int{-1, 228, 0}
            v21 += (v21 >> ((2) & 63))
            _ = v24
            _ = v25
        }
    }
    var v26 uint = 4294967295
    _ = v21
    _ = v26
    return 148
}

func main() {
    var v1 float64 = 3.14159
    var v2 int32 = 1
    var v3 uint16 = 942
    var v4 float64 = (float64(-2.25) - v1)
    var v5 int64 = -1
    v6 := []string{"\n\t", "hello"}
    var v27 int8 = 127
    var v28 uint64 = 18446744073709551615
    fmt.Println(v27)
    if (((2147483647 <= 4294967295) || true) && true) {
        var v29 uint8 = 42
        v30 := GenS1{F0: v29}
        if ((^(v27)) > (v27 / ((0) | 1))) {
            _ = len(v6)
        }
        _ = v29
        _ = v30
    } else {
        v5 = (v5 / ((v5) | 1))
    }
    switch -583 {
    case 331:
        if (65535 == 65535) {
            switch "a" {
            case "":
                v1 += ((0.1 - (float64(((float32(1e10) - 0.0) / -2.25)) / 1e10)) - 1.5)
            }
            var v31 float64 = (v1 * (v4 / 1.0))
            _ = v31
        }
    case 4294967295:
        fmt.Println(v2)
    default:
        switch 32767 {
        case -9223372036854775808:
            fmt.Println(v5)
        }
    }
    fmt.Println(v4, v3)
    var v32 uint8 = 65
    v33 := GenS1{F0: 127}
    v2 -= int32(v28)
    v2 += v2
    if ("a" >= ("a" + "\n\t")) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        for fi34 := 0; fi34 < 9; fi34++ {
            v3 += v3
            _ = fi34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
