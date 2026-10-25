package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64; F1 int32; F2 string }
type GenS7 struct { F0 int8 }
func genFn0(p0 uint) int64 {
    var v7 float64 = 0.0
    var v8 uint8 = 42
    _ = p0
    _ = v7
    _ = v8
    return 690
}
func genFn1(p0 int8) int32 {
    for fi9 := 0; fi9 < 8; fi9++ {
        p0 = (^(p0))
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 uint = 0
            v12 := []int{}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    p0 += p0
    _ = p0
    return -1
}

func main() {
    var v1 float64 = ((1e10 * (0.1 / 1.5)) * 0.1)
    v2 := make(map[string]bool)
    var v3 uint8 = 17
    v4 := make(map[string]bool)
    v5 := make(map[int]uint32)
    var v6 int16 = -1
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 string = ""
        var v15 uint = 2
        var v16 int16 = v6
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    fmt.Println(v3)
    switch 804 {
    case 127:
        for fi17 := 0; fi17 < 20; fi17++ {
            if (!false) {
                switch 42 {
                case 661:
                    for fi18 := 0; fi18 < 20; fi18++ {
                        var v19 int = (fi18 % ((fi18) | 1))
                        v1 = (v1 * v1)
                        v20 := []int{-9223372036854775808, (v19 | (fi17 % ((248) | 1))), fi17, 2147483647}
                        v6 -= int16(v3)
                        if ("\n\t" > ("a" + "hello")) {
                            v19 -= fi17
                        } else {
                            var v21 uint64 = 1
                            if false {
                                var v22 string = ""
                                _ = v22
                            }
                            if len(v20) > 0 {
                                v20[0] = 1
                            }
                            v23 := GenS1{F0: -913, F1: -767, F2: v6}
                            _ = v21
                            _ = v23
                        }
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                default:
                    if (719 > uint64(v3)) {
                        fmt.Println(v6)
                    }
                }
                _ = len(v4)
            }
            fmt.Println(fi17, fi17)
            if true {
                if (int16(v3) < (v6 & v6)) {
                    v24 := GenS2{F0: 516, F1: (v3 / ((v3) | 1)), F2: 225}
                    _ = v24
                }
                for fi25 := 0; fi25 < 5; fi25++ {
                    if (v1 == float64(3.14159)) {
                        var v26 uint32 = 785
                        v27 := GenS2{F0: 4294967295, F1: v3, F2: 156}
                        var v28 int32 = genFn1(1)
                        var v29 uint8 = 127
                        _ = v26
                        _ = v27
                        _ = v28
                        _ = v29
                    } else {
                        if false {
                            v30 := v5[(fi17 * (fi25 - -9223372036854775808))]
                            var v31 int32 = genFn1(42)
                            v31 = v31
                            _ = v30
                            _ = v31
                        } else {
                            var v32 int32 = 42
                            v33 := GenS2{F0: 1, F1: (v3 - v3), F2: (v3 - (v3 + (^(v3))))}
                            var v34 uint64 = 279
                            fmt.Println(v34)
                            var v35 int = int(v3)
                            _ = v32
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        var v36 uint32 = 952
                        v3 -= v3
                        _ = v36
                    }
                    v3 = (v3 + (v3 & v3))
                    var v37 uint16 = 42
                    var v38 int64 = 42
                    _ = fi25
                    _ = v37
                    _ = v38
                }
                for fi39 := 0; fi39 < 7; fi39++ {
                    if false {
                        v40 := GenS1{F0: (v6 * 0), F1: v6, F2: (v6 & (v6 >> ((27) & 63)))}
                        switch "" {
                        case "hello":
                            if (18446744073709551615 < 138) {
                                var v41 int16 = 32767
                                var v42 uint16 = 0
                                v43 := GenS2{F0: 0, F1: v3, F2: (^(v3))}
                                _ = v41
                                _ = v42
                                _ = v43
                            } else {
                                v44 := GenS2{F0: 42, F1: (^(v3)), F2: (v3 % ((v3) | 1))}
                                _ = v44
                            }
                        case "":
                            v45 := GenS1{F0: v6, F1: (v6 ^ -1), F2: v6}
                            _ = v45
                        }
                        var v46 int8 = 0
                        fmt.Println(fi39)
                        _ = v40
                        _ = v46
                    } else {
                        fmt.Println(v6, fi17, v1)
                        if true {
                            v47 := GenS3{F0: ("" + ("" + "a"))}
                            var v48 int = fi39
                            _ = v47
                            _ = v48
                        }
                    }
                    v49 := GenS4{F0: -587, F1: 2}
                    _ = fi39
                    _ = v49
                }
            } else {
                for fi50 := 0; fi50 < 12; fi50++ {
                    v51 := GenS5{F0: 357}
                    var v52 uint32 = 266
                    _ = fi50
                    _ = v51
                    _ = v52
                }
                switch "\n\t" {
                case "hello":
                    for fi53 := 0; fi53 < 11; fi53++ {
                        fmt.Println(v6)
                        _ = fi53
                    }
                case "foo bar":
                    var v54 int = fi17
                    _ = v54
                default:
                    v55 := GenS6{F0: 4294967295, F1: 42, F2: ("a" + ("a" + "a"))}
                    _ = v55
                }
            }
            _ = fi17
        }
    case -243:
        for fi56 := 0; fi56 < 14; fi56++ {
            v1 -= v1
            var v57 uint16 = 970
            _ = fi56
            _ = v57
        }
    default:
        v6 += (v6 - 1000)
    }
    v6 += v6
    v3 -= 42
    for fi58 := 0; fi58 < 15; fi58++ {
        v6 += (v6 ^ v6)
        v1 -= v1
        var v59 bool = false
        _ = fi58
        _ = v59
    }
    v60 := GenS7{F0: 100}
    var v61 string = ("\n\t" + ((("a" + "a") + "foo bar") + "a"))
    fmt.Println(v3, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
}
