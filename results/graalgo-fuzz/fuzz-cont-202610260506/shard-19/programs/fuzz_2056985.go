package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 string }

func main() {
    var v1 string = "foo bar"
    var v2 int8 = 42
    var v3 int32 = 2147483647
    v4 := []string{v1, v1, "", "foo bar"}
    var v5 string = (v1 + v1)
    v6 := []string{}
    if false {
        for fi7 := 0; fi7 < 3; fi7++ {
            v2 = (v2 + (v2 / ((-128) | 1)))
            _ = fi7
        }
        switch 703 {
        case 1:
            if (415 < 259) {
                v5 = (v1 + v1)
                fmt.Println(v5)
            }
        case -9223372036854775808:
            if false {
                var v8 int8 = 42
                switch 2147483647 {
                case 128:
                    v9 := GenS1{F0: 260, F1: 18446744073709551615}
                    _ = v9
                default:
                    v10 := GenS1{F0: 0, F1: 321}
                    _ = v10
                }
                _ = v8
            } else {
                for fi11 := 0; fi11 < 2; fi11++ {
                    fmt.Println(fi11)
                    if (3.14159 > float64(float32(-2.25))) {
                        fmt.Println(v5)
                        v1 = v1
                    }
                    var v12 int8 = (^(v2))
                    _ = fi11
                    _ = v12
                }
                for fi13 := 0; fi13 < 6; fi13++ {
                    if (false || (848 <= 18446744073709551615)) {
                        switch v5 {
                        case "":
                            _ = len(v6)
                        case "foo bar":
                            v6 = append(v6, v1)
                        default:
                            var v14 uint8 = 255
                            _ = v14
                        }
                        var v15 bool = false
                        v2 -= v2
                        _ = v15
                    }
                    _ = fi13
                }
            }
        }
    } else {
        var v16 float32 = 1e10
        v2 -= v2
        _ = v16
    }
    var v17 int8 = 127
    var v18 float64 = -2.25
    v18 -= v18
    var v19 uint = 65535
    v20 := GenS2{F0: 127, F1: (v3 / (((v3 % (((v3 & -1)) | 1))) | 1))}
    switch v5 {
    case "foo bar":
        for fi21 := 0; fi21 < 4; fi21++ {
            v22 := GenS2{F0: v17, F1: v3}
            fmt.Println(fi21, v5, v19)
            _ = fi21
            _ = v22
        }
    case "a":
        v2 = v17
    }
    fmt.Println(v17, v2)
    v5 = "foo bar"
    fmt.Println(v19)
    v23 := GenS2{F0: v2, F1: v3}
    for fi24 := 0; fi24 < 10; fi24++ {
        v25 := GenS3{F0: v1}
        _ = fi24
        _ = v25
    }
    fmt.Println(v3, v5, v5)
    v18 = 1e10
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 uint32 = 65535
        for fi28 := 0; fi28 < 15; fi28++ {
            var v29 int = fi28
            v29 = v29
            var v30 uint8 = 127
            _ = fi28
            _ = v29
            _ = v30
        }
        switch "hello" {
        case "foo bar":
            var v31 int32 = (-(v3))
            _ = v31
        case "\n\t":
            v18 = (float64((0.0 * float32(v18))) - v18)
        default:
            fmt.Println(v17)
        }
        _ = fi26
        _ = v27
    }
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 uint16 = 780
        v2 = 14
        var v34 uint64 = 1
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
