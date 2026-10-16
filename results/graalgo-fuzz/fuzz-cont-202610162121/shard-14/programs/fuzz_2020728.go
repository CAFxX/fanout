package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 string; F1 bool; F2 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[uint32]int)
    var v3 int8 = 58
    var v4 uint8 = 2
    var v5 uint8 = v4
    var v6 uint32 = 2147483648
    var v7 uint8 = v4
    v8 := []int{862, -519}
    if len(v8) > 0 {
        v8[0] = -2147483648
    }
    v7 -= (v4 + v4)
    if true {
        if false {
            switch ("foo bar" + "a") {
            case "":
                if true {
                    for fi9 := 0; fi9 < 2; fi9++ {
                        v10 := GenS1{F0: (v5 >> ((18446744073709551615) & 63))}
                        if (!(true || ("foo bar" != ""))) {
                            v5 = (^(v7))
                            if false {
                                _ = len(v2)
                            }
                            var v11 int8 = (^(v3))
                            _ = v11
                        }
                        _ = fi9
                        _ = v10
                    }
                    v12 := GenS2{F0: v6, F1: 126}
                    var v13 float32 = ((1e10 - 0.0) + 1.5)
                    _ = v12
                    _ = v13
                } else {
                    var v14 float64 = (0.1 / -2.25)
                    _ = v14
                }
            }
            switch -9223372036854775808 {
            case -1:
                var v15 uint16 = 2
                _ = v15
            case 977:
                v6 = 2147483648
            case 1:
                var v16 float32 = 1e10
                _ = v16
            default:
                if false {
                    var v17 uint32 = (^(v6))
                    _ = v17
                } else {
                    v18 := GenS2{F0: 496, F1: int(v1)}
                    v19 := GenS2{F0: (v6 << ((18446744073709551615) & 63)), F1: 1}
                    _ = len(v2)
                    _ = v18
                    _ = v19
                }
            }
        } else {
            v3 += int8(v1)
            for fi20 := 0; fi20 < 17; fi20++ {
                var v21 int64 = 1
                _ = fi20
                _ = v21
            }
        }
        var v22 uint32 = 116
        fmt.Println(v1, v1, v4)
        _ = v22
    } else {
        for fi23 := 0; fi23 < 14; fi23++ {
            for fi24 := 0; fi24 < 17; fi24++ {
                switch "" {
                case "\n\t":
                    if false {
                        switch fi24 {
                        case 42:
                            v1 += v1
                        case 0:
                            var v25 float64 = float64(-2.25)
                            _ = v25
                        }
                        v26 := GenS1{F0: v4}
                        switch "" {
                        case "":
                            fmt.Println(fi24, v3)
                        case "\n\t":
                            fmt.Println(fi23, v7)
                        default:
                            v5 = v4
                        }
                        _ = v26
                    } else {
                        v27 := GenS3{F0: "", F1: false, F2: (-(v1))}
                        _ = v27
                    }
                }
                _ = fi24
            }
            if (!(255 == 4294967295)) {
                v28 := GenS4{F0: 354}
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                switch 128 {
                case 4294967295:
                    _ = len(v8)
                case 9223372036854775807:
                    switch 0 {
                    case -9223372036854775808:
                        switch ("foo bar" + "") {
                        case "a":
                            v6 -= v6
                        default:
                            v3 = (v3 * -53)
                        }
                    case 0:
                        if true {
                            v6 = 4294967295
                        }
                    case -1:
                        var v29 int32 = v1
                        _ = v29
                    default:
                        v4 = 90
                    }
                }
                _ = v28
            }
            _ = fi23
        }
        switch 128 {
        case 32768:
            v4 += v7
        case 42:
            var v30 string = ("a" + ("a" + ""))
            _ = v30
        default:
            if (785 >= 32767) {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                v31 := GenS1{F0: v5}
                _ = v31
            } else {
                var v32 uint = 18446744073709551615
                switch 4294967295 {
                case -227:
                    v33 := GenS4{F0: 1}
                    _ = v33
                }
                v34 := GenS3{F0: (("foo bar" + "foo bar") + "\n\t"), F1: (((float64(0.0) + (3.14159 * -2.25)) + 0.1) == 3.14159), F2: -860}
                _ = v32
                _ = v34
            }
        }
    }
    v3 = v3
    if false {
        v4 = uint8(v6)
        v1 += v1
        var v35 bool = (602 <= 0)
        _ = v35
    }
    delete(v2, 352)
    for fi36 := 0; fi36 < 13; fi36++ {
        v3 += -128
        if (true && false) {
            v8 = append(v8, (fi36 ^ (fi36 - (fi36 / ((127) | 1)))))
            v37 := GenS2{F0: v6, F1: fi36}
            v38 := GenS2{F0: v6, F1: fi36}
            _ = v37
            _ = v38
        }
        _ = fi36
    }
    for fi39 := 0; fi39 < 5; fi39++ {
        fmt.Println(v3)
        for fi40 := 0; fi40 < 14; fi40++ {
            var v41 int8 = (v3 << ((668) & 63))
            if (1 == 42) {
                fmt.Println(v7, v41)
            } else {
                v42 := GenS3{F0: ("" + ""), F1: true, F2: (v1 % (((v1 % (((v1 * v1)) | 1))) | 1))}
                _ = v42
            }
            _ = fi40
            _ = v41
        }
        _ = fi39
    }
    var v43 string = "foo bar"
    if (v43 == v43) {
        for fi44 := 0; fi44 < 15; fi44++ {
            delete(v2, 2147483647)
            v45 := GenS5{F0: 1, F1: (v3 & -48)}
            _ = fi44
            _ = v45
        }
        fmt.Println(v5, v5, v4)
        if (18446744073709551615 > 65535) {
            v46 := GenS6{F0: v5}
            if len(v8) > 0 {
                v8[0] = 32767
            }
            _ = v46
        } else {
            fmt.Println(v1)
        }
    } else {
        v47 := GenS7{F0: 2}
        var v48 uint64 = 0
        _ = v47
        _ = v48
    }
    if false {
        var v49 bool = ((166 >= 18446744073709551615) && false)
        _ = v49
    }
    if true {
        if len(v8) > 0 {
            v8[0] = -888
        }
    } else {
        v5 = v7
        var v50 uint = 1
        _ = v50
    }
    v6 -= 145
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
