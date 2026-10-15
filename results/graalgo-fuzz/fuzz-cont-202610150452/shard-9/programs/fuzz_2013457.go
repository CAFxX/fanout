package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint16; F2 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
func genFn0() uint32 {
    v5 := []int{}
    for fi6 := 0; fi6 < 1; fi6++ {
        _ = len(v5)
        var v7 int = fi6
        _ = fi6
        _ = v7
    }
    var v8 float64 = 0.0
    _ = v5
    _ = v8
    return 2147483647
}
func genFn1() int16 {
    v9 := []int{2147483647}
    _ = len(v9)
    var v10 uint16 = 42
    if true {
        if (2 < 0) {
            if (float32(0.1) > 0.1) {
                for fi11 := 0; fi11 < 4; fi11++ {
                    for fi12 := 0; fi12 < 1; fi12++ {
                        v10 = 2
                        v10 = v10
                        _ = fi12
                    }
                    var v13 int8 = 1
                    fmt.Println(fi11)
                    _ = fi11
                    _ = v13
                }
                v14 := GenS1{F0: 42}
                v15 := GenS1{F0: 12}
                _ = v14
                _ = v15
            }
            if true {
                for fi16 := 0; fi16 < 16; fi16++ {
                    v17 := GenS2{F0: 4294967295, F1: (^(v10)), F2: 42}
                    v10 += 255
                    var v18 bool = true
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                if true {
                    fmt.Println(v10)
                }
            } else {
                _ = len(v9)
                switch 0 {
                case -2147483648:
                    if (!(-2147483648 > -600)) {
                        fmt.Println(v10)
                        if (1 >= 255) {
                            for fi19 := 0; fi19 < 13; fi19++ {
                                switch fi19 {
                                case -715:
                                    switch (("foo bar" + (("hello" + ("a" + ("hello" + ("a" + ("a" + ""))))) + "\n\t")) + "foo bar") {
                                    case "hello":
                                        v10 -= v10
                                    default:
                                        switch ("" + ("" + ("" + ""))) {
                                        case "a":
                                            v10 = (v10 % ((v10) | 1))
                                        case "foo bar":
                                            fmt.Println(v10, fi19)
                                        }
                                    }
                                }
                                v9 = append(v9, (fi19 * (fi19 >> ((1) & 63))))
                                var v20 uint16 = (v10 % (((v10 >> ((255) & 63))) | 1))
                                _ = fi19
                                _ = v20
                            }
                            for fi21 := 0; fi21 < 14; fi21++ {
                                if true {
                                    for fi22 := 0; fi22 < 4; fi22++ {
                                        var v23 uint16 = (^(v10))
                                        switch ("hello" + "\n\t") {
                                        case "foo bar":
                                            v24 := GenS1{F0: -128}
                                            _ = v24
                                        case "":
                                            fmt.Println(v23, fi21)
                                        case "a":
                                            switch "hello" {
                                            case "foo bar":
                                                v10 = v10
                                            case "hello":
                                                switch int(v23) {
                                                case 784:
                                                    var v25 float64 = 1.5
                                                    _ = v25
                                                default:
                                                    var v26 uint8 = 255
                                                    _ = v26
                                                }
                                            default:
                                                fmt.Println(fi22)
                                            }
                                        default:
                                            v9 = append(v9, fi22)
                                        }
                                        if true {
                                            var v27 string = "hello"
                                            v28 := GenS1{F0: 1}
                                            fmt.Println(fi21, v27)
                                            v29 := GenS1{F0: 48}
                                            v10 = v23
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        } else {
                                            fmt.Println(v10, fi21)
                                            if len(v9) > 0 {
                                                v9[0] = fi21
                                            }
                                        }
                                        v9 = append(v9, 838)
                                        _ = fi22
                                        _ = v23
                                    }
                                    v10 += v10
                                    fmt.Println(v10)
                                } else {
                                    for fi30 := 0; fi30 < 15; fi30++ {
                                        if false {
                                            if (0 <= 9223372036854775807) {
                                                v31 := GenS3{F0: -128}
                                                if (2 >= 612) {
                                                    var v32 int8 = 100
                                                    var v33 string = ((((("\n\t" + ("" + "foo bar")) + "\n\t") + "foo bar") + "hello") + "")
                                                    switch v33 {
                                                    case "hello":
                                                        _ = len(v9)
                                                    case "foo bar":
                                                        var v34 uint32 = 65535
                                                        _ = v34
                                                    default:
                                                        var v35 uint32 = 2147483648
                                                        _ = v35
                                                    }
                                                    var v36 uint32 = 4294967295
                                                    v33 = v33
                                                    _ = v32
                                                    _ = v33
                                                    _ = v36
                                                }
                                                var v37 float64 = 1.5
                                                v10 = 32767
                                                _ = v31
                                                _ = v37
                                            }
                                            v9 = append(v9, fi30)
                                            switch (fi21 % ((32768) | 1)) {
                                            case 0:
                                                var v38 bool = true
                                                _ = v38
                                            default:
                                                fmt.Println(fi30, fi21, fi30)
                                            }
                                        }
                                        var v39 uint32 = 2147483647
                                        var v40 bool = true
                                        v9 = append(v9, (fi30 | -451))
                                        _ = fi30
                                        _ = v39
                                        _ = v40
                                    }
                                    v10 = (v10 ^ v10)
                                    if true {
                                        fmt.Println(v10, v10)
                                        for fi41 := 0; fi41 < 2; fi41++ {
                                            v10 += (v10 << ((42) & 63))
                                            v10 += v10
                                            var v42 int16 = -32768
                                            _ = fi41
                                            _ = v42
                                        }
                                    } else {
                                        if false {
                                            fmt.Println(v10)
                                        }
                                        v43 := GenS4{F0: 2147483647}
                                        if (true || false) {
                                            fmt.Println(v10, v10)
                                            var v44 int = fi21
                                            var v45 uint = 1
                                            _ = v44
                                            _ = v45
                                        }
                                        _ = v43
                                    }
                                }
                                _ = fi21
                            }
                        } else {
                            var v46 string = ("" + "")
                            var v47 float32 = (float32(float64((-2.25 / 3.14159))) - (-2.25 * float32((((3.14159 - 1.5) - (((3.14159 / 1e10) / -2.25) - ((3.14159 / 1.5) / 1.5))) - (3.14159 * 0.1)))))
                            _ = v46
                            _ = v47
                        }
                        if false {
                            v48 := GenS1{F0: 42}
                            var v49 uint16 = 0
                            if true {
                                var v50 int64 = 0
                                _ = v50
                            } else {
                                for fi51 := 0; fi51 < 11; fi51++ {
                                    fmt.Println(v10, v49)
                                    var v52 uint64 = 0
                                    _ = fi51
                                    _ = v52
                                }
                                v53 := GenS2{F0: 4294967295, F1: v10, F2: 0}
                                for fi54 := 0; fi54 < 1; fi54++ {
                                    switch "foo bar" {
                                    case "hello":
                                        switch -1 {
                                        case -903:
                                            v49 = v49
                                        case 977:
                                            var v55 uint32 = genFn0()
                                            _ = v55
                                        case 128:
                                            switch (fi54 % ((-2147483648) | 1)) {
                                            case -399:
                                                var v56 string = ""
                                                _ = v56
                                            case -9223372036854775808:
                                                var v57 bool = (!true)
                                                _ = v57
                                            }
                                        default:
                                            var v58 int32 = -571
                                            _ = v58
                                        }
                                    case "foo bar":
                                        if ((^(v10)) >= 334) {
                                            v49 -= (^(v10))
                                            v59 := GenS3{F0: -128}
                                            switch (fi54 << ((1) & 63)) {
                                            case 1:
                                                var v60 uint = 18446744073709551615
                                                _ = v60
                                            case -268:
                                                v49 = (v49 ^ 2)
                                            case -9223372036854775808:
                                                fmt.Println(v10, v10)
                                            default:
                                                switch ("" + "a") {
                                                case "hello":
                                                    var v61 string = ((("" + "hello") + "\n\t") + "hello")
                                                    _ = v61
                                                case "\n\t":
                                                    for fi62 := 0; fi62 < 12; fi62++ {
                                                        fmt.Println(fi62, v10)
                                                        v63 := GenS2{F0: 4294967295, F1: 2, F2: -32768}
                                                        _ = fi62
                                                        _ = v63
                                                    }
                                                case "":
                                                    var v64 float64 = -2.25
                                                    _ = v64
                                                }
                                            }
                                            _ = v59
                                        }
                                    }
                                    if true {
                                        if (2147483647 <= 0) {
                                            v49 = 255
                                            fmt.Println(v10, v10)
                                            v49 -= 255
                                        }
                                        fmt.Println(v49, v49, v49)
                                        var v65 string = "a"
                                        _ = v65
                                    }
                                    _ = fi54
                                }
                                _ = v53
                            }
                            _ = v48
                            _ = v49
                        }
                    }
                case 32768:
                    v10 += 1
                case 42:
                    var v66 int8 = 1
                    _ = v66
                }
            }
        }
        var v67 int64 = 42
        _ = v67
    } else {
        for fi68 := 0; fi68 < 4; fi68++ {
            v10 = v10
            _ = len(v9)
            fmt.Println(fi68, v10)
            _ = fi68
        }
        var v69 int32 = 42
        if false {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            for fi70 := 0; fi70 < 5; fi70++ {
                v69 = (v69 - -520)
                var v71 uint64 = 2
                v71 = v71
                _ = fi70
                _ = v71
            }
        }
        _ = v69
    }
    _ = v10
    _ = v9
    return -637
}

func main() {
    var v1 float32 = ((0.0 + (0.0 * (float32(float64(1e10)) / 1e10))) * 3.14159)
    var v2 uint = 988
    var v3 int16 = -32768
    v4 := []string{("foo bar" + "foo bar"), ("a" + ("foo bar" + "hello")), "hello"}
    fmt.Println(v1, v2)
    var v72 int32 = -1
    v2 = v2
    if ("foo bar" != "hello") {
        if false {
            switch 42 {
            case -9223372036854775808:
                fmt.Println(v1)
            }
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v73 int = 471
        _ = v73
    } else {
        var v74 int16 = v3
        _ = v74
    }
    var v75 uint = v2
    v1 += v1
    v76 := GenS2{F0: (v75 | v75), F1: 1, F2: genFn1()}
    if false {
        fmt.Println(v75, v3, v75)
        fmt.Println(v72, v2, v3)
        v3 = genFn1()
    }
    if false {
        if ((v3 | 0) > v3) {
            v72 = (-(v72))
            var v77 int32 = v72
            v2 += (v2 + (^(v75)))
            _ = v77
        } else {
            var v78 uint = v75
            v78 -= v2
            if (true && (!true)) {
                v79 := GenS2{F0: 1, F1: 2, F2: v3}
                _ = v79
            } else {
                for fi80 := 0; fi80 < 3; fi80++ {
                    switch 0 {
                    case 16:
                        v72 -= v72
                    }
                    if len(v4) > 0 {
                        v4[0] = "foo bar"
                    }
                    v3 = genFn1()
                    _ = fi80
                }
                v2 = 533
            }
            _ = v78
        }
        _ = len(v4)
    }
    var v81 uint16 = 2
    switch ("foo bar" + "a") {
    case "a":
        v2 += uint(v81)
    case "\n\t":
        v2 = v75
    case "hello":
        for fi82 := 0; fi82 < 1; fi82++ {
            var v83 uint64 = 18446744073709551615
            var v84 int8 = 127
            if ((4294967295 == 2) && false) {
                if true {
                    v1 = v1
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                    var v85 uint8 = 128
                    _ = v85
                } else {
                    for fi86 := 0; fi86 < 12; fi86++ {
                        var v87 uint = 204
                        var v88 bool = false
                        var v89 int8 = (v84 / (((v84 * (v84 - (v84 ^ 42)))) | 1))
                        _ = fi86
                        _ = v87
                        _ = v88
                        _ = v89
                    }
                }
            } else {
                var v90 int64 = -9223372036854775808
                v72 += v72
                var v91 int32 = v72
                _ = v90
                _ = v91
            }
            _ = fi82
            _ = v83
            _ = v84
        }
    default:
        v2 = (v75 % ((65535) | 1))
    }
    v92 := GenS1{F0: -34}
    if (2147483647 != 34) {
        fmt.Println(v81, v1)
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v72:", v72)
    fmt.Println("v75:", v75)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v72
    _ = v75
    _ = v76
    _ = v81
    _ = v92
}
