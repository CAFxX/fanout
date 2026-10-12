package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 uint32; F2 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint16; F1 int64 }

func main() {
    var v1 string = ("\n\t" + "")
    var v2 float64 = 0.1
    var v3 float64 = (0.0 / 1.5)
    var v4 bool = true
    var v5 uint64 = 0
    switch (v1 + v1) {
    case "foo bar":
        var v6 uint = 65535
        _ = v6
    default:
        v7 := []int{}
        _ = v7
    }
    switch 42 {
    case -308:
        if (121 < 0) {
            v8 := GenS1{F0: ("" + "hello"), F1: v1}
            _ = v8
        }
    case 1:
        v3 = v3
    }
    v9 := GenS1{F0: (v1 + v1), F1: v1}
    var v10 uint64 = (^(v5))
    switch 158 {
    case 127:
        var v11 float32 = -2.25
        _ = v11
    case 1:
        v12 := []int{}
        _ = v12
    default:
        var v13 int32 = 42
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        fmt.Println(v2, v4)
        v1 = v1
        _ = fi14
    }
    var v15 uint8 = 128
    v10 = v10
    switch 9223372036854775807 {
    case 4294967295:
        for fi16 := 0; fi16 < 10; fi16++ {
            v3 = (v3 / v3)
            _ = fi16
        }
    case 916:
        switch -486 {
        case -9223372036854775808:
            fmt.Println(v3)
        case -1:
            var v17 string = v1
            _ = v17
        case 128:
            fmt.Println(v4, v4, v2)
        default:
            fmt.Println(v2)
        }
    }
    v3 += float64(((float32(0.1) - 0.1) - 3.14159))
    if true {
        for fi18 := 0; fi18 < 3; fi18++ {
            v19 := []int{fi18, -152}
            v15 -= v15
            v20 := GenS1{F0: "foo bar", F1: (v1 + ("\n\t" + v1))}
            _ = fi18
            _ = v19
            _ = v20
        }
        v21 := GenS2{F0: (v10 >> ((18446744073709551615) & 63)), F1: 42}
        switch -1 {
        case 1:
            if (127 <= 9223372036854775807) {
                v22 := GenS2{F0: (v10 >> ((255) & 63)), F1: 1000}
                v23 := []int{-9223372036854775808, -9223372036854775808, -760}
                fmt.Println(v10)
                _ = v22
                _ = v23
            }
        case -580:
            var v24 string = v1
            _ = v24
        default:
            v4 = (v4 && (881 > 9223372036854775807))
        }
        _ = v21
    }
    if true {
        var v25 int8 = 95
        for fi26 := 0; fi26 < 19; fi26++ {
            switch fi26 {
            case -71:
                if v4 {
                    v3 -= float64(((1e10 + -2.25) - 3.14159))
                } else {
                    switch v1 {
                    case "hello":
                        v1 = v1
                    }
                    v27 := GenS1{F0: "foo bar", F1: v1}
                    _ = v27
                }
            case 32767:
                v28 := []int{fi26, (-(fi26)), (^(fi26))}
                _ = v28
            case 4294967295:
                switch 1 {
                case 127:
                    v15 += v15
                case -882:
                    v29 := []int{fi26, (fi26 | 32768), fi26}
                    _ = v29
                }
            default:
                v30 := GenS1{F0: "a", F1: "a"}
                _ = v30
            }
            v31 := []int{(fi26 >> ((4294967295) & 63)), (fi26 >> ((42) & 63)), (fi26 & 4294967295)}
            if (!v4) {
                v15 = v15
                switch "foo bar" {
                case "":
                    var v32 int32 = -1
                    _ = v32
                case "foo bar":
                    v15 -= (v15 ^ v15)
                }
                if false {
                    fmt.Println(v25, v10, v4)
                    if (!(42 != 0)) {
                        switch v1 {
                        case "\n\t":
                            v4 = v4
                        case "hello":
                            v25 = v25
                        default:
                            v33 := GenS3{F0: (^(v5))}
                            _ = v33
                        }
                    } else {
                        var v34 int32 = -249
                        _ = len(v31)
                        _ = v34
                    }
                    switch v1 {
                    case "foo bar":
                        var v35 uint16 = 42
                        _ = v35
                    default:
                        v2 += ((float64(-2.25) + v2) - (-2.25 * float64(0.0)))
                    }
                }
            } else {
                v2 = (v3 * v2)
            }
            _ = fi26
            _ = v31
        }
        _ = v25
    } else {
        var v36 float32 = 1e10
        switch ("" + ((v1 + ("hello" + ("a" + (v1 + v1)))) + v1)) {
        case "\n\t":
            var v37 float32 = float32(float64((v36 - ((v36 + (1e10 - (v36 + (v36 + float32((1e10 - v3)))))) / (v36 / ((v36 - v36) - float32(v2)))))))
            _ = v37
        case "hello":
            for fi38 := 0; fi38 < 9; fi38++ {
                var v39 uint32 = 700
                _ = fi38
                _ = v39
            }
        case "":
            v36 -= v36
        }
        fmt.Println(v3, v5, v2)
        _ = v36
    }
    for fi40 := 0; fi40 < 12; fi40++ {
        v5 -= v10
        if v4 {
            v41 := []int{128}
            _ = v41
        } else {
            if v4 {
                for fi42 := 0; fi42 < 1; fi42++ {
                    if false {
                        switch v1 {
                        case "\n\t":
                            fmt.Println(v3, v2, v5)
                        default:
                            v1 = v1
                        }
                        if (1.5 < (((0.0 * 1.5) + (0.0 / 1.5)) - 1e10)) {
                            if (32767 < 0) {
                                switch "\n\t" {
                                case "a":
                                    if v4 {
                                        if false {
                                            v43 := []int{(fi42 ^ (fi40 >> ((66) & 63))), fi42}
                                            _ = v43
                                        }
                                    } else {
                                        v1 = "\n\t"
                                    }
                                case "foo bar":
                                    v44 := []int{fi40, 0, (-(fi42))}
                                    _ = v44
                                case "":
                                    if ((42 < 0) && ("a" == ("" + ("hello" + v1)))) {
                                        var v45 int16 = -1
                                        v46 := GenS1{F0: "\n\t", F1: v1}
                                        _ = v45
                                        _ = v46
                                    }
                                }
                                fmt.Println(v4, fi42, v1)
                            } else {
                                v47 := []int{(fi40 - (fi40 ^ fi40)), fi42, (^(fi40)), (fi40 ^ -1)}
                                var v48 uint64 = v10
                                _ = v47
                                _ = v48
                            }
                            if v4 {
                                v2 += v3
                                v49 := GenS4{F0: 65535}
                                switch ("foo bar" + (((v1 + v1) + "a") + (v1 + ((("hello" + v1) + v1) + v1)))) {
                                case "hello":
                                    v50 := GenS1{F0: v1, F1: v1}
                                    _ = v50
                                case "a":
                                    v51 := GenS5{F0: 0, F1: 4294967295, F2: 42}
                                    _ = v51
                                case "foo bar":
                                    v52 := GenS6{F0: -2147483648}
                                    _ = v52
                                }
                                _ = v49
                            } else {
                                v53 := GenS2{F0: uint64(v15), F1: -32768}
                                v54 := []int{(fi42 - fi42)}
                                v10 = 349
                                _ = v53
                                _ = v54
                            }
                        }
                        v3 += (v2 / v2)
                    } else {
                        var v55 float32 = (3.14159 - -2.25)
                        _ = v55
                    }
                    _ = fi42
                }
            } else {
                v56 := GenS7{F0: 405, F1: 0}
                for fi57 := 0; fi57 < 14; fi57++ {
                    switch "a" {
                    case "\n\t":
                        var v58 uint16 = 1
                        _ = v58
                    case "a":
                        var v59 int = (fi57 / ((0) | 1))
                        _ = v59
                    }
                    v1 = ("\n\t" + v1)
                    _ = fi57
                }
                _ = v56
            }
        }
        _ = fi40
    }
    var v60 uint32 = 2
    if false {
        v61 := GenS7{F0: 0, F1: 9223372036854775807}
        fmt.Println(v60)
        _ = v61
    } else {
        if true {
            fmt.Println(v60)
            v60 += v60
            v2 -= v3
        } else {
            var v62 float64 = float64(float32((-2.25 - v3)))
            _ = v62
        }
    }
    v63 := []int{}
    v64 := GenS2{F0: (v10 | (^(v5))), F1: 1000}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v63:", len(v63), v63)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v63
    _ = v64
    _ = v9
}
