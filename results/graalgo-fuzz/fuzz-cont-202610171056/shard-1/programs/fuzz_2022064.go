package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint16; F1 int16; F2 int8 }
type GenS5 struct { F0 uint64; F1 bool; F2 int8 }
type GenS6 struct { F0 int }
func genFn0(p0 uint8) string {
    var v4 uint16 = 65535
    for fi5 := 0; fi5 < 10; fi5++ {
        fmt.Println(p0)
        var v6 int8 = 0
        for fi7 := 0; fi7 < 6; fi7++ {
            p0 = p0
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return (("\n\t" + "") + ("" + ("foo bar" + "\n\t")))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 499
    var v3 uint8 = 128
    if true {
        var v8 int16 = 499
        fmt.Println(v8, v3, v3)
        _ = v8
    } else {
        var v9 int16 = 42
        for fi10 := 0; fi10 < 17; fi10++ {
            v9 -= v9
            _ = fi10
        }
        v9 = v9
        _ = v9
    }
    v11 := GenS1{F0: v2}
    v2 = 1
    v12 := GenS1{F0: 1}
    fmt.Println(v2, v2)
    var v13 string = ("" + ("\n\t" + ""))
    if (471 == 2) {
        if (2147483647 != 2147483647) {
            if true {
                v2 = v2
                v14 := []int{2147483647, 0, 9223372036854775807}
                fmt.Println(v13, v13)
                _ = v14
            } else {
                switch "foo bar" {
                case "foo bar":
                    v13 = v13
                case "a":
                    v3 = v3
                case "":
                    var v15 uint64 = v2
                    _ = v15
                }
                var v16 float64 = (1.5 / ((3.14159 / (3.14159 / 3.14159)) * 0.1))
                fmt.Println(v2, v3, v3)
                _ = v16
            }
        } else {
            for fi17 := 0; fi17 < 5; fi17++ {
                v3 += 127
                if (18446744073709551615 > 785) {
                    v13 = "\n\t"
                }
                v13 = ("a" + (v13 + v13))
                _ = fi17
            }
            switch v13 {
            case "hello":
                var v18 int32 = -2147483648
                _ = v18
            }
            for fi19 := 0; fi19 < 10; fi19++ {
                if ((fi19 << ((2) & 63)) > fi19) {
                    fmt.Println(fi19)
                }
                _ = fi19
            }
        }
        if true {
            v20 := GenS1{F0: v2}
            if ((1 <= 42) || (4294967295 == -107)) {
                v21 := GenS1{F0: v2}
                v2 += v2
                if true {
                    v22 := GenS1{F0: (v2 << ((567) & 63))}
                    var v23 uint = 701
                    _ = v22
                    _ = v23
                }
                _ = v21
            }
            _ = v20
        }
        v24 := GenS2{F0: -892}
        _ = v24
    } else {
        if false {
            v2 += v2
        }
    }
    v25 := GenS1{F0: v2}
    if (!false) {
        for fi26 := 0; fi26 < 2; fi26++ {
            if ((127 == 42) || (v2 <= 1)) {
                var v27 bool = (!(3.14159 != 1.5))
                for fi28 := 0; fi28 < 5; fi28++ {
                    v13 = v13
                    _ = fi28
                }
                for fi29 := 0; fi29 < 9; fi29++ {
                    if (v27 && (fi29 < -280)) {
                        if true {
                            fmt.Println(fi29, fi29, v3)
                            if (v3 != 0) {
                                v3 = (v3 >> ((42) & 63))
                                v2 = (v2 % (((v2 ^ (v2 * (v2 + (^(v2)))))) | 1))
                                v30 := GenS3{F0: ((fi29 ^ fi29) != (fi26 ^ -377)), F1: (v2 == (v2 >> ((0) & 63)))}
                                _ = v30
                            }
                        }
                        v31 := []int{fi29, (fi26 | 2147483647), 1}
                        fmt.Println(v3, v2)
                        _ = v31
                    } else {
                        if v27 {
                            if v27 {
                                switch 2147483647 {
                                case 0:
                                    var v32 string = v13
                                    _ = v32
                                case 4294967295:
                                    v33 := GenS3{F0: v27, F1: (0.0 == 1.5)}
                                    _ = v33
                                }
                                v34 := GenS2{F0: fi29}
                                switch ("hello" + (v13 + "\n\t")) {
                                case "foo bar":
                                    if (!v27) {
                                        v35 := GenS3{F0: (!v27), F1: v27}
                                        v36 := GenS4{F0: 65535, F1: 0, F2: 1}
                                        _ = v35
                                        _ = v36
                                    }
                                case "\n\t":
                                    v37 := []int{(fi29 + 127), (fi26 >> ((uint(v2)) & 63))}
                                    _ = v37
                                case "a":
                                    v38 := []int{fi29, fi26, (^(fi29)), fi26}
                                    _ = v38
                                }
                                _ = v34
                            }
                            var v39 uint32 = 138
                            switch fi29 {
                            case 42:
                                if (false && (v13 > (v13 + v13))) {
                                    var v40 string = "\n\t"
                                    var v41 uint64 = v2
                                    switch fi26 {
                                    case 514:
                                        if (!((^(v39)) <= (v39 * v39))) {
                                            v42 := GenS4{F0: 86, F1: -32768, F2: 100}
                                            _ = v42
                                        }
                                    case 9223372036854775807:
                                        fmt.Println(v41, v40, v41)
                                    }
                                    _ = v40
                                    _ = v41
                                }
                            }
                            _ = v39
                        }
                        v27 = (2147483647 > 544)
                        if (uint32(v3) == 0) {
                            var v43 uint = 31
                            _ = v43
                        }
                        delete(v1, fi29)
                        v44 := []int{fi29, -9223372036854775808, fi29}
                        _ = v44
                    }
                    fmt.Println(v3, v27, fi26)
                    if ((((0.0 + float64((-2.25 + 0.1))) + 3.14159) / (((float64((1.5 + (0.0 / 1.5))) + 3.14159) - 0.0) + 3.14159)) > 0.0) {
                        fmt.Println(v27, fi26, fi29)
                        v3 = v3
                        var v45 uint32 = 4294967295
                        if v27 {
                            v27 = (641 != 355)
                            v46 := []int{}
                            if false {
                                switch v13 {
                                case "foo bar":
                                    var v47 bool = (701 >= fi26)
                                    _ = v47
                                }
                                var v48 int64 = 9223372036854775807
                                v49 := GenS2{F0: 982}
                                _ = v48
                                _ = v49
                            }
                            _ = v46
                        } else {
                            var v50 string = ((v13 + v13) + v13)
                            _ = v50
                        }
                        _ = v45
                    }
                    v13 = ("foo bar" + v13)
                    if (0.0 == (0.1 * 0.1)) {
                        switch (v13 + v13) {
                        case "\n\t":
                            _ = len(v1)
                        default:
                            switch ("foo bar" + ((v13 + "foo bar") + v13)) {
                            case "hello":
                                var v51 uint64 = v2
                                _ = v51
                            default:
                                v52 := GenS4{F0: 0, F1: -32768, F2: 127}
                                _ = v52
                            }
                        }
                    }
                    _ = fi29
                }
                _ = v27
            } else {
                var v53 string = v13
                _ = v53
            }
            var v54 bool = (2147483647 < 571)
            _ = fi26
            _ = v54
        }
    }
    var v55 uint64 = uint64(v3)
    switch 32768 {
    case 128:
        v56 := GenS5{F0: 571, F1: (42 == v2), F2: 0}
        _ = v56
    }
    if (626 <= 0) {
        v57 := GenS3{F0: false, F1: (("" + v13) >= ("hello" + v13))}
        var v58 string = ""
        _ = v57
        _ = v58
    }
    v59 := GenS6{F0: 1}
    v3 = uint8(v2)
    delete(v1, -946)
    v60 := []int{-64, 673}
    if (2 <= v3) {
        for fi61 := 0; fi61 < 6; fi61++ {
            if (65535 != 665) {
                if ("hello" <= ("" + "a")) {
                    if len(v60) >= 2 {
                        v60 = v60[1:len(v60)]
                    }
                }
                for fi62 := 0; fi62 < 1; fi62++ {
                    var v63 int32 = 42
                    v63 = (-(v63))
                    if true {
                        var v64 int32 = v63
                        _ = v64
                    } else {
                        v3 = (^(v3))
                    }
                    fmt.Println(v13, v3, v13)
                    if ((v2 << ((4294967295) & 63)) <= (v2 >> ((652) & 63))) {
                        v55 = v2
                    } else {
                        if len(v60) >= 2 {
                            v60 = v60[1:len(v60)]
                        }
                        v55 -= v2
                    }
                    _ = len(v60)
                    v60 = append(v60, -9223372036854775808)
                    var v65 int64 = -1
                    _ = fi62
                    _ = v63
                    _ = v65
                }
            } else {
                for fi66 := 0; fi66 < 9; fi66++ {
                    if false {
                        var v67 string = v13
                        var v68 int32 = 820
                        var v69 uint64 = (v55 | (v55 & 2))
                        var v70 int16 = 42
                        _ = v67
                        _ = v68
                        _ = v69
                        _ = v70
                    }
                    v2 = 789
                    _ = fi66
                }
                if len(v60) > 0 {
                    v60[0] = fi61
                }
            }
            _ = fi61
        }
        v2 = (^(v55))
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v60:", len(v60), v60)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v55
    _ = v59
    _ = v60
}
