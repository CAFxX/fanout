package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int32; F2 int64 }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS8 struct { F0 string }
func genFn0(p0 uint64) int32 {
    fmt.Println(p0)
    p0 = (p0 % ((p0) | 1))
    v6 := GenS1{F0: 1, F1: 335}
    v7 := GenS1{F0: 0, F1: 4294967295}
    _ = p0
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1() int16 {
    var v8 uint = 175
    v9 := GenS1{F0: 947, F1: 32767}
    _ = v8
    _ = v9
    return -32768
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 0.1
    var v3 int32 = -49
    var v4 uint8 = 0
    var v5 uint64 = 0
    for fi10 := 0; fi10 < 9; fi10++ {
        v5 -= v5
        if ((!(2147483647 == 78)) || false) {
            if (float64(v2) >= ((0.1 * (0.0 * 3.14159)) / (1.5 - 0.0))) {
                var v11 int32 = v3
                v11 -= 303
                if (238 < v5) {
                    if (!true) {
                        var v12 int = (fi10 / ((fi10) | 1))
                        fmt.Println(v5, v11, v3)
                        v11 = v11
                        _ = v12
                    }
                }
                _ = v11
            }
            fmt.Println(v3, fi10, fi10)
            v13 := GenS2{F0: 80, F1: ("foo bar" + "\n\t")}
            _ = v13
        } else {
            delete(v1, (fi10 & fi10))
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 float64 = -2.25
                _ = fi14
                _ = v15
            }
        }
        _ = fi10
    }
    if true {
        var v16 int = 1
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int64 = 9223372036854775807
            v18 = 9223372036854775807
            switch ("a" + "foo bar") {
            case "":
                var v19 int = (^(v16))
                _ = v19
            case "foo bar":
                switch "\n\t" {
                case "hello":
                    for fi20 := 0; fi20 < 10; fi20++ {
                        if (true || false) {
                            fmt.Println(v2, v18, v3)
                            v21 := GenS1{F0: -718, F1: (v16 << ((1) & 63))}
                            _ = v21
                        } else {
                            v4 -= v4
                            v22 := GenS2{F0: genFn1(), F1: "hello"}
                            var v23 int8 = -1
                            _ = v22
                            _ = v23
                        }
                        _ = fi20
                    }
                case "a":
                    if true {
                        v24 := []int{fi17, (fi17 | (v16 + v16))}
                        _ = v24
                    }
                default:
                    v2 = v2
                }
            default:
                switch "hello" {
                case "":
                    fmt.Println(v5, v5, v2)
                }
            }
            _ = fi17
            _ = v18
        }
        if ((^(v5)) != 311) {
            var v25 int16 = 32767
            v26 := v1[32768]
            var v27 int64 = 449
            _ = v25
            _ = v26
            _ = v27
        }
        _ = v16
    } else {
        v28 := []int{0, 1, 928}
        fmt.Println(v3, v4, v3)
        _ = v28
    }
    v4 -= v4
    if true {
        switch ("hello" + ("\n\t" + "foo bar")) {
        case "\n\t":
            fmt.Println(v3)
        case "hello":
            if (1.5 < 3.14159) {
                for fi29 := 0; fi29 < 5; fi29++ {
                    switch fi29 {
                    case 2147483647:
                        var v30 uint32 = 2147483647
                        _ = v30
                    case 32767:
                        var v31 int8 = 127
                        _ = v31
                    case 127:
                        _ = len(v1)
                    }
                    _ = fi29
                }
                v32 := GenS2{F0: genFn1(), F1: ("hello" + ("a" + ("a" + (("" + "foo bar") + "a"))))}
                if (!(int8(v5) <= 42)) {
                    var v33 bool = true
                    delete(v1, 127)
                    for fi34 := 0; fi34 < 13; fi34++ {
                        v2 += float32(((float64((v2 / float32(0.1))) * (-2.25 + ((float64(0.1) - (((0.1 - 0.1) + (3.14159 * float64(v2))) + ((0.0 / 0.1) * ((1.5 - -2.25) - ((3.14159 + ((((((0.0 / -2.25) + 3.14159) / 1.0) + float64((v2 - v2))) / 1e10) + 0.1)) * -2.25))))) - 1.5))) * 3.14159))
                        var v35 int64 = 42
                        _ = fi34
                        _ = v35
                    }
                    _ = v33
                }
                _ = v32
            } else {
                v36 := GenS3{F0: 0}
                _ = v36
            }
        case "a":
            v2 += v2
        }
        v37 := GenS1{F0: 0, F1: -295}
        if true {
            v2 = v2
            v38 := []int{}
            _ = v38
        }
        _ = v37
    }
    v39 := v1[32768]
    v2 += v2
    fmt.Println(v2)
    v3 = (v3 * genFn0(v5))
    var v40 int32 = v3
    fmt.Println(v40)
    v41 := GenS4{F0: 1, F1: (v3 << ((255) & 63)), F2: -233}
    switch ((v39 + "\n\t") + "a") {
    case "\n\t":
        switch v39 {
        case "foo bar":
            var v42 int32 = (v40 & -1)
            _ = v42
        case "\n\t":
            var v43 uint16 = uint16(v4)
            _ = v43
        case "hello":
            v44 := GenS4{F0: 127, F1: (v40 << ((42) & 63)), F2: -1}
            _ = v44
        }
    default:
        v3 += 2147483647
    }
    for fi45 := 0; fi45 < 1; fi45++ {
        switch v39 {
        case "hello":
            v46 := GenS4{F0: -1, F1: (v3 + (v40 + (-(v40)))), F2: -1}
            _ = v46
        }
        v4 = 255
        _ = fi45
    }
    switch ((v39 + v39) + v39) {
    case "a":
        switch (v39 + (v39 + v39)) {
        case "hello":
            switch ((v39 + "hello") + "hello") {
            case "a":
                v3 -= v3
            case "\n\t":
                for fi47 := 0; fi47 < 20; fi47++ {
                    fmt.Println(v2, v2, fi47)
                    if (-1 <= 42) {
                        v48 := GenS5{F0: -459, F1: v39}
                        if true {
                            v4 -= (v4 + (v4 % (((v4 - (v4 & 1))) | 1)))
                            var v49 int64 = 0
                            _ = v49
                        }
                        _ = v48
                    } else {
                        v4 -= v4
                        if true {
                            fmt.Println(v39, v3, v39)
                            switch "" {
                            case "\n\t":
                                for fi50 := 0; fi50 < 19; fi50++ {
                                    var v51 string = (("hello" + "\n\t") + v39)
                                    _ = fi50
                                    _ = v51
                                }
                            default:
                                v52 := v1[(fi47 * fi47)]
                                _ = v52
                            }
                            for fi53 := 0; fi53 < 13; fi53++ {
                                if true {
                                    v2 += (-2.25 / v2)
                                }
                                _ = len(v1)
                                _ = fi53
                            }
                        } else {
                            for fi54 := 0; fi54 < 9; fi54++ {
                                v4 = v4
                                var v55 uint16 = 646
                                _ = fi54
                                _ = v55
                            }
                            var v56 uint64 = 2
                            v56 = v56
                            _ = v56
                        }
                        v57 := []int{fi47, fi47}
                        _ = v57
                    }
                    _ = fi47
                }
            case "foo bar":
                fmt.Println(v40, v39)
            default:
                v2 += v2
            }
        case "":
            v39 = ""
        default:
            if (!(42 != 42)) {
                v58 := []int{128, int(v40)}
                switch 0 {
                case 492:
                    _ = len(v1)
                case -2147483648:
                    v59 := GenS6{F0: 18446744073709551615}
                    _ = v59
                case 32767:
                    _ = len(v58)
                default:
                    v60 := v1[42]
                    _ = v60
                }
                switch v39 {
                case "hello":
                    for fi61 := 0; fi61 < 5; fi61++ {
                        if true {
                            if (!false) {
                                for fi62 := 0; fi62 < 3; fi62++ {
                                    _ = len(v58)
                                    _ = fi62
                                }
                                _ = len(v1)
                            }
                            _ = len(v58)
                            var v63 uint16 = 32767
                            _ = v63
                        } else {
                            var v64 float64 = 3.14159
                            if len(v58) >= 2 {
                                v58 = v58[1:len(v58)]
                            }
                            if (-1 < 1) {
                                v2 += (float32(v64) * (3.14159 * (((v2 - (1e10 - float32(v64))) * float32(((v64 - (((v64 + ((float64(v2) - v64) / 1.5)) + v64) * float64(-2.25))) * v64))) + (3.14159 / v2))))
                                for fi65 := 0; fi65 < 8; fi65++ {
                                    var v66 float32 = v2
                                    fmt.Println(v5, v4)
                                    v67 := GenS6{F0: v5}
                                    _ = fi65
                                    _ = v66
                                    _ = v67
                                }
                            }
                            _ = v64
                        }
                        v39 = v39
                        _ = fi61
                    }
                case "foo bar":
                    v39 = v39
                case "\n\t":
                    fmt.Println(v5)
                default:
                    if true {
                        v68 := GenS7{F0: 65535, F1: (^(v40)), F2: 543}
                        if (4294967295 == uint(v5)) {
                            v69 := GenS8{F0: (v39 + "\n\t")}
                            for fi70 := 0; fi70 < 12; fi70++ {
                                v40 += genFn0(v5)
                                var v71 uint32 = 1
                                _ = fi70
                                _ = v71
                            }
                            _ = v69
                        }
                        _ = v68
                    } else {
                        v72 := GenS8{F0: v39}
                        _ = v72
                    }
                }
                _ = v58
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
