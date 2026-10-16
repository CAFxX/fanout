package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 uint32 }
func genFn0() bool {
    v5 := GenS1{F0: 1, F1: 2, F2: 1}
    if false {
        if true {
            var v6 int64 = 42
            _ = v6
        }
    }
    _ = v5
    return (0 > 0)
}

func main() {
    v1 := []string{"\n\t"}
    var v2 bool = (2147483647 < 42)
    v3 := []int{}
    v4 := []string{"a", "a", "foo bar", ""}
    v2 = (173 != 0)
    v2 = v2
    v2 = ((!(v2 || ((!v2) || (1 <= 127)))) || v2)
    v2 = (v2 || true)
    v2 = ("foo bar" <= "")
    v7 := GenS1{F0: -1, F1: 2, F2: -627}
    var v8 int16 = 1
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            if (!(((v8 - (v8 >> ((0) & 63))) >= 402) || (v8 == (v8 + v8)))) {
                var v11 uint16 = 0
                fmt.Println(v11, v11, fi9)
                _ = v11
            }
            v12 := GenS1{F0: -1, F1: 1, F2: (^(v8))}
            v8 += v8
            _ = fi10
            _ = v12
        }
        var v13 int16 = v8
        _ = fi9
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        switch "" {
        case "foo bar":
            v4 = append(v4, "\n\t")
        case "":
            for fi15 := 0; fi15 < 4; fi15++ {
                _ = len(v1)
                _ = fi15
            }
        case "a":
            for fi16 := 0; fi16 < 15; fi16++ {
                switch 4294967295 {
                case -242:
                    v17 := GenS1{F0: -17, F1: 255, F2: v8}
                    _ = v17
                case 0:
                    v18 := GenS1{F0: -1, F1: 2, F2: -1}
                    _ = v18
                }
                switch 515 {
                case 32767:
                    var v19 int16 = v8
                    _ = v19
                case -9223372036854775808:
                    v20 := GenS1{F0: 42, F1: 128, F2: (v8 & v8)}
                    _ = v20
                case 0:
                    v8 += v8
                }
                v21 := GenS2{F0: -78}
                _ = fi16
                _ = v21
            }
        default:
            fmt.Println(v8, v2, v8)
        }
        fmt.Println(fi14)
        var v22 int16 = v8
        _ = fi14
        _ = v22
    }
    switch "\n\t" {
    case "":
        if true {
            if (0 != 0) {
                switch "\n\t" {
                case "foo bar":
                    v8 = 480
                default:
                    for fi23 := 0; fi23 < 4; fi23++ {
                        var v24 string = ("" + (("" + ("" + ("foo bar" + ""))) + "hello"))
                        _ = fi23
                        _ = v24
                    }
                }
                var v25 float64 = 0.0
                _ = v25
            }
        } else {
            v26 := GenS1{F0: 0, F1: 127, F2: v8}
            var v27 string = "foo bar"
            _ = len(v1)
            _ = v26
            _ = v27
        }
    case "\n\t":
        var v28 uint64 = 0
        _ = v28
    case "a":
        switch -842 {
        case 2147483647:
            v29 := GenS3{F0: 811, F1: 2}
            _ = v29
        case -556:
            _ = len(v4)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
