package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS5 struct { F0 int32; F1 int; F2 bool }
type GenS6 struct { F0 int; F1 int8; F2 bool }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int16; F1 string }
type GenS10 struct { F0 int8 }
func genFn0(p0 int32) uint8 {
    var v7 uint16 = 0
    if (!false) {
        v8 := GenS1{F0: 42}
        p0 -= -2147483648
        _ = v8
    }
    var v9 float64 = 0.0
    if true {
        switch (("" + "foo bar") + ("a" + ("hello" + ""))) {
        case "foo bar":
            switch "a" {
            case "foo bar":
                v9 += (-2.25 * float64(0.0))
            case "a":
                v10 := []int{32768, 32768, 9223372036854775807, 32767}
                _ = v10
            case "hello":
                if (!true) {
                    v11 := GenS1{F0: 73}
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        v13 := []int{fi12, (^(fi12))}
                        if ((1e10 > v9) && (42 >= 247)) {
                            p0 = (^(p0))
                            var v14 bool = true
                            v7 += 0
                            _ = v14
                        } else {
                            for fi15 := 0; fi15 < 8; fi15++ {
                                switch -2147483648 {
                                case 32768:
                                    if (false && (true || true)) {
                                        var v16 bool = (127 > 42)
                                        _ = len(v13)
                                        _ = v16
                                    }
                                case 6:
                                    v7 += 639
                                case 2147483647:
                                    var v17 int16 = 0
                                    _ = v17
                                }
                                _ = fi15
                            }
                            var v18 int32 = p0
                            _ = v18
                        }
                        _ = fi12
                        _ = v13
                    }
                    v19 := []int{28}
                    _ = v19
                }
            default:
                v20 := GenS1{F0: 42}
                _ = v20
            }
        case "\n\t":
            var v21 int8 = 42
            _ = v21
        default:
            switch 32767 {
            case -134:
                v7 += (v7 >> ((2) & 63))
            default:
                v9 = (v9 - (v9 + (float64(-2.25) + (1.5 + 0.0))))
            }
        }
        v22 := []int{}
        _ = v22
    } else {
        var v23 int8 = -58
        _ = v23
    }
    _ = p0
    _ = v7
    _ = v9
    return 0
}

func main() {
    var v1 float64 = float64(0.0)
    var v2 int8 = 42
    var v3 uint = 539
    var v4 int8 = (v2 + (v2 << ((v3) & 63)))
    var v5 int16 = 996
    v6 := []int{}
    v24 := GenS1{F0: (v4 >> ((309) & 63))}
    switch "foo bar" {
    case "":
        var v25 uint32 = 625
        _ = v25
    default:
        for fi26 := 0; fi26 < 7; fi26++ {
            for fi27 := 0; fi27 < 1; fi27++ {
                if len(v6) > 0 {
                    v6[0] = 4294967295
                }
                _ = fi27
            }
            _ = fi26
        }
    }
    _ = len(v6)
    var v28 uint16 = 0
    for fi29 := 0; fi29 < 12; fi29++ {
        if (-2.25 <= v1) {
            v4 = -128
            v28 = (v28 / ((1) | 1))
        } else {
            if len(v6) > 0 {
                v6[0] = (^(fi29))
            }
            var v30 uint16 = v28
            if len(v6) > 0 {
                v6[0] = (fi29 * 32767)
            }
            _ = v30
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v31 int64 = 42
        _ = fi29
        _ = v31
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if true {
        var v32 uint32 = 42
        _ = v32
    } else {
        var v33 float64 = v1
        v28 = v28
        var v34 uint8 = 0
        _ = v33
        _ = v34
    }
    v28 = v28
    var v35 float64 = v1
    v35 += (v35 - (-2.25 * ((v1 - v35) * float64(3.14159))))
    switch -724 {
    case -497:
        for fi36 := 0; fi36 < 6; fi36++ {
            v37 := GenS2{F0: (v5 % ((v5) | 1))}
            v38 := GenS1{F0: v4}
            for fi39 := 0; fi39 < 13; fi39++ {
                var v40 bool = false
                _ = fi39
                _ = v40
            }
            _ = fi36
            _ = v37
            _ = v38
        }
    case 32768:
        var v41 int64 = -1
        _ = v41
    case 127:
        v42 := GenS2{F0: 456}
        _ = v42
    default:
        for fi43 := 0; fi43 < 17; fi43++ {
            var v44 float64 = v1
            _ = fi43
            _ = v44
        }
    }
    var v45 string = ("foo bar" + "")
    var v46 uint32 = 65535
    if true {
        if true {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            switch (v45 + v45) {
            case "":
                _ = len(v6)
            case "a":
                if len(v6) > 0 {
                    v6[0] = -333
                }
            }
        } else {
            fmt.Println(v2, v4)
        }
        v47 := GenS1{F0: v2}
        _ = v47
    } else {
        fmt.Println(v46, v5)
    }
    for fi48 := 0; fi48 < 19; fi48++ {
        if ((v2 / (((v2 << (((v3 - v3)) & 63))) | 1)) != (v2 >> (((v3 & (^(v3)))) & 63))) {
            if true {
                var v49 bool = (false || false)
                _ = v49
            } else {
                switch v45 {
                case "foo bar":
                    v28 -= v28
                default:
                    switch v45 {
                    case "a":
                        v2 = 0
                    case "\n\t":
                        for fi50 := 0; fi50 < 1; fi50++ {
                            v2 = v2
                            v51 := GenS3{F0: uint8(v3)}
                            if ((42 > -847) && true) {
                                v3 = v3
                                var v52 float32 = float32(-2.25)
                                fmt.Println(v4, v52)
                                v53 := GenS4{F0: 117, F1: 128, F2: (v4 * 1)}
                                _ = v52
                                _ = v53
                            } else {
                                v3 = v3
                                v54 := GenS5{F0: -1, F1: -968, F2: (!false)}
                                v2 = (^(v4))
                                v55 := GenS3{F0: 2}
                                _ = v54
                                _ = v55
                            }
                            v56 := GenS6{F0: (fi50 - fi50), F1: v2, F2: (true && true)}
                            if (v45 < (v45 + v45)) {
                                v57 := GenS7{F0: v28, F1: (!false)}
                                switch 127 {
                                case 9223372036854775807:
                                    v35 -= v35
                                case 4294967295:
                                    if (int16(v46) >= (v5 / ((v5) | 1))) {
                                        fmt.Println(v2)
                                    } else {
                                        _ = len(v6)
                                        var v58 bool = false
                                        _ = v58
                                    }
                                case 0:
                                    v6 = append(v6, (^(fi48)))
                                default:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                v46 += 2147483648
                                _ = v57
                            }
                            var v59 uint64 = 42
                            _ = fi50
                            _ = v51
                            _ = v56
                            _ = v59
                        }
                    case "hello":
                        v60 := GenS8{F0: v46}
                        _ = v60
                    default:
                        v61 := GenS3{F0: 1}
                        _ = v61
                    }
                }
                var v62 bool = true
                for fi63 := 0; fi63 < 19; fi63++ {
                    if ((v46 ^ (v46 + (^(v46)))) > 894) {
                        if false {
                            v4 -= v4
                            v3 -= (v3 % ((v3) | 1))
                            v28 += (v28 % (((^(v28))) | 1))
                            fmt.Println(v1)
                            v2 -= 0
                            v64 := GenS9{F0: (v5 % ((int16(v2)) | 1)), F1: (("hello" + (v45 + "hello")) + "hello")}
                            _ = v64
                        }
                    } else {
                        fmt.Println(v46)
                        switch (fi63 * 635) {
                        case 0:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        case -2147483648:
                            v65 := GenS10{F0: 42}
                            _ = v65
                        }
                        var v66 string = v45
                        _ = v66
                    }
                    _ = fi63
                }
                _ = v62
            }
        } else {
            switch ("a" + "a") {
            case "\n\t":
                var v67 int64 = -9223372036854775808
                _ = v67
            case "foo bar":
                v3 = v3
            case "a":
                fmt.Println(v4)
            default:
                v68 := GenS2{F0: v5}
                _ = v68
            }
        }
        for fi69 := 0; fi69 < 16; fi69++ {
            fmt.Println(v46, v3, v46)
            if len(v6) > 0 {
                v6[0] = fi69
            }
            fmt.Println(v1, v45)
            _ = fi69
        }
        _ = fi48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v35
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
