package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint32; F1 int; F2 uint64 }
type GenS5 struct { F0 int8; F1 uint16; F2 int32 }
type GenS6 struct { F0 int; F1 int16; F2 uint32 }
type GenS7 struct { F0 uint; F1 uint }
type GenS8 struct { F0 uint8; F1 uint32 }
func genFn0() int32 {
    var v4 uint16 = 4
    v4 = (v4 - (^(v4)))
    for fi5 := 0; fi5 < 7; fi5++ {
        switch "" {
        case "\n\t":
            var v6 int = (fi5 << ((18446744073709551615) & 63))
            _ = v6
        case "foo bar":
            if (!false) {
                if ((128 >= 2) && false) {
                    v4 = v4
                    for fi7 := 0; fi7 < 10; fi7++ {
                        if (4294967295 != 18446744073709551615) {
                            v4 += (v4 * v4)
                            fmt.Println(fi7, v4, fi7)
                        } else {
                            fmt.Println(fi5, fi7, fi5)
                            v4 += (^(v4))
                            var v8 string = (("\n\t" + "\n\t") + "a")
                            _ = v8
                        }
                        fmt.Println(v4, v4)
                        v9 := []int{(fi5 * fi5), (fi7 & (fi7 << ((0) & 63))), fi7, fi5}
                        _ = fi7
                        _ = v9
                    }
                    var v10 bool = (1 >= 65535)
                    _ = v10
                }
            }
        default:
            var v11 float64 = (1e10 * 3.14159)
            _ = v11
        }
        if (!false) {
            if (2 < 42) {
                switch "foo bar" {
                case "\n\t":
                    v4 -= v4
                case "hello":
                    if (!(127 == 127)) {
                        v4 -= v4
                    } else {
                        var v12 string = (("a" + "foo bar") + (("\n\t" + "\n\t") + "a"))
                        var v13 float64 = (1.5 - 0.1)
                        _ = v12
                        _ = v13
                    }
                default:
                    v4 = v4
                }
            } else {
                v14 := GenS1{F0: 2147483648}
                if (v4 >= v4) {
                    var v15 uint16 = (v4 >> ((1) & 63))
                    var v16 bool = false
                    fmt.Println(v15)
                    _ = v15
                    _ = v16
                } else {
                    v17 := []int{fi5, (fi5 & (fi5 & fi5)), fi5, (fi5 / ((fi5) | 1))}
                    fmt.Println(v4, v4)
                    for fi18 := 0; fi18 < 13; fi18++ {
                        if (2147483647 < uint32(v4)) {
                            v17 = append(v17, 775)
                        }
                        if (!((!true) && true)) {
                            var v19 uint32 = 1
                            if (0.0 <= 3.14159) {
                                fmt.Println(v4)
                                v19 = v19
                            }
                            v20 := GenS2{F0: 255, F1: fi18}
                            _ = v19
                            _ = v20
                        } else {
                            v4 = v4
                        }
                        _ = fi18
                    }
                    _ = v17
                }
                _ = v14
            }
            v4 += (v4 << ((339) & 63))
            v21 := GenS2{F0: 65535, F1: fi5}
            _ = v21
        }
        var v22 int16 = 42
        _ = fi5
        _ = v22
    }
    v4 += v4
    v23 := GenS3{F0: 978, F1: 0}
    _ = v23
    _ = v4
    return 1
}
func genFn1(p0 int8, p1 string) int64 {
    switch -80 {
    case -1:
        p0 = (p0 / ((p0) | 1))
    }
    p0 = (p0 | (p0 / (((^(p0))) | 1)))
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := GenS4{F0: 2147483647, F1: fi24, F2: 18446744073709551615}
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = p1
    return 9223372036854775807
}

func main() {
    var v1 uint = 42
    var v2 uint16 = 688
    v3 := []int{4294967295}
    v1 = (v1 & (v1 << ((v1) & 63)))
    v26 := GenS5{F0: 43, F1: 1, F2: -409}
    v1 += (v1 >> ((v1) & 63))
    var v27 float32 = -2.25
    if ((v2 % ((v2) | 1)) <= v2) {
        fmt.Println(v1, v1, v27)
        if (false || false) {
            if (!true) {
                v2 = v2
                var v28 int64 = -1
                var v29 uint = v1
                _ = v28
                _ = v29
            }
        } else {
            var v30 uint32 = 42
            _ = v30
        }
        for fi31 := 0; fi31 < 18; fi31++ {
            v32 := GenS6{F0: 0, F1: 1, F2: 42}
            _ = fi31
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        for fi34 := 0; fi34 < 19; fi34++ {
            v35 := GenS3{F0: uint32(v1), F1: 2}
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    if (false || (1 != 743)) {
        for fi36 := 0; fi36 < 18; fi36++ {
            var v37 int64 = 9223372036854775807
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi36
            _ = v37
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi38 := 0; fi38 < 14; fi38++ {
            v2 = 65535
            _ = fi38
        }
    }
    fmt.Println(v27, v2, v2)
    var v39 bool = false
    v40 := GenS7{F0: v1, F1: (v1 ^ (v1 - uint(v2)))}
    var v41 uint16 = v2
    if ((v2 ^ (v41 - 1)) >= (v2 & (v41 + (^(v2))))) {
        switch 624 {
        case 42:
            var v42 uint16 = (v41 << ((2) & 63))
            _ = v42
        case -497:
            for fi43 := 0; fi43 < 10; fi43++ {
                var v44 float32 = v27
                _ = fi43
                _ = v44
            }
        default:
            switch -589 {
            case 4294967295:
                var v45 int64 = -9223372036854775808
                _ = v45
            default:
                var v46 bool = (((!v39) && false) && v39)
                _ = v46
            }
        }
        var v47 float64 = 1e10
        for fi48 := 0; fi48 < 18; fi48++ {
            v41 = (^(v2))
            var v49 uint16 = v41
            v50 := GenS8{F0: 63, F1: 4294967295}
            _ = fi48
            _ = v49
            _ = v50
        }
        _ = v47
    }
    v3 = append(v3, 127)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v39
    _ = v40
    _ = v41
}
