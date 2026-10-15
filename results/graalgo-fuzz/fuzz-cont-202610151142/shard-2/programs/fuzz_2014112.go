package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8; F1 string }
type GenS8 struct { F0 int; F1 bool; F2 uint32 }
type GenS9 struct { F0 bool; F1 int32; F2 uint }
type GenS10 struct { F0 uint64; F1 int64 }
type GenS11 struct { F0 int32; F1 int32 }
func genFn0(p0 uint32, p1 int8) string {
    v6 := GenS1{F0: (!(false || (2 == 451)))}
    fmt.Println(p0, p0)
    for fi7 := 0; fi7 < 10; fi7++ {
        p1 -= (p1 ^ p1)
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v6
    return ("foo bar" + "a")
}
func genFn1(p0 bool, p1 int) uint16 {
    p1 = -1
    v8 := []int{(p1 % ((p1) | 1)), (-(p1)), (p1 & (p1 & p1))}
    _ = p0
    _ = p1
    _ = v8
    return 546
}

func main() {
    v1 := make(map[string]int)
    var v2 string = ""
    var v3 uint32 = 418
    var v4 int = 4294967295
    v5 := []uint32{v3, v3}
    v4 += v4
    fmt.Println(v4)
    fmt.Println(v2, v2)
    v9 := GenS2{F0: v3}
    v3 -= v3
    switch v2 {
    case "hello":
        v4 += (v4 | v4)
    case "\n\t":
        v10 := GenS3{F0: -444, F1: 128}
        _ = v10
    }
    v11 := GenS1{F0: (329 == -1)}
    if ((true || (false && (true || false))) && true) {
        fmt.Println(v3, v3)
    } else {
        if ((!false) || true) {
            if false {
                if ((((0.1 * float64(1.5)) < 1.5) || false) || true) {
                    v2 = ("foo bar" + ((v2 + "") + ((v2 + "hello") + (v2 + v2))))
                    if true {
                        if (42 > -32768) {
                            if false {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v4 = (v4 - (v4 | (v4 * (v4 % (((v4 + (v4 * v4))) | 1)))))
                                for fi12 := 0; fi12 < 2; fi12++ {
                                    var v13 int16 = 32767
                                    var v14 uint = 42
                                    _ = fi12
                                    _ = v13
                                    _ = v14
                                }
                            } else {
                                v3 += v3
                                var v15 uint16 = 1
                                v15 = v15
                                _ = v15
                            }
                        } else {
                            v4 = (v4 & -166)
                            for fi16 := 0; fi16 < 17; fi16++ {
                                v3 = v3
                                for fi17 := 0; fi17 < 5; fi17++ {
                                    v4 += v4
                                    var v18 int32 = 0
                                    _ = fi17
                                    _ = v18
                                }
                                _ = fi16
                            }
                        }
                    } else {
                        if false {
                            v19 := GenS1{F0: false}
                            _ = v19
                        }
                        v3 += (v3 & 298)
                    }
                }
            } else {
                if ((v4 % ((v4) | 1)) >= (v4 + v4)) {
                    var v20 uint8 = 128
                    if true {
                        var v21 int32 = 2147483647
                        _ = v21
                    }
                    if (!(!false)) {
                        switch v2 {
                        case "a":
                            if true {
                                if (1.5 >= 0.1) {
                                    for fi22 := 0; fi22 < 15; fi22++ {
                                        v5 = append(v5, (v3 >> ((1) & 63)))
                                        fmt.Println(v4, v3, v20)
                                        _ = fi22
                                    }
                                    for fi23 := 0; fi23 < 9; fi23++ {
                                        v24 := GenS4{F0: true, F1: 255}
                                        var v25 int64 = -9223372036854775808
                                        _ = fi23
                                        _ = v24
                                        _ = v25
                                    }
                                    var v26 uint32 = (v3 - 272)
                                    _ = v26
                                } else {
                                    var v27 int64 = 270
                                    v28 := GenS5{F0: v3}
                                    switch (((v2 + v2) + "foo bar") + "\n\t") {
                                    case "a":
                                        for fi29 := 0; fi29 < 15; fi29++ {
                                            var v30 uint16 = 1
                                            for fi31 := 0; fi31 < 18; fi31++ {
                                                if len(v5) > 0 {
                                                    v5[0] = v3
                                                }
                                                v32 := GenS5{F0: v3}
                                                switch v2 {
                                                case "a":
                                                    v33 := GenS2{F0: (v3 / ((v3) | 1))}
                                                    _ = v33
                                                case "foo bar":
                                                    if false {
                                                        var v34 int32 = 1
                                                        v3 = (v3 / ((0) | 1))
                                                        _ = v34
                                                    }
                                                case "\n\t":
                                                    switch v2 {
                                                    case "\n\t":
                                                        v35 := GenS6{F0: 42}
                                                        _ = v35
                                                    case "a":
                                                        var v36 uint64 = 662
                                                        _ = v36
                                                    case "hello":
                                                        v3 -= 521
                                                    }
                                                default:
                                                    v27 += 0
                                                }
                                                var v37 uint = 255
                                                if true {
                                                    var v38 uint = v37
                                                    v2 = v2
                                                    _ = v38
                                                }
                                                fmt.Println(fi31, fi29)
                                                var v39 uint = 1
                                                var v40 bool = true
                                                v37 = v39
                                                _ = fi31
                                                _ = v32
                                                _ = v37
                                                _ = v39
                                                _ = v40
                                            }
                                            _ = fi29
                                            _ = v30
                                        }
                                    }
                                    _ = v27
                                    _ = v28
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v2 = ("foo bar" + v2)
                            }
                        case "foo bar":
                            v20 = 173
                        }
                    }
                    _ = v20
                }
                v41 := GenS2{F0: v3}
                var v42 bool = false
                _ = v41
                _ = v42
            }
            switch v2 {
            case "":
                v2 = ""
            case "a":
                var v43 uint = 2
                _ = v43
            case "foo bar":
                v5 = append(v5, v3)
            default:
                v4 = v4
            }
            v3 = (^(v3))
        }
        for fi44 := 0; fi44 < 17; fi44++ {
            var v45 uint16 = 65535
            switch (v2 + (((v2 + (v2 + ((v2 + "") + v2))) + ((((v2 + "") + v2) + (v2 + (v2 + "hello"))) + v2)) + v2)) {
            case "":
                var v46 uint16 = genFn1(true, (v4 / ((v4) | 1)))
                _ = v46
            default:
                for fi47 := 0; fi47 < 17; fi47++ {
                    v48 := GenS3{F0: 1000, F1: 2}
                    v4 = fi47
                    fmt.Println(v2, v2)
                    _ = fi47
                    _ = v48
                }
            }
            _ = fi44
            _ = v45
        }
    }
    if (((0.1 + float64(float32(0.0))) / ((3.14159 / float64(1e10)) - float64(float32((3.14159 + 3.14159))))) > 0.1) {
        if (255 != genFn1((true || false), (v4 ^ 842))) {
            fmt.Println(v3, v3)
            var v49 string = v2
            if false {
                fmt.Println(v49)
                v2 = ""
            }
            _ = v49
        } else {
            v4 -= (^(v4))
            var v50 uint16 = 148
            _ = v50
        }
        switch v2 {
        case "foo bar":
            for fi51 := 0; fi51 < 15; fi51++ {
                v2 = v2
                if (true || true) {
                    v5 = append(v5, (v3 << ((867) & 63)))
                    if true {
                        fmt.Println(v3, v4)
                        v52 := GenS7{F0: 127, F1: "a"}
                        _ = v52
                    }
                    switch v2 {
                    case "hello":
                        var v53 int32 = 2147483647
                        _ = v53
                    case "\n\t":
                        for fi54 := 0; fi54 < 7; fi54++ {
                            v55 := GenS8{F0: fi54, F1: false, F2: v3}
                            _ = fi54
                            _ = v55
                        }
                    case "a":
                        for fi56 := 0; fi56 < 7; fi56++ {
                            switch (((((((v2 + v2) + "\n\t") + v2) + "\n\t") + "foo bar") + v2) + ("hello" + "")) {
                            case "hello":
                                v4 = v4
                            case "\n\t":
                                v2 = ("a" + "a")
                            }
                            v57 := GenS9{F0: false, F1: 0, F2: 1}
                            if ("foo bar" != (v2 + v2)) {
                                v3 = v3
                                switch (fi51 / ((32767) | 1)) {
                                case -598:
                                    fmt.Println(v2, fi51)
                                    v58 := GenS10{F0: 681, F1: 0}
                                    _ = v58
                                case 2147483647:
                                    v2 = v2
                                default:
                                    switch -9223372036854775808 {
                                    case -94:
                                        var v59 uint32 = 0
                                        _ = v59
                                    case 127:
                                        v5 = append(v5, (^(v3)))
                                        if true {
                                            v2 = v2
                                            var v60 int16 = 1
                                            v60 = v60
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            _ = v60
                                        }
                                    }
                                }
                            } else {
                                var v61 uint8 = 128
                                fmt.Println(fi56, v61)
                                _ = v61
                            }
                            _ = fi56
                            _ = v57
                        }
                    default:
                        if true {
                            fmt.Println(v2)
                        } else {
                            fmt.Println(v4, v4)
                        }
                    }
                } else {
                    v62 := GenS7{F0: 1, F1: ((("\n\t" + v2) + v2) + (v2 + "hello"))}
                    for fi63 := 0; fi63 < 18; fi63++ {
                        switch int(v3) {
                        case 32767:
                            if (!false) {
                                switch (v4 ^ (fi63 / ((v4) | 1))) {
                                case -294:
                                    v5 = append(v5, (^(v3)))
                                case 0:
                                    fmt.Println(v2, fi51, v3)
                                case 42:
                                    if (0 < 102) {
                                        var v64 float32 = -2.25
                                        var v65 bool = (0 <= 42)
                                        _ = v64
                                        _ = v65
                                    } else {
                                        var v66 int32 = 42
                                        switch fi63 {
                                        case 0:
                                            if false {
                                                v2 = ("hello" + (v2 + ((v2 + (((v2 + "\n\t") + v2) + "")) + v2)))
                                                var v67 uint16 = 0
                                                v4 = fi63
                                                _ = v67
                                            } else {
                                                var v68 string = v2
                                                v69 := GenS8{F0: (-(fi51)), F1: (!(1 < -64)), F2: (v3 & 65535)}
                                                _ = len(v1)
                                                _ = v68
                                                _ = v69
                                            }
                                        case 127:
                                            switch (^(fi63)) {
                                            case 32767:
                                                var v70 float32 = 1e10
                                                _ = v70
                                            case 0:
                                                switch -1 {
                                                case -276:
                                                    v3 -= (^(v3))
                                                case -805:
                                                    if false {
                                                        v71 := GenS9{F0: (864 == 0), F1: v66, F2: 807}
                                                        var v72 bool = false
                                                        v3 = (^(v3))
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                default:
                                                    _ = len(v1)
                                                }
                                            }
                                        default:
                                            fmt.Println(v4)
                                        }
                                        v1[v2] = 9223372036854775807
                                        _ = v66
                                    }
                                default:
                                    _ = len(v5)
                                }
                                var v73 int64 = int64(fi63)
                                fmt.Println(fi63)
                                _ = v73
                            }
                        case 2147483647:
                            if true {
                                fmt.Println(fi51, v3, v2)
                                switch "hello" {
                                case "":
                                    v74 := v1["foo bar"]
                                    _ = v74
                                default:
                                    v4 += v4
                                }
                            }
                        }
                        _ = fi63
                    }
                    _ = v62
                }
                for fi75 := 0; fi75 < 8; fi75++ {
                    if (41 > 4294967295) {
                        v76 := GenS6{F0: 18446744073709551615}
                        _ = v76
                    } else {
                        var v77 uint16 = 42
                        if (750 == -9223372036854775808) {
                            v2 = v2
                            fmt.Println(v3, v2, fi51)
                        } else {
                            var v78 int8 = 1
                            v79 := GenS5{F0: v3}
                            switch v2 {
                            case "hello":
                                v3 = (v3 % ((2147483647) | 1))
                            case "":
                                fmt.Println(v4)
                            case "\n\t":
                                v3 -= (v3 * 242)
                            }
                            _ = v78
                            _ = v79
                        }
                        v1[""] = v4
                        v77 += (^(v77))
                        var v80 int64 = 42
                        _ = v77
                        _ = v80
                    }
                    _ = fi75
                }
                _ = fi51
            }
        }
        var v81 uint64 = 42
        _ = v81
    } else {
        for fi82 := 0; fi82 < 16; fi82++ {
            var v83 uint32 = v3
            var v84 uint32 = (v3 * (v83 / (((v3 >> ((42) & 63))) | 1)))
            _ = fi82
            _ = v83
            _ = v84
        }
        delete(v1, "hello")
    }
    if false {
        var v85 float32 = (float32((0.0 + 1e10)) / 0.1)
        v3 = (v3 + (v3 / (((v3 % ((v3) | 1))) | 1)))
        _ = v85
    } else {
        v4 = v4
        for fi86 := 0; fi86 < 16; fi86++ {
            v87 := GenS11{F0: 2147483647, F1: -1}
            _ = fi86
            _ = v87
        }
        v1[v2] = v4
    }
    var v88 uint = 0
    v88 = (v88 ^ (^(v88)))
    v2 = "foo bar"
    v89 := GenS4{F0: (-627 < 42), F1: (v88 & v88)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v88
    _ = v89
    _ = v9
}
