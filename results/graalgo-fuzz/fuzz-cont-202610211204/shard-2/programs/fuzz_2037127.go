package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 int32; F2 uint }
type GenS6 struct { F0 uint; F1 int; F2 uint64 }
type GenS7 struct { F0 uint; F1 int64 }
type GenS8 struct { F0 bool; F1 uint64; F2 string }
type GenS9 struct { F0 int16 }

func main() {
    var v1 float32 = 1e10
    var v2 uint32 = 65535
    var v3 float64 = (1e10 / 1.5)
    var v4 int = 42
    v3 -= ((v3 + v3) - v3)
    var v5 uint8 = 2
    var v6 int16 = -1
    var v7 int = -389
    for fi8 := 0; fi8 < 16; fi8++ {
        v4 = -1
        _ = fi8
    }
    v9 := []int{(v7 << ((773) & 63)), v7, -978, v7}
    if (true && true) {
        var v10 string = ("hello" + "foo bar")
        _ = v10
    }
    if ((v6 >> ((2) & 63)) >= v6) {
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := GenS1{F0: 988, F1: true}
            var v13 bool = true
            _ = fi11
            _ = v12
            _ = v13
        }
        switch v7 {
        case -2147483648:
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
        var v14 uint16 = 0
        _ = v14
    }
    if false {
        if len(v9) > 0 {
            v9[0] = (v7 >> ((65535) & 63))
        }
        v1 = v1
        if len(v9) > 0 {
            v9[0] = v7
        }
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        switch 32768 {
        case 9223372036854775807:
            v16 := GenS2{F0: v5, F1: 2, F2: -924}
            _ = v16
        case 42:
            if false {
                var v17 uint8 = 38
                fmt.Println(fi15, v17)
                _ = v17
            }
        default:
            if ("\n\t" != ("" + ("a" + "a"))) {
                v18 := GenS3{F0: (v6 + v6), F1: 65535}
                v6 = (v6 + 1000)
                _ = v18
            } else {
                if false {
                    v19 := GenS1{F0: 2, F1: false}
                    switch "foo bar" {
                    case "\n\t":
                        for fi20 := 0; fi20 < 11; fi20++ {
                            v2 = (v2 % ((v2) | 1))
                            v6 = -703
                            _ = fi20
                        }
                    case "a":
                        for fi21 := 0; fi21 < 1; fi21++ {
                            v7 += (fi15 / ((fi21) | 1))
                            if true {
                                fmt.Println(v5)
                                v22 := GenS3{F0: v6, F1: 255}
                                v2 -= 0
                                _ = v22
                            }
                            v3 = 0.1
                            _ = len(v9)
                            _ = fi21
                        }
                    case "":
                        for fi23 := 0; fi23 < 10; fi23++ {
                            v2 = v2
                            var v24 int64 = -487
                            fmt.Println(v6)
                            _ = fi23
                            _ = v24
                        }
                    }
                    _ = v19
                }
                v5 = v5
            }
        }
        _ = fi15
    }
    fmt.Println(v2, v6)
    for fi25 := 0; fi25 < 3; fi25++ {
        v6 += 7
        v4 = (v7 / (((^(fi25))) | 1))
        var v26 bool = (-128 != -42)
        _ = fi25
        _ = v26
    }
    var v27 bool = ((true && false) || false)
    v28 := GenS1{F0: 18446744073709551615, F1: (40 >= -80)}
    if v27 {
        if (!v27) {
            var v29 float32 = 1e10
            if false {
                switch (("foo bar" + ("" + "")) + "") {
                case "\n\t":
                    if len(v9) > 0 {
                        v9[0] = v7
                    }
                case "a":
                    for fi30 := 0; fi30 < 7; fi30++ {
                        for fi31 := 0; fi31 < 17; fi31++ {
                            if len(v9) > 0 {
                                v9[0] = (^(v4))
                            }
                            var v32 uint16 = 65535
                            switch "\n\t" {
                            case "foo bar":
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                            case "\n\t":
                                fmt.Println(v32, fi31)
                            case "hello":
                                var v33 bool = (v5 >= 255)
                                _ = v33
                            default:
                                v1 -= v1
                            }
                            if (!(1 != 4294967295)) {
                                if (v2 >= (^(v2))) {
                                    v4 -= fi31
                                    v34 := GenS1{F0: 759, F1: (!(false || (826 != uint64(v2))))}
                                    var v35 uint16 = 737
                                    v36 := GenS4{F0: v6}
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                }
                                v29 += v1
                                v2 += 2147483647
                            }
                            v37 := GenS5{F0: 674, F1: -448, F2: 9}
                            _ = fi31
                            _ = v32
                            _ = v37
                        }
                        v38 := GenS6{F0: 42, F1: v4, F2: 42}
                        v39 := GenS7{F0: 18446744073709551615, F1: 761}
                        _ = fi30
                        _ = v38
                        _ = v39
                    }
                case "foo bar":
                    var v40 uint16 = 2
                    _ = v40
                }
            } else {
                v5 -= v5
            }
            v41 := GenS5{F0: 65535, F1: 491, F2: 647}
            _ = v29
            _ = v41
        }
    } else {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        v3 += ((v3 / v3) * ((v3 / 3.14159) + 0.0))
    }
    if false {
        v27 = (!(!v27))
        v4 += v4
    } else {
        for fi42 := 0; fi42 < 12; fi42++ {
            v43 := GenS8{F0: (v27 && (v4 < (fi42 * 9223372036854775807))), F1: 42, F2: "hello"}
            v3 = float64((float32(v3) / v1))
            var v44 int8 = 127
            _ = fi42
            _ = v43
            _ = v44
        }
        v9 = append(v9, (v7 * v4))
        for fi45 := 0; fi45 < 3; fi45++ {
            var v46 int8 = -128
            _ = fi45
            _ = v46
        }
    }
    for fi47 := 0; fi47 < 5; fi47++ {
        v27 = ((^(v5)) > uint8(v2))
        for fi48 := 0; fi48 < 13; fi48++ {
            if (!true) {
                var v49 uint = 0
                v50 := GenS9{F0: v6}
                _ = v49
                _ = v50
            } else {
                var v51 uint16 = 391
                _ = v51
            }
            _ = len(v9)
            _ = fi48
        }
        _ = fi47
    }
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
