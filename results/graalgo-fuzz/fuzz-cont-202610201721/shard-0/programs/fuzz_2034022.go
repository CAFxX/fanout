package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0(p0 int64) int32 {
    switch -205 {
    case -9223372036854775808:
        var v6 uint32 = 431
        _ = v6
    case 883:
        switch 2147483647 {
        case 128:
            v7 := GenS1{F0: 2, F1: (!(float64(1e10) != 1.5))}
            _ = v7
        default:
            v8 := []int{-9223372036854775808}
            _ = v8
        }
    }
    v9 := []int{128, 4294967295}
    _ = p0
    _ = v9
    return 2147483647
}
func genFn1(p0 uint8, p1 uint8) uint64 {
    v10 := []int{4294967295}
    v11 := GenS1{F0: 1, F1: (!((!((!(!true)) || false)) || true))}
    p1 = (p1 | 128)
    var v12 int16 = 1000
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 2
}

func main() {
    var v1 int64 = 1
    var v2 bool = (false && false)
    var v3 uint64 = 4294967295
    v4 := []string{""}
    var v5 uint8 = 184
    v3 += (v3 + v3)
    if v2 {
        for fi13 := 0; fi13 < 15; fi13++ {
            switch "a" {
            case "hello":
                v2 = (((1.5 + float64(1e10)) / float64(3.14159)) != (-2.25 - 1.5))
            case "a":
                if len(v4) > 0 {
                    v4[0] = ""
                }
            }
            switch fi13 {
            case 4294967295:
                v4 = append(v4, "")
            case 32767:
                if true {
                    v14 := GenS1{F0: 168, F1: (0.1 <= (1e10 / 3.14159))}
                    _ = v14
                } else {
                    v15 := GenS2{F0: 0, F1: (-(v1))}
                    var v16 int = (fi13 / ((-1) | 1))
                    _ = v15
                    _ = v16
                }
            case 42:
                if false {
                    switch "\n\t" {
                    case "a":
                        var v17 uint16 = uint16(v3)
                        _ = v17
                    case "\n\t":
                        if v2 {
                            for fi18 := 0; fi18 < 2; fi18++ {
                                fmt.Println(v5)
                                _ = len(v4)
                                var v19 int16 = -990
                                _ = fi18
                                _ = v19
                            }
                        }
                    }
                    var v20 int8 = 77
                    var v21 float32 = (3.14159 / ((1e10 + 3.14159) + 3.14159))
                    _ = v20
                    _ = v21
                } else {
                    var v22 int64 = -1
                    if v2 {
                        for fi23 := 0; fi23 < 3; fi23++ {
                            v22 += 9223372036854775807
                            _ = fi23
                        }
                        for fi24 := 0; fi24 < 14; fi24++ {
                            var v25 uint = 4294967295
                            _ = len(v4)
                            v22 += v1
                            _ = fi24
                            _ = v25
                        }
                        var v26 string = "foo bar"
                        _ = v26
                    }
                    _ = v22
                }
            }
            _ = fi13
        }
        var v27 int = 2147483647
        var v28 float64 = 3.14159
        _ = v27
        _ = v28
    }
    var v29 uint8 = 255
    fmt.Println(v1)
    for fi30 := 0; fi30 < 11; fi30++ {
        v2 = (1e10 == 1e10)
        var v31 int = -2147483648
        _ = fi30
        _ = v31
    }
    v32 := GenS2{F0: 2147483648, F1: v1}
    fmt.Println(v5, v5, v5)
    v29 = (v29 - (v5 - (^(v29))))
    switch ("\n\t" + (("\n\t" + "hello") + "hello")) {
    case "hello":
        var v33 uint = 42
        _ = v33
    case "\n\t":
        v4 = append(v4, "hello")
    case "":
        fmt.Println(v1, v3, v1)
    default:
        v1 += (v1 | v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
