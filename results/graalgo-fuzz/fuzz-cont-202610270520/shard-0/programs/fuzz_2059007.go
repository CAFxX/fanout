package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int8; F1 int8; F2 int64 }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 uint8; F1 bool; F2 int }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 bool }
func genFn0(p0 uint8, p1 uint16, p2 uint16) int16 {
    var v7 uint32 = 1
    v8 := []int{42, -114, 127, 32767}
    v9 := GenS1{F0: uint(v7), F1: 42, F2: -100}
    for fi10 := 0; fi10 < 10; fi10++ {
        if ((p2 - 42) != (^(p1))) {
            p2 = 411
            if len(v8) > 0 {
                v8[0] = (fi10 << ((290) & 63))
            }
        } else {
            fmt.Println(p1, v7, p0)
        }
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        p2 += p1
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -437
}
func genFn1(p0 uint64, p1 string) int32 {
    v11 := []int{599, -311, -9223372036854775808}
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(p1, fi12, fi12)
        _ = fi12
    }
    if true {
        var v13 int8 = -62
        switch 4294967295 {
        case 128:
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 uint32 = 867
                _ = fi14
                _ = v15
            }
        case 42:
            if (0 <= v13) {
                p1 = p1
            } else {
                if true {
                    fmt.Println(p0)
                } else {
                    v16 := GenS2{F0: 42, F1: (v13 / ((-48) | 1))}
                    var v17 int8 = v13
                    _ = v16
                    _ = v17
                }
            }
        case -923:
            _ = len(v11)
        default:
            if ((3.14159 - 0.0) < 1e10) {
                if len(v11) > 0 {
                    v11[0] = 2147483647
                }
                var v18 bool = (0.0 > 3.14159)
                _ = v18
            } else {
                v19 := GenS1{F0: 18446744073709551615, F1: p0, F2: (v13 | v13)}
                _ = v19
            }
        }
        _ = v13
    } else {
        if true {
            var v20 float64 = (1e10 * 0.1)
            if len(v11) > 0 {
                v11[0] = 485
            }
            _ = v20
        } else {
            for fi21 := 0; fi21 < 11; fi21++ {
                if len(v11) > 0 {
                    v11[0] = (fi21 % ((fi21) | 1))
                }
                var v22 int = fi21
                p1 = "hello"
                _ = fi21
                _ = v22
            }
            var v23 float64 = 0.0
            _ = v23
        }
        v24 := GenS3{F0: p0, F1: 46}
        var v25 int32 = 1
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        p0 += (p0 ^ 2)
        var v27 string = ("\n\t" + "a")
        _ = fi26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = v11
    return 2147483647
}

func main() {
    v1 := []uint32{1}
    var v2 float64 = 0.1
    var v3 int32 = 1
    v4 := make(map[int]string)
    var v5 float32 = (((0.0 - -2.25) - 1.5) + 1.5)
    var v6 float64 = v2
    v5 += (((v5 / (v5 - (-2.25 * ((v5 * v5) * v5)))) + 3.14159) * v5)
    v28 := GenS1{F0: 42, F1: 2, F2: -1}
    if (65535 < 2) {
        switch -2147483648 {
        case 128:
            switch int(v3) {
            case 978:
                fmt.Println(v6)
            default:
                for fi29 := 0; fi29 < 4; fi29++ {
                    switch fi29 {
                    case 1:
                        v5 = v5
                    case -1:
                        v3 = (v3 * (v3 * v3))
                    default:
                        var v30 int16 = 42
                        _ = v30
                    }
                    if (915 == 845) {
                        for fi31 := 0; fi31 < 19; fi31++ {
                            var v32 float64 = v2
                            v32 -= float64(v5)
                            fmt.Println(v3)
                            v3 = v3
                            v2 += ((-2.25 + (((v32 / v6) * -2.25) / v6)) * float64((v5 / v5)))
                            v33 := v4[(fi29 % ((fi31) | 1))]
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        var v34 int64 = -487
                        switch (fi29 ^ fi29) {
                        case 4294967295:
                            v35 := GenS4{F0: 42, F1: int8(v34), F2: v34}
                            _ = v35
                        case 2147483647:
                            if true {
                                v2 += v6
                                var v36 string = ((("\n\t" + (("hello" + "\n\t") + "\n\t")) + ("a" + "")) + "foo bar")
                                _ = v36
                            } else {
                                var v37 bool = false
                                v38 := GenS5{F0: fi29, F1: -32768}
                                _ = v37
                                _ = v38
                            }
                        default:
                            switch (fi29 & fi29) {
                            case 127:
                                for fi39 := 0; fi39 < 2; fi39++ {
                                    var v40 bool = false
                                    _ = fi39
                                    _ = v40
                                }
                            case 729:
                                v2 -= v6
                            default:
                                if ((-1 > v34) || (126 >= 2)) {
                                    var v41 int64 = v34
                                    for fi42 := 0; fi42 < 13; fi42++ {
                                        v1 = append(v1, 42)
                                        _ = len(v1)
                                        _ = fi42
                                    }
                                    _ = v41
                                } else {
                                    v43 := GenS6{F0: genFn0(0, 0, 648)}
                                    _ = v43
                                }
                            }
                        }
                        _ = v34
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                        for fi44 := 0; fi44 < 16; fi44++ {
                            var v45 float32 = (v5 + v5)
                            var v46 uint = 2
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                    }
                    _ = fi29
                }
            }
        case -727:
            v47 := GenS1{F0: 166, F1: 4294967295, F2: 0}
            _ = v47
        case 2147483647:
            var v48 uint8 = 14
            _ = v48
        }
        v49 := GenS4{F0: -53, F1: -116, F2: 0}
        _ = v49
    }
    if false {
        for fi50 := 0; fi50 < 19; fi50++ {
            if ((7 < 26) || false) {
                v51 := GenS3{F0: 2, F1: -30}
                _ = v51
            }
            if (!(((-2.25 - 1e10) - (-2.25 / v6)) < v6)) {
                var v52 int32 = v3
                _ = v52
            }
            _ = fi50
        }
    }
    v53 := GenS2{F0: 65535, F1: 127}
    v54 := GenS7{F0: (true && (genFn1(2, "") != -2147483648)), F1: 563}
    v5 = float32(((float64((v5 / 1.0)) * (v2 / (v2 * (-2.25 + (0.0 + v2))))) + (v6 - (v2 / v2))))
    v55 := GenS8{F0: 2, F1: (2147483647 <= v3), F2: 9223372036854775807}
    var v56 int8 = -1
    var v57 uint32 = 65535
    v58 := GenS9{F0: ((v56 >= (v56 - v56)) && true)}
    var v59 uint16 = 2
    v60 := GenS9{F0: false}
    v56 = (v56 - 42)
    for fi61 := 0; fi61 < 14; fi61++ {
        var v62 int = (fi61 - (fi61 & (fi61 << ((2) & 63))))
        if (!false) {
            v63 := GenS4{F0: 1, F1: 127, F2: -505}
            var v64 int8 = 127
            _ = v63
            _ = v64
        } else {
            if false {
                var v65 int64 = 1
                _ = v65
            }
            v66 := GenS10{F0: false}
            _ = v66
        }
        _ = fi61
        _ = v62
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
}
