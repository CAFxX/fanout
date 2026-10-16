package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 string; F1 string; F2 int16 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 int8; F1 bool; F2 int8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int }
func genFn0() int {
    var v4 uint16 = 255
    v4 = v4
    fmt.Println(v4)
    v5 := GenS1{F0: 127}
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            v4 = (v4 - v4)
            v4 += v4
            _ = fi7
        }
        _ = fi6
    }
    _ = v4
    _ = v5
    return 127
}

func main() {
    var v1 int8 = 69
    var v2 float32 = 3.14159
    var v3 float64 = -2.25
    if true {
        for fi8 := 0; fi8 < 12; fi8++ {
            if true {
                for fi9 := 0; fi9 < 12; fi9++ {
                    var v10 int32 = 42
                    _ = fi9
                    _ = v10
                }
            }
            _ = fi8
        }
        switch -50 {
        case 921:
            v11 := GenS2{F0: 587}
            _ = v11
        case -852:
            v1 = (-(v1))
        case 661:
            v12 := GenS3{F0: -1, F1: (v1 ^ v1)}
            _ = v12
        }
    } else {
        if (v2 < (3.14159 - float32((float64(v2) - 0.1)))) {
            if true {
                var v13 bool = (63 <= 65535)
                fmt.Println(v13, v2)
                for fi14 := 0; fi14 < 11; fi14++ {
                    v3 += (v3 - 1e10)
                    v15 := GenS4{F0: (("a" + "a") + ("\n\t" + "a")), F1: ("\n\t" + ("\n\t" + ("foo bar" + ""))), F2: 0}
                    fmt.Println(v13, v13, fi14)
                    _ = fi14
                    _ = v15
                }
                _ = v13
            } else {
                if false {
                    var v16 int32 = 1
                    _ = v16
                }
                var v17 int16 = -871
                _ = v17
            }
            var v18 uint16 = 32767
            _ = v18
        } else {
            var v19 int8 = v1
            _ = v19
        }
        fmt.Println(v3, v2)
        v3 = ((float64(v2) + v3) / v3)
    }
    v3 += 1.5
    v20 := GenS5{F0: "\n\t", F1: 127}
    if true {
        if ((42 < 42) && (!false)) {
            switch "" {
            case "":
                fmt.Println(v3, v1, v1)
            default:
                if true {
                    var v21 int32 = -2147483648
                    for fi22 := 0; fi22 < 17; fi22++ {
                        v23 := []int{-459}
                        for fi24 := 0; fi24 < 13; fi24++ {
                            var v25 uint32 = 441
                            v21 = 2147483647
                            v3 = v3
                            _ = fi24
                            _ = v25
                        }
                        v2 += (v2 / 3.14159)
                        _ = fi22
                        _ = v23
                    }
                    for fi26 := 0; fi26 < 4; fi26++ {
                        v1 = 1
                        v27 := []int{(^(fi26)), fi26, (^(fi26))}
                        v3 = float64((v2 - (-2.25 + float32(v3))))
                        _ = fi26
                        _ = v27
                    }
                    _ = v21
                } else {
                    var v28 int64 = 42
                    _ = v28
                }
            }
            v29 := []int{genFn0()}
            var v30 float64 = 3.14159
            _ = v29
            _ = v30
        }
        v31 := GenS6{F0: (-632 > -457), F1: 65535}
        _ = v31
    }
    v1 -= (v1 << ((1) & 63))
    var v32 float32 = float32((0.1 * v3))
    if false {
        for fi33 := 0; fi33 < 19; fi33++ {
            v1 = v1
            v34 := []int{}
            if (!((-2.25 + v2) >= (3.14159 - 1.5))) {
                switch fi33 {
                case 128:
                    fmt.Println(v2)
                case 138:
                    v2 = (0.1 - 1e10)
                case 0:
                    v34 = append(v34, 746)
                }
            }
            _ = fi33
            _ = v34
        }
        for fi35 := 0; fi35 < 10; fi35++ {
            switch genFn0() {
            case 9223372036854775807:
                var v36 uint = 65535
                _ = v36
            default:
                v37 := []int{int(v1)}
                _ = v37
            }
            _ = fi35
        }
        var v38 uint8 = 2
        _ = v38
    } else {
        for fi39 := 0; fi39 < 12; fi39++ {
            if true {
                if false {
                    v40 := GenS7{F0: v1, F1: (!true), F2: (-(v1))}
                    var v41 uint8 = 131
                    v32 -= 1.5
                    _ = v40
                    _ = v41
                } else {
                    fmt.Println(v2)
                    if (genFn0() >= fi39) {
                        var v42 int8 = v1
                        _ = v42
                    } else {
                        v43 := []int{-1}
                        for fi44 := 0; fi44 < 5; fi44++ {
                            var v45 int32 = 2147483647
                            v45 -= (^(v45))
                            var v46 float64 = v3
                            v43 = append(v43, (fi44 << ((964) & 63)))
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                        var v47 int64 = 0
                        _ = v43
                        _ = v47
                    }
                    v48 := []int{(fi39 * 1), fi39, 167}
                    _ = v48
                }
                v32 = (v2 * float32(v3))
                for fi49 := 0; fi49 < 15; fi49++ {
                    v50 := GenS7{F0: v1, F1: true, F2: v1}
                    _ = fi49
                    _ = v50
                }
            } else {
                if (2 != 127) {
                    switch genFn0() {
                    case 901:
                        v2 = float32((float64(v2) / v3))
                    case 0:
                        if false {
                            v51 := GenS7{F0: 42, F1: false, F2: (v1 >> ((4294967295) & 63))}
                            _ = v51
                        } else {
                            v1 = v1
                            if (1 < 0) {
                                if true {
                                    var v52 int16 = 1000
                                    v1 -= (v1 ^ v1)
                                    var v53 string = "hello"
                                    _ = v52
                                    _ = v53
                                }
                                var v54 bool = true
                                v1 -= (v1 * v1)
                                _ = v54
                            }
                        }
                    default:
                        v55 := GenS6{F0: false, F1: 3}
                        _ = v55
                    }
                    fmt.Println(v32, v32, v1)
                    v1 = v1
                }
                var v56 int64 = 0
                _ = v56
            }
            _ = fi39
        }
        var v57 uint16 = 42
        fmt.Println(v57, v2)
        _ = v57
    }
    var v58 int64 = 9223372036854775807
    v32 = float32(-2.25)
    fmt.Println(v1, v58)
    for fi59 := 0; fi59 < 3; fi59++ {
        v60 := GenS8{F0: true}
        fmt.Println(fi59)
        _ = fi59
        _ = v60
    }
    var v61 uint64 = 1
    switch 1 {
    case 32768:
        v62 := GenS9{F0: genFn0()}
        _ = v62
    }
    switch "" {
    case "foo bar":
        var v63 uint8 = 128
        _ = v63
    case "":
        var v64 float32 = ((v2 * -2.25) + 3.14159)
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v58:", v58)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v58
    _ = v61
}
