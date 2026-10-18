package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 int16; F1 int64; F2 uint }
func genFn0() uint {
    var v4 int64 = 42
    switch 1 {
    case -9223372036854775808:
        v5 := GenS1{F0: 329}
        _ = v5
    case 127:
        v6 := []int{9223372036854775807}
        _ = v6
    case -376:
        switch -765 {
        case 708:
            fmt.Println(v4)
        default:
            v4 += v4
        }
    default:
        v7 := GenS1{F0: 255}
        _ = v7
    }
    _ = v4
    return 362
}

func main() {
    v1 := []int{-9223372036854775808}
    var v2 float64 = 0.0
    var v3 uint16 = 255
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint64 = 515
        _ = fi8
        _ = v9
    }
    v1 = append(v1, -2147483648)
    if false {
        switch -859 {
        case 4294967295:
            for fi10 := 0; fi10 < 13; fi10++ {
                v2 = v2
                if false {
                    v3 += 0
                    var v11 bool = true
                    _ = v11
                } else {
                    if len(v1) > 0 {
                        v1[0] = fi10
                    }
                }
                v2 -= v2
                _ = fi10
            }
        case 128:
            v12 := GenS2{F0: false, F1: 0, F2: -1}
            _ = v12
        case 127:
            var v13 int = 413
            _ = v13
        default:
            switch "\n\t" {
            case "":
                var v14 uint8 = 157
                _ = v14
            case "foo bar":
                var v15 float64 = 3.14159
                _ = v15
            case "\n\t":
                v16 := GenS3{F0: (v3 << ((198) & 63))}
                _ = v16
            }
        }
    }
    if (((("hello" + (("\n\t" + "a") + (("a" + ("hello" + "a")) + "hello"))) + "\n\t") + ("\n\t" + "foo bar")) < ("" + ("a" + "\n\t"))) {
        v2 += float64(0.1)
    } else {
        var v17 uint32 = 365
        if true {
            var v18 uint32 = 4294967295
            _ = v18
        }
        for fi19 := 0; fi19 < 2; fi19++ {
            switch (-(fi19)) {
            case -1:
                for fi20 := 0; fi20 < 15; fi20++ {
                    if (!true) {
                        _ = len(v1)
                    } else {
                        v21 := GenS3{F0: v3}
                        _ = v21
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v22 := GenS2{F0: true, F1: (v17 / ((1) | 1)), F2: 1}
                    var v23 int16 = -180
                    v24 := GenS2{F0: (2 != 42), F1: (v17 - (v17 + (v17 ^ 392))), F2: v23}
                    _ = fi20
                    _ = v22
                    _ = v23
                    _ = v24
                }
            case 4294967295:
                var v25 uint = 307
                _ = v25
            default:
                switch fi19 {
                case 677:
                    v26 := GenS4{F0: (v17 & v17)}
                    _ = v26
                case -956:
                    if false {
                        v17 += 1
                    }
                case 791:
                    switch (fi19 & (fi19 << ((42) & 63))) {
                    case -2147483648:
                        v17 = 2
                    default:
                        var v27 int64 = 0
                        _ = v27
                    }
                default:
                    if true {
                        if (9223372036854775807 != 0) {
                            switch fi19 {
                            case 1:
                                var v28 int16 = 29
                                _ = v28
                            case 4294967295:
                                switch -1 {
                                case 2147483647:
                                    var v29 string = ("\n\t" + (("hello" + "") + ("" + "\n\t")))
                                    _ = v29
                                }
                            case 163:
                                fmt.Println(v17, v17)
                            }
                        } else {
                            var v30 uint64 = 1
                            switch fi19 {
                            case 4294967295:
                                v3 = 965
                            case 32768:
                                var v31 float64 = v2
                                _ = v31
                            case 0:
                                var v32 uint32 = (v17 + 0)
                                _ = v32
                            }
                            switch -2147483648 {
                            case 9223372036854775807:
                                fmt.Println(v17, v17, v30)
                            }
                            _ = v30
                        }
                    } else {
                        v33 := GenS3{F0: 42}
                        fmt.Println(v17)
                        _ = v33
                    }
                }
            }
            _ = fi19
        }
        _ = v17
    }
    v2 = ((v2 / v2) * (v2 * 0.0))
    var v34 uint8 = 255
    switch 127 {
    case -2147483648:
        v34 += 128
    case 9223372036854775807:
        for fi35 := 0; fi35 < 3; fi35++ {
            v34 += (v34 / ((5) | 1))
            for fi36 := 0; fi36 < 5; fi36++ {
                fmt.Println(v2, v3)
                v37 := GenS5{F0: 510, F1: v3}
                v38 := GenS6{F0: -1, F1: -1}
                _ = fi36
                _ = v37
                _ = v38
            }
            if (!(4294967295 <= 2147483647)) {
                v39 := GenS7{F0: -273, F1: 42, F2: 4294967295}
                _ = v39
            }
            _ = fi35
        }
    case 128:
        v3 -= 65535
    }
    for fi40 := 0; fi40 < 16; fi40++ {
        _ = len(v1)
        var v41 int16 = -97
        _ = fi40
        _ = v41
    }
    var v42 string = "\n\t"
    v2 += 0.0
    v1 = append(v1, 443)
    var v43 uint8 = v34
    for fi44 := 0; fi44 < 10; fi44++ {
        switch "a" {
        case "a":
            v45 := GenS7{F0: 107, F1: 0, F2: 539}
            _ = v45
        default:
            v43 += 42
        }
        if true {
            fmt.Println(v43)
        }
        if len(v1) > 0 {
            v1[0] = fi44
        }
        _ = fi44
    }
    v43 -= v34
    if false {
        if (false && (123 > uint32(v3))) {
            var v46 int = 128
            v34 -= (v43 * (v43 + (v43 / (((v43 << ((255) & 63))) | 1))))
            fmt.Println(v43)
            _ = v46
        } else {
            if true {
                v34 = (^(v43))
                v34 -= v34
            }
        }
        for fi47 := 0; fi47 < 2; fi47++ {
            for fi48 := 0; fi48 < 12; fi48++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(fi47, v42)
                var v49 int32 = 0
                _ = fi48
                _ = v49
            }
            v42 = "hello"
            _ = fi47
        }
        v42 = ""
    } else {
        var v50 uint = 15
        for fi51 := 0; fi51 < 12; fi51++ {
            var v52 int8 = -34
            v1 = append(v1, -336)
            for fi53 := 0; fi53 < 16; fi53++ {
                v1 = append(v1, (-(fi53)))
                if (v34 <= (v43 << ((v50) & 63))) {
                    var v54 bool = (!(v50 <= v50))
                    if v54 {
                        switch v42 {
                        case "":
                            v34 += (v43 & (^(v43)))
                        }
                        var v55 float32 = 1e10
                        v54 = (v54 || true)
                        _ = v55
                    }
                    _ = len(v1)
                    v1 = append(v1, -379)
                    _ = v54
                }
                v1 = append(v1, (fi51 - (fi53 | fi51)))
                _ = fi53
            }
            _ = fi51
            _ = v52
        }
        v3 = (v3 % ((v3) | 1))
        _ = v50
    }
    var v56 float64 = (3.14159 / (0.1 + (0.1 + v2)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v42
    _ = v43
    _ = v56
}
