package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint; F2 int64 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint }
type GenS5 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS6 struct { F0 int32 }
func genFn0() int8 {
    if true {
        var v7 uint = 255
        v7 = v7
        if (-139 > -145) {
            var v8 float32 = (1e10 * ((1e10 + (1.5 + (float32((-2.25 + float64(float32(0.1)))) + 0.1))) - (((float32(3.14159) * (0.0 / 1e10)) / -2.25) - (1e10 / ((1.5 / 1.0) * -2.25)))))
            var v9 float64 = (1e10 + 1.5)
            _ = v8
            _ = v9
        }
        _ = v7
    }
    if true {
        switch ((("hello" + "hello") + "\n\t") + "hello") {
        case "a":
            var v10 int8 = 123
            _ = v10
        case "foo bar":
            var v11 int32 = 161
            _ = v11
        case "hello":
            var v12 int32 = -1
            _ = v12
        }
        v13 := []int{-2147483648, 32768, 128, 128}
        for fi14 := 0; fi14 < 18; fi14++ {
            if len(v13) > 0 {
                v13[0] = fi14
            }
            _ = fi14
        }
        _ = v13
    } else {
        var v15 int64 = -9223372036854775808
        _ = v15
    }
    if true {
        var v16 int16 = -32768
        v16 -= -380
        _ = v16
    }
    return 75
}
func genFn1() uint16 {
    v17 := GenS1{F0: -128, F1: 1}
    var v18 uint8 = 127
    v19 := GenS1{F0: 100, F1: v18}
    if (0 >= 255) {
        for fi20 := 0; fi20 < 15; fi20++ {
            v21 := GenS2{F0: false}
            _ = fi20
            _ = v21
        }
        switch -398 {
        case 4294967295:
            fmt.Println(v18)
        case 9223372036854775807:
            for fi22 := 0; fi22 < 16; fi22++ {
                if ((1.5 - (1.5 + ((float32(((1e10 - 3.14159) * (float64(1.5) * float64(0.0)))) + ((0.1 / 0.1) / ((-2.25 + (1.5 / 0.1)) - -2.25))) - float32(3.14159)))) < 0.1) {
                    var v23 bool = (false || false)
                    v23 = (!v23)
                    _ = v23
                } else {
                    if (128 >= (fi22 / ((-1) | 1))) {
                        for fi24 := 0; fi24 < 2; fi24++ {
                            switch fi22 {
                            case 642:
                                v25 := GenS2{F0: true}
                                _ = v25
                            case 127:
                                v18 = (v18 / (((v18 >> ((2) & 63))) | 1))
                            case 0:
                                v26 := []int{fi24}
                                switch ("\n\t" + "") {
                                case "\n\t":
                                    var v27 uint8 = 0
                                    _ = v27
                                case "hello":
                                    v18 = v18
                                case "":
                                    if (74 == -12) {
                                        var v28 uint = 18446744073709551615
                                        v26 = append(v26, 128)
                                        fmt.Println(fi22)
                                        _ = v28
                                    } else {
                                        v29 := GenS1{F0: 1, F1: v18}
                                        switch fi24 {
                                        case 127:
                                            v18 += (v18 * v18)
                                        case -1:
                                            v18 += 255
                                        case -691:
                                            var v30 float64 = (0.1 + (-2.25 / 0.1))
                                            _ = v30
                                        }
                                        _ = v29
                                    }
                                default:
                                    v31 := GenS3{F0: 42, F1: 390, F2: 42}
                                    _ = v31
                                }
                                switch 0 {
                                case 32767:
                                    v26 = append(v26, (fi22 - fi24))
                                case 9223372036854775807:
                                    v32 := GenS1{F0: 124, F1: v18}
                                    v26 = append(v26, -170)
                                    _ = v32
                                default:
                                    v18 += v18
                                }
                                _ = v26
                            default:
                                v33 := GenS2{F0: ((fi22 >> ((683) & 63)) == (-(fi22)))}
                                _ = v33
                            }
                            switch "\n\t" {
                            case "":
                                var v34 uint16 = 0
                                _ = v34
                            default:
                                var v35 uint16 = 172
                                _ = v35
                            }
                            _ = fi24
                        }
                    }
                }
                _ = fi22
            }
        }
    }
    _ = v17
    _ = v18
    _ = v19
    return 65535
}

func main() {
    var v1 bool = true
    v2 := []string{"\n\t", "hello", ("foo bar" + ""), "\n\t"}
    var v3 uint8 = 19
    v4 := make(map[uint32]string)
    var v5 float64 = (1e10 - -2.25)
    v6 := make(map[string]bool)
    fmt.Println(v5, v3, v5)
    delete(v4, 42)
    var v36 float32 = (3.14159 + ((1.5 + 0.0) * 0.0))
    v3 -= 42
    for fi37 := 0; fi37 < 6; fi37++ {
        switch "hello" {
        case "foo bar":
            if v1 {
                var v38 uint32 = 2
                v39 := GenS4{F0: 42, F1: 127, F2: 0}
                v40 := GenS4{F0: 1, F1: -89, F2: 390}
                _ = v38
                _ = v39
                _ = v40
            } else {
                switch (fi37 - (fi37 * (fi37 >> ((871) & 63)))) {
                case 42:
                    for fi41 := 0; fi41 < 20; fi41++ {
                        switch ("a" + "\n\t") {
                        case "a":
                            v36 = float32(v5)
                        case "hello":
                            if false {
                                var v42 float32 = 1e10
                                var v43 int16 = 32767
                                v44 := GenS5{F0: 343, F1: (false && ("" <= "")), F2: genFn1()}
                                var v45 string = "foo bar"
                                _ = v42
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                        default:
                            switch 9223372036854775807 {
                            case -9223372036854775808:
                                var v46 int32 = 1
                                _ = v46
                            case 4294967295:
                                var v47 uint = 377
                                _ = v47
                            case 493:
                                switch ("hello" + "hello") {
                                case "foo bar":
                                    var v48 uint = 1
                                    _ = v48
                                case "\n\t":
                                    switch "foo bar" {
                                    case "hello":
                                        var v49 float32 = float32(v5)
                                        _ = v49
                                    }
                                case "hello":
                                    var v50 uint16 = 657
                                    _ = v50
                                default:
                                    fmt.Println(v36)
                                }
                            }
                        }
                        if (v1 && true) {
                            var v51 float32 = v36
                            if (v1 && false) {
                                if (!v1) {
                                    var v52 uint8 = 57
                                    _ = v52
                                }
                                var v53 uint64 = 4294967295
                                _ = v53
                            } else {
                                v54 := GenS1{F0: -1, F1: (v3 / (((v3 | 127)) | 1))}
                                var v55 float32 = 1e10
                                fmt.Println(fi41)
                                _ = v54
                                _ = v55
                            }
                            _ = v51
                        }
                        v5 -= -2.25
                        _ = fi41
                    }
                case -432:
                    if (v1 || v1) {
                        for fi56 := 0; fi56 < 7; fi56++ {
                            if true {
                                var v57 uint32 = 2147483647
                                v3 += 128
                                if true {
                                    fmt.Println(v1, v1)
                                    switch fi37 {
                                    case 128:
                                        switch (fi37 + fi56) {
                                        case -631:
                                            if (-2147483648 < -621) {
                                                v6[""] = (v1 && v1)
                                            }
                                        }
                                    case 507:
                                        v3 = v3
                                    default:
                                        v3 += v3
                                    }
                                    switch ("hello" + "\n\t") {
                                    case "hello":
                                        v58 := GenS4{F0: 4294967295, F1: 61, F2: 18446744073709551615}
                                        _ = v58
                                    default:
                                        v36 += -2.25
                                    }
                                }
                                _ = v57
                            }
                            v59 := GenS3{F0: 0, F1: 1, F2: int64(fi56)}
                            _ = fi56
                            _ = v59
                        }
                        var v60 uint64 = 42
                        var v61 uint8 = v3
                        _ = v60
                        _ = v61
                    } else {
                        var v62 int = (-(fi37))
                        switch "\n\t" {
                        case "hello":
                            for fi63 := 0; fi63 < 8; fi63++ {
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                                var v64 string = ((("\n\t" + "foo bar") + (("a" + ("" + (("foo bar" + (("foo bar" + "") + "")) + ""))) + "a")) + "foo bar")
                                if v1 {
                                    v36 += 3.14159
                                } else {
                                    v3 += (v3 << ((65535) & 63))
                                    var v65 uint8 = v3
                                    _ = v65
                                }
                                _ = fi63
                                _ = v64
                            }
                        case "a":
                            if false {
                                var v66 uint32 = 382
                                var v67 float32 = v36
                                _ = v66
                                _ = v67
                            } else {
                                v68 := GenS1{F0: -1, F1: 1}
                                _ = v68
                            }
                        case "\n\t":
                            v69 := GenS6{F0: -2147483648}
                            _ = v69
                        default:
                            v62 = 32768
                        }
                        var v70 uint32 = 943
                        _ = v62
                        _ = v70
                    }
                case 9223372036854775807:
                    var v71 uint32 = 4294967295
                    _ = v71
                }
            }
        case "hello":
            var v72 string = ""
            _ = v72
        }
        v3 += 42
        if v1 {
            v3 = v3
        }
        _ = fi37
    }
    if true {
        fmt.Println(v1)
        v73 := v4[1]
        fmt.Println(v1, v1, v5)
        _ = v73
    }
    v1 = v1
    for fi74 := 0; fi74 < 15; fi74++ {
        if v1 {
            var v75 uint64 = 2
            var v76 uint16 = 2
            _ = v75
            _ = v76
        } else {
            v1 = true
        }
        var v77 int8 = 127
        _ = fi74
        _ = v77
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
