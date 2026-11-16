package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 string }

func main() {
    v1 := []string{}
    var v2 string = "\n\t"
    var v3 int32 = 42
    v4 := []int{9223372036854775807, -590, 127, 127}
    v2 = ""
    var v5 int = -1
    fmt.Println(v5)
    var v6 uint = 4294967295
    v4 = append(v4, v5)
    v5 -= v5
    fmt.Println(v5, v2, v2)
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            if false {
                v9 := GenS1{F0: v3, F1: 42}
                if ((v3 * v3) < v3) {
                    switch v2 {
                    case "":
                        v10 := GenS2{F0: (true || ((v2 + v2) == (v2 + ""))), F1: 2}
                        v4 = append(v4, fi7)
                        fmt.Println(fi8)
                        _ = v10
                    case "foo bar":
                        v11 := GenS2{F0: false, F1: 32767}
                        _ = v11
                    case "hello":
                        fmt.Println(v5, v6, v2)
                    default:
                        var v12 int16 = 1
                        _ = v12
                    }
                } else {
                    var v13 uint64 = 0
                    v4 = append(v4, -2147483648)
                    if false {
                        v6 = v6
                        v2 = v2
                    } else {
                        v14 := GenS3{F0: v2}
                        if true {
                            if ((3.14159 == 0.0) || (42 <= 320)) {
                                if (128 <= int(v6)) {
                                    v2 = "hello"
                                    if len(v1) > 0 {
                                        v1[0] = (v2 + v2)
                                    }
                                }
                            }
                        } else {
                            v6 -= 255
                        }
                        switch fi8 {
                        case -2147483648:
                            var v15 bool = (!(-1 < 0))
                            _ = v15
                        case 2147483647:
                            v5 = (fi7 - (v5 | (^(fi8))))
                        case -413:
                            var v16 float64 = (3.14159 * ((0.1 * 1.5) / -2.25))
                            _ = v16
                        }
                        _ = v14
                    }
                    _ = v13
                }
                v2 = v2
                _ = v9
            } else {
                var v17 float32 = 1.5
                _ = v17
            }
            _ = fi8
        }
        var v18 string = v2
        var v19 uint8 = 255
        _ = fi7
        _ = v18
        _ = v19
    }
    v1 = append(v1, v2)
    v1 = append(v1, v2)
    fmt.Println(v5, v3, v5)
    if (4294967295 == 2147483647) {
        for fi20 := 0; fi20 < 1; fi20++ {
            fmt.Println(fi20)
            _ = fi20
        }
    }
    switch v2 {
    case "foo bar":
        v2 = "a"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
