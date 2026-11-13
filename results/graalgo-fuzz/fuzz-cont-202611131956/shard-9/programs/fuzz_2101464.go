package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }

func main() {
    var v1 uint8 = 128
    v2 := []uint32{465, 2147483647}
    var v3 int8 = 100
    v4 := []string{("a" + ""), "a"}
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 15; fi6++ {
            switch fi6 {
            case 128:
                v3 = 127
                var v7 float32 = float32(-2.25)
                var v8 uint = 610
                _ = v7
                _ = v8
            case 1:
                v9 := GenS1{F0: -553, F1: -9223372036854775808}
                _ = v9
            default:
                v3 -= 1
            }
            _ = fi6
        }
        for fi10 := 0; fi10 < 10; fi10++ {
            switch "hello" {
            case "\n\t":
                v2 = append(v2, 65535)
            case "":
                if (-523 != 687) {
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(fi5, v1, fi10)
                    default:
                        var v11 uint = 255
                        _ = v11
                    }
                    v2 = append(v2, 677)
                } else {
                    var v12 uint16 = 42
                    _ = v12
                }
            }
            var v13 uint64 = 1
            var v14 float64 = 0.0
            _ = fi10
            _ = v13
            _ = v14
        }
        var v15 uint = 255
        _ = fi5
        _ = v15
    }
    v16 := GenS1{F0: 42, F1: -765}
    if true {
        fmt.Println(v1)
        for fi17 := 0; fi17 < 16; fi17++ {
            if false {
                v18 := GenS1{F0: 32767, F1: fi17}
                v3 = 127
                fmt.Println(fi17, v3, fi17)
                _ = v18
            } else {
                v19 := GenS1{F0: 1000, F1: (fi17 - fi17)}
                v3 += 127
                _ = v19
            }
            _ = fi17
        }
        for fi20 := 0; fi20 < 12; fi20++ {
            fmt.Println(fi20, v3, v1)
            var v21 int32 = 2147483647
            v3 = (v3 & 100)
            _ = fi20
            _ = v21
        }
    }
    var v22 int8 = v3
    for fi23 := 0; fi23 < 9; fi23++ {
        var v24 float32 = 1e10
        _ = fi23
        _ = v24
    }
    var v25 int64 = -1
    if true {
        v1 = v1
        v22 -= (v3 - v22)
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        for fi27 := 0; fi27 < 10; fi27++ {
            v1 = (v1 ^ 127)
            _ = fi27
        }
        for fi28 := 0; fi28 < 10; fi28++ {
            var v29 uint64 = 18446744073709551615
            v22 -= (v3 - (v22 << ((997) & 63)))
            var v30 float64 = float64(0.1)
            _ = fi28
            _ = v29
            _ = v30
        }
        if false {
            fmt.Println(fi26, fi26)
        } else {
            switch "foo bar" {
            case "foo bar":
                switch 847 {
                case 128:
                    v3 -= (v22 / ((-128) | 1))
                case -984:
                    if false {
                        for fi31 := 0; fi31 < 1; fi31++ {
                            var v32 int32 = -508
                            var v33 uint16 = 0
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                    }
                case -740:
                    fmt.Println(v3)
                }
            case "a":
                v3 += v22
            default:
                _ = len(v4)
            }
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
}
