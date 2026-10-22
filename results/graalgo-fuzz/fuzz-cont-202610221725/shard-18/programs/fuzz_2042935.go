package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint8; F2 string }
type GenS4 struct { F0 int }
func genFn0(p0 uint64, p1 uint16) bool {
    fmt.Println(p0)
    if false {
        p1 += (p1 * p1)
    }
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            p1 -= p1
            _ = fi6
        }
        for fi7 := 0; fi7 < 4; fi7++ {
            p1 = (p1 + (p1 | (p1 / ((42) | 1))))
            p1 -= (p1 << ((0) & 63))
            switch fi5 {
            case 744:
                var v8 string = (("foo bar" + "hello") + ("a" + "foo bar"))
                _ = v8
            case 127:
                var v9 uint64 = 2
                _ = v9
            default:
                p1 = uint16(p0)
            }
            var v10 int64 = 0
            _ = fi7
            _ = v10
        }
        v11 := []int{-313, 32768, fi5}
        _ = fi5
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS1{F0: ""}
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    return false
}
func genFn1() string {
    switch ("" + ("foo bar" + "a")) {
    case "":
        for fi14 := 0; fi14 < 5; fi14++ {
            fmt.Println(fi14)
            _ = fi14
        }
    case "a":
        var v15 uint8 = 255
        _ = v15
    case "\n\t":
        var v16 uint8 = 104
        _ = v16
    }
    var v17 float32 = (-2.25 - 0.1)
    var v18 float64 = -2.25
    _ = v17
    _ = v18
    return "hello"
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = (3.14159 + (3.14159 / ((0.1 * -2.25) + (-2.25 / float32(1e10)))))
    v3 := make(map[string]string)
    var v4 uint32 = 65535
    v19 := []int{-182, 0}
    if true {
        var v20 uint32 = v4
        _ = v20
    } else {
        var v21 int32 = 2147483647
        _ = v21
    }
    var v22 uint = 0
    for fi23 := 0; fi23 < 2; fi23++ {
        var v24 uint8 = uint8(v22)
        if false {
            if false {
                fmt.Println(v2, v22, v22)
                if true {
                    if true {
                        v1[(^(v4))] = false
                        v25 := GenS2{F0: 674}
                        if (v4 > 1) {
                            v24 = v24
                        }
                        _ = v25
                    } else {
                        var v26 float64 = (1.5 / (0.1 / 1.0))
                        _ = v26
                    }
                    switch (fi23 % ((int(v22)) | 1)) {
                    case -852:
                        v24 = 155
                    case -1:
                        for fi27 := 0; fi27 < 18; fi27++ {
                            if false {
                                if true {
                                    v4 = (v4 / (((v4 & v4)) | 1))
                                }
                                var v28 int16 = -404
                                _ = v28
                            }
                            _ = fi27
                        }
                    case 128:
                        if (true && true) {
                            if (("hello" + "hello") == "a") {
                                if (799 >= 849) {
                                    fmt.Println(v24, v4)
                                    var v29 uint64 = 4294967295
                                    v22 += (v22 >> ((4294967295) & 63))
                                    _ = v29
                                }
                            }
                            v4 += 0
                            for fi30 := 0; fi30 < 2; fi30++ {
                                var v31 int32 = 1
                                fmt.Println(v24, fi23, fi30)
                                var v32 uint16 = 32767
                                v33 := v1[65535]
                                delete(v3, "")
                                v34 := GenS1{F0: ("hello" + "a")}
                                _ = fi30
                                _ = v31
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                        } else {
                            switch "" {
                            case "":
                                if ((true && (1 >= uint16(v4))) && true) {
                                    var v35 uint64 = 42
                                    v36 := GenS1{F0: "foo bar"}
                                    var v37 float32 = v2
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                } else {
                                    v19 = append(v19, (fi23 / (((^(fi23))) | 1)))
                                }
                            case "hello":
                                fmt.Println(v22, v4, v22)
                            default:
                                if (!(2 > 255)) {
                                    v4 += (v4 & v4)
                                    if (42 != 934) {
                                        if true {
                                            var v38 int64 = 1
                                            var v39 uint8 = (v24 & 127)
                                            _ = v38
                                            _ = v39
                                        }
                                        var v40 int8 = 100
                                        _ = v40
                                    } else {
                                        delete(v1, uint32(v24))
                                    }
                                } else {
                                    v2 -= (v2 * float32(1.5))
                                    v41 := GenS1{F0: "\n\t"}
                                    var v42 uint16 = 65535
                                    _ = v41
                                    _ = v42
                                }
                            }
                        }
                    }
                }
                v43 := GenS3{F0: 100, F1: (^(v24)), F2: "hello"}
                _ = v43
            }
            v4 -= 42
        }
        _ = fi23
        _ = v24
    }
    var v44 int = 32768
    for fi45 := 0; fi45 < 3; fi45++ {
        switch v44 {
        case 521:
            if true {
                switch "\n\t" {
                case "":
                    v46 := GenS3{F0: 48, F1: 26, F2: "\n\t"}
                    _ = v46
                default:
                    v4 = v4
                }
            } else {
                delete(v1, (v4 ^ v4))
                v4 = 196
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
            }
        }
        v47 := v1[v4]
        _ = fi45
        _ = v47
    }
    if false {
        var v48 bool = (0 != 42)
        v3["foo bar"] = "hello"
        _ = v48
    } else {
        v4 += 2147483648
    }
    for fi49 := 0; fi49 < 10; fi49++ {
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        v22 -= (^(v22))
        if (989 <= 255) {
            v44 = (^(v44))
            if (true && (0.1 < ((-2.25 / float64((0.1 + v2))) / 1e10))) {
                delete(v3, "")
            } else {
                fmt.Println(v22)
                if (!false) {
                    var v50 uint64 = 4294967295
                    var v51 uint16 = 1
                    v52 := GenS2{F0: -2147483648}
                    _ = v50
                    _ = v51
                    _ = v52
                }
                fmt.Println(v4, fi49)
            }
            v22 = 42
        }
        _ = fi49
    }
    if true {
        var v53 float32 = v2
        for fi54 := 0; fi54 < 14; fi54++ {
            if true {
                var v55 int = -1
                var v56 uint8 = 42
                v55 += (fi54 | -808)
                _ = v55
                _ = v56
            } else {
                var v57 float32 = v53
                var v58 int16 = 1
                _ = v57
                _ = v58
            }
            switch "a" {
            case "hello":
                var v59 uint8 = 127
                _ = v59
            case "foo bar":
                switch "\n\t" {
                case "":
                    var v60 int64 = -1
                    _ = v60
                case "\n\t":
                    v53 -= 1e10
                }
            case "\n\t":
                for fi61 := 0; fi61 < 5; fi61++ {
                    var v62 int16 = 620
                    _ = fi61
                    _ = v62
                }
            }
            _ = fi54
        }
        _ = v53
    } else {
        if len(v19) > 0 {
            v19[0] = (-(v44))
        }
    }
    v3["foo bar"] = "foo bar"
    v4 -= (v4 - (v4 % (((v4 ^ v4)) | 1)))
    for fi63 := 0; fi63 < 14; fi63++ {
        v64 := GenS3{F0: 1, F1: 71, F2: "\n\t"}
        for fi65 := 0; fi65 < 16; fi65++ {
            v66 := GenS3{F0: 0, F1: 2, F2: "\n\t"}
            if true {
                v67 := GenS2{F0: 42}
                switch ("hello" + "foo bar") {
                case "\n\t":
                    if (!(true || false)) {
                        fmt.Println(v22, fi65)
                        v44 = (fi65 ^ v44)
                        v19 = append(v19, v44)
                    }
                case "a":
                    fmt.Println(v2)
                }
                _ = len(v1)
                var v68 int64 = 0
                _ = v67
                _ = v68
            } else {
                fmt.Println(v44)
                v69 := GenS2{F0: 1}
                _ = v69
            }
            fmt.Println(v2)
            _ = fi65
            _ = v66
        }
        for fi70 := 0; fi70 < 18; fi70++ {
            var v71 uint16 = 930
            v44 -= (fi70 - (v44 % (((fi63 * 2147483647)) | 1)))
            v72 := GenS2{F0: 873}
            _ = fi70
            _ = v71
            _ = v72
        }
        _ = fi63
        _ = v64
    }
    v44 = 2147483647
    var v73 int8 = 100
    if true {
        if (((1 > (v22 / ((1) | 1))) && true) && false) {
            var v74 uint32 = v4
            switch (((("" + "\n\t") + "") + "") + ("\n\t" + "foo bar")) {
            case "":
                v75 := GenS4{F0: -49}
                _ = v75
            case "\n\t":
                v22 = (v22 / ((65535) | 1))
            }
            _ = v74
        }
        fmt.Println(v2, v4)
    }
    switch (v44 >> ((v22) & 63)) {
    case -2147483648:
        if ((("hello" <= "foo bar") && true) && true) {
            v73 += (^(v73))
        } else {
            var v76 int16 = 0
            _ = v76
        }
    }
    delete(v3, "")
    for fi77 := 0; fi77 < 2; fi77++ {
        var v78 float64 = float64(((((v2 * (0.0 - 0.1)) * v2) - float32(0.0)) * float32((1e10 / 1e10))))
        _ = fi77
        _ = v78
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v44
    _ = v73
}
