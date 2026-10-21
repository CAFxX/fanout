package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int16; F1 int32; F2 bool }
type GenS4 struct { F0 int; F1 int64; F2 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int16 }
func genFn0(p0 int64, p1 bool) uint8 {
    p0 -= p0
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: 842}
        _ = fi4
        _ = v5
    }
    v6 := []int{1, 9223372036854775807, -520}
    var v7 string = ""
    if len(v6) > 0 {
        v6[0] = 4294967295
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 128
}

func main() {
    var v1 int32 = 218
    v2 := []int{4294967295, 32767, 32767}
    var v3 uint8 = 2
    var v8 uint16 = 0
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            v11 := GenS1{F0: 1000}
            v8 = (v8 - (v8 << ((18446744073709551615) & 63)))
            v12 := GenS1{F0: 42}
            v13 := GenS2{F0: 124, F1: (1.5 != float32(float64(1.5)))}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 17; fi14++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if false {
                v2 = append(v2, (fi9 ^ (fi9 & -406)))
                if (255 == 1) {
                    if true {
                        var v15 uint = 919
                        _ = len(v2)
                        _ = v15
                    }
                } else {
                    v16 := GenS2{F0: 2, F1: (598 >= int64(fi14))}
                    var v17 int64 = -9223372036854775808
                    _ = v16
                    _ = v17
                }
                switch "hello" {
                case "":
                    var v18 bool = ((-2.25 - 1e10) >= -2.25)
                    _ = v18
                case "foo bar":
                    var v19 int16 = -931
                    _ = v19
                case "hello":
                    if true {
                        var v20 uint8 = 52
                        if true {
                            var v21 int64 = 1
                            if false {
                                v3 += v3
                                var v22 uint8 = v3
                                v20 -= genFn0((v21 - v21), (18446744073709551615 >= 0))
                                _ = v22
                            }
                            _ = v21
                        }
                        switch "a" {
                        case "a":
                            var v23 int32 = (v1 / ((v1) | 1))
                            _ = v23
                        case "":
                            var v24 uint16 = (v8 * 1)
                            _ = v24
                        case "foo bar":
                            if (fi14 >= fi9) {
                                v2 = append(v2, fi9)
                                if true {
                                    switch fi9 {
                                    case -1:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    default:
                                        v25 := GenS3{F0: 1, F1: 216, F2: (127 > -128)}
                                        _ = v25
                                    }
                                    switch "a" {
                                    case "hello":
                                        _ = len(v2)
                                    case "a":
                                        v20 = 127
                                    }
                                } else {
                                    v20 += (v3 >> ((2) & 63))
                                    if ((v1 | (v1 * (v1 >> ((0) & 63)))) < 363) {
                                        if (32767 <= -32768) {
                                            if (!(!(((v1 - -91) >= 1) || (((int16(v1) > -32768) || (65535 >= 0)) && false)))) {
                                                if true {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    switch "a" {
                                                    case "hello":
                                                        var v26 int32 = 1
                                                        _ = v26
                                                    default:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                } else {
                                                    v27 := GenS2{F0: 2, F1: (!(("\n\t" + "") < ("\n\t" + "\n\t")))}
                                                    v28 := GenS4{F0: int(v8), F1: 837, F2: 846}
                                                    if len(v2) > 0 {
                                                        v2[0] = fi9
                                                    }
                                                    _ = v27
                                                    _ = v28
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = len(v2)
                                            }
                                            var v29 uint8 = (v3 | (^(v20)))
                                            if len(v2) > 0 {
                                                v2[0] = (fi14 | (fi9 - 449))
                                            }
                                            _ = v29
                                        }
                                    }
                                }
                            }
                        default:
                            v30 := GenS1{F0: 42}
                            v2 = append(v2, (fi9 ^ fi9))
                            v3 += 0
                            _ = v30
                        }
                        _ = v20
                    } else {
                        v3 -= uint8(v8)
                    }
                default:
                    var v31 float32 = 1e10
                    _ = v31
                }
                if len(v2) > 0 {
                    v2[0] = (fi14 * fi9)
                }
                if len(v2) > 0 {
                    v2[0] = 1
                }
            }
            _ = fi14
        }
        _ = fi9
    }
    if false {
        var v32 int32 = v1
        _ = v32
    } else {
        if true {
            switch -1 {
            case 32767:
                fmt.Println(v1, v8)
            }
            var v33 float32 = 0.0
            if ((!(!true)) && ("foo bar" > "\n\t")) {
                switch 397 {
                case -121:
                    fmt.Println(v8, v33)
                }
                var v34 int8 = 8
                _ = v34
            } else {
                v35 := GenS3{F0: -715, F1: 112, F2: false}
                switch "a" {
                case "hello":
                    switch "hello" {
                    case "a":
                        if (463 > 0) {
                            var v36 int16 = 1
                            fmt.Println(v36)
                            _ = v36
                        }
                    }
                case "\n\t":
                    for fi37 := 0; fi37 < 9; fi37++ {
                        var v38 uint16 = v8
                        v39 := GenS5{F0: "a"}
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                case "":
                    if len(v2) > 0 {
                        v2[0] = -2147483648
                    }
                }
                _ = v35
            }
            _ = v33
        }
        switch "\n\t" {
        case "":
            v1 += v1
        }
    }
    v8 += v8
    fmt.Println(v3)
    v3 -= (v3 ^ (^(v3)))
    if len(v2) > 0 {
        v2[0] = -177
    }
    fmt.Println(v8, v1, v1)
    v3 = v3
    v1 = v1
    if len(v2) > 0 {
        v2[0] = 242
    }
    if (1.5 == (-2.25 - (-2.25 - -2.25))) {
        switch int(v3) {
        case -493:
            v40 := GenS6{F0: genFn0(42, true), F1: -375}
            _ = v40
        case -163:
            for fi41 := 0; fi41 < 4; fi41++ {
                v8 = v8
                v2 = append(v2, fi41)
                switch fi41 {
                case -2147483648:
                    v8 -= v8
                default:
                    v8 = (^(v8))
                }
                _ = fi41
            }
        }
    } else {
        v42 := GenS3{F0: 1000, F1: v1, F2: (true || (true && (!false)))}
        var v43 string = ("foo bar" + "foo bar")
        _ = v42
        _ = v43
    }
    fmt.Println(v1, v3, v1)
    switch 32768 {
    case -2147483648:
        if true {
            var v44 int16 = 1
            _ = v44
        }
    }
    v3 += 41
    v45 := GenS7{F0: ("\n\t" + "foo bar")}
    var v46 uint = 1
    if ((!(true || true)) || (!true)) {
        if false {
            for fi47 := 0; fi47 < 3; fi47++ {
                v1 -= -138
                v48 := GenS8{F0: 649}
                v1 -= v1
                _ = fi47
                _ = v48
            }
            var v49 uint = v46
            _ = v49
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v8
}
