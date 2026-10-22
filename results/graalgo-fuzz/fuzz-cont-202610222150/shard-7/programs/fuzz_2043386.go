package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int64; F1 int64; F2 int64 }
func genFn0(p0 uint, p1 int, p2 uint8) uint8 {
    var v6 int32 = 159
    var v7 string = ("foo bar" + (("foo bar" + "a") + "hello"))
    fmt.Println(p1, v7, v7)
    v8 := GenS1{F0: 127, F1: (v7 + "foo bar")}
    if ((4294967295 == 0) || true) {
        var v9 int64 = 619
        _ = v9
    } else {
        var v10 uint8 = (p2 % ((127) | 1))
        fmt.Println(v10)
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return p2
}

func main() {
    v1 := []string{"a", ("hello" + ""), ("" + "hello"), "\n\t"}
    var v2 uint8 = 231
    var v3 uint = 2
    var v4 float32 = float32(1e10)
    var v5 uint16 = 433
    var v11 uint16 = v5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v12 int16 = 42
    _ = len(v1)
    v5 += v5
    for fi13 := 0; fi13 < 1; fi13++ {
        if false {
            fmt.Println(v2, fi13, v4)
        }
        if (!(false && true)) {
            v3 = 255
            if false {
                v14 := GenS1{F0: 42, F1: "hello"}
                v3 += v3
                v15 := GenS2{F0: 508}
                _ = v14
                _ = v15
            } else {
                v3 += (v3 * v3)
                if (0 > -128) {
                    v16 := GenS3{F0: 722, F1: 65535}
                    v17 := GenS4{F0: 656, F1: 100}
                    switch "a" {
                    case "hello":
                        for fi18 := 0; fi18 < 4; fi18++ {
                            v5 += v5
                            v19 := GenS2{F0: 923}
                            if false {
                                var v20 string = "foo bar"
                                _ = v20
                            }
                            _ = fi18
                            _ = v19
                        }
                    case "a":
                        fmt.Println(v3, v2, v3)
                    case "":
                        v12 = v12
                    }
                    _ = v16
                    _ = v17
                }
                var v21 float64 = 0.0
                _ = v21
            }
            v22 := GenS5{F0: 42, F1: 80, F2: -1}
            _ = v22
        } else {
            if true {
                v4 = v4
            } else {
                fmt.Println(v12)
                switch ("\n\t" + ("foo bar" + "foo bar")) {
                case "a":
                    _ = len(v1)
                case "hello":
                    v4 = v4
                case "foo bar":
                    switch "hello" {
                    case "a":
                        var v23 int16 = (v12 ^ (v12 / ((v12) | 1)))
                        _ = v23
                    case "":
                        var v24 int64 = 42
                        _ = v24
                    case "\n\t":
                        var v25 int = (fi13 | 9223372036854775807)
                        _ = v25
                    }
                default:
                    v26 := GenS5{F0: -712, F1: -9223372036854775808, F2: 1}
                    _ = v26
                }
            }
            var v27 int32 = 2147483647
            var v28 int64 = -9223372036854775808
            _ = v27
            _ = v28
        }
        _ = fi13
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        v30 := GenS3{F0: 9223372036854775807, F1: v11}
        _ = fi29
        _ = v30
    }
    fmt.Println(v4, v5, v12)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
