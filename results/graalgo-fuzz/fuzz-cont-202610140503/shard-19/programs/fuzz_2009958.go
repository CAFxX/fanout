package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint) int16 {
    var v7 int8 = 0
    var v8 uint32 = 833
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := []int{(^(fi9)), int(p0), fi9, (fi9 ^ fi9)}
        _ = fi9
        _ = v10
    }
    if (!(0 >= 65535)) {
        p0 = (p0 ^ p0)
        if false {
            var v11 string = "foo bar"
            var v12 uint16 = 65535
            _ = v11
            _ = v12
        } else {
            var v13 int = 1
            v14 := GenS1{F0: -32768, F1: 211}
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS2{F0: p0}
        var v17 int32 = -1
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := []uint32{65535, 4294967295, 2147483648}
    v2 := make(map[int]string)
    v3 := []string{"hello", ("a" + "\n\t"), "\n\t"}
    v4 := make(map[int]int)
    v5 := make(map[string]uint32)
    v6 := make(map[string]bool)
    var v18 int16 = -89
    v18 = (v18 & 701)
    fmt.Println(v18)
    if true {
        v19 := GenS1{F0: -1, F1: 32767}
        fmt.Println(v18)
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 uint = 0
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    v18 = genFn0(1, 892)
    if true {
        switch 32768 {
        case 4294967295:
            var v22 string = ""
            _ = v22
        case 40:
            var v23 uint32 = 2147483647
            _ = v23
        case -9223372036854775808:
            var v24 float32 = 1e10
            _ = v24
        default:
            switch 127 {
            case 4294967295:
                if ((9223372036854775807 >= -1) && false) {
                    v18 = -977
                } else {
                    v18 = (v18 + (v18 << ((959) & 63)))
                    if false {
                        v25 := GenS2{F0: 1}
                        _ = v25
                    } else {
                        switch "" {
                        case "hello":
                            v18 = v18
                        case "foo bar":
                            if false {
                                if true {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    fmt.Println(v18)
                                }
                                var v26 uint32 = 1
                                _ = v26
                            }
                        case "a":
                            var v27 uint64 = 18446744073709551615
                            _ = v27
                        }
                        switch -2147483648 {
                        case 2147483647:
                            v18 = v18
                        default:
                            v1 = append(v1, 809)
                        }
                    }
                    v3 = append(v3, "\n\t")
                }
            default:
                v18 += v18
            }
        }
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        for fi29 := 0; fi29 < 12; fi29++ {
            if true {
                fmt.Println(fi29, fi29, v18)
            } else {
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
                if false {
                    fmt.Println(fi29, fi29, v18)
                    v18 -= -478
                    if false {
                        v18 -= v18
                    }
                }
                switch (fi29 & (fi28 % ((fi29) | 1))) {
                case 0:
                    var v30 uint = 2
                    _ = v30
                case 2147483647:
                    v18 += genFn0(421, 1)
                }
            }
            _ = fi29
        }
        v18 += -32768
        fmt.Println(v18)
        _ = fi28
    }
    var v31 float64 = (0.1 * 3.14159)
    fmt.Println(v31, v31)
    fmt.Println(v31, v18)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
