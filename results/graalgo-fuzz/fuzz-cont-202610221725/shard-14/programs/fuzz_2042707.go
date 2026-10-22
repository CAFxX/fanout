package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS6 struct { F0 int32; F1 int8; F2 int16 }
func genFn0(p0 uint32, p1 string) string {
    var v5 uint = 1
    if true {
        fmt.Println(p1, p0, v5)
        v6 := GenS1{F0: (p1 + "foo bar"), F1: 65535, F2: 127}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := GenS2{F0: (v5 << ((18446744073709551615) & 63))}
            if true {
                p0 -= p0
                var v9 float32 = 0.0
                fmt.Println(p0, v5, fi7)
                _ = v9
            } else {
                p0 = p0
            }
            _ = fi7
            _ = v8
        }
    }
    v10 := []int{}
    v5 = (v5 / (((v5 - (v5 * v5))) | 1))
    if (1 >= 255) {
        v11 := GenS1{F0: ((p1 + (p1 + (p1 + ((p1 + "hello") + "\n\t")))) + ((((p1 + "") + p1) + "a") + "\n\t")), F1: 65535, F2: 127}
        _ = v11
    } else {
        v12 := GenS3{F0: 65535, F1: 433}
        var v13 uint16 = 32767
        switch p1 {
        case "a":
            v14 := GenS3{F0: v13, F1: (v13 / ((42) | 1))}
            _ = v14
        default:
            var v15 int32 = 2147483647
            _ = v15
        }
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    return (p1 + ("a" + ""))
}

func main() {
    v1 := []string{"a", ""}
    var v2 uint64 = 278
    v3 := make(map[int]int)
    v4 := []string{"a", "foo bar"}
    switch "hello" {
    case "a":
        for fi16 := 0; fi16 < 3; fi16++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v2 = (v2 * (^(v2)))
            _ = fi16
        }
    case "":
        switch -14 {
        case -605:
            var v17 uint16 = 1
            _ = v17
        default:
            switch -672 {
            case 32768:
                var v18 uint64 = (v2 ^ v2)
                _ = v18
            case -2147483648:
                var v19 uint16 = 255
                _ = v19
            }
        }
    case "foo bar":
        switch (("\n\t" + ("a" + "a")) + (("a" + "\n\t") + "\n\t")) {
        case "a":
            fmt.Println(v2)
        case "hello":
            v20 := GenS2{F0: 432}
            _ = v20
        case "\n\t":
            var v21 int = 32767
            _ = v21
        default:
            v3[2147483647] = -124
        }
    default:
        fmt.Println(v2)
    }
    if (true && (((0 < 12) && true) && (0 < 127))) {
        v2 -= 892
        v2 = 0
        v22 := GenS3{F0: 255, F1: 52}
        _ = v22
    } else {
        v2 = v2
        fmt.Println(v2)
        delete(v3, 876)
    }
    switch 2147483647 {
    case -879:
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 uint16 = 32767
            _ = fi23
            _ = v24
        }
    case -818:
        v2 = v2
    default:
        switch ("foo bar" + "") {
        case "":
            var v25 int16 = -439
            _ = v25
        case "foo bar":
            switch "\n\t" {
            case "hello":
                if true {
                    var v26 float32 = -2.25
                    _ = v26
                } else {
                    v27 := GenS4{F0: 42}
                    _ = v27
                }
            }
        case "\n\t":
            switch -135 {
            case -756:
                delete(v3, 9223372036854775807)
            case -2147483648:
                v28 := GenS2{F0: 484}
                _ = v28
            }
        default:
            if false {
                for fi29 := 0; fi29 < 4; fi29++ {
                    var v30 int8 = -20
                    if (!(65535 >= 65535)) {
                        for fi31 := 0; fi31 < 3; fi31++ {
                            var v32 int = fi31
                            _ = fi31
                            _ = v32
                        }
                        v33 := GenS4{F0: 803}
                        switch ((((("\n\t" + "foo bar") + "\n\t") + "\n\t") + (("\n\t" + "foo bar") + "a")) + "foo bar") {
                        case "\n\t":
                            v34 := GenS5{F0: 1, F1: 825, F2: 747}
                            _ = v34
                        case "a":
                            switch fi29 {
                            case -372:
                                for fi35 := 0; fi35 < 16; fi35++ {
                                    v30 = v30
                                    _ = fi35
                                }
                            case -9223372036854775808:
                                var v36 int32 = int32(v30)
                                _ = v36
                            case 4294967295:
                                if true {
                                    v2 = (v2 - 18446744073709551615)
                                } else {
                                    v30 -= 127
                                }
                            }
                        case "hello":
                            if (!(1.5 > 0.1)) {
                                fmt.Println(fi29, v2, v2)
                                fmt.Println(v30, fi29, v30)
                                var v37 int = fi29
                                _ = v37
                            } else {
                                if (!((42 != 771) || (0.1 > (((1e10 + 1.5) * 1.5) + (-2.25 * 3.14159))))) {
                                    var v38 string = "\n\t"
                                    _ = v38
                                } else {
                                    v2 = v2
                                }
                                v2 = 42
                            }
                        default:
                            if false {
                                var v39 uint = 1
                                _ = v39
                            } else {
                                if (174 < 128) {
                                    var v40 int8 = 39
                                    for fi41 := 0; fi41 < 4; fi41++ {
                                        switch (fi41 % ((-276) | 1)) {
                                        case 32768:
                                            var v42 float32 = (3.14159 * 0.1)
                                            _ = v42
                                        default:
                                            var v43 int = (fi41 * fi41)
                                            _ = v43
                                        }
                                        var v44 uint8 = 255
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = fi41
                                        _ = v44
                                    }
                                    _ = v40
                                } else {
                                    var v45 uint32 = 1
                                    _ = v45
                                }
                                v46 := GenS5{F0: -50, F1: 65535, F2: 946}
                                v2 -= v2
                                _ = v46
                            }
                        }
                        _ = v33
                    }
                    _ = fi29
                    _ = v30
                }
            }
        }
    }
    var v47 uint = 0
    var v48 uint32 = 2147483648
    var v49 bool = true
    var v50 uint = 0
    for fi51 := 0; fi51 < 3; fi51++ {
        for fi52 := 0; fi52 < 10; fi52++ {
            switch fi52 {
            case 0:
                v47 += (v50 >> ((v50) & 63))
            case -2147483648:
                var v53 uint = v47
                _ = v53
            }
            var v54 int16 = 426
            if true {
                if v49 {
                    var v55 string = "hello"
                    v1 = append(v1, v55)
                    if (!true) {
                        v56 := v3[(fi51 >> (((v47 - (v50 | 42))) & 63))]
                        var v57 float32 = ((3.14159 - -2.25) + (0.0 + float32(1e10)))
                        _ = v56
                        _ = v57
                    }
                    _ = v55
                } else {
                    switch (("\n\t" + "hello") + ("\n\t" + ("foo bar" + ("" + ("\n\t" + ""))))) {
                    case "":
                        fmt.Println(v48, v2, v50)
                    case "foo bar":
                        var v58 int16 = (v54 ^ v54)
                        _ = v58
                    case "a":
                        var v59 uint64 = (v2 ^ 42)
                        if (v59 == (^(v59))) {
                            switch "hello" {
                            case "\n\t":
                                v48 = 2
                            }
                        }
                        v60 := GenS3{F0: 32767, F1: 272}
                        _ = v59
                        _ = v60
                    default:
                        var v61 int16 = int16(v50)
                        _ = v61
                    }
                }
                var v62 int8 = -128
                v63 := v3[fi51]
                _ = v62
                _ = v63
            } else {
                v64 := GenS6{F0: 243, F1: -30, F2: v54}
                _ = len(v3)
                var v65 float64 = 1e10
                _ = v64
                _ = v65
            }
            _ = fi52
            _ = v54
        }
        _ = fi51
    }
    v66 := GenS4{F0: 731}
    v47 = (^(v47))
    v47 = 0
    var v67 uint32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v66
    _ = v67
}
