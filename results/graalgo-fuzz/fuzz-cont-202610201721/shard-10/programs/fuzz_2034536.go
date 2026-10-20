package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int8; F2 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint; F1 bool }
type GenS9 struct { F0 int16; F1 int8 }
func genFn0() uint16 {
    v5 := []int{}
    v5 = append(v5, 0)
    var v6 int8 = 0
    v7 := GenS1{F0: -32768}
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int64 = 42
    v2 := []uint32{698, 4294967295}
    var v3 int32 = 214
    var v4 int64 = (v1 / ((v1) | 1))
    var v8 int32 = v3
    fmt.Println(v4)
    v8 = int32(v4)
    var v9 int8 = -1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3, v4, v9)
    v3 = 42
    var v10 int8 = v9
    fmt.Println(v10, v8, v3)
    v3 -= v3
    v11 := GenS1{F0: 0}
    var v12 int64 = 1
    fmt.Println(v4, v4, v3)
    var v13 float32 = 0.0
    var v14 uint = 255
    var v15 uint8 = 0
    fmt.Println(v10, v8)
    if true {
        if false {
            v10 -= (-(v10))
        } else {
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 uint = 0
                if (-2.25 <= 3.14159) {
                    v13 = float32((-2.25 + (float64((1.5 / v13)) - 0.0)))
                    for fi18 := 0; fi18 < 5; fi18++ {
                        v17 += (v17 | v17)
                        var v19 uint32 = 71
                        var v20 float64 = 0.0
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    v12 -= v1
                }
                var v21 uint64 = 18446744073709551615
                _ = fi16
                _ = v17
                _ = v21
            }
            v1 = v1
            var v22 uint = v14
            _ = v22
        }
        if true {
            var v23 string = "\n\t"
            if false {
                var v24 int32 = v8
                for fi25 := 0; fi25 < 16; fi25++ {
                    v13 = v13
                    if (0 <= 2) {
                        v9 += -1
                        for fi26 := 0; fi26 < 13; fi26++ {
                            if len(v2) > 0 {
                                v2[0] = 2147483648
                            }
                            _ = fi26
                        }
                    } else {
                        var v27 uint32 = 0
                        for fi28 := 0; fi28 < 20; fi28++ {
                            v4 += -9223372036854775808
                            v24 = (v24 - v3)
                            if (1.5 >= (1e10 * 1e10)) {
                                v29 := GenS1{F0: -32768}
                                _ = v29
                            }
                            var v30 uint = (^(v14))
                            if (v12 >= (v12 >> ((v30) & 63))) {
                                v10 = (v9 / (((v9 << ((65535) & 63))) | 1))
                                if ((3.14159 - (-2.25 * 0.0)) <= 1.5) {
                                    if ((((v13 / ((v13 - v13) - 0.0)) - v13) + v13) <= (v13 / 1.5)) {
                                        v10 = int8(v8)
                                        v15 = v15
                                    }
                                    if (!(!(!(2 < 32767)))) {
                                        var v31 int = (fi25 / ((4294967295) | 1))
                                        v23 = v23
                                        _ = v31
                                    } else {
                                        v2 = append(v2, 2)
                                        if ("foo bar" != v23) {
                                            if len(v2) > 0 {
                                                v2[0] = (v27 & v27)
                                            }
                                        } else {
                                            switch fi25 {
                                            case -1:
                                                var v32 uint8 = (^(v15))
                                                _ = v32
                                            case -44:
                                                var v33 int8 = 51
                                                _ = v33
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = (^(v27))
                                            }
                                            v24 -= (-(v8))
                                            v23 = v23
                                            if len(v2) > 0 {
                                                v2[0] = v27
                                            }
                                        }
                                    }
                                } else {
                                    _ = len(v2)
                                    if false {
                                        fmt.Println(v4)
                                        var v34 int16 = 71
                                        var v35 float32 = ((v13 * (1e10 / float32(1.5))) + v13)
                                        if (!true) {
                                            v14 = (^(v30))
                                            fmt.Println(fi25, v3)
                                        }
                                        _ = v34
                                        _ = v35
                                    }
                                    switch fi25 {
                                    case 2147483647:
                                        switch (fi25 << (((v14 + (v14 ^ 138))) & 63)) {
                                        case 32767:
                                            v36 := GenS2{F0: 255}
                                            _ = v36
                                        case -1:
                                            var v37 uint8 = (v15 & (v15 >> (((v14 << (((v14 - v14)) & 63))) & 63)))
                                            _ = v37
                                        default:
                                            fmt.Println(v12, v27)
                                            if true {
                                                switch fi28 {
                                                case -767:
                                                    v38 := GenS3{F0: (v8 ^ (v8 / ((-207) | 1))), F1: v10, F2: 74}
                                                    _ = v38
                                                default:
                                                    if (-1 != 1) {
                                                        v12 = (v1 & -1)
                                                    }
                                                }
                                            } else {
                                                if false {
                                                    var v39 string = ""
                                                    switch fi25 {
                                                    case -2147483648:
                                                        v27 = (^(v27))
                                                    }
                                                    _ = v39
                                                }
                                            }
                                        }
                                    case 128:
                                        switch v23 {
                                        case "foo bar":
                                            var v40 uint64 = 0
                                            _ = v40
                                        case "":
                                            if false {
                                                fmt.Println(v27, v10)
                                                if false {
                                                    if false {
                                                        fmt.Println(v8, v3, v12)
                                                    } else {
                                                        var v41 uint = 2
                                                        v13 = (v13 * ((((v13 * 0.1) - 0.0) / v13) - 3.14159))
                                                        _ = v41
                                                    }
                                                    if true {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        if (false || (v23 > "")) {
                                                            v2 = append(v2, (v27 >> ((v14) & 63)))
                                                        } else {
                                                            v42 := GenS3{F0: (^(v8)), F1: (v10 / (((v10 + v9)) | 1)), F2: -128}
                                                            var v43 float64 = 3.14159
                                                            v2 = append(v2, 0)
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                        v44 := GenS4{F0: (v9 % (((-(v9))) | 1))}
                                                        _ = v44
                                                    }
                                                } else {
                                                    v23 = ((v23 + (v23 + "a")) + (v23 + "a"))
                                                }
                                                var v45 uint64 = 1
                                                _ = v45
                                            } else {
                                                if (!(!true)) {
                                                    switch ((v23 + v23) + ("" + (v23 + (((v23 + (v23 + (v23 + ("hello" + "\n\t")))) + "") + v23)))) {
                                                    case "hello":
                                                        v2 = append(v2, (v27 * (v27 >> (((v30 / (((v30 % ((v14) | 1))) | 1))) & 63))))
                                                    case "a":
                                                        v8 += -356
                                                    default:
                                                        v23 = v23
                                                    }
                                                    v46 := GenS3{F0: (v24 * v24), F1: -114, F2: 1}
                                                    v47 := GenS3{F0: v8, F1: v10, F2: -128}
                                                    _ = v46
                                                    _ = v47
                                                } else {
                                                    var v48 uint64 = 1
                                                    v27 += 4294967295
                                                    if true {
                                                        if true {
                                                            switch fi25 {
                                                            case 42:
                                                                var v49 int64 = 199
                                                                _ = v49
                                                            case -845:
                                                                var v50 bool = (42 >= 477)
                                                                v51 := GenS5{F0: v15, F1: v30}
                                                                _ = v50
                                                                _ = v51
                                                            case 1:
                                                                var v52 uint16 = 255
                                                                _ = v52
                                                            default:
                                                                var v53 bool = false
                                                                _ = v53
                                                            }
                                                        } else {
                                                            v54 := GenS5{F0: v15, F1: 65535}
                                                            var v55 string = v23
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                        v56 := GenS6{F0: -64}
                                                        var v57 int = (fi28 / ((9223372036854775807) | 1))
                                                        v12 -= 42
                                                        if true {
                                                            _ = len(v2)
                                                        } else {
                                                            var v58 uint64 = (v48 / (((v48 >> ((491) & 63))) | 1))
                                                            v59 := GenS7{F0: (v27 & (^(v27)))}
                                                            v48 += (v58 / (((v48 ^ v58)) | 1))
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                        v60 := GenS6{F0: v9}
                                                        _ = v56
                                                        _ = v57
                                                        _ = v60
                                                    }
                                                    _ = v48
                                                }
                                            }
                                        }
                                    }
                                    _ = len(v2)
                                }
                            }
                            v9 -= int8(v8)
                            var v61 uint8 = v15
                            _ = fi28
                            _ = v30
                            _ = v61
                        }
                        v62 := GenS8{F0: v14, F1: (1 != 209)}
                        _ = v27
                        _ = v62
                    }
                    if true {
                        v63 := GenS7{F0: 42}
                        v64 := GenS8{F0: (v14 % (((v14 | (v14 % ((327) | 1)))) | 1)), F1: false}
                        fmt.Println(v24, v15, v3)
                        _ = v63
                        _ = v64
                    }
                    _ = fi25
                }
                _ = v24
            } else {
                var v65 int32 = (v3 | v8)
                if true {
                    if false {
                        if false {
                            fmt.Println(v9, v3)
                            v66 := GenS1{F0: 32767}
                            _ = v66
                        } else {
                            fmt.Println(v14)
                            var v67 float32 = float32((float64(v13) * 0.0))
                            if (0.0 <= (1.5 * v67)) {
                                v9 = (v10 & (v10 + 100))
                                var v68 uint64 = 72
                                _ = v68
                            } else {
                                if (34 >= 1) {
                                    var v69 uint8 = v15
                                    _ = v69
                                } else {
                                    var v70 int8 = 127
                                    var v71 float64 = float64(v67)
                                    v72 := GenS3{F0: (^(v8)), F1: v10, F2: -1}
                                    _ = v70
                                    _ = v71
                                    _ = v72
                                }
                                var v73 float32 = v67
                                v4 += (^(v12))
                                _ = v73
                            }
                            _ = v67
                        }
                    } else {
                        v10 = 127
                    }
                    fmt.Println(v15)
                } else {
                    if false {
                        if len(v2) > 0 {
                            v2[0] = 929
                        }
                    }
                    for fi74 := 0; fi74 < 11; fi74++ {
                        for fi75 := 0; fi75 < 12; fi75++ {
                            fmt.Println(v14, v1)
                            fmt.Println(v13)
                            _ = fi75
                        }
                        _ = fi74
                    }
                }
                _ = v65
            }
            var v76 uint32 = 4294967295
            _ = v23
            _ = v76
        } else {
            v15 = v15
            v1 += (v12 * v4)
        }
        for fi77 := 0; fi77 < 1; fi77++ {
            var v78 int8 = (v9 * -128)
            _ = fi77
            _ = v78
        }
    } else {
        if (false && (true || (42 >= 429))) {
            v79 := GenS9{F0: -1, F1: -35}
            _ = v79
        } else {
            var v80 uint = v14
            v81 := GenS7{F0: 0}
            _ = v80
            _ = v81
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
