package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint, p1 int64, p2 uint) uint16 {
    p2 -= (p0 >> (((p0 | 18446744073709551615)) & 63))
    var v6 int64 = (p1 + 1)
    p0 -= (p2 | (p0 - (p2 + 279)))
    for fi7 := 0; fi7 < 18; fi7++ {
        switch (fi7 / ((fi7) | 1)) {
        case 32767:
            v8 := []int{(fi7 + fi7), fi7, fi7, -1}
            _ = v8
        case 128:
            p2 = 255
        default:
            v9 := GenS1{F0: 0, F1: 0, F2: 0}
            _ = v9
        }
        switch "hello" {
        case "\n\t":
            if false {
                fmt.Println(v6, p2)
                v10 := []int{fi7}
                _ = v10
            } else {
                v11 := []int{(fi7 * int(p1))}
                _ = v11
            }
        case "":
            var v12 string = "\n\t"
            _ = v12
        }
        switch (fi7 - 32767) {
        case 2147483647:
            fmt.Println(p0, p0, p2)
        case -1:
            var v13 string = "\n\t"
            _ = v13
        case 0:
            p0 += 556
        }
        _ = fi7
    }
    switch 0 {
    case 2147483647:
        v14 := []int{127, -908, 9223372036854775807, 4294967295}
        _ = v14
    case 128:
        var v15 bool = (0 < 2147483647)
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}
func genFn1(p0 int64, p1 int8, p2 bool) int8 {
    var v16 uint8 = 255
    v16 = v16
    v17 := []int{-9223372036854775808, 746, 1}
    if (!p2) {
        v18 := GenS1{F0: 1, F1: 1, F2: 0}
        _ = v18
    } else {
        if p2 {
            if (((0.0 - (0.0 + ((-2.25 / 1e10) / -2.25))) + (3.14159 + (-2.25 + 1e10))) >= 1.5) {
                v19 := GenS2{F0: 418, F1: 42}
                for fi20 := 0; fi20 < 15; fi20++ {
                    v21 := GenS2{F0: 162, F1: uint64(v16)}
                    _ = fi20
                    _ = v21
                }
                v22 := GenS2{F0: 65535, F1: 2}
                _ = v19
                _ = v22
            } else {
                v16 += 42
            }
            _ = len(v17)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 8
}

func main() {
    var v1 int16 = 42
    var v2 int = -457
    var v3 int16 = 0
    v4 := []uint32{}
    var v5 int64 = 683
    var v23 int = (v2 >> ((18446744073709551615) & 63))
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 int8 = 62
        _ = fi24
        _ = v25
    }
    var v26 int8 = 88
    if false {
        var v27 string = "a"
        var v28 int64 = v5
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v27
        _ = v28
    } else {
        if len(v4) > 0 {
            v4[0] = 872
        }
        var v29 int16 = int16(v23)
        _ = v29
    }
    v30 := GenS1{F0: 42, F1: 687, F2: 0}
    if (!(!(v5 > v5))) {
        var v31 uint = 2
        if ((float32(1e10) + -2.25) <= 3.14159) {
            v31 -= v31
        } else {
            switch (-(v23)) {
            case 318:
                v5 -= (v5 & (v5 ^ (v5 / ((-9223372036854775808) | 1))))
            default:
                for fi32 := 0; fi32 < 16; fi32++ {
                    fmt.Println(v26, fi32)
                    v3 = (-(v3))
                    _ = fi32
                }
            }
            if true {
                var v33 uint16 = 1
                v3 -= 1
                _ = len(v4)
                _ = v33
            } else {
                v2 = v23
            }
        }
        _ = v31
    } else {
        for fi34 := 0; fi34 < 5; fi34++ {
            switch 1 {
            case -210:
                v35 := GenS2{F0: 2, F1: 42}
                _ = v35
            case 2147483647:
                if (fi34 > v2) {
                    var v36 uint = 0
                    _ = v36
                }
            default:
                var v37 uint32 = 254
                _ = v37
            }
            var v38 int16 = v3
            _ = fi34
            _ = v38
        }
        var v39 int16 = 0
        _ = v39
    }
    if len(v4) > 0 {
        v4[0] = 42
    }
    var v40 int16 = 0
    switch "a" {
    case "a":
        switch (v2 / ((-652) | 1)) {
        case 42:
            for fi41 := 0; fi41 < 17; fi41++ {
                var v42 float32 = (float32(1e10) + 0.1)
                _ = fi41
                _ = v42
            }
        case -9223372036854775808:
            if false {
                v43 := GenS2{F0: 4294967295, F1: 1}
                _ = v43
            } else {
                _ = len(v4)
                if len(v4) > 0 {
                    v4[0] = 42
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        default:
            fmt.Println(v1)
        }
    case "\n\t":
        if ("hello" > "a") {
            switch "a" {
            case "":
                v40 -= -400
            case "foo bar":
                v44 := GenS2{F0: 757, F1: 648}
                _ = v44
            case "\n\t":
                switch 32768 {
                case -1:
                    if false {
                        fmt.Println(v1, v1)
                        v45 := GenS2{F0: 542, F1: 42}
                        if (true && true) {
                            fmt.Println(v1)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        _ = v45
                    } else {
                        v23 = 32767
                        for fi46 := 0; fi46 < 15; fi46++ {
                            v5 += v5
                            _ = fi46
                        }
                    }
                default:
                    if false {
                        for fi47 := 0; fi47 < 18; fi47++ {
                            v23 = v23
                            var v48 float32 = 0.0
                            v49 := GenS1{F0: 1, F1: 2, F2: 42}
                            _ = fi47
                            _ = v48
                            _ = v49
                        }
                        var v50 int8 = (-(v26))
                        for fi51 := 0; fi51 < 7; fi51++ {
                            if len(v4) > 0 {
                                v4[0] = 1
                            }
                            _ = fi51
                        }
                        _ = v50
                    } else {
                        v52 := GenS3{F0: "\n\t", F1: genFn0(42, (v5 % ((v5) | 1)), 65535)}
                        v2 -= (^(v23))
                        _ = v52
                    }
                }
            }
            v53 := GenS4{F0: "a"}
            _ = v53
        } else {
            if (1.5 > 1.5) {
                v23 = v23
            }
            v54 := GenS1{F0: 127, F1: 47, F2: -1}
            switch 2147483647 {
            case 42:
                if (!(true && (!true))) {
                    for fi55 := 0; fi55 < 18; fi55++ {
                        for fi56 := 0; fi56 < 19; fi56++ {
                            var v57 uint16 = 255
                            v3 += v40
                            _ = fi56
                            _ = v57
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v58 := GenS3{F0: "hello", F1: 0}
                        _ = fi55
                        _ = v58
                    }
                    v2 = v23
                } else {
                    _ = len(v4)
                    switch "\n\t" {
                    case "foo bar":
                        v59 := GenS1{F0: 128, F1: 542, F2: 2147483647}
                        _ = v59
                    case "a":
                        v40 = v1
                    default:
                        for fi60 := 0; fi60 < 7; fi60++ {
                            v23 -= -2147483648
                            _ = fi60
                        }
                    }
                    for fi61 := 0; fi61 < 10; fi61++ {
                        for fi62 := 0; fi62 < 19; fi62++ {
                            var v63 int64 = 1
                            v64 := GenS5{F0: 667, F1: 1}
                            _ = fi62
                            _ = v63
                            _ = v64
                        }
                        _ = fi61
                    }
                }
            }
            _ = v54
        }
    default:
        var v65 bool = false
        _ = v65
    }
    if false {
        v5 = v5
        var v66 uint16 = 2
        var v67 int64 = (v5 % ((v5) | 1))
        _ = v66
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v5
}
