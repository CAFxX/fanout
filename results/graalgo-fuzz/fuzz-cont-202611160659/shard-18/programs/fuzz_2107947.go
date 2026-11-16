package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32 }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 0
    v3 := make(map[uint32]int)
    var v4 string = (("a" + "\n\t") + "\n\t")
    var v5 uint64 = 543
    var v6 uint8 = 1
    v7 := GenS1{F0: 42, F1: 877, F2: 0}
    v8 := []int{-9223372036854775808, 4294967295, 4294967295}
    v9 := GenS1{F0: 1, F1: 4294967295, F2: 1}
    for fi10 := 0; fi10 < 1; fi10++ {
        if (!false) {
            fmt.Println(v6, v6, v5)
            v6 -= 127
        }
        v6 += v6
        switch "hello" {
        case "hello":
            if (false && (!false)) {
                for fi11 := 0; fi11 < 14; fi11++ {
                    v5 -= v5
                    v2 = 557
                    _ = fi11
                }
                var v12 int64 = 42
                _ = v12
            }
        case "":
            var v13 uint8 = v6
            _ = v13
        case "\n\t":
            var v14 float32 = 1e10
            _ = v14
        default:
            var v15 bool = true
            _ = v15
        }
        _ = fi10
    }
    v16 := GenS2{F0: (v2 | v2), F1: false}
    if false {
        if true {
            switch 9223372036854775807 {
            case -2147483648:
                v6 = 127
            case 32767:
                if ((0.0 + ((-2.25 / -2.25) / 1.0)) >= 1.5) {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    v4 = "hello"
                    var v17 uint8 = (v6 ^ (v6 | v6))
                    _ = v17
                }
            }
            v18 := GenS3{F0: v2, F1: -1}
            v5 -= (v5 * v5)
            _ = v18
        }
        v19 := GenS3{F0: v2, F1: 1}
        _ = v19
    } else {
        if (65535 == 42) {
            for fi20 := 0; fi20 < 14; fi20++ {
                v5 += v5
                var v21 uint64 = v5
                var v22 int16 = -1
                _ = fi20
                _ = v21
                _ = v22
            }
            var v23 uint32 = 23
            if (9223372036854775807 != 230) {
                for fi24 := 0; fi24 < 3; fi24++ {
                    v25 := GenS4{F0: int32(v5)}
                    v6 = 128
                    _ = fi24
                    _ = v25
                }
                var v26 int64 = 9223372036854775807
                _ = v26
            } else {
                fmt.Println(v2, v4, v5)
                v27 := GenS2{F0: 1, F1: false}
                v23 = v23
                _ = v27
            }
            _ = v23
        }
        for fi28 := 0; fi28 < 7; fi28++ {
            v4 = v4
            v2 += (v2 << ((4294967295) & 63))
            switch (^(fi28)) {
            case -592:
                var v29 float32 = 0.1
                _ = v29
            case 1:
                v4 = ("\n\t" + "\n\t")
            case 493:
                for fi30 := 0; fi30 < 18; fi30++ {
                    if (v6 < v6) {
                        v6 = (v6 + v6)
                        var v31 int32 = v2
                        if false {
                            v32 := GenS5{F0: (v31 % ((-1) | 1)), F1: 255}
                            var v33 int16 = 1000
                            if true {
                                v6 = (v6 % ((30) | 1))
                            } else {
                                v34 := GenS2{F0: (v2 >> ((65535) & 63)), F1: (true || true)}
                                if true {
                                    delete(v1, "hello")
                                }
                                _ = v34
                            }
                            _ = v32
                            _ = v33
                        } else {
                            fmt.Println(v31, fi30)
                            switch "a" {
                            case "":
                                if true {
                                    v5 = v5
                                }
                                if (992 >= 32767) {
                                    v35 := GenS5{F0: 770, F1: 2}
                                    _ = v35
                                } else {
                                    fmt.Println(v2, v5, v2)
                                    delete(v1, "a")
                                    var v36 int32 = (v2 / (((v31 - v2)) | 1))
                                    _ = v36
                                }
                            case "foo bar":
                                if ((2 < uint(v5)) && (((v4 + v4) + v4) >= ("foo bar" + v4))) {
                                    v4 = ("hello" + (("hello" + "hello") + v4))
                                    if false {
                                        var v37 uint8 = 128
                                        if (!true) {
                                            var v38 uint8 = (v6 | 1)
                                            fmt.Println(fi28, fi30, v37)
                                            _ = v38
                                        }
                                        var v39 int8 = 33
                                        var v40 int = -1
                                        _ = v37
                                        _ = v39
                                        _ = v40
                                    } else {
                                        var v41 float64 = 1e10
                                        var v42 bool = ((v41 / 1e10) == -2.25)
                                        if (v42 || v42) {
                                            var v43 float64 = (v41 - ((v41 / float64((1e10 - float32(((3.14159 / v41) - ((1e10 - v41) / v41)))))) / 1.0))
                                            _ = v43
                                        }
                                        _ = v41
                                        _ = v42
                                    }
                                    if (0.1 == 3.14159) {
                                        if true {
                                            var v44 uint8 = v6
                                            v45 := GenS2{F0: (v31 | (v31 | v2)), F1: (true && (32767 > fi30))}
                                            _ = v44
                                            _ = v45
                                        }
                                    } else {
                                        v46 := GenS2{F0: v31, F1: (65535 > 18446744073709551615)}
                                        _ = v46
                                    }
                                }
                            }
                        }
                        _ = v31
                    } else {
                        if false {
                            var v47 uint16 = 32767
                            fmt.Println(v47, fi28, fi28)
                            _ = v47
                        } else {
                            _ = len(v3)
                            v5 -= 375
                            switch fi30 {
                            case -1:
                                v48 := GenS1{F0: 460, F1: 500, F2: 100}
                                _ = v48
                            case 32767:
                                v49 := GenS6{F0: fi30}
                                _ = v49
                            }
                        }
                        if true {
                            fmt.Println(fi30)
                        }
                    }
                    v50 := GenS4{F0: (v2 / ((v2) | 1))}
                    _ = fi30
                    _ = v50
                }
            default:
                fmt.Println(v5)
            }
            _ = fi28
        }
    }
    switch v4 {
    case "":
        var v51 string = ((v4 + (v4 + v4)) + "")
        _ = v51
    case "foo bar":
        for fi52 := 0; fi52 < 18; fi52++ {
            if true {
                fmt.Println(v5)
            } else {
                for fi53 := 0; fi53 < 10; fi53++ {
                    v54 := GenS7{F0: 127}
                    v5 = 18446744073709551615
                    v4 = ((v4 + (v4 + (((v4 + (("foo bar" + v4) + v4)) + (((v4 + "a") + v4) + v4)) + "hello"))) + "foo bar")
                    _ = fi53
                    _ = v54
                }
            }
            var v55 int16 = 1
            _ = fi52
            _ = v55
        }
    default:
        for fi56 := 0; fi56 < 13; fi56++ {
            fmt.Println(v6)
            _ = fi56
        }
    }
    if (!(255 >= v6)) {
        _ = len(v8)
        if len(v8) > 0 {
            v8[0] = 452
        }
        var v57 uint = 42
        _ = v57
    } else {
        var v58 int8 = -128
        _ = v58
    }
    for fi59 := 0; fi59 < 13; fi59++ {
        v60 := GenS8{F0: v2}
        _ = fi59
        _ = v60
    }
    v61 := GenS7{F0: v6}
    v62 := v1["hello"]
    for fi63 := 0; fi63 < 8; fi63++ {
        v4 = (v4 + "foo bar")
        for fi64 := 0; fi64 < 14; fi64++ {
            v6 += (v6 + (v6 >> ((0) & 63)))
            var v65 float32 = 1e10
            _ = fi64
            _ = v65
        }
        _ = fi63
    }
    if (false || ((1 != 32767) && true)) {
        v5 = v5
    }
    var v66 int8 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v66:", v66)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
    _ = v66
    _ = v7
    _ = v8
    _ = v9
}
