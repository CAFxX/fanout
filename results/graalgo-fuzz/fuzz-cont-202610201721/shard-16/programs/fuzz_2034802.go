package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint64 }
type GenS2 struct { F0 int8; F1 bool; F2 string }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint8, p1 uint64) int16 {
    var v7 int16 = 1
    v8 := []int{-16, 42}
    p0 = (p0 + p0)
    if ((p0 == 0) && false) {
        fmt.Println(p0, v7, p1)
        v9 := GenS1{F0: 56, F1: 2, F2: (p1 * 1)}
        _ = v9
    } else {
        v7 = 84
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (-(v7))
}
func genFn1() int64 {
    var v10 int = 2147483647
    fmt.Println(v10)
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := []int{(fi11 * v10), (fi11 % ((fi11) | 1)), v10, fi11}
        _ = len(v12)
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 uint64 = 429
        _ = fi13
        _ = v14
    }
    _ = v10
    return 895
}

func main() {
    var v1 float64 = (float64(0.0) / 1.0)
    var v2 uint16 = 120
    var v3 uint = 255
    v4 := make(map[string]int)
    var v5 float64 = 3.14159
    v6 := make(map[string]int)
    v15 := GenS2{F0: -8, F1: (!false), F2: "a"}
    v16 := []int{-2147483648, 0, -9223372036854775808}
    var v17 string = "a"
    var v18 uint8 = 1
    delete(v4, v17)
    v17 = v17
    switch ("a" + "foo bar") {
    case "foo bar":
        v19 := GenS1{F0: 1, F1: (v3 << ((v3) & 63)), F2: 27}
        _ = v19
    case "a":
        for fi20 := 0; fi20 < 5; fi20++ {
            v1 = 3.14159
            _ = fi20
        }
    case "":
        v2 -= (^(v2))
    default:
        v5 = v1
    }
    fmt.Println(v5, v5)
    for fi21 := 0; fi21 < 11; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 int32 = -1
            switch v17 {
            case "foo bar":
                v23 -= (v23 % (((v23 + v23)) | 1))
            case "":
                fmt.Println(v2, v17, v1)
            case "hello":
                v24 := GenS2{F0: 15, F1: true, F2: ("" + "foo bar")}
                var v25 string = (((v17 + "hello") + ((v17 + v17) + "hello")) + "foo bar")
                _ = v24
                _ = v25
            default:
                switch (v17 + (("foo bar" + v17) + v17)) {
                case "":
                    v5 = (1.5 - -2.25)
                case "\n\t":
                    var v26 int = fi22
                    _ = v26
                }
            }
            v5 += 3.14159
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        switch (v17 + "") {
        case "hello":
            for fi28 := 0; fi28 < 17; fi28++ {
                v5 += (float64((float32(v5) * 0.1)) - ((0.1 / 0.1) / v5))
                _ = fi28
            }
        case "foo bar":
            fmt.Println(v17, v17, v3)
        case "\n\t":
            var v29 bool = true
            _ = v29
        default:
            v18 -= (^(v18))
        }
        _ = fi27
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        for fi31 := 0; fi31 < 9; fi31++ {
            switch (fi30 & fi31) {
            case 1:
                switch fi30 {
                case 17:
                    if (!(!(!((v18 >> ((0) & 63)) >= 1)))) {
                        var v32 int64 = -632
                        v1 -= float64(3.14159)
                        _ = v32
                    }
                    v17 = v17
                    v33 := GenS2{F0: -1, F1: false, F2: (("hello" + ("\n\t" + ("a" + v17))) + v17)}
                    _ = v33
                }
            case 32767:
                if false {
                    switch "a" {
                    case "":
                        fmt.Println(v5, fi30, v3)
                    default:
                        v34 := GenS3{F0: 655, F1: 1}
                        _ = v34
                    }
                    var v35 uint16 = 103
                    v1 += v5
                    v16 = append(v16, (fi30 + fi30))
                    v5 = (1.5 * v5)
                    _ = v35
                }
            default:
                v36 := GenS3{F0: 9223372036854775807, F1: 1000}
                _ = v36
            }
            var v37 float64 = v5
            _ = fi31
            _ = v37
        }
        _ = fi30
    }
    for fi38 := 0; fi38 < 6; fi38++ {
        v17 = (((v17 + "a") + "a") + v17)
        _ = fi38
    }
    if true {
        v2 -= 51
    } else {
        if (!(1 != 905)) {
            switch (v17 + v17) {
            case "\n\t":
                v18 = v18
            default:
                v5 = float64(0.0)
            }
            switch v17 {
            case "":
                var v39 uint16 = 1
                _ = v39
            case "a":
                if (127 == 1) {
                    v40 := GenS4{F0: -306, F1: 127}
                    _ = v40
                } else {
                    v3 = v3
                    if false {
                        var v41 int16 = 42
                        for fi42 := 0; fi42 < 2; fi42++ {
                            var v43 int32 = -2147483648
                            var v44 uint = v3
                            if (!((false || ((!false) && true)) && (int(v18) == (fi42 >> (((^(v3))) & 63))))) {
                                v45 := GenS4{F0: (v43 % ((v43) | 1)), F1: (v18 % (((v18 - 65)) | 1))}
                                v4[v17] = fi42
                                _ = v45
                            }
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                        for fi46 := 0; fi46 < 1; fi46++ {
                            if true {
                                var v47 uint = 0
                                switch (("\n\t" + ((("" + "hello") + v17) + ("" + (v17 + v17)))) + v17) {
                                case "foo bar":
                                    var v48 string = v17
                                    _ = v48
                                default:
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                }
                                _ = v47
                            }
                            switch 4294967295 {
                            case 32768:
                                if false {
                                    switch 0 {
                                    case 1:
                                        if (186 > v18) {
                                            v17 = "foo bar"
                                            v41 -= (v41 + v41)
                                            v49 := GenS2{F0: 100, F1: (!(!false)), F2: v17}
                                            _ = v49
                                        }
                                    case -330:
                                        if true {
                                            v16 = append(v16, -2147483648)
                                        }
                                    }
                                    if (uint16(v3) < v2) {
                                        for fi50 := 0; fi50 < 19; fi50++ {
                                            var v51 int16 = v41
                                            v52 := GenS2{F0: 56, F1: ((false || false) && false), F2: ""}
                                            _ = fi50
                                            _ = v51
                                            _ = v52
                                        }
                                    }
                                    for fi53 := 0; fi53 < 16; fi53++ {
                                        switch 4294967295 {
                                        case 127:
                                            v2 += (v2 & (v2 + 122))
                                        }
                                        _ = fi53
                                    }
                                } else {
                                    var v54 uint16 = v2
                                    if len(v16) > 0 {
                                        v16[0] = (fi46 ^ (^(fi46)))
                                    }
                                    _ = v54
                                }
                            case -658:
                                v17 = "a"
                            case 9223372036854775807:
                                v55 := GenS2{F0: 0, F1: ((127 == -105) && true), F2: ("\n\t" + ((v17 + "") + "\n\t"))}
                                _ = v55
                            }
                            _ = fi46
                        }
                        _ = v41
                    } else {
                        v2 -= (v2 ^ 1)
                        var v56 int = 2147483647
                        _ = v56
                    }
                    for fi57 := 0; fi57 < 7; fi57++ {
                        v3 = 998
                        var v58 uint64 = 1
                        _ = fi57
                        _ = v58
                    }
                }
            case "hello":
                fmt.Println(v3, v5)
            }
        } else {
            for fi59 := 0; fi59 < 9; fi59++ {
                v18 = 42
                v18 = 109
                var v60 uint = (^(v3))
                _ = fi59
                _ = v60
            }
            v3 += v3
        }
        var v61 int = -96
        _ = v61
    }
    for fi62 := 0; fi62 < 17; fi62++ {
        switch (("hello" + v17) + v17) {
        case "foo bar":
            v18 = 106
        case "\n\t":
            var v63 int32 = -2147483648
            _ = v63
        default:
            var v64 float32 = float32(v1)
            _ = v64
        }
        if (true && false) {
            v65 := GenS5{F0: 4294967295, F1: -2147483648}
            delete(v4, (("" + "foo bar") + v17))
            _ = v65
        }
        _ = fi62
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
