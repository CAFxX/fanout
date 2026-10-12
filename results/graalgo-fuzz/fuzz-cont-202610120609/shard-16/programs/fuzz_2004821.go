package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int; F2 int64 }

func main() {
    var v1 int = -752
    v2 := make(map[int]uint32)
    var v3 int64 = -9223372036854775808
    var v4 int16 = -629
    var v5 int8 = -1
    v6 := GenS1{F0: 143}
    v7 := []int{}
    v7 = append(v7, int(v3))
    var v8 float64 = 1.5
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS2{F0: (v3 % (((v3 + v3)) | 1)), F1: (("" + "foo bar") + "hello")}
        for fi11 := 0; fi11 < 7; fi11++ {
            if false {
                v12 := GenS1{F0: 32767}
                var v13 uint32 = 0
                _ = v12
                _ = v13
            } else {
                v14 := GenS2{F0: (v3 & (v3 & (v3 - v3))), F1: "hello"}
                if ("foo bar" < ("foo bar" + "a")) {
                    _ = len(v2)
                    fmt.Println(fi9, fi11, v1)
                    if (!false) {
                        v15 := GenS1{F0: 193}
                        v8 -= 0.0
                        var v16 uint = 682
                        _ = v15
                        _ = v16
                    } else {
                        if (0 > 255) {
                            var v17 bool = true
                            v18 := GenS1{F0: 1}
                            switch 1 {
                            case 0:
                                switch 42 {
                                case -9223372036854775808:
                                    if true {
                                        _ = len(v2)
                                    } else {
                                        var v19 int = (fi9 ^ (fi9 - 32768))
                                        _ = v19
                                    }
                                case 32768:
                                    _ = len(v7)
                                default:
                                    v20 := GenS1{F0: 420}
                                    _ = v20
                                }
                            case 2147483647:
                                var v21 float32 = (1e10 + float32(((float64(0.0) - v8) - v8)))
                                _ = v21
                            }
                            _ = v17
                            _ = v18
                        } else {
                            if (v8 >= (((v8 + v8) / float64(0.1)) * float64(-2.25))) {
                                if (-584 != fi9) {
                                    v4 = 42
                                }
                                var v22 bool = (v3 > v3)
                                if v22 {
                                    v5 += (v5 * v5)
                                } else {
                                    if false {
                                        v22 = ((-(v1)) == fi11)
                                        if v22 {
                                            var v23 int16 = v4
                                            fmt.Println(v1, v3, v23)
                                            _ = v23
                                        } else {
                                            var v24 float32 = 1.5
                                            v25 := GenS2{F0: -920, F1: "hello"}
                                            _ = v24
                                            _ = v25
                                        }
                                    }
                                    if v22 {
                                        v4 = (v4 ^ 32767)
                                    }
                                    switch "a" {
                                    case "\n\t":
                                        if v22 {
                                            var v26 int64 = (v3 >> ((2) & 63))
                                            var v27 float64 = v8
                                            v22 = v22
                                            _ = v26
                                            _ = v27
                                        } else {
                                            var v28 bool = ((("a" + "foo bar") + "a") > (("a" + ("hello" + ("foo bar" + "\n\t"))) + "hello"))
                                            _ = v28
                                        }
                                    case "a":
                                        _ = len(v2)
                                    case "hello":
                                        var v29 bool = false
                                        _ = v29
                                    }
                                    switch ("hello" + "foo bar") {
                                    case "foo bar":
                                        if len(v7) > 0 {
                                            v7[0] = v1
                                        }
                                    case "\n\t":
                                        var v30 int64 = 0
                                        _ = v30
                                    case "a":
                                        _ = len(v7)
                                    default:
                                        v4 = -588
                                    }
                                    switch "hello" {
                                    case "a":
                                        if (!v22) {
                                            switch 2147483647 {
                                            case 42:
                                                if len(v7) > 0 {
                                                    v7[0] = 0
                                                }
                                                v4 += (^(v4))
                                                v31 := GenS1{F0: 2}
                                                _ = v31
                                            }
                                        }
                                    }
                                }
                                _ = v22
                            }
                        }
                    }
                }
                fmt.Println(v8, v4, v4)
                _ = v14
            }
            var v32 uint64 = 1
            switch (fi11 * -1) {
            case -843:
                v7 = append(v7, (fi11 - fi11))
            case 127:
                if (!true) {
                    if true {
                        v7 = append(v7, v1)
                        var v33 uint32 = 0
                        switch (fi9 >> ((646) & 63)) {
                        case -554:
                            switch (^(v1)) {
                            case 2147483647:
                                var v34 int64 = -9223372036854775808
                                _ = v34
                            case -1:
                                v8 += (v8 / 0.1)
                            default:
                                if false {
                                    var v35 uint8 = 212
                                    fmt.Println(fi11)
                                    v36 := GenS2{F0: (v3 >> ((4294967295) & 63)), F1: "hello"}
                                    _ = v35
                                    _ = v36
                                }
                            }
                        default:
                            if len(v7) > 0 {
                                v7[0] = fi9
                            }
                        }
                        _ = v33
                    } else {
                        v37 := v2[32767]
                        switch "\n\t" {
                        case "foo bar":
                            v32 = (v32 / ((v32) | 1))
                        case "a":
                            v38 := v2[(fi11 >> ((2) & 63))]
                            _ = v38
                        case "":
                            delete(v2, v1)
                        }
                        _ = v37
                    }
                    if false {
                        fmt.Println(v4, v1)
                    } else {
                        fmt.Println(fi11, v4, fi9)
                        switch ("hello" + "\n\t") {
                        case "foo bar":
                            var v39 int64 = (-(v3))
                            _ = v39
                        case "a":
                            v40 := v2[v1]
                            _ = v40
                        case "\n\t":
                            fmt.Println(fi11)
                        }
                    }
                }
            default:
                var v41 float32 = 1.5
                _ = v41
            }
            _ = fi11
            _ = v32
        }
        var v42 bool = ("a" != "hello")
        _ = fi9
        _ = v10
        _ = v42
    }
    var v43 int = v1
    v44 := GenS2{F0: 928, F1: ("hello" + "")}
    var v45 uint16 = 322
    var v46 int = -758
    var v47 int64 = (v3 | v3)
    v3 += 42
    fmt.Println(v45)
    if ((true && (1 < -2147483648)) && false) {
        for fi48 := 0; fi48 < 8; fi48++ {
            if (-2147483648 > -189) {
                var v49 string = "\n\t"
                var v50 float64 = v8
                for fi51 := 0; fi51 < 14; fi51++ {
                    v2[4294967295] = 0
                    _ = fi51
                }
                _ = v49
                _ = v50
            } else {
                for fi52 := 0; fi52 < 14; fi52++ {
                    v43 = fi52
                    var v53 int16 = (v4 + 32767)
                    v46 += 71
                    if len(v7) > 0 {
                        v7[0] = -9223372036854775808
                    }
                    if (v43 > fi48) {
                        delete(v2, v1)
                        var v54 int64 = (v47 * (v47 % ((v3) | 1)))
                        fmt.Println(v5, v53)
                        _ = v54
                    } else {
                        if false {
                            var v55 bool = true
                            v1 = (^(fi52))
                            if v55 {
                                switch "foo bar" {
                                case "foo bar":
                                    fmt.Println(fi48)
                                case "a":
                                    var v56 int64 = v47
                                    _ = v56
                                case "hello":
                                    var v57 int = -2147483648
                                    _ = v57
                                default:
                                    fmt.Println(v55, v8, v5)
                                }
                                v55 = (false && ((v4 - -1) >= (v4 & v4)))
                                var v58 uint8 = 1
                                _ = v58
                            } else {
                                switch (v43 & (fi52 << ((748) & 63))) {
                                case 42:
                                    v59 := GenS2{F0: v47, F1: ""}
                                    fmt.Println(v46, v55, v5)
                                    v55 = (65535 <= 2147483648)
                                    _ = v59
                                default:
                                    var v60 float32 = (1e10 + float32(1.5))
                                    _ = v60
                                }
                                if v55 {
                                    if v55 {
                                        if (!(598 <= 296)) {
                                            if (0.0 == 0.0) {
                                                v61 := GenS3{F0: (v3 | -9223372036854775808)}
                                                switch (fi52 ^ v46) {
                                                case -9223372036854775808:
                                                    fmt.Println(v5, v5, v46)
                                                case 32767:
                                                    var v62 uint8 = 127
                                                    _ = v62
                                                }
                                                _ = v61
                                            }
                                            v8 += -2.25
                                            v53 -= v4
                                        }
                                    }
                                    v63 := GenS4{F0: (v5 ^ -1), F1: fi48, F2: v47}
                                    var v64 uint16 = (v45 + (v45 / ((v45) | 1)))
                                    _ = v63
                                    _ = v64
                                }
                            }
                            _ = v55
                        } else {
                            var v65 bool = false
                            var v66 float64 = v8
                            _ = v65
                            _ = v66
                        }
                    }
                    _ = len(v7)
                    _ = fi52
                    _ = v53
                }
                for fi67 := 0; fi67 < 16; fi67++ {
                    switch 32767 {
                    case -573:
                        v2[fi67] = 1
                    }
                    delete(v2, (v1 % ((v46) | 1)))
                    v68 := GenS1{F0: v45}
                    _ = fi67
                    _ = v68
                }
                fmt.Println(v43, fi48, v4)
            }
            _ = fi48
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    } else {
        var v69 float32 = 1e10
        v45 = (v45 & 32767)
        _ = v69
    }
    if true {
        v70 := GenS2{F0: v3, F1: ""}
        v45 -= (v45 & 822)
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
