package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int; F1 uint; F2 string }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 uint; F2 int8 }
type GenS7 struct { F0 uint16; F1 int }
func genFn0() uint64 {
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{128, (fi7 - (fi7 << ((4294967295) & 63))), fi7}
        _ = fi7
        _ = v8
    }
    var v9 int = 494
    v10 := []int{(v9 * (v9 - -748)), (v9 % ((v9) | 1))}
    var v11 int = v9
    _ = v10
    _ = v11
    _ = v9
    return 421
}
func genFn1() uint32 {
    switch ("" + "") {
    case "a":
        var v12 uint32 = 2
        _ = v12
    case "foo bar":
        var v13 int = -2147483648
        _ = v13
    case "\n\t":
        var v14 bool = true
        _ = v14
    }
    var v15 int32 = 2147483647
    var v16 uint = 2
    var v17 int32 = -88
    _ = v15
    _ = v16
    _ = v17
    return 0
}

func main() {
    v1 := []string{"", "foo bar"}
    var v2 int16 = 648
    var v3 uint16 = 255
    var v4 uint = 1
    var v5 int64 = -9223372036854775808
    v6 := make(map[int]string)
    var v18 uint16 = (v3 * (v3 | v3))
    if (!((v3 & v18) != 2)) {
        switch ((("foo bar" + "\n\t") + ("" + "foo bar")) + "foo bar") {
        case "":
            var v19 uint16 = (^(v18))
            _ = v19
        case "a":
            v20 := GenS1{F0: -1, F1: 1, F2: 0}
            _ = v20
        }
    }
    v1 = append(v1, "foo bar")
    if true {
        for fi21 := 0; fi21 < 1; fi21++ {
            for fi22 := 0; fi22 < 10; fi22++ {
                switch "foo bar" {
                case "\n\t":
                    if (!(!(false || (true && false)))) {
                        var v23 uint64 = 18446744073709551615
                        var v24 float64 = (-2.25 - 0.1)
                        v5 += v5
                        _ = v23
                        _ = v24
                    } else {
                        if (0.1 != 0.0) {
                            if false {
                                v5 = (v5 + 9223372036854775807)
                                var v25 bool = ((fi21 | (-(fi21))) == 127)
                                if (4294967295 < 255) {
                                    v6[127] = "foo bar"
                                } else {
                                    if v25 {
                                        if true {
                                            var v26 int64 = (v5 - v5)
                                            if true {
                                                var v27 string = ("foo bar" + ((("a" + "hello") + "foo bar") + ""))
                                                v1 = append(v1, "hello")
                                                var v28 int8 = 0
                                                _ = v27
                                                _ = v28
                                            }
                                            if (!(!true)) {
                                                v2 += -329
                                                v2 = (v2 ^ (^(v2)))
                                            } else {
                                                _ = len(v6)
                                                switch fi21 {
                                                case 9223372036854775807:
                                                    if (((255 < (v4 & v4)) && (!(v25 && v25))) || v25) {
                                                        v4 -= v4
                                                    } else {
                                                        var v29 uint16 = v18
                                                        v30 := GenS2{F0: (fi22 & int(v18)), F1: (v4 * v4), F2: "hello"}
                                                        var v31 int16 = (v2 >> ((171) & 63))
                                                        var v32 uint64 = 0
                                                        _ = v29
                                                        _ = v30
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                case 348:
                                                    if v25 {
                                                        v33 := GenS2{F0: (-(fi21)), F1: (v4 >> ((v4) & 63)), F2: "a"}
                                                        var v34 uint64 = 18446744073709551615
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                case -9223372036854775808:
                                                    if len(v1) > 0 {
                                                        v1[0] = "\n\t"
                                                    }
                                                }
                                                if ((^(v4)) >= (v4 >> (((v4 - v4)) & 63))) {
                                                    var v35 uint = v4
                                                    var v36 uint64 = 965
                                                    v37 := GenS2{F0: (fi21 | fi21), F1: v35, F2: "a"}
                                                    delete(v6, fi22)
                                                    v38 := GenS1{F0: fi21, F1: 145, F2: 2}
                                                    _ = v35
                                                    _ = v36
                                                    _ = v37
                                                    _ = v38
                                                }
                                            }
                                            fmt.Println(v18, v2)
                                            _ = v26
                                        }
                                        fmt.Println(v5)
                                        var v39 int64 = (v5 - (v5 + (v5 % (((^(v5))) | 1))))
                                        _ = v39
                                    } else {
                                        v40 := GenS1{F0: fi21, F1: 42, F2: 18446744073709551615}
                                        _ = v40
                                    }
                                    v5 -= (v5 - (v5 / (((v5 % ((767) | 1))) | 1)))
                                }
                                v18 = 890
                                _ = v25
                            }
                            v18 -= v3
                            v18 = (v3 << ((4294967295) & 63))
                        } else {
                            v4 = v4
                            v6[-9223372036854775808] = "hello"
                            var v41 int = -9223372036854775808
                            _ = v41
                        }
                    }
                case "hello":
                    var v42 int = (fi22 & (fi22 + fi21))
                    _ = v42
                case "a":
                    fmt.Println(v3)
                default:
                    v18 += v18
                }
                if (95 != 156) {
                    if ((v5 >> ((v4) & 63)) >= v5) {
                        fmt.Println(v2)
                        switch fi22 {
                        case 733:
                            switch 2147483647 {
                            case 9223372036854775807:
                                v43 := GenS3{F0: 217, F1: 42, F2: 65535}
                                _ = v43
                            default:
                                var v44 string = (("a" + "hello") + "a")
                                _ = v44
                            }
                            switch ("a" + "a") {
                            case "\n\t":
                                var v45 int16 = -32768
                                _ = v45
                            case "hello":
                                if len(v1) > 0 {
                                    v1[0] = "a"
                                }
                            default:
                                v4 += (v4 | 112)
                            }
                            if true {
                                if true {
                                    var v46 int64 = (v5 | 1)
                                    v4 += (^(v4))
                                    _ = v46
                                } else {
                                    var v47 uint64 = 42
                                    switch "\n\t" {
                                    case "foo bar":
                                        fmt.Println(v47, v4, v3)
                                    case "a":
                                        if ((128 >= 9223372036854775807) || true) {
                                            v4 -= uint(v18)
                                        } else {
                                            switch "a" {
                                            case "\n\t":
                                                if true {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                } else {
                                                    switch -103 {
                                                    case 32767:
                                                        var v48 int = (fi21 >> ((42) & 63))
                                                        _ = v48
                                                    case 2147483647:
                                                        v47 = v47
                                                    case 1:
                                                        if true {
                                                            var v49 uint32 = 2147483648
                                                            v49 = genFn1()
                                                            _ = v49
                                                        }
                                                    }
                                                }
                                            }
                                            var v50 int = (^(fi21))
                                            _ = v50
                                        }
                                    }
                                    _ = v47
                                }
                                v6[fi21] = "\n\t"
                                var v51 int16 = v2
                                _ = v51
                            }
                        default:
                            v52 := GenS2{F0: (fi21 * -5), F1: (^(v4)), F2: ("" + "a")}
                            _ = v52
                        }
                    }
                } else {
                    v53 := GenS4{F0: 1000, F1: 1000}
                    v6[fi21] = "foo bar"
                    _ = v53
                }
                _ = fi22
            }
            var v54 uint = v4
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi21
            _ = v54
        }
        var v55 int8 = -1
        _ = v55
    }
    v6[236] = "a"
    for fi56 := 0; fi56 < 20; fi56++ {
        if (!(-2147483648 == 42)) {
            v57 := GenS5{F0: genFn0()}
            _ = len(v6)
            _ = v57
        } else {
            v58 := GenS6{F0: fi56, F1: v4, F2: -128}
            v4 = (v4 * (^(v4)))
            _ = len(v6)
            _ = v58
        }
        var v59 float64 = 3.14159
        _ = fi56
        _ = v59
    }
    if len(v1) > 0 {
        v1[0] = ("foo bar" + ("" + ""))
    }
    v60 := GenS7{F0: v3, F1: 140}
    if len(v1) > 0 {
        v1[0] = ("foo bar" + "")
    }
    switch 32767 {
    case 128:
        fmt.Println(v4, v18)
    case 4294967295:
        v61 := GenS3{F0: 182, F1: 130, F2: 286}
        _ = v61
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
