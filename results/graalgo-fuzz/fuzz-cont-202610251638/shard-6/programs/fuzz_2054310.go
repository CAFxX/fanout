package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 string }
type GenS2 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS3 struct { F0 string; F1 int64; F2 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 int8; F1 bool }
func genFn0(p0 uint8) uint8 {
    for fi5 := 0; fi5 < 10; fi5++ {
        p0 -= (p0 << ((255) & 63))
        var v6 float64 = (1.5 - float64(1e10))
        p0 = 0
        _ = fi5
        _ = v6
    }
    p0 -= p0
    p0 = (p0 % ((1) | 1))
    p0 += p0
    _ = p0
    return p0
}

func main() {
    var v1 int16 = 50
    var v2 int = -370
    var v3 float64 = ((0.1 + 1e10) * 0.1)
    var v4 uint16 = 1
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 2147483647, F1: "a", F2: ("foo bar" + ("hello" + "a"))}
        var v9 int16 = (-(v1))
        if true {
            v1 -= (v1 & v9)
            v10 := []int{}
            _ = v10
        } else {
            var v11 bool = (!false)
            v12 := GenS2{F0: 42, F1: 42, F2: 113}
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v13 bool = false
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 float32 = 0.0
        v2 -= (v2 | (^(fi14)))
        _ = fi14
        _ = v15
    }
    fmt.Println(v13, v2, v4)
    if false {
        v16 := GenS2{F0: 1, F1: 75, F2: 1}
        if v13 {
            var v17 uint8 = 217
            v18 := GenS3{F0: "\n\t", F1: 0, F2: 9223372036854775807}
            var v19 int8 = -2
            _ = v17
            _ = v18
            _ = v19
        }
        _ = v16
    }
    v20 := GenS4{F0: 42}
    for fi21 := 0; fi21 < 5; fi21++ {
        for fi22 := 0; fi22 < 4; fi22++ {
            v23 := GenS2{F0: 2147483647, F1: -128, F2: 2}
            v13 = true
            _ = fi22
            _ = v23
        }
        fmt.Println(v2, v4, v4)
        _ = fi21
    }
    switch (v2 + (v2 & v2)) {
    case -348:
        fmt.Println(v1, v3, v13)
    case 32767:
        var v24 int = (-(v2))
        _ = v24
    case -341:
        v25 := GenS3{F0: "a", F1: 42, F2: 0}
        _ = v25
    default:
        v1 = int16(v4)
    }
    v26 := GenS5{F0: 42}
    v27 := GenS3{F0: ("a" + ""), F1: 42, F2: -9223372036854775808}
    v28 := []int{v2}
    _ = len(v28)
    for fi29 := 0; fi29 < 19; fi29++ {
        if len(v28) > 0 {
            v28[0] = fi29
        }
        var v30 int8 = 70
        _ = fi29
        _ = v30
    }
    v31 := GenS1{F0: 42, F1: "hello", F2: ""}
    switch "hello" {
    case "\n\t":
        if len(v28) >= 2 {
            v28 = v28[1:len(v28)]
        }
    case "hello":
        v1 = (v1 & (v1 * (^(v1))))
    default:
        if v13 {
            var v32 uint16 = 0
            v2 -= (v2 - (v2 + (v2 | (v2 % ((578) | 1)))))
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        if v13 {
            fmt.Println(fi33)
            switch ("hello" + "foo bar") {
            case "":
                if v13 {
                    v34 := GenS6{F0: 2, F1: fi33}
                    for fi35 := 0; fi35 < 1; fi35++ {
                        var v36 int8 = -128
                        _ = fi35
                        _ = v36
                    }
                    _ = v34
                }
            case "a":
                v13 = ((false || (!(!(637 <= 1)))) || false)
            case "foo bar":
                v28 = append(v28, v2)
            default:
                if (!v13) {
                    if ((0 == 2) || v13) {
                        v37 := GenS2{F0: 570, F1: 0, F2: 4294967295}
                        for fi38 := 0; fi38 < 5; fi38++ {
                            fmt.Println(fi33, v13)
                            v39 := GenS1{F0: int32(v4), F1: "a", F2: ("\n\t" + "hello")}
                            _ = fi38
                            _ = v39
                        }
                        _ = v37
                    }
                    v2 = 4294967295
                    if (!v13) {
                        var v40 string = "foo bar"
                        _ = v40
                    } else {
                        if (true && v13) {
                            v41 := GenS5{F0: -1}
                            var v42 uint32 = 42
                            if len(v28) > 0 {
                                v28[0] = (v2 | 305)
                            }
                            _ = v41
                            _ = v42
                        } else {
                            var v43 int = (fi33 << ((69) & 63))
                            for fi44 := 0; fi44 < 12; fi44++ {
                                if len(v28) > 0 {
                                    v28[0] = (v43 | fi33)
                                }
                                var v45 bool = (-2.25 >= 1.5)
                                if (false && ((-1 >= 9223372036854775807) || false)) {
                                    if false {
                                        var v46 int32 = -945
                                        _ = v46
                                    }
                                } else {
                                    v45 = (42 > 42)
                                }
                                _ = fi44
                                _ = v45
                            }
                            switch ("foo bar" + "") {
                            case "hello":
                                switch fi33 {
                                case -817:
                                    if v13 {
                                        if v13 {
                                            v47 := GenS2{F0: 0, F1: 84, F2: 65535}
                                            switch ((("hello" + ("" + ("a" + (("hello" + "foo bar") + "")))) + "hello") + "foo bar") {
                                            case "foo bar":
                                                if v13 {
                                                    for fi48 := 0; fi48 < 13; fi48++ {
                                                        v49 := GenS5{F0: 678}
                                                        _ = fi48
                                                        _ = v49
                                                    }
                                                    var v50 string = "a"
                                                    for fi51 := 0; fi51 < 18; fi51++ {
                                                        switch (-(fi33)) {
                                                        case 894:
                                                            var v52 int = int(v4)
                                                            _ = v52
                                                        case 850:
                                                            v43 = fi51
                                                        }
                                                        _ = fi51
                                                    }
                                                    _ = v50
                                                } else {
                                                    for fi53 := 0; fi53 < 1; fi53++ {
                                                        switch fi33 {
                                                        case -9223372036854775808:
                                                            switch (-(fi33)) {
                                                            case -575:
                                                                v4 += 272
                                                            default:
                                                                fmt.Println(fi53, v4, v43)
                                                            }
                                                        case 1:
                                                            switch (("hello" + (("foo bar" + "\n\t") + "foo bar")) + ("hello" + "\n\t")) {
                                                            case "a":
                                                                _ = len(v28)
                                                            case "hello":
                                                                v4 -= (v4 >> ((1) & 63))
                                                            default:
                                                                v4 = (v4 ^ 65535)
                                                            }
                                                        case -504:
                                                            v54 := GenS6{F0: 0, F1: (fi53 >> ((235) & 63))}
                                                            _ = v54
                                                        }
                                                        var v55 string = ("" + (("a" + "hello") + "foo bar"))
                                                        _ = fi53
                                                        _ = v55
                                                    }
                                                    var v56 int32 = -311
                                                    _ = v56
                                                }
                                            default:
                                                if ((fi33 | (fi33 >> ((498) & 63))) >= 32768) {
                                                    fmt.Println(v43, v3, fi33)
                                                }
                                            }
                                            _ = v47
                                        }
                                    }
                                case 128:
                                    v57 := GenS5{F0: -887}
                                    _ = v57
                                default:
                                    var v58 float64 = v3
                                    _ = v58
                                }
                            case "\n\t":
                                var v59 bool = v13
                                _ = v59
                            case "foo bar":
                                if len(v28) >= 2 {
                                    v28 = v28[1:len(v28)]
                                }
                            default:
                                for fi60 := 0; fi60 < 9; fi60++ {
                                    if len(v28) > 0 {
                                        v28[0] = (v2 - fi33)
                                    }
                                    v61 := GenS2{F0: 1, F1: 42, F2: 4294967295}
                                    _ = fi60
                                    _ = v61
                                }
                            }
                            _ = v43
                        }
                        v62 := GenS3{F0: "\n\t", F1: -715, F2: 0}
                        var v63 uint32 = 648
                        _ = v62
                        _ = v63
                    }
                }
            }
            if true {
                var v64 int32 = 0
                _ = v64
            } else {
                v65 := GenS7{F0: -112, F1: (-73 <= 127)}
                _ = v65
            }
        }
        v28 = append(v28, (fi33 << ((42) & 63)))
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
