package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 int8 = 1
    v2 := make(map[uint32]int)
    var v3 int32 = 42
    var v4 int64 = 9223372036854775807
    var v5 uint8 = 7
    var v6 bool = (!false)
    var v7 int32 = 1
    v7 -= v3
    switch "foo bar" {
    case "":
        v8 := GenS1{F0: 2, F1: v6}
        _ = v8
    case "hello":
        v3 = (^(v3))
    default:
        v9 := []int{32767, 32767, 603, 32768}
        _ = v9
    }
    v5 -= v5
    if (!v6) {
        v10 := GenS2{F0: 42, F1: 1, F2: v5}
        if false {
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 string = "hello"
                v4 = (v4 / ((v4) | 1))
                _ = fi11
                _ = v12
            }
        }
        v13 := GenS1{F0: 4294967295, F1: v6}
        _ = v10
        _ = v13
    }
    v4 = int64(v3)
    if false {
        fmt.Println(v4)
    } else {
        v5 = (v5 >> ((65535) & 63))
    }
    if v6 {
        fmt.Println(v1, v1)
        var v14 float32 = (0.0 / ((0.0 * 3.14159) - -2.25))
        _ = v14
    } else {
        v15 := GenS2{F0: 65535, F1: 1, F2: 2}
        switch -2147483648 {
        case 846:
            for fi16 := 0; fi16 < 10; fi16++ {
                v1 += -1
                for fi17 := 0; fi17 < 1; fi17++ {
                    v18 := GenS1{F0: 0, F1: v6}
                    if (false && ((!(float32(-2.25) < 1e10)) || v6)) {
                        fmt.Println(v6)
                    }
                    _ = fi17
                    _ = v18
                }
                fmt.Println(fi16)
                _ = fi16
            }
        case 42:
            for fi19 := 0; fi19 < 6; fi19++ {
                v7 += int32(v5)
                var v20 uint32 = 2147483647
                if ((v6 || (108 > v5)) || (v6 || (1 < 760))) {
                    fmt.Println(v1, fi19, v5)
                    fmt.Println(v20, v5, v1)
                    for fi21 := 0; fi21 < 13; fi21++ {
                        v22 := GenS2{F0: 1, F1: 42, F2: (v5 | 127)}
                        if true {
                            var v23 int64 = v4
                            var v24 uint16 = 65535
                            _ = v23
                            _ = v24
                        }
                        v25 := GenS3{F0: 2147483647}
                        _ = fi21
                        _ = v22
                        _ = v25
                    }
                } else {
                    v4 = v4
                    var v26 int64 = (v4 ^ (v4 | (v4 % ((v4) | 1))))
                    _ = v26
                }
                _ = fi19
                _ = v20
            }
        }
        v27 := GenS1{F0: 65535, F1: v6}
        _ = v15
        _ = v27
    }
    v5 = v5
    switch -814 {
    case 32767:
        fmt.Println(v4)
    case 9223372036854775807:
        v28 := []int{-2147483648, 363}
        _ = v28
    }
    fmt.Println(v4, v1)
    fmt.Println(v5)
    fmt.Println(v1, v5)
    v3 += (v7 - (v3 | v7))
    if true {
        switch 42 {
        case 745:
            fmt.Println(v6, v6)
        }
    }
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
