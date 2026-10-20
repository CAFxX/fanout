package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 bool; F2 int }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint }
func genFn0(p0 bool, p1 uint16) int32 {
    switch "\n\t" {
    case "\n\t":
        switch ("" + "\n\t") {
        case "foo bar":
            p0 = true
        case "\n\t":
            switch "a" {
            case "a":
                p0 = (212 <= 2)
            default:
                v5 := []int{171, 4294967295}
                _ = v5
            }
        }
    case "":
        switch -311 {
        case 4294967295:
            var v6 bool = (0.1 < float32(3.14159))
            _ = v6
        case -2147483648:
            if (42 < 9223372036854775807) {
                var v7 int8 = 1
                var v8 int32 = 1
                v9 := []int{-914, 594}
                _ = v7
                _ = v8
                _ = v9
            }
        case 0:
            fmt.Println(p0)
        default:
            var v10 int64 = 0
            _ = v10
        }
    case "a":
        p0 = p0
    }
    v11 := GenS1{F0: -1, F1: 482, F2: 4294967295}
    p0 = (196 != 0)
    var v12 float64 = 1e10
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return -845
}

func main() {
    v1 := []int{-2147483648, -956}
    var v2 uint16 = 2
    v3 := make(map[uint32]bool)
    var v4 int = 2147483647
    if false {
        var v13 int = (v4 % ((v4) | 1))
        _ = v13
    }
    var v14 uint64 = 42
    var v15 uint64 = v14
    if true {
        if ((!false) && true) {
            if (4294967295 <= 1) {
                switch v4 {
                case 626:
                    v4 = v4
                case -1:
                    v16 := GenS2{F0: 695, F1: 42}
                    _ = v16
                case 549:
                    v17 := GenS1{F0: 42, F1: 540, F2: (v4 | (v4 * 4294967295))}
                    _ = v17
                }
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v18 := GenS3{F0: 701}
                _ = v18
            }
        } else {
            switch "hello" {
            case "foo bar":
                switch 127 {
                case 42:
                    fmt.Println(v14, v4, v14)
                case 4294967295:
                    if (!true) {
                        for fi19 := 0; fi19 < 2; fi19++ {
                            var v20 uint32 = 2
                            _ = fi19
                            _ = v20
                        }
                        v21 := GenS2{F0: 853, F1: 1}
                        _ = v21
                    } else {
                        if true {
                            for fi22 := 0; fi22 < 20; fi22++ {
                                v2 += 463
                                _ = fi22
                            }
                            for fi23 := 0; fi23 < 4; fi23++ {
                                delete(v3, 2)
                                _ = fi23
                            }
                            fmt.Println(v14, v4, v14)
                        }
                        v2 += (v2 & 560)
                    }
                case 1:
                    var v24 uint32 = 2
                    _ = v24
                default:
                    var v25 uint = 18446744073709551615
                    _ = v25
                }
            }
            for fi26 := 0; fi26 < 14; fi26++ {
                var v27 uint = 1
                fmt.Println(v4, v15, v14)
                for fi28 := 0; fi28 < 20; fi28++ {
                    if false {
                        v4 = (fi26 + 32768)
                        v14 = (v14 / ((v15) | 1))
                        v29 := GenS1{F0: -2147483648, F1: (v15 % ((v15) | 1)), F2: v4}
                        _ = v29
                    }
                    delete(v3, 0)
                    _ = fi28
                }
                _ = fi26
                _ = v27
            }
        }
        for fi30 := 0; fi30 < 18; fi30++ {
            _ = len(v1)
            for fi31 := 0; fi31 < 17; fi31++ {
                v15 += (v15 << ((359) & 63))
                var v32 bool = true
                var v33 uint8 = 0
                _ = fi31
                _ = v32
                _ = v33
            }
            if (!true) {
                v15 -= (v15 % (((v14 / ((v15) | 1))) | 1))
                v34 := GenS4{F0: 0, F1: true, F2: -48}
                _ = v34
            }
            _ = fi30
        }
    } else {
        for fi35 := 0; fi35 < 15; fi35++ {
            if ("foo bar" == "foo bar") {
                for fi36 := 0; fi36 < 7; fi36++ {
                    v14 -= 2
                    v14 = (v14 << ((65535) & 63))
                    switch "foo bar" {
                    case "":
                        delete(v3, 42)
                    case "hello":
                        fmt.Println(fi36, fi35, v15)
                    }
                    _ = fi36
                }
            }
            var v37 uint16 = v2
            _ = fi35
            _ = v37
        }
        for fi38 := 0; fi38 < 6; fi38++ {
            for fi39 := 0; fi39 < 18; fi39++ {
                var v40 uint8 = 29
                v40 -= (v40 << ((0) & 63))
                fmt.Println(v4, v15, fi39)
                _ = fi39
                _ = v40
            }
            _ = fi38
        }
        var v41 uint16 = (^(v2))
        _ = v41
    }
    v14 = (v15 << ((191) & 63))
    v2 = (v2 ^ 249)
    if false {
        v15 -= v14
    }
    v2 = (v2 ^ 0)
    var v42 uint64 = v15
    v4 -= v4
    v43 := GenS5{F0: 544, F1: 1, F2: 464}
    var v44 int64 = 958
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
}
