package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int64; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8; F1 string; F2 string }
type GenS9 struct { F0 uint; F1 int16; F2 int64 }
func genFn0(p0 uint) uint64 {
    v5 := []int{}
    v6 := GenS1{F0: 32767}
    fmt.Println(p0)
    _ = p0
    _ = v5
    _ = v6
    return uint64(p0)
}
func genFn1(p0 int32) uint {
    var v7 uint32 = 2147483648
    for fi8 := 0; fi8 < 10; fi8++ {
        switch "" {
        case "\n\t":
            var v9 int32 = -916
            _ = v9
        default:
            v7 += v7
        }
        v10 := GenS2{F0: 50}
        _ = fi8
        _ = v10
    }
    fmt.Println(p0, v7)
    v11 := GenS3{F0: 100, F1: 778, F2: 1}
    _ = p0
    _ = v11
    _ = v7
    return 667
}

func main() {
    v1 := []uint32{344}
    var v2 int16 = 1000
    v3 := []int{42}
    var v4 uint64 = 207
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 bool = true
        if true {
            v2 = (v2 >> ((genFn1(42)) & 63))
            var v14 int64 = -104
            var v15 uint16 = 65535
            _ = v14
            _ = v15
        } else {
            if true {
                fmt.Println(v2, v4)
                v16 := GenS4{F0: 599}
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 13; fi17++ {
            v4 = (v4 << ((2) & 63))
            _ = fi17
        }
        _ = fi12
        _ = v13
    }
    if ((0 > 127) || (!(!("" < ("foo bar" + "a"))))) {
        var v18 bool = (false || (!true))
        fmt.Println(v2, v2, v18)
        if false {
            if (599 < 65535) {
                if (0.1 < 1.5) {
                    v19 := GenS3{F0: -128, F1: 9223372036854775807, F2: 42}
                    var v20 int8 = int8(v2)
                    v4 = genFn0(660)
                    _ = v19
                    _ = v20
                } else {
                    fmt.Println(v18)
                }
            }
            if (9223372036854775807 != -2147483648) {
                switch -188 {
                case -1:
                    v21 := GenS5{F0: 4294967295}
                    _ = v21
                }
                for fi22 := 0; fi22 < 13; fi22++ {
                    v2 = -32768
                    _ = fi22
                }
            } else {
                for fi23 := 0; fi23 < 10; fi23++ {
                    if (v18 || (-9223372036854775808 > 1)) {
                        var v24 uint32 = 0
                        if (((0.1 * float64(0.1)) + 0.1) < -2.25) {
                            switch (fi23 + fi23) {
                            case 32767:
                                var v25 int16 = (^(v2))
                                _ = v25
                            case 9223372036854775807:
                                v18 = v18
                            case 42:
                                fmt.Println(v24)
                            default:
                                v2 -= v2
                            }
                        }
                        _ = v24
                    } else {
                        v2 = (v2 | -32768)
                        v2 -= 1
                        var v26 int16 = v2
                        _ = v26
                    }
                    _ = fi23
                }
            }
            if ((!("\n\t" == "")) && v18) {
                for fi27 := 0; fi27 < 20; fi27++ {
                    for fi28 := 0; fi28 < 14; fi28++ {
                        _ = len(v1)
                        _ = fi28
                    }
                    _ = fi27
                }
                v29 := GenS5{F0: 127}
                v30 := GenS6{F0: v18}
                _ = v29
                _ = v30
            }
        } else {
            fmt.Println(v4, v18, v4)
            v1 = append(v1, 2147483647)
            fmt.Println(v18)
        }
        _ = v18
    }
    v31 := GenS7{F0: (v4 % ((981) | 1))}
    for fi32 := 0; fi32 < 2; fi32++ {
        v4 -= (v4 - (v4 * 4294967295))
        switch (fi32 << ((608) & 63)) {
        case -949:
            switch "\n\t" {
            case "\n\t":
                switch "foo bar" {
                case "a":
                    v4 -= 282
                case "":
                    if ((false && true) || false) {
                        var v33 string = "\n\t"
                        _ = v33
                    }
                }
            case "":
                v4 = (v4 / (((^(v4))) | 1))
            case "a":
                v2 = int16(fi32)
            default:
                var v34 uint32 = 821
                _ = v34
            }
        case -1:
            v2 += -688
        case -335:
            v35 := GenS8{F0: 42, F1: "a", F2: "a"}
            _ = v35
        }
        v36 := GenS9{F0: 0, F1: -1, F2: 0}
        _ = fi32
        _ = v36
    }
    var v37 bool = false
    _ = len(v1)
    var v38 float64 = 3.14159
    var v39 uint64 = (v4 % ((genFn0(907)) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
