package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int64 }
type GenS2 struct { F0 int; F1 string; F2 bool }
type GenS3 struct { F0 string; F1 bool; F2 uint16 }
type GenS4 struct { F0 int16; F1 string; F2 int }

func main() {
    var v1 uint32 = 0
    var v2 float64 = (3.14159 - 0.0)
    v3 := make(map[string]int)
    var v4 uint8 = 0
    var v5 int8 = 1
    fmt.Println(v4, v2)
    if ((510 >= uint16(v4)) && (549 <= 959)) {
        for fi6 := 0; fi6 < 17; fi6++ {
            if false {
                var v7 int32 = -37
                _ = v7
            } else {
                if true {
                    v3["hello"] = (fi6 / (((fi6 << ((687) & 63))) | 1))
                }
            }
            switch fi6 {
            case -1:
                for fi8 := 0; fi8 < 5; fi8++ {
                    v1 = (v1 ^ (v1 << ((42) & 63)))
                    var v9 float64 = ((v2 * v2) + 0.1)
                    if false {
                        var v10 int8 = v5
                        v11 := GenS1{F0: false, F1: 42, F2: -467}
                        fmt.Println(v9)
                        _ = v10
                        _ = v11
                    }
                    _ = fi8
                    _ = v9
                }
            case 2147483647:
                var v12 uint16 = 2
                _ = v12
            }
            switch fi6 {
            case 128:
                fmt.Println(v4, v4)
            case 9223372036854775807:
                if false {
                    v13 := []int{fi6}
                    _ = v13
                }
            case 0:
                var v14 string = "foo bar"
                _ = v14
            default:
                if true {
                    v4 -= (v4 | (v4 << ((4294967295) & 63)))
                    for fi15 := 0; fi15 < 16; fi15++ {
                        v16 := v3["\n\t"]
                        switch ("\n\t" + ((("a" + ("\n\t" + ("hello" + (("hello" + ("foo bar" + "")) + (((("foo bar" + (("" + "\n\t") + "a")) + "foo bar") + "\n\t") + "hello"))))) + ("a" + "a")) + "\n\t")) {
                        case "a":
                            v3["foo bar"] = (fi15 & v16)
                        }
                        _ = fi15
                        _ = v16
                    }
                    v1 += v1
                } else {
                    if true {
                        for fi17 := 0; fi17 < 2; fi17++ {
                            switch "hello" {
                            case "\n\t":
                                v18 := []int{(fi6 & 237), fi6, (fi6 - (fi6 % (((-(fi6))) | 1))), 1}
                                _ = v18
                            case "hello":
                                fmt.Println(fi17, v2)
                            default:
                                v19 := GenS2{F0: (-(fi17)), F1: "\n\t", F2: false}
                                _ = v19
                            }
                            _ = fi17
                        }
                        if (v1 != v1) {
                            switch (fi6 * fi6) {
                            case 128:
                                v2 += 3.14159
                            case -9223372036854775808:
                                v20 := GenS1{F0: ((v1 >> ((18446744073709551615) & 63)) >= (v1 / (((^(v1))) | 1))), F1: 0, F2: -101}
                                _ = v20
                            default:
                                fmt.Println(v2, v5, v4)
                            }
                        }
                        if false {
                            for fi21 := 0; fi21 < 20; fi21++ {
                                delete(v3, "a")
                                _ = fi21
                            }
                            if false {
                                if (!false) {
                                    v1 = v1
                                    v22 := GenS2{F0: (fi6 << ((1) & 63)), F1: ("a" + "hello"), F2: (1 <= 332)}
                                    var v23 uint64 = 42
                                    _ = v22
                                    _ = v23
                                }
                                var v24 int32 = -2147483648
                                var v25 int16 = 383
                                _ = v24
                                _ = v25
                            }
                        } else {
                            v26 := []int{}
                            _ = v26
                        }
                    } else {
                        if (65535 != 203) {
                            var v27 string = ""
                            v28 := GenS1{F0: false, F1: 42, F2: 42}
                            for fi29 := 0; fi29 < 15; fi29++ {
                                switch fi6 {
                                case 128:
                                    v2 = v2
                                default:
                                    fmt.Println(fi6)
                                }
                                _ = fi29
                            }
                            _ = v27
                            _ = v28
                        }
                    }
                }
            }
            _ = fi6
        }
        v4 = (v4 + (v4 >> ((0) & 63)))
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        if false {
            v31 := []int{fi30, fi30}
            _ = v31
        }
        v32 := []int{}
        _ = fi30
        _ = v32
    }
    var v33 bool = (!(188 < 18446744073709551615))
    v34 := []int{}
    var v35 uint8 = (v4 - (^(v4)))
    if v33 {
        for fi36 := 0; fi36 < 16; fi36++ {
            for fi37 := 0; fi37 < 18; fi37++ {
                v38 := GenS3{F0: ("a" + ""), F1: ((v4 - v35) <= 255), F2: 899}
                v39 := GenS4{F0: 1000, F1: "\n\t", F2: fi36}
                _ = len(v34)
                switch "\n\t" {
                case "":
                    v4 -= (v35 / ((255) | 1))
                case "foo bar":
                    fmt.Println(v2, fi36)
                }
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = fi36
        }
        switch ((("foo bar" + ((("foo bar" + "foo bar") + ("hello" + "")) + "hello")) + "foo bar") + "\n\t") {
        case "hello":
            if v33 {
                v34 = append(v34, 0)
                v1 = v1
            } else {
                fmt.Println(v2)
                for fi40 := 0; fi40 < 5; fi40++ {
                    delete(v3, (("foo bar" + "\n\t") + ("a" + "a")))
                    v34 = append(v34, -2147483648)
                    v41 := GenS1{F0: false, F1: -32768, F2: 207}
                    _ = fi40
                    _ = v41
                }
            }
        case "a":
            _ = len(v3)
        default:
            if len(v34) >= 2 {
                v34 = v34[1:len(v34)]
            }
        }
        if v33 {
            if v33 {
                var v42 uint = 0
                _ = v42
            }
            var v43 int32 = 670
            v44 := GenS1{F0: v33, F1: 0, F2: 1}
            _ = v43
            _ = v44
        }
    }
    v33 = (!(1000 != 0))
    if len(v34) >= 2 {
        v34 = v34[1:len(v34)]
    }
    for fi45 := 0; fi45 < 13; fi45++ {
        if (32767 > fi45) {
            v33 = ((v33 && v33) || v33)
            fmt.Println(v4, v5, fi45)
        }
        var v46 uint64 = 42
        _ = fi45
        _ = v46
    }
    fmt.Println(v35, v5)
    for fi47 := 0; fi47 < 17; fi47++ {
        fmt.Println(v1, v5, v35)
        if ((^(fi47)) >= fi47) {
            var v48 int32 = 318
            _ = v48
        } else {
            _ = len(v3)
            fmt.Println(v33)
            if (0 > 32767) {
                var v49 bool = (0 > 1)
                var v50 uint64 = 42
                v35 = v35
                _ = v49
                _ = v50
            } else {
                var v51 float64 = (0.0 / (float64((3.14159 - (3.14159 + 0.1))) / v2))
                if len(v34) >= 2 {
                    v34 = v34[1:len(v34)]
                }
                _ = v51
            }
        }
        fmt.Println(v4, v1)
        _ = fi47
    }
    v5 = v5
    var v52 uint64 = 472
    for fi53 := 0; fi53 < 17; fi53++ {
        fmt.Println(fi53, fi53)
        v2 = v2
        _ = fi53
    }
    if len(v34) > 0 {
        v34[0] = -566
    }
    var v54 int16 = 711
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v52
    _ = v54
}
