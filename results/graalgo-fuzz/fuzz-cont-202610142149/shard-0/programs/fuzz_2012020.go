package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 string }
func genFn0() string {
    var v4 bool = false
    v4 = false
    v5 := []int{}
    v5 = append(v5, -1)
    _ = v4
    _ = v5
    return ("a" + "\n\t")
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{2147483647}
    var v3 uint64 = 4294967295
    var v6 bool = true
    v7 := GenS1{F0: 258}
    fmt.Println(v3)
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: 221}
        if len(v2) > 0 {
            v2[0] = 287
        }
        v3 -= 1
        _ = fi8
        _ = v9
    }
    _ = len(v2)
    fmt.Println(v3, v3)
    v6 = false
    if (1 < 424) {
        for fi10 := 0; fi10 < 13; fi10++ {
            for fi11 := 0; fi11 < 1; fi11++ {
                _ = len(v1)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi11
            }
            if false {
                v12 := GenS2{F0: 65535, F1: 32767}
                if (v6 || false) {
                    fmt.Println(v3, fi10)
                } else {
                    v3 -= v3
                    for fi13 := 0; fi13 < 10; fi13++ {
                        v3 = 0
                        var v14 uint16 = 2
                        v6 = v6
                        _ = fi13
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 20; fi15++ {
                        v16 := GenS2{F0: 4294967295, F1: 2}
                        switch "a" {
                        case "":
                            v3 += (v3 % ((2) | 1))
                        case "hello":
                            v2 = append(v2, 0)
                        case "a":
                            v3 -= v3
                        }
                        _ = fi15
                        _ = v16
                    }
                }
                var v17 int64 = 9223372036854775807
                _ = v12
                _ = v17
            } else {
                var v18 int8 = 1
                var v19 int16 = 0
                var v20 bool = v6
                _ = v18
                _ = v19
                _ = v20
            }
            _ = fi10
        }
        if false {
            v21 := GenS2{F0: 460, F1: 1}
            var v22 uint = 345
            var v23 uint64 = 4294967295
            _ = v21
            _ = v22
            _ = v23
        } else {
            v24 := GenS1{F0: 42}
            var v25 int64 = 271
            fmt.Println(v6, v3)
            _ = v24
            _ = v25
        }
        fmt.Println(v3, v3)
    } else {
        for fi26 := 0; fi26 < 3; fi26++ {
            v27 := GenS3{F0: 255}
            var v28 uint32 = 1
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    v2 = append(v2, 194)
    for fi29 := 0; fi29 < 18; fi29++ {
        switch fi29 {
        case 2147483647:
            var v30 uint8 = 230
            _ = v30
        case -854:
            v6 = (!((v3 <= v3) || false))
        }
        if v6 {
            for fi31 := 0; fi31 < 15; fi31++ {
                if v6 {
                    var v32 int64 = -1
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                    v33 := GenS2{F0: 0, F1: 32767}
                    _ = v32
                    _ = v33
                }
                if (624 != int16(v3)) {
                    v34 := GenS4{F0: 1}
                    _ = v34
                } else {
                    if v6 {
                        v35 := GenS3{F0: 219}
                        var v36 string = "\n\t"
                        if (0 != 42) {
                            var v37 uint = 0
                            switch fi29 {
                            case 127:
                                fmt.Println(v37, v36, fi29)
                            case -1:
                                v1[("foo bar" + v36)] = v6
                            case 49:
                                v38 := GenS5{F0: 65535, F1: 18446744073709551615}
                                _ = v38
                            }
                            _ = v37
                        } else {
                            if v6 {
                                if v6 {
                                    if len(v2) > 0 {
                                        v2[0] = 42
                                    }
                                    var v39 uint32 = 2147483648
                                    _ = v39
                                } else {
                                    v1["a"] = true
                                    _ = len(v2)
                                    if v6 {
                                        if false {
                                            var v40 int64 = -274
                                            fmt.Println(v40, fi29, v6)
                                            v40 = (v40 << ((274) & 63))
                                            _ = v40
                                        } else {
                                            var v41 uint64 = v3
                                            _ = v41
                                        }
                                    } else {
                                        if (fi31 <= fi29) {
                                            v3 = (v3 << ((246) & 63))
                                            v42 := GenS5{F0: 4294967295, F1: (^(v3))}
                                            _ = v42
                                        } else {
                                            fmt.Println(v3, v3)
                                        }
                                        if false {
                                            if v6 {
                                                fmt.Println(v3, fi31)
                                                if (255 >= 617) {
                                                    v3 += v3
                                                    v6 = v6
                                                    v36 = (v36 + v36)
                                                }
                                                v43 := GenS4{F0: 687}
                                                _ = v43
                                            } else {
                                                if (!(v6 || true)) {
                                                    var v44 string = v36
                                                    v44 = "foo bar"
                                                    _ = v44
                                                }
                                                v3 -= v3
                                                var v45 float64 = 0.0
                                                var v46 int16 = -32768
                                                if (v6 && (3.14159 == 3.14159)) {
                                                    fmt.Println(v45, fi31, v6)
                                                } else {
                                                    v47 := GenS6{F0: 0, F1: 65535}
                                                    switch (("\n\t" + ((v36 + v36) + "")) + "\n\t") {
                                                    case "a":
                                                        v45 += -2.25
                                                    case "hello":
                                                        if ((v6 || (!false)) || v6) {
                                                            if ((!true) || v6) {
                                                                fmt.Println(v46)
                                                                v36 = "a"
                                                            }
                                                            var v48 float64 = v45
                                                            v3 += v3
                                                            _ = v48
                                                        } else {
                                                            if (100 != 0) {
                                                                v49 := GenS6{F0: (v3 % ((v3) | 1)), F1: 43}
                                                                var v50 int = (fi29 & (fi29 % ((fi31) | 1)))
                                                                v46 -= v46
                                                                _ = v49
                                                                _ = v50
                                                            } else {
                                                                var v51 bool = (((v46 >> ((1) & 63)) != 42) || v6)
                                                                var v52 int = fi29
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                            delete(v1, v36)
                                                            if (18446744073709551615 >= 65535) {
                                                                fmt.Println(fi29)
                                                                v36 = v36
                                                                var v53 uint16 = 1
                                                                _ = v53
                                                            } else {
                                                                switch fi31 {
                                                                case 32767:
                                                                    v6 = v6
                                                                default:
                                                                    var v54 int = fi31
                                                                    _ = v54
                                                                }
                                                                if false {
                                                                    v3 += v3
                                                                    v55 := GenS4{F0: 42}
                                                                    v56 := GenS5{F0: 2, F1: v3}
                                                                    v57 := GenS4{F0: 42}
                                                                    _ = v55
                                                                    _ = v56
                                                                    _ = v57
                                                                }
                                                                var v58 uint16 = 0
                                                                _ = v58
                                                            }
                                                        }
                                                    default:
                                                        var v59 int16 = (v46 % (((^(v46))) | 1))
                                                        _ = v59
                                                    }
                                                    var v60 uint64 = v3
                                                    _ = v47
                                                    _ = v60
                                                }
                                                _ = v45
                                                _ = v46
                                            }
                                        }
                                    }
                                    var v61 float64 = -2.25
                                    switch (fi29 ^ (fi29 | (fi31 % (((fi31 >> ((2) & 63))) | 1)))) {
                                    case -9223372036854775808:
                                        _ = len(v2)
                                    case -1:
                                        v62 := GenS3{F0: 127}
                                        _ = v62
                                    case -2147483648:
                                        var v63 uint32 = 65535
                                        _ = v63
                                    default:
                                        fmt.Println(v6, v6, fi29)
                                    }
                                    _ = v61
                                }
                            }
                        }
                        _ = v35
                        _ = v36
                    }
                    switch "foo bar" {
                    case "":
                        fmt.Println(v6)
                    }
                    _ = len(v1)
                }
                v6 = (707 < 4294967295)
                _ = fi31
            }
            for fi64 := 0; fi64 < 9; fi64++ {
                v6 = (((3.14159 / 1.5) * ((0.0 * 0.0) + float64(1e10))) == 1.5)
                delete(v1, "\n\t")
                switch fi64 {
                case 0:
                    fmt.Println(fi29)
                case 9223372036854775807:
                    v6 = (2147483648 >= 65535)
                case -1:
                    if (127 == 242) {
                        v3 = (v3 / ((887) | 1))
                    }
                }
                _ = fi64
            }
            switch (fi29 - (fi29 / ((fi29) | 1))) {
            case 4294967295:
                fmt.Println(fi29, v6, v6)
            case 128:
                if (!(!(!v6))) {
                    v6 = v6
                } else {
                    for fi65 := 0; fi65 < 19; fi65++ {
                        switch fi29 {
                        case 0:
                            v3 -= v3
                        }
                        fmt.Println(fi65, v6)
                        var v66 float64 = (((0.0 + 1e10) * 3.14159) + (1.5 + -2.25))
                        _ = fi65
                        _ = v66
                    }
                    fmt.Println(v6)
                    _ = len(v1)
                }
            }
        }
        v6 = v6
        _ = fi29
    }
    if (v6 && false) {
        v3 = v3
        v67 := GenS2{F0: 0, F1: 65535}
        _ = v67
    } else {
        for fi68 := 0; fi68 < 4; fi68++ {
            v69 := GenS2{F0: 560, F1: 199}
            fmt.Println(fi68, fi68, fi68)
            fmt.Println(v6, v3, v6)
            _ = fi68
            _ = v69
        }
        switch "a" {
        case "":
            switch 290 {
            case 127:
                if v6 {
                    v3 = 2
                    if false {
                        var v70 bool = v6
                        for fi71 := 0; fi71 < 5; fi71++ {
                            var v72 string = "\n\t"
                            v73 := GenS6{F0: v3, F1: 750}
                            _ = fi71
                            _ = v72
                            _ = v73
                        }
                        v74 := GenS1{F0: 708}
                        _ = v70
                        _ = v74
                    } else {
                        v3 = (^(v3))
                        fmt.Println(v6)
                    }
                }
            }
        case "\n\t":
            var v75 bool = false
            _ = v75
        default:
            v76 := GenS7{F0: ("a" + (("\n\t" + ("\n\t" + "")) + "a"))}
            _ = v76
        }
    }
    for fi77 := 0; fi77 < 12; fi77++ {
        switch "hello" {
        case "hello":
            v6 = v6
        }
        _ = fi77
    }
    v6 = (v6 || (1.5 > 1e10))
    for fi78 := 0; fi78 < 8; fi78++ {
        var v79 uint16 = 32767
        _ = fi78
        _ = v79
    }
    for fi80 := 0; fi80 < 11; fi80++ {
        v6 = true
        _ = fi80
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
