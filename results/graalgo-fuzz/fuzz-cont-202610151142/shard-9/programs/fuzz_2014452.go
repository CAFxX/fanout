package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int16; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 int; F1 uint8; F2 string }
func genFn0(p0 string, p1 int64, p2 uint8) int16 {
    fmt.Println(p1, p2, p0)
    v5 := GenS1{F0: 18446744073709551615, F1: 2147483647, F2: p1}
    var v6 uint8 = p2
    v7 := GenS2{F0: 42, F1: 65535, F2: 770}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 32767
}

func main() {
    var v1 int32 = 42
    v2 := []int{32768, 42}
    v3 := []int{4294967295}
    v4 := []int{32768, 127}
    var v8 uint8 = 0
    fmt.Println(v8)
    switch 32768 {
    case 710:
        switch (("" + "hello") + "") {
        case "foo bar":
            var v9 int64 = -785
            _ = v9
        case "":
            if true {
                v10 := GenS3{F0: 0, F1: 672, F2: 100}
                v8 = v8
                _ = v10
            } else {
                var v11 int64 = 42
                if false {
                    v1 += (v1 - v1)
                } else {
                    if false {
                        if ((!false) || (673 == 907)) {
                            v8 += (v8 - v8)
                        }
                        v11 = v11
                    } else {
                        if true {
                            if ((v11 + v11) > v11) {
                                var v12 int16 = 31
                                _ = v12
                            } else {
                                var v13 int64 = 9223372036854775807
                                _ = v13
                            }
                        } else {
                            var v14 uint32 = 846
                            fmt.Println(v11, v14, v8)
                            _ = v14
                        }
                        var v15 float32 = (1.5 * ((-2.25 / (float32(0.1) / ((0.1 / (3.14159 / 1.0)) / (1e10 / 3.14159)))) - 1e10))
                        _ = v15
                    }
                    if len(v2) > 0 {
                        v2[0] = 32767
                    }
                    v11 += v11
                }
                _ = v11
            }
        case "a":
            var v16 uint8 = 255
            _ = v16
        }
    }
    v17 := GenS1{F0: 2, F1: 476, F2: -354}
    fmt.Println(v8, v1)
    v2 = append(v2, 4294967295)
    _ = len(v2)
    v18 := GenS4{F0: 676}
    v1 -= (v1 + (v1 | 407))
    switch 4294967295 {
    case 9223372036854775807:
        for fi19 := 0; fi19 < 7; fi19++ {
            v8 = (v8 * (v8 - (^(v8))))
            v1 = (^(v1))
            _ = fi19
        }
    case 32768:
        if ((9223372036854775807 > 42) || false) {
            v1 -= 0
            v1 += v1
            switch 0 {
            case -2147483648:
                v1 -= -2147483648
            case 127:
                v1 -= (v1 * v1)
            case 4294967295:
                var v20 uint8 = (v8 % (((^(v8))) | 1))
                _ = v20
            default:
                fmt.Println(v8, v1)
            }
        } else {
            v21 := GenS4{F0: 134}
            if (0 < 2147483647) {
                switch "" {
                case "hello":
                    if false {
                        var v22 int16 = 1000
                        v23 := GenS5{F0: 255}
                        _ = v22
                        _ = v23
                    }
                case "a":
                    fmt.Println(v1, v1)
                case "\n\t":
                    for fi24 := 0; fi24 < 5; fi24++ {
                        v1 = v1
                        if (1.5 <= (-2.25 + (1e10 - ((0.0 - 1.5) + 1e10)))) {
                            var v25 uint = 2
                            var v26 uint = (v25 >> ((v25) & 63))
                            if true {
                                var v27 uint8 = v8
                                v28 := GenS2{F0: v1, F1: 65535, F2: 1}
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v27
                                _ = v28
                            }
                            _ = v25
                            _ = v26
                        }
                        _ = fi24
                    }
                }
                if false {
                    var v29 uint16 = 65535
                    if false {
                        var v30 int32 = (-(v1))
                        var v31 int8 = -128
                        _ = v30
                        _ = v31
                    }
                    switch 1 {
                    case 355:
                        _ = len(v4)
                    }
                    _ = v29
                } else {
                    v1 = 42
                }
            }
            var v32 int64 = 0
            _ = v21
            _ = v32
        }
    default:
        if false {
            v8 -= v8
        } else {
            v33 := GenS4{F0: 1}
            if len(v2) > 0 {
                v2[0] = -723
            }
            _ = v33
        }
    }
    if (2 >= 523) {
        var v34 int16 = int16(v8)
        _ = v34
    } else {
        _ = len(v3)
        v35 := GenS6{F0: v8, F1: 9223372036854775807}
        var v36 bool = true
        _ = v35
        _ = v36
    }
    v37 := GenS7{F0: -1, F1: (v8 - (^(v8))), F2: ("\n\t" + "a")}
    v1 += v1
    if (-2.25 < 3.14159) {
        var v38 int = 9223372036854775807
        var v39 uint32 = 2147483648
        fmt.Println(v38, v38, v39)
        _ = v38
        _ = v39
    } else {
        var v40 uint32 = 2147483647
        fmt.Println(v1)
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v8
}
