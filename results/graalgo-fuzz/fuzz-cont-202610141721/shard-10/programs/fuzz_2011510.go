package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int32 }
type GenS2 struct { F0 int64; F1 int8; F2 uint64 }
type GenS3 struct { F0 int8; F1 bool; F2 string }

func main() {
    var v1 uint = 0
    var v2 int8 = 100
    v3 := make(map[int]string)
    var v4 float32 = 0.0
    v5 := []string{"hello", "\n\t", ("a" + "\n\t"), ""}
    for fi6 := 0; fi6 < 19; fi6++ {
        if ("foo bar" < ("" + "foo bar")) {
            fmt.Println(v4, fi6)
            fmt.Println(v2)
            var v7 int64 = 1
            _ = v7
        }
        _ = fi6
    }
    switch 32767 {
    case 127:
        if ((!false) && ((!(0.0 > v4)) || (float64(v4) >= 1e10))) {
            v8 := GenS1{F0: false, F1: "foo bar", F2: 2147483647}
            v2 = v2
            switch ("" + (("foo bar" + ("" + "")) + "")) {
            case "":
                for fi9 := 0; fi9 < 14; fi9++ {
                    var v10 int64 = 999
                    v2 += (-(v2))
                    _ = fi9
                    _ = v10
                }
            case "foo bar":
                var v11 int8 = (v2 * (v2 * v2))
                _ = v11
            case "\n\t":
                var v12 float32 = v4
                _ = v12
            }
            _ = v8
        }
    case 9223372036854775807:
        for fi13 := 0; fi13 < 18; fi13++ {
            switch "foo bar" {
            case "":
                switch ((("\n\t" + "a") + ("a" + ("foo bar" + ""))) + "") {
                case "foo bar":
                    v14 := GenS1{F0: true, F1: ("a" + ("\n\t" + ("a" + "foo bar"))), F2: -1}
                    _ = v14
                }
            case "a":
                v4 = v4
            case "foo bar":
                if (fi13 >= (fi13 % ((42) | 1))) {
                    v15 := GenS1{F0: true, F1: "a", F2: -2147483648}
                    switch "hello" {
                    case "a":
                        v16 := GenS2{F0: 42, F1: (-(v2)), F2: 945}
                        _ = v16
                    case "\n\t":
                        fmt.Println(fi13, v4, fi13)
                    case "hello":
                        var v17 uint = (v1 ^ (v1 % ((v1) | 1)))
                        _ = v17
                    default:
                        v2 = v2
                    }
                    var v18 uint8 = 128
                    _ = v15
                    _ = v18
                } else {
                    v19 := GenS1{F0: (42 != 32767), F1: "a", F2: -594}
                    _ = v19
                }
            }
            _ = fi13
        }
    case 32767:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        for fi20 := 0; fi20 < 2; fi20++ {
            for fi21 := 0; fi21 < 1; fi21++ {
                var v22 float32 = (0.1 * (v4 + 1e10))
                v23 := GenS2{F0: -9223372036854775808, F1: (v2 & v2), F2: 0}
                _ = fi21
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 string = "hello"
                v5 = append(v5, (v25 + "hello"))
                v26 := GenS2{F0: 0, F1: (^(v2)), F2: 42}
                v27 := GenS1{F0: (302 == -194), F1: "foo bar", F2: -1}
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
            _ = fi20
        }
    }
    var v28 int32 = -626
    v2 += -1
    v4 += v4
    var v29 int64 = 9223372036854775807
    _ = len(v3)
    v2 -= 1
    v30 := GenS3{F0: v2, F1: (0.0 != ((v4 + (v4 - 1.5)) * v4)), F2: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
