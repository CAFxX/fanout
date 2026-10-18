package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 uint16 = 1
    var v2 int32 = 42
    var v3 bool = true
    var v4 float64 = 1.5
    if v3 {
        var v5 int64 = 0
        v6 := GenS1{F0: v5, F1: "\n\t"}
        v1 += v1
        _ = v5
        _ = v6
    } else {
        var v7 int = -179
        for fi8 := 0; fi8 < 12; fi8++ {
            for fi9 := 0; fi9 < 15; fi9++ {
                v3 = v3
                _ = fi9
            }
            _ = fi8
        }
        if true {
            v1 -= 65535
        } else {
            var v10 int32 = v2
            var v11 int64 = 42
            _ = v10
            _ = v11
        }
        _ = v7
    }
    if (32768 > -322) {
        var v12 string = "foo bar"
        v13 := GenS1{F0: 42, F1: (v12 + v12)}
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            fmt.Println(fi15, fi14, v3)
            v3 = v3
            _ = fi15
        }
        _ = fi14
    }
    v16 := []int{}
    v2 += v2
    if (!v3) {
        for fi17 := 0; fi17 < 4; fi17++ {
            v2 -= (v2 * v2)
            _ = fi17
        }
    } else {
        fmt.Println(v4)
        v3 = true
    }
    v18 := GenS1{F0: 9223372036854775807, F1: ""}
    if v3 {
        fmt.Println(v2, v1, v3)
    }
    switch -2147483648 {
    case -2147483648:
        v19 := GenS1{F0: 790, F1: "\n\t"}
        _ = v19
    default:
        switch 4294967295 {
        case -9223372036854775808:
            if (1 != 127) {
                for fi20 := 0; fi20 < 14; fi20++ {
                    fmt.Println(v1, v3)
                    v1 = 32767
                    var v21 uint64 = 2
                    _ = fi20
                    _ = v21
                }
                var v22 float32 = 1e10
                _ = v22
            } else {
                _ = len(v16)
                var v23 uint32 = 42
                for fi24 := 0; fi24 < 8; fi24++ {
                    if v3 {
                        v2 = v2
                        v4 = v4
                    }
                    var v25 float32 = (-2.25 + 1e10)
                    _ = fi24
                    _ = v25
                }
                _ = v23
            }
        }
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        for fi27 := 0; fi27 < 12; fi27++ {
            var v28 int8 = -89
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 6; fi29++ {
            v30 := GenS1{F0: 1, F1: "a"}
            _ = fi29
            _ = v30
        }
        if (!(!v3)) {
            fmt.Println(v3, fi26)
            v31 := GenS1{F0: 939, F1: ("" + "foo bar")}
            fmt.Println(fi26, v3)
            _ = v31
        }
        _ = fi26
    }
    if true {
        if len(v16) > 0 {
            v16[0] = -9223372036854775808
        }
        for fi32 := 0; fi32 < 1; fi32++ {
            v33 := GenS1{F0: 253, F1: "a"}
            v2 -= v2
            v16 = append(v16, fi32)
            _ = fi32
            _ = v33
        }
    } else {
        if (!(295 != 65535)) {
            for fi34 := 0; fi34 < 20; fi34++ {
                v35 := GenS2{F0: 42, F1: 0}
                fmt.Println(v2)
                switch "foo bar" {
                case "hello":
                    if v3 {
                        var v36 int8 = 100
                        v37 := GenS1{F0: -637, F1: ("hello" + "a")}
                        _ = v36
                        _ = v37
                    }
                case "\n\t":
                    var v38 uint = 65535
                    _ = v38
                case "":
                    v2 = (v2 + v2)
                default:
                    fmt.Println(v1)
                }
                _ = fi34
                _ = v35
            }
        }
        v2 -= 126
    }
    var v39 uint = 307
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v39
    _ = v4
}
