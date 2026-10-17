package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 string }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint64, p1 uint8) bool {
    var v7 uint8 = p1
    v8 := []int{}
    for fi9 := 0; fi9 < 10; fi9++ {
        p1 = uint8(p0)
        p1 -= 128
        _ = fi9
    }
    switch 2147483647 {
    case -954:
        p0 -= (^(p0))
    case 2147483647:
        for fi10 := 0; fi10 < 3; fi10++ {
            v7 -= (^(v7))
            v11 := GenS1{F0: 32767, F1: 1, F2: 65535}
            _ = fi10
            _ = v11
        }
    default:
        var v12 uint64 = (^(p0))
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (1 <= uint(p0))
}

func main() {
    v1 := []int{-2147483648}
    var v2 uint64 = 613
    v3 := make(map[int]int)
    var v4 int16 = 1000
    var v5 uint16 = 255
    v6 := make(map[string]uint32)
    _ = len(v1)
    switch "\n\t" {
    case "a":
        fmt.Println(v5, v4)
    default:
        if len(v1) > 0 {
            v1[0] = 818
        }
    }
    if false {
        switch -9223372036854775808 {
        case 0:
            var v13 uint64 = (v2 >> ((65535) & 63))
            _ = v13
        }
    } else {
        v5 = 984
        _ = len(v1)
    }
    v14 := GenS2{F0: 42, F1: 65535, F2: ""}
    switch "" {
    case "\n\t":
        v15 := GenS1{F0: (v4 + -1), F1: (v5 * v5), F2: 1}
        _ = v15
    case "":
        var v16 int8 = 0
        _ = v16
    default:
        v2 -= 4294967295
    }
    var v17 float64 = float64(-2.25)
    if true {
        if ((false || (1.5 == -2.25)) || (!false)) {
            v17 -= v17
            switch ("hello" + "") {
            case "\n\t":
                if (!(1 < 641)) {
                    switch "\n\t" {
                    case "foo bar":
                        v1 = append(v1, 127)
                    case "a":
                        if ((^(v4)) != v4) {
                            v2 -= (v2 / ((v2) | 1))
                            if true {
                                var v18 uint64 = 42
                                v19 := GenS3{F0: 127}
                                _ = v18
                                _ = v19
                            } else {
                                v5 = (^(v5))
                                v5 = 317
                                switch "foo bar" {
                                case "\n\t":
                                    _ = len(v6)
                                default:
                                    v5 = v5
                                }
                            }
                            for fi20 := 0; fi20 < 10; fi20++ {
                                v5 += 254
                                _ = fi20
                            }
                        } else {
                            switch "\n\t" {
                            case "":
                                switch "\n\t" {
                                case "\n\t":
                                    for fi21 := 0; fi21 < 1; fi21++ {
                                        for fi22 := 0; fi22 < 10; fi22++ {
                                            var v23 uint = 619
                                            v4 += v4
                                            _ = fi22
                                            _ = v23
                                        }
                                        var v24 float32 = (((0.0 + float32((-2.25 + v17))) - (0.1 - 3.14159)) / 0.1)
                                        _ = fi21
                                        _ = v24
                                    }
                                case "foo bar":
                                    if true {
                                        var v25 uint64 = 42
                                        var v26 int32 = 42
                                        v25 = (v25 * v25)
                                        _ = v25
                                        _ = v26
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v27 uint16 = v5
                                        fmt.Println(v4, v4)
                                        _ = v27
                                    }
                                }
                            }
                            fmt.Println(v4, v17)
                            var v28 int = -2147483648
                            _ = v28
                        }
                    case "hello":
                        if true {
                            v29 := GenS3{F0: 42}
                            _ = v29
                        }
                    default:
                        var v30 int = 32768
                        _ = v30
                    }
                    var v31 uint16 = 2
                    var v32 uint32 = 686
                    _ = v31
                    _ = v32
                }
            case "a":
                if false {
                    var v33 int64 = 9223372036854775807
                    _ = v33
                }
            case "hello":
                v34 := GenS3{F0: -51}
                _ = v34
            }
        }
        switch 128 {
        case -9223372036854775808:
            fmt.Println(v2)
        default:
            v35 := GenS3{F0: 42}
            _ = v35
        }
        for fi36 := 0; fi36 < 19; fi36++ {
            v4 = -662
            var v37 float64 = ((1e10 * v17) - -2.25)
            _ = fi36
            _ = v37
        }
    } else {
        var v38 int32 = -448
        _ = v38
    }
    for fi39 := 0; fi39 < 11; fi39++ {
        var v40 int16 = v4
        v1 = append(v1, (fi39 / (((fi39 ^ 127)) | 1)))
        v41 := GenS1{F0: v40, F1: (v5 * v5), F2: 65535}
        _ = fi39
        _ = v40
        _ = v41
    }
    v5 = v5
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
