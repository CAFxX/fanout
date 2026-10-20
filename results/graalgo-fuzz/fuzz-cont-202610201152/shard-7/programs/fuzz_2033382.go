package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint8; F1 int; F2 int16 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 string; F1 string; F2 string }
func genFn0(p0 uint, p1 int8, p2 int8) int16 {
    p2 += p1
    v6 := GenS1{F0: 2, F1: 0}
    switch -1 {
    case 1:
        v7 := []int{128}
        _ = v7
    default:
        p2 = -128
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1000
}
func genFn1(p0 uint32) uint16 {
    p0 += 42
    switch 32767 {
    case -872:
        v8 := []int{128, 4294967295}
        _ = v8
    case 32768:
        switch -1 {
        case 4294967295:
            p0 -= p0
        }
    case 127:
        v9 := []int{32767, -29, -681, 42}
        _ = v9
    }
    p0 = p0
    _ = p0
    return 65535
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = -1
    var v3 bool = true
    v4 := make(map[int]uint32)
    v5 := []uint32{2, 1}
    var v10 uint64 = 374
    v2 = (v2 & v2)
    if v3 {
        v3 = v3
        v11 := GenS1{F0: 138, F1: 42}
        v12 := GenS1{F0: 1, F1: 9223372036854775807}
        _ = v11
        _ = v12
    }
    _ = len(v5)
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 uint = 18446744073709551615
        v15 := GenS2{F0: 137, F1: (fi13 - fi13), F2: -32768}
        _ = fi13
        _ = v14
        _ = v15
    }
    if v3 {
        fmt.Println(v2)
    } else {
        var v16 float32 = float32((0.1 / 0.1))
        var v17 float32 = v16
        fmt.Println(v3, v3, v2)
        _ = v16
        _ = v17
    }
    var v18 uint8 = 3
    var v19 uint32 = 0
    fmt.Println(v10, v10, v3)
    if (9223372036854775807 >= 1) {
        v18 += 128
        v10 += (v10 << ((65535) & 63))
    }
    v2 -= -1
    switch "a" {
    case "":
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 int32 = -516
            _ = fi20
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        if v3 {
            v5 = append(v5, v19)
        } else {
            var v23 int64 = 0
            _ = v23
        }
        for fi24 := 0; fi24 < 14; fi24++ {
            fmt.Println(fi24, v2, v3)
            if (42 > 42) {
                v18 -= v18
                v25 := GenS1{F0: 50, F1: -691}
                switch "" {
                case "a":
                    v26 := GenS1{F0: uint8(v10), F1: -1}
                    _ = v26
                }
                var v27 bool = false
                v28 := GenS3{F0: (v19 - (v19 / ((v19) | 1))), F1: (-2147483648 >= -2147483648)}
                _ = v25
                _ = v27
                _ = v28
            } else {
                v5 = append(v5, (v19 + (v19 & 0)))
                v5 = append(v5, (v19 * 1))
                if (1e10 != 3.14159) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    var v29 int8 = v2
                    var v30 uint64 = (v10 >> ((190) & 63))
                    var v31 int16 = -1
                    _ = v29
                    _ = v30
                    _ = v31
                }
            }
            fmt.Println(v19, v2)
            _ = fi24
        }
        v32 := GenS4{F0: (("" + "a") + ("foo bar" + "\n\t")), F1: "a", F2: ("foo bar" + (("\n\t" + "") + ("foo bar" + "foo bar")))}
        _ = fi22
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
