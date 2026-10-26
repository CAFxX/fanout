package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int8; F2 uint32 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 uint64 }
func genFn0() bool {
    var v7 uint32 = 1
    v7 = 910
    v8 := GenS1{F0: 9223372036854775807}
    _ = v7
    _ = v8
    return true
}
func genFn1(p0 string) uint64 {
    if true {
        p0 = "foo bar"
    } else {
        p0 = (p0 + p0)
        switch p0 {
        case "a":
            fmt.Println(p0)
        default:
            var v9 uint64 = 18446744073709551615
            _ = v9
        }
        v10 := []int{-191, 2147483647, 0, 945}
        _ = v10
    }
    var v11 int = 811
    v12 := GenS1{F0: 1}
    v13 := []int{532}
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 822
}

func main() {
    var v1 int16 = 1000
    v2 := []string{"foo bar"}
    v3 := []int{-9223372036854775808, 32768}
    var v4 uint32 = 4294967295
    var v5 int8 = 100
    v6 := make(map[string]uint32)
    v14 := GenS1{F0: 0}
    var v15 string = ("foo bar" + "a")
    switch ("" + ("" + ("foo bar" + ((("foo bar" + v15) + v15) + ((v15 + "") + v15))))) {
    case "":
        v16 := GenS1{F0: 42}
        _ = v16
    case "hello":
        switch v15 {
        case "foo bar":
            _ = len(v3)
        case "\n\t":
            if true {
                v4 = 76
                if false {
                    switch 4294967295 {
                    case 4294967295:
                        var v17 string = "\n\t"
                        _ = v17
                    case -2147483648:
                        var v18 float64 = 1.5
                        _ = v18
                    default:
                        _ = len(v3)
                    }
                } else {
                    v5 = (v5 * v5)
                }
            }
        default:
            if ((-2.25 - 3.14159) < -2.25) {
                switch v15 {
                case "a":
                    v15 = v15
                }
                var v19 float32 = (1e10 - 0.0)
                if false {
                    switch -25 {
                    case 1:
                        var v20 bool = true
                        _ = v20
                    case 2147483647:
                        v21 := GenS2{F0: 2147483647, F1: -1, F2: 4294967295}
                        _ = v21
                    case -2147483648:
                        switch (((v15 + (v15 + (v15 + v15))) + v15) + v15) {
                        case "\n\t":
                            fmt.Println(v5, v15)
                        }
                    default:
                        var v22 uint = 0
                        _ = v22
                    }
                    fmt.Println(v4, v1)
                    var v23 int32 = -816
                    _ = v23
                }
                _ = v19
            } else {
                v1 += v1
            }
        }
    }
    var v24 string = "foo bar"
    if (128 > 127) {
        v25 := GenS3{F0: (v1 & (^(v1))), F1: (v1 % ((v1) | 1))}
        var v26 float64 = float64(1.5)
        for fi27 := 0; fi27 < 18; fi27++ {
            switch int(v1) {
            case -9223372036854775808:
                v24 = ("hello" + (v24 + ((("" + "foo bar") + "foo bar") + v15)))
            case 128:
                if true {
                    if (!(17 >= (v5 | (v5 << ((861) & 63))))) {
                        for fi28 := 0; fi28 < 9; fi28++ {
                            if true {
                                v1 = (v1 << ((18446744073709551615) & 63))
                                if (2147483647 <= -364) {
                                    var v29 int32 = -2147483648
                                    v6[("a" + v15)] = v4
                                    switch (v24 + "\n\t") {
                                    case "a":
                                        var v30 float64 = ((0.0 / (v26 - (v26 - 3.14159))) + float64(-2.25))
                                        _ = v30
                                    }
                                    _ = v29
                                }
                                v5 = 0
                                v6["a"] = (^(v4))
                                var v31 bool = false
                                _ = v31
                            } else {
                                v5 = v5
                                var v32 uint8 = 1
                                switch ("hello" + v15) {
                                case "a":
                                    if ("\n\t" >= ((v24 + v15) + v15)) {
                                        var v33 uint64 = 18446744073709551615
                                        v34 := GenS4{F0: (v4 * v4), F1: v33}
                                        if (false && (v5 < v5)) {
                                            fmt.Println(v24, fi27)
                                            v35 := GenS4{F0: uint32(v32), F1: (v33 / (((^(v33))) | 1))}
                                            _ = v35
                                        } else {
                                            fmt.Println(v24, v15)
                                        }
                                        _ = v33
                                        _ = v34
                                    } else {
                                        fmt.Println(v26, v24, v4)
                                        _ = len(v3)
                                    }
                                    v24 = v15
                                    var v36 int = fi28
                                    _ = v36
                                default:
                                    v37 := GenS5{F0: 2, F1: v1}
                                    _ = v37
                                }
                                _ = v32
                            }
                            v26 = float64(3.14159)
                            var v38 int64 = 9223372036854775807
                            _ = fi28
                            _ = v38
                        }
                        if (255 <= 237) {
                            v24 = (v24 + (((v24 + v15) + "foo bar") + ("hello" + ("\n\t" + "a"))))
                        } else {
                            switch "" {
                            case "":
                                v26 = (v26 / 1.5)
                            }
                            delete(v6, v15)
                        }
                    } else {
                        fmt.Println(v26, v24, v5)
                    }
                    v39 := GenS6{F0: (v4 + 65535), F1: 42}
                    _ = v39
                } else {
                    v1 = v1
                    var v40 int8 = v5
                    _ = v40
                }
            default:
                var v41 float64 = v26
                _ = v41
            }
            _ = len(v3)
            _ = fi27
        }
        _ = v25
        _ = v26
    }
    v24 = ((v24 + (v24 + v24)) + ("\n\t" + v15))
    var v42 uint64 = genFn1("a")
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (32768 < 0) {
        v5 = (^(v5))
        var v43 float64 = 0.0
        switch "" {
        case "":
            if (!(!false)) {
                var v44 float64 = (float64(1e10) * ((float64(0.0) + -2.25) + v43))
                for fi45 := 0; fi45 < 11; fi45++ {
                    fmt.Println(v24, v1)
                    var v46 int32 = 42
                    fmt.Println(v5)
                    _ = fi45
                    _ = v46
                }
                var v47 uint = 65535
                _ = v44
                _ = v47
            } else {
                v4 -= v4
                var v48 int64 = 42
                _ = v48
            }
        case "a":
            switch 951 {
            case 1:
                v43 -= (v43 + ((v43 * v43) + (float64((1.5 * (1e10 / (0.1 * 1e10)))) + (v43 / v43))))
            case 0:
                v49 := GenS2{F0: 127, F1: v5, F2: 316}
                _ = v49
            default:
                for fi50 := 0; fi50 < 16; fi50++ {
                    for fi51 := 0; fi51 < 18; fi51++ {
                        v52 := GenS2{F0: (fi50 / (((-(fi51))) | 1)), F1: (v5 - 29), F2: 456}
                        fmt.Println(fi51, v42, v43)
                        _ = fi51
                        _ = v52
                    }
                    var v53 string = v24
                    v54 := GenS4{F0: (v4 & v4), F1: genFn1("a")}
                    _ = fi50
                    _ = v53
                    _ = v54
                }
            }
        case "hello":
            var v55 string = ""
            _ = v55
        }
        _ = v43
    }
    switch int(v4) {
    case -9223372036854775808:
        if false {
            fmt.Println(v1, v5)
        } else {
            v2 = append(v2, (v24 + v24))
            var v56 int16 = (v1 + (v1 & v1))
            _ = v56
        }
    case 9223372036854775807:
        var v57 string = v24
        _ = v57
    case 957:
        fmt.Println(v42, v1, v4)
    default:
        if ((v42 >> ((1) & 63)) >= (v42 << ((1) & 63))) {
            switch -9223372036854775808 {
            case -9223372036854775808:
                for fi58 := 0; fi58 < 4; fi58++ {
                    if (int(v5) <= (^(fi58))) {
                        _ = len(v6)
                    } else {
                        v4 -= 1
                        v59 := GenS2{F0: (fi58 | -9223372036854775808), F1: 34, F2: 42}
                        v24 = ("" + v15)
                        _ = v59
                    }
                    switch v24 {
                    case "":
                        if (!((v1 | 32767) > (^(v1)))) {
                            var v60 uint32 = 2147483648
                            fmt.Println(v42, v60, v24)
                            _ = v60
                        }
                    case "a":
                        for fi61 := 0; fi61 < 19; fi61++ {
                            var v62 int8 = v5
                            var v63 uint8 = 1
                            _ = fi61
                            _ = v62
                            _ = v63
                        }
                    default:
                        var v64 int8 = v5
                        _ = v64
                    }
                    v6[(v15 + "a")] = v4
                    _ = fi58
                }
            case 4294967295:
                fmt.Println(v42, v15)
            default:
                for fi65 := 0; fi65 < 2; fi65++ {
                    v66 := GenS7{F0: genFn1("")}
                    switch fi65 {
                    case -1:
                        _ = len(v6)
                    case 1:
                        fmt.Println(fi65)
                    }
                    var v67 uint32 = uint32(v42)
                    _ = fi65
                    _ = v66
                    _ = v67
                }
            }
            fmt.Println(v15)
        } else {
            if len(v2) > 0 {
                v2[0] = ((v24 + v24) + v15)
            }
            v68 := GenS6{F0: (v4 << ((4294967295) & 63)), F1: v4}
            var v69 int8 = (v5 >> ((18446744073709551615) & 63))
            _ = v68
            _ = v69
        }
    }
    v15 = ("a" + (v15 + (v24 + v24)))
    var v70 string = ((v24 + (v24 + (v15 + "\n\t"))) + ((v15 + "a") + ("\n\t" + (("hello" + (v15 + "\n\t")) + v24))))
    if false {
        fmt.Println(v70, v15)
        if ((1e10 - 3.14159) >= 3.14159) {
            var v71 uint32 = (v4 * (v4 ^ (v4 & (v4 / ((v4) | 1)))))
            _ = v71
        }
    } else {
        v70 = v24
    }
    v24 = v15
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v70:", v70)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v70
}
