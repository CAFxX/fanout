package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64; F1 int32 }

func main() {
    var v1 int = 9223372036854775807
    var v2 float64 = (float64(1.5) * -2.25)
    var v3 int8 = 127
    var v4 uint64 = 18446744073709551615
    var v5 uint64 = (^(v4))
    v3 -= (-(v3))
    fmt.Println(v4)
    fmt.Println(v3, v2)
    var v6 uint = 854
    var v7 float64 = v2
    for fi8 := 0; fi8 < 3; fi8++ {
        if true {
            for fi9 := 0; fi9 < 9; fi9++ {
                switch "" {
                case "":
                    v10 := []int{1}
                    _ = v10
                case "foo bar":
                    var v11 uint = (v6 >> ((v6) & 63))
                    _ = v11
                case "a":
                    v3 = (^(v3))
                    v12 := GenS1{F0: false, F1: v5}
                    _ = v12
                }
                _ = fi9
            }
        } else {
            if true {
                for fi13 := 0; fi13 < 7; fi13++ {
                    var v14 int64 = 9223372036854775807
                    _ = fi13
                    _ = v14
                }
            }
            v7 -= v2
        }
        for fi15 := 0; fi15 < 14; fi15++ {
            fmt.Println(fi15, fi8, fi15)
            fmt.Println(fi15)
            v16 := GenS2{F0: 0, F1: "\n\t"}
            _ = fi15
            _ = v16
        }
        _ = fi8
    }
    var v17 int32 = 42
    v18 := GenS1{F0: (v6 == 624), F1: (v5 << (((v6 / ((v6) | 1))) & 63))}
    for fi19 := 0; fi19 < 13; fi19++ {
        if (!true) {
            v20 := []int{int(v6), (v1 % ((fi19) | 1)), fi19}
            fmt.Println(v4, v5, v7)
            _ = v20
        }
        _ = fi19
    }
    var v21 int64 = -1
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := GenS2{F0: v21, F1: "hello"}
        _ = fi22
        _ = v23
    }
    var v24 int64 = int64(v17)
    v5 = 1
    var v25 uint16 = 32767
    v26 := []int{}
    if false {
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
        if ((32767 > 32767) && true) {
            var v27 float32 = 3.14159
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
            var v28 int8 = v3
            _ = v27
            _ = v28
        }
        if false {
            v26 = append(v26, (v1 - int(v3)))
            if false {
                var v29 uint32 = 825
                fmt.Println(v5, v29)
                for fi30 := 0; fi30 < 19; fi30++ {
                    switch int(v25) {
                    case -9223372036854775808:
                        var v31 int16 = 42
                        _ = v31
                    case 32767:
                        v25 = (v25 / (((v25 | v25)) | 1))
                    case 127:
                        v21 = v24
                    default:
                        if (true && ((v24 * v24) > (^(v24)))) {
                            if (0 <= uint8(v5)) {
                                if (-1 == -160) {
                                    fmt.Println(v3)
                                } else {
                                    v32 := GenS3{F0: 14}
                                    _ = v32
                                }
                                v4 = v4
                            }
                            v1 -= fi30
                            var v33 string = (("foo bar" + ("foo bar" + "hello")) + "a")
                            _ = v33
                        } else {
                            v24 -= v24
                            v26 = append(v26, fi30)
                            if ((v3 - (v3 + (v3 >> (((v6 / ((516) | 1))) & 63)))) < (v3 % ((v3) | 1))) {
                                v34 := GenS1{F0: (((((!false) || (false || (true || (!(157 > 0))))) && false) || ((((true || false) || true) && (false && false)) || (((fi30 & (-(fi30))) > (fi30 / (((v1 | fi30)) | 1))) && (42 >= uint8(v25))))) || true), F1: v5}
                                var v35 int = (fi30 / ((v1) | 1))
                                v36 := GenS4{F0: 0}
                                _ = v34
                                _ = v35
                                _ = v36
                            } else {
                                v37 := GenS2{F0: 9223372036854775807, F1: "foo bar"}
                                for fi38 := 0; fi38 < 17; fi38++ {
                                    v39 := GenS5{F0: 2147483648, F1: (v25 << ((v6) & 63)), F2: v21}
                                    _ = fi38
                                    _ = v39
                                }
                                v40 := GenS6{F0: v3}
                                _ = v37
                                _ = v40
                            }
                        }
                    }
                    v1 = 1
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                    _ = fi30
                }
                _ = v29
            }
            v41 := GenS5{F0: 2147483648, F1: v25, F2: v21}
            _ = v41
        } else {
            var v42 uint = 4294967295
            v43 := GenS7{F0: (v21 & v24), F1: v17}
            _ = v42
            _ = v43
        }
    } else {
        switch (("foo bar" + "foo bar") + "\n\t") {
        case "hello":
            v5 += (v5 * 4294967295)
        default:
            if (1e10 != 3.14159) {
                fmt.Println(v7, v5, v2)
                var v44 uint16 = v25
                for fi45 := 0; fi45 < 9; fi45++ {
                    switch (("a" + "\n\t") + "a") {
                    case "hello":
                        if ((0.1 / ((v2 - (1e10 * ((v7 / (float64(0.0) + v7)) * v2))) / float64(-2.25))) <= (1.5 / (v2 / v2))) {
                            v17 -= v17
                            for fi46 := 0; fi46 < 11; fi46++ {
                                v44 += v44
                                _ = fi46
                            }
                        } else {
                            var v47 uint64 = (v4 % ((1) | 1))
                            var v48 float32 = 1.5
                            _ = v47
                            _ = v48
                        }
                    }
                    _ = fi45
                }
                _ = v44
            } else {
                var v49 int8 = v3
                _ = v49
            }
        }
    }
    var v50 int = -348
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v7
}
