package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 int }

func main() {
    var v1 int = 1
    var v2 float64 = 1e10
    var v3 uint16 = 0
    v4 := make(map[int]bool)
    v5 := make(map[uint32]string)
    v6 := []uint32{0, 638, 773}
    var v7 int16 = 32767
    var v8 float64 = (v2 + 1e10)
    fmt.Println(v1, v3, v2)
    v9 := v4[v1]
    if v9 {
        if true {
            if false {
                fmt.Println(v7)
                v9 = (v9 || v9)
            } else {
                fmt.Println(v9, v2, v8)
                for fi10 := 0; fi10 < 9; fi10++ {
                    v2 = v2
                    var v11 uint = 0
                    for fi12 := 0; fi12 < 8; fi12++ {
                        switch (v1 + 2147483647) {
                        case 0:
                            var v13 int8 = -128
                            _ = v13
                        case 811:
                            _ = len(v6)
                        case 42:
                            if len(v6) > 0 {
                                v6[0] = 338
                            }
                        default:
                            fmt.Println(v7, v7)
                        }
                        if ((true && true) || (93 < 684)) {
                            var v14 int = (fi10 << ((v11) & 63))
                            _ = v14
                        } else {
                            if (((!("foo bar" != "hello")) && v9) && v9) {
                                if true {
                                    switch (("foo bar" + ((("a" + "a") + "foo bar") + ("\n\t" + (("a" + ("" + "hello")) + ("\n\t" + "a"))))) + ("" + "")) {
                                    case "\n\t":
                                        v15 := GenS1{F0: 42}
                                        _ = v15
                                    case "foo bar":
                                        v8 += v8
                                    case "":
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v16 uint8 = 128
                                            _ = v16
                                        case "\n\t":
                                            if (-1 >= 0) {
                                                var v17 int16 = (^(v7))
                                                v2 += (float64(0.0) - v2)
                                                v3 += v3
                                                _ = v17
                                            }
                                        default:
                                            var v18 uint = (v11 % ((v11) | 1))
                                            _ = v18
                                        }
                                    }
                                    switch -2147483648 {
                                    case 127:
                                        var v19 uint = (v11 - (v11 / ((1) | 1)))
                                        _ = v19
                                    case 2147483647:
                                        v20 := v4[fi12]
                                        _ = v20
                                    case 42:
                                        v2 = float64((0.0 - (0.1 * 0.0)))
                                    default:
                                        fmt.Println(v8)
                                    }
                                    var v21 float64 = v2
                                    _ = v21
                                }
                            } else {
                                switch -1 {
                                case 42:
                                    v22 := GenS2{F0: (("hello" + "\n\t") + ("a" + "a"))}
                                    _ = v22
                                case 128:
                                    fmt.Println(fi12, v8)
                                default:
                                    var v23 bool = (18446744073709551615 != 321)
                                    _ = v23
                                }
                                var v24 int64 = -9223372036854775808
                                _ = v24
                            }
                        }
                        _ = fi12
                    }
                    _ = fi10
                    _ = v11
                }
                v25 := v5[639]
                _ = v25
            }
        }
    }
    v26 := GenS3{F0: (-(v1)), F1: v1}
    var v27 int = v1
    for fi28 := 0; fi28 < 12; fi28++ {
        if v9 {
            var v29 int = -2147483648
            _ = v29
        }
        switch ((("hello" + "hello") + ("\n\t" + ("hello" + ""))) + ("foo bar" + "a")) {
        case "":
            for fi30 := 0; fi30 < 5; fi30++ {
                var v31 uint64 = 18446744073709551615
                v1 = (v27 * 9223372036854775807)
                v8 = ((3.14159 / 1e10) * v2)
                _ = fi30
                _ = v31
            }
        case "a":
            v9 = v9
        case "\n\t":
            _ = len(v6)
        }
        var v32 bool = (v9 && ((-128 != -9223372036854775808) && (v9 || v9)))
        _ = fi28
        _ = v32
    }
    switch (("foo bar" + "\n\t") + "") {
    case "\n\t":
        v5[2147483648] = ""
    }
    _ = len(v4)
    fmt.Println(v9)
    var v33 float32 = 1e10
    v34 := v4[(v27 / ((v27) | 1))]
    v8 = (-2.25 - (v8 - 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
