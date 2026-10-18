package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int8 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint64; F1 int64; F2 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 int32; F2 uint64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int8; F1 bool; F2 uint8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint8; F1 int64 }
type GenS12 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS13 struct { F0 uint16 }
func genFn0(p0 string) bool {
    v7 := GenS1{F0: 85, F1: 1, F2: 42}
    v8 := GenS1{F0: 109, F1: 127, F2: 110}
    var v9 int8 = 42
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return false
}

func main() {
    var v1 uint = 4294967295
    var v2 string = "hello"
    var v3 int8 = 0
    var v4 uint = v1
    var v5 uint32 = 2147483648
    v6 := make(map[uint32]bool)
    v1 -= v1
    var v10 uint = 835
    for fi11 := 0; fi11 < 6; fi11++ {
        if (!(v5 < v5)) {
            var v12 float32 = (3.14159 * 0.0)
            var v13 int16 = -563
            fmt.Println(fi11, v10, v10)
            _ = v12
            _ = v13
        } else {
            v14 := GenS2{F0: 65535, F1: false}
            _ = v14
        }
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 bool = true
            v17 := GenS1{F0: (^(v3)), F1: -305, F2: -1}
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 uint16 = 65535
        _ = fi11
        _ = v18
    }
    var v19 float32 = -2.25
    var v20 int8 = (^(v3))
    v4 += uint(v5)
    for fi21 := 0; fi21 < 16; fi21++ {
        switch "\n\t" {
        case "\n\t":
            for fi22 := 0; fi22 < 8; fi22++ {
                var v23 bool = (int64(v20) <= -1)
                _ = fi22
                _ = v23
            }
        default:
            v20 += 127
        }
        _ = fi21
    }
    v4 -= v10
    if ((v5 ^ v5) == 2) {
        var v24 int16 = -32768
        var v25 uint32 = v5
        _ = v24
        _ = v25
    } else {
        v2 = (("foo bar" + v2) + v2)
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        for fi27 := 0; fi27 < 15; fi27++ {
            fmt.Println(v1)
            switch fi27 {
            case 0:
                v2 = (v2 + (("hello" + v2) + "\n\t"))
            case 300:
                v2 = "foo bar"
            }
            _ = fi27
        }
        switch v2 {
        case "":
            v2 = (v2 + (v2 + ((v2 + v2) + v2)))
        case "foo bar":
            fmt.Println(v20)
        case "a":
            for fi28 := 0; fi28 < 18; fi28++ {
                var v29 int16 = 344
                v19 = ((float32((1.5 - -2.25)) + v19) + v19)
                v4 -= (v1 % ((v4) | 1))
                _ = fi28
                _ = v29
            }
        default:
            v20 -= v20
        }
        v30 := GenS3{F0: v5, F1: 0}
        _ = fi26
        _ = v30
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        v19 += v19
        _ = fi31
    }
    fmt.Println(v2, v3, v20)
    v2 = "foo bar"
    var v32 int32 = -1
    v20 = 42
    v33 := v6[v5]
    if v33 {
        if v33 {
            v34 := GenS4{F0: 1, F1: -392, F2: ("a" + (v2 + (v2 + v2)))}
            if v33 {
                if (42 <= 101) {
                    var v35 uint32 = v5
                    if (float64(v19) == (0.0 * (3.14159 + (((1.5 / (-2.25 / (float64(v19) - (1e10 / -2.25)))) + 1e10) + 0.0)))) {
                        v3 = v20
                        v19 = ((v19 - ((float32(0.1) / (v19 + (0.1 - 3.14159))) + v19)) + 1e10)
                        fmt.Println(v33)
                    }
                    if (("" <= "hello") || (0 <= 240)) {
                        v36 := []int{9223372036854775807}
                        for fi37 := 0; fi37 < 9; fi37++ {
                            for fi38 := 0; fi38 < 12; fi38++ {
                                if v33 {
                                    if false {
                                        if len(v36) >= 2 {
                                            v36 = v36[1:len(v36)]
                                        }
                                        var v39 uint = uint(v35)
                                        if (100 == 100) {
                                            v1 = v4
                                        } else {
                                            if true {
                                                var v40 uint64 = 628
                                                if ((v1 & 1) > v39) {
                                                    switch v2 {
                                                    case "\n\t":
                                                        v41 := GenS4{F0: v40, F1: int64(v40), F2: v2}
                                                        _ = v41
                                                    case "hello":
                                                        v10 += uint(v35)
                                                    }
                                                }
                                                fmt.Println(v35, v4)
                                                v42 := GenS1{F0: int8(v5), F1: 127, F2: -126}
                                                _ = len(v6)
                                                _ = v40
                                                _ = v42
                                            } else {
                                                switch (fi38 ^ fi38) {
                                                case -258:
                                                    if false {
                                                        fmt.Println(v32, fi37, fi38)
                                                        var v43 int64 = -1
                                                        _ = v43
                                                    } else {
                                                        var v44 int8 = (-(v3))
                                                        v20 += -26
                                                        if ((v33 && v33) && (1 <= (v44 * -128))) {
                                                            switch v2 {
                                                            case "\n\t":
                                                                var v45 uint64 = 1
                                                                _ = v45
                                                            case "a":
                                                                switch ((v2 + "foo bar") + (v2 + v2)) {
                                                                case "hello":
                                                                    _ = len(v36)
                                                                case "":
                                                                    var v46 uint64 = 2
                                                                    _ = v46
                                                                case "foo bar":
                                                                    var v47 string = v2
                                                                    _ = v47
                                                                default:
                                                                    v10 -= (^(v10))
                                                                }
                                                            case "":
                                                                v48 := GenS5{F0: 1}
                                                                _ = v48
                                                            default:
                                                                switch (^(fi37)) {
                                                                case -886:
                                                                    v49 := GenS6{F0: fi37, F1: v32, F2: 42}
                                                                    _ = v49
                                                                default:
                                                                    if v33 {
                                                                        var v50 uint64 = 942
                                                                        _ = v50
                                                                    }
                                                                }
                                                            }
                                                            v51 := GenS1{F0: (v20 >> (((v1 - v4)) & 63)), F1: int(v1), F2: -1}
                                                            v39 = (^(v39))
                                                            _ = v51
                                                        } else {
                                                            if false {
                                                                v52 := GenS7{F0: (^(v5))}
                                                                fmt.Println(v5, v4)
                                                                if v33 {
                                                                    v53 := GenS2{F0: 65535, F1: true}
                                                                    if (v33 || false) {
                                                                        delete(v6, v35)
                                                                        var v54 float32 = (v19 - v19)
                                                                        var v55 int16 = -1
                                                                        _ = v54
                                                                        _ = v55
                                                                    }
                                                                    _ = v53
                                                                } else {
                                                                    var v56 uint8 = uint8(v39)
                                                                    var v57 bool = v33
                                                                    v6[920] = (2 == 4294967295)
                                                                    _ = v56
                                                                    _ = v57
                                                                }
                                                                _ = v52
                                                            } else {
                                                                fmt.Println(v5)
                                                                v58 := GenS8{F0: 42}
                                                                v59 := GenS9{F0: (v44 ^ -1), F1: v33, F2: 1}
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                        }
                                                        _ = v44
                                                    }
                                                }
                                                var v60 int32 = (^(v32))
                                                fmt.Println(v35, v60)
                                                _ = v60
                                            }
                                        }
                                        _ = v39
                                    }
                                    var v61 uint64 = 994
                                    v4 -= v4
                                    _ = v61
                                }
                                switch (fi37 / (((fi38 % (((fi38 >> ((v4) & 63))) | 1))) | 1)) {
                                case -259:
                                    v36 = append(v36, fi37)
                                default:
                                    v62 := GenS9{F0: 0, F1: (v32 == v32), F2: 42}
                                    var v63 uint8 = 83
                                    v64 := GenS10{F0: (v2 + (("a" + "a") + (v2 + ("" + ((v2 + v2) + "hello")))))}
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                }
                                _ = fi38
                            }
                            _ = fi37
                        }
                        _ = v36
                    } else {
                        var v65 int16 = 0
                        _ = v65
                    }
                    _ = v35
                }
            } else {
                v33 = v33
            }
            _ = v34
        }
        switch 42 {
        case 4294967295:
            if v33 {
                var v66 uint64 = 42
                var v67 uint8 = 157
                v19 -= float32(-2.25)
                _ = v66
                _ = v67
            } else {
                var v68 int32 = (v32 | (v32 ^ 0))
                _ = v68
            }
        }
        v4 -= v4
    } else {
        v69 := GenS11{F0: 37, F1: 518}
        for fi70 := 0; fi70 < 9; fi70++ {
            var v71 int32 = v32
            _ = fi70
            _ = v71
        }
        fmt.Println(v33, v20)
        _ = v69
    }
    for fi72 := 0; fi72 < 17; fi72++ {
        for fi73 := 0; fi73 < 9; fi73++ {
            delete(v6, v5)
            var v74 int8 = 1
            var v75 bool = false
            _ = fi73
            _ = v74
            _ = v75
        }
        v76 := GenS12{F0: 42, F1: -189, F2: 2}
        for fi77 := 0; fi77 < 3; fi77++ {
            v78 := GenS9{F0: (v20 % (((v3 % ((v3) | 1))) | 1)), F1: ((v19 != (1.5 - v19)) || v33), F2: 128}
            if v33 {
                if (!(1 <= 18446744073709551615)) {
                    delete(v6, (v5 ^ v5))
                    if true {
                        v79 := v6[(^(v5))]
                        if v33 {
                            var v80 uint16 = 2
                            _ = v80
                        } else {
                            var v81 float32 = (v19 * v19)
                            v82 := GenS13{F0: 255}
                            v83 := []int{(fi77 | fi72), -726}
                            _ = v81
                            _ = v82
                            _ = v83
                        }
                        _ = v79
                    }
                }
            } else {
                v1 += (v4 & v1)
                fmt.Println(v4, fi77)
                var v84 int64 = 1
                _ = v84
            }
            _ = fi77
            _ = v78
        }
        _ = fi72
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
