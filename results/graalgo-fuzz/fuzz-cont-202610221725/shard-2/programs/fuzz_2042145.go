package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint64; F1 string; F2 uint }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int32; F1 int8; F2 uint16 }
type GenS5 struct { F0 int16; F1 int32; F2 uint32 }
func genFn0(p0 int8, p1 int16, p2 uint) int32 {
    fmt.Println(p2, p1, p0)
    var v6 int32 = 83
    var v7 int32 = 200
    var v8 int = 128
    switch "a" {
    case "\n\t":
        v9 := []int{v8, v8}
        _ = v9
    case "":
        for fi10 := 0; fi10 < 3; fi10++ {
            v7 = (v6 % (((-(v7))) | 1))
            v11 := []int{v8}
            if (65535 > (p2 - p2)) {
                var v12 uint16 = 255
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 float64 = ((1e10 / ((-2.25 * float64((-2.25 / 1e10))) / (((1e10 + 1e10) + (1e10 / (1e10 / float64(0.1)))) + (float64(3.14159) - ((3.14159 / 1.5) / 1.5))))) / 1.5)
                    if (p1 < 32767) {
                        v7 -= v6
                    }
                    _ = fi13
                    _ = v14
                }
                _ = v12
            }
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (v7 % ((733) | 1))
}
func genFn1(p0 uint32, p1 uint) uint {
    switch "foo bar" {
    case "foo bar":
        v15 := GenS1{F0: -1, F1: 42}
        _ = v15
    case "hello":
        v16 := GenS2{F0: 4294967295, F1: ("\n\t" + ""), F2: 0}
        _ = v16
    default:
        v17 := []int{322}
        _ = v17
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        v19 := []int{fi18}
        if (2147483647 != (p0 & (p0 - 42))) {
            var v20 int64 = 1
            v21 := GenS1{F0: 1, F1: -215}
            var v22 int32 = 261
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi18
        _ = v19
    }
    if true {
        var v23 int64 = -9223372036854775808
        var v24 string = "foo bar"
        _ = v23
        _ = v24
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -1
    var v3 int64 = -9223372036854775808
    var v4 bool = (((1.5 / (1.5 - 3.14159)) > (0.0 - 3.14159)) || (!(("hello" != ((("\n\t" + ((((("" + "\n\t") + "foo bar") + (("foo bar" + ("a" + ("\n\t" + ("\n\t" + (("a" + "hello") + "a"))))) + ("\n\t" + (("\n\t" + "") + "foo bar")))) + ("\n\t" + ("\n\t" + ("" + "foo bar")))) + "")) + "") + "foo bar")) && (32767 < 2))))
    var v5 int64 = (v2 | v3)
    for fi25 := 0; fi25 < 3; fi25++ {
        v4 = v4
        _ = fi25
    }
    v4 = false
    var v26 uint = genFn1(65535, 7)
    for fi27 := 0; fi27 < 20; fi27++ {
        for fi28 := 0; fi28 < 2; fi28++ {
            v29 := GenS1{F0: -2147483648, F1: 234}
            if v4 {
                if ((!(1 == 575)) && true) {
                    switch ("\n\t" + "hello") {
                    case "hello":
                        v30 := []int{fi28, fi27, (fi28 * (fi28 | (fi27 * (fi27 << ((65535) & 63)))))}
                        _ = v30
                    case "\n\t":
                        v2 = v3
                    case "foo bar":
                        v5 = 1
                    default:
                        v1[""] = v4
                    }
                    if (!v4) {
                        if (!(!(v26 >= (v26 | (v26 | (v26 * (v26 >> ((v26) & 63)))))))) {
                            var v31 uint16 = 562
                            if (-2.25 >= 1e10) {
                                v3 -= v3
                                _ = len(v1)
                                if (false && v4) {
                                    fmt.Println(fi28)
                                    switch (("hello" + "foo bar") + "foo bar") {
                                    case "hello":
                                        v3 += v3
                                    case "":
                                        var v32 uint8 = 0
                                        _ = v32
                                    case "\n\t":
                                        switch ("hello" + "a") {
                                        case "a":
                                            var v33 uint16 = (v31 | (^(v31)))
                                            _ = v33
                                        case "\n\t":
                                            v3 -= v3
                                        case "foo bar":
                                            switch (fi27 >> (((v26 << (((^(v26))) & 63))) & 63)) {
                                            case -933:
                                                if (!((0.0 - -2.25) == (float32(-2.25) + ((1e10 / 1e10) - float32(1e10))))) {
                                                    var v34 uint32 = 4294967295
                                                    var v35 int16 = -1
                                                    v36 := []int{-710, (fi27 | (^(fi28))), fi27}
                                                    _ = v34
                                                    _ = v35
                                                    _ = v36
                                                } else {
                                                    v37 := []int{}
                                                    _ = v37
                                                }
                                            case 32768:
                                                if v4 {
                                                    v38 := GenS2{F0: 0, F1: (("" + "\n\t") + "foo bar"), F2: genFn1(65535, v26)}
                                                    _ = v38
                                                } else {
                                                    v39 := GenS1{F0: 42, F1: -272}
                                                    v40 := GenS2{F0: 516, F1: "\n\t", F2: v26}
                                                    v26 = (v26 << (((v26 + 2)) & 63))
                                                    _ = v39
                                                    _ = v40
                                                }
                                            case 997:
                                                var v41 uint64 = 0
                                                _ = v41
                                            default:
                                                var v42 string = "a"
                                                _ = v42
                                            }
                                        }
                                    default:
                                        v5 += (v3 - v5)
                                    }
                                } else {
                                    v43 := GenS2{F0: 0, F1: "foo bar", F2: (^(v26))}
                                    v44 := []int{(fi27 + 1), fi27, fi27, fi27}
                                    var v45 int64 = v3
                                    v2 -= 0
                                    var v46 uint64 = 1
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                }
                            } else {
                                v47 := []int{fi28}
                                _ = v47
                            }
                            var v48 int = (fi27 ^ fi28)
                            var v49 uint16 = (v31 * v31)
                            _ = v31
                            _ = v48
                            _ = v49
                        }
                        if true {
                            switch (fi28 | fi28) {
                            case 9223372036854775807:
                                fmt.Println(v26)
                            case 32767:
                                v5 -= v5
                            default:
                                switch "a" {
                                case "foo bar":
                                    delete(v1, ("foo bar" + ((("a" + ("a" + (("foo bar" + "\n\t") + "hello"))) + "a") + "foo bar")))
                                case "hello":
                                    v50 := GenS1{F0: -735, F1: 42}
                                    _ = v50
                                case "\n\t":
                                    if false {
                                        var v51 string = "foo bar"
                                        var v52 uint64 = 42
                                        _ = v51
                                        _ = v52
                                    } else {
                                        var v53 int32 = 1
                                        var v54 uint16 = 988
                                        _ = v53
                                        _ = v54
                                    }
                                }
                            }
                            v55 := GenS3{F0: v3, F1: 2}
                            v56 := []int{fi28, 127, fi28, 127}
                            _ = v55
                            _ = v56
                        } else {
                            var v57 uint8 = 2
                            _ = v57
                        }
                    }
                }
                if v4 {
                    fmt.Println(v2)
                } else {
                    _ = len(v1)
                    fmt.Println(v2, v26, fi28)
                }
                v58 := []int{(fi27 | (fi27 % (((fi27 + fi28)) | 1))), fi27}
                v1[("a" + "hello")] = (127 < 42)
                _ = v58
            }
            _ = fi28
            _ = v29
        }
        if (2147483647 > 4294967295) {
            delete(v1, "")
            v59 := []int{}
            if (((2 < 65535) || (!v4)) || v4) {
                var v60 uint = v26
                _ = v60
            }
            _ = v59
        } else {
            for fi61 := 0; fi61 < 4; fi61++ {
                v62 := GenS2{F0: 4294967295, F1: ("hello" + ""), F2: 555}
                fmt.Println(fi27, v3)
                switch ("foo bar" + ("" + "hello")) {
                case "foo bar":
                    v63 := []int{(fi27 ^ fi61), 4294967295, (fi27 << (((v26 / ((v26) | 1))) & 63))}
                    _ = v63
                }
                _ = fi61
                _ = v62
            }
            switch ("a" + "") {
            case "":
                var v64 int = fi27
                _ = v64
            default:
                v4 = false
            }
            fmt.Println(v2)
        }
        _ = fi27
    }
    v65 := GenS4{F0: -2147483648, F1: -108, F2: 2}
    v66 := GenS2{F0: 810, F1: "foo bar", F2: v26}
    if false {
        fmt.Println(v5, v2, v2)
        delete(v1, ("\n\t" + ""))
    } else {
        v67 := []int{1, 42, -9223372036854775808}
        delete(v1, "a")
        _ = v67
    }
    switch "\n\t" {
    case "a":
        v68 := GenS5{F0: -32768, F1: genFn0(-128, -891, genFn1(2147483647, (v26 | v26))), F2: 0}
        _ = v68
    case "hello":
        if (((((-2.25 * (3.14159 + 0.0)) / 3.14159) / (0.1 + 0.0)) != 1e10) || (!(!v4))) {
            var v69 float64 = ((0.0 / (3.14159 - (3.14159 * 0.0))) / (float64(0.1) * 0.1))
            delete(v1, "hello")
            _ = v69
        } else {
            fmt.Println(v4)
            v3 += (^(v5))
            var v70 uint32 = 2147483648
            _ = v70
        }
    }
    v71 := []int{451, 32768}
    v72 := GenS3{F0: -1, F1: 128}
    v26 -= genFn1(434, 18446744073709551615)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v71:", len(v71), v71)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v71
    _ = v72
}
