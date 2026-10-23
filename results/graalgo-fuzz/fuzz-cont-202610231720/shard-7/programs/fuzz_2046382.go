package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 bool; F2 int32 }
type GenS6 struct { F0 int64; F1 uint16; F2 string }
type GenS7 struct { F0 uint32; F1 uint16; F2 bool }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint; F1 int32 }
type GenS10 struct { F0 uint; F1 string; F2 int16 }
type GenS11 struct { F0 string }
func genFn0(p0 string, p1 int16, p2 uint8) uint {
    v7 := GenS1{F0: int32(p1), F1: 32767}
    p2 = (p2 * (p2 | p2))
    v8 := GenS2{F0: (1e10 >= -2.25)}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 0
    var v4 uint = 4294967295
    var v5 string = "hello"
    var v6 int8 = 127
    if (100 < 127) {
        v5 = (v5 + "a")
        v9 := GenS3{F0: 59, F1: (!((v4 / ((v4) | 1)) < (v4 + (v4 ^ (v4 >> ((255) & 63))))))}
        _ = v9
    } else {
        switch -9223372036854775808 {
        case -9223372036854775808:
            if true {
                v10 := GenS3{F0: 42, F1: true}
                if (2147483648 <= 757) {
                    v3 -= v3
                    for fi11 := 0; fi11 < 19; fi11++ {
                        for fi12 := 0; fi12 < 12; fi12++ {
                            v13 := GenS4{F0: v3}
                            _ = len(v2)
                            v14 := GenS5{F0: v5, F1: false, F2: 1}
                            var v15 uint16 = 65535
                            _ = fi12
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                        _ = fi11
                    }
                }
                if false {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        if (1 <= 846) {
                            if ((0 <= 1) || (!(0 != 669))) {
                                var v17 uint64 = 188
                                v18 := GenS3{F0: -412, F1: (fi16 >= fi16)}
                                _ = v17
                                _ = v18
                            }
                            v19 := GenS6{F0: 0, F1: 255, F2: ""}
                            _ = v19
                        } else {
                            var v20 uint64 = 4294967295
                            var v21 int32 = 0
                            _ = v20
                            _ = v21
                        }
                        var v22 int16 = -32768
                        _ = fi16
                        _ = v22
                    }
                    if false {
                        v23 := GenS7{F0: 4294967295, F1: 2, F2: true}
                        var v24 int64 = 1
                        if ((1.5 * (1e10 + -2.25)) != (3.14159 * (float64(0.1) / (-2.25 / (0.0 - 1.5))))) {
                            v25 := []int{42, 42, 128}
                            if (2147483647 > 127) {
                                v26 := GenS4{F0: (v3 - v3)}
                                _ = v26
                            }
                            v2[2147483647] = 1
                            _ = v25
                        } else {
                            for fi27 := 0; fi27 < 9; fi27++ {
                                if ((v24 | v24) < v24) {
                                    if ((float64(-2.25) / 3.14159) != 0.1) {
                                        v3 += 28
                                    }
                                    var v28 int32 = 0
                                    var v29 uint8 = (v3 | 2)
                                    _ = v28
                                    _ = v29
                                }
                                _ = fi27
                            }
                            for fi30 := 0; fi30 < 8; fi30++ {
                                for fi31 := 0; fi31 < 14; fi31++ {
                                    v32 := GenS8{F0: ((528 == 1) || (genFn0("foo bar", 88, v3) > v4))}
                                    _ = fi31
                                    _ = v32
                                }
                                _ = fi30
                            }
                            v5 = ""
                        }
                        _ = v23
                        _ = v24
                    } else {
                        var v33 uint16 = 65535
                        fmt.Println(v33, v6, v3)
                        _ = v33
                    }
                } else {
                    var v34 uint64 = 1
                    v35 := GenS5{F0: v5, F1: ((1000 <= 32767) || (859 >= 2)), F2: -1}
                    _ = v34
                    _ = v35
                }
                _ = v10
            }
        case 9223372036854775807:
            v36 := GenS9{F0: (v4 | 667), F1: 42}
            _ = v36
        case -331:
            var v37 int = 9223372036854775807
            _ = v37
        default:
            if (-32768 == 42) {
                var v38 string = v5
                fmt.Println(v4, v4)
                for fi39 := 0; fi39 < 11; fi39++ {
                    v3 = (v3 << ((65535) & 63))
                    _ = fi39
                }
                _ = v38
            } else {
                v40 := []int{}
                v5 = "foo bar"
                _ = v40
            }
        }
    }
    var v41 bool = true
    v5 = v5
    fmt.Println(v6, v4)
    v5 = ("foo bar" + v5)
    v42 := GenS7{F0: 2, F1: 619, F2: (32768 < -1)}
    switch 42 {
    case -330:
        if (42 > 1000) {
            v3 -= v3
            switch 9223372036854775807 {
            case 1:
                v43 := []int{127}
                _ = v43
            case 128:
                switch -2147483648 {
                case -626:
                    v6 += (v6 - v6)
                case -650:
                    if v41 {
                        for fi44 := 0; fi44 < 12; fi44++ {
                            if (v3 <= v3) {
                                var v45 float64 = 1e10
                                _ = v45
                            }
                            var v46 int = fi44
                            _ = fi44
                            _ = v46
                        }
                    }
                case -1:
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v3, v41, v41)
                    case "\n\t":
                        fmt.Println(v6, v3)
                    case "":
                        v4 = v4
                    }
                }
            default:
                if ((float32(float64(((float32(float64(1.5)) * 1e10) / (1.5 - 0.1)))) + (1.5 * float32(1.5))) > (1.5 - -2.25)) {
                    if (42 >= -477) {
                        if (v41 && (2 < 0)) {
                            for fi47 := 0; fi47 < 11; fi47++ {
                                var v48 float64 = -2.25
                                _ = fi47
                                _ = v48
                            }
                            if v41 {
                                v5 = (v5 + v5)
                                var v49 float64 = 0.1
                                var v50 int32 = 152
                                _ = v49
                                _ = v50
                            } else {
                                var v51 uint8 = (v3 / ((v3) | 1))
                                v52 := GenS10{F0: genFn0("\n\t", 1000, v51), F1: (((v5 + v5) + (v5 + v5)) + "a"), F2: -792}
                                _ = v51
                                _ = v52
                            }
                        }
                        var v53 string = v5
                        _ = v53
                    } else {
                        v3 = (v3 & v3)
                        switch "\n\t" {
                        case "a":
                            delete(v1, uint32(v3))
                        }
                    }
                    switch -569 {
                    case 9223372036854775807:
                        if (v41 || v41) {
                            v5 = ""
                            v5 = v5
                        }
                    case 2147483647:
                        switch 128 {
                        case 0:
                            delete(v2, 42)
                        case -1:
                            if v41 {
                                v54 := GenS11{F0: v5}
                                _ = v54
                            } else {
                                delete(v1, 65535)
                            }
                        }
                    case -502:
                        var v55 uint8 = 81
                        _ = v55
                    }
                }
            }
        } else {
            v2[2147483648] = 2147483648
            var v56 string = v5
            switch "" {
            case "hello":
                v57 := GenS3{F0: 789, F1: false}
                _ = v57
            case "foo bar":
                v58 := GenS6{F0: 42, F1: 2, F2: v5}
                _ = v58
            default:
                for fi59 := 0; fi59 < 8; fi59++ {
                    switch fi59 {
                    case 4294967295:
                        var v60 float64 = 0.1
                        _ = v60
                    case 32767:
                        v61 := GenS2{F0: v41}
                        _ = v61
                    case -862:
                        var v62 float32 = (0.1 - 1e10)
                        _ = v62
                    }
                    for fi63 := 0; fi63 < 5; fi63++ {
                        v64 := GenS6{F0: 9223372036854775807, F1: 42, F2: v56}
                        _ = fi63
                        _ = v64
                    }
                    _ = fi59
                }
            }
            _ = v56
        }
    case -109:
        v41 = v41
    case 128:
        v65 := GenS10{F0: v4, F1: "a", F2: 850}
        _ = v65
    }
    if ((-(v6)) == 1) {
        var v66 float32 = 1e10
        var v67 int = 2147483647
        var v68 int8 = -105
        _ = v66
        _ = v67
        _ = v68
    }
    for fi69 := 0; fi69 < 17; fi69++ {
        v2[65535] = 2147483647
        _ = fi69
    }
    var v70 string = (("a" + v5) + v5)
    v41 = false
    var v71 float64 = (3.14159 - 1e10)
    v72 := []int{-1, -9223372036854775808, 32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("slicelen v72:", len(v72), v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v70
    _ = v71
    _ = v72
}
