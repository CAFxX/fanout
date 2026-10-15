package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 string; F1 string; F2 int }

func main() {
    var v1 uint16 = 65535
    v2 := make(map[uint32]string)
    v3 := []uint32{220, 904}
    v4 := make(map[string]int)
    var v5 uint = 65535
    switch 0 {
    case 234:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case -9223372036854775808:
        if ((v5 / (((v5 & 2)) | 1)) != (v5 % (((^(v5))) | 1))) {
            var v6 uint32 = 596
            _ = v6
        } else {
            fmt.Println(v1, v5)
        }
    case 0:
        var v7 string = ""
        _ = v7
    default:
        switch 519 {
        case 32767:
            if true {
                v1 = uint16(v5)
                v8 := GenS1{F0: 1, F1: (("hello" + "hello") + "foo bar")}
                v5 += 4294967295
                _ = v8
            } else {
                _ = len(v3)
                var v9 bool = true
                _ = v9
            }
        case 42:
            var v10 uint = (v5 ^ (v5 << (((v5 >> ((v5) & 63))) & 63)))
            _ = v10
        case 1:
            if true {
                var v11 int32 = 914
                _ = v11
            }
        default:
            var v12 uint64 = 18446744073709551615
            _ = v12
        }
    }
    var v13 bool = (-71 > -1)
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 uint = 909
        var v16 int16 = -702
        _ = fi14
        _ = v15
        _ = v16
    }
    if ((v13 && (!(-677 > 42))) && ((false && (-9223372036854775808 > -807)) || v13)) {
        if false {
            v17 := GenS1{F0: 42, F1: "a"}
            var v18 float32 = 0.0
            if v13 {
                v5 = uint(v1)
                v19 := GenS1{F0: -1, F1: "foo bar"}
                _ = v19
            } else {
                if false {
                    if len(v3) > 0 {
                        v3[0] = 1
                    }
                } else {
                    for fi20 := 0; fi20 < 1; fi20++ {
                        if v13 {
                            v21 := GenS1{F0: 9223372036854775807, F1: "foo bar"}
                            _ = v21
                        }
                        var v22 uint16 = (v1 << ((v5) & 63))
                        _ = fi20
                        _ = v22
                    }
                    var v23 uint = v5
                    _ = v23
                }
            }
            _ = v17
            _ = v18
        }
        var v24 uint = (v5 + v5)
        if v13 {
            fmt.Println(v24, v24, v1)
            v1 -= 952
        } else {
            v25 := GenS2{F0: "hello", F1: ("hello" + "\n\t"), F2: 32768}
            var v26 int8 = 127
            _ = v25
            _ = v26
        }
        _ = v24
    }
    var v27 int32 = 2147483647
    fmt.Println(v27, v1)
    fmt.Println(v13, v13, v27)
    var v28 int16 = 42
    if (!v13) {
        switch "" {
        case "hello":
            v27 += v27
        }
        fmt.Println(v27)
        v28 = v28
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        v27 = v27
        _ = fi29
    }
    v3 = append(v3, 42)
    for fi30 := 0; fi30 < 12; fi30++ {
        for fi31 := 0; fi31 < 4; fi31++ {
            v28 = (^(v28))
            v5 = v5
            _ = fi31
        }
        if false {
            if false {
                v32 := GenS2{F0: "\n\t", F1: "hello", F2: (fi30 | (fi30 >> ((v5) & 63)))}
                v28 += 684
                _ = v32
            } else {
                v27 += (v27 * -2147483648)
                var v33 uint = (v5 % ((v5) | 1))
                if v13 {
                    v4[("a" + "a")] = fi30
                }
                _ = v33
            }
            fmt.Println(v27)
        }
        _ = fi30
    }
    var v34 float64 = 0.0
    if (!false) {
        v1 -= v1
        if (!v13) {
            for fi35 := 0; fi35 < 6; fi35++ {
                var v36 int = fi35
                v5 += 255
                var v37 uint32 = 2147483647
                _ = fi35
                _ = v36
                _ = v37
            }
            fmt.Println(v1)
            v1 = v1
        }
        fmt.Println(v1, v13, v28)
    }
    var v38 int64 = 1
    fmt.Println(v34)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v5
}
