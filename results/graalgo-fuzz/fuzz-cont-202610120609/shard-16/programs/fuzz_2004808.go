package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int8; F2 int16 }
type GenS4 struct { F0 uint8; F1 int16; F2 int8 }
type GenS5 struct { F0 int32; F1 uint64; F2 string }
type GenS6 struct { F0 int8; F1 uint16 }
func genFn0(p0 int16, p1 int32) int32 {
    v6 := []int{4294967295}
    var v7 uint8 = 2
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint8 = (v7 >> ((567) & 63))
        if false {
            if (uint16(v9) <= uint16(v7)) {
                switch (^(fi8)) {
                case -2147483648:
                    p1 = (-(p1))
                case 127:
                    if false {
                        fmt.Println(p0, p0, p0)
                        var v10 uint8 = (v7 * v7)
                        _ = v10
                    } else {
                        var v11 int8 = -128
                        if false {
                            v12 := GenS1{F0: 18446744073709551615, F1: v11, F2: ("" + "foo bar")}
                            _ = v12
                        }
                        _ = v11
                    }
                case 0:
                    p0 = (p0 ^ (p0 / ((p0) | 1)))
                default:
                    v9 -= v7
                }
            } else {
                var v13 uint32 = 42
                var v14 bool = (0 != 42)
                _ = v13
                _ = v14
            }
            v15 := GenS2{F0: (^(fi8))}
            _ = v15
        } else {
            if len(v6) > 0 {
                v6[0] = fi8
            }
        }
        if true {
            var v16 uint16 = 2
            var v17 uint16 = v16
            _ = v16
            _ = v17
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p1 - (p1 * (p1 | p1)))
}
func genFn1(p0 uint, p1 int8) int64 {
    p1 = p1
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 int32 = 436
        v20 := []int{(fi18 >> (((p0 / (((p0 / (((p0 | (p0 | (^(p0))))) | 1))) | 1))) & 63)), (^(fi18)), fi18}
        if true {
            fmt.Println(v19)
            var v21 uint8 = 0
            _ = v21
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := []uint32{4294967295, 192, 2147483648, 49}
    var v2 float32 = float32((-2.25 * 0.1))
    var v3 uint32 = 0
    var v4 string = "a"
    v5 := make(map[string]uint32)
    for fi22 := 0; fi22 < 17; fi22++ {
        for fi23 := 0; fi23 < 8; fi23++ {
            switch (v4 + ("foo bar" + "")) {
            case "hello":
                var v24 uint32 = v3
                _ = v24
            case "foo bar":
                var v25 float32 = (v2 * (float32(float64(3.14159)) * float32(3.14159)))
                _ = v25
            case "\n\t":
                var v26 int16 = 32767
                _ = v26
            default:
                if true {
                    var v27 uint16 = 366
                    var v28 bool = (52 < 100)
                    _ = len(v5)
                    _ = v27
                    _ = v28
                }
            }
            _ = fi23
        }
        var v29 uint32 = v3
        _ = fi22
        _ = v29
    }
    v1 = append(v1, v3)
    var v30 bool = false
    var v31 int8 = -1
    if v30 {
        switch 4294967295 {
        case -306:
            var v32 uint16 = 963
            _ = v32
        case -2147483648:
            switch v4 {
            case "\n\t":
                for fi33 := 0; fi33 < 13; fi33++ {
                    v31 -= (^(v31))
                    _ = fi33
                }
            case "a":
                v30 = (400 == 496)
            case "hello":
                v4 = ("foo bar" + (((v4 + "\n\t") + v4) + "a"))
            default:
                var v34 int32 = 0
                _ = v34
            }
        }
        v30 = (!v30)
        v31 = (v31 * v31)
    } else {
        v2 -= ((float32(1.5) - v2) + 0.0)
        fmt.Println(v31)
    }
    v35 := GenS1{F0: 0, F1: (v31 & 0), F2: ((v4 + (v4 + ("a" + v4))) + v4)}
    _ = len(v1)
    if (505 >= 693) {
        for fi36 := 0; fi36 < 15; fi36++ {
            if v30 {
                if (v30 || (float64(v2) > -2.25)) {
                    for fi37 := 0; fi37 < 11; fi37++ {
                        if false {
                            v38 := GenS1{F0: 18446744073709551615, F1: (v31 * (v31 | 42)), F2: (v4 + v4)}
                            v39 := GenS3{F0: 255, F1: (-(v31)), F2: 42}
                            fmt.Println(v2, v3)
                            _ = v38
                            _ = v39
                        }
                        var v40 uint16 = 32767
                        _ = fi37
                        _ = v40
                    }
                    for fi41 := 0; fi41 < 12; fi41++ {
                        v42 := GenS4{F0: 1, F1: 42, F2: (v31 & v31)}
                        v43 := GenS5{F0: -1, F1: 1, F2: v4}
                        switch v4 {
                        case "foo bar":
                            v3 = 4294967295
                        case "\n\t":
                            v4 = v4
                        case "":
                            var v44 uint16 = 255
                            _ = v44
                        }
                        v3 += v3
                        var v45 uint8 = 42
                        _ = fi41
                        _ = v42
                        _ = v43
                        _ = v45
                    }
                    var v46 uint8 = 1
                    _ = v46
                } else {
                    v47 := GenS1{F0: 0, F1: v31, F2: "hello"}
                    var v48 uint64 = 1
                    v30 = (v30 || false)
                    _ = v47
                    _ = v48
                }
            }
            _ = fi36
        }
    } else {
        v30 = v30
        v49 := GenS6{F0: -128, F1: 0}
        v31 = -1
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
    _ = v5
}
