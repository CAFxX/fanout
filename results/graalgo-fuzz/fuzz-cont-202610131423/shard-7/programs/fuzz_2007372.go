package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int; F2 uint64 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 int16; F1 uint; F2 uint32 }
type GenS6 struct { F0 int32; F1 int16; F2 int64 }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 uint16 }

func main() {
    v1 := []string{"hello", "a"}
    v2 := make(map[uint32]int)
    var v3 int = -573
    var v4 int = (^(v3))
    v5 := make(map[int]int)
    v6 := make(map[uint32]uint32)
    if false {
        var v7 int8 = 0
        var v8 uint8 = 154
        var v9 int16 = -1
        _ = v7
        _ = v8
        _ = v9
    }
    if (false && true) {
        if true {
            v4 -= v3
            var v10 int16 = -493
            _ = v10
        }
        v4 += -9223372036854775808
        if (!(0.1 > (3.14159 / (-2.25 / 1.5)))) {
            v11 := GenS1{F0: true}
            v12 := GenS2{F0: 65535}
            v13 := GenS3{F0: "", F1: (v3 - 2147483647), F2: 18446744073709551615}
            _ = v11
            _ = v12
            _ = v13
        } else {
            switch "a" {
            case "a":
                if (-1 <= 967) {
                    delete(v6, 35)
                    var v14 int = (^(v3))
                    _ = v14
                } else {
                    fmt.Println(v3, v3)
                    var v15 int16 = -32768
                    for fi16 := 0; fi16 < 3; fi16++ {
                        switch "" {
                        case "a":
                            var v17 int = v4
                            _ = v17
                        case "foo bar":
                            var v18 uint8 = 255
                            _ = v18
                        }
                        v4 = 239
                        v19 := GenS1{F0: (113 == 0)}
                        _ = fi16
                        _ = v19
                    }
                    _ = v15
                }
            case "hello":
                v20 := GenS3{F0: "foo bar", F1: (^(v3)), F2: 18446744073709551615}
                _ = v20
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                v4 += v3
            }
            fmt.Println(v4)
        }
    } else {
        var v21 int64 = 0
        _ = v21
    }
    var v22 uint64 = 4294967295
    v22 += v22
    switch "foo bar" {
    case "a":
        v3 = v4
    case "hello":
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    default:
        var v23 int64 = 827
        _ = v23
    }
    fmt.Println(v3)
    v4 = -2147483648
    if false {
        for fi24 := 0; fi24 < 1; fi24++ {
            if (13 > -1) {
                var v25 uint = 212
                var v26 int8 = 100
                if (((-998 >= 736) || ("foo bar" < ("\n\t" + ((("hello" + "a") + "hello") + "\n\t")))) || (false && false)) {
                    fmt.Println(v22, v4)
                }
                _ = v25
                _ = v26
            } else {
                v4 = fi24
            }
            v27 := GenS4{F0: 18446744073709551615, F1: 1}
            _ = fi24
            _ = v27
        }
        v28 := GenS5{F0: -814, F1: 65535, F2: 4294967295}
        v29 := GenS6{F0: -1, F1: 475, F2: -9223372036854775808}
        _ = v28
        _ = v29
    }
    v30 := GenS7{F0: "hello", F1: 1}
    for fi31 := 0; fi31 < 10; fi31++ {
        var v32 uint32 = 447
        v22 += 42
        v33 := GenS8{F0: 2}
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
