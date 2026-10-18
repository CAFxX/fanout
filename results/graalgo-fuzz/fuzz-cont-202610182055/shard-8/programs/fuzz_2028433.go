package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int16; F1 int64; F2 int32 }
type GenS5 struct { F0 uint; F1 uint; F2 int8 }
type GenS6 struct { F0 int; F1 int32; F2 uint8 }
func genFn0(p0 uint16, p1 uint64) int32 {
    var v6 int = 128
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    var v1 bool = ("\n\t" <= (("foo bar" + "foo bar") + "a"))
    var v2 uint16 = 32767
    v3 := make(map[uint32]int)
    v4 := make(map[string]string)
    var v5 uint8 = 198
    fmt.Println(v5, v5, v2)
    v5 = (^(v5))
    v5 -= v5
    if false {
        v1 = true
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 uint32 = 718
            _ = fi7
            _ = v8
        }
        v5 -= v5
    } else {
        v9 := GenS1{F0: "hello"}
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        v5 = (v5 & (^(v5)))
        if v1 {
            v11 := v3[702]
            _ = v11
        } else {
            v12 := []int{fi10}
            _ = v12
        }
        var v13 string = "foo bar"
        _ = fi10
        _ = v13
    }
    v4[""] = "hello"
    var v14 uint16 = v2
    if (2 > 2) {
        var v15 string = "a"
        if false {
            var v16 float32 = (0.1 * (0.0 * -2.25))
            v17 := GenS1{F0: ("foo bar" + "foo bar")}
            _ = v16
            _ = v17
        }
        var v18 int = -9223372036854775808
        _ = v15
        _ = v18
    } else {
        if v1 {
            switch (("foo bar" + "hello") + ("hello" + "a")) {
            case "hello":
                fmt.Println(v2, v5, v14)
            case "\n\t":
                var v19 string = ((("hello" + "") + "\n\t") + "\n\t")
                _ = v19
            case "a":
                var v20 int64 = -935
                _ = v20
            }
        }
        switch "hello" {
        case "a":
            switch ("a" + "\n\t") {
            case "\n\t":
                v1 = true
            default:
                for fi21 := 0; fi21 < 20; fi21++ {
                    switch fi21 {
                    case 42:
                        var v22 int32 = 589
                        _ = v22
                    case 127:
                        for fi23 := 0; fi23 < 1; fi23++ {
                            fmt.Println(v1, fi23, v14)
                            var v24 uint8 = v5
                            delete(v3, 1)
                            _ = fi23
                            _ = v24
                        }
                    }
                    _ = fi21
                }
            }
        case "foo bar":
            for fi25 := 0; fi25 < 2; fi25++ {
                if (float32((-2.25 / 1.5)) == (-2.25 / -2.25)) {
                    for fi26 := 0; fi26 < 9; fi26++ {
                        var v27 uint64 = 2
                        switch "" {
                        case "\n\t":
                            v2 -= v2
                        case "hello":
                            if (v1 || v1) {
                                switch "" {
                                case "":
                                    fmt.Println(fi25, v5)
                                case "a":
                                    v28 := []int{fi26, fi25, (fi25 + (fi25 << ((4294967295) & 63))), -2147483648}
                                    _ = v28
                                }
                            }
                        default:
                            v29 := GenS2{F0: -9223372036854775808}
                            _ = v29
                        }
                        var v30 bool = ((!v1) && (v1 || ("hello" == "foo bar")))
                        _ = fi26
                        _ = v27
                        _ = v30
                    }
                    v31 := GenS2{F0: 836}
                    _ = v31
                }
                for fi32 := 0; fi32 < 11; fi32++ {
                    v1 = true
                    var v33 int16 = 241
                    v2 -= uint16(v5)
                    v34 := GenS2{F0: 9223372036854775807}
                    var v35 int32 = -293
                    var v36 int = fi25
                    if (v1 || v1) {
                        if ("a" != "foo bar") {
                            fmt.Println(v35, v33)
                            if true {
                                var v37 int64 = 1
                                v38 := GenS1{F0: "a"}
                                v39 := v3[5]
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                            var v40 float64 = 1e10
                            v41 := GenS3{F0: v35, F1: (^(v14))}
                            var v42 float32 = 1e10
                            if v1 {
                                v40 += ((v40 / (v40 - (v40 * float64(v42)))) + v40)
                                var v43 float64 = ((v40 / v40) / v40)
                                _ = v43
                            } else {
                                if (!(true || (4294967295 != uint32(v14)))) {
                                    switch fi25 {
                                    case -1:
                                        switch "hello" {
                                        case "foo bar":
                                            var v44 int = -212
                                            _ = v44
                                        case "":
                                            v45 := GenS4{F0: v33, F1: -749, F2: (-(v35))}
                                            _ = v45
                                        case "\n\t":
                                            var v46 int32 = (v35 * v35)
                                            _ = v46
                                        }
                                    case 127:
                                        fmt.Println(v14, v40, v36)
                                        fmt.Println(v42, v40, v33)
                                    default:
                                        switch "a" {
                                        case "hello":
                                            v14 = v14
                                        case "":
                                            var v47 uint16 = (v2 - v14)
                                            _ = v47
                                        case "foo bar":
                                            delete(v3, 2147483647)
                                        default:
                                            v5 -= v5
                                        }
                                        var v48 uint = 4294967295
                                        v49 := GenS4{F0: v33, F1: 0, F2: (^(v35))}
                                        _ = v48
                                        _ = v49
                                    }
                                    var v50 int32 = v35
                                    v51 := GenS2{F0: 0}
                                    _ = v50
                                    _ = v51
                                }
                                var v52 int32 = 42
                                _ = v52
                            }
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                    }
                    v53 := []int{(-(v36)), fi25, int(v2), (v36 / (((v36 / ((1) | 1))) | 1))}
                    fmt.Println(fi25, v5)
                    _ = fi32
                    _ = v33
                    _ = v34
                    _ = v35
                    _ = v36
                    _ = v53
                }
                _ = fi25
            }
        }
        switch "" {
        case "foo bar":
            var v54 bool = (uint32(v14) <= 2147483647)
            _ = v54
        case "a":
            var v55 string = "\n\t"
            _ = v55
        default:
            fmt.Println(v1, v2)
        }
    }
    switch 32767 {
    case 127:
        if v1 {
            if (1000 > -65) {
                _ = len(v3)
                for fi56 := 0; fi56 < 20; fi56++ {
                    v57 := GenS1{F0: (("" + "hello") + "foo bar")}
                    if v1 {
                        if (95 <= 614) {
                            fmt.Println(v1)
                        } else {
                            v1 = (1 < -1)
                        }
                        v14 = (v2 & 255)
                        var v58 int32 = 0
                        _ = v58
                    }
                    fmt.Println(v2, fi56)
                    _ = fi56
                    _ = v57
                }
                switch "hello" {
                case "hello":
                    var v59 uint32 = 2
                    _ = v59
                }
            }
        }
    case 1:
        v60 := GenS3{F0: -1, F1: v14}
        _ = v60
    }
    v61 := GenS5{F0: 1, F1: 65535, F2: 127}
    v62 := []int{-9223372036854775808, -1, 4294967295}
    var v63 uint8 = (^(v5))
    var v64 int32 = int32(v5)
    var v65 uint64 = uint64(v2)
    for fi66 := 0; fi66 < 16; fi66++ {
        v62 = append(v62, (fi66 | fi66))
        _ = fi66
    }
    for fi67 := 0; fi67 < 9; fi67++ {
        var v68 uint32 = 42
        var v69 uint16 = uint16(v5)
        v70 := GenS4{F0: 60, F1: 180, F2: int32(v2)}
        _ = fi67
        _ = v68
        _ = v69
        _ = v70
    }
    v65 = (v65 / ((18446744073709551615) | 1))
    v71 := GenS6{F0: -2147483648, F1: genFn0(v14, (^(v65))), F2: v63}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v62:", len(v62), v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v65
    _ = v71
}
