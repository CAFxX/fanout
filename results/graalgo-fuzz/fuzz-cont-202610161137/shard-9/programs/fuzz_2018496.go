package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8; F1 int; F2 int32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64 }
func genFn0(p0 uint16) uint32 {
    v7 := []int{876, 42, -128}
    var v8 int32 = -2147483648
    _ = p0
    _ = v7
    _ = v8
    return 2147483648
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 1
    v3 := []uint32{42, uint32(v2), 4294967295}
    var v4 uint32 = 2147483648
    v5 := []int{-130, 4294967295}
    v6 := []string{((("a" + "hello") + "") + "foo bar")}
    v2 += v2
    for fi9 := 0; fi9 < 15; fi9++ {
        v2 += v2
        switch fi9 {
        case 127:
            var v10 uint16 = 1
            _ = v10
        case 893:
            v11 := GenS1{F0: 0, F1: 4294967295}
            _ = v11
        case 32768:
            switch ("hello" + ("a" + ("" + ""))) {
            case "hello":
                switch ("" + "a") {
                case "foo bar":
                    var v12 int16 = 906
                    _ = v12
                default:
                    if false {
                        fmt.Println(v4, fi9)
                        v4 = (v4 | genFn0(255))
                    }
                }
            }
        default:
            _ = len(v1)
        }
        v1[(-(fi9))] = 385
        _ = fi9
    }
    if false {
        var v13 int8 = 1
        _ = v13
    } else {
        if false {
            var v14 int32 = -360
            if true {
                v14 += v14
                v3 = append(v3, 2147483647)
            }
            var v15 uint = 2
            _ = v14
            _ = v15
        } else {
            v6 = append(v6, ("hello" + ""))
        }
        for fi16 := 0; fi16 < 19; fi16++ {
            switch "" {
            case "":
                v17 := GenS1{F0: -966, F1: 4294967295}
                _ = v17
            case "\n\t":
                switch fi16 {
                case 137:
                    v2 = (v2 / ((v2) | 1))
                }
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = fi16
        }
    }
    v18 := GenS2{F0: (!true)}
    if true {
        if true {
            v4 -= v4
            var v19 uint64 = v2
            _ = v19
        } else {
            v2 -= v2
            v20 := GenS1{F0: 863, F1: 255}
            var v21 uint32 = v4
            _ = v20
            _ = v21
        }
        var v22 int8 = 0
        _ = v22
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        if ((fi23 % ((fi23) | 1)) == fi23) {
            switch "\n\t" {
            case "\n\t":
                var v24 int64 = int64(v4)
                _ = v24
            }
        } else {
            for fi25 := 0; fi25 < 14; fi25++ {
                switch "\n\t" {
                case "hello":
                    if (!(false || false)) {
                        fmt.Println(fi25, fi23)
                        var v26 int8 = 100
                        switch (("" + "\n\t") + "\n\t") {
                        case "\n\t":
                            fmt.Println(v4, fi23, fi25)
                        case "":
                            switch int(v2) {
                            case 9223372036854775807:
                                var v27 int32 = -2147483648
                                _ = v27
                            case 127:
                                switch "foo bar" {
                                case "\n\t":
                                    if true {
                                        v4 += v4
                                        v6 = append(v6, "")
                                    } else {
                                        v4 -= 4294967295
                                    }
                                    switch (("" + "foo bar") + ((("\n\t" + ("\n\t" + ("hello" + ""))) + "foo bar") + "foo bar")) {
                                    case "\n\t":
                                        v26 -= v26
                                    case "foo bar":
                                        v4 = (^(v4))
                                    }
                                    fmt.Println(v4, v4)
                                case "hello":
                                    v28 := GenS3{F0: uint64(v4)}
                                    _ = v28
                                case "a":
                                    if (0 == 0) {
                                        if false {
                                            switch ("hello" + (("" + "hello") + ("hello" + ("a" + "\n\t")))) {
                                            case "\n\t":
                                                if ((v26 % ((v26) | 1)) < (v26 << ((18446744073709551615) & 63))) {
                                                    v29 := GenS3{F0: v2}
                                                    _ = v29
                                                }
                                            case "foo bar":
                                                switch (-(fi25)) {
                                                case -9223372036854775808:
                                                    if true {
                                                        if len(v3) > 0 {
                                                            v3[0] = genFn0(0)
                                                        }
                                                        v30 := GenS1{F0: -218, F1: 2}
                                                        switch ("hello" + "") {
                                                        case "hello":
                                                            var v31 uint64 = v2
                                                            _ = v31
                                                        case "":
                                                            var v32 int8 = v26
                                                            _ = v32
                                                        case "a":
                                                            fmt.Println(v2, v26)
                                                        default:
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        }
                                                        _ = v30
                                                    }
                                                case -199:
                                                    v2 = 4294967295
                                                case 237:
                                                    var v33 bool = false
                                                    _ = v33
                                                }
                                            case "hello":
                                                if true {
                                                    var v34 uint16 = 255
                                                    _ = v34
                                                } else {
                                                    v4 -= v4
                                                }
                                            }
                                        }
                                    }
                                }
                            case -929:
                                _ = len(v1)
                            default:
                                var v35 uint = 18446744073709551615
                                _ = v35
                            }
                        }
                        _ = v26
                    } else {
                        var v36 int32 = 42
                        var v37 string = (("foo bar" + "a") + "a")
                        _ = v36
                        _ = v37
                    }
                default:
                    fmt.Println(fi25, fi25)
                }
                fmt.Println(fi25, fi23)
                if (377 < 531) {
                    v4 -= v4
                    if true {
                        if true {
                            v2 = v2
                            v38 := GenS4{F0: v2, F1: 997}
                            _ = v38
                        }
                        v4 -= 2147483647
                        v2 += (v2 / ((v2) | 1))
                        fmt.Println(v2, v2, v4)
                        if true {
                            switch (fi25 >> ((42) & 63)) {
                            case 405:
                                v2 -= 18446744073709551615
                            default:
                                var v39 int32 = -1
                                _ = v39
                            }
                        } else {
                            v40 := GenS5{F0: true}
                            _ = v40
                        }
                    }
                    if true {
                        fmt.Println(v2, v2, fi25)
                        var v41 uint = 2
                        fmt.Println(fi25, v2)
                        _ = v41
                    } else {
                        if len(v6) > 0 {
                            v6[0] = ""
                        }
                    }
                } else {
                    v6 = append(v6, "hello")
                }
                var v42 int16 = 0
                fmt.Println(v2, v42)
                _ = fi25
                _ = v42
            }
            if ((fi23 | fi23) >= (fi23 + 42)) {
                fmt.Println(v2)
                if false {
                    for fi43 := 0; fi43 < 3; fi43++ {
                        v44 := GenS1{F0: 469, F1: 4294967295}
                        _ = fi43
                        _ = v44
                    }
                } else {
                    for fi45 := 0; fi45 < 6; fi45++ {
                        v2 = (v2 >> ((4294967295) & 63))
                        v4 -= genFn0(65535)
                        _ = fi45
                    }
                    for fi46 := 0; fi46 < 15; fi46++ {
                        v47 := GenS6{F0: 132, F1: (fi46 & fi46), F2: 42}
                        if false {
                            v4 += uint32(v2)
                            v4 -= 42
                            v2 = v2
                            switch "hello" {
                            case "a":
                                v48 := GenS7{F0: -201}
                                _ = v48
                            }
                        }
                        v4 = genFn0(uint16(v2))
                        _ = fi46
                        _ = v47
                    }
                }
            } else {
                switch (^(fi23)) {
                case -9223372036854775808:
                    delete(v1, (fi23 >> ((583) & 63)))
                case 1000:
                    v49 := GenS3{F0: v2}
                    _ = v49
                }
            }
        }
        fmt.Println(v2, fi23, v2)
        fmt.Println(v2, fi23, fi23)
        _ = fi23
    }
    var v50 int8 = -1
    if true {
        if (42 < 1) {
            v50 = v50
        } else {
            v5 = append(v5, 517)
            if false {
                var v51 uint32 = v4
                v52 := GenS8{F0: -9223372036854775808}
                _ = v51
                _ = v52
            } else {
                var v53 uint8 = 42
                _ = v53
            }
            for fi54 := 0; fi54 < 4; fi54++ {
                for fi55 := 0; fi55 < 18; fi55++ {
                    fmt.Println(fi54, v4, v4)
                    _ = fi55
                }
                _ = fi54
            }
        }
        switch "a" {
        case "hello":
            v56 := GenS8{F0: int64(v4)}
            _ = v56
        default:
            v57 := GenS2{F0: false}
            _ = v57
        }
    } else {
        var v58 int64 = 838
        if (1 >= 135) {
            v1[32768] = -9223372036854775808
            fmt.Println(v50, v50)
            for fi59 := 0; fi59 < 11; fi59++ {
                v60 := GenS5{F0: true}
                for fi61 := 0; fi61 < 14; fi61++ {
                    if true {
                        var v62 uint32 = (v4 % ((genFn0(255)) | 1))
                        var v63 uint32 = 195
                        var v64 uint16 = 32767
                        var v65 int = fi59
                        v66 := GenS5{F0: ("hello" != ("hello" + ("hello" + "a")))}
                        _ = v62
                        _ = v63
                        _ = v64
                        _ = v65
                        _ = v66
                    }
                    _ = fi61
                }
                _ = fi59
                _ = v60
            }
        } else {
            for fi67 := 0; fi67 < 9; fi67++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v58 += v58
                _ = fi67
            }
            v50 -= v50
            fmt.Println(v50, v58)
        }
        v58 = -9223372036854775808
        _ = v58
    }
    v68 := GenS1{F0: 1000, F1: 0}
    if true {
        v4 = (v4 | v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v68
}
