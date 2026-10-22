package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint; F2 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 uint32; F1 uint32 }
func genFn0() uint8 {
    var v5 string = "\n\t"
    if false {
        var v6 int32 = 42
        _ = v6
    }
    v7 := []int{4294967295}
    _ = v5
    _ = v7
    return 40
}
func genFn1() int32 {
    var v8 bool = false
    fmt.Println(v8)
    _ = v8
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]int)
    var v3 int8 = 100
    v4 := []uint32{760, 0}
    var v9 int = 9223372036854775807
    if len(v4) > 0 {
        v4[0] = 2147483648
    }
    v1[42] = ("hello" + "a")
    fmt.Println(v3)
    var v10 float32 = 1.5
    var v11 bool = ((float64((v10 + v10)) * 3.14159) == 1e10)
    switch "foo bar" {
    case "hello":
        v3 += v3
    case "a":
        if v11 {
            v12 := GenS1{F0: 1000, F1: -715}
            if v11 {
                delete(v1, 2147483648)
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 int16 = -1
                    delete(v1, 2147483647)
                    _ = fi13
                    _ = v14
                }
                fmt.Println(v9)
            }
            _ = v12
        } else {
            v3 = v3
            var v15 int = v9
            _ = v15
        }
    case "":
        v16 := GenS2{F0: -954}
        _ = v16
    default:
        if v11 {
            var v17 int8 = (v3 * v3)
            _ = v17
        } else {
            if (0 != 42) {
                var v18 int = (^(v9))
                v11 = (-5 > 1)
                _ = v18
            } else {
                var v19 string = "a"
                fmt.Println(v19)
                if (0 != 2) {
                    if false {
                        for fi20 := 0; fi20 < 19; fi20++ {
                            v21 := GenS1{F0: 0, F1: 205}
                            _ = fi20
                            _ = v21
                        }
                        switch v19 {
                        case "foo bar":
                            v22 := GenS2{F0: 0}
                            _ = v22
                        case "":
                            v10 = v10
                        case "hello":
                            var v23 uint = 255
                            _ = v23
                        }
                        if v11 {
                            for fi24 := 0; fi24 < 6; fi24++ {
                                v25 := GenS3{F0: -2147483648, F1: 255, F2: "foo bar"}
                                var v26 uint8 = 42
                                var v27 uint32 = 2147483648
                                _ = fi24
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                        }
                    }
                    if true {
                        fmt.Println(v19)
                    } else {
                        for fi28 := 0; fi28 < 19; fi28++ {
                            v29 := GenS1{F0: 42, F1: 2147483647}
                            _ = fi28
                            _ = v29
                        }
                        var v30 string = (v19 + "\n\t")
                        _ = v30
                    }
                    if (1 != 4294967295) {
                        v19 = "a"
                        fmt.Println(v10, v3, v10)
                        var v31 uint8 = genFn0()
                        _ = v31
                    }
                } else {
                    fmt.Println(v3, v10, v19)
                    switch (v9 | (v9 * (v9 | v9))) {
                    case 0:
                        if v11 {
                            var v32 bool = v11
                            var v33 uint16 = 783
                            _ = v32
                            _ = v33
                        } else {
                            var v34 float64 = float64(float32(0.0))
                            v35 := GenS2{F0: 2147483647}
                            _ = v34
                            _ = v35
                        }
                    case -165:
                        switch "hello" {
                        case "a":
                            for fi36 := 0; fi36 < 20; fi36++ {
                                var v37 uint = 493
                                v11 = ((v3 > int8(fi36)) && (v3 >= v3))
                                v9 = v9
                                _ = fi36
                                _ = v37
                            }
                        }
                    case -2147483648:
                        v4 = append(v4, 2147483647)
                    }
                    v10 = (0.1 * ((0.1 * 0.1) - 1.5))
                }
                _ = v19
            }
            for fi38 := 0; fi38 < 20; fi38++ {
                v11 = (v10 > (3.14159 + float32((float64(((v10 * (v10 * v10)) - ((1e10 + 3.14159) * float32(((0.1 / 1e10) / 1e10))))) - 1.5))))
                if (v9 <= -2147483648) {
                    for fi39 := 0; fi39 < 17; fi39++ {
                        if (2147483647 == 2147483647) {
                            var v40 int64 = -215
                            v41 := GenS3{F0: -248, F1: 0, F2: (("\n\t" + "foo bar") + "a")}
                            v9 = (fi39 + (v9 ^ (fi38 << ((18446744073709551615) & 63))))
                            fmt.Println(fi39, v11, v3)
                            var v42 int = v9
                            var v43 string = "a"
                            _ = v40
                            _ = v41
                            _ = v42
                            _ = v43
                        } else {
                            if v11 {
                                if v11 {
                                    switch "hello" {
                                    case "foo bar":
                                        var v44 int8 = v3
                                        _ = v44
                                    default:
                                        fmt.Println(fi39, fi39)
                                    }
                                    var v45 bool = false
                                    _ = v45
                                }
                                switch "\n\t" {
                                case "hello":
                                    var v46 int64 = 1
                                    _ = v46
                                }
                                if v11 {
                                    v11 = v11
                                    fmt.Println(fi38)
                                    v10 -= v10
                                } else {
                                    var v47 int = -288
                                    _ = v47
                                }
                            }
                            v11 = false
                            v48 := GenS1{F0: 0, F1: -2147483648}
                            _ = v48
                        }
                        _ = len(v4)
                        if v11 {
                            v9 += v9
                            if len(v4) > 0 {
                                v4[0] = 895
                            }
                            if v11 {
                                switch (("hello" + "a") + "\n\t") {
                                case "":
                                    delete(v2, fi38)
                                case "a":
                                    var v49 uint16 = 847
                                    _ = v49
                                }
                                switch "a" {
                                case "a":
                                    if false {
                                        var v50 string = "a"
                                        v51 := GenS1{F0: 42, F1: -2147483648}
                                        var v52 int32 = 0
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                    }
                                case "":
                                    if (!(42 <= -128)) {
                                        v10 -= (-2.25 * 3.14159)
                                    } else {
                                        delete(v2, (fi38 % (((-(fi38))) | 1)))
                                    }
                                    if (1000 <= 685) {
                                        if (v11 && true) {
                                            _ = len(v1)
                                        } else {
                                            _ = len(v1)
                                            v53 := GenS4{F0: 1}
                                            _ = v53
                                        }
                                    }
                                case "\n\t":
                                    var v54 uint8 = 127
                                    _ = v54
                                default:
                                    v9 -= (fi39 % ((9223372036854775807) | 1))
                                }
                                if (true && ((!true) && (v11 || v11))) {
                                    v11 = v11
                                    var v55 uint64 = 4294967295
                                    _ = v55
                                }
                            } else {
                                _ = len(v2)
                                var v56 uint16 = 42
                                _ = v56
                            }
                            var v57 uint64 = 4294967295
                            _ = v57
                        }
                        _ = fi39
                    }
                    for fi58 := 0; fi58 < 5; fi58++ {
                        fmt.Println(v10, fi38, v3)
                        var v59 uint16 = 255
                        _ = fi58
                        _ = v59
                    }
                    v11 = false
                } else {
                    for fi60 := 0; fi60 < 12; fi60++ {
                        v61 := GenS1{F0: 0, F1: 42}
                        _ = fi60
                        _ = v61
                    }
                    v62 := GenS5{F0: (!((v3 ^ 1) >= -1)), F1: (fi38 - (-(v9)))}
                    _ = v62
                }
                _ = fi38
            }
            for fi63 := 0; fi63 < 18; fi63++ {
                v11 = v11
                var v64 int32 = 2147483647
                var v65 uint32 = 1
                _ = fi63
                _ = v64
                _ = v65
            }
        }
    }
    v66 := GenS6{F0: 0, F1: 65535}
    var v67 uint = 1
    v9 = (v9 | v9)
    if v11 {
        v3 += v3
        _ = len(v4)
    }
    _ = len(v1)
    v9 -= (-(v9))
    v67 -= 65535
    for fi68 := 0; fi68 < 6; fi68++ {
        for fi69 := 0; fi69 < 15; fi69++ {
            fmt.Println(fi69, fi68)
            _ = fi69
        }
        var v70 string = "hello"
        _ = fi68
        _ = v70
    }
    var v71 int = int(v67)
    delete(v2, v71)
    if len(v4) > 0 {
        v4[0] = 818
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v67:", v67)
    fmt.Println("v71:", v71)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v66
    _ = v67
    _ = v71
    _ = v9
}
