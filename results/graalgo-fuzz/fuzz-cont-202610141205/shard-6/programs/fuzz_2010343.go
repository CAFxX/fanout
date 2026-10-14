package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint16; F2 bool }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int8; F1 int; F2 uint8 }
type GenS6 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS7 struct { F0 bool; F1 int32; F2 int }
type GenS8 struct { F0 uint64; F1 bool }
func genFn0(p0 uint8, p1 uint16, p2 int16) int8 {
    var v6 float32 = (1.5 * 0.1)
    p1 = (p1 - (p1 << ((4294967295) & 63)))
    var v7 uint32 = 699
    var v8 bool = false
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 uint8 = (p0 ^ p0)
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1() uint32 {
    var v11 uint16 = 892
    var v12 string = ""
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    v4 := []string{"", "\n\t"}
    v5 := make(map[int]uint32)
    if false {
        if (((!(255 != 2)) && false) && true) {
            if (-2147483648 >= 991) {
                v13 := GenS1{F0: 2}
                _ = v13
            } else {
                var v14 float32 = 3.14159
                v15 := GenS2{F0: 99, F1: 1, F2: 2}
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 6; fi16++ {
                var v17 float64 = (1.5 * 0.1)
                v18 := GenS3{F0: 0, F1: 32767, F2: (!false)}
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 float64 = 1e10
            _ = v19
        } else {
            switch "foo bar" {
            case "a":
                switch 0 {
                case -2147483648:
                    if true {
                        var v20 int = 537
                        for fi21 := 0; fi21 < 7; fi21++ {
                            var v22 uint = 1
                            for fi23 := 0; fi23 < 5; fi23++ {
                                _ = len(v5)
                                if (-637 != 171) {
                                    if true {
                                        switch "\n\t" {
                                        case "":
                                            v22 = (^(v22))
                                        case "hello":
                                            v20 += fi23
                                        case "\n\t":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        var v24 uint = v22
                                        v20 = -1
                                        _ = v24
                                    }
                                    _ = len(v2)
                                }
                                _ = fi23
                            }
                            v25 := GenS4{F0: true, F1: 1}
                            _ = fi21
                            _ = v22
                            _ = v25
                        }
                        if (127 >= 182) {
                            v4 = append(v4, (("a" + "a") + "foo bar"))
                            switch v20 {
                            case -1:
                                v26 := GenS5{F0: genFn0(213, 69, 0), F1: v20, F2: 1}
                                _ = v26
                            case -874:
                                _ = len(v4)
                            }
                            v27 := GenS6{F0: 42, F1: 42, F2: 217}
                            _ = v27
                        }
                        _ = v20
                    } else {
                        var v28 int32 = -1
                        v29 := GenS7{F0: (0.0 == (-2.25 + 1.5)), F1: 42, F2: 893}
                        _ = v28
                        _ = v29
                    }
                case 0:
                    var v30 float64 = 1.5
                    _ = v30
                default:
                    var v31 int8 = -117
                    _ = v31
                }
            case "hello":
                var v32 uint8 = 32
                _ = v32
            default:
                for fi33 := 0; fi33 < 3; fi33++ {
                    fmt.Println(fi33)
                    _ = fi33
                }
            }
        }
    } else {
        switch "hello" {
        case "a":
            for fi34 := 0; fi34 < 15; fi34++ {
                var v35 string = ("foo bar" + ("hello" + ("hello" + ("" + (("foo bar" + "hello") + "\n\t")))))
                if false {
                    v36 := v2[fi34]
                    v36 = 65535
                    _ = v36
                }
                if false {
                    v37 := GenS6{F0: -32768, F1: 18446744073709551615, F2: 653}
                    v35 = (v35 + "foo bar")
                    _ = v37
                } else {
                    for fi38 := 0; fi38 < 13; fi38++ {
                        v5[(^(fi34))] = 768
                        v35 = v35
                        _ = fi38
                    }
                }
                _ = fi34
                _ = v35
            }
        case "\n\t":
            var v39 uint = 400
            _ = v39
        default:
            switch (("a" + (("\n\t" + "a") + "foo bar")) + "hello") {
            case "":
                switch 396 {
                case -479:
                    v40 := GenS5{F0: 100, F1: -925, F2: 2}
                    _ = v40
                case 42:
                    var v41 uint32 = 0
                    _ = v41
                default:
                    var v42 uint16 = 0
                    _ = v42
                }
            case "foo bar":
                var v43 uint32 = 784
                _ = v43
            case "a":
                var v44 int16 = 32767
                _ = v44
            }
        }
    }
    var v45 uint32 = 74
    var v46 int64 = 42
    var v47 float32 = 1.5
    v46 -= v46
    var v48 int8 = -19
    var v49 int64 = v46
    for fi50 := 0; fi50 < 8; fi50++ {
        v51 := GenS8{F0: 1, F1: false}
        _ = fi50
        _ = v51
    }
    if ((true && (8 >= v48)) && true) {
        var v52 uint8 = 127
        for fi53 := 0; fi53 < 19; fi53++ {
            switch fi53 {
            case -1:
                v47 -= 1e10
            case 2147483647:
                switch 128 {
                case 434:
                    if false {
                        _ = len(v4)
                    }
                case -1:
                    _ = len(v3)
                case 9223372036854775807:
                    var v54 float64 = 0.1
                    _ = v54
                default:
                    for fi55 := 0; fi55 < 4; fi55++ {
                        var v56 int = int(v52)
                        _ = fi55
                        _ = v56
                    }
                }
            case 32768:
                fmt.Println(v45, v52)
            }
            delete(v1, (("\n\t" + "foo bar") + ""))
            _ = fi53
        }
        _ = v52
    } else {
        v49 -= 588
    }
    _ = len(v4)
    var v57 int64 = (v49 / ((v49) | 1))
    v46 += (v57 ^ (v46 | (-(v49))))
    if (int(v49) > 42) {
        v58 := GenS3{F0: 83, F1: 42, F2: (252 < 317)}
        if len(v4) > 0 {
            v4[0] = "a"
        }
        _ = v58
    }
    var v59 uint16 = 665
    _ = len(v3)
    var v60 uint = 4294967295
    var v61 uint32 = (v45 & v45)
    v62 := v2[-1]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v57
    _ = v59
    _ = v60
    _ = v61
    _ = v62
}
