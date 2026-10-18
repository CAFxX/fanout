package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS3 struct { F0 bool; F1 bool; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint; F2 uint }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 uint32; F1 int }
func genFn0(p0 int32, p1 bool, p2 uint32) bool {
    p0 -= (p0 ^ -1)
    var v5 int = 719
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 2
    v3 := make(map[string]int)
    var v4 bool = true
    var v6 string = "a"
    fmt.Println(v6)
    switch 148 {
    case 470:
        var v7 string = "a"
        _ = v7
    case -660:
        fmt.Println(v2, v2)
    case 4294967295:
        v4 = (false && true)
    }
    if v4 {
        if v4 {
            switch ((v6 + v6) + v6) {
            case "":
                v4 = ((1.5 - 0.1) >= 0.1)
            }
            if (976 < 4294967295) {
                v8 := GenS1{F0: 1}
                switch "" {
                case "hello":
                    var v9 uint = 952
                    _ = v9
                case "a":
                    v10 := GenS1{F0: 1}
                    _ = v10
                case "\n\t":
                    var v11 bool = v4
                    _ = v11
                }
                var v12 string = (v6 + "")
                _ = v12
                _ = v8
            } else {
                var v13 bool = (!v4)
                v4 = true
                v14 := GenS1{F0: 46}
                _ = v13
                _ = v14
            }
        }
    }
    v15 := GenS2{F0: (v2 >> ((18446744073709551615) & 63)), F1: (v2 % ((v2) | 1)), F2: -94}
    v16 := GenS3{F0: ((v4 && true) && ((0 >= 42) && v4)), F1: (1.5 != -2.25), F2: uint8(v2)}
    v17 := []int{-1, -90}
    if true {
        for fi18 := 0; fi18 < 3; fi18++ {
            if (32 == 1) {
                if v4 {
                    var v19 float64 = (0.1 - 1e10)
                    fmt.Println(v4)
                    _ = v19
                } else {
                    switch (fi18 + (fi18 ^ -9223372036854775808)) {
                    case 128:
                        var v20 int64 = 1
                        _ = v20
                    case -479:
                        for fi21 := 0; fi21 < 15; fi21++ {
                            if true {
                                v22 := GenS4{F0: -642, F1: 255, F2: 255}
                                _ = v22
                            } else {
                                v4 = true
                                var v23 uint32 = v2
                                v2 -= 2
                                _ = v23
                            }
                            _ = fi21
                        }
                    case 633:
                        v6 = v6
                    default:
                        for fi24 := 0; fi24 < 1; fi24++ {
                            v6 = v6
                            _ = len(v17)
                            v2 = 813
                            _ = fi24
                        }
                    }
                    v4 = ((520 != 65535) && ((2 != 18446744073709551615) || v4))
                    v3[v6] = 4294967295
                }
                switch v6 {
                case "":
                    if (v4 || (v4 || (!v4))) {
                        switch v6 {
                        case "foo bar":
                            switch v6 {
                            case "a":
                                fmt.Println(fi18, fi18)
                            default:
                                switch v6 {
                                case "":
                                    v6 = v6
                                case "a":
                                    var v25 int = (fi18 * 0)
                                    _ = v25
                                case "foo bar":
                                    v2 -= v2
                                }
                            }
                        case "a":
                            v2 -= v2
                        }
                        fmt.Println(v6)
                    }
                case "a":
                    fmt.Println(v2, v4, fi18)
                case "hello":
                    var v26 uint16 = 42
                    _ = v26
                default:
                    var v27 uint64 = 2
                    _ = v27
                }
                v4 = ((-852 == -9223372036854775808) || false)
            } else {
                v4 = (!(127 > fi18))
            }
            _ = fi18
        }
        for fi28 := 0; fi28 < 2; fi28++ {
            for fi29 := 0; fi29 < 16; fi29++ {
                fmt.Println(fi28, v2, v2)
                _ = fi29
            }
            v6 = v6
            if len(v17) > 0 {
                v17[0] = (fi28 % ((97) | 1))
            }
            _ = fi28
        }
    }
    v4 = ((!((!false) || (1000 == 1000))) && v4)
    fmt.Println(v2, v6, v6)
    if false {
        if (981 >= 889) {
            if v4 {
                var v30 int = 32768
                v2 -= (v2 % ((4294967295) | 1))
                _ = v30
            }
            v31 := GenS5{F0: 129, F1: 255}
            v32 := GenS3{F0: (2 > 71), F1: (1000 > 764), F2: 171}
            _ = v31
            _ = v32
        } else {
            v33 := GenS4{F0: 42, F1: 65535, F2: 0}
            var v34 uint8 = 42
            _ = v33
            _ = v34
        }
        if true {
            var v35 uint32 = 42
            if len(v17) > 0 {
                v17[0] = 4294967295
            }
            if false {
                v4 = (v4 || v4)
            }
            _ = v35
        } else {
            v6 = ("hello" + "\n\t")
        }
        v36 := v3[("" + (v6 + "foo bar"))]
        _ = v36
    } else {
        for fi37 := 0; fi37 < 15; fi37++ {
            fmt.Println(v6)
            v38 := GenS4{F0: 1, F1: 112, F2: 1}
            for fi39 := 0; fi39 < 15; fi39++ {
                if true {
                    v40 := GenS5{F0: 1, F1: 127}
                    v3[(((v6 + v6) + "hello") + (v6 + v6))] = fi37
                    var v41 int16 = 618
                    _ = v40
                    _ = v41
                }
                switch v6 {
                case "a":
                    var v42 float64 = (1e10 + 0.1)
                    _ = v42
                case "\n\t":
                    v2 -= (v2 * 1)
                case "":
                    if len(v17) > 0 {
                        v17[0] = 32768
                    }
                    var v43 uint = 747
                    var v44 uint32 = (v2 % (((v2 - (v2 / (((^(v2))) | 1)))) | 1))
                    _ = v43
                    _ = v44
                }
                _ = fi39
            }
            _ = fi37
            _ = v38
        }
        for fi45 := 0; fi45 < 16; fi45++ {
            v6 = ("hello" + "foo bar")
            v6 = v6
            v2 -= 158
            _ = fi45
        }
    }
    for fi46 := 0; fi46 < 12; fi46++ {
        switch v6 {
        case "hello":
            v47 := GenS3{F0: (!((fi46 | fi46) <= (fi46 * -1))), F1: v4, F2: 2}
            _ = v47
        case "a":
            v48 := GenS4{F0: 32767, F1: 65535, F2: 915}
            _ = v48
        default:
            var v49 int16 = 42
            _ = v49
        }
        v50 := GenS3{F0: true, F1: ((42 > uint(v2)) && (v4 && (!(!false)))), F2: 151}
        _ = fi46
        _ = v50
    }
    v4 = true
    for fi51 := 0; fi51 < 8; fi51++ {
        if v4 {
            fmt.Println(v2, v6)
            var v52 int8 = -1
            switch ((v6 + "a") + v6) {
            case "\n\t":
                for fi53 := 0; fi53 < 4; fi53++ {
                    fmt.Println(v52, v4, fi53)
                    _ = fi53
                }
            case "hello":
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            case "":
                if v4 {
                    fmt.Println(v6, v4)
                    v4 = (v2 <= (v2 + 4294967295))
                    var v54 int64 = 0
                    _ = v54
                }
            default:
                for fi55 := 0; fi55 < 9; fi55++ {
                    if (float32(float64(-2.25)) >= (-2.25 / ((1e10 / 1.0) / 3.14159))) {
                        if (741 == 255) {
                            v56 := GenS4{F0: -698, F1: 42, F2: 65535}
                            _ = v56
                        }
                        if (v4 && ((0.0 - 0.0) >= 3.14159)) {
                            v52 = v52
                        } else {
                            switch (fi55 << ((2) & 63)) {
                            case -706:
                                var v57 uint16 = 1
                                var v58 int64 = -1
                                _ = v57
                                _ = v58
                            case 128:
                                var v59 string = v6
                                _ = v59
                            }
                        }
                        if false {
                            v2 = (v2 % ((v2) | 1))
                            v2 = (v2 % (((^(v2))) | 1))
                        }
                    } else {
                        var v60 uint = 850
                        _ = v60
                    }
                    _ = fi55
                }
            }
            _ = v52
        }
        if (v4 && (v6 != (v6 + v6))) {
            v4 = v4
        } else {
            if v4 {
                if (-1 < 1) {
                    var v61 bool = v4
                    v62 := GenS2{F0: (^(v2)), F1: (v2 / (((^(v2))) | 1)), F2: 0}
                    fmt.Println(v2)
                    _ = v61
                    _ = v62
                }
            } else {
                v63 := GenS6{F0: (!(true || false)), F1: "\n\t"}
                _ = v63
            }
            switch -750 {
            case 976:
                v64 := GenS7{F0: (^(v2)), F1: fi51}
                _ = v64
            case 0:
                for fi65 := 0; fi65 < 19; fi65++ {
                    v66 := v1[-192]
                    var v67 int = 2147483647
                    _ = fi65
                    _ = v66
                    _ = v67
                }
            }
            v68 := v1[fi51]
            _ = v68
        }
        for fi69 := 0; fi69 < 13; fi69++ {
            v17 = append(v17, fi69)
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            _ = fi69
        }
        _ = fi51
    }
    for fi70 := 0; fi70 < 16; fi70++ {
        for fi71 := 0; fi71 < 19; fi71++ {
            var v72 string = v6
            _ = fi71
            _ = v72
        }
        _ = fi70
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
