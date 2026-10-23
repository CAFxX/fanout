package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 int64; F2 int64 }
func genFn0() uint16 {
    var v6 int = 4294967295
    for fi7 := 0; fi7 < 7; fi7++ {
        if (2 >= 1) {
            if (true || false) {
                fmt.Println(v6, v6)
            }
        } else {
            if false {
                if true {
                    v8 := []int{(v6 & (^(fi7))), v6, 51, v6}
                    v6 -= (fi7 << ((65535) & 63))
                    _ = v8
                } else {
                    v9 := GenS1{F0: (v6 | 32767), F1: 0, F2: (v6 + 1)}
                    _ = v9
                }
            }
            switch "foo bar" {
            case "a":
                v6 -= 759
            case "\n\t":
                var v10 int = 499
                _ = v10
            case "foo bar":
                var v11 uint8 = 127
                _ = v11
            }
            for fi12 := 0; fi12 < 1; fi12++ {
                v13 := GenS2{F0: fi12}
                if (0 == 0) {
                    v14 := []int{(-(fi12)), (fi7 * v6), fi7, (-(v6))}
                    v6 = 32768
                    var v15 int32 = 42
                    _ = len(v14)
                    v15 += (v15 * v15)
                    _ = v14
                    _ = v15
                } else {
                    v6 = 0
                }
                if true {
                    v6 += (fi12 | 9223372036854775807)
                    var v16 bool = (!(1 == 878))
                    _ = v16
                }
                _ = fi12
                _ = v13
            }
        }
        _ = fi7
    }
    v6 -= v6
    _ = v6
    return 0
}
func genFn1() int32 {
    var v17 int64 = -474
    if ("hello" > "foo bar") {
        v17 += (v17 >> ((1) & 63))
    }
    _ = v17
    return 1
}

func main() {
    var v1 uint16 = 1
    var v2 float32 = -2.25
    var v3 uint16 = v1
    v4 := []string{}
    var v5 int8 = -1
    if ((((3.14159 - (0.1 / 1.0)) * (1.5 * 0.0)) * (1e10 + (1e10 - ((0.1 / 1.0) - -2.25)))) >= 0.1) {
        if len(v4) > 0 {
            v4[0] = ("" + "\n\t")
        }
    }
    v2 -= 3.14159
    var v18 uint16 = 32767
    switch 236 {
    case 9223372036854775807:
        switch 56 {
        case 94:
            v19 := GenS3{F0: 42}
            _ = v19
        case 4294967295:
            switch -647 {
            case 923:
                for fi20 := 0; fi20 < 9; fi20++ {
                    v21 := GenS3{F0: uint8(v1)}
                    _ = fi20
                    _ = v21
                }
            case -214:
                v22 := GenS1{F0: -2147483648, F1: -32768, F2: 127}
                _ = v22
            default:
                for fi23 := 0; fi23 < 12; fi23++ {
                    fmt.Println(v1, v3, v3)
                    for fi24 := 0; fi24 < 18; fi24++ {
                        switch fi24 {
                        case 752:
                            switch (fi23 % (((fi23 | 893)) | 1)) {
                            case 127:
                                var v25 float64 = -2.25
                                _ = v25
                            case 9223372036854775807:
                                fmt.Println(fi23, v5)
                            case 32767:
                                if (!((((9223372036854775807 != -9223372036854775808) || (!false)) || true) || false)) {
                                    if (!(3.14159 > (float32((float64((v2 * v2)) + 1e10)) + v2))) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v2 -= v2
                                        v26 := GenS3{F0: 255}
                                        _ = v26
                                    } else {
                                        switch "" {
                                        case "":
                                            v27 := GenS1{F0: 0, F1: -1, F2: 4294967295}
                                            _ = v27
                                        }
                                    }
                                    var v28 int = 1
                                    _ = len(v4)
                                    _ = v28
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            default:
                                var v29 float64 = 0.1
                                _ = v29
                            }
                        case 42:
                            v3 = (v1 * 2)
                        default:
                            var v30 uint8 = 0
                            _ = v30
                        }
                        var v31 int = fi24
                        var v32 int64 = 9223372036854775807
                        _ = fi24
                        _ = v31
                        _ = v32
                    }
                    v33 := GenS4{F0: 31}
                    _ = fi23
                    _ = v33
                }
            }
        }
    default:
        if (185 >= 45) {
            for fi34 := 0; fi34 < 6; fi34++ {
                v2 = 3.14159
                _ = fi34
            }
        } else {
            var v35 int = 127
            fmt.Println(v3)
            _ = v35
        }
    }
    var v36 uint64 = 1
    fmt.Println(v2, v5, v5)
    for fi37 := 0; fi37 < 1; fi37++ {
        if (0 < 333) {
            for fi38 := 0; fi38 < 14; fi38++ {
                var v39 uint8 = 127
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi38
                _ = v39
            }
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi40 := 0; fi40 < 1; fi40++ {
                if false {
                    v41 := GenS3{F0: 2}
                    var v42 int64 = -727
                    var v43 uint64 = v36
                    _ = v41
                    _ = v42
                    _ = v43
                }
                v3 = v18
                _ = fi40
            }
            _ = len(v4)
        }
        switch 32767 {
        case 127:
            v44 := GenS5{F0: genFn0(), F1: 42, F2: 42}
            _ = v44
        default:
            for fi45 := 0; fi45 < 19; fi45++ {
                v2 = -2.25
                if ((^(v36)) < 1) {
                    var v46 uint32 = 0
                    v4 = append(v4, ("a" + ("a" + "foo bar")))
                    fmt.Println(v36)
                    _ = v46
                } else {
                    v47 := GenS4{F0: 52}
                    fmt.Println(v5)
                    _ = v47
                }
                fmt.Println(v18)
                _ = fi45
            }
        }
        _ = fi37
    }
    var v48 float64 = (0.0 + float64((-2.25 + 3.14159)))
    if (-1 < int32(v18)) {
        v5 = (v5 >> ((18446744073709551615) & 63))
        v48 -= v48
        var v49 bool = (((true || true) && false) && (!false))
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v48
    _ = v5
}
