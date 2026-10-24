package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 uint64; F2 int }

func main() {
    v1 := []string{"foo bar", "hello"}
    var v2 bool = (-955 <= 0)
    v3 := make(map[uint32]int)
    var v4 int = -76
    var v5 int8 = -128
    v4 += -116
    fmt.Println(v2)
    v2 = true
    if (32767 >= 104) {
        var v6 int64 = -585
        var v7 float32 = 0.1
        _ = v6
        _ = v7
    }
    var v8 int32 = 36
    var v9 uint8 = 150
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    switch (v4 - (v4 % ((32767) | 1))) {
    case -2147483648:
        fmt.Println(v5)
    case -604:
        if (((v8 ^ v8) < v8) && ((v8 >> ((1) & 63)) > (v8 << ((0) & 63)))) {
            v10 := GenS1{F0: v2, F1: (((("a" + "hello") + "") + (("foo bar" + ("" + "")) + "a")) + "")}
            v11 := GenS2{F0: "hello", F1: v9}
            var v12 int64 = 491
            _ = v10
            _ = v11
            _ = v12
        } else {
            v13 := GenS1{F0: (((!true) && (((0.0 - float32((-2.25 + float64(-2.25)))) * 1.5) == 1e10)) || (((840 >= 178) && (0 < 316)) || (884 > 2))), F1: (("\n\t" + "hello") + "")}
            v9 += 50
            _ = v13
        }
    case 0:
        v14 := GenS3{F0: v2, F1: -9223372036854775808}
        _ = v14
    }
    if true {
        v1 = append(v1, (("\n\t" + "\n\t") + "foo bar"))
        switch ("a" + "foo bar") {
        case "hello":
            switch 1 {
            case -96:
                fmt.Println(v5, v5)
            case -1:
                v15 := GenS2{F0: ("hello" + "hello"), F1: 17}
                _ = v15
            case 2147483647:
                v5 += 127
            default:
                v16 := GenS1{F0: (997 == 1), F1: "hello"}
                _ = v16
            }
        case "a":
            switch "foo bar" {
            case "":
                v17 := v3[0]
                _ = v17
            case "hello":
                var v18 uint32 = 2
                _ = v18
            case "\n\t":
                var v19 float64 = 1.5
                _ = v19
            default:
                fmt.Println(v2)
            }
        case "\n\t":
            v2 = (-2.25 >= 1e10)
        }
        var v20 int32 = 2147483647
        _ = v20
    }
    v5 = v5
    v9 = v9
    v21 := GenS4{F0: 359}
    fmt.Println(v4, v5, v5)
    var v22 uint32 = 78
    var v23 uint8 = 59
    for fi24 := 0; fi24 < 15; fi24++ {
        if false {
            if (740 <= 42) {
                var v25 int32 = 42
                _ = v25
            }
            if true {
                v4 += (fi24 ^ (v4 ^ v4))
                fmt.Println(v4)
                if (v23 >= (^(v23))) {
                    _ = len(v1)
                    for fi26 := 0; fi26 < 7; fi26++ {
                        if true {
                            v2 = ((31 >= 255) || ((^(v8)) <= v8))
                            var v27 uint8 = (v9 + uint8(v22))
                            v1 = append(v1, "\n\t")
                            _ = v27
                        }
                        _ = fi26
                    }
                } else {
                    fmt.Println(fi24)
                }
            } else {
                v28 := GenS5{F0: (^(v22)), F1: 1, F2: (^(fi24))}
                v8 += v8
                for fi29 := 0; fi29 < 10; fi29++ {
                    v8 = v8
                    _ = len(v3)
                    var v30 int16 = 227
                    _ = fi29
                    _ = v30
                }
                _ = v28
            }
            v31 := GenS3{F0: (v22 >= (v22 ^ 42)), F1: 0}
            _ = v31
        }
        var v32 int16 = -676
        _ = fi24
        _ = v32
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
