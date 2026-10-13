package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint16 }
func genFn0(p0 bool) uint16 {
    v5 := GenS1{F0: 18446744073709551615, F1: 2147483648}
    p0 = p0
    if p0 {
        v6 := []int{-1, -890, 32767, -2147483648}
        for fi7 := 0; fi7 < 9; fi7++ {
            p0 = true
            _ = fi7
        }
        _ = v6
    } else {
        var v8 uint32 = 2
        p0 = p0
        p0 = (p0 && p0)
        _ = v8
    }
    _ = p0
    _ = v5
    return 1
}
func genFn1(p0 bool) uint64 {
    var v9 int16 = 1000
    v10 := GenS2{F0: 791, F1: -1}
    v11 := []int{42, -125, 1}
    v9 += v9
    p0 = (32767 > 260)
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 441
}

func main() {
    var v1 uint = 203
    v2 := make(map[uint32]int)
    var v3 bool = true
    v4 := make(map[string]string)
    switch "" {
    case "a":
        fmt.Println(v1, v1)
    case "":
        v12 := []int{-2147483648, 42}
        _ = v12
    default:
        v13 := GenS3{F0: 255, F1: 255}
        _ = v13
    }
    v1 += v1
    var v14 uint32 = 594
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 bool = false
        if ((!(!v16)) && false) {
            for fi17 := 0; fi17 < 6; fi17++ {
                fmt.Println(fi15, v16, v14)
                _ = fi17
            }
        } else {
            v14 += v14
            v18 := []int{(fi15 & 2147483647), fi15}
            _ = v18
        }
        var v19 uint32 = v14
        _ = fi15
        _ = v16
        _ = v19
    }
    v14 += 65535
    v1 += v1
    if v3 {
        if false {
            v20 := []int{-1, 4294967295, 902}
            v4["\n\t"] = "foo bar"
            var v21 int8 = 0
            _ = v20
            _ = v21
        }
        var v22 int = 128
        _ = v22
    } else {
        v14 = (v14 | v14)
        switch -1 {
        case 2147483647:
            if v3 {
                _ = len(v2)
                var v23 uint = 121
                _ = v23
            } else {
                v1 = (v1 & 1)
                v4[("" + ("hello" + ""))] = "a"
            }
        case 128:
            v1 = v1
        case 9223372036854775807:
            for fi24 := 0; fi24 < 15; fi24++ {
                switch "hello" {
                case "":
                    v3 = v3
                default:
                    if false {
                        v3 = (v14 != 2147483648)
                        var v25 uint32 = (^(v14))
                        var v26 uint16 = 0
                        _ = v25
                        _ = v26
                    }
                }
                _ = fi24
            }
        default:
            v1 -= v1
        }
        _ = len(v2)
    }
    switch 32768 {
    case 1:
        for fi27 := 0; fi27 < 7; fi27++ {
            v1 += v1
            v1 = (v1 % ((2) | 1))
            delete(v4, ("hello" + "\n\t"))
            _ = fi27
        }
    }
    fmt.Println(v3, v1, v3)
    for fi28 := 0; fi28 < 9; fi28++ {
        switch "hello" {
        case "\n\t":
            v14 -= (v14 / (((v14 & 494)) | 1))
        }
        if ((42 > 42) && ((v3 && (!(genFn0((1 <= -1)) >= 264))) && v3)) {
            fmt.Println(v1, v1)
        } else {
            for fi29 := 0; fi29 < 8; fi29++ {
                var v30 float64 = 0.1
                _ = fi29
                _ = v30
            }
        }
        _ = fi28
    }
    var v31 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
