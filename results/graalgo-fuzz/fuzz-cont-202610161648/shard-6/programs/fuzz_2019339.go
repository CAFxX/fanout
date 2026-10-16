package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 int64; F2 int8 }
type GenS6 struct { F0 int32; F1 int8 }

func main() {
    v1 := []int{42, 9223372036854775807, 4294967295, -351}
    var v2 int64 = -1
    v3 := make(map[uint32]bool)
    if true {
        var v4 uint = 236
        _ = v4
    } else {
        switch "a" {
        case "a":
            for fi5 := 0; fi5 < 5; fi5++ {
                v2 = (^(v2))
                _ = fi5
            }
        case "foo bar":
            fmt.Println(v2)
        default:
            v2 = (v2 >> ((18446744073709551615) & 63))
        }
    }
    fmt.Println(v2)
    if (276 == 964) {
        _ = len(v1)
    } else {
        if (1 <= 2147483647) {
            switch 42 {
            case -9223372036854775808:
                v6 := GenS1{F0: 753, F1: -259, F2: 255}
                _ = v6
            case 9223372036854775807:
                switch 42 {
                case -2147483648:
                    fmt.Println(v2)
                case -964:
                    if false {
                        v1 = append(v1, 32767)
                    }
                }
            case 32768:
                switch ("" + "\n\t") {
                case "\n\t":
                    switch ("\n\t" + "a") {
                    case "\n\t":
                        v7 := GenS2{F0: 65535}
                        _ = v7
                    case "hello":
                        v8 := GenS3{F0: 127, F1: 1}
                        _ = v8
                    case "a":
                        var v9 uint32 = 767
                        _ = v9
                    }
                case "a":
                    switch -897 {
                    case -1:
                        v10 := v3[147]
                        _ = v10
                    case 32767:
                        v11 := GenS4{F0: "a"}
                        _ = v11
                    default:
                        for fi12 := 0; fi12 < 3; fi12++ {
                            for fi13 := 0; fi13 < 12; fi13++ {
                                v2 = (^(v2))
                                var v14 int32 = 0
                                v14 = v14
                                _ = fi13
                                _ = v14
                            }
                            _ = fi12
                        }
                    }
                case "foo bar":
                    if (!(0 > 127)) {
                        v2 -= 1
                        fmt.Println(v2)
                    }
                default:
                    v15 := GenS5{F0: "\n\t", F1: (v2 + v2), F2: 127}
                    _ = v15
                }
            }
            fmt.Println(v2)
            v2 = v2
        }
    }
    var v16 string = ("hello" + "a")
    v3[2147483648] = false
    var v17 uint = 18446744073709551615
    var v18 string = "foo bar"
    v18 = "\n\t"
    var v19 uint32 = 0
    switch 42 {
    case 762:
        if (1 > 0) {
            for fi20 := 0; fi20 < 13; fi20++ {
                for fi21 := 0; fi21 < 16; fi21++ {
                    if true {
                        if (22 == 0) {
                            fmt.Println(v18)
                            var v22 float64 = 1e10
                            var v23 int16 = -1
                            if (!(228 < v23)) {
                                if len(v1) > 0 {
                                    v1[0] = -948
                                }
                            }
                            v24 := GenS3{F0: (fi20 * fi20), F1: 643}
                            if (-68 <= -1) {
                                if false {
                                    if (!(!false)) {
                                        v16 = "foo bar"
                                        v25 := GenS6{F0: 0, F1: 15}
                                        fmt.Println(v17)
                                        _ = v25
                                    } else {
                                        _ = len(v1)
                                        v17 += 482
                                        v22 = (v22 * v22)
                                    }
                                    v18 = (v16 + v18)
                                }
                                var v26 bool = true
                                _ = v26
                            }
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            v19 = 0
                        }
                        fmt.Println(fi21)
                        if true {
                            fmt.Println(v16)
                        } else {
                            fmt.Println(fi21)
                        }
                    }
                    fmt.Println(v19, v16)
                    _ = fi21
                }
                _ = fi20
            }
            if (4294967295 >= 18446744073709551615) {
                switch -680 {
                case 1:
                    var v27 int32 = 42
                    _ = v27
                }
            }
        }
    case 42:
        v28 := GenS2{F0: v17}
        _ = v28
    }
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    _ = len(v3)
    v29 := GenS5{F0: v18, F1: v2, F2: 0}
    if len(v1) > 0 {
        v1[0] = 367
    }
    v18 = v16
    if true {
        v17 = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
}
