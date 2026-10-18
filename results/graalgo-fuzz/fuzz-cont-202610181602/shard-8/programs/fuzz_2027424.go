package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint32; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 int; F2 uint64 }
type GenS5 struct { F0 bool; F1 uint8; F2 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8; F1 string }
func genFn0(p0 bool) bool {
    var v5 int32 = 42
    var v6 int = -167
    for fi7 := 0; fi7 < 13; fi7++ {
        p0 = (p0 && (!(true || (32767 == 814))))
        _ = fi7
    }
    fmt.Println(v5, v5)
    if (false || false) {
        var v8 uint16 = 2
        switch (v6 & (v6 ^ (v6 >> ((65535) & 63)))) {
        case 4294967295:
            v9 := GenS1{F0: 0, F1: (!true), F2: v8}
            _ = v9
        default:
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := GenS2{F0: (v5 ^ int32(v8)), F1: 2147483647, F2: 32767}
                _ = fi10
                _ = v11
            }
        }
        _ = v8
    } else {
        if (!(4294967295 != 1)) {
            switch v6 {
            case 993:
                if (127 == 98) {
                    fmt.Println(p0, v5)
                }
            case 32768:
                v6 += (v6 * 42)
            case -132:
                v12 := []int{}
                _ = v12
            default:
                v13 := GenS2{F0: -1, F1: 297, F2: 42}
                _ = v13
            }
            v5 = -2147483648
            p0 = (!((!(486 < 255)) && (!true)))
        } else {
            var v14 int = 0
            v15 := GenS3{F0: v5}
            var v16 int32 = (v5 << ((1) & 63))
            _ = v14
            _ = v15
            _ = v16
        }
        var v17 uint64 = 18446744073709551615
        _ = v17
    }
    _ = p0
    _ = v5
    _ = v6
    return p0
}

func main() {
    v1 := []uint32{}
    var v2 int16 = -32768
    var v3 uint64 = 544
    var v4 int8 = 100
    for fi18 := 0; fi18 < 8; fi18++ {
        for fi19 := 0; fi19 < 20; fi19++ {
            if (true && false) {
                var v20 int = (^(fi18))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v21 int16 = (v2 >> ((65535) & 63))
                var v22 int32 = 484
                _ = v20
                _ = v21
                _ = v22
            } else {
                v1 = append(v1, 2147483647)
                var v23 float32 = 1.5
                _ = v23
            }
            _ = fi19
        }
        var v24 uint16 = 538
        _ = fi18
        _ = v24
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        if (0 == 861) {
            v26 := GenS3{F0: 42}
            v27 := GenS3{F0: -1}
            switch "" {
            case "hello":
                if len(v1) > 0 {
                    v1[0] = 203
                }
            case "a":
                fmt.Println(fi25)
            case "\n\t":
                fmt.Println(fi25)
            }
            _ = v26
            _ = v27
        }
        if false {
            var v28 int64 = 9223372036854775807
            var v29 float32 = 3.14159
            var v30 bool = (true && true)
            _ = v28
            _ = v29
            _ = v30
        }
        var v31 uint64 = (v3 | (v3 + v3))
        _ = fi25
        _ = v31
    }
    v32 := GenS2{F0: 42, F1: 2, F2: (-(v2))}
    switch -495 {
    case -507:
        var v33 int16 = (v2 | v2)
        _ = v33
    case -182:
        fmt.Println(v4, v2, v4)
    case 4294967295:
        switch 420 {
        case 1:
            for fi34 := 0; fi34 < 17; fi34++ {
                fmt.Println(v2)
                var v35 int32 = -2147483648
                v1 = append(v1, 2147483647)
                _ = fi34
                _ = v35
            }
        case -2147483648:
            var v36 uint = 4294967295
            _ = v36
        case -822:
            v4 = v4
        default:
            var v37 string = ("a" + "\n\t")
            _ = v37
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v1 = append(v1, 310)
    v3 += (v3 * 517)
    for fi38 := 0; fi38 < 6; fi38++ {
        v39 := GenS4{F0: 1, F1: 0, F2: v3}
        v4 = v4
        _ = fi38
        _ = v39
    }
    v40 := GenS1{F0: v3, F1: true, F2: 42}
    var v41 int16 = v2
    for fi42 := 0; fi42 < 11; fi42++ {
        for fi43 := 0; fi43 < 4; fi43++ {
            var v44 uint32 = 376
            _ = fi43
            _ = v44
        }
        switch "foo bar" {
        case "hello":
            if false {
                var v45 uint = 42
                v4 += (v4 - (v4 % ((v4) | 1)))
                fmt.Println(fi42, v45)
                _ = v45
            }
        case "foo bar":
            v46 := GenS5{F0: true, F1: 140, F2: "foo bar"}
            _ = v46
        }
        if (true && (2 <= 18446744073709551615)) {
            v2 -= 0
        } else {
            fmt.Println(fi42)
            if ("" == ("hello" + "")) {
                if true {
                    if true {
                        var v47 int16 = (v41 % (((v2 | -155)) | 1))
                        switch fi42 {
                        case 42:
                            v48 := GenS4{F0: v4, F1: fi42, F2: (^(v3))}
                            _ = v48
                        case 9223372036854775807:
                            if len(v1) > 0 {
                                v1[0] = 4294967295
                            }
                        case -2147483648:
                            v41 = v2
                        default:
                            var v49 int = fi42
                            _ = v49
                        }
                        _ = v47
                    } else {
                        var v50 float64 = float64(1.5)
                        _ = v50
                    }
                    v4 += (v4 * 0)
                    if (v3 != (v3 & v3)) {
                        var v51 int32 = 1
                        for fi52 := 0; fi52 < 15; fi52++ {
                            switch "foo bar" {
                            case "\n\t":
                                if len(v1) > 0 {
                                    v1[0] = 4294967295
                                }
                            }
                            if true {
                                if false {
                                    var v53 int8 = -35
                                    var v54 uint8 = 101
                                    _ = v53
                                    _ = v54
                                }
                                if true {
                                    var v55 string = ""
                                    if true {
                                        v56 := GenS6{F0: 4294967295}
                                        v57 := GenS2{F0: (v51 << ((2) & 63)), F1: 2147483647, F2: (^(v2))}
                                        _ = v56
                                        _ = v57
                                    }
                                    fmt.Println(v41)
                                    var v58 uint8 = 225
                                    switch 2147483647 {
                                    case 0:
                                        v2 += -32768
                                    default:
                                        v55 = v55
                                    }
                                    _ = v55
                                    _ = v58
                                } else {
                                    var v59 int32 = (^(v51))
                                    var v60 int = (fi42 & -1)
                                    if len(v1) > 0 {
                                        v1[0] = 332
                                    }
                                    _ = v59
                                    _ = v60
                                }
                            }
                            v61 := GenS7{F0: 0, F1: "\n\t"}
                            _ = fi52
                            _ = v61
                        }
                        _ = v51
                    } else {
                        if true {
                            v1 = append(v1, 65535)
                            v3 += 2
                            for fi62 := 0; fi62 < 1; fi62++ {
                                if len(v1) > 0 {
                                    v1[0] = 0
                                }
                                _ = fi62
                            }
                        }
                        v3 -= (v3 % (((^(v3))) | 1))
                    }
                }
                fmt.Println(v3, v41)
            } else {
                switch (^(fi42)) {
                case 4294967295:
                    var v63 int8 = v4
                    _ = v63
                }
            }
        }
        _ = fi42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v41
}
