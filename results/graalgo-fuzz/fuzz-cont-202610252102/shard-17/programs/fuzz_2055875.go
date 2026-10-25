package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int8 }

func main() {
    var v1 int = 128
    var v2 bool = true
    var v3 int8 = 73
    var v4 string = "a"
    var v5 uint64 = 560
    switch v1 {
    case 2147483647:
        var v6 int16 = 440
        _ = v6
    }
    v7 := GenS1{F0: v4}
    switch v1 {
    case -9223372036854775808:
        v2 = v2
    default:
        v8 := []int{}
        _ = v8
    }
    v2 = (!false)
    v9 := []int{(v1 + (-(v1))), (v1 / ((32767) | 1)), (v1 & v1), -426}
    if false {
        v10 := GenS2{F0: 945, F1: 0, F2: 2}
        for fi11 := 0; fi11 < 13; fi11++ {
            v2 = (!false)
            switch v4 {
            case "":
                var v12 bool = true
                _ = v12
            case "hello":
                switch 128 {
                case 247:
                    switch 42 {
                    case 0:
                        v13 := GenS2{F0: 255, F1: 1, F2: 0}
                        _ = v13
                    case 32767:
                        for fi14 := 0; fi14 < 1; fi14++ {
                            switch -830 {
                            case -92:
                                if (!v2) {
                                    var v15 float32 = 1e10
                                    _ = v15
                                } else {
                                    _ = len(v9)
                                    fmt.Println(v1)
                                    if v2 {
                                        fmt.Println(v2, v3)
                                        if false {
                                            fmt.Println(v2, v5, v1)
                                            _ = len(v9)
                                        } else {
                                            v9 = append(v9, 2147483647)
                                        }
                                        v3 = v3
                                    }
                                }
                            case 127:
                                if len(v9) > 0 {
                                    v9[0] = (fi14 % ((fi14) | 1))
                                }
                            case 2147483647:
                                switch (^(v1)) {
                                case 101:
                                    fmt.Println(v2, v1)
                                case 406:
                                    v4 = v4
                                }
                            }
                            _ = fi14
                        }
                    case 52:
                        v1 = fi11
                    }
                case -562:
                    switch v4 {
                    case "hello":
                        if v2 {
                            if (!false) {
                                v3 -= v3
                                if v2 {
                                    if v2 {
                                        v3 -= v3
                                    }
                                    v9 = append(v9, v1)
                                }
                            } else {
                                for fi16 := 0; fi16 < 9; fi16++ {
                                    fmt.Println(v1, fi11, fi11)
                                    v1 = v1
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    _ = fi16
                                }
                                switch (fi11 >> ((4294967295) & 63)) {
                                case -537:
                                    v5 = (^(v5))
                                case -222:
                                    v2 = (-2.25 >= (1e10 * -2.25))
                                case 9223372036854775807:
                                    var v17 uint8 = 0
                                    _ = v17
                                default:
                                    v4 = (v4 + ("hello" + ("" + (v4 + (("hello" + ((v4 + v4) + v4)) + ("foo bar" + v4))))))
                                }
                            }
                            if (!false) {
                                v1 = (fi11 / ((v1) | 1))
                                if len(v9) > 0 {
                                    v9[0] = 2147483647
                                }
                            } else {
                                var v18 int16 = 0
                                for fi19 := 0; fi19 < 16; fi19++ {
                                    var v20 int32 = -1
                                    var v21 uint = 0
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                }
                                fmt.Println(v2)
                                _ = v18
                            }
                        } else {
                            var v22 uint32 = 943
                            _ = v22
                        }
                    case "foo bar":
                        fmt.Println(v1, fi11)
                    case "":
                        v23 := GenS2{F0: 0, F1: 9223372036854775807, F2: 1}
                        _ = v23
                    default:
                        fmt.Println(v5, v3)
                    }
                case 4294967295:
                    if ((v4 + (v4 + v4)) == ("" + v4)) {
                        v24 := GenS3{F0: v5, F1: 23}
                        _ = v24
                    }
                }
            case "\n\t":
                var v25 int16 = 1
                _ = v25
            default:
                var v26 float64 = 0.1
                _ = v26
            }
            _ = fi11
        }
        _ = len(v9)
        _ = v10
    } else {
        var v27 int8 = -128
        if (38 == int16(v27)) {
            for fi28 := 0; fi28 < 17; fi28++ {
                v5 = 1
                for fi29 := 0; fi29 < 18; fi29++ {
                    if len(v9) > 0 {
                        v9[0] = (fi29 * v1)
                    }
                    _ = fi29
                }
                v30 := GenS2{F0: 32767, F1: 9223372036854775807, F2: 414}
                _ = fi28
                _ = v30
            }
            var v31 int8 = v27
            var v32 uint8 = 83
            _ = v31
            _ = v32
        } else {
            v2 = (!(!v2))
        }
        _ = v27
    }
    if v2 {
        var v33 uint16 = 729
        v4 = v4
        var v34 uint8 = 42
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
