package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int8; F1 string }

func main() {
    var v1 uint16 = 65535
    var v2 int32 = 42
    v3 := make(map[string]string)
    v4 := []string{"\n\t", ("a" + "")}
    v5 := []uint32{2147483648, 2147483648, 65535}
    if false {
        var v6 int32 = -2147483648
        if (1e10 > 1e10) {
            v5 = append(v5, 42)
        } else {
            var v7 int = 9223372036854775807
            var v8 int32 = v2
            _ = len(v3)
            _ = v7
            _ = v8
        }
        _ = v6
    }
    if (2 >= 1) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi9 := 0; fi9 < 5; fi9++ {
            v2 += 42
            v10 := GenS1{F0: 2, F1: 42, F2: 32767}
            if (false || (0 <= 974)) {
                v2 = v2
                v1 += 32767
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: 18446744073709551615, F1: 2, F2: v1}
        _ = v11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        v1 = (v1 & 799)
        if false {
            if false {
                _ = len(v3)
                if ((!(v2 <= -2147483648)) || (false && false)) {
                    for fi13 := 0; fi13 < 13; fi13++ {
                        if false {
                            var v14 int8 = 89
                            v15 := GenS1{F0: 4294967295, F1: 128, F2: (v1 * 963)}
                            v1 = (v1 * (^(v1)))
                            if true {
                                if (v2 >= v2) {
                                    v16 := GenS2{F0: (v14 % ((v14) | 1)), F1: (("" + ("\n\t" + "")) + ("\n\t" + ("hello" + "\n\t")))}
                                    v1 -= v1
                                    if (!(2 != 419)) {
                                        var v17 uint8 = 2
                                        var v18 uint8 = v17
                                        v19 := GenS2{F0: (v14 / ((101) | 1)), F1: ("" + ("" + "hello"))}
                                        _ = v17
                                        _ = v18
                                        _ = v19
                                    } else {
                                        switch ("foo bar" + "foo bar") {
                                        case "":
                                            v1 += (^(v1))
                                        }
                                        if (!(42 != 42)) {
                                            fmt.Println(v14, v2)
                                            var v20 float64 = ((3.14159 / 1.5) / (0.0 + (1e10 - (0.1 + 1.5))))
                                            _ = v20
                                        } else {
                                            if false {
                                                var v21 uint = 2
                                                _ = v21
                                            } else {
                                                v2 += v2
                                            }
                                            var v22 bool = true
                                            v4 = append(v4, ((("foo bar" + "a") + ("hello" + "")) + (("\n\t" + "\n\t") + ("\n\t" + "hello"))))
                                            _ = v22
                                        }
                                    }
                                    if (!false) {
                                        v1 = v1
                                        v1 += (v1 | 255)
                                        var v23 int16 = 32767
                                        _ = v23
                                    }
                                    _ = v16
                                }
                                var v24 uint32 = 2147483648
                                v14 -= v14
                                _ = v24
                            } else {
                                switch "a" {
                                case "foo bar":
                                    var v25 int16 = 1000
                                    if true {
                                        v26 := GenS1{F0: 42, F1: 98, F2: (v1 << ((42) & 63))}
                                        var v27 int8 = v14
                                        _ = v26
                                        _ = v27
                                    } else {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    _ = v25
                                }
                                v1 += v1
                            }
                            _ = v14
                            _ = v15
                        } else {
                            _ = len(v4)
                            if len(v5) > 0 {
                                v5[0] = 0
                            }
                        }
                        v2 -= 430
                        _ = fi13
                    }
                }
                fmt.Println(v2)
            }
        }
        _ = fi12
    }
    v28 := GenS1{F0: 230, F1: 189, F2: (v1 | (^(v1)))}
    if true {
        var v29 uint = 513
        v1 = 496
        var v30 float64 = float64(0.0)
        _ = v29
        _ = v30
    } else {
        v31 := GenS1{F0: 1, F1: 2, F2: v1}
        _ = v31
    }
    _ = len(v4)
    v2 = (v2 + v2)
    var v32 int64 = 9223372036854775807
    var v33 uint8 = 128
    fmt.Println(v32, v33, v33)
    fmt.Println(v33, v2)
    v32 = v32
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
