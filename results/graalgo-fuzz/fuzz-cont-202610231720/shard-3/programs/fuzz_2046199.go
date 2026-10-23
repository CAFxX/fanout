package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 int8; F1 int; F2 string }
type GenS6 struct { F0 string; F1 string; F2 int64 }
type GenS7 struct { F0 uint32; F1 uint16 }
type GenS8 struct { F0 string; F1 int }
func genFn0() int64 {
    v6 := []int{4294967295, 42, -693, 32767}
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint = 2
        v8 -= v8
        _ = fi7
        _ = v8
    }
    var v9 float32 = float32(0.0)
    _ = v6
    _ = v9
    return 0
}
func genFn1() uint {
    var v10 int = 2147483647
    for fi11 := 0; fi11 < 5; fi11++ {
        v10 += (v10 | (fi11 * 128))
        v10 -= (fi11 * v10)
        var v12 uint8 = 190
        _ = fi11
        _ = v12
    }
    switch "a" {
    case "foo bar":
        if (true && true) {
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 uint16 = 32767
                _ = fi13
                _ = v14
            }
            switch (v10 & v10) {
            case -453:
                for fi15 := 0; fi15 < 18; fi15++ {
                    var v16 int = (^(fi15))
                    var v17 int16 = -1
                    for fi18 := 0; fi18 < 1; fi18++ {
                        var v19 uint16 = 705
                        if true {
                            v20 := GenS1{F0: "\n\t", F1: "\n\t"}
                            _ = v20
                        }
                        _ = fi18
                        _ = v19
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            case 128:
                if (255 < 0) {
                    v10 += (-(v10))
                    v21 := []int{(v10 % ((-1) | 1))}
                    _ = v21
                } else {
                    if (2 > 4294967295) {
                        var v22 int8 = 127
                        _ = v22
                    }
                    var v23 uint32 = 2147483648
                    _ = v23
                }
            case 42:
                fmt.Println(v10)
            }
        } else {
            switch v10 {
            case 32767:
                for fi24 := 0; fi24 < 18; fi24++ {
                    var v25 int16 = 376
                    _ = fi24
                    _ = v25
                }
            case 42:
                var v26 float64 = 0.1
                _ = v26
            default:
                switch "a" {
                case "foo bar":
                    v27 := GenS2{F0: 2147483647, F1: 1}
                    _ = v27
                default:
                    if (!true) {
                        for fi28 := 0; fi28 < 18; fi28++ {
                            for fi29 := 0; fi29 < 5; fi29++ {
                                v30 := []int{}
                                if len(v30) > 0 {
                                    v30[0] = v10
                                }
                                if true {
                                    v31 := GenS3{F0: 4294967295}
                                    v10 -= fi29
                                    if false {
                                        switch "hello" {
                                        case "hello":
                                            if false {
                                                var v32 uint8 = 2
                                                v10 += 2147483647
                                                _ = v32
                                            } else {
                                                switch fi28 {
                                                case 9223372036854775807:
                                                    var v33 uint32 = 2
                                                    _ = v33
                                                case 127:
                                                    fmt.Println(fi29)
                                                }
                                            }
                                        case "":
                                            _ = len(v30)
                                            fmt.Println(v10)
                                            v10 = fi28
                                        case "a":
                                            var v34 uint64 = 0
                                            _ = v34
                                        }
                                    }
                                    _ = v31
                                }
                                v35 := GenS4{F0: 0, F1: 65535}
                                _ = fi29
                                _ = v30
                                _ = v35
                            }
                            _ = fi28
                        }
                        fmt.Println(v10)
                        v10 += 128
                    } else {
                        v36 := []int{-442, (v10 ^ v10), (v10 >> ((65535) & 63)), v10}
                        v37 := GenS5{F0: 0, F1: 4294967295, F2: ""}
                        _ = v36
                        _ = v37
                    }
                }
            }
            fmt.Println(v10)
        }
    case "hello":
        v38 := GenS4{F0: 0, F1: 508}
        _ = v38
    default:
        switch -70 {
        case 4294967295:
            fmt.Println(v10)
        case -217:
            v39 := GenS5{F0: 100, F1: (v10 + v10), F2: "a"}
            _ = v39
        }
    }
    _ = v10
    return 207
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]uint32)
    var v3 bool = (!(false || true))
    var v4 float64 = v1
    var v5 int32 = 742
    v40 := []int{66}
    if v3 {
        for fi41 := 0; fi41 < 2; fi41++ {
            fmt.Println(v3)
            _ = fi41
        }
        v1 = 1.5
    } else {
        switch 32768 {
        case -440:
            switch -1 {
            case 32767:
                var v42 uint32 = 42
                _ = v42
            case 32768:
                for fi43 := 0; fi43 < 19; fi43++ {
                    v44 := GenS6{F0: ("\n\t" + "\n\t"), F1: "foo bar", F2: -324}
                    for fi45 := 0; fi45 < 7; fi45++ {
                        v46 := GenS6{F0: ("" + ("\n\t" + "foo bar")), F1: "a", F2: 885}
                        v1 = v1
                        var v47 int32 = (v5 * v5)
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                    _ = fi43
                    _ = v44
                }
            default:
                v48 := GenS7{F0: 2147483647, F1: 65535}
                _ = v48
            }
        case 32767:
            for fi49 := 0; fi49 < 10; fi49++ {
                fmt.Println(v4)
                v1 = 0.1
                for fi50 := 0; fi50 < 12; fi50++ {
                    v51 := GenS3{F0: fi49}
                    v1 = v4
                    _ = fi50
                    _ = v51
                }
                _ = fi49
            }
        }
        var v52 uint8 = 128
        v52 += v52
        _ = v52
    }
    for fi53 := 0; fi53 < 8; fi53++ {
        v4 = v4
        _ = fi53
    }
    var v54 uint64 = 0
    var v55 uint32 = 0
    var v56 int = 1
    v5 = v5
    if ((491 >= 18446744073709551615) && (v3 || v3)) {
        switch v56 {
        case -9223372036854775808:
            var v57 float64 = (float64((0.0 / 1.5)) * v4)
            _ = v57
        case 42:
            v58 := v2[v56]
            _ = v58
        case -772:
            fmt.Println(v4, v56)
        default:
            v59 := GenS6{F0: "\n\t", F1: "", F2: -1}
            _ = v59
        }
        v1 = ((1e10 * (0.1 + (1e10 / 0.1))) * v1)
    } else {
        fmt.Println(v55, v55)
        fmt.Println(v1, v54, v56)
    }
    if false {
        for fi60 := 0; fi60 < 11; fi60++ {
            v54 -= (v54 << ((1) & 63))
            _ = fi60
        }
        if (v54 == 2) {
            v61 := GenS5{F0: int8(v5), F1: v56, F2: "a"}
            _ = v61
        } else {
            var v62 float32 = -2.25
            var v63 float32 = v62
            _ = v62
            _ = v63
        }
        v64 := GenS5{F0: 127, F1: v56, F2: ("a" + ("a" + "\n\t"))}
        _ = v64
    } else {
        if len(v40) >= 2 {
            v40 = v40[1:len(v40)]
        }
        v65 := GenS8{F0: (("" + ("" + "hello")) + ""), F1: 4294967295}
        _ = v65
    }
    v66 := GenS4{F0: genFn0(), F1: v55}
    _ = len(v2)
    var v67 float32 = 0.1
    var v68 int16 = 0
    var v69 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v66
    _ = v67
    _ = v68
    _ = v69
}
