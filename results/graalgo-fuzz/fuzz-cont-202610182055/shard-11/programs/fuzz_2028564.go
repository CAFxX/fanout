package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 int; F1 uint16; F2 uint32 }
type GenS6 struct { F0 uint; F1 bool; F2 int16 }
func genFn0(p0 uint32, p1 bool, p2 int16) int16 {
    var v6 int = 0
    v7 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p2
}

func main() {
    var v1 string = ("" + "foo bar")
    var v2 int32 = 0
    v3 := []int{2147483647, 328, 1}
    v4 := []int{-9223372036854775808, 562, 32767}
    var v5 int64 = -321
    var v8 bool = (v2 == v2)
    _ = len(v3)
    if (-2.25 >= 3.14159) {
        switch 815 {
        case -1:
            for fi9 := 0; fi9 < 19; fi9++ {
                v2 += v2
                v10 := GenS1{F0: 1}
                _ = fi9
                _ = v10
            }
        case 2147483647:
            var v11 int64 = (^(v5))
            _ = v11
        case 128:
            v1 = ("a" + "foo bar")
        default:
            var v12 int8 = -1
            _ = v12
        }
    }
    if v8 {
        fmt.Println(v8, v5, v1)
    } else {
        var v13 uint16 = 554
        v2 -= (v2 | v2)
        if (float32(((-2.25 / 1e10) / 1e10)) > 1.5) {
            if len(v4) > 0 {
                v4[0] = 128
            }
            if true {
                var v14 uint16 = (v13 | 32767)
                var v15 bool = ((!(3.14159 != (float64(((-2.25 * (0.1 * 1.5)) / 3.14159)) / (float64((-2.25 + 1.5)) / 0.1)))) && v8)
                _ = v14
                _ = v15
            } else {
                for fi16 := 0; fi16 < 4; fi16++ {
                    v1 = v1
                    v8 = (53 < 1)
                    for fi17 := 0; fi17 < 5; fi17++ {
                        fmt.Println(v1, v5)
                        var v18 int32 = (v2 | v2)
                        _ = fi17
                        _ = v18
                    }
                    _ = fi16
                }
                for fi19 := 0; fi19 < 4; fi19++ {
                    switch (^(fi19)) {
                    case -2147483648:
                        var v20 int8 = -1
                        _ = v20
                    }
                    v1 = v1
                    if (2147483647 < 42) {
                        for fi21 := 0; fi21 < 5; fi21++ {
                            if len(v4) > 0 {
                                v4[0] = 127
                            }
                            switch 42 {
                            case 174:
                                var v22 int16 = 42
                                _ = v22
                            case -497:
                                var v23 bool = (42 <= (-(v2)))
                                _ = v23
                            }
                            v13 -= 982
                            _ = fi21
                        }
                        if v8 {
                            v8 = (int8(v5) != 0)
                            var v24 int8 = -128
                            _ = v24
                        }
                    } else {
                        var v25 int = (fi19 >> ((42) & 63))
                        v8 = true
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v25
                    }
                    _ = fi19
                }
                v1 = ""
            }
        } else {
            if (4294967295 != 0) {
                var v26 int = 128
                if v8 {
                    v1 = v1
                } else {
                    if true {
                        v26 = v26
                    }
                }
                _ = v26
            }
            for fi27 := 0; fi27 < 10; fi27++ {
                if (42 < 255) {
                    v28 := GenS1{F0: -930}
                    for fi29 := 0; fi29 < 9; fi29++ {
                        v30 := GenS1{F0: 1000}
                        if (((v5 > v5) || true) || true) {
                            var v31 int = (fi29 * fi29)
                            switch v1 {
                            case "foo bar":
                                v5 = (v5 | 123)
                            case "":
                                if (!((2 <= 2) || v8)) {
                                    var v32 uint32 = 42
                                    _ = v32
                                }
                            case "\n\t":
                                if (v2 <= (^(v2))) {
                                    v33 := GenS2{F0: 42, F1: -1}
                                    _ = v33
                                } else {
                                    switch ("foo bar" + ((v1 + "") + v1)) {
                                    case "hello":
                                        switch ((v1 + (v1 + (v1 + v1))) + v1) {
                                        case "":
                                            fmt.Println(fi27, v13)
                                        }
                                    case "":
                                        v2 += int32(v13)
                                    }
                                }
                                switch v31 {
                                case 32767:
                                    fmt.Println(fi29, v2, v2)
                                    if (((fi29 ^ fi27) != v31) || (3.14159 < 1e10)) {
                                        if true {
                                            fmt.Println(v1, fi29, v2)
                                            v2 = v2
                                        }
                                        switch (fi27 ^ int(v13)) {
                                        case 2147483647:
                                            var v34 int64 = v5
                                            _ = v34
                                        case 42:
                                            var v35 uint64 = 0
                                            _ = v35
                                        case 128:
                                            if v8 {
                                                if len(v4) > 0 {
                                                    v4[0] = fi29
                                                }
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                if (v8 || true) {
                                                    if v8 {
                                                        fmt.Println(fi29, fi29)
                                                        var v36 int64 = v5
                                                        _ = v36
                                                    } else {
                                                        var v37 bool = (1e10 < 0.0)
                                                        v5 = -68
                                                        v38 := GenS2{F0: 869, F1: 58}
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                    _ = len(v4)
                                                }
                                                v4 = append(v4, 42)
                                            } else {
                                                v8 = v8
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v39 int8 = 0
                                                v31 -= v31
                                                v2 -= v2
                                                _ = v39
                                            }
                                        default:
                                            v2 = 42
                                        }
                                    } else {
                                        v1 = (v1 + v1)
                                        var v40 uint16 = v13
                                        _ = v40
                                    }
                                }
                                v1 = "a"
                                var v41 uint = 1
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v41
                            default:
                                switch v1 {
                                case "\n\t":
                                    fmt.Println(v1, v2, v8)
                                case "a":
                                    var v42 uint8 = 2
                                    _ = v42
                                default:
                                    var v43 uint = 42
                                    _ = v43
                                }
                            }
                            fmt.Println(fi27, v5)
                            if v8 {
                                switch (fi27 ^ (^(v31))) {
                                case 529:
                                    v2 = v2
                                    if (v8 || (!false)) {
                                        v2 = (-(v2))
                                        v13 = (^(v13))
                                    }
                                }
                            } else {
                                fmt.Println(v5, v1)
                                v2 -= -1
                                var v44 int8 = 1
                                _ = v44
                            }
                            _ = v31
                        }
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                } else {
                    v2 += (v2 + (v2 ^ -1))
                    switch v1 {
                    case "":
                        var v45 int8 = -128
                        _ = v45
                    case "a":
                        var v46 int64 = 9223372036854775807
                        _ = v46
                    }
                }
                switch (("" + "hello") + "foo bar") {
                case "hello":
                    v8 = v8
                default:
                    v47 := GenS3{F0: v13}
                    _ = v47
                }
                _ = fi27
            }
        }
        _ = v13
    }
    if true {
        switch 9223372036854775807 {
        case 2147483647:
            v2 = v2
        case 254:
            v8 = v8
        case 865:
            v48 := GenS2{F0: 129, F1: -30}
            _ = v48
        default:
            v49 := GenS4{F0: 2, F1: 1}
            _ = v49
        }
    } else {
        v50 := GenS5{F0: 127, F1: 58, F2: 2147483647}
        var v51 uint = 661
        _ = v50
        _ = v51
    }
    if v8 {
        if v8 {
            var v52 uint32 = 42
            _ = v52
        }
        v53 := GenS4{F0: 33, F1: 235}
        _ = v53
    }
    for fi54 := 0; fi54 < 2; fi54++ {
        v5 -= v5
        _ = fi54
    }
    var v55 int64 = -9223372036854775808
    v56 := GenS3{F0: 0}
    v57 := GenS6{F0: 0, F1: (255 > 42), F2: -908}
    v5 = v5
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v8
}
