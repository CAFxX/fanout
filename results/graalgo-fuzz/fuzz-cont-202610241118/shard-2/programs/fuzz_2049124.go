package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint; F1 string; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 int; F2 uint }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 string; F1 bool; F2 string }
type GenS7 struct { F0 int8; F1 bool; F2 uint32 }
type GenS8 struct { F0 uint16; F1 int }
type GenS9 struct { F0 int32; F1 uint; F2 bool }
func genFn0(p0 int64, p1 int8, p2 int8) uint {
    p2 = 122
    var v6 uint = 105
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[string]uint32)
    var v3 string = "hello"
    var v4 int = 32767
    var v5 int16 = 42
    v7 := GenS1{F0: -128, F1: -9223372036854775808}
    v8 := GenS1{F0: 100, F1: -1}
    if true {
        if true {
            switch ((("\n\t" + "a") + "foo bar") + v3) {
            case "hello":
                v3 = v3
            case "":
                if (true && (false || ((-2.25 < (1e10 / 1.5)) || (false || false)))) {
                    var v9 int64 = -9223372036854775808
                    v3 = v3
                    var v10 uint32 = 657
                    _ = v10
                    _ = v9
                }
            case "\n\t":
                if (!(0 == 0)) {
                    _ = len(v2)
                    switch "foo bar" {
                    case "a":
                        if true {
                            for fi11 := 0; fi11 < 12; fi11++ {
                                delete(v2, ("foo bar" + v3))
                                if false {
                                    var v12 uint16 = 914
                                    for fi13 := 0; fi13 < 15; fi13++ {
                                        v5 = (^(v5))
                                        v4 = fi11
                                        _ = fi13
                                    }
                                    _ = v12
                                } else {
                                    var v14 string = v3
                                    v3 = ((("" + ("foo bar" + (v3 + "hello"))) + v3) + "foo bar")
                                    _ = v14
                                }
                                for fi15 := 0; fi15 < 12; fi15++ {
                                    v16 := []int{32767, (fi15 + 0), fi15}
                                    v2[(v3 + v3)] = 42
                                    delete(v2, "hello")
                                    _ = fi15
                                    _ = v16
                                }
                                _ = fi11
                            }
                            var v17 uint8 = 94
                            var v18 float32 = 1.5
                            _ = v17
                            _ = v18
                        }
                    }
                }
            }
        } else {
            v19 := []int{(v4 ^ v4), v4, (v4 * (v4 | -255)), -2147483648}
            _ = v19
        }
    } else {
        for fi20 := 0; fi20 < 2; fi20++ {
            if false {
                v3 = "foo bar"
                var v21 uint64 = 0
                var v22 string = "a"
                _ = v21
                _ = v22
            } else {
                var v23 uint32 = 460
                for fi24 := 0; fi24 < 20; fi24++ {
                    if false {
                        fmt.Println(fi24, fi24, v1)
                    }
                    v3 = (v3 + v3)
                    var v25 int8 = 58
                    _ = fi24
                    _ = v25
                }
                v26 := GenS1{F0: -60, F1: -392}
                _ = v23
                _ = v26
            }
            _ = fi20
        }
        switch (v4 << ((918) & 63)) {
        case 1:
            var v27 int64 = 0
            _ = v27
        case 9223372036854775807:
            v3 = "\n\t"
        case 127:
            switch (v4 - 127) {
            case 0:
                if true {
                    v28 := GenS1{F0: -128, F1: 514}
                    if false {
                        v29 := GenS1{F0: -128, F1: -130}
                        v30 := GenS1{F0: 1, F1: -9223372036854775808}
                        var v31 float32 = (((-2.25 + 3.14159) - -2.25) / 1.5)
                        _ = v29
                        _ = v30
                        _ = v31
                    } else {
                        v32 := GenS2{F0: genFn0(-9223372036854775808, -62, 127), F1: v3, F2: -1}
                        _ = v32
                    }
                    v33 := GenS3{F0: v3}
                    _ = v28
                    _ = v33
                }
            case 519:
                switch v3 {
                case "hello":
                    v5 -= (v5 - (v1 / ((1000) | 1)))
                default:
                    fmt.Println(v5)
                }
            case 2147483647:
                if (-1 != 186) {
                    v34 := []int{(^(v4)), (v4 ^ (v4 + (v4 | 0)))}
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 13; fi35++ {
                        v36 := GenS4{F0: 1, F1: v4, F2: 18446744073709551615}
                        v1 = 1
                        _ = fi35
                        _ = v36
                    }
                    var v37 uint16 = 551
                    v38 := GenS2{F0: 18446744073709551615, F1: "hello", F2: int64(v5)}
                    _ = v37
                    _ = v38
                }
            default:
                if false {
                    v39 := GenS5{F0: 1, F1: (^(v1))}
                    switch -632 {
                    case 127:
                        v3 = "hello"
                    }
                    _ = v39
                } else {
                    v5 += v1
                    for fi40 := 0; fi40 < 3; fi40++ {
                        v41 := GenS5{F0: 42, F1: v1}
                        v4 -= (fi40 >> ((464) & 63))
                        v42 := v2[("hello" + v3)]
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                }
            }
        default:
            var v43 uint64 = 2
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 1; fi44++ {
        var v45 int16 = -1
        v46 := []int{32767, v4, (fi44 | v4)}
        var v47 uint16 = 1
        _ = fi44
        _ = v45
        _ = v46
        _ = v47
    }
    if (1 == 284) {
        if ((true || false) || false) {
            if true {
                v4 = 2147483647
                if false {
                    v5 = 589
                } else {
                    v48 := GenS6{F0: "hello", F1: (!((-(v4)) < (v4 + (v4 | (v4 - (^(v4))))))), F2: ("" + v3)}
                    _ = v48
                }
                v4 -= 128
            } else {
                if (28 == -99) {
                    var v49 uint16 = 657
                    fmt.Println(v5, v4)
                    _ = v49
                }
            }
            for fi50 := 0; fi50 < 5; fi50++ {
                if (100 >= 42) {
                    for fi51 := 0; fi51 < 1; fi51++ {
                        fmt.Println(fi51, fi51, v3)
                        _ = fi51
                    }
                    var v52 uint32 = 42
                    v53 := []int{fi50}
                    _ = v52
                    _ = v53
                }
                _ = fi50
            }
            v5 -= (v1 << ((1) & 63))
        } else {
            var v54 bool = false
            _ = len(v2)
            v2[v3] = 42
            _ = v54
        }
    } else {
        v55 := GenS7{F0: 1, F1: true, F2: 57}
        v56 := GenS4{F0: 0, F1: v4, F2: 1}
        _ = v55
        _ = v56
    }
    var v57 int32 = -2147483648
    for fi58 := 0; fi58 < 17; fi58++ {
        if false {
            if (2 > 42) {
                if false {
                    v3 = "a"
                }
                if true {
                    var v59 string = "foo bar"
                    v4 = (v4 ^ (^(v4)))
                    _ = v59
                } else {
                    if true {
                        var v60 float32 = 3.14159
                        _ = v60
                    }
                    var v61 uint = 1
                    _ = v61
                }
                v4 = (v4 - fi58)
            }
        } else {
            fmt.Println(v5)
            for fi62 := 0; fi62 < 18; fi62++ {
                var v63 uint16 = 65535
                var v64 bool = (false || ((v57 << ((42) & 63)) < (^(v57))))
                if false {
                    if true {
                        v65 := GenS8{F0: v63, F1: (v4 - (v4 | v4))}
                        var v66 int64 = 212
                        v64 = (((fi62 >> ((16) & 63)) != v4) && v64)
                        _ = v65
                        _ = v66
                    }
                }
                if (!((471 <= 1) && (v57 < (v57 | 42)))) {
                    fmt.Println(v64)
                }
                _ = fi62
                _ = v63
                _ = v64
            }
        }
        v57 -= (v57 >> ((1) & 63))
        _ = fi58
    }
    v5 += 962
    for fi67 := 0; fi67 < 16; fi67++ {
        var v68 string = "\n\t"
        _ = fi67
        _ = v68
    }
    v69 := GenS3{F0: (v3 + v3)}
    v70 := GenS2{F0: 4294967295, F1: "a", F2: 1}
    _ = len(v2)
    v4 = v4
    v71 := GenS9{F0: (v57 * (v57 - v57)), F1: 2, F2: (true && (!true))}
    var v72 float64 = (1.5 - float64(float32(3.14159)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v69
    _ = v7
    _ = v70
    _ = v71
    _ = v72
    _ = v8
}
