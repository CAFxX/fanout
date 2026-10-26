package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int64; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int16; F2 uint32 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint16; F1 int32; F2 int8 }
type GenS7 struct { F0 int32 }

func main() {
    v1 := []int{127, 0}
    var v2 uint32 = 0
    v3 := make(map[string]uint32)
    fmt.Println(v2)
    for fi4 := 0; fi4 < 3; fi4++ {
        switch (fi4 >> ((190) & 63)) {
        case 660:
            fmt.Println(v2, fi4)
        case 4294967295:
            if false {
                switch -717 {
                case -476:
                    for fi5 := 0; fi5 < 4; fi5++ {
                        fmt.Println(fi4, v2)
                        fmt.Println(v2)
                        if (4294967295 < 4294967295) {
                            switch ("foo bar" + "\n\t") {
                            case "a":
                                if (!false) {
                                    fmt.Println(fi4, fi4, v2)
                                    var v6 uint64 = 2
                                    _ = v6
                                } else {
                                    var v7 float32 = (float32(0.1) + float32((float64((0.0 / 1.5)) - (0.0 - -2.25))))
                                    v8 := GenS1{F0: 255}
                                    if false {
                                        var v9 string = "hello"
                                        var v10 int16 = 32767
                                        v9 = "foo bar"
                                        v7 += 1e10
                                        v11 := GenS1{F0: 623}
                                        v2 = v2
                                        var v12 uint = 4294967295
                                        _ = v10
                                        _ = v11
                                        _ = v12
                                        _ = v9
                                    } else {
                                        v2 -= (v2 + v2)
                                        v7 = (v7 / 3.14159)
                                        v13 := v3[("" + ("foo bar" + "a"))]
                                        _ = v13
                                    }
                                    _ = v7
                                    _ = v8
                                }
                            }
                            v2 += v2
                            var v14 int8 = 127
                            _ = v14
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        _ = fi5
                    }
                case 32767:
                    v2 = v2
                case 42:
                    var v15 uint32 = v2
                    _ = v15
                default:
                    v1 = append(v1, fi4)
                }
            }
        case 771:
            var v16 uint32 = (v2 << ((842) & 63))
            _ = v16
        }
        _ = fi4
    }
    v17 := GenS1{F0: 836}
    if (false || (0 > -56)) {
        switch 596 {
        case -187:
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 int32 = 550
                _ = fi18
                _ = v19
            }
        case 0:
            v20 := GenS1{F0: 1}
            _ = v20
        }
        v21 := GenS1{F0: 2}
        v1 = append(v1, 807)
        _ = v21
    }
    _ = len(v3)
    if false {
        if (-1 == 1) {
            if (4294967295 < 4294967295) {
                v2 = (v2 | v2)
                fmt.Println(v2)
            } else {
                v2 += (v2 + v2)
                for fi22 := 0; fi22 < 9; fi22++ {
                    v23 := GenS1{F0: 2}
                    if len(v1) > 0 {
                        v1[0] = 0
                    }
                    v2 += (^(v2))
                    _ = fi22
                    _ = v23
                }
            }
            v24 := v3["\n\t"]
            if (1 != -9223372036854775808) {
                var v25 int64 = 1
                _ = v25
            } else {
                switch -9223372036854775808 {
                case -788:
                    v2 -= v24
                }
            }
            _ = v24
        }
    } else {
        var v26 uint32 = 2147483648
        var v27 uint32 = v2
        for fi28 := 0; fi28 < 6; fi28++ {
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 uint = 4294967295
                _ = fi29
                _ = v30
            }
            v1 = append(v1, fi28)
            _ = fi28
        }
        _ = v26
        _ = v27
    }
    var v31 int64 = 9223372036854775807
    v2 += (v2 | (v2 / ((v2) | 1)))
    for fi32 := 0; fi32 < 9; fi32++ {
        if (128 != 2) {
            switch "a" {
            case "":
                var v33 int = (-(fi32))
                _ = v33
            case "hello":
                for fi34 := 0; fi34 < 10; fi34++ {
                    v31 = 850
                    _ = fi34
                }
            case "foo bar":
                if (1 <= 461) {
                    v31 = v31
                    for fi35 := 0; fi35 < 3; fi35++ {
                        v2 += v2
                        v31 = -1
                        _ = fi35
                    }
                    if len(v1) > 0 {
                        v1[0] = fi32
                    }
                } else {
                    v36 := GenS2{F0: v2, F1: (v31 ^ v31), F2: -9}
                    var v37 int16 = 1000
                    _ = v36
                    _ = v37
                }
            }
            fmt.Println(v31)
        }
        for fi38 := 0; fi38 < 19; fi38++ {
            fmt.Println(fi32)
            _ = fi38
        }
        _ = fi32
    }
    if false {
        v31 += v31
        if (!((3.14159 != (1e10 / (float64((0.1 * (((0.1 / 3.14159) * 0.1) / -2.25))) * 0.1))) && false)) {
            fmt.Println(v31)
            v39 := GenS3{F0: v31}
            _ = v39
        }
    } else {
        v40 := GenS4{F0: -53, F1: 1000, F2: v2}
        v41 := GenS3{F0: -9223372036854775808}
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        var v43 int = fi42
        switch fi42 {
        case 0:
            v44 := GenS2{F0: (^(v2)), F1: v31, F2: -2147483648}
            _ = v44
        case -1:
            for fi45 := 0; fi45 < 7; fi45++ {
                switch ("\n\t" + "a") {
                case "a":
                    fmt.Println(fi45)
                case "":
                    var v46 int = v43
                    _ = v46
                }
                _ = fi45
            }
        case -391:
            fmt.Println(v43, v43)
        }
        _ = fi42
        _ = v43
    }
    switch 64 {
    case -2147483648:
        for fi47 := 0; fi47 < 13; fi47++ {
            var v48 string = "\n\t"
            var v49 int8 = 42
            _ = fi47
            _ = v48
            _ = v49
        }
    case 32768:
        v31 += v31
    case 9223372036854775807:
        var v50 string = ((("\n\t" + "hello") + "") + "hello")
        _ = v50
    }
    for fi51 := 0; fi51 < 6; fi51++ {
        switch "\n\t" {
        case "hello":
            v2 = 451
        case "":
            if (false || (false || true)) {
                v52 := GenS1{F0: 32767}
                if (4294967295 <= 4294967295) {
                    v53 := GenS5{F0: (v31 << ((816) & 63)), F1: 255}
                    switch fi51 {
                    case -2147483648:
                        if false {
                            v54 := GenS6{F0: 65535, F1: 460, F2: 121}
                            fmt.Println(v31)
                            _ = v54
                        }
                    default:
                        switch ("\n\t" + "a") {
                        case "foo bar":
                            v55 := v3["hello"]
                            _ = v55
                        default:
                            switch fi51 {
                            case 4294967295:
                                v56 := GenS7{F0: -269}
                                _ = v56
                            case 42:
                                v57 := GenS4{F0: 42, F1: 1000, F2: v2}
                                _ = v57
                            default:
                                fmt.Println(fi51, fi51, v31)
                            }
                        }
                    }
                    _ = v53
                } else {
                    _ = len(v1)
                    if len(v1) > 0 {
                        v1[0] = (fi51 + (-(fi51)))
                    }
                }
                _ = v52
            }
        default:
            for fi58 := 0; fi58 < 2; fi58++ {
                v2 -= (^(v2))
                if true {
                    v2 = v2
                }
                v2 = 1
                _ = fi58
            }
        }
        _ = fi51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
}
