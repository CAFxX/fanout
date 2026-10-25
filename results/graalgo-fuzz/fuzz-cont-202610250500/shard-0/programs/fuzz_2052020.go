package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int64; F1 bool; F2 uint64 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 int; F1 uint64 }

func main() {
    var v1 int32 = -1
    var v2 int16 = -32768
    var v3 int16 = (v2 >> ((1) & 63))
    v4 := []int{10}
    var v5 uint64 = 4294967295
    v6 := make(map[string]int)
    v2 -= (v3 * (v2 * -32768))
    _ = len(v4)
    v7 := GenS1{F0: 903}
    switch 9223372036854775807 {
    case 802:
        v1 = 2147483647
    case 9223372036854775807:
        var v8 uint = 255
        _ = v8
    }
    switch 127 {
    case 4294967295:
        fmt.Println(v1)
    case 127:
        switch "hello" {
        case "":
            var v9 float64 = 1.5
            _ = v9
        case "hello":
            v10 := GenS2{F0: 0, F1: 744, F2: 91}
            _ = v10
        default:
            if false {
                v1 = (v1 / ((1) | 1))
                var v11 string = ""
                for fi12 := 0; fi12 < 17; fi12++ {
                    if (2147483647 > (fi12 / ((fi12) | 1))) {
                        v3 += v3
                        var v13 uint8 = 127
                        v14 := GenS1{F0: 42}
                        _ = v13
                        _ = v14
                    } else {
                        v4 = append(v4, -320)
                    }
                    var v15 bool = (65535 > 32767)
                    switch "hello" {
                    case "":
                        v16 := GenS2{F0: 255, F1: 32767, F2: 255}
                        _ = v16
                    case "\n\t":
                        v17 := GenS3{F0: uint32(v5), F1: 2, F2: 2}
                        _ = v17
                    }
                    _ = fi12
                    _ = v15
                }
                _ = v11
            }
        }
    case 927:
        var v18 int32 = v1
        _ = v18
    }
    var v19 bool = (!false)
    var v20 int8 = 1
    switch "a" {
    case "\n\t":
        if (!v19) {
            v21 := GenS3{F0: 2, F1: 42, F2: 42}
            var v22 uint64 = v5
            v5 -= v22
            _ = v21
            _ = v22
        } else {
            var v23 uint16 = 32767
            fmt.Println(v3, v5)
            _ = v23
        }
    case "foo bar":
        switch "" {
        case "\n\t":
            for fi24 := 0; fi24 < 19; fi24++ {
                if v19 {
                    var v25 float64 = -2.25
                    v26 := GenS4{F0: -837, F1: (65535 > 18446744073709551615), F2: 4294967295}
                    _ = v25
                    _ = v26
                } else {
                    v2 = (v3 | (v2 * v2))
                }
                v5 += (v5 + 42)
                _ = fi24
            }
        case "a":
            v27 := GenS3{F0: 195, F1: 1, F2: 857}
            _ = v27
        case "foo bar":
            for fi28 := 0; fi28 < 11; fi28++ {
                var v29 uint8 = 127
                switch "a" {
                case "\n\t":
                    fmt.Println(v2, v5)
                default:
                    var v30 int = fi28
                    _ = v30
                }
                _ = fi28
                _ = v29
            }
        }
    case "a":
        v31 := GenS5{F0: (v20 + v20), F1: 42}
        _ = v31
    }
    fmt.Println(v19, v19, v1)
    if (v19 || (-2.25 > 0.0)) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v32 uint64 = (v5 << ((171) & 63))
    var v33 int16 = v2
    v34 := GenS6{F0: 9223372036854775807, F1: 18446744073709551615}
    for fi35 := 0; fi35 < 2; fi35++ {
        v6["hello"] = fi35
        switch ("" + ("foo bar" + ("foo bar" + ""))) {
        case "foo bar":
            v36 := GenS1{F0: 2}
            _ = v36
        default:
            if false {
                v1 -= -2147483648
                fmt.Println(v1, v32)
                var v37 uint8 = 14
                _ = v37
            } else {
                fmt.Println(fi35, fi35)
                for fi38 := 0; fi38 < 12; fi38++ {
                    v4 = append(v4, (fi38 + int(v3)))
                    var v39 int16 = v2
                    _ = fi38
                    _ = v39
                }
                v40 := GenS2{F0: 255, F1: fi35, F2: 1}
                _ = v40
            }
        }
        v5 += v32
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
