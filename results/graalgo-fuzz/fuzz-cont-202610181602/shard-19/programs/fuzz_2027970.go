package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint64 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 string; F1 bool; F2 int64 }
type GenS5 struct { F0 int8; F1 uint16 }
func genFn0(p0 bool) int {
    if p0 {
        if p0 {
            for fi5 := 0; fi5 < 16; fi5++ {
                for fi6 := 0; fi6 < 5; fi6++ {
                    var v7 int64 = 0
                    if (18446744073709551615 <= 18446744073709551615) {
                        v7 -= (v7 << ((1) & 63))
                        p0 = p0
                        switch "foo bar" {
                        case "a":
                            var v8 int16 = 42
                            _ = v8
                        case "\n\t":
                            var v9 uint = 885
                            _ = v9
                        case "hello":
                            if p0 {
                                v10 := []int{}
                                p0 = p0
                                p0 = ((!p0) && (!false))
                                _ = v10
                            }
                        }
                    }
                    var v11 int8 = -128
                    _ = fi6
                    _ = v11
                    _ = v7
                }
                var v12 int16 = 42
                _ = fi5
                _ = v12
            }
            p0 = p0
        } else {
            v13 := GenS1{F0: -1, F1: 1, F2: 42}
            _ = v13
        }
        v14 := []int{}
        _ = v14
    }
    if ((1e10 / ((-2.25 - (3.14159 * 3.14159)) * 1.5)) > 3.14159) {
        v15 := GenS1{F0: 1, F1: -6, F2: 422}
        v16 := []int{0, -9223372036854775808, 607, 2147483647}
        _ = v15
        _ = v16
    }
    var v17 int8 = 1
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 uint = 0
        for fi20 := 0; fi20 < 5; fi20++ {
            v19 = (v19 % ((255) | 1))
            var v21 int16 = 840
            p0 = p0
            fmt.Println(fi18)
            var v22 int = -218
            v23 := GenS1{F0: (^(v21)), F1: -1, F2: 2}
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        v24 := []int{}
        _ = fi18
        _ = v19
        _ = v24
    }
    _ = p0
    _ = v17
    return 9223372036854775807
}
func genFn1(p0 int) uint64 {
    if false {
        for fi25 := 0; fi25 < 18; fi25++ {
            v26 := []int{fi25}
            _ = fi25
            _ = v26
        }
        for fi27 := 0; fi27 < 2; fi27++ {
            if true {
                v28 := []int{32767}
                _ = v28
            } else {
                var v29 float64 = 1e10
                _ = v29
            }
            _ = fi27
        }
        p0 += p0
    } else {
        for fi30 := 0; fi30 < 5; fi30++ {
            var v31 int16 = 0
            _ = fi30
            _ = v31
        }
        p0 += -1
        if ((p0 << ((18446744073709551615) & 63)) > 2147483647) {
            p0 = (p0 & genFn0(true))
            for fi32 := 0; fi32 < 12; fi32++ {
                v33 := GenS2{F0: 42, F1: "\n\t"}
                v34 := GenS2{F0: 0, F1: "hello"}
                if (genFn0(false) >= p0) {
                    var v35 float64 = 1.5
                    _ = v35
                }
                _ = fi32
                _ = v33
                _ = v34
            }
        } else {
            var v36 uint8 = 2
            _ = v36
        }
    }
    p0 -= (p0 % ((p0) | 1))
    _ = p0
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 18
    var v2 float64 = 0.0
    var v3 int64 = 550
    v4 := []string{("a" + ("" + ("hello" + "foo bar"))), ("foo bar" + ""), ""}
    v37 := GenS3{F0: (v1 & (v1 / ((v1) | 1))), F1: "a"}
    var v38 string = "a"
    if len(v4) > 0 {
        v4[0] = ((v38 + "\n\t") + "a")
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v38 = "hello"
    for fi39 := 0; fi39 < 1; fi39++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi39
    }
    if false {
        if (0.1 == 1.5) {
            fmt.Println(v2)
            fmt.Println(v38)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    } else {
        switch v38 {
        case "a":
            fmt.Println(v3)
        default:
            v38 = ("foo bar" + v38)
        }
        v40 := GenS3{F0: (v1 ^ v1), F1: "hello"}
        fmt.Println(v3, v38, v3)
        _ = v40
    }
    v3 += v3
    switch -1 {
    case 983:
        for fi41 := 0; fi41 < 11; fi41++ {
            for fi42 := 0; fi42 < 19; fi42++ {
                v3 = v3
                var v43 uint32 = 65535
                _ = len(v4)
                _ = fi42
                _ = v43
            }
            _ = fi41
        }
    case -162:
        _ = len(v4)
    case 127:
        for fi44 := 0; fi44 < 14; fi44++ {
            if len(v4) > 0 {
                v4[0] = "a"
            }
            v45 := GenS3{F0: 255, F1: v38}
            v38 = "foo bar"
            _ = fi44
            _ = v45
        }
    default:
        for fi46 := 0; fi46 < 3; fi46++ {
            v3 += 42
            _ = fi46
        }
    }
    if true {
        for fi47 := 0; fi47 < 16; fi47++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch "" {
            case "a":
                if false {
                    fmt.Println(v3, v1)
                } else {
                    var v48 bool = (genFn0(false) >= (fi47 + genFn0((!(!true)))))
                    v38 = "a"
                    _ = v48
                }
            case "":
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(fi47)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            default:
                var v49 string = ("" + (v38 + v38))
                _ = v49
            }
            _ = fi47
        }
        v3 -= (v3 / ((v3) | 1))
        for fi50 := 0; fi50 < 2; fi50++ {
            v2 = (0.0 * v2)
            v51 := GenS3{F0: v1, F1: "foo bar"}
            _ = fi50
            _ = v51
        }
    }
    fmt.Println(v3, v1)
    fmt.Println(v3)
    fmt.Println(v1, v38)
    v3 += -1
    switch 52 {
    case -368:
        var v52 uint16 = 65535
        _ = v52
    case 653:
        if (!false) {
            v4 = append(v4, v38)
        }
    case -2147483648:
        if (float64(0.1) < 1e10) {
            var v53 int32 = int32(v3)
            var v54 int64 = v3
            _ = v53
            _ = v54
        } else {
            if true {
                var v55 int = 815
                if false {
                    _ = len(v4)
                } else {
                    var v56 bool = (!(v2 != (v2 + ((3.14159 / -2.25) - ((float64(3.14159) - 1e10) * v2)))))
                    v57 := GenS4{F0: ((v38 + v38) + v38), F1: (127 != v1), F2: (-(v3))}
                    for fi58 := 0; fi58 < 10; fi58++ {
                        v59 := GenS5{F0: -128, F1: 255}
                        _ = fi58
                        _ = v59
                    }
                    _ = v56
                    _ = v57
                }
                _ = v55
            } else {
                fmt.Println(v2, v38, v2)
                v60 := GenS4{F0: v38, F1: false, F2: (v3 | v3)}
                _ = v60
            }
            switch v38 {
            case "\n\t":
                var v61 uint8 = v1
                _ = v61
            default:
                v2 -= 0.1
            }
        }
    default:
        var v62 int16 = -1
        _ = v62
    }
    for fi63 := 0; fi63 < 17; fi63++ {
        var v64 uint8 = (^(v1))
        var v65 uint8 = 42
        var v66 int = fi63
        _ = fi63
        _ = v64
        _ = v65
        _ = v66
    }
    v38 = v38
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
