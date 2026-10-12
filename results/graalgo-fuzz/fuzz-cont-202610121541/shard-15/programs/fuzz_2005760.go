package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]string)
    var v6 float64 = 0.1
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int8 = -128
        _ = fi7
        _ = v8
    }
    v9 := []int{128, 42, -125, 1}
    v6 += 0.0
    var v10 int64 = 624
    if ((0.0 * float64(1e10)) < v6) {
        for fi11 := 0; fi11 < 5; fi11++ {
            for fi12 := 0; fi12 < 10; fi12++ {
                delete(v2, fi12)
                v10 -= (^(v10))
                _ = fi12
            }
            _ = fi11
        }
    }
    v10 += (v10 % ((v10) | 1))
    var v13 int = -9223372036854775808
    switch (v13 >> ((4294967295) & 63)) {
    case 42:
        if len(v9) > 0 {
            v9[0] = 989
        }
    case -725:
        v13 = v13
    case 204:
        v10 -= (^(v10))
    }
    var v14 int32 = 743
    if true {
        v15 := GenS1{F0: (v10 & (v10 % (((v10 - v10)) | 1)))}
        _ = v15
    }
    switch (v13 >> ((245) & 63)) {
    case 4294967295:
        switch "\n\t" {
        case "hello":
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        case "\n\t":
            v16 := GenS1{F0: (v10 ^ -1)}
            _ = v16
        case "a":
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
    case 597:
        var v17 float64 = v6
        _ = v17
    }
    v14 += v14
    v10 = 42
    if true {
        if ((v10 & v10) == (v10 - 449)) {
            if false {
                for fi18 := 0; fi18 < 16; fi18++ {
                    v19 := GenS1{F0: int64(v13)}
                    v6 -= 0.0
                    v9 = append(v9, (fi18 + 2147483647))
                    _ = fi18
                    _ = v19
                }
                var v20 int32 = v14
                for fi21 := 0; fi21 < 16; fi21++ {
                    v20 -= (v20 % ((v20) | 1))
                    _ = len(v9)
                    _ = fi21
                }
                _ = v20
            } else {
                v3[2147483648] = 1
            }
        } else {
            v22 := GenS2{F0: 18446744073709551615}
            for fi23 := 0; fi23 < 16; fi23++ {
                for fi24 := 0; fi24 < 2; fi24++ {
                    var v25 uint = 2
                    switch "a" {
                    case "hello":
                        var v26 string = ((((("a" + "foo bar") + (("" + "") + "\n\t")) + ("a" + "a")) + ("" + ("a" + "a"))) + ("\n\t" + "foo bar"))
                        _ = v26
                    case "a":
                        if true {
                            fmt.Println(fi24, fi24, v6)
                            v27 := GenS1{F0: -9223372036854775808}
                            _ = v27
                        } else {
                            v6 = -2.25
                        }
                    case "\n\t":
                        fmt.Println(v10)
                    }
                    _ = fi24
                    _ = v25
                }
                _ = fi23
            }
            _ = v22
        }
        fmt.Println(v6)
    }
    if (((v13 % ((v13) | 1)) > v13) && false) {
        fmt.Println(v13, v6, v13)
        v3[2147483647] = 2147483648
    } else {
        v9 = append(v9, (v13 >> ((2) & 63)))
        if ((!false) && true) {
            var v28 bool = false
            _ = v28
        } else {
            v9 = append(v9, 42)
            v2[(v13 & (v13 % ((-535) | 1)))] = "a"
            for fi29 := 0; fi29 < 17; fi29++ {
                v14 = -540
                if true {
                    for fi30 := 0; fi30 < 7; fi30++ {
                        var v31 uint32 = 2147483648
                        if (!(!true)) {
                            var v32 float32 = (3.14159 / (((1.5 * -2.25) + 1e10) / (float32(v6) + 1e10)))
                            var v33 float64 = (v6 - (1e10 / v6))
                            switch fi29 {
                            case -670:
                                var v34 uint32 = 614
                                _ = v34
                            case -2147483648:
                                v10 = (v10 & v10)
                            case 849:
                                if (true || false) {
                                    var v35 string = ("\n\t" + "hello")
                                    fmt.Println(v10, v31, v6)
                                    v36 := GenS1{F0: v10}
                                    if false {
                                        v13 = v13
                                    } else {
                                        v37 := v1[(v35 + "foo bar")]
                                        var v38 uint32 = (v31 - (v31 ^ v31))
                                        _ = v37
                                        _ = v38
                                    }
                                    _ = v35
                                    _ = v36
                                } else {
                                    var v39 float32 = v32
                                    v6 -= v6
                                    fmt.Println(v10, v6, v31)
                                    _ = v39
                                }
                            }
                            v40 := GenS2{F0: 4294967295}
                            _ = v32
                            _ = v33
                            _ = v40
                        }
                        _ = fi30
                        _ = v31
                    }
                    var v41 int = fi29
                    fmt.Println(v13, fi29)
                    _ = v41
                } else {
                    var v42 int = v13
                    v6 = v6
                    _ = v42
                }
                _ = fi29
            }
        }
    }
    for fi43 := 0; fi43 < 1; fi43++ {
        var v44 int = fi43
        var v45 int32 = -2147483648
        for fi46 := 0; fi46 < 20; fi46++ {
            if ((-(v45)) >= v14) {
                _ = len(v9)
                fmt.Println(fi46, v10, fi43)
                v47 := GenS1{F0: 42}
                _ = len(v4)
                switch ("\n\t" + "a") {
                case "hello":
                    switch fi43 {
                    case -595:
                        if false {
                            var v48 uint8 = 42
                            var v49 uint8 = v48
                            _ = v48
                            _ = v49
                        } else {
                            if (v6 != (((float64(-2.25) - v6) / v6) + 0.1)) {
                                v50 := GenS3{F0: 4294967295, F1: 65535}
                                v51 := GenS4{F0: 546}
                                v9 = append(v9, v13)
                                var v52 int8 = -77
                                var v53 float64 = 1.5
                                _ = v50
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                            switch v44 {
                            case 191:
                                var v54 uint = 255
                                _ = v54
                            default:
                                switch ("hello" + "foo bar") {
                                case "foo bar":
                                    v55 := v4[v13]
                                    _ = v55
                                }
                            }
                            v9 = append(v9, v13)
                        }
                    case 1:
                        var v56 float64 = (0.1 + (((1.5 - float64(0.1)) * (v6 + float64(0.1))) - (0.0 + 1e10)))
                        _ = v56
                    }
                case "a":
                    v57 := v1["hello"]
                    var v58 int = (v13 - 127)
                    switch "foo bar" {
                    case "hello":
                        if true {
                            v13 -= v57
                            fmt.Println(v13)
                            var v59 float64 = 1.5
                            _ = v59
                        }
                    case "":
                        v60 := GenS3{F0: 853, F1: 65535}
                        _ = v60
                    }
                    var v61 int64 = v10
                    if (!false) {
                        switch (v57 & fi46) {
                        case 42:
                            var v62 int16 = 0
                            v58 = fi46
                            _ = v62
                        case -2147483648:
                            fmt.Println(v6)
                        case -9223372036854775808:
                            if false {
                                v14 += v45
                                if true {
                                    v61 -= v10
                                } else {
                                    switch "hello" {
                                    case "a":
                                        fmt.Println(v57, v10, v58)
                                        _ = len(v2)
                                        v10 = v10
                                    case "":
                                        fmt.Println(v45, v10)
                                    case "\n\t":
                                        var v63 int32 = (v45 + v14)
                                        _ = v63
                                    }
                                    delete(v1, "a")
                                    if false {
                                        var v64 bool = ("foo bar" == ("" + "\n\t"))
                                        v5[72] = (((("a" + ((("foo bar" + "") + ("\n\t" + "a")) + (("\n\t" + "") + "foo bar"))) + "foo bar") + (("\n\t" + "\n\t") + "")) + ("hello" + ("a" + ("a" + ("\n\t" + ("foo bar" + ("hello" + "")))))))
                                        v13 = v58
                                        _ = v64
                                    }
                                }
                                v65 := GenS1{F0: (v61 >> ((443) & 63))}
                                switch "hello" {
                                case "hello":
                                    v58 = 32767
                                case "foo bar":
                                    var v66 int32 = 2147483647
                                    _ = v66
                                case "\n\t":
                                    var v67 float32 = 3.14159
                                    _ = v67
                                }
                                _ = v65
                            }
                        }
                    }
                    _ = v57
                    _ = v58
                    _ = v61
                default:
                    if false {
                        if ((3.14159 != float32(v6)) && (412 >= 42)) {
                            var v68 string = "\n\t"
                            _ = v68
                        }
                        v69 := GenS1{F0: (v10 + v10)}
                        switch -9223372036854775808 {
                        case 128:
                            var v70 uint32 = 9
                            _ = v70
                        case -532:
                            delete(v5, 408)
                        default:
                            switch (((("foo bar" + "\n\t") + "hello") + "foo bar") + "hello") {
                            case "hello":
                                v45 -= v14
                                if (721 >= 4294967295) {
                                    switch "a" {
                                    case "foo bar":
                                        v14 = v45
                                    default:
                                        v71 := GenS3{F0: 1, F1: 1}
                                        _ = v71
                                    }
                                    fmt.Println(v10, v10)
                                    var v72 uint8 = 0
                                    _ = v72
                                }
                            }
                        }
                        _ = v69
                    }
                }
                var v73 int = 4294967295
                _ = v47
                _ = v73
            }
            _ = fi46
        }
        _ = fi43
        _ = v44
        _ = v45
    }
    switch (v13 + 32767) {
    case 1:
        switch (v13 ^ 4294967295) {
        case 0:
            for fi74 := 0; fi74 < 10; fi74++ {
                fmt.Println(v14, v13)
                _ = fi74
            }
        default:
            v75 := GenS4{F0: 295}
            _ = v75
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
