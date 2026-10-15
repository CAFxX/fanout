package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint16 }
func genFn0(p0 int8) uint {
    if ((1.5 / (1.5 - -2.25)) != 0.0) {
        for fi7 := 0; fi7 < 12; fi7++ {
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 int = (fi8 | (fi8 + (fi7 / (((fi7 + (fi8 / ((-2147483648) | 1)))) | 1))))
                var v10 float64 = 1e10
                _ = fi8
                _ = v10
                _ = v9
            }
            if false {
                var v11 int = fi7
                p0 = (p0 | p0)
                for fi12 := 0; fi12 < 12; fi12++ {
                    v11 = fi7
                    p0 = (p0 | (p0 * p0))
                    var v13 uint8 = 175
                    fmt.Println(v11, p0)
                    _ = fi12
                    _ = v13
                }
                _ = v11
            } else {
                fmt.Println(fi7, p0)
                p0 -= -128
                v14 := []int{32768, fi7, fi7, fi7}
                _ = v14
            }
            v15 := []int{fi7, fi7}
            _ = fi7
            _ = v15
        }
        p0 = (p0 % (((^(p0))) | 1))
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        v17 := []int{42, 45}
        if true {
            var v18 int = fi16
            _ = v18
        } else {
            v17 = append(v17, fi16)
        }
        _ = fi16
        _ = v17
    }
    p0 = 127
    _ = p0
    return 1
}

func main() {
    var v1 bool = false
    v2 := []uint32{179, 65535}
    v3 := make(map[uint32]bool)
    var v4 uint8 = 42
    var v5 uint8 = 22
    var v6 int64 = 247
    if (genFn0(127) > genFn0(-80)) {
        var v19 uint = 619
        if v1 {
            var v20 int32 = 2147483647
            _ = v20
        }
        _ = v19
    }
    fmt.Println(v6, v5)
    v5 = (v4 * (v5 ^ v4))
    var v21 float64 = 1.5
    for fi22 := 0; fi22 < 18; fi22++ {
        for fi23 := 0; fi23 < 13; fi23++ {
            if v1 {
                switch "a" {
                case "foo bar":
                    if v1 {
                        var v24 uint64 = 2
                        fmt.Println(v4, fi23)
                        _ = v24
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    v21 -= float64(3.14159)
                    var v25 uint16 = 491
                    _ = v25
                }
            } else {
                if ((v6 & (v6 * (v6 >> ((genFn0(-45)) & 63)))) == (v6 + v6)) {
                    v1 = (4294967295 != fi22)
                    if v1 {
                        if v1 {
                            v4 = v5
                            if false {
                                var v26 float32 = float32(float64(-2.25))
                                var v27 uint32 = 2
                                var v28 uint16 = 321
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                        }
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case "":
                            v29 := GenS1{F0: 648, F1: -2147483648, F2: 255}
                            _ = v29
                        }
                    }
                } else {
                    v4 = v4
                    fmt.Println(fi22, v5)
                    v30 := v3[2147483648]
                    _ = v30
                }
            }
            fmt.Println(fi22, v4, v5)
            _ = fi23
        }
        v3[745] = v1
        _ = fi22
    }
    v21 = (1e10 - -2.25)
    var v31 uint8 = v4
    if (!(!v1)) {
        v4 = v31
    }
    if (v1 && (v1 || (!(435 > 2)))) {
        v4 = v4
        v2 = append(v2, 2147483648)
    } else {
        switch "a" {
        case "a":
            if v1 {
                v31 = (v4 + (v31 & v5))
                v21 = v21
            }
        case "hello":
            for fi32 := 0; fi32 < 5; fi32++ {
                v33 := v3[2]
                v2 = append(v2, uint32(v5))
                _ = fi32
                _ = v33
            }
        }
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
