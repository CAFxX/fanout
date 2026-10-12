package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int8; F1 uint16; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 int16; F2 int32 }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int16; F1 uint32 }
type GenS10 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 uint64 = 482
    var v2 int32 = 2147483647
    v3 := make(map[int]int)
    v1 += (v1 << ((610) & 63))
    for fi4 := 0; fi4 < 2; fi4++ {
        fmt.Println(fi4, v2, v2)
        var v5 uint16 = 1
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 299, F1: "hello"}
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS2{F0: 127, F1: 42, F2: (fi7 | -1)}
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 string = (("" + ("" + (("foo bar" + "\n\t") + "hello"))) + "")
            if true {
                v2 -= (v2 >> ((817) & 63))
            } else {
                v11 := []int{(-(fi7)), fi9, (fi9 >> ((65535) & 63)), (fi9 % ((0) | 1))}
                _ = len(v11)
                v12 := GenS3{F0: 0}
                _ = v11
                _ = v12
            }
            if ((v1 >> ((384) & 63)) == 28) {
                var v13 uint64 = (v1 | v1)
                v1 -= (^(v13))
                var v14 uint8 = 255
                _ = v13
                _ = v14
            } else {
                v1 = (v1 ^ v1)
            }
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    var v15 int64 = 870
    var v16 uint32 = 0
    if false {
        v17 := GenS4{F0: 2, F1: 100}
        v18 := []int{9223372036854775807}
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: 127, F1: 0, F2: 736}
    if true {
        switch 32767 {
        case 32767:
            if false {
                v20 := []int{-549}
                _ = v20
            } else {
                if false {
                    fmt.Println(v15)
                }
                var v21 float64 = 3.14159
                _ = v21
            }
        }
        v15 = 1
        for fi22 := 0; fi22 < 3; fi22++ {
            v15 = 42
            var v23 int32 = v2
            v24 := GenS4{F0: v1, F1: 0}
            _ = fi22
            _ = v23
            _ = v24
        }
    } else {
        var v25 bool = (3.14159 <= (1e10 + 1.5))
        _ = v25
    }
    v16 -= 1
    v16 -= v16
    v26 := GenS5{F0: true}
    switch 9223372036854775807 {
    case -2147483648:
        v2 = (v2 / (((-(v2))) | 1))
    }
    if ((("a" + "a") + "hello") != ("hello" + "foo bar")) {
        var v27 uint = 95
        for fi28 := 0; fi28 < 4; fi28++ {
            for fi29 := 0; fi29 < 17; fi29++ {
                switch fi29 {
                case 42:
                    var v30 string = ""
                    _ = v30
                case 9223372036854775807:
                    fmt.Println(fi29, v2, v2)
                case 973:
                    if (1 > (v27 - uint(v1))) {
                        if (!(fi28 <= -777)) {
                            fmt.Println(v1, v15, v2)
                            if (!false) {
                                _ = len(v3)
                                fmt.Println(fi29, v15, v16)
                                v27 += (v27 >> (((^(v27))) & 63))
                            }
                            var v31 uint32 = (^(v16))
                            _ = v31
                        }
                        v32 := []int{fi28, fi29, (fi28 / ((fi29) | 1)), (fi29 >> ((v27) & 63))}
                        v33 := GenS3{F0: 42}
                        _ = v32
                        _ = v33
                    }
                default:
                    var v34 uint = v27
                    _ = v34
                }
                var v35 uint64 = v1
                v36 := GenS1{F0: (v16 & v16), F1: "a"}
                _ = fi29
                _ = v35
                _ = v36
            }
            v16 += (^(v16))
            switch ("foo bar" + ("a" + "a")) {
            case "foo bar":
                var v37 float32 = (-2.25 + (1e10 + 3.14159))
                _ = v37
            case "":
                v38 := GenS3{F0: 255}
                _ = v38
            case "\n\t":
                v39 := GenS2{F0: 100, F1: 65535, F2: fi28}
                _ = v39
            }
            _ = fi28
        }
        v27 += (v27 & (v27 + v27))
        _ = v27
    }
    fmt.Println(v2)
    switch "hello" {
    case "foo bar":
        var v40 uint16 = 0
        _ = v40
    case "":
        var v41 int8 = 127
        _ = v41
    case "\n\t":
        if false {
            for fi42 := 0; fi42 < 14; fi42++ {
                v43 := GenS2{F0: 42, F1: 65535, F2: fi42}
                v44 := GenS6{F0: 213, F1: -32768, F2: v2}
                if false {
                    for fi45 := 0; fi45 < 3; fi45++ {
                        var v46 uint64 = (v1 | 18446744073709551615)
                        if (!(!false)) {
                            fmt.Println(v1, v2, v16)
                            if (!(0 < -128)) {
                                var v47 int8 = 1
                                var v48 int = (fi45 + fi42)
                                fmt.Println(v2)
                                _ = v47
                                _ = v48
                            } else {
                                v15 = v15
                            }
                            var v49 string = ("\n\t" + "a")
                            _ = v49
                        } else {
                            var v50 int64 = v15
                            v51 := GenS7{F0: (v50 << ((562) & 63)), F1: int8(v1)}
                            _ = v50
                            _ = v51
                        }
                        if ((!(true || (((0.0 - (0.0 * -2.25)) < float32(((((-2.25 / 1.5) - (((((1e10 + (0.1 - (1.5 + 1e10))) * 0.0) + 0.0) / 1.0) + (1.5 * 3.14159))) / (3.14159 + (0.0 - ((1e10 + 0.0) + 0.0)))) / 1.0))) && false))) || true) {
                            v52 := GenS4{F0: v1, F1: 127}
                            if (v46 <= v1) {
                                var v53 uint8 = 1
                                var v54 string = ("" + ("foo bar" + ""))
                                switch v54 {
                                case "hello":
                                    var v55 bool = (!(2147483647 >= v16))
                                    _ = v55
                                }
                                _ = v53
                                _ = v54
                            }
                            v46 -= (v1 % (((v1 % ((v1) | 1))) | 1))
                            _ = v52
                        } else {
                            v16 = (v16 + (v16 | 1))
                        }
                        _ = fi45
                        _ = v46
                    }
                } else {
                    v1 = v1
                    v3[fi42] = int(v2)
                }
                _ = fi42
                _ = v43
                _ = v44
            }
            v56 := GenS8{F0: int32(v15)}
            _ = v56
        } else {
            if true {
                if true {
                    v57 := v3[9223372036854775807]
                    v58 := GenS9{F0: 1, F1: v16}
                    _ = v57
                    _ = v58
                } else {
                    switch "" {
                    case "a":
                        if false {
                            var v59 bool = true
                            _ = v59
                        }
                    default:
                        var v60 string = ""
                        _ = v60
                    }
                }
                if false {
                    for fi61 := 0; fi61 < 3; fi61++ {
                        v16 = 0
                        _ = fi61
                    }
                    if false {
                        v16 -= (v16 & (v16 >> ((4294967295) & 63)))
                        v16 = 65535
                    }
                    v1 -= (v1 & (^(v1)))
                }
                var v62 uint = 42
                _ = v62
            } else {
                var v63 uint32 = v16
                var v64 float32 = 1e10
                _ = v63
                _ = v64
            }
        }
    default:
        if false {
            for fi65 := 0; fi65 < 1; fi65++ {
                fmt.Println(v1)
                var v66 uint16 = 445
                v67 := GenS10{F0: v1, F1: (v16 | (v16 | v16))}
                _ = fi65
                _ = v66
                _ = v67
            }
            v68 := GenS10{F0: (^(v1)), F1: (v16 / ((v16) | 1))}
            if true {
                var v69 float32 = (0.1 * -2.25)
                var v70 string = ""
                switch v70 {
                case "a":
                    if true {
                        v71 := []int{}
                        fmt.Println(v16, v15, v69)
                        _ = v71
                    } else {
                        var v72 int8 = 42
                        _ = v72
                    }
                case "hello":
                    v1 -= (v1 % (((v1 - v1)) | 1))
                }
                _ = v69
                _ = v70
            }
            _ = v68
        } else {
            v1 -= v1
        }
    }
    v1 = 691
    v15 = (-(v15))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v6
}
