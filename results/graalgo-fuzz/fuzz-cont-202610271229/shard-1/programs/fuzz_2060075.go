package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint8, p1 int16, p2 int16) bool {
    for fi4 := 0; fi4 < 14; fi4++ {
        p2 -= (p2 ^ (^(p2)))
        fmt.Println(p0, p1, p1)
        _ = fi4
    }
    p2 += p1
    v5 := []int{423, 4294967295}
    v6 := GenS1{F0: (p0 << ((18446744073709551615) & 63))}
    var v7 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (-2.25 < 3.14159)
}

func main() {
    var v1 int = 42
    v2 := []int{127, 32767, v1, (v1 << ((42) & 63))}
    var v3 int64 = -1
    v3 += (v3 % (((-(v3))) | 1))
    fmt.Println(v1, v3)
    v8 := GenS1{F0: 1}
    v2 = append(v2, (v1 / ((42) | 1)))
    v3 += v3
    switch (v1 * (v1 * (v1 | 127))) {
    case 32768:
        v1 = -141
    case -9223372036854775808:
        var v9 int32 = 0
        _ = v9
    default:
        if len(v2) > 0 {
            v2[0] = (v1 << ((297) & 63))
        }
    }
    var v10 uint64 = 656
    var v11 float32 = 1.5
    v12 := GenS1{F0: 128}
    fmt.Println(v1, v3, v10)
    if (3.14159 < 0.1) {
        var v13 int64 = (v3 & (v3 * 42))
        v10 += (v10 * v10)
        _ = v13
    }
    v14 := GenS2{F0: -1}
    v15 := GenS3{F0: 42, F1: v3, F2: (^(v10))}
    for fi16 := 0; fi16 < 5; fi16++ {
        v11 = v11
        if true {
            v1 = v1
            var v17 float64 = ((0.0 * 3.14159) / -2.25)
            v3 += -517
            _ = v17
        } else {
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 string = ""
                fmt.Println(v19, v3, v1)
                v3 = 42
                _ = fi18
                _ = v19
            }
            if false {
                if len(v2) > 0 {
                    v2[0] = (v1 + (-(fi16)))
                }
                var v20 int16 = -32768
                if len(v2) > 0 {
                    v2[0] = fi16
                }
                _ = v20
            }
        }
        var v21 int = (fi16 | (v1 / ((int(v10)) | 1)))
        _ = fi16
        _ = v21
    }
    if true {
        if true {
            for fi22 := 0; fi22 < 2; fi22++ {
                if false {
                    for fi23 := 0; fi23 < 9; fi23++ {
                        var v24 float32 = float32(3.14159)
                        var v25 string = ""
                        switch (v25 + "hello") {
                        case "":
                            fmt.Println(v10, v25, v25)
                        case "\n\t":
                            v2 = append(v2, fi22)
                        case "foo bar":
                            fmt.Println(v3)
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 16; fi26++ {
                        var v27 float32 = -2.25
                        var v28 int16 = int16(fi22)
                        v29 := GenS4{F0: 32767, F1: (v28 + 502)}
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 16; fi30++ {
                        v2 = append(v2, (fi30 % ((v1) | 1)))
                        _ = fi30
                    }
                } else {
                    v1 = fi22
                }
                _ = fi22
            }
            if false {
                v1 = v1
                v10 += v10
                switch "a" {
                case "foo bar":
                    if (int64(v1) != v3) {
                        var v31 bool = false
                        _ = v31
                    } else {
                        for fi32 := 0; fi32 < 12; fi32++ {
                            fmt.Println(fi32, v10)
                            v3 += v3
                            fmt.Println(v1, fi32, v11)
                            _ = fi32
                        }
                    }
                case "\n\t":
                    v33 := GenS1{F0: 0}
                    _ = v33
                default:
                    switch -2147483648 {
                    case 469:
                        v11 -= v11
                    default:
                        var v34 int32 = 2147483647
                        _ = v34
                    }
                }
            } else {
                v35 := GenS3{F0: 128, F1: (v3 & v3), F2: 4294967295}
                _ = v35
            }
            if true {
                var v36 bool = false
                v37 := GenS3{F0: 128, F1: v3, F2: v10}
                _ = v36
                _ = v37
            }
        }
    } else {
        fmt.Println(v1, v3, v3)
        if len(v2) > 0 {
            v2[0] = v1
        }
        fmt.Println(v1, v11)
    }
    var v38 uint8 = 128
    var v39 uint8 = (v38 * v38)
    v1 -= (v1 ^ v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v8
}
