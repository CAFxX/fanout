package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint16 }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 4; fi6++ {
        switch "foo bar" {
        case "":
            v7 := GenS1{F0: 255, F1: "foo bar", F2: 2}
            _ = v7
        case "a":
            for fi8 := 0; fi8 < 16; fi8++ {
                var v9 uint32 = 42
                switch ("foo bar" + "hello") {
                case "\n\t":
                    v10 := GenS1{F0: uint16(v9), F1: "hello", F2: 2}
                    _ = v10
                case "hello":
                    if ((1.5 * -2.25) > 0.1) {
                        var v11 uint32 = 65535
                        var v12 float32 = (float32(0.0) / (1.5 + 1e10))
                        _ = v11
                        _ = v12
                    } else {
                        fmt.Println(fi6, fi6)
                    }
                }
                _ = fi8
                _ = v9
            }
        case "foo bar":
            for fi13 := 0; fi13 < 14; fi13++ {
                v14 := GenS2{F0: 65535, F1: 42}
                fmt.Println(fi6, fi13)
                _ = fi13
                _ = v14
            }
        default:
            fmt.Println(fi6)
        }
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 uint32 = 993
            v16 = 2147483648
            _ = fi15
            _ = v16
        }
        switch fi6 {
        case -1:
            v17 := []int{fi6, (^(fi6)), 128, fi6}
            _ = v17
        case 4294967295:
            v18 := GenS3{F0: (("foo bar" + "a") + "")}
            _ = v18
        case 2147483647:
            if false {
                var v19 string = "hello"
                v19 = "\n\t"
                _ = v19
            } else {
                v20 := []int{(fi6 % ((32767) | 1))}
                var v21 float32 = 0.0
                _ = v20
                _ = v21
            }
        default:
            var v22 int8 = 127
            _ = v22
        }
        _ = fi6
    }
    var v23 int64 = 480
    _ = v23
    return 255
}
func genFn1(p0 int, p1 uint8, p2 int32) int8 {
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 int8 = -117
        p0 = p0
        var v26 int = 2147483647
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 uint16 = 181
    var v28 uint = 65535
    for fi29 := 0; fi29 < 18; fi29++ {
        switch p0 {
        case -417:
            switch p0 {
            case -982:
                var v30 float64 = (-2.25 / 3.14159)
                _ = v30
            }
        case 42:
            var v31 uint64 = uint64(v28)
            _ = v31
        case -1:
            var v32 uint8 = p1
            _ = v32
        default:
            p1 += genFn0()
        }
        if (!((p1 == 127) || ((1.5 + 3.14159) > 1e10))) {
            switch "\n\t" {
            case "":
                for fi33 := 0; fi33 < 5; fi33++ {
                    switch (fi29 / ((42) | 1)) {
                    case -1:
                        var v34 uint32 = 65535
                        _ = v34
                    case 2147483647:
                        p2 = p2
                    }
                    if true {
                        v28 += uint(v27)
                        fmt.Println(p0, p0, v27)
                    } else {
                        p1 -= p1
                        if (1.5 <= float64(3.14159)) {
                            var v35 int16 = 32767
                            _ = v35
                        }
                    }
                    if true {
                        var v36 float64 = 3.14159
                        switch "hello" {
                        case "":
                            var v37 int32 = p2
                            _ = v37
                        case "\n\t":
                            if true {
                                v38 := GenS2{F0: 222, F1: 1}
                                v39 := []int{618, (p0 / (((p0 << ((1) & 63))) | 1))}
                                _ = v38
                                _ = v39
                            } else {
                                v40 := []int{}
                                p1 -= genFn0()
                                var v41 int64 = -9223372036854775808
                                _ = v40
                                _ = v41
                            }
                        }
                        var v42 uint64 = 2
                        _ = v36
                        _ = v42
                    }
                    var v43 int16 = -1
                    _ = fi33
                    _ = v43
                }
            case "hello":
                var v44 int32 = (p2 << ((v28) & 63))
                _ = v44
            case "\n\t":
                for fi45 := 0; fi45 < 10; fi45++ {
                    switch "" {
                    case "\n\t":
                        var v46 uint = v28
                        _ = v46
                    default:
                        var v47 uint8 = genFn0()
                        _ = v47
                    }
                    var v48 float32 = (1e10 * 0.0)
                    _ = fi45
                    _ = v48
                }
            default:
                fmt.Println(p0)
            }
            p2 -= p2
        }
        p1 = 1
        _ = fi29
    }
    for fi49 := 0; fi49 < 4; fi49++ {
        var v50 uint32 = 4294967295
        var v51 string = ("a" + ("" + "\n\t"))
        _ = fi49
        _ = v50
        _ = v51
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v28
    return -128
}

func main() {
    var v1 string = (("a" + "a") + (("foo bar" + "a") + "\n\t"))
    v2 := []int{4294967295, 42}
    v3 := make(map[string]string)
    var v4 uint = 255
    var v5 uint64 = 375
    _ = len(v2)
    v5 = (v5 & v5)
    if true {
        if len(v2) > 0 {
            v2[0] = 0
        }
        var v52 string = ("\n\t" + v1)
        _ = v52
    } else {
        v53 := v3[(v1 + "a")]
        v1 = ("a" + v53)
        fmt.Println(v5)
        _ = v53
    }
    switch (v1 + "foo bar") {
    case "hello":
        if len(v2) > 0 {
            v2[0] = -2147483648
        }
    default:
        var v54 int32 = 1
        _ = v54
    }
    v55 := GenS1{F0: 908, F1: ("hello" + v1), F2: v5}
    v56 := GenS2{F0: 42, F1: 787}
    delete(v3, (("\n\t" + "a") + ((v1 + "") + v1)))
    v57 := GenS4{F0: 0, F1: false}
    var v58 uint64 = 340
    v1 = ("foo bar" + v1)
    var v59 uint16 = 138
    v59 += uint16(v58)
    var v60 string = (v1 + (("a" + "") + ("foo bar" + "")))
    switch "\n\t" {
    case "\n\t":
        for fi61 := 0; fi61 < 9; fi61++ {
            v62 := GenS3{F0: "a"}
            _ = fi61
            _ = v62
        }
    case "hello":
        var v63 uint64 = v58
        _ = v63
    }
    v64 := v3[""]
    if true {
        if (!false) {
            if false {
                fmt.Println(v64)
                v60 = v1
            }
        } else {
            v3["foo bar"] = ("" + (v64 + v60))
            v5 -= 514
            v65 := GenS5{F0: v59}
            _ = v65
        }
        var v66 int = 127
        _ = v66
    }
    fmt.Println(v5, v64, v5)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v60
    _ = v64
}
