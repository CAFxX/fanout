package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 bool; F2 uint }

func main() {
    var v1 int64 = 1
    var v2 uint8 = 2
    var v3 int8 = -4
    v4 := make(map[int]bool)
    v5 := make(map[string]string)
    var v6 uint = 1
    v7 := GenS1{F0: v3}
    var v8 uint8 = (^(v2))
    v9 := GenS2{F0: 0, F1: -2147483648, F2: 65535}
    var v10 uint16 = 65535
    for fi11 := 0; fi11 < 10; fi11++ {
        fmt.Println(v6, v3)
        _ = fi11
    }
    v12 := GenS3{F0: (^(v1)), F1: 18446744073709551615}
    v13 := GenS3{F0: (v1 % ((v1) | 1)), F1: 4294967295}
    if (1 > v3) {
        var v14 uint64 = 2
        delete(v4, 9223372036854775807)
        v4[-1] = (false || true)
        _ = v14
    }
    v15 := GenS4{F0: 42, F1: 2147483648}
    v16 := []int{510, -2147483648}
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 uint16 = uint16(v6)
        if len(v16) > 0 {
            v16[0] = (fi17 % (((^(fi17))) | 1))
        }
        v19 := GenS3{F0: (v1 << ((341) & 63)), F1: 0}
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(v3, v3, v1)
    switch ("a" + "hello") {
    case "a":
        v8 -= v2
    case "\n\t":
        delete(v5, "a")
    case "hello":
        var v20 int16 = 1
        _ = v20
    }
    var v21 uint8 = (^(v2))
    for fi22 := 0; fi22 < 3; fi22++ {
        v3 = (v3 % ((int8(v21)) | 1))
        _ = fi22
    }
    v23 := v4[4294967295]
    if v23 {
        if false {
            if false {
                for fi24 := 0; fi24 < 1; fi24++ {
                    var v25 uint8 = (v8 << (((v6 ^ v6)) & 63))
                    _ = fi24
                    _ = v25
                }
            } else {
                v26 := GenS5{F0: 42}
                _ = v26
            }
        } else {
            var v27 uint8 = (v2 * (v2 + (v8 % ((v21) | 1))))
            var v28 uint = (^(v6))
            switch "foo bar" {
            case "hello":
                for fi29 := 0; fi29 < 14; fi29++ {
                    v30 := GenS6{F0: (-(v3)), F1: v23, F2: 0}
                    switch "" {
                    case "a":
                        var v31 int8 = (-(v3))
                        _ = v31
                    case "foo bar":
                        v28 -= v6
                    default:
                        v27 = v27
                    }
                    var v32 int16 = 565
                    _ = fi29
                    _ = v30
                    _ = v32
                }
            default:
                for fi33 := 0; fi33 < 18; fi33++ {
                    fmt.Println(v3, v27)
                    v21 = 249
                    v23 = false
                    _ = fi33
                }
            }
            _ = v27
            _ = v28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
