package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint16; F1 bool; F2 int32 }
type GenS5 struct { F0 uint64; F1 int64; F2 int }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int; F1 uint64 }
func genFn0(p0 int16, p1 uint32) bool {
    v7 := GenS1{F0: -1, F1: (^(p1)), F2: (p0 >> ((255) & 63))}
    fmt.Println(p0, p1)
    p1 = 1
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return (-484 > -2147483648)
}
func genFn1() int16 {
    if (42 > 814) {
        v8 := GenS2{F0: 42}
        if true {
            switch ("\n\t" + "foo bar") {
            case "":
                var v9 bool = true
                _ = v9
            case "\n\t":
                var v10 float32 = 3.14159
                _ = v10
            case "foo bar":
                var v11 int16 = -1
                _ = v11
            default:
                for fi12 := 0; fi12 < 15; fi12++ {
                    switch fi12 {
                    case 127:
                        if false {
                            if false {
                                fmt.Println(fi12)
                                v13 := []int{(fi12 % ((fi12) | 1)), (fi12 / ((4294967295) | 1)), -9223372036854775808, 32768}
                                var v14 uint32 = 2147483647
                                _ = v13
                                _ = v14
                            } else {
                                var v15 uint8 = 0
                                _ = v15
                            }
                        } else {
                            if true {
                                if false {
                                    var v16 uint8 = 255
                                    _ = v16
                                }
                                v17 := GenS3{F0: 0, F1: 0}
                                switch "foo bar" {
                                case "a":
                                    var v18 int8 = 89
                                    _ = v18
                                case "hello":
                                    if (true && (!false)) {
                                        var v19 string = "\n\t"
                                        for fi20 := 0; fi20 < 18; fi20++ {
                                            switch (("a" + "hello") + (("\n\t" + (v19 + ("\n\t" + v19))) + ((v19 + v19) + v19))) {
                                            case "\n\t":
                                                var v21 uint64 = 18446744073709551615
                                                _ = v21
                                            case "hello":
                                                fmt.Println(fi12, fi20)
                                            }
                                            _ = fi20
                                        }
                                        if true {
                                            for fi22 := 0; fi22 < 3; fi22++ {
                                                if true {
                                                    fmt.Println(v19)
                                                } else {
                                                    if true {
                                                        var v23 int32 = 1
                                                        _ = v23
                                                    } else {
                                                        v24 := GenS2{F0: 0}
                                                        switch (fi22 ^ (-(fi22))) {
                                                        case 127:
                                                            v19 = ((("a" + (v19 + ("foo bar" + ("\n\t" + v19)))) + v19) + v19)
                                                        default:
                                                            switch v19 {
                                                            case "\n\t":
                                                                v25 := GenS4{F0: 1, F1: true, F2: -1}
                                                                _ = v25
                                                            case "foo bar":
                                                                var v26 int32 = 0
                                                                _ = v26
                                                            case "":
                                                                var v27 int32 = 0
                                                                _ = v27
                                                            default:
                                                                v28 := []int{42, fi22}
                                                                _ = v28
                                                            }
                                                        }
                                                        v19 = v19
                                                        _ = v24
                                                    }
                                                    switch (fi12 & fi22) {
                                                    case 2147483647:
                                                        v19 = "hello"
                                                    case -2147483648:
                                                        var v29 uint16 = 0
                                                        _ = v29
                                                    }
                                                }
                                                v19 = (v19 + v19)
                                                _ = fi22
                                            }
                                        } else {
                                            var v30 int = (^(fi12))
                                            var v31 int8 = 1
                                            _ = v30
                                            _ = v31
                                        }
                                        _ = v19
                                    } else {
                                        var v32 int = (fi12 + (fi12 + 128))
                                        _ = v32
                                    }
                                }
                                _ = v17
                            }
                        }
                    case 32767:
                        var v33 int64 = -1
                        _ = v33
                    case 1:
                        fmt.Println(fi12)
                    default:
                        switch fi12 {
                        case -9223372036854775808:
                            if (4294967295 >= 332) {
                                v34 := GenS2{F0: 820}
                                for fi35 := 0; fi35 < 11; fi35++ {
                                    v36 := []int{(^(fi12)), fi35, (-(fi12))}
                                    _ = fi35
                                    _ = v36
                                }
                                _ = v34
                            }
                        default:
                            v37 := GenS5{F0: 1, F1: -549, F2: 441}
                            _ = v37
                        }
                    }
                    _ = fi12
                }
            }
            var v38 float32 = float32((1e10 - (1e10 / -2.25)))
            v39 := []int{}
            _ = v38
            _ = v39
        }
        v40 := []int{-339, 0, -2147483648}
        _ = v40
        _ = v8
    } else {
        var v41 int64 = 898
        _ = v41
    }
    v42 := GenS6{F0: "foo bar"}
    var v43 float64 = 1e10
    var v44 string = "foo bar"
    _ = v42
    _ = v43
    _ = v44
    return 377
}

func main() {
    var v1 uint64 = 60
    var v2 int = 42
    var v3 int64 = 1
    var v4 int32 = -2147483648
    v5 := make(map[string]bool)
    var v6 int = -1
    var v45 int16 = 1000
    v2 -= (v2 / ((32768) | 1))
    var v46 uint8 = 42
    for fi47 := 0; fi47 < 14; fi47++ {
        if (false || false) {
            switch (v2 - 0) {
            case -9223372036854775808:
                var v48 int = fi47
                _ = v48
            default:
                for fi49 := 0; fi49 < 14; fi49++ {
                    fmt.Println(v1, v4, v4)
                    _ = fi49
                }
            }
        } else {
            v4 = (v4 | (v4 - (v4 | v4)))
            if true {
                var v50 int8 = 0
                _ = v50
            } else {
                fmt.Println(fi47)
            }
        }
        fmt.Println(v1)
        _ = fi47
    }
    v5["foo bar"] = true
    for fi51 := 0; fi51 < 16; fi51++ {
        v4 = v4
        _ = fi51
    }
    v52 := GenS3{F0: 42, F1: (v46 >> ((18446744073709551615) & 63))}
    v2 = v2
    v53 := v5[((("\n\t" + "\n\t") + "a") + "hello")]
    var v54 int8 = -128
    if v53 {
        switch "hello" {
        case "hello":
            if false {
                if (!((^(v1)) > (v1 + v1))) {
                    for fi55 := 0; fi55 < 12; fi55++ {
                        v56 := []int{}
                        for fi57 := 0; fi57 < 1; fi57++ {
                            v58 := GenS7{F0: v2, F1: (v1 + (v1 * (v1 >> ((640) & 63))))}
                            _ = fi57
                            _ = v58
                        }
                        v59 := GenS6{F0: (("\n\t" + ((("\n\t" + ("foo bar" + ("" + ((("" + "") + "hello") + "foo bar")))) + "a") + "hello")) + "")}
                        _ = fi55
                        _ = v56
                        _ = v59
                    }
                    var v60 int16 = -490
                    fmt.Println(v45, v3)
                    _ = v60
                }
                if v53 {
                    fmt.Println(v46)
                    if (43 > 275) {
                        switch "a" {
                        case "foo bar":
                            switch (("" + ("" + "hello")) + "foo bar") {
                            case "foo bar":
                                v45 += (v45 & genFn1())
                            case "":
                                var v61 uint32 = 2
                                _ = v61
                            default:
                                v62 := GenS2{F0: 150}
                                _ = v62
                            }
                        }
                    } else {
                        v63 := []int{}
                        v54 -= (v54 ^ (^(v54)))
                        _ = v63
                    }
                } else {
                    fmt.Println(v54, v46)
                    v1 -= v1
                }
            } else {
                v6 -= v6
                v2 = v2
            }
        case "":
            var v64 uint8 = v46
            _ = v64
        }
    }
    v65 := []int{v2}
    v66 := GenS1{F0: v3, F1: 2147483648, F2: 42}
    var v67 uint8 = v46
    if ((^(v1)) == (^(v1))) {
        var v68 int32 = (v4 - int32(v46))
        var v69 float64 = -2.25
        _ = v68
        _ = v69
    } else {
        v70 := v5["hello"]
        v71 := GenS5{F0: 42, F1: (v3 << ((1) & 63)), F2: (v2 >> ((255) & 63))}
        if len(v65) > 0 {
            v65[0] = -9223372036854775808
        }
        _ = v70
        _ = v71
    }
    var v72 int16 = 687
    var v73 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v65:", len(v65), v65)
    fmt.Println("v67:", v67)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v6
    _ = v65
    _ = v66
    _ = v67
    _ = v72
    _ = v73
}
