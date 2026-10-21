package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int8; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 int64; F2 uint8 }
type GenS5 struct { F0 int64 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []string{}
    var v3 uint64 = 42
    v1 = 0
    var v4 int = 35
    v4 += v4
    _ = len(v2)
    var v5 uint64 = v1
    _ = len(v2)
    fmt.Println(v5)
    v6 := GenS1{F0: -780}
    switch 128 {
    case 42:
        v1 = v5
    }
    if true {
        if true {
            var v7 float32 = 3.14159
            _ = len(v2)
            var v8 uint16 = 0
            _ = v7
            _ = v8
        }
        fmt.Println(v3)
    }
    if true {
        if (2 != 2) {
            var v9 int16 = 1
            _ = v9
        }
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := GenS1{F0: -451}
            fmt.Println(fi10)
            if (71 >= 255) {
                if (42 >= 302) {
                    var v12 float64 = 0.0
                    for fi13 := 0; fi13 < 16; fi13++ {
                        if false {
                            if len(v2) > 0 {
                                v2[0] = "a"
                            }
                            var v14 bool = (v1 > (^(v1)))
                            fmt.Println(v4)
                            _ = v14
                        } else {
                            var v15 float32 = ((0.1 + -2.25) + 0.0)
                            fmt.Println(v3, v4)
                            _ = v15
                        }
                        if (uint16(v1) <= 217) {
                            if (false || (false && true)) {
                                var v16 uint64 = 18446744073709551615
                                if true {
                                    if false {
                                        var v17 float64 = 1.5
                                        v18 := GenS2{F0: 2, F1: -77, F2: 42}
                                        v19 := GenS3{F0: 32767}
                                        _ = v17
                                        _ = v18
                                        _ = v19
                                    }
                                    v2 = append(v2, "\n\t")
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if true {
                                        if true {
                                            v3 = (v1 & (^(v1)))
                                            v20 := GenS2{F0: 128, F1: 100, F2: -9223372036854775808}
                                            var v21 int32 = 42
                                            _ = v20
                                            _ = v21
                                        }
                                        if ((!(false && (false && (67 >= 42)))) || (-641 < 0)) {
                                            var v22 int = v4
                                            _ = v22
                                        }
                                        var v23 int8 = 1
                                        _ = v23
                                    }
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = ("\n\t" + "\n\t")
                                    }
                                    fmt.Println(v12, v3)
                                }
                                _ = v16
                            } else {
                                v12 = ((0.0 / v12) / (float64((0.1 * 0.1)) / v12))
                            }
                            var v24 string = (("hello" + ("" + "hello")) + (("foo bar" + ("\n\t" + ("" + "foo bar"))) + ""))
                            if true {
                                switch v24 {
                                case "\n\t":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case "a":
                                    _ = len(v2)
                                case "hello":
                                    var v25 uint16 = 32767
                                    var v26 float32 = 0.0
                                    v25 = (^(v25))
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v25
                                    _ = v26
                                default:
                                    var v27 int64 = 9223372036854775807
                                    _ = v27
                                }
                                var v28 float32 = -2.25
                                _ = v28
                            }
                            _ = v24
                        } else {
                            v5 = v3
                            var v29 string = "\n\t"
                            var v30 uint8 = 42
                            _ = v29
                            _ = v30
                        }
                        _ = len(v2)
                        _ = fi13
                    }
                    _ = v12
                }
                v31 := GenS1{F0: -9223372036854775808}
                if false {
                    v32 := GenS1{F0: 42}
                    _ = v32
                } else {
                    if true {
                        if true {
                            var v33 float64 = 0.0
                            _ = v33
                        }
                        v34 := GenS4{F0: ((v1 * (v3 | v5)) != (v5 | (v3 % ((v1) | 1)))), F1: -917, F2: 0}
                        if (4294967295 >= (fi10 ^ -2147483648)) {
                            var v35 float64 = 0.1
                            _ = v35
                        } else {
                            v36 := GenS4{F0: (!true), F1: -1, F2: 69}
                            var v37 bool = ((255 != 0) && false)
                            _ = v36
                            _ = v37
                        }
                        _ = v34
                    } else {
                        v1 = (v1 * v1)
                        if (1 == 255) {
                            fmt.Println(v3, v5, v4)
                        }
                        v38 := GenS5{F0: -1}
                        _ = v38
                    }
                    for fi39 := 0; fi39 < 14; fi39++ {
                        v40 := GenS2{F0: 128, F1: 42, F2: -9223372036854775808}
                        var v41 uint8 = 216
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                }
                _ = v31
            } else {
                var v42 bool = (42 <= -9223372036854775808)
                v1 += (v5 - (v3 << ((255) & 63)))
                _ = v42
            }
            _ = fi10
            _ = v11
        }
    }
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    v4 = (v4 >> ((53) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
