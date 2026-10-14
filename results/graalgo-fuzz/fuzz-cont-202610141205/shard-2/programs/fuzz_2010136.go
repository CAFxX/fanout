package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint32; F1 uint; F2 int64 }
func genFn0(p0 uint64, p1 uint) int {
    v4 := GenS1{F0: true}
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 uint = 1
        v7 := GenS1{F0: (!false)}
        v6 = (p1 >> ((v6) & 63))
        _ = fi5
        _ = v6
        _ = v7
    }
    if true {
        var v8 uint = (p1 | p1)
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v4
    return -9223372036854775808
}

func main() {
    var v1 string = (("a" + ("hello" + ("hello" + "a"))) + (("hello" + "a") + "foo bar"))
    v2 := []uint32{135, 14, 338, 1}
    v3 := make(map[int]uint32)
    var v9 int8 = 100
    fmt.Println(v1)
    switch 167 {
    case 32767:
        if len(v2) > 0 {
            v2[0] = 2
        }
    case 128:
        v9 = v9
    default:
        var v10 uint64 = 0
        _ = v10
    }
    v9 += 0
    v9 = v9
    var v11 int16 = 1
    _ = len(v2)
    if false {
        if true {
            switch 88 {
            case 902:
                switch v1 {
                case "\n\t":
                    v1 = ""
                case "":
                    for fi12 := 0; fi12 < 5; fi12++ {
                        v11 -= (v11 + v11)
                        _ = fi12
                    }
                default:
                    var v13 float32 = 3.14159
                    _ = v13
                }
            }
        }
    }
    v14 := GenS1{F0: (!false)}
    if false {
        switch -537 {
        case -205:
            var v15 int32 = 2147483647
            _ = v15
        case 32767:
            switch -504 {
            case 874:
                v16 := GenS1{F0: (!(42 > v11))}
                _ = v16
            default:
                if true {
                    var v17 uint64 = 1
                    var v18 string = v1
                    _ = v17
                    _ = v18
                } else {
                    var v19 int8 = v9
                    fmt.Println(v19, v1)
                    _ = v19
                }
            }
        }
        v20 := GenS1{F0: true}
        v21 := GenS1{F0: false}
        _ = v20
        _ = v21
    }
    v22 := GenS2{F0: (!(1 != v9)), F1: 42}
    if ((((-1 == v9) && false) || true) || false) {
        for fi23 := 0; fi23 < 16; fi23++ {
            if ((v9 & (^(v9))) == 127) {
                v1 = v1
                var v24 uint64 = 4294967295
                _ = v24
            } else {
                switch (fi23 * fi23) {
                case 32767:
                    v9 -= v9
                case 128:
                    if false {
                        for fi25 := 0; fi25 < 16; fi25++ {
                            if false {
                                fmt.Println(v9, v11, fi23)
                                v9 = v9
                                v9 = (v9 / (((v9 / ((v9) | 1))) | 1))
                                v26 := GenS1{F0: (true || false)}
                                v1 = ((v1 + v1) + v1)
                                _ = v26
                            }
                            _ = fi25
                        }
                        var v27 uint64 = 0
                        _ = v27
                    } else {
                        switch v1 {
                        case "\n\t":
                            var v28 uint16 = 0
                            _ = v28
                        case "":
                            fmt.Println(fi23, v11, v9)
                        default:
                            v29 := v3[0]
                            _ = v29
                        }
                    }
                case -9223372036854775808:
                    v30 := GenS2{F0: ((!true) && false), F1: (^(v11))}
                    _ = v30
                }
            }
            switch (fi23 % ((-218) | 1)) {
            case -2147483648:
                switch fi23 {
                case -863:
                    fmt.Println(v11, v11, v11)
                case 4294967295:
                    v31 := GenS1{F0: false}
                    _ = v31
                case 32768:
                    if (1e10 <= 0.0) {
                        v1 = ""
                        for fi32 := 0; fi32 < 19; fi32++ {
                            var v33 int = genFn0(235, 4294967295)
                            v1 = ("hello" + ("" + (v1 + "a")))
                            _ = fi32
                            _ = v33
                        }
                        var v34 bool = ((^(v11)) != v11)
                        _ = v34
                    }
                }
            default:
                for fi35 := 0; fi35 < 19; fi35++ {
                    if (127 == 42) {
                        v36 := GenS2{F0: (!(!true)), F1: v11}
                        fmt.Println(fi23)
                        _ = v36
                    }
                    v37 := GenS3{F0: (fi23 / (((fi35 >> ((2) & 63))) | 1)), F1: 4294967295}
                    if (!(!(true || false))) {
                        if (1 >= 2) {
                            if true {
                                fmt.Println(fi23, v9, v9)
                                var v38 int64 = -1
                                _ = v38
                            }
                            if true {
                                v39 := GenS1{F0: true}
                                _ = v39
                            }
                        } else {
                            var v40 float64 = (float64(-2.25) + 3.14159)
                            if (false || true) {
                                var v41 uint8 = 19
                                v40 = float64(1e10)
                                _ = v41
                            } else {
                                fmt.Println(fi23, fi35, v40)
                                if (!(1 > 323)) {
                                    v42 := GenS2{F0: false, F1: (v11 * (v11 + 1))}
                                    v43 := v3[fi23]
                                    _ = v42
                                    _ = v43
                                }
                                v2 = append(v2, 2147483648)
                                if true {
                                    v3[(fi35 & 42)] = 2147483647
                                    var v44 int32 = 1
                                    if (9223372036854775807 != 1) {
                                        v11 = v11
                                        switch (fi23 - (fi23 / ((int(v9)) | 1))) {
                                        case 0:
                                            var v45 uint = 612
                                            _ = v45
                                        case 127:
                                            var v46 int8 = v9
                                            _ = v46
                                        default:
                                            if true {
                                                switch (fi23 + genFn0(837, 441)) {
                                                case 32768:
                                                    fmt.Println(v9, v44, v44)
                                                }
                                            }
                                        }
                                        if true {
                                            if (false || true) {
                                                v44 = v44
                                                v9 -= (v9 & (v9 >> ((1) & 63)))
                                                switch fi35 {
                                                case -9223372036854775808:
                                                    v40 = v40
                                                case 127:
                                                    v1 = v1
                                                }
                                            }
                                            var v47 int16 = (v11 >> ((65535) & 63))
                                            _ = v47
                                        } else {
                                            var v48 int16 = v11
                                            _ = v48
                                        }
                                    } else {
                                        var v49 float64 = v40
                                        var v50 int16 = 42
                                        _ = v49
                                        _ = v50
                                    }
                                    _ = v44
                                } else {
                                    v11 = v11
                                    v51 := GenS4{F0: 42, F1: 65535, F2: -1}
                                    _ = v51
                                }
                                v1 = "a"
                                if false {
                                    var v52 int16 = v11
                                    if (!(1 == -2147483648)) {
                                        switch ((v1 + "hello") + "a") {
                                        case "hello":
                                            switch "" {
                                            case "":
                                                v40 += v40
                                            }
                                        case "a":
                                            if false {
                                                var v53 float64 = 0.1
                                                fmt.Println(v40, fi23)
                                                _ = v53
                                            }
                                        default:
                                            var v54 uint8 = 2
                                            _ = v54
                                        }
                                    } else {
                                        v9 -= v9
                                    }
                                    _ = v52
                                } else {
                                    if true {
                                        fmt.Println(v1, v1)
                                    }
                                    var v55 float32 = 3.14159
                                    var v56 int8 = 100
                                    _ = v55
                                    _ = v56
                                }
                                _ = len(v3)
                                v57 := v3[2147483647]
                                v2 = append(v2, (v57 | v57))
                                v58 := GenS2{F0: (1e10 <= (0.0 - 1.5)), F1: -32768}
                                _ = v57
                                _ = v58
                            }
                            _ = v40
                        }
                    } else {
                        var v59 bool = false
                        var v60 int = 921
                        _ = v59
                        _ = v60
                    }
                    _ = fi35
                    _ = v37
                }
            }
            for fi61 := 0; fi61 < 16; fi61++ {
                var v62 uint16 = 161
                var v63 uint = 255
                var v64 uint16 = 1
                _ = fi61
                _ = v62
                _ = v63
                _ = v64
            }
            _ = fi23
        }
    } else {
        fmt.Println(v11, v9)
        fmt.Println(v9)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v9
}
