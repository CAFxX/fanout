package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint64; F1 int16; F2 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 bool = true
    v3 := []uint32{0}
    v2 = v2
    v4 := GenS1{F0: 1, F1: (!false)}
    v2 = (995 < 369)
    v5 := v1[""]
    if (2 < 18446744073709551615) {
        var v6 string = "foo bar"
        var v7 int8 = 37
        v6 = "a"
        _ = v6
        _ = v7
    }
    var v8 int8 = -1
    v3 = append(v3, 42)
    switch 9223372036854775807 {
    case 752:
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 string = ("\n\t" + ("\n\t" + ((("a" + "\n\t") + "hello") + "")))
            v11 := GenS1{F0: v8, F1: (4294967295 < 2)}
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        if (v2 && (1.5 > (3.14159 + ((1.5 + 3.14159) - 1.5)))) {
            var v12 uint = 490
            _ = v12
        } else {
            v13 := GenS1{F0: (-(v8)), F1: v2}
            v8 += (^(v8))
            v14 := v1[("a" + "foo bar")]
            _ = v13
            _ = v14
        }
    }
    switch "a" {
    case "":
        switch "" {
        case "a":
            switch -712 {
            case 32767:
                if (("\n\t" < "a") && (-919 > 970)) {
                    if (v2 && (v2 || (4294967295 != 18446744073709551615))) {
                        fmt.Println(v8)
                    } else {
                        switch (v5 / ((-843) | 1)) {
                        case 4294967295:
                            v5 = 128
                        case -407:
                            v5 += v5
                        case 127:
                            switch "hello" {
                            case "hello":
                                v15 := GenS1{F0: v8, F1: ((-1 != -19) || v2)}
                                _ = v15
                            }
                        default:
                            v3 = append(v3, 4294967295)
                        }
                    }
                    switch "a" {
                    case "":
                        _ = len(v3)
                    case "\n\t":
                        if (!(18446744073709551615 <= 941)) {
                            fmt.Println(v2, v5)
                        }
                    case "hello":
                        v16 := GenS1{F0: 1, F1: (false || (v2 && (v2 && v2)))}
                        _ = v16
                    default:
                        fmt.Println(v5)
                    }
                }
            }
        }
    case "hello":
        v5 += (v5 | (v5 / ((v5) | 1)))
    case "\n\t":
        v17 := GenS2{F0: 543, F1: "hello"}
        _ = v17
    default:
        var v18 int16 = 843
        _ = v18
    }
    v5 = v5
    if (!v2) {
        if (42 == 2147483647) {
            v5 -= (v5 % ((v5) | 1))
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 uint16 = 299
                _ = fi19
                _ = v20
            }
        } else {
            if false {
                if ((-2.25 != (-2.25 + (1.5 - float64(-2.25)))) && v2) {
                    var v21 int64 = -240
                    v5 += v5
                    _ = v21
                }
            }
            switch -294 {
            case 128:
                var v22 int64 = -411
                _ = v22
            case 127:
                v5 += v5
            case -68:
                fmt.Println(v2)
            }
            v5 = 128
        }
        v23 := GenS3{F0: 4294967295, F1: 42, F2: v8}
        v24 := GenS3{F0: 0, F1: -167, F2: v8}
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        for fi26 := 0; fi26 < 3; fi26++ {
            v2 = (!(126 >= 147))
            v27 := GenS3{F0: 772, F1: 1, F2: v8}
            _ = fi26
            _ = v27
        }
        if (4294967295 == 42) {
            for fi28 := 0; fi28 < 10; fi28++ {
                fmt.Println(fi28, v8, fi25)
                var v29 int32 = 2147483647
                if (1 != 0) {
                    v1["foo bar"] = (fi28 ^ (fi25 - v5))
                } else {
                    delete(v1, ("foo bar" + ("a" + "hello")))
                }
                _ = fi28
                _ = v29
            }
        } else {
            fmt.Println(v2)
            v8 = v8
        }
        _ = fi25
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        delete(v1, ("foo bar" + "\n\t"))
        _ = fi30
    }
    var v31 int64 = -429
    var v32 uint = 18446744073709551615
    v33 := GenS2{F0: 829, F1: "hello"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v8
}
