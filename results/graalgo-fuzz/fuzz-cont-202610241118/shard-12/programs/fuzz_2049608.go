package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS5 struct { F0 uint64; F1 uint32; F2 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int16; F1 int16 }
type GenS9 struct { F0 uint64; F1 uint; F2 uint }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]bool)
    v3 := []string{("foo bar" + "\n\t"), "hello", "\n\t", ("hello" + "a")}
    v4 := []int{-869}
    v5 := []string{}
    v6 := make(map[int]int)
    var v7 string = (((("hello" + "foo bar") + "\n\t") + ("a" + "foo bar")) + "")
    var v8 float64 = 0.0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v9 := GenS1{F0: 32767, F1: (v1 >> ((1) & 63))}
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(v1, v8)
        switch (v7 + v7) {
        case "hello":
            v11 := GenS1{F0: 0, F1: v1}
            _ = v11
        }
        _ = fi10
    }
    v2[1] = (!(32768 != 27))
    if true {
        switch -752 {
        case 1:
            v12 := GenS1{F0: 0, F1: 1}
            _ = v12
        case -561:
            v13 := GenS1{F0: 32767, F1: -9223372036854775808}
            _ = v13
        case -689:
            var v14 uint64 = 0
            _ = v14
        default:
            fmt.Println(v7, v7, v1)
        }
    } else {
        v15 := GenS1{F0: 0, F1: v1}
        _ = v15
    }
    if ((false || false) && true) {
        if (false || false) {
            var v16 uint = 1
            fmt.Println(v16, v8)
            v16 = (^(v16))
            _ = v16
        } else {
            switch -2147483648 {
            case 288:
                switch 32768 {
                case -2147483648:
                    var v17 uint64 = 569
                    _ = v17
                default:
                    for fi18 := 0; fi18 < 19; fi18++ {
                        switch v7 {
                        case "\n\t":
                            var v19 int16 = -1
                            _ = v19
                        case "a":
                            if (!((!((902 == 4294967295) && (!false))) && true)) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if (!(v1 >= (v1 << ((736) & 63)))) {
                                    if true {
                                        v1 = v1
                                        if (876 <= -1) {
                                            v1 = (v1 ^ (^(v1)))
                                        } else {
                                            var v20 int = (^(fi18))
                                            v21 := GenS2{F0: 42}
                                            v1 -= v1
                                            _ = v20
                                            _ = v21
                                        }
                                        var v22 uint64 = 0
                                        _ = v22
                                    }
                                    v6[9223372036854775807] = (fi18 & fi18)
                                } else {
                                    for fi23 := 0; fi23 < 10; fi23++ {
                                        if true {
                                            v24 := GenS1{F0: 0, F1: -116}
                                            _ = v24
                                        }
                                        if (42 >= 386) {
                                            var v25 uint8 = 2
                                            v25 = v25
                                            _ = v25
                                        }
                                        var v26 int32 = 1
                                        _ = fi23
                                        _ = v26
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v1 += (v1 >> ((65535) & 63))
                                }
                                v8 -= 3.14159
                            } else {
                                fmt.Println(v7, v7, fi18)
                            }
                        case "foo bar":
                            var v27 int16 = -1
                            _ = v27
                        default:
                            var v28 int32 = 1
                            _ = v28
                        }
                        v7 = "a"
                        v29 := GenS3{F0: (0 <= 62)}
                        _ = fi18
                        _ = v29
                    }
                }
            case 9223372036854775807:
                if (false || true) {
                    _ = len(v3)
                    v30 := GenS4{F0: ((false && false) || (321 != 274)), F1: 255, F2: 1}
                    v1 -= (v1 << ((0) & 63))
                    _ = v30
                }
            case 32768:
                var v31 uint8 = 0
                _ = v31
            }
            var v32 int16 = -32768
            _ = v32
        }
    } else {
        for fi33 := 0; fi33 < 3; fi33++ {
            switch 127 {
            case -125:
                _ = len(v6)
            case -9223372036854775808:
                var v34 int64 = (v1 & 1)
                _ = v34
            default:
                var v35 int32 = 0
                _ = v35
            }
            _ = fi33
        }
        if (true || (420 >= (v1 - v1))) {
            if (0.1 <= v8) {
                switch v7 {
                case "\n\t":
                    if false {
                        var v36 uint8 = 127
                        if ((-243 == -1) || (49 <= 0)) {
                            v2[2147483648] = (1.5 <= (0.0 * float64((3.14159 / -2.25))))
                            switch v7 {
                            case "foo bar":
                                v37 := GenS5{F0: 18446744073709551615, F1: 2, F2: (13 < 0)}
                                _ = v37
                            case "":
                                var v38 int64 = 9223372036854775807
                                _ = v38
                            case "\n\t":
                                for fi39 := 0; fi39 < 13; fi39++ {
                                    var v40 int16 = 42
                                    _ = fi39
                                    _ = v40
                                }
                            default:
                                if false {
                                    for fi41 := 0; fi41 < 8; fi41++ {
                                        fmt.Println(fi41)
                                        v7 = "\n\t"
                                        var v42 float64 = (((float64((((0.1 / 1.5) / -2.25) - 3.14159)) * 0.1) - ((v8 - float64((3.14159 - -2.25))) * -2.25)) / float64(0.1))
                                        _ = fi41
                                        _ = v42
                                    }
                                    v7 = (v7 + ("\n\t" + ""))
                                    var v43 int16 = 293
                                    _ = v43
                                } else {
                                    if true {
                                        if false {
                                            for fi44 := 0; fi44 < 20; fi44++ {
                                                v7 = (v7 + "")
                                                fmt.Println(v8, fi44)
                                                var v45 float64 = (v8 - v8)
                                                _ = fi44
                                                _ = v45
                                            }
                                            var v46 uint32 = 781
                                            _ = v46
                                        } else {
                                            for fi47 := 0; fi47 < 12; fi47++ {
                                                v48 := GenS4{F0: false, F1: (^(v36)), F2: 4294967295}
                                                _ = fi47
                                                _ = v48
                                            }
                                        }
                                    }
                                    switch -928 {
                                    case 128:
                                        v2[4294967295] = (0 > 0)
                                    case 0:
                                        var v49 uint8 = (v36 | v36)
                                        _ = v49
                                    default:
                                        var v50 int32 = -1
                                        _ = v50
                                    }
                                }
                            }
                            fmt.Println(v7)
                        } else {
                            if (!(2147483647 == -1)) {
                                v8 -= 0.0
                                var v51 uint8 = (v36 >> ((942) & 63))
                                _ = v51
                            }
                            _ = len(v5)
                        }
                        var v52 int16 = 1000
                        _ = v36
                        _ = v52
                    } else {
                        fmt.Println(v8, v8)
                    }
                }
                var v53 int64 = (^(v1))
                _ = v53
            } else {
                v54 := GenS6{F0: (((v7 + v7) + v7) + v7)}
                var v55 int32 = 42
                _ = v54
                _ = v55
            }
            var v56 int16 = 42
            _ = v56
        }
        if (-2147483648 > -926) {
            for fi57 := 0; fi57 < 2; fi57++ {
                v1 = v1
                v8 += 1.5
                var v58 int64 = v1
                _ = fi57
                _ = v58
            }
        } else {
            if true {
                var v59 uint8 = 88
                _ = v59
            } else {
                if len(v4) > 0 {
                    v4[0] = 899
                }
            }
            _ = len(v6)
        }
    }
    if (18446744073709551615 >= 917) {
        for fi60 := 0; fi60 < 20; fi60++ {
            fmt.Println(v1)
            _ = fi60
        }
    }
    v1 += v1
    v61 := GenS7{F0: 0}
    v62 := GenS8{F0: 42, F1: 295}
    v63 := GenS2{F0: 0}
    switch 127 {
    case 1:
        v64 := GenS9{F0: 0, F1: 738, F2: 656}
        _ = v64
    case 128:
        v65 := v6[128]
        _ = v65
    case 247:
        v8 += v8
    default:
        fmt.Println(v7, v8, v8)
    }
    if (false || true) {
        var v66 bool = true
        fmt.Println(v7)
        _ = v66
    } else {
        var v67 float64 = (float64(0.1) + v8)
        _ = v67
    }
    var v68 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v68:", v68)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
    _ = v63
    _ = v68
    _ = v7
    _ = v8
    _ = v9
}
