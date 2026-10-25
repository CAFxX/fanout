package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint }
func genFn0() string {
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := []int{fi6, 370, (fi6 / (((fi6 / ((fi6) | 1))) | 1)), (fi6 >> ((65) & 63))}
        _ = fi6
        _ = v7
    }
    switch 229 {
    case 9223372036854775807:
        var v8 float32 = 0.1
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: 42}
    _ = v9
    return ""
}
func genFn1(p0 bool) uint8 {
    v10 := GenS1{F0: 42, F1: 647}
    v11 := GenS1{F0: 1, F1: 413}
    var v12 bool = (781 < 2147483647)
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return 255
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]int)
    v3 := []int{-1, -9223372036854775808, 42, 591}
    v4 := make(map[uint32]bool)
    var v5 uint32 = 310
    v1 = v1
    v5 = v5
    var v13 int32 = 0
    for fi14 := 0; fi14 < 2; fi14++ {
        v5 -= (v5 & v5)
        v15 := GenS1{F0: 0, F1: 154}
        _ = fi14
        _ = v15
    }
    v13 = (v13 - v13)
    switch 42 {
    case -306:
        for fi16 := 0; fi16 < 19; fi16++ {
            if false {
                v17 := GenS2{F0: 18446744073709551615, F1: 255}
                for fi18 := 0; fi18 < 3; fi18++ {
                    v1 = 127
                    var v19 int = (fi16 * fi16)
                    fmt.Println(v19, v19, fi18)
                    _ = fi18
                    _ = v19
                }
                v13 = v13
                _ = v17
            } else {
                fmt.Println(v5)
                for fi20 := 0; fi20 < 20; fi20++ {
                    var v21 uint8 = 182
                    switch fi20 {
                    case 32768:
                        fmt.Println(fi16, v13)
                    }
                    var v22 float32 = (((((1e10 - 3.14159) + 1.5) / 0.1) + 0.0) / 1.5)
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
            _ = len(v3)
            v1 = uint8(v5)
            _ = fi16
        }
    }
    v23 := GenS2{F0: 1, F1: 1}
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 uint8 = v1
        _ = fi24
        _ = v25
    }
    switch "foo bar" {
    case "a":
        var v26 uint16 = 32767
        _ = v26
    }
    v27 := v2["hello"]
    for fi28 := 0; fi28 < 20; fi28++ {
        for fi29 := 0; fi29 < 7; fi29++ {
            v13 += -401
            var v30 uint32 = v5
            _ = len(v2)
            v1 = genFn1((false || false))
            if true {
                var v31 uint8 = (^(v1))
                delete(v4, (^(v30)))
                _ = v31
            }
            var v32 bool = (true && (!((1 < 158) || false)))
            var v33 int64 = 718
            _ = fi29
            _ = v30
            _ = v32
            _ = v33
        }
        if (1 <= 66) {
            var v34 int32 = v13
            _ = v34
        }
        v3 = append(v3, (v27 / ((int(v5)) | 1)))
        _ = fi28
    }
    v2["\n\t"] = v27
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
