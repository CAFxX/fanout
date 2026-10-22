package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int32; F1 int16; F2 uint64 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 514
    v3 := []string{"hello", "\n\t", "foo bar"}
    var v4 uint = v2
    var v5 string = "foo bar"
    var v6 uint8 = 255
    var v7 int32 = 79
    var v8 uint = (v2 | v2)
    v9 := GenS1{F0: (v7 ^ v7), F1: 1000}
    v6 += v6
    fmt.Println(v2, v8, v8)
    var v10 int8 = 105
    if false {
        v11 := GenS1{F0: v7, F1: 214}
        _ = v11
    } else {
        delete(v1, ("a" + v5))
        var v12 uint = (^(v2))
        v7 -= v7
        _ = v12
    }
    if true {
        fmt.Println(v8, v10)
        switch 32767 {
        case 4294967295:
            v5 = "\n\t"
        case -725:
            for fi13 := 0; fi13 < 19; fi13++ {
                _ = len(v1)
                var v14 int8 = v10
                switch fi13 {
                case -227:
                    if (!true) {
                        v6 = (^(v6))
                    }
                case 2147483647:
                    for fi15 := 0; fi15 < 7; fi15++ {
                        v16 := GenS2{F0: 0, F1: v5}
                        var v17 float64 = 0.1
                        v4 += v8
                        if (!true) {
                            v18 := GenS1{F0: -1, F1: -1}
                            var v19 int8 = v10
                            _ = v18
                            _ = v19
                        }
                        if false {
                            fmt.Println(fi15)
                            var v20 int8 = v10
                            _ = v20
                        } else {
                            var v21 string = (v5 + "")
                            v14 = -128
                            v10 -= v14
                            _ = v21
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                default:
                    fmt.Println(v6)
                }
                _ = fi13
                _ = v14
            }
        case -9223372036854775808:
            v22 := GenS3{F0: (v7 - 368), F1: 792, F2: 42}
            _ = v22
        }
    } else {
        v10 = v10
        var v23 uint16 = 695
        _ = v23
    }
    v6 = 162
    switch -31 {
    case 42:
        for fi24 := 0; fi24 < 4; fi24++ {
            var v25 float64 = ((((1.5 * float64(-2.25)) * 3.14159) + -2.25) + 1.5)
            v6 = (v6 % (((v6 >> (((v4 + 186)) & 63))) | 1))
            _ = fi24
            _ = v25
        }
    case -2147483648:
        fmt.Println(v8, v2, v10)
    case -9223372036854775808:
        if len(v3) > 0 {
            v3[0] = ("hello" + "hello")
        }
    default:
        _ = len(v3)
    }
    v2 -= (v8 % ((v8) | 1))
    var v26 int32 = (v7 + (v7 * v7))
    switch 128 {
    case -444:
        v27 := GenS3{F0: (v7 % ((2147483647) | 1)), F1: 32767, F2: 2}
        _ = v27
    default:
        var v28 uint16 = 32767
        _ = v28
    }
    if (42 >= -9223372036854775808) {
        if false {
            for fi29 := 0; fi29 < 19; fi29++ {
                v30 := GenS4{F0: -1, F1: (v5 + v5)}
                _ = fi29
                _ = v30
            }
            fmt.Println(v7, v6, v10)
            var v31 uint32 = 2
            _ = v31
        } else {
            v32 := GenS5{F0: 2}
            if len(v3) > 0 {
                v3[0] = ("a" + ("hello" + (v5 + v5)))
            }
            if false {
                fmt.Println(v4, v10)
                fmt.Println(v8, v2, v4)
                for fi33 := 0; fi33 < 2; fi33++ {
                    var v34 int8 = (v10 * v10)
                    v4 = 65535
                    v35 := GenS3{F0: (^(v26)), F1: 42, F2: 4294967295}
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            } else {
                v26 += 80
                var v36 string = "foo bar"
                _ = v36
            }
            _ = v32
        }
        switch (((("a" + v5) + (((("\n\t" + v5) + "hello") + (("hello" + v5) + ("" + v5))) + v5)) + v5) + v5) {
        case "hello":
            for fi37 := 0; fi37 < 1; fi37++ {
                v38 := v1[v5]
                var v39 uint64 = 173
                v38 += (-(fi37))
                _ = fi37
                _ = v38
                _ = v39
            }
        case "\n\t":
            v8 += (^(v4))
        default:
            if true {
                v10 += (v10 + (v10 << ((18446744073709551615) & 63)))
            } else {
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
            }
        }
    }
    v40 := GenS2{F0: 9223372036854775807, F1: "\n\t"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
