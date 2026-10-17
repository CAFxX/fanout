package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 bool }
type GenS2 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint }

func main() {
    v1 := []int{9223372036854775807}
    var v2 uint64 = 847
    v3 := make(map[string]bool)
    v4 := []uint32{65535}
    var v5 int64 = -9223372036854775808
    v2 += 424
    switch "foo bar" {
    case "foo bar":
        if true {
            v5 = v5
            var v6 float64 = 1e10
            fmt.Println(v2, v6, v6)
            _ = v6
        }
    default:
        switch -337 {
        case -9223372036854775808:
            switch -1 {
            case 27:
                fmt.Println(v2, v5)
            default:
                v5 = -1
            }
        default:
            fmt.Println(v5, v5)
        }
    }
    if false {
        v5 = (v5 - (v5 & v5))
    }
    v5 -= 747
    v7 := GenS1{F0: 128, F1: "a", F2: true}
    switch "foo bar" {
    case "foo bar":
        var v8 int8 = -99
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v2 = v2
        _ = fi9
    }
    v10 := GenS1{F0: -9223372036854775808, F1: "a", F2: true}
    _ = len(v4)
    v5 += (v5 << ((254) & 63))
    v2 += (v2 & v2)
    switch 42 {
    case -1:
        v5 = v5
    case 104:
        for fi11 := 0; fi11 < 5; fi11++ {
            v1 = append(v1, 128)
            _ = fi11
        }
    }
    switch "\n\t" {
    case "a":
        var v12 float64 = 1.5
        _ = v12
    case "foo bar":
        switch "a" {
        case "hello":
            if ((((v2 < v2) && true) && (856 == 1)) || false) {
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 uint32 = 913
                    _ = fi13
                    _ = v14
                }
            } else {
                if true {
                    if false {
                        var v15 string = "hello"
                        _ = v15
                    } else {
                        v5 -= 0
                    }
                    for fi16 := 0; fi16 < 19; fi16++ {
                        if true {
                            if (((-2.25 - 1.5) - float64(1e10)) < 1.5) {
                                var v17 int32 = 2147483647
                                v17 -= (v17 % (((v17 >> ((2) & 63))) | 1))
                                v1 = append(v1, (fi16 & fi16))
                                _ = v17
                            } else {
                                for fi18 := 0; fi18 < 10; fi18++ {
                                    switch 499 {
                                    case 42:
                                        if (!false) {
                                            var v19 uint16 = 65535
                                            _ = v19
                                        } else {
                                            v3[("a" + "a")] = false
                                            v5 -= v5
                                        }
                                    case 631:
                                        var v20 int32 = -920
                                        _ = v20
                                    case 32767:
                                        v5 += 530
                                    }
                                    _ = fi18
                                }
                                v2 += v2
                            }
                        }
                        if ("" != "foo bar") {
                            v5 = (v5 / (((v5 % (((-(v5))) | 1))) | 1))
                        }
                        _ = fi16
                    }
                }
                v21 := GenS2{F0: 255, F1: -1, F2: 2}
                fmt.Println(v5, v2)
                _ = v21
            }
        case "":
            switch 32767 {
            case 128:
                var v22 uint8 = 0
                _ = v22
            case 127:
                v23 := v3["hello"]
                _ = v23
            case 0:
                v24 := GenS3{F0: 292, F1: 30, F2: 0}
                _ = v24
            default:
                if true {
                    switch 282 {
                    case -2147483648:
                        var v25 int8 = -1
                        _ = v25
                    case -241:
                        for fi26 := 0; fi26 < 12; fi26++ {
                            switch ("a" + "") {
                            case "":
                                if false {
                                    v27 := GenS3{F0: 0, F1: 0, F2: 65535}
                                    if len(v4) > 0 {
                                        v4[0] = 755
                                    }
                                    v5 = (v5 + v5)
                                    _ = v27
                                } else {
                                    var v28 float32 = 0.0
                                    if (false || (!false)) {
                                        fmt.Println(v5, v2, v5)
                                    }
                                    var v29 int64 = (v5 ^ -781)
                                    _ = v28
                                    _ = v29
                                }
                            }
                            _ = fi26
                        }
                    default:
                        switch "hello" {
                        case "foo bar":
                            var v30 uint32 = 2
                            _ = v30
                        case "hello":
                            v2 += (v2 - (v2 * 259))
                        }
                    }
                    if false {
                        fmt.Println(v5, v2)
                    } else {
                        if true {
                            for fi31 := 0; fi31 < 12; fi31++ {
                                fmt.Println(fi31, v2)
                                _ = fi31
                            }
                            var v32 string = "a"
                            _ = v32
                        }
                    }
                } else {
                    v5 -= (v5 / ((-143) | 1))
                    if true {
                        var v33 string = "a"
                        _ = v33
                    }
                    var v34 uint32 = 273
                    _ = v34
                }
            }
        case "a":
            v2 = (v2 * 931)
        }
    }
    var v35 uint8 = 128
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v7
}
