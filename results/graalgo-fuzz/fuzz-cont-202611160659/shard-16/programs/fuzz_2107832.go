package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint16 = 65535
    var v2 int = 127
    var v3 uint16 = v1
    var v4 int8 = 73
    v1 = v1
    v5 := []int{-823}
    v2 += 141
    v2 = v2
    for fi6 := 0; fi6 < 9; fi6++ {
        if len(v5) > 0 {
            v5[0] = 295
        }
        v2 -= 4294967295
        if false {
            switch "\n\t" {
            case "a":
                switch "foo bar" {
                case "hello":
                    var v7 int32 = 42
                    _ = v7
                case "":
                    fmt.Println(v1)
                default:
                    _ = len(v5)
                }
            case "":
                for fi8 := 0; fi8 < 14; fi8++ {
                    switch "" {
                    case "\n\t":
                        if false {
                            if true {
                                switch (("a" + "") + ("a" + ((("" + "a") + "hello") + ("\n\t" + "foo bar")))) {
                                case "hello":
                                    var v9 uint64 = 453
                                    _ = v9
                                case "a":
                                    var v10 uint64 = 18446744073709551615
                                    _ = v10
                                }
                                if (!(("hello" + ("hello" + "")) == "hello")) {
                                    if ((v2 >= (fi6 * fi8)) || ("hello" < "\n\t")) {
                                        _ = len(v5)
                                        fmt.Println(v4)
                                    } else {
                                        var v11 int16 = -1
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = v11
                                    }
                                }
                                if false {
                                    v1 = v3
                                    v4 += -128
                                    fmt.Println(fi6, v4)
                                } else {
                                    switch "\n\t" {
                                    case "":
                                        v1 -= v3
                                    case "hello":
                                        var v12 uint16 = (v1 >> ((42) & 63))
                                        _ = v12
                                    case "foo bar":
                                        v5 = append(v5, (fi8 / ((v2) | 1)))
                                    default:
                                        v4 = 100
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    case "hello":
                                        if true {
                                            var v13 int = v2
                                            _ = v13
                                        }
                                    }
                                }
                                v1 -= v1
                                var v14 float32 = 0.1
                                _ = v14
                            } else {
                                v15 := GenS1{F0: false, F1: int8(v1), F2: (v3 >> ((1) & 63))}
                                var v16 uint16 = (v3 | v1)
                                _ = v15
                                _ = v16
                            }
                            var v17 float64 = -2.25
                            _ = v17
                        } else {
                            if true {
                                var v18 int = (fi8 - v2)
                                v18 = 2147483647
                                var v19 float32 = ((3.14159 - 3.14159) - -2.25)
                                var v20 int = (-(v2))
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                        }
                    case "hello":
                        var v21 uint32 = 1
                        _ = v21
                    }
                    v22 := GenS1{F0: true, F1: v4, F2: v3}
                    var v23 int8 = (v4 >> ((550) & 63))
                    _ = fi8
                    _ = v22
                    _ = v23
                }
            case "foo bar":
                v24 := GenS1{F0: (2 != 42), F1: v4, F2: v1}
                _ = v24
            }
            switch ("" + "") {
            case "foo bar":
                v1 = v1
            case "hello":
                v1 -= v1
            case "":
                for fi25 := 0; fi25 < 7; fi25++ {
                    v2 = (fi25 ^ 32767)
                    _ = len(v5)
                    _ = fi25
                }
            default:
                if true {
                    fmt.Println(fi6, v1)
                    fmt.Println(v2)
                } else {
                    if (-9223372036854775808 < (fi6 / ((752) | 1))) {
                        v4 += (v4 & (-(v4)))
                    }
                }
            }
        } else {
            var v26 string = "\n\t"
            _ = v26
        }
        _ = fi6
    }
    v2 -= v2
    for fi27 := 0; fi27 < 2; fi27++ {
        fmt.Println(v4, fi27, v4)
        _ = fi27
    }
    v28 := GenS2{F0: 1}
    var v29 uint = 42
    if len(v5) > 0 {
        v5[0] = 128
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
