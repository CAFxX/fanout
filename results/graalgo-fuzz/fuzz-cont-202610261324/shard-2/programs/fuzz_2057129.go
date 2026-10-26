package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 bool; F1 uint8; F2 int8 }
type GenS6 struct { F0 uint; F1 uint64; F2 int64 }
type GenS7 struct { F0 int64; F1 string; F2 uint }
func genFn0(p0 int, p1 uint64) int {
    v5 := []int{(p0 - p0)}
    var v6 uint = 568
    var v7 int16 = 131
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (p0 << (((v6 % ((v6) | 1))) & 63))
}

func main() {
    var v1 bool = false
    v2 := []string{}
    var v3 uint = 493
    var v4 int64 = 1
    if (v4 > (v4 - (v4 + -9223372036854775808))) {
        if (0.1 != (1e10 - -2.25)) {
            if v1 {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v8 := GenS1{F0: 0, F1: (v4 ^ -466)}
                _ = v8
            }
        }
    } else {
        for fi9 := 0; fi9 < 3; fi9++ {
            switch "\n\t" {
            case "foo bar":
                v4 = (v4 + (-(v4)))
            case "a":
                switch "\n\t" {
                case "\n\t":
                    v1 = v1
                }
            case "":
                fmt.Println(v1, fi9, v1)
            }
            var v10 int64 = (v4 / ((v4) | 1))
            _ = fi9
            _ = v10
        }
        var v11 int16 = 32767
        _ = v11
    }
    if (-1 >= -611) {
        var v12 uint16 = 2
        if ((-(v4)) <= 14) {
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := GenS2{F0: -9223372036854775808, F1: -9223372036854775808}
                var v15 string = "foo bar"
                v16 := GenS2{F0: (v4 >> (((v3 - v3)) & 63)), F1: (v4 - (v4 % ((9223372036854775807) | 1)))}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            for fi17 := 0; fi17 < 1; fi17++ {
                if (1.5 != 3.14159) {
                    if true {
                        v18 := GenS3{F0: 814, F1: v3, F2: v12}
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 18; fi19++ {
                        v3 += (v3 ^ (v3 << ((v3) & 63)))
                        var v20 bool = false
                        v3 += v3
                        var v21 bool = (!(!v20))
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    v22 := GenS1{F0: 544, F1: (-(v4))}
                    _ = v22
                } else {
                    if ((!((false || (!(58 >= 943))) && v1)) && v1) {
                        v3 += v3
                        for fi23 := 0; fi23 < 3; fi23++ {
                            if (2 != uint32(v12)) {
                                switch "hello" {
                                case "foo bar":
                                    var v24 uint = v3
                                    _ = v24
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            } else {
                                v25 := GenS4{F0: -7, F1: true}
                                fmt.Println(v3, v3, v1)
                                var v26 uint8 = 2
                                _ = v25
                                _ = v26
                            }
                            v12 -= v12
                            var v27 uint16 = v12
                            var v28 int = (^(fi17))
                            _ = fi23
                            _ = v27
                            _ = v28
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if (genFn0(42, 213) > (fi17 >> (((^(v3))) & 63))) {
                        fmt.Println(fi17)
                        switch 835 {
                        case 32768:
                            fmt.Println(v3, fi17)
                        case 2147483647:
                            var v29 float64 = 0.1
                            _ = v29
                        case -230:
                            switch (fi17 & (^(fi17))) {
                            case 128:
                                v4 = -399
                            case 32767:
                                if ((!((-649 < 1) || (v4 == v4))) || false) {
                                    var v30 int16 = 32767
                                    fmt.Println(v3, v1)
                                    _ = v30
                                }
                            case 32768:
                                var v31 float32 = ((0.1 + (0.1 + float32(float64(0.1)))) - -2.25)
                                _ = v31
                            default:
                                var v32 int8 = 13
                                _ = v32
                            }
                        default:
                            v3 = v3
                        }
                        var v33 bool = (!v1)
                        _ = v33
                    } else {
                        var v34 uint32 = 65535
                        _ = v34
                    }
                }
                var v35 uint32 = 42
                fmt.Println(v3, v12)
                _ = fi17
                _ = v35
            }
        }
        _ = v12
    }
    var v36 uint16 = 32767
    var v37 bool = (255 > 2)
    for fi38 := 0; fi38 < 14; fi38++ {
        for fi39 := 0; fi39 < 3; fi39++ {
            v40 := GenS1{F0: 1, F1: (v4 | (-(v4)))}
            fmt.Println(v36, v36)
            _ = fi39
            _ = v40
        }
        var v41 uint16 = 803
        _ = fi38
        _ = v41
    }
    v42 := GenS5{F0: v37, F1: 128, F2: 0}
    v4 = v4
    v3 = (v3 + (v3 * v3))
    _ = len(v2)
    if (587 != v4) {
        var v43 uint8 = 42
        v44 := GenS4{F0: 42, F1: v37}
        var v45 bool = v37
        _ = v43
        _ = v44
        _ = v45
    }
    if (!(304 > 889)) {
        if (float32((3.14159 * 3.14159)) != 1.5) {
            var v46 int64 = (v4 / ((v4) | 1))
            if v37 {
                var v47 uint16 = 255
                for fi48 := 0; fi48 < 6; fi48++ {
                    switch "hello" {
                    case "":
                        v49 := GenS3{F0: 4294967295, F1: v3, F2: (v47 ^ (v36 - (v36 << ((255) & 63))))}
                        _ = v49
                    }
                    v36 = (^(v36))
                    v50 := GenS1{F0: 28, F1: v4}
                    _ = fi48
                    _ = v50
                }
                _ = v47
            }
            switch "" {
            case "foo bar":
                for fi51 := 0; fi51 < 6; fi51++ {
                    fmt.Println(v37, fi51)
                    _ = fi51
                }
            case "":
                if (127 != 42) {
                    var v52 uint8 = 127
                    v53 := GenS5{F0: v37, F1: 63, F2: 31}
                    v54 := GenS6{F0: v3, F1: 1, F2: v4}
                    _ = v52
                    _ = v53
                    _ = v54
                } else {
                    v55 := GenS7{F0: v4, F1: "", F2: (^(v3))}
                    v36 -= 42
                    _ = v55
                }
            case "a":
                var v56 int8 = 127
                _ = v56
            default:
                v2 = append(v2, "a")
            }
            _ = v46
        } else {
            var v57 int32 = 0
            if v1 {
                v4 -= v4
            }
            _ = v57
        }
    } else {
        switch "\n\t" {
        case "foo bar":
            v37 = v1
        }
        var v58 int = -456
        _ = v58
    }
    switch genFn0(-223, 0) {
    case 217:
        switch -2147483648 {
        case 611:
            if false {
                var v59 int16 = 0
                _ = v59
            }
        case 4294967295:
            for fi60 := 0; fi60 < 13; fi60++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi60
            }
        default:
            v3 -= (v3 ^ 1)
        }
    case 9223372036854775807:
        var v61 string = "\n\t"
        _ = v61
    case -305:
        fmt.Println(v3)
    }
    var v62 int32 = 2147483647
    switch "a" {
    case "foo bar":
        var v63 bool = true
        _ = v63
    case "hello":
        switch 127 {
        case 32767:
            for fi64 := 0; fi64 < 20; fi64++ {
                _ = len(v2)
                v62 = (v62 | (v62 << (((v3 % (((^(v3))) | 1))) & 63)))
                _ = fi64
            }
        case 2147483647:
            var v65 int8 = 42
            _ = v65
        case -339:
            var v66 int8 = 1
            _ = v66
        }
    }
    if (1 >= 34) {
        v2 = append(v2, ("\n\t" + "foo bar"))
        _ = len(v2)
    } else {
        if v37 {
            v62 += v62
        } else {
            for fi67 := 0; fi67 < 19; fi67++ {
                v4 += 1
                fmt.Println(fi67)
                _ = fi67
            }
            _ = len(v2)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v42
    _ = v62
}
