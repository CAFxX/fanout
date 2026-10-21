package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 bool; F1 int; F2 int }
type GenS3 struct { F0 bool; F1 int16; F2 string }
type GenS4 struct { F0 int }
func genFn0(p0 uint64, p1 uint) int32 {
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := []int{fi6, (fi6 >> ((p1) & 63)), -2147483648}
        v8 := GenS1{F0: 87, F1: 229}
        _ = fi6
        _ = v7
        _ = v8
    }
    p0 = (^(p0))
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 float64 = (-2.25 + ((3.14159 / 1e10) / -2.25))
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    return int32(p0)
}
func genFn1() int {
    v11 := []int{498, 32768, 32767}
    v12 := GenS1{F0: 128, F1: 94}
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 int = (fi13 | (fi13 % ((4294967295) | 1)))
        var v15 int = (fi13 + fi13)
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := []string{"a"}
    var v2 string = ((("hello" + "") + (("hello" + "foo bar") + "hello")) + "\n\t")
    var v3 bool = false
    var v4 string = v2
    var v5 int64 = 9223372036854775807
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := GenS2{F0: v3, F1: fi16, F2: genFn1()}
        fmt.Println(v5, fi16)
        _ = fi16
        _ = v17
    }
    switch "\n\t" {
    case "foo bar":
        v18 := GenS1{F0: 42, F1: 1}
        _ = v18
    case "a":
        v3 = false
    }
    v3 = (!v3)
    var v19 int32 = genFn0(2, 65535)
    if (genFn0(218, 153) >= v19) {
        v3 = true
        switch 9223372036854775807 {
        case -2147483648:
            var v20 int64 = v5
            _ = v20
        case 1:
            for fi21 := 0; fi21 < 18; fi21++ {
                if (!true) {
                    for fi22 := 0; fi22 < 20; fi22++ {
                        if (((-2.25 - 0.0) >= 1.5) || (!((!false) || ("hello" >= v4)))) {
                            var v23 float64 = -2.25
                            _ = len(v1)
                            _ = v23
                        } else {
                            var v24 float64 = (((-2.25 * -2.25) / -2.25) + 1e10)
                            _ = v24
                        }
                        switch fi21 {
                        case 232:
                            switch (v4 + "") {
                            case "\n\t":
                                v3 = ((v3 || true) && v3)
                            case "hello":
                                if true {
                                    fmt.Println(v4, fi21, v5)
                                    v3 = (!(1 <= 49))
                                } else {
                                    if (v3 || v3) {
                                        v5 -= (-(v5))
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                    v3 = v3
                                }
                            }
                        case 1:
                            if v3 {
                                var v25 string = ""
                                v25 = v4
                                _ = v25
                            } else {
                                var v26 int32 = v19
                                fmt.Println(v4, v19, v26)
                                _ = v26
                            }
                        case -30:
                            v1 = append(v1, v4)
                        default:
                            var v27 int = fi22
                            _ = v27
                        }
                        v3 = (v3 && ((867 >= -1) && (0.0 <= 0.1)))
                        _ = fi22
                    }
                    var v28 float32 = float32(float64(0.0))
                    for fi29 := 0; fi29 < 3; fi29++ {
                        var v30 bool = ("foo bar" >= (v2 + v2))
                        var v31 bool = true
                        fmt.Println(v5)
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = v28
                } else {
                    for fi32 := 0; fi32 < 9; fi32++ {
                        var v33 uint64 = 943
                        _ = fi32
                        _ = v33
                    }
                }
                var v34 uint32 = 0
                _ = fi21
                _ = v34
            }
        default:
            var v35 uint64 = 4294967295
            _ = v35
        }
    }
    _ = len(v1)
    fmt.Println(v4, v4, v3)
    v2 = (v4 + "a")
    var v36 int64 = -639
    if (528 < 2) {
        v5 += 42
    }
    fmt.Println(v2, v19)
    switch -386 {
    case -2147483648:
        v1 = append(v1, "foo bar")
    default:
        switch v2 {
        case "foo bar":
            var v37 int32 = (v19 ^ v19)
            _ = v37
        case "hello":
            var v38 string = v2
            _ = v38
        case "":
            for fi39 := 0; fi39 < 8; fi39++ {
                v1 = append(v1, v2)
                if v3 {
                    v40 := GenS3{F0: (0 != 1), F1: 560, F2: v2}
                    for fi41 := 0; fi41 < 7; fi41++ {
                        v42 := GenS3{F0: ((v5 | v5) != (v5 + v36)), F1: -32768, F2: (v2 + (("foo bar" + v4) + "\n\t"))}
                        _ = fi41
                        _ = v42
                    }
                    switch "" {
                    case "\n\t":
                        v5 = int64(v19)
                    case "":
                        v2 = ((v2 + v2) + (v2 + v4))
                    }
                    _ = v40
                }
                _ = fi39
            }
        default:
            switch v4 {
            case "\n\t":
                if (((1e10 / 3.14159) + float64(1.5)) == (3.14159 + -2.25)) {
                    for fi43 := 0; fi43 < 20; fi43++ {
                        var v44 uint16 = 129
                        for fi45 := 0; fi45 < 15; fi45++ {
                            switch genFn1() {
                            case 9223372036854775807:
                                var v46 uint8 = 255
                                _ = v46
                            case 32767:
                                var v47 int16 = -1
                                _ = v47
                            }
                            v3 = false
                            _ = fi45
                        }
                        fmt.Println(v5, v3, v4)
                        _ = fi43
                        _ = v44
                    }
                } else {
                    v19 = 293
                }
            case "a":
                var v48 uint16 = 961
                _ = v48
            default:
                switch "a" {
                case "\n\t":
                    v49 := GenS3{F0: v3, F1: 1, F2: v2}
                    _ = v49
                case "foo bar":
                    if v3 {
                        v5 = v5
                        var v50 uint8 = 2
                        for fi51 := 0; fi51 < 4; fi51++ {
                            switch (v2 + (((("foo bar" + "") + v4) + v2) + v2)) {
                            case "\n\t":
                                _ = len(v1)
                            }
                            if v3 {
                                var v52 float64 = 3.14159
                                for fi53 := 0; fi53 < 13; fi53++ {
                                    v4 = v4
                                    if true {
                                        v50 -= v50
                                        fmt.Println(v50, v52)
                                        v3 = true
                                    } else {
                                        switch (v4 + v2) {
                                        case "":
                                            if (v3 || (((0.1 * float32((v52 / (((0.1 - ((v52 / v52) - (1.5 / v52))) + float64(((1.5 - float32((v52 / (3.14159 - float64(0.1))))) * (float32(0.1) * 0.0)))) * ((((v52 / 3.14159) / v52) * 0.0) - (v52 + float64(3.14159))))))) - (1e10 / (-2.25 - 1e10))) != 1.5)) {
                                                var v54 uint = 1
                                                var v55 int8 = -1
                                                v1 = append(v1, v2)
                                                v56 := GenS1{F0: (v50 / ((128) | 1)), F1: (^(v50))}
                                                fmt.Println(v50, v5)
                                                _ = v54
                                                _ = v55
                                                _ = v56
                                            }
                                        case "\n\t":
                                            fmt.Println(fi53)
                                        case "foo bar":
                                            if v3 {
                                                var v57 float32 = 0.1
                                                switch (v4 + v2) {
                                                case "a":
                                                    var v58 uint16 = 65535
                                                    _ = v58
                                                case "":
                                                    if ((v50 ^ (^(v50))) > v50) {
                                                        fmt.Println(v52)
                                                        var v59 int64 = (v5 | (v5 % (((^(v5))) | 1)))
                                                        _ = v59
                                                    } else {
                                                        v1 = append(v1, (v2 + (v4 + v4)))
                                                        var v60 uint32 = 6
                                                        _ = v60
                                                    }
                                                }
                                                _ = v57
                                            }
                                            var v61 int = (-(fi51))
                                            if len(v1) > 0 {
                                                v1[0] = (v2 + ("\n\t" + ("\n\t" + "hello")))
                                            }
                                            fmt.Println(v19, fi53, v52)
                                            _ = v61
                                        default:
                                            v62 := GenS1{F0: (v50 | v50), F1: v50}
                                            _ = v62
                                        }
                                    }
                                    _ = fi53
                                }
                                _ = v52
                            } else {
                                v3 = (-268 >= -1)
                                var v63 uint = 18446744073709551615
                                if len(v1) > 0 {
                                    v1[0] = v4
                                }
                                _ = v63
                            }
                            _ = fi51
                        }
                        _ = v50
                    } else {
                        v2 = "foo bar"
                        var v64 int64 = v36
                        _ = v64
                    }
                }
            }
        }
    }
    if v3 {
        var v65 string = (v2 + v2)
        for fi66 := 0; fi66 < 9; fi66++ {
            v67 := GenS1{F0: 235, F1: 42}
            v65 = v65
            if len(v1) > 0 {
                v1[0] = (((v65 + "foo bar") + v4) + ("" + (v65 + v2)))
            }
            _ = fi66
            _ = v67
        }
        _ = v65
    }
    v68 := GenS4{F0: -192}
    var v69 float64 = -2.25
    fmt.Println(v69, v3)
    v1 = append(v1, "a")
    v4 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v68
    _ = v69
}
