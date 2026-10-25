package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 uint64; F2 uint32 }
func genFn0() int32 {
    var v4 int16 = -1
    if (true || true) {
        fmt.Println(v4)
    }
    v4 = v4
    v4 = (v4 & 32767)
    _ = v4
    return -177
}
func genFn1() int32 {
    switch "hello" {
    case "\n\t":
        v5 := GenS1{F0: "", F1: -1, F2: 622}
        _ = v5
    case "hello":
        var v6 int64 = 1
        _ = v6
    default:
        var v7 float64 = 1e10
        _ = v7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        if (!false) {
            if (127 >= 0) {
                for fi9 := 0; fi9 < 18; fi9++ {
                    v10 := []int{(-(fi9)), (fi9 % ((fi9) | 1))}
                    _ = fi9
                    _ = v10
                }
                for fi11 := 0; fi11 < 18; fi11++ {
                    var v12 bool = false
                    v12 = v12
                    v12 = true
                    _ = fi11
                    _ = v12
                }
            }
        } else {
            var v13 int = (fi8 / (((fi8 & fi8)) | 1))
            _ = v13
        }
        v14 := GenS2{F0: 820, F1: 2147483647, F2: 18446744073709551615}
        _ = fi8
        _ = v14
    }
    var v15 int = -656
    _ = v15
    return -2147483648
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{"\n\t"}
    v3 := make(map[string]bool)
    switch (("" + "\n\t") + "foo bar") {
    case "foo bar":
        var v16 uint32 = 4294967295
        _ = v16
    case "\n\t":
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int32 = genFn1()
            fmt.Println(fi17, v18)
            _ = fi17
            _ = v18
        }
    case "":
        for fi19 := 0; fi19 < 17; fi19++ {
            v20 := GenS2{F0: 161, F1: 0, F2: 789}
            if (-996 < 1) {
                var v21 uint32 = 65535
                var v22 bool = (false && true)
                v21 -= 2
                _ = v21
                _ = v22
            } else {
                fmt.Println(fi19)
                if (!(617 < 2)) {
                    if ((true || (!false)) || false) {
                        if false {
                            if false {
                                switch ("foo bar" + "\n\t") {
                                case "hello":
                                    var v23 uint16 = 0
                                    _ = v23
                                default:
                                    switch (fi19 % (((-(fi19))) | 1)) {
                                    case -109:
                                        for fi24 := 0; fi24 < 13; fi24++ {
                                            var v25 uint16 = 154
                                            v26 := GenS3{F0: 13}
                                            var v27 string = ""
                                            switch v27 {
                                            case "a":
                                                if (false && (false || true)) {
                                                    var v28 int32 = -932
                                                    switch -756 {
                                                    case 32767:
                                                        v28 = (^(v28))
                                                    case 42:
                                                        v29 := GenS1{F0: ("hello" + "a"), F1: 42, F2: v25}
                                                        _ = v29
                                                    default:
                                                        if (false && (1 < -106)) {
                                                            switch (fi19 + fi19) {
                                                            case 640:
                                                                v30 := GenS3{F0: -73}
                                                                _ = v30
                                                            }
                                                        } else {
                                                            if false {
                                                                var v31 uint64 = 42
                                                                var v32 int = 450
                                                                if (!false) {
                                                                    var v33 uint16 = (v25 % (((v25 / ((v25) | 1))) | 1))
                                                                    _ = v33
                                                                } else {
                                                                    v28 -= genFn1()
                                                                }
                                                                _ = v31
                                                                _ = v32
                                                            } else {
                                                                var v34 int64 = -1
                                                                fmt.Println(v34, v28, v34)
                                                                _ = len(v2)
                                                                _ = v34
                                                            }
                                                            var v35 int16 = 473
                                                            var v36 float32 = (3.14159 * 0.1)
                                                            switch fi24 {
                                                            case 1:
                                                                var v37 uint64 = 2
                                                                _ = v37
                                                            default:
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                            if (18446744073709551615 >= 1) {
                                                                fmt.Println(v25)
                                                                if true {
                                                                    if (42 < 65535) {
                                                                        fmt.Println(v28)
                                                                        var v38 bool = ((v27 + "foo bar") <= "\n\t")
                                                                        _ = v38
                                                                    } else {
                                                                        v28 += -1
                                                                    }
                                                                }
                                                                v39 := GenS4{F0: 9223372036854775807, F1: (!(18446744073709551615 == 2))}
                                                                _ = v39
                                                            } else {
                                                                var v40 float64 = ((float64(v36) / (0.1 - 0.0)) / (3.14159 + -2.25))
                                                                v41 := v3[v27]
                                                                var v42 uint8 = 2
                                                                if v41 {
                                                                    fmt.Println(fi24, v40, fi24)
                                                                    switch 1 {
                                                                    case 0:
                                                                        var v43 int8 = 0
                                                                        _ = v43
                                                                    default:
                                                                        var v44 int = fi24
                                                                        _ = v44
                                                                    }
                                                                }
                                                                _ = v40
                                                                _ = v41
                                                                _ = v42
                                                            }
                                                            v45 := GenS5{F0: (128 > 255)}
                                                            v25 = v25
                                                            _ = v35
                                                            _ = v36
                                                            _ = v45
                                                        }
                                                        var v46 float32 = (float32((1e10 / 0.1)) - 1e10)
                                                        _ = v46
                                                    }
                                                    v28 -= (^(v28))
                                                    _ = v28
                                                }
                                                v47 := GenS1{F0: "hello", F1: -128, F2: 65535}
                                                var v48 uint32 = 65535
                                                _ = v47
                                                _ = v48
                                            case "":
                                                v25 = (^(v25))
                                            case "\n\t":
                                                fmt.Println(fi24)
                                            }
                                            _ = fi24
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                        }
                                    case -2147483648:
                                        delete(v1, "a")
                                    default:
                                        var v49 uint64 = 497
                                        _ = v49
                                    }
                                }
                            }
                            v50 := GenS1{F0: (("hello" + "foo bar") + (("hello" + "\n\t") + (("foo bar" + "hello") + "foo bar"))), F1: -52, F2: 42}
                            for fi51 := 0; fi51 < 6; fi51++ {
                                if false {
                                    switch "hello" {
                                    case "foo bar":
                                        var v52 string = ("" + "a")
                                        _ = v52
                                    default:
                                        var v53 int8 = 127
                                        _ = v53
                                    }
                                } else {
                                    switch ("" + "foo bar") {
                                    case "":
                                        var v54 uint16 = 42
                                        _ = v54
                                    case "foo bar":
                                        if false {
                                            if (((!true) || (true && true)) || true) {
                                                var v55 bool = (false && true)
                                                _ = v55
                                            } else {
                                                var v56 uint = 42
                                                _ = v56
                                            }
                                        }
                                    default:
                                        fmt.Println(fi51)
                                    }
                                    var v57 float32 = (-2.25 / 0.1)
                                    _ = v57
                                }
                                _ = fi51
                            }
                            _ = v50
                        }
                    }
                    var v58 int = fi19
                    switch "a" {
                    case "foo bar":
                        v58 = (fi19 - fi19)
                    }
                    _ = v58
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi19
            _ = v20
        }
    default:
        if (!(255 > 255)) {
            var v59 uint64 = 1
            _ = v59
        } else {
            var v60 float64 = (0.1 + -2.25)
            _ = v60
        }
    }
    if true {
        delete(v1, ("foo bar" + "foo bar"))
        v61 := GenS5{F0: true}
        var v62 string = "foo bar"
        _ = v61
        _ = v62
    }
    switch ("a" + ("" + "\n\t")) {
    case "foo bar":
        var v63 int8 = -89
        _ = v63
    default:
        var v64 uint8 = 127
        _ = v64
    }
    v65 := GenS6{F0: true, F1: 2, F2: 174}
    var v66 uint16 = 65535
    for fi67 := 0; fi67 < 3; fi67++ {
        fmt.Println(v66, v66)
        var v68 uint16 = 32767
        _ = fi67
        _ = v68
    }
    _ = len(v3)
    switch ("hello" + ("foo bar" + "a")) {
    case "foo bar":
        v66 -= 42
    }
    fmt.Println(v66)
    v66 -= v66
    v66 -= 822
    v69 := GenS1{F0: "a", F1: -128, F2: v66}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v65
    _ = v66
    _ = v69
}
