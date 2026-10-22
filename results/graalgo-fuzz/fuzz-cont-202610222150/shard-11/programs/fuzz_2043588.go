package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int64; F2 int32 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32; F1 int8; F2 uint }

func main() {
    var v1 bool = true
    var v2 string = ((("a" + ("" + "a")) + "") + "a")
    v3 := make(map[uint32]string)
    var v4 uint16 = 255
    var v5 int16 = 32767
    v6 := []string{("a" + v2), "hello", "hello"}
    switch v2 {
    case "a":
        _ = len(v6)
    }
    if false {
        var v7 int8 = -1
        _ = v7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int64 = 9223372036854775807
        if (0 >= 879) {
            for fi10 := 0; fi10 < 5; fi10++ {
                if v1 {
                    fmt.Println(v1, v5)
                    var v11 float32 = float32((0.1 + (0.1 * (1.5 * 0.0))))
                    v12 := GenS1{F0: (v9 + 1), F1: 0, F2: 151}
                    _ = v11
                    _ = v12
                } else {
                    switch fi10 {
                    case 42:
                        switch v2 {
                        case "a":
                            v13 := GenS1{F0: (-(v9)), F1: v4, F2: 2}
                            _ = v13
                        case "":
                            v1 = (v4 < (v4 / ((0) | 1)))
                        }
                    case -1:
                        fmt.Println(v4)
                    default:
                        v14 := GenS2{F0: v5}
                        _ = v14
                    }
                }
                _ = fi10
            }
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(v5)
    var v15 float32 = 0.1
    fmt.Println(v5, v2)
    if (true && true) {
        switch v2 {
        case "a":
            if true {
                if false {
                    var v16 float64 = (1e10 - (1.5 / 1e10))
                    var v17 int = 42
                    v15 += (v15 - (float32(0.0) - float32(v16)))
                    _ = v16
                    _ = v17
                }
            }
        case "":
            v18 := GenS1{F0: 614, F1: 2, F2: 30}
            _ = v18
        case "hello":
            _ = len(v6)
        default:
            for fi19 := 0; fi19 < 4; fi19++ {
                switch v2 {
                case "a":
                    v6 = append(v6, v2)
                case "\n\t":
                    var v20 uint8 = 2
                    _ = v20
                case "":
                    if (false || false) {
                        var v21 int64 = 42
                        _ = v21
                    } else {
                        v22 := GenS2{F0: (v5 % (((v5 - 42)) | 1))}
                        for fi23 := 0; fi23 < 1; fi23++ {
                            v15 = (1.5 - ((0.1 + -2.25) - float32(float64(v15))))
                            _ = fi23
                        }
                        if (!v1) {
                            if false {
                                var v24 uint64 = 42
                                var v25 float32 = ((1e10 / ((0.1 + 1.5) + 0.1)) / 0.1)
                                _ = v24
                                _ = v25
                            }
                            v26 := GenS1{F0: -1, F1: v4, F2: 831}
                            _ = v26
                        } else {
                            v3[2147483648] = ""
                            for fi27 := 0; fi27 < 10; fi27++ {
                                v2 = (v2 + v2)
                                switch fi27 {
                                case 32767:
                                    delete(v3, 0)
                                case 82:
                                    _ = len(v3)
                                default:
                                    v2 = ""
                                }
                                _ = len(v3)
                                var v28 uint16 = v4
                                _ = fi27
                                _ = v28
                            }
                        }
                        _ = v22
                    }
                }
                v4 = (^(v4))
                _ = fi19
            }
        }
        switch 2147483647 {
        case 127:
            for fi29 := 0; fi29 < 1; fi29++ {
                v1 = (4294967295 != 2)
                switch 4294967295 {
                case 127:
                    var v30 string = v2
                    _ = v30
                case 128:
                    var v31 int8 = -118
                    _ = v31
                case -712:
                    fmt.Println(v2, fi29)
                default:
                    fmt.Println(v1)
                }
                v32 := GenS3{F0: ((v2 + (v2 + "")) + "hello"), F1: 0, F2: 567}
                _ = fi29
                _ = v32
            }
        case 4294967295:
            var v33 uint64 = 18446744073709551615
            _ = v33
        case 0:
            switch ((((v2 + (v2 + ("hello" + ""))) + v2) + v2) + ("foo bar" + (v2 + "hello"))) {
            case "\n\t":
                v34 := GenS1{F0: -1, F1: (^(v4)), F2: 42}
                _ = v34
            case "":
                var v35 uint64 = 1
                _ = v35
            default:
                var v36 uint16 = (v4 - (v4 + 42))
                _ = v36
            }
        }
        for fi37 := 0; fi37 < 9; fi37++ {
            v1 = true
            _ = fi37
        }
    } else {
        for fi38 := 0; fi38 < 16; fi38++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            for fi39 := 0; fi39 < 19; fi39++ {
                v1 = true
                v40 := GenS1{F0: -9223372036854775808, F1: 65535, F2: 34}
                var v41 float32 = v15
                v15 = 3.14159
                v42 := GenS1{F0: 1, F1: 465, F2: 1}
                var v43 bool = (true && false)
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
                _ = v43
            }
            _ = fi38
        }
        switch 127 {
        case -952:
            v44 := GenS1{F0: -9223372036854775808, F1: v4, F2: 4294967295}
            _ = v44
        case 185:
            var v45 bool = (0 == -128)
            _ = v45
        case 427:
            v46 := GenS1{F0: -9223372036854775808, F1: 65535, F2: 2}
            _ = v46
        default:
            switch v2 {
            case "a":
                var v47 int = 32767
                _ = v47
            case "foo bar":
                for fi48 := 0; fi48 < 7; fi48++ {
                    for fi49 := 0; fi49 < 10; fi49++ {
                        switch v2 {
                        case "foo bar":
                            v4 = v4
                        case "":
                            var v50 int64 = -9223372036854775808
                            _ = v50
                        }
                        v1 = v1
                        switch fi49 {
                        case -9223372036854775808:
                            if true {
                                v4 += 255
                            }
                        default:
                            v5 -= (v5 & v5)
                        }
                        _ = fi49
                    }
                    _ = fi48
                }
            case "hello":
                var v51 int16 = v5
                _ = v51
            default:
                v52 := GenS2{F0: v5}
                _ = v52
            }
        }
    }
    var v53 int8 = 127
    for fi54 := 0; fi54 < 17; fi54++ {
        delete(v3, 0)
        var v55 int32 = -659
        _ = fi54
        _ = v55
    }
    for fi56 := 0; fi56 < 6; fi56++ {
        v57 := GenS1{F0: -1, F1: v4, F2: uint64(v4)}
        _ = fi56
        _ = v57
    }
    if ((v5 % ((v5) | 1)) == 42) {
        if (3.14159 != 0.1) {
            var v58 float32 = (v15 * (v15 * ((1e10 / (3.14159 * float32(1.5))) - float32(0.1))))
            _ = v58
        } else {
            v5 = v5
            var v59 uint16 = v4
            _ = v59
        }
        v60 := GenS4{F0: 177, F1: 4294967295}
        var v61 int16 = (v5 / ((v5) | 1))
        _ = v60
        _ = v61
    } else {
        v6 = append(v6, v2)
        v62 := GenS5{F0: 42}
        _ = v62
    }
    fmt.Println(v2, v53)
    fmt.Println(v5)
    switch -377 {
    case 32768:
        v63 := v3[17]
        _ = v63
    default:
        if v1 {
            if v1 {
                v15 += 0.0
                if v1 {
                    for fi64 := 0; fi64 < 5; fi64++ {
                        v65 := GenS6{F0: 799, F1: v53, F2: 1}
                        _ = fi64
                        _ = v65
                    }
                    var v66 int = 42
                    _ = v66
                } else {
                    if ((v2 + v2) == v2) {
                        if (false || ((v1 || (v15 < (v15 - v15))) || (v4 > (v4 - v4)))) {
                            var v67 uint64 = 1
                            _ = v67
                        }
                    } else {
                        v15 = v15
                        var v68 int = 128
                        _ = v68
                    }
                    v69 := v3[65535]
                    var v70 int64 = int64(v53)
                    _ = v69
                    _ = v70
                }
                fmt.Println(v5, v53, v4)
            }
        }
    }
    if len(v6) > 0 {
        v6[0] = ""
    }
    for fi71 := 0; fi71 < 9; fi71++ {
        var v72 uint32 = 42
        if (true || false) {
            v72 -= (v72 + v72)
            switch fi71 {
            case 1:
                v53 = v53
            case 128:
                v53 = (v53 + v53)
            default:
                var v73 uint16 = (v4 - uint16(v72))
                _ = v73
            }
        } else {
            for fi74 := 0; fi74 < 7; fi74++ {
                v5 = int16(fi71)
                v53 = -1
                _ = fi74
            }
        }
        fmt.Println(v4, v15, v5)
        _ = fi71
        _ = v72
    }
    for fi75 := 0; fi75 < 5; fi75++ {
        fmt.Println(v15, v1, v2)
        delete(v3, uint32(v4))
        _ = fi75
    }
    for fi76 := 0; fi76 < 12; fi76++ {
        v77 := GenS4{F0: fi76, F1: 688}
        var v78 int64 = -675
        _ = fi76
        _ = v77
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
}
