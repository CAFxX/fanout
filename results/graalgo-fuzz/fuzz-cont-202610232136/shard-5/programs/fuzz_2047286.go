package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 uint; F2 uint32 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint64) int32 {
    if (-2147483648 >= -9223372036854775808) {
        p0 = p0
        fmt.Println(p0)
        if (191 != 65535) {
            for fi6 := 0; fi6 < 15; fi6++ {
                for fi7 := 0; fi7 < 19; fi7++ {
                    if (false && (false || false)) {
                        fmt.Println(fi6)
                        var v8 uint16 = 65535
                        v8 = (^(v8))
                        _ = v8
                    }
                    var v9 uint16 = 405
                    v9 -= v9
                    _ = fi7
                    _ = v9
                }
                _ = fi6
            }
            v10 := GenS1{F0: 42, F1: 83}
            _ = v10
        }
    } else {
        p0 = (^(p0))
        if false {
            fmt.Println(p0)
            var v11 uint32 = 349
            _ = v11
        } else {
            var v12 uint = 4294967295
            _ = v12
        }
    }
    switch -1 {
    case -9223372036854775808:
        var v13 float32 = ((3.14159 / ((0.1 / 1.0) + 0.0)) * 1.5)
        _ = v13
    case 1:
        p0 = 1
    default:
        fmt.Println(p0)
    }
    _ = p0
    return int32(p0)
}

func main() {
    v1 := []uint32{796, 976, 835}
    var v2 uint64 = 1
    var v3 int16 = 1000
    var v4 uint32 = 269
    v5 := make(map[uint32]uint32)
    switch ("a" + ("" + "foo bar")) {
    case "\n\t":
        switch "hello" {
        case "":
            delete(v5, v4)
        case "foo bar":
            if true {
                v4 = v4
                fmt.Println(v2, v2)
                var v14 string = (("hello" + "a") + (("hello" + "\n\t") + (((("a" + "") + "\n\t") + "") + "hello")))
                _ = v14
            } else {
                switch "hello" {
                case "a":
                    for fi15 := 0; fi15 < 19; fi15++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v16 := GenS1{F0: 0, F1: 42}
                        fmt.Println(v4, fi15)
                        _ = fi15
                        _ = v16
                    }
                case "hello":
                    for fi17 := 0; fi17 < 20; fi17++ {
                        fmt.Println(v2)
                        _ = fi17
                    }
                case "":
                    if false {
                        v18 := GenS1{F0: 127, F1: 1}
                        for fi19 := 0; fi19 < 1; fi19++ {
                            if ("foo bar" != "a") {
                                delete(v5, v4)
                                var v20 uint32 = v4
                                v4 = v4
                                _ = v20
                            } else {
                                v21 := GenS2{F0: (v4 - (v4 - 65535)), F1: 1, F2: (fi19 - fi19)}
                                _ = v21
                            }
                            v4 = v4
                            _ = fi19
                        }
                        for fi22 := 0; fi22 < 12; fi22++ {
                            v23 := GenS3{F0: true}
                            if (true || (9223372036854775807 == -211)) {
                                var v24 string = "\n\t"
                                v4 = (v4 + v4)
                                _ = v24
                            }
                            _ = fi22
                            _ = v23
                        }
                        _ = v18
                    }
                default:
                    v25 := GenS1{F0: 129, F1: -128}
                    _ = v25
                }
                if false {
                    switch "\n\t" {
                    case "":
                        v4 = (v4 / (((v4 % (((v4 + v4)) | 1))) | 1))
                    case "hello":
                        var v26 uint64 = v2
                        _ = v26
                    }
                    var v27 float32 = -2.25
                    if (!false) {
                        switch "" {
                        case "hello":
                            var v28 bool = false
                            _ = v28
                        default:
                            var v29 int = 32767
                            _ = v29
                        }
                        v27 -= ((1.5 + v27) + float32(1.5))
                    } else {
                        v3 = v3
                    }
                    _ = v27
                }
            }
        default:
            var v30 uint = 0
            _ = v30
        }
    default:
        v3 = 434
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        v2 = (v2 % ((642) | 1))
        v3 += (-(v3))
        v32 := GenS4{F0: 42, F1: 614, F2: (v4 >> ((255) & 63))}
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        var v34 int8 = -1
        var v35 string = "hello"
        _ = fi33
        _ = v34
        _ = v35
    }
    v1 = append(v1, 2)
    v36 := GenS5{F0: (((true && true) || (true || (0.0 == 1e10))) && (!(!(42 < 128))))}
    fmt.Println(v4, v3)
    v37 := GenS3{F0: ((v3 == (^(v3))) && false)}
    for fi38 := 0; fi38 < 7; fi38++ {
        v1 = append(v1, v4)
        _ = fi38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
