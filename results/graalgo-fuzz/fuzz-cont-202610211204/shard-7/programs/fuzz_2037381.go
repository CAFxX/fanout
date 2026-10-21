package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int16; F1 int32 }
func genFn0() uint32 {
    switch 0 {
    case -1:
        var v6 int8 = -1
        _ = v6
    case 549:
        var v7 int64 = 9223372036854775807
        _ = v7
    default:
        v8 := GenS1{F0: 1, F1: -253, F2: 76}
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS1{F0: -1, F1: -2147483648, F2: 100}
        switch fi9 {
        case -1:
            switch fi9 {
            case 128:
                v11 := GenS1{F0: 503, F1: -1, F2: 1}
                _ = v11
            case -1:
                v12 := []int{fi9, fi9, 64, (^(fi9))}
                _ = v12
            case -448:
                var v13 bool = (18446744073709551615 == 437)
                _ = v13
            }
        default:
            for fi14 := 0; fi14 < 12; fi14++ {
                if true {
                    fmt.Println(fi9, fi9)
                    v15 := []int{}
                    _ = v15
                } else {
                    var v16 int16 = 32767
                    _ = v16
                }
                fmt.Println(fi14)
                if false {
                    switch fi14 {
                    case 128:
                        fmt.Println(fi14, fi14)
                    case -1:
                        v17 := []int{}
                        _ = v17
                    case -9223372036854775808:
                        var v18 int = (^(fi9))
                        _ = v18
                    }
                    if true {
                        if (50 < -86) {
                            var v19 uint64 = 2
                            _ = v19
                        } else {
                            v20 := GenS1{F0: -519, F1: -1, F2: 42}
                            v21 := GenS2{F0: 92, F1: -9223372036854775808}
                            _ = v20
                            _ = v21
                        }
                        v22 := []int{}
                        _ = v22
                    }
                }
                var v23 float64 = 1e10
                _ = fi14
                _ = v23
            }
        }
        _ = fi9
        _ = v10
    }
    switch 579 {
    case 691:
        var v24 float32 = 0.0
        _ = v24
    }
    return 715
}
func genFn1(p0 int) uint8 {
    v25 := []int{p0, p0, -2147483648, p0}
    if true {
        switch "hello" {
        case "foo bar":
            v25 = append(v25, p0)
        }
    } else {
        p0 -= (p0 + p0)
    }
    v26 := GenS2{F0: 1, F1: int64(p0)}
    var v27 uint16 = 20
    _ = p0
    _ = v25
    _ = v26
    _ = v27
    return 105
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]int)
    var v3 int64 = -9223372036854775808
    var v4 uint16 = 65535
    var v5 int = 32768
    v4 = 461
    var v28 int = v5
    switch 9223372036854775807 {
    case 42:
        v5 = v5
    }
    if false {
        switch v28 {
        case 32767:
            v1 = (v1 - (v1 & (v1 & genFn1(-873))))
        }
        var v29 bool = ((603 <= 18446744073709551615) && false)
        _ = v29
    }
    fmt.Println(v1)
    v30 := GenS3{F0: 42, F1: 42}
    v5 = -585
    var v31 float64 = 1.5
    for fi32 := 0; fi32 < 16; fi32++ {
        fmt.Println(v5)
        _ = fi32
    }
    v28 = 32767
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
