package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 string; F1 uint }

func main() {
    var v1 uint = 0
    var v2 uint16 = 154
    var v3 float64 = float64(((1.5 * 1.5) / -2.25))
    v4 := []int{32767, 2147483647, -860}
    var v5 uint64 = 147
    switch -570 {
    case -451:
        v2 = (v2 & (v2 % (((v2 / ((723) | 1))) | 1)))
    }
    v6 := GenS1{F0: v2, F1: 914}
    fmt.Println(v1)
    if (false && (!false)) {
        v5 = (v5 ^ v5)
        _ = len(v4)
        if ((!true) && true) {
            fmt.Println(v2, v5)
        } else {
            var v7 int8 = 0
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 int16 = -32768
                for fi10 := 0; fi10 < 7; fi10++ {
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(fi8, fi10, v7)
                    case "foo bar":
                        v2 += (^(v2))
                    }
                    v11 := GenS2{F0: 541, F1: (^(v5))}
                    fmt.Println(v9, v7)
                    _ = fi10
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
            if ((v7 % ((v7) | 1)) > 42) {
                v12 := GenS1{F0: (^(v2)), F1: -2}
                for fi13 := 0; fi13 < 5; fi13++ {
                    fmt.Println(v7)
                    if len(v4) > 0 {
                        v4[0] = fi13
                    }
                    switch fi13 {
                    case 128:
                        v14 := GenS3{F0: "hello", F1: v1}
                        _ = v14
                    case 32768:
                        for fi15 := 0; fi15 < 1; fi15++ {
                            var v16 string = "hello"
                            _ = fi15
                            _ = v16
                        }
                    }
                    _ = fi13
                }
                _ = v12
            } else {
                var v17 float32 = 1.5
                v17 = v17
                _ = v17
            }
            _ = v7
        }
    } else {
        v4 = append(v4, 128)
        v1 = v1
    }
    if true {
        var v18 int64 = -348
        _ = v18
    } else {
        fmt.Println(v1, v2, v5)
    }
    if true {
        _ = len(v4)
        if len(v4) > 0 {
            v4[0] = 128
        }
    }
    var v19 float32 = 0.0
    if false {
        var v20 uint64 = 915
        _ = v20
    } else {
        v4 = append(v4, 42)
    }
    v2 -= 65535
    for fi21 := 0; fi21 < 9; fi21++ {
        if ((978 >= 1) && ((1 != 388) && true)) {
            v19 = v19
        }
        _ = fi21
    }
    switch 127 {
    case 9223372036854775807:
        v22 := GenS1{F0: 2, F1: int16(v5)}
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
