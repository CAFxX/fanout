package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 bool; F2 int32 }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 int64; F1 int8; F2 int8 }
func genFn0(p0 int) bool {
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := GenS1{F0: 2}
        _ = fi4
        _ = v5
    }
    p0 -= -301
    fmt.Println(p0)
    var v6 string = ("a" + "foo bar")
    _ = p0
    _ = v6
    return true
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]int)
    var v3 float32 = float32(-2.25)
    var v7 uint = 770
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := []int{fi8, (fi8 % ((fi8) | 1))}
        _ = fi8
        _ = v9
    }
    var v10 int64 = 0
    v7 = v7
    v11 := GenS1{F0: (v7 % (((v7 % ((v7) | 1))) | 1))}
    if (true || (42 >= v10)) {
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 float64 = (1.5 + 0.1)
            v14 := []int{68}
            v14 = append(v14, fi12)
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 int64 = 42
        if (0 > -1) {
            if (2 >= 1) {
                if false {
                    for fi16 := 0; fi16 < 3; fi16++ {
                        fmt.Println(fi16, fi16, v7)
                        for fi17 := 0; fi17 < 8; fi17++ {
                            v18 := GenS1{F0: v7}
                            delete(v2, fi16)
                            var v19 uint32 = 4294967295
                            var v20 int8 = -128
                            v20 += v20
                            _ = fi17
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        var v21 int16 = 0
                        _ = fi16
                        _ = v21
                    }
                    switch -61 {
                    case -2147483648:
                        for fi22 := 0; fi22 < 7; fi22++ {
                            for fi23 := 0; fi23 < 10; fi23++ {
                                if true {
                                    var v24 string = ("" + ("hello" + ("\n\t" + "foo bar")))
                                    var v25 uint = (v7 >> ((v7) & 63))
                                    v25 = 949
                                    _ = v24
                                    _ = v25
                                } else {
                                    switch -581 {
                                    case -1:
                                        if false {
                                            switch (fi22 % ((fi23) | 1)) {
                                            case -159:
                                                var v26 int = (fi22 / ((fi23) | 1))
                                                _ = v26
                                            default:
                                                fmt.Println(v15)
                                            }
                                            v7 = v7
                                            v27 := GenS2{F0: 425}
                                            _ = v27
                                        } else {
                                            switch "a" {
                                            case "\n\t":
                                                switch fi23 {
                                                case 0:
                                                    v3 = v3
                                                case -377:
                                                    v10 += v10
                                                case 32768:
                                                    v28 := GenS1{F0: v7}
                                                    _ = v28
                                                default:
                                                    if true {
                                                        v15 -= (^(v15))
                                                        v10 = int64(fi23)
                                                        v29 := GenS1{F0: (v7 | (v7 | v7))}
                                                        v15 = 200
                                                        var v30 float32 = v3
                                                        _ = v29
                                                        _ = v30
                                                    } else {
                                                        fmt.Println(fi22, v15)
                                                    }
                                                }
                                            case "hello":
                                                var v31 int64 = (v15 * (v15 | v10))
                                                _ = v31
                                            case "foo bar":
                                                v7 += (v7 / ((1) | 1))
                                            }
                                        }
                                    }
                                    fmt.Println(v15)
                                    _ = len(v2)
                                }
                                v10 = 9223372036854775807
                                v32 := []int{fi23, 42}
                                _ = fi23
                                _ = v32
                            }
                            if false {
                                v3 += v3
                                for fi33 := 0; fi33 < 18; fi33++ {
                                    var v34 int32 = 2147483647
                                    if false {
                                        if true {
                                            fmt.Println(fi33, v10)
                                            v15 = (v10 >> (((v7 % ((v7) | 1))) & 63))
                                        } else {
                                            var v35 uint16 = 223
                                            _ = v35
                                        }
                                        if (42 != 18446744073709551615) {
                                            v36 := GenS3{F0: 140}
                                            switch "foo bar" {
                                            case "hello":
                                                var v37 string = ""
                                                _ = v37
                                            case "foo bar":
                                                switch "a" {
                                                case "a":
                                                    v7 = 389
                                                }
                                            case "a":
                                                v7 = 42
                                            default:
                                                var v38 uint = v7
                                                _ = v38
                                            }
                                            _ = v36
                                        } else {
                                            switch -301 {
                                            case -317:
                                                v34 -= (v34 | v34)
                                            case 32768:
                                                var v39 float32 = 0.0
                                                _ = v39
                                            }
                                            var v40 int16 = 869
                                            v3 = v3
                                            _ = v40
                                        }
                                    } else {
                                        v41 := GenS4{F0: 4294967295}
                                        fmt.Println(v3, v3)
                                        v42 := GenS2{F0: (-(v34))}
                                        _ = v41
                                        _ = v42
                                    }
                                    _ = fi33
                                    _ = v34
                                }
                                switch "foo bar" {
                                case "foo bar":
                                    _ = len(v1)
                                default:
                                    switch fi22 {
                                    case -1:
                                        for fi43 := 0; fi43 < 2; fi43++ {
                                            _ = len(v2)
                                            _ = fi43
                                        }
                                    }
                                }
                            }
                            if true {
                                v44 := GenS1{F0: v7}
                                v7 += 408
                                v45 := []int{}
                                _ = v44
                                _ = v45
                            }
                            _ = fi22
                        }
                    case 42:
                        switch 297 {
                        case 9223372036854775807:
                            if true {
                                for fi46 := 0; fi46 < 11; fi46++ {
                                    v1[""] = ((("foo bar" + "") == "\n\t") && true)
                                    for fi47 := 0; fi47 < 20; fi47++ {
                                        if (!((885 <= (-(fi46))) && (true || ("foo bar" >= "")))) {
                                            v48 := GenS5{F0: 1}
                                            v10 = (v10 << ((v7) & 63))
                                            var v49 int64 = -126
                                            _ = v48
                                            _ = v49
                                        } else {
                                            var v50 uint = 0
                                            var v51 int8 = -1
                                            _ = v50
                                            _ = v51
                                        }
                                        v7 = v7
                                        _ = fi47
                                    }
                                    v52 := v1["foo bar"]
                                    _ = fi46
                                    _ = v52
                                }
                                v53 := GenS1{F0: v7}
                                if false {
                                    switch ("" + ((("foo bar" + ("" + "\n\t")) + ("foo bar" + "")) + "foo bar")) {
                                    case "foo bar":
                                        for fi54 := 0; fi54 < 8; fi54++ {
                                            v3 = float32(-2.25)
                                            v55 := GenS6{F0: ("a" + ""), F1: (!true), F2: -1}
                                            for fi56 := 0; fi56 < 13; fi56++ {
                                                v7 += (v7 >> ((v7) & 63))
                                                fmt.Println(v7, v7)
                                                if (((fi54 <= (fi54 << ((v7) & 63))) || (!((1e10 / 1.0) != -2.25))) && (true && true)) {
                                                    fmt.Println(v3)
                                                    v57 := v1[(("foo bar" + "hello") + "\n\t")]
                                                    v7 -= (v7 & v7)
                                                    _ = v57
                                                } else {
                                                    var v58 uint = (v7 ^ v7)
                                                    v10 += v10
                                                    v59 := []int{fi56}
                                                    v60 := GenS7{F0: "", F1: true}
                                                    v61 := GenS4{F0: fi56}
                                                    _ = v58
                                                    _ = v59
                                                    _ = v60
                                                    _ = v61
                                                }
                                                _ = fi56
                                            }
                                            _ = fi54
                                            _ = v55
                                        }
                                    case "hello":
                                        var v62 string = ""
                                        _ = v62
                                    }
                                }
                                _ = v53
                            } else {
                                var v63 string = "\n\t"
                                _ = v63
                            }
                        case -921:
                            var v64 uint8 = 42
                            _ = v64
                        case 149:
                            var v65 float32 = v3
                            _ = v65
                        default:
                            if false {
                                var v66 int = 128
                                var v67 uint64 = 42
                                for fi68 := 0; fi68 < 10; fi68++ {
                                    v69 := []int{int(v7), v66, fi68}
                                    _ = fi68
                                    _ = v69
                                }
                                _ = v66
                                _ = v67
                            }
                        }
                    default:
                        var v70 int = 9223372036854775807
                        _ = v70
                    }
                } else {
                    for fi71 := 0; fi71 < 12; fi71++ {
                        var v72 int32 = -1
                        _ = fi71
                        _ = v72
                    }
                    v7 = 159
                    v3 = v3
                }
            }
        } else {
            _ = len(v1)
            switch -866 {
            case 9223372036854775807:
                for fi73 := 0; fi73 < 10; fi73++ {
                    v10 = v15
                    v15 += (v15 - v10)
                    _ = fi73
                }
            }
            var v74 string = "\n\t"
            _ = v74
        }
        _ = v15
    } else {
        fmt.Println(v10)
        for fi75 := 0; fi75 < 20; fi75++ {
            var v76 uint64 = 2
            _ = fi75
            _ = v76
        }
    }
    for fi77 := 0; fi77 < 9; fi77++ {
        fmt.Println(v10)
        v78 := []int{127, 1, (fi77 >> (((v7 - v7)) & 63))}
        _ = fi77
        _ = v78
    }
    for fi79 := 0; fi79 < 15; fi79++ {
        v80 := []int{128}
        _ = fi79
        _ = v80
    }
    v7 = 4294967295
    switch 128 {
    case 128:
        for fi81 := 0; fi81 < 7; fi81++ {
            var v82 int32 = -1
            for fi83 := 0; fi83 < 10; fi83++ {
                v82 = (v82 & -809)
                _ = fi83
            }
            v84 := GenS8{F0: (v10 & v10), F1: 42, F2: -105}
            _ = fi81
            _ = v82
            _ = v84
        }
    case 273:
        switch 650 {
        case -887:
            var v85 uint64 = 0
            _ = v85
        }
    }
    var v86 uint64 = 18446744073709551615
    var v87 uint32 = 2147483647
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v7
    _ = v86
    _ = v87
}
