package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 bool; F2 uint32 }
type GenS3 struct { F0 int64; F1 bool }
func genFn0(p0 uint8, p1 string) int64 {
    switch "hello" {
    case "foo bar":
        p0 -= p0
    case "hello":
        fmt.Println(p0, p1)
    case "a":
        if (127 > 1) {
            if false {
                fmt.Println(p1, p1)
                var v7 float32 = ((0.1 - (3.14159 / 1.0)) * float32(1.5))
                _ = v7
            }
            var v8 uint8 = (^(p0))
            _ = v8
        } else {
            var v9 int64 = 1
            _ = v9
        }
    default:
        p0 = (^(p0))
    }
    if (p1 == p1) {
        v10 := GenS1{F0: 986}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 3; fi11++ {
            if ((4294967295 > 0) && (!false)) {
                v12 := GenS1{F0: 65535}
                var v13 int64 = -9223372036854775808
                var v14 int32 = -1
                _ = v12
                _ = v13
                _ = v14
            } else {
                for fi15 := 0; fi15 < 19; fi15++ {
                    v16 := GenS2{F0: -9223372036854775808, F1: false, F2: 2147483647}
                    v17 := []int{0, (fi11 | 32768), 2147483647}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                p1 = p1
                var v18 float32 = 1e10
                _ = v18
            }
            p1 = p1
            p0 += p0
            _ = fi11
        }
        fmt.Println(p0, p0)
        var v19 uint = 42
        _ = v19
    }
    v20 := []int{-9223372036854775808, 838}
    _ = p0
    _ = p1
    _ = v20
    return -9223372036854775808
}

func main() {
    var v1 bool = ((1 > -535) || true)
    var v2 int32 = 277
    var v3 float64 = 1.5
    v4 := []uint32{65535, 2, 65535, 394}
    var v5 uint16 = 1
    var v6 uint = 508
    if len(v4) > 0 {
        v4[0] = 2
    }
    fmt.Println(v6, v6)
    var v21 int8 = 0
    v4 = append(v4, 2147483648)
    fmt.Println(v6, v6)
    v2 -= v2
    if (-128 < v21) {
        var v22 float32 = -2.25
        fmt.Println(v2, v3, v21)
        _ = v22
    } else {
        if len(v4) > 0 {
            v4[0] = 65535
        }
        if false {
            switch "\n\t" {
            case "":
                v23 := GenS2{F0: -9223372036854775808, F1: v1, F2: 2147483648}
                _ = v23
            case "a":
                var v24 int8 = v21
                _ = v24
            case "foo bar":
                var v25 string = "hello"
                _ = v25
            }
        }
    }
    fmt.Println(v5)
    var v26 int64 = -9223372036854775808
    for fi27 := 0; fi27 < 2; fi27++ {
        var v28 uint16 = (^(v5))
        var v29 uint64 = 523
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 int8 = -118
    for fi31 := 0; fi31 < 14; fi31++ {
        v30 += v30
        v32 := GenS1{F0: (v6 * (v6 & v6))}
        v1 = v1
        _ = fi31
        _ = v32
    }
    if false {
        v5 -= (^(v5))
        if (!(v1 || (("hello" + "foo bar") != ""))) {
            switch 0 {
            case 0:
                v2 = (v2 & (v2 ^ 1))
            case 110:
                switch "" {
                case "\n\t":
                    switch ("" + "hello") {
                    case "":
                        v33 := GenS1{F0: v6}
                        _ = v33
                    case "a":
                        fmt.Println(v21, v26, v30)
                    default:
                        for fi34 := 0; fi34 < 1; fi34++ {
                            var v35 int64 = v26
                            _ = fi34
                            _ = v35
                        }
                    }
                }
            default:
                var v36 int64 = v26
                _ = v36
            }
        } else {
            if len(v4) > 0 {
                v4[0] = 0
            }
        }
        var v37 int32 = v2
        _ = v37
    } else {
        if (!((v5 / (((v5 + 562)) | 1)) <= 32767)) {
            v30 -= (-(v21))
            v38 := GenS3{F0: int64(v21), F1: v1}
            for fi39 := 0; fi39 < 11; fi39++ {
                v5 += (v5 - 1)
                _ = fi39
            }
            _ = v38
        } else {
            v6 -= (v6 / ((v6) | 1))
            if (v26 != -868) {
                var v40 int32 = 37
                fmt.Println(v21)
                _ = v40
            } else {
                switch "\n\t" {
                case "":
                    for fi41 := 0; fi41 < 6; fi41++ {
                        v42 := GenS1{F0: (v6 * v6)}
                        v43 := GenS3{F0: int64(fi41), F1: (v21 > -1)}
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                case "hello":
                    v44 := GenS1{F0: (^(v6))}
                    _ = v44
                }
            }
            _ = len(v4)
        }
        v5 = 960
        v21 += -19
    }
    var v45 uint = v6
    if ((!v1) || true) {
        v6 = v45
    } else {
        var v46 int32 = 940
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
