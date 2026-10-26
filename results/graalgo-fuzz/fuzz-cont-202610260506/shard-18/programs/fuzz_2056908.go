package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 int8; F2 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64; F1 int }
func genFn0() uint {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 float32 = 0.0
        v8 := GenS1{F0: -78, F1: -869}
        if false {
            v7 -= v7
            v9 := GenS1{F0: 1, F1: -3}
            _ = v9
        } else {
            var v10 bool = true
            _ = v10
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v11 int16 = -32768
    if (true && true) {
        var v12 uint64 = 588
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 bool = (false && (true || true))
            _ = fi13
            _ = v14
        }
        v15 := []int{916, 728, 0, -2147483648}
        _ = v12
        _ = v15
    }
    v16 := []int{-9223372036854775808, 402, -1, -2147483648}
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    _ = v11
    _ = v16
    return 4294967295
}
func genFn1(p0 int64, p1 uint32) int {
    if false {
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 int16 = -1
            _ = fi17
            _ = v18
        }
        if true {
            v19 := []int{32767}
            _ = len(v19)
            if ((!(42 > 1)) || false) {
                if len(v19) > 0 {
                    v19[0] = 9223372036854775807
                }
                var v20 float64 = 0.0
                _ = v20
            } else {
                var v21 uint = genFn0()
                if true {
                    p0 -= -848
                    v22 := GenS1{F0: 33, F1: p0}
                    for fi23 := 0; fi23 < 17; fi23++ {
                        var v24 uint8 = 41
                        var v25 int8 = 127
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = v22
                }
                for fi26 := 0; fi26 < 18; fi26++ {
                    fmt.Println(v21, fi26, v21)
                    switch (fi26 << ((v21) & 63)) {
                    case 736:
                        p1 = 888
                    case 128:
                        fmt.Println(v21)
                    }
                    p0 = 9223372036854775807
                    _ = fi26
                }
                _ = v21
            }
            _ = v19
        } else {
            var v27 uint32 = 65535
            _ = v27
        }
    }
    fmt.Println(p0, p0)
    var v28 float32 = 1e10
    _ = p0
    _ = p1
    _ = v28
    return -188
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    v3 := []string{"hello", "\n\t"}
    var v4 int8 = 42
    v5 := make(map[uint32]int)
    if (32768 > genFn1(-7, 0)) {
        if (true || true) {
            _ = len(v3)
        }
        if true {
            v4 = 1
        } else {
            switch 42 {
            case 128:
                fmt.Println(v4)
            case 315:
                var v29 uint8 = 127
                _ = v29
            }
        }
        var v30 bool = (!(-9223372036854775808 >= 0))
        _ = v30
    } else {
        for fi31 := 0; fi31 < 9; fi31++ {
            if false {
                if true {
                    for fi32 := 0; fi32 < 11; fi32++ {
                        v4 = v4
                        v4 = v4
                        delete(v5, 2147483648)
                        _ = fi32
                    }
                    var v33 uint = 84
                    _ = v33
                }
            } else {
                var v34 bool = false
                _ = v34
            }
            var v35 float64 = 3.14159
            _ = fi31
            _ = v35
        }
    }
    if true {
        var v36 int32 = 2147483647
        for fi37 := 0; fi37 < 14; fi37++ {
            for fi38 := 0; fi38 < 3; fi38++ {
                var v39 string = "hello"
                _ = fi38
                _ = v39
            }
            fmt.Println(v36, v4, fi37)
            _ = fi37
        }
        v40 := GenS2{F0: 0, F1: 0}
        _ = v36
        _ = v40
    } else {
        var v41 uint16 = 804
        if true {
            for fi42 := 0; fi42 < 4; fi42++ {
                switch "" {
                case "hello":
                    delete(v5, 65535)
                case "\n\t":
                    for fi43 := 0; fi43 < 2; fi43++ {
                        var v44 int16 = 1000
                        if true {
                            v41 = v41
                            v4 += (v4 << ((832) & 63))
                            if true {
                                fmt.Println(fi42, v44)
                                v45 := GenS3{F0: v44, F1: -62}
                                fmt.Println(fi42)
                                var v46 uint16 = (v41 % ((v41) | 1))
                                if len(v3) > 0 {
                                    v3[0] = (("a" + "") + "a")
                                }
                                _ = v45
                                _ = v46
                            }
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi43
                        _ = v44
                    }
                case "foo bar":
                    fmt.Println(fi42)
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                switch (fi42 << ((42) & 63)) {
                case 0:
                    switch genFn1(9223372036854775807, 1) {
                    case 681:
                        if (!false) {
                            fmt.Println(v41, v4)
                        }
                    case -1:
                        for fi47 := 0; fi47 < 14; fi47++ {
                            v48 := GenS4{F0: 848}
                            _ = fi47
                            _ = v48
                        }
                    case 1:
                        var v49 uint = 42
                        _ = v49
                    default:
                        switch (fi42 + (fi42 & (fi42 / ((fi42) | 1)))) {
                        case 4294967295:
                            switch (fi42 % ((fi42) | 1)) {
                            case 42:
                                var v50 uint32 = 65535
                                _ = v50
                            case 128:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                v41 = (v41 - v41)
                            }
                        case 0:
                            fmt.Println(fi42)
                        }
                    }
                case 1:
                    for fi51 := 0; fi51 < 1; fi51++ {
                        v41 += v41
                        _ = fi51
                    }
                default:
                    for fi52 := 0; fi52 < 16; fi52++ {
                        var v53 bool = (!false)
                        _ = fi52
                        _ = v53
                    }
                }
                _ = fi42
            }
        } else {
            v54 := v1[977]
            v1[42] = false
            _ = v54
        }
        _ = v41
    }
    var v55 float32 = ((float32(0.0) / 1e10) + ((((0.0 * 0.1) * -2.25) - -2.25) + (0.0 - ((-2.25 / 1.0) + (0.1 * ((float32(1.5) + 3.14159) - float32(-2.25)))))))
    if (false && true) {
        var v56 uint8 = 134
        _ = v56
    }
    v57 := v2[2147483647]
    for fi58 := 0; fi58 < 3; fi58++ {
        v59 := GenS5{F0: 127, F1: (v4 - (v4 & (v4 * -59))), F2: 1}
        v60 := GenS6{F0: v57}
        _ = fi58
        _ = v59
        _ = v60
    }
    var v61 float32 = v55
    v62 := GenS7{F0: 1, F1: v57}
    v3 = append(v3, "a")
    fmt.Println(v61, v61)
    for fi63 := 0; fi63 < 15; fi63++ {
        v64 := v5[455]
        v4 = (-(v4))
        for fi65 := 0; fi65 < 15; fi65++ {
            v5[65535] = (fi63 - (fi65 | 1))
            _ = fi65
        }
        _ = fi63
        _ = v64
    }
    v61 = (((v61 / ((float32((((1e10 / (1.5 + 0.1)) * 1e10) / -2.25)) + v55) * v61)) * 1.5) - -2.25)
    delete(v5, 2)
    var v66 uint16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v61:", v61)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v57
    _ = v61
    _ = v62
    _ = v66
}
