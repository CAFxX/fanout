package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 uint16; F1 uint32; F2 uint16 }
func genFn0(p0 int64) int64 {
    var v6 int16 = -32768
    if false {
        p0 = 0
    } else {
        var v7 float32 = 3.14159
        var v8 uint16 = 580
        p0 = (p0 >> ((4294967295) & 63))
        _ = v7
        _ = v8
    }
    var v9 bool = (!false)
    _ = p0
    _ = v6
    _ = v9
    return p0
}
func genFn1(p0 uint8, p1 uint16) int8 {
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 uint8 = 42
        var v12 int8 = 127
        v13 := GenS1{F0: 1, F1: false}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    p0 -= (p0 * p0)
    var v14 uint = 622
    v15 := GenS2{F0: -1}
    p1 += p1
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 100
}

func main() {
    var v1 int64 = -421
    v2 := make(map[string]uint32)
    v3 := []uint32{42, 0}
    v4 := []string{"a", ("foo bar" + "foo bar"), "\n\t"}
    var v5 int32 = -812
    var v16 int16 = 1000
    var v17 int64 = (v1 - (v1 & (v1 + 1)))
    for fi18 := 0; fi18 < 6; fi18++ {
        v17 = (^(v17))
        var v19 string = ("a" + "hello")
        v20 := GenS1{F0: (v5 + v5), F1: (genFn0(genFn0((-(v17)))) >= v1)}
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        v22 := GenS1{F0: (v5 % ((v5) | 1)), F1: false}
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: v5, F1: false}
    v17 -= genFn0(v17)
    for fi24 := 0; fi24 < 17; fi24++ {
        switch 30 {
        case 4294967295:
            switch fi24 {
            case 32767:
                switch 32768 {
                case 32768:
                    if (!false) {
                        var v25 float64 = 1e10
                        switch "\n\t" {
                        case "a":
                            var v26 int64 = v1
                            _ = v26
                        default:
                            if true {
                                for fi27 := 0; fi27 < 18; fi27++ {
                                    switch -996 {
                                    case 127:
                                        if true {
                                            if (1 <= 32767) {
                                                var v28 string = "hello"
                                                switch fi27 {
                                                case 870:
                                                    var v29 float64 = 3.14159
                                                    _ = v29
                                                default:
                                                    v30 := GenS2{F0: int64(fi27)}
                                                    _ = v30
                                                }
                                                _ = v28
                                            } else {
                                                v1 = 843
                                                v17 -= v1
                                                var v31 float64 = float64(1.5)
                                                _ = v31
                                            }
                                            var v32 int64 = 873
                                            var v33 int32 = (v5 | int32(v1))
                                            _ = v32
                                            _ = v33
                                        }
                                    default:
                                        v17 = (v1 + int64(v16))
                                    }
                                    _ = fi27
                                }
                                v34 := GenS1{F0: (v5 - (v5 >> ((255) & 63))), F1: true}
                                _ = v34
                            } else {
                                delete(v2, "hello")
                                v4 = append(v4, ("\n\t" + ""))
                                if false {
                                    for fi35 := 0; fi35 < 16; fi35++ {
                                        if (!(0 < 255)) {
                                            v36 := GenS1{F0: 42, F1: (294 >= 2)}
                                            _ = v36
                                        } else {
                                            fmt.Println(v5, v16, fi35)
                                            var v37 uint = 42
                                            v38 := GenS3{F0: (fi35 > fi24), F1: 32767}
                                            _ = v37
                                            _ = v38
                                        }
                                        v17 -= (v17 * v17)
                                        v39 := GenS1{F0: v5, F1: ((v1 | v1) >= -675)}
                                        fmt.Println(fi35, fi35, v1)
                                        _ = fi35
                                        _ = v39
                                    }
                                    v40 := GenS4{F0: 50}
                                    v1 += 41
                                    _ = v40
                                }
                            }
                        }
                        _ = v25
                    }
                case 1:
                    var v41 int8 = genFn1(224, 65535)
                    _ = v41
                case 32767:
                    if ((2 == 2) && (("" + ("\n\t" + "hello")) <= ("hello" + "a"))) {
                        fmt.Println(v5, v1, v17)
                        v42 := GenS5{F0: int(v5)}
                        switch "foo bar" {
                        case "hello":
                            for fi43 := 0; fi43 < 17; fi43++ {
                                _ = len(v2)
                                if (v5 == v5) {
                                    var v44 int8 = genFn1(255, 0)
                                    if true {
                                        switch (^(fi43)) {
                                        case 227:
                                            var v45 uint = 458
                                            _ = v45
                                        case -2147483648:
                                            var v46 bool = (!true)
                                            _ = v46
                                        case 222:
                                            if (289 != 2) {
                                                var v47 uint32 = 75
                                                var v48 int32 = (v5 << ((42) & 63))
                                                v17 = (v17 * 133)
                                                _ = v47
                                                _ = v48
                                            }
                                        default:
                                            var v49 int8 = -97
                                            _ = v49
                                        }
                                        fmt.Println(v44, v44)
                                        var v50 uint16 = 898
                                        _ = v50
                                    } else {
                                        var v51 uint32 = 42
                                        v16 -= (v16 / ((-32768) | 1))
                                        var v52 uint64 = 18446744073709551615
                                        _ = v51
                                        _ = v52
                                    }
                                    _ = v44
                                }
                                _ = fi43
                            }
                        }
                        _ = v42
                    } else {
                        delete(v2, "foo bar")
                        v16 = -1
                        switch ("\n\t" + ("" + ("hello" + (("" + "a") + ("\n\t" + "a"))))) {
                        case "\n\t":
                            for fi53 := 0; fi53 < 13; fi53++ {
                                v54 := v2[""]
                                if true {
                                    if (float64(1e10) >= 1e10) {
                                        v55 := GenS2{F0: (v17 & -674)}
                                        _ = v55
                                    }
                                } else {
                                    if true {
                                        v17 = (v17 | v1)
                                        switch fi24 {
                                        case 746:
                                            if false {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                fmt.Println(v5)
                                                _ = len(v2)
                                            } else {
                                                var v56 string = (("a" + "") + "\n\t")
                                                _ = v56
                                            }
                                        case 127:
                                            switch int(v5) {
                                            case -932:
                                                var v57 float64 = (3.14159 / float64(1.5))
                                                _ = v57
                                            }
                                        default:
                                            switch (-(fi53)) {
                                            case 9223372036854775807:
                                                v58 := v2["foo bar"]
                                                _ = v58
                                            case 32767:
                                                var v59 int16 = v16
                                                _ = v59
                                            default:
                                                v60 := GenS6{F0: true, F1: 42}
                                                _ = v60
                                            }
                                        }
                                        var v61 uint64 = 2
                                        _ = v61
                                    }
                                    v3 = append(v3, (v54 + 225))
                                    v2["a"] = v54
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi53
                                _ = v54
                            }
                        default:
                            switch "hello" {
                            case "\n\t":
                                fmt.Println(v16, v1, v1)
                            default:
                                switch "" {
                                case "\n\t":
                                    for fi62 := 0; fi62 < 2; fi62++ {
                                        var v63 int32 = (v5 + 0)
                                        switch 702 {
                                        case 2147483647:
                                            fmt.Println(fi24, v63, v63)
                                        case -363:
                                            v5 = v5
                                        case 102:
                                            switch (fi24 >> ((2) & 63)) {
                                            case -1:
                                                v1 = (v17 * 9223372036854775807)
                                            case 4294967295:
                                                fmt.Println(fi62)
                                            case 2147483647:
                                                var v64 uint16 = 42
                                                _ = v64
                                            }
                                        }
                                        _ = fi62
                                        _ = v63
                                    }
                                }
                            }
                        }
                    }
                default:
                    if false {
                        for fi65 := 0; fi65 < 15; fi65++ {
                            var v66 uint8 = 19
                            _ = fi65
                            _ = v66
                        }
                    }
                }
            case 0:
                fmt.Println(fi24)
            case 127:
                v67 := GenS7{F0: 32767, F1: 2147483648, F2: 818}
                _ = v67
            }
        case -9223372036854775808:
            for fi68 := 0; fi68 < 12; fi68++ {
                if (fi68 <= (fi24 * 42)) {
                    fmt.Println(v17, fi24)
                    switch "hello" {
                    case "\n\t":
                        var v69 uint64 = 199
                        _ = v69
                    case "a":
                        var v70 int = fi68
                        _ = v70
                    }
                }
                _ = fi68
            }
        }
        var v71 uint32 = 42
        _ = fi24
        _ = v71
    }
    v16 = 1000
    var v72 float64 = (0.1 + 1.5)
    switch 1 {
    case -1:
        var v73 uint16 = 0
        _ = v73
    case 9223372036854775807:
        var v74 uint64 = 1
        _ = v74
    }
    _ = len(v2)
    v16 = v16
    v16 += (v16 | v16)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v72
}
