package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS5 struct { F0 uint32; F1 uint }
func genFn0(p0 uint8, p1 int32, p2 uint64) uint16 {
    var v5 int16 = 160
    v5 += -32768
    fmt.Println(p2, p1, p1)
    p2 = (p2 ^ p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 32767
}
func genFn1() int16 {
    var v6 float32 = 3.14159
    if true {
        v6 = v6
    } else {
        fmt.Println(v6)
        v6 = 1e10
        var v7 string = "hello"
        _ = v7
    }
    _ = v6
    return -32768
}

func main() {
    v1 := []string{"hello", "a"}
    v2 := make(map[string]string)
    var v3 string = "foo bar"
    v4 := make(map[uint32]bool)
    if (-1 == 62) {
        v4[2147483648] = (-2.25 >= (1e10 * -2.25))
    }
    v8 := GenS1{F0: 375, F1: -1, F2: 112}
    for fi9 := 0; fi9 < 2; fi9++ {
        v3 = "hello"
        _ = fi9
    }
    v10 := GenS1{F0: genFn0(128, 0, 42), F1: 9223372036854775807, F2: 0}
    if true {
        switch 1 {
        case -838:
            v3 = "a"
        case 1:
            switch 2147483647 {
            case 396:
                switch -9223372036854775808 {
                case -1:
                    fmt.Println(v3)
                case 127:
                    v1 = append(v1, (v3 + "\n\t"))
                case 9223372036854775807:
                    fmt.Println(v3)
                }
            case 2147483647:
                var v11 int16 = 42
                _ = v11
            }
        case 32767:
            if false {
                if false {
                    var v12 uint16 = 1
                    var v13 float64 = -2.25
                    v14 := GenS2{F0: 98, F1: 931, F2: -9223372036854775808}
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    v2[(v3 + v3)] = "\n\t"
                    v3 = "\n\t"
                    var v15 uint32 = 0
                    _ = v15
                }
                if ((0 == 515) || true) {
                    if false {
                        if (!true) {
                            fmt.Println(v3)
                            var v16 uint64 = 923
                            var v17 string = v3
                            _ = v16
                            _ = v17
                        }
                        v3 = ("" + v3)
                        var v18 int = 427
                        _ = v18
                    } else {
                        var v19 uint = 42
                        var v20 int64 = -9223372036854775808
                        _ = v19
                        _ = v20
                    }
                    switch v3 {
                    case "":
                        switch v3 {
                        case "foo bar":
                            v21 := GenS3{F0: 0, F1: 2}
                            _ = v21
                        }
                    case "a":
                        switch -2147483648 {
                        case 0:
                            var v22 int16 = -32768
                            _ = v22
                        case 128:
                            v1 = append(v1, v3)
                        case 127:
                            switch 9223372036854775807 {
                            case 1:
                                if (false || (0 <= 4294967295)) {
                                    v3 = "foo bar"
                                    for fi23 := 0; fi23 < 7; fi23++ {
                                        _ = len(v1)
                                        switch v3 {
                                        case "hello":
                                            v3 = v3
                                        case "foo bar":
                                            switch ("a" + (v3 + "foo bar")) {
                                            case "":
                                                if false {
                                                    v24 := GenS1{F0: genFn0(209, 0, 18446744073709551615), F1: 455, F2: 18}
                                                    v3 = (v3 + "hello")
                                                    _ = v24
                                                } else {
                                                    var v25 int = -307
                                                    for fi26 := 0; fi26 < 8; fi26++ {
                                                        v3 = v3
                                                        if true {
                                                            v27 := GenS4{F0: 438, F1: 2, F2: 840}
                                                            _ = v27
                                                        } else {
                                                            v28 := GenS3{F0: 191, F1: genFn0(0, 474, 0)}
                                                            v3 = "\n\t"
                                                            v25 = fi23
                                                            _ = v28
                                                        }
                                                        var v29 uint = 4294967295
                                                        _ = fi26
                                                        _ = v29
                                                    }
                                                    _ = v25
                                                }
                                            case "\n\t":
                                                var v30 uint16 = 715
                                                _ = v30
                                            case "hello":
                                                _ = len(v4)
                                            }
                                        case "":
                                            var v31 uint32 = 2147483648
                                            _ = v31
                                        }
                                        fmt.Println(fi23, v3)
                                        _ = fi23
                                    }
                                    v3 = "\n\t"
                                } else {
                                    v3 = v3
                                    var v32 int16 = genFn1()
                                    _ = v32
                                }
                            default:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        default:
                            if len(v1) > 0 {
                                v1[0] = v3
                            }
                        }
                    case "hello":
                        for fi33 := 0; fi33 < 5; fi33++ {
                            _ = len(v1)
                            var v34 int = -1
                            if (false && true) {
                                v34 += v34
                                var v35 int8 = int8(fi33)
                                var v36 bool = true
                                _ = v35
                                _ = v36
                            }
                            _ = fi33
                            _ = v34
                        }
                    }
                    v3 = ("hello" + v3)
                } else {
                    if (!(genFn1() > -32768)) {
                        fmt.Println(v3)
                    }
                }
                v3 = v3
            }
        }
    }
    if (1 != 965) {
        var v37 bool = true
        if (false || v37) {
            var v38 string = "\n\t"
            _ = v38
        }
        _ = v37
    } else {
        var v39 float64 = -2.25
        if (!true) {
            if true {
                switch 128 {
                case -995:
                    delete(v2, (v3 + "hello"))
                default:
                    if (!(0 <= -9223372036854775808)) {
                        var v40 float64 = ((float64(0.0) - v39) / v39)
                        if true {
                            v3 = (v3 + (v3 + ("foo bar" + v3)))
                        }
                        _ = v40
                    }
                }
                fmt.Println(v3, v3)
                var v41 uint8 = 2
                _ = v41
            }
            if (2 >= 250) {
                fmt.Println(v3, v39)
            }
        }
        _ = v39
    }
    v42 := GenS2{F0: 127, F1: 4294967295, F2: 785}
    switch "foo bar" {
    case "foo bar":
        v3 = "hello"
    default:
        v43 := GenS5{F0: 128, F1: 2}
        _ = v43
    }
    var v44 uint = 2
    v3 = ((v3 + v3) + v3)
    v44 -= v44
    fmt.Println(v3)
    if true {
        v44 = (^(v44))
    }
    var v45 uint = 42
    v44 = v45
    fmt.Println(v44, v45, v45)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v44
    _ = v45
    _ = v8
}
