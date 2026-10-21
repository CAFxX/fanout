package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint64; F1 uint32; F2 bool }
type GenS4 struct { F0 uint16; F1 string; F2 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 int64; F1 int64; F2 uint }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int; F1 uint; F2 int8 }
func genFn0() uint8 {
    if (42 != 1) {
        v5 := []int{-538, 32767}
        if (((4294967295 <= 2) && true) && (!(!true))) {
            var v6 int16 = 42
            _ = v6
        }
        var v7 int16 = -32768
        _ = v5
        _ = v7
    }
    v8 := GenS1{F0: 1000, F1: -26}
    if true {
        v9 := GenS1{F0: 1, F1: -128}
        switch "hello" {
        case "foo bar":
            var v10 float32 = 0.0
            _ = v10
        default:
            var v11 float32 = 3.14159
            _ = v11
        }
        _ = v9
    } else {
        var v12 uint64 = 18446744073709551615
        var v13 int = -1
        _ = v12
        _ = v13
    }
    _ = v8
    return 137
}
func genFn1(p0 uint16) uint8 {
    p0 -= p0
    var v14 int16 = 724
    if true {
        v14 = (v14 * (^(v14)))
        if true {
            if (v14 != v14) {
                v15 := GenS1{F0: v14, F1: 42}
                if (1.5 != ((1e10 * float64((0.0 - (3.14159 - float32(float64(0.0)))))) + 3.14159)) {
                    v16 := GenS1{F0: v14, F1: -2}
                    _ = v16
                } else {
                    fmt.Println(v14, p0)
                }
                if true {
                    if ((!((^(p0)) >= (p0 | p0))) && (!((-(v14)) > (-(v14))))) {
                        for fi17 := 0; fi17 < 18; fi17++ {
                            v18 := GenS2{F0: 2, F1: (fi17 | -346)}
                            var v19 float32 = (1e10 - 1e10)
                            if false {
                                v14 = (^(v14))
                                v20 := GenS3{F0: 4294967295, F1: 0, F2: (v14 >= v14)}
                                v21 := GenS3{F0: 1, F1: 65535, F2: (!false)}
                                _ = v20
                                _ = v21
                            } else {
                                for fi22 := 0; fi22 < 4; fi22++ {
                                    var v23 int16 = v14
                                    var v24 int16 = v23
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                }
                                fmt.Println(fi17, v14, v19)
                                if false {
                                    var v25 uint32 = 2
                                    switch (("a" + (("a" + "a") + ("a" + ("foo bar" + "")))) + ("a" + "")) {
                                    case "foo bar":
                                        p0 = (p0 >> ((255) & 63))
                                    case "\n\t":
                                        switch "" {
                                        case "a":
                                            var v26 string = "a"
                                            _ = v26
                                        case "hello":
                                            for fi27 := 0; fi27 < 2; fi27++ {
                                                v14 -= 0
                                                if true {
                                                    if (18446744073709551615 >= 255) {
                                                        fmt.Println(v19, v19, v25)
                                                    }
                                                    var v28 string = "\n\t"
                                                    if ((!(false || (v28 >= v28))) || ((42 >= 1) && (!((42 > 0) || ((!true) || ((1e10 + (1.5 / 1.0)) < 0.1)))))) {
                                                        v19 += (float32(float64(v19)) * v19)
                                                    } else {
                                                        if false {
                                                            v29 := GenS3{F0: 698, F1: (v25 | v25), F2: ((v25 / (((v25 << ((uint(p0)) & 63))) | 1)) <= (v25 >> ((4294967295) & 63)))}
                                                            v30 := GenS1{F0: v14, F1: 13}
                                                            v25 = (v25 % (((v25 / (((v25 << ((966) & 63))) | 1))) | 1))
                                                            _ = v29
                                                            _ = v30
                                                        }
                                                    }
                                                    _ = v28
                                                } else {
                                                    var v31 int = 32768
                                                    if true {
                                                        if (false || true) {
                                                            switch v31 {
                                                            case 128:
                                                                v25 = (v25 << ((0) & 63))
                                                            case -864:
                                                                switch fi17 {
                                                                case 0:
                                                                    fmt.Println(fi17, fi17, v19)
                                                                default:
                                                                    v32 := GenS4{F0: (p0 >> ((65535) & 63)), F1: ("\n\t" + "hello"), F2: ("" + "")}
                                                                    _ = v32
                                                                }
                                                                v33 := GenS1{F0: v14, F1: 42}
                                                                _ = v33
                                                            case 127:
                                                                var v34 int64 = 42
                                                                _ = v34
                                                            default:
                                                                v31 = (fi27 | 417)
                                                            }
                                                            v19 += v19
                                                            var v35 uint8 = 164
                                                            _ = v35
                                                        }
                                                    }
                                                    _ = v31
                                                }
                                                if (!false) {
                                                    switch fi27 {
                                                    case 534:
                                                        switch fi17 {
                                                        case 32767:
                                                            v36 := GenS5{F0: -9223372036854775808}
                                                            _ = v36
                                                        case 423:
                                                            v37 := GenS6{F0: 423, F1: 2}
                                                            _ = v37
                                                        case 303:
                                                            p0 = p0
                                                        default:
                                                            v38 := GenS1{F0: v14, F1: -128}
                                                            v19 = (((1.5 + v19) * v19) * float32(-2.25))
                                                            v14 -= (v14 % ((v14) | 1))
                                                            _ = v38
                                                        }
                                                    case 4294967295:
                                                        var v39 uint64 = 146
                                                        fmt.Println(v25, v39, fi27)
                                                        switch fi17 {
                                                        case 641:
                                                            v40 := []int{128, (fi27 & fi27), fi27}
                                                            _ = v40
                                                        case 340:
                                                            v39 += v39
                                                        case 926:
                                                            if false {
                                                                v41 := GenS7{F0: 0, F1: 980, F2: 42}
                                                                v14 = v14
                                                                v42 := GenS1{F0: v14, F1: 46}
                                                                _ = v41
                                                                _ = v42
                                                            }
                                                        }
                                                        _ = v39
                                                    case -1:
                                                        fmt.Println(v25, fi17, v14)
                                                    default:
                                                        p0 -= p0
                                                    }
                                                }
                                                v19 = 3.14159
                                                fmt.Println(p0)
                                                if (1.5 == float32((3.14159 * (-2.25 - 1.5)))) {
                                                    v19 = 1e10
                                                    var v43 uint16 = p0
                                                    v44 := GenS8{F0: v25}
                                                    _ = v43
                                                    _ = v44
                                                }
                                                _ = fi27
                                            }
                                        case "":
                                            var v45 float64 = 1.5
                                            _ = v45
                                        default:
                                            v25 += v25
                                        }
                                    case "a":
                                        v25 = 242
                                    default:
                                        var v46 uint = 0
                                        _ = v46
                                    }
                                    _ = v25
                                } else {
                                    v47 := []int{2147483647, (fi17 / ((fi17) | 1))}
                                    v19 = (0.0 - (3.14159 - float32(((-2.25 / ((3.14159 * 1.5) / (((-2.25 / 1.0) * 1e10) * 1e10))) / 1.0))))
                                    fmt.Println(p0)
                                    _ = v47
                                }
                            }
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    } else {
                        v48 := []int{32768}
                        p0 += 870
                        for fi49 := 0; fi49 < 8; fi49++ {
                            var v50 uint16 = 65535
                            _ = fi49
                            _ = v50
                        }
                        _ = v48
                    }
                }
                _ = v15
            }
        }
    }
    var v51 bool = (150 >= 2)
    _ = p0
    _ = v14
    _ = v51
    return 0
}

func main() {
    v1 := []int{0, 2147483647}
    v2 := make(map[string]int)
    var v3 int = 127
    var v4 string = "hello"
    v4 = v4
    delete(v2, "a")
    var v52 uint16 = 2
    v52 = (v52 % ((v52) | 1))
    var v53 uint64 = 394
    switch ((v4 + v4) + ("a" + ("a" + (v4 + v4)))) {
    case "a":
        var v54 int64 = -9223372036854775808
        _ = v54
    case "\n\t":
        fmt.Println(v52, v53)
    }
    fmt.Println(v53)
    if (0 == -1) {
        fmt.Println(v3, v53)
        var v55 float32 = (1.5 * 1.5)
        _ = v55
    }
    v2[("foo bar" + v4)] = (v3 / ((v3) | 1))
    v56 := v2[("hello" + "hello")]
    switch (-(v3)) {
    case 9223372036854775807:
        var v57 int16 = 42
        _ = v57
    case -414:
        v58 := GenS9{F0: (v3 << ((42) & 63)), F1: 255, F2: 0}
        _ = v58
    case 32768:
        if (2 >= 455) {
            v59 := GenS9{F0: v56, F1: 42, F2: 114}
            for fi60 := 0; fi60 < 11; fi60++ {
                if false {
                    var v61 int = (fi60 - -1)
                    v2[("\n\t" + "hello")] = v61
                    _ = v61
                } else {
                    v53 = v53
                }
                var v62 float32 = ((((((0.1 * 3.14159) - float32(3.14159)) + 1e10) - 3.14159) / 1.5) + 3.14159)
                fmt.Println(v53, v53, v52)
                _ = fi60
                _ = v62
            }
            v63 := GenS5{F0: 42}
            _ = v59
            _ = v63
        } else {
            v4 = v4
        }
    default:
        if len(v1) > 0 {
            v1[0] = (v56 << ((0) & 63))
        }
    }
    v64 := GenS1{F0: 42, F1: 24}
    for fi65 := 0; fi65 < 13; fi65++ {
        var v66 uint = 0
        v67 := GenS4{F0: v52, F1: "", F2: ((("" + "") + (v4 + ((v4 + ("foo bar" + v4)) + "a"))) + "\n\t")}
        _ = fi65
        _ = v66
        _ = v67
    }
    v52 = (v52 & uint16(v53))
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v56
    _ = v64
}
