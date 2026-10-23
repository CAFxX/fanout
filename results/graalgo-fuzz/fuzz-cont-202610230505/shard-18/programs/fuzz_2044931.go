package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint32; F2 int }
type GenS4 struct { F0 bool; F1 int64; F2 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16; F1 int64 }
type GenS8 struct { F0 string; F1 int64; F2 int8 }
func genFn0(p0 uint16) string {
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{fi6, (fi6 << ((1) & 63)), (fi6 / ((fi6) | 1)), (fi6 / ((fi6) | 1))}
        v8 := GenS1{F0: 128, F1: ("hello" != "a"), F2: 15}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{}
    _ = p0
    _ = v9
    return (("\n\t" + "") + "a")
}

func main() {
    var v1 int8 = 0
    var v2 int8 = v1
    var v3 uint = 255
    var v4 int = 2147483647
    var v5 int32 = 2147483647
    switch ("a" + "") {
    case "foo bar":
        fmt.Println(v3, v1, v1)
    case "hello":
        switch "foo bar" {
        case "foo bar":
            v10 := GenS2{F0: 42}
            _ = v10
        case "\n\t":
            for fi11 := 0; fi11 < 16; fi11++ {
                switch v4 {
                case 0:
                    v12 := GenS1{F0: (v4 << (((v3 ^ 749)) & 63)), F1: (v3 == (v3 ^ 42)), F2: 0}
                    _ = v12
                case 9:
                    v13 := []int{(v4 >> ((4294967295) & 63)), 32768, (-(fi11))}
                    _ = v13
                }
                if true {
                    var v14 uint64 = 0
                    switch ("a" + ("a" + "")) {
                    case "":
                        fmt.Println(v2, v14, fi11)
                    default:
                        v15 := GenS3{F0: int32(v2), F1: 191, F2: v4}
                        _ = v15
                    }
                    fmt.Println(fi11, v4)
                    _ = v14
                }
                var v16 int = 32768
                _ = fi11
                _ = v16
            }
        }
    case "":
        v17 := []int{}
        _ = v17
    default:
        fmt.Println(v1, v5, v2)
    }
    v3 = v3
    if (!(("" + ("a" + ("hello" + (("\n\t" + (("hello" + (("foo bar" + ("" + ("foo bar" + "hello"))) + "a")) + "\n\t")) + "\n\t")))) > "\n\t")) {
        v1 = v2
        v18 := GenS4{F0: true, F1: -9223372036854775808, F2: 42}
        v3 = v3
        _ = v18
    } else {
        v19 := []int{(v4 & (v4 - -9223372036854775808)), v4, (v4 % (((v4 & -2147483648)) | 1))}
        _ = v19
    }
    v3 -= (v3 | v3)
    v20 := GenS4{F0: false, F1: -9223372036854775808, F2: -859}
    fmt.Println(v4)
    if true {
        switch 42 {
        case 0:
            switch v4 {
            case 0:
                if false {
                    switch (v4 - v4) {
                    case 115:
                        v21 := GenS2{F0: 2}
                        _ = v21
                    case -2147483648:
                        for fi22 := 0; fi22 < 18; fi22++ {
                            for fi23 := 0; fi23 < 20; fi23++ {
                                v2 = v1
                                _ = fi23
                            }
                            _ = fi22
                        }
                    case -1:
                        if true {
                            v24 := []int{v4, v4}
                            var v25 int64 = -726
                            var v26 int32 = int32(v3)
                            _ = v24
                            _ = v25
                            _ = v26
                        } else {
                            if (9223372036854775807 > 42) {
                                if ((float64((1e10 + (1.5 * 1.5))) + (3.14159 - ((float64(1.5) - float64((1.5 / 0.1))) * 1.5))) <= 3.14159) {
                                    switch v4 {
                                    case 522:
                                        if true {
                                            v4 = (v4 + 9223372036854775807)
                                            v1 = (v1 >> ((18446744073709551615) & 63))
                                        } else {
                                            var v27 uint16 = uint16(v3)
                                            _ = v27
                                        }
                                    case 32768:
                                        var v28 float64 = (-2.25 / ((float64((-2.25 - (float32(-2.25) / 1.0))) - 1e10) / 1.5))
                                        _ = v28
                                    default:
                                        if (false || (2 >= 18446744073709551615)) {
                                            v2 = (v2 * int8(v4))
                                            var v29 float32 = 3.14159
                                            for fi30 := 0; fi30 < 5; fi30++ {
                                                for fi31 := 0; fi31 < 12; fi31++ {
                                                    var v32 int64 = 1
                                                    var v33 string = "\n\t"
                                                    var v34 string = v33
                                                    _ = fi31
                                                    _ = v32
                                                    _ = v33
                                                    _ = v34
                                                }
                                                fmt.Println(fi30)
                                                _ = fi30
                                            }
                                            _ = v29
                                        }
                                    }
                                    for fi35 := 0; fi35 < 3; fi35++ {
                                        var v36 float64 = 3.14159
                                        for fi37 := 0; fi37 < 9; fi37++ {
                                            var v38 int16 = -32768
                                            _ = fi37
                                            _ = v38
                                        }
                                        _ = fi35
                                        _ = v36
                                    }
                                } else {
                                    switch (v4 + v4) {
                                    case 32768:
                                        var v39 int = v4
                                        _ = v39
                                    case 127:
                                        v40 := GenS5{F0: (v4 % (((v4 << ((v3) & 63))) | 1))}
                                        _ = v40
                                    case 4294967295:
                                        var v41 uint = v3
                                        _ = v41
                                    }
                                }
                                var v42 int8 = (-(v2))
                                if (v1 != v2) {
                                    if false {
                                        var v43 uint64 = 2
                                        v44 := []int{v4, (v4 + v4)}
                                        var v45 int16 = 1000
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    } else {
                                        v46 := GenS6{F0: 1000}
                                        _ = v46
                                    }
                                    switch 32768 {
                                    case 4294967295:
                                        var v47 int8 = v42
                                        _ = v47
                                    case -139:
                                        var v48 float32 = 0.0
                                        _ = v48
                                    default:
                                        if true {
                                            for fi49 := 0; fi49 < 13; fi49++ {
                                                var v50 float32 = (1.5 + 0.0)
                                                var v51 int32 = (v5 * v5)
                                                fmt.Println(v4, v2)
                                                _ = fi49
                                                _ = v50
                                                _ = v51
                                            }
                                            if (!(!false)) {
                                                v5 = 1
                                                for fi52 := 0; fi52 < 20; fi52++ {
                                                    v5 = (v5 << ((v3) & 63))
                                                    _ = fi52
                                                }
                                                v53 := []int{(v4 + 1), (v4 * -9223372036854775808)}
                                                _ = v53
                                            } else {
                                                v54 := []int{(v4 << ((v3) & 63)), 32768, v4, v4}
                                                v2 = (^(v42))
                                                _ = v54
                                            }
                                        }
                                    }
                                } else {
                                    var v55 uint64 = 2
                                    switch v4 {
                                    case -2147483648:
                                        for fi56 := 0; fi56 < 16; fi56++ {
                                            v57 := GenS7{F0: int16(v1), F1: 0}
                                            if false {
                                                var v58 uint64 = v55
                                                _ = v58
                                            } else {
                                                v4 -= (fi56 ^ (v4 << ((217) & 63)))
                                                if (42 <= 4294967295) {
                                                    var v59 uint32 = 0
                                                    _ = v59
                                                }
                                                var v60 uint16 = 2
                                                _ = v60
                                            }
                                            _ = fi56
                                            _ = v57
                                        }
                                    case 2147483647:
                                        if false {
                                            if (v4 == v4) {
                                                for fi61 := 0; fi61 < 5; fi61++ {
                                                    v62 := GenS7{F0: -1, F1: 1}
                                                    _ = fi61
                                                    _ = v62
                                                }
                                                for fi63 := 0; fi63 < 6; fi63++ {
                                                    var v64 uint32 = 65535
                                                    v4 -= v4
                                                    _ = fi63
                                                    _ = v64
                                                }
                                            } else {
                                                fmt.Println(v3, v42, v5)
                                            }
                                            if (128 != 71) {
                                                fmt.Println(v42)
                                                if false {
                                                    var v65 uint8 = 42
                                                    _ = v65
                                                }
                                                var v66 float32 = (float32(3.14159) * float32(float64(1e10)))
                                                _ = v66
                                            } else {
                                                switch (v4 ^ (v4 ^ v4)) {
                                                case 140:
                                                    for fi67 := 0; fi67 < 7; fi67++ {
                                                        v68 := GenS8{F0: "\n\t", F1: 18, F2: (-(v2))}
                                                        v4 += 127
                                                        var v69 int8 = -1
                                                        _ = fi67
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                case 128:
                                                    v3 = (v3 * 18446744073709551615)
                                                }
                                                var v70 int8 = v2
                                                _ = v70
                                            }
                                            var v71 uint8 = 128
                                            _ = v71
                                        }
                                    case 32767:
                                        v5 = -29
                                    default:
                                        if (1.5 > -2.25) {
                                            v1 = (-(v1))
                                            v72 := GenS6{F0: -1}
                                            v5 = 596
                                            _ = v72
                                        }
                                    }
                                    _ = v55
                                }
                                _ = v42
                            } else {
                                v3 = (v3 >> (((^(v3))) & 63))
                            }
                        }
                    }
                    for fi73 := 0; fi73 < 13; fi73++ {
                        fmt.Println(v2, fi73, v4)
                        switch v4 {
                        case 32768:
                            v74 := []int{v4, 937, fi73, (v4 * (fi73 % ((4294967295) | 1)))}
                            _ = v74
                        case 9223372036854775807:
                            fmt.Println(v5, v5)
                        case -482:
                            v75 := []int{(fi73 % (((fi73 % (((fi73 ^ fi73)) | 1))) | 1))}
                            _ = v75
                        }
                        _ = fi73
                    }
                } else {
                    v76 := []int{(v4 * (v4 & (^(v4)))), v4}
                    _ = len(v76)
                    _ = v76
                }
            default:
                var v77 uint = v3
                _ = v77
            }
        }
    }
    var v78 int16 = 1000
    v1 += (v1 * v1)
    for fi79 := 0; fi79 < 6; fi79++ {
        var v80 int32 = v5
        _ = fi79
        _ = v80
    }
    var v81 uint16 = 1
    v5 = (v5 / ((v5) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v78:", v78)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v78
    _ = v81
}
