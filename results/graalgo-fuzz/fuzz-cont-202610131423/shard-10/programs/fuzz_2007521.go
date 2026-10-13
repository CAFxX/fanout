package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int32; F2 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS8 struct { F0 int; F1 int32 }
func genFn0(p0 uint64, p1 string) int {
    v6 := GenS1{F0: 165}
    if false {
        v7 := GenS2{F0: 18446744073709551615, F1: 0, F2: 4294967295}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return 32767
}

func main() {
    var v1 uint16 = 906
    var v2 int16 = 1000
    var v3 string = ("\n\t" + "foo bar")
    v4 := []uint32{711, 2147483647, 0, 2147483647}
    var v5 uint32 = 168
    v8 := GenS1{F0: (v2 / ((1000) | 1))}
    v3 = (v3 + ((v3 + v3) + v3))
    v9 := GenS3{F0: 128}
    for fi10 := 0; fi10 < 20; fi10++ {
        v5 = 0
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 int32 = int32(v1)
            v2 = (v2 * (v2 - -32768))
            if ((((v12 << ((268) & 63)) == v12) || false) || (false || ((v1 >> ((0) & 63)) != 255))) {
                switch "" {
                case "a":
                    v13 := GenS3{F0: fi11}
                    _ = v13
                case "":
                    if false {
                        switch (fi11 * fi11) {
                        case 1:
                            v4 = append(v4, (v5 | 1))
                        case 596:
                            _ = len(v4)
                        case 4294967295:
                            switch fi10 {
                            case 9223372036854775807:
                                if ((v2 > v2) || true) {
                                    v14 := GenS2{F0: 521, F1: v12, F2: (v5 % ((v5) | 1))}
                                    _ = v14
                                } else {
                                    if true {
                                        _ = len(v4)
                                        v2 -= 712
                                        if (1e10 > ((-2.25 + 0.0) + (3.14159 - ((((float64(float32(1e10)) * 0.0) * 0.0) + (1e10 * -2.25)) - 3.14159)))) {
                                            _ = len(v4)
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            switch fi11 {
                                            case -848:
                                                v5 -= (v5 * v5)
                                            case 32767:
                                                var v15 float32 = -2.25
                                                _ = v15
                                            }
                                        } else {
                                            v16 := GenS2{F0: 2, F1: (-(v12)), F2: uint32(v1)}
                                            v17 := GenS2{F0: 0, F1: v12, F2: (v5 >> ((18446744073709551615) & 63))}
                                            v1 += v1
                                            var v18 string = (v3 + v3)
                                            var v19 uint = 582
                                            fmt.Println(fi10)
                                            _ = v16
                                            _ = v17
                                            _ = v18
                                            _ = v19
                                        }
                                    }
                                }
                            case -9223372036854775808:
                                var v20 uint16 = 65535
                                _ = v20
                            }
                        }
                    } else {
                        v1 = 359
                        _ = len(v4)
                        if true {
                            v1 = v1
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            fmt.Println(fi10)
                            var v21 int32 = (v12 + v12)
                            _ = v21
                        }
                    }
                    var v22 bool = (v3 > v3)
                    _ = v22
                case "hello":
                    var v23 uint32 = v5
                    _ = v23
                }
                v3 = "foo bar"
                fmt.Println(v2, v1)
                v24 := GenS1{F0: -1}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v24
            }
            v1 = 65535
            if true {
                var v25 uint = 255
                _ = v25
            } else {
                if (((433 < v1) || (65535 < 0)) || false) {
                    var v26 int32 = (-(v12))
                    if false {
                        if ((!true) || false) {
                            var v27 bool = ((!(0.1 >= (float32(3.14159) / 3.14159))) && (!((v1 % ((665) | 1)) == (v1 + v1))))
                            _ = v27
                        } else {
                            _ = len(v4)
                            switch (-(fi10)) {
                            case 2147483647:
                                if len(v4) > 0 {
                                    v4[0] = (^(v5))
                                }
                            case -916:
                                fmt.Println(v2, v26)
                                _ = len(v4)
                                v28 := GenS4{F0: (v5 / ((2) | 1)), F1: true}
                                _ = v28
                            default:
                                v1 = v1
                            }
                        }
                        var v29 int = (fi10 ^ (fi10 + (fi11 * fi10)))
                        _ = v29
                    } else {
                        _ = len(v4)
                        var v30 uint16 = v1
                        _ = v30
                    }
                    var v31 int64 = -741
                    if ((!(!false)) || ((v26 | -769) > 0)) {
                        if (1 >= v12) {
                            var v32 int = (fi11 << ((uint(v1)) & 63))
                            v33 := GenS4{F0: v5, F1: false}
                            _ = v32
                            _ = v33
                        }
                    } else {
                        switch ("foo bar" + ("hello" + v3)) {
                        case "a":
                            if false {
                                v3 = (v3 + "hello")
                            } else {
                                fmt.Println(fi10)
                                var v34 int = genFn0(1, (v3 + (v3 + (("hello" + (v3 + ("\n\t" + v3))) + (v3 + (((v3 + "") + v3) + v3))))))
                                if len(v4) > 0 {
                                    v4[0] = 65535
                                }
                                _ = v34
                            }
                        case "":
                            if (1 != 0) {
                                var v35 bool = true
                                _ = v35
                            } else {
                                v36 := GenS5{F0: 210}
                                v37 := GenS6{F0: (("" + (("hello" + (v3 + v3)) + v3)) + v3)}
                                _ = v36
                                _ = v37
                            }
                        default:
                            fmt.Println(v3, v31, v26)
                        }
                        if len(v4) > 0 {
                            v4[0] = (v5 >> ((1) & 63))
                        }
                    }
                    _ = v26
                    _ = v31
                }
                v1 += v1
            }
            _ = fi11
            _ = v12
        }
        if true {
            for fi38 := 0; fi38 < 16; fi38++ {
                _ = len(v4)
                v4 = append(v4, v5)
                var v39 float32 = 1.5
                _ = fi38
                _ = v39
            }
            v40 := GenS6{F0: ((("" + ((v3 + v3) + v3)) + (v3 + v3)) + "foo bar")}
            _ = v40
        } else {
            v5 += (v5 - 0)
        }
        _ = fi10
    }
    if len(v4) > 0 {
        v4[0] = v5
    }
    var v41 int8 = 100
    v3 = "foo bar"
    fmt.Println(v41)
    v42 := GenS5{F0: 0}
    _ = len(v4)
    if (-650 < 32767) {
        for fi43 := 0; fi43 < 17; fi43++ {
            if len(v4) > 0 {
                v4[0] = (v5 ^ (v5 >> ((18446744073709551615) & 63)))
            }
            if ((((^(v5)) > v5) || false) && ((1e10 != 1.5) && false)) {
                if (("hello" + v3) >= ((v3 + (v3 + (v3 + ("foo bar" + "foo bar")))) + "foo bar")) {
                    if (495 != 4294967295) {
                        v3 = "foo bar"
                        fmt.Println(v1, v3)
                        v41 = 127
                    } else {
                        if true {
                            var v44 int8 = 42
                            _ = v44
                        }
                        for fi45 := 0; fi45 < 10; fi45++ {
                            v4 = append(v4, v5)
                            if ((v1 <= v1) && (false || (v41 >= (v41 & (v41 & v41))))) {
                                v41 -= (v41 | v41)
                                var v46 uint64 = 4294967295
                                v2 -= v2
                                _ = v46
                            } else {
                                if (!false) {
                                    var v47 int32 = -955
                                    var v48 float64 = 0.1
                                    var v49 bool = false
                                    _ = v47
                                    _ = v48
                                    _ = v49
                                } else {
                                    v50 := GenS4{F0: v5, F1: (1 < -1)}
                                    fmt.Println(v2, v3, v5)
                                    _ = v50
                                }
                                var v51 int8 = -128
                                var v52 int32 = 1
                                _ = v51
                                _ = v52
                            }
                            _ = fi45
                        }
                        for fi53 := 0; fi53 < 10; fi53++ {
                            var v54 uint = 1
                            v55 := GenS7{F0: 1, F1: v5, F2: 142}
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                    }
                    var v56 bool = true
                    var v57 uint16 = (v1 - (^(v1)))
                    _ = v56
                    _ = v57
                } else {
                    v1 += 2
                    v2 -= 0
                    v4 = append(v4, 42)
                }
                if false {
                    var v58 int8 = v41
                    switch -2147483648 {
                    case 1:
                        v59 := GenS8{F0: fi43, F1: -2147483648}
                        _ = v59
                    case -469:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case 9223372036854775807:
                        fmt.Println(v58)
                    }
                    _ = v58
                } else {
                    if (!true) {
                        v3 = v3
                        for fi60 := 0; fi60 < 3; fi60++ {
                            switch v3 {
                            case "\n\t":
                                v1 -= 42
                            }
                            _ = fi60
                        }
                        v61 := GenS1{F0: 1000}
                        _ = v61
                    }
                    switch (("\n\t" + v3) + (v3 + v3)) {
                    case "foo bar":
                        v62 := GenS7{F0: 42, F1: (v5 - v5), F2: 2}
                        _ = v62
                    case "hello":
                        var v63 uint64 = 0
                        _ = v63
                    }
                }
                for fi64 := 0; fi64 < 18; fi64++ {
                    _ = len(v4)
                    if false {
                        if len(v4) > 0 {
                            v4[0] = (v5 & v5)
                        }
                        v65 := GenS7{F0: 42, F1: 2147483647, F2: 154}
                        var v66 uint32 = v5
                        _ = v65
                        _ = v66
                    } else {
                        v5 = (v5 % (((v5 / ((267) | 1))) | 1))
                    }
                    v41 -= v41
                    _ = fi64
                }
            } else {
                v5 = (v5 - (v5 / ((v5) | 1)))
                var v67 float64 = 0.0
                _ = v67
            }
            _ = fi43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v8
    _ = v9
}
