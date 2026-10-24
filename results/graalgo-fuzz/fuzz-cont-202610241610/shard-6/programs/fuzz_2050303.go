package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint32 = 42
    var v3 int64 = 42
    v4 := GenS1{F0: 32767, F1: 2147483647}
    fmt.Println(v1)
    var v5 uint8 = 123
    fmt.Println(v2, v1)
    v6 := []int{9223372036854775807}
    v2 -= (v2 / ((890) | 1))
    if true {
        switch 1 {
        case 1:
            v7 := GenS2{F0: -493, F1: true}
            _ = v7
        case -2147483648:
            v1 = (v1 % ((uint64(v2)) | 1))
        default:
            var v8 int64 = v3
            _ = v8
        }
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint8 = 0
            v11 := GenS3{F0: (fi9 + (fi9 % (((fi9 / (((fi9 | fi9)) | 1))) | 1)))}
            _ = fi9
            _ = v10
            _ = v11
        }
        switch "\n\t" {
        case "hello":
            var v12 uint = uint(v2)
            _ = v12
        case "foo bar":
            _ = len(v6)
        case "a":
            if false {
                v13 := GenS3{F0: 32767}
                if true {
                    if true {
                        var v14 int32 = -399
                        switch "\n\t" {
                        case "\n\t":
                            switch ((("hello" + "a") + "hello") + ("" + "")) {
                            case "hello":
                                var v15 float64 = -2.25
                                _ = v15
                            case "":
                                v6 = append(v6, -9223372036854775808)
                            default:
                                v2 = (^(v2))
                            }
                        case "hello":
                            var v16 int32 = v14
                            _ = v16
                        default:
                            var v17 int16 = 0
                            _ = v17
                        }
                        _ = v14
                    }
                    v18 := GenS1{F0: 1000, F1: -546}
                    _ = v18
                } else {
                    switch "hello" {
                    case "foo bar":
                        v6 = append(v6, 9223372036854775807)
                    case "hello":
                        fmt.Println(v2, v5)
                    }
                    for fi19 := 0; fi19 < 19; fi19++ {
                        v20 := GenS1{F0: -215, F1: 347}
                        _ = fi19
                        _ = v20
                    }
                }
                var v21 float64 = -2.25
                _ = v13
                _ = v21
            } else {
                var v22 uint = 255
                _ = v22
            }
        default:
            switch ("hello" + (("\n\t" + ("a" + "foo bar")) + "hello")) {
            case "a":
                fmt.Println(v5, v3, v5)
            case "\n\t":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            default:
                if false {
                    var v23 uint = 1
                    var v24 bool = false
                    _ = len(v6)
                    _ = v23
                    _ = v24
                } else {
                    var v25 int = 513
                    _ = v25
                }
            }
        }
    } else {
        if (false || (0.1 < 1.5)) {
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 uint8 = v5
                for fi28 := 0; fi28 < 3; fi28++ {
                    v2 = (v2 % (((v2 & (v2 % ((v2) | 1)))) | 1))
                    _ = fi28
                }
                _ = fi26
                _ = v27
            }
        } else {
            fmt.Println(v2)
        }
        v6 = append(v6, -9223372036854775808)
        v29 := GenS2{F0: 601, F1: false}
        _ = v29
    }
    var v30 int = 4294967295
    var v31 int32 = -873
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi32 := 0; fi32 < 11; fi32++ {
        if len(v6) > 0 {
            v6[0] = 897
        }
        var v33 bool = (((v3 / ((v3) | 1)) >= v3) && (fi32 >= v30))
        v30 = (v30 / ((fi32) | 1))
        _ = fi32
        _ = v33
    }
    var v34 int32 = 2147483647
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
