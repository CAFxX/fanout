package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint64; F2 uint16 }

func main() {
    var v1 uint32 = 42
    v2 := []uint32{v1, 42, 65535}
    var v3 int16 = -32768
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 int32 = 0
        if false {
            switch fi4 {
            case -407:
                v6 := GenS1{F0: 18446744073709551615, F1: 1}
                _ = v6
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            fmt.Println(fi4, v3)
        } else {
            v3 -= 0
            var v7 uint16 = 334
            _ = v7
        }
        v1 = v1
        _ = fi4
        _ = v5
    }
    var v8 int32 = 2147483647
    v9 := GenS2{F0: (!true)}
    v3 = -662
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint64 = 538
        _ = fi10
        _ = v11
    }
    fmt.Println(v1, v8, v3)
    v1 = (^(v1))
    switch "\n\t" {
    case "":
        if (false && true) {
            switch "hello" {
            case "foo bar":
                var v12 uint16 = 255
                _ = v12
            case "":
                v2 = append(v2, (^(v1)))
            default:
                var v13 int64 = -465
                _ = v13
            }
        } else {
            var v14 float64 = 1e10
            fmt.Println(v8, v1)
            _ = v14
        }
    case "foo bar":
        var v15 uint = 646
        _ = v15
    default:
        if (!false) {
            var v16 uint8 = 127
            v2 = append(v2, 65535)
            switch 471 {
            case -555:
                v17 := GenS1{F0: uint(v1), F1: (v16 ^ (v16 & 192))}
                _ = v17
            default:
                var v18 int = -329
                _ = v18
            }
            _ = v16
        } else {
            var v19 int64 = -9223372036854775808
            switch "a" {
            case "hello":
                v20 := GenS3{F0: 9223372036854775807, F1: 356, F2: 687}
                _ = v20
            }
            _ = v19
        }
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        if len(v2) > 0 {
            v2[0] = v1
        }
        _ = fi21
    }
    var v22 int8 = 26
    v22 = (^(v22))
    v2 = append(v2, v1)
    v23 := GenS1{F0: 109, F1: 42}
    v24 := GenS3{F0: 1, F1: 0, F2: 195}
    for fi25 := 0; fi25 < 17; fi25++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v8 -= (v8 & (v8 << ((483) & 63)))
        _ = fi25
    }
    var v26 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v8
    _ = v9
}
