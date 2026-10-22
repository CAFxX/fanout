package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 int8; F1 int }
func genFn0(p0 uint8) int {
    var v4 int16 = 1
    if (v4 < 1000) {
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 uint = 65535
            p0 = 2
            v7 := GenS1{F0: 42, F1: p0}
            _ = fi5
            _ = v6
            _ = v7
        }
        for fi8 := 0; fi8 < 20; fi8++ {
            p0 += 162
            var v9 int8 = 1
            _ = fi8
            _ = v9
        }
        var v10 uint64 = 42
        _ = v10
    }
    fmt.Println(v4, v4)
    var v11 int32 = -479
    _ = p0
    _ = v11
    _ = v4
    return 128
}
func genFn1() int8 {
    for fi12 := 0; fi12 < 8; fi12++ {
        if ((!false) && true) {
            var v13 uint = 2
            v13 = (v13 / ((v13) | 1))
            var v14 uint32 = 2
            _ = v13
            _ = v14
        }
        var v15 uint32 = 2
        if (-9223372036854775808 == 762) {
            var v16 uint32 = 2
            v15 = v15
            for fi17 := 0; fi17 < 13; fi17++ {
                if true {
                    var v18 int32 = -153
                    v19 := GenS1{F0: 1, F1: 42}
                    _ = v18
                    _ = v19
                } else {
                    if ((v15 ^ (v15 + v15)) > (v15 & v16)) {
                        v20 := []int{}
                        v21 := GenS2{F0: 32767}
                        _ = v20
                        _ = v21
                    } else {
                        if ((1 == 2) && true) {
                            var v22 uint64 = 555
                            v16 = 65535
                            v23 := []int{(fi12 / ((fi12) | 1)), -92, (^(fi17)), fi17}
                            _ = v22
                            _ = v23
                        } else {
                            var v24 int64 = 375
                            var v25 int8 = 100
                            v26 := []int{(-(fi12)), fi17, genFn0(242)}
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        var v27 int8 = -26
                        v28 := GenS1{F0: 1, F1: 127}
                        _ = v27
                        _ = v28
                    }
                    if false {
                        v29 := []int{genFn0(1), fi17, genFn0(128), (fi12 / ((678) | 1))}
                        var v30 int16 = 1
                        fmt.Println(fi12, v16)
                        _ = v29
                        _ = v30
                    }
                    switch (("a" + (("a" + "hello") + ("" + "foo bar"))) + "hello") {
                    case "foo bar":
                        switch "hello" {
                        case "a":
                            var v31 uint64 = 4294967295
                            _ = v31
                        case "\n\t":
                            v16 += 467
                        }
                    case "hello":
                        switch fi12 {
                        case 127:
                            v16 -= 2147483647
                            var v32 bool = ((float64(0.1) / 1.5) != 1e10)
                            _ = v32
                        case -9223372036854775808:
                            var v33 int = (fi17 ^ genFn0(1))
                            _ = v33
                        case -1:
                            if (true || (!(!true))) {
                                v34 := GenS2{F0: int16(v16)}
                                v35 := GenS3{F0: (fi12 << ((2) & 63)), F1: 380}
                                _ = v34
                                _ = v35
                            }
                        }
                    }
                    v36 := GenS1{F0: 65535, F1: 25}
                    _ = v36
                }
                _ = fi17
            }
            _ = v16
        } else {
            switch ("foo bar" + ("foo bar" + ("" + "a"))) {
            case "hello":
                v37 := GenS1{F0: 2, F1: 104}
                _ = v37
            default:
                fmt.Println(fi12, fi12)
            }
            if (342 < 4294967295) {
                v15 = (v15 * v15)
                if true {
                    v38 := GenS3{F0: fi12, F1: 65535}
                    _ = v38
                }
                var v39 int = fi12
                _ = v39
            } else {
                v15 += (v15 | (v15 % ((v15) | 1)))
                for fi40 := 0; fi40 < 16; fi40++ {
                    v41 := GenS1{F0: 854, F1: 255}
                    var v42 uint64 = 815
                    fmt.Println(v15, v15)
                    var v43 string = "a"
                    _ = fi40
                    _ = v41
                    _ = v42
                    _ = v43
                }
                switch fi12 {
                case 617:
                    var v44 uint16 = 42
                    _ = v44
                }
            }
        }
        _ = fi12
        _ = v15
    }
    var v45 int64 = 1
    v46 := []int{904}
    var v47 uint8 = 177
    for fi48 := 0; fi48 < 6; fi48++ {
        var v49 uint64 = 4294967295
        v50 := GenS1{F0: 42, F1: (v47 * v47)}
        v51 := GenS4{F0: 9223372036854775807}
        _ = fi48
        _ = v49
        _ = v50
        _ = v51
    }
    _ = v45
    _ = v46
    _ = v47
    return 0
}

func main() {
    var v1 uint32 = 856
    v2 := []uint32{}
    var v3 string = "foo bar"
    fmt.Println(v3, v1)
    if len(v2) > 0 {
        v2[0] = v1
    }
    var v52 int16 = -181
    _ = len(v2)
    v2 = append(v2, v1)
    if true {
        var v53 int16 = (v52 / (((v52 << ((65535) & 63))) | 1))
        v3 = (v3 + v3)
        _ = v53
    } else {
        var v54 int16 = v52
        var v55 int64 = 981
        if (!false) {
            if true {
                _ = len(v2)
            } else {
                switch 494 {
                case -2147483648:
                    for fi56 := 0; fi56 < 12; fi56++ {
                        fmt.Println(v52, v52)
                        _ = fi56
                    }
                case 1:
                    v57 := GenS5{F0: (v1 / (((v1 - (v1 | v1))) | 1)), F1: (v54 | (v52 + (v52 % ((121) | 1))))}
                    _ = v57
                }
                _ = len(v2)
                v58 := GenS4{F0: v55}
                _ = v58
            }
        } else {
            if len(v2) > 0 {
                v2[0] = 42
            }
        }
        _ = v54
        _ = v55
    }
    switch v3 {
    case "a":
        v59 := GenS4{F0: -181}
        _ = v59
    case "hello":
        v60 := GenS1{F0: 975, F1: 2}
        _ = v60
    case "foo bar":
        if (true && false) {
            v2 = append(v2, v1)
        }
    default:
        for fi61 := 0; fi61 < 4; fi61++ {
            _ = len(v2)
            v62 := GenS5{F0: (v1 * v1), F1: (v52 << ((4294967295) & 63))}
            _ = fi61
            _ = v62
        }
    }
    v63 := GenS6{F0: -128, F1: 32768}
    v3 = v3
    var v64 bool = ((!((3.14159 != 1.5) && true)) && true)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
    _ = v63
    _ = v64
}
