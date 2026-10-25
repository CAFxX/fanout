package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 bool; F1 uint64; F2 string }
func genFn0() int8 {
    var v5 int16 = 42
    v5 = v5
    v6 := GenS1{F0: 117, F1: 2147483647}
    _ = v5
    _ = v6
    return 127
}

func main() {
    var v1 uint64 = 2
    var v2 uint8 = 128
    var v3 float64 = 1.5
    v4 := make(map[uint32]bool)
    var v7 uint8 = 4
    _ = len(v4)
    v1 += 1
    v8 := []int{9223372036854775807, 134, -916, 9223372036854775807}
    fmt.Println(v3)
    for fi9 := 0; fi9 < 11; fi9++ {
        fmt.Println(v1, v7)
        switch (fi9 & -65) {
        case 982:
            if false {
                v4[0] = (v1 >= v1)
            } else {
                v10 := GenS1{F0: 1, F1: 1}
                var v11 int16 = 0
                _ = len(v4)
                _ = v10
                _ = v11
            }
        }
        _ = fi9
    }
    var v12 uint = 65535
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS1{F0: 79, F1: 2}
        v15 := GenS2{F0: (fi13 ^ 643), F1: (v12 >> ((86) & 63)), F2: v1}
        _ = fi13
        _ = v14
        _ = v15
    }
    if false {
        v7 = (v2 % (((v2 ^ v2)) | 1))
        var v16 bool = (v7 < v2)
        v8 = append(v8, 9223372036854775807)
        _ = v16
    } else {
        switch -664 {
        case 32767:
            var v17 int16 = 666
            _ = v17
        }
        v12 += (v12 * (v12 * v12))
        _ = len(v4)
    }
    switch "hello" {
    case "hello":
        v3 -= (-2.25 / ((v3 - (v3 - (v3 / (-2.25 / v3)))) + -2.25))
    default:
        if true {
            var v18 int32 = -2147483648
            var v19 string = "\n\t"
            if (127 >= -2) {
                v4[4294967295] = (int8(v2) > genFn0())
            }
            _ = v18
            _ = v19
        }
    }
    switch 32768 {
    case 128:
        if len(v8) > 0 {
            v8[0] = -1
        }
    case 487:
        var v20 int64 = 1
        _ = v20
    default:
        switch ("a" + (((("foo bar" + ("a" + ("foo bar" + "foo bar"))) + "a") + "a") + (("foo bar" + ("a" + (((("foo bar" + ("foo bar" + ("" + (("hello" + "a") + "\n\t")))) + "foo bar") + (("hello" + ("a" + "hello")) + "\n\t")) + ""))) + "foo bar"))) {
        case "a":
            if (1 > 516) {
                if (!false) {
                    if true {
                        v21 := GenS3{F0: (v12 << ((v12) & 63)), F1: "\n\t"}
                        if false {
                            var v22 int64 = -9223372036854775808
                            switch "hello" {
                            case "a":
                                var v23 string = "hello"
                                _ = v23
                            case "hello":
                                for fi24 := 0; fi24 < 6; fi24++ {
                                    v1 = (^(v1))
                                    v25 := GenS2{F0: fi24, F1: v12, F2: 349}
                                    v26 := GenS3{F0: 897, F1: ""}
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                            case "\n\t":
                                var v27 uint8 = v2
                                _ = v27
                            }
                            _ = v22
                        }
                        v28 := GenS4{F0: (true && false), F1: (v1 + v1), F2: "foo bar"}
                        _ = v21
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 6; fi29++ {
                        v3 += float64(3.14159)
                        var v30 int32 = -309
                        _ = fi29
                        _ = v30
                    }
                }
                var v31 int = 9223372036854775807
                var v32 uint16 = 42
                _ = v31
                _ = v32
            }
        case "hello":
            fmt.Println(v7, v2, v1)
        case "\n\t":
            if true {
                v33 := v4[224]
                var v34 int64 = 822
                if true {
                    var v35 uint64 = uint64(v12)
                    _ = v35
                }
                _ = v33
                _ = v34
            }
        default:
            var v36 float32 = -2.25
            _ = v36
        }
    }
    var v37 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v7
    _ = v8
}
