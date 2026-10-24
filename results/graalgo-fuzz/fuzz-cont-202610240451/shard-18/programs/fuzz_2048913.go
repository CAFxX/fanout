package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint32) uint16 {
    for fi6 := 0; fi6 < 6; fi6++ {
        p0 = p0
        fmt.Println(p0, p0)
        var v7 uint32 = p0
        _ = fi6
        _ = v7
    }
    v8 := []int{}
    if (((p0 * 65535) != (p0 / ((601) | 1))) || (347 > 2)) {
        if false {
            v9 := GenS1{F0: 42, F1: -1, F2: 42}
            _ = v9
        }
        p0 -= (^(p0))
    } else {
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 uint = 1
            var v12 uint16 = 65535
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 float64 = 1e10
        _ = v13
    }
    v14 := GenS2{F0: 255, F1: false}
    for fi15 := 0; fi15 < 16; fi15++ {
        _ = len(v8)
        fmt.Println(p0)
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 uint64 = 42
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    _ = p0
    _ = v14
    _ = v8
    return 1
}
func genFn1() bool {
    var v18 float64 = 1e10
    var v19 int16 = 940
    switch 2147483647 {
    case -2147483648:
        v18 = 1e10
    case 2147483647:
        v20 := GenS3{F0: 0, F1: 148}
        _ = v20
    default:
        var v21 uint32 = 42
        _ = v21
    }
    _ = v18
    _ = v19
    return (!false)
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{}
    v3 := make(map[int]bool)
    var v4 int8 = 42
    var v5 int = 2147483647
    if len(v2) > 0 {
        v2[0] = 65535
    }
    var v22 int32 = 388
    if true {
        var v23 int8 = (v4 ^ (v4 | v4))
        _ = v23
    }
    switch v5 {
    case 2147483647:
        fmt.Println(v4, v22)
    case 979:
        v2 = append(v2, 65535)
    case -118:
        v4 += v4
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        v5 = 760
        if (42 >= 128) {
            v4 -= v4
            var v25 int64 = 934
            var v26 int16 = -1
            _ = v25
            _ = v26
        }
        if (!(-32768 <= 0)) {
            var v27 int16 = 1
            v4 += v4
            _ = v27
        } else {
            v28 := GenS1{F0: 640, F1: 1000, F2: 4294967295}
            if false {
                if false {
                    v22 = v22
                    v4 = (v4 % ((127) | 1))
                    switch (("foo bar" + "") + "hello") {
                    case "":
                        var v29 float64 = -2.25
                        _ = v29
                    case "foo bar":
                        v30 := GenS4{F0: v4}
                        _ = v30
                    default:
                        for fi31 := 0; fi31 < 11; fi31++ {
                            v5 = (fi31 >> ((65535) & 63))
                            var v32 int = (v5 % ((4294967295) | 1))
                            _ = len(v1)
                            v5 -= (fi24 | (^(fi24)))
                            _ = fi31
                            _ = v32
                        }
                    }
                }
                if (545 != 160) {
                    var v33 uint16 = genFn0(2147483647)
                    _ = v33
                } else {
                    for fi34 := 0; fi34 < 17; fi34++ {
                        v35 := GenS3{F0: 1, F1: -9223372036854775808}
                        switch "foo bar" {
                        case "a":
                            v22 = 1
                        case "\n\t":
                            switch "\n\t" {
                            case "\n\t":
                                var v36 int64 = -975
                                _ = v36
                            case "a":
                                fmt.Println(v5)
                            case "hello":
                                switch v5 {
                                case 2147483647:
                                    v5 -= fi24
                                    v2 = append(v2, 4294967295)
                                case -1:
                                    var v37 int8 = -117
                                    _ = v37
                                }
                            }
                        case "hello":
                            v1[2] = 65535
                        default:
                            v2 = append(v2, 4294967295)
                        }
                        if false {
                            v1[2147483647] = 241
                            v38 := GenS3{F0: 42, F1: -340}
                            _ = v38
                        }
                        _ = fi34
                        _ = v35
                    }
                    var v39 int = v5
                    _ = v39
                }
            } else {
                var v40 uint32 = 2147483647
                _ = v40
            }
            _ = v28
        }
        _ = fi24
    }
    v5 = v5
    v22 += -1
    v4 -= 127
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
