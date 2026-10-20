package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint8; F1 string; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 uint32; F2 int8 }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 string; F1 uint32 }
type GenS8 struct { F0 int32; F1 uint64; F2 int64 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int32; F1 uint16; F2 int }
func genFn0(p0 int64, p1 string, p2 int64) int16 {
    var v4 float64 = 0.1
    v5 := []int{}
    var v6 uint32 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 1000
}

func main() {
    var v1 int16 = 650
    var v2 string = ("hello" + "\n\t")
    var v3 float32 = (float32(-2.25) / (1e10 / 0.1))
    fmt.Println(v2, v1, v3)
    v7 := GenS1{F0: 907, F1: 42}
    v8 := []int{0, 4294967295}
    _ = len(v8)
    var v9 string = v2
    v1 += -214
    for fi10 := 0; fi10 < 3; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            v8 = append(v8, fi10)
            _ = fi11
        }
        if false {
            fmt.Println(fi10, v3, v3)
            if (18446744073709551615 <= 2) {
                var v12 int = 2147483647
                v13 := GenS2{F0: 9, F1: "a", F2: v9}
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = v12
                _ = v13
            }
        } else {
            var v14 bool = (!false)
            _ = v14
        }
        _ = fi10
    }
    fmt.Println(v3, v3)
    for fi15 := 0; fi15 < 5; fi15++ {
        fmt.Println(v2, v3)
        switch v2 {
        case "":
            v9 = v9
        default:
            switch 0 {
            case -9223372036854775808:
                var v16 int = fi15
                _ = v16
            case 1:
                fmt.Println(fi15, v3, v3)
            }
        }
        v9 = v2
        _ = fi15
    }
    var v17 uint32 = 548
    switch ("hello" + "hello") {
    case "foo bar":
        if true {
            if len(v8) > 0 {
                v8[0] = 2147483647
            }
        } else {
            fmt.Println(v2)
            var v18 float32 = (0.1 - 3.14159)
            _ = v18
        }
    case "":
        v9 = v2
    case "a":
        if len(v8) > 0 {
            v8[0] = 32768
        }
    }
    v8 = append(v8, 313)
    v19 := GenS1{F0: 1, F1: 255}
    if false {
        if (v3 == ((0.1 / v3) + ((v3 / (v3 + v3)) - ((3.14159 / v3) * v3)))) {
            if false {
                v20 := GenS1{F0: 18446744073709551615, F1: 255}
                v9 = (("" + (v9 + ("foo bar" + (v9 + ("" + (v2 + v2)))))) + "\n\t")
                _ = v20
            }
            v21 := GenS2{F0: 128, F1: v9, F2: "\n\t"}
            _ = v21
        }
        var v22 float64 = float64(1.5)
        _ = v22
    } else {
        if (false && false) {
            var v23 uint64 = 1
            var v24 uint32 = v17
            for fi25 := 0; fi25 < 17; fi25++ {
                if (((42 < -187) || (v17 != (v17 / ((633) | 1)))) || false) {
                    if len(v8) > 0 {
                        v8[0] = (fi25 | int(v1))
                    }
                } else {
                    switch (fi25 * (fi25 + (-(fi25)))) {
                    case 0:
                        v26 := GenS3{F0: uint16(v23)}
                        _ = v26
                    case -1:
                        if (100 < 100) {
                            var v27 float64 = 1e10
                            _ = v27
                        } else {
                            v28 := GenS3{F0: 65535}
                            v29 := GenS3{F0: uint16(v23)}
                            if (-2.25 < (v3 - v3)) {
                                fmt.Println(v9, v2, v9)
                                fmt.Println(v1, v17, v3)
                                switch v2 {
                                case "":
                                    if false {
                                        v3 = ((v3 / ((-2.25 / 0.1) / 0.1)) * v3)
                                        switch (fi25 >> ((1) & 63)) {
                                        case -9223372036854775808:
                                            for fi30 := 0; fi30 < 6; fi30++ {
                                                if ((false && false) && (true && (!false))) {
                                                    switch v2 {
                                                    case "a":
                                                        if true {
                                                            if (!true) {
                                                                fmt.Println(v9, v17, v23)
                                                            }
                                                            switch "hello" {
                                                            case "":
                                                                var v31 string = v2
                                                                _ = v31
                                                            case "\n\t":
                                                                if false {
                                                                    fmt.Println(v23, v17)
                                                                    _ = len(v8)
                                                                    v24 = v24
                                                                }
                                                            case "a":
                                                                v23 = 1
                                                            default:
                                                                var v32 float32 = 0.0
                                                                v33 := GenS1{F0: (v23 & (v23 * 0)), F1: 128}
                                                                _ = v32
                                                                _ = v33
                                                            }
                                                        } else {
                                                            var v34 uint = 255
                                                            v35 := GenS3{F0: 428}
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                    }
                                                    v8 = append(v8, -749)
                                                    v3 = v3
                                                }
                                                var v36 bool = (-983 == -2147483648)
                                                v37 := GenS4{F0: ((v2 + "hello") + "foo bar")}
                                                _ = fi30
                                                _ = v36
                                                _ = v37
                                            }
                                        case 9223372036854775807:
                                            v38 := GenS2{F0: 255, F1: v2, F2: "\n\t"}
                                            _ = v38
                                        case 32768:
                                            v9 = v2
                                        default:
                                            for fi39 := 0; fi39 < 20; fi39++ {
                                                if (!false) {
                                                    var v40 int = 32768
                                                    _ = v40
                                                } else {
                                                    v23 -= 15
                                                    var v41 int32 = 2147483647
                                                    v42 := GenS1{F0: v23, F1: 0}
                                                    if false {
                                                        switch v9 {
                                                        case "a":
                                                            v8 = append(v8, (fi39 - int(v1)))
                                                        case "foo bar":
                                                            if true {
                                                                fmt.Println(fi39, v3, v24)
                                                                var v43 string = v9
                                                                var v44 int64 = -1
                                                                _ = v43
                                                                _ = v44
                                                            }
                                                        case "hello":
                                                            if false {
                                                                var v45 uint64 = v23
                                                                fmt.Println(v1)
                                                                _ = v45
                                                            }
                                                        }
                                                        v1 += -1
                                                        var v46 int32 = v41
                                                        v47 := GenS5{F0: genFn0(0, v9, 0), F1: (v24 ^ 4294967295), F2: 100}
                                                        _ = v46
                                                        _ = v47
                                                    } else {
                                                        v24 = (^(v24))
                                                        v9 = v9
                                                    }
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                    _ = v41
                                                    _ = v42
                                                }
                                                _ = fi39
                                            }
                                        }
                                    }
                                default:
                                    if false {
                                        if (0 >= -3) {
                                            v9 = v9
                                        }
                                    }
                                }
                            } else {
                                for fi48 := 0; fi48 < 3; fi48++ {
                                    v49 := GenS6{F0: (v9 + (("hello" + (v2 + "foo bar")) + v9)), F1: 112}
                                    if true {
                                        switch fi25 {
                                        case 476:
                                            var v50 bool = (!false)
                                            _ = v50
                                        default:
                                            switch fi25 {
                                            case 403:
                                                if false {
                                                    switch -2147483648 {
                                                    case -495:
                                                        switch fi25 {
                                                        case 998:
                                                            var v51 int = (fi25 + (fi25 + (^(fi25))))
                                                            _ = v51
                                                        case -278:
                                                            fmt.Println(fi48, fi48, v23)
                                                            var v52 int = (fi25 >> ((20) & 63))
                                                            if (true || (!false)) {
                                                                if (973 <= 445) {
                                                                    v24 = (v17 & (v24 * 854))
                                                                }
                                                                v17 = 1
                                                            }
                                                            _ = v52
                                                        case 0:
                                                            _ = len(v8)
                                                            v8 = append(v8, fi48)
                                                            v53 := GenS7{F0: v9, F1: v17}
                                                            _ = v53
                                                        }
                                                    case 2147483647:
                                                        v8 = append(v8, (^(fi48)))
                                                    case 32768:
                                                        v1 += v1
                                                    default:
                                                        fmt.Println(v24)
                                                    }
                                                    v24 -= 65535
                                                } else {
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                    v1 = 665
                                                    fmt.Println(v23, fi48)
                                                }
                                            case 127:
                                                v24 -= v24
                                            default:
                                                v23 = (^(v23))
                                            }
                                        }
                                    } else {
                                        fmt.Println(fi25, v1, v24)
                                        var v54 uint32 = v24
                                        if ((v17 >> ((65535) & 63)) == (^(v17))) {
                                            v55 := GenS4{F0: "hello"}
                                            if ((true && false) && true) {
                                                fmt.Println(fi25, v23)
                                                switch ("a" + v2) {
                                                case "a":
                                                    var v56 uint8 = 2
                                                    _ = v56
                                                }
                                            }
                                            switch ((v9 + "\n\t") + v2) {
                                            case "hello":
                                                if (!(42 == 590)) {
                                                    if true {
                                                        v57 := GenS8{F0: 1, F1: (v23 << ((255) & 63)), F2: -798}
                                                        if (162 > 118) {
                                                            if false {
                                                                if ((v24 - v24) >= v17) {
                                                                    switch (fi48 * fi25) {
                                                                    case 9223372036854775807:
                                                                        switch 487 {
                                                                        case -1:
                                                                            var v58 int64 = 9223372036854775807
                                                                            v9 = ""
                                                                            _ = v58
                                                                        default:
                                                                            if len(v8) >= 2 {
                                                                                v8 = v8[1:len(v8)]
                                                                            }
                                                                        }
                                                                    }
                                                                    v59 := GenS9{F0: 2}
                                                                    _ = v59
                                                                } else {
                                                                    var v60 float64 = ((-2.25 - float64(1.5)) / (1e10 / -2.25))
                                                                    _ = v60
                                                                }
                                                            }
                                                        }
                                                        _ = v57
                                                    }
                                                    v54 = (v17 & v24)
                                                } else {
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                }
                                            }
                                            v23 -= (^(v23))
                                            _ = v55
                                        }
                                        _ = v54
                                    }
                                    v61 := GenS10{F0: 0, F1: 65535, F2: (-(fi25))}
                                    _ = fi48
                                    _ = v49
                                    _ = v61
                                }
                                var v62 int8 = 127
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                _ = v62
                            }
                            _ = v28
                            _ = v29
                        }
                    case -859:
                        v3 -= v3
                    default:
                        var v63 int64 = int64(v24)
                        _ = v63
                    }
                    v1 -= (v1 | (v1 - v1))
                }
                _ = fi25
            }
            _ = v23
            _ = v24
        } else {
            var v64 int8 = 127
            v17 -= (v17 + 2147483647)
            _ = v64
        }
        var v65 int32 = -2147483648
        v8 = append(v8, 2147483647)
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
