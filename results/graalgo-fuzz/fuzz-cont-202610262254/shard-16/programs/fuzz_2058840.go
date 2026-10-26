package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int8; F2 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 uint8 }

func main() {
    var v1 uint16 = 523
    v2 := []string{}
    var v3 bool = true
    var v4 uint8 = 127
    var v5 uint8 = v4
    fmt.Println(v1, v1)
    v3 = (100 <= 100)
    if (v3 || v3) {
        var v6 int = 1
        _ = v6
    } else {
        var v7 float32 = float32(3.14159)
        switch -9223372036854775808 {
        case 128:
            if (!v3) {
                if len(v2) > 0 {
                    v2[0] = ("hello" + "")
                }
                var v8 uint64 = 381
                if ((0.0 + (-2.25 - (1.5 * 0.1))) != 3.14159) {
                    fmt.Println(v8)
                    if v3 {
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                        v3 = ((!v3) || (true && v3))
                        for fi9 := 0; fi9 < 6; fi9++ {
                            v10 := GenS1{F0: 10}
                            if len(v2) > 0 {
                                v2[0] = (("a" + "\n\t") + ("a" + "foo bar"))
                            }
                            var v11 int64 = 757
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                    } else {
                        if (0 == 2147483647) {
                            for fi12 := 0; fi12 < 18; fi12++ {
                                var v13 int8 = 42
                                if v3 {
                                    switch ("hello" + "\n\t") {
                                    case "hello":
                                        v14 := GenS2{F0: 65535}
                                        _ = v14
                                    default:
                                        if (false && ((v8 * 4294967295) >= 42)) {
                                            var v15 int64 = -565
                                            _ = v15
                                        }
                                    }
                                }
                                v16 := GenS3{F0: "foo bar", F1: (v13 % (((^(v13))) | 1)), F2: 255}
                                _ = fi12
                                _ = v13
                                _ = v16
                            }
                            if v3 {
                                if false {
                                    if (!true) {
                                        v17 := GenS2{F0: 4294967295}
                                        v18 := GenS3{F0: "a", F1: -100, F2: 18446744073709551615}
                                        var v19 int64 = 9223372036854775807
                                        _ = v17
                                        _ = v18
                                        _ = v19
                                    } else {
                                        v20 := GenS1{F0: 100}
                                        _ = v20
                                    }
                                } else {
                                    if (2 < 65535) {
                                        if ((1e10 - -2.25) > 1.5) {
                                            if (1 <= 509) {
                                                for fi21 := 0; fi21 < 11; fi21++ {
                                                    if (!true) {
                                                        var v22 int8 = -128
                                                        if v3 {
                                                            for fi23 := 0; fi23 < 11; fi23++ {
                                                                var v24 int = fi21
                                                                _ = fi23
                                                                _ = v24
                                                            }
                                                            v2 = append(v2, ("foo bar" + ("foo bar" + "hello")))
                                                            v2 = append(v2, "a")
                                                        } else {
                                                            var v25 uint64 = v8
                                                            v26 := GenS4{F0: -892}
                                                            var v27 int8 = (^(v22))
                                                            _ = v25
                                                            _ = v26
                                                            _ = v27
                                                        }
                                                        _ = v22
                                                    } else {
                                                        for fi28 := 0; fi28 < 15; fi28++ {
                                                            var v29 int = (fi28 - (fi21 % ((fi28) | 1)))
                                                            if (1 == 1) {
                                                                var v30 int = (v29 - (fi28 | fi28))
                                                                var v31 uint = 42
                                                                _ = v30
                                                                _ = v31
                                                            }
                                                            _ = fi28
                                                            _ = v29
                                                        }
                                                    }
                                                    fmt.Println(v1, v4)
                                                    _ = fi21
                                                }
                                                v3 = (764 <= 4294967295)
                                            } else {
                                                v3 = v3
                                                for fi32 := 0; fi32 < 4; fi32++ {
                                                    var v33 bool = (-9223372036854775808 < 42)
                                                    v5 = (v5 * v4)
                                                    _ = fi32
                                                    _ = v33
                                                }
                                            }
                                        }
                                        if false {
                                            switch "a" {
                                            case "hello":
                                                v1 -= (v1 % (((v1 >> ((714) & 63))) | 1))
                                            case "\n\t":
                                                v7 = ((1e10 + v7) - v7)
                                            case "":
                                                fmt.Println(v1, v1)
                                            }
                                        }
                                    } else {
                                        v3 = v3
                                    }
                                    var v34 bool = (-2147483648 < 1)
                                    v2 = append(v2, "")
                                    _ = v34
                                }
                                v35 := GenS2{F0: 2}
                                _ = v35
                            }
                            if v3 {
                                if v3 {
                                    v36 := GenS5{F0: "\n\t", F1: v4}
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if v3 {
                                        v4 = 127
                                        switch -534 {
                                        case 2147483647:
                                            fmt.Println(v7)
                                        case 0:
                                            for fi37 := 0; fi37 < 15; fi37++ {
                                                v1 = (^(v1))
                                                v2 = append(v2, ("a" + (((("a" + "foo bar") + ("hello" + "hello")) + "a") + "hello")))
                                                for fi38 := 0; fi38 < 3; fi38++ {
                                                    v5 = v5
                                                    _ = fi38
                                                }
                                                _ = fi37
                                            }
                                        case 495:
                                            if len(v2) > 0 {
                                                v2[0] = "a"
                                            }
                                        default:
                                            _ = len(v2)
                                        }
                                        v2 = append(v2, "a")
                                    } else {
                                        var v39 uint32 = 366
                                        if (v5 <= 1) {
                                            v40 := GenS2{F0: uint32(v5)}
                                            _ = v40
                                        } else {
                                            for fi41 := 0; fi41 < 7; fi41++ {
                                                v39 = 42
                                                _ = fi41
                                            }
                                            v1 = v1
                                            var v42 int16 = 1000
                                            _ = v42
                                        }
                                        _ = v39
                                    }
                                    _ = v36
                                }
                                v7 = 3.14159
                                for fi43 := 0; fi43 < 2; fi43++ {
                                    for fi44 := 0; fi44 < 13; fi44++ {
                                        v7 += (v7 - v7)
                                        _ = fi44
                                    }
                                    v45 := GenS4{F0: 57}
                                    var v46 int16 = 394
                                    _ = fi43
                                    _ = v45
                                    _ = v46
                                }
                            }
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    var v47 uint = 757
                    _ = v47
                }
                _ = v8
            }
        case 397:
            for fi48 := 0; fi48 < 1; fi48++ {
                v7 += 1e10
                _ = len(v2)
                _ = fi48
            }
        }
        fmt.Println(v4, v3)
        _ = v7
    }
    for fi49 := 0; fi49 < 13; fi49++ {
        switch 896 {
        case 32768:
            switch fi49 {
            case -546:
                if true {
                    fmt.Println(fi49, fi49, v5)
                } else {
                    switch (fi49 >> ((255) & 63)) {
                    case 583:
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                    case 128:
                        v3 = (!v3)
                    case 127:
                        v3 = ((fi49 | fi49) != (fi49 & (fi49 ^ (fi49 % (((fi49 + -536)) | 1)))))
                    }
                }
            case 238:
                v4 = (v4 ^ (v5 / ((v5) | 1)))
            }
        case 9223372036854775807:
            switch (fi49 * -128) {
            case 2147483647:
                var v50 float32 = 3.14159
                _ = v50
            case 0:
                if (v1 > v1) {
                    var v51 bool = (false && (!(("hello" + "\n\t") > ("hello" + "foo bar"))))
                    v52 := GenS3{F0: "hello", F1: -124, F2: 0}
                    if (!(3.14159 == 0.1)) {
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                        for fi53 := 0; fi53 < 2; fi53++ {
                            v1 = v1
                            _ = fi53
                        }
                        var v54 uint64 = 4294967295
                        _ = v54
                    }
                    _ = v51
                    _ = v52
                } else {
                    v3 = v3
                }
            case -1:
                var v55 int = (fi49 | fi49)
                _ = v55
            default:
                v3 = v3
            }
        case 0:
            var v56 int8 = 127
            _ = v56
        }
        v5 -= (v5 % ((80) | 1))
        _ = fi49
    }
    var v57 uint = 18446744073709551615
    if true {
        v2 = append(v2, "foo bar")
    } else {
        v5 += v4
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi58 := 0; fi58 < 19; fi58++ {
            if ("hello" <= "") {
                v2 = append(v2, "")
            }
            v57 -= v57
            if ((0.0 == 3.14159) || ((fi58 ^ fi58) < 9223372036854775807)) {
                var v59 uint = v57
                _ = v59
            } else {
                var v60 int64 = 1
                if (v60 > v60) {
                    fmt.Println(v60, v4, v1)
                    switch ("hello" + "") {
                    case "a":
                        v61 := GenS3{F0: "a", F1: 100, F2: (^(v57))}
                        _ = v61
                    case "":
                        v2 = append(v2, (("hello" + "") + ""))
                    case "\n\t":
                        v1 = (v1 + v1)
                    }
                    switch "\n\t" {
                    case "a":
                        switch fi58 {
                        case 128:
                            if len(v2) > 0 {
                                v2[0] = ("hello" + "")
                            }
                        case 660:
                            if v3 {
                                fmt.Println(v5)
                            } else {
                                var v62 int64 = v60
                                _ = v62
                            }
                        case 2147483647:
                            var v63 uint16 = v1
                            _ = v63
                        default:
                            v57 = v57
                        }
                    default:
                        v57 -= v57
                    }
                }
                _ = v60
            }
            _ = fi58
        }
    }
    v2 = append(v2, "")
    var v64 string = ("foo bar" + "hello")
    v57 -= (^(v57))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v64
}
