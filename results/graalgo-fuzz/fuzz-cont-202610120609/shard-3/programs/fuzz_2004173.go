package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint32; F1 int64; F2 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 int64 = 42
    var v2 int8 = 88
    v3 := []int{}
    v4 := make(map[int]bool)
    v2 = -128
    v5 := GenS1{F0: 2, F1: 42, F2: 455}
    if true {
        v2 += v2
        v6 := GenS1{F0: 0, F1: 42, F2: 2}
        var v7 uint16 = 32767
        _ = v6
        _ = v7
    }
    switch -2147483648 {
    case -354:
        delete(v4, 2147483647)
    case 0:
        if false {
            if (true || (false && false)) {
                var v8 uint = 4294967295
                _ = v8
            } else {
                v9 := GenS2{F0: v1, F1: 230, F2: v2}
                v1 = (v1 | -9223372036854775808)
                if (1 <= 519) {
                    v1 += v1
                    if (("" + "hello") < "\n\t") {
                        v2 = 3
                        v1 = v1
                    }
                    v1 = (v1 / (((v1 ^ v1)) | 1))
                } else {
                    var v10 uint64 = 807
                    _ = v10
                }
                _ = v9
            }
            fmt.Println(v1, v1)
        }
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint32 = 0
        _ = fi11
        _ = v12
    }
    v1 = 9223372036854775807
    var v13 int16 = 32767
    if false {
        v4[127] = (!(18446744073709551615 >= 0))
        v3 = append(v3, 32768)
        fmt.Println(v2)
    } else {
        var v14 uint64 = 1
        _ = v14
    }
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    switch -580 {
    case 128:
        switch ("\n\t" + "foo bar") {
        case "hello":
            if (v2 < (^(v2))) {
                v2 = (v2 * (v2 ^ v2))
            }
        case "a":
            v2 = (v2 << ((630) & 63))
        default:
            var v15 int64 = (v1 % ((v1) | 1))
            _ = v15
        }
    case 4294967295:
        v4[0] = true
    case 369:
        if (v13 == (v13 | v13)) {
            if true {
                var v16 uint = 1
                _ = v16
            } else {
                var v17 uint64 = 526
                _ = v17
            }
            for fi18 := 0; fi18 < 4; fi18++ {
                v2 += v2
                var v19 int = (fi18 | (fi18 ^ (fi18 & (-(fi18)))))
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 20; fi20++ {
                for fi21 := 0; fi21 < 8; fi21++ {
                    delete(v4, (^(fi20)))
                    switch (("\n\t" + "foo bar") + (("hello" + ("foo bar" + "")) + "hello")) {
                    case "a":
                        var v22 float32 = (1e10 * -2.25)
                        _ = v22
                    case "hello":
                        if false {
                            if false {
                                v4[fi21] = (!(!false))
                                v23 := GenS3{F0: 42, F1: 0, F2: "foo bar"}
                                _ = v23
                            }
                            v24 := GenS4{F0: fi20}
                            _ = v24
                        } else {
                            fmt.Println(v2, fi21)
                            v2 -= (v2 / (((v2 + v2)) | 1))
                        }
                    }
                    v13 += 42
                    _ = fi21
                }
                if (true && false) {
                    var v25 float64 = 1.5
                    var v26 uint16 = 0
                    _ = v25
                    _ = v26
                } else {
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    if (!((v1 ^ -9223372036854775808) < 1)) {
                        _ = len(v4)
                    } else {
                        if true {
                            v13 -= v13
                        }
                    }
                    if (!(1 <= 65535)) {
                        if false {
                            delete(v4, fi20)
                            switch "a" {
                            case "":
                                v27 := v4[int(v2)]
                                _ = v27
                            case "hello":
                                v2 -= 127
                            case "foo bar":
                                if ((float32((-2.25 / float64((float32(float64((1e10 / (3.14159 + 0.0)))) - (0.0 - (1e10 * 3.14159)))))) > 1.5) || false) {
                                    v1 = -1
                                }
                            default:
                                fmt.Println(v2, fi20)
                            }
                        }
                        switch (fi20 | fi20) {
                        case 4294967295:
                            v13 += v13
                        case 924:
                            var v28 int32 = 2147483647
                            _ = v28
                        default:
                            if (18446744073709551615 < 65535) {
                                v1 -= v1
                                if true {
                                    v2 += v2
                                } else {
                                    for fi29 := 0; fi29 < 19; fi29++ {
                                        var v30 int8 = (v2 ^ v2)
                                        v13 = v13
                                        v31 := GenS5{F0: (v1 % ((v1) | 1))}
                                        _ = fi29
                                        _ = v30
                                        _ = v31
                                    }
                                    switch fi20 {
                                    case -387:
                                        fmt.Println(v1, v2)
                                    }
                                    var v32 uint8 = 127
                                    _ = v32
                                }
                            }
                        }
                    }
                }
                _ = fi20
            }
        }
    default:
        var v33 int8 = v2
        _ = v33
    }
    v2 -= v2
    fmt.Println(v1, v1)
    for fi34 := 0; fi34 < 5; fi34++ {
        var v35 uint16 = 231
        v4[fi34] = true
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        var v37 uint8 = 0
        _ = fi36
        _ = v37
    }
    switch "\n\t" {
    case "":
        v38 := GenS6{F0: 1}
        _ = v38
    case "\n\t":
        fmt.Println(v13, v2)
    case "foo bar":
        v39 := GenS1{F0: 42, F1: 42, F2: 92}
        _ = v39
    default:
        v40 := GenS7{F0: 42, F1: 2}
        _ = v40
    }
    for fi41 := 0; fi41 < 9; fi41++ {
        var v42 float32 = float32(1e10)
        var v43 int = fi41
        var v44 bool = ((true && false) || (false && (127 > 99)))
        _ = fi41
        _ = v42
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
