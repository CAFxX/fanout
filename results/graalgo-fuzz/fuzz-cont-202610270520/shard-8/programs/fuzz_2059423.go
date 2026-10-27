package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 int; F2 int8 }
type GenS4 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint64; F1 int64 }
func genFn0() int {
    v4 := GenS1{F0: 0, F1: -2147483648, F2: 4294967295}
    for fi5 := 0; fi5 < 15; fi5++ {
        switch "hello" {
        case "":
            var v6 uint16 = 42
            _ = v6
        case "foo bar":
            var v7 uint16 = 814
            _ = v7
        case "hello":
            v8 := GenS1{F0: 1, F1: -1, F2: 2}
            _ = v8
        }
        _ = fi5
    }
    var v9 uint8 = 255
    _ = v4
    _ = v9
    return 32768
}
func genFn1(p0 int16) int32 {
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(fi10, p0)
        _ = fi10
    }
    p0 += p0
    v11 := GenS1{F0: 0, F1: 0, F2: 0}
    _ = p0
    _ = v11
    return -150
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint8 = 62
    var v3 float64 = ((1e10 * 1.5) / float64(3.14159))
    var v12 int32 = (v1 - v1)
    switch 42 {
    case 909:
        var v13 string = ("hello" + (("a" + "hello") + ("" + "\n\t")))
        _ = v13
    case -989:
        for fi14 := 0; fi14 < 2; fi14++ {
            v3 -= 1e10
            v15 := []int{(fi14 + (^(fi14))), 9223372036854775807, (fi14 ^ 701)}
            _ = fi14
            _ = v15
        }
    case -235:
        switch 237 {
        case 128:
            var v16 uint = 980
            _ = v16
        case -306:
            if (true || false) {
                v1 = 0
                var v17 int8 = 0
                v17 = -128
                _ = v17
            }
        case 0:
            var v18 int16 = -278
            _ = v18
        }
    }
    if (!(!false)) {
        switch -1 {
        case -670:
            var v19 string = "hello"
            _ = v19
        case 237:
            switch 4294967295 {
            case -1:
                if (int16(v12) >= 1) {
                    v3 -= v3
                    if true {
                        if false {
                            v20 := GenS2{F0: -9223372036854775808}
                            _ = v20
                        } else {
                            for fi21 := 0; fi21 < 2; fi21++ {
                                v3 -= v3
                                for fi22 := 0; fi22 < 20; fi22++ {
                                    var v23 uint16 = 2
                                    switch "a" {
                                    case "\n\t":
                                        if false {
                                            switch "hello" {
                                            case "hello":
                                                v23 += (v23 | v23)
                                            case "\n\t":
                                                switch ("foo bar" + "a") {
                                                case "\n\t":
                                                    var v24 bool = (!true)
                                                    _ = v24
                                                case "hello":
                                                    fmt.Println(fi22)
                                                case "foo bar":
                                                    var v25 uint = 1
                                                    _ = v25
                                                }
                                            case "foo bar":
                                                switch genFn0() {
                                                case -2147483648:
                                                    v26 := []int{fi21}
                                                    _ = v26
                                                case 445:
                                                    if false {
                                                        v27 := []int{(fi22 >> ((94) & 63)), genFn0()}
                                                        _ = v27
                                                    }
                                                default:
                                                    fmt.Println(fi21, v1)
                                                }
                                            }
                                        }
                                    case "a":
                                        var v28 int64 = -332
                                        _ = v28
                                    default:
                                        if ((^(v12)) >= v12) {
                                            v29 := []int{}
                                            v30 := GenS2{F0: 0}
                                            _ = v29
                                            _ = v30
                                        } else {
                                            v31 := []int{-808, (fi21 >> ((1) & 63))}
                                            _ = v31
                                        }
                                    }
                                    _ = fi22
                                    _ = v23
                                }
                                _ = fi21
                            }
                            switch 0 {
                            case 309:
                                v32 := GenS2{F0: 9223372036854775807}
                                _ = v32
                            case -9223372036854775808:
                                v33 := GenS3{F0: 182, F1: -934, F2: 31}
                                _ = v33
                            default:
                                switch 127 {
                                case 328:
                                    v34 := []int{0}
                                    _ = v34
                                case 32767:
                                    v12 -= (^(v12))
                                case -9223372036854775808:
                                    v1 -= (v12 >> ((65535) & 63))
                                default:
                                    v35 := GenS4{F0: 928, F1: 175, F2: 42}
                                    _ = v35
                                }
                            }
                            switch genFn0() {
                            case 2147483647:
                                var v36 float32 = 1.5
                                _ = v36
                            default:
                                for fi37 := 0; fi37 < 16; fi37++ {
                                    v38 := GenS3{F0: (v2 - v2), F1: (fi37 & 0), F2: 100}
                                    var v39 string = ("a" + ("a" + ("\n\t" + "")))
                                    _ = fi37
                                    _ = v38
                                    _ = v39
                                }
                            }
                        }
                        for fi40 := 0; fi40 < 9; fi40++ {
                            fmt.Println(v1, v12, v1)
                            for fi41 := 0; fi41 < 13; fi41++ {
                                if false {
                                    fmt.Println(fi40)
                                    if true {
                                        v42 := []int{(fi41 + (^(fi41))), genFn0(), (fi40 - (fi40 + -778))}
                                        _ = v42
                                    }
                                } else {
                                    v43 := []int{}
                                    var v44 string = "a"
                                    _ = v43
                                    _ = v44
                                }
                                var v45 uint8 = v2
                                var v46 int8 = 100
                                _ = fi41
                                _ = v45
                                _ = v46
                            }
                            var v47 int8 = 38
                            _ = fi40
                            _ = v47
                        }
                        if true {
                            for fi48 := 0; fi48 < 3; fi48++ {
                                var v49 int = fi48
                                var v50 string = "foo bar"
                                for fi51 := 0; fi51 < 13; fi51++ {
                                    var v52 int64 = 42
                                    v53 := []int{fi51}
                                    _ = fi51
                                    _ = v52
                                    _ = v53
                                }
                                _ = fi48
                                _ = v49
                                _ = v50
                            }
                            if ("hello" != ("a" + "\n\t")) {
                                if true {
                                    v3 = ((((1.5 * 0.0) - v3) + 0.0) - (v3 * v3))
                                    fmt.Println(v12)
                                }
                            } else {
                                v54 := []int{genFn0(), -489, 32767, -391}
                                _ = v54
                            }
                            var v55 string = "\n\t"
                            _ = v55
                        }
                    }
                }
            case 32768:
                switch -9223372036854775808 {
                case 9223372036854775807:
                    v3 -= v3
                case 724:
                    v56 := []int{}
                    _ = v56
                }
            case 127:
                v57 := GenS1{F0: 1, F1: (v1 >> ((0) & 63)), F2: 4294967295}
                _ = v57
            default:
                if (65535 <= 2) {
                    var v58 int8 = 127
                    v2 += v2
                    _ = v58
                } else {
                    v59 := []int{genFn0(), genFn0()}
                    v60 := GenS3{F0: v2, F1: genFn0(), F2: -85}
                    var v61 float32 = float32(v3)
                    _ = v59
                    _ = v60
                    _ = v61
                }
            }
        }
        v2 -= v2
        v1 = genFn1(42)
    }
    v62 := GenS5{F0: 2, F1: 9223372036854775807}
    var v63 int16 = 32767
    for fi64 := 0; fi64 < 13; fi64++ {
        for fi65 := 0; fi65 < 4; fi65++ {
            v66 := []int{(-(fi64)), (fi64 >> ((0) & 63)), (fi64 ^ 42), fi65}
            _ = fi65
            _ = v66
        }
        v67 := GenS2{F0: 0}
        _ = fi64
        _ = v67
    }
    var v68 bool = ((187 == 597) && false)
    var v69 int = -1
    var v70 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v63:", v63)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v62
    _ = v63
    _ = v68
    _ = v69
    _ = v70
}
