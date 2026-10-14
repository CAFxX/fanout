package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint64; F2 int }
type GenS5 struct { F0 int16; F1 uint }
func genFn0(p0 int8) uint8 {
    var v6 string = ("\n\t" + (("" + ("" + "")) + ""))
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            if (fi7 == (fi8 + (fi7 | (fi7 + (-(fi8)))))) {
                fmt.Println(p0)
                var v9 float64 = 0.0
                v6 = (("foo bar" + ("foo bar" + ((v6 + v6) + (v6 + v6)))) + "\n\t")
                _ = v9
            } else {
                var v10 int16 = -32768
                _ = v10
            }
            fmt.Println(v6, fi8, v6)
            var v11 string = ""
            _ = fi8
            _ = v11
        }
        if ((1 > 2) && (0 > 1000)) {
            switch (fi7 % ((-149) | 1)) {
            case 2147483647:
                for fi12 := 0; fi12 < 14; fi12++ {
                    if (true && ((42 <= 1) || (!true))) {
                        v6 = v6
                        var v13 int16 = 1
                        _ = v13
                    } else {
                        var v14 int32 = 2147483647
                        v15 := GenS1{F0: v6, F1: -743}
                        v6 = v6
                        _ = v14
                        _ = v15
                    }
                    _ = fi12
                }
            default:
                var v16 bool = false
                _ = v16
            }
            if false {
                v17 := GenS1{F0: v6, F1: 9223372036854775807}
                _ = v17
            } else {
                var v18 int32 = 1
                v19 := []int{fi7, (fi7 - (fi7 << ((42) & 63)))}
                for fi20 := 0; fi20 < 3; fi20++ {
                    v21 := GenS1{F0: (v6 + v6), F1: 42}
                    _ = fi20
                    _ = v21
                }
                _ = v18
                _ = v19
            }
        }
        _ = fi7
    }
    switch "" {
    case "\n\t":
        fmt.Println(p0, v6)
    case "hello":
        v22 := GenS1{F0: ("\n\t" + (((v6 + "") + "hello") + ((v6 + "\n\t") + v6))), F1: 1}
        _ = v22
    case "":
        if true {
            v23 := []int{-9223372036854775808}
            _ = v23
        } else {
            for fi24 := 0; fi24 < 9; fi24++ {
                p0 = (-(p0))
                switch fi24 {
                case -743:
                    v25 := GenS2{F0: 2, F1: 2}
                    _ = v25
                }
                var v26 uint16 = 255
                _ = fi24
                _ = v26
            }
        }
    }
    v6 = ("\n\t" + v6)
    _ = p0
    _ = v6
    return 113
}

func main() {
    var v1 uint32 = 2
    var v2 bool = (false || false)
    var v3 string = "a"
    var v4 float32 = float32(-2.25)
    var v5 uint8 = 128
    if true {
        var v27 int32 = -382
        _ = v27
    } else {
        v28 := GenS1{F0: (v3 + "\n\t"), F1: -1}
        fmt.Println(v5, v4, v4)
        v29 := []int{}
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        switch (fi30 - int(v5)) {
        case 32768:
            v31 := GenS3{F0: v5, F1: -1}
            _ = v31
        case 32767:
            v4 += 0.0
        default:
            for fi32 := 0; fi32 < 1; fi32++ {
                v33 := GenS1{F0: (v3 + v3), F1: 1}
                var v34 int64 = 193
                _ = fi32
                _ = v33
                _ = v34
            }
        }
        _ = fi30
    }
    v1 = v1
    for fi35 := 0; fi35 < 18; fi35++ {
        switch ("" + v3) {
        case "foo bar":
            v36 := GenS3{F0: genFn0(-1), F1: -69}
            _ = v36
        }
        v5 = (v5 ^ (v5 & 0))
        _ = fi35
    }
    fmt.Println(v3, v4, v2)
    v37 := GenS1{F0: "hello", F1: -545}
    v2 = v2
    v5 -= v5
    v38 := []int{}
    var v39 int = 1
    if true {
        switch (v3 + (("hello" + (v3 + ("a" + ("foo bar" + "a")))) + "")) {
        case "a":
            for fi40 := 0; fi40 < 10; fi40++ {
                v5 = (v5 / ((v5) | 1))
                _ = fi40
            }
        case "":
            for fi41 := 0; fi41 < 17; fi41++ {
                v39 += fi41
                switch v3 {
                case "\n\t":
                    if (true && ((v5 - 1) < (v5 - (^(v5))))) {
                        if (2 <= 2) {
                            if (-2.25 <= float64(-2.25)) {
                                var v42 int32 = 1
                                _ = v42
                            }
                            var v43 int8 = 42
                            if ((((v43 ^ (v43 / (((^(v43))) | 1))) >= v43) || ((!(255 < 4294967295)) || v2)) && (true && false)) {
                                fmt.Println(v5, v39)
                                v44 := GenS1{F0: "hello", F1: 0}
                                _ = v44
                            } else {
                                fmt.Println(fi41, v1)
                            }
                            _ = v43
                        }
                    }
                case "":
                    switch "a" {
                    case "\n\t":
                        switch v3 {
                        case "hello":
                            v45 := GenS2{F0: 127, F1: (v1 ^ (v1 + v1))}
                            _ = v45
                        case "foo bar":
                            v4 = 1.5
                        case "\n\t":
                            var v46 int = 9223372036854775807
                            _ = v46
                        default:
                            v3 = v3
                        }
                    case "hello":
                        if v2 {
                            if len(v38) > 0 {
                                v38[0] = fi41
                            }
                        } else {
                            fmt.Println(fi41)
                            v47 := GenS1{F0: "foo bar", F1: -9223372036854775808}
                            fmt.Println(fi41, v1, v3)
                            _ = v47
                        }
                    default:
                        switch 109 {
                        case 42:
                            if len(v38) > 0 {
                                v38[0] = (fi41 ^ fi41)
                            }
                        case 32767:
                            v2 = true
                        default:
                            fmt.Println(v3, v1, fi41)
                        }
                    }
                case "a":
                    v48 := GenS3{F0: (v5 % (((v5 | v5)) | 1)), F1: 42}
                    _ = v48
                default:
                    var v49 float32 = float32(1e10)
                    _ = v49
                }
                _ = fi41
            }
        case "\n\t":
            var v50 int16 = 1
            _ = v50
        }
        var v51 uint64 = 42
        v3 = v3
        _ = v51
    } else {
        if (!v2) {
            v2 = v2
            v52 := GenS4{F0: (v1 << ((739) & 63)), F1: 339, F2: (v39 ^ (v39 >> ((218) & 63)))}
            var v53 float32 = (-2.25 + v4)
            _ = v52
            _ = v53
        } else {
            fmt.Println(v5)
        }
        var v54 int32 = -2147483648
        fmt.Println(v54, v3)
        _ = v54
    }
    for fi55 := 0; fi55 < 14; fi55++ {
        if false {
            var v56 uint16 = 0
            v5 = 0
            if len(v38) > 0 {
                v38[0] = -839
            }
            _ = v56
        } else {
            for fi57 := 0; fi57 < 11; fi57++ {
                switch (fi57 ^ 32767) {
                case -86:
                    if (18446744073709551615 > 2) {
                        var v58 float32 = v4
                        _ = v58
                    } else {
                        v59 := GenS2{F0: v5, F1: (v1 - (v1 % (((v1 << ((749) & 63))) | 1)))}
                        v5 += v5
                        if true {
                            v60 := GenS4{F0: 1, F1: 578, F2: fi57}
                            v39 = (^(fi57))
                            _ = v60
                        }
                        v3 = "hello"
                        if v2 {
                            _ = len(v38)
                            if len(v38) >= 2 {
                                v38 = v38[1:len(v38)]
                            }
                            var v61 int64 = 42
                            _ = v61
                        }
                        v2 = (32767 < 203)
                        _ = v59
                    }
                case 1:
                    v62 := GenS5{F0: 1, F1: 640}
                    _ = v62
                case -476:
                    v4 -= (float32((0.1 * 0.0)) / 1.0)
                }
                var v63 uint16 = 319
                _ = fi57
                _ = v63
            }
            switch fi55 {
            case 32768:
                v3 = (v3 + "hello")
            default:
                if len(v38) >= 2 {
                    v38 = v38[1:len(v38)]
                }
            }
        }
        for fi64 := 0; fi64 < 7; fi64++ {
            v3 = "hello"
            _ = len(v38)
            _ = fi64
        }
        _ = fi55
    }
    var v65 uint32 = 2
    v1 += (v1 << ((4294967295) & 63))
    var v66 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v65
    _ = v66
}
