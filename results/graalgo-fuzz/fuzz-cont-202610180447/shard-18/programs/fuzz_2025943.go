package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
func genFn0(p0 uint16, p1 int16, p2 bool) int64 {
    switch "a" {
    case "foo bar":
        v7 := []int{32768, 1, -754, 2147483647}
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 10; fi8++ {
            p0 += 42
            _ = fi8
        }
    case "hello":
        var v9 uint16 = (p0 + p0)
        _ = v9
    }
    p0 = (p0 ^ p0)
    var v10 int32 = 2147483647
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 651
}
func genFn1() uint16 {
    var v11 int = 59
    var v12 int8 = -128
    var v13 int32 = 0
    _ = v11
    _ = v12
    _ = v13
    return 2
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]string)
    var v3 int8 = 1
    v4 := make(map[int]uint32)
    var v5 uint32 = 318
    v6 := make(map[string]int)
    v14 := GenS1{F0: 42, F1: 9223372036854775807}
    v15 := []int{4294967295}
    fmt.Println(v1, v1)
    var v16 uint8 = 42
    switch int(v16) {
    case 838:
        var v17 string = "\n\t"
        _ = v17
    }
    if false {
        v3 += v3
    } else {
        v5 += (v5 | (v5 & (^(v5))))
        if false {
            for fi18 := 0; fi18 < 13; fi18++ {
                v15 = append(v15, (-(fi18)))
                _ = fi18
            }
            fmt.Println(v3, v3)
        } else {
            fmt.Println(v16, v16)
        }
        var v19 int8 = v3
        _ = v19
    }
    var v20 uint32 = 2
    switch 75 {
    case -9223372036854775808:
        for fi21 := 0; fi21 < 8; fi21++ {
            if ("" > "") {
                var v22 uint16 = 32767
                v4[(fi21 / (((fi21 * 32768)) | 1))] = (^(v5))
                _ = v22
            }
            _ = fi21
        }
    case 9223372036854775807:
        v5 += 2
    default:
        _ = len(v2)
    }
    var v23 uint16 = 65535
    switch "" {
    case "\n\t":
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    default:
        v20 = (v20 << ((0) & 63))
    }
    if false {
        var v24 uint16 = v23
        var v25 int16 = 220
        v20 = (^(v20))
        _ = v24
        _ = v25
    } else {
        if false {
            v16 += (v16 - v16)
        }
        var v26 int8 = v3
        _ = v26
    }
    switch 2147483647 {
    case 42:
        var v27 bool = (float32(0.1) != 0.0)
        _ = v27
    default:
        for fi28 := 0; fi28 < 10; fi28++ {
            if true {
                var v29 int32 = -783
                if (!((v16 + 255) <= v16)) {
                    _ = len(v15)
                    v29 += v29
                    fmt.Println(v20, v23)
                } else {
                    var v30 int = int(v29)
                    v23 += v23
                    _ = v30
                }
                _ = v29
            } else {
                v31 := GenS1{F0: 9223372036854775807, F1: fi28}
                _ = v31
            }
            fmt.Println(v1)
            v20 += 752
            _ = fi28
        }
    }
    _ = len(v15)
    v5 = uint32(v23)
    var v32 bool = ((((v3 + (v3 * 41)) <= v3) && (485 > genFn0(v23, -390, false))) && false)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
