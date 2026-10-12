package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 bool }

func main() {
    var v1 uint16 = 0
    var v2 float64 = 3.14159
    v3 := make(map[uint32]bool)
    var v4 float64 = (3.14159 + (1.5 / 1.5))
    var v5 bool = (460 < 582)
    var v6 float32 = (3.14159 - (0.0 * (3.14159 + -2.25)))
    fmt.Println(v2)
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint16 = 1
        var v9 float64 = (((float64(v6) * ((float64((v6 - 1e10)) - v2) + 0.1)) * v4) / v2)
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint64 = 18446744073709551615
    var v11 string = "hello"
    fmt.Println(v11, v1)
    switch 468 {
    case 128:
        fmt.Println(v2)
    case 32768:
        v12 := GenS1{F0: -338}
        _ = v12
    case 2147483647:
        v1 -= v1
    }
    var v13 int16 = 32767
    v14 := GenS1{F0: -760}
    for fi15 := 0; fi15 < 14; fi15++ {
        if (((1e10 < (-2.25 * v6)) && v5) || v5) {
            v2 += (v4 * (-2.25 - (float64(v6) - -2.25)))
        }
        _ = fi15
    }
    if (false && (1 != 233)) {
        v1 = (v1 + (v1 ^ (v1 & (v1 / ((v1) | 1)))))
        _ = len(v3)
    } else {
        if v5 {
            fmt.Println(v1, v6)
            if true {
                var v16 uint = 42
                _ = v16
            }
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 uint = 491
                var v19 float32 = (v6 - (-2.25 - float32((v4 / v2))))
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        if (!v5) {
            var v20 bool = v5
            v21 := v3[1]
            if (976 != 242) {
                v3[152] = v21
                switch 128 {
                case 0:
                    var v22 float64 = v4
                    _ = v22
                case 1:
                    var v23 float32 = v6
                    _ = v23
                case 127:
                    v2 = float64(1e10)
                default:
                    var v24 bool = v20
                    _ = v24
                }
                fmt.Println(v2)
            } else {
                var v25 uint64 = v10
                v26 := []int{32767, 212}
                _ = v25
                _ = v26
            }
            _ = v20
            _ = v21
        }
    }
    if v5 {
        switch v11 {
        case "a":
            var v27 uint64 = (v10 + (v10 + (v10 - (v10 << ((65535) & 63)))))
            _ = v27
        case "\n\t":
            var v28 int32 = 42
            _ = v28
        default:
            v29 := GenS1{F0: -9223372036854775808}
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 16; fi30++ {
        for fi31 := 0; fi31 < 1; fi31++ {
            v11 = ((v11 + ("foo bar" + "foo bar")) + "foo bar")
            var v32 int64 = 42
            var v33 int32 = -926
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    var v34 uint8 = 0
    for fi35 := 0; fi35 < 4; fi35++ {
        switch (fi35 >> ((614) & 63)) {
        case 128:
            var v36 int = fi35
            _ = v36
        case 32768:
            switch "hello" {
            case "":
                switch (fi35 - fi35) {
                case 42:
                    var v37 int32 = -1
                    _ = v37
                }
            case "\n\t":
                var v38 uint64 = (v10 * (v10 + v10))
                _ = v38
            case "foo bar":
                for fi39 := 0; fi39 < 18; fi39++ {
                    switch "a" {
                    case "hello":
                        v40 := GenS1{F0: 278}
                        _ = v40
                    case "foo bar":
                        if true {
                            var v41 uint = 255
                            _ = v41
                        } else {
                            v34 = 255
                            if true {
                                var v42 int32 = -273
                                _ = v42
                            }
                        }
                    case "\n\t":
                        fmt.Println(v2, v10)
                    default:
                        switch fi39 {
                        case 0:
                            v43 := GenS2{F0: v11}
                            fmt.Println(v2, v2)
                            _ = v43
                        case 4294967295:
                            if ((v10 | (v10 + (v10 ^ v10))) > (v10 ^ v10)) {
                                var v44 int32 = -1
                                _ = v44
                            }
                        case -486:
                            v45 := GenS3{F0: 42, F1: 42, F2: -1}
                            _ = v45
                        }
                    }
                    _ = fi39
                }
            }
        case 829:
            var v46 uint64 = v10
            _ = v46
        }
        v10 = (v10 % ((v10) | 1))
        if (!(!true)) {
            fmt.Println(v10, v4, v11)
            var v47 int64 = -9223372036854775808
            _ = v47
        } else {
            fmt.Println(v11, v13)
            if v5 {
                switch "foo bar" {
                case "\n\t":
                    v13 = 42
                case "":
                    for fi48 := 0; fi48 < 1; fi48++ {
                        v49 := GenS4{F0: 991, F1: -9223372036854775808, F2: v10}
                        v50 := []int{}
                        var v51 int8 = 100
                        if v5 {
                            var v52 int8 = v51
                            var v53 uint = 217
                            _ = v52
                            _ = v53
                        } else {
                            fmt.Println(v34, fi48, v2)
                            if false {
                                v54 := GenS1{F0: 9223372036854775807}
                                var v55 uint64 = v10
                                _ = v54
                                _ = v55
                            } else {
                                fmt.Println(fi35)
                                v1 -= v1
                                switch fi35 {
                                case -1:
                                    _ = len(v3)
                                }
                                if ((((1e10 / ((v2 / v2) * float64((v6 + (v6 + 0.1))))) / 1.5) * (0.0 - 0.1)) < v4) {
                                    switch -2147483648 {
                                    case 0:
                                        var v56 uint = 65535
                                        _ = v56
                                    default:
                                        v5 = false
                                    }
                                }
                                var v57 float32 = v6
                                _ = v57
                            }
                        }
                        _ = fi48
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                default:
                    v3[2147483648] = v5
                }
            } else {
                v13 += (v13 & (v13 >> ((uint(v1)) & 63)))
                for fi58 := 0; fi58 < 20; fi58++ {
                    v5 = ((-(fi58)) == fi58)
                    switch v11 {
                    case "a":
                        var v59 int = (fi35 | fi58)
                        _ = v59
                    default:
                        fmt.Println(v10)
                    }
                    _ = fi58
                }
                for fi60 := 0; fi60 < 1; fi60++ {
                    v61 := GenS5{F0: 1, F1: ("a" + "")}
                    if v5 {
                        v2 = v2
                        if (100 >= -61) {
                            var v62 uint8 = v34
                            if (1.5 < (3.14159 / -2.25)) {
                                var v63 uint64 = v10
                                var v64 uint16 = v1
                                fmt.Println(v11, v6, v62)
                                _ = v63
                                _ = v64
                            } else {
                                switch ("hello" + ("" + "foo bar")) {
                                case "":
                                    var v65 uint64 = v10
                                    _ = v65
                                }
                                var v66 int16 = 0
                                _ = v66
                            }
                            if true {
                                v5 = (v2 != (v2 * v4))
                                var v67 int = 9223372036854775807
                                _ = v67
                            }
                            delete(v3, 4294967295)
                            _ = v62
                        }
                        var v68 uint = 18446744073709551615
                        if v5 {
                            v34 = v34
                            v69 := v3[1]
                            fmt.Println(fi60, fi60)
                            _ = v69
                        } else {
                            v70 := GenS6{F0: ((v5 && (!true)) && (true || v5))}
                            _ = v70
                        }
                        v71 := v3[0]
                        _ = v68
                        _ = v71
                    }
                    v72 := []int{2147483647, (-(fi35))}
                    _ = fi60
                    _ = v61
                    _ = v72
                }
            }
            v3[1] = (!true)
        }
        _ = fi35
    }
    switch -2147483648 {
    case 431:
        v1 -= 7
    case 1:
        if (((v13 - int16(v34)) >= v13) || (!true)) {
            fmt.Println(v34, v34, v6)
        }
    case 9223372036854775807:
        v11 = ((((v11 + v11) + (("foo bar" + "") + v11)) + "\n\t") + "")
    }
    var v73 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v73
}
