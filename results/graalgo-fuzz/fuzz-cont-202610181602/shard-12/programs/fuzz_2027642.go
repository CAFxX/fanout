package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint; F2 int16 }
type GenS4 struct { F0 string; F1 int8; F2 uint32 }
func genFn0(p0 int16) uint8 {
    var v4 bool = false
    v4 = (!(255 <= 42))
    _ = p0
    _ = v4
    return 2
}

func main() {
    v1 := []uint32{}
    var v2 float64 = 3.14159
    var v3 int8 = 127
    var v5 uint32 = 1
    var v6 float64 = v2
    fmt.Println(v2, v2)
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint = 0
        var v9 uint8 = 1
        var v10 int64 = 1
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 bool = false
    if v11 {
        if v11 {
            v6 += (v2 / (float64(-2.25) / (float64(0.1) / 3.14159)))
            var v12 float64 = float64(-2.25)
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 int64 = -9223372036854775808
                v15 := GenS1{F0: -680, F1: (fi13 << ((0) & 63)), F2: "\n\t"}
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        } else {
            switch (("" + (("a" + (("foo bar" + "\n\t") + "a")) + "hello")) + "hello") {
            case "a":
                v16 := GenS1{F0: -2147483648, F1: 977, F2: ("\n\t" + ((("" + ("\n\t" + "foo bar")) + "foo bar") + ("foo bar" + "")))}
                _ = v16
            case "":
                if v11 {
                    var v17 float32 = float32(v6)
                    _ = v17
                } else {
                    v18 := GenS2{F0: 184}
                    for fi19 := 0; fi19 < 15; fi19++ {
                        var v20 float64 = -2.25
                        v1 = append(v1, v5)
                        for fi21 := 0; fi21 < 11; fi21++ {
                            var v22 float32 = 1.5
                            _ = fi21
                            _ = v22
                        }
                        _ = fi19
                        _ = v20
                    }
                    _ = v18
                }
            case "hello":
                v1 = append(v1, v5)
            default:
                var v23 float32 = 3.14159
                _ = v23
            }
        }
        fmt.Println(v3, v11, v2)
        v24 := GenS3{F0: -825, F1: 42, F2: -947}
        _ = v24
    } else {
        for fi25 := 0; fi25 < 4; fi25++ {
            switch ("a" + "foo bar") {
            case "hello":
                v11 = (!v11)
            case "":
                for fi26 := 0; fi26 < 15; fi26++ {
                    switch -1 {
                    case 127:
                        fmt.Println(fi26, v5, v11)
                    default:
                        switch (fi26 >> ((18446744073709551615) & 63)) {
                        case 128:
                            v27 := GenS4{F0: "", F1: v3, F2: (v5 & (^(v5)))}
                            _ = v27
                        case 9223372036854775807:
                            var v28 uint8 = 127
                            _ = v28
                        case -996:
                            if false {
                                if len(v1) > 0 {
                                    v1[0] = (^(v5))
                                }
                                var v29 int16 = 761
                                _ = v29
                            }
                        }
                    }
                    _ = fi26
                }
            case "\n\t":
                v3 -= 4
            }
            _ = fi25
        }
        v6 = v6
    }
    v1 = append(v1, v5)
    v30 := GenS4{F0: "", F1: v3, F2: v5}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v5
    _ = v6
}
