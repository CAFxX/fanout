package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }

func main() {
    v1 := make(map[int]string)
    var v2 int = -957
    v3 := []uint32{45, 0, 1}
    v4 := make(map[string]int)
    fmt.Println(v2)
    if ((v2 | (v2 >> ((4294967295) & 63))) == (v2 << ((1) & 63))) {
        v5 := GenS1{F0: 0, F1: -1}
        _ = v5
    } else {
        fmt.Println(v2)
        delete(v4, "")
        if true {
            v3 = append(v3, 42)
        } else {
            for fi6 := 0; fi6 < 12; fi6++ {
                for fi7 := 0; fi7 < 17; fi7++ {
                    v2 = (fi7 >> ((2) & 63))
                    v2 += fi7
                    _ = fi7
                }
                var v8 bool = (42 == 42)
                for fi9 := 0; fi9 < 16; fi9++ {
                    if v8 {
                        var v10 uint8 = 127
                        _ = v10
                    }
                    var v11 float64 = ((1.5 - ((-2.25 * 1e10) / -2.25)) - ((-2.25 - 1.5) / 3.14159))
                    fmt.Println(fi9)
                    _ = fi9
                    _ = v11
                }
                _ = fi6
                _ = v8
            }
            v2 += v2
            fmt.Println(v2)
        }
    }
    var v12 string = "a"
    fmt.Println(v2, v2)
    _ = len(v3)
    var v13 string = "hello"
    var v14 int8 = -76
    switch v2 {
    case -2147483648:
        switch "" {
        case "hello":
            v2 = (v2 % ((v2) | 1))
        case "foo bar":
            var v15 int16 = -1
            _ = v15
        default:
            fmt.Println(v13, v2)
        }
    case 4294967295:
        var v16 uint64 = 956
        _ = v16
    case 9223372036854775807:
        v13 = "foo bar"
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 float64 = (-2.25 + 0.1)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if ("\n\t" > "\n\t") {
            for fi19 := 0; fi19 < 18; fi19++ {
                if (!true) {
                    v12 = (((v13 + v12) + v13) + v12)
                    v14 = (v14 ^ -118)
                }
                if true {
                    if true {
                        v18 = v18
                        switch "foo bar" {
                        case "":
                            fmt.Println(fi17, v18)
                        case "hello":
                            v18 += 0.1
                        }
                        v3 = append(v3, 105)
                    }
                    switch (v12 + v13) {
                    case "foo bar":
                        var v20 float32 = 3.14159
                        _ = v20
                    default:
                        if true {
                            var v21 int = fi17
                            v12 = ""
                            var v22 int8 = (v14 % (((-(v14))) | 1))
                            _ = v21
                            _ = v22
                        } else {
                            if len(v3) > 0 {
                                v3[0] = 229
                            }
                            v23 := GenS1{F0: 2147483647, F1: 42}
                            _ = v23
                        }
                    }
                }
                _ = fi19
            }
            fmt.Println(v14)
        }
        _ = fi17
        _ = v18
    }
    var v24 bool = true
    var v25 int64 = -584
    if v24 {
        fmt.Println(v12, v13)
        var v26 float32 = -2.25
        var v27 uint = 91
        _ = v26
        _ = v27
    }
    if (239 == 1) {
        var v28 int16 = -1
        if (!((float32(1.5) * (float32(1.5) - (0.0 * (-2.25 - 3.14159)))) >= (-2.25 + 3.14159))) {
            if len(v3) > 0 {
                v3[0] = 602
            }
            var v29 uint32 = 0
            v3 = append(v3, (^(v29)))
            _ = v29
        }
        v3 = append(v3, 481)
        _ = v28
    }
    var v30 int32 = 404
    for fi31 := 0; fi31 < 1; fi31++ {
        for fi32 := 0; fi32 < 1; fi32++ {
            if true {
                fmt.Println(v24, v13)
                fmt.Println(v14)
            } else {
                v25 += 42
            }
            _ = fi32
        }
        _ = fi31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
}
