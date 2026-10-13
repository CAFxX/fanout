package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint64; F1 int16 }
func genFn0(p0 int16) int8 {
    var v4 float32 = 0.1
    var v5 uint64 = 42
    v6 := []int{}
    v7 := GenS1{F0: (3.14159 > 1e10), F1: 882}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 100
}

func main() {
    var v1 float32 = 0.1
    var v2 uint = 0
    v3 := []int{-223, -2147483648, -421, -817}
    if true {
        v2 += v2
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(fi8, v2)
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 string = "a"
            _ = fi9
            _ = v10
        }
        var v11 uint16 = 65535
        _ = fi8
        _ = v11
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        if true {
            v1 = float32((0.1 * (0.1 / 0.1)))
        }
        if ((false || (-391 == -306)) && true) {
            if len(v3) > 0 {
                v3[0] = fi12
            }
            var v13 int8 = 127
            v13 = (v13 & 42)
            _ = v13
        } else {
            v2 += (^(v2))
            if true {
                v2 = (v2 * v2)
                if false {
                    v2 = (v2 / (((v2 - v2)) | 1))
                    v2 -= (v2 & 255)
                    v1 -= (((((v1 * v1) * v1) * v1) / (v1 - v1)) / float32(1.5))
                }
            } else {
                fmt.Println(fi12)
                var v14 int = 4294967295
                _ = v14
            }
            v15 := GenS2{F0: 42}
            _ = v15
        }
        if true {
            switch "foo bar" {
            case "hello":
                switch -283 {
                case 42:
                    for fi16 := 0; fi16 < 2; fi16++ {
                        var v17 uint8 = 2
                        _ = fi16
                        _ = v17
                    }
                case 9223372036854775807:
                    if ((2147483647 != (fi12 - (fi12 >> (((v2 - v2)) & 63)))) && false) {
                        for fi18 := 0; fi18 < 11; fi18++ {
                            var v19 uint16 = 255
                            v19 = v19
                            v20 := GenS3{F0: (v2 << ((v2) & 63))}
                            fmt.Println(v1, fi18, fi12)
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        switch "hello" {
                        case "foo bar":
                            v21 := GenS4{F0: 42}
                            _ = v21
                        case "hello":
                            for fi22 := 0; fi22 < 15; fi22++ {
                                switch fi12 {
                                case -9223372036854775808:
                                    v3 = append(v3, (fi22 / (((fi12 & -179)) | 1)))
                                case 225:
                                    var v23 uint8 = 0
                                    _ = v23
                                }
                                var v24 int16 = 42
                                v1 -= 1.5
                                _ = fi22
                                _ = v24
                            }
                        case "":
                            v25 := GenS5{F0: -114, F1: false}
                            _ = v25
                        default:
                            v26 := GenS6{F0: 4294967295, F1: 42}
                            _ = v26
                        }
                        fmt.Println(v2)
                    }
                }
            case "\n\t":
                v2 += v2
            default:
                fmt.Println(v2, fi12, v1)
            }
            for fi27 := 0; fi27 < 7; fi27++ {
                v3 = append(v3, int(v2))
                _ = fi27
            }
        }
        _ = fi12
    }
    var v28 int64 = -9223372036854775808
    v1 = ((v1 / float32(3.14159)) - v1)
    v28 = v28
    var v29 int16 = 1000
    for fi30 := 0; fi30 < 11; fi30++ {
        for fi31 := 0; fi31 < 4; fi31++ {
            v28 = -9223372036854775808
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(fi30, v2)
            _ = fi31
        }
        _ = fi30
    }
    v1 -= 0.0
    if false {
        var v32 float32 = v1
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v32
    }
    switch 9223372036854775807 {
    case -1:
        for fi33 := 0; fi33 < 16; fi33++ {
            var v34 uint8 = 128
            v29 = 290
            _ = fi33
            _ = v34
        }
    case 844:
        v29 = (v29 >> ((v2) & 63))
    case 1:
        fmt.Println(v29, v1, v28)
    }
    v29 = (v29 >> ((v2) & 63))
    var v35 uint16 = 870
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v35
}
