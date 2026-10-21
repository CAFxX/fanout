package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int }
type GenS2 struct { F0 uint; F1 uint64; F2 int64 }
type GenS3 struct { F0 int16; F1 int8; F2 int32 }
type GenS4 struct { F0 string }
func genFn0(p0 uint8, p1 int) int16 {
    v4 := []int{(p1 & 197), (p1 & 32767)}
    v4 = append(v4, 54)
    _ = p0
    _ = p1
    _ = v4
    return 608
}
func genFn1(p0 string, p1 int8, p2 uint8) string {
    v5 := GenS1{F0: 0, F1: -1, F2: 0}
    v6 := GenS1{F0: 2, F1: 0, F2: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p0
}

func main() {
    var v1 float32 = (3.14159 + -2.25)
    var v2 uint64 = 0
    v3 := make(map[uint32]int)
    switch 2147483647 {
    case 1:
        v7 := []int{157, 515}
        _ = v7
    default:
        if false {
            if false {
                v2 += (v2 >> ((2) & 63))
            }
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 float32 = (v1 / v1)
                _ = fi8
                _ = v9
            }
            v1 -= v1
        } else {
            switch 1 {
            case -9223372036854775808:
                for fi10 := 0; fi10 < 7; fi10++ {
                    v2 = (v2 - 4294967295)
                    switch 2147483647 {
                    case 32768:
                        var v11 uint = 255
                        _ = v11
                    case -542:
                        v12 := GenS1{F0: 368, F1: 2147483647, F2: (fi10 & (fi10 & fi10))}
                        _ = v12
                    }
                    _ = fi10
                }
            case -869:
                var v13 float64 = 0.1
                _ = v13
            case 4294967295:
                v2 = (^(v2))
            default:
                v2 = v2
            }
            switch -577 {
            case 128:
                v2 += (^(v2))
            }
            var v14 uint64 = 0
            _ = v14
        }
    }
    v2 = (v2 << ((2) & 63))
    switch 128 {
    case 97:
        var v15 uint16 = 42
        _ = v15
    case 127:
        v1 = 0.0
    default:
        for fi16 := 0; fi16 < 20; fi16++ {
            switch "a" {
            case "a":
                switch "" {
                case "":
                    for fi17 := 0; fi17 < 6; fi17++ {
                        var v18 int = fi16
                        v19 := GenS2{F0: 4294967295, F1: 0, F2: -1}
                        var v20 string = ""
                        var v21 uint32 = 914
                        v1 = v1
                        var v22 uint16 = 255
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                case "foo bar":
                    var v23 int16 = 32767
                    _ = v23
                case "a":
                    v2 = (v2 ^ (^(v2)))
                }
            }
            _ = fi16
        }
    }
    v24 := GenS2{F0: 65535, F1: 1, F2: 42}
    v1 = (((v1 + v1) * 0.1) * (v1 - 0.0))
    var v25 int32 = 0
    v26 := GenS3{F0: 135, F1: 42, F2: 2147483647}
    var v27 uint16 = 1
    switch 10 {
    case 128:
        v28 := GenS4{F0: ""}
        _ = v28
    }
    switch ("foo bar" + (("a" + "") + "a")) {
    case "hello":
        fmt.Println(v1)
    case "":
        switch "a" {
        case "\n\t":
            v27 = v27
        case "a":
            var v29 uint = 18446744073709551615
            _ = v29
        }
    }
    if true {
        var v30 int = int(v25)
        _ = v30
    } else {
        var v31 int16 = 1
        v2 = (^(v2))
        v3[321] = 240
        _ = v31
    }
    v32 := GenS3{F0: int16(v25), F1: 42, F2: (v25 + (v25 | 2147483647))}
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 int32 = v25
        var v35 float64 = 0.0
        v25 -= (v25 * (v34 + v34))
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
}
