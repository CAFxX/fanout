package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS4 struct { F0 string; F1 uint8; F2 uint64 }
type GenS5 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint) int16 {
    fmt.Println(p0)
    var v6 uint64 = 18446744073709551615
    v7 := GenS1{F0: p0}
    _ = p0
    _ = v6
    _ = v7
    return -32768
}
func genFn1(p0 bool) bool {
    p0 = (601 > 139)
    switch (("hello" + "foo bar") + "\n\t") {
    case "hello":
        p0 = (p0 && (18446744073709551615 >= 1))
    case "a":
        switch ("" + "foo bar") {
        case "a":
            switch ("a" + ("hello" + (("" + "\n\t") + ""))) {
            case "":
                v8 := []int{}
                _ = v8
            case "\n\t":
                if false {
                    switch (((((("" + "a") + (("foo bar" + (("hello" + "foo bar") + ("" + ("a" + "foo bar")))) + ("foo bar" + ("a" + "")))) + ("" + "a")) + ("\n\t" + "foo bar")) + ("foo bar" + "a")) + "\n\t") {
                    case "\n\t":
                        fmt.Println(p0)
                    case "foo bar":
                        var v9 int = 9223372036854775807
                        _ = v9
                    case "hello":
                        v10 := GenS1{F0: 178}
                        _ = v10
                    default:
                        v11 := []int{1}
                        _ = v11
                    }
                    p0 = ((18446744073709551615 >= 52) && false)
                    var v12 uint8 = 2
                    _ = v12
                } else {
                    p0 = p0
                    var v13 int32 = 266
                    _ = v13
                }
            case "a":
                for fi14 := 0; fi14 < 2; fi14++ {
                    var v15 uint = 255
                    for fi16 := 0; fi16 < 10; fi16++ {
                        switch "foo bar" {
                        case "":
                            v15 = (v15 >> (((v15 % (((v15 ^ v15)) | 1))) & 63))
                        case "hello":
                            switch (fi16 / ((2147483647) | 1)) {
                            case -9223372036854775808:
                                var v17 float64 = (3.14159 - (3.14159 + (-2.25 / 1e10)))
                                _ = v17
                            }
                        case "\n\t":
                            var v18 int64 = -28
                            _ = v18
                        }
                        _ = fi16
                    }
                    v15 = v15
                    _ = fi14
                    _ = v15
                }
            }
        case "\n\t":
            v19 := []int{32768, 127, -9223372036854775808}
            _ = v19
        }
    case "foo bar":
        for fi20 := 0; fi20 < 1; fi20++ {
            switch ("foo bar" + "\n\t") {
            case "":
                v21 := GenS2{F0: 863, F1: true}
                _ = v21
            case "\n\t":
                if (4294967295 != 65535) {
                    switch fi20 {
                    case 798:
                        p0 = p0
                    case 2147483647:
                        var v22 uint64 = 1
                        _ = v22
                    case 32767:
                        for fi23 := 0; fi23 < 7; fi23++ {
                            v24 := []int{9223372036854775807, fi20, 2147483647}
                            p0 = p0
                            _ = fi23
                            _ = v24
                        }
                    default:
                        switch "foo bar" {
                        case "foo bar":
                            switch "\n\t" {
                            case "":
                                fmt.Println(fi20, fi20)
                            case "a":
                                p0 = (true || p0)
                            case "foo bar":
                                for fi25 := 0; fi25 < 17; fi25++ {
                                    v26 := GenS3{F0: p0, F1: 13, F2: 2147483647}
                                    v27 := GenS3{F0: p0, F1: 255, F2: 42}
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                            }
                        case "":
                            for fi28 := 0; fi28 < 2; fi28++ {
                                if (!p0) {
                                    fmt.Println(fi20)
                                    fmt.Println(p0)
                                    var v29 bool = (!p0)
                                    _ = v29
                                } else {
                                    var v30 uint8 = 151
                                    _ = v30
                                }
                                var v31 uint = 42
                                v32 := GenS2{F0: 1, F1: (("foo bar" + "foo bar") >= ("foo bar" + ("hello" + "")))}
                                _ = fi28
                                _ = v31
                                _ = v32
                            }
                        case "a":
                            var v33 int64 = 858
                            _ = v33
                        default:
                            p0 = p0
                        }
                    }
                    v34 := GenS4{F0: ("" + ""), F1: 57, F2: 309}
                    if (0 <= 2) {
                        switch "a" {
                        case "hello":
                            var v35 uint = 317
                            _ = v35
                        case "a":
                            v36 := []int{fi20, (fi20 / (((fi20 - (fi20 - fi20))) | 1)), fi20, fi20}
                            _ = v36
                        }
                    } else {
                        var v37 uint16 = 0
                        _ = v37
                    }
                    _ = v34
                } else {
                    v38 := []int{fi20, (-(fi20)), 42}
                    _ = v38
                }
            default:
                if ((!p0) || p0) {
                    var v39 float32 = 1.5
                    p0 = p0
                    _ = v39
                } else {
                    var v40 int8 = 1
                    fmt.Println(fi20)
                    _ = v40
                }
            }
            for fi41 := 0; fi41 < 1; fi41++ {
                v42 := GenS1{F0: 255}
                var v43 uint = 18446744073709551615
                _ = fi41
                _ = v42
                _ = v43
            }
            _ = fi20
        }
    }
    _ = p0
    return (0 >= 100)
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 128
    v3 := []uint32{393, 0}
    v4 := []string{("a" + ((("a" + "\n\t") + ("a" + "foo bar")) + ("\n\t" + "hello")))}
    v5 := make(map[string]uint32)
    if (false || true) {
        var v44 bool = false
        _ = v44
    } else {
        if true {
            switch ("foo bar" + "\n\t") {
            case "a":
                if false {
                    for fi45 := 0; fi45 < 8; fi45++ {
                        var v46 bool = (fi45 < (fi45 | v2))
                        var v47 uint64 = 1
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                    if false {
                        var v48 int16 = 1000
                        v49 := GenS5{F0: 1, F1: 515}
                        var v50 uint64 = 599
                        _ = v48
                        _ = v49
                        _ = v50
                    }
                } else {
                    var v51 int64 = 1
                    if len(v3) > 0 {
                        v3[0] = 65535
                    }
                    _ = v51
                }
            default:
                for fi52 := 0; fi52 < 3; fi52++ {
                    fmt.Println(v2, v2)
                    _ = fi52
                }
            }
            var v53 int8 = 42
            var v54 int8 = (-(v53))
            _ = v53
            _ = v54
        } else {
            fmt.Println(v2)
            v55 := GenS2{F0: 42, F1: (false || false)}
            _ = v55
        }
        fmt.Println(v2)
        v2 = v2
    }
    if ((v2 & v2) > v2) {
        for fi56 := 0; fi56 < 17; fi56++ {
            v2 += -2147483648
            v2 += (fi56 ^ (v2 + (fi56 << ((65535) & 63))))
            _ = len(v5)
            _ = fi56
        }
        for fi57 := 0; fi57 < 17; fi57++ {
            var v58 uint = 65535
            if (451 != 1) {
                if (true && true) {
                    if (!(850 >= 508)) {
                        switch (v2 << ((v58) & 63)) {
                        case -68:
                            var v59 uint = v58
                            _ = v59
                        }
                    } else {
                        var v60 uint8 = 223
                        _ = v60
                    }
                }
            }
            var v61 bool = (0 == 593)
            _ = fi57
            _ = v58
            _ = v61
        }
        if true {
            v5[("a" + ("" + "hello"))] = 2
        }
    }
    var v62 int = v2
    if (true || false) {
        var v63 string = ("foo bar" + "hello")
        v2 += 128
        _ = v63
    } else {
        v62 += (-(v62))
        v2 = (^(v2))
    }
    var v64 uint32 = 0
    var v65 uint = 917
    v62 += (v2 << (((v65 * 255)) & 63))
    v2 = 1
    if len(v3) > 0 {
        v3[0] = v64
    }
    fmt.Println(v2, v2, v2)
    var v66 uint8 = 165
    var v67 uint8 = (v66 ^ v66)
    var v68 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v62:", v62)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v68
}
