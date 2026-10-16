package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int16; F1 uint32; F2 int }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int64; F1 int8 }

func main() {
    var v1 int8 = -128
    v2 := []uint32{936, 858, 322, 4294967295}
    var v3 uint64 = 18446744073709551615
    switch ("a" + "hello") {
    case "\n\t":
        v4 := GenS1{F0: "foo bar", F1: 18446744073709551615}
        _ = v4
    case "foo bar":
        var v5 uint16 = 1
        _ = v5
    default:
        _ = len(v2)
    }
    if (false && (false && (2 < 0))) {
        var v6 float32 = (3.14159 * -2.25)
        if true {
            fmt.Println(v3, v1)
            switch "" {
            case "":
                var v7 string = "foo bar"
                _ = v7
            case "foo bar":
                v8 := GenS2{F0: 31, F1: uint8(v3), F2: 0}
                _ = v8
            case "hello":
                switch "a" {
                case "":
                    if false {
                        v3 = v3
                    }
                case "a":
                    v6 += v6
                default:
                    v9 := GenS1{F0: ("" + "hello"), F1: 255}
                    _ = v9
                }
            default:
                if len(v2) > 0 {
                    v2[0] = 42
                }
            }
        } else {
            if true {
                v3 = v3
                v10 := GenS3{F0: 1000, F1: 126, F2: 552}
                _ = v10
            } else {
                if false {
                    v11 := GenS4{F0: 49, F1: 2147483648}
                    _ = v11
                }
            }
            for fi12 := 0; fi12 < 19; fi12++ {
                v1 += (-(v1))
                _ = fi12
            }
            if true {
                v6 = v6
            }
        }
        for fi13 := 0; fi13 < 18; fi13++ {
            if true {
                var v14 float32 = v6
                _ = v14
            }
            _ = fi13
        }
        _ = v6
    } else {
        v1 -= (v1 << ((1) & 63))
        v1 = v1
        var v15 int8 = (v1 - v1)
        _ = v15
    }
    if true {
        for fi16 := 0; fi16 < 6; fi16++ {
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 int = (fi16 | (fi16 - (-(fi17))))
                v19 := GenS4{F0: 255, F1: 776}
                var v20 int16 = 32767
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            v3 += v3
            _ = fi16
        }
        for fi21 := 0; fi21 < 20; fi21++ {
            if (42 >= 552) {
                v1 -= (^(v1))
            }
            var v22 uint8 = 128
            if len(v2) > 0 {
                v2[0] = 1
            }
            _ = fi21
            _ = v22
        }
        fmt.Println(v1, v3)
    }
    v23 := GenS2{F0: -1, F1: 255, F2: 177}
    var v24 int8 = v1
    var v25 string = ("" + "")
    v1 += -34
    v25 = (v25 + ((v25 + "") + v25))
    var v26 int64 = 9223372036854775807
    if len(v2) > 0 {
        v2[0] = 65535
    }
    v2 = append(v2, 1)
    if true {
        var v27 int8 = 100
        _ = v27
    } else {
        v28 := GenS5{F0: (-(v26)), F1: int8(v3)}
        fmt.Println(v25, v24)
        fmt.Println(v24, v25)
        _ = v28
    }
    fmt.Println(v25)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
