package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS2 struct { F0 string; F1 int16; F2 uint8 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int16; F1 uint32; F2 uint }
type GenS6 struct { F0 uint8; F1 uint }
func genFn0(p0 uint16) uint16 {
    var v6 uint16 = 240
    for fi7 := 0; fi7 < 20; fi7++ {
        p0 += 430
        var v8 int64 = 9223372036854775807
        if true {
            v9 := GenS1{F0: 1, F1: 42, F2: 89}
            _ = v9
        } else {
            switch (fi7 - 2147483647) {
            case -187:
                switch "hello" {
                case "":
                    v10 := GenS2{F0: "hello", F1: 42, F2: 127}
                    _ = v10
                case "foo bar":
                    fmt.Println(fi7)
                case "a":
                    v11 := GenS1{F0: 0, F1: -952, F2: 26}
                    _ = v11
                }
            case -657:
                var v12 uint8 = 128
                _ = v12
            default:
                switch fi7 {
                case -2147483648:
                    v13 := GenS3{F0: (603 != uint(p0)), F1: true}
                    _ = v13
                case -9223372036854775808:
                    v14 := GenS4{F0: 12, F1: 128, F2: 255}
                    _ = v14
                case -1:
                    switch ("foo bar" + "a") {
                    case "\n\t":
                        v6 = v6
                    }
                default:
                    switch fi7 {
                    case 4294967295:
                        for fi15 := 0; fi15 < 3; fi15++ {
                            if true {
                                fmt.Println(fi7, v6, p0)
                                if false {
                                    var v16 int8 = 0
                                    _ = v16
                                }
                                p0 -= 65535
                            }
                            var v17 uint = 2
                            fmt.Println(fi15, v17, v17)
                            _ = fi15
                            _ = v17
                        }
                    case 0:
                        v6 = p0
                    default:
                        if (int8(fi7) == -1) {
                            v6 -= 2
                        } else {
                            fmt.Println(v8, p0)
                        }
                    }
                }
            }
        }
        _ = fi7
        _ = v8
    }
    var v18 float64 = 1.5
    v6 -= (p0 * 740)
    var v19 int16 = 1000
    _ = p0
    _ = v18
    _ = v19
    _ = v6
    return 378
}

func main() {
    v1 := []string{"", ("a" + ""), "foo bar"}
    var v2 uint64 = 10
    var v3 int8 = -69
    var v4 int16 = 42
    v5 := make(map[int]int)
    var v20 uint = 4294967295
    var v21 uint64 = v2
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 int32 = -2147483648
        v21 -= v21
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        if (fi24 != 396) {
            for fi25 := 0; fi25 < 15; fi25++ {
                v21 = 0
                _ = fi25
            }
            v26 := GenS1{F0: v2, F1: (v4 % ((v4) | 1)), F2: (^(v2))}
            v27 := GenS5{F0: v4, F1: 761, F2: v20}
            _ = v26
            _ = v27
        } else {
            v4 -= (-(v4))
            for fi28 := 0; fi28 < 1; fi28++ {
                if false {
                    v4 -= 0
                    if (-528 > (fi24 - (^(fi28)))) {
                        v20 -= uint(v2)
                        var v29 int32 = 0
                        _ = v29
                    }
                    var v30 int32 = 0
                    _ = v30
                } else {
                    v20 -= 362
                }
                v31 := GenS3{F0: (100 == -123), F1: (v3 > v3)}
                fmt.Println(v20)
                _ = fi28
                _ = v31
            }
            if true {
                _ = len(v1)
                v1 = append(v1, "\n\t")
            }
        }
        v32 := GenS3{F0: true, F1: false}
        _ = fi24
        _ = v32
    }
    v33 := GenS5{F0: (v4 & v4), F1: 4294967295, F2: 2}
    if (-1 < 339) {
        if false {
            if (true || true) {
                fmt.Println(v3, v21)
            } else {
                if (42 != 87) {
                    for fi34 := 0; fi34 < 1; fi34++ {
                        var v35 bool = false
                        var v36 float64 = 0.1
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                }
            }
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            var v37 int8 = (v3 >> (((v20 | v20)) & 63))
            _ = v37
        } else {
            if false {
                var v38 uint8 = 0
                delete(v5, 2147483647)
                switch "hello" {
                case "a":
                    var v39 int32 = 42
                    _ = v39
                case "":
                    v4 += -32768
                case "foo bar":
                    v2 -= v2
                default:
                    v20 -= v20
                }
                _ = v38
            } else {
                fmt.Println(v4, v21)
                switch int(v3) {
                case 334:
                    v40 := GenS1{F0: 738, F1: v4, F2: 4294967295}
                    _ = v40
                case -719:
                    switch "hello" {
                    case "foo bar":
                        if false {
                            var v41 bool = (!((3.14159 - float64(-2.25)) <= 0.0))
                            for fi42 := 0; fi42 < 12; fi42++ {
                                v4 += 260
                                v2 = (^(v2))
                                _ = fi42
                            }
                            var v43 float32 = 3.14159
                            _ = v41
                            _ = v43
                        } else {
                            v5[-2147483648] = 127
                            fmt.Println(v20, v2, v2)
                        }
                    default:
                        v44 := GenS6{F0: 128, F1: 255}
                        _ = v44
                    }
                }
                if (-2147483648 < 42) {
                    fmt.Println(v2)
                }
            }
        }
        var v45 int8 = 68
        switch 32768 {
        case 1:
            var v46 uint = v20
            _ = v46
        }
        _ = v45
    } else {
        if true {
            if (-91 != 430) {
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
            }
        } else {
            if true {
                var v47 int = 1
                v20 += v20
                _ = v47
            }
            if (1e10 > 3.14159) {
                fmt.Println(v4, v2, v2)
            } else {
                v21 = 2
                var v48 int64 = 190
                if true {
                    v49 := GenS3{F0: true, F1: (4294967295 != 238)}
                    var v50 string = "hello"
                    _ = v49
                    _ = v50
                }
                _ = v48
            }
            for fi51 := 0; fi51 < 11; fi51++ {
                for fi52 := 0; fi52 < 5; fi52++ {
                    switch (fi51 & -2147483648) {
                    case 32768:
                        var v53 uint = 42
                        _ = v53
                    case 624:
                        fmt.Println(fi52, v2)
                    case -1:
                        var v54 string = "a"
                        _ = v54
                    }
                    _ = fi52
                }
                v55 := GenS5{F0: v4, F1: 923, F2: (v20 - v20)}
                if false {
                    if (32767 <= 42) {
                        var v56 int32 = -34
                        var v57 string = "hello"
                        _ = len(v1)
                        _ = v56
                        _ = v57
                    } else {
                        for fi58 := 0; fi58 < 6; fi58++ {
                            var v59 uint64 = v21
                            switch "a" {
                            case "\n\t":
                                v2 += v21
                            case "hello":
                                v60 := GenS6{F0: 128, F1: v20}
                                _ = v60
                            case "":
                                if true {
                                    var v61 string = ""
                                    delete(v5, (fi58 / (((^(fi51))) | 1)))
                                    switch ("foo bar" + ((v61 + "\n\t") + "foo bar")) {
                                    case "foo bar":
                                        v3 += v3
                                    case "a":
                                        fmt.Println(v21, v20)
                                    case "\n\t":
                                        v62 := GenS3{F0: (v21 != (v21 >> (((^(v20))) & 63))), F1: (false && true)}
                                        _ = v62
                                    }
                                    _ = v61
                                } else {
                                    fmt.Println(v21, fi51, v3)
                                }
                            default:
                                if (2 >= 1) {
                                    fmt.Println(v59)
                                    switch fi51 {
                                    case 32768:
                                        v4 += v4
                                    case -2147483648:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    default:
                                        v21 = v2
                                    }
                                } else {
                                    var v63 uint32 = 4294967295
                                    var v64 uint16 = 32767
                                    _ = v63
                                    _ = v64
                                }
                            }
                            v65 := GenS3{F0: true, F1: (1.5 != (((3.14159 * -2.25) / 1e10) / (float32(3.14159) + float32(1e10))))}
                            v59 += 18446744073709551615
                            switch (^(fi51)) {
                            case 245:
                                var v66 int64 = 42
                                _ = v66
                            case 127:
                                fmt.Println(v20, fi58)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v67 uint64 = v21
                                _ = v67
                            case -389:
                                delete(v5, (fi58 * (fi51 * (fi58 | -314))))
                            }
                            _ = fi58
                            _ = v59
                            _ = v65
                        }
                    }
                    v4 = v4
                    fmt.Println(v4, v20)
                }
                _ = fi51
                _ = v55
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3 = (v3 + 127)
    }
    for fi68 := 0; fi68 < 17; fi68++ {
        var v69 float64 = -2.25
        fmt.Println(v3, v20, v4)
        for fi70 := 0; fi70 < 10; fi70++ {
            var v71 int = (fi70 % ((int(v21)) | 1))
            var v72 uint8 = 255
            var v73 int32 = 2147483647
            v73 = 2147483647
            fmt.Println(v69)
            _ = fi70
            _ = v71
            _ = v72
            _ = v73
        }
        _ = fi68
        _ = v69
    }
    v74 := GenS6{F0: 42, F1: (v20 ^ v20)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v74
}
