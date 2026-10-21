package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 int16; F2 int64 }
type GenS4 struct { F0 string; F1 int8; F2 bool }
type GenS5 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS6 struct { F0 int; F1 uint16; F2 int }
type GenS7 struct { F0 string; F1 string; F2 uint8 }
type GenS8 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS9 struct { F0 uint8; F1 int64; F2 int32 }
type GenS10 struct { F0 bool; F1 int64 }

func main() {
    var v1 uint64 = 0
    v2 := []string{"", "hello", "hello"}
    v3 := make(map[string]int)
    if (-9223372036854775808 > -118) {
        fmt.Println(v1)
        if (((289 > 255) || ((!true) && (v1 <= v1))) && ((!(271 < 770)) && false)) {
            _ = len(v2)
            v4 := GenS1{F0: true, F1: 143, F2: 597}
            for fi5 := 0; fi5 < 3; fi5++ {
                var v6 uint = 622
                var v7 bool = (0 < 65535)
                _ = fi5
                _ = v6
                _ = v7
            }
            _ = v4
        }
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint = 2
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: false, F1: 125, F2: 4294967295}
    v11 := GenS2{F0: 127}
    _ = len(v2)
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := GenS3{F0: -115, F1: 1, F2: -9223372036854775808}
            if (((-1 == -1) || false) && ("a" <= ("hello" + (("a" + "") + "\n\t")))) {
                _ = len(v3)
                v1 += (v1 + v1)
                if (1e10 != 0.0) {
                    v15 := GenS4{F0: "\n\t", F1: 1, F2: (-101 >= 1)}
                    v16 := GenS3{F0: 42, F1: -1, F2: 1}
                    if false {
                        var v17 float32 = 0.1
                        v17 += v17
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v17
                    } else {
                        v1 = (^(v1))
                    }
                    _ = v15
                    _ = v16
                } else {
                    switch fi13 {
                    case -235:
                        v1 = (v1 | (v1 & 198))
                    case -367:
                        if false {
                            fmt.Println(fi13, v1)
                        }
                    case 678:
                        var v18 int16 = 42
                        _ = v18
                    default:
                        if false {
                            if (true && false) {
                                var v19 uint8 = 1
                                switch ("hello" + ("foo bar" + "a")) {
                                case "foo bar":
                                    var v20 int = (fi13 | fi13)
                                    _ = v20
                                case "a":
                                    if true {
                                        var v21 int16 = 32767
                                        if (true || true) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v22 uint8 = v19
                                            _ = v22
                                        }
                                        v1 += v1
                                        _ = v21
                                    } else {
                                        if ((false && false) && (-434 == 0)) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        } else {
                                            var v23 bool = true
                                            v2 = append(v2, "")
                                            _ = v23
                                        }
                                        switch fi12 {
                                        case 266:
                                            var v24 int8 = 5
                                            _ = v24
                                        case 0:
                                            switch "\n\t" {
                                            case "foo bar":
                                                switch ("hello" + ("foo bar" + "hello")) {
                                                case "hello":
                                                    _ = len(v3)
                                                }
                                            }
                                        case -9223372036854775808:
                                            if true {
                                                var v25 uint32 = uint32(v1)
                                                _ = v25
                                            } else {
                                                v1 += v1
                                                v26 := GenS3{F0: 1, F1: 1, F2: 0}
                                                fmt.Println(v1, v19, v19)
                                                _ = len(v2)
                                                _ = v26
                                            }
                                        }
                                        v19 = 255
                                    }
                                    var v27 string = "foo bar"
                                    if (true && true) {
                                        var v28 float64 = float64(3.14159)
                                        v2 = append(v2, v27)
                                        _ = v28
                                    }
                                    v29 := GenS5{F0: 1000, F1: 2, F2: 4294967295}
                                    switch (-(fi13)) {
                                    case 2147483647:
                                        if (false && false) {
                                            if (928 <= 42) {
                                                v27 = "foo bar"
                                                if true {
                                                    var v30 uint64 = v1
                                                    _ = v30
                                                } else {
                                                    switch "a" {
                                                    case "\n\t":
                                                        v2 = append(v2, ((v27 + v27) + v27))
                                                    }
                                                }
                                                var v31 int16 = -659
                                                _ = v31
                                            }
                                            switch "" {
                                            case "foo bar":
                                                _ = len(v2)
                                            case "hello":
                                                var v32 int = fi12
                                                _ = v32
                                            case "\n\t":
                                                v33 := GenS6{F0: (fi12 + (fi13 % ((4294967295) | 1))), F1: uint16(v19), F2: fi13}
                                                _ = v33
                                            }
                                            v1 += (v1 & v1)
                                            v34 := GenS7{F0: v27, F1: ((v27 + ("foo bar" + "")) + "hello"), F2: (v19 * (v19 ^ (^(v19))))}
                                            if false {
                                                v1 = (v1 / ((v1) | 1))
                                            } else {
                                                fmt.Println(fi12)
                                                v19 = v19
                                            }
                                            _ = v34
                                        } else {
                                            switch (("" + (v27 + "hello")) + (v27 + v27)) {
                                            case "":
                                                var v35 bool = true
                                                _ = v35
                                            case "hello":
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v27 = v27
                                                }
                                            case "foo bar":
                                                fmt.Println(fi12)
                                            }
                                            var v36 uint64 = (v1 - 836)
                                            _ = v36
                                        }
                                    }
                                    v37 := GenS1{F0: true, F1: 0, F2: 1}
                                    if len(v2) > 0 {
                                        v2[0] = (v27 + v27)
                                    }
                                    _ = v27
                                    _ = v29
                                    _ = v37
                                default:
                                    v38 := GenS8{F0: (^(v19)), F1: 4294967295, F2: (v19 / ((255) | 1))}
                                    _ = v38
                                }
                                _ = v19
                            }
                            fmt.Println(fi13, fi13, fi12)
                            var v39 int64 = 694
                            _ = v39
                        } else {
                            v3["hello"] = (fi13 / ((fi13) | 1))
                            fmt.Println(v1, v1)
                            var v40 string = "foo bar"
                            switch ("a" + "") {
                            case "\n\t":
                                v1 += v1
                            case "":
                                var v41 bool = false
                                _ = v41
                            default:
                                _ = len(v2)
                            }
                            if true {
                                v42 := GenS5{F0: -122, F1: 4294967295, F2: 2147483648}
                                v43 := GenS9{F0: 2, F1: 1, F2: 2147483647}
                                _ = v42
                                _ = v43
                            }
                            _ = v40
                        }
                    }
                    v44 := GenS1{F0: true, F1: 42, F2: 1}
                    var v45 int16 = 0
                    _ = v44
                    _ = v45
                }
            } else {
                if (!(0 > 1)) {
                    switch fi12 {
                    case -9223372036854775808:
                        switch (-(fi13)) {
                        case 128:
                            var v46 int32 = -1
                            _ = v46
                        default:
                            v1 -= v1
                        }
                    case -961:
                        v3["\n\t"] = -2147483648
                    case 4294967295:
                        if false {
                            v47 := GenS10{F0: false, F1: -585}
                            _ = v47
                        }
                    }
                    var v48 bool = false
                    v2 = append(v2, ("hello" + "a"))
                    _ = v48
                } else {
                    v49 := GenS8{F0: 0, F1: 4294967295, F2: 127}
                    fmt.Println(v1)
                    switch "foo bar" {
                    case "":
                        if (585 >= 0) {
                            v1 += 682
                            var v50 string = "foo bar"
                            v51 := GenS1{F0: ((!false) && ((0.1 * 1e10) == 1.5)), F1: 42, F2: 908}
                            _ = v50
                            _ = v51
                        } else {
                            var v52 uint16 = 32767
                            _ = v52
                        }
                    case "\n\t":
                        delete(v3, "")
                    case "hello":
                        var v53 float64 = (-2.25 * 0.1)
                        _ = v53
                    }
                    _ = v49
                }
                fmt.Println(fi13, fi12)
                if (196 != 616) {
                    var v54 int16 = 1
                    var v55 bool = false
                    v56 := GenS3{F0: 127, F1: (-(v54)), F2: -9223372036854775808}
                    _ = v54
                    _ = v55
                    _ = v56
                } else {
                    v1 = v1
                    var v57 int = (fi13 | (fi13 & (fi13 >> ((145) & 63))))
                    _ = v57
                }
            }
            fmt.Println(v1, v1, v1)
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    v1 += v1
    for fi58 := 0; fi58 < 18; fi58++ {
        if true {
            v1 += v1
        }
        v1 -= 18446744073709551615
        var v59 bool = true
        _ = fi58
        _ = v59
    }
    var v60 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v60
}
