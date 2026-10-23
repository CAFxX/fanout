package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := []uint32{42}
    v2 := make(map[int]uint32)
    v3 := []int{-2147483648, 32767}
    var v4 string = "hello"
    var v5 int64 = -458
    if len(v3) > 0 {
        v3[0] = 42
    }
    fmt.Println(v5, v5)
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int16 = 634
        _ = fi6
        _ = v7
    }
    delete(v2, 0)
    _ = len(v3)
    v8 := v2[2147483647]
    fmt.Println(v4, v8)
    switch 879 {
    case -1:
        var v9 bool = true
        _ = v9
    case -896:
        if (-32768 < 32767) {
            if true {
                v5 = 1
                _ = len(v1)
                v5 -= v5
            }
            v10 := GenS1{F0: 42, F1: -1}
            _ = v10
        } else {
            v4 = "a"
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 int = (^(fi11))
                _ = fi11
                _ = v12
            }
            v4 = "\n\t"
        }
    case -767:
        var v13 bool = (1 <= 2)
        _ = v13
    default:
        fmt.Println(v5, v5, v8)
    }
    switch ("foo bar" + v4) {
    case "foo bar":
        v14 := GenS2{F0: true, F1: 32767}
        _ = v14
    case "hello":
        var v15 uint8 = 208
        _ = v15
    default:
        v16 := GenS1{F0: 2, F1: 245}
        _ = v16
    }
    if ((uint64(v8) <= 999) && true) {
        for fi17 := 0; fi17 < 6; fi17++ {
            if (!((false && (false || true)) && false)) {
                var v18 int16 = 808
                var v19 float32 = ((0.0 + 1.5) - 1e10)
                _ = v18
                _ = v19
            }
            fmt.Println(v4, fi17)
            v20 := GenS1{F0: 255, F1: 997}
            _ = fi17
            _ = v20
        }
        v4 = (((v4 + "a") + v4) + "foo bar")
    } else {
        v5 = -1
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        v8 += 566
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        v4 = (v4 + v4)
        if false {
            switch fi22 {
            case -1:
                if len(v1) > 0 {
                    v1[0] = v8
                }
            }
            switch v4 {
            case "hello":
                switch v4 {
                case "\n\t":
                    v5 -= v5
                case "":
                    v23 := GenS1{F0: 2, F1: 32767}
                    _ = v23
                case "a":
                    for fi24 := 0; fi24 < 4; fi24++ {
                        var v25 string = "foo bar"
                        switch "foo bar" {
                        case "foo bar":
                            v5 = v5
                        case "hello":
                            if len(v3) > 0 {
                                v3[0] = (fi24 ^ fi24)
                            }
                        case "a":
                            if true {
                                switch (v25 + ((v25 + (((v4 + (v25 + ("\n\t" + "hello"))) + v4) + ("" + v25))) + ((((v25 + (v25 + "foo bar")) + ((v4 + "a") + "hello")) + v25) + "a"))) {
                                case "\n\t":
                                    var v26 int = 32768
                                    _ = v26
                                }
                                if (false || true) {
                                    if true {
                                        if false {
                                            var v27 uint8 = 0
                                            _ = v27
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = 250
                                            }
                                            if true {
                                                var v28 bool = ((0.1 + 0.1) < 1.5)
                                                var v29 int16 = -32768
                                                v5 = -9223372036854775808
                                                _ = v28
                                                _ = v29
                                            }
                                            var v30 uint = 255
                                            v8 -= (v8 % ((v8) | 1))
                                            var v31 bool = false
                                            v1 = append(v1, v8)
                                            var v32 int64 = (v5 * v5)
                                            if v31 {
                                                var v33 uint = 253
                                                if ((!v31) && true) {
                                                    switch (fi22 % (((fi24 >> ((v33) & 63))) | 1)) {
                                                    case 128:
                                                        switch v25 {
                                                        case "":
                                                            v34 := GenS2{F0: v31, F1: 32767}
                                                            _ = v34
                                                        case "foo bar":
                                                            var v35 uint8 = 42
                                                            _ = v35
                                                        }
                                                    case 2147483647:
                                                        if false {
                                                            switch v4 {
                                                            case "foo bar":
                                                                if ((v4 == "\n\t") || true) {
                                                                    var v36 uint = (v30 | 152)
                                                                    if len(v1) > 0 {
                                                                        v1[0] = (^(v8))
                                                                    }
                                                                    _ = v36
                                                                }
                                                            }
                                                        } else {
                                                            v4 = (("\n\t" + (v25 + (v25 + (v4 + v25)))) + v4)
                                                        }
                                                    }
                                                    fmt.Println(v31)
                                                }
                                                v4 = ("foo bar" + v25)
                                                _ = v33
                                            } else {
                                                var v37 string = v25
                                                _ = v37
                                            }
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                        }
                                        if (-2147483648 <= 690) {
                                            var v38 int = (fi24 & 522)
                                            _ = v38
                                        }
                                        var v39 int8 = 127
                                        _ = v39
                                    }
                                    if true {
                                        v25 = (("a" + v4) + v4)
                                        var v40 uint = 42
                                        _ = v40
                                    } else {
                                        var v41 int = fi24
                                        switch ("\n\t" + v25) {
                                        case "hello":
                                            var v42 int = (fi22 / (((fi24 + fi24)) | 1))
                                            _ = v42
                                        case "foo bar":
                                            v41 += fi24
                                            var v43 int8 = 127
                                            switch fi24 {
                                            case -1:
                                                if false {
                                                    if false {
                                                        var v44 float64 = 3.14159
                                                        v5 = 9223372036854775807
                                                        v45 := GenS2{F0: false, F1: 1}
                                                        var v46 float64 = 0.0
                                                        v47 := GenS2{F0: (float64(-2.25) <= (v46 * float64(0.1))), F1: 560}
                                                        _ = v44
                                                        _ = v45
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    if (true || (v4 < v4)) {
                                                        delete(v2, (fi24 >> ((1) & 63)))
                                                    }
                                                }
                                            case 32767:
                                                v8 = (^(v8))
                                            case 127:
                                                if (2 < 1) {
                                                    if true {
                                                        v48 := GenS1{F0: 127, F1: 357}
                                                        v1 = append(v1, (v8 ^ v8))
                                                        _ = v48
                                                    }
                                                } else {
                                                    v5 -= (v5 * v5)
                                                }
                                            default:
                                                v49 := GenS3{F0: 379, F1: 4294967295}
                                                _ = v49
                                            }
                                            v50 := GenS3{F0: 153, F1: 966}
                                            _ = v43
                                            _ = v50
                                        default:
                                            var v51 int32 = -1
                                            _ = v51
                                        }
                                        if false {
                                            v41 -= fi24
                                            var v52 bool = (!false)
                                            if (v52 && (!true)) {
                                                v5 += (v5 % ((-1) | 1))
                                                var v53 int32 = 2147483647
                                                _ = len(v2)
                                                _ = v53
                                            } else {
                                                v41 = (v41 << ((65535) & 63))
                                                v5 += (v5 + v5)
                                            }
                                            _ = v52
                                        } else {
                                            fmt.Println(fi22, v25, fi24)
                                            _ = len(v1)
                                            v8 -= (v8 >> ((808) & 63))
                                        }
                                        _ = v41
                                    }
                                }
                            }
                        }
                        v5 -= v5
                        _ = fi24
                        _ = v25
                    }
                }
            case "a":
                v54 := GenS2{F0: (uint8(v8) <= 127), F1: 1}
                _ = v54
            default:
                var v55 uint16 = 32767
                _ = v55
            }
        } else {
            for fi56 := 0; fi56 < 20; fi56++ {
                if true {
                    var v57 uint32 = v8
                    if true {
                        v8 = (v8 | v57)
                        if ("" <= v4) {
                            var v58 uint16 = 2
                            var v59 int8 = 1
                            _ = v58
                            _ = v59
                        } else {
                            v5 -= v5
                            v4 = (((v4 + v4) + v4) + v4)
                            var v60 int8 = -106
                            _ = v60
                        }
                    }
                    _ = v57
                } else {
                    var v61 float32 = 0.1
                    switch v4 {
                    case "foo bar":
                        v8 -= 2147483647
                    default:
                        _ = len(v2)
                    }
                    v61 = float32(3.14159)
                    _ = v61
                }
                if true {
                    v62 := GenS3{F0: 25, F1: 42}
                    _ = v62
                }
                _ = fi56
            }
            fmt.Println(v5, v4)
            var v63 string = v4
            _ = v63
        }
        _ = fi22
    }
    var v64 float32 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v8
}
