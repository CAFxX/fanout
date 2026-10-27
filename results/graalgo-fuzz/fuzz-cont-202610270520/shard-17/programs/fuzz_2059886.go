package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint) int64 {
    p0 = (p0 - p0)
    fmt.Println(p0)
    for fi5 := 0; fi5 < 12; fi5++ {
        if false {
            p0 -= (p0 - p0)
            var v6 uint32 = 65535
            _ = v6
        } else {
            fmt.Println(p0, p0)
            var v7 uint32 = 4294967295
            v7 -= v7
            _ = v7
        }
        switch "" {
        case "foo bar":
            p0 -= (p0 << ((905) & 63))
        }
        _ = fi5
    }
    p0 = (p0 - 18446744073709551615)
    _ = p0
    return 1
}

func main() {
    var v1 uint8 = 42
    var v2 uint = 2
    var v3 int32 = 242
    var v4 uint8 = (v1 >> ((v2) & 63))
    v8 := []int{21, -76, 1}
    v1 = (^(v4))
    v9 := GenS1{F0: "hello", F1: 111, F2: genFn0((v2 >> ((v2) & 63)))}
    var v10 uint32 = 221
    switch "hello" {
    case "hello":
        if false {
            if (true && false) {
                var v11 int = -9223372036854775808
                v12 := GenS2{F0: (v1 * 128)}
                if false {
                    v13 := GenS3{F0: false}
                    _ = v13
                }
                _ = v11
                _ = v12
            }
            switch -329 {
            case 9223372036854775807:
                if (v2 == (v2 - uint(v1))) {
                    if false {
                        switch "" {
                        case "\n\t":
                            var v14 uint64 = 4294967295
                            _ = v14
                        case "":
                            var v15 int64 = 0
                            _ = v15
                        case "hello":
                            v16 := GenS2{F0: (v4 % ((v1) | 1))}
                            _ = v16
                        }
                    } else {
                        var v17 int16 = -1
                        v10 += v10
                        _ = v17
                    }
                }
            case -1:
                if true {
                    var v18 uint32 = uint32(v1)
                    _ = v18
                } else {
                    if len(v8) > 0 {
                        v8[0] = 587
                    }
                    v8 = append(v8, 32767)
                }
            default:
                v1 = v4
            }
            v8 = append(v8, 127)
        }
    case "\n\t":
        fmt.Println(v4, v1)
    case "a":
        var v19 uint8 = v4
        _ = v19
    default:
        v4 = (v1 & v1)
    }
    switch (((("\n\t" + "") + "hello") + "\n\t") + ("\n\t" + "")) {
    case "":
        v20 := GenS1{F0: "\n\t", F1: -128, F2: 328}
        _ = v20
    case "\n\t":
        var v21 bool = (-94 < 0)
        _ = v21
    default:
        switch "foo bar" {
        case "hello":
            v22 := GenS3{F0: (false || (0 == 42))}
            _ = v22
        }
    }
    if true {
        var v23 int16 = -684
        _ = v23
    }
    v24 := GenS3{F0: ((!false) && false)}
    v10 += (v10 << (((^(v2))) & 63))
    v3 += v3
    v25 := GenS4{F0: -1}
    v2 += 2
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 float32 = -2.25
        _ = fi26
        _ = v27
    }
    fmt.Println(v1, v4, v2)
    v10 -= 384
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
