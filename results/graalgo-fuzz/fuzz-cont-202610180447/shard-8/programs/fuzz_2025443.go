package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 bool }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 string; F1 uint16 }
func genFn0(p0 uint64) bool {
    v5 := GenS1{F0: 0, F1: 0, F2: false}
    fmt.Println(p0)
    v6 := []int{32767, -529}
    _ = len(v6)
    _ = p0
    _ = v5
    _ = v6
    return false
}
func genFn1(p0 uint, p1 int32, p2 int8) uint32 {
    switch "foo bar" {
    case "foo bar":
        switch -589 {
        case 0:
            v7 := []int{32767}
            _ = v7
        }
    }
    switch 0 {
    case -9223372036854775808:
        p1 += -2147483648
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 1
    var v3 bool = false
    v4 := make(map[int]uint32)
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint16 = 1
        fmt.Println(v3, fi8, v2)
        fmt.Println(v3, v2)
        _ = fi8
        _ = v9
    }
    v3 = ((v2 - (v2 & (v2 + v2))) > v2)
    _ = len(v4)
    switch "\n\t" {
    case "foo bar":
        var v10 int32 = 0
        _ = v10
    case "":
        v3 = v3
    case "a":
        v11 := GenS2{F0: 42, F1: 42}
        _ = v11
    }
    v12 := v4[128]
    switch "a" {
    case "":
        fmt.Println(v3)
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        v3 = v3
    case 127:
        delete(v4, -2147483648)
    case 9223372036854775807:
        if (((531 >= 2) || (genFn1(18446744073709551615, -2147483648, 42) == v12)) || (float64((0.0 / 1.0)) != (1e10 + 0.1))) {
            for fi13 := 0; fi13 < 5; fi13++ {
                fmt.Println(v3)
                fmt.Println(v2, fi13)
                _ = fi13
            }
            v3 = v3
        } else {
            for fi14 := 0; fi14 < 13; fi14++ {
                fmt.Println(v3)
                _ = fi14
            }
            if false {
                if (128 != 128) {
                    v15 := []int{}
                    for fi16 := 0; fi16 < 11; fi16++ {
                        for fi17 := 0; fi17 < 7; fi17++ {
                            var v18 int = (fi16 % ((9223372036854775807) | 1))
                            _ = fi17
                            _ = v18
                        }
                        _ = fi16
                    }
                    fmt.Println(v2)
                    _ = v15
                } else {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        var v20 float32 = 1e10
                        v12 -= (v12 % ((v12) | 1))
                        var v21 int16 = 1000
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    var v22 uint8 = 17
                    _ = v22
                }
                delete(v4, 867)
            }
        }
    default:
        var v23 int = 42
        _ = v23
    }
    v24 := []int{-571, 9223372036854775807}
    for fi25 := 0; fi25 < 14; fi25++ {
        v2 = (v2 & -9223372036854775808)
        _ = fi25
    }
    fmt.Println(v3, v3)
    v2 += v2
    if v3 {
        var v26 int16 = 42
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 uint16 = 32767
            _ = fi27
            _ = v28
        }
        _ = v26
    } else {
        fmt.Println(v2, v3, v3)
        var v29 int64 = v2
        _ = v29
    }
    v12 += v12
    v30 := GenS3{F0: -128}
    if true {
        var v31 int16 = -1
        var v32 uint16 = 0
        _ = v31
        _ = v32
    } else {
        var v33 uint64 = 909
        v2 += v2
        var v34 int8 = 1
        _ = v33
        _ = v34
    }
    v2 = v2
    switch 32767 {
    case 42:
        v3 = (456 >= (v12 % ((genFn1(4294967295, -2147483648, 46)) | 1)))
    case 2147483647:
        var v35 bool = true
        _ = v35
    case 9223372036854775807:
        v12 += v12
    }
    switch "hello" {
    case "":
        v3 = (42 < (v2 ^ (v2 - -148)))
    case "\n\t":
        if ((4294967295 <= 970) || (100 == 127)) {
            if (141 != 0) {
                v36 := GenS4{F0: 339, F1: 9223372036854775807}
                v37 := GenS1{F0: -1, F1: -795, F2: (false && (-2.25 != 3.14159))}
                _ = v36
                _ = v37
            }
            v38 := GenS5{F0: "", F1: 65535}
            _ = v38
        } else {
            for fi39 := 0; fi39 < 13; fi39++ {
                fmt.Println(fi39)
                _ = fi39
            }
            fmt.Println(v2, v12, v3)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v4
}
