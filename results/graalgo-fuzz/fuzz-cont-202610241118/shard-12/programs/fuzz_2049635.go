package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int32 }
type GenS2 struct { F0 int32; F1 int8; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 int32; F1 string; F2 int8 }
type GenS7 struct { F0 int8; F1 int8 }
func genFn0(p0 int8, p1 uint64) uint16 {
    var v4 bool = false
    fmt.Println(p1)
    v5 := GenS1{F0: v4, F1: (((!(255 == 0)) || (!(p0 < (p0 % ((0) | 1))))) || false), F2: -870}
    var v6 uint64 = p1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 65535
}
func genFn1() int32 {
    var v7 uint8 = 16
    v8 := GenS1{F0: ((!true) || (506 < 2)), F1: false, F2: -2147483648}
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 uint16 = 536
        if (-2.25 < (1e10 - (1.5 + 0.1))) {
            if (!false) {
                switch int(v10) {
                case 0:
                    v11 := []int{fi9, -115, 32767, (fi9 % ((fi9) | 1))}
                    _ = v11
                default:
                    for fi12 := 0; fi12 < 2; fi12++ {
                        var v13 uint16 = v10
                        if true {
                            v10 = v13
                            v14 := []int{(fi12 % (((fi12 + 9223372036854775807)) | 1)), -244, fi12}
                            if (false || (false && true)) {
                                v10 = 1
                                v15 := GenS2{F0: -1, F1: -121, F2: ("\n\t" == "\n\t")}
                                var v16 float32 = 1.5
                                v10 += (v13 & v10)
                                v7 = (v7 % ((v7) | 1))
                                _ = v15
                                _ = v16
                            } else {
                                if true {
                                    v17 := GenS1{F0: (v7 >= (v7 << ((2) & 63))), F1: true, F2: 42}
                                    v18 := GenS3{F0: v10}
                                    if false {
                                        v13 = (v13 + 42)
                                        fmt.Println(v7, v7)
                                        fmt.Println(v7, v10, v7)
                                    } else {
                                        var v19 uint64 = 4294967295
                                        var v20 int8 = 42
                                        _ = v19
                                        _ = v20
                                    }
                                    if len(v14) > 0 {
                                        v14[0] = fi12
                                    }
                                    _ = v17
                                    _ = v18
                                } else {
                                    switch ("" + "") {
                                    case "foo bar":
                                        if false {
                                            var v21 uint16 = (v13 + 32767)
                                            var v22 int8 = 2
                                            _ = v21
                                            _ = v22
                                        }
                                    case "":
                                        if (!(true && true)) {
                                            v10 = (v10 & 65535)
                                            if (false || false) {
                                                v23 := GenS3{F0: (^(v10))}
                                                var v24 uint16 = (v13 << ((1) & 63))
                                                _ = v23
                                                _ = v24
                                            } else {
                                                if false {
                                                    fmt.Println(fi12, v10)
                                                }
                                                v7 -= v7
                                                var v25 int16 = -1
                                                _ = v25
                                            }
                                            fmt.Println(v7)
                                        } else {
                                            if (!true) {
                                                v26 := GenS1{F0: ((!true) || false), F1: (!true), F2: 244}
                                                var v27 uint16 = (v13 & genFn0(1, 87))
                                                var v28 int64 = -9223372036854775808
                                                _ = v26
                                                _ = v27
                                                _ = v28
                                            } else {
                                                v14 = append(v14, 0)
                                            }
                                            var v29 float64 = -2.25
                                            _ = v29
                                        }
                                    default:
                                        v30 := GenS1{F0: true, F1: false, F2: -692}
                                        _ = v30
                                    }
                                }
                            }
                            _ = v14
                        }
                        _ = fi12
                        _ = v13
                    }
                }
            } else {
                for fi31 := 0; fi31 < 5; fi31++ {
                    v32 := GenS2{F0: int32(v7), F1: 127, F2: true}
                    _ = fi31
                    _ = v32
                }
                switch "a" {
                case "foo bar":
                    v10 = genFn0(-37, 18446744073709551615)
                case "":
                    fmt.Println(v7, v7)
                }
                var v33 int8 = -1
                _ = v33
            }
        }
        for fi34 := 0; fi34 < 13; fi34++ {
            v7 = (v7 << ((65535) & 63))
            _ = fi34
        }
        _ = fi9
        _ = v10
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        var v36 int8 = 42
        fmt.Println(fi35, v7)
        for fi37 := 0; fi37 < 10; fi37++ {
            v7 -= v7
            var v38 int = 1
            _ = fi37
            _ = v38
        }
        _ = fi35
        _ = v36
    }
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    var v1 uint8 = 128
    var v2 uint64 = 868
    var v3 int = 1
    fmt.Println(v1)
    if ("foo bar" > "foo bar") {
        if true {
            if false {
                v39 := GenS4{F0: 1, F1: -391, F2: 411}
                for fi40 := 0; fi40 < 12; fi40++ {
                    if false {
                        switch v3 {
                        case -925:
                            fmt.Println(v2, v1)
                        }
                        if true {
                            var v41 int32 = -2147483648
                            var v42 int = -690
                            _ = v41
                            _ = v42
                        } else {
                            for fi43 := 0; fi43 < 4; fi43++ {
                                var v44 int16 = 727
                                var v45 uint8 = v1
                                v2 = (^(v2))
                                v46 := []int{(fi43 | fi40)}
                                _ = fi43
                                _ = v44
                                _ = v45
                                _ = v46
                            }
                        }
                    }
                    _ = fi40
                }
                _ = v39
            }
            var v47 string = "hello"
            v48 := GenS5{F0: (^(v1)), F1: v47}
            _ = v47
            _ = v48
        }
    } else {
        switch "\n\t" {
        case "hello":
            var v49 float64 = 0.0
            _ = v49
        case "":
            var v50 float64 = -2.25
            _ = v50
        case "foo bar":
            var v51 float64 = float64((-2.25 / (3.14159 - (0.1 - float32(0.1)))))
            _ = v51
        }
        v52 := GenS5{F0: v1, F1: ""}
        switch ((("a" + ("\n\t" + ("foo bar" + "foo bar"))) + "hello") + ("" + "foo bar")) {
        case "a":
            v53 := []int{v3, (v3 * (v3 + (^(v3)))), v3}
            _ = v53
        case "hello":
            if (true || true) {
                for fi54 := 0; fi54 < 10; fi54++ {
                    v55 := []int{(fi54 - 42), (fi54 ^ 32767)}
                    for fi56 := 0; fi56 < 14; fi56++ {
                        v57 := GenS2{F0: 0, F1: -47, F2: false}
                        _ = fi56
                        _ = v57
                    }
                    if true {
                        if (v3 <= (v3 >> ((1) & 63))) {
                            v58 := GenS6{F0: 870, F1: "\n\t", F2: 127}
                            for fi59 := 0; fi59 < 10; fi59++ {
                                v2 -= v2
                                var v60 uint16 = genFn0(0, (v2 + v2))
                                fmt.Println(fi59, fi54, fi59)
                                _ = fi59
                                _ = v60
                            }
                            _ = v58
                        }
                        v3 -= int(v1)
                        var v61 int32 = 1
                        _ = v61
                    }
                    _ = fi54
                    _ = v55
                }
                switch "\n\t" {
                case "foo bar":
                    var v62 uint32 = 0
                    _ = v62
                case "":
                    var v63 uint8 = (v1 + v1)
                    _ = v63
                case "hello":
                    switch (v3 % ((128) | 1)) {
                    case -2147483648:
                        v64 := []int{v3, v3}
                        _ = v64
                    case -32:
                        for fi65 := 0; fi65 < 5; fi65++ {
                            v3 += (v3 % ((fi65) | 1))
                            _ = fi65
                        }
                    }
                default:
                    var v66 uint = 4294967295
                    _ = v66
                }
                v1 += 191
            }
        default:
            if (((42 == 1) || (18446744073709551615 < 2)) && (((float64(1.5) / 1.0) - 0.0) < 3.14159)) {
                v67 := []int{(v3 | v3), (^(v3)), (-(v3))}
                v3 -= -1
                fmt.Println(v2, v2, v2)
                _ = v67
            } else {
                for fi68 := 0; fi68 < 2; fi68++ {
                    if true {
                        v3 = v3
                    }
                    var v69 float32 = -2.25
                    if false {
                        var v70 float32 = float32(0.0)
                        var v71 float32 = float32(1.5)
                        v3 = v3
                        _ = v70
                        _ = v71
                    }
                    _ = fi68
                    _ = v69
                }
                fmt.Println(v2, v3, v1)
                var v72 int64 = -427
                _ = v72
            }
        }
        _ = v52
    }
    v73 := GenS7{F0: 0, F1: -128}
    v1 -= 207
    var v74 float64 = float64(-2.25)
    v75 := GenS1{F0: false, F1: (771 <= genFn1()), F2: genFn1()}
    if false {
        v76 := []int{v3}
        _ = v76
    } else {
        switch (v3 * 639) {
        case 32768:
            for fi77 := 0; fi77 < 7; fi77++ {
                v78 := GenS1{F0: (((("foo bar" + (((("hello" + ("" + "")) + "foo bar") + "hello") + "foo bar")) + "\n\t") + "a") <= ""), F1: (42 < 42), F2: genFn1()}
                _ = fi77
                _ = v78
            }
        }
        for fi79 := 0; fi79 < 9; fi79++ {
            var v80 float64 = (v74 - float64(0.0))
            v81 := GenS3{F0: genFn0(127, (v2 | (v2 & (v2 << ((1) & 63)))))}
            _ = fi79
            _ = v80
            _ = v81
        }
        v3 = (v3 % ((v3) | 1))
    }
    v74 = (v74 - v74)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v73
    _ = v74
    _ = v75
}
