package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int64; F1 int16; F2 bool }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 uint64; F1 uint64; F2 uint }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 uint16; F1 int64; F2 uint16 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    var v4 uint = 18446744073709551615
    for fi5 := 0; fi5 < 8; fi5++ {
        v1 = v1
        v4 -= v4
        v6 := GenS1{F0: fi5}
        _ = fi5
        _ = v6
    }
    fmt.Println(v4, v4)
    v7 := GenS1{F0: -2147483648}
    v4 += 2
    switch ("\n\t" + "a") {
    case "\n\t":
        fmt.Println(v1, v4)
    default:
        fmt.Println(v1, v4)
    }
    fmt.Println(v4, v4)
    switch -1 {
    case -495:
        if (1 >= 128) {
            switch 32767 {
            case -9223372036854775808:
                v4 -= v4
            case -870:
                var v8 bool = ((0 == 2) && true)
                _ = v8
            case 32768:
                for fi9 := 0; fi9 < 18; fi9++ {
                    for fi10 := 0; fi10 < 9; fi10++ {
                        v11 := GenS2{F0: 42, F1: 358, F2: 107}
                        v12 := GenS3{F0: "hello", F1: (v4 * (v4 ^ 42))}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    var v13 int8 = 100
                    _ = fi9
                    _ = v13
                }
            default:
                switch 32767 {
                case -899:
                    v14 := []int{}
                    _ = v14
                case -2147483648:
                    v15 := []int{}
                    _ = v15
                default:
                    v4 -= 255
                }
            }
        }
    }
    var v16 uint8 = 1
    for fi17 := 0; fi17 < 16; fi17++ {
        v18 := GenS3{F0: ("foo bar" + "foo bar"), F1: (v4 * (^(v4)))}
        v4 -= (v4 + v4)
        _ = fi17
        _ = v18
    }
    if (!(false || (-128 == 127))) {
        v19 := GenS1{F0: 32768}
        _ = v19
    } else {
        var v20 int8 = 42
        v21 := GenS2{F0: (v1 >> ((v4) & 63)), F1: 65535, F2: 32}
        _ = v20
        _ = v21
    }
    if false {
        var v22 uint = v4
        if true {
            if false {
                var v23 string = "hello"
                _ = v23
            }
        } else {
            v24 := GenS2{F0: -400, F1: 0, F2: 420}
            var v25 float32 = 0.1
            if (false && false) {
                if false {
                    v26 := GenS1{F0: 128}
                    _ = v26
                } else {
                    if false {
                        var v27 uint8 = 32
                        if true {
                            var v28 bool = true
                            var v29 uint = (v4 - v22)
                            switch "hello" {
                            case "":
                                for fi30 := 0; fi30 < 13; fi30++ {
                                    v25 += float32(0.0)
                                    v31 := []int{fi30, fi30, fi30}
                                    _ = fi30
                                    _ = v31
                                }
                            case "\n\t":
                                var v32 int32 = v1
                                _ = v32
                            default:
                                for fi33 := 0; fi33 < 15; fi33++ {
                                    v28 = (v28 && v28)
                                    var v34 uint32 = 2
                                    _ = fi33
                                    _ = v34
                                }
                            }
                            _ = v28
                            _ = v29
                        } else {
                            v25 += v25
                        }
                        _ = v27
                    } else {
                        var v35 uint64 = 4294967295
                        _ = v35
                    }
                    for fi36 := 0; fi36 < 1; fi36++ {
                        v22 = v4
                        _ = fi36
                    }
                }
                v4 = v4
            } else {
                v37 := GenS3{F0: ("" + ("\n\t" + ("a" + ("a" + "foo bar")))), F1: v22}
                _ = v37
            }
            _ = v24
            _ = v25
        }
        var v38 uint8 = v16
        _ = v22
        _ = v38
    } else {
        switch -9223372036854775808 {
        case 4294967295:
            if (106 != 42) {
                v39 := GenS4{F0: 42, F1: 1, F2: (uint64(v16) == 4294967295)}
                v40 := GenS5{F0: 0, F1: 192}
                for fi41 := 0; fi41 < 3; fi41++ {
                    _ = len(v3)
                    v42 := GenS6{F0: 384, F1: 1, F2: uint(v16)}
                    fmt.Println(v4)
                    _ = fi41
                    _ = v42
                }
                _ = v39
                _ = v40
            }
        case 0:
            for fi43 := 0; fi43 < 8; fi43++ {
                fmt.Println(fi43, fi43, v4)
                var v44 int32 = (v1 | (v1 - v1))
                if false {
                    switch fi43 {
                    case -9223372036854775808:
                        v1 = (-(v44))
                    case -787:
                        v1 = -2147483648
                    case 0:
                        for fi45 := 0; fi45 < 10; fi45++ {
                            var v46 float32 = 3.14159
                            if (!(!true)) {
                                var v47 string = ("hello" + "")
                                _ = v47
                            } else {
                                switch (fi43 - fi43) {
                                case 0:
                                    v48 := GenS5{F0: 2, F1: v16}
                                    _ = v48
                                case 128:
                                    var v49 bool = true
                                    _ = v49
                                case 32767:
                                    v16 -= (v16 + 0)
                                }
                                var v50 uint64 = 2
                                var v51 float64 = float64(((3.14159 * v46) + v46))
                                var v52 uint64 = v50
                                if (!(false && (!(true && (!false))))) {
                                    switch (-(fi45)) {
                                    case -1:
                                        if true {
                                            v53 := []int{(-(fi43)), fi43}
                                            fmt.Println(v44)
                                            v44 = int32(v4)
                                            _ = v53
                                        } else {
                                            var v54 uint = 4294967295
                                            fmt.Println(fi43, v4)
                                            v55 := GenS1{F0: 516}
                                            _ = v54
                                            _ = v55
                                        }
                                    case 4294967295:
                                        v44 += v44
                                    }
                                    if false {
                                        var v56 uint16 = 956
                                        _ = v56
                                    } else {
                                        v46 -= v46
                                        v51 = v51
                                        v57 := GenS6{F0: v50, F1: (v52 | 0), F2: (v4 % ((18446744073709551615) | 1))}
                                        v58 := []int{(fi43 + (fi45 & fi43)), (fi43 & (-(fi43))), int(v4)}
                                        if true {
                                            switch "\n\t" {
                                            case "":
                                                fmt.Println(v16)
                                            }
                                            if false {
                                                delete(v3, 818)
                                                if (false || ((v52 >= v52) || false)) {
                                                    switch 127 {
                                                    case 50:
                                                        switch (fi43 * fi43) {
                                                        case 2147483647:
                                                            _ = len(v58)
                                                        }
                                                    case 0:
                                                        var v59 int32 = -1
                                                        _ = v59
                                                    case -9223372036854775808:
                                                        var v60 bool = false
                                                        _ = v60
                                                    default:
                                                        var v61 uint16 = 0
                                                        _ = v61
                                                    }
                                                }
                                                v50 += (v50 % ((v50) | 1))
                                            }
                                            fmt.Println(v52, v16)
                                            v50 = (^(v50))
                                            v62 := GenS6{F0: v52, F1: (v52 % ((v50) | 1)), F2: (v4 / (((v4 % (((v4 / ((18446744073709551615) | 1))) | 1))) | 1))}
                                            _ = v62
                                        } else {
                                            var v63 uint8 = 1
                                            v64 := GenS7{F0: 433, F1: 1000}
                                            _ = v63
                                            _ = v64
                                        }
                                        _ = v57
                                        _ = v58
                                    }
                                    fmt.Println(v44, v1, v50)
                                    switch "" {
                                    case "a":
                                        fmt.Println(fi45, fi43, v51)
                                    case "foo bar":
                                        v51 -= v51
                                    }
                                } else {
                                    v44 += 42
                                    v52 -= (v52 * 154)
                                    v46 = v46
                                    v52 -= 279
                                }
                                _ = v50
                                _ = v51
                                _ = v52
                            }
                            if ((v44 + v44) < -2147483648) {
                                v4 += v4
                                v65 := GenS8{F0: true, F1: 2147483648}
                                var v66 int32 = (^(v1))
                                _ = v65
                                _ = v66
                            } else {
                                v16 = v16
                            }
                            _ = fi45
                            _ = v46
                        }
                    default:
                        fmt.Println(v16)
                    }
                    v67 := []int{fi43}
                    var v68 int16 = -489
                    _ = v67
                    _ = v68
                }
                _ = fi43
                _ = v44
            }
        case 1:
            v4 = v4
        }
    }
    v69 := GenS6{F0: 18446744073709551615, F1: 0, F2: (v4 >> ((v4) & 63))}
    var v70 uint16 = 0
    if (-128 == 1) {
        switch "a" {
        case "hello":
            var v71 int64 = 289
            _ = v71
        default:
            v2[("foo bar" + "a")] = -1
        }
        if false {
            fmt.Println(v16)
            v72 := []int{32768, 32767, 739}
            if false {
                fmt.Println(v16)
                for fi73 := 0; fi73 < 17; fi73++ {
                    var v74 int64 = 1
                    _ = fi73
                    _ = v74
                }
                v2["hello"] = 100
            } else {
                v1 -= v1
                v75 := GenS9{F0: (v70 - v70), F1: -9223372036854775808, F2: 29}
                _ = v75
            }
            _ = v72
        }
        v16 -= v16
    } else {
        _ = len(v2)
    }
    fmt.Println(v16, v4, v4)
    v4 = (^(v4))
    v76 := GenS5{F0: 520, F1: 96}
    v77 := v3[1]
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v70:", v70)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v69
    _ = v7
    _ = v70
    _ = v76
    _ = v77
}
