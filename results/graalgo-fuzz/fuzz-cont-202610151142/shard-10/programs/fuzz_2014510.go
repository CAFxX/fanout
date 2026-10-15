package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint16 }

func main() {
    var v1 string = ""
    v2 := make(map[string]int)
    var v3 int8 = 127
    var v4 uint = 4294967295
    if false {
        v5 := GenS1{F0: 1}
        fmt.Println(v3, v1, v4)
        v6 := GenS1{F0: -2147483648}
        _ = v5
        _ = v6
    }
    var v7 int8 = 61
    switch 32768 {
    case 0:
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := v2["\n\t"]
            var v10 int64 = 1
            _ = fi8
            _ = v10
            _ = v9
        }
    case 2147483647:
        if true {
            v3 = v3
            var v11 bool = true
            _ = v11
        }
    case -848:
        var v12 float64 = 1e10
        _ = v12
    default:
        fmt.Println(v7, v1, v7)
    }
    var v13 float32 = 3.14159
    v14 := []int{1}
    for fi15 := 0; fi15 < 19; fi15++ {
        v13 = 0.0
        fmt.Println(v4, v4, v7)
        _ = fi15
    }
    v16 := GenS1{F0: -952}
    v7 = v7
    var v17 uint8 = 0
    v13 = (0.0 * 1e10)
    v3 -= v7
    v2[""] = 42
    if false {
        v18 := GenS1{F0: -2147483648}
        v19 := v2[(v1 + "a")]
        _ = v18
        _ = v19
    } else {
        switch 2147483647 {
        case -270:
            var v20 float64 = ((0.0 - 0.0) * ((3.14159 - 0.0) + 3.14159))
            _ = v20
        case -9223372036854775808:
            fmt.Println(v13)
        default:
            var v21 uint8 = (^(v17))
            _ = v21
        }
    }
    fmt.Println(v1, v17, v7)
    if false {
        var v22 float64 = -2.25
        _ = len(v14)
        _ = v22
    } else {
        var v23 bool = (v17 <= v17)
        if (32767 <= -274) {
            if (v23 && (("a" >= v1) && false)) {
                var v24 float64 = (0.1 + (float64(v13) - -2.25))
                v25 := GenS2{F0: 42}
                _ = v24
                _ = v25
            } else {
                v3 -= (-(v3))
                for fi26 := 0; fi26 < 17; fi26++ {
                    switch (fi26 + fi26) {
                    case 32768:
                        v27 := GenS1{F0: -1}
                        _ = v27
                    case 2147483647:
                        if (!true) {
                            if len(v14) > 0 {
                                v14[0] = fi26
                            }
                            switch v1 {
                            case "a":
                                v7 = v7
                            case "foo bar":
                                var v28 int16 = 580
                                _ = v28
                            case "":
                                v13 += (v13 - (1e10 + v13))
                            default:
                                v29 := GenS3{F0: fi26, F1: 130}
                                _ = v29
                            }
                        }
                    }
                    _ = fi26
                }
                v13 = v13
            }
            var v30 bool = false
            _ = v30
        }
        _ = v23
    }
    v13 -= v13
    v13 -= v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
