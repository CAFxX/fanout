package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 string; F1 string; F2 bool }
type GenS6 struct { F0 uint16; F1 uint; F2 int8 }

func main() {
    var v1 uint = 0
    v2 := make(map[string]bool)
    var v3 uint64 = 18446744073709551615
    var v4 int64 = 428
    if false {
        v3 = 2
        v4 -= -602
        delete(v2, "\n\t")
    }
    delete(v2, "foo bar")
    v5 := GenS1{F0: 0, F1: ("foo bar" + "\n\t")}
    if true {
        v6 := []int{-1, 9223372036854775807}
        fmt.Println(v1, v1, v3)
        _ = v6
    } else {
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := []int{970, (fi7 - 9223372036854775807)}
            v4 -= 1
            v4 -= -9223372036854775808
            _ = fi7
            _ = v8
        }
    }
    var v9 uint64 = 157
    if false {
        switch 792 {
        case -2147483648:
            v10 := []int{}
            _ = v10
        case -899:
            if true {
                v11 := []int{-9223372036854775808, 127}
                var v12 int8 = 127
                _ = v11
                _ = v12
            }
        }
    } else {
        fmt.Println(v4)
    }
    fmt.Println(v9)
    if true {
        switch ("\n\t" + (("foo bar" + "\n\t") + "")) {
        case "\n\t":
            var v13 uint16 = 65535
            _ = v13
        case "hello":
            var v14 uint32 = 483
            _ = v14
        }
        v1 = 285
    } else {
        if (false || (false || false)) {
            if false {
                v4 = v4
            }
            for fi15 := 0; fi15 < 1; fi15++ {
                for fi16 := 0; fi16 < 6; fi16++ {
                    switch "\n\t" {
                    case "a":
                        switch fi16 {
                        case -2147483648:
                            fmt.Println(v9, v9)
                        case -222:
                            v17 := []int{fi15, fi15}
                            _ = v17
                        case -418:
                            var v18 int8 = 1
                            _ = v18
                        default:
                            v9 -= 18446744073709551615
                            v4 = v4
                        }
                    default:
                        switch (fi15 * fi16) {
                        case -499:
                            var v19 string = "foo bar"
                            _ = v19
                        }
                    }
                    v20 := GenS2{F0: ((fi15 | fi15) > (fi16 | (fi16 + (fi16 - (fi15 & (fi15 - fi15)))))), F1: -2147483648}
                    switch "foo bar" {
                    case "\n\t":
                        v9 -= 18446744073709551615
                    default:
                        v21 := []int{}
                        _ = v21
                    }
                    _ = fi16
                    _ = v20
                }
                if false {
                    for fi22 := 0; fi22 < 18; fi22++ {
                        var v23 int8 = 95
                        v9 += v9
                        if (43 > 1) {
                            v1 = (v1 & 668)
                            v3 += v9
                        }
                        _ = fi22
                        _ = v23
                    }
                } else {
                    v24 := GenS3{F0: v1}
                    var v25 string = "hello"
                    v26 := GenS3{F0: v1}
                    _ = v24
                    _ = v25
                    _ = v26
                }
                var v27 string = ("hello" + ("a" + "\n\t"))
                _ = fi15
                _ = v27
            }
            v1 -= (v1 | 42)
        } else {
            v28 := []int{-724}
            _ = v28
        }
        v29 := GenS4{F0: "hello", F1: 127}
        var v30 string = ("hello" + "foo bar")
        _ = v29
        _ = v30
    }
    fmt.Println(v9, v9, v9)
    if ((42 < -1) && (false || true)) {
        if (1 != int32(v1)) {
            v31 := GenS1{F0: 42, F1: "a"}
            v32 := []int{127}
            _ = v31
            _ = v32
        }
        for fi33 := 0; fi33 < 5; fi33++ {
            var v34 int = fi33
            var v35 uint32 = 583
            _ = fi33
            _ = v34
            _ = v35
        }
    } else {
        var v36 int64 = (-(v4))
        _ = v36
    }
    if false {
        if false {
            v37 := GenS4{F0: "foo bar", F1: 99}
            v38 := GenS4{F0: (("\n\t" + "a") + "foo bar"), F1: 255}
            if (-9 <= -81) {
                var v39 int = 0
                v4 = v4
                _ = v39
            } else {
                v2["\n\t"] = (false || true)
                if (-1 > 346) {
                    v40 := []int{0}
                    if len(v40) >= 2 {
                        v40 = v40[1:len(v40)]
                    }
                    _ = v40
                }
            }
            _ = v37
            _ = v38
        }
        if false {
            if false {
                v41 := GenS4{F0: ("foo bar" + (("\n\t" + ("hello" + ("hello" + ((("\n\t" + ("\n\t" + "")) + "hello") + "a")))) + ("foo bar" + (((("\n\t" + "a") + ("foo bar" + (("foo bar" + ("" + "a")) + ("\n\t" + "hello")))) + ("hello" + "")) + ("foo bar" + "foo bar"))))), F1: 130}
                if (!(1e10 > 1e10)) {
                    v42 := v2["foo bar"]
                    fmt.Println(v9)
                    _ = v42
                } else {
                    if (((-48 > 32) && (!false)) && (2147483647 < -2147483648)) {
                        v43 := GenS4{F0: "foo bar", F1: 144}
                        for fi44 := 0; fi44 < 3; fi44++ {
                            v1 -= v1
                            v45 := GenS5{F0: "", F1: "", F2: (true && ((v9 ^ v9) <= v9))}
                            for fi46 := 0; fi46 < 20; fi46++ {
                                v1 -= 1
                                var v47 uint = 255
                                v48 := GenS5{F0: "a", F1: "hello", F2: false}
                                _ = fi46
                                _ = v47
                                _ = v48
                            }
                            _ = fi44
                            _ = v45
                        }
                        var v49 uint = (v1 ^ v1)
                        _ = v43
                        _ = v49
                    }
                }
                _ = v41
            }
            v50 := []int{500, 108}
            _ = v50
        } else {
            v1 += v1
        }
        var v51 float64 = 0.1
        _ = v51
    }
    var v52 float32 = (((1.5 + ((((((-2.25 / 1e10) * (0.0 - 1e10)) / (((((1.5 + (1.5 / float32((0.1 / (((1e10 - 0.0) - 0.0) - 1.5))))) / float32(3.14159)) * -2.25) + 1.5) * 1.5)) + 0.1) - (1e10 / 0.1)) + (float32((((0.0 - 0.0) * 3.14159) / 1.5)) + 3.14159))) / float32(-2.25)) - -2.25)
    for fi53 := 0; fi53 < 20; fi53++ {
        v2[("a" + ((("foo bar" + "") + ("a" + "hello")) + "\n\t"))] = (575 > 42)
        _ = fi53
    }
    switch int(v4) {
    case -1:
        var v54 int8 = 52
        _ = v54
    case 475:
        v9 = v3
    case 677:
        v55 := GenS2{F0: true, F1: -1}
        _ = v55
    }
    switch -2147483648 {
    case 4294967295:
        if false {
            var v56 bool = false
            v3 = (v3 - 2)
            v57 := GenS1{F0: 127, F1: (("" + "hello") + "")}
            _ = v56
            _ = v57
        } else {
            switch 32768 {
            case 4294967295:
                v58 := []int{-9223372036854775808}
                _ = v58
            case 755:
                v59 := []int{4294967295}
                _ = v59
            case -667:
                v1 = v1
            default:
                if true {
                    for fi60 := 0; fi60 < 20; fi60++ {
                        v52 += v52
                        _ = fi60
                    }
                    if (-336 != 42) {
                        if (0.1 != v52) {
                            v9 = (v9 << ((v1) & 63))
                            v61 := GenS6{F0: 42, F1: v1, F2: 98}
                            if false {
                                for fi62 := 0; fi62 < 18; fi62++ {
                                    for fi63 := 0; fi63 < 14; fi63++ {
                                        if (fi62 == fi62) {
                                            var v64 float64 = float64(v52)
                                            _ = v64
                                        }
                                        delete(v2, ("" + (("" + ("hello" + "hello")) + ("foo bar" + (("" + "a") + "a")))))
                                        _ = fi63
                                    }
                                    v65 := []int{32767, (fi62 | fi62), fi62}
                                    _ = fi62
                                    _ = v65
                                }
                                if (false || (false || (((v1 * 255) != (^(v1))) || (0 <= 42)))) {
                                    var v66 bool = (!((v4 & v4) >= v4))
                                    v4 = v4
                                    var v67 int32 = -1
                                    _ = v66
                                    _ = v67
                                }
                            } else {
                                v9 += v3
                                v68 := []int{284, 1}
                                _ = v68
                            }
                            _ = v61
                        } else {
                            var v69 int = -468
                            _ = v69
                        }
                    }
                }
            }
            var v70 bool = (2147483647 < 2147483647)
            _ = v70
        }
    case -277:
        if true {
            v52 += (v52 / 1.5)
            for fi71 := 0; fi71 < 8; fi71++ {
                v1 += 218
                if false {
                    var v72 string = "hello"
                    var v73 float64 = 0.1
                    _ = v72
                    _ = v73
                } else {
                    fmt.Println(v1, fi71, v9)
                }
                _ = fi71
            }
        }
    case 32767:
        var v74 string = ("hello" + "")
        _ = v74
    }
    fmt.Println(v9, v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v9
}
