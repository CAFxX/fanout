package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint32 }
func genFn0(p0 string) int32 {
    fmt.Println(p0)
    switch (p0 + p0) {
    case "a":
        v5 := GenS1{F0: 50}
        _ = v5
    case "hello":
        v6 := GenS1{F0: 0}
        _ = v6
    case "foo bar":
        if (1 >= 763) {
            var v7 int64 = 1
            for fi8 := 0; fi8 < 15; fi8++ {
                if true {
                    if true {
                        var v9 uint64 = 18446744073709551615
                        var v10 int32 = 2147483647
                        if true {
                            v11 := []int{}
                            _ = v11
                        }
                        _ = v10
                        _ = v9
                    }
                    var v12 int32 = -1
                    switch "\n\t" {
                    case "hello":
                        v13 := GenS1{F0: 0}
                        _ = v13
                    }
                    _ = v12
                } else {
                    v14 := []int{}
                    _ = v14
                }
                switch fi8 {
                case 9223372036854775807:
                    switch (fi8 ^ (fi8 ^ fi8)) {
                    case 2147483647:
                        v15 := []int{128, (fi8 >> ((18446744073709551615) & 63)), (^(fi8)), (fi8 / ((fi8) | 1))}
                        _ = v15
                    case 32768:
                        if ((p0 != "\n\t") || false) {
                            v7 -= v7
                            p0 = p0
                            if (!false) {
                                switch p0 {
                                case "":
                                    v7 += (v7 & (v7 ^ -331))
                                case "a":
                                    v16 := []int{978, fi8}
                                    _ = v16
                                case "hello":
                                    p0 = ""
                                default:
                                    switch fi8 {
                                    case -277:
                                        v7 = v7
                                    }
                                }
                                v17 := []int{4294967295, fi8, fi8, fi8}
                                _ = len(v17)
                                _ = v17
                            } else {
                                if false {
                                    v7 += 541
                                } else {
                                    var v18 uint8 = 42
                                    v19 := []int{(fi8 & 742), fi8}
                                    _ = v18
                                    _ = v19
                                }
                            }
                        } else {
                            v20 := GenS2{F0: v7}
                            _ = v20
                        }
                    case 32767:
                        v21 := []int{}
                        _ = v21
                    default:
                        fmt.Println(fi8)
                    }
                case -374:
                    p0 = p0
                case 32768:
                    v22 := []int{fi8, 0}
                    _ = v22
                }
                _ = fi8
            }
            switch 778 {
            case 2147483647:
                if false {
                    var v23 uint = 18446744073709551615
                    var v24 float64 = 0.1
                    _ = v23
                    _ = v24
                }
            case 714:
                p0 = p0
            }
            _ = v7
        }
    }
    switch 1 {
    case 0:
        for fi25 := 0; fi25 < 8; fi25++ {
            fmt.Println(fi25)
            p0 = "foo bar"
            _ = fi25
        }
    case 32768:
        var v26 int32 = 42
        _ = v26
    case 42:
        p0 = p0
    }
    _ = p0
    return 2147483647
}
func genFn1(p0 int32) int32 {
    if true {
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 uint16 = 308
            fmt.Println(fi27)
            for fi29 := 0; fi29 < 14; fi29++ {
                v28 -= v28
                _ = fi29
            }
            _ = fi27
            _ = v28
        }
        if ((3.14159 + (0.1 - 3.14159)) > 1e10) {
            var v30 int8 = 1
            _ = v30
        } else {
            var v31 float32 = 0.1
            for fi32 := 0; fi32 < 20; fi32++ {
                v31 = 0.0
                _ = fi32
            }
            switch -2147483648 {
            case 127:
                p0 -= p0
            default:
                v33 := GenS3{F0: 4294967295}
                _ = v33
            }
            _ = v31
        }
    }
    var v34 uint32 = 203
    if ((false && (!true)) || (2147483647 < 4294967295)) {
        p0 += p0
        var v35 uint32 = v34
        fmt.Println(v34)
        _ = v35
    }
    _ = p0
    _ = v34
    return p0
}

func main() {
    var v1 int32 = 2147483647
    v2 := []int{4294967295, 42}
    var v3 uint = 0
    var v4 uint = (v3 ^ (v3 ^ v3))
    for fi36 := 0; fi36 < 10; fi36++ {
        fmt.Println(v4, fi36)
        for fi37 := 0; fi37 < 6; fi37++ {
            var v38 float32 = (((((-2.25 + float32((0.0 + ((0.1 / 1e10) * 1.5)))) - (3.14159 / 0.1)) / (0.1 / 0.1)) * float32(1.5)) - 0.1)
            if ((!(1 > 255)) && true) {
                var v39 uint32 = 2
                _ = v39
            }
            _ = fi37
            _ = v38
        }
        fmt.Println(fi36)
        _ = fi36
    }
    var v40 uint = (v3 >> ((v4) & 63))
    v40 += v40
    for fi41 := 0; fi41 < 5; fi41++ {
        for fi42 := 0; fi42 < 10; fi42++ {
            v43 := GenS4{F0: v40, F1: 32767, F2: 238}
            _ = fi42
            _ = v43
        }
        _ = fi41
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        v3 -= 18446744073709551615
        var v45 int64 = -9223372036854775808
        _ = fi44
        _ = v45
    }
    fmt.Println(v40, v1, v3)
    var v46 int8 = 100
    if len(v2) > 0 {
        v2[0] = 321
    }
    v2 = append(v2, 32768)
    if len(v2) > 0 {
        v2[0] = 872
    }
    v46 = v46
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v46
}
