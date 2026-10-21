package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 uint16; F2 int }
type GenS5 struct { F0 int64; F1 uint32; F2 int64 }
type GenS6 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{}
    v3 := []int{-457, -1, -1}
    var v4 uint = 1
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch 32767 {
        case 187:
            _ = len(v1)
        }
        var v5 int64 = -64
        _ = v5
    } else {
        var v6 uint16 = 2
        var v7 string = "foo bar"
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        v4 -= v4
        _ = fi8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 17; fi10++ {
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            _ = fi10
        }
        switch "foo bar" {
        case "a":
            var v11 int16 = 831
            _ = v11
        case "foo bar":
            v12 := GenS1{F0: 1, F1: 0}
            _ = v12
        }
        var v13 float32 = ((float32(((0.0 * (3.14159 - (3.14159 * -2.25))) - float64((3.14159 - 3.14159)))) + float32(1.5)) * 1e10)
        _ = fi9
        _ = v13
    }
    var v14 int64 = 9223372036854775807
    if (true && false) {
        var v15 uint64 = 323
        var v16 uint16 = 65535
        fmt.Println(v15, v16)
        _ = v15
        _ = v16
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v14, v14)
    }
    v14 += v14
    var v17 uint16 = 65535
    fmt.Println(v17, v4)
    if (!(1e10 <= float64(((0.0 - (3.14159 - (3.14159 + -2.25))) / 1.0)))) {
        v14 = int64(v17)
        if (0 != v14) {
            fmt.Println(v17)
        }
        var v18 string = ""
        _ = v18
    } else {
        var v19 int16 = 0
        var v20 uint32 = 1
        _ = v19
        _ = v20
    }
    if (false && (!(42 >= v4))) {
        var v21 uint8 = 255
        for fi22 := 0; fi22 < 19; fi22++ {
            if false {
                fmt.Println(v17, v4, v21)
                var v23 uint8 = (v21 % ((255) | 1))
                if false {
                    var v24 int16 = 42
                    _ = v24
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        for fi25 := 0; fi25 < 16; fi25++ {
                            v26 := GenS1{F0: 1, F1: 42}
                            fmt.Println(v17, v14)
                            if len(v3) > 0 {
                                v3[0] = 42
                            }
                            _ = fi25
                            _ = v26
                        }
                    case "a":
                        var v27 int = (fi22 + (fi22 + fi22))
                        _ = v27
                    }
                    var v28 uint64 = 42
                    v29 := GenS1{F0: 1, F1: -721}
                    _ = v28
                    _ = v29
                }
                _ = v23
            }
            v1[42] = 42
            _ = fi22
        }
        _ = v21
    } else {
        if false {
            v30 := GenS1{F0: 42, F1: 32767}
            v31 := GenS2{F0: 423, F1: 0, F2: v17}
            for fi32 := 0; fi32 < 9; fi32++ {
                var v33 int16 = -551
                if true {
                    var v34 int16 = v33
                    if (!false) {
                        v14 = v14
                    }
                    _ = v34
                } else {
                    switch (fi32 ^ (fi32 / ((4294967295) | 1))) {
                    case -2147483648:
                        if false {
                            v17 = (v17 ^ 32767)
                        } else {
                            for fi35 := 0; fi35 < 8; fi35++ {
                                v14 = v14
                                var v36 int = -9223372036854775808
                                _ = fi35
                                _ = v36
                            }
                        }
                    case 77:
                        switch "\n\t" {
                        case "":
                            if (v4 <= 4294967295) {
                                switch fi32 {
                                case 4294967295:
                                    var v37 uint64 = 2
                                    _ = v37
                                case 651:
                                    v38 := GenS3{F0: (-(fi32))}
                                    _ = v38
                                }
                            } else {
                                var v39 float32 = 0.0
                                switch fi32 {
                                case 32767:
                                    if ((3.14159 < ((((-2.25 / float64(v39)) / 1e10) * 0.0) * float64(((0.1 * 0.1) + v39)))) || (((v4 << (((v4 >> ((v4) & 63))) & 63)) <= v4) && false)) {
                                        var v40 bool = (!(!true))
                                        if (56 < -113) {
                                            v41 := GenS4{F0: v4, F1: 530, F2: int(v14)}
                                            _ = v41
                                        } else {
                                            v3 = append(v3, fi32)
                                            v40 = ((v40 && true) || v40)
                                            if (v14 <= 9223372036854775807) {
                                                for fi42 := 0; fi42 < 4; fi42++ {
                                                    var v43 uint = v4
                                                    _ = fi42
                                                    _ = v43
                                                }
                                                var v44 uint16 = v17
                                                _ = v44
                                            } else {
                                                var v45 uint = 157
                                                _ = len(v1)
                                                switch (fi32 - fi32) {
                                                case 2147483647:
                                                    if (v40 || (444 <= 0)) {
                                                        v46 := GenS5{F0: (^(v14)), F1: 42, F2: (v14 / ((v14) | 1))}
                                                        if false {
                                                            _ = len(v1)
                                                            if v40 {
                                                                fmt.Println(v14)
                                                            }
                                                            for fi47 := 0; fi47 < 12; fi47++ {
                                                                if (127 > 0) {
                                                                    _ = len(v1)
                                                                    _ = len(v2)
                                                                    var v48 uint32 = 695
                                                                    _ = v48
                                                                }
                                                                v49 := v1[231]
                                                                _ = fi47
                                                                _ = v49
                                                            }
                                                        }
                                                        _ = v46
                                                    }
                                                case 32768:
                                                    var v50 int16 = v33
                                                    _ = v50
                                                default:
                                                    var v51 string = ("hello" + "foo bar")
                                                    _ = v51
                                                }
                                                _ = v45
                                            }
                                        }
                                        _ = v40
                                    }
                                default:
                                    for fi52 := 0; fi52 < 8; fi52++ {
                                        v14 = v14
                                        v33 = 0
                                        v39 = ((v39 / (v39 - v39)) * (v39 / v39))
                                        _ = fi52
                                    }
                                }
                                _ = v39
                            }
                        case "a":
                            v14 += (v14 * v14)
                        }
                    }
                    v53 := GenS6{F0: -1}
                    _ = v53
                }
                _ = fi32
                _ = v33
            }
            _ = v30
            _ = v31
        }
        var v54 int8 = 127
        _ = v54
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
