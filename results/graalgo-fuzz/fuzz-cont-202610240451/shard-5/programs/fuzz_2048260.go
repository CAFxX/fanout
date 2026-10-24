package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int64; F2 int }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    v3 := []string{"hello", ("hello" + ("\n\t" + "\n\t")), "hello"}
    var v4 int8 = 18
    v5 := GenS1{F0: 434}
    v4 = 1
    switch 458 {
    case 9223372036854775807:
        v4 = -128
    case -502:
        var v6 uint8 = 42
        _ = v6
    }
    switch "a" {
    case "\n\t":
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := GenS1{F0: (fi7 ^ 32768)}
            v4 = -50
            var v9 uint32 = 42
            _ = fi7
            _ = v8
            _ = v9
        }
    case "a":
        switch "a" {
        case "":
            if true {
                v4 += 54
                var v10 uint32 = 2147483647
                _ = v10
            }
        case "foo bar":
            v4 = 100
        case "\n\t":
            v4 = (^(v4))
        default:
            for fi11 := 0; fi11 < 18; fi11++ {
                if true {
                    switch fi11 {
                    case 42:
                        fmt.Println(v4, v4)
                    }
                    fmt.Println(fi11, v4)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                if (-81 >= -58) {
                    v12 := GenS1{F0: 32768}
                    v4 = v4
                    _ = v12
                }
                if false {
                    if false {
                        v4 = (v4 >> ((255) & 63))
                        v4 -= v4
                        switch 128 {
                        case 32767:
                            fmt.Println(fi11)
                        }
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    if (1e10 >= (((float64(0.1) - (1.5 / (1.5 * (3.14159 * 1e10)))) / 1e10) / (1.5 + 3.14159))) {
                        v4 += (v4 ^ v4)
                    } else {
                        var v13 float64 = 3.14159
                        _ = v13
                    }
                }
                _ = fi11
            }
        }
    case "hello":
        v14 := GenS2{F0: 42, F1: 258, F2: 776}
        _ = v14
    default:
        var v15 uint32 = 388
        _ = v15
    }
    switch -9223372036854775808 {
    case 284:
        if len(v3) > 0 {
            v3[0] = ("hello" + (("hello" + "foo bar") + ("hello" + "a")))
        }
    case -1:
        if (!(!(-668 < -9223372036854775808))) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 int8 = (v4 | (v4 / ((v4) | 1)))
                fmt.Println(fi16, v17, fi16)
                _ = fi16
                _ = v17
            }
        } else {
            fmt.Println(v4)
            if true {
                fmt.Println(v4)
            }
            for fi18 := 0; fi18 < 17; fi18++ {
                var v19 int16 = 0
                v4 = (^(v4))
                _ = fi18
                _ = v19
            }
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    delete(v1, 65535)
    v20 := GenS1{F0: 187}
    v21 := v2[-2147483648]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
