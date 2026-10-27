package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint32 }
func genFn0(p0 int32, p1 bool, p2 string) int32 {
    var v4 string = p2
    for fi5 := 0; fi5 < 16; fi5++ {
        p2 = ""
        _ = fi5
    }
    v6 := GenS1{F0: (p0 ^ p0), F1: 286}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 2
    v3 := []uint32{552}
    switch "\n\t" {
    case "foo bar":
        v1[4294967295] = (("" + "foo bar") + "")
    default:
        v7 := v1[9223372036854775807]
        _ = v7
    }
    v8 := GenS2{F0: 264}
    v2 = 4294967295
    if (3.14159 == 3.14159) {
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := GenS1{F0: -2147483648, F1: 1}
            v2 += (^(v2))
            v11 := GenS1{F0: -946, F1: 1}
            _ = fi9
            _ = v10
            _ = v11
        }
        v2 -= (v2 + v2)
        v2 -= (v2 >> ((255) & 63))
    } else {
        if len(v3) > 0 {
            v3[0] = 1
        }
        var v12 int32 = int32(v2)
        _ = v12
    }
    var v13 float64 = ((((float64(0.1) / 3.14159) + ((3.14159 / 1.5) - (((1.5 * (1e10 * ((1e10 - 1.5) * 1.5))) + -2.25) - (-2.25 + -2.25)))) / (0.0 + 1.5)) - 1.5)
    v14 := GenS2{F0: -2147483648}
    v15 := GenS3{F0: 403, F1: (("" + ("" + "\n\t")) + ("\n\t" + "hello"))}
    switch 636 {
    case 32768:
        for fi16 := 0; fi16 < 3; fi16++ {
            for fi17 := 0; fi17 < 9; fi17++ {
                switch ("\n\t" + "hello") {
                case "":
                    v13 = v13
                }
                var v18 uint8 = 255
                var v19 int8 = 1
                _ = fi17
                _ = v18
                _ = v19
            }
            v20 := GenS2{F0: 2147483647}
            switch (fi16 | (fi16 << ((700) & 63))) {
            case 4294967295:
                v13 += 0.0
            }
            _ = fi16
            _ = v20
        }
    case -703:
        var v21 float64 = v13
        _ = v21
    case 32767:
        v13 -= v13
    default:
        v22 := GenS1{F0: -968, F1: 793}
        _ = v22
    }
    v2 -= v2
    for fi23 := 0; fi23 < 10; fi23++ {
        v13 -= float64(float32(v13))
        for fi24 := 0; fi24 < 10; fi24++ {
            fmt.Println(fi24, v2, v13)
            fmt.Println(fi23, fi24)
            v25 := GenS3{F0: 1, F1: "a"}
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    var v26 float64 = 1e10
    for fi27 := 0; fi27 < 2; fi27++ {
        if false {
            switch "" {
            case "":
                if (255 <= 1) {
                    for fi28 := 0; fi28 < 15; fi28++ {
                        if (0 >= -1) {
                            v26 -= v26
                            v29 := GenS4{F0: true, F1: (fi27 - -581)}
                            _ = v29
                        } else {
                            fmt.Println(v26, v13, v13)
                            var v30 int8 = 42
                            fmt.Println(fi27, v2)
                            v30 -= 1
                            _ = v30
                        }
                        fmt.Println(fi27, v2)
                        _ = fi28
                    }
                    var v31 uint = 1
                    _ = v31
                } else {
                    v3 = append(v3, 2)
                    v2 = (v2 - (v2 & (v2 & 382)))
                    v32 := GenS5{F0: false}
                    _ = v32
                }
            case "foo bar":
                v2 = (v2 >> ((2) & 63))
            }
        }
        switch 250 {
        case 2147483647:
            var v33 uint64 = v2
            _ = v33
        case 810:
            switch -2147483648 {
            case 844:
                var v34 int32 = 2147483647
                _ = v34
            case 128:
                var v35 int64 = 928
                _ = v35
            }
        case 128:
            if true {
                var v36 int32 = 1
                _ = v36
            } else {
                v37 := GenS3{F0: 42, F1: "a"}
                _ = v37
            }
        default:
            var v38 uint = 18446744073709551615
            _ = v38
        }
        if false {
            if (((!(true && (((((0.0 + 1.5) + 0.1) != ((v26 * v26) - v26)) && false) && false))) || (int8(v2) != -1)) && true) {
                v39 := GenS4{F0: false, F1: 82}
                switch "hello" {
                case "foo bar":
                    _ = len(v1)
                default:
                    for fi40 := 0; fi40 < 11; fi40++ {
                        if (100 <= 0) {
                            if true {
                                v41 := v1[339]
                                var v42 float32 = (1e10 - 1e10)
                                _ = v41
                                _ = v42
                            } else {
                                if ((false && true) || true) {
                                    if (!(!false)) {
                                        var v43 float64 = -2.25
                                        if (3.14159 > v43) {
                                            fmt.Println(v13, v13, v13)
                                        }
                                        _ = len(v1)
                                        _ = v43
                                    } else {
                                        if (!true) {
                                            var v44 bool = true
                                            v45 := GenS6{F0: 65535, F1: -1}
                                            _ = v44
                                            _ = v45
                                        } else {
                                            v46 := GenS7{F0: 801}
                                            v26 = v26
                                            _ = v46
                                        }
                                        v47 := GenS8{F0: 2147483647}
                                        var v48 int = (fi27 >> ((42) & 63))
                                        _ = v47
                                        _ = v48
                                    }
                                } else {
                                    fmt.Println(v13, v13)
                                }
                            }
                            var v49 int = fi40
                            _ = v49
                        } else {
                            if true {
                                v26 = (v13 - (v26 * 0.0))
                                fmt.Println(v26, fi27)
                                v26 = v13
                                v26 = float64((0.1 - (-2.25 * 0.0)))
                                v2 += 0
                            }
                            v2 -= v2
                            var v50 float32 = (1.5 / 0.1)
                            _ = v50
                        }
                        _ = fi40
                    }
                }
                _ = v39
            } else {
                var v51 int64 = -1
                fmt.Println(v13)
                _ = v51
            }
            v52 := GenS8{F0: 4294967295}
            v13 -= ((v26 - 0.0) / (3.14159 / v13))
            _ = v52
        }
        _ = fi27
    }
    if (42 != 918) {
        var v53 int8 = 127
        _ = v53
    }
    fmt.Println(v13, v2)
    fmt.Println(v13)
    var v54 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v54
    _ = v8
}
