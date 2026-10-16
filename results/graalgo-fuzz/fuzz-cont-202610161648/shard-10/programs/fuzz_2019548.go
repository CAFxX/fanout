package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int16; F1 int }

func main() {
    v1 := []uint32{42}
    var v2 uint64 = 0
    var v3 uint16 = 42
    var v4 uint64 = v2
    var v5 int32 = -229
    v5 -= -2147483648
    v3 -= (v3 % (((v3 | uint16(v4))) | 1))
    switch -2147483648 {
    case 32767:
        switch int(v3) {
        case 0:
            var v6 int32 = 797
            _ = v6
        case -745:
            fmt.Println(v3, v5, v4)
        case 2147483647:
            v2 = (v4 + (^(v4)))
        default:
            v5 -= (v5 + (v5 ^ -329))
        }
    case 127:
        for fi7 := 0; fi7 < 10; fi7++ {
            _ = len(v1)
            v8 := GenS1{F0: 776}
            _ = fi7
            _ = v8
        }
    }
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        fmt.Println(v4)
    case "":
        if (!(0.0 == 0.1)) {
            for fi9 := 0; fi9 < 15; fi9++ {
                v10 := GenS1{F0: 392}
                if true {
                    var v11 int32 = v5
                    v12 := GenS2{F0: 405}
                    _ = v11
                    _ = v12
                }
                _ = fi9
                _ = v10
            }
            if false {
                if true {
                    v1 = append(v1, 2147483648)
                }
            } else {
                var v13 float64 = 1e10
                var v14 uint = 1
                v15 := GenS2{F0: 238}
                _ = v13
                _ = v14
                _ = v15
            }
        } else {
            if true {
                if true {
                    switch 32768 {
                    case 4294967295:
                        v5 -= v5
                    case 202:
                        v3 = v3
                    }
                    switch "hello" {
                    case "foo bar":
                        v3 = v3
                    }
                }
            }
            if true {
                if (((v3 << ((588) & 63)) >= v3) && (!true)) {
                    v16 := GenS2{F0: 65535}
                    v2 += v2
                    if (!(((v4 << ((488) & 63)) >= (v2 + 489)) || true)) {
                        var v17 float32 = 1.5
                        var v18 uint32 = 972
                        fmt.Println(v18)
                        _ = v17
                        _ = v18
                    }
                    _ = v16
                }
                v19 := GenS3{F0: ("foo bar" >= "hello"), F1: -152}
                _ = v19
            }
        }
    }
    v2 = 2
    v20 := GenS4{F0: 717, F1: 0}
    var v21 string = "a"
    switch v21 {
    case "hello":
        var v22 float32 = (1e10 + 0.0)
        _ = v22
    default:
        for fi23 := 0; fi23 < 4; fi23++ {
            v2 -= v2
            var v24 int8 = -1
            _ = fi23
            _ = v24
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
