package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint64, p1 uint) int8 {
    v6 := []int{}
    for fi7 := 0; fi7 < 18; fi7++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi7
    }
    var v8 int32 = -245
    var v9 int = 127
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 bool = (-2147483648 >= -2147483648)
    var v2 float32 = 3.14159
    v3 := make(map[uint32]bool)
    var v4 uint = 461
    v5 := make(map[int]int)
    var v10 int8 = genFn0(324, (v4 + 2))
    if false {
        var v11 uint32 = 78
        var v12 int = 266
        _ = v11
        _ = v12
    } else {
        var v13 float64 = (3.14159 / float64(v2))
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 uint8 = 251
        v1 = v1
        _ = fi14
        _ = v15
    }
    var v16 float32 = (1.5 + v2)
    v17 := []int{4294967295, -39, -592, 1}
    var v18 int32 = -1
    if (42 >= 0) {
        if false {
            for fi19 := 0; fi19 < 17; fi19++ {
                v16 = 1e10
                _ = fi19
            }
            fmt.Println(v18, v10, v10)
        } else {
            _ = len(v17)
            v5[0] = 4294967295
            var v20 uint32 = 768
            _ = v20
        }
    } else {
        switch 9223372036854775807 {
        case 4294967295:
            for fi21 := 0; fi21 < 11; fi21++ {
                var v22 uint32 = 0
                var v23 bool = false
                _ = fi21
                _ = v22
                _ = v23
            }
        default:
            for fi24 := 0; fi24 < 11; fi24++ {
                v4 = v4
                v17 = append(v17, (fi24 | fi24))
                for fi25 := 0; fi25 < 11; fi25++ {
                    fmt.Println(fi24, v1)
                    v26 := GenS1{F0: (v10 / ((genFn0(42, (v4 % (((v4 << ((255) & 63))) | 1)))) | 1)), F1: 548, F2: fi24}
                    switch (fi24 * 127) {
                    case 4294967295:
                        v27 := GenS2{F0: int64(fi24)}
                        _ = v27
                    case -1:
                        v10 = (-(v10))
                    case 9223372036854775807:
                        v28 := v3[42]
                        _ = v28
                    default:
                        v10 = (v10 >> ((v4) & 63))
                    }
                    var v29 string = "a"
                    _ = fi25
                    _ = v26
                    _ = v29
                }
                _ = fi24
            }
        }
        var v30 int32 = (^(v18))
        _ = v30
    }
    v31 := GenS3{F0: genFn0(18446744073709551615, (v4 >> ((v4) & 63))), F1: 42}
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 uint = v4
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        var v35 uint32 = 476
        if v1 {
            v1 = (-32768 >= 42)
        }
        switch fi34 {
        case 2147483647:
            var v36 int16 = 42
            _ = v36
        default:
            fmt.Println(v2, v1, fi34)
        }
        _ = fi34
        _ = v35
    }
    for fi37 := 0; fi37 < 9; fi37++ {
        v2 -= ((1e10 - (v2 - (v16 * v2))) / float32((-2.25 / 1.5)))
        _ = fi37
    }
    var v38 int8 = (v10 + genFn0(0, v4))
    for fi39 := 0; fi39 < 2; fi39++ {
        var v40 bool = true
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 10; fi41++ {
        if (872 >= 4294967295) {
            v16 = (v2 / 1e10)
            var v42 uint32 = 2
            v16 -= v16
            _ = v42
        } else {
            v43 := GenS1{F0: (v38 & v38), F1: 0, F2: fi41}
            if (65535 > 379) {
                v44 := GenS2{F0: 513}
                if len(v17) > 0 {
                    v17[0] = int(v10)
                }
                _ = v44
            }
            _ = v43
        }
        if v1 {
            v38 = v10
        }
        if len(v17) > 0 {
            v17[0] = fi41
        }
        _ = fi41
    }
    var v45 bool = v1
    delete(v3, 188)
    var v46 int = 0
    for fi47 := 0; fi47 < 7; fi47++ {
        if v1 {
            fmt.Println(fi47)
            for fi48 := 0; fi48 < 8; fi48++ {
                v49 := GenS3{F0: (v10 >> ((134) & 63)), F1: 42}
                switch (fi48 | int(v18)) {
                case 128:
                    fmt.Println(v1)
                }
                v50 := GenS4{F0: 940, F1: -9223372036854775808}
                _ = fi48
                _ = v49
                _ = v50
            }
        } else {
            var v51 uint64 = 643
            if false {
                var v52 uint = (v4 >> (((v4 * 2)) & 63))
                var v53 int64 = -345
                if ((v53 + (v53 | v53)) == v53) {
                    v54 := GenS3{F0: v10, F1: 697}
                    switch fi47 {
                    case -397:
                        v2 = v2
                    case -9223372036854775808:
                        for fi55 := 0; fi55 < 16; fi55++ {
                            if v45 {
                                v1 = (!(uint8(v4) <= 127))
                                fmt.Println(v53, v38)
                                var v56 uint32 = 4294967295
                                var v57 uint64 = v51
                                var v58 float64 = 3.14159
                                v17 = append(v17, v46)
                                v38 = (v38 ^ v10)
                                if (!false) {
                                    if false {
                                        v1 = true
                                        v17 = append(v17, (v46 / ((1) | 1)))
                                    } else {
                                        var v59 uint = 4294967295
                                        _ = len(v17)
                                        _ = len(v17)
                                        _ = v59
                                    }
                                }
                                _ = v56
                                _ = v57
                                _ = v58
                            }
                            v1 = (!false)
                            _ = fi55
                        }
                    case 42:
                        fmt.Println(v45, fi47)
                    default:
                        v16 += v16
                    }
                    if true {
                        if true {
                            fmt.Println(v16)
                        } else {
                            var v60 int16 = -1
                            _ = v60
                        }
                        v2 -= v16
                    }
                    _ = v54
                } else {
                    v61 := GenS4{F0: (v51 | v51), F1: (v53 * (v53 - 42))}
                    v2 += (3.14159 - (v2 * ((v2 * v16) + 1.5)))
                    _ = v61
                }
                _ = v52
                _ = v53
            } else {
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                if true {
                    var v62 int8 = (^(v10))
                    var v63 float32 = float32(1e10)
                    var v64 string = ""
                    _ = v62
                    _ = v63
                    _ = v64
                }
            }
            var v65 uint16 = 255
            _ = v51
            _ = v65
        }
        v66 := GenS5{F0: -354}
        v5[fi47] = v46
        _ = fi47
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v38
    _ = v4
    _ = v45
    _ = v46
    _ = v5
}
