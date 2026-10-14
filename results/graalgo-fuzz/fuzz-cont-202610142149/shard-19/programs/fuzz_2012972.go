package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 int16; F1 uint8 }
type GenS7 struct { F0 uint8; F1 bool; F2 uint16 }

func main() {
    var v1 int32 = 2147483647
    var v2 uint8 = 99
    var v3 uint16 = 42
    var v4 int64 = 796
    var v5 string = ("" + "a")
    if ((!(!(1e10 >= 1e10))) && true) {
        v4 -= (v4 & (v4 % (((v4 * (v4 + v4))) | 1)))
    }
    if (-2.25 != ((3.14159 / 1e10) - 1e10)) {
        fmt.Println(v1)
        fmt.Println(v2)
        for fi6 := 0; fi6 < 14; fi6++ {
            for fi7 := 0; fi7 < 20; fi7++ {
                fmt.Println(v2, v5)
                var v8 bool = (-1 != (-(fi6)))
                var v9 int8 = 1
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = fi6
        }
    }
    fmt.Println(v2)
    v10 := GenS1{F0: v2, F1: 544}
    if false {
        if (42 < 2) {
            fmt.Println(v2, v1)
            var v11 uint8 = 1
            v12 := GenS2{F0: 4294967295}
            _ = v11
            _ = v12
        } else {
            switch v5 {
            case "a":
                v5 = (v5 + v5)
            case "foo bar":
                if false {
                    v13 := GenS1{F0: (^(v2)), F1: 321}
                    var v14 int = 127
                    fmt.Println(v5, v4, v14)
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS3{F0: 1, F1: v2, F2: 0}
                    if (0.0 > -2.25) {
                        var v16 int = 42
                        fmt.Println(v2)
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 16; fi17++ {
                            v18 := GenS2{F0: 0}
                            if false {
                                fmt.Println(v3, v5)
                                if (!true) {
                                    v19 := []int{32767, fi17, (fi17 - (^(fi17)))}
                                    fmt.Println(v4, v2, v5)
                                    var v20 uint64 = 1
                                    _ = v19
                                    _ = v20
                                } else {
                                    switch fi17 {
                                    case 1:
                                        for fi21 := 0; fi21 < 2; fi21++ {
                                            v22 := []int{(fi17 - fi17)}
                                            var v23 uint8 = (^(v2))
                                            if (true || true) {
                                                var v24 int64 = -9223372036854775808
                                                var v25 uint32 = 4294967295
                                                var v26 uint16 = (v3 << ((2) & 63))
                                                fmt.Println(v2)
                                                if (!false) {
                                                    v27 := GenS1{F0: v2, F1: 42}
                                                    var v28 uint32 = v25
                                                    _ = v27
                                                    _ = v28
                                                } else {
                                                    fmt.Println(v1)
                                                }
                                                _ = v24
                                                _ = v25
                                                _ = v26
                                            } else {
                                                v1 = v1
                                            }
                                            v29 := GenS2{F0: 691}
                                            if (636 != 42) {
                                                v30 := GenS1{F0: (v2 + (v2 >> ((630) & 63))), F1: 722}
                                                var v31 float64 = 0.1
                                                fmt.Println(fi17)
                                                _ = len(v22)
                                                var v32 float64 = v31
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            } else {
                                                var v33 uint = 18446744073709551615
                                                if true {
                                                    v5 = "\n\t"
                                                    var v34 string = "a"
                                                    v35 := GenS1{F0: (v23 ^ (v23 >> ((4294967295) & 63))), F1: (^(fi21))}
                                                    v36 := GenS1{F0: v23, F1: (fi21 % (((fi17 >> ((255) & 63))) | 1))}
                                                    v1 -= (v1 >> ((225) & 63))
                                                    v2 = (v2 - v2)
                                                    v2 = v23
                                                    _ = v34
                                                    _ = v35
                                                    _ = v36
                                                } else {
                                                    if true {
                                                        if false {
                                                            v1 = (v1 % (((v1 / ((v1) | 1))) | 1))
                                                            var v37 int = fi17
                                                            _ = len(v22)
                                                            _ = v37
                                                        } else {
                                                            if true {
                                                                switch v5 {
                                                                case "\n\t":
                                                                    if false {
                                                                        v33 += v33
                                                                        v22 = append(v22, (fi17 >> ((825) & 63)))
                                                                    }
                                                                case "":
                                                                    var v38 int32 = (v1 >> (((v33 << ((v33) & 63))) & 63))
                                                                    _ = v38
                                                                }
                                                                v39 := GenS4{F0: (^(v33)), F1: v2}
                                                                v22 = append(v22, fi17)
                                                                var v40 uint16 = (v3 >> (((^(v33))) & 63))
                                                                v41 := GenS3{F0: (v1 ^ v1), F1: v2, F2: 2}
                                                                _ = v39
                                                                _ = v40
                                                                _ = v41
                                                            } else {
                                                                v22 = append(v22, (fi21 >> ((2) & 63)))
                                                                v33 += 0
                                                            }
                                                        }
                                                        v42 := GenS5{F0: (v33 / (((^(v33))) | 1)), F1: 127}
                                                        v23 = v23
                                                        _ = v42
                                                    } else {
                                                        v2 -= v2
                                                        var v43 uint8 = 127
                                                        _ = v43
                                                    }
                                                    v44 := GenS4{F0: (v33 * 18446744073709551615), F1: uint8(v33)}
                                                    fmt.Println(fi21, v33, v3)
                                                    _ = v44
                                                }
                                                var v45 uint16 = (^(v3))
                                                _ = v33
                                                _ = v45
                                            }
                                            v4 = (v4 & (v4 >> ((1) & 63)))
                                            _ = fi21
                                            _ = v22
                                            _ = v23
                                            _ = v29
                                        }
                                    }
                                    var v46 string = (((((("foo bar" + "foo bar") + v5) + ("hello" + ((v5 + v5) + v5))) + v5) + v5) + (v5 + v5))
                                    if false {
                                        v5 = (v46 + v5)
                                        v5 = v46
                                    } else {
                                        v4 -= (-(v4))
                                    }
                                    _ = v46
                                }
                            }
                            _ = fi17
                            _ = v18
                        }
                        for fi47 := 0; fi47 < 1; fi47++ {
                            for fi48 := 0; fi48 < 1; fi48++ {
                                switch v5 {
                                case "\n\t":
                                    switch v5 {
                                    case "":
                                        v49 := []int{-14, (fi47 % (((-(fi47))) | 1))}
                                        _ = v49
                                    case "foo bar":
                                        if false {
                                            switch fi48 {
                                            case 0:
                                                var v50 int = (fi47 << ((18446744073709551615) & 63))
                                                _ = v50
                                            case 1:
                                                var v51 float32 = -2.25
                                                var v52 int16 = 1000
                                                v52 = (v52 & (v52 - v52))
                                                _ = v51
                                                _ = v52
                                            case -349:
                                                switch (fi47 + (fi47 - (fi47 % ((fi47) | 1)))) {
                                                case 0:
                                                    v53 := []int{}
                                                    _ = v53
                                                case 1:
                                                    var v54 int8 = int8(fi48)
                                                    _ = v54
                                                default:
                                                    switch ("\n\t" + v5) {
                                                    case "":
                                                        var v55 uint32 = 57
                                                        _ = v55
                                                    case "foo bar":
                                                        var v56 float32 = 3.14159
                                                        _ = v56
                                                    }
                                                }
                                            default:
                                                v57 := []int{(fi48 * (fi47 >> ((1) & 63))), fi48, fi47}
                                                _ = v57
                                            }
                                        } else {
                                            switch "\n\t" {
                                            case "":
                                                v58 := []int{(-(fi48)), 32768, fi47}
                                                _ = v58
                                            }
                                        }
                                    case "\n\t":
                                        fmt.Println(v3)
                                    }
                                default:
                                    v5 = (v5 + ((("hello" + v5) + ((((v5 + v5) + ((v5 + v5) + "foo bar")) + (v5 + (("" + v5) + (((v5 + ((v5 + v5) + "")) + (v5 + "hello")) + (((v5 + v5) + "a") + v5))))) + "\n\t")) + v5))
                                }
                                var v59 int32 = (v1 * (v1 << ((18446744073709551615) & 63)))
                                _ = fi48
                                _ = v59
                            }
                            _ = fi47
                        }
                        v60 := []int{}
                        _ = v60
                    }
                    _ = v15
                }
            case "\n\t":
                for fi61 := 0; fi61 < 11; fi61++ {
                    v62 := GenS1{F0: v2, F1: fi61}
                    var v63 int16 = -278
                    _ = fi61
                    _ = v62
                    _ = v63
                }
            }
        }
        var v64 uint = 81
        _ = v64
    } else {
        for fi65 := 0; fi65 < 10; fi65++ {
            var v66 uint8 = 2
            v2 = (v66 % ((42) | 1))
            _ = fi65
            _ = v66
        }
    }
    v1 = v1
    fmt.Println(v3, v3, v5)
    fmt.Println(v4, v4, v2)
    v67 := []int{-298}
    v68 := GenS2{F0: uint64(v3)}
    var v69 uint16 = v3
    for fi70 := 0; fi70 < 1; fi70++ {
        if true {
            switch (fi70 * (-(fi70))) {
            case -897:
                v71 := GenS5{F0: 248, F1: (^(v2))}
                _ = v71
            case 0:
                var v72 float64 = 0.0
                _ = v72
            default:
                if false {
                    for fi73 := 0; fi73 < 12; fi73++ {
                        var v74 float32 = 0.1
                        var v75 uint8 = (v2 ^ v2)
                        var v76 bool = false
                        switch ("\n\t" + (v5 + v5)) {
                        case "":
                            _ = len(v67)
                        case "a":
                            fmt.Println(v76)
                        case "\n\t":
                            _ = len(v67)
                        }
                        fmt.Println(fi73)
                        var v77 uint64 = 3
                        switch (v5 + (((v5 + v5) + ((v5 + "a") + "foo bar")) + ("foo bar" + ("foo bar" + ("a" + (v5 + v5)))))) {
                        case "foo bar":
                            v78 := GenS6{F0: -485, F1: v2}
                            _ = v78
                        }
                        _ = fi73
                        _ = v74
                        _ = v75
                        _ = v76
                        _ = v77
                    }
                    for fi79 := 0; fi79 < 9; fi79++ {
                        var v80 uint32 = 2147483648
                        var v81 uint = 18446744073709551615
                        fmt.Println(v3, v1)
                        _ = fi79
                        _ = v80
                        _ = v81
                    }
                    v82 := GenS7{F0: (v2 - (v2 + v2)), F1: false, F2: (v3 / ((v3) | 1))}
                    _ = v82
                }
            }
            var v83 uint64 = 2
            _ = v83
        }
        _ = fi70
    }
    v1 -= v1
    v1 += (v1 / (((v1 ^ v1)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v67:", len(v67), v67)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v67
    _ = v68
    _ = v69
}
