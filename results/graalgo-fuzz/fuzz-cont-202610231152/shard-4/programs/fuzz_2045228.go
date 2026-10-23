package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 string; F1 int64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    v3 := []int{9223372036854775807}
    var v4 int16 = 0
    fmt.Println(v4)
    var v5 float32 = float32(1e10)
    v6 := GenS1{F0: (!true)}
    var v7 uint8 = 0
    fmt.Println(v4, v5, v7)
    var v8 float64 = (((3.14159 + (0.1 / -2.25)) * 1.5) / float64(v5))
    if false {
        _ = len(v1)
        switch 9223372036854775807 {
        case 666:
            v7 = 174
        case 910:
            fmt.Println(v7, v5)
        case 127:
            for fi9 := 0; fi9 < 17; fi9++ {
                for fi10 := 0; fi10 < 17; fi10++ {
                    switch (fi10 - (fi9 / (((fi9 << ((832) & 63))) | 1))) {
                    case -1:
                        v8 = 1e10
                    case 491:
                        v11 := GenS2{F0: v7, F1: 704, F2: v4}
                        _ = v11
                    case 154:
                        if (!false) {
                            delete(v1, 744)
                            var v12 float32 = v5
                            if (!true) {
                                if ((v12 - v12) >= v12) {
                                    v7 = v7
                                } else {
                                    v8 = 0.0
                                    v13 := GenS1{F0: (false && true)}
                                    var v14 uint8 = 159
                                    _ = v13
                                    _ = v14
                                }
                                var v15 string = "hello"
                                v5 = 1.5
                                _ = v15
                            } else {
                                _ = len(v1)
                                v8 += (v8 * 0.0)
                                switch "a" {
                                case "\n\t":
                                    var v16 int64 = -133
                                    _ = v16
                                case "hello":
                                    switch -9223372036854775808 {
                                    case 2147483647:
                                        var v17 uint16 = 32767
                                        var v18 bool = true
                                        _ = v17
                                        _ = v18
                                    case -13:
                                        v19 := GenS3{F0: (-(fi10))}
                                        _ = v19
                                    }
                                }
                                v7 = (^(v7))
                            }
                            v20 := GenS4{F0: "a", F1: -850}
                            if false {
                                fmt.Println(v5, v7, v4)
                            } else {
                                v21 := GenS3{F0: (^(fi9))}
                                _ = v21
                            }
                            _ = v12
                            _ = v20
                        }
                    }
                    var v22 uint8 = v7
                    switch fi10 {
                    case 32768:
                        if false {
                            _ = len(v1)
                            v23 := GenS5{F0: false, F1: (float32(1e10) < v5)}
                            if false {
                                var v24 float32 = (float32(-2.25) / (-2.25 - v5))
                                if (!(!(((v22 & v22) > 122) || true))) {
                                    v25 := GenS4{F0: "hello", F1: 724}
                                    v5 += (1e10 * (-2.25 + (v5 + v5)))
                                    _ = v25
                                }
                                _ = v24
                            } else {
                                var v26 string = "a"
                                fmt.Println(v22, v7)
                                v7 += (v22 % ((v7) | 1))
                                if ((fi9 % ((32768) | 1)) != (fi9 >> ((304) & 63))) {
                                    v1[42] = (1 != 1)
                                } else {
                                    v8 = 1e10
                                }
                                if (!true) {
                                    v7 -= v22
                                    v22 = (v22 << ((4294967295) & 63))
                                } else {
                                    v22 -= (v22 + 244)
                                    _ = len(v3)
                                    v26 = v26
                                    var v27 int8 = 0
                                    var v28 int8 = v27
                                    _ = v27
                                    _ = v28
                                }
                                _ = v26
                            }
                            _ = v23
                        } else {
                            var v29 uint = 42
                            _ = v29
                        }
                    case -9223372036854775808:
                        var v30 string = "a"
                        _ = v30
                    case -194:
                        if ((v22 << ((4294967295) & 63)) == (v7 * v22)) {
                            fmt.Println(v4)
                            if false {
                                v7 = (v7 / ((v22) | 1))
                            } else {
                                var v31 uint32 = 2147483648
                                var v32 bool = true
                                _ = v31
                                _ = v32
                            }
                            if (!false) {
                                var v33 int16 = (v4 ^ (v4 % ((v4) | 1)))
                                if true {
                                    v4 -= (v33 | v4)
                                }
                                var v34 int = -227
                                if (v22 <= v7) {
                                    var v35 int64 = 869
                                    fmt.Println(v8, v8)
                                    switch 42 {
                                    case -111:
                                        if false {
                                            var v36 int64 = v35
                                            _ = v36
                                        }
                                    case 128:
                                        v3 = append(v3, fi10)
                                    case 42:
                                        v22 = 70
                                    }
                                    v33 = v4
                                    v37 := GenS1{F0: (((v4 >> ((255) & 63)) == v33) && ((v33 * (v33 ^ v33)) != 42))}
                                    if (-9223372036854775808 <= (v35 / (((-(v35))) | 1))) {
                                        v38 := GenS4{F0: ("" + (((("hello" + "a") + "hello") + "foo bar") + "hello")), F1: 879}
                                        v3 = append(v3, v34)
                                        v8 = v8
                                        _ = v38
                                    } else {
                                        var v39 uint64 = 42
                                        v39 += (v39 / ((42) | 1))
                                        v3 = append(v3, fi10)
                                        _ = v39
                                    }
                                    switch "" {
                                    case "a":
                                        switch (v34 ^ 2147483647) {
                                        case 946:
                                            v40 := GenS5{F0: false, F1: false}
                                            _ = v40
                                        default:
                                            switch "a" {
                                            case "\n\t":
                                                var v41 uint = 2
                                                v42 := GenS4{F0: ("\n\t" + "foo bar"), F1: 1}
                                                v43 := GenS4{F0: ("\n\t" + "\n\t"), F1: int32(v41)}
                                                if true {
                                                    if true {
                                                        var v44 uint = v41
                                                        if len(v3) > 0 {
                                                            v3[0] = -1
                                                        }
                                                        _ = v44
                                                    } else {
                                                        if (!true) {
                                                            v41 -= (v41 ^ (^(v41)))
                                                            delete(v1, 477)
                                                        }
                                                        var v45 bool = ((!true) && true)
                                                        v1[1] = true
                                                        v8 = 0.0
                                                        _ = v45
                                                    }
                                                } else {
                                                    var v46 int16 = (v33 >> ((42) & 63))
                                                    if ((!true) && false) {
                                                        fmt.Println(fi9)
                                                    }
                                                    var v47 string = "foo bar"
                                                    _ = v46
                                                    _ = v47
                                                }
                                                var v48 int16 = (v33 >> ((65535) & 63))
                                                var v49 int64 = v35
                                                _ = v41
                                                _ = v42
                                                _ = v43
                                                _ = v48
                                                _ = v49
                                            default:
                                                var v50 float64 = (1e10 - v8)
                                                _ = v50
                                            }
                                        }
                                    case "foo bar":
                                        v4 -= v4
                                    case "\n\t":
                                        v3 = append(v3, -522)
                                    }
                                    _ = v35
                                    _ = v37
                                } else {
                                    fmt.Println(v5, v33, fi9)
                                }
                                _ = len(v2)
                                _ = v33
                                _ = v34
                            } else {
                                v22 += v7
                                v5 = 1.5
                                var v51 uint8 = 128
                                _ = v51
                            }
                        }
                    }
                    v7 -= (v22 - 0)
                    fmt.Println(fi10, fi9)
                    _ = fi10
                    _ = v22
                }
                for fi52 := 0; fi52 < 9; fi52++ {
                    if (-847 < -9223372036854775808) {
                        fmt.Println(fi52, v4, v8)
                        v4 += (v4 >> ((255) & 63))
                        v53 := GenS6{F0: "\n\t", F1: -276}
                        _ = v53
                    }
                    _ = fi52
                }
                var v54 uint8 = (v7 / ((v7) | 1))
                _ = fi9
                _ = v54
            }
        default:
            fmt.Println(v8)
        }
    }
    v3 = append(v3, 32767)
    v8 += (v8 + ((float64(v5) * v8) * (v8 / float64(v5))))
    var v55 int32 = 202
    v4 += -689
    if (!true) {
        v56 := v2[587]
        var v57 uint = 967
        _ = v56
        _ = v57
    } else {
        _ = len(v3)
        var v58 int = -2147483648
        switch (("" + ("" + "hello")) + "\n\t") {
        case "a":
            _ = len(v3)
        default:
            fmt.Println(v58)
        }
        _ = v58
    }
    var v59 bool = (42 != 1)
    var v60 string = "a"
    var v61 string = v60
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v7
    _ = v8
}
