package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint32, p1 uint, p2 int8) uint32 {
    var v7 int32 = 0
    var v8 int32 = (^(v7))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p0 >> ((878) & 63))
}
func genFn1() uint64 {
    var v9 uint32 = genFn0(65535, 2, 1)
    var v10 float64 = 3.14159
    if false {
        var v11 int64 = -735
        v9 -= v9
        _ = v11
    } else {
        var v12 bool = true
        _ = v12
    }
    v10 -= 1e10
    switch "\n\t" {
    case "foo bar":
        if false {
            v10 += float64(0.1)
        }
    }
    _ = v10
    _ = v9
    return 106
}

func main() {
    var v1 int8 = 1
    var v2 uint = 42
    var v3 float64 = 1e10
    v4 := []int{9223372036854775807, 0, 32767, 42}
    v5 := make(map[string]uint32)
    v6 := make(map[int]string)
    for fi13 := 0; fi13 < 4; fi13++ {
        v3 += v3
        _ = fi13
    }
    v1 = (v1 * v1)
    if false {
        for fi14 := 0; fi14 < 10; fi14++ {
            for fi15 := 0; fi15 < 13; fi15++ {
                if (false && (42 >= 2147483647)) {
                    if true {
                        var v16 bool = true
                        v16 = (!false)
                        v1 = v1
                        _ = v16
                    } else {
                        if ("" < "") {
                            if true {
                                v2 = (v2 ^ v2)
                                fmt.Println(v1)
                                _ = len(v4)
                            }
                            var v17 int = (fi15 & (fi14 & (fi14 >> ((v2) & 63))))
                            switch v17 {
                            case -9223372036854775808:
                                var v18 int = (^(fi14))
                                _ = v18
                            case 0:
                                v1 = v1
                                v17 = (fi15 - 127)
                                v19 := GenS1{F0: 0, F1: (v1 << (((v2 ^ v2)) & 63))}
                                _ = v19
                            }
                            _ = v17
                        } else {
                            v2 = (v2 & 4294967295)
                            var v20 uint = (^(v2))
                            _ = v20
                        }
                        v6[9223372036854775807] = "foo bar"
                    }
                }
                v5[""] = 2147483647
                v21 := GenS1{F0: 2147483648, F1: (^(v1))}
                fmt.Println(fi15, v3, v3)
                _ = fi15
                _ = v21
            }
            v22 := GenS1{F0: 106, F1: v1}
            if false {
                v2 -= (v2 ^ (v2 >> (((v2 & 1)) & 63)))
            }
            _ = fi14
            _ = v22
        }
    }
    switch 2147483647 {
    case 0:
        var v23 int64 = -380
        _ = v23
    case 355:
        v4 = append(v4, 42)
    default:
        _ = len(v6)
    }
    fmt.Println(v3, v1, v3)
    var v24 uint = v2
    var v25 uint32 = 2147483648
    for fi26 := 0; fi26 < 5; fi26++ {
        if true {
            v4 = append(v4, (-(fi26)))
            v1 = v1
        }
        var v27 float64 = 3.14159
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
