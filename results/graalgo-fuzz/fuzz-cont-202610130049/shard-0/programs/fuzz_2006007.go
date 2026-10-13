package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8 }
func genFn0() uint32 {
    v4 := GenS1{F0: -1, F1: 65535}
    v5 := []int{4294967295}
    var v6 bool = true
    switch 4294967295 {
    case 0:
        _ = len(v5)
    default:
        _ = len(v5)
    }
    _ = v4
    _ = v5
    _ = v6
    return 65535
}
func genFn1(p0 bool, p1 uint, p2 uint32) uint {
    v7 := []int{-2147483648, 4294967295, -1}
    var v8 float64 = 1.5
    switch 127 {
    case -2147483648:
        for fi9 := 0; fi9 < 11; fi9++ {
            p2 = (p2 + p2)
            v10 := GenS1{F0: -32768, F1: p2}
            _ = fi9
            _ = v10
        }
    case 356:
        if (624 == 42) {
            p2 += (p2 + p2)
            var v11 int32 = 2147483647
            _ = v11
        } else {
            p2 = p2
            fmt.Println(v8, p1)
            if (p0 || p0) {
                for fi12 := 0; fi12 < 10; fi12++ {
                    v13 := GenS1{F0: 411, F1: p2}
                    for fi14 := 0; fi14 < 17; fi14++ {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        switch ("hello" + "foo bar") {
                        case "":
                            v15 := GenS1{F0: -32768, F1: (^(p2))}
                            _ = v15
                        case "\n\t":
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        case "foo bar":
                            if false {
                                v7 = append(v7, (fi14 - (-(fi12))))
                                var v16 int8 = -1
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                p1 = (p1 & p1)
                                _ = v16
                            } else {
                                var v17 int = 127
                                _ = v17
                            }
                        default:
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        }
                        _ = fi14
                    }
                    _ = fi12
                    _ = v13
                }
                if ((2 <= (p1 & p1)) || ((!false) || ((false && p0) && false))) {
                    p1 = p1
                }
                switch "a" {
                case "":
                    for fi18 := 0; fi18 < 8; fi18++ {
                        v19 := GenS1{F0: 1, F1: genFn0()}
                        if p0 {
                            switch (fi18 >> (((p1 / ((p1) | 1))) & 63)) {
                            case -327:
                                var v20 string = "\n\t"
                                _ = v20
                            case -331:
                                if len(v7) > 0 {
                                    v7[0] = fi18
                                }
                            default:
                                p1 += (p1 ^ p1)
                            }
                        } else {
                            var v21 uint16 = 322
                            _ = v21
                        }
                        _ = fi18
                        _ = v19
                    }
                case "\n\t":
                    for fi22 := 0; fi22 < 6; fi22++ {
                        p2 = p2
                        v23 := GenS2{F0: 42, F1: 1000}
                        _ = fi22
                        _ = v23
                    }
                case "a":
                    v24 := GenS1{F0: 42, F1: (p2 << (((p1 - 2)) & 63))}
                    _ = v24
                }
            }
        }
    case 1:
        switch -9223372036854775808 {
        case -9223372036854775808:
            switch "" {
            case "\n\t":
                switch -892 {
                case 1:
                    v7 = append(v7, -1)
                }
            case "":
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            case "a":
                v25 := GenS2{F0: 127, F1: 32767}
                _ = v25
            }
        case 1:
            if p0 {
                var v26 int32 = 42
                _ = v26
            } else {
                fmt.Println(p1, p2, p1)
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
        default:
            v27 := GenS3{F0: 42}
            _ = v27
        }
    default:
        if p0 {
            fmt.Println(v8, p2, p0)
            switch -384 {
            case 32767:
                if (2 > 32767) {
                    v28 := GenS3{F0: 255}
                    fmt.Println(p1, p0)
                    v29 := GenS3{F0: 255}
                    _ = v28
                    _ = v29
                }
            default:
                p0 = (65535 != 32767)
            }
            for fi30 := 0; fi30 < 3; fi30++ {
                if ((((v8 / 3.14159) + (v8 - (float64(float32((float64(((-2.25 / float32(v8)) / -2.25)) * 0.1))) - v8))) == v8) || (!((!p0) && p0))) {
                    switch "" {
                    case "\n\t":
                        p1 = (^(p1))
                    case "foo bar":
                        for fi31 := 0; fi31 < 6; fi31++ {
                            v32 := GenS2{F0: -128, F1: 1000}
                            var v33 int64 = -634
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                    }
                    v8 -= 1.5
                    fmt.Println(v8, p1, fi30)
                } else {
                    v8 = 1.5
                    var v34 uint32 = (^(p2))
                    _ = v34
                }
                _ = fi30
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 0
    var v3 bool = true
    var v35 int8 = 100
    if true {
        switch "" {
        case "a":
            if (false && v3) {
                if (!(false && v3)) {
                    v35 = v35
                    var v36 bool = (-32768 >= 42)
                    _ = v36
                } else {
                    v37 := []int{}
                    _ = v37
                }
            }
        case "":
            v38 := GenS4{F0: (v35 >> ((1) & 63))}
            _ = v38
        }
    }
    v3 = (1e10 > -2.25)
    v2 += 255
    for fi39 := 0; fi39 < 20; fi39++ {
        v40 := GenS4{F0: 42}
        if v3 {
            if false {
                v41 := GenS4{F0: v35}
                _ = v41
            }
        } else {
            v42 := []int{}
            _ = v42
        }
        _ = fi39
        _ = v40
    }
    if false {
        var v43 int16 = -32768
        var v44 int8 = v35
        _ = v43
        _ = v44
    }
    fmt.Println(v2, v35)
    v35 -= v35
    fmt.Println(v35)
    fmt.Println(v2)
    switch 32767 {
    case 4294967295:
        fmt.Println(v3, v2)
    case -2147483648:
        var v45 int32 = 0
        _ = v45
    case -9223372036854775808:
        for fi46 := 0; fi46 < 16; fi46++ {
            if true {
                v2 = v2
                v35 = v35
            }
            _ = fi46
        }
    default:
        for fi47 := 0; fi47 < 5; fi47++ {
            fmt.Println(v35, v2, fi47)
            for fi48 := 0; fi48 < 13; fi48++ {
                var v49 int8 = v35
                _ = fi48
                _ = v49
            }
            _ = fi47
        }
    }
    var v50 uint = 86
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v50
}
