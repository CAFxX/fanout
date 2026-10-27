package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: 0, F1: -128}
        if false {
            var v7 int32 = 1
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 int32 = v7
                var v10 uint8 = 128
                v9 = -374
                var v11 uint = 18446744073709551615
                if (true && true) {
                    if (32767 >= 0) {
                        var v12 uint8 = (v10 * (v10 * (v10 >> ((v11) & 63))))
                        _ = len(v1)
                        switch "a" {
                        case "hello":
                            v13 := []int{(fi8 % (((fi5 - (fi5 / ((fi8) | 1)))) | 1)), (fi8 ^ (^(fi8))), int(v10), int(v11)}
                            _ = v13
                        case "foo bar":
                            var v14 uint = (^(v11))
                            v15 := GenS1{F0: -14, F1: -1}
                            _ = v14
                            _ = v15
                        case "\n\t":
                            v16 := GenS1{F0: -1, F1: 42}
                            _ = v16
                        }
                        v17 := GenS1{F0: 0, F1: 42}
                        _ = v12
                        _ = v17
                    } else {
                        fmt.Println(v7, fi8, v9)
                        var v18 uint32 = 780
                        _ = v18
                    }
                    var v19 string = "hello"
                    _ = v19
                } else {
                    switch "hello" {
                    case "\n\t":
                        var v20 int64 = 63
                        _ = v20
                    default:
                        if false {
                            switch "\n\t" {
                            case "a":
                                var v21 bool = ((v10 >= 42) || false)
                                _ = v21
                            case "":
                                var v22 int8 = 100
                                _ = v22
                            default:
                                v23 := []int{(fi5 % ((127) | 1))}
                                _ = v23
                            }
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                var v24 string = (("" + "foo bar") + "\n\t")
                                _ = v24
                            }
                            if true {
                                v25 := GenS2{F0: fi8, F1: -690, F2: v10}
                                fmt.Println(fi5, v10, v9)
                                _ = v25
                            } else {
                                v11 -= (v11 + (v11 * v11))
                            }
                            v26 := []int{}
                            if (-862 >= 489) {
                                switch "" {
                                case "a":
                                    v7 -= 846
                                case "hello":
                                    var v27 float64 = 3.14159
                                    _ = v27
                                case "":
                                    var v28 uint64 = 0
                                    _ = v28
                                default:
                                    v29 := GenS1{F0: -128, F1: 100}
                                    _ = v29
                                }
                                v9 -= (^(v9))
                                if ((1e10 / 1e10) < 0.0) {
                                    var v30 float64 = float64((((3.14159 - 1e10) * float32((3.14159 * 0.1))) / 1e10))
                                    v30 = v30
                                    _ = v30
                                } else {
                                    var v31 string = ("a" + "\n\t")
                                    switch 659 {
                                    case 398:
                                        if ((^(v9)) >= 1) {
                                            v9 -= v9
                                            var v32 int16 = int16(v11)
                                            _ = v32
                                        }
                                    default:
                                        var v33 int16 = 32767
                                        _ = v33
                                    }
                                    _ = v31
                                }
                            }
                            var v34 uint = v11
                            _ = v26
                            _ = v34
                        }
                    }
                }
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            for fi35 := 0; fi35 < 18; fi35++ {
                delete(v1, 2147483647)
                v36 := v1[4294967295]
                switch fi35 {
                case 4294967295:
                    var v37 int32 = (v7 / ((-1) | 1))
                    _ = v37
                case 128:
                    switch (fi5 + fi35) {
                    case 0:
                        v38 := GenS1{F0: -62, F1: 97}
                        _ = v38
                    case 42:
                        v39 := GenS3{F0: 128, F1: v7}
                        _ = v39
                    case 4294967295:
                        v36 += 2147483648
                    }
                case -2147483648:
                    if (false && true) {
                        if false {
                            v40 := []int{0, (fi35 | fi35), (fi5 << ((473) & 63))}
                            _ = v40
                        }
                        v36 = v36
                    } else {
                        var v41 uint16 = 42
                        _ = v41
                    }
                }
                _ = fi35
                _ = v36
            }
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    var v42 float32 = 1e10
    for fi43 := 0; fi43 < 12; fi43++ {
        for fi44 := 0; fi44 < 19; fi44++ {
            if false {
                switch "\n\t" {
                case "a":
                    fmt.Println(fi44, v42)
                default:
                    var v45 bool = false
                    _ = v45
                }
            } else {
                delete(v3, (fi44 | -181))
                var v46 int32 = 959
                _ = v46
            }
            var v47 float64 = -2.25
            var v48 uint64 = 808
            var v49 int8 = 127
            _ = fi44
            _ = v47
            _ = v48
            _ = v49
        }
        var v50 uint32 = 1
        _ = fi43
        _ = v50
    }
    v51 := []int{471, 282}
    var v52 uint64 = 18446744073709551615
    fmt.Println(v52, v42)
    if true {
        v52 += v52
    } else {
        v53 := GenS2{F0: 9223372036854775807, F1: 1, F2: 127}
        _ = v53
    }
    v42 = v42
    var v54 uint32 = 0
    switch 32767 {
    case -1:
        for fi55 := 0; fi55 < 12; fi55++ {
            var v56 string = ("hello" + "")
            if ((false || (-2147483648 >= 9)) || (false || true)) {
                var v57 int = fi55
                _ = v57
            }
            _ = fi55
            _ = v56
        }
    default:
        if false {
            v58 := GenS3{F0: 42, F1: 2147483647}
            v2[(v54 * v54)] = "foo bar"
            fmt.Println(v42)
            _ = v58
        } else {
            if (true && false) {
                var v59 int16 = -900
                if false {
                    v52 -= (v52 >> ((2) & 63))
                    for fi60 := 0; fi60 < 18; fi60++ {
                        v59 += v59
                        _ = fi60
                    }
                    v61 := GenS2{F0: -461, F1: (v59 ^ v59), F2: 42}
                    _ = v61
                } else {
                    fmt.Println(v59)
                    if ((-216 >= -9223372036854775808) && (("foo bar" + "") != ("a" + ("a" + "a")))) {
                        for fi62 := 0; fi62 < 8; fi62++ {
                            var v63 int = fi62
                            var v64 uint8 = 42
                            var v65 uint16 = uint16(v52)
                            _ = fi62
                            _ = v63
                            _ = v64
                            _ = v65
                        }
                    } else {
                        if true {
                            _ = len(v4)
                        } else {
                            for fi66 := 0; fi66 < 15; fi66++ {
                                _ = len(v3)
                                var v67 int16 = v59
                                _ = fi66
                                _ = v67
                            }
                            var v68 float32 = v42
                            if false {
                                v42 = v68
                                _ = len(v4)
                                v69 := GenS2{F0: -2147483648, F1: (^(v59)), F2: 127}
                                _ = v69
                            } else {
                                v70 := GenS4{F0: 1}
                                v71 := v1[(v54 ^ 4294967295)]
                                _ = v70
                                _ = v71
                            }
                            _ = v68
                        }
                    }
                }
                _ = v59
            } else {
                fmt.Println(v52, v52)
                var v72 int16 = 32767
                if len(v51) >= 2 {
                    v51 = v51[1:len(v51)]
                }
                _ = v72
            }
        }
    }
    var v73 uint64 = 2
    if true {
        var v74 string = ""
        _ = v74
    }
    if true {
        if len(v51) >= 2 {
            v51 = v51[1:len(v51)]
        }
        var v75 int = 871
        _ = v75
    }
    fmt.Println(v73)
    if true {
        v42 = v42
    } else {
        var v76 bool = (9223372036854775807 <= -949)
        _ = v76
    }
    var v77 int8 = -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    fmt.Println("v73:", v73)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v51
    _ = v52
    _ = v54
    _ = v73
    _ = v77
}
