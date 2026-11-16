package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int; F2 int64 }
type GenS3 struct { F0 uint32 }
func genFn0() int64 {
    var v5 int64 = 1
    v6 := []int{}
    if (4294967295 != 2) {
        for fi7 := 0; fi7 < 8; fi7++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi7
        }
    }
    var v8 uint8 = 25
    _ = v5
    _ = v6
    _ = v8
    return v5
}
func genFn1(p0 string, p1 uint8) int16 {
    var v9 int8 = -69
    v10 := GenS1{F0: 1}
    p0 = (p0 + p0)
    var v11 uint32 = 604
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 32767
}

func main() {
    v1 := []uint32{429, 2147483648}
    v2 := []uint32{0}
    v3 := []string{("" + ""), "hello", "\n\t"}
    var v4 uint16 = 255
    switch "\n\t" {
    case "a":
        switch 249 {
        case 42:
            switch -377 {
            case -9223372036854775808:
                v4 = (v4 * (^(v4)))
            case 32767:
                var v12 int16 = -436
                _ = v12
            default:
                v4 += v4
            }
        }
    case "\n\t":
        fmt.Println(v4)
    case "hello":
        v13 := GenS1{F0: 1}
        _ = v13
    }
    if (false || (!(83 != 71))) {
        var v14 uint64 = 42
        v15 := GenS2{F0: -1, F1: 128, F2: 1}
        var v16 uint64 = (v14 << ((0) & 63))
        _ = v14
        _ = v15
        _ = v16
    }
    v4 = (v4 - (v4 | v4))
    if (false || (!(!true))) {
        v4 += (v4 * (v4 ^ v4))
        if true {
            var v17 uint8 = 255
            v17 += 141
            _ = v17
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch 32768 {
            case -999:
                if false {
                    var v18 float64 = -2.25
                    _ = v18
                } else {
                    if len(v1) > 0 {
                        v1[0] = 951
                    }
                    v4 = (v4 / ((v4) | 1))
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                }
            case 128:
                if len(v1) > 0 {
                    v1[0] = 42
                }
            default:
                var v19 uint64 = 617
                _ = v19
            }
        }
    }
    v20 := GenS3{F0: 2147483647}
    if len(v2) > 0 {
        v2[0] = 429
    }
    v4 = v4
    if (4294967295 >= 42) {
        var v21 uint64 = 2
        if true {
            for fi22 := 0; fi22 < 14; fi22++ {
                if len(v3) > 0 {
                    v3[0] = ""
                }
                switch ("a" + "hello") {
                case "":
                    if len(v1) > 0 {
                        v1[0] = 287
                    }
                case "hello":
                    var v23 uint8 = 255
                    _ = v23
                case "\n\t":
                    v24 := GenS1{F0: -384}
                    _ = v24
                default:
                    v21 = v21
                }
                _ = fi22
            }
        }
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
