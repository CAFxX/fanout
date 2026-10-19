package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int16; F2 bool }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := []string{"a", ("a" + ("a" + "hello")), "foo bar"}
    var v3 int = 4294967295
    v4 := []int{(v3 | -827), (v3 & -14)}
    var v5 int32 = -453
    v6 := []string{"a"}
    if true {
        v7 := GenS1{F0: 127, F1: v5, F2: 255}
        fmt.Println(v3)
        _ = v7
    }
    if true {
        v5 = -2
        fmt.Println(v5, v3)
        var v8 bool = true
        _ = v8
    }
    v3 -= (^(v3))
    v5 = 0
    fmt.Println(v3, v3)
    if len(v4) > 0 {
        v4[0] = v3
    }
    var v9 int8 = 58
    v9 -= (v9 + v9)
    v10 := v1[int(v9)]
    if false {
        for fi11 := 0; fi11 < 16; fi11++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch "a" {
            case "":
                _ = len(v1)
            case "hello":
                v3 -= (v3 % ((-1) | 1))
            }
            _ = fi11
        }
        switch (v3 * (v10 / ((v10) | 1))) {
        case -1:
            fmt.Println(v3, v3, v5)
        case 32767:
            var v12 int8 = v9
            _ = v12
        }
    }
    var v13 float32 = 0.1
    for fi14 := 0; fi14 < 7; fi14++ {
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        var v15 uint16 = 65535
        _ = fi14
        _ = v15
    }
    v3 = (v3 >> ((255) & 63))
    switch v10 {
    case 4294967295:
        for fi16 := 0; fi16 < 20; fi16++ {
            fmt.Println(v9, fi16, v9)
            switch -371 {
            case 128:
                var v17 uint16 = 42
                _ = v17
            }
            _ = fi16
        }
    case -1:
        if (1 >= 2) {
            if true {
                if (1 < (v5 / (((v5 % (((v5 + (v5 << ((1) & 63)))) | 1))) | 1))) {
                    for fi18 := 0; fi18 < 20; fi18++ {
                        v19 := GenS2{F0: true, F1: 65535, F2: 1}
                        _ = fi18
                        _ = v19
                    }
                    delete(v1, v10)
                    for fi20 := 0; fi20 < 2; fi20++ {
                        if (!false) {
                            switch v10 {
                            case 858:
                                var v21 uint64 = 2
                                _ = v21
                            default:
                                v3 -= 4294967295
                            }
                            v22 := GenS1{F0: 254, F1: (v5 | v5), F2: 0}
                            _ = v22
                        }
                        _ = fi20
                    }
                } else {
                    v10 += -293
                    var v23 uint32 = 65535
                    v23 = v23
                    _ = v23
                }
            }
        } else {
            var v24 int = 9223372036854775807
            v10 += v24
            _ = v24
        }
    }
    v5 -= int32(v3)
    var v25 string = ("foo bar" + "")
    if true {
        var v26 uint = 4294967295
        v27 := v1[-129]
        for fi28 := 0; fi28 < 16; fi28++ {
            switch fi28 {
            case 207:
                for fi29 := 0; fi29 < 12; fi29++ {
                    fmt.Println(fi29)
                    _ = len(v6)
                    _ = fi29
                }
            case 2147483647:
                v27 = (v10 - (v3 & (v3 - (v3 / (((fi28 | (fi28 & v3))) | 1)))))
            case 4294967295:
                v6 = append(v6, v25)
            default:
                switch v25 {
                case "\n\t":
                    fmt.Println(v10)
                case "a":
                    for fi30 := 0; fi30 < 6; fi30++ {
                        switch fi28 {
                        case 1:
                            if true {
                                v10 -= 32767
                                var v31 uint16 = 255
                                _ = v31
                            } else {
                                var v32 string = v25
                                _ = v32
                            }
                        case 2147483647:
                            v33 := GenS3{F0: 1}
                            _ = v33
                        case -1:
                            v9 = v9
                        }
                        var v34 uint8 = 128
                        var v35 int64 = 0
                        var v36 bool = (1 > 1)
                        switch fi28 {
                        case 302:
                            if v36 {
                                if ((int8(v5) != (^(v9))) && (!false)) {
                                    if len(v2) > 0 {
                                        v2[0] = v25
                                    }
                                    fmt.Println(v13, v9, v3)
                                } else {
                                    v34 -= (^(v34))
                                    v6 = append(v6, ((v25 + ("\n\t" + v25)) + ("foo bar" + (v25 + "a"))))
                                    var v37 float32 = ((1.5 / v13) * ((3.14159 - v13) - v13))
                                    var v38 int32 = v5
                                    v37 -= v13
                                    v25 = ("hello" + v25)
                                    _ = v37
                                    _ = v38
                                }
                                var v39 uint16 = 1
                                v6 = append(v6, v25)
                                _ = v39
                            } else {
                                v25 = v25
                                if v36 {
                                    var v40 uint16 = 1
                                    if v36 {
                                        fmt.Println(v34, v27)
                                        if true {
                                            if v36 {
                                                switch v25 {
                                                case "foo bar":
                                                    if (32767 >= 0) {
                                                        v41 := GenS4{F0: v34, F1: -181, F2: true}
                                                        _ = v41
                                                    } else {
                                                        if true {
                                                            if false {
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                v26 -= v26
                                                                switch "" {
                                                                case "hello":
                                                                    v1[(v10 + v27)] = fi28
                                                                }
                                                                switch (((v25 + (v25 + v25)) + ("\n\t" + "")) + v25) {
                                                                case "a":
                                                                    var v42 int32 = (v5 - (v5 << ((337) & 63)))
                                                                    _ = v42
                                                                case "foo bar":
                                                                    v43 := GenS4{F0: (v34 ^ 42), F1: 77, F2: (!(73 == 18446744073709551615))}
                                                                    _ = v43
                                                                }
                                                            } else {
                                                                if true {
                                                                    var v44 uint8 = (v34 - v34)
                                                                    fmt.Println(v44)
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    _ = v44
                                                                } else {
                                                                    fmt.Println(fi28, v40, fi28)
                                                                    v13 += (v13 - v13)
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                }
                                                                v13 = ((v13 * (v13 * float32(1e10))) + v13)
                                                                switch v25 {
                                                                case "\n\t":
                                                                    var v45 int = (v27 % ((54) | 1))
                                                                    _ = v45
                                                                default:
                                                                    var v46 int16 = 42
                                                                    _ = v46
                                                                }
                                                            }
                                                            fmt.Println(fi30, v25)
                                                            var v47 uint32 = 645
                                                            _ = v47
                                                        } else {
                                                            fmt.Println(v10, v27)
                                                            if (v13 < (((1.5 + ((v13 - v13) + 1.5)) * v13) / 1.5)) {
                                                                var v48 float32 = (float32((((1.5 / 0.1) + (0.1 + 0.1)) + -2.25)) * ((v13 - (1e10 / ((1.5 + 3.14159) + ((v13 / v13) - v13)))) / 0.1))
                                                                var v49 bool = false
                                                                switch ((v25 + (v25 + v25)) + v25) {
                                                                case "\n\t":
                                                                    var v50 float64 = float64(-2.25)
                                                                    _ = v50
                                                                case "foo bar":
                                                                    v48 = v13
                                                                }
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                        }
                                                        var v51 uint16 = (v40 % ((v40) | 1))
                                                        var v52 uint16 = (v51 >> ((v26) & 63))
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                case "":
                                                    fmt.Println(v3, v5, v25)
                                                default:
                                                    var v53 uint32 = 2
                                                    _ = v53
                                                }
                                            }
                                            v40 += v40
                                            v25 = ((v25 + v25) + (v25 + ("hello" + ("hello" + ""))))
                                        }
                                        switch -902 {
                                        case -134:
                                            var v54 bool = v36
                                            _ = v54
                                        case 521:
                                            v9 += v9
                                        }
                                    }
                                    v40 = 65535
                                    _ = v40
                                } else {
                                    switch 42 {
                                    case 2147483647:
                                        switch (v25 + "") {
                                        case "":
                                            v3 -= v3
                                        default:
                                            if v36 {
                                                var v55 int16 = -331
                                                switch fi28 {
                                                case -881:
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                case 2147483647:
                                                    if v36 {
                                                        fmt.Println(fi30, fi30)
                                                        var v56 uint64 = 18446744073709551615
                                                        var v57 int32 = (v5 >> ((565) & 63))
                                                        _ = v56
                                                        _ = v57
                                                    } else {
                                                        var v58 uint = (v26 & (v26 | v26))
                                                        _ = v58
                                                    }
                                                }
                                                _ = v55
                                            }
                                        }
                                    default:
                                        v59 := GenS1{F0: v34, F1: v5, F2: 141}
                                        _ = v59
                                    }
                                    fmt.Println(v3, v13, v36)
                                }
                                v3 -= (v3 / ((v10) | 1))
                            }
                        case -1:
                            fmt.Println(v3, v5, v27)
                        }
                        _ = fi30
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                default:
                    v60 := GenS5{F0: 0, F1: v25}
                    _ = v60
                }
            }
            if (128 == 2) {
                delete(v1, v3)
                if (2147483647 > v5) {
                    switch (fi28 | fi28) {
                    case -2147483648:
                        fmt.Println(v9, fi28, v27)
                    case 811:
                        v61 := GenS2{F0: true, F1: v26, F2: 42}
                        _ = v61
                    case -550:
                        var v62 float64 = 0.1
                        _ = v62
                    default:
                        v10 -= v27
                    }
                } else {
                    v10 -= 4294967295
                    var v63 string = (v25 + "\n\t")
                    var v64 int32 = (v5 & v5)
                    _ = v63
                    _ = v64
                }
            } else {
                v3 -= 32767
                if false {
                    var v65 string = v25
                    v66 := GenS1{F0: 1, F1: v5, F2: 1}
                    _ = v65
                    _ = v66
                } else {
                    if false {
                        v9 += v9
                    } else {
                        var v67 int32 = v5
                        _ = v67
                    }
                    fmt.Println(v25)
                    v68 := GenS6{F0: 42}
                    _ = v68
                }
                for fi69 := 0; fi69 < 10; fi69++ {
                    if ((!(false || (171 != 135))) && false) {
                        fmt.Println(v26, v5, v27)
                    }
                    switch (v25 + v25) {
                    case "foo bar":
                        _ = len(v2)
                    }
                    _ = fi69
                }
            }
            _ = fi28
        }
        _ = v26
        _ = v27
    } else {
        v70 := GenS5{F0: -1, F1: v25}
        var v71 uint8 = 0
        var v72 float32 = (v13 / ((v13 - float32((1e10 - -2.25))) / 1e10))
        _ = v70
        _ = v71
        _ = v72
    }
    v73 := GenS3{F0: (v5 & (v5 & 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v9
}
