package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint8; F1 int64 }

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[uint32]int)
    v3 := []string{"\n\t", (("foo bar" + ("" + "a")) + "foo bar")}
    var v4 float32 = 1.5
    var v5 float32 = ((v4 - ((v4 + v4) - v4)) / 1.0)
    v6 := []int{127, (^(v1)), v1, (v1 | 42)}
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 float32 = 1e10
        _ = fi7
        _ = v8
    }
    var v9 float64 = 0.1
    _ = len(v3)
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            fmt.Println(fi11, fi11)
            var v12 string = ("\n\t" + "")
            v4 -= 0.1
            _ = fi11
            _ = v12
        }
        if true {
            var v13 float64 = 3.14159
            _ = v13
        } else {
            var v14 float32 = v5
            _ = v14
        }
        _ = fi10
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        fmt.Println(fi15)
        var v16 int8 = 79
        switch v1 {
        case 0:
            v5 += ((v4 / ((((0.0 * v5) / ((((((v5 + (v4 - v5)) + -2.25) * (((v5 * -2.25) / float32(v9)) + v4)) - 3.14159) - v5) + 1e10)) + 1e10) * (-2.25 / ((((((1.5 * (0.0 * v4)) + v5) - (v4 - (v5 + 3.14159))) * (((v4 - float32(1.5)) / v4) / (v4 / v5))) * float32((v9 / v9))) * v5)))) * v4)
        }
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        v18 := GenS1{F0: -782, F1: (!false)}
        v19 := GenS1{F0: -124, F1: (374 < 42)}
        v2[4294967295] = (^(fi17))
        _ = fi17
        _ = v18
        _ = v19
    }
    v9 = v9
    for fi20 := 0; fi20 < 1; fi20++ {
        if (!(!true)) {
            if (127 <= 0) {
                v21 := GenS2{F0: 143}
                var v22 string = "\n\t"
                _ = v21
                _ = v22
            } else {
                switch "\n\t" {
                case "hello":
                    var v23 int = (fi20 - 42)
                    _ = v23
                default:
                    switch fi20 {
                    case 32768:
                        v9 += -2.25
                    case -9223372036854775808:
                        if (!(!true)) {
                            switch fi20 {
                            case -9223372036854775808:
                                switch "a" {
                                case "":
                                    if (2147483647 >= fi20) {
                                        if ((-20 < 0) || ((!true) || (-39 < 0))) {
                                            v24 := v2[4294967295]
                                            _ = v24
                                        } else {
                                            var v25 float32 = (-2.25 * (0.1 / v4))
                                            _ = v25
                                        }
                                    }
                                case "hello":
                                    v2[2] = -9223372036854775808
                                case "\n\t":
                                    for fi26 := 0; fi26 < 3; fi26++ {
                                        fmt.Println(v4)
                                        v9 -= (1e10 * 0.0)
                                        v5 += (float32(3.14159) * (v4 / ((v4 - 1.5) + v5)))
                                        v27 := GenS3{F0: 4294967295, F1: 0}
                                        _ = fi26
                                        _ = v27
                                    }
                                }
                            default:
                                fmt.Println(v9)
                            }
                            if len(v6) > 0 {
                                v6[0] = fi20
                            }
                            v1 = 32768
                        }
                    case -114:
                        for fi28 := 0; fi28 < 12; fi28++ {
                            fmt.Println(v9)
                            v5 += v4
                            switch fi20 {
                            case 4294967295:
                                v2[1] = v1
                            case -582:
                                v9 -= (float64(v5) * float64(v4))
                            }
                            if len(v3) > 0 {
                                v3[0] = "\n\t"
                            }
                            _ = fi28
                        }
                    default:
                        if (!false) {
                            v5 += 1e10
                            v5 -= float32(v9)
                        } else {
                            var v29 string = ("a" + "foo bar")
                            _ = v29
                        }
                    }
                }
                if (393 > 557) {
                    var v30 string = (("" + (("a" + ("\n\t" + "\n\t")) + "\n\t")) + "\n\t")
                    for fi31 := 0; fi31 < 18; fi31++ {
                        if (-7 >= 1) {
                            if false {
                                var v32 uint64 = 42
                                switch (fi20 + fi31) {
                                case -167:
                                    v33 := GenS4{F0: 42, F1: 42, F2: 0}
                                    _ = v33
                                case 9223372036854775807:
                                    fmt.Println(v4, fi31, v30)
                                default:
                                    v5 -= 0.1
                                }
                                _ = v32
                            } else {
                                fmt.Println(v5, v4)
                                v34 := GenS5{F0: 0, F1: 65535, F2: 42}
                                _ = v34
                            }
                            v5 += ((3.14159 + (0.0 - ((float32(0.0) - (float32(v9) + (v5 * ((v5 - float32(float64(((v5 - v5) - v4)))) / (v5 + v4))))) * 0.0))) / float32(float64(1e10)))
                            fmt.Println(v4)
                            switch v30 {
                            case "foo bar":
                                var v35 uint = 42
                                _ = v35
                            case "\n\t":
                                v2[2147483647] = (fi31 + 190)
                                var v36 int64 = 42
                                delete(v2, 262)
                                var v37 uint8 = 32
                                _ = v36
                                _ = v37
                            case "hello":
                                var v38 uint16 = 42
                                _ = v38
                            }
                            v4 = (3.14159 * v4)
                            _ = len(v6)
                        }
                        _ = fi31
                    }
                    if false {
                        v5 = v4
                        switch 348 {
                        case -371:
                            for fi39 := 0; fi39 < 19; fi39++ {
                                v4 += 1.5
                                v40 := v2[1]
                                _ = fi39
                                _ = v40
                            }
                        case 127:
                            v5 = v4
                        }
                        v1 = 32768
                    } else {
                        v41 := v2[366]
                        _ = v41
                    }
                    _ = v30
                } else {
                    _ = len(v2)
                    switch "foo bar" {
                    case "foo bar":
                        _ = len(v2)
                    case "":
                        if (true && (true && (128 < 32))) {
                            var v42 int32 = 740
                            var v43 uint32 = 2147483647
                            v9 += 0.0
                            _ = v42
                            _ = v43
                        } else {
                            fmt.Println(v1)
                            v1 = 907
                        }
                    case "a":
                        var v44 int16 = 586
                        _ = v44
                    }
                }
                v45 := GenS2{F0: 127}
                _ = v45
            }
            var v46 float32 = float32((1e10 - -2.25))
            v2[4294967295] = (-(fi20))
            _ = v46
        }
        _ = fi20
    }
    var v47 int32 = 942
    switch "\n\t" {
    case "\n\t":
        v5 -= (0.0 - v5)
    case "hello":
        delete(v2, 2)
    case "":
        fmt.Println(v47)
    }
    var v48 string = "hello"
    fmt.Println(v1, v48)
    v9 = (((v9 + v9) - (0.0 * v9)) / v9)
    if true {
        for fi49 := 0; fi49 < 10; fi49++ {
            var v50 float64 = 1e10
            v6 = append(v6, fi49)
            _ = fi49
            _ = v50
        }
        var v51 int = (v1 - int(v47))
        _ = v51
    } else {
        var v52 uint64 = 0
        if true {
            v9 = 1e10
            v53 := GenS5{F0: 576, F1: 8, F2: (v52 % (((v52 | v52)) | 1))}
            if (false || (!(false && (v52 <= v52)))) {
                fmt.Println(v48)
                var v54 string = "\n\t"
                _ = v54
            }
            _ = v53
        }
        _ = v52
    }
    if (!(((((v9 / v9) + v9) + float64(float32(0.1))) / v9) > v9)) {
        switch "\n\t" {
        case "":
            var v55 float64 = (v9 - float64(v4))
            _ = v55
        case "a":
            v1 = (v1 / (((-(v1))) | 1))
        }
        if true {
            var v56 float64 = ((v9 - ((0.1 * 0.1) + ((float64(v4) + 3.14159) / v9))) - 1.5)
            _ = v56
        } else {
            if len(v3) > 0 {
                v3[0] = v48
            }
            delete(v2, 97)
            var v57 string = v48
            _ = v57
        }
        if (1 == 0) {
            v58 := GenS2{F0: 2}
            var v59 uint64 = 4294967295
            _ = v58
            _ = v59
        }
    } else {
        for fi60 := 0; fi60 < 10; fi60++ {
            switch v1 {
            case -1:
                v4 = (-2.25 / (((0.0 * (((float32(v9) + 0.1) + v4) + (1e10 + v4))) - 1e10) * 1.5))
            case 4294967295:
                switch "a" {
                case "a":
                    v4 -= v5
                default:
                    var v61 int64 = 0
                    _ = v61
                }
            case -113:
                var v62 int64 = 754
                _ = v62
            }
            fmt.Println(fi60)
            _ = fi60
        }
        var v63 float64 = v9
        v64 := GenS6{F0: 42, F1: -201}
        _ = v63
        _ = v64
    }
    if (32767 <= 143) {
        var v65 float64 = v9
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v9
}
