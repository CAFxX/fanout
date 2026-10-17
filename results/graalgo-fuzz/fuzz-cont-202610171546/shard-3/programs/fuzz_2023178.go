package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint16; F1 uint; F2 int64 }
type GenS4 struct { F0 uint32; F1 int8; F2 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint8; F2 string }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int16; F1 uint8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int; F1 int8; F2 uint }

func main() {
    var v1 uint8 = 105
    var v2 int = 852
    var v3 uint32 = 1
    v4 := make(map[string]int)
    v5 := make(map[uint32]int)
    var v6 float32 = (float32(-2.25) / 1.5)
    if (false || (false && true)) {
        v7 := GenS1{F0: (32767 < -1), F1: 0}
        v8 := GenS2{F0: -809, F1: 2}
        v6 += v6
        _ = v7
        _ = v8
    }
    v9 := []int{}
    _ = len(v5)
    if true {
        switch "foo bar" {
        case "foo bar":
            switch v2 {
            case 42:
                v10 := GenS3{F0: 255, F1: 65535, F2: int64(v1)}
                _ = v10
            case 867:
                fmt.Println(v1, v2)
            case 127:
                if false {
                    if len(v9) > 0 {
                        v9[0] = -242
                    }
                    for fi11 := 0; fi11 < 17; fi11++ {
                        v12 := GenS4{F0: 202, F1: 1, F2: -2147483648}
                        if false {
                            for fi13 := 0; fi13 < 3; fi13++ {
                                v14 := GenS1{F0: false, F1: int32(v1)}
                                v15 := v5[4294967295]
                                switch fi11 {
                                case 128:
                                    v3 = (v3 / ((v3) | 1))
                                default:
                                    v2 -= (^(v15))
                                }
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                        } else {
                            for fi16 := 0; fi16 < 5; fi16++ {
                                if (42 == 403) {
                                    v17 := GenS3{F0: 1, F1: 0, F2: 9223372036854775807}
                                    var v18 int8 = 25
                                    switch "hello" {
                                    case "a":
                                        if true {
                                            v3 = v3
                                            v2 = v2
                                            v18 = (v18 % ((v18) | 1))
                                        } else {
                                            switch -59 {
                                            case 279:
                                                var v19 int16 = 512
                                                _ = v19
                                            case 2147483647:
                                                var v20 string = ("a" + ("foo bar" + ""))
                                                _ = v20
                                            case -646:
                                                v18 = v18
                                            }
                                            fmt.Println(v2)
                                            v1 -= 129
                                            v5[v3] = fi16
                                        }
                                    default:
                                        v3 = 2147483647
                                    }
                                    _ = v17
                                    _ = v18
                                } else {
                                    var v21 uint8 = 1
                                    var v22 float32 = (v6 / v6)
                                    _ = v21
                                    _ = v22
                                }
                                var v23 int32 = 0
                                if true {
                                    v24 := GenS4{F0: (v3 + 1), F1: 100, F2: (v23 * v23)}
                                    if true {
                                        if (714 == 614) {
                                            var v25 int8 = -1
                                            if ((v6 - float32((0.0 * (float64((v6 + 0.1)) + (1.5 * 0.0))))) > v6) {
                                                v26 := GenS5{F0: true}
                                                _ = v26
                                            } else {
                                                v3 += (v3 / ((v3) | 1))
                                                var v27 uint = 4294967295
                                                v28 := GenS1{F0: (65535 < (^(v27))), F1: (^(v23))}
                                                v6 -= v6
                                                fmt.Println(fi16, v25, v3)
                                                _ = v27
                                                _ = v28
                                            }
                                            _ = v25
                                        } else {
                                            if false {
                                                _ = len(v5)
                                                var v29 int8 = 59
                                                v30 := GenS3{F0: 32767, F1: uint(v3), F2: 0}
                                                _ = v29
                                                _ = v30
                                            } else {
                                                delete(v5, (v3 >> ((65535) & 63)))
                                                if (v3 == 2147483647) {
                                                    if true {
                                                        if (32767 <= 418) {
                                                            fmt.Println(v1)
                                                            fmt.Println(v6)
                                                            var v31 uint16 = 815
                                                            _ = v31
                                                        }
                                                        fmt.Println(v1, v2)
                                                    }
                                                    if (!true) {
                                                        v1 += v1
                                                        fmt.Println(v23)
                                                    } else {
                                                        fmt.Println(v23, fi16, v3)
                                                    }
                                                } else {
                                                    if len(v9) > 0 {
                                                        v9[0] = (fi11 - v2)
                                                    }
                                                }
                                                v23 = v23
                                            }
                                            var v32 uint16 = 741
                                            v33 := GenS6{F0: "a", F1: (v1 << ((2) & 63)), F2: "foo bar"}
                                            fmt.Println(v6)
                                            _ = v32
                                            _ = v33
                                        }
                                    } else {
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                        if false {
                                            switch (((("foo bar" + "") + "\n\t") + ("" + ("\n\t" + ("foo bar" + "foo bar")))) + "a") {
                                            case "a":
                                                v1 = v1
                                            default:
                                                var v34 int16 = 32767
                                                _ = v34
                                            }
                                        }
                                    }
                                    _ = v24
                                } else {
                                    var v35 uint32 = (v3 + v3)
                                    _ = v35
                                }
                                _ = fi16
                                _ = v23
                            }
                        }
                        _ = fi11
                        _ = v12
                    }
                    v1 += 42
                } else {
                    switch ("foo bar" + "foo bar") {
                    case "a":
                        fmt.Println(v6, v2, v1)
                    case "hello":
                        v1 += 255
                    case "foo bar":
                        if false {
                            if (0.1 != 1.5) {
                                if false {
                                    var v36 uint = uint(v1)
                                    var v37 int8 = 0
                                    fmt.Println(v3, v3)
                                    _ = v36
                                    _ = v37
                                }
                                var v38 int32 = 0
                                _ = v38
                            }
                        } else {
                            if (true || (-128 <= 0)) {
                                for fi39 := 0; fi39 < 1; fi39++ {
                                    v3 += v3
                                    for fi40 := 0; fi40 < 7; fi40++ {
                                        var v41 float64 = 0.0
                                        _ = fi40
                                        _ = v41
                                    }
                                    _ = fi39
                                }
                                var v42 uint8 = 128
                                if false {
                                    if len(v9) > 0 {
                                        v9[0] = v2
                                    }
                                }
                                _ = v42
                            }
                        }
                    }
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    var v43 uint32 = v3
                    _ = v43
                }
            default:
                var v44 int = (v2 | v2)
                _ = v44
            }
        case "a":
            var v45 int64 = 1
            _ = v45
        case "\n\t":
            v46 := GenS2{F0: 1, F1: 250}
            _ = v46
        }
        _ = len(v4)
    }
    if (false || false) {
        v47 := GenS7{F0: 32768}
        delete(v4, "hello")
        fmt.Println(v3)
        _ = v47
    } else {
        v6 = 0.1
    }
    for fi48 := 0; fi48 < 15; fi48++ {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        if true {
            switch (-(fi48)) {
            case 127:
                if (!(1e10 <= (0.1 * ((-2.25 + -2.25) + ((0.1 + 1.5) * (1.5 * 1e10)))))) {
                    var v49 int32 = 42
                    v50 := GenS8{F0: 1000, F1: (v1 >> ((18446744073709551615) & 63))}
                    _ = v49
                    _ = v50
                }
            case -1:
                for fi51 := 0; fi51 < 12; fi51++ {
                    v1 += v1
                    if false {
                        v52 := GenS6{F0: "\n\t", F1: v1, F2: ("foo bar" + "foo bar")}
                        _ = v52
                    }
                    _ = fi51
                }
            }
            v53 := GenS4{F0: v3, F1: 1, F2: 0}
            _ = v53
        } else {
            delete(v5, (v3 / ((uint32(v1)) | 1)))
            for fi54 := 0; fi54 < 4; fi54++ {
                v55 := GenS9{F0: 42}
                fmt.Println(fi48)
                v1 += 255
                _ = fi54
                _ = v55
            }
        }
        _ = fi48
    }
    v56 := GenS10{F0: v2, F1: 10, F2: 255}
    var v57 float64 = -2.25
    var v58 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v6
    _ = v9
}
