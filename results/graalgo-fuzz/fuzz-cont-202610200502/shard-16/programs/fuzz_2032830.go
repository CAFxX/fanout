package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int16; F1 bool }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint; F1 uint64 }
type GenS10 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS11 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS12 struct { F0 uint64; F1 int8; F2 int32 }
type GenS13 struct { F0 uint64; F1 int32; F2 int32 }

func main() {
    var v1 int8 = 82
    var v2 int8 = (v1 << ((251) & 63))
    v3 := make(map[string]uint32)
    var v4 uint64 = 868
    v5 := make(map[int]bool)
    for fi6 := 0; fi6 < 11; fi6++ {
        if (!(255 != 255)) {
            if false {
                switch 462 {
                case -2147483648:
                    if false {
                        for fi7 := 0; fi7 < 7; fi7++ {
                            v8 := GenS1{F0: 65535, F1: v4, F2: uint16(v4)}
                            v9 := GenS1{F0: 79, F1: 2, F2: 255}
                            fmt.Println(v4, fi7, fi6)
                            _ = fi7
                            _ = v8
                            _ = v9
                        }
                        fmt.Println(v2, v1)
                        v10 := GenS2{F0: 2147483648, F1: v4, F2: 0}
                        _ = v10
                    }
                }
                v2 = v1
            }
            fmt.Println(v2, v2)
        } else {
            for fi11 := 0; fi11 < 16; fi11++ {
                v12 := GenS3{F0: (("a" + "hello") + "")}
                _ = fi11
                _ = v12
            }
            switch "" {
            case "":
                v13 := GenS4{F0: v1}
                _ = v13
            case "hello":
                if (!true) {
                    v5[(fi6 + fi6)] = (254 >= 974)
                    if (false && true) {
                        if (false || true) {
                            var v14 bool = true
                            v15 := GenS5{F0: (^(v4))}
                            _ = v14
                            _ = v15
                        } else {
                            if false {
                                for fi16 := 0; fi16 < 15; fi16++ {
                                    if ((3.14159 * float64((1e10 * 0.1))) == ((((0.0 - ((1e10 * 3.14159) - -2.25)) * 3.14159) / 1.0) / 1.0)) {
                                        v2 = v1
                                        v17 := []int{}
                                        _ = v17
                                    }
                                    v18 := GenS6{F0: -1, F1: (415 != 65535)}
                                    v19 := GenS2{F0: 4294967295, F1: 4294967295, F2: 143}
                                    _ = fi16
                                    _ = v18
                                    _ = v19
                                }
                            }
                            var v20 uint64 = v4
                            fmt.Println(fi6)
                            _ = v20
                        }
                        for fi21 := 0; fi21 < 17; fi21++ {
                            var v22 uint64 = v4
                            var v23 uint64 = v22
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                        if (v1 != v2) {
                            var v24 int8 = (v2 ^ v1)
                            _ = v24
                        }
                    } else {
                        _ = len(v3)
                        for fi25 := 0; fi25 < 19; fi25++ {
                            var v26 int16 = -32768
                            var v27 bool = (!(432 != 451))
                            if v27 {
                                switch fi6 {
                                case 32767:
                                    _ = len(v5)
                                case -2147483648:
                                    var v28 bool = ((!(false && (65535 == 42))) || false)
                                    _ = v28
                                default:
                                    if v27 {
                                        v29 := GenS7{F0: 127, F1: 32767}
                                        switch ("foo bar" + ("" + "")) {
                                        case "foo bar":
                                            var v30 uint = 18446744073709551615
                                            _ = v30
                                        case "a":
                                            if ((fi25 >> ((0) & 63)) >= fi25) {
                                                v31 := GenS1{F0: 775, F1: (v4 - 4294967295), F2: 65535}
                                                switch ("foo bar" + "") {
                                                case "\n\t":
                                                    if true {
                                                        var v32 int16 = (v26 << ((45) & 63))
                                                        _ = v32
                                                    }
                                                case "foo bar":
                                                    v33 := GenS2{F0: 1, F1: (v4 & v4), F2: 4294967295}
                                                    _ = v33
                                                case "":
                                                    v34 := GenS1{F0: 65535, F1: v4, F2: 412}
                                                    _ = v34
                                                }
                                                v2 -= -14
                                                switch "\n\t" {
                                                case "":
                                                    v1 = (v1 * (v2 * v1))
                                                }
                                                var v35 uint64 = 615
                                                v36 := []int{(fi6 % ((fi25) | 1)), (fi25 & -730), int(v26)}
                                                delete(v3, "\n\t")
                                                switch ("" + "hello") {
                                                case "foo bar":
                                                    switch (("hello" + "hello") + "\n\t") {
                                                    case "a":
                                                        v37 := GenS4{F0: 0}
                                                        _ = v37
                                                    }
                                                case "hello":
                                                    v3[("foo bar" + (("\n\t" + (("hello" + "foo bar") + "foo bar")) + ""))] = 65535
                                                case "\n\t":
                                                    v26 += (-(v26))
                                                default:
                                                    fmt.Println(v1)
                                                }
                                                _ = v31
                                                _ = v35
                                                _ = v36
                                            } else {
                                                if (!v27) {
                                                    var v38 int = fi25
                                                    _ = v38
                                                } else {
                                                    if ("a" >= "\n\t") {
                                                        switch 2147483647 {
                                                        case 851:
                                                            fmt.Println(fi6, fi25)
                                                            v2 += v1
                                                            fmt.Println(v2, v2, v4)
                                                        case 127:
                                                            delete(v3, "foo bar")
                                                        default:
                                                            fmt.Println(fi6)
                                                        }
                                                    }
                                                    fmt.Println(fi25, v4, v26)
                                                }
                                                v4 += (v4 & v4)
                                            }
                                        case "":
                                            if (fi25 < fi6) {
                                                v4 += (v4 % ((v4) | 1))
                                            } else {
                                                var v39 uint8 = 0
                                                _ = v39
                                            }
                                        default:
                                            v40 := GenS3{F0: "a"}
                                            _ = v40
                                        }
                                        v41 := []int{83}
                                        v42 := GenS8{F0: -1}
                                        _ = v29
                                        _ = v41
                                        _ = v42
                                    } else {
                                        var v43 uint32 = 1
                                        _ = len(v5)
                                        fmt.Println(v2, v27)
                                        _ = v43
                                    }
                                }
                                var v44 int = fi6
                                v5[(fi25 * int(v2))] = (("foo bar" + "") >= "hello")
                                _ = v44
                            }
                            switch "a" {
                            case "a":
                                switch 45 {
                                case -377:
                                    if v27 {
                                        var v45 int16 = -415
                                        _ = v45
                                    }
                                }
                            case "hello":
                                if (v26 < v26) {
                                    v46 := v5[128]
                                    _ = v46
                                } else {
                                    if false {
                                        var v47 int32 = -1
                                        var v48 float32 = (((0.1 - 3.14159) / -2.25) + 0.0)
                                        _ = v47
                                        _ = v48
                                    } else {
                                        switch fi6 {
                                        case 128:
                                            v49 := GenS4{F0: (v1 % (((-(v1))) | 1))}
                                            _ = v49
                                        default:
                                            var v50 uint = 687
                                            _ = v50
                                        }
                                        var v51 float64 = 1.5
                                        switch "foo bar" {
                                        case "hello":
                                            v5[(fi25 ^ 2147483647)] = (9 >= 1)
                                        case "":
                                            var v52 int64 = -9223372036854775808
                                            _ = v52
                                        }
                                        _ = v51
                                    }
                                    var v53 int8 = (-(v1))
                                    var v54 string = "\n\t"
                                    _ = v53
                                    _ = v54
                                }
                            case "\n\t":
                                fmt.Println(v2, v2)
                            }
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        for fi55 := 0; fi55 < 3; fi55++ {
                            if false {
                                var v56 uint32 = 4294967295
                                var v57 int = (fi55 % (((-(fi55))) | 1))
                                _ = v56
                                _ = v57
                            } else {
                                fmt.Println(v4)
                            }
                            switch -9223372036854775808 {
                            case -409:
                                _ = len(v5)
                            case -1:
                                if (uint8(v4) != 156) {
                                    var v58 int8 = v1
                                    if (!(-2147483648 <= 750)) {
                                        v59 := []int{fi6}
                                        _ = v59
                                    }
                                    _ = v58
                                }
                            }
                            v60 := GenS9{F0: 520, F1: (v4 * 4294967295)}
                            var v61 string = "hello"
                            if true {
                                delete(v5, 259)
                                v4 = (v4 & (v4 + v4))
                            }
                            _ = fi55
                            _ = v60
                            _ = v61
                        }
                    }
                }
            }
            v62 := v5[(fi6 << ((115) & 63))]
            _ = v62
        }
        v63 := GenS10{F0: 0, F1: 42, F2: 697}
        v64 := GenS11{F0: 81, F1: 1, F2: (^(v4))}
        _ = fi6
        _ = v63
        _ = v64
    }
    var v65 uint8 = 127
    v66 := GenS12{F0: (^(v4)), F1: v2, F2: 219}
    var v67 uint32 = 65535
    v68 := []int{32767, -1, -9223372036854775808}
    for fi69 := 0; fi69 < 2; fi69++ {
        v1 = v2
        _ = fi69
    }
    _ = len(v68)
    v70 := GenS13{F0: 526, F1: -2147483648, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v65:", v65)
    fmt.Println("v67:", v67)
    fmt.Println("slicelen v68:", len(v68), v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v67
    _ = v68
    _ = v70
}
