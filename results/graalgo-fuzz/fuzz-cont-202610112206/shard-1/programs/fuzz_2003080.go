package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint16; F1 bool; F2 int64 }
type GenS4 struct { F0 int8; F1 uint32; F2 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string }

func main() {
    var v1 uint32 = 462
    v2 := []int{int(v1), 0, int(v1)}
    var v3 string = ""
    var v4 uint = 255
    v5 := []uint32{2, v1, (^(v1))}
    var v6 float64 = 0.1
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint8 = 127
        var v9 int64 = 1
        var v10 uint32 = (v1 >> ((v4) & 63))
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    fmt.Println(v6)
    var v11 uint8 = 14
    if len(v5) > 0 {
        v5[0] = (v1 >> (((v4 & 42)) & 63))
    }
    if (false && true) {
        switch 32767 {
        case -449:
            switch 32768 {
            case 127:
                if ((0.1 - float32(v6)) == (0.0 - ((-2.25 * 3.14159) + 1e10))) {
                    v4 = (^(v4))
                    v11 = (v11 * 128)
                    switch ("hello" + "foo bar") {
                    case "foo bar":
                        if false {
                            switch 1 {
                            case 2147483647:
                                v12 := GenS1{F0: 1}
                                _ = v12
                            }
                            _ = len(v2)
                            v4 = v4
                        } else {
                            fmt.Println(v11)
                            for fi13 := 0; fi13 < 15; fi13++ {
                                fmt.Println(v3)
                                _ = fi13
                            }
                        }
                    }
                }
            case 774:
                if ((^(v4)) == 1) {
                    v14 := GenS2{F0: (v11 + (^(v11))), F1: 18}
                    if (true && false) {
                        var v15 int = 1
                        _ = v15
                    }
                    v6 -= (((v6 - ((v6 * float64(1e10)) - (v6 / ((v6 / ((3.14159 * (v6 * ((v6 / v6) - -2.25))) - 0.0)) / (v6 + (v6 / ((float64((0.1 * 0.1)) + v6) + v6))))))) + v6) + (v6 / (1.5 * float64(3.14159))))
                    _ = v14
                } else {
                    var v16 bool = true
                    var v17 uint64 = 916
                    if v16 {
                        if (v16 || v16) {
                            v1 += v1
                        }
                        if true {
                            v16 = (v1 >= v1)
                        }
                    }
                    _ = v16
                    _ = v17
                }
            default:
                fmt.Println(v6, v6)
            }
        default:
            if false {
                var v18 uint16 = 285
                v5 = append(v5, (v1 + uint32(v4)))
                _ = v18
            }
        }
        switch (((("a" + (v3 + v3)) + "a") + (v3 + "hello")) + v3) {
        case "foo bar":
            if (true || false) {
                if true {
                    var v19 int = 126
                    _ = v19
                }
            } else {
                if (!(100 < 48)) {
                    v20 := GenS2{F0: v11, F1: 1}
                    _ = v20
                } else {
                    var v21 float32 = 0.0
                    _ = v21
                }
                var v22 float32 = (0.1 + ((0.1 * (3.14159 - 1e10)) - 0.0))
                _ = v22
            }
        case "":
            var v23 string = ("a" + v3)
            _ = v23
        }
        if ((^(v11)) == (v11 & v11)) {
            if true {
                if (3.14159 > 0.1) {
                    fmt.Println(v4, v11)
                    v24 := GenS1{F0: 926}
                    _ = v24
                } else {
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                    var v25 int64 = 1
                    _ = v25
                }
            }
            var v26 int64 = 0
            var v27 string = v3
            _ = v26
            _ = v27
        } else {
            if true {
                v28 := GenS1{F0: 0}
                _ = v28
            } else {
                if (false || (false && true)) {
                    if false {
                        for fi29 := 0; fi29 < 19; fi29++ {
                            for fi30 := 0; fi30 < 14; fi30++ {
                                switch v3 {
                                case "\n\t":
                                    v1 -= 2
                                default:
                                    fmt.Println(v4)
                                }
                                v4 = uint(v11)
                                _ = fi30
                            }
                            fmt.Println(v3, v4, v6)
                            if true {
                                switch (fi29 & fi29) {
                                case 0:
                                    var v31 float32 = 1.5
                                    _ = v31
                                default:
                                    v32 := GenS2{F0: v11, F1: -128}
                                    _ = v32
                                }
                            } else {
                                var v33 uint16 = 32767
                                _ = v33
                            }
                            _ = fi29
                        }
                        v34 := GenS1{F0: 9223372036854775807}
                        var v35 int32 = 42
                        _ = v34
                        _ = v35
                    } else {
                        v6 += (v6 + ((v6 - v6) / (0.1 * 0.1)))
                        v36 := GenS3{F0: 0, F1: (!((((false || (!(-32768 > 176))) || (!true)) || true) && true)), F2: -1}
                        _ = v36
                    }
                    if true {
                        for fi37 := 0; fi37 < 11; fi37++ {
                            fmt.Println(v4, v3)
                            v38 := GenS4{F0: 100, F1: v1, F2: -1}
                            v4 = 681
                            _ = fi37
                            _ = v38
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v39 := GenS5{F0: (v4 >> ((65535) & 63))}
                        _ = v39
                    }
                    v40 := GenS5{F0: (v4 + v4)}
                    _ = v40
                }
                _ = len(v2)
                v41 := GenS6{F0: (v4 > v4)}
                _ = v41
            }
            v6 -= 3.14159
            for fi42 := 0; fi42 < 9; fi42++ {
                v3 = v3
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v11 += 127
                _ = fi42
            }
        }
    }
    var v43 uint32 = (^(v1))
    v43 -= v43
    var v44 float64 = v6
    switch -39 {
    case -304:
        v11 = (v11 | v11)
    case -466:
        v5 = append(v5, (v43 % (((^(v1))) | 1)))
    default:
        if true {
            v3 = v3
            for fi45 := 0; fi45 < 20; fi45++ {
                var v46 uint16 = 213
                var v47 uint = v4
                _ = fi45
                _ = v46
                _ = v47
            }
            var v48 uint8 = (v11 << ((65535) & 63))
            _ = v48
        } else {
            v49 := GenS2{F0: 2, F1: 8}
            _ = v49
        }
    }
    var v50 float64 = (0.1 + 3.14159)
    if true {
        var v51 string = ""
        switch v3 {
        case "":
            fmt.Println(v3)
        case "hello":
            fmt.Println(v1, v6, v44)
        default:
            v44 = v6
        }
        _ = v51
    }
    if false {
        if (!false) {
            v2 = append(v2, -167)
            v11 += v11
            var v52 uint = (^(v4))
            _ = v52
        } else {
            v5 = append(v5, (v43 * (v43 / ((v43) | 1))))
            v50 += (v50 + (1e10 * 1.5))
        }
        if false {
            switch 1 {
            case 2147483647:
                v11 = 0
            }
            if (!false) {
                for fi53 := 0; fi53 < 14; fi53++ {
                    v54 := GenS3{F0: 65535, F1: false, F2: 0}
                    var v55 int = 0
                    var v56 int64 = 9223372036854775807
                    _ = fi53
                    _ = v54
                    _ = v55
                    _ = v56
                }
            }
            v57 := GenS7{F0: v3}
            _ = v57
        }
    } else {
        v44 = float64((1.5 * (0.0 - 0.1)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v50
    _ = v6
}
