package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint; F1 uint64; F2 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 int }
func genFn0(p0 int, p1 uint8) int16 {
    switch 32767 {
    case -159:
        p0 = (-(p0))
    case 2147483647:
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 uint64 = 42
            _ = fi6
            _ = v7
        }
    }
    fmt.Println(p1, p0)
    v8 := GenS1{F0: 42}
    if (2 >= 2147483647) {
        v9 := GenS2{F0: 2}
        v10 := GenS3{F0: 42, F1: 0}
        v11 := []int{(p0 / ((p0) | 1)), 32767}
        _ = v10
        _ = v11
        _ = v9
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v8
    return -32768
}
func genFn1(p0 bool, p1 int8, p2 int32) int8 {
    var v12 int16 = 1000
    v13 := []int{-9223372036854775808}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return p1
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[int]int)
    var v3 uint8 = 127
    var v4 uint8 = (v3 % (((v3 & 191)) | 1))
    var v5 uint = 2
    if false {
        if (!true) {
            switch ("hello" + "hello") {
            case "hello":
                var v14 uint32 = 2147483647
                _ = v14
            default:
                var v15 int32 = -1
                _ = v15
            }
        } else {
            var v16 int64 = 42
            _ = v16
        }
        fmt.Println(v1)
    }
    switch "foo bar" {
    case "foo bar":
        v4 = v3
    case "":
        v1 += v1
    default:
        var v17 uint32 = 1
        _ = v17
    }
    var v18 float64 = (0.0 * -2.25)
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 string = ""
        switch fi19 {
        case -1:
            if true {
                if (2 >= 32767) {
                    v21 := GenS1{F0: (v3 - (v4 << ((v5) & 63)))}
                    fmt.Println(v18)
                    _ = v21
                } else {
                    var v22 float32 = (1.5 / 0.1)
                    var v23 int64 = 1
                    _ = v22
                    _ = v23
                }
                fmt.Println(v20)
                var v24 bool = false
                _ = v24
            }
        }
        _ = fi19
        _ = v20
    }
    var v25 int32 = 42
    switch "hello" {
    case "":
        v25 = (^(v25))
    case "\n\t":
        for fi26 := 0; fi26 < 15; fi26++ {
            v27 := GenS4{F0: v5, F1: 4294967295, F2: (v25 << ((42) & 63))}
            for fi28 := 0; fi28 < 5; fi28++ {
                v29 := GenS5{F0: -9223372036854775808}
                _ = fi28
                _ = v29
            }
            _ = fi26
            _ = v27
        }
    case "foo bar":
        for fi30 := 0; fi30 < 12; fi30++ {
            fmt.Println(v5, v4)
            var v31 uint = (v5 * v5)
            _ = fi30
            _ = v31
        }
    }
    v32 := GenS6{F0: ("" + "hello"), F1: -2147483648}
    for fi33 := 0; fi33 < 8; fi33++ {
        fmt.Println(v4, fi33, v5)
        var v34 int = (fi33 & (fi33 | (fi33 >> (((v5 - v5)) & 63))))
        _ = fi33
        _ = v34
    }
    if true {
        v5 -= 226
        var v35 float32 = ((1e10 - 1.5) - (((0.1 * -2.25) * float32(v18)) / 1.0))
        if (false && true) {
            v1 = 18446744073709551615
            for fi36 := 0; fi36 < 6; fi36++ {
                if ((828 >= 0) || false) {
                    switch "" {
                    case "hello":
                        v3 -= 0
                    case "":
                        var v37 int32 = 811
                        _ = v37
                    }
                    v1 = 42
                    var v38 int64 = 9223372036854775807
                    _ = v38
                }
                fmt.Println(v25)
                v39 := []int{2147483647}
                _ = fi36
                _ = v39
            }
            for fi40 := 0; fi40 < 8; fi40++ {
                var v41 uint16 = 786
                _ = fi40
                _ = v41
            }
        }
        _ = v35
    } else {
        switch "hello" {
        case "foo bar":
            if (127 <= 42) {
                v3 = (v4 * v4)
            }
        case "\n\t":
            for fi42 := 0; fi42 < 12; fi42++ {
                if (false && true) {
                    v18 = float64(3.14159)
                    if (int(v1) <= (fi42 | fi42)) {
                        for fi43 := 0; fi43 < 19; fi43++ {
                            v1 = (v1 << ((v5) & 63))
                            _ = fi43
                        }
                        for fi44 := 0; fi44 < 4; fi44++ {
                            v18 += (3.14159 + (v18 * (v18 + v18)))
                            if true {
                                switch -2147483648 {
                                case 42:
                                    v3 = v4
                                case -9223372036854775808:
                                    v1 = (^(v1))
                                    var v45 int16 = 1
                                    _ = v45
                                case -946:
                                    fmt.Println(fi42, v18)
                                }
                                delete(v2, -2147483648)
                                v46 := GenS2{F0: (v1 / (((v1 | (v1 % ((v1) | 1)))) | 1))}
                                _ = v46
                            } else {
                                var v47 float32 = 3.14159
                                v3 += v3
                                _ = v47
                            }
                            _ = fi44
                        }
                    } else {
                        var v48 float64 = 0.0
                        v49 := []int{fi42, 32767, (fi42 ^ (fi42 | fi42)), fi42}
                        var v50 uint8 = (v4 >> (((v5 | (^(v5)))) & 63))
                        _ = v48
                        _ = v49
                        _ = v50
                    }
                    var v51 int8 = -1
                    _ = v51
                }
                v3 = (v4 % (((v3 & (^(v4)))) | 1))
                _ = fi42
            }
        default:
            _ = len(v2)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
