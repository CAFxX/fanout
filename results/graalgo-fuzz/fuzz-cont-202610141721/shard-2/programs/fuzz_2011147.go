package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 bool; F1 string; F2 int }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 int8 }
func genFn0(p0 bool) uint8 {
    var v4 uint32 = 1
    var v5 string = "\n\t"
    _ = p0
    _ = v4
    _ = v5
    return 128
}
func genFn1(p0 uint32, p1 int16) int64 {
    v6 := []int{-2147483648, -182, 9223372036854775807}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 float64 = 3.14159
        v9 := GenS1{F0: p0, F1: p0, F2: int(p1)}
        fmt.Println(p1, fi7)
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return -680
}

func main() {
    var v1 int8 = -8
    var v2 int64 = -990
    var v3 uint = 294
    v2 = 489
    if false {
        var v10 int8 = v1
        var v11 float32 = 0.0
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := GenS1{F0: 2147483648, F1: 1, F2: 974}
            for fi14 := 0; fi14 < 18; fi14++ {
                fmt.Println(v3, fi12)
                fmt.Println(fi12, v1, fi14)
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
        var v15 uint = 609
        v16 := GenS1{F0: 2147483647, F1: 2, F2: -1}
        _ = v15
        _ = v16
    }
    v2 -= v2
    v17 := GenS1{F0: 2147483647, F1: 2147483648, F2: 163}
    var v18 uint = (v3 >> ((0) & 63))
    v19 := GenS2{F0: -2147483648, F1: -128}
    fmt.Println(v1, v3, v18)
    for fi20 := 0; fi20 < 18; fi20++ {
        fmt.Println(v1, v1, fi20)
        v21 := GenS2{F0: 1, F1: 127}
        var v22 bool = true
        _ = fi20
        _ = v21
        _ = v22
    }
    switch "a" {
    case "\n\t":
        if (65535 >= 255) {
            var v23 uint16 = 0
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 uint32 = 950
                switch "" {
                case "\n\t":
                    v26 := []int{}
                    _ = v26
                case "foo bar":
                    switch (("\n\t" + "") + "a") {
                    case "a":
                        for fi27 := 0; fi27 < 6; fi27++ {
                            v28 := GenS3{F0: true, F1: (v2 / (((^(v2))) | 1))}
                            switch fi24 {
                            case 9223372036854775807:
                                v29 := []int{(fi24 ^ fi24)}
                                _ = v29
                            default:
                                v30 := []int{(fi27 << (((v3 * (v3 % (((v18 - v18)) | 1)))) & 63)), fi27, fi24, (fi24 >> (((v3 | 65535)) & 63))}
                                _ = v30
                            }
                            v31 := GenS3{F0: (!true), F1: v2}
                            _ = fi27
                            _ = v28
                            _ = v31
                        }
                    case "hello":
                        var v32 uint8 = 255
                        _ = v32
                    }
                }
                for fi33 := 0; fi33 < 1; fi33++ {
                    if (true && ((0.1 >= (1.5 * -2.25)) || ((4294967295 <= 2) && false))) {
                        v25 = (v25 & 678)
                        v34 := GenS3{F0: (((v2 >> ((13) & 63)) <= (v2 >> ((979) & 63))) && false), F1: v2}
                        v35 := []int{-685, fi33}
                        _ = v34
                        _ = v35
                    }
                    var v36 float64 = 1e10
                    v23 += v23
                    v23 = (v23 % ((v23) | 1))
                    _ = fi33
                    _ = v36
                }
                _ = fi24
                _ = v25
            }
            _ = v23
        } else {
            fmt.Println(v18)
            for fi37 := 0; fi37 < 14; fi37++ {
                if (v1 >= (v1 ^ int8(v18))) {
                    v3 += (^(v18))
                }
                var v38 float32 = (0.1 - 1.5)
                if true {
                    v39 := GenS4{F0: false, F1: "foo bar", F2: (fi37 ^ fi37)}
                    _ = v39
                }
                _ = fi37
                _ = v38
            }
        }
    case "":
        var v40 int16 = 42
        _ = v40
    case "a":
        if (18446744073709551615 <= 4294967295) {
            for fi41 := 0; fi41 < 6; fi41++ {
                v42 := []int{-720, (fi41 / (((fi41 >> ((916) & 63))) | 1)), fi41}
                v43 := GenS5{F0: "\n\t", F1: 472}
                var v44 uint8 = 255
                _ = fi41
                _ = v42
                _ = v43
                _ = v44
            }
        }
    }
    for fi45 := 0; fi45 < 5; fi45++ {
        v46 := GenS3{F0: true, F1: -139}
        _ = fi45
        _ = v46
    }
    var v47 uint = v18
    for fi48 := 0; fi48 < 14; fi48++ {
        for fi49 := 0; fi49 < 6; fi49++ {
            v50 := GenS2{F0: 2147483647, F1: (v1 & (v1 >> ((v18) & 63)))}
            if (399 >= 0) {
                if false {
                    if false {
                        if (true && (!false)) {
                            switch (fi48 | (fi49 ^ 9223372036854775807)) {
                            case 2147483647:
                                switch (("foo bar" + "\n\t") + ("hello" + "foo bar")) {
                                case "a":
                                    v51 := GenS6{F0: v1}
                                    _ = v51
                                case "\n\t":
                                    if (42 <= 1) {
                                        v52 := []int{fi48, 2147483647, (fi48 / ((fi49) | 1))}
                                        _ = v52
                                    } else {
                                        if true {
                                            switch "foo bar" {
                                            case "a":
                                                v3 += (v18 | v18)
                                            case "hello":
                                                switch fi48 {
                                                case 127:
                                                    var v53 uint64 = 732
                                                    _ = v53
                                                }
                                            case "foo bar":
                                                var v54 int32 = 0
                                                _ = v54
                                            default:
                                                fmt.Println(v1, v47)
                                            }
                                            v55 := []int{-9223372036854775808}
                                            var v56 int16 = 42
                                            v3 += v3
                                            _ = v55
                                            _ = v56
                                        }
                                        var v57 bool = false
                                        _ = v57
                                    }
                                case "foo bar":
                                    v1 = (^(v1))
                                }
                            case -219:
                                v58 := []int{fi48}
                                _ = v58
                            case 4294967295:
                                v2 += (v2 >> ((v47) & 63))
                            }
                        }
                        v18 += (v47 | v18)
                        v59 := []int{fi48, 127, fi49}
                        _ = v59
                    } else {
                        v1 = (v1 & int8(v18))
                    }
                    v18 += (v18 ^ (v18 ^ 2))
                } else {
                    v47 = 42
                }
            }
            _ = fi49
            _ = v50
        }
        _ = fi48
    }
    var v60 int16 = 1000
    v61 := []int{9223372036854775807, 129}
    var v62 int = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v47
    _ = v60
    _ = v61
    _ = v62
}
