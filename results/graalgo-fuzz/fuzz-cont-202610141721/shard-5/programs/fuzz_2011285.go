package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 int64; F2 uint16 }
type GenS6 struct { F0 int; F1 int64; F2 uint32 }
type GenS7 struct { F0 uint16; F1 string }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int8; F1 string }
func genFn0(p0 int64) int16 {
    var v4 int = 32767
    v4 = v4
    v5 := []int{}
    _ = p0
    _ = v4
    _ = v5
    return -32768
}

func main() {
    v1 := []string{}
    v2 := make(map[int]uint32)
    var v3 string = ""
    fmt.Println(v3)
    var v6 uint8 = 255
    v7 := GenS1{F0: 1, F1: ("hello" + (("hello" + (("\n\t" + "foo bar") + (("a" + v3) + "foo bar"))) + ""))}
    v8 := GenS1{F0: -32768, F1: (("\n\t" + "") + (("\n\t" + v3) + "a"))}
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS2{F0: genFn0(681)}
        var v11 int = 9223372036854775807
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint16 = 832
    switch "\n\t" {
    case "\n\t":
        _ = len(v1)
    case "hello":
        if true {
            if (-1 > 32767) {
                fmt.Println(v12, v3)
            }
            for fi13 := 0; fi13 < 15; fi13++ {
                var v14 uint8 = v6
                v15 := GenS1{F0: -69, F1: (v3 + "foo bar")}
                _ = fi13
                _ = v14
                _ = v15
            }
            if (!false) {
                for fi16 := 0; fi16 < 6; fi16++ {
                    if (2 != 4294967295) {
                        var v17 float32 = ((((0.1 * (3.14159 + 0.0)) - -2.25) + (0.0 / 0.1)) - 3.14159)
                        var v18 int64 = 42
                        v19 := GenS1{F0: 32767, F1: (v3 + ("a" + v3))}
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = len(v1)
                    if (true && (32767 >= 0)) {
                        v20 := v2[683]
                        v21 := GenS2{F0: genFn0(348)}
                        switch "hello" {
                        case "hello":
                            fmt.Println(v3, v3)
                        case "a":
                            v2[int(v20)] = v20
                        default:
                            var v22 uint = 18446744073709551615
                            _ = v22
                        }
                        _ = v20
                        _ = v21
                    }
                    _ = fi16
                }
                v23 := GenS3{F0: (v12 | (^(v12)))}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v23
            } else {
                v12 = v12
            }
        } else {
            if (v6 == (v6 >> ((255) & 63))) {
                for fi24 := 0; fi24 < 1; fi24++ {
                    var v25 uint32 = 323
                    _ = fi24
                    _ = v25
                }
                v3 = ((v3 + v3) + ("a" + "\n\t"))
            } else {
                fmt.Println(v12, v3, v6)
                var v26 uint8 = (v6 >> ((719) & 63))
                _ = v26
            }
        }
    }
    fmt.Println(v12, v3)
    v27 := v2[32767]
    var v28 int32 = 2147483647
    if (42 != 0) {
        for fi29 := 0; fi29 < 18; fi29++ {
            var v30 int64 = 9223372036854775807
            _ = fi29
            _ = v30
        }
        v31 := GenS3{F0: uint16(v27)}
        _ = v31
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        fmt.Println(v28, v27, v28)
        for fi33 := 0; fi33 < 17; fi33++ {
            switch (v3 + v3) {
            case "hello":
                if true {
                    switch 4294967295 {
                    case 32767:
                        var v34 uint64 = 251
                        _ = v34
                    case 2147483647:
                        var v35 uint32 = (v27 >> ((992) & 63))
                        _ = v35
                    default:
                        v36 := GenS3{F0: v12}
                        _ = v36
                    }
                }
            case "\n\t":
                if (v3 < "foo bar") {
                    if (1 == 9223372036854775807) {
                        var v37 float32 = (3.14159 * -2.25)
                        _ = v37
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            var v38 float64 = 0.1
                            _ = v38
                        }
                    }
                    switch fi32 {
                    case 1:
                        if false {
                            switch fi32 {
                            case 1:
                                switch fi33 {
                                case 4294967295:
                                    if (v28 != -2147483648) {
                                        var v39 string = "\n\t"
                                        var v40 int32 = 42
                                        v28 = (v40 - (v40 - -417))
                                        if false {
                                            if false {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v12 = (v12 & v12)
                                                v3 = ""
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v41 := GenS4{F0: 42}
                                            _ = v41
                                        } else {
                                            switch fi32 {
                                            case 32768:
                                                fmt.Println(v28)
                                            case 440:
                                                switch int(v27) {
                                                case 4294967295:
                                                    switch "hello" {
                                                    case "hello":
                                                        v42 := GenS5{F0: (v28 - v40), F1: -22, F2: (v12 - (^(v12)))}
                                                        _ = v42
                                                    }
                                                case 0:
                                                    var v43 uint16 = v12
                                                    _ = v43
                                                case 32767:
                                                    fmt.Println(fi33)
                                                }
                                            case 0:
                                                if false {
                                                    v27 = v27
                                                    var v44 bool = (!false)
                                                    _ = v44
                                                }
                                            }
                                            fmt.Println(v3, v40)
                                        }
                                        v45 := GenS4{F0: -128}
                                        _ = v39
                                        _ = v40
                                        _ = v45
                                    }
                                case 32767:
                                    if (-9223372036854775808 > 91) {
                                        delete(v2, fi33)
                                        var v46 float32 = -2.25
                                        var v47 int8 = -123
                                        _ = v46
                                        _ = v47
                                    } else {
                                        fmt.Println(v12)
                                        fmt.Println(fi32, fi33, v12)
                                        fmt.Println(v27, fi33, v6)
                                        if true {
                                            var v48 string = v3
                                            v49 := GenS5{F0: v28, F1: 0, F2: v12}
                                            if true {
                                                v27 -= 65535
                                            } else {
                                                var v50 uint64 = 0
                                                _ = v50
                                            }
                                            _ = v48
                                            _ = v49
                                        } else {
                                            v1 = append(v1, v3)
                                            fmt.Println(fi32, v12)
                                            if (0 > 42) {
                                                switch v3 {
                                                case "a":
                                                    if true {
                                                        v12 += (v12 * v12)
                                                    }
                                                default:
                                                    v51 := v2[fi32]
                                                    _ = v51
                                                }
                                            } else {
                                                if true {
                                                    var v52 int64 = -850
                                                    v53 := GenS6{F0: 4294967295, F1: (^(v52)), F2: (v27 ^ 0)}
                                                    _ = v52
                                                    _ = v53
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        }
                                        var v54 int = 127
                                        _ = v54
                                    }
                                case 2147483647:
                                    v55 := GenS5{F0: (v28 | -1), F1: -695, F2: v12}
                                    _ = v55
                                }
                            case 610:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case 42:
                                var v56 uint = 459
                                _ = v56
                            }
                            var v57 int8 = -1
                            if (-1 > v57) {
                                delete(v2, (fi33 ^ (fi32 / ((fi32) | 1))))
                                if len(v1) > 0 {
                                    v1[0] = v3
                                }
                                fmt.Println(v3)
                            } else {
                                v58 := GenS2{F0: -80}
                                var v59 uint32 = (v27 >> ((868) & 63))
                                v60 := GenS7{F0: (^(v12)), F1: v3}
                                _ = v58
                                _ = v59
                                _ = v60
                            }
                            _ = v57
                        } else {
                            var v61 int64 = 0
                            _ = v61
                        }
                    case 206:
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                    default:
                        if (fi32 < fi32) {
                            v62 := GenS8{F0: 1}
                            v6 += 2
                            _ = v62
                        } else {
                            _ = len(v2)
                            delete(v2, (fi32 + -9223372036854775808))
                        }
                    }
                    v27 = (^(v27))
                } else {
                    fmt.Println(v12, fi33, v12)
                    var v63 float32 = -2.25
                    var v64 uint8 = (v6 << ((690) & 63))
                    _ = v63
                    _ = v64
                }
            case "":
                var v65 float64 = 0.1
                _ = v65
            default:
                v66 := GenS9{F0: 100, F1: ""}
                switch "hello" {
                case "a":
                    v12 += v12
                case "":
                    var v67 int16 = -32768
                    _ = v67
                }
                _ = v66
            }
            var v68 int32 = v28
            var v69 int32 = -2147483648
            _ = fi33
            _ = v68
            _ = v69
        }
        v28 -= v28
        _ = fi32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
