package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS6 struct { F0 string; F1 uint16; F2 uint16 }
func genFn0(p0 int8) uint8 {
    v4 := GenS1{F0: -759}
    v5 := GenS1{F0: 32767}
    var v6 int64 = -9223372036854775808
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 128
}
func genFn1(p0 uint32, p1 int8) int64 {
    v7 := GenS2{F0: false, F1: 1}
    if true {
        if (506 > -2147483648) {
            var v8 float64 = float64(1.5)
            _ = v8
        }
    } else {
        var v9 float64 = (0.1 - ((0.1 - float64((float32(1.5) * (0.1 - (-2.25 + 1e10))))) + 3.14159))
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 int = fi10
        v12 := GenS1{F0: 32767}
        _ = fi10
        _ = v11
        _ = v12
    }
    switch 633 {
    case 248:
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 uint16 = 42
            var v15 float64 = ((-2.25 - float64((0.1 / 1.5))) + 1.5)
            _ = fi13
            _ = v14
            _ = v15
        }
    case 805:
        for fi16 := 0; fi16 < 11; fi16++ {
            for fi17 := 0; fi17 < 5; fi17++ {
                v18 := []int{484, fi16, (fi16 | (fi17 / ((-2147483648) | 1))), (fi16 + fi16)}
                p1 -= (-(p1))
                p1 = (p1 % ((-128) | 1))
                _ = fi17
                _ = v18
            }
            switch "\n\t" {
            case "\n\t":
                v19 := GenS2{F0: false, F1: 1}
                _ = v19
            case "":
                p0 -= (p0 >> ((4294967295) & 63))
            default:
                if true {
                    var v20 int = fi16
                    v20 -= 0
                    v20 = (v20 + fi16)
                    _ = v20
                } else {
                    switch "" {
                    case "foo bar":
                        if false {
                            v21 := []int{(fi16 % ((fi16) | 1)), (fi16 - (-(fi16))), (^(fi16)), 854}
                            switch 597 {
                            case 9223372036854775807:
                                if (float64(1e10) < float64(0.0)) {
                                    if (fi16 == (-(fi16))) {
                                        var v22 int64 = -9223372036854775808
                                        _ = v22
                                    }
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                }
                            default:
                                v23 := GenS2{F0: true, F1: 680}
                                _ = v23
                            }
                            fmt.Println(p0)
                            _ = v21
                        } else {
                            v24 := GenS1{F0: 326}
                            var v25 bool = (!(!true))
                            _ = v24
                            _ = v25
                        }
                    }
                    var v26 uint64 = 2
                    for fi27 := 0; fi27 < 4; fi27++ {
                        var v28 float64 = 0.1
                        v29 := []int{(fi16 - -2147483648), (fi27 / ((fi27) | 1)), fi27}
                        if (128 <= 128) {
                            var v30 int = (fi27 | 0)
                            _ = v30
                        } else {
                            fmt.Println(fi27)
                            var v31 float32 = 1e10
                            _ = v31
                        }
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    _ = v26
                }
            }
            var v32 int = (^(fi16))
            _ = fi16
            _ = v32
        }
    case 4294967295:
        fmt.Println(p0, p0)
    }
    if false {
        v33 := []int{-1, 128, 969}
        v34 := GenS1{F0: -32768}
        _ = v33
        _ = v34
    } else {
        for fi35 := 0; fi35 < 8; fi35++ {
            for fi36 := 0; fi36 < 9; fi36++ {
                switch fi36 {
                case 2147483647:
                    v37 := []int{fi36, fi36}
                    _ = v37
                default:
                    p1 = p1
                }
                _ = fi36
            }
            v38 := []int{(fi35 + fi35), fi35}
            if (0 > 4294967295) {
                for fi39 := 0; fi39 < 19; fi39++ {
                    fmt.Println(fi39)
                    var v40 uint16 = 255
                    _ = fi39
                    _ = v40
                }
                p1 -= (^(p1))
                v41 := GenS3{F0: (("\n\t" + "\n\t") + ("a" + "hello")), F1: (^(p0))}
                _ = v41
            } else {
                var v42 float32 = (1e10 + (0.1 + ((0.1 - (float32(float64(1e10)) + 3.14159)) + 3.14159)))
                v43 := GenS3{F0: ("" + (("a" + "") + "a")), F1: (^(p0))}
                if (true || (("" + "hello") < "a")) {
                    v44 := GenS2{F0: (false || (!true)), F1: -9223372036854775808}
                    p0 = p0
                    _ = v44
                }
                _ = v42
                _ = v43
            }
            _ = fi35
            _ = v38
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return -161
}

func main() {
    v1 := []string{"hello", ("hello" + "\n\t")}
    var v2 int8 = -87
    var v3 uint8 = 42
    for fi45 := 0; fi45 < 17; fi45++ {
        var v46 int32 = 422
        v47 := GenS4{F0: -32768}
        _ = fi45
        _ = v46
        _ = v47
    }
    for fi48 := 0; fi48 < 9; fi48++ {
        var v49 uint16 = 1
        if len(v1) > 0 {
            v1[0] = "a"
        }
        for fi50 := 0; fi50 < 7; fi50++ {
            v51 := GenS5{F0: 236, F1: 1, F2: 4294967295}
            v2 = v2
            _ = fi50
            _ = v51
        }
        _ = fi48
        _ = v49
    }
    v3 -= genFn0(v2)
    var v52 float32 = ((3.14159 * 1.5) + (1e10 + 1.5))
    v3 = (v3 & genFn0((v2 / (((v2 >> ((868) & 63))) | 1))))
    v53 := GenS3{F0: "a", F1: 1}
    if true {
        if false {
            var v54 uint64 = 1
            _ = v54
        }
        for fi55 := 0; fi55 < 4; fi55++ {
            for fi56 := 0; fi56 < 17; fi56++ {
                if true {
                    fmt.Println(fi55, fi55, fi55)
                    var v57 float32 = ((v52 / float32(float64(((1e10 - float32(-2.25)) - float32(float64(float32(1.5))))))) * v52)
                    if (false && true) {
                        v57 += 0.1
                    }
                    _ = v57
                }
                switch "foo bar" {
                case "hello":
                    switch int(v2) {
                    case 42:
                        var v58 int16 = -32768
                        _ = v58
                    default:
                        if (!(true && ((2 == 258) || (18446744073709551615 == 817)))) {
                            fmt.Println(fi56)
                        }
                        v59 := GenS6{F0: "", F1: 0, F2: 65535}
                        if false {
                            v2 = (v2 ^ int8(v3))
                        }
                        _ = v59
                    }
                }
                _ = fi56
            }
            fmt.Println(v2, v52, fi55)
            _ = fi55
        }
    }
    for fi60 := 0; fi60 < 5; fi60++ {
        v52 = v52
        _ = fi60
    }
    v2 = (v2 % ((v2) | 1))
    switch -9223372036854775808 {
    case 42:
        v2 = -1
    case 127:
        if (false && (!(float32(1.5) == ((0.1 - -2.25) - ((v52 * v52) - v52))))) {
            v2 += v2
            var v61 int64 = 309
            _ = v61
        }
    }
    if (!(3.14159 > float64((float32(0.1) / v52)))) {
        if (0 == 0) {
            fmt.Println(v3, v3)
        } else {
            v52 -= (1e10 - v52)
        }
        v2 += (v2 + v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
    _ = v53
}
