package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int }
func genFn0(p0 int, p1 string) int32 {
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 float32 = 3.14159
        _ = fi6
        _ = v7
    }
    if (0 <= 42) {
        var v8 bool = (false && ((!true) && true))
        var v9 int16 = 42
        fmt.Println(v8, v8)
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 uint64 = 42
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := []int{}
            v14 := GenS1{F0: v11, F1: 1, F2: 32767}
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 int64 = -9223372036854775808
        _ = fi10
        _ = v11
        _ = v15
    }
    var v16 bool = true
    _ = p0
    _ = p1
    _ = v16
    return -2147483648
}

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    var v4 uint16 = 255
    v5 := []uint32{}
    switch ((((("" + "foo bar") + "") + "\n\t") + "\n\t") + "a") {
    case "":
        if false {
            v1 = v1
            var v17 uint16 = v4
            v18 := GenS2{F0: 42, F1: 42}
            _ = v17
            _ = v18
        } else {
            for fi19 := 0; fi19 < 3; fi19++ {
                v4 -= (v4 - v4)
                _ = fi19
            }
        }
    case "\n\t":
        if (!false) {
            fmt.Println(v4, v1)
            fmt.Println(v1, v4)
        } else {
            fmt.Println(v4, v1)
            v1 = (761 != 255)
            var v20 int = 127
            _ = v20
        }
    case "foo bar":
        delete(v3, 42)
    default:
        v4 += v4
    }
    switch -1 {
    case 744:
        for fi21 := 0; fi21 < 2; fi21++ {
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 uint = 65535
                v24 := GenS3{F0: 2147483647, F1: 242, F2: 0}
                _ = fi22
                _ = v23
                _ = v24
            }
            if (1 > 1) {
                switch "foo bar" {
                case "":
                    v25 := GenS1{F0: 791, F1: v4, F2: 0}
                    _ = v25
                case "\n\t":
                    fmt.Println(fi21)
                default:
                    if (v1 && true) {
                        if true {
                            switch "a" {
                            case "":
                                var v26 float64 = ((-2.25 - 1e10) - (float64(0.0) * (1e10 + 1.5)))
                                _ = v26
                            default:
                                v27 := GenS3{F0: -847, F1: 42, F2: 2}
                                _ = v27
                            }
                            switch (fi21 >> ((65535) & 63)) {
                            case 399:
                                var v28 int64 = -9223372036854775808
                                _ = v28
                            case 127:
                                var v29 int = fi21
                                _ = v29
                            }
                        }
                    } else {
                        v30 := GenS4{F0: (^(fi21))}
                        var v31 int32 = 865
                        if false {
                            var v32 uint8 = 2
                            v5 = append(v5, 4294967295)
                            _ = v32
                        } else {
                            v33 := GenS4{F0: (fi21 >> ((2) & 63))}
                            fmt.Println(v4, fi21)
                            v1 = false
                            _ = v33
                        }
                        _ = v30
                        _ = v31
                    }
                }
            }
            for fi34 := 0; fi34 < 6; fi34++ {
                _ = len(v3)
                fmt.Println(fi21, v1, fi34)
                _ = fi34
            }
            _ = fi21
        }
    }
    v4 += 32767
    if (156 > -1) {
        fmt.Println(v1)
    }
    v35 := GenS1{F0: 18446744073709551615, F1: (v4 - v4), F2: 652}
    for fi36 := 0; fi36 < 5; fi36++ {
        var v37 uint8 = 178
        _ = fi36
        _ = v37
    }
    v4 = v4
    var v38 int32 = 1
    var v39 bool = false
    v38 = v38
    v38 = (v38 >> ((827) & 63))
    v38 = -2147483648
    var v40 uint64 = 706
    v1 = (float32(1.5) > float32(-2.25))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
