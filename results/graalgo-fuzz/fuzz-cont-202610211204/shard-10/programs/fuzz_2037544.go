package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint16; F1 string; F2 uint8 }
func genFn0(p0 bool) int8 {
    fmt.Println(p0)
    v5 := GenS1{F0: 38}
    if (!p0) {
        for fi6 := 0; fi6 < 4; fi6++ {
            p0 = (42 >= 42)
            v7 := []int{32767, (fi6 | (fi6 | -848))}
            _ = len(v7)
            _ = fi6
            _ = v7
        }
        fmt.Println(p0)
    } else {
        var v8 uint32 = 2
        p0 = p0
        fmt.Println(v8, p0)
        _ = v8
    }
    _ = p0
    _ = v5
    return 1
}

func main() {
    var v1 int8 = 0
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    var v4 int16 = 1
    if (true || false) {
        var v9 int = -9223372036854775808
        v10 := GenS1{F0: 128}
        _ = v10
        _ = v9
    }
    delete(v2, 2)
    switch 2147483647 {
    case -149:
        delete(v3, 2)
    default:
        var v11 float64 = (3.14159 - 0.1)
        _ = v11
    }
    switch "foo bar" {
    case "foo bar":
        if (2 < 155) {
            var v12 uint8 = 100
            _ = v12
        }
    }
    var v13 uint32 = 1
    v1 += v1
    v14 := []int{42, -334}
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 int8 = (v1 | (v1 & (v1 >> ((65535) & 63))))
        switch (fi15 + (fi15 / (((fi15 & fi15)) | 1))) {
        case 9223372036854775807:
            v1 += 1
        case 32767:
            v13 = v13
        case 42:
            v17 := GenS1{F0: 127}
            _ = v17
        default:
            switch ("foo bar" + "hello") {
            case "hello":
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            default:
                switch "hello" {
                case "a":
                    var v18 int16 = v4
                    _ = v18
                case "hello":
                    for fi19 := 0; fi19 < 13; fi19++ {
                        if ((!(-255 <= 1)) && false) {
                            if true {
                                var v20 float64 = (3.14159 - (1.5 * (3.14159 * -2.25)))
                                v1 = genFn0(("hello" == (((("hello" + "hello") + "hello") + "\n\t") + "hello")))
                                var v21 string = ("hello" + "")
                                v22 := GenS2{F0: 873, F1: 1}
                                v4 -= (v4 + v4)
                                v23 := v2[374]
                                _ = v20
                                _ = v21
                                _ = v22
                                _ = v23
                            } else {
                                if len(v14) > 0 {
                                    v14[0] = (^(fi15))
                                }
                            }
                            v16 -= (v16 + (v16 * v1))
                        } else {
                            if len(v14) > 0 {
                                v14[0] = fi19
                            }
                            var v24 uint16 = 1
                            var v25 int = 2147483647
                            _ = v24
                            _ = v25
                        }
                        if len(v14) > 0 {
                            v14[0] = (fi15 + fi19)
                        }
                        _ = fi19
                    }
                default:
                    switch -9223372036854775808 {
                    case 9223372036854775807:
                        var v26 int32 = 42
                        _ = v26
                    default:
                        v2[2147483647] = 0
                    }
                }
            }
        }
        if len(v14) > 0 {
            v14[0] = (fi15 / ((fi15) | 1))
        }
        _ = fi15
        _ = v16
    }
    v1 = v1
    if (("hello" + "a") < ("foo bar" + "\n\t")) {
        var v27 uint64 = 4294967295
        var v28 uint = 448
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        switch fi29 {
        case 32767:
            v13 = (v13 >> ((1) & 63))
        case -1:
            v30 := GenS2{F0: 42, F1: v4}
            _ = v30
        case -252:
            if true {
                if len(v14) > 0 {
                    v14[0] = (-(fi29))
                }
            }
        }
        if false {
            switch "a" {
            case "":
                if true {
                    if (true && true) {
                        switch (("a" + "\n\t") + "a") {
                        case "\n\t":
                            var v31 uint = 255
                            _ = v31
                        case "hello":
                            var v32 uint = 18446744073709551615
                            _ = v32
                        case "foo bar":
                            var v33 uint16 = 0
                            _ = v33
                        }
                    } else {
                        fmt.Println(v4, fi29, v1)
                    }
                } else {
                    v1 -= (v1 | -128)
                    var v34 uint = uint(v13)
                    if len(v14) > 0 {
                        v14[0] = (fi29 | 1)
                    }
                    _ = v34
                }
            case "hello":
                for fi35 := 0; fi35 < 14; fi35++ {
                    var v36 int8 = (v1 - (v1 * genFn0((4294967295 < 851))))
                    _ = fi35
                    _ = v36
                }
            case "foo bar":
                v1 = v1
            }
            fmt.Println(fi29, v4, v1)
            var v37 float64 = (float64(0.1) * 1.5)
            _ = v37
        } else {
            v13 = (v13 + (v13 + (^(v13))))
            v38 := GenS2{F0: 692, F1: 684}
            _ = v38
        }
        _ = fi29
    }
    if (0 <= 42) {
        fmt.Println(v13, v1, v13)
        for fi39 := 0; fi39 < 15; fi39++ {
            for fi40 := 0; fi40 < 13; fi40++ {
                v41 := GenS1{F0: 255}
                v1 = int8(v13)
                _ = fi40
                _ = v41
            }
            if true {
                v4 = v4
                var v42 int32 = 165
                v43 := GenS2{F0: uint(v13), F1: 42}
                _ = v42
                _ = v43
            } else {
                _ = len(v14)
                var v44 bool = (!(true && false))
                _ = v44
            }
            _ = fi39
        }
    } else {
        v1 = v1
    }
    if (!true) {
        v13 = (^(v13))
    } else {
        switch "hello" {
        case "foo bar":
            v4 = v4
        }
        for fi45 := 0; fi45 < 16; fi45++ {
            for fi46 := 0; fi46 < 5; fi46++ {
                if true {
                    switch (fi45 + -2147483648) {
                    case -2147483648:
                        var v47 int64 = -1
                        _ = v47
                    default:
                        v1 -= (v1 - v1)
                    }
                }
                switch fi46 {
                case -2147483648:
                    var v48 int32 = 2147483647
                    _ = v48
                case 861:
                    v4 -= v4
                    var v49 float64 = 1e10
                    _ = v49
                case 1:
                    v50 := GenS3{F0: 1, F1: ("foo bar" + ("\n\t" + "a")), F2: 127}
                    _ = v50
                default:
                    var v51 int64 = 1
                    _ = v51
                }
                var v52 int16 = -1
                _ = fi46
                _ = v52
            }
            if len(v14) > 0 {
                v14[0] = 127
            }
            if true {
                for fi53 := 0; fi53 < 13; fi53++ {
                    v54 := GenS1{F0: uint8(v13)}
                    var v55 uint32 = (v13 + (^(v13)))
                    v55 -= 2147483647
                    _ = fi53
                    _ = v54
                    _ = v55
                }
                switch int(v4) {
                case 62:
                    for fi56 := 0; fi56 < 9; fi56++ {
                        var v57 int16 = (v4 / ((0) | 1))
                        _ = len(v2)
                        fmt.Println(fi56)
                        _ = fi56
                        _ = v57
                    }
                }
            }
            _ = fi45
        }
        var v58 int32 = -207
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
