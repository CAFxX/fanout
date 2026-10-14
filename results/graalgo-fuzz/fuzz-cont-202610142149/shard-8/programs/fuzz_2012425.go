package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 int; F1 uint8; F2 int16 }
type GenS7 struct { F0 uint16; F1 uint; F2 string }
type GenS8 struct { F0 string; F1 int64; F2 int16 }
func genFn0() int32 {
    v7 := GenS1{F0: 702, F1: -2147483648, F2: -634}
    var v8 int = 32767
    v9 := []int{(v8 & (v8 % ((v8) | 1)))}
    _ = len(v9)
    _ = v7
    _ = v8
    _ = v9
    return -2147483648
}
func genFn1(p0 uint) uint64 {
    var v10 uint64 = 2
    v10 += (v10 / (((v10 & v10)) | 1))
    for fi11 := 0; fi11 < 17; fi11++ {
        p0 = (p0 + p0)
        _ = fi11
    }
    _ = p0
    _ = v10
    return (v10 + v10)
}

func main() {
    var v1 uint64 = 0
    v2 := []string{(((("a" + "foo bar") + "a") + "foo bar") + "hello"), "hello", "a", "\n\t"}
    v3 := make(map[int]uint32)
    var v4 uint16 = 1
    v5 := []string{}
    var v6 float32 = 3.14159
    if (true || false) {
        var v12 int16 = 59
        _ = v12
    }
    var v13 float64 = (0.1 / 1.5)
    var v14 uint16 = 32767
    v13 = -2.25
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 float64 = v13
        _ = len(v3)
        v17 := GenS1{F0: 47, F1: genFn0(), F2: fi15}
        _ = fi15
        _ = v16
        _ = v17
    }
    if false {
        var v18 int64 = -736
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if true {
            if false {
                for fi19 := 0; fi19 < 13; fi19++ {
                    if (v18 < -950) {
                        v20 := GenS2{F0: 42, F1: (v18 + v18)}
                        if (!true) {
                            switch ("\n\t" + "hello") {
                            case "":
                                if ("foo bar" < "foo bar") {
                                    switch (-(fi19)) {
                                    case -497:
                                        v14 -= (v4 + v4)
                                    case 0:
                                        v14 = v4
                                    default:
                                        v21 := GenS3{F0: 72}
                                        _ = v21
                                    }
                                } else {
                                    v22 := GenS4{F0: (v1 ^ 655), F1: ("foo bar" + ("hello" + ""))}
                                    _ = v22
                                }
                            case "hello":
                                fmt.Println(v18, v14, v1)
                            default:
                                if (-2147483648 != -703) {
                                    v23 := GenS3{F0: 127}
                                    v3[fi19] = 2147483647
                                    v24 := GenS3{F0: 250}
                                    _ = v23
                                    _ = v24
                                } else {
                                    var v25 float64 = ((v13 / ((1e10 * ((v13 * v13) + v13)) * 1.5)) + (v13 / v13))
                                    fmt.Println(fi19, v18, v25)
                                    _ = len(v2)
                                    _ = v25
                                }
                            }
                            v26 := GenS5{F0: true, F1: (-(fi19))}
                            _ = v26
                        }
                        _ = v20
                    }
                    _ = fi19
                }
            }
            v27 := GenS4{F0: v1, F1: (("\n\t" + "\n\t") + "")}
            _ = v27
        } else {
            fmt.Println(v1, v18)
            switch (("foo bar" + "hello") + ("a" + "\n\t")) {
            case "\n\t":
                var v28 uint8 = 0
                _ = v28
            case "a":
                fmt.Println(v4, v14, v6)
            }
        }
        _ = v18
    } else {
        v29 := GenS1{F0: v4, F1: -1, F2: -1}
        switch "hello" {
        case "hello":
            if (false || true) {
                switch 1 {
                case -2147483648:
                    fmt.Println(v14, v6)
                case 112:
                    v1 = genFn1(1)
                case -1:
                    v30 := GenS2{F0: 2, F1: 9223372036854775807}
                    _ = v30
                }
                v4 -= (v14 / (((^(v4))) | 1))
                for fi31 := 0; fi31 < 14; fi31++ {
                    switch "foo bar" {
                    case "\n\t":
                        for fi32 := 0; fi32 < 18; fi32++ {
                            var v33 uint8 = 0
                            switch "foo bar" {
                            case "\n\t":
                                v1 = uint64(v33)
                            case "foo bar":
                                v34 := GenS6{F0: fi31, F1: (^(v33)), F2: -1}
                                _ = v34
                            case "a":
                                v14 -= 255
                            }
                            _ = fi32
                            _ = v33
                        }
                    case "foo bar":
                        var v35 int32 = 42
                        _ = v35
                    case "a":
                        var v36 uint16 = (v4 | (v4 & (^(v14))))
                        _ = v36
                    default:
                        v14 -= 42
                    }
                    v13 += (float64(v6) * v13)
                    _ = fi31
                }
            }
        default:
            v13 = v13
        }
        _ = v29
    }
    for fi37 := 0; fi37 < 13; fi37++ {
        v1 = uint64(v4)
        _ = fi37
    }
    var v38 int8 = -128
    if false {
        for fi39 := 0; fi39 < 20; fi39++ {
            v13 -= float64(v6)
            switch "hello" {
            case "a":
                var v40 int32 = -45
                _ = v40
            default:
                var v41 float64 = 3.14159
                _ = v41
            }
            var v42 float32 = (3.14159 * (v6 * (1e10 * v6)))
            _ = fi39
            _ = v42
        }
        for fi43 := 0; fi43 < 7; fi43++ {
            fmt.Println(fi43, v1, v38)
            v44 := GenS4{F0: (v1 / ((v1) | 1)), F1: (("hello" + "foo bar") + "hello")}
            if (true || true) {
                delete(v3, -1)
            } else {
                v6 -= float32(((float64((3.14159 + v6)) / v13) * 0.0))
                var v45 float32 = ((v6 + v6) * 0.0)
                _ = v45
            }
            _ = fi43
            _ = v44
        }
        var v46 float64 = v13
        _ = v46
    }
    if false {
        v47 := GenS7{F0: (v4 / (((v4 + (v4 * (v14 % ((0) | 1))))) | 1)), F1: 2, F2: "hello"}
        v14 = v4
        var v48 uint = 43
        _ = v47
        _ = v48
    }
    fmt.Println(v1, v6)
    v4 += (^(v14))
    fmt.Println(v1, v14)
    fmt.Println(v4, v14, v13)
    if (v14 >= (v4 - (v14 ^ 0))) {
        fmt.Println(v6, v4, v4)
        if false {
            switch -9223372036854775808 {
            case -1:
                for fi49 := 0; fi49 < 15; fi49++ {
                    if (float64(v6) >= v13) {
                        var v50 bool = true
                        _ = v50
                    }
                    _ = fi49
                }
            }
        } else {
            fmt.Println(v14)
        }
        if (v1 >= v1) {
            _ = len(v2)
            if true {
                for fi51 := 0; fi51 < 19; fi51++ {
                    v52 := GenS7{F0: (^(v14)), F1: 1, F2: ("\n\t" + "")}
                    fmt.Println(v38, v4)
                    _ = fi51
                    _ = v52
                }
                for fi53 := 0; fi53 < 2; fi53++ {
                    fmt.Println(v6, v13, v1)
                    switch fi53 {
                    case 0:
                        v54 := GenS5{F0: false, F1: -9223372036854775808}
                        _ = v54
                    }
                    v13 = float64((v6 - 3.14159))
                    _ = fi53
                }
            } else {
                v4 -= (v4 ^ v14)
            }
        } else {
            v3[32767] = 65535
            v6 -= v6
        }
    } else {
        if (v38 > (v38 >> ((42) & 63))) {
            fmt.Println(v4)
        } else {
            fmt.Println(v1, v1)
        }
    }
    if (4294967295 == 65535) {
        for fi55 := 0; fi55 < 17; fi55++ {
            if false {
                v56 := GenS8{F0: "", F1: -1, F2: -852}
                var v57 int64 = 148
                var v58 uint64 = (v1 | v1)
                _ = v56
                _ = v57
                _ = v58
            }
            fmt.Println(v13, v4, fi55)
            var v59 int16 = -963
            _ = fi55
            _ = v59
        }
    } else {
        for fi60 := 0; fi60 < 5; fi60++ {
            v1 = (v1 >> ((114) & 63))
            _ = fi60
        }
    }
    v4 += (v14 + v14)
    if (("hello" + "\n\t") == "") {
        for fi61 := 0; fi61 < 12; fi61++ {
            var v62 int16 = 0
            _ = fi61
            _ = v62
        }
        switch "hello" {
        case "a":
            var v63 uint8 = 2
            _ = v63
        case "foo bar":
            switch 2147483647 {
            case -1:
                if false {
                    v4 = 2
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if true {
                        var v64 uint64 = v1
                        _ = v64
                    }
                }
            case -833:
                v14 += v4
            default:
                v1 += genFn1(42)
            }
        default:
            var v65 string = ""
            _ = v65
        }
        fmt.Println(v1, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
