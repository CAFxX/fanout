package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 int8; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 uint32; F1 uint16; F2 int }
type GenS8 struct { F0 uint16; F1 uint8 }
type GenS9 struct { F0 uint16; F1 int8; F2 string }
func genFn0(p0 bool) int64 {
    if (!(0 > -1)) {
        var v4 uint8 = 136
        v5 := GenS1{F0: 346}
        _ = v4
        _ = v5
    }
    if (float64((3.14159 - (0.0 / 1.0))) > ((((0.1 + 3.14159) - 0.0) + 1e10) + (1e10 / 1.5))) {
        var v6 int16 = 1
        p0 = true
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 uint8 = 85
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    var v9 uint8 = 1
    _ = p0
    _ = v9
    return 9223372036854775807
}
func genFn1(p0 int32, p1 int) int {
    var v10 float32 = 3.14159
    fmt.Println(p0, p1, p0)
    for fi11 := 0; fi11 < 15; fi11++ {
        switch "\n\t" {
        case "foo bar":
            p0 -= p0
        case "":
            v12 := GenS1{F0: 255}
            _ = v12
        default:
            p0 = 277
        }
        v13 := GenS2{F0: (0.0 == (v10 / v10))}
        _ = fi11
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    return p1
}

func main() {
    var v1 int32 = 450
    var v2 int16 = -32768
    var v3 bool = true
    v14 := GenS1{F0: 18446744073709551615}
    v15 := []int{128, 32768, 127, 32768}
    v3 = v3
    var v16 bool = (965 <= 18446744073709551615)
    var v17 string = "a"
    v18 := GenS2{F0: true}
    var v19 string = ""
    v16 = (((1e10 / 1.0) + 1e10) >= (3.14159 - ((float64(float32(1e10)) * 1.5) - (-2.25 + 1.5))))
    for fi20 := 0; fi20 < 1; fi20++ {
        switch fi20 {
        case -1:
            if (!v3) {
                fmt.Println(v17, v1)
                for fi21 := 0; fi21 < 7; fi21++ {
                    if ((!true) && (float32(-2.25) != -2.25)) {
                        switch "hello" {
                        case "a":
                            v1 = v1
                        case "foo bar":
                            v22 := GenS3{F0: 0, F1: 100, F2: 42}
                            _ = v22
                        }
                        if true {
                            v2 += (^(v2))
                        } else {
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            fmt.Println(fi21, v16)
                            v23 := GenS1{F0: 255}
                            _ = v23
                        }
                    } else {
                        var v24 uint32 = 385
                        v1 -= v1
                        switch (fi20 + (fi21 | (^(fi21)))) {
                        case 698:
                            v19 = (v17 + v17)
                        }
                        v17 = "a"
                        var v25 uint64 = 2
                        var v26 int64 = int64(v25)
                        v24 = 2147483648
                        if (100 < 127) {
                            v2 += v2
                            fmt.Println(v17)
                            switch v17 {
                            case "\n\t":
                                if v3 {
                                    if (false && (fi21 == -904)) {
                                        fmt.Println(v19)
                                        v27 := GenS4{F0: (v2 | 107)}
                                        var v28 float64 = float64(((0.1 + 1e10) / -2.25))
                                        _ = len(v15)
                                        fmt.Println(fi20, v19, v16)
                                        _ = v27
                                        _ = v28
                                    } else {
                                        switch v19 {
                                        case "hello":
                                            fmt.Println(v25, v2, v17)
                                            _ = len(v15)
                                            switch v19 {
                                            case "\n\t":
                                                v29 := GenS4{F0: (v2 * v2)}
                                                _ = v29
                                            case "a":
                                                var v30 uint16 = 65535
                                                _ = v30
                                            case "":
                                                var v31 float64 = float64(1e10)
                                                switch fi20 {
                                                case 789:
                                                    var v32 int = (fi21 - -95)
                                                    _ = v32
                                                default:
                                                    fmt.Println(v31)
                                                }
                                                v24 += v24
                                                _ = v31
                                            default:
                                                if (genFn1((v1 ^ (^(v1))), fi20) <= fi21) {
                                                    v15 = append(v15, genFn1(v1, (fi20 - (fi20 ^ fi20))))
                                                    switch 42 {
                                                    case 2147483647:
                                                        var v33 int8 = -1
                                                        _ = v33
                                                    default:
                                                        var v34 int = (fi21 + -2147483648)
                                                        _ = v34
                                                    }
                                                    switch (v19 + v17) {
                                                    case "\n\t":
                                                        if false {
                                                            switch "hello" {
                                                            case "hello":
                                                                v19 = (v19 + v19)
                                                            case "a":
                                                                var v35 float32 = 3.14159
                                                                _ = v35
                                                            case "":
                                                                var v36 float32 = 1e10
                                                                _ = v36
                                                            default:
                                                                switch (fi20 >> ((4294967295) & 63)) {
                                                                case 1:
                                                                    fmt.Println(v17)
                                                                case -2147483648:
                                                                    v25 = v25
                                                                case 4294967295:
                                                                    if (true || true) {
                                                                        v26 = v26
                                                                        var v37 uint64 = (^(v25))
                                                                        _ = v37
                                                                    }
                                                                default:
                                                                    v38 := GenS2{F0: false}
                                                                    _ = v38
                                                                }
                                                            }
                                                        } else {
                                                            if len(v15) > 0 {
                                                                v15[0] = (fi21 | (fi20 + genFn1(v1, 42)))
                                                            }
                                                            v16 = (v3 && (v17 >= v19))
                                                        }
                                                    case "":
                                                        v39 := GenS1{F0: 4294967295}
                                                        _ = v39
                                                    default:
                                                        if ("a" <= v17) {
                                                            _ = len(v15)
                                                        }
                                                    }
                                                }
                                            }
                                        default:
                                            v3 = (-2.25 <= 3.14159)
                                        }
                                        if (!(!(2 == 128))) {
                                            if (genFn0((v1 != (v1 * (v1 % ((v1) | 1))))) > genFn0(((2147483647 != v1) || v16))) {
                                                fmt.Println(v16, fi20)
                                            }
                                            v40 := GenS2{F0: true}
                                            var v41 int64 = genFn0(true)
                                            _ = v40
                                            _ = v41
                                        } else {
                                            v15 = append(v15, 32767)
                                            v3 = false
                                            switch (fi21 * 32768) {
                                            case 962:
                                                if (!((v24 <= v24) && (1e10 < 0.0))) {
                                                    fmt.Println(v19, v19, fi20)
                                                }
                                                fmt.Println(v25)
                                            case 127:
                                                _ = len(v15)
                                            case 9223372036854775807:
                                                if ((v24 >> ((747) & 63)) == 505) {
                                                    v42 := GenS5{F0: 65535, F1: 32767}
                                                    _ = v42
                                                } else {
                                                    switch fi21 {
                                                    case 2147483647:
                                                        var v43 string = (v17 + (v19 + (v17 + ((v17 + ("hello" + v17)) + (("foo bar" + (v17 + "a")) + ((v17 + ("hello" + v19)) + "\n\t"))))))
                                                        _ = v43
                                                    case 9223372036854775807:
                                                        v44 := GenS1{F0: 4294967295}
                                                        _ = v44
                                                    case -490:
                                                        var v45 float64 = 1e10
                                                        _ = v45
                                                    }
                                                    fmt.Println(v2)
                                                    if v3 {
                                                        if v16 {
                                                            if len(v15) > 0 {
                                                                v15[0] = (fi21 - fi20)
                                                            }
                                                            v46 := GenS6{F0: v26, F1: v25}
                                                            _ = v46
                                                        } else {
                                                            v19 = v19
                                                            var v47 int8 = -55
                                                            v24 -= (v24 * (^(v24)))
                                                            _ = v47
                                                        }
                                                        v48 := GenS4{F0: v2}
                                                        _ = v48
                                                    }
                                                    v2 = (v2 + -1)
                                                    switch (v19 + "\n\t") {
                                                    case "":
                                                        v19 = "a"
                                                    case "hello":
                                                        if (!(42 >= 124)) {
                                                            v1 = (^(v1))
                                                            v49 := GenS7{F0: v24, F1: 1, F2: 158}
                                                            v17 = v19
                                                            _ = v49
                                                        }
                                                    default:
                                                        v50 := GenS4{F0: 1}
                                                        _ = v50
                                                    }
                                                }
                                            default:
                                                switch (((v17 + ("\n\t" + v19)) + v17) + "\n\t") {
                                                case "a":
                                                    v3 = ((v1 & (v1 % ((v1) | 1))) == (v1 - (v1 / ((v1) | 1))))
                                                default:
                                                    v17 = v17
                                                }
                                            }
                                        }
                                        var v51 float32 = 0.0
                                        _ = v51
                                    }
                                    v52 := GenS8{F0: 1, F1: 128}
                                    _ = len(v15)
                                    _ = v52
                                }
                            case "":
                                v19 = v17
                            default:
                                v15 = append(v15, (fi20 / ((int(v26)) | 1)))
                            }
                        } else {
                            switch (v17 + (v19 + v17)) {
                            case "a":
                                var v53 uint = 0
                                _ = v53
                            case "\n\t":
                                var v54 uint64 = v25
                                _ = v54
                            case "":
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                            default:
                                v19 = "\n\t"
                                if len(v15) > 0 {
                                    v15[0] = 727
                                }
                            }
                            if (v16 || true) {
                                switch v17 {
                                case "foo bar":
                                    v25 -= (v25 & (^(v25)))
                                }
                                if (v16 || v16) {
                                    if (!((genFn1(0, int(v25)) != fi20) || ((!((v24 >> ((2) & 63)) >= (v24 + 42))) || (42 < v1)))) {
                                        var v55 uint32 = 65535
                                        var v56 uint8 = 127
                                        if len(v15) >= 2 {
                                            v15 = v15[1:len(v15)]
                                        }
                                        _ = v55
                                        _ = v56
                                    } else {
                                        v15 = append(v15, -2147483648)
                                    }
                                } else {
                                    v1 += v1
                                }
                            } else {
                                v26 = v26
                                v57 := GenS9{F0: 32767, F1: -1, F2: v19}
                                _ = v57
                            }
                        }
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    if (4294967295 >= 620) {
                        switch fi20 {
                        case 148:
                            fmt.Println(v19, v17)
                        case -2147483648:
                            _ = len(v15)
                        }
                        var v58 string = v17
                        v1 = 0
                        var v59 int16 = 0
                        _ = v58
                        _ = v59
                    }
                    _ = fi21
                }
            } else {
                for fi60 := 0; fi60 < 9; fi60++ {
                    v61 := GenS4{F0: 0}
                    v1 -= v1
                    if true {
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                    } else {
                        fmt.Println(v3, v1)
                    }
                    _ = fi60
                    _ = v61
                }
                if (v17 != v17) {
                    var v62 int64 = -1
                    switch 42 {
                    case 0:
                        var v63 int16 = (v2 % (((v2 / ((v2) | 1))) | 1))
                        _ = v63
                    case 9223372036854775807:
                        for fi64 := 0; fi64 < 6; fi64++ {
                            fmt.Println(v17, fi20)
                            _ = fi64
                        }
                    case 42:
                        var v65 int8 = 1
                        _ = v65
                    default:
                        switch (("" + (v19 + "")) + "hello") {
                        case "hello":
                            v66 := GenS6{F0: genFn0((!((v16 && true) || ((-(v2)) == v2)))), F1: 141}
                            _ = v66
                        case "foo bar":
                            v16 = (1e10 > 1.5)
                        default:
                            fmt.Println(v3, v62)
                        }
                    }
                    _ = v62
                } else {
                    v67 := GenS4{F0: 1000}
                    _ = v67
                }
            }
        default:
            for fi68 := 0; fi68 < 16; fi68++ {
                var v69 int8 = 1
                v15 = append(v15, genFn1(749, (^(fi68))))
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = fi68
                _ = v69
            }
        }
        switch -2147483648 {
        case 127:
            if (true || ((!v3) || v3)) {
                var v70 float64 = (0.0 + -2.25)
                v17 = (v19 + "")
                _ = v70
            }
        case 42:
            v16 = v3
        default:
            for fi71 := 0; fi71 < 8; fi71++ {
                var v72 float32 = ((0.0 * (1.5 - (-2.25 - ((-2.25 / 1e10) / (1e10 / 1.0))))) - (3.14159 * float32((((-2.25 + (1e10 - float64((-2.25 - 0.1)))) * 0.0) * 1.5))))
                v19 = (v17 + v19)
                v16 = v16
                _ = fi71
                _ = v72
            }
        }
        v3 = v16
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
