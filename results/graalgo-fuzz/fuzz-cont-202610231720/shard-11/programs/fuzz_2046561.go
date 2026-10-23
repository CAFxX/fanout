package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    var v4 float32 = 3.14159
    v5 := make(map[int]string)
    var v6 bool = (435 > -70)
    if (-1 < 1) {
        var v7 uint64 = 18446744073709551615
        _ = v7
    }
    v4 += v4
    if (0 > 869) {
        if v6 {
            if v6 {
                var v8 int64 = 1
                for fi9 := 0; fi9 < 1; fi9++ {
                    var v10 bool = (-25 == 0)
                    _ = fi9
                    _ = v10
                }
                for fi11 := 0; fi11 < 7; fi11++ {
                    v12 := []int{(fi11 | (^(fi11))), fi11}
                    v13 := GenS1{F0: 42, F1: 9223372036854775807, F2: 8}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = v8
            }
        }
        switch -725 {
        case 32767:
            var v14 uint64 = 290
            _ = v14
        case -9223372036854775808:
            var v15 uint16 = 1
            _ = v15
        case 32768:
            if true {
                var v16 uint16 = 2
                v16 -= (v16 - (v16 << ((0) & 63)))
                switch "hello" {
                case "a":
                    fmt.Println(v4)
                default:
                    v17 := []int{-2147483648, 4294967295, 2147483647, -381}
                    _ = v17
                }
                _ = v16
            } else {
                fmt.Println(v6)
                for fi18 := 0; fi18 < 8; fi18++ {
                    fmt.Println(v4, fi18)
                    _ = fi18
                }
            }
        }
        var v19 int32 = 59
        _ = v19
    }
    v6 = v6
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 int = 0
        var v22 float32 = (0.0 / v4)
        _ = fi20
        _ = v21
        _ = v22
    }
    v6 = v6
    if (1 == 1) {
        v4 = (float32(0.1) + (0.0 - -2.25))
    }
    var v23 bool = v6
    fmt.Println(v23, v6)
    var v24 int16 = 0
    var v25 float64 = 1e10
    if false {
        switch 0 {
        case -737:
            v26 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: 42}
            _ = v26
        default:
            v25 = float64(v4)
        }
        if v6 {
            for fi27 := 0; fi27 < 8; fi27++ {
                var v28 string = ""
                v29 := []int{fi27, 9223372036854775807, fi27, (fi27 | fi27)}
                _ = fi27
                _ = v28
                _ = v29
            }
            var v30 int8 = -8
            _ = v30
        } else {
            var v31 uint32 = 4294967295
            _ = v31
        }
        v23 = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
