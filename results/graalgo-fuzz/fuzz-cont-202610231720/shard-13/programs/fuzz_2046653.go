package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool }
func genFn0() uint {
    var v7 uint8 = 127
    v7 = v7
    v7 = 127
    v8 := []int{-442, -918}
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int64, p1 uint8) uint8 {
    v9 := []int{127, 9223372036854775807}
    var v10 bool = (!true)
    switch "" {
    case "foo bar":
        if len(v9) > 0 {
            v9[0] = 32767
        }
    case "\n\t":
        switch -714 {
        case -1:
            if (-669 <= 127) {
                for fi11 := 0; fi11 < 15; fi11++ {
                    for fi12 := 0; fi12 < 4; fi12++ {
                        fmt.Println(p1, fi12)
                        _ = fi12
                    }
                    _ = fi11
                }
            } else {
                var v13 string = (("" + "") + ("foo bar" + "hello"))
                _ = v13
            }
        case 42:
            fmt.Println(p0, v10, p0)
        case 128:
            p1 = (p1 << ((836) & 63))
        }
    }
    v14 := GenS1{F0: 128}
    _ = p0
    _ = p1
    _ = v10
    _ = v14
    _ = v9
    return (p1 ^ 128)
}

func main() {
    var v1 uint64 = 117
    v2 := []string{"hello", "", ("hello" + ("a" + ("foo bar" + ("\n\t" + (("\n\t" + ("a" + "a")) + ""))))), "foo bar"}
    v3 := make(map[uint32]string)
    var v4 int = 32768
    v5 := make(map[int]int)
    v6 := []uint32{328, 42}
    switch v4 {
    case 128:
        var v15 uint16 = 1
        _ = v15
    case -548:
        var v16 int = 1
        _ = v16
    }
    if (false && (uint32(v1) > 1)) {
        v4 -= v4
        if (218 <= 96) {
            if false {
                for fi17 := 0; fi17 < 3; fi17++ {
                    v18 := GenS1{F0: fi17}
                    _ = fi17
                    _ = v18
                }
            }
            delete(v3, 260)
        }
        v1 = 18446744073709551615
    }
    v19 := GenS1{F0: v4}
    for fi20 := 0; fi20 < 15; fi20++ {
        if ("hello" == ("\n\t" + "")) {
            v21 := GenS2{F0: 42, F1: -96, F2: 121}
            var v22 bool = (!(65535 == 2))
            var v23 float32 = 0.0
            _ = v21
            _ = v22
            _ = v23
        } else {
            switch -1 {
            case -709:
                fmt.Println(fi20, v1)
            }
            v1 += (v1 | (v1 | (v1 + 2)))
            v1 -= (v1 ^ v1)
        }
        _ = fi20
    }
    v4 = (v4 / (((v4 ^ v4)) | 1))
    var v24 string = ""
    if ((42 > 32767) && false) {
        v4 = 128
        if false {
            var v25 uint64 = (v1 + v1)
            _ = v25
        }
        var v26 int16 = 0
        _ = v26
    } else {
        if ((false || false) || (v24 != (((("\n\t" + ("a" + v24)) + v24) + v24) + v24))) {
            var v27 uint8 = 128
            _ = v27
        } else {
            if (v24 == v24) {
                v4 += (^(v4))
                if ((-1 >= 42) && false) {
                    delete(v3, 238)
                    v24 = v24
                    v4 = (v4 << ((0) & 63))
                }
            } else {
                fmt.Println(v1, v1, v1)
            }
            var v28 uint16 = 255
            var v29 uint = 2
            _ = v28
            _ = v29
        }
        if (0 < 270) {
            for fi30 := 0; fi30 < 14; fi30++ {
                switch v4 {
                case -712:
                    if (42 > -1) {
                        var v31 int64 = 9223372036854775807
                        v32 := GenS3{F0: 1}
                        v33 := GenS2{F0: 42, F1: 127, F2: 1}
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                case 2147483647:
                    v34 := GenS4{F0: 1, F1: (v4 * v4)}
                    _ = v34
                case -1:
                    var v35 int32 = -235
                    _ = v35
                default:
                    if true {
                        v36 := v5[(v4 << ((18446744073709551615) & 63))]
                        _ = v36
                    }
                }
                if (((v24 + ("\n\t" + v24)) + "foo bar") >= ("hello" + ("\n\t" + "hello"))) {
                    for fi37 := 0; fi37 < 10; fi37++ {
                        v4 += fi37
                        fmt.Println(v4)
                        _ = fi37
                    }
                    var v38 uint64 = v1
                    fmt.Println(v4, v1, v24)
                    _ = v38
                }
                v39 := GenS2{F0: 127, F1: -1, F2: 1}
                _ = fi30
                _ = v39
            }
            if (0 > v1) {
                if (!(283 < 2147483647)) {
                    fmt.Println(v4, v24)
                    var v40 uint = 39
                    _ = v40
                } else {
                    v24 = v24
                }
            } else {
                switch v24 {
                case "a":
                    for fi41 := 0; fi41 < 10; fi41++ {
                        v42 := GenS5{F0: (v1 / (((^(v1))) | 1))}
                        _ = fi41
                        _ = v42
                    }
                case "foo bar":
                    var v43 uint32 = 726
                    _ = v43
                case "hello":
                    for fi44 := 0; fi44 < 13; fi44++ {
                        for fi45 := 0; fi45 < 8; fi45++ {
                            var v46 uint16 = 2
                            v24 = "a"
                            if (!true) {
                                fmt.Println(fi45, v1)
                                v1 = 42
                                if ((-552 <= 0) && (2147483647 >= 0)) {
                                    var v47 int16 = -589
                                    _ = v47
                                } else {
                                    if false {
                                        if (fi44 <= (fi45 & (fi44 * fi45))) {
                                            v46 = 0
                                        } else {
                                            switch v24 {
                                            case "hello":
                                                v1 = (v1 * 313)
                                            default:
                                                v46 = v46
                                            }
                                        }
                                        delete(v5, (fi45 << ((4294967295) & 63)))
                                        v48 := GenS6{F0: (true && true)}
                                        _ = v48
                                    }
                                }
                            }
                            _ = fi45
                            _ = v46
                        }
                        for fi49 := 0; fi49 < 10; fi49++ {
                            var v50 int = v4
                            _ = fi49
                            _ = v50
                        }
                        v4 += -895
                        _ = fi44
                    }
                default:
                    v24 = v24
                }
            }
        }
        v4 = (v4 - (v4 / (((v4 / ((32768) | 1))) | 1)))
    }
    var v51 uint32 = 945
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
