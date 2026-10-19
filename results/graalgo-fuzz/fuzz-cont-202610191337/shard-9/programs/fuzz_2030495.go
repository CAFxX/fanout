package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int64 }
func genFn0(p0 int8) uint {
    switch 128 {
    case 4294967295:
        for fi5 := 0; fi5 < 13; fi5++ {
            p0 = p0
            var v6 int32 = int32(p0)
            var v7 int8 = p0
            _ = fi5
            _ = v6
            _ = v7
        }
    default:
        if false {
            p0 += 1
            if false {
                var v8 int16 = 1000
                v9 := []int{32767, 206, -1, -1}
                _ = v8
                _ = v9
            } else {
                var v10 uint = 4294967295
                if (0 == (p0 / (((^(p0))) | 1))) {
                    var v11 uint32 = 787
                    switch "a" {
                    case "a":
                        v12 := GenS1{F0: 42}
                        _ = v12
                    case "hello":
                        v13 := GenS1{F0: 431}
                        _ = v13
                    default:
                        v14 := []int{-2147483648, 32767}
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 12; fi15++ {
                        v16 := GenS1{F0: 189}
                        _ = fi15
                        _ = v16
                    }
                    _ = v11
                } else {
                    p0 = (p0 ^ (p0 * (p0 & 42)))
                    fmt.Println(v10, v10)
                }
                v17 := GenS1{F0: 162}
                _ = v10
                _ = v17
            }
            p0 += p0
        } else {
            v18 := []int{1, 42, 127, 127}
            _ = v18
        }
    }
    v19 := []int{127, 4294967295, -997, 1}
    _ = p0
    _ = v19
    return 4294967295
}

func main() {
    v1 := []int{}
    var v2 int16 = 1000
    var v3 int8 = 1
    v4 := []int{128}
    if (764 > 0) {
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 float64 = 1.5
            _ = fi20
            _ = v21
        }
    }
    v22 := GenS2{F0: 2}
    if ((2 < 4294967295) || true) {
        fmt.Println(v3, v2)
    } else {
        var v23 float32 = (3.14159 / (1.5 - 0.0))
        v24 := GenS2{F0: 0}
        var v25 uint16 = 859
        _ = v23
        _ = v24
        _ = v25
    }
    var v26 int8 = (v3 * (v3 | v3))
    switch "" {
    case "foo bar":
        switch "\n\t" {
        case "":
            for fi27 := 0; fi27 < 13; fi27++ {
                v28 := GenS3{F0: v26, F1: -9223372036854775808}
                if false {
                    var v29 string = ("foo bar" + "\n\t")
                    var v30 int8 = (v3 << ((2) & 63))
                    _ = v29
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
        }
    }
    v26 = v3
    var v31 uint32 = 0
    for fi32 := 0; fi32 < 14; fi32++ {
        switch "foo bar" {
        case "foo bar":
            var v33 uint64 = 18446744073709551615
            _ = v33
        case "hello":
            if ((false || (2 < 65535)) && (v26 < v3)) {
                var v34 float32 = 0.0
                _ = v34
            } else {
                var v35 string = "hello"
                var v36 float64 = 0.0
                _ = v35
                _ = v36
            }
        case "":
            if false {
                if (-503 != 0) {
                    switch 128 {
                    case 2147483647:
                        if len(v1) > 0 {
                            v1[0] = (fi32 & fi32)
                        }
                    case 32768:
                        v31 = v31
                    default:
                        for fi37 := 0; fi37 < 13; fi37++ {
                            v3 = v26
                            fmt.Println(fi32, v3, fi32)
                            _ = fi37
                        }
                    }
                }
            }
        }
        _ = fi32
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        for fi39 := 0; fi39 < 11; fi39++ {
            var v40 float32 = -2.25
            v3 -= (v26 % ((42) | 1))
            var v41 bool = (42 < 0)
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = fi38
    }
    var v42 float32 = 0.0
    fmt.Println(v31)
    switch -950 {
    case 1:
        _ = len(v1)
    default:
        v43 := GenS1{F0: 65535}
        _ = v43
    }
    for fi44 := 0; fi44 < 4; fi44++ {
        var v45 float32 = v42
        fmt.Println(v2)
        v45 = v45
        _ = fi44
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v42
}
