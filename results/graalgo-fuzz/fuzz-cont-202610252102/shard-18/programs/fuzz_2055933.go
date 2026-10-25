package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 uint64; F1 int32; F2 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 uint; F1 string }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64; F1 string; F2 int32 }
type GenS10 struct { F0 uint64; F1 int }
func genFn0() uint64 {
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(fi5)
        fmt.Println(fi5)
        var v6 bool = true
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 float32 = 0.0
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 2147483647
    switch -1 {
    case 951:
        var v10 uint64 = 605
        _ = v10
    case 67:
        for fi11 := 0; fi11 < 7; fi11++ {
            v12 := GenS1{F0: 1}
            var v13 int64 = -658
            _ = fi11
            _ = v12
            _ = v13
        }
    case -1:
        var v14 int64 = 9223372036854775807
        _ = v14
    }
    v9 -= (v9 % ((909) | 1))
    _ = v9
    return 42
}

func main() {
    v1 := []string{}
    var v2 uint64 = 0
    v3 := make(map[int]string)
    var v4 bool = true
    if (3.14159 >= (-2.25 - 0.0)) {
        fmt.Println(v2)
        v15 := GenS2{F0: 1, F1: 612}
        _ = v15
    }
    var v16 uint16 = 32767
    var v17 string = "foo bar"
    fmt.Println(v4, v16)
    switch 1 {
    case -1:
        var v18 int64 = 0
        _ = v18
    case 0:
        v17 = ("" + v17)
    default:
        v4 = true
    }
    v4 = (!(4294967295 != 0))
    switch ("a" + "") {
    case "hello":
        v4 = (!(v4 && (!v4)))
    default:
        v16 += 2
    }
    if ((166 == 124) || v4) {
        v19 := GenS2{F0: 42, F1: 2}
        _ = v19
    }
    var v20 uint32 = 538
    for fi21 := 0; fi21 < 13; fi21++ {
        delete(v3, -165)
        v22 := GenS3{F0: 42, F1: -369}
        var v23 uint16 = v16
        _ = fi21
        _ = v22
        _ = v23
    }
    v17 = ("foo bar" + ((("" + (v17 + v17)) + ("" + "a")) + v17))
    if ((!false) || v4) {
        for fi24 := 0; fi24 < 15; fi24++ {
            v1 = append(v1, v17)
            _ = fi24
        }
        switch 9223372036854775807 {
        case 1:
            var v25 int16 = -204
            _ = v25
        case 127:
            for fi26 := 0; fi26 < 18; fi26++ {
                if false {
                    if (fi26 <= (-(fi26))) {
                        for fi27 := 0; fi27 < 1; fi27++ {
                            v20 = (v20 + (v20 / ((2) | 1)))
                            var v28 int16 = 32767
                            v20 = v20
                            v1 = append(v1, "\n\t")
                            if (161 == 42) {
                                if v4 {
                                    var v29 bool = false
                                    switch "a" {
                                    case "hello":
                                        if len(v1) > 0 {
                                            v1[0] = (((v17 + v17) + v17) + v17)
                                        }
                                    case "a":
                                        v28 = (v28 & -32768)
                                    }
                                    _ = v29
                                } else {
                                    fmt.Println(v16)
                                    v2 = 1
                                }
                            }
                            if v4 {
                                var v30 string = ((("foo bar" + v17) + v17) + v17)
                                v28 = (v28 % (((v28 | (v28 - (v28 & (v28 / ((v28) | 1)))))) | 1))
                                switch fi27 {
                                case -9223372036854775808:
                                    if (!v4) {
                                        fmt.Println(v20, v16, v28)
                                        switch "foo bar" {
                                        case "\n\t":
                                            v31 := GenS1{F0: 1}
                                            _ = v31
                                        }
                                        v32 := GenS2{F0: 42, F1: 1}
                                        _ = v32
                                    }
                                case 429:
                                    v2 += v2
                                    v33 := GenS4{F0: (^(v2)), F1: 2147483647, F2: v28}
                                    switch (fi26 % ((fi27) | 1)) {
                                    case 9223372036854775807:
                                        var v34 bool = true
                                        _ = v34
                                    case 32768:
                                        var v35 float64 = (3.14159 - float64(0.0))
                                        _ = v35
                                    default:
                                        var v36 int = (fi26 & 1)
                                        v36 = (v36 / (((v36 * v36)) | 1))
                                        v37 := GenS3{F0: -11, F1: 1}
                                        _ = v36
                                        _ = v37
                                    }
                                    _ = v33
                                default:
                                    fmt.Println(v20)
                                    v30 = v30
                                }
                                v1 = append(v1, ("" + v17))
                                _ = len(v3)
                                _ = v30
                            } else {
                                v4 = (1 <= 42)
                            }
                            fmt.Println(v4, v4)
                            if (true || (739 < v20)) {
                                v20 = 2
                                fmt.Println(v28, v16)
                            }
                            switch -2147483648 {
                            case 32767:
                                if false {
                                    if (v4 && false) {
                                        if v4 {
                                            switch (-(fi26)) {
                                            case 42:
                                                if len(v1) > 0 {
                                                    v1[0] = v17
                                                }
                                            default:
                                                switch "hello" {
                                                case "":
                                                    var v38 float64 = ((((3.14159 * ((float64(1e10) - float64((3.14159 / 1.5))) + 3.14159)) - 3.14159) * 0.0) + (3.14159 - -2.25))
                                                    _ = v38
                                                case "a":
                                                    fmt.Println(v17, v4, fi26)
                                                    v20 += uint32(v2)
                                                    switch v17 {
                                                    case "\n\t":
                                                        if v4 {
                                                            if v4 {
                                                                var v39 int64 = -302
                                                                _ = v39
                                                            }
                                                            v16 -= (v16 - v16)
                                                        }
                                                    case "a":
                                                        v40 := GenS3{F0: -9223372036854775808, F1: 42}
                                                        _ = v40
                                                    case "":
                                                        v28 += v28
                                                    }
                                                case "hello":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v28 = (v28 ^ v28)
                                                    if v4 {
                                                        v41 := GenS3{F0: 9223372036854775807, F1: 42}
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        _ = v41
                                                    }
                                                    v1 = append(v1, v17)
                                                    v42 := GenS5{F0: 0}
                                                    _ = v42
                                                default:
                                                    v43 := GenS1{F0: 132}
                                                    _ = v43
                                                }
                                            }
                                        } else {
                                            var v44 int8 = 0
                                            _ = v44
                                        }
                                        var v45 uint8 = 42
                                        var v46 uint = 42
                                        _ = v45
                                        _ = v46
                                    } else {
                                        switch (v17 + v17) {
                                        case "":
                                            v47 := GenS4{F0: v2, F1: 42, F2: 32767}
                                            _ = v47
                                        case "\n\t":
                                            v4 = v4
                                        }
                                        v48 := GenS3{F0: 9223372036854775807, F1: 0}
                                        if v4 {
                                            var v49 uint8 = 255
                                            var v50 int = fi27
                                            switch 2147483647 {
                                            case 127:
                                                switch ("\n\t" + ((v17 + v17) + "\n\t")) {
                                                case "foo bar":
                                                    var v51 int16 = -32768
                                                    _ = v51
                                                case "a":
                                                    var v52 bool = false
                                                    _ = v52
                                                case "hello":
                                                    v53 := GenS6{F0: v17, F1: (v49 + (^(v49)))}
                                                    _ = v53
                                                default:
                                                    var v54 uint = 255
                                                    _ = v54
                                                }
                                            case 128:
                                                var v55 uint = 4294967295
                                                _ = v55
                                            case -196:
                                                v56 := v3[v50]
                                                _ = v56
                                            }
                                            _ = v49
                                            _ = v50
                                        }
                                        _ = v48
                                    }
                                    var v57 uint32 = v20
                                    v17 = "\n\t"
                                    v57 += v57
                                    _ = v57
                                }
                            case 2147483647:
                                v28 = (v28 | v28)
                            case 128:
                                v58 := v3[1]
                                _ = v58
                            default:
                                v59 := GenS7{F0: 1, F1: v17}
                                _ = v59
                            }
                            _ = fi27
                            _ = v28
                        }
                    } else {
                        var v60 bool = (1e10 < 1e10)
                        v61 := GenS8{F0: 321}
                        _ = v60
                        _ = v61
                    }
                    v62 := GenS7{F0: uint(v20), F1: ""}
                    _ = v62
                } else {
                    if (1 == 276) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if (v20 != v20) {
                            if false {
                                v2 -= uint64(v16)
                                v20 -= uint32(v16)
                            }
                            v1 = append(v1, "")
                        } else {
                            switch fi26 {
                            case 9223372036854775807:
                                var v63 float32 = 1e10
                                _ = v63
                            case 2147483647:
                                fmt.Println(v20, v4)
                            }
                        }
                    } else {
                        var v64 uint16 = 32767
                        _ = v64
                    }
                    v65 := GenS6{F0: ((v17 + "a") + ("foo bar" + ("a" + v17))), F1: uint8(v2)}
                    _ = v65
                }
                _ = fi26
            }
        }
    } else {
        v66 := GenS9{F0: -9223372036854775808, F1: (v17 + v17), F2: 2147483647}
        fmt.Println(v20, v16)
        var v67 int64 = 0
        _ = v66
        _ = v67
    }
    if (float32(0.0) > -2.25) {
        fmt.Println(v20, v20)
    } else {
        v68 := GenS10{F0: v2, F1: 128}
        _ = v68
    }
    v1 = append(v1, ("\n\t" + v17))
    var v69 int64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v69
}
