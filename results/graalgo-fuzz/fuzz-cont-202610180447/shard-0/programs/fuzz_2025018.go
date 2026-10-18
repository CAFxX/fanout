package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint64; F1 int; F2 bool }
type GenS4 struct { F0 int; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint32; F1 uint; F2 bool }
type GenS6 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS7 struct { F0 uint8; F1 string }
func genFn0(p0 int16) int {
    v5 := GenS1{F0: 1}
    for fi6 := 0; fi6 < 3; fi6++ {
        switch (fi6 - (fi6 - fi6)) {
        case -9223372036854775808:
            fmt.Println(fi6, p0)
        case 127:
            v7 := []int{4294967295, (fi6 * -1), (fi6 ^ (fi6 ^ fi6))}
            _ = v7
        }
        if true {
            fmt.Println(p0, fi6)
            p0 = p0
        }
        for fi8 := 0; fi8 < 18; fi8++ {
            p0 -= (p0 + p0)
            v9 := GenS1{F0: 185}
            fmt.Println(fi8)
            _ = fi8
            _ = v9
        }
        _ = fi6
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 bool = ((false || (-1 == 0)) && true)
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v5
    return -960
}
func genFn1(p0 int64, p1 int, p2 string) int8 {
    v12 := GenS2{F0: 42, F1: 0}
    var v13 int16 = -1
    v14 := []int{(p1 & (p1 * (p1 * (p1 * (-(p1)))))), (-(p1)), genFn0((v13 | (^(v13)))), 0}
    fmt.Println(v13, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    return -116
}

func main() {
    var v1 uint = 4294967295
    var v2 string = ("hello" + "a")
    v3 := make(map[uint32]int)
    var v4 uint32 = 178
    var v15 int = genFn0(32767)
    if (-9223372036854775808 < 263) {
        if (true && (!(true && true))) {
            v16 := []int{genFn0(-1), genFn0(32767)}
            v15 -= v15
            _ = v16
        }
        v17 := []int{v15, -9223372036854775808, (v15 * v15)}
        _ = v17
    } else {
        for fi18 := 0; fi18 < 5; fi18++ {
            fmt.Println(v4, fi18)
            if ((v4 < (v4 - 1)) || (v4 != (v4 - (v4 << (((v1 * 535)) & 63))))) {
                v1 += v1
                var v19 float64 = (0.1 / 1.0)
                for fi20 := 0; fi20 < 14; fi20++ {
                    if true {
                        if true {
                            var v21 float64 = v19
                            fmt.Println(v19, v21)
                            v22 := []int{-901}
                            _ = v21
                            _ = v22
                        }
                        v23 := v3[v4]
                        switch genFn0(-117) {
                        case -315:
                            _ = len(v3)
                        case 631:
                            v24 := []int{(v15 ^ (v23 % ((4294967295) | 1))), 1, v15}
                            _ = v24
                        default:
                            v19 -= v19
                        }
                        _ = v23
                    } else {
                        fmt.Println(v19)
                    }
                    v25 := GenS3{F0: 0, F1: genFn0(32767), F2: false}
                    v26 := GenS4{F0: -1, F1: uint8(v4), F2: genFn1(-1, genFn0(0), ("" + v2))}
                    var v27 uint16 = 2
                    _ = fi20
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = v19
            }
            var v28 int64 = -9223372036854775808
            _ = fi18
            _ = v28
        }
        v29 := GenS5{F0: 990, F1: (v1 | (^(v1))), F2: true}
        _ = v29
    }
    var v30 int = (v15 + (v15 * genFn0(-1)))
    v31 := GenS6{F0: (v1 + (v1 >> ((v1) & 63))), F1: 205, F2: 2}
    v1 -= v1
    v4 += uint32(v1)
    v4 = (v4 ^ (v4 % (((v4 | (v4 ^ (v4 * 65535)))) | 1)))
    for fi32 := 0; fi32 < 10; fi32++ {
        var v33 int64 = 127
        _ = fi32
        _ = v33
    }
    v30 = 0
    var v34 int16 = 899
    if true {
        for fi35 := 0; fi35 < 17; fi35++ {
            v4 = v4
            for fi36 := 0; fi36 < 9; fi36++ {
                v30 = int(v34)
                _ = fi36
            }
            v37 := []int{v15, 274, (v15 / (((v15 >> ((v1) & 63))) | 1))}
            _ = fi35
            _ = v37
        }
        for fi38 := 0; fi38 < 7; fi38++ {
            v1 = 0
            fmt.Println(v2)
            _ = fi38
        }
    } else {
        var v39 uint = 18446744073709551615
        v40 := GenS5{F0: (v4 * (v4 * v4)), F1: 950, F2: true}
        var v41 uint = (^(v39))
        _ = v39
        _ = v40
        _ = v41
    }
    var v42 int = int(v4)
    for fi43 := 0; fi43 < 19; fi43++ {
        if true {
            var v44 int8 = 79
            var v45 string = (v2 + v2)
            fmt.Println(v2, v44, v44)
            _ = v44
            _ = v45
        } else {
            var v46 uint32 = 65535
            var v47 bool = false
            v48 := GenS7{F0: 0, F1: ""}
            _ = v46
            _ = v47
            _ = v48
        }
        v4 = (v4 % (((v4 & 1)) | 1))
        v49 := GenS5{F0: (v4 | v4), F1: v1, F2: (false || true)}
        _ = fi43
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v42
}
