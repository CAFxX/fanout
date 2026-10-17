package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint32 }

func main() {
    var v1 float64 = 0.1
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    fmt.Println(v1)
    var v4 uint32 = 436
    var v5 string = "a"
    var v6 uint8 = 2
    var v7 int64 = 1
    var v8 float32 = (1.5 / (3.14159 + 1.5))
    if false {
        v9 := GenS1{F0: -1}
        switch "\n\t" {
        case "a":
            var v10 int16 = 1000
            _ = v10
        case "\n\t":
            v11 := GenS2{F0: 1, F1: 32767}
            _ = v11
        default:
            if true {
                switch "" {
                case "a":
                    switch ("" + v5) {
                    case "":
                        for fi12 := 0; fi12 < 10; fi12++ {
                            fmt.Println(v8, v7)
                            _ = fi12
                        }
                    default:
                        switch v5 {
                        case "\n\t":
                            v8 = v8
                        case "foo bar":
                            switch -1 {
                            case 32767:
                                if (32767 > -792) {
                                    var v13 uint16 = 1
                                    for fi14 := 0; fi14 < 10; fi14++ {
                                        var v15 string = "foo bar"
                                        v16 := GenS2{F0: -157, F1: int(v4)}
                                        _ = fi14
                                        _ = v15
                                        _ = v16
                                    }
                                    _ = v13
                                } else {
                                    var v17 float32 = v8
                                    if false {
                                        var v18 int64 = -575
                                        v19 := GenS1{F0: -268}
                                        var v20 uint16 = 2
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    }
                                    v4 = (v4 ^ (v4 ^ 2))
                                    _ = v17
                                }
                            }
                        case "a":
                            switch "hello" {
                            case "\n\t":
                                fmt.Println(v6)
                            case "a":
                                switch v5 {
                                case "\n\t":
                                    var v21 float32 = float32(v1)
                                    _ = v21
                                }
                            default:
                                var v22 bool = false
                                _ = v22
                            }
                        default:
                            v1 += ((v1 * 0.1) / v1)
                        }
                    }
                case "\n\t":
                    v23 := GenS3{F0: 534}
                    _ = v23
                case "hello":
                    switch v5 {
                    case "\n\t":
                        var v24 int8 = 104
                        _ = v24
                    case "":
                        var v25 int32 = 0
                        _ = v25
                    default:
                        for fi26 := 0; fi26 < 20; fi26++ {
                            fmt.Println(v6, v7)
                            _ = fi26
                        }
                    }
                default:
                    if (((-2.25 + v1) * 1.5) > (-2.25 + (v1 + (((v1 / v1) - v1) + v1)))) {
                        var v27 int64 = v7
                        _ = v27
                    } else {
                        if true {
                            if true {
                                if (v7 > (v7 << ((42) & 63))) {
                                    v28 := []int{}
                                    switch 9223372036854775807 {
                                    case -9223372036854775808:
                                        fmt.Println(v7)
                                    case 9223372036854775807:
                                        fmt.Println(v8)
                                    }
                                    switch 246 {
                                    case -215:
                                        v29 := v3[(v5 + v5)]
                                        _ = v29
                                    case -2147483648:
                                        switch "\n\t" {
                                        case "a":
                                            if len(v28) > 0 {
                                                v28[0] = 32767
                                            }
                                        case "\n\t":
                                            var v30 int8 = 1
                                            _ = v30
                                        default:
                                            v7 += 9223372036854775807
                                        }
                                    }
                                    _ = v28
                                }
                                var v31 int8 = 127
                                _ = v31
                            }
                            if false {
                                v32 := GenS4{F0: (2147483647 != 1), F1: 4294967295}
                                v1 = 3.14159
                                _ = v32
                            } else {
                                var v33 uint16 = 1
                                var v34 uint32 = v4
                                _ = v33
                                _ = v34
                            }
                        } else {
                            var v35 bool = (((!false) && (v6 < v6)) && true)
                            v36 := []int{898, int(v6), int(v7)}
                            _ = v35
                            _ = v36
                        }
                        delete(v3, v5)
                        v6 = v6
                    }
                }
            }
        }
        _ = v9
    }
    if true {
        var v37 int64 = (-(v7))
        if (false && (65535 <= 261)) {
            var v38 uint32 = v4
            _ = v38
        }
        for fi39 := 0; fi39 < 5; fi39++ {
            v40 := []int{}
            fmt.Println(v4)
            fmt.Println(v4)
            _ = fi39
            _ = v40
        }
        _ = v37
    } else {
        var v41 int16 = -558
        v2[(v4 << ((uint(v6)) & 63))] = ((v5 + "a") + v5)
        for fi42 := 0; fi42 < 14; fi42++ {
            fmt.Println(v4, v5)
            v7 += (v7 >> ((42) & 63))
            _ = fi42
        }
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
