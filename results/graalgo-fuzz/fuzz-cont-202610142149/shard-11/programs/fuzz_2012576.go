package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int32; F2 int16 }
type GenS3 struct { F0 int64; F1 uint32 }
func genFn0(p0 int64) int32 {
    var v4 string = "a"
    var v5 float64 = 1.5
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := []int{}
        v4 = (("foo bar" + v4) + "\n\t")
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 1}
    fmt.Println(p0, v5)
    _ = p0
    _ = v4
    _ = v5
    _ = v8
    return 2147483647
}
func genFn1() uint8 {
    var v9 uint64 = 18446744073709551615
    v9 = v9
    switch "" {
    case "":
        v9 += 18446744073709551615
    default:
        fmt.Println(v9)
    }
    if true {
        v9 = v9
    } else {
        v10 := GenS1{F0: 18446744073709551615}
        _ = v10
    }
    var v11 uint32 = 65535
    _ = v11
    _ = v9
    return 108
}

func main() {
    var v1 int8 = 1
    v2 := []int{413, -629, 42}
    var v3 float64 = (float64(1e10) / 3.14159)
    v1 -= -1
    switch 9223372036854775807 {
    case 127:
        _ = len(v2)
    case 9223372036854775807:
        switch 9223372036854775807 {
        case 2147483647:
            v1 -= (v1 % ((v1) | 1))
        }
    case 1:
        var v12 int16 = 32767
        _ = v12
    default:
        for fi13 := 0; fi13 < 19; fi13++ {
            for fi14 := 0; fi14 < 4; fi14++ {
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
                v15 := GenS2{F0: (^(v1)), F1: 1, F2: 32767}
                _ = fi14
                _ = v15
            }
            v1 = v1
            _ = fi13
        }
    }
    if true {
        var v16 int = 378
        _ = v16
    } else {
        var v17 float64 = (v3 / -2.25)
        _ = v17
    }
    v1 -= v1
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := GenS3{F0: 0, F1: 65535}
        if false {
            v3 = ((v3 / ((float64(0.1) + v3) + v3)) * v3)
        }
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        v1 = 127
        _ = fi20
    }
    if false {
        for fi21 := 0; fi21 < 2; fi21++ {
            v2 = append(v2, fi21)
            v22 := GenS2{F0: v1, F1: 2147483647, F2: -311}
            v2 = append(v2, fi21)
            _ = fi21
            _ = v22
        }
        var v23 uint32 = 2147483648
        var v24 uint8 = 128
        _ = v23
        _ = v24
    }
    switch "a" {
    case "":
        for fi25 := 0; fi25 < 4; fi25++ {
            v1 -= v1
            _ = fi25
        }
    case "hello":
        v2 = append(v2, -1)
    }
    var v26 uint64 = 134
    var v27 uint32 = 4294967295
    v1 = (v1 | (-(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
