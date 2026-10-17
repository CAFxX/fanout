package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string }
func genFn0(p0 uint8, p1 uint, p2 int32) uint16 {
    for fi4 := 0; fi4 < 2; fi4++ {
        if (3.14159 == (float32(0.1) / ((1.5 / -2.25) * -2.25))) {
            fmt.Println(p2, p2)
            var v5 uint32 = 1
            _ = v5
        } else {
            v6 := GenS1{F0: 9223372036854775807, F1: 523}
            v7 := []int{32768, fi4, fi4}
            v8 := GenS2{F0: 0, F1: 497, F2: 0}
            _ = v6
            _ = v7
            _ = v8
        }
        _ = fi4
    }
    var v9 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 65535
}
func genFn1(p0 uint32, p1 uint16) int64 {
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 int8 = 127
        v12 := []int{(fi10 + fi10), (fi10 / ((fi10) | 1))}
        _ = fi10
        _ = v11
        _ = v12
    }
    switch "\n\t" {
    case "a":
        var v13 string = "a"
        _ = v13
    default:
        if true {
            v14 := GenS3{F0: -9223372036854775808, F1: 974}
            var v15 uint = 604
            for fi16 := 0; fi16 < 9; fi16++ {
                if true {
                    v17 := GenS4{F0: (^(v15))}
                    _ = v17
                }
                var v18 bool = (p0 == p0)
                v19 := GenS4{F0: (v15 | v15)}
                _ = fi16
                _ = v18
                _ = v19
            }
            _ = v14
            _ = v15
        }
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint = 1
    v3 := make(map[int]uint32)
    v20 := GenS5{F0: "\n\t"}
    v21 := GenS3{F0: genFn1(886, 32767), F1: 2147483647}
    switch "a" {
    case "foo bar":
        switch -273 {
        case 0:
            if true {
                v2 = v2
                v2 = 4294967295
            } else {
                var v22 string = ("foo bar" + (("" + "\n\t") + ("" + ("foo bar" + "hello"))))
                delete(v3, 2147483647)
                _ = v22
            }
        case -9223372036854775808:
            fmt.Println(v2, v1)
        case -1:
            var v23 int8 = 42
            _ = v23
        default:
            v24 := GenS3{F0: v1, F1: 2147483647}
            _ = v24
        }
    case "":
        v1 -= (v1 / ((-9223372036854775808) | 1))
    case "\n\t":
        v25 := []int{}
        _ = v25
    default:
        v26 := v3[4294967295]
        _ = v26
    }
    v27 := []int{122, 32767, 32768, 1}
    fmt.Println(v1, v1)
    fmt.Println(v2, v1)
    var v28 uint32 = 453
    if true {
        v1 -= (v1 * v1)
    } else {
        if ("foo bar" < "a") {
            v29 := v3[128]
            var v30 int32 = 2147483647
            _ = v29
            _ = v30
        }
    }
    fmt.Println(v1, v28, v2)
    fmt.Println(v1, v2, v1)
    if (1 < 0) {
        if len(v27) >= 2 {
            v27 = v27[1:len(v27)]
        }
        for fi31 := 0; fi31 < 1; fi31++ {
            var v32 int8 = 127
            var v33 float64 = 1e10
            var v34 float64 = v33
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        v2 += (^(v2))
    }
    switch 1 {
    case 32767:
        v1 = 0
    default:
        var v35 int32 = 42
        _ = v35
    }
    var v36 uint32 = v28
    var v37 uint32 = 2147483648
    v28 = v36
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v37
}
