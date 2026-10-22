package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int32 }
type GenS2 struct { F0 int32; F1 int16; F2 uint16 }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = (-9223372036854775808 < 42)
    v3 := make(map[int]int)
    var v4 int64 = -1
    v5 := []string{"", ("" + ("" + "foo bar")), ("hello" + "hello")}
    var v6 int64 = v4
    var v7 bool = ((v6 - (v6 % ((v6) | 1))) != 0)
    if v7 {
        v8 := GenS1{F0: -9223372036854775808, F1: 8, F2: 1}
        v6 = v4
        var v9 float32 = (1e10 * 1e10)
        _ = v8
        _ = v9
    }
    fmt.Println(v2, v7, v2)
    var v10 float64 = 3.14159
    v11 := GenS1{F0: -1, F1: 0, F2: -2147483648}
    v4 = v6
    var v12 uint16 = 255
    switch "hello" {
    case "":
        v5 = append(v5, ("" + ""))
    case "\n\t":
        delete(v3, 128)
    case "hello":
        var v13 uint64 = 117
        _ = v13
    default:
        switch -485 {
        case -667:
            fmt.Println(v6, v2, v12)
        case 4294967295:
            for fi14 := 0; fi14 < 6; fi14++ {
                v15 := GenS2{F0: 2147483647, F1: 0, F2: (v12 | (v12 ^ (^(v12))))}
                for fi16 := 0; fi16 < 12; fi16++ {
                    v2 = (v12 >= 529)
                    if v7 {
                        _ = len(v5)
                    }
                    _ = fi16
                }
                _ = fi14
                _ = v15
            }
        default:
            v12 = v12
        }
    }
    v1[65535] = "foo bar"
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if true {
            v2 = v7
        } else {
            fmt.Println(v7, v2, v10)
            if v7 {
                if (92 == 9223372036854775807) {
                    for fi17 := 0; fi17 < 6; fi17++ {
                        if (!v7) {
                            if (v7 && (1 < 200)) {
                                v18 := GenS3{F0: 127}
                                v2 = (251 > 42)
                                var v19 uint16 = v12
                                _ = v18
                                _ = v19
                            } else {
                                var v20 int64 = 42
                                _ = v20
                            }
                            v7 = ((fi17 - (fi17 * (^(fi17)))) != fi17)
                            if v2 {
                                var v21 string = ""
                                _ = v21
                            } else {
                                var v22 uint8 = 189
                                var v23 int64 = (-(v4))
                                _ = v22
                                _ = v23
                            }
                        } else {
                            var v24 int = -9223372036854775808
                            fmt.Println(v6, v7, fi17)
                            for fi25 := 0; fi25 < 19; fi25++ {
                                if (158 == 48) {
                                    if false {
                                        v12 = 1
                                        switch "foo bar" {
                                        case "hello":
                                            v26 := GenS2{F0: -2147483648, F1: 32767, F2: (^(v12))}
                                            _ = v26
                                        case "foo bar":
                                            fmt.Println(fi17, v7)
                                        default:
                                            switch ("foo bar" + "") {
                                            case "\n\t":
                                                var v27 int = 916
                                                _ = v27
                                            case "":
                                                v28 := v1[912]
                                                _ = v28
                                            case "foo bar":
                                                var v29 int64 = (^(v6))
                                                _ = v29
                                            default:
                                                v4 -= (v6 / ((v6) | 1))
                                            }
                                        }
                                        var v30 int8 = -127
                                        _ = v30
                                    }
                                    v4 -= v6
                                    var v31 string = "\n\t"
                                    var v32 int8 = -1
                                    var v33 int8 = (v32 + -128)
                                    delete(v1, 501)
                                    var v34 int32 = -1
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                                _ = fi25
                            }
                            _ = v24
                        }
                        v10 = v10
                        v4 = (v6 - 0)
                        _ = fi17
                    }
                }
            } else {
                v5 = append(v5, ("" + (("hello" + "foo bar") + "hello")))
            }
            var v35 uint8 = 68
            _ = v35
        }
        v12 += v12
    } else {
        v12 = (v12 >> ((4294967295) & 63))
        v2 = v2
    }
    if (true || v2) {
        if (18446744073709551615 != 218) {
            v4 -= (v4 >> ((65535) & 63))
            for fi36 := 0; fi36 < 13; fi36++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi36
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
