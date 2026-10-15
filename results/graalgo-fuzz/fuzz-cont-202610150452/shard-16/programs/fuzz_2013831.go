package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 int; F1 uint16; F2 int16 }
type GenS8 struct { F0 uint8; F1 int64; F2 int }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool; F1 uint64 }
type GenS11 struct { F0 uint8; F1 int }
type GenS12 struct { F0 string; F1 uint64 }

func main() {
    var v1 int16 = 42
    var v2 float32 = (-2.25 / 1e10)
    var v3 string = ("foo bar" + ("foo bar" + "foo bar"))
    var v4 uint8 = 66
    fmt.Println(v4, v3, v2)
    fmt.Println(v4, v2)
    var v5 bool = false
    v5 = false
    v6 := GenS1{F0: 183}
    v7 := []int{127, 865, 1}
    v8 := GenS1{F0: 0}
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 int32 = 1
            v5 = (fi9 <= (-(fi10)))
            _ = fi10
            _ = v11
        }
        v4 -= v4
        _ = fi9
    }
    switch (v3 + v3) {
    case "\n\t":
        v12 := GenS1{F0: 603}
        _ = v12
    case "foo bar":
        for fi13 := 0; fi13 < 18; fi13++ {
            if len(v7) > 0 {
                v7[0] = (-(fi13))
            }
            if false {
                switch 0 {
                case -270:
                    var v14 uint64 = 42
                    _ = v14
                }
                if v5 {
                    var v15 uint8 = (v4 * (v4 >> ((65535) & 63)))
                    v7 = append(v7, -1)
                    _ = v15
                }
            } else {
                if true {
                    fmt.Println(v4, v1, v3)
                    v16 := GenS2{F0: -747, F1: 255, F2: 18446744073709551615}
                    _ = v16
                } else {
                    v17 := GenS3{F0: 2, F1: 1, F2: 1}
                    _ = v17
                }
                for fi18 := 0; fi18 < 8; fi18++ {
                    var v19 uint = 8
                    switch (-(fi18)) {
                    case 9223372036854775807:
                        var v20 string = v3
                        var v21 float32 = v2
                        if v5 {
                            var v22 int32 = -2147483648
                            _ = v22
                        } else {
                            v21 = v21
                            var v23 int = 127
                            var v24 bool = ((!v5) && true)
                            _ = v23
                            _ = v24
                        }
                        v2 = 0.1
                        _ = v20
                        _ = v21
                    case 158:
                        switch ("hello" + v3) {
                        case "hello":
                            fmt.Println(v1)
                            switch v3 {
                            case "\n\t":
                                v1 -= 588
                            default:
                                v1 += 1000
                            }
                            v2 -= (v2 - (v2 + (v2 + 1e10)))
                        default:
                            v4 += (v4 ^ v4)
                        }
                    case 428:
                        v7 = append(v7, (-(fi18)))
                        if ((v1 >> ((v19) & 63)) >= (v1 * (v1 >> (((v19 >> (((v19 * (v19 + 413))) & 63))) & 63)))) {
                            fmt.Println(v2)
                            if (-2.25 == v2) {
                                var v25 uint32 = 1
                                var v26 bool = ((1.5 / 1.0) != (float64(v2) / (1e10 * float64(v2))))
                                if (((v4 * 1) >= (v4 | (^(v4)))) || v26) {
                                    if true {
                                        _ = len(v7)
                                        v5 = (v5 || false)
                                        v1 -= 99
                                    } else {
                                        if v26 {
                                            var v27 string = v3
                                            if len(v7) > 0 {
                                                v7[0] = (fi18 & fi18)
                                            }
                                            v28 := GenS4{F0: fi18}
                                            v26 = v26
                                            v3 = (v27 + v3)
                                            _ = v27
                                            _ = v28
                                        }
                                        switch v3 {
                                        case "hello":
                                            v3 = ((v3 + v3) + v3)
                                        }
                                    }
                                }
                                if ((^(v19)) > v19) {
                                    if ((((^(v1)) != (^(v1))) || (false && v26)) && true) {
                                        if (!v5) {
                                            v7 = append(v7, -9223372036854775808)
                                            v25 = v25
                                        }
                                        v26 = true
                                    } else {
                                        var v29 uint8 = uint8(v19)
                                        if v5 {
                                            var v30 uint = 42
                                            if v5 {
                                                if (!false) {
                                                    if len(v7) > 0 {
                                                        v7[0] = 2147483647
                                                    }
                                                    if len(v7) > 0 {
                                                        v7[0] = 32767
                                                    }
                                                    var v31 uint32 = v25
                                                    v25 = (^(v31))
                                                    v7 = append(v7, (-(fi18)))
                                                    switch ((v3 + ((v3 + v3) + "\n\t")) + v3) {
                                                    case "\n\t":
                                                        if (!(!true)) {
                                                            if len(v7) >= 2 {
                                                                v7 = v7[1:len(v7)]
                                                            }
                                                            if v5 {
                                                                v32 := GenS5{F0: ""}
                                                                switch fi13 {
                                                                case -205:
                                                                    var v33 int16 = (v1 - 32767)
                                                                    _ = v33
                                                                case 906:
                                                                    v4 = (^(v4))
                                                                default:
                                                                    v31 -= (v31 << ((v30) & 63))
                                                                }
                                                                switch (fi13 / (((-(fi13))) | 1)) {
                                                                case -218:
                                                                    if v5 {
                                                                        v34 := GenS4{F0: fi13}
                                                                        v1 += (v1 << ((v30) & 63))
                                                                        if v5 {
                                                                            var v35 int32 = 42
                                                                            if (!((!true) || (!((v30 ^ (v19 * 1)) != 610)))) {
                                                                                if len(v7) > 0 {
                                                                                    v7[0] = (-(fi18))
                                                                                }
                                                                                var v36 uint = v30
                                                                                _ = v36
                                                                            }
                                                                            var v37 int64 = -1
                                                                            _ = v35
                                                                            _ = v37
                                                                        }
                                                                        _ = v34
                                                                    } else {
                                                                        switch 4294967295 {
                                                                        case 128:
                                                                            if ((v31 + v25) <= (v31 - 582)) {
                                                                                if v26 {
                                                                                    var v38 float32 = 1.5
                                                                                    if (42 == 100) {
                                                                                        v1 = 370
                                                                                        v39 := GenS4{F0: (fi13 - (fi18 / (((fi18 - (fi18 * (fi18 & -878)))) | 1)))}
                                                                                        if v5 {
                                                                                            v40 := GenS6{F0: 1, F1: 65535}
                                                                                            v38 = v38
                                                                                            _ = v40
                                                                                        }
                                                                                        _ = v39
                                                                                    }
                                                                                    _ = v38
                                                                                } else {
                                                                                    v1 += v1
                                                                                    var v41 bool = (!false)
                                                                                    _ = v41
                                                                                }
                                                                                if false {
                                                                                    _ = len(v7)
                                                                                    v1 = (v1 / ((v1) | 1))
                                                                                }
                                                                            } else {
                                                                                v3 = ("foo bar" + ("hello" + (v3 + v3)))
                                                                                v19 += v30
                                                                            }
                                                                        case -21:
                                                                            if ((!(v26 || (false && (int8(v25) < 1)))) || (v25 < (v25 - 556))) {
                                                                                var v42 string = (v3 + v3)
                                                                                v43 := GenS7{F0: (fi13 & (fi18 & fi18)), F1: 17, F2: (v1 >> ((2) & 63))}
                                                                                _ = v42
                                                                                _ = v43
                                                                            } else {
                                                                                if (v5 && (4294967295 > 2)) {
                                                                                    var v44 uint = v30
                                                                                    var v45 float64 = 0.1
                                                                                    var v46 string = ""
                                                                                    _ = v44
                                                                                    _ = v45
                                                                                    _ = v46
                                                                                }
                                                                            }
                                                                        }
                                                                        var v47 int32 = -404
                                                                        v30 -= (^(v19))
                                                                        _ = v47
                                                                    }
                                                                case 32767:
                                                                    v5 = ((v31 ^ (v31 % ((uint32(v19)) | 1))) <= v25)
                                                                default:
                                                                    fmt.Println(v1, fi13)
                                                                }
                                                                fmt.Println(fi13, fi13)
                                                                var v48 uint8 = 169
                                                                _ = v32
                                                                _ = v48
                                                            } else {
                                                                var v49 uint64 = 563
                                                                if len(v7) >= 2 {
                                                                    v7 = v7[1:len(v7)]
                                                                }
                                                                var v50 int32 = 2147483647
                                                                _ = v49
                                                                _ = v50
                                                            }
                                                            v51 := GenS8{F0: v4, F1: 387, F2: -9223372036854775808}
                                                            _ = v51
                                                        } else {
                                                            v4 -= (v29 / ((uint8(v31)) | 1))
                                                        }
                                                    case "a":
                                                        switch 127 {
                                                        case 1:
                                                            if true {
                                                                v31 -= 530
                                                            } else {
                                                                v52 := GenS8{F0: (v4 ^ v29), F1: -9223372036854775808, F2: 42}
                                                                v7 = append(v7, (fi18 << ((1) & 63)))
                                                                _ = v52
                                                            }
                                                        case 9223372036854775807:
                                                            v53 := GenS4{F0: int(v30)}
                                                            _ = v53
                                                        case 4294967295:
                                                            fmt.Println(fi18, fi18)
                                                            v54 := GenS9{F0: 255}
                                                            if (908 != 2147483647) {
                                                                v26 = v5
                                                            } else {
                                                                v55 := GenS8{F0: (v29 * v29), F1: 117, F2: fi18}
                                                                var v56 string = (v3 + v3)
                                                                v57 := GenS10{F0: (!((^(fi18)) >= (fi13 - (fi13 << (((^(v30))) & 63))))), F1: 511}
                                                                _ = v55
                                                                _ = v56
                                                                _ = v57
                                                            }
                                                            _ = v54
                                                        }
                                                        if v5 {
                                                            if (v2 <= 0.0) {
                                                                if false {
                                                                    v58 := GenS10{F0: ((v26 || v26) && (!v26)), F1: 841}
                                                                    if true {
                                                                        fmt.Println(v31, v4, v31)
                                                                    }
                                                                    _ = v58
                                                                }
                                                                v7 = append(v7, fi18)
                                                                var v59 string = ("hello" + v3)
                                                                _ = v59
                                                            }
                                                        } else {
                                                            if v26 {
                                                                var v60 bool = v5
                                                                fmt.Println(v5)
                                                                v61 := GenS11{F0: 1, F1: fi18}
                                                                _ = v60
                                                                _ = v61
                                                            } else {
                                                                var v62 uint64 = 2
                                                                _ = v62
                                                            }
                                                            _ = len(v7)
                                                        }
                                                    }
                                                    _ = v31
                                                } else {
                                                    switch (-(fi13)) {
                                                    case -158:
                                                        var v63 uint8 = v29
                                                        _ = v63
                                                    case 127:
                                                        v2 = (v2 * ((v2 / v2) + v2))
                                                    case 1:
                                                        v1 -= v1
                                                    default:
                                                        if len(v7) > 0 {
                                                            v7[0] = fi18
                                                        }
                                                    }
                                                }
                                                v3 = ""
                                            } else {
                                                var v64 uint64 = 450
                                                _ = v64
                                            }
                                            _ = v30
                                        }
                                        switch (^(fi18)) {
                                        case -61:
                                            var v65 uint64 = 18446744073709551615
                                            if false {
                                                v66 := GenS8{F0: v29, F1: 0, F2: fi18}
                                                v7 = append(v7, (^(fi18)))
                                                var v67 bool = false
                                                _ = v66
                                                _ = v67
                                            } else {
                                                var v68 string = "hello"
                                                var v69 uint64 = 18446744073709551615
                                                _ = v68
                                                _ = v69
                                            }
                                            _ = v65
                                        case 0:
                                            fmt.Println(v3, v29, v4)
                                        default:
                                            v7 = append(v7, (fi18 & 0))
                                        }
                                        _ = v29
                                    }
                                    v70 := GenS6{F0: 1, F1: 42}
                                    v5 = (-9223372036854775808 >= -1)
                                    _ = v70
                                } else {
                                    fmt.Println(v1)
                                }
                                _ = v25
                                _ = v26
                            } else {
                                v1 -= v1
                                var v71 int64 = 1
                                _ = v71
                            }
                        } else {
                            var v72 float32 = v2
                            v73 := GenS2{F0: 0, F1: v19, F2: 969}
                            v19 -= (v19 << (((v19 << (((v19 % ((v19) | 1))) & 63))) & 63))
                            _ = v72
                            _ = v73
                        }
                        fmt.Println(v19, v2, v5)
                    }
                    _ = fi18
                    _ = v19
                }
                if (v5 && true) {
                    for fi74 := 0; fi74 < 8; fi74++ {
                        var v75 int = fi74
                        _ = fi74
                        _ = v75
                    }
                } else {
                    v4 = v4
                    for fi76 := 0; fi76 < 19; fi76++ {
                        var v77 uint8 = (v4 & v4)
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        _ = fi76
                        _ = v77
                    }
                    v78 := GenS8{F0: v4, F1: 1, F2: (fi13 << ((4294967295) & 63))}
                    _ = v78
                }
            }
            if true {
                v79 := GenS2{F0: -75, F1: 65535, F2: 18446744073709551615}
                _ = v79
            }
            _ = fi13
        }
    }
    v80 := GenS12{F0: (((v3 + v3) + "a") + ((v3 + "foo bar") + v3)), F1: 4294967295}
    for fi81 := 0; fi81 < 2; fi81++ {
        switch 0 {
        case 1:
            v7 = append(v7, 874)
        case -1:
            for fi82 := 0; fi82 < 17; fi82++ {
                v2 = -2.25
                _ = fi82
            }
        case 2147483647:
            fmt.Println(v2)
        default:
            v1 += (v1 << ((2) & 63))
        }
        fmt.Println(fi81)
        _ = fi81
    }
    v1 += v1
    if len(v7) > 0 {
        v7[0] = -1
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v80
}
