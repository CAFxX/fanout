package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 int16; F1 bool }
type GenS8 struct { F0 int16; F1 uint; F2 int }
type GenS9 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint8, p1 uint32, p2 string) int8 {
    var v5 int64 = 9223372036854775807
    fmt.Println(p0)
    for fi6 := 0; fi6 < 16; fi6++ {
        if true {
            var v7 int8 = -128
            var v8 float32 = 1e10
            _ = v7
            _ = v8
        } else {
            var v9 bool = (true && false)
            v10 := []int{(^(fi6))}
            var v11 int8 = 40
            _ = v10
            _ = v11
            _ = v9
        }
        var v12 float64 = 0.1
        _ = fi6
        _ = v12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS1{F0: 32768}
        v15 := GenS1{F0: fi13}
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 100
}
func genFn1(p0 uint16) int8 {
    v16 := []int{-517, 32768}
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    switch 128 {
    case 32767:
        v17 := GenS2{F0: 0, F1: 1}
        _ = v17
    case -9223372036854775808:
        switch (("" + "") + "a") {
        case "":
            v18 := GenS3{F0: 255, F1: 0}
            _ = v18
        case "\n\t":
            fmt.Println(p0)
        }
    case 202:
        v19 := GenS4{F0: 2, F1: 42}
        _ = v19
    }
    _ = len(v16)
    var v20 int64 = 168
    _ = p0
    _ = v16
    _ = v20
    return 2
}

func main() {
    var v1 uint32 = 42
    var v2 uint16 = 255
    var v3 uint32 = (v1 << ((4) & 63))
    v4 := []uint32{579}
    var v21 int8 = -57
    if len(v4) > 0 {
        v4[0] = v1
    }
    v2 = v2
    v4 = append(v4, (v1 | v3))
    v1 = 0
    v22 := GenS2{F0: genFn1(v2), F1: -1}
    if true {
        var v23 uint = 4294967295
        var v24 float32 = (float32((0.1 * 0.1)) * float32((0.0 / 1e10)))
        _ = v23
        _ = v24
    }
    var v25 int32 = 2147483647
    var v26 int16 = -32768
    var v27 float32 = (float32(3.14159) + 3.14159)
    var v28 bool = false
    if true {
        v2 = 752
        v25 = (v25 | int32(v26))
    } else {
        var v29 int8 = (v21 >> ((1) & 63))
        _ = v29
    }
    var v30 int8 = 42
    switch -98 {
    case 2147483647:
        switch -2147483648 {
        case 283:
            if true {
                v4 = append(v4, 608)
                var v31 uint32 = 853
                var v32 string = "hello"
                _ = v31
                _ = v32
            } else {
                fmt.Println(v1, v30, v28)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v2 -= v2
            }
        case 32768:
            _ = len(v4)
        }
    case 4294967295:
        var v33 bool = ((v3 | 2147483647) > v3)
        _ = v33
    default:
        v28 = ((42 == 42) && ((-(v21)) <= (v21 ^ v30)))
    }
    v21 = genFn1(v2)
    for fi34 := 0; fi34 < 4; fi34++ {
        v4 = append(v4, (v1 - v1))
        v35 := GenS1{F0: (fi34 >> ((255) & 63))}
        _ = fi34
        _ = v35
    }
    if false {
        if (42 != 0) {
            for fi36 := 0; fi36 < 6; fi36++ {
                if (v30 == (-(v21))) {
                    if len(v4) > 0 {
                        v4[0] = v3
                    }
                    fmt.Println(v3, v21, v2)
                }
                fmt.Println(v21, v25)
                _ = fi36
            }
            v37 := GenS5{F0: v30}
            if (!(v26 != (v26 << ((374) & 63)))) {
                var v38 int16 = v26
                v27 += 0.1
                _ = v38
            }
            _ = v37
        } else {
            v27 -= v27
            for fi39 := 0; fi39 < 4; fi39++ {
                var v40 int = fi39
                v30 += v30
                switch 2147483647 {
                case -2147483648:
                    v21 -= (v30 >> ((4294967295) & 63))
                case 9223372036854775807:
                    if v28 {
                        if v28 {
                            fmt.Println(v26, v28, v2)
                            if len(v4) > 0 {
                                v4[0] = v1
                            }
                            for fi41 := 0; fi41 < 8; fi41++ {
                                var v42 string = "\n\t"
                                var v43 uint64 = 18446744073709551615
                                v2 += (v2 >> ((146) & 63))
                                var v44 int16 = (v26 - (v26 - (^(v26))))
                                _ = fi41
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                        }
                        for fi45 := 0; fi45 < 14; fi45++ {
                            switch 128 {
                            case 32768:
                                var v46 float32 = (3.14159 / ((float32(-2.25) / v27) + ((((v27 - v27) + float32(0.0)) - v27) - float32(1.5))))
                                switch 9223372036854775807 {
                                case 4294967295:
                                    v27 = (v46 - v46)
                                    var v47 int8 = (-(v21))
                                    _ = v47
                                case 0:
                                    v4 = append(v4, (v3 | 0))
                                    var v48 int8 = -128
                                    switch 9223372036854775807 {
                                    case -539:
                                        v4 = append(v4, v3)
                                    }
                                    _ = v48
                                case -1:
                                    fmt.Println(v25, v30)
                                }
                                if v28 {
                                    if ((!true) && v28) {
                                        if (254 != int64(v1)) {
                                            v1 -= 2
                                            switch 930 {
                                            case 0:
                                                if v28 {
                                                    v49 := GenS6{F0: v26, F1: (v1 & (v1 << ((42) & 63)))}
                                                    _ = v49
                                                }
                                            case -2147483648:
                                                var v50 int32 = int32(v2)
                                                _ = v50
                                            }
                                            if true {
                                                v51 := GenS7{F0: -32768, F1: v28}
                                                var v52 bool = (v46 == 3.14159)
                                                v46 = (v46 / float32((float64(1e10) - ((float64(3.14159) / 3.14159) + (float64((((v27 - (((v46 - v27) * 1.5) + (-2.25 / 1.5))) + ((v27 - v27) - v27)) + 0.0)) + 3.14159)))))
                                                _ = v51
                                                _ = v52
                                            }
                                        }
                                        fmt.Println(v46, v46, v27)
                                        v53 := GenS5{F0: int8(v26)}
                                        _ = v53
                                    }
                                    var v54 int8 = (-(v30))
                                    _ = v54
                                }
                                if v28 {
                                    v25 -= v25
                                    v1 = (^(v3))
                                    if (false || true) {
                                        if ((!false) || v28) {
                                            var v55 float32 = (v46 / v46)
                                            _ = v55
                                        }
                                        var v56 uint16 = v2
                                        _ = v56
                                    } else {
                                        v25 -= 2147483647
                                        switch "a" {
                                        case "foo bar":
                                            switch fi45 {
                                            case 1:
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case -1:
                                                switch (fi45 >> ((2) & 63)) {
                                                case 127:
                                                    var v57 uint8 = 128
                                                    _ = v57
                                                case -9223372036854775808:
                                                    if v28 {
                                                        var v58 uint8 = 42
                                                        v4 = append(v4, v1)
                                                        _ = v58
                                                    } else {
                                                        if v28 {
                                                            if true {
                                                                var v59 bool = (4294967295 >= 872)
                                                                v28 = true
                                                                _ = v59
                                                            } else {
                                                                var v60 string = "a"
                                                                var v61 uint8 = 128
                                                                if len(v4) > 0 {
                                                                    v4[0] = v1
                                                                }
                                                                var v62 uint8 = (v61 - (v61 & uint8(v2)))
                                                                _ = v60
                                                                _ = v61
                                                                _ = v62
                                                            }
                                                            var v63 uint64 = 0
                                                            _ = v63
                                                        } else {
                                                            if v28 {
                                                                if v28 {
                                                                    v30 += v21
                                                                }
                                                                fmt.Println(v28, fi39)
                                                            }
                                                        }
                                                    }
                                                case -2147483648:
                                                    if v28 {
                                                        v64 := GenS3{F0: 255, F1: v25}
                                                        v65 := GenS3{F0: 744, F1: (^(v25))}
                                                        fmt.Println(v40)
                                                        if len(v4) > 0 {
                                                            v4[0] = 546
                                                        }
                                                        v46 = 0.0
                                                        if ((v26 >= 1) && v28) {
                                                            switch (v40 << ((42) & 63)) {
                                                            case 32767:
                                                                var v66 uint64 = 532
                                                                v4 = append(v4, v1)
                                                                _ = v66
                                                            case 4294967295:
                                                                v67 := GenS5{F0: 42}
                                                                _ = v67
                                                            case 914:
                                                                v28 = (!(((("" + ("" + "")) + "hello") + "") >= "\n\t"))
                                                                v68 := GenS8{F0: -649, F1: 255, F2: int(v2)}
                                                                if v28 {
                                                                    v69 := GenS2{F0: v30, F1: -32768}
                                                                    var v70 uint16 = 0
                                                                    var v71 float32 = float32(0.1)
                                                                    _ = v69
                                                                    _ = v70
                                                                    _ = v71
                                                                }
                                                                _ = v68
                                                            }
                                                            if len(v4) > 0 {
                                                                v4[0] = v1
                                                            }
                                                        } else {
                                                            var v72 string = "\n\t"
                                                            v26 -= v26
                                                            _ = v72
                                                        }
                                                        var v73 uint16 = v2
                                                        var v74 float32 = 0.0
                                                        var v75 float64 = 0.0
                                                        v76 := GenS7{F0: (v26 * v26), F1: false}
                                                        _ = v64
                                                        _ = v65
                                                        _ = v73
                                                        _ = v74
                                                        _ = v75
                                                        _ = v76
                                                    } else {
                                                        v77 := GenS1{F0: fi45}
                                                        _ = v77
                                                    }
                                                default:
                                                    if false {
                                                        v3 = (v1 >> ((332) & 63))
                                                        var v78 uint32 = 42
                                                        _ = v78
                                                    } else {
                                                        var v79 bool = ((v25 + v25) >= v25)
                                                        _ = v79
                                                    }
                                                }
                                            }
                                            v80 := GenS6{F0: (v26 - 0), F1: (v3 + (v3 * v3))}
                                            var v81 bool = v28
                                            _ = v80
                                            _ = v81
                                        default:
                                            v4 = append(v4, (v3 | v1))
                                        }
                                        v82 := GenS9{F0: v28, F1: (v2 / ((v2) | 1))}
                                        v2 = v2
                                        _ = v82
                                    }
                                }
                                if true {
                                    var v83 uint8 = 255
                                    _ = v83
                                }
                                var v84 int64 = 0
                                var v85 uint32 = (v1 - v3)
                                _ = v46
                                _ = v84
                                _ = v85
                            case 32767:
                                v86 := GenS1{F0: v40}
                                _ = v86
                            default:
                                if (v28 && (v30 > (v30 / ((v30) | 1)))) {
                                    var v87 uint64 = 2
                                    _ = v87
                                }
                            }
                            if len(v4) > 0 {
                                v4[0] = v3
                            }
                            var v88 string = "\n\t"
                            _ = fi45
                            _ = v88
                        }
                        v89 := GenS3{F0: 129, F1: 2147483647}
                        _ = v89
                    }
                case -919:
                    if false {
                        var v90 int8 = (v21 + 115)
                        if v28 {
                            if v28 {
                                v2 += (^(v2))
                                fmt.Println(fi39, v25, v26)
                                v3 = v1
                            }
                            if (true || (!(-1 >= 42))) {
                                var v91 float64 = 0.1
                                _ = v91
                            }
                        }
                        _ = v90
                    }
                }
                _ = fi39
                _ = v40
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
}
