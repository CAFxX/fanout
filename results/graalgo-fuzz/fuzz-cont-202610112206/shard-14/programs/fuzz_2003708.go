package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS5 struct { F0 bool; F1 uint16 }
func genFn0() int16 {
    var v7 float32 = 3.14159
    v7 += v7
    v8 := []int{-444, 4294967295, -2147483648}
    _ = v7
    _ = v8
    return 0
}
func genFn1() string {
    v9 := []int{1, 9223372036854775807}
    v10 := GenS1{F0: 127}
    _ = len(v9)
    v11 := GenS1{F0: 128}
    var v12 uint64 = 42
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return ("foo bar" + ("a" + "\n\t"))
}

func main() {
    var v1 uint64 = 374
    var v2 uint8 = 30
    var v3 uint32 = 42
    var v4 uint32 = (v3 << ((399) & 63))
    var v5 int = -847
    var v6 float64 = 3.14159
    if (!false) {
        var v13 bool = (false || (v1 < v1))
        fmt.Println(v2, v6)
        _ = v13
    } else {
        fmt.Println(v5, v5, v4)
        v14 := GenS2{F0: false, F1: 274, F2: 1}
        var v15 int16 = -512
        _ = v14
        _ = v15
    }
    v16 := []int{int(v1), v5}
    fmt.Println(v1)
    switch (v5 / ((v5) | 1)) {
    case 320:
        var v17 uint64 = (v1 % (((v1 | v1)) | 1))
        _ = v17
    case 464:
        var v18 uint16 = 715
        _ = v18
    case 163:
        switch v5 {
        case 4294967295:
            v19 := GenS3{F0: v3, F1: 255, F2: v3}
            _ = v19
        case 9223372036854775807:
            var v20 uint64 = 0
            _ = v20
        case 659:
            for fi21 := 0; fi21 < 15; fi21++ {
                for fi22 := 0; fi22 < 8; fi22++ {
                    if false {
                        var v23 int64 = int64(v2)
                        switch (fi21 * -265) {
                        case 9223372036854775807:
                            v16 = append(v16, fi22)
                        case -9223372036854775808:
                            v24 := GenS4{F0: -32768, F1: (v1 & (v1 / (((v1 >> ((201) & 63))) | 1))), F2: 127}
                            _ = v24
                        case -491:
                            fmt.Println(fi21)
                        }
                        if true {
                            var v25 int = 0
                            if len(v16) > 0 {
                                v16[0] = (fi22 * fi22)
                            }
                            _ = v25
                        }
                        _ = v23
                    } else {
                        v6 = (v6 + v6)
                        fmt.Println(fi21, fi22, v6)
                        v5 = v5
                    }
                    if true {
                        if len(v16) > 0 {
                            v16[0] = (-(fi22))
                        }
                    } else {
                        fmt.Println(v4)
                        switch "\n\t" {
                        case "":
                            fmt.Println(v3, fi22, v5)
                            _ = len(v16)
                            v5 -= (v5 - (fi22 - fi21))
                        default:
                            var v26 int16 = 260
                            _ = v26
                        }
                        var v27 float64 = (float64(float32(v6)) - (-2.25 - 0.1))
                        _ = v27
                    }
                    switch "a" {
                    case "\n\t":
                        fmt.Println(fi21, v4, v1)
                        v6 -= 1e10
                        v6 += (1e10 * 0.1)
                    }
                    v6 -= 1.5
                    v4 -= v3
                    _ = fi22
                }
                v28 := GenS4{F0: -1, F1: 42, F2: (v2 << ((18446744073709551615) & 63))}
                _ = fi21
                _ = v28
            }
        }
    default:
        var v29 uint8 = (v2 * (v2 << ((566) & 63)))
        _ = v29
    }
    var v30 uint = 38
    if (true || true) {
        _ = len(v16)
    } else {
        var v31 int = 735
        var v32 float32 = 3.14159
        _ = v31
        _ = v32
    }
    fmt.Println(v5, v6)
    if len(v16) > 0 {
        v16[0] = (-(v5))
    }
    v1 -= (v1 % ((0) | 1))
    var v33 uint16 = 65535
    if true {
        fmt.Println(v1)
        v34 := GenS4{F0: 1000, F1: v1, F2: v2}
        _ = v34
    }
    if true {
        if (!(!false)) {
            fmt.Println(v4)
            v30 -= uint(v1)
        }
        for fi35 := 0; fi35 < 5; fi35++ {
            for fi36 := 0; fi36 < 5; fi36++ {
                v2 = (v2 ^ (v2 % ((v2) | 1)))
                switch "hello" {
                case "foo bar":
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    v37 := GenS5{F0: (!("foo bar" < "\n\t")), F1: v33}
                    _ = v37
                default:
                    fmt.Println(fi35, v1)
                }
                if ((!true) && (("" + "\n\t") == ("hello" + "hello"))) {
                    var v38 bool = false
                    _ = v38
                }
                fmt.Println(fi35, v6)
                _ = fi36
            }
            _ = fi35
        }
    }
    for fi39 := 0; fi39 < 17; fi39++ {
        var v40 int64 = -430
        switch (fi39 >> ((42) & 63)) {
        case 0:
            v5 -= fi39
        case -31:
            v3 = (^(v3))
        }
        switch "a" {
        case "hello":
            var v41 uint16 = v33
            _ = v41
        case "":
            fmt.Println(v6, v40)
        default:
            for fi42 := 0; fi42 < 19; fi42++ {
                var v43 uint = 614
                var v44 string = "a"
                switch v44 {
                case "hello":
                    if len(v16) > 0 {
                        v16[0] = (-(v5))
                    }
                case "foo bar":
                    var v45 uint = 4294967295
                    _ = v45
                case "a":
                    switch (^(fi39)) {
                    case 32768:
                        var v46 bool = (false || false)
                        _ = v46
                    case 483:
                        v43 += 255
                    }
                default:
                    var v47 int16 = -759
                    _ = v47
                }
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                _ = fi42
                _ = v43
                _ = v44
            }
        }
        _ = fi39
        _ = v40
    }
    _ = len(v16)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
