package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int8; F2 string }
type GenS3 struct { F0 int32; F1 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 2147483648
    v3 := make(map[int]int)
    var v4 float32 = ((0.1 - (float32((1.5 - 1e10)) - 3.14159)) / 3.14159)
    var v5 int8 = 0
    if (65535 < 12) {
        v6 := GenS1{F0: 1}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 17; fi7++ {
            fmt.Println(v2, v4, v4)
            v8 := GenS1{F0: 32767}
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 5; fi9++ {
            v1[2] = (!true)
            var v10 float64 = (-2.25 - 0.0)
            if true {
                if (!false) {
                    fmt.Println(v10, v10)
                    var v11 bool = (400 == -946)
                    _ = v11
                }
            }
            _ = fi9
            _ = v10
        }
        fmt.Println(v4, v5, v2)
    }
    var v12 bool = true
    if v12 {
        fmt.Println(v4, v2)
        switch 647 {
        case 9223372036854775807:
            var v13 int32 = 0
            _ = v13
        case 599:
            if false {
                var v14 int64 = 9223372036854775807
                switch "hello" {
                case "hello":
                    var v15 uint32 = 1
                    _ = v15
                case "\n\t":
                    v5 = v5
                case "":
                    for fi16 := 0; fi16 < 10; fi16++ {
                        v17 := GenS1{F0: 42}
                        if (-2147483648 <= 451) {
                            var v18 int64 = v14
                            _ = v18
                        }
                        _ = fi16
                        _ = v17
                    }
                }
                v2 = (v2 % ((42) | 1))
                _ = v14
            } else {
                switch -1 {
                case -1:
                    _ = len(v3)
                default:
                    v5 = v5
                }
                v19 := GenS2{F0: 1, F1: 0, F2: "\n\t"}
                fmt.Println(v5, v2, v5)
                _ = v19
            }
        }
        if v12 {
            if v12 {
                var v20 uint64 = 578
                v4 -= 0.1
                v21 := GenS3{F0: 1, F1: 178}
                _ = v20
                _ = v21
            } else {
                if ((-384 >= -1) && v12) {
                    v5 -= (^(v5))
                    _ = len(v1)
                } else {
                    if v12 {
                        var v22 float64 = 0.0
                        _ = v22
                    }
                }
                _ = len(v3)
            }
        } else {
            switch ("" + "hello") {
            case "\n\t":
                fmt.Println(v5)
            case "hello":
                var v23 uint64 = 400
                _ = v23
            case "a":
                fmt.Println(v5, v12)
            default:
                var v24 int8 = (v5 | (v5 << ((255) & 63)))
                _ = v24
            }
        }
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 uint = 0
        _ = fi25
        _ = v26
    }
    var v27 int32 = 501
    v2 = v2
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v28 int32 = v27
        _ = v28
    case -396:
        var v29 float64 = -2.25
        _ = v29
    default:
        var v30 uint = 2
        _ = v30
    }
    if v12 {
        var v31 float64 = 3.14159
        _ = v31
    }
    var v32 float64 = 3.14159
    for fi33 := 0; fi33 < 4; fi33++ {
        v34 := GenS3{F0: (v27 << ((1) & 63)), F1: 1}
        _ = fi33
        _ = v34
    }
    v32 += (float64(v4) + 1e10)
    fmt.Println(v12)
    v5 -= 0
    var v35 bool = (128 == -2147483648)
    var v36 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
