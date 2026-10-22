package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint64, p1 int64, p2 uint) int {
    p2 -= (p2 - uint(p0))
    fmt.Println(p1, p2)
    var v7 uint = 675
    v8 := GenS1{F0: 255}
    v9 := GenS1{F0: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 32767
}
func genFn1(p0 int, p1 string, p2 uint16) uint64 {
    v10 := GenS1{F0: 255}
    var v11 uint64 = 0
    p0 = (p0 | p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return (^(v11))
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint = 65535
    var v3 float64 = 0.0
    var v4 int16 = -32768
    v5 := make(map[string]bool)
    v6 := make(map[string]string)
    var v12 uint64 = 920
    v1 -= v3
    if false {
        var v13 uint32 = 2
        var v14 int = genFn0(v12, 1, 545)
        v15 := v5[""]
        _ = v13
        _ = v14
        _ = v15
    } else {
        var v16 int = -2147483648
        if false {
            for fi17 := 0; fi17 < 8; fi17++ {
                switch ("" + "foo bar") {
                case "foo bar":
                    var v18 int32 = 1
                    _ = v18
                case "\n\t":
                    for fi19 := 0; fi19 < 6; fi19++ {
                        switch "hello" {
                        case "":
                            if (false || (false || (("hello" + "\n\t") != (("\n\t" + (("\n\t" + "foo bar") + "hello")) + ("hello" + (("a" + "a") + (("" + "hello") + "a"))))))) {
                                v12 -= 262
                                var v20 bool = (v4 != (v4 - 0))
                                _ = v20
                            }
                        default:
                            switch (v16 - fi19) {
                            case 128:
                                v21 := GenS1{F0: 0}
                                _ = v21
                            default:
                                switch genFn0(v12, 42, (^(v2))) {
                                case -747:
                                    fmt.Println(v16, fi19)
                                case 9223372036854775807:
                                    if false {
                                        switch (fi19 - 1) {
                                        case 566:
                                            v16 += v16
                                        case -268:
                                            switch "a" {
                                            case "a":
                                                v16 -= 42
                                            case "foo bar":
                                                v1 = v1
                                            }
                                        case 32767:
                                            v4 = (v4 % ((v4) | 1))
                                        }
                                        v22 := GenS2{F0: 746}
                                        if (true || false) {
                                            switch (fi17 - 0) {
                                            case -2147483648:
                                                v23 := v5[""]
                                                _ = v23
                                            case 80:
                                                v2 = (v2 >> (((v2 >> ((v2) & 63))) & 63))
                                            default:
                                                v24 := GenS3{F0: 42}
                                                _ = v24
                                            }
                                            switch v16 {
                                            case -9223372036854775808:
                                                fmt.Println(v2, v12, v4)
                                            default:
                                                v25 := GenS1{F0: 65535}
                                                _ = v25
                                            }
                                            var v26 string = "a"
                                            _ = v26
                                        } else {
                                            v27 := []int{-531}
                                            _ = v27
                                        }
                                        _ = v22
                                    }
                                case 42:
                                    fmt.Println(v1, fi19)
                                }
                            }
                        }
                        var v28 uint64 = v12
                        v29 := []int{fi19}
                        _ = fi19
                        _ = v28
                        _ = v29
                    }
                default:
                    switch ("" + "") {
                    case "hello":
                        for fi30 := 0; fi30 < 20; fi30++ {
                            if true {
                                fmt.Println(v2, fi30)
                                v31 := GenS1{F0: 418}
                                _ = v31
                            }
                            delete(v6, "a")
                            if ("foo bar" != "") {
                                var v32 uint16 = 65535
                                _ = v32
                            }
                            var v33 float32 = 1e10
                            switch ("foo bar" + "a") {
                            case "\n\t":
                                if true {
                                    var v34 int32 = -2147483648
                                    v2 = 18446744073709551615
                                    _ = v34
                                } else {
                                    v1 += v1
                                    v6["foo bar"] = ""
                                    v3 += v3
                                    switch fi30 {
                                    case -2147483648:
                                        if false {
                                            var v35 uint32 = 42
                                            fmt.Println(v33, v4, fi30)
                                            v36 := GenS1{F0: 827}
                                            _ = v35
                                            _ = v36
                                        }
                                    case 32767:
                                        var v37 float32 = v33
                                        _ = v37
                                    case 42:
                                        fmt.Println(v16)
                                    }
                                    switch int(v2) {
                                    case 42:
                                        v33 += float32(3.14159)
                                    default:
                                        var v38 string = "\n\t"
                                        var v39 uint32 = 919
                                        v4 = (v4 & v4)
                                        if true {
                                            var v40 int = (fi30 ^ v16)
                                            _ = v40
                                        } else {
                                            var v41 int = 1
                                            fmt.Println(fi30)
                                            v16 = (v41 * (fi30 - (v16 ^ fi30)))
                                            _ = v41
                                        }
                                        _ = v38
                                        _ = v39
                                    }
                                }
                            case "":
                                v42 := GenS3{F0: -9223372036854775808}
                                _ = v42
                            }
                            _ = fi30
                            _ = v33
                        }
                    case "foo bar":
                        if true {
                            switch "foo bar" {
                            case "":
                                if true {
                                    var v43 uint16 = 839
                                    v1 = -2.25
                                    v44 := []int{genFn0(v12, 477, v2), 438, fi17, int(v2)}
                                    _ = v43
                                    _ = v44
                                }
                            case "hello":
                                for fi45 := 0; fi45 < 3; fi45++ {
                                    v12 = genFn1(fi17, ((("hello" + "a") + "hello") + "a"), 255)
                                    var v46 uint16 = 873
                                    _ = fi45
                                    _ = v46
                                }
                            }
                        } else {
                            v47 := GenS2{F0: genFn0(v12, 707, (^(v2)))}
                            switch (v16 ^ 2147483647) {
                            case 128:
                                v48 := v5["\n\t"]
                                _ = v48
                            case 259:
                                if (1 <= -835) {
                                    v12 = (v12 ^ 529)
                                    v4 = v4
                                }
                            case 1:
                                v49 := v6["foo bar"]
                                _ = v49
                            }
                            _ = v47
                        }
                    default:
                        var v50 string = "foo bar"
                        _ = v50
                    }
                }
                _ = fi17
            }
            var v51 float64 = 3.14159
            fmt.Println(v3, v51, v3)
            _ = v51
        }
        _ = v16
    }
    switch ("hello" + "\n\t") {
    case "foo bar":
        v3 = ((float64(float32(float64(0.1))) + (float64((-2.25 - 1.5)) + v3)) / (v1 - 0.1))
    case "\n\t":
        var v52 uint8 = 97
        _ = v52
    case "a":
        var v53 uint8 = 2
        _ = v53
    }
    v3 = 1.5
    v54 := GenS3{F0: 0}
    v55 := GenS4{F0: (v2 & (v2 | (v2 | 1))), F1: (v2 % (((^(v2))) | 1))}
    var v56 float32 = -2.25
    var v57 uint = v2
    for fi58 := 0; fi58 < 11; fi58++ {
        if true {
            fmt.Println(v1)
            var v59 uint = v57
            _ = v59
        } else {
            switch (fi58 ^ fi58) {
            case 4294967295:
                var v60 float32 = v56
                _ = v60
            case 32767:
                if false {
                    v2 += 0
                    v12 += genFn1(fi58, ("foo bar" + ("hello" + ("\n\t" + "foo bar"))), uint16(v57))
                } else {
                    for fi61 := 0; fi61 < 15; fi61++ {
                        if ((^(v12)) != v12) {
                            v3 = v3
                            if false {
                                fmt.Println(v56, fi58)
                            } else {
                                v56 -= (float32(v1) * 0.1)
                                v62 := GenS4{F0: (v2 | (v57 >> ((610) & 63))), F1: (v57 ^ (v2 / (((v2 | v57)) | 1)))}
                                _ = v62
                            }
                            var v63 uint16 = 32767
                            var v64 uint32 = 0
                            fmt.Println(v1, v56)
                            if false {
                                _ = len(v6)
                                var v65 int8 = -127
                                v56 = (1e10 - -2.25)
                                _ = v65
                            } else {
                                if true {
                                    v66 := []int{fi58}
                                    v2 += (^(v2))
                                    v4 = v4
                                    _ = v66
                                }
                                if ((^(v64)) >= v64) {
                                    v1 = (v3 * -2.25)
                                    v56 = v56
                                } else {
                                    v67 := GenS1{F0: (v63 % ((42) | 1))}
                                    switch fi58 {
                                    case -9223372036854775808:
                                        var v68 int16 = v4
                                        _ = v68
                                    case -363:
                                        switch "a" {
                                        case "":
                                            if ((!("foo bar" <= "foo bar")) && ("" == ("hello" + "a"))) {
                                                var v69 bool = ((("hello" + "") == "foo bar") && false)
                                                var v70 uint16 = v63
                                                _ = v69
                                                _ = v70
                                            }
                                            v4 = v4
                                        }
                                    }
                                    _ = v67
                                }
                                var v71 int8 = 42
                                _ = v71
                            }
                            _ = v63
                            _ = v64
                        }
                        var v72 uint8 = 127
                        v73 := GenS5{F0: 1}
                        _ = fi61
                        _ = v72
                        _ = v73
                    }
                }
            case -2147483648:
                var v74 uint64 = 449
                _ = v74
            default:
                switch fi58 {
                case -1:
                    if (42 < 65535) {
                        v3 += v1
                        var v75 uint = (v57 | (v57 & v2))
                        _ = v75
                    } else {
                        v76 := GenS2{F0: -9223372036854775808}
                        _ = v76
                    }
                case -2147483648:
                    fmt.Println(v1, v1, v3)
                }
            }
        }
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v3, v1, v4)
        case "hello":
            v1 = 3.14159
        }
        for fi77 := 0; fi77 < 8; fi77++ {
            v56 -= (1e10 / (((v56 - ((float32(0.1) * v56) / v56)) + v56) * (v56 * v56)))
            v78 := GenS4{F0: (^(v2)), F1: v2}
            fmt.Println(v56)
            _ = fi77
            _ = v78
        }
        _ = fi58
    }
    var v79 uint = (v2 / (((v57 ^ (v57 / ((v2) | 1)))) | 1))
    v79 = v79
    v79 -= (v79 / (((v79 << ((v2) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v79:", v79)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v6
    _ = v79
}
