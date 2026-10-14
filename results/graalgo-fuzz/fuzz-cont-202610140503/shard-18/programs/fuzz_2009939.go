package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int64 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32; F1 string }
type GenS7 struct { F0 uint8; F1 int }
type GenS8 struct { F0 uint32; F1 bool; F2 int32 }
type GenS9 struct { F0 uint64; F1 int64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 bool }

func main() {
    v1 := make(map[int]int)
    var v2 uint = 255
    v3 := make(map[string]bool)
    v4 := []int{}
    var v5 int32 = 42
    v6 := GenS1{F0: 32767, F1: (v2 >> (((v2 >> ((65535) & 63))) & 63)), F2: -1}
    if ((-1 > 127) || true) {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 uint16 = 65535
            _ = fi7
            _ = v8
        }
    } else {
        var v9 int64 = 42
        if (!true) {
            var v10 float64 = -2.25
            if (!true) {
                fmt.Println(v5, v9, v9)
                v11 := GenS2{F0: v9, F1: (!(((3.14159 + (v10 + v10)) / float64(1e10)) <= float64(((0.1 * (0.0 - float32(v10))) / ((1e10 / 1.0) + 3.14159)))))}
                _ = v11
            } else {
                var v12 int = 42
                switch v12 {
                case 42:
                    var v13 int16 = 32767
                    _ = v13
                case 0:
                    for fi14 := 0; fi14 < 7; fi14++ {
                        if true {
                            var v15 uint64 = 1
                            var v16 float64 = (v10 / ((v10 - v10) * v10))
                            v17 := GenS2{F0: (v9 / ((v9) | 1)), F1: true}
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        if (4294967295 < (v2 >> (((^(v2))) & 63))) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            fmt.Println(v5)
                            if false {
                                for fi18 := 0; fi18 < 6; fi18++ {
                                    var v19 bool = false
                                    var v20 int8 = 42
                                    if (((((-32768 <= -333) || v19) || (((!((v2 << ((v2) & 63)) >= v2)) || (!("\n\t" <= "\n\t"))) && false)) && (!(3.14159 <= 0.1))) || false) {
                                        switch (^(v12)) {
                                        case 659:
                                            if false {
                                                var v21 uint16 = 217
                                                var v22 uint16 = (^(v21))
                                                _ = v21
                                                _ = v22
                                            } else {
                                                var v23 string = ("hello" + "hello")
                                                v4 = append(v4, 32768)
                                                if (fi14 == fi18) {
                                                    fmt.Println(v12)
                                                } else {
                                                    if v19 {
                                                        v19 = (v23 >= "\n\t")
                                                        var v24 uint8 = 255
                                                        _ = v24
                                                    } else {
                                                        var v25 uint32 = 102
                                                        v12 += fi14
                                                        v26 := v1[(fi14 + (v12 & fi14))]
                                                        _ = v25
                                                        _ = v26
                                                    }
                                                    v9 = v9
                                                }
                                                var v27 float32 = 0.1
                                                if v19 {
                                                    fmt.Println(v27, v20, v5)
                                                    if v19 {
                                                        v19 = true
                                                        fmt.Println(v27, fi14)
                                                        var v28 float32 = v27
                                                        _ = v28
                                                    } else {
                                                        var v29 uint8 = 127
                                                        _ = v29
                                                    }
                                                    v2 += (v2 << (((v2 - (v2 | v2))) & 63))
                                                }
                                                var v30 uint16 = 65535
                                                _ = v23
                                                _ = v27
                                                _ = v30
                                            }
                                        default:
                                            switch "a" {
                                            case "":
                                                if v19 {
                                                    switch ("\n\t" + "hello") {
                                                    case "\n\t":
                                                        v31 := GenS2{F0: (v9 ^ (-(v9))), F1: true}
                                                        _ = v31
                                                    }
                                                    fmt.Println(v2, fi14)
                                                } else {
                                                    fmt.Println(v9, v19, v10)
                                                    switch int(v9) {
                                                    case 128:
                                                        var v32 uint8 = 1
                                                        _ = v32
                                                    default:
                                                        v33 := GenS3{F0: ((v20 << ((v2) & 63)) > v20)}
                                                        _ = len(v3)
                                                        _ = v33
                                                    }
                                                    var v34 uint64 = 300
                                                    _ = v34
                                                }
                                            case "a":
                                                var v35 uint32 = 0
                                                _ = v35
                                            default:
                                                if ((!false) || v19) {
                                                    if (1 <= fi18) {
                                                        v36 := GenS3{F0: false}
                                                        _ = v36
                                                    } else {
                                                        switch ("\n\t" + "foo bar") {
                                                        case "hello":
                                                            var v37 uint32 = 2147483647
                                                            _ = v37
                                                        default:
                                                            fmt.Println(v20, v19)
                                                            var v38 uint16 = 65535
                                                            if (335 == uint64(v38)) {
                                                                if ((4294967295 == 4294967295) && false) {
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        v39 := v1[(-(v12))]
                                                                        if true {
                                                                            if v19 {
                                                                                v10 += float64(((-2.25 + (0.1 * 0.0)) / 0.1))
                                                                                _ = len(v3)
                                                                                v40 := GenS2{F0: 42, F1: false}
                                                                                _ = v40
                                                                            }
                                                                            var v41 uint = (v2 * (v2 + 42))
                                                                            v42 := GenS3{F0: (v9 < -765)}
                                                                            _ = v41
                                                                            _ = v42
                                                                        }
                                                                        _ = v39
                                                                    }
                                                                } else {
                                                                    var v43 float32 = 1.5
                                                                    var v44 float32 = float32((3.14159 / 3.14159))
                                                                    _ = v43
                                                                    _ = v44
                                                                }
                                                            }
                                                            _ = v38
                                                        }
                                                        v45 := GenS4{F0: (!v19), F1: 1}
                                                        _ = v45
                                                    }
                                                } else {
                                                    v46 := v1[(fi14 + fi14)]
                                                    _ = v46
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    v47 := GenS5{F0: v2}
                                                    _ = v47
                                                case "":
                                                    if ((128 == 127) && (int32(v20) < v5)) {
                                                        v5 = v5
                                                        v19 = v19
                                                        if ("" >= "") {
                                                            var v48 int32 = v5
                                                            v49 := GenS6{F0: v48, F1: ("\n\t" + ("" + ("foo bar" + ("\n\t" + "hello"))))}
                                                            _ = v48
                                                            _ = v49
                                                        } else {
                                                            v50 := GenS2{F0: v9, F1: ((v19 && v19) || v19)}
                                                            _ = v50
                                                        }
                                                    }
                                                case "a":
                                                    if false {
                                                        v12 += 2147483647
                                                        _ = len(v1)
                                                        var v51 uint32 = 2147483648
                                                        _ = v51
                                                    }
                                                }
                                                var v52 bool = ((v2 - (v2 | (^(v2)))) >= 0)
                                                var v53 float64 = 3.14159
                                                var v54 string = "hello"
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                            }
                                        }
                                        v12 += fi18
                                        v19 = v19
                                    } else {
                                        if false {
                                            var v55 uint64 = 424
                                            _ = v55
                                        } else {
                                            v56 := GenS7{F0: 42, F1: v12}
                                            _ = v56
                                        }
                                        if ((-(v5)) <= 1) {
                                            switch "\n\t" {
                                            case "foo bar":
                                                v57 := GenS1{F0: -293, F1: (^(v2)), F2: (v9 % (((v9 / ((-1) | 1))) | 1))}
                                                _ = v57
                                            case "\n\t":
                                                var v58 uint32 = 288
                                                _ = v58
                                            }
                                            v5 = -2147483648
                                        }
                                    }
                                    _ = fi18
                                    _ = v19
                                    _ = v20
                                }
                            }
                        } else {
                            v59 := GenS6{F0: (v5 & int32(v2)), F1: ("\n\t" + "foo bar")}
                            v60 := GenS8{F0: 1, F1: false, F2: v5}
                            _ = v59
                            _ = v60
                        }
                        _ = fi14
                    }
                }
                _ = v12
            }
            _ = v10
        }
        var v61 int = -92
        _ = v61
        _ = v9
    }
    delete(v1, 1)
    var v62 float32 = 1e10
    v63 := GenS9{F0: 18446744073709551615, F1: -9223372036854775808}
    v62 = 1.5
    var v64 bool = false
    if v64 {
        v64 = false
        v2 = v2
    }
    var v65 string = "\n\t"
    for fi66 := 0; fi66 < 8; fi66++ {
        v67 := GenS10{F0: 1000}
        var v68 uint16 = 65535
        for fi69 := 0; fi69 < 16; fi69++ {
            v70 := GenS11{F0: (v2 > (v2 | 807))}
            if false {
                v62 = v62
                _ = len(v4)
                if true {
                    v4 = append(v4, int(v5))
                    v4 = append(v4, (fi69 | (fi66 | fi69)))
                    v2 = (v2 << ((v2) & 63))
                } else {
                    v65 = v65
                    v71 := GenS1{F0: 9, F1: (v2 * v2), F2: -9223372036854775808}
                    _ = len(v4)
                    _ = v71
                }
            } else {
                var v72 uint32 = 2
                v72 -= (^(v72))
                fmt.Println(v65, v62, v64)
                _ = v72
            }
            _ = fi69
            _ = v70
        }
        _ = fi66
        _ = v67
        _ = v68
    }
    if (4294967295 > 2) {
        fmt.Println(v65)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        switch v65 {
        case "hello":
            v73 := GenS1{F0: -1, F1: 0, F2: 642}
            _ = v73
        case "foo bar":
            for fi74 := 0; fi74 < 10; fi74++ {
                for fi75 := 0; fi75 < 4; fi75++ {
                    if ((v5 - v5) <= (v5 - v5)) {
                        v76 := v3[("foo bar" + (v65 + ((v65 + ((v65 + v65) + v65)) + ("hello" + (v65 + "a")))))]
                        _ = v76
                    }
                    v5 += v5
                    _ = fi75
                }
                fmt.Println(v64, v62, v62)
                _ = fi74
            }
        }
        delete(v1, 846)
        for fi77 := 0; fi77 < 4; fi77++ {
            v62 -= float32(3.14159)
            v2 += (^(v2))
            _ = fi77
        }
    }
    v5 += (v5 * 2147483647)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
