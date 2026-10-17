package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint64, p1 uint16) uint32 {
    v4 := []int{437, 32768, -9223372036854775808}
    var v5 uint = 4294967295
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 51
}
func genFn1(p0 int32) string {
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := []int{fi6, -671}
        switch fi6 {
        case 0:
            p0 -= p0
        case -112:
            if true {
                p0 = p0
                p0 = (^(p0))
            }
        case 181:
            var v8 int16 = 1
            _ = v8
        }
        switch fi6 {
        case 32768:
            fmt.Println(fi6, fi6)
        case 141:
            for fi9 := 0; fi9 < 17; fi9++ {
                if true {
                    var v10 float32 = (0.0 * (3.14159 * float32(3.14159)))
                    _ = v10
                } else {
                    var v11 int32 = (p0 | 1)
                    if (9223372036854775807 > -1) {
                        v12 := GenS1{F0: 602, F1: 0}
                        var v13 int8 = 0
                        p0 += p0
                        _ = v12
                        _ = v13
                    }
                    _ = v11
                }
                p0 -= (p0 * 42)
                _ = fi9
            }
        }
        _ = fi6
        _ = v7
    }
    var v14 float32 = -2.25
    _ = p0
    _ = v14
    return ("a" + "\n\t")
}

func main() {
    v1 := []uint32{164}
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v15 int16 = -999
    var v16 int32 = -2147483648
    v16 = 2147483647
    if (-1 <= -2147483648) {
        fmt.Println(v15, v15)
        v17 := GenS1{F0: v15, F1: 255}
        v1 = append(v1, 1)
        _ = v17
    } else {
        var v18 int8 = 30
        _ = v18
    }
    v15 = v15
    switch "foo bar" {
    case "a":
        var v19 string = "hello"
        _ = v19
    default:
        var v20 string = ""
        _ = v20
    }
    v15 = (v15 & (v15 >> ((18446744073709551615) & 63)))
    if true {
        v21 := v3[2147483647]
        _ = v21
    }
    v22 := v3[810]
    v16 += v16
    switch 32768 {
    case 247:
        var v23 int = 42
        _ = v23
    case -1:
        fmt.Println(v16)
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 uint32 = v22
        v26 := GenS1{F0: -922, F1: 0}
        if false {
            var v27 uint32 = v22
            _ = v27
        } else {
            for fi28 := 0; fi28 < 12; fi28++ {
                v22 -= (v25 - v25)
                v25 -= (v22 + (v25 >> ((2) & 63)))
                v25 -= v22
                _ = fi28
            }
            for fi29 := 0; fi29 < 8; fi29++ {
                v30 := GenS1{F0: 42, F1: 65535}
                _ = fi29
                _ = v30
            }
            switch "hello" {
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = v25
                }
            case "a":
                var v31 uint64 = 618
                _ = v31
            }
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
}
