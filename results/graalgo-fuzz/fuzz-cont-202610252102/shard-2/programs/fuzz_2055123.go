package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 uint32; F1 int; F2 int8 }
type GenS9 struct { F0 uint32; F1 string }
func genFn0(p0 uint, p1 int16, p2 int64) int16 {
    var v7 float64 = -2.25
    var v8 int16 = p1
    v9 := []int{-94, -2147483648, -364}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return v8
}
func genFn1(p0 bool, p1 uint, p2 uint16) uint32 {
    for fi10 := 0; fi10 < 7; fi10++ {
        if (42 < uint8(p1)) {
            v11 := []int{(-(fi10)), fi10}
            v12 := GenS1{F0: (!p0)}
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 uint16 = p2
                var v15 bool = p0
                _ = fi13
                _ = v14
                _ = v15
            }
            var v16 float32 = -2.25
            switch -829 {
            case -698:
                for fi17 := 0; fi17 < 13; fi17++ {
                    switch (("\n\t" + "\n\t") + ("\n\t" + "a")) {
                    case "":
                        v18 := GenS1{F0: (0 < 0)}
                        _ = v18
                    case "\n\t":
                        fmt.Println(fi17, v16, v16)
                    case "a":
                        fmt.Println(p2)
                    }
                    v19 := []int{-1, (fi17 & 4294967295), fi17}
                    v20 := GenS2{F0: (("\n\t" + (("\n\t" + "\n\t") + ("foo bar" + "foo bar"))) > ("" + "")), F1: -2147483648}
                    var v21 float64 = float64((v16 - v16))
                    _ = len(v19)
                    _ = fi17
                    _ = v19
                    _ = v20
                    _ = v21
                }
            }
            _ = v16
        }
        _ = fi10
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        for fi23 := 0; fi23 < 8; fi23++ {
            v24 := GenS2{F0: true, F1: -1}
            if true {
                if (!p0) {
                    v25 := GenS1{F0: false}
                    _ = v25
                }
            }
            _ = fi23
            _ = v24
        }
        if (1.5 > (0.0 - -2.25)) {
            p2 -= 239
            var v26 float32 = float32(3.14159)
            _ = v26
        }
        _ = fi22
    }
    fmt.Println(p1)
    for fi27 := 0; fi27 < 5; fi27++ {
        for fi28 := 0; fi28 < 17; fi28++ {
            v29 := GenS3{F0: 938, F1: -1}
            if (127 == -32) {
                p1 = (p1 / ((620) | 1))
            } else {
                p1 -= p1
                var v30 uint64 = 42
                v31 := GenS1{F0: p0}
                _ = v30
                _ = v31
            }
            _ = fi28
            _ = v29
        }
        p2 = (p2 & (p2 - (p2 & p2)))
        for fi32 := 0; fi32 < 20; fi32++ {
            v33 := GenS2{F0: (!((42 != -128) && true)), F1: -2147483648}
            switch "\n\t" {
            case "foo bar":
                fmt.Println(p1, fi27)
            default:
                var v34 bool = p0
                _ = v34
            }
            var v35 uint = (p1 << (((p1 ^ 4294967295)) & 63))
            v36 := []int{(fi27 >> ((v35) & 63)), (fi27 >> (((^(p1))) & 63))}
            var v37 uint = (p1 + (v35 >> (((v35 << ((2) & 63))) & 63)))
            _ = fi32
            _ = v33
            _ = v35
            _ = v36
            _ = v37
        }
        _ = fi27
    }
    var v38 int8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v38
    return 65535
}

func main() {
    var v1 string = ((("\n\t" + "") + "foo bar") + (("foo bar" + "hello") + ("foo bar" + ((("" + "foo bar") + "hello") + ""))))
    v2 := make(map[int]int)
    var v3 string = v1
    var v4 int64 = 1
    var v5 string = ""
    v6 := []int{}
    delete(v2, -2147483648)
    for fi39 := 0; fi39 < 7; fi39++ {
        fmt.Println(fi39, v5, fi39)
        fmt.Println(v3)
        fmt.Println(v3)
        _ = fi39
    }
    if false {
        var v40 float64 = 3.14159
        switch -9223372036854775808 {
        case 4294967295:
            switch 374 {
            case 2147483647:
                var v41 float32 = 0.0
                _ = v41
            default:
                v5 = v5
            }
        case 32767:
            v42 := GenS4{F0: true}
            _ = v42
        case 1:
            var v43 uint8 = 38
            _ = v43
        }
        _ = v40
    } else {
        var v44 int8 = -128
        _ = v44
    }
    v45 := GenS5{F0: 127}
    if true {
        var v46 int64 = v4
        _ = v46
    }
    if true {
        if (506 == 42) {
            for fi47 := 0; fi47 < 17; fi47++ {
                v1 = v5
                _ = fi47
            }
            var v48 int64 = (v4 * v4)
            _ = v48
        }
        v49 := v2[167]
        switch (v49 | (v49 | v49)) {
        case 42:
            fmt.Println(v49, v5)
        default:
            v50 := GenS5{F0: 127}
            _ = v50
        }
        _ = v49
    }
    switch -1 {
    case 42:
        switch 4294967295 {
        case 0:
            for fi51 := 0; fi51 < 14; fi51++ {
                delete(v2, (fi51 - fi51))
                var v52 int32 = 0
                for fi53 := 0; fi53 < 12; fi53++ {
                    switch (v1 + ("hello" + v1)) {
                    case "foo bar":
                        switch -300 {
                        case 1:
                            if false {
                                if false {
                                    v1 = v1
                                    v54 := GenS6{F0: genFn0(42, genFn0(581, 42, 1), v4), F1: "hello"}
                                    if (!((!false) || true)) {
                                        v2[(fi53 * (fi51 ^ (fi53 ^ (fi53 % (((fi53 | fi51)) | 1)))))] = (-(fi51))
                                    } else {
                                        v55 := GenS1{F0: true}
                                        if (42 < 1000) {
                                            if ((v4 * (v4 / (((v4 << ((42) & 63))) | 1))) >= v4) {
                                                _ = len(v6)
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            }
                                            var v56 uint64 = 42
                                            _ = v56
                                        }
                                        v1 = ""
                                        _ = v55
                                    }
                                    v4 = v4
                                    _ = v54
                                }
                            } else {
                                v57 := GenS6{F0: genFn0(811, 1, v4), F1: v1}
                                v58 := GenS3{F0: 2, F1: 42}
                                if ((4294967295 <= 808) && true) {
                                    v5 = v5
                                }
                                v3 = v3
                                v3 = (v1 + v3)
                                v1 = ("" + "a")
                                _ = v57
                                _ = v58
                            }
                            v59 := GenS7{F0: v4, F1: 0}
                            if (42 < 5) {
                                v1 = v3
                            } else {
                                v60 := GenS1{F0: (908 >= 1)}
                                _ = v60
                            }
                            _ = v59
                        case 287:
                            fmt.Println(v1, fi51)
                        case 32767:
                            _ = len(v2)
                        default:
                            _ = len(v2)
                        }
                    case "a":
                        if len(v6) > 0 {
                            v6[0] = fi51
                        }
                    default:
                        fmt.Println(v4, v3)
                    }
                    if (((!(331 == 914)) || false) && (false && true)) {
                        switch (fi51 + (^(fi53))) {
                        case 32768:
                            v61 := GenS2{F0: ((1 <= 42) || false), F1: 1}
                            _ = v61
                        case 4294967295:
                            var v62 float32 = float32(0.1)
                            _ = v62
                        default:
                            fmt.Println(fi53)
                        }
                    } else {
                        fmt.Println(fi51, v52)
                        v1 = ""
                    }
                    if (!false) {
                        if (!false) {
                            switch (fi51 + fi51) {
                            case 892:
                                fmt.Println(v3)
                            case 265:
                                if false {
                                    v4 -= 42
                                    switch v3 {
                                    case "hello":
                                        var v63 int8 = 0
                                        _ = v63
                                    case "a":
                                        var v64 bool = (v1 <= v1)
                                        _ = v64
                                    case "foo bar":
                                        var v65 int32 = (v52 >> ((563) & 63))
                                        _ = v65
                                    default:
                                        v4 = (v4 ^ (v4 / ((v4) | 1)))
                                    }
                                    var v66 float64 = -2.25
                                    _ = v66
                                } else {
                                    fmt.Println(v1)
                                    v52 = -2147483648
                                }
                            }
                            if (!(true || false)) {
                                var v67 uint8 = 127
                                _ = v67
                            }
                        }
                    } else {
                        v3 = v3
                        v68 := GenS6{F0: int16(fi53), F1: v3}
                        var v69 int32 = -1
                        _ = v68
                        _ = v69
                    }
                    _ = fi53
                }
                _ = fi51
                _ = v52
            }
        case 944:
            if false {
                v5 = v1
                v6 = append(v6, 42)
                v5 = "\n\t"
            }
        }
    case 9223372036854775807:
        v4 = 42
    }
    if (!false) {
        if ((float64(0.0) / 1e10) != float64(1e10)) {
            var v70 int16 = -1
            for fi71 := 0; fi71 < 19; fi71++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi71
            }
            _ = v70
        }
        switch v5 {
        case "hello":
            for fi72 := 0; fi72 < 12; fi72++ {
                v73 := GenS8{F0: 33, F1: (fi72 / ((-389) | 1)), F2: 42}
                v74 := GenS9{F0: 2147483647, F1: v5}
                v75 := GenS8{F0: 42, F1: fi72, F2: -7}
                _ = fi72
                _ = v73
                _ = v74
                _ = v75
            }
        case "":
            v76 := GenS8{F0: 4294967295, F1: 128, F2: -1}
            _ = v76
        default:
            for fi77 := 0; fi77 < 4; fi77++ {
                v3 = v5
                v1 = v1
                v3 = v3
                _ = fi77
            }
        }
    } else {
        var v78 int64 = (v4 | v4)
        _ = v78
    }
    v79 := GenS2{F0: (-1 > 100), F1: -2147483648}
    fmt.Println(v4, v4, v5)
    v1 = ("\n\t" + ("" + v1))
    v4 = v4
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v79
}
