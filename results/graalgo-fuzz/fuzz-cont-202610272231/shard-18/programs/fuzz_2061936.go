package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint64; F1 int; F2 bool }
type GenS4 struct { F0 bool; F1 int; F2 string }
type GenS5 struct { F0 int; F1 string; F2 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 int32; F1 uint64 }
type GenS9 struct { F0 int64 }
func genFn0(p0 uint8, p1 uint32) string {
    p0 += (^(p0))
    if (((-2.25 / 1.5) + 3.14159) <= (-2.25 / 1.5)) {
        v6 := GenS1{F0: 32768}
        _ = v6
    } else {
        p1 = p1
    }
    var v7 float32 = 0.0
    v7 += 3.14159
    var v8 int = 127
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return ("hello" + "foo bar")
}
func genFn1(p0 int32, p1 uint64, p2 int32) uint {
    var v9 uint16 = 0
    v10 := GenS1{F0: 128}
    fmt.Println(p2, v9, p0)
    var v11 float64 = ((1e10 + ((1e10 / -2.25) - -2.25)) - (0.0 * 1.5))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 255
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []int{0, 4294967295, 467, 9223372036854775807}
    var v3 string = "hello"
    var v4 float64 = -2.25
    var v5 float64 = float64(1e10)
    _ = len(v2)
    v5 = (v4 / (float64(float32((v4 + ((3.14159 + (3.14159 / v4)) * v5)))) + v4))
    v12 := GenS1{F0: -2147483648}
    fmt.Println(v1, v4, v3)
    v5 -= v4
    var v13 uint16 = 65535
    v1 -= v1
    v5 -= 0.1
    if ((v13 & v13) == (v13 >> ((42) & 63))) {
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 bool = false
            v13 += (^(v13))
            v16 := GenS1{F0: 1}
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 bool = (!false)
        _ = v17
    }
    switch 0 {
    case 4294967295:
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 int64 = v1
            var v20 int16 = -379
            if (false || false) {
                if (-317 != (fi18 - fi18)) {
                    if (!true) {
                        if (42 <= 42) {
                            fmt.Println(v13, fi18, v20)
                        } else {
                            var v21 string = ""
                            var v22 bool = false
                            _ = v21
                            _ = v22
                        }
                        for fi23 := 0; fi23 < 1; fi23++ {
                            v24 := GenS1{F0: 127}
                            _ = fi23
                            _ = v24
                        }
                    } else {
                        _ = len(v2)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
                switch (^(fi18)) {
                case -811:
                    v3 = (("" + "\n\t") + v3)
                default:
                    _ = len(v2)
                }
            } else {
                if (!true) {
                    v25 := GenS2{F0: 255, F1: 223, F2: -330}
                    _ = v25
                } else {
                    v26 := GenS3{F0: 0, F1: 1, F2: true}
                    if false {
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            var v27 int64 = v1
                            var v28 int16 = (^(v20))
                            var v29 uint16 = (v13 << ((genFn1(2147483647, 945, -1)) & 63))
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        var v30 uint64 = 18446744073709551615
                        for fi31 := 0; fi31 < 13; fi31++ {
                            var v32 uint = 2
                            _ = fi31
                            _ = v32
                        }
                        _ = v30
                    } else {
                        v33 := GenS3{F0: 1, F1: (fi18 ^ fi18), F2: false}
                        if true {
                            if true {
                                var v34 int = fi18
                                fmt.Println(v13)
                                v20 = (v20 | (v20 & 843))
                                _ = v34
                            } else {
                                if false {
                                    switch "hello" {
                                    case "a":
                                        fmt.Println(v20, v20)
                                    }
                                    v3 = "a"
                                    switch ("hello" + "") {
                                    case "foo bar":
                                        v13 = (v13 * v13)
                                    default:
                                        fmt.Println(v3, v13)
                                    }
                                } else {
                                    var v35 uint = 42
                                    _ = len(v2)
                                    switch (-(fi18)) {
                                    case 282:
                                        for fi36 := 0; fi36 < 17; fi36++ {
                                            v5 = 1.5
                                            _ = fi36
                                        }
                                    case 684:
                                        v37 := GenS1{F0: fi18}
                                        _ = v37
                                    case -2147483648:
                                        switch v3 {
                                        case "\n\t":
                                            switch 128 {
                                            case 32768:
                                                v3 = v3
                                            case -818:
                                                fmt.Println(v20, v3)
                                            default:
                                                if false {
                                                    _ = len(v2)
                                                }
                                            }
                                        case "":
                                            var v38 string = "foo bar"
                                            _ = v38
                                        }
                                    default:
                                        if true {
                                            v2 = append(v2, fi18)
                                            if true {
                                                if (4294967295 != 0) {
                                                    var v39 float64 = float64((3.14159 * 0.1))
                                                    if true {
                                                        for fi40 := 0; fi40 < 13; fi40++ {
                                                            v5 = 3.14159
                                                            if false {
                                                                var v41 string = "hello"
                                                                var v42 bool = ((v1 + v19) >= (v1 / ((v1) | 1)))
                                                                fmt.Println(v39)
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                if true {
                                                                    var v43 uint16 = 1
                                                                    _ = v43
                                                                }
                                                                v39 -= v5
                                                                var v44 int8 = 127
                                                                _ = v41
                                                                _ = v42
                                                                _ = v44
                                                            } else {
                                                                fmt.Println(v35)
                                                                if (v1 < (v19 & (v19 | (v19 | v1)))) {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = fi40
                                                                    }
                                                                }
                                                            }
                                                            v45 := GenS4{F0: true, F1: fi18, F2: (((v3 + v3) + "a") + (v3 + v3))}
                                                            if true {
                                                                var v46 uint8 = 128
                                                                _ = v46
                                                            }
                                                            v13 = (v13 & v13)
                                                            if len(v2) > 0 {
                                                                v2[0] = fi40
                                                            }
                                                            var v47 int64 = v19
                                                            if false {
                                                                fmt.Println(v19, v3, v4)
                                                            }
                                                            _ = fi40
                                                            _ = v45
                                                            _ = v47
                                                        }
                                                        var v48 int32 = -768
                                                        _ = v48
                                                    }
                                                    _ = v39
                                                } else {
                                                    v4 += (float64(0.0) * ((1e10 * ((0.1 + v4) / (float64(3.14159) / v5))) / -2.25))
                                                    switch int(v35) {
                                                    case 9223372036854775807:
                                                        v13 -= 42
                                                    case -2147483648:
                                                        v2 = append(v2, (fi18 & fi18))
                                                    case 32767:
                                                        fmt.Println(v1, v35, v3)
                                                    }
                                                }
                                                var v49 uint = 1
                                                _ = v49
                                            }
                                        } else {
                                            switch (fi18 / ((fi18) | 1)) {
                                            case 32768:
                                                var v50 int = 2147483647
                                                _ = v50
                                            default:
                                                v4 += 0.0
                                            }
                                            v51 := GenS5{F0: (fi18 & (fi18 ^ fi18)), F1: v3, F2: v13}
                                            v3 = "hello"
                                            _ = v51
                                        }
                                    }
                                    _ = v35
                                }
                                switch (v3 + v3) {
                                case "hello":
                                    for fi52 := 0; fi52 < 7; fi52++ {
                                        var v53 int64 = 42
                                        _ = fi52
                                        _ = v53
                                    }
                                case "a":
                                    fmt.Println(v3, v5, v3)
                                case "":
                                    var v54 int8 = 127
                                    _ = v54
                                }
                            }
                        } else {
                            var v55 int16 = -1
                            _ = v55
                        }
                        var v56 uint16 = v13
                        _ = v33
                        _ = v56
                    }
                    if false {
                        var v57 int64 = int64(v20)
                        _ = v57
                    } else {
                        switch fi18 {
                        case 9223372036854775807:
                            fmt.Println(v5, v5)
                        }
                        fmt.Println(v19, v4, v20)
                        v58 := GenS6{F0: 42}
                        _ = v58
                    }
                    _ = v26
                }
                v3 = ((((v3 + ("foo bar" + (v3 + v3))) + v3) + v3) + v3)
            }
            _ = fi18
            _ = v19
            _ = v20
        }
    case 490:
        switch 9223372036854775807 {
        case -9223372036854775808:
            _ = len(v2)
        }
    case 128:
        for fi59 := 0; fi59 < 14; fi59++ {
            switch "\n\t" {
            case "hello":
                for fi60 := 0; fi60 < 8; fi60++ {
                    if ((!((true || (-1 <= 32767)) && (!false))) && false) {
                        v61 := GenS4{F0: ((^(fi60)) >= fi60), F1: (fi59 & -415), F2: v3}
                        if (fi60 >= 2147483647) {
                            fmt.Println(v1, fi60)
                        }
                        _ = v61
                    }
                    if (0 != 0) {
                        v62 := GenS5{F0: int(v1), F1: "\n\t", F2: (v13 - v13)}
                        _ = v62
                    } else {
                        v63 := GenS7{F0: 18446744073709551615, F1: (fi60 + (fi60 ^ fi59))}
                        _ = v63
                    }
                    _ = fi60
                }
            case "foo bar":
                for fi64 := 0; fi64 < 10; fi64++ {
                    v65 := GenS8{F0: 668, F1: 0}
                    v4 += v4
                    _ = fi64
                    _ = v65
                }
            case "\n\t":
                if false {
                    v2 = append(v2, int(v13))
                    v13 = v13
                    for fi66 := 0; fi66 < 12; fi66++ {
                        fmt.Println(v1, v13)
                        fmt.Println(v5, fi59)
                        _ = fi66
                    }
                } else {
                    var v67 int16 = 42
                    _ = v67
                }
            default:
                v68 := GenS9{F0: (-(v1))}
                _ = v68
            }
            for fi69 := 0; fi69 < 12; fi69++ {
                if (!false) {
                    var v70 uint64 = 4294967295
                    var v71 int16 = 339
                    _ = v70
                    _ = v71
                }
                fmt.Println(v1, v4, fi69)
                var v72 int64 = v1
                _ = fi69
                _ = v72
            }
            _ = fi59
        }
    }
    v73 := GenS6{F0: -1}
    fmt.Println(v5, v13)
    fmt.Println(v5, v1, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v73
}
