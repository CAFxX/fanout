package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int64; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint64; F1 string; F2 int8 }
type GenS5 struct { F0 int32; F1 int }
type GenS6 struct { F0 uint64; F1 int64 }
func genFn0() int {
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 bool = true
        v8 := []int{fi6}
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(fi9)
        var v10 bool = false
        var v11 uint = 255
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int8 = -128
    var v13 int16 = -32768
    fmt.Println(v13, v13)
    _ = v12
    _ = v13
    return -248
}

func main() {
    var v1 float64 = -2.25
    var v2 int32 = 1
    var v3 uint16 = 1
    v4 := []int{}
    var v5 uint64 = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v14 int16 = 0
    v14 -= -32768
    for fi15 := 0; fi15 < 1; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            v3 += v3
            if false {
                v2 = int32(v3)
                switch 106 {
                case -1:
                    v1 -= 0.0
                case 32768:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case -278:
                    var v17 uint64 = v5
                    _ = v17
                }
            } else {
                v4 = append(v4, (fi15 << ((11) & 63)))
                var v18 int = 2147483647
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v18
            }
            switch (fi15 << ((2) & 63)) {
            case 32767:
                v3 -= (v3 & 2)
            case -356:
                v19 := GenS1{F0: 952}
                _ = v19
            case 42:
                if ((!(128 <= 128)) && ((fi15 * fi16) > (fi16 >> ((1) & 63)))) {
                    switch "hello" {
                    case "foo bar":
                        v2 = v2
                    case "a":
                        var v20 int = (fi15 - (^(fi15)))
                        _ = v20
                    }
                }
            }
            fmt.Println(v1)
            _ = fi16
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            v2 = (v2 / ((42) | 1))
            if true {
                v2 = 42
                fmt.Println(v1)
                fmt.Println(v1, v2)
                fmt.Println(fi21)
            }
            _ = fi21
        }
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := GenS1{F0: v5}
            _ = fi22
            _ = v23
        }
        _ = fi15
    }
    v14 += v14
    for fi24 := 0; fi24 < 12; fi24++ {
        var v25 bool = ((v2 >> ((424) & 63)) != v2)
        if (32767 <= (fi24 & (fi24 << ((2) & 63)))) {
            if v25 {
                fmt.Println(v2, v2, v25)
                if ((v2 << ((255) & 63)) > (v2 ^ (^(v2)))) {
                    v3 -= 42
                } else {
                    for fi26 := 0; fi26 < 6; fi26++ {
                        var v27 int64 = 0
                        var v28 uint = 2
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    v14 += v14
                }
                switch ("" + "a") {
                case "foo bar":
                    v29 := GenS1{F0: (v5 | v5)}
                    _ = v29
                case "a":
                    if v25 {
                        for fi30 := 0; fi30 < 12; fi30++ {
                            if (v25 && v25) {
                                v5 = v5
                                if v25 {
                                    var v31 int8 = int8(v3)
                                    var v32 int64 = int64(v5)
                                    var v33 int8 = (v31 | v31)
                                    if false {
                                        v4 = append(v4, (fi30 << ((385) & 63)))
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = fi30
                                        }
                                    }
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                            } else {
                                if ("foo bar" != "") {
                                    v3 = 42
                                    if len(v4) > 0 {
                                        v4[0] = int(v3)
                                    }
                                    fmt.Println(fi30, v25, fi24)
                                } else {
                                    var v34 uint8 = 1
                                    var v35 uint32 = 1
                                    switch "" {
                                    case "foo bar":
                                        if len(v4) > 0 {
                                            v4[0] = (fi24 % ((127) | 1))
                                        }
                                        var v36 uint16 = (^(v3))
                                        _ = v36
                                    case "hello":
                                        v37 := GenS1{F0: (v5 * v5)}
                                        _ = v37
                                    case "":
                                        v38 := GenS2{F0: 235, F1: v25}
                                        _ = v38
                                    }
                                    _ = v34
                                    _ = v35
                                }
                                _ = len(v4)
                                var v39 string = "hello"
                                _ = v39
                            }
                            if ((fi24 | (^(fi30))) <= fi30) {
                                fmt.Println(fi30, v3)
                            }
                            var v40 int8 = 0
                            v25 = (v25 || ((^(v3)) >= (v3 + v3)))
                            _ = fi30
                            _ = v40
                        }
                        if len(v4) > 0 {
                            v4[0] = (^(fi24))
                        }
                        if v25 {
                            v41 := GenS2{F0: 573, F1: v25}
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v5 -= (^(v5))
                            _ = v41
                        }
                    }
                case "\n\t":
                    _ = len(v4)
                }
            }
            switch fi24 {
            case -9223372036854775808:
                v42 := GenS2{F0: 255, F1: ((("" + "a") + "foo bar") == "hello")}
                _ = v42
            default:
                fmt.Println(v3, v25, v3)
            }
        }
        _ = fi24
        _ = v25
    }
    if ((v14 << ((4294967295) & 63)) <= v14) {
        switch -9223372036854775808 {
        case 281:
            _ = len(v4)
        case -2147483648:
            v14 = -32768
        case -462:
            v5 = v5
        }
        fmt.Println(v2)
    } else {
        v2 = v2
    }
    v43 := GenS2{F0: 229, F1: (!(v14 != v14))}
    v44 := GenS3{F0: -1, F1: v2, F2: (v3 | v3)}
    var v45 uint64 = 155
    if false {
        if (32767 < -833) {
            if true {
                var v46 int32 = (v2 >> ((291) & 63))
                v47 := GenS4{F0: (v5 - (v5 + v5)), F1: (("foo bar" + "\n\t") + ("foo bar" + "hello")), F2: 100}
                for fi48 := 0; fi48 < 5; fi48++ {
                    v2 -= 1
                    if (!(42 != 491)) {
                        v45 -= v45
                        fmt.Println(v14)
                    }
                    _ = fi48
                }
                _ = v46
                _ = v47
            }
            if len(v4) > 0 {
                v4[0] = 32767
            }
        }
    } else {
        var v49 int64 = -856
        fmt.Println(v5, v3)
        _ = v49
    }
    if true {
        v5 -= v45
        v50 := GenS5{F0: (-(v2)), F1: -2147483648}
        if false {
            if false {
                _ = len(v4)
            }
            v51 := GenS4{F0: (v45 ^ 18446744073709551615), F1: "hello", F2: 127}
            if (1 < 42) {
                switch genFn0() {
                case 1:
                    v14 += v14
                case 941:
                    v52 := GenS6{F0: 2, F1: -9223372036854775808}
                    _ = v52
                case 128:
                    fmt.Println(v45, v1)
                }
                fmt.Println(v1)
            } else {
                fmt.Println(v1)
                v1 = (v1 / (((float64(3.14159) / v1) - ((float64(1.5) + 1e10) - v1)) + v1))
            }
            _ = v51
        } else {
            for fi53 := 0; fi53 < 13; fi53++ {
                fmt.Println(v1, fi53)
                for fi54 := 0; fi54 < 18; fi54++ {
                    v55 := GenS3{F0: 9223372036854775807, F1: (v2 >> ((334) & 63)), F2: (v3 << ((872) & 63))}
                    if len(v4) > 0 {
                        v4[0] = (fi53 % ((fi53) | 1))
                    }
                    _ = fi54
                    _ = v55
                }
                _ = fi53
            }
            var v56 int8 = int8(v2)
            for fi57 := 0; fi57 < 7; fi57++ {
                var v58 int64 = 0
                v59 := GenS3{F0: v58, F1: v2, F2: v3}
                _ = fi57
                _ = v58
                _ = v59
            }
            _ = v56
        }
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
