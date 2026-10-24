package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int; F2 uint64 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 int16 }
func genFn0() uint8 {
    var v7 int64 = -1
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := []int{(fi8 | fi8), fi8, (fi8 << ((255) & 63))}
        switch "\n\t" {
        case "foo bar":
            v10 := GenS1{F0: 36, F1: 18446744073709551615}
            _ = v10
        case "a":
            if (0 != 2147483647) {
                v9 = append(v9, (fi8 / ((fi8) | 1)))
                for fi11 := 0; fi11 < 1; fi11++ {
                    var v12 int64 = int64(fi8)
                    v13 := GenS1{F0: 0, F1: 4294967295}
                    switch "hello" {
                    case "a":
                        v12 += v12
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            } else {
                for fi14 := 0; fi14 < 15; fi14++ {
                    v7 += v7
                    var v15 float32 = ((-2.25 * float32(0.0)) - 0.1)
                    if false {
                        if len(v9) > 0 {
                            v9[0] = 9223372036854775807
                        }
                        var v16 uint = 85
                        fmt.Println(v7, v7, v7)
                        if ((false && true) && true) {
                            v17 := GenS2{F0: 42, F1: v7}
                            switch -1 {
                            case 32767:
                                v9 = append(v9, fi8)
                            case -710:
                                var v18 bool = (124 >= 2)
                                _ = v18
                            }
                            fmt.Println(v7)
                            _ = v17
                        }
                        if len(v9) > 0 {
                            v9[0] = -9223372036854775808
                        }
                        _ = v16
                    } else {
                        v19 := GenS3{F0: 2}
                        v7 -= -755
                        _ = v19
                    }
                    fmt.Println(v7)
                    v7 += 42
                    if false {
                        switch (("a" + ("a" + ((((("hello" + "a") + "hello") + "") + "") + "a"))) + "foo bar") {
                        case "a":
                            v20 := GenS3{F0: 0}
                            _ = v20
                        case "hello":
                            v15 = (((v15 + v15) / 3.14159) - float32(((1.5 + 0.1) - -2.25)))
                        case "":
                            switch (fi14 / (((fi14 % ((fi14) | 1))) | 1)) {
                            case -2147483648:
                                v21 := GenS4{F0: 0, F1: 32768, F2: 398}
                                _ = v21
                            case 1:
                                var v22 uint = 387
                                _ = v22
                            default:
                                fmt.Println(fi8, v7)
                            }
                        }
                        fmt.Println(fi14, v15, v7)
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    }
                    _ = fi14
                    _ = v15
                }
            }
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi8
        _ = v9
    }
    v23 := GenS3{F0: 127}
    _ = v23
    _ = v7
    return 127
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int8 = -128
    var v3 uint = (v1 / (((v1 & v1)) | 1))
    v4 := []int{}
    var v5 uint16 = 1
    var v6 uint16 = 56
    v24 := GenS1{F0: 255, F1: v1}
    fmt.Println(v2, v6, v3)
    var v25 string = "a"
    var v26 int = -679
    if true {
        for fi27 := 0; fi27 < 16; fi27++ {
            var v28 int64 = 1
            v29 := GenS4{F0: 2147483647, F1: v26, F2: 4294967295}
            for fi30 := 0; fi30 < 10; fi30++ {
                if true {
                    var v31 uint32 = 8
                    v4 = append(v4, v26)
                    _ = v31
                } else {
                    if ((!true) && (328 < 0)) {
                        _ = len(v4)
                        _ = len(v4)
                        var v32 bool = (v3 > v3)
                        _ = v32
                    }
                }
                if len(v4) > 0 {
                    v4[0] = 0
                }
                _ = fi30
            }
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    v26 += 4294967295
    for fi33 := 0; fi33 < 1; fi33++ {
        if (1 < 2) {
            v5 += 42
            if false {
                _ = len(v4)
                if true {
                    switch (v26 >> ((42) & 63)) {
                    case -1:
                        v5 = v5
                    case 32767:
                        fmt.Println(v1, v2, v2)
                    }
                }
            }
            v4 = append(v4, 2147483647)
        } else {
            fmt.Println(v1, v1, v2)
        }
        for fi34 := 0; fi34 < 1; fi34++ {
            var v35 int32 = 1
            if (!false) {
                if true {
                    fmt.Println(v35)
                } else {
                    var v36 float64 = (-2.25 / 3.14159)
                    var v37 bool = (v35 < (^(v35)))
                    if (v37 || (!(!v37))) {
                        if v37 {
                            v38 := GenS3{F0: 254}
                            _ = v38
                        } else {
                            var v39 int16 = 769
                            if (uint64(v1) <= 18446744073709551615) {
                                v40 := GenS2{F0: (v6 & v6), F1: 1}
                                v41 := GenS5{F0: 0, F1: (v2 * v2)}
                                _ = v40
                                _ = v41
                            }
                            _ = v39
                        }
                    }
                    _ = v36
                    _ = v37
                }
                if (-2147483648 < 714) {
                    if false {
                        var v42 int32 = v35
                        _ = v42
                    } else {
                        v4 = append(v4, 1)
                        if ((true && (!(1 > 9223372036854775807))) && false) {
                            fmt.Println(v2, v2, v35)
                            var v43 uint32 = 42
                            fmt.Println(fi33)
                            _ = v43
                        } else {
                            var v44 uint32 = 65535
                            switch v26 {
                            case 4294967295:
                                var v45 int = (fi34 * 0)
                                _ = v45
                            case 32767:
                                var v46 int8 = v2
                                _ = v46
                            default:
                                fmt.Println(v2, v2, v5)
                            }
                            v47 := GenS2{F0: 32767, F1: -9223372036854775808}
                            if len(v4) > 0 {
                                v4[0] = v26
                            }
                            switch "\n\t" {
                            case "\n\t":
                                var v48 uint64 = 4294967295
                                fmt.Println(fi33, v26, fi34)
                                fmt.Println(v35)
                                _ = v48
                            case "":
                                var v49 int64 = 845
                                _ = v49
                            default:
                                var v50 uint8 = 255
                                _ = v50
                            }
                            _ = v44
                            _ = v47
                        }
                        var v51 uint = 2
                        _ = v51
                    }
                } else {
                    if len(v4) > 0 {
                        v4[0] = -1
                    }
                }
                v26 = (fi34 << ((502) & 63))
                if (v26 <= 42) {
                    fmt.Println(v3, v6, v1)
                    var v52 int16 = -297
                    var v53 uint64 = 985
                    _ = v52
                    _ = v53
                } else {
                    var v54 float32 = 0.0
                    _ = len(v4)
                    fmt.Println(v3, v3)
                    v5 -= (v5 % ((v5) | 1))
                    _ = v54
                }
                fmt.Println(v1)
                var v55 int = (fi34 & -622)
                var v56 float32 = float32(3.14159)
                _ = v55
                _ = v56
            } else {
                fmt.Println(fi33)
                v1 -= (^(v1))
                v57 := GenS1{F0: genFn0(), F1: v1}
                _ = v57
            }
            fmt.Println(fi33)
            _ = fi34
            _ = v35
        }
        var v58 float64 = 3.14159
        _ = fi33
        _ = v58
    }
    var v59 int8 = (v2 / ((int8(v1)) | 1))
    if (true && ((v26 + 4294967295) <= (v26 >> ((v3) & 63)))) {
        var v60 float64 = 1e10
        v61 := GenS4{F0: -2147483648, F1: (^(v26)), F2: 0}
        var v62 int32 = 1
        _ = v60
        _ = v61
        _ = v62
    }
    v2 += v2
    var v63 bool = true
    switch (-(v26)) {
    case -2147483648:
        var v64 uint8 = genFn0()
        _ = v64
    case 409:
        v65 := GenS3{F0: 215}
        _ = v65
    }
    if (0 >= 1) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi66 := 0; fi66 < 20; fi66++ {
            var v67 int32 = 2147483647
            _ = fi66
            _ = v67
        }
    }
    var v68 uint16 = 2
    var v69 int32 = 2147483647
    v70 := GenS6{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v63
    _ = v68
    _ = v69
    _ = v70
}
