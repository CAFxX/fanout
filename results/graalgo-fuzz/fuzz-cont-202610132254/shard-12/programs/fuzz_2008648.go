package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int32 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 string; F1 int32; F2 int32 }

func main() {
    var v1 uint64 = 42
    var v2 int64 = 1
    v3 := make(map[string]bool)
    var v4 int16 = 42
    var v5 bool = (31 > 32767)
    v5 = false
    v6 := GenS1{F0: 26, F1: v4, F2: 42}
    if (false || v5) {
        if v5 {
            fmt.Println(v4)
            v1 -= 710
        } else {
            _ = len(v3)
            v1 += v1
            for fi7 := 0; fi7 < 10; fi7++ {
                delete(v3, ("\n\t" + "foo bar"))
                var v8 uint16 = 255
                for fi9 := 0; fi9 < 13; fi9++ {
                    v10 := GenS1{F0: 2, F1: (v4 + v4), F2: -2147483648}
                    if false {
                        v11 := []int{(fi9 / ((fi9) | 1)), fi7}
                        _ = v11
                    } else {
                        v12 := []int{int(v1), (fi7 >> ((uint(v1)) & 63)), -9223372036854775808, fi9}
                        v4 = v4
                        var v13 int = fi9
                        _ = v12
                        _ = v13
                    }
                    var v14 float64 = -2.25
                    _ = fi9
                    _ = v10
                    _ = v14
                }
                _ = fi7
                _ = v8
            }
        }
        for fi15 := 0; fi15 < 8; fi15++ {
            switch "a" {
            case "foo bar":
                fmt.Println(v4, v1)
            case "\n\t":
                switch 42 {
                case -2147483648:
                    fmt.Println(fi15)
                case 0:
                    if false {
                        for fi16 := 0; fi16 < 18; fi16++ {
                            v4 += v4
                            if false {
                                var v17 uint8 = 127
                                v18 := GenS2{F0: (v5 && (-464 == (v2 * 543))), F1: -2147483648}
                                var v19 string = "a"
                                delete(v3, "\n\t")
                                _ = v17
                                _ = v18
                                _ = v19
                            }
                            _ = fi16
                        }
                        v1 = v1
                    } else {
                        v20 := []int{(fi15 ^ fi15), fi15, 202, fi15}
                        v3[""] = (!false)
                        v21 := GenS1{F0: 255, F1: v4, F2: 359}
                        _ = v20
                        _ = v21
                    }
                }
            case "hello":
                var v22 bool = (fi15 >= fi15)
                _ = v22
            }
            var v23 int8 = 104
            _ = fi15
            _ = v23
        }
        fmt.Println(v1, v4)
    }
    if (v5 || (v5 && v5)) {
        fmt.Println(v1, v4)
        v2 = (v2 | v2)
        v1 -= (v1 / ((2) | 1))
    } else {
        if (0 >= -1) {
            v24 := GenS2{F0: v5, F1: 42}
            var v25 float64 = 3.14159
            _ = v24
            _ = v25
        }
    }
    if (v2 < 846) {
        delete(v3, "foo bar")
    } else {
        _ = len(v3)
    }
    var v26 float32 = float32(0.1)
    v1 -= 4294967295
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 string = ""
        _ = fi27
        _ = v28
    }
    if ((129 != 255) && v5) {
        v29 := GenS1{F0: 136, F1: (v4 - v4), F2: 42}
        _ = v29
    }
    var v30 int = 2147483647
    if ("" <= "\n\t") {
        v31 := GenS3{F0: "foo bar", F1: 632, F2: 0}
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
