package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 bool }

func main() {
    var v1 int8 = -75
    var v2 uint16 = 499
    v3 := []string{(("a" + "a") + "a"), "a", "hello"}
    var v4 float64 = ((0.1 * 3.14159) + float64(1e10))
    var v5 bool = (((1 == 128) && false) || (true && true))
    v6 := []int{2147483647, -1, 9223372036854775807, -9223372036854775808}
    if (1 <= 127) {
        if ((0 > 18446744073709551615) && (-1 <= 9223372036854775807)) {
            var v7 float64 = v4
            if (!((!false) || v5)) {
                var v8 uint8 = 2
                _ = v8
            }
            if (!((-522 == int64(v1)) && ("" != ""))) {
                v1 = (v1 | (v1 << ((2) & 63)))
                v9 := GenS1{F0: 2147483647, F1: 119}
                var v10 float32 = 0.1
                _ = v10
                _ = v9
            }
            _ = v7
        }
        if true {
            if true {
                _ = len(v3)
                switch 2147483647 {
                case 42:
                    fmt.Println(v1, v1, v1)
                default:
                    fmt.Println(v5, v4, v4)
                }
                if (0 >= 1) {
                    var v11 int = 168
                    _ = v11
                }
            } else {
                if false {
                    v4 -= v4
                    fmt.Println(v4, v4)
                } else {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        switch "" {
                        case "\n\t":
                            if (590 == 65535) {
                                v5 = (v5 && v5)
                                for fi13 := 0; fi13 < 19; fi13++ {
                                    if v5 {
                                        if len(v6) > 0 {
                                            v6[0] = (^(fi12))
                                        }
                                        if ((((1e10 * 3.14159) / 1.0) * -2.25) == ((0.0 + 1e10) / (1e10 + (float32(1e10) + 3.14159)))) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v5 = v5
                                            if len(v6) > 0 {
                                                v6[0] = 32767
                                            }
                                            var v14 uint16 = 477
                                            var v15 float32 = (3.14159 - (-2.25 - float32(v4)))
                                            _ = v14
                                            _ = v15
                                        }
                                    } else {
                                        fmt.Println(v1)
                                    }
                                    v6 = append(v6, fi13)
                                    var v16 uint = 0
                                    var v17 uint8 = 108
                                    fmt.Println(v4, v2, v1)
                                    _ = fi13
                                    _ = v16
                                    _ = v17
                                }
                            }
                        case "a":
                            v18 := GenS1{F0: 1, F1: -1}
                            _ = v18
                        case "":
                            if ((^(v2)) >= 226) {
                                fmt.Println(v4)
                            }
                        default:
                            for fi19 := 0; fi19 < 6; fi19++ {
                                if false {
                                    var v20 uint = 0
                                    var v21 int = (fi19 & fi12)
                                    switch (fi19 | fi19) {
                                    case 4294967295:
                                        var v22 uint = (^(v20))
                                        _ = v22
                                    case -773:
                                        switch (fi19 * fi12) {
                                        case -9223372036854775808:
                                            if len(v3) > 0 {
                                                v3[0] = "\n\t"
                                            }
                                        case 1:
                                            v21 = v21
                                        default:
                                            v23 := GenS2{F0: v1}
                                            _ = v23
                                        }
                                    default:
                                        if (0 == 762) {
                                            v2 += v2
                                            var v24 int8 = v1
                                            if (fi19 == 127) {
                                                v24 = v1
                                                var v25 int32 = -586
                                                v20 = (v20 % (((v20 >> ((211) & 63))) | 1))
                                                _ = v25
                                            }
                                            _ = v24
                                        } else {
                                            v3 = append(v3, ((("\n\t" + "\n\t") + "hello") + (("\n\t" + "") + ("" + ((("" + "hello") + ((("foo bar" + "a") + ("foo bar" + "foo bar")) + "foo bar")) + "hello")))))
                                        }
                                        v26 := GenS2{F0: (v1 * (v1 ^ (v1 | v1)))}
                                        _ = v26
                                    }
                                    _ = v20
                                    _ = v21
                                } else {
                                    if true {
                                        if (((1 > 0) || true) || v5) {
                                            fmt.Println(fi12)
                                            var v27 uint8 = 1
                                            v5 = (v1 < v1)
                                            _ = v27
                                        }
                                        var v28 uint8 = 127
                                        var v29 int16 = -32768
                                        _ = v28
                                        _ = v29
                                    }
                                    var v30 int16 = -32768
                                    if (2 <= 1) {
                                        fmt.Println(v4, v30)
                                    }
                                    _ = v30
                                }
                                _ = fi19
                            }
                        }
                        _ = fi12
                    }
                }
            }
            if false {
                var v31 float32 = (0.0 + (0.0 / (1.5 + 1.5)))
                _ = v31
            }
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    } else {
        if len(v6) > 0 {
            v6[0] = 2147483647
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    var v32 uint32 = 2147483648
    v4 = ((v4 * 1e10) + float64((1e10 + -2.25)))
    var v33 uint16 = v2
    var v34 int = 0
    switch v34 {
    case 2147483647:
        var v35 int = 4294967295
        _ = v35
    case 1:
        var v36 int = v34
        _ = v36
    case -721:
        v5 = (!v5)
    default:
        v1 = v1
    }
    var v37 string = ("hello" + "")
    var v38 uint64 = 18446744073709551615
    var v39 string = v37
    v40 := GenS3{F0: -32768, F1: 127}
    v41 := GenS4{F0: false}
    switch (v34 >> ((0) & 63)) {
    case -1:
        v33 -= 0
    }
    var v42 uint8 = 42
    var v43 uint64 = 42
    var v44 float64 = v4
    var v45 bool = v5
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
