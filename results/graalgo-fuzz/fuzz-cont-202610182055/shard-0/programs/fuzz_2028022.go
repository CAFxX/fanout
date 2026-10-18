package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int32 }
type GenS2 struct { F0 string; F1 bool; F2 uint64 }
func genFn0(p0 int32, p1 uint8, p2 int16) int64 {
    var v4 uint32 = 2147483647
    v5 := GenS1{F0: p2, F1: (v4 * v4), F2: (p0 * p0)}
    fmt.Println(p2, p1)
    v6 := []int{-1, 9223372036854775807, 32768}
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return -59
}
func genFn1(p0 int8, p1 int16, p2 int32) int8 {
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint8 = 127
        if false {
            fmt.Println(v8)
            for fi9 := 0; fi9 < 5; fi9++ {
                v10 := []int{fi7, -677, (fi9 & fi9), fi7}
                v11 := GenS1{F0: (p1 * -985), F1: 2147483647, F2: (^(p2))}
                fmt.Println(fi7)
                var v12 uint64 = 530
                var v13 int8 = (p0 + (p0 / ((p0) | 1)))
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
                _ = v13
            }
        } else {
            fmt.Println(p0, p1)
        }
        _ = fi7
        _ = v8
    }
    var v14 int64 = 576
    v15 := []int{-9223372036854775808, int(v14)}
    p2 = (p2 ^ 444)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return p0
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]uint32)
    v3 := []string{("\n\t" + ("" + "foo bar")), ""}
    _ = len(v2)
    v16 := GenS1{F0: 1, F1: 113, F2: 42}
    if true {
        v3 = append(v3, "\n\t")
        fmt.Println(v1)
        v17 := GenS1{F0: -879, F1: 1, F2: -2147483648}
        _ = v17
    } else {
        switch "\n\t" {
        case "foo bar":
            v1 += v1
        case "\n\t":
            v18 := GenS2{F0: "", F1: true, F2: 42}
            _ = v18
        case "":
            for fi19 := 0; fi19 < 19; fi19++ {
                var v20 float32 = (0.1 - 0.1)
                _ = fi19
                _ = v20
            }
        default:
            var v21 int32 = 0
            _ = v21
        }
    }
    _ = len(v2)
    var v22 int = 9223372036854775807
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (!(0 != 0)) {
        for fi23 := 0; fi23 < 19; fi23++ {
            if false {
                v24 := v2["hello"]
                var v25 uint = 65535
                _ = v24
                _ = v25
            }
            v22 = (v22 % ((fi23) | 1))
            _ = fi23
        }
        var v26 int = v22
        _ = v26
    } else {
        if (((v1 * v1) > v1) || false) {
            delete(v2, "\n\t")
            for fi27 := 0; fi27 < 13; fi27++ {
                if false {
                    fmt.Println(v22, v22)
                    var v28 uint8 = v1
                    if false {
                        fmt.Println(v28, fi27, v28)
                    }
                    _ = v28
                }
                _ = fi27
            }
        }
        var v29 int = 851
        _ = v29
    }
    fmt.Println(v22, v22)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v3
}
