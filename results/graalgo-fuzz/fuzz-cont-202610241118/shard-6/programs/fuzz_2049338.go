package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 int; F2 uint64 }
func genFn0(p0 int) int32 {
    v4 := GenS1{F0: 127}
    p0 = -9223372036854775808
    _ = p0
    _ = v4
    return 0
}
func genFn1() int8 {
    switch "foo bar" {
    case "foo bar":
        var v5 string = "foo bar"
        _ = v5
    case "a":
        var v6 float64 = 3.14159
        _ = v6
    case "hello":
        var v7 uint32 = 947
        _ = v7
    default:
        var v8 float32 = ((float32(1.5) + 1.5) / 1.5)
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int32 = 1
        _ = fi9
        _ = v10
    }
    if true {
        v11 := []int{462}
        v12 := GenS1{F0: 71}
        _ = v11
        _ = v12
    } else {
        var v13 int64 = 1
        v13 = 42
        _ = v13
    }
    if false {
        var v14 uint64 = 159
        _ = v14
    }
    return 100
}

func main() {
    v1 := []uint32{678}
    v2 := []int{4294967295, 32768}
    var v3 uint64 = 42
    switch "a" {
    case "foo bar":
        var v15 float64 = 1.5
        _ = v15
    case "":
        for fi16 := 0; fi16 < 11; fi16++ {
            if true {
                v3 += 4294967295
                v17 := GenS1{F0: 255}
                _ = v17
            } else {
                v3 += v3
            }
            _ = fi16
        }
    case "\n\t":
        var v18 float64 = 0.0
        _ = v18
    }
    var v19 int16 = 1
    var v20 int16 = (v19 + (v19 ^ v19))
    v21 := GenS1{F0: 42}
    if ((v20 | 244) >= (^(v20))) {
        for fi22 := 0; fi22 < 9; fi22++ {
            for fi23 := 0; fi23 < 19; fi23++ {
                v19 = (v20 / (((v20 * v20)) | 1))
                if false {
                    if true {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v20 -= -1
                    } else {
                        var v24 int32 = -2147483648
                        var v25 int64 = -614
                        _ = v24
                        _ = v25
                    }
                    if true {
                        var v26 int64 = 42
                        _ = v26
                    }
                    fmt.Println(v20, fi22)
                }
                _ = fi23
            }
            _ = fi22
        }
    }
    v20 = v19
    if false {
        var v27 uint32 = uint32(v3)
        for fi28 := 0; fi28 < 18; fi28++ {
            _ = len(v2)
            if ((v19 >> ((2) & 63)) == -447) {
                var v29 int32 = 436
                v30 := GenS1{F0: 128}
                _ = v29
                _ = v30
            }
            v3 = (^(v3))
            _ = fi28
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v27
    }
    switch "hello" {
    case "\n\t":
        v3 = (v3 / ((v3) | 1))
    case "hello":
        v20 -= (v19 << ((1) & 63))
    case "a":
        v20 -= (v19 | (v20 ^ (v19 | (v19 + (v20 % (((v19 & v20)) | 1))))))
    default:
        if (((((0.0 / 1e10) - 1.5) * 3.14159) - (float64(((-2.25 / 1.5) + (1e10 * 1e10))) + 0.0)) == float64((1.5 / 1e10))) {
            if (!true) {
                v3 += v3
                for fi31 := 0; fi31 < 7; fi31++ {
                    var v32 int = (^(fi31))
                    var v33 int16 = v19
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            } else {
                fmt.Println(v20)
                v3 = 42
                v34 := GenS1{F0: 37}
                _ = v34
            }
            if len(v2) > 0 {
                v2[0] = 128
            }
            if (!(255 < 128)) {
                v1 = append(v1, 739)
            }
        } else {
            if ((v20 ^ (v19 * (-(v19)))) != v19) {
                v35 := GenS2{F0: v3, F1: uint(v3), F2: (v3 - (^(v3)))}
                v1 = append(v1, 2147483647)
                var v36 uint = 981
                _ = v35
                _ = v36
            }
            switch "hello" {
            case "hello":
                v19 += v19
            case "":
                v20 += -510
            case "\n\t":
                if (!(false && ((("\n\t" + ("foo bar" + "hello")) + "hello") == "hello"))) {
                    v19 = (^(v19))
                    switch -2147483648 {
                    case -260:
                        if false {
                            for fi37 := 0; fi37 < 9; fi37++ {
                                for fi38 := 0; fi38 < 9; fi38++ {
                                    v39 := GenS1{F0: 0}
                                    v40 := GenS1{F0: 2}
                                    if len(v2) > 0 {
                                        v2[0] = fi37
                                    }
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                                for fi41 := 0; fi41 < 5; fi41++ {
                                    v19 -= v20
                                    if (false && false) {
                                        switch "a" {
                                        case "hello":
                                            v19 -= v20
                                        case "\n\t":
                                            v20 = (v19 % ((-275) | 1))
                                        default:
                                            switch ("\n\t" + "hello") {
                                            case "foo bar":
                                                var v42 uint16 = 0
                                                _ = v42
                                            case "":
                                                if false {
                                                    var v43 int = fi41
                                                    v44 := GenS3{F0: v3, F1: 0, F2: 83}
                                                    _ = v43
                                                    _ = v44
                                                } else {
                                                    switch 9223372036854775807 {
                                                    case 2147483647:
                                                        if false {
                                                            var v45 uint16 = 65535
                                                            _ = v45
                                                        } else {
                                                            _ = len(v1)
                                                            var v46 bool = (1e10 == 1.5)
                                                            var v47 string = "foo bar"
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                    case 0:
                                                        var v48 bool = false
                                                        _ = v48
                                                    case 127:
                                                        var v49 uint64 = 753
                                                        v50 := GenS4{F0: -9223372036854775808}
                                                        if true {
                                                            var v51 uint = 447
                                                            switch fi37 {
                                                            case -907:
                                                                if len(v2) > 0 {
                                                                    v2[0] = 1
                                                                }
                                                            default:
                                                                var v52 uint64 = (^(v3))
                                                                _ = v52
                                                            }
                                                            var v53 int16 = (v19 * (^(v20)))
                                                            fmt.Println(v49)
                                                            _ = v51
                                                            _ = v53
                                                        } else {
                                                            v19 -= 42
                                                            var v54 uint32 = 937
                                                            _ = v54
                                                        }
                                                        _ = v49
                                                        _ = v50
                                                    default:
                                                        var v55 uint16 = 255
                                                        _ = v55
                                                    }
                                                }
                                            }
                                        }
                                    } else {
                                        var v56 int8 = -128
                                        _ = v56
                                    }
                                    _ = fi41
                                }
                                v57 := GenS5{F0: "a", F1: (fi37 & 0), F2: v3}
                                _ = fi37
                                _ = v57
                            }
                            v20 = (v19 * v19)
                        }
                    }
                }
            default:
                if (false || true) {
                    for fi58 := 0; fi58 < 5; fi58++ {
                        if len(v2) > 0 {
                            v2[0] = (fi58 % ((fi58) | 1))
                        }
                        _ = fi58
                    }
                } else {
                    switch 32768 {
                    case 1:
                        fmt.Println(v20, v20)
                    case 985:
                        var v59 float64 = (1.5 / 1e10)
                        _ = v59
                    }
                }
            }
            v60 := GenS4{F0: 1}
            _ = v60
        }
    }
    v61 := GenS1{F0: 2}
    var v62 uint8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v61
    _ = v62
}
