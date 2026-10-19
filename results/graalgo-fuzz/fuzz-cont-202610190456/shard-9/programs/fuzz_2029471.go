package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint; F2 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint16, p1 uint16, p2 uint16) int16 {
    var v6 int32 = -836
    for fi7 := 0; fi7 < 11; fi7++ {
        p0 = p0
        _ = fi7
    }
    if (-153 >= (v6 - (v6 % (((^(v6))) | 1)))) {
        switch "\n\t" {
        case "a":
            p0 -= (p0 | (p1 - 294))
        case "":
            if false {
                v8 := GenS1{F0: 4294967295}
                var v9 uint32 = 2147483648
                _ = v8
                _ = v9
            } else {
                v6 = v6
                switch "\n\t" {
                case "foo bar":
                    for fi10 := 0; fi10 < 2; fi10++ {
                        v11 := []int{(fi10 & (fi10 & 128))}
                        if true {
                            v12 := GenS2{F0: (fi10 >> ((814) & 63)), F1: 255, F2: 42}
                            _ = v12
                        }
                        if true {
                            p0 = 0
                            var v13 float64 = 1e10
                            if (v13 >= v13) {
                                p2 = (p0 / ((p2) | 1))
                                var v14 uint = 65535
                                if ((v14 << (((v14 + 42)) & 63)) != (v14 % ((v14) | 1))) {
                                    for fi15 := 0; fi15 < 15; fi15++ {
                                        if len(v11) > 0 {
                                            v11[0] = fi15
                                        }
                                        switch 128 {
                                        case -891:
                                            v16 := GenS1{F0: 65535}
                                            _ = v16
                                        case 806:
                                            v11 = append(v11, fi15)
                                        }
                                        _ = fi15
                                    }
                                    v17 := GenS2{F0: fi10, F1: 4294967295, F2: -133}
                                    for fi18 := 0; fi18 < 6; fi18++ {
                                        var v19 string = (("\n\t" + "a") + "\n\t")
                                        _ = len(v11)
                                        v20 := GenS3{F0: 2}
                                        _ = fi18
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = v17
                                }
                                _ = v14
                            } else {
                                _ = len(v11)
                                p0 = p1
                                var v21 int8 = -1
                                _ = v21
                            }
                            _ = v13
                        } else {
                            if (-598 > 42) {
                                var v22 int32 = (v6 * 0)
                                var v23 int32 = v22
                                v6 += (v23 * v6)
                                _ = v22
                                _ = v23
                            }
                        }
                        _ = fi10
                        _ = v11
                    }
                case "":
                    fmt.Println(p0)
                case "a":
                    for fi24 := 0; fi24 < 16; fi24++ {
                        switch "a" {
                        case "\n\t":
                            var v25 uint16 = p1
                            _ = v25
                        case "foo bar":
                            for fi26 := 0; fi26 < 11; fi26++ {
                                var v27 uint32 = 329
                                if false {
                                    p0 = 42
                                }
                                p1 += p2
                                v28 := GenS3{F0: 42}
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                        case "hello":
                            if ((^(fi24)) != -950) {
                                var v29 int32 = 808
                                var v30 int64 = 9223372036854775807
                                _ = v29
                                _ = v30
                            }
                        }
                        if ((fi24 + fi24) != (fi24 * (^(fi24)))) {
                            var v31 float64 = float64(0.1)
                            v31 -= ((v31 / v31) / v31)
                            fmt.Println(v6, p2, p0)
                            _ = v31
                        } else {
                            fmt.Println(fi24, p0, p1)
                        }
                        _ = fi24
                    }
                }
                var v32 bool = false
                _ = v32
            }
        case "foo bar":
            v33 := GenS2{F0: 127, F1: 552, F2: -489}
            _ = v33
        }
        var v34 int8 = 100
        _ = v34
    } else {
        fmt.Println(p0, p0)
        for fi35 := 0; fi35 < 1; fi35++ {
            if (!true) {
                v6 += v6
                for fi36 := 0; fi36 < 12; fi36++ {
                    fmt.Println(fi35, p2)
                    _ = fi36
                }
                if false {
                    var v37 int8 = 88
                    _ = v37
                }
            }
            for fi38 := 0; fi38 < 14; fi38++ {
                var v39 float64 = 1.5
                _ = fi38
                _ = v39
            }
            var v40 int8 = 3
            _ = fi35
            _ = v40
        }
        if (1 <= 75) {
            v41 := GenS4{F0: 42, F1: 4294967295}
            var v42 bool = (-1 > 1000)
            var v43 string = "a"
            _ = v41
            _ = v42
            _ = v43
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 713
}

func main() {
    var v1 uint32 = 0
    v2 := []uint32{}
    var v3 int32 = -2147483648
    v4 := make(map[uint32]bool)
    v5 := make(map[int]int)
    var v44 int32 = (v3 + v3)
    var v45 int16 = 0
    var v46 uint8 = 2
    var v47 bool = true
    for fi48 := 0; fi48 < 2; fi48++ {
        if (!(v46 < 127)) {
            v49 := GenS4{F0: 65535, F1: 2147483647}
            v50 := GenS3{F0: v46}
            v51 := GenS5{F0: -1}
            _ = v49
            _ = v50
            _ = v51
        } else {
            fmt.Println(v46, v45, fi48)
        }
        _ = fi48
    }
    var v52 string = ("a" + "a")
    v53 := GenS6{F0: (^(v1))}
    var v54 int8 = -31
    if ((v52 + ((v52 + v52) + "foo bar")) >= (v52 + v52)) {
        if (!(v45 == (v45 ^ (v45 >> ((1) & 63))))) {
            switch "\n\t" {
            case "a":
                v55 := GenS6{F0: v1}
                _ = v55
            default:
                switch v52 {
                case "":
                    var v56 int8 = v54
                    _ = v56
                case "hello":
                    fmt.Println(v44, v1, v44)
                default:
                    switch 32768 {
                    case 9223372036854775807:
                        fmt.Println(v1)
                    case 2147483647:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case 1:
                        v2 = append(v2, (v1 | (^(v1))))
                    }
                }
            }
        }
    } else {
        switch (("hello" + v52) + "a") {
        case "hello":
            for fi57 := 0; fi57 < 18; fi57++ {
                var v58 bool = (!v47)
                _ = fi57
                _ = v58
            }
        case "":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        var v59 uint8 = (v46 - v46)
        _ = v59
    }
    v2 = append(v2, (v1 | (v1 << ((42) & 63))))
    var v60 bool = v47
    for fi61 := 0; fi61 < 10; fi61++ {
        v52 = (("foo bar" + v52) + (v52 + "foo bar"))
        for fi62 := 0; fi62 < 7; fi62++ {
            if (((473 != 80) && (!(!((1e10 >= -2.25) && ((v46 | uint8(v1)) != v46))))) && true) {
                v63 := GenS2{F0: (fi62 << ((495) & 63)), F1: 18446744073709551615, F2: v45}
                v46 -= 239
                _ = v63
            }
            var v64 uint = 65535
            _ = fi62
            _ = v64
        }
        if v47 {
            for fi65 := 0; fi65 < 7; fi65++ {
                v66 := GenS2{F0: (fi61 * (fi65 | (fi61 >> ((4294967295) & 63)))), F1: 1, F2: int16(v1)}
                v52 = ("" + v52)
                _ = fi65
                _ = v66
            }
            v44 = v44
            if (!(42 == 486)) {
                v67 := GenS2{F0: (fi61 << ((2) & 63)), F1: 42, F2: 223}
                if (1.5 == (1e10 - float64(3.14159))) {
                    if v60 {
                        if false {
                            v2 = append(v2, (^(v1)))
                            for fi68 := 0; fi68 < 5; fi68++ {
                                v60 = (v52 >= v52)
                                v4[2147483647] = false
                                if len(v2) > 0 {
                                    v2[0] = (v1 >> ((255) & 63))
                                }
                                _ = fi68
                            }
                        } else {
                            var v69 string = v52
                            if v60 {
                                if ((0 == 446) && (v46 >= (v46 * v46))) {
                                    for fi70 := 0; fi70 < 5; fi70++ {
                                        var v71 int = fi70
                                        v45 -= (-(v45))
                                        _ = fi70
                                        _ = v71
                                    }
                                }
                                v72 := GenS5{F0: -797}
                                v73 := v5[(fi61 / ((fi61) | 1))]
                                _ = v72
                                _ = v73
                            }
                            var v74 int64 = 208
                            _ = v69
                            _ = v74
                        }
                    }
                }
                _ = v67
            }
        } else {
            if true {
                if v47 {
                    v2 = append(v2, (v1 - v1))
                    for fi75 := 0; fi75 < 16; fi75++ {
                        v54 = v54
                        v44 = v44
                        _ = fi75
                    }
                }
                for fi76 := 0; fi76 < 17; fi76++ {
                    v1 += (v1 - v1)
                    _ = fi76
                }
                for fi77 := 0; fi77 < 5; fi77++ {
                    v78 := GenS1{F0: (^(v1))}
                    _ = len(v4)
                    var v79 float32 = 3.14159
                    _ = fi77
                    _ = v78
                    _ = v79
                }
            }
            var v80 string = (("" + (("hello" + (("a" + v52) + ((v52 + v52) + ("" + "foo bar")))) + "a")) + v52)
            v60 = false
            _ = v80
        }
        _ = fi61
    }
    var v81 bool = (557 == 2)
    v60 = true
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    fmt.Println("v60:", v60)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v60
    _ = v81
}
