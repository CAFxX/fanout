package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint16; F2 bool }
type GenS3 struct { F0 uint8; F1 int8; F2 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 488
    v3 := make(map[uint32]int)
    v4 := make(map[string]int)
    v5 := []string{}
    v6 := make(map[string]string)
    v7 := GenS1{F0: 154}
    if false {
        var v8 int32 = 0
        _ = len(v1)
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint8 = 194
        _ = fi9
        _ = v10
    }
    switch "foo bar" {
    case "\n\t":
        var v11 int = 32767
        _ = v11
    case "a":
        v12 := GenS1{F0: 127}
        _ = v12
    case "":
        v2 += v2
    }
    v2 -= (^(v2))
    var v13 int64 = -42
    var v14 uint16 = 210
    switch -2147483648 {
    case 687:
        v2 = (v2 * (v2 >> ((uint(v14)) & 63)))
    case -2147483648:
        var v15 int = 1
        _ = v15
    case 4294967295:
        if true {
            var v16 uint32 = 1
            _ = v16
        }
    default:
        for fi17 := 0; fi17 < 16; fi17++ {
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
            var v18 int32 = 1
            _ = fi17
            _ = v18
        }
    }
    switch "hello" {
    case "hello":
        if true {
            v13 = -1
            for fi19 := 0; fi19 < 7; fi19++ {
                if (v2 >= (v2 >> ((1) & 63))) {
                    v13 -= v13
                    v14 = (v14 - (v14 + v14))
                } else {
                    var v20 int16 = v2
                    fmt.Println(v14, fi19, v2)
                    _ = v20
                }
                _ = fi19
            }
        }
    case "\n\t":
        var v21 int8 = 44
        _ = v21
    case "a":
        for fi22 := 0; fi22 < 11; fi22++ {
            fmt.Println(v13, fi22, v13)
            v23 := GenS1{F0: 44}
            _ = fi22
            _ = v23
        }
    }
    var v24 uint32 = 1
    v13 = (v13 * (^(v13)))
    switch "\n\t" {
    case "\n\t":
        var v25 uint = 1
        _ = v25
    case "a":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        switch 32768 {
        case 32767:
            v26 := GenS2{F0: 42, F1: 1, F2: false}
            _ = v26
        case 0:
            _ = len(v3)
        case 127:
            var v27 uint64 = 337
            _ = v27
        }
    }
    switch "\n\t" {
    case "hello":
        if (((("a" + ("a" + ("" + "a"))) + (("\n\t" + "hello") + "hello")) + ("foo bar" + "hello")) >= "a") {
            v2 = (^(v2))
            var v28 int64 = -1
            _ = v28
        } else {
            for fi29 := 0; fi29 < 7; fi29++ {
                v14 = (^(v14))
                _ = len(v4)
                var v30 float32 = float32(-2.25)
                _ = fi29
                _ = v30
            }
            if ((v14 >> ((1) & 63)) >= uint16(v24)) {
                fmt.Println(v14, v13)
                v14 = v14
            }
        }
    case "":
        var v31 bool = (false && ((!(false || (!true))) && true))
        _ = v31
    case "a":
        var v32 uint16 = v14
        _ = v32
    default:
        for fi33 := 0; fi33 < 17; fi33++ {
            v14 += 231
            _ = fi33
        }
    }
    if false {
        for fi34 := 0; fi34 < 1; fi34++ {
            if len(v5) > 0 {
                v5[0] = "foo bar"
            }
            v2 = (v2 * int16(fi34))
            v14 += v14
            _ = fi34
        }
        switch -629 {
        case -309:
            var v35 int16 = v2
            _ = v35
        case -1:
            if (-980 >= 9223372036854775807) {
                var v36 int64 = v13
                for fi37 := 0; fi37 < 11; fi37++ {
                    v24 = v24
                    v36 = (v13 & v13)
                    if (1 > v24) {
                        v14 += 65535
                        for fi38 := 0; fi38 < 12; fi38++ {
                            v13 += (v13 ^ (v36 & (v36 % ((0) | 1))))
                            _ = fi38
                        }
                    }
                    _ = fi37
                }
                v24 += v24
                _ = v36
            } else {
                v39 := GenS3{F0: 128, F1: -114, F2: 2}
                var v40 uint64 = 4294967295
                var v41 uint16 = v14
                _ = v39
                _ = v40
                _ = v41
            }
        case 2147483647:
            var v42 uint16 = v14
            _ = v42
        }
        v13 = v13
    } else {
        var v43 uint64 = 487
        v44 := GenS2{F0: -933, F1: (v14 - v14), F2: false}
        v13 = (v13 % ((0) | 1))
        _ = v43
        _ = v44
    }
    fmt.Println(v2)
    for fi45 := 0; fi45 < 8; fi45++ {
        _ = len(v1)
        fmt.Println(fi45)
        var v46 int = (^(fi45))
        _ = fi45
        _ = v46
    }
    var v47 uint64 = 18446744073709551615
    fmt.Println(v47, v14, v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
