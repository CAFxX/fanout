package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }

func main() {
    var v1 uint16 = 0
    v2 := []int{578, -1}
    var v3 uint32 = 2147483647
    v1 += 1
    v1 += 42
    v1 = (v1 ^ 933)
    if (!(false && true)) {
        var v4 uint32 = (v3 | v3)
        _ = v4
    } else {
        fmt.Println(v1, v3)
        var v5 uint16 = uint16(v3)
        for fi6 := 0; fi6 < 2; fi6++ {
            for fi7 := 0; fi7 < 3; fi7++ {
                switch fi7 {
                case 4294967295:
                    v3 = 4294967295
                default:
                    var v8 int64 = -346
                    _ = v8
                }
                var v9 bool = true
                v1 += v5
                _ = fi7
                _ = v9
            }
            _ = fi6
        }
        _ = v5
    }
    v3 += (v3 | v3)
    switch "foo bar" {
    case "":
        var v10 float32 = 1e10
        _ = v10
    case "hello":
        v11 := GenS1{F0: -2147483648, F1: (717 >= 1)}
        _ = v11
    case "\n\t":
        v12 := GenS1{F0: -1, F1: false}
        _ = v12
    default:
        switch -282 {
        case 32767:
            fmt.Println(v1, v3)
        case -2147483648:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 1:
            fmt.Println(v3, v3)
        }
    }
    var v13 int16 = 1000
    var v14 int32 = 2147483647
    v2 = append(v2, -1)
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 int = 32768
        fmt.Println(v13)
        _ = fi15
        _ = v16
    }
    if false {
        _ = len(v2)
        v13 = (v13 % (((v13 & v13)) | 1))
        switch ("hello" + "hello") {
        case "foo bar":
            var v17 uint = 18446744073709551615
            _ = v17
        case "\n\t":
            if true {
                v18 := GenS1{F0: v14, F1: false}
                _ = v18
            } else {
                var v19 float64 = float64(-2.25)
                for fi20 := 0; fi20 < 17; fi20++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v21 int = 1
                    var v22 bool = (!true)
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                var v23 int8 = -6
                _ = v19
                _ = v23
            }
        }
    } else {
        v24 := GenS1{F0: v14, F1: false}
        v25 := GenS1{F0: (^(v14)), F1: (!true)}
        v1 -= (v1 | 42)
        _ = v24
        _ = v25
    }
    var v26 string = ("a" + "\n\t")
    v14 -= int32(v1)
    if ((v14 * (v14 + v14)) > v14) {
        var v27 int16 = (-(v13))
        v13 -= v13
        var v28 uint = 1
        _ = v27
        _ = v28
    }
    switch v26 {
    case "":
        var v29 int64 = 767
        _ = v29
    case "a":
        v1 = (^(v1))
    case "\n\t":
        fmt.Println(v26, v26, v14)
    default:
        var v30 uint8 = 2
        _ = v30
    }
    var v31 int = -1
    if false {
        v3 = v3
    } else {
        fmt.Println(v31)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v31
}
