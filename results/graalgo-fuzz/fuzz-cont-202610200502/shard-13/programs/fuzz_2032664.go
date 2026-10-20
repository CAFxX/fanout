package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 bool; F1 int16; F2 uint16 }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0(p0 int, p1 bool, p2 uint64) uint8 {
    var v5 uint = 4294967295
    fmt.Println(p1, p1)
    var v6 string = "a"
    var v7 float32 = float32((-2.25 / (0.0 - 0.1)))
    for fi8 := 0; fi8 < 2; fi8++ {
        p0 -= (p0 % (((^(p0))) | 1))
        var v9 int64 = 42
        var v10 int64 = 42
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int32, p1 uint32, p2 int) int32 {
    var v11 uint64 = uint64(p1)
    var v12 uint32 = (p1 + p1)
    p2 -= (p2 % ((p2) | 1))
    var v13 int16 = 653
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return (p0 << ((104) & 63))
}

func main() {
    var v1 uint16 = 1
    var v2 uint = 0
    var v3 int = 42
    v4 := []uint32{518, 2147483648, 2147483647}
    var v14 uint8 = 127
    v2 = v2
    if true {
        switch v3 {
        case -943:
            var v15 bool = (true || true)
            _ = v15
        default:
            v2 = v2
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := GenS1{F0: 1, F1: (^(v14))}
                _ = fi17
                _ = v18
            }
            fmt.Println(v1)
            _ = fi16
        }
    }
    if true {
        v19 := GenS1{F0: 32767, F1: (v14 % (((^(v14))) | 1))}
        for fi20 := 0; fi20 < 2; fi20++ {
            for fi21 := 0; fi21 < 14; fi21++ {
                if (!false) {
                    v22 := GenS2{F0: false, F1: 1000, F2: (v1 << ((v2) & 63))}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v22
                } else {
                    if true {
                        var v23 int = (v3 & 9223372036854775807)
                        _ = v23
                    }
                    if true {
                        if (!true) {
                            v24 := GenS3{F0: -9223372036854775808, F1: 1000}
                            var v25 float64 = 0.0
                            _ = v24
                            _ = v25
                        }
                    }
                    var v26 string = "hello"
                    v3 = fi20
                    switch (v26 + ("\n\t" + "a")) {
                    case "":
                        v26 = "foo bar"
                    default:
                        v2 += (v2 & v2)
                    }
                    _ = v26
                }
                _ = fi21
            }
            _ = fi20
        }
        var v27 float32 = ((0.1 * float32(0.0)) * 3.14159)
        _ = v19
        _ = v27
    } else {
        fmt.Println(v3)
        v28 := GenS3{F0: -158, F1: -1}
        _ = v28
    }
    var v29 int16 = 1000
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if ((!true) && (!(genFn0(int(v14), false, 4294967295) <= genFn0(int(v14), (true && true), 1)))) {
        v30 := GenS3{F0: 42, F1: (v29 | 1)}
        var v31 int8 = 109
        var v32 uint8 = (v14 | (v14 >> ((uint(v1)) & 63)))
        _ = v30
        _ = v31
        _ = v32
    } else {
        var v33 string = (("a" + "a") + "a")
        for fi34 := 0; fi34 < 2; fi34++ {
            for fi35 := 0; fi35 < 12; fi35++ {
                v33 = v33
                fmt.Println(v33, fi34, v33)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi35
            }
            _ = fi34
        }
        v33 = v33
        _ = v33
    }
    fmt.Println(v14)
    v36 := GenS1{F0: (^(v29)), F1: genFn0(32767, false, 4294967295)}
    fmt.Println(v2, v29)
    if len(v4) > 0 {
        v4[0] = 0
    }
    if (!(true || true)) {
        if true {
            var v37 uint64 = 2
            _ = v37
        } else {
            for fi38 := 0; fi38 < 11; fi38++ {
                switch "foo bar" {
                case "":
                    v3 = (^(fi38))
                default:
                    for fi39 := 0; fi39 < 9; fi39++ {
                        var v40 uint16 = v1
                        var v41 string = (("foo bar" + ("hello" + "hello")) + "foo bar")
                        if (true && (true || false)) {
                            v29 -= -1
                        }
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                }
                _ = fi38
            }
        }
        for fi42 := 0; fi42 < 13; fi42++ {
            var v43 uint8 = (v14 % ((42) | 1))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if true {
                v2 += 34
                v3 += v3
                v43 += (v43 >> ((v2) & 63))
            }
            _ = fi42
            _ = v43
        }
        v1 = v1
    }
    fmt.Println(v1, v1)
    fmt.Println(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v36
    _ = v4
}
