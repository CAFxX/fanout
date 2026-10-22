package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 bool; F1 bool; F2 uint }
type GenS3 struct { F0 uint8; F1 string; F2 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 uint64 }
func genFn0(p0 int, p1 bool) uint {
    if ((p1 && false) || true) {
        switch "" {
        case "hello":
            var v4 uint = 337
            _ = v4
        default:
            v5 := GenS1{F0: 560, F1: p0}
            _ = v5
        }
        v6 := []int{p0}
        _ = v6
    } else {
        var v7 int = p0
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := []int{(fi8 * 42)}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        switch p0 {
        case -44:
            if p1 {
                v10 := GenS2{F0: ((p0 ^ p0) != p0), F1: (152 == 2), F2: 2}
                for fi11 := 0; fi11 < 17; fi11++ {
                    switch fi8 {
                    case 0:
                        p1 = ((!p1) && (p1 || p1))
                    case 806:
                        fmt.Println(p1)
                    default:
                        if (p1 && ((32767 > 519) && p1)) {
                            p1 = (!p1)
                        } else {
                            p1 = p1
                        }
                    }
                    p0 -= fi8
                    switch fi11 {
                    case 42:
                        v9 = append(v9, (fi8 & 32767))
                    case 32768:
                        if len(v9) > 0 {
                            v9[0] = (p0 ^ 331)
                        }
                    case 32767:
                        p0 = fi8
                    }
                    _ = fi11
                }
                v12 := GenS3{F0: 2, F1: "\n\t", F2: 797}
                _ = v10
                _ = v12
            } else {
                p1 = (!false)
            }
        default:
            var v13 uint16 = 32767
            _ = v13
        }
        _ = fi8
        _ = v9
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            fmt.Println(p0)
            _ = fi15
        }
        var v16 int64 = -9223372036854775808
        _ = fi14
        _ = v16
    }
    _ = p0
    _ = p1
    return 65535
}
func genFn1(p0 string, p1 uint64, p2 bool) int16 {
    fmt.Println(p2, p1)
    v17 := GenS2{F0: (true || p2), F1: (1 != 727), F2: 42}
    v18 := GenS4{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return 42
}

func main() {
    var v1 int64 = 42
    var v2 uint8 = 2
    var v3 int64 = v1
    v19 := []int{127, 1, 282, 42}
    if (false || false) {
        v19 = append(v19, -385)
        var v20 uint64 = 1
        _ = v20
    } else {
        fmt.Println(v2, v2, v1)
    }
    v2 -= (v2 & (^(v2)))
    if false {
        var v21 float32 = 1.5
        _ = v21
    }
    var v22 string = ("\n\t" + "\n\t")
    for fi23 := 0; fi23 < 2; fi23++ {
        var v24 int8 = 1
        _ = len(v19)
        _ = fi23
        _ = v24
    }
    if true {
        for fi25 := 0; fi25 < 13; fi25++ {
            if false {
                if true {
                    v26 := GenS5{F0: 580}
                    _ = v26
                } else {
                    var v27 int64 = (v3 << ((18446744073709551615) & 63))
                    _ = v27
                }
                v2 = (v2 * (v2 ^ v2))
                if true {
                    v3 += (v3 * v1)
                    fmt.Println(v3, v2)
                }
            } else {
                var v28 int8 = 100
                if true {
                    fmt.Println(v3)
                } else {
                    for fi29 := 0; fi29 < 16; fi29++ {
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        _ = len(v19)
                        fmt.Println(v2, v22, v1)
                        _ = fi29
                    }
                    switch v22 {
                    case "":
                        var v30 int16 = 1000
                        _ = v30
                    default:
                        var v31 int64 = (v1 % ((9223372036854775807) | 1))
                        _ = v31
                    }
                }
                _ = v28
            }
            _ = fi25
        }
        v22 = "foo bar"
    } else {
        var v32 int = 2147483647
        var v33 uint16 = 255
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 9; fi34++ {
        if true {
            var v35 bool = true
            v1 = (v1 >> ((0) & 63))
            v36 := GenS6{F0: 1, F1: 4294967295}
            _ = v35
            _ = v36
        }
        fmt.Println(v22)
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
}
