package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 int8; F1 uint64 }
func genFn0() uint32 {
    v6 := []int{32767, 32768}
    if (!(9223372036854775807 <= 42)) {
        var v7 uint32 = 1
        switch 671 {
        case -9223372036854775808:
            v8 := GenS1{F0: v7, F1: (v7 % ((v7) | 1)), F2: v7}
            _ = v8
        }
        if len(v6) > 0 {
            v6[0] = 32768
        }
        _ = v7
    } else {
        v9 := GenS1{F0: 757, F1: 0, F2: 726}
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS1{F0: 6, F1: 65535, F2: 26}
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi10
            _ = v11
        }
        if (((3.14159 + 0.1) - float32(float64((1e10 - (1.5 / (float32(-2.25) / -2.25)))))) != (float32(3.14159) - 0.0)) {
            if false {
                var v12 uint64 = 386
                _ = v12
            }
            var v13 uint64 = 4294967295
            _ = v13
        } else {
            var v14 int64 = -1
            if true {
                v15 := GenS2{F0: 128, F1: 510, F2: 1}
                _ = v15
            } else {
                var v16 bool = (!(-811 != 1))
                if (v16 && v16) {
                    if len(v6) > 0 {
                        v6[0] = 9223372036854775807
                    }
                    v16 = (v16 && (true && false))
                }
                _ = v16
            }
            var v17 uint64 = 4294967295
            _ = v14
            _ = v17
        }
        _ = v9
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        if ((!false) && false) {
            var v19 float32 = (-2.25 * 1.5)
            var v20 int8 = 115
            fmt.Println(v20, fi18, v20)
            _ = v19
            _ = v20
        }
        _ = fi18
    }
    var v21 int = 2147483647
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := GenS2{F0: 2, F1: 526, F2: 532}
        v21 -= v21
        if true {
            v24 := GenS1{F0: 1, F1: 65535, F2: 536}
            _ = len(v6)
            _ = v24
        }
        _ = fi22
        _ = v23
    }
    _ = v21
    _ = v6
    return 1
}

func main() {
    var v1 string = ("" + "hello")
    var v2 int = 128
    var v3 int = (-(v2))
    var v4 float64 = (1.5 + -2.25)
    v5 := []int{32767}
    for fi25 := 0; fi25 < 3; fi25++ {
        var v26 uint32 = 42
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        v3 = (fi27 ^ v2)
        switch "" {
        case "foo bar":
            switch fi27 {
            case 4294967295:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case 32768:
                for fi28 := 0; fi28 < 10; fi28++ {
                    var v29 uint8 = 42
                    v30 := GenS3{F0: 809}
                    var v31 float32 = (((1e10 + (float32(((v4 - -2.25) - v4)) / float32((v4 * v4)))) * float32((v4 / ((v4 / float64(0.1)) + 0.1)))) * 1e10)
                    switch (v1 + "foo bar") {
                    case "":
                        var v32 uint8 = 237
                        _ = v32
                    case "a":
                        var v33 int = (-(v3))
                        _ = v33
                    case "hello":
                        var v34 uint8 = (v29 | (v29 % ((v29) | 1)))
                        _ = v34
                    default:
                        var v35 int16 = 0
                        _ = v35
                    }
                    _ = fi28
                    _ = v29
                    _ = v30
                    _ = v31
                }
            }
        case "":
            switch 42 {
            case -9223372036854775808:
                fmt.Println(v4)
            default:
                var v36 uint16 = 257
                _ = v36
            }
        case "hello":
            for fi37 := 0; fi37 < 10; fi37++ {
                fmt.Println(fi37)
                v38 := GenS3{F0: 42}
                if (false || (!(95 == 97))) {
                    v4 -= v4
                    if (-2147483648 <= 0) {
                        v5 = append(v5, (v2 & (v2 << ((5) & 63))))
                        var v39 int64 = 42
                        if false {
                            v40 := GenS2{F0: 2, F1: 1, F2: -1}
                            v1 = "\n\t"
                            if false {
                                var v41 int64 = (v39 >> ((2) & 63))
                                _ = v41
                            }
                            _ = v40
                        }
                        _ = v39
                    }
                    v2 = 2147483647
                    v1 = "foo bar"
                    v4 += (float64(-2.25) + float64(0.1))
                    fmt.Println(v3)
                } else {
                    if (!false) {
                        if false {
                            if false {
                                v3 = v3
                            } else {
                                v3 += (-(v3))
                                fmt.Println(v1, fi37, v4)
                                v4 = 1e10
                                fmt.Println(fi37)
                                fmt.Println(fi27)
                            }
                        }
                        v3 = v3
                    } else {
                        v4 = v4
                        if true {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        } else {
                            v42 := GenS4{F0: -1}
                            _ = v42
                        }
                    }
                    v2 += fi37
                }
                _ = fi37
                _ = v38
            }
        default:
            v5 = append(v5, fi27)
        }
        _ = fi27
    }
    v43 := GenS4{F0: -340}
    var v44 uint8 = 1
    var v45 int8 = -58
    for fi46 := 0; fi46 < 4; fi46++ {
        v44 += v44
        for fi47 := 0; fi47 < 4; fi47++ {
            var v48 int16 = -1
            _ = fi47
            _ = v48
        }
        var v49 int8 = (v45 << ((0) & 63))
        _ = fi46
        _ = v49
    }
    if (0.1 <= 0.1) {
        switch v3 {
        case 0:
            var v50 uint32 = genFn0()
            _ = v50
        }
        if false {
            var v51 uint = 993
            _ = len(v5)
            if len(v5) > 0 {
                v5[0] = (v3 & (v2 / ((v2) | 1)))
            }
            _ = v51
        } else {
            v2 = (v3 << ((42) & 63))
            fmt.Println(v2, v3)
        }
        switch v1 {
        case "\n\t":
            v52 := GenS5{F0: 42, F1: v45}
            _ = v52
        case "":
            for fi53 := 0; fi53 < 1; fi53++ {
                fmt.Println(v4, v2)
                var v54 uint32 = genFn0()
                _ = fi53
                _ = v54
            }
        }
    } else {
        v1 = ("foo bar" + v1)
    }
    v55 := GenS2{F0: (v44 | (^(v44))), F1: 42, F2: -2147483648}
    for fi56 := 0; fi56 < 12; fi56++ {
        for fi57 := 0; fi57 < 8; fi57++ {
            if true {
                v58 := GenS6{F0: 0, F1: 4294967295}
                _ = v58
            } else {
                var v59 uint = 18446744073709551615
                _ = v59
            }
            _ = fi57
        }
        switch "a" {
        case "a":
            var v60 float64 = 1.5
            _ = v60
        case "foo bar":
            v44 -= v44
        }
        v44 = 2
        _ = fi56
    }
    v1 = v1
    fmt.Println(v3, v3)
    v5 = append(v5, (v2 * (v3 ^ 127)))
    _ = len(v5)
    v4 = ((v4 * float64(0.0)) + 1.5)
    if len(v5) > 0 {
        v5[0] = (v3 / ((-707) | 1))
    }
    switch -233 {
    case -1:
        switch v2 {
        case -1:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case 920:
            var v61 uint = 4294967295
            _ = v61
        case 42:
            for fi62 := 0; fi62 < 3; fi62++ {
                fmt.Println(fi62, v1)
                if (v1 > "") {
                    var v63 string = "hello"
                    var v64 uint32 = 0
                    _ = v63
                    _ = v64
                }
                _ = fi62
            }
        default:
            v4 -= v4
        }
    case 128:
        var v65 int8 = (^(v45))
        _ = v65
    default:
        var v66 uint32 = 2147483648
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v55
}
