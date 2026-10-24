package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int64 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 int64; F2 uint64 }
func genFn0() int8 {
    var v5 uint16 = 32767
    fmt.Println(v5)
    v6 := GenS1{F0: 21, F1: 127, F2: 1}
    _ = v5
    _ = v6
    return -69
}
func genFn1() uint32 {
    var v7 uint64 = 567
    var v8 uint64 = 102
    fmt.Println(v8, v8)
    v9 := GenS1{F0: 100, F1: -2147483648, F2: -9223372036854775808}
    _ = v7
    _ = v8
    _ = v9
    return 393
}

func main() {
    var v1 int64 = -188
    v2 := make(map[uint32]string)
    var v3 string = ("\n\t" + "foo bar")
    var v4 uint32 = 0
    var v10 int32 = 0
    if false {
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := GenS2{F0: -32768, F1: 503}
            var v13 string = ((((v3 + v3) + v3) + v3) + ("hello" + ("\n\t" + v3)))
            var v14 int = (fi11 + (fi11 | (fi11 / (((fi11 - (fi11 * fi11))) | 1))))
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        if false {
            v15 := GenS1{F0: genFn0(), F1: 1, F2: 42}
            _ = v15
        }
        v16 := GenS1{F0: 127, F1: 42, F2: (v1 >> ((830) & 63))}
        _ = v16
    }
    var v17 uint8 = 127
    v18 := []int{}
    if false {
        var v19 int = 1
        _ = v19
    }
    if (true || (!true)) {
        var v20 float32 = 0.1
        var v21 string = "a"
        _ = v20
        _ = v21
    }
    if true {
        var v22 bool = ((^(v4)) == 0)
        if (!v22) {
            for fi23 := 0; fi23 < 9; fi23++ {
                v24 := GenS3{F0: -1}
                v17 = 127
                _ = fi23
                _ = v24
            }
            if len(v18) > 0 {
                v18[0] = 127
            }
        }
        var v25 int = 0
        _ = v22
        _ = v25
    } else {
        switch v3 {
        case "\n\t":
            v26 := GenS2{F0: int16(v10), F1: 725}
            _ = v26
        case "foo bar":
            v27 := GenS2{F0: 42, F1: 32767}
            _ = v27
        default:
            var v28 bool = (v17 > 4)
            _ = v28
        }
        for fi29 := 0; fi29 < 17; fi29++ {
            for fi30 := 0; fi30 < 3; fi30++ {
                var v31 uint = 930
                v3 = v3
                _ = fi30
                _ = v31
            }
            _ = fi29
        }
    }
    v4 = (v4 ^ (v4 | (v4 / ((v4) | 1))))
    for fi32 := 0; fi32 < 11; fi32++ {
        switch v3 {
        case "hello":
            for fi33 := 0; fi33 < 12; fi33++ {
                v1 += v1
                v18 = append(v18, 42)
                _ = fi33
            }
        case "":
            fmt.Println(v17)
        case "foo bar":
            var v34 uint16 = 32767
            _ = v34
        }
        _ = len(v2)
        v1 -= (-(v1))
        _ = fi32
    }
    var v35 int8 = -128
    _ = len(v18)
    if true {
        var v36 uint64 = uint64(v4)
        _ = v36
    } else {
        for fi37 := 0; fi37 < 11; fi37++ {
            if (!(!false)) {
                for fi38 := 0; fi38 < 20; fi38++ {
                    if (true && (!true)) {
                        v39 := GenS2{F0: 42, F1: 42}
                        v4 -= genFn1()
                        _ = v39
                    } else {
                        var v40 uint8 = 115
                        var v41 uint8 = v17
                        fmt.Println(fi37, v17)
                        delete(v2, 2147483648)
                        var v42 float32 = -2.25
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                    _ = fi38
                }
                fmt.Println(fi37)
            } else {
                v1 += (^(v1))
                v17 += v17
                v1 = -226
            }
            _ = fi37
        }
        for fi43 := 0; fi43 < 19; fi43++ {
            var v44 int = int(v17)
            switch (fi43 - (v44 | (^(fi43)))) {
            case 4294967295:
                v10 = (v10 & v10)
            case 128:
                if (146 < 995) {
                    var v45 string = v3
                    _ = v45
                } else {
                    if true {
                        for fi46 := 0; fi46 < 10; fi46++ {
                            var v47 int32 = (-(v10))
                            _ = fi46
                            _ = v47
                        }
                        var v48 bool = false
                        _ = len(v18)
                        _ = v48
                    } else {
                        if len(v18) > 0 {
                            v18[0] = v44
                        }
                        v49 := GenS4{F0: 61}
                        _ = v49
                    }
                    v50 := GenS4{F0: uint16(v17)}
                    v51 := GenS1{F0: v35, F1: -1, F2: -903}
                    _ = v50
                    _ = v51
                }
            case 1:
                v1 = 42
            default:
                v2[1] = v3
            }
            var v52 int64 = 0
            _ = fi43
            _ = v44
            _ = v52
        }
    }
    for fi53 := 0; fi53 < 6; fi53++ {
        var v54 uint = 337
        _ = fi53
        _ = v54
    }
    v55 := GenS5{F0: 65535, F1: (v1 ^ 42), F2: 0}
    if false {
        if false {
            v56 := GenS5{F0: 492, F1: int64(v17), F2: 4294967295}
            if true {
                if true {
                    var v57 uint32 = 245
                    _ = v57
                } else {
                    var v58 float32 = (1.5 + 3.14159)
                    _ = v58
                }
                if false {
                    switch 32768 {
                    case 4294967295:
                        for fi59 := 0; fi59 < 3; fi59++ {
                            v60 := v2[genFn1()]
                            _ = fi59
                            _ = v60
                        }
                    case -2147483648:
                        v17 += (v17 << ((4294967295) & 63))
                    case 1:
                        var v61 uint64 = 2
                        _ = v61
                    }
                    switch v3 {
                    case "\n\t":
                        var v62 uint64 = 4294967295
                        _ = v62
                    case "hello":
                        if len(v18) > 0 {
                            v18[0] = int(v10)
                        }
                    case "a":
                        switch v3 {
                        case "\n\t":
                            v35 = (v35 & v35)
                        case "hello":
                            var v63 int32 = (^(v10))
                            _ = v63
                        }
                    default:
                        for fi64 := 0; fi64 < 4; fi64++ {
                            var v65 int64 = (v1 - (v1 * (^(v1))))
                            if (!(v3 != ("" + v3))) {
                                var v66 uint32 = (v4 % ((v4) | 1))
                                for fi67 := 0; fi67 < 7; fi67++ {
                                    var v68 string = "a"
                                    _ = fi67
                                    _ = v68
                                }
                                _ = v66
                            }
                            _ = fi64
                            _ = v65
                        }
                    }
                } else {
                    for fi69 := 0; fi69 < 3; fi69++ {
                        var v70 uint = 18446744073709551615
                        _ = fi69
                        _ = v70
                    }
                    if false {
                        var v71 string = v3
                        _ = v71
                    } else {
                        for fi72 := 0; fi72 < 8; fi72++ {
                            if false {
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                                var v73 float32 = 1e10
                                _ = v73
                            } else {
                                if len(v18) > 0 {
                                    v18[0] = 193
                                }
                            }
                            fmt.Println(v35, v35, fi72)
                            fmt.Println(v3)
                            _ = fi72
                        }
                        fmt.Println(v17, v4, v1)
                        var v74 string = "hello"
                        _ = v74
                    }
                }
                var v75 float32 = 1.5
                _ = v75
            } else {
                var v76 int64 = -1
                _ = v76
            }
            for fi77 := 0; fi77 < 12; fi77++ {
                v10 = -2147483648
                var v78 uint32 = (v4 << ((788) & 63))
                v1 = v1
                _ = fi77
                _ = v78
            }
            _ = v56
        }
    }
    var v79 uint = 1
    if (true || true) {
        if (!(!true)) {
            var v80 uint32 = v4
            var v81 float32 = 0.0
            var v82 int32 = v10
            _ = v80
            _ = v81
            _ = v82
        }
        if true {
            var v83 uint16 = 1
            fmt.Println(v35, v1, v3)
            _ = v83
        } else {
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        }
    } else {
        for fi84 := 0; fi84 < 6; fi84++ {
            _ = len(v18)
            _ = fi84
        }
        for fi85 := 0; fi85 < 6; fi85++ {
            switch v3 {
            case "":
                v79 -= (v79 / ((2) | 1))
            case "\n\t":
                var v86 float64 = ((1e10 + 3.14159) / ((float64(0.1) / 1.0) - 0.0))
                _ = v86
            }
            _ = fi85
        }
        fmt.Println(v4, v1)
    }
    fmt.Println(v79, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v55
    _ = v79
}
