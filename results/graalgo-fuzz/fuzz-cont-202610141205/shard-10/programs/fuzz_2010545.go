package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int16; F1 int16; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 int8; F2 uint }
type GenS7 struct { F0 string; F1 int8 }
type GenS8 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint64) int8 {
    var v7 uint64 = (p0 & (p0 & (p0 >> ((65535) & 63))))
    v8 := []int{128, 127, 32768, -412}
    var v9 int32 = -115
    v10 := GenS1{F0: 1000, F1: v7, F2: 448}
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 10; fi12++ {
            v8 = append(v8, fi12)
            var v13 float64 = ((1.5 / (1.5 - -2.25)) + float64((0.1 - (((-2.25 - (1.5 * 0.1)) * ((-2.25 + 1e10) * ((3.14159 - 0.0) + 1e10))) * (3.14159 / 1e10)))))
            v14 := GenS2{F0: 1, F1: 0, F2: 2}
            v15 := GenS3{F0: (fi11 << ((879) & 63))}
            p0 = p0
            fmt.Println(fi11, p0, v7)
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi11
    }
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -83
}

func main() {
    v1 := []string{(((("foo bar" + "hello") + "\n\t") + (("\n\t" + "\n\t") + ("hello" + "foo bar"))) + ("hello" + ("\n\t" + "a"))), "a", ""}
    var v2 bool = (true && false)
    v3 := []int{-467, 32768, 4294967295, 32768}
    var v4 float64 = 1.5
    v5 := make(map[int]bool)
    v6 := []uint32{107}
    v16 := GenS2{F0: 1000, F1: 415, F2: 195}
    var v17 int16 = 1000
    _ = len(v6)
    var v18 uint8 = 216
    for fi19 := 0; fi19 < 5; fi19++ {
        v4 = 0.1
        fmt.Println(v4, v17)
        _ = fi19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        fmt.Println(fi20, v4, v18)
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 int16 = (v17 * int16(fi21))
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    var v23 uint = 0
    if v2 {
        fmt.Println(v4, v4, v4)
        for fi24 := 0; fi24 < 7; fi24++ {
            v25 := GenS4{F0: 691, F1: (32767 != 0)}
            if (genFn0(407) < 100) {
                fmt.Println(v4, v18)
                switch "foo bar" {
                case "hello":
                    var v26 float32 = -2.25
                    _ = v26
                case "foo bar":
                    v18 -= (v18 ^ (v18 % ((v18) | 1)))
                }
                switch "" {
                case "foo bar":
                    v2 = false
                case "a":
                    for fi27 := 0; fi27 < 10; fi27++ {
                        var v28 uint16 = 32767
                        v17 = (v17 << (((v23 / ((v23) | 1))) & 63))
                        switch "\n\t" {
                        case "a":
                            switch "" {
                            case "hello":
                                if true {
                                    var v29 uint8 = 128
                                    var v30 uint = v23
                                    var v31 uint64 = 1
                                    delete(v5, fi27)
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                } else {
                                    switch fi24 {
                                    case -610:
                                        if ((1.5 <= 3.14159) || (("" + (("foo bar" + ("\n\t" + "a")) + "\n\t")) >= "foo bar")) {
                                            v17 = v17
                                            var v32 int8 = int8(v17)
                                            _ = v32
                                        } else {
                                            switch "hello" {
                                            case "hello":
                                                v2 = v2
                                            case "\n\t":
                                                if false {
                                                    switch "a" {
                                                    case "hello":
                                                        v2 = true
                                                    case "":
                                                        if ((v23 | 18446744073709551615) != (v23 >> ((v23) & 63))) {
                                                            switch -9223372036854775808 {
                                                            case 613:
                                                                if ((!(!(1 > 889))) || v2) {
                                                                    var v33 uint16 = (^(v28))
                                                                    if v2 {
                                                                        v33 = (v28 & v33)
                                                                    }
                                                                    _ = v33
                                                                }
                                                                _ = len(v6)
                                                            case 4294967295:
                                                                _ = len(v3)
                                                            case 42:
                                                                var v34 int8 = 100
                                                                _ = v34
                                                            }
                                                            v23 = 1
                                                        } else {
                                                            v35 := GenS5{F0: 42}
                                                            fmt.Println(v23, v28)
                                                            var v36 bool = (genFn0(2) <= 100)
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                    case "\n\t":
                                                        v37 := GenS4{F0: -1, F1: (1 >= 42)}
                                                        _ = v37
                                                    default:
                                                        v38 := GenS6{F0: "", F1: -1, F2: 4294967295}
                                                        _ = v38
                                                    }
                                                } else {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v39 uint32 = 73
                                                        var v40 uint64 = 0
                                                        if len(v1) > 0 {
                                                            v1[0] = ""
                                                        }
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                    v41 := GenS6{F0: ((("\n\t" + "\n\t") + "a") + "foo bar"), F1: genFn0(621), F2: v23}
                                                    _ = v41
                                                }
                                            default:
                                                if (-9223372036854775808 < 9223372036854775807) {
                                                    v2 = ((v18 - v18) == (^(v18)))
                                                } else {
                                                    v2 = v2
                                                }
                                            }
                                        }
                                    case -258:
                                        if (1 >= 0) {
                                            switch (("\n\t" + "a") + "hello") {
                                            case "":
                                                fmt.Println(v2, v2, fi27)
                                            default:
                                                v18 = 42
                                                v2 = (42 != -1)
                                                var v42 bool = (true || false)
                                                _ = v42
                                            }
                                            var v43 bool = ((-1 == 0) && (false && (89 >= 42)))
                                            _ = v43
                                        }
                                    case -621:
                                        var v44 int = (fi24 & (fi27 * 793))
                                        _ = v44
                                    }
                                }
                            case "":
                                v45 := GenS4{F0: -1, F1: (false || false)}
                                _ = v45
                            case "\n\t":
                                v18 -= (v18 << (((v23 >> ((v23) & 63))) & 63))
                            default:
                                v23 -= 86
                            }
                        case "foo bar":
                            if (!v2) {
                                v46 := v5[fi24]
                                v17 = -493
                                if v2 {
                                    v18 = (v18 | v18)
                                    v17 = v17
                                } else {
                                    v46 = v46
                                }
                                if (1e10 <= 0.1) {
                                    v47 := GenS7{F0: "hello", F1: 100}
                                    _ = v47
                                } else {
                                    fmt.Println(v17, v2)
                                    v23 = (v23 % ((v23) | 1))
                                    var v48 uint = (v23 / ((18446744073709551615) | 1))
                                    _ = v48
                                }
                                _ = v46
                            } else {
                                v4 += float64(1.5)
                            }
                        }
                        _ = fi27
                        _ = v28
                    }
                default:
                    v18 = v18
                }
            }
            _ = fi24
            _ = v25
        }
    } else {
        var v49 float64 = v4
        _ = len(v3)
        _ = v49
    }
    var v50 float32 = 0.0
    var v51 uint16 = 22
    for fi52 := 0; fi52 < 9; fi52++ {
        v53 := GenS2{F0: v17, F1: 42, F2: v51}
        if (!v2) {
            v18 += v18
            v4 = (v4 / v4)
        } else {
            var v54 uint32 = 2147483648
            var v55 uint64 = 18446744073709551615
            var v56 uint16 = 297
            _ = v54
            _ = v55
            _ = v56
        }
        _ = fi52
        _ = v53
    }
    v57 := GenS8{F0: 42, F1: -1}
    fmt.Println(v4, v51, v51)
    if ((^(v51)) <= (^(v51))) {
        if v2 {
            var v58 float32 = v50
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v59 bool = (v2 || (!true))
            _ = v58
            _ = v59
        }
    }
    var v60 string = "foo bar"
    _ = len(v6)
    var v61 bool = (v2 || v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v57
    _ = v6
    _ = v60
    _ = v61
}
