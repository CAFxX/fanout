package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int16 }
type GenS2 struct { F0 int8; F1 int16; F2 uint8 }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -512
    var v3 string = "foo bar"
    var v4 uint8 = 1
    v5 := make(map[string]string)
    v3 = ("" + v3)
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 uint8 = 1
        v2 = (v2 / ((-32768) | 1))
        v2 = v2
        _ = fi6
        _ = v7
    }
    v8 := v5[v3]
    v4 = (v4 >> ((2) & 63))
    if false {
        switch 326 {
        case 915:
            v4 += v4
        case -9223372036854775808:
            if ((v8 + (v8 + ((v3 + v8) + "a"))) != "\n\t") {
                v3 = v3
                _ = len(v1)
            }
        case 32768:
            switch 127 {
            case 4294967295:
                _ = len(v1)
            default:
                if false {
                    delete(v1, 1)
                    v9 := v1[127]
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 15; fi10++ {
                        var v11 uint16 = 42
                        v12 := []int{fi10, 0, (fi10 ^ fi10), fi10}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 12; fi13++ {
                        for fi14 := 0; fi14 < 19; fi14++ {
                            switch "\n\t" {
                            case "":
                                fmt.Println(v4, fi14, v2)
                            case "a":
                                v8 = ""
                            case "foo bar":
                                if (!false) {
                                    if (v3 == v8) {
                                        if false {
                                            delete(v1, (fi14 >> ((18446744073709551615) & 63)))
                                        }
                                        v15 := GenS1{F0: v8, F1: v4, F2: v2}
                                        var v16 int8 = -100
                                        _ = v15
                                        _ = v16
                                    } else {
                                        v1[923] = fi13
                                        switch "a" {
                                        case "\n\t":
                                            _ = len(v5)
                                        case "a":
                                            v8 = "a"
                                        }
                                    }
                                    v17 := GenS1{F0: v3, F1: v4, F2: int16(fi13)}
                                    switch "a" {
                                    case "a":
                                        var v18 int16 = 1
                                        _ = v18
                                    case "":
                                        if (true && false) {
                                            v2 -= v2
                                            v19 := GenS1{F0: v8, F1: v4, F2: v2}
                                            _ = v19
                                        }
                                    default:
                                        v20 := GenS1{F0: "foo bar", F1: (v4 ^ v4), F2: v2}
                                        _ = v20
                                    }
                                    _ = v17
                                } else {
                                    fmt.Println(v8, v4)
                                }
                            }
                            v8 = v3
                            v21 := []int{(fi13 >> ((4294967295) & 63)), (fi14 | fi13), int(v2), (fi13 - (fi14 ^ (fi13 - (fi13 | -9223372036854775808))))}
                            _ = fi14
                            _ = v21
                        }
                        _ = fi13
                    }
                    v2 = 42
                }
            }
        default:
            var v22 uint = 141
            _ = v22
        }
        for fi23 := 0; fi23 < 8; fi23++ {
            for fi24 := 0; fi24 < 3; fi24++ {
                var v25 int8 = -1
                var v26 int64 = 9223372036854775807
                _ = fi24
                _ = v25
                _ = v26
            }
            switch -2147483648 {
            case -1:
                var v27 int64 = 1
                _ = v27
            case 9223372036854775807:
                v28 := GenS2{F0: 42, F1: v2, F2: 75}
                _ = v28
            case -9223372036854775808:
                if false {
                    v29 := GenS2{F0: -1, F1: (v2 - v2), F2: (^(v4))}
                    v30 := GenS2{F0: -128, F1: -936, F2: (v4 - (v4 & (^(v4))))}
                    _ = v29
                    _ = v30
                }
            }
            delete(v1, fi23)
            _ = fi23
        }
    }
    switch 2147483647 {
    case 9223372036854775807:
        var v31 uint8 = 2
        _ = v31
    case 973:
        fmt.Println(v3, v3)
    case -948:
        var v32 uint = 611
        _ = v32
    }
    switch ("" + "hello") {
    case "foo bar":
        fmt.Println(v4, v4, v4)
    case "":
        var v33 uint8 = 0
        _ = v33
    case "\n\t":
        var v34 uint = 255
        _ = v34
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        if false {
            var v36 int16 = 42
            var v37 int8 = 42
            _ = v36
            _ = v37
        } else {
            var v38 uint = 4294967295
            var v39 int16 = v2
            _ = v38
            _ = v39
        }
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
