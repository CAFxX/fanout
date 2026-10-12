package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint }
type GenS4 struct { F0 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = 29
    v3 := []string{("" + ("a" + "")), "\n\t", "hello"}
    v4 := []string{}
    var v5 uint8 = 128
    v6 := make(map[int]string)
    v2 -= (v2 | v2)
    v7 := GenS1{F0: 255, F1: 1}
    v5 = v5
    v5 -= 55
    var v8 int8 = (v2 + v2)
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS2{F0: 1, F1: -336}
        if true {
            v11 := GenS3{F0: (!false), F1: 721, F2: 965}
            _ = v11
        }
        v12 := GenS4{F0: 65535}
        _ = fi9
        _ = v10
        _ = v12
    }
    v5 -= (v5 * 127)
    v8 += (v8 >> ((255) & 63))
    if false {
        switch "foo bar" {
        case "":
            delete(v6, -304)
        case "a":
            if false {
                switch 999 {
                case -1:
                    v13 := v6[1]
                    _ = v13
                case 32767:
                    v14 := GenS1{F0: 558, F1: 18446744073709551615}
                    _ = v14
                case -9223372036854775808:
                    var v15 string = ("a" + "foo bar")
                    _ = v15
                default:
                    v8 = (v2 + v2)
                }
                fmt.Println(v8, v2)
            }
        }
        var v16 uint32 = 42
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            if false {
                if false {
                    var v19 int8 = 0
                    v20 := GenS2{F0: v5, F1: -71}
                    v5 = (v5 + (v5 + (v5 ^ 2)))
                    _ = v19
                    _ = v20
                } else {
                    if ((v2 << ((433) & 63)) < (v8 + v2)) {
                        v2 += 1
                        var v21 int8 = v2
                        _ = v21
                    }
                }
            } else {
                v22 := GenS4{F0: 65535}
                v1[((("a" + "foo bar") + "hello") + ("" + (("\n\t" + "hello") + "a")))] = true
                _ = v22
            }
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            fmt.Println(v2, v8)
            delete(v1, "hello")
            _ = fi18
        }
        var v23 int8 = (v8 ^ 42)
        fmt.Println(v23, v8)
        _ = fi17
        _ = v23
    }
    var v24 uint16 = 2
    var v25 uint8 = 112
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
