package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 uint64; F1 int8; F2 uint }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint64; F1 bool }
func genFn0(p0 uint64) uint64 {
    p0 -= (p0 << ((18446744073709551615) & 63))
    for fi7 := 0; fi7 < 17; fi7++ {
        switch "hello" {
        case "a":
            switch 42 {
            case -1:
                for fi8 := 0; fi8 < 7; fi8++ {
                    fmt.Println(fi7)
                    v9 := GenS1{F0: 2}
                    var v10 int8 = -20
                    var v11 float64 = 0.0
                    v12 := GenS1{F0: 1}
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v12
                    _ = v9
                }
            default:
                switch (("\n\t" + "") + "a") {
                case "foo bar":
                    fmt.Println(fi7, fi7)
                case "a":
                    var v13 int32 = -2147483648
                    _ = v13
                case "\n\t":
                    v14 := []int{32768}
                    _ = v14
                }
            }
        default:
            v15 := GenS2{F0: 1, F1: (fi7 << ((42) & 63))}
            _ = v15
        }
        _ = fi7
    }
    _ = p0
    return p0
}
func genFn1(p0 uint8) int16 {
    p0 -= (p0 >> ((18446744073709551615) & 63))
    v16 := GenS2{F0: 1, F1: 128}
    _ = p0
    _ = v16
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 32767
    var v3 float64 = 1e10
    v4 := make(map[string]uint32)
    var v5 uint = 4294967295
    var v6 uint8 = 218
    v17 := GenS2{F0: v2, F1: 9223372036854775807}
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 int64 = -1
        var v20 uint32 = 2
        var v21 float64 = -2.25
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    if false {
        v1[65535] = int(v6)
        for fi22 := 0; fi22 < 10; fi22++ {
            for fi23 := 0; fi23 < 20; fi23++ {
                fmt.Println(v6, v2)
                v3 = (v3 - float64(0.0))
                v5 += 255
                _ = fi23
            }
            if false {
                v2 -= genFn1((^(v6)))
            }
            _ = fi22
        }
    } else {
        v2 += v2
        switch 42 {
        case 4294967295:
            v6 = 127
        case -9223372036854775808:
            v24 := GenS3{F0: (!true)}
            _ = v24
        }
    }
    if false {
        if (4294967295 >= 4294967295) {
            var v25 uint8 = (^(v6))
            for fi26 := 0; fi26 < 2; fi26++ {
                if (18446744073709551615 == genFn0(0)) {
                    v6 = (v25 - (^(v25)))
                    v25 -= 127
                } else {
                    v27 := GenS4{F0: v2, F1: ("hello" >= "")}
                    var v28 int8 = 0
                    _ = v27
                    _ = v28
                }
                for fi29 := 0; fi29 < 15; fi29++ {
                    var v30 bool = (!(-128 == 1))
                    _ = fi29
                    _ = v30
                }
                if (true || true) {
                    var v31 float32 = -2.25
                    _ = v31
                }
                _ = fi26
            }
            _ = v25
        }
        switch "foo bar" {
        case "hello":
            if false {
                var v32 float32 = (float32(v3) + 3.14159)
                v3 = 1e10
                _ = v32
            }
        }
        for fi33 := 0; fi33 < 4; fi33++ {
            if (1 > 255) {
                var v34 float64 = 1.5
                for fi35 := 0; fi35 < 10; fi35++ {
                    v2 = v2
                    _ = fi35
                }
                for fi36 := 0; fi36 < 6; fi36++ {
                    if ((fi33 * (fi36 ^ 32767)) == 0) {
                        v37 := GenS3{F0: (42 <= 42)}
                        if false {
                            v2 += v2
                            fmt.Println(fi33, v3, v3)
                            v6 = 2
                        } else {
                            if false {
                                switch fi36 {
                                case 42:
                                    if (false || false) {
                                        if true {
                                            v38 := GenS5{F0: 406, F1: 4294967295}
                                            var v39 uint16 = 1
                                            switch (fi36 % (((fi36 - (fi33 & (^(fi36))))) | 1)) {
                                            case 4294967295:
                                                v2 = int16(v5)
                                            case 32767:
                                                if (true || (!false)) {
                                                    if true {
                                                        var v40 uint32 = 42
                                                        v1[938] = fi33
                                                        _ = len(v4)
                                                        _ = v40
                                                    }
                                                    v34 -= (((1.5 * v34) - v34) / v3)
                                                    v41 := []int{fi33, (fi33 - (fi36 - fi33)), (fi36 + fi33), fi33}
                                                    _ = v41
                                                }
                                            case -1:
                                                v42 := GenS6{F0: 0, F1: 1, F2: (v5 % ((v5) | 1))}
                                                _ = v42
                                            default:
                                                v6 = 2
                                            }
                                            _ = v38
                                            _ = v39
                                        } else {
                                            v3 += (((v34 / ((-2.25 - ((0.0 * v34) / v34)) * v34)) / (v34 * -2.25)) - 1.5)
                                            var v43 uint = 42
                                            v44 := GenS1{F0: 42}
                                            _ = v43
                                            _ = v44
                                        }
                                    }
                                    v6 -= v6
                                    v5 += (^(v5))
                                }
                            }
                            switch "hello" {
                            case "\n\t":
                                v45 := []int{(fi36 >> ((v5) & 63)), fi33, (fi36 >> ((4294967295) & 63)), fi36}
                                _ = v45
                            case "foo bar":
                                if true {
                                    fmt.Println(v34)
                                    v46 := []int{}
                                    v6 = v6
                                    _ = v46
                                }
                            }
                            fmt.Println(v6, fi33, fi33)
                        }
                        _ = v37
                    } else {
                        if true {
                            fmt.Println(v34)
                        }
                        var v47 string = "hello"
                        _ = v47
                    }
                    _ = fi36
                }
                _ = v34
            }
            _ = fi33
        }
    }
    switch 4294967295 {
    case 237:
        fmt.Println(v5, v6, v6)
    case -169:
        v3 = 0.1
    default:
        v1[2147483647] = 32767
    }
    if ((true && (((-1 >= v2) && (true && false)) || false)) && true) {
        v2 = (v2 - v2)
    } else {
        var v48 int32 = 2147483647
        _ = v48
    }
    var v49 uint = (v5 | v5)
    var v50 int32 = 42
    switch 32767 {
    case 127:
        for fi51 := 0; fi51 < 16; fi51++ {
            for fi52 := 0; fi52 < 14; fi52++ {
                var v53 int16 = v2
                v5 += 117
                v54 := GenS7{F0: v5}
                _ = fi52
                _ = v53
                _ = v54
            }
            fmt.Println(fi51)
            switch fi51 {
            case 32767:
                v2 += (v2 / ((v2) | 1))
            case 1:
                v2 -= (v2 ^ (v2 + v2))
            }
            _ = fi51
        }
    case 32768:
        v55 := GenS1{F0: 65535}
        _ = v55
    }
    v4["a"] = 330
    switch 9223372036854775807 {
    case 9223372036854775807:
        fmt.Println(v5, v3)
    case -1:
        for fi56 := 0; fi56 < 15; fi56++ {
            v57 := GenS8{F0: 42, F1: true}
            _ = fi56
            _ = v57
        }
    default:
        for fi58 := 0; fi58 < 16; fi58++ {
            v4[("a" + "foo bar")] = 2
            for fi59 := 0; fi59 < 8; fi59++ {
                var v60 int64 = 509
                v49 += (v5 << (((^(v5))) & 63))
                _ = fi59
                _ = v60
            }
            var v61 int32 = (v50 << ((18446744073709551615) & 63))
            _ = fi58
            _ = v61
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
