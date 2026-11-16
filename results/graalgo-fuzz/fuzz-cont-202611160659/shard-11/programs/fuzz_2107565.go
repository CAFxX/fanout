package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 bool; F1 int; F2 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int32, p1 uint16) int32 {
    var v4 uint8 = 25
    var v5 bool = (1 != -1)
    for fi6 := 0; fi6 < 4; fi6++ {
        if v5 {
            fmt.Println(v5, p0, v4)
        }
        _ = fi6
    }
    v7 := []int{0, 4294967295}
    var v8 float64 = -2.25
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    return (p0 / (((p0 << ((4294967295) & 63))) | 1))
}

func main() {
    var v1 int8 = 62
    var v2 uint8 = 255
    var v3 uint16 = 0
    var v9 int64 = -1
    var v10 uint64 = 4294967295
    v10 = (v10 ^ uint64(v3))
    if true {
        var v11 bool = (!false)
        _ = v11
    }
    v10 = (v10 ^ (v10 << ((65535) & 63)))
    if ((v1 << ((901) & 63)) >= v1) {
        v3 -= v3
    }
    if ((0 < 1) && true) {
        if false {
            v3 -= v3
            if (!((v10 | (v10 - 194)) > 506)) {
                v12 := GenS1{F0: (v2 & v2), F1: 65535}
                switch ("a" + "hello") {
                case "a":
                    if true {
                        if false {
                            for fi13 := 0; fi13 < 13; fi13++ {
                                if (genFn0(42, (v3 ^ (v3 - v3))) >= 1) {
                                    switch (((("" + ("\n\t" + "\n\t")) + (("hello" + (("hello" + ("a" + ("\n\t" + "\n\t"))) + "\n\t")) + "a")) + "a") + "a") {
                                    case "\n\t":
                                        var v14 int16 = -1
                                        _ = v14
                                    }
                                }
                                switch (fi13 - fi13) {
                                case 412:
                                    for fi15 := 0; fi15 < 20; fi15++ {
                                        var v16 int16 = -568
                                        var v17 uint16 = v3
                                        v18 := GenS2{F0: ((1e10 * (3.14159 * ((0.0 / ((0.1 - 0.1) + 3.14159)) * -2.25))) != 0.1), F1: 42, F2: -1}
                                        _ = fi15
                                        _ = v16
                                        _ = v17
                                        _ = v18
                                    }
                                case -696:
                                    v19 := GenS2{F0: false, F1: int(v9), F2: 1}
                                    _ = v19
                                }
                                _ = fi13
                            }
                        }
                        v3 -= v3
                        v20 := []int{32767, 32768, 32767, int(v1)}
                        _ = v20
                    } else {
                        v10 = (^(v10))
                    }
                case "":
                    var v21 int16 = 42
                    _ = v21
                default:
                    switch ("\n\t" + ("" + "\n\t")) {
                    case "a":
                        for fi22 := 0; fi22 < 10; fi22++ {
                            switch "hello" {
                            case "":
                                v9 = (v9 % ((v9) | 1))
                            case "a":
                                for fi23 := 0; fi23 < 14; fi23++ {
                                    var v24 int64 = -9223372036854775808
                                    _ = fi23
                                    _ = v24
                                }
                            default:
                                v9 -= v9
                            }
                            var v25 int16 = -630
                            fmt.Println(v9, v9)
                            _ = fi22
                            _ = v25
                        }
                    default:
                        var v26 float64 = 0.0
                        _ = v26
                    }
                }
                var v27 bool = true
                _ = v12
                _ = v27
            }
        } else {
            for fi28 := 0; fi28 < 15; fi28++ {
                v1 += 127
                v3 += v3
                if true {
                    fmt.Println(fi28, v10)
                    fmt.Println(v9, v3, v2)
                }
                _ = fi28
            }
            if true {
                for fi29 := 0; fi29 < 7; fi29++ {
                    if (!(true || (-792 != 0))) {
                        v1 = v1
                        v30 := GenS2{F0: ((false && true) || ((("" + ("a" + "hello")) + "\n\t") != "")), F1: 42, F2: -1}
                        _ = v30
                    }
                    v31 := []int{-197}
                    _ = fi29
                    _ = v31
                }
                v3 = 962
            } else {
                switch ("hello" + "foo bar") {
                case "":
                    v32 := []int{-1, -2147483648, -303, 2147483647}
                    _ = v32
                case "hello":
                    v33 := GenS1{F0: 1, F1: 217}
                    _ = v33
                default:
                    v34 := GenS3{F0: v10}
                    _ = v34
                }
                if false {
                    fmt.Println(v9, v3)
                }
            }
            var v35 uint32 = 1
            _ = v35
        }
        fmt.Println(v9, v2, v3)
    } else {
        var v36 int16 = 371
        if true {
            fmt.Println(v1)
        } else {
            v37 := []int{int(v2), -1, 546}
            _ = v37
        }
        v10 -= v10
        _ = v36
    }
    for fi38 := 0; fi38 < 19; fi38++ {
        v3 = v3
        v39 := []int{fi38, fi38, (fi38 ^ fi38), 412}
        for fi40 := 0; fi40 < 4; fi40++ {
            _ = len(v39)
            v9 = v9
            v41 := GenS2{F0: (605 < 600), F1: (fi38 >> ((255) & 63)), F2: -388}
            _ = fi40
            _ = v41
        }
        _ = fi38
        _ = v39
    }
    fmt.Println(v9)
    for fi42 := 0; fi42 < 9; fi42++ {
        fmt.Println(v1)
        for fi43 := 0; fi43 < 15; fi43++ {
            v44 := GenS4{F0: 63}
            if (-32768 >= 32767) {
                if false {
                    if false {
                        v45 := GenS1{F0: 140, F1: 428}
                        v46 := GenS5{F0: (v3 * 42)}
                        var v47 uint8 = (v2 & (^(v2)))
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    v10 = (v10 + v10)
                }
            }
            v2 -= (v2 - v2)
            _ = fi43
            _ = v44
        }
        for fi48 := 0; fi48 < 11; fi48++ {
            var v49 string = ("foo bar" + "\n\t")
            switch v49 {
            case "":
                v50 := GenS3{F0: (v10 - v10)}
                _ = v50
            case "\n\t":
                fmt.Println(v1, v2, v3)
            case "a":
                fmt.Println(fi42, v3)
            }
            _ = fi48
            _ = v49
        }
        _ = fi42
    }
    var v51 bool = ((true && (v1 != v1)) || true)
    v3 = v3
    v52 := GenS3{F0: (v10 % ((v10) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v51:", v51)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v51
    _ = v52
    _ = v9
}
