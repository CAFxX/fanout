package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 uint32; F2 bool }
type GenS4 struct { F0 int16; F1 int16; F2 uint16 }
type GenS5 struct { F0 int }
func genFn0(p0 uint16, p1 uint8, p2 int64) string {
    if (!false) {
        p2 = 9223372036854775807
        var v6 string = "\n\t"
        _ = v6
    } else {
        for fi7 := 0; fi7 < 18; fi7++ {
            p2 = (p2 / ((p2) | 1))
            _ = fi7
        }
        if (-116 >= 1) {
            v8 := GenS1{F0: (p2 + p2), F1: 128}
            v9 := []int{32768, -1}
            _ = v8
            _ = v9
        } else {
            if false {
                var v10 float64 = float64((0.0 / 1.5))
                v10 = 0.1
                _ = v10
            } else {
                v11 := GenS1{F0: p2, F1: 1}
                _ = v11
            }
            var v12 float64 = float64(0.1)
            _ = v12
        }
        v13 := []int{9223372036854775807, 0, 32768}
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := []int{}
        var v16 int8 = 127
        if true {
            v17 := GenS2{F0: (v16 / ((0) | 1))}
            v18 := GenS3{F0: 2147483647, F1: 2, F2: false}
            p2 = (p2 | 0)
            _ = v17
            _ = v18
        } else {
            if true {
                for fi19 := 0; fi19 < 12; fi19++ {
                    var v20 uint = 42
                    var v21 uint = v20
                    var v22 uint32 = 2147483647
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                v23 := GenS3{F0: 0, F1: 42, F2: true}
                _ = v23
            }
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        if (1000 >= -149) {
            var v25 uint8 = (^(p1))
            var v26 uint32 = 65535
            v27 := []int{(fi24 ^ (fi24 ^ 0)), (fi24 * 127), (-(fi24)), int(p0)}
            _ = v25
            _ = v26
            _ = v27
        }
        var v28 uint8 = p1
        _ = fi24
        _ = v28
    }
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    return "hello"
}

func main() {
    var v1 uint = 600
    v2 := []int{}
    var v3 int64 = -104
    var v4 uint16 = 42
    var v5 int16 = 1000
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v29 float64 = (1.5 / -2.25)
    v30 := GenS4{F0: (v5 * 1000), F1: (-(v5)), F2: v4}
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v3 -= -1
        v2 = append(v2, -9223372036854775808)
        v5 = v5
    }
    v5 = (v5 * 32767)
    v3 += 1
    var v31 int8 = 42
    var v32 int32 = -1
    if ((!false) && (uint64(v1) > 0)) {
        v33 := GenS5{F0: -1}
        switch "a" {
        case "foo bar":
            for fi34 := 0; fi34 < 19; fi34++ {
                v32 += v32
                _ = fi34
            }
        case "":
            v3 = v3
        case "\n\t":
            switch "hello" {
            case "":
                v35 := GenS3{F0: (v32 % (((-(v32))) | 1)), F1: 62, F2: (65535 >= 247)}
                _ = v35
            case "hello":
                v32 = -1
            case "a":
                var v36 int8 = v31
                _ = v36
            }
        default:
            v5 = (v5 % (((v5 / (((v5 % (((v5 - -108)) | 1))) | 1))) | 1))
        }
        var v37 int8 = (v31 + (v31 ^ (v31 % ((v31) | 1))))
        _ = v33
        _ = v37
    } else {
        for fi38 := 0; fi38 < 5; fi38++ {
            var v39 float64 = ((v29 * v29) * v29)
            _ = fi38
            _ = v39
        }
        for fi40 := 0; fi40 < 11; fi40++ {
            for fi41 := 0; fi41 < 19; fi41++ {
                v5 += (v5 % (((v5 & v5)) | 1))
                fmt.Println(v1, v31)
                _ = fi41
            }
            _ = fi40
        }
        v42 := GenS5{F0: 128}
        _ = v42
    }
    v5 = v5
    v3 = v3
    v43 := GenS4{F0: v5, F1: -574, F2: 0}
    if false {
        var v44 uint = v1
        v5 -= -1
        var v45 int32 = -1
        _ = v44
        _ = v45
    } else {
        fmt.Println(v1, v32)
        for fi46 := 0; fi46 < 10; fi46++ {
            fmt.Println(fi46, v4)
            fmt.Println(v3, v29, fi46)
            v1 -= (v1 * (v1 << (((v1 | (v1 - v1))) & 63)))
            _ = fi46
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v43
    _ = v5
}
