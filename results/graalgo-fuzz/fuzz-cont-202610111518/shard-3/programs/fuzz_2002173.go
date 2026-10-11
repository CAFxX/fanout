package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 string; F1 string; F2 int }

func main() {
    var v1 uint64 = 281
    var v2 uint64 = (^(v1))
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 65535
    var v5 float64 = (0.0 + (1.5 + ((3.14159 / (-2.25 - 0.1)) * (((float64((1.5 * 1.5)) - 0.1) - -2.25) - (1.5 - -2.25)))))
    var v6 uint8 = 128
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            if false {
                v2 = 18446744073709551615
            } else {
                v5 += 1.5
                var v9 int64 = -817
                var v10 uint32 = 2147483647
                if false {
                    _ = len(v3)
                    v6 -= (v6 << ((876) & 63))
                    v11 := GenS1{F0: 0, F1: ((fi8 | (-(fi7))) >= (fi8 % (((-(fi8))) | 1)))}
                    switch 32767 {
                    case -2147483648:
                        v2 += v2
                    case 42:
                        if true {
                            v12 := GenS1{F0: 85, F1: false}
                            v2 = (v2 | uint64(v10))
                            _ = v12
                        } else {
                            v13 := GenS1{F0: 0, F1: (!true)}
                            _ = v13
                        }
                    case 394:
                        _ = len(v3)
                    }
                    v14 := GenS2{F0: false, F1: 124}
                    _ = v11
                    _ = v14
                } else {
                    switch "" {
                    case "a":
                        v15 := []int{}
                        _ = v15
                    case "":
                        v16 := []int{4294967295}
                        _ = v16
                    case "foo bar":
                        var v17 int = (fi8 / ((0) | 1))
                        _ = v17
                    default:
                        if (-1 != -9223372036854775808) {
                            v6 = v6
                            if (-2.25 == -2.25) {
                                v4 += v4
                                var v18 uint64 = uint64(v4)
                                v18 = 2
                                _ = v18
                            } else {
                                var v19 uint64 = 18446744073709551615
                                v4 -= v4
                                fmt.Println(fi7, fi7)
                                _ = v19
                            }
                        }
                    }
                    var v20 uint = uint(v1)
                    switch "hello" {
                    case "foo bar":
                        v2 = v2
                    default:
                        var v21 uint8 = v6
                        _ = v21
                    }
                    _ = v20
                }
                if (1 != -2147483648) {
                    if true {
                        if (!true) {
                            if false {
                                if true {
                                    var v22 float64 = 0.0
                                    _ = v22
                                } else {
                                    fmt.Println(v2)
                                    if false {
                                        switch "hello" {
                                        case "\n\t":
                                            if false {
                                                v23 := v3[v10]
                                                v5 += 0.1
                                                _ = v23
                                            } else {
                                                var v24 float64 = ((1e10 * v5) * v5)
                                                var v25 uint = 255
                                                _ = v24
                                                _ = v25
                                            }
                                        }
                                        switch 2147483647 {
                                        case 547:
                                            fmt.Println(v9)
                                            v26 := GenS3{F0: uint64(v4), F1: "hello"}
                                            if ((v10 & uint32(v6)) == uint32(v4)) {
                                                v5 = 3.14159
                                                v6 = (v6 / ((v6) | 1))
                                                v27 := []int{(^(fi7)), 32768, -2147483648}
                                                _ = v27
                                            }
                                            _ = v26
                                        }
                                    }
                                }
                            }
                        } else {
                            switch "foo bar" {
                            case "a":
                                if false {
                                    if false {
                                        var v28 int = 32767
                                        var v29 uint = 159
                                        v3[v10] = (^(v10))
                                        delete(v3, v10)
                                        v29 += (v29 - 42)
                                        _ = v28
                                        _ = v29
                                    } else {
                                        var v30 float32 = (1.5 / 1.5)
                                        _ = v30
                                    }
                                    if (true && (!((false || false) && (true && false)))) {
                                        var v31 bool = true
                                        var v32 string = ""
                                        switch fi8 {
                                        case 32767:
                                            v9 -= (v9 * v9)
                                            fmt.Println(v4, fi7)
                                        case 32768:
                                            var v33 int16 = 0
                                            _ = v33
                                        case 1:
                                            fmt.Println(fi8, v31)
                                        }
                                        v32 = "\n\t"
                                        _ = v31
                                        _ = v32
                                    } else {
                                        var v34 uint16 = uint16(v6)
                                        _ = v34
                                    }
                                    if (false || (-511 != 0)) {
                                        delete(v3, v10)
                                        v5 = v5
                                    } else {
                                        v35 := GenS4{F0: (("a" + "hello") + "foo bar"), F1: "a", F2: fi7}
                                        v36 := GenS2{F0: ((((v5 + v5) / 3.14159) < float64(0.0)) || ((^(v1)) < v2)), F1: 374}
                                        v37 := []int{fi8, 2147483647, fi8}
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                } else {
                                    switch fi7 {
                                    case 32767:
                                        v38 := []int{}
                                        _ = v38
                                    case 128:
                                        if (1 <= 127) {
                                            v10 += (^(v10))
                                            var v39 uint64 = 42
                                            _ = v39
                                        } else {
                                            delete(v3, v10)
                                            v40 := v3[v10]
                                            v5 -= 3.14159
                                            var v41 int = 127
                                            _ = v40
                                            _ = v41
                                        }
                                    default:
                                        v3[4294967295] = (v10 + v10)
                                    }
                                }
                            case "hello":
                                if ((v2 > v1) && ((v4 % ((v4) | 1)) == v4)) {
                                    var v42 uint16 = (^(v4))
                                    _ = v42
                                }
                            }
                            v10 = v10
                            fmt.Println(v5)
                        }
                        if (true && ((1 == 42) || false)) {
                            v43 := []int{0, fi8, 9223372036854775807}
                            v44 := GenS1{F0: 1, F1: (65535 >= 2)}
                            var v45 float64 = float64((3.14159 / 3.14159))
                            _ = v43
                            _ = v44
                            _ = v45
                        }
                        var v46 uint64 = 18446744073709551615
                        switch "foo bar" {
                        case "a":
                            var v47 int8 = -128
                            _ = v47
                        case "\n\t":
                            v48 := []int{fi8, fi8, fi8}
                            _ = v48
                        default:
                            v49 := GenS3{F0: 18446744073709551615, F1: "hello"}
                            _ = v49
                        }
                        _ = v46
                    }
                    var v50 int64 = 42
                    _ = v50
                }
                _ = v10
                _ = v9
            }
            v51 := v3[1]
            _ = fi8
            _ = v51
        }
        for fi52 := 0; fi52 < 1; fi52++ {
            v53 := v3[937]
            var v54 int16 = 1
            _ = fi52
            _ = v53
            _ = v54
        }
        _ = fi7
    }
    var v55 int = 2147483647
    fmt.Println(v4)
    for fi56 := 0; fi56 < 3; fi56++ {
        fmt.Println(v4, fi56, v2)
        var v57 float64 = (v5 - -2.25)
        _ = fi56
        _ = v57
    }
    v5 += v5
    var v58 int16 = 1
    var v59 bool = true
    fmt.Println(v6, v59)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v58
    _ = v59
    _ = v6
}
