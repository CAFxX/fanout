package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 int16; F2 string }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 708
    var v3 uint16 = 835
    v4 := []int{32768, -2147483648, 32767, 547}
    v5 := v1["\n\t"]
    v6 := GenS1{F0: 1}
    var v7 string = "a"
    if len(v4) > 0 {
        v4[0] = (v5 - (v5 >> ((993) & 63)))
    }
    fmt.Println(v5, v5)
    _ = len(v4)
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            if (2 == v2) {
                switch (fi9 * (fi8 << ((1) & 63))) {
                case 127:
                    v2 -= (v2 >> ((0) & 63))
                case 2147483647:
                    if true {
                        switch fi9 {
                        case 42:
                            if false {
                                switch 32768 {
                                case -1:
                                    v7 = ("foo bar" + v7)
                                    var v10 bool = (3.14159 < ((1.5 * 3.14159) * ((3.14159 + 0.0) * 0.1)))
                                    _ = v10
                                case 723:
                                    var v11 uint16 = (v3 * uint16(v2))
                                    _ = v11
                                case 32768:
                                    v12 := GenS2{F0: 42}
                                    _ = v12
                                }
                                v1["\n\t"] = (fi9 + 0)
                                var v13 int64 = -168
                                _ = v13
                            } else {
                                v14 := GenS2{F0: 4294967295}
                                v15 := GenS3{F0: -1, F1: 163}
                                _ = v14
                                _ = v15
                            }
                            v16 := GenS4{F0: 2147483647}
                            v5 += (fi9 ^ v5)
                            if (1 > 0) {
                                if true {
                                    v5 = v5
                                    _ = len(v1)
                                } else {
                                    delete(v1, "a")
                                    v17 := GenS5{F0: (v2 & v2)}
                                    _ = v17
                                }
                            }
                            switch (v5 + (-(fi9))) {
                            case 2147483647:
                                if (false || (!true)) {
                                    v2 -= 2
                                    v3 -= (v3 & v3)
                                    switch "a" {
                                    case "hello":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    if true {
                                        if len(v4) > 0 {
                                            v4[0] = 2147483647
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v18 int8 = 109
                                        _ = v18
                                    }
                                    if (false || true) {
                                        v19 := GenS4{F0: -590}
                                        if true {
                                            var v20 uint8 = 128
                                            v20 = (v20 + v20)
                                            v2 = v2
                                            switch (("\n\t" + "hello") + "hello") {
                                            case "":
                                                if (42 >= -1) {
                                                    var v21 int16 = 1000
                                                    fmt.Println(v21, v20)
                                                    _ = len(v4)
                                                    _ = v21
                                                } else {
                                                    v22 := GenS1{F0: 949}
                                                    var v23 float64 = (-2.25 * 0.0)
                                                    _ = v22
                                                    _ = v23
                                                }
                                            case "hello":
                                                v2 -= (v2 % ((v2) | 1))
                                            case "\n\t":
                                                v2 = v2
                                            default:
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                            _ = v20
                                        } else {
                                            v24 := GenS6{F0: (^(fi8)), F1: -1, F2: ((v7 + v7) + ("a" + ""))}
                                            if (-1 >= 127) {
                                                if true {
                                                    var v25 int64 = -1
                                                    fmt.Println(fi9, fi8)
                                                    v26 := GenS3{F0: v25, F1: 2147483648}
                                                    _ = len(v1)
                                                    var v27 float64 = 1e10
                                                    _ = v25
                                                    _ = v26
                                                    _ = v27
                                                }
                                                v28 := GenS2{F0: 42}
                                                _ = v28
                                            }
                                            var v29 int16 = -560
                                            _ = v24
                                            _ = v29
                                        }
                                        v2 += uint64(v3)
                                        _ = v19
                                    }
                                    var v30 uint16 = uint16(v2)
                                    _ = v30
                                }
                            case 4294967295:
                                v31 := GenS1{F0: 561}
                                _ = v31
                            }
                            _ = v16
                        }
                    } else {
                        v5 += int(v3)
                    }
                }
                fmt.Println(v5, v7, v7)
                if (255 > v3) {
                    switch (fi8 - (fi8 + fi8)) {
                    case 127:
                        var v32 uint16 = v3
                        _ = v32
                    case -656:
                        if false {
                            var v33 bool = false
                            var v34 float32 = 3.14159
                            v35 := v1[""]
                            _ = v33
                            _ = v34
                            _ = v35
                        } else {
                            v36 := v1[((("" + (v7 + "a")) + v7) + ("hello" + ("a" + "a")))]
                            v37 := GenS2{F0: 0}
                            _ = v36
                            _ = v37
                        }
                    case 0:
                        fmt.Println(v5)
                        var v38 int16 = 1
                        if (0 <= 4294967295) {
                            fmt.Println(v2)
                            if true {
                                fmt.Println(v3, fi9, v2)
                            }
                            var v39 int64 = 1
                            _ = v39
                        }
                        if true {
                            var v40 int16 = v38
                            switch (v7 + "foo bar") {
                            case "":
                                v4 = append(v4, fi8)
                            case "hello":
                                if ((0.0 - 0.0) != float32(0.0)) {
                                    if (!true) {
                                        var v41 int8 = -1
                                        _ = v41
                                    }
                                } else {
                                    var v42 float32 = 1e10
                                    switch (^(fi9)) {
                                    case 961:
                                        if (!(!false)) {
                                            v43 := GenS7{F0: 2147483647, F1: v40}
                                            switch 32768 {
                                            case 2147483647:
                                                v44 := GenS2{F0: 527}
                                                _ = v44
                                            case 127:
                                                v38 += -882
                                            }
                                            _ = v43
                                        } else {
                                            delete(v1, v7)
                                            var v45 string = v7
                                            _ = v45
                                        }
                                    default:
                                        if (false && (1 >= 1)) {
                                            var v46 int16 = v38
                                            v40 += v38
                                            var v47 string = ((v7 + ((v7 + (v7 + v7)) + v7)) + v7)
                                            fmt.Println(v2, v42, v5)
                                            var v48 int = fi8
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                        }
                                    }
                                    if true {
                                        v5 = fi9
                                        var v49 float64 = 3.14159
                                        _ = len(v4)
                                        _ = v49
                                    } else {
                                        v50 := v1[(v7 + v7)]
                                        switch v7 {
                                        case "a":
                                            var v51 int32 = 2147483647
                                            _ = v51
                                        case "hello":
                                            var v52 string = (v7 + "\n\t")
                                            _ = v52
                                        }
                                        _ = v50
                                    }
                                    _ = v42
                                }
                            }
                            v1[v7] = (fi8 << ((1) & 63))
                            _ = v40
                        }
                        switch ((v7 + ("hello" + "hello")) + (v7 + "\n\t")) {
                        case "":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case "\n\t":
                            fmt.Println(v38, v38)
                        case "a":
                            delete(v1, ("hello" + v7))
                        default:
                            fmt.Println(v38, v7)
                        }
                        var v53 uint8 = 235
                        switch (("foo bar" + v7) + v7) {
                        case "":
                            var v54 float32 = (-2.25 - 0.0)
                            _ = v54
                        case "hello":
                            if false {
                                v55 := GenS3{F0: -1, F1: uint32(v2)}
                                if (-32768 <= 563) {
                                    var v56 string = "hello"
                                    _ = v56
                                } else {
                                    v53 = v53
                                }
                                _ = v55
                            }
                        default:
                            var v57 float64 = 0.1
                            _ = v57
                        }
                        if false {
                            v1[""] = fi9
                            _ = len(v4)
                            v58 := GenS8{F0: 18446744073709551615}
                            _ = v58
                        } else {
                            var v59 int64 = -961
                            _ = v59
                        }
                        _ = len(v1)
                        v7 = "foo bar"
                        _ = len(v1)
                        v60 := v1[((v7 + v7) + "")]
                        var v61 uint64 = 18446744073709551615
                        v62 := GenS9{F0: false}
                        v63 := GenS4{F0: 482}
                        _ = v38
                        _ = v53
                        _ = v60
                        _ = v61
                        _ = v62
                        _ = v63
                    default:
                        if false {
                            var v64 uint64 = v2
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v65 := GenS2{F0: 255}
                            _ = v64
                            _ = v65
                        } else {
                            fmt.Println(fi8, v5, fi8)
                        }
                    }
                    v66 := GenS3{F0: -385, F1: uint32(v2)}
                    v7 = "foo bar"
                    _ = v66
                }
            } else {
                delete(v1, ("hello" + v7))
            }
            var v67 uint = 2
            _ = fi9
            _ = v67
        }
        var v68 uint = 95
        _ = fi8
        _ = v68
    }
    v7 = v7
    _ = len(v4)
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
