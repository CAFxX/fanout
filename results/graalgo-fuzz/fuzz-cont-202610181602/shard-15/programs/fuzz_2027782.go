package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint32; F2 string }
type GenS3 struct { F0 int32; F1 uint16; F2 int64 }
func genFn0(p0 int16, p1 int64) int64 {
    v7 := []int{}
    for fi8 := 0; fi8 < 6; fi8++ {
        p1 += p1
        _ = fi8
    }
    var v9 bool = false
    var v10 int64 = (p1 | 1)
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v9
    return 9223372036854775807
}
func genFn1() int16 {
    v11 := []int{}
    if true {
        var v12 uint16 = 1
        if false {
            v13 := GenS1{F0: 65535, F1: 18446744073709551615}
            _ = v13
        }
        _ = v12
    }
    if false {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        switch "" {
        case "":
            var v14 uint32 = 4294967295
            _ = v14
        case "foo bar":
            var v15 int8 = 93
            _ = v15
        }
        var v16 int64 = -597
        _ = v16
    } else {
        var v17 string = "foo bar"
        _ = v17
    }
    switch 654 {
    case 271:
        var v18 string = "\n\t"
        _ = v18
    case 456:
        var v19 float64 = 0.0
        _ = v19
    case 670:
        var v20 uint64 = 588
        _ = v20
    }
    var v21 string = "hello"
    _ = v11
    _ = v21
    return 1
}

func main() {
    v1 := []string{"\n\t", "hello", "foo bar", (("\n\t" + ("hello" + ("hello" + ""))) + "")}
    var v2 int16 = 0
    var v3 float32 = (-2.25 + -2.25)
    v4 := make(map[int]int)
    v5 := make(map[int]int)
    var v6 float32 = v3
    if false {
        if (!("foo bar" > "a")) {
            delete(v4, 1)
            if len(v1) > 0 {
                v1[0] = "hello"
            }
        } else {
            v3 += 3.14159
        }
    } else {
        if false {
            switch 9223372036854775807 {
            case 128:
                if ((!("hello" < "hello")) && (-90 >= 0)) {
                    var v22 int16 = v2
                    fmt.Println(v22, v3, v2)
                    v22 += v22
                    _ = v22
                }
            case -599:
                delete(v4, 78)
            case -1:
                if (4294967295 == 2) {
                    v23 := GenS1{F0: 10, F1: 918}
                    _ = v23
                } else {
                    v6 += v3
                    switch "" {
                    case "":
                        switch "a" {
                        case "hello":
                            if ((!(!true)) || (v3 == v6)) {
                                v24 := GenS1{F0: 4294967295, F1: 4294967295}
                                _ = v24
                            } else {
                                for fi25 := 0; fi25 < 3; fi25++ {
                                    v26 := GenS1{F0: 255, F1: 2}
                                    _ = fi25
                                    _ = v26
                                }
                                fmt.Println(v2)
                            }
                        case "":
                            var v27 float32 = ((float32(1e10) / v3) / 0.1)
                            _ = v27
                        case "a":
                            switch "\n\t" {
                            case "a":
                                if len(v1) > 0 {
                                    v1[0] = "a"
                                }
                            default:
                                for fi28 := 0; fi28 < 2; fi28++ {
                                    v29 := GenS2{F0: true, F1: 4294967295, F2: "\n\t"}
                                    _ = fi28
                                    _ = v29
                                }
                            }
                        }
                    case "hello":
                        for fi30 := 0; fi30 < 6; fi30++ {
                            var v31 uint16 = 65535
                            _ = fi30
                            _ = v31
                        }
                    case "foo bar":
                        _ = len(v4)
                    }
                    if false {
                        v6 += (-2.25 + v6)
                        fmt.Println(v3, v2)
                    } else {
                        switch "a" {
                        case "hello":
                            delete(v4, 32767)
                        case "a":
                            if true {
                                v6 += 1.5
                            } else {
                                _ = len(v4)
                                var v32 uint8 = 136
                                _ = v32
                            }
                        }
                    }
                }
            }
            v33 := GenS3{F0: 543, F1: 2, F2: genFn0((v2 | (v2 % ((v2) | 1))), 1)}
            _ = v33
        }
        for fi34 := 0; fi34 < 1; fi34++ {
            if (300 <= 2147483647) {
                v6 = (float32(0.1) * float32((0.0 - (float64(0.0) + 1e10))))
                fmt.Println(v2, fi34)
            }
            for fi35 := 0; fi35 < 10; fi35++ {
                var v36 uint = 42
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
    }
    if (!(128 != 220)) {
        var v37 uint = 18446744073709551615
        var v38 float64 = ((0.0 + (float64(float32((1.5 - 0.0))) / 1e10)) - 1e10)
        _ = v37
        _ = v38
    } else {
        v6 = float32(3.14159)
        switch -647 {
        case -2147483648:
            v6 += v6
        case 32767:
            var v39 uint64 = 1
            _ = v39
        case -1:
            v2 += (v2 / ((v2) | 1))
        default:
            switch ("" + "a") {
            case "\n\t":
                var v40 int = 9223372036854775807
                _ = v40
            case "":
                v3 += -2.25
            case "foo bar":
                fmt.Println(v2)
            default:
                var v41 int32 = -1
                _ = v41
            }
        }
    }
    var v42 bool = (false && true)
    if (!v42) {
        v3 = (float32((0.1 * 3.14159)) - 1e10)
    }
    var v43 int64 = -1
    if ((-1 < 9) && v42) {
        fmt.Println(v6, v6)
        var v44 uint = 255
        for fi45 := 0; fi45 < 16; fi45++ {
            fmt.Println(v3, v44)
            var v46 int = (fi45 % ((fi45) | 1))
            _ = fi45
            _ = v46
        }
        _ = v44
    }
    v43 = -1
    v6 = 1e10
    v6 += (float32(0.1) - (float32(0.0) + v6))
    var v47 int16 = (v2 * (v2 - 1))
    fmt.Println(v47, v43, v43)
    v48 := v4[42]
    if false {
        switch "" {
        case "foo bar":
            for fi49 := 0; fi49 < 9; fi49++ {
                v2 += (v2 & v47)
                if v42 {
                    var v50 int = v48
                    fmt.Println(v6, v6)
                    v1 = append(v1, "hello")
                    _ = v50
                } else {
                    _ = len(v4)
                    v51 := GenS2{F0: false, F1: 724, F2: "hello"}
                    _ = v51
                }
                _ = fi49
            }
        case "hello":
            var v52 int8 = 0
            _ = v52
        default:
            switch "" {
            case "":
                fmt.Println(v2, v43)
            case "a":
                for fi53 := 0; fi53 < 14; fi53++ {
                    _ = len(v4)
                    _ = fi53
                }
            }
        }
    }
    fmt.Println(v48)
    v54 := GenS2{F0: (v42 && ((1e10 / 1.0) == (((-2.25 - v6) + v6) + v6))), F1: 427, F2: ""}
    var v55 uint64 = 4294967295
    var v56 uint8 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v47
    _ = v48
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v6
}
