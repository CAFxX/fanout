package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 uint8; F2 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 string; F1 uint }
type GenS8 struct { F0 string; F1 int8; F2 uint8 }
func genFn0() int64 {
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := GenS1{F0: 1}
        _ = fi5
        _ = v6
    }
    if (true && (2147483648 > 2147483648)) {
        v7 := GenS2{F0: 255, F1: 0, F2: ("\n\t" + "hello")}
        var v8 uint8 = 32
        fmt.Println(v8)
        _ = v7
        _ = v8
    } else {
        if true {
            var v9 uint64 = 42
            var v10 bool = (!(9223372036854775807 > 800))
            v10 = (v10 && (127 == -112))
            _ = v10
            _ = v9
        } else {
            var v11 int = 32768
            _ = v11
        }
    }
    var v12 int32 = 1
    _ = v12
    return 0
}
func genFn1() string {
    switch "\n\t" {
    case "a":
        v13 := []int{-2147483648}
        _ = v13
    case "":
        switch "a" {
        case "":
            var v14 float32 = 1e10
            _ = v14
        case "hello":
            v15 := GenS1{F0: 88}
            _ = v15
        case "foo bar":
            if true {
                v16 := GenS3{F0: ("\n\t" + "foo bar")}
                _ = v16
            } else {
                var v17 int16 = -530
                switch "" {
                case "hello":
                    v17 -= -26
                case "foo bar":
                    var v18 uint32 = 0
                    _ = v18
                }
                if (!(-1 == 0)) {
                    v19 := []int{}
                    if false {
                        v20 := GenS2{F0: 0, F1: 620, F2: "a"}
                        _ = v20
                    } else {
                        v17 -= (v17 << ((18446744073709551615) & 63))
                    }
                    _ = v19
                }
                _ = v17
            }
        }
    case "\n\t":
        var v21 uint16 = 2
        _ = v21
    default:
        v22 := GenS3{F0: ("\n\t" + "foo bar")}
        _ = v22
    }
    switch "" {
    case "":
        if false {
            var v23 float32 = (1e10 - -2.25)
            _ = v23
        }
    case "foo bar":
        v24 := GenS4{F0: (("" + "a") + "hello"), F1: 0, F2: 255}
        _ = v24
    default:
        v25 := GenS1{F0: 2147483647}
        _ = v25
    }
    return "\n\t"
}

func main() {
    var v1 float32 = 0.1
    var v2 string = "a"
    v3 := []string{v2, (v2 + v2), v2}
    var v4 uint64 = 42
    if len(v3) > 0 {
        v3[0] = v2
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 uint32 = 408
        var v28 uint = uint(v4)
        var v29 uint8 = 230
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    var v30 float32 = 1e10
    v4 -= 18446744073709551615
    v4 += v4
    v31 := GenS5{F0: true}
    var v32 string = ("hello" + "hello")
    var v33 int8 = 43
    if false {
        fmt.Println(v33, v1)
        v1 += (v1 - ((0.0 - 1.5) * (v1 + v30)))
        v34 := GenS6{F0: 42, F1: true}
        _ = v34
    }
    switch "\n\t" {
    case "hello":
        switch v2 {
        case "a":
            var v35 float64 = 0.0
            _ = v35
        case "\n\t":
            var v36 float32 = 0.1
            _ = v36
        case "":
            var v37 int64 = 42
            _ = v37
        default:
            v38 := GenS5{F0: true}
            _ = v38
        }
    case "foo bar":
        for fi39 := 0; fi39 < 14; fi39++ {
            if (true || ((float32(3.14159) >= v30) && (2147483647 >= 4294967295))) {
                if (false || true) {
                    v32 = "\n\t"
                    v3 = append(v3, v2)
                    var v40 int64 = 743
                    _ = v40
                } else {
                    for fi41 := 0; fi41 < 1; fi41++ {
                        switch v2 {
                        case "hello":
                            v30 += v30
                        }
                        if (128 < 255) {
                            if true {
                                if (true && (3.14159 >= 3.14159)) {
                                    var v42 uint = 718
                                    var v43 int = 32768
                                    _ = v42
                                    _ = v43
                                }
                                switch "a" {
                                case "hello":
                                    var v44 uint = 1
                                    _ = v44
                                case "":
                                    var v45 float32 = 1.5
                                    _ = v45
                                default:
                                    if (true && (int64(fi41) >= genFn0())) {
                                        v32 = v2
                                    } else {
                                        var v46 float64 = -2.25
                                        var v47 int = (fi39 - (fi41 >> ((851) & 63)))
                                        _ = v46
                                        _ = v47
                                    }
                                }
                            }
                            var v48 uint32 = 42
                            _ = len(v3)
                            v49 := GenS7{F0: v2, F1: 1}
                            if true {
                                v30 = (v30 + (((v1 * ((v30 + ((0.1 / (v30 * v1)) / -2.25)) * v30)) / v30) + 1.5))
                                v2 = v32
                            } else {
                                v1 = ((((v30 - v30) / 1.0) - 1.5) / v1)
                                var v50 uint64 = (v4 ^ 413)
                                if false {
                                    v4 -= v4
                                } else {
                                    v2 = ""
                                    if true {
                                        if false {
                                            v51 := GenS8{F0: "hello", F1: v33, F2: 0}
                                            switch fi41 {
                                            case -9223372036854775808:
                                                var v52 int32 = 0
                                                _ = v52
                                            case -16:
                                                v2 = "hello"
                                            case 9223372036854775807:
                                                v33 = (v33 + v33)
                                            }
                                            _ = v51
                                        }
                                        if (true && (9223372036854775807 == 1)) {
                                            if false {
                                                var v53 float32 = (float32(0.0) / 0.1)
                                                _ = v53
                                            } else {
                                                v32 = ("" + ("a" + "a"))
                                                fmt.Println(fi39, v33)
                                                var v54 int8 = (-(v33))
                                                _ = v54
                                            }
                                        }
                                    }
                                    if (42 > 726) {
                                        var v55 float64 = (float64((v30 - 3.14159)) + 1.5)
                                        _ = v55
                                    }
                                }
                                _ = v50
                            }
                            _ = v48
                            _ = v49
                        }
                        switch (("foo bar" + v2) + "\n\t") {
                        case "hello":
                            v30 += v30
                        case "\n\t":
                            v2 = ((((("foo bar" + (v32 + "\n\t")) + "a") + v32) + ((v2 + v32) + "hello")) + v32)
                        case "a":
                            var v56 bool = true
                            _ = v56
                        default:
                            v57 := GenS5{F0: true}
                            _ = v57
                        }
                        fmt.Println(v1)
                        fmt.Println(v2)
                        var v58 bool = (255 >= 128)
                        if (161 <= 2) {
                            var v59 int16 = 708
                            var v60 float32 = v30
                            _ = v59
                            _ = v60
                        } else {
                            var v61 int64 = 1
                            var v62 uint = 42
                            var v63 int16 = 494
                            _ = v61
                            _ = v62
                            _ = v63
                        }
                        _ = fi41
                        _ = v58
                    }
                }
            }
            fmt.Println(v32, v33)
            var v64 int32 = 42
            _ = fi39
            _ = v64
        }
    case "a":
        var v65 int8 = (v33 - v33)
        _ = v65
    default:
        _ = len(v3)
    }
    fmt.Println(v32, v4)
    var v66 int64 = -1
    fmt.Println(v30, v4)
    v2 = v32
    var v67 uint8 = 138
    if false {
        var v68 float64 = 1.5
        var v69 uint32 = 945
        _ = v68
        _ = v69
    }
    if (182 >= 412) {
        var v70 float64 = 3.14159
        _ = v70
    }
    for fi71 := 0; fi71 < 6; fi71++ {
        fmt.Println(fi71, v66)
        _ = fi71
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v66
    _ = v67
}
