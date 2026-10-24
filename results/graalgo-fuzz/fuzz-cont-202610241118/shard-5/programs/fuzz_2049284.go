package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint64) int16 {
    v6 := GenS1{F0: -1, F1: 100}
    p0 = p0
    p0 += (^(p0))
    v7 := []int{-9223372036854775808, -171, -1, 127}
    switch ("\n\t" + "hello") {
    case "a":
        if (127 == 0) {
            switch 42 {
            case -9223372036854775808:
                for fi8 := 0; fi8 < 18; fi8++ {
                    v9 := GenS1{F0: 24, F1: 73}
                    fmt.Println(p0)
                    _ = fi8
                    _ = v9
                }
            case 42:
                for fi10 := 0; fi10 < 12; fi10++ {
                    v11 := GenS2{F0: ""}
                    _ = fi10
                    _ = v11
                }
            case 32767:
                var v12 uint8 = 2
                _ = v12
            }
            for fi13 := 0; fi13 < 6; fi13++ {
                fmt.Println(p0)
                fmt.Println(p0, p0)
                _ = fi13
            }
        }
    default:
        v14 := GenS2{F0: "\n\t"}
        _ = v14
    }
    _ = p0
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 32767
    var v3 uint = 0
    var v4 uint16 = 32767
    v5 := make(map[uint32]uint32)
    var v15 int = 765
    v3 -= 170
    for fi16 := 0; fi16 < 15; fi16++ {
        if true {
            if true {
                var v17 uint = v3
                var v18 float64 = (1e10 / 1.0)
                _ = v17
                _ = v18
            } else {
                var v19 int8 = -81
                v20 := []int{int(v2), 0}
                v19 = v19
                _ = v19
                _ = v20
            }
        } else {
            fmt.Println(v15, v3)
            switch v15 {
            case 0:
                switch "a" {
                case "a":
                    v2 = 42
                default:
                    if (0.1 >= 3.14159) {
                        switch v15 {
                        case 534:
                            switch (fi16 ^ fi16) {
                            case 32767:
                                fmt.Println(v4, v4, v15)
                            case -1:
                                v2 = -1
                            case 128:
                                v15 = fi16
                            default:
                                var v21 float32 = 0.1
                                _ = v21
                            }
                        case 93:
                            v22 := []int{(v15 - v15), (-(fi16)), (fi16 + (v15 | fi16)), v15}
                            _ = v22
                        case 42:
                            v3 -= v3
                        default:
                            v5[209] = 65535
                        }
                        v23 := GenS1{F0: -2147483648, F1: 1}
                        _ = v23
                    }
                }
            default:
                fmt.Println(v2)
            }
        }
        switch fi16 {
        case 2147483647:
            v24 := GenS2{F0: ("a" + ("foo bar" + ""))}
            _ = v24
        case 735:
            if true {
                for fi25 := 0; fi25 < 19; fi25++ {
                    v15 += (v15 | fi25)
                    switch fi25 {
                    case -784:
                        var v26 int8 = -128
                        _ = v26
                    default:
                        fmt.Println(v15, fi16)
                    }
                    fmt.Println(v15, fi25, fi16)
                    v3 = 1
                    _ = fi25
                }
            }
        default:
            if false {
                for fi27 := 0; fi27 < 6; fi27++ {
                    v4 = v4
                    _ = fi27
                }
            }
        }
        _ = fi16
    }
    var v28 int8 = 89
    if false {
        fmt.Println(v3, v28, v4)
    }
    switch "\n\t" {
    case "a":
        v29 := []int{(v15 * (v15 << ((v3) & 63))), v15, (v15 % ((v15) | 1))}
        _ = v29
    default:
        v2 = -712
    }
    fmt.Println(v28, v28)
    v2 = v2
    fmt.Println(v28, v28, v28)
    v4 = (^(v4))
    v30 := GenS3{F0: 2}
    for fi31 := 0; fi31 < 6; fi31++ {
        v4 -= v4
        v32 := []int{804}
        v2 += (v2 % ((genFn0(0)) | 1))
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        var v34 int64 = -1
        v35 := GenS4{F0: fi33, F1: 4294967295}
        v3 += uint(v4)
        _ = fi33
        _ = v34
        _ = v35
    }
    var v36 string = "hello"
    fmt.Println(v36, v15, v4)
    var v37 int16 = int16(v3)
    var v38 uint8 = 29
    switch v15 {
    case -913:
        v39 := GenS5{F0: 1, F1: 2147483647}
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
