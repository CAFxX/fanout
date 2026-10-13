package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int32; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS5 struct { F0 int; F1 uint8 }
func genFn0() uint64 {
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 float64 = 0.1
        v9 := GenS1{F0: 109}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 uint16 = 255
        if true {
            v11 = (^(v11))
            fmt.Println(v11, fi10)
        } else {
            if (true && (true || (1.5 < (float64(((float32(1.5) * (3.14159 * 0.1)) - 1e10)) * 0.0)))) {
                switch "a" {
                case "foo bar":
                    var v12 string = "foo bar"
                    _ = v12
                case "\n\t":
                    if (4294967295 < 4294967295) {
                        v13 := GenS2{F0: 2, F1: 65535}
                        for fi14 := 0; fi14 < 9; fi14++ {
                            v15 := GenS1{F0: v11}
                            if true {
                                var v16 uint8 = 110
                                var v17 int64 = 80
                                v18 := []int{}
                                var v19 int16 = 42
                                var v20 uint64 = 2
                                _ = v16
                                _ = v17
                                _ = v18
                                _ = v19
                                _ = v20
                            } else {
                                if (uint(v11) != 4294967295) {
                                    if false {
                                        var v21 float64 = (0.1 + -2.25)
                                        var v22 uint = 574
                                        var v23 bool = (!(!true))
                                        if v23 {
                                            if (2 == 2147483647) {
                                                if (false && v23) {
                                                    v24 := GenS2{F0: 65535, F1: (^(v22))}
                                                    v11 = v11
                                                    _ = v24
                                                } else {
                                                    v25 := []int{}
                                                    _ = v25
                                                }
                                                v23 = (117 > -2147483648)
                                                v23 = ("hello" <= "a")
                                            }
                                            v21 += -2.25
                                        }
                                        v26 := GenS1{F0: v11}
                                        _ = v21
                                        _ = v22
                                        _ = v23
                                        _ = v26
                                    } else {
                                        fmt.Println(fi14, v11)
                                        v11 = (v11 + (v11 >> ((42) & 63)))
                                        if false {
                                            var v27 uint8 = 0
                                            _ = v27
                                        } else {
                                            switch 4294967295 {
                                            case 32768:
                                                v11 += (v11 + 32767)
                                            }
                                            v11 = v11
                                            var v28 float64 = float64(-2.25)
                                            var v29 int8 = 42
                                            _ = v28
                                            _ = v29
                                        }
                                        if ("foo bar" == "\n\t") {
                                            if false {
                                                fmt.Println(fi14, fi14)
                                                v30 := []int{fi14, 32767}
                                                if (!((128 > 1) || (1000 != 1))) {
                                                    if len(v30) >= 2 {
                                                        v30 = v30[1:len(v30)]
                                                    }
                                                    v11 -= (v11 >> ((65535) & 63))
                                                    var v31 float64 = 1e10
                                                    _ = v31
                                                }
                                                _ = v30
                                            }
                                        }
                                    }
                                    if false {
                                        if true {
                                            v11 += v11
                                            v32 := []int{(fi10 << ((271) & 63)), (fi14 % ((fi14) | 1))}
                                            var v33 uint16 = v11
                                            _ = v32
                                            _ = v33
                                        }
                                        v34 := []int{0, (fi10 ^ 32768), fi14, fi10}
                                        switch (fi10 % ((0) | 1)) {
                                        case 699:
                                            v11 = v11
                                        case -913:
                                            var v35 bool = (1 > 4294967295)
                                            _ = v35
                                        case 128:
                                            if false {
                                                var v36 uint = 4294967295
                                                _ = v36
                                            } else {
                                                var v37 bool = false
                                                _ = v37
                                            }
                                        default:
                                            switch fi14 {
                                            case 340:
                                                var v38 int16 = int16(v11)
                                                _ = v38
                                            case -2147483648:
                                                v11 = 0
                                            case 9223372036854775807:
                                                fmt.Println(fi10)
                                            }
                                        }
                                        _ = v34
                                    } else {
                                        var v39 uint16 = 774
                                        _ = v39
                                    }
                                } else {
                                    var v40 int32 = 1
                                    var v41 float64 = (0.1 / -2.25)
                                    _ = v40
                                    _ = v41
                                }
                                v42 := GenS1{F0: (v11 << ((255) & 63))}
                                v43 := []int{(fi14 % (((fi14 / ((128) | 1))) | 1)), fi14, fi10}
                                _ = v42
                                _ = v43
                            }
                            _ = fi14
                            _ = v15
                        }
                        var v44 int64 = 1
                        _ = v13
                        _ = v44
                    }
                }
                var v45 int = 9223372036854775807
                v46 := GenS1{F0: v11}
                _ = v45
                _ = v46
            } else {
                fmt.Println(fi10, fi10)
            }
            var v47 string = "\n\t"
            _ = v47
        }
        _ = fi10
        _ = v11
    }
    v48 := GenS1{F0: 255}
    var v49 bool = false
    v50 := []int{941}
    _ = v48
    _ = v49
    _ = v50
    return 42
}

func main() {
    var v1 int64 = -253
    var v2 uint8 = 2
    v3 := make(map[int]int)
    var v4 float64 = 1e10
    v5 := make(map[uint32]string)
    var v6 uint64 = 117
    v51 := GenS1{F0: 42}
    v6 += v6
    switch "foo bar" {
    case "hello":
        for fi52 := 0; fi52 < 3; fi52++ {
            v1 += v1
            switch ("\n\t" + "foo bar") {
            case "foo bar":
                var v53 uint = 4294967295
                _ = v53
            default:
                for fi54 := 0; fi54 < 8; fi54++ {
                    if (true || (-110 >= -128)) {
                        switch ("hello" + (("hello" + ("" + "foo bar")) + ("a" + "\n\t"))) {
                        case "":
                            fmt.Println(fi54)
                        case "a":
                            switch "foo bar" {
                            case "":
                                var v55 uint16 = 593
                                _ = v55
                            default:
                                if ("a" < "\n\t") {
                                    var v56 float32 = 0.0
                                    v57 := GenS3{F0: int32(v6), F1: 2147483647, F2: v2}
                                    _ = v56
                                    _ = v57
                                } else {
                                    v58 := []int{fi54, 128, fi52}
                                    v59 := v5[947]
                                    _ = v58
                                    _ = v59
                                }
                            }
                        case "hello":
                            v60 := GenS2{F0: 661, F1: 2}
                            _ = v60
                        default:
                            fmt.Println(v4, v1, fi54)
                        }
                    } else {
                        v61 := []int{32767, fi52}
                        var v62 int32 = int32(v6)
                        v63 := GenS4{F0: 210, F1: 34, F2: (v62 % (((v62 + v62)) | 1))}
                        _ = v61
                        _ = v62
                        _ = v63
                    }
                    switch (fi52 << ((42) & 63)) {
                    case 0:
                        var v64 uint16 = 230
                        switch (("a" + "a") + (("\n\t" + "hello") + "hello")) {
                        case "hello":
                            v65 := GenS1{F0: v64}
                            _ = v65
                        case "":
                            v66 := GenS1{F0: 32767}
                            _ = v66
                        default:
                            v1 += v1
                        }
                        switch "foo bar" {
                        case "":
                            v64 = v64
                        }
                        v4 += float64(0.0)
                        _ = v64
                    default:
                        v67 := GenS1{F0: uint16(v2)}
                        _ = v67
                    }
                    var v68 float32 = (((1e10 / float32(v4)) + ((1e10 * 3.14159) - -2.25)) * 1.5)
                    _ = fi54
                    _ = v68
                }
            }
            _ = fi52
        }
    case "\n\t":
        v2 += (v2 << ((2) & 63))
    case "foo bar":
        v6 = v6
    default:
        for fi69 := 0; fi69 < 18; fi69++ {
            v6 -= v6
            var v70 float64 = v4
            _ = fi69
            _ = v70
        }
    }
    for fi71 := 0; fi71 < 5; fi71++ {
        var v72 int = (fi71 / ((-2147483648) | 1))
        var v73 int64 = v1
        for fi74 := 0; fi74 < 16; fi74++ {
            if false {
                var v75 float32 = -2.25
                var v76 int16 = 32767
                _ = v75
                _ = v76
            }
            v2 += (^(v2))
            _ = fi74
        }
        _ = fi71
        _ = v72
        _ = v73
    }
    delete(v5, 0)
    var v77 int16 = -907
    v78 := GenS5{F0: -2147483648, F1: (^(v2))}
    v79 := v5[0]
    fmt.Println(v77)
    for fi80 := 0; fi80 < 19; fi80++ {
        if false {
            switch fi80 {
            case 42:
                v81 := []int{(-(fi80))}
                _ = v81
            default:
                var v82 bool = (127 < (v2 >> ((65535) & 63)))
                _ = v82
            }
        } else {
            switch "foo bar" {
            case "a":
                var v83 int64 = v1
                _ = v83
            default:
                delete(v5, 65535)
            }
        }
        if (((-(fi80)) <= fi80) && (v4 != -2.25)) {
            if true {
                if false {
                    switch (v79 + v79) {
                    case "foo bar":
                        switch v79 {
                        case "hello":
                            if (127 > -1) {
                                switch v79 {
                                case "":
                                    var v84 float64 = v4
                                    _ = v84
                                case "\n\t":
                                    v4 = 0.0
                                case "foo bar":
                                    var v85 float64 = 0.0
                                    _ = v85
                                default:
                                    v86 := []int{128, int(v6), (fi80 - fi80)}
                                    _ = v86
                                }
                            }
                        case "foo bar":
                            fmt.Println(v79)
                        case "a":
                            if true {
                                v87 := GenS1{F0: 2}
                                _ = v87
                            }
                        default:
                            for fi88 := 0; fi88 < 1; fi88++ {
                                v4 -= v4
                                _ = fi88
                            }
                        }
                    default:
                        fmt.Println(v6, v2)
                    }
                    var v89 uint16 = 32767
                    switch (fi80 + (fi80 % ((fi80) | 1))) {
                    case 1:
                        v77 += -1
                    default:
                        var v90 int32 = -1
                        _ = v90
                    }
                    _ = v89
                } else {
                    if (v2 > v2) {
                        if ((255 > 0) && true) {
                            var v91 float64 = float64(0.0)
                            _ = v91
                        }
                    }
                    v77 -= -146
                    v92 := []int{fi80, (fi80 & (fi80 % ((fi80) | 1)))}
                    _ = v92
                }
            }
        }
        _ = fi80
    }
    var v93 uint64 = v6
    v93 = v93
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    fmt.Println("v79:", v79)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v77
    _ = v78
    _ = v79
    _ = v93
}
