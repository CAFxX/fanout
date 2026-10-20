package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 uint; F1 bool; F2 string }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 bool; F1 bool; F2 uint32 }
func genFn0() bool {
    var v4 string = "hello"
    var v5 uint = 65535
    v6 := []int{-2147483648, 703, 0, -1}
    fmt.Println(v4, v5)
    _ = v4
    _ = v5
    _ = v6
    return ((true && (!false)) && false)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = "foo bar"
    var v3 uint8 = 182
    if (false || (104 <= 817)) {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 uint16 = 0
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 255, F1: 0, F2: -1}
        if false {
            var v10 int = -1
            _ = v10
        }
        _ = v9
    } else {
        if ((98 < 65535) && (!true)) {
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 int8 = 63
                var v13 bool = false
                var v14 uint = 65535
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                for fi16 := 0; fi16 < 12; fi16++ {
                    switch (fi16 % (((fi15 >> ((838) & 63))) | 1)) {
                    case 128:
                        if (true && true) {
                            var v17 int64 = 1
                            if false {
                                fmt.Println(fi16, fi15, fi16)
                                if (false && false) {
                                    var v18 int64 = (v17 / (((v17 - (v17 % (((-(v17))) | 1)))) | 1))
                                    var v19 uint64 = 0
                                    _ = v18
                                    _ = v19
                                } else {
                                    var v20 uint8 = v3
                                    v2 = v2
                                    var v21 string = "a"
                                    _ = v20
                                    _ = v21
                                }
                                v3 -= (^(v3))
                            }
                            _ = v17
                        } else {
                            var v22 uint32 = 65535
                            _ = v22
                        }
                    default:
                        var v23 bool = (true || (true && true))
                        _ = v23
                    }
                    var v24 string = v2
                    fmt.Println(v2)
                    switch v24 {
                    case "foo bar":
                        v25 := GenS1{F0: 403, F1: 651, F2: 32767}
                        _ = v25
                    case "hello":
                        v3 += (v3 - 0)
                    default:
                        if ((v3 | 2) <= v3) {
                            v26 := []int{(fi16 * fi15), (fi15 & -2147483648), (-(fi16)), (fi15 >> ((379) & 63))}
                            _ = v26
                        } else {
                            v1[927] = fi15
                        }
                    }
                    _ = fi16
                    _ = v24
                }
                if (1000 <= 0) {
                    var v27 int32 = 263
                    v2 = v2
                    _ = v27
                } else {
                    v28 := GenS2{F0: 32767, F1: 1}
                    v29 := GenS3{F0: 0, F1: (((41 >= (v3 >> ((5) & 63))) || (false && (true || (0.1 >= (float64(-2.25) / 1.0))))) && (!true)), F2: 1}
                    for fi30 := 0; fi30 < 4; fi30++ {
                        switch ("foo bar" + v2) {
                        case "foo bar":
                            if true {
                                v31 := GenS2{F0: 144, F1: -1}
                                _ = v31
                            } else {
                                v2 = "\n\t"
                                v2 = "foo bar"
                                switch (fi30 >> ((255) & 63)) {
                                case 127:
                                    v3 -= v3
                                case 337:
                                    if (-883 != 2147483647) {
                                        var v32 uint = 1
                                        if true {
                                            v32 = v32
                                        }
                                        switch 415 {
                                        case 0:
                                            var v33 uint64 = 4294967295
                                            _ = v33
                                        default:
                                            if ((v3 ^ v3) <= 108) {
                                                v3 = (v3 & (^(v3)))
                                                switch (fi30 + 9223372036854775807) {
                                                case 0:
                                                    switch v2 {
                                                    case "foo bar":
                                                        v3 -= (v3 | v3)
                                                    case "hello":
                                                        fmt.Println(fi15)
                                                    case "a":
                                                        v34 := GenS4{F0: (v3 / ((v3) | 1)), F1: 1}
                                                        _ = v34
                                                    }
                                                case 414:
                                                    v35 := GenS5{F0: 18446744073709551615, F1: (^(v32))}
                                                    _ = v35
                                                }
                                                var v36 float64 = ((0.0 / ((float64(0.0) + 1e10) + 3.14159)) - float64((0.1 - 0.1)))
                                                _ = v36
                                            }
                                        }
                                        _ = v32
                                    }
                                }
                            }
                        }
                        v37 := []int{}
                        var v38 int16 = -442
                        _ = fi30
                        _ = v37
                        _ = v38
                    }
                    _ = v28
                    _ = v29
                }
                _ = fi15
            }
        }
        v39 := []int{9223372036854775807, -2147483648, 2147483647, -9223372036854775808}
        _ = v39
    }
    v2 = ("\n\t" + v2)
    for fi40 := 0; fi40 < 16; fi40++ {
        v41 := GenS4{F0: v3, F1: 33}
        _ = len(v1)
        if (!true) {
            v3 -= (v3 & v3)
            v42 := []int{801}
            _ = v42
        }
        _ = fi40
        _ = v41
    }
    v3 = 0
    var v43 uint64 = 2
    v44 := GenS2{F0: 1, F1: 0}
    v45 := GenS5{F0: v43, F1: 0}
    v3 -= v3
    _ = len(v1)
    for fi46 := 0; fi46 < 7; fi46++ {
        v1[806] = fi46
        var v47 float64 = (float64(-2.25) * 1e10)
        fmt.Println(fi46, v2)
        _ = fi46
        _ = v47
    }
    if false {
        v48 := GenS6{F0: 1, F1: true, F2: ("hello" + v2)}
        var v49 int32 = -1
        _ = v48
        _ = v49
    }
    fmt.Println(v3, v43)
    fmt.Println(v2, v3, v43)
    for fi50 := 0; fi50 < 7; fi50++ {
        var v51 uint16 = 2
        var v52 int8 = 42
        for fi53 := 0; fi53 < 17; fi53++ {
            switch fi53 {
            case -1:
                fmt.Println(fi50)
                v52 += v52
                var v54 float32 = 0.1
                _ = v54
            }
            if false {
                v55 := GenS7{F0: 2147483647, F1: (!(!false))}
                v52 += (v52 % (((v52 << ((182) & 63))) | 1))
                _ = v55
            } else {
                v56 := []int{}
                var v57 uint = 217
                v58 := GenS2{F0: v51, F1: 9}
                var v59 float64 = (1.5 / 0.1)
                fmt.Println(v3, fi53)
                _ = v56
                _ = v57
                _ = v58
                _ = v59
            }
            _ = fi53
        }
        _ = fi50
        _ = v51
        _ = v52
    }
    switch -1 {
    case -242:
        for fi60 := 0; fi60 < 1; fi60++ {
            for fi61 := 0; fi61 < 11; fi61++ {
                _ = len(v1)
                var v62 int8 = -34
                v43 -= (^(v43))
                _ = fi61
                _ = v62
            }
            switch (fi60 | -9223372036854775808) {
            case 127:
                v63 := GenS8{F0: (!(1 == 0)), F1: true, F2: 0}
                _ = v63
            }
            _ = fi60
        }
    case -547:
        if (32768 < 32767) {
            var v64 uint64 = (v43 & (v43 * (v43 | v43)))
            v65 := GenS8{F0: (true || false), F1: (!(!true)), F2: 47}
            for fi66 := 0; fi66 < 5; fi66++ {
                if false {
                    var v67 uint64 = 170
                    _ = v67
                } else {
                    var v68 uint8 = 162
                    _ = v68
                }
                v64 = 888
                _ = fi66
            }
            _ = v64
            _ = v65
        } else {
            v3 = v3
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v44
    _ = v45
}
