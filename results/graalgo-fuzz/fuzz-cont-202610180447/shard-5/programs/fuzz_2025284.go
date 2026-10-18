package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 int32; F2 int16 }
type GenS5 struct { F0 int; F1 int16 }

func main() {
    var v1 uint = 42
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]bool)
    var v4 uint64 = 4294967295
    switch ("hello" + "hello") {
    case "foo bar":
        if (!(128 <= 2)) {
            switch ("hello" + "") {
            case "a":
                switch ("foo bar" + "\n\t") {
                case "foo bar":
                    v5 := GenS1{F0: (v4 / ((v4) | 1)), F1: v4, F2: -2147483648}
                    _ = v5
                case "\n\t":
                    fmt.Println(v4, v4, v4)
                default:
                    v6 := GenS2{F0: -377}
                    _ = v6
                }
            case "":
                fmt.Println(v4)
            case "\n\t":
                switch "a" {
                case "":
                    var v7 bool = true
                    _ = v7
                case "hello":
                    v8 := GenS1{F0: v4, F1: 18446744073709551615, F2: 1}
                    _ = v8
                default:
                    if true {
                        var v9 string = "hello"
                        _ = v9
                    }
                }
            default:
                v2 = 608
            }
            var v10 float64 = 3.14159
            _ = v10
        } else {
            v11 := []int{}
            fmt.Println(v4)
            fmt.Println(v1)
            _ = v11
        }
    case "a":
        switch "hello" {
        case "a":
            v4 = (^(v4))
        case "hello":
            v3[(^(v2))] = false
        case "":
            if ((754 != 4294967295) && false) {
                v2 -= v2
                v12 := GenS1{F0: v4, F1: 42, F2: 128}
                switch "\n\t" {
                case "a":
                    for fi13 := 0; fi13 < 9; fi13++ {
                        var v14 int32 = -860
                        var v15 int32 = (^(v14))
                        var v16 float64 = -2.25
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                case "\n\t":
                    v17 := GenS1{F0: 18446744073709551615, F1: (v4 >> ((v1) & 63)), F2: 549}
                    _ = v17
                case "hello":
                    if false {
                        v1 = 255
                    } else {
                        var v18 int32 = -332
                        var v19 int64 = 1
                        fmt.Println(v18)
                        _ = v18
                        _ = v19
                    }
                }
                _ = v12
            } else {
                switch 32768 {
                case -364:
                    v20 := GenS3{F0: 39}
                    _ = v20
                case 9223372036854775807:
                    if (true || false) {
                        var v21 uint32 = 2147483648
                        for fi22 := 0; fi22 < 9; fi22++ {
                            fmt.Println(v2, v1)
                            _ = fi22
                        }
                        _ = v21
                    } else {
                        v1 -= (v1 / (((v1 << ((0) & 63))) | 1))
                        v3[v2] = false
                    }
                case 325:
                    switch 0 {
                    case 42:
                        v23 := []int{86, -137, 128, 32767}
                        _ = v23
                    case 2147483647:
                        v4 = 2
                    }
                }
                fmt.Println(v4, v4)
                v1 -= 1
            }
        }
    }
    var v24 float32 = 0.1
    if (true || true) {
        var v25 int64 = -480
        _ = v25
    } else {
        v4 = uint64(v1)
    }
    if false {
        v24 += v24
    }
    var v26 uint = v1
    if (-1 != -631) {
        v1 -= 511
    } else {
        for fi27 := 0; fi27 < 4; fi27++ {
            if false {
                var v28 bool = ((0 > 2) || false)
                switch fi27 {
                case 948:
                    fmt.Println(v24, fi27, v4)
                case 1:
                    v29 := []int{42, -56}
                    _ = v29
                case -1:
                    if v28 {
                        if (fi27 >= 4294967295) {
                            v26 -= 4294967295
                            v2 = v2
                        } else {
                            if v28 {
                                v30 := GenS4{F0: ((((((("foo bar" + "hello") + ("a" + "foo bar")) + "\n\t") + "\n\t") + "\n\t") + "") + "foo bar"), F1: -546, F2: -32768}
                                if ((119 != v4) || (v26 <= 255)) {
                                    var v31 int = fi27
                                    fmt.Println(v4, v4, v24)
                                    _ = v31
                                } else {
                                    for fi32 := 0; fi32 < 20; fi32++ {
                                        v33 := []int{}
                                        v34 := v3[v2]
                                        switch (fi32 | (^(fi27))) {
                                        case 2147483647:
                                            v35 := GenS3{F0: 42}
                                            _ = v35
                                        case 127:
                                            switch int(v4) {
                                            case 537:
                                                var v36 int8 = -52
                                                _ = v36
                                            case 42:
                                                if len(v33) > 0 {
                                                    v33[0] = (fi32 | (-(fi27)))
                                                }
                                            }
                                        }
                                        _ = fi32
                                        _ = v33
                                        _ = v34
                                    }
                                }
                                var v37 int32 = 2147483647
                                _ = v30
                                _ = v37
                            } else {
                                v4 = v4
                                v26 = v1
                                switch fi27 {
                                case 1:
                                    v28 = (!((v28 || true) || v28))
                                case -9223372036854775808:
                                    if v28 {
                                        v38 := GenS5{F0: (fi27 ^ 141), F1: 32767}
                                        _ = v38
                                    } else {
                                        v28 = v28
                                        v39 := []int{(-(fi27))}
                                        _ = v39
                                    }
                                case 0:
                                    if false {
                                        var v40 int = (-(fi27))
                                        _ = v40
                                    } else {
                                        var v41 int8 = -128
                                        _ = v41
                                    }
                                }
                            }
                        }
                    }
                default:
                    var v42 int8 = 100
                    _ = v42
                }
                var v43 uint32 = (v2 % (((v2 | (v2 & v2))) | 1))
                _ = v28
                _ = v43
            }
            if (!true) {
                v26 = (v26 & v26)
                v26 = (^(v1))
            }
            _ = fi27
        }
        v44 := []int{-55, -612, -985, 0}
        _ = v44
    }
    for fi45 := 0; fi45 < 18; fi45++ {
        v2 = v2
        var v46 bool = true
        if v46 {
            for fi47 := 0; fi47 < 2; fi47++ {
                var v48 float32 = ((v24 - 1.5) * v24)
                _ = fi47
                _ = v48
            }
            switch "hello" {
            case "foo bar":
                var v49 bool = false
                _ = v49
            case "\n\t":
                v24 = ((0.1 * float32(3.14159)) - 0.0)
            default:
                if ((0.1 != ((1.5 - float64(-2.25)) + (0.1 + -2.25))) || v46) {
                    v1 += (v26 >> ((uint(v2)) & 63))
                    var v50 float64 = 0.0
                    _ = v50
                }
            }
        } else {
            fmt.Println(v26, v26)
            for fi51 := 0; fi51 < 15; fi51++ {
                v2 -= v2
                v3[4294967295] = v46
                _ = fi51
            }
        }
        _ = fi45
        _ = v46
    }
    delete(v3, 466)
    v52 := []int{4294967295, 1, -401, -388}
    var v53 bool = ((false && false) || true)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v52
    _ = v53
}
