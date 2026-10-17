package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 uint }
func genFn0() uint16 {
    v7 := GenS1{F0: 127}
    var v8 uint32 = 1
    v9 := GenS2{F0: -1}
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    var v1 int8 = 127
    var v2 int8 = (v1 / ((1) | 1))
    var v3 uint16 = 0
    v4 := []uint32{4294967295, 1, 1, 723}
    var v5 int32 = 2147483647
    var v6 uint32 = 42
    var v10 uint8 = 236
    _ = len(v4)
    v5 = (-(v5))
    var v11 uint8 = 2
    _ = len(v4)
    switch 128 {
    case -9223372036854775808:
        if (!true) {
            switch "a" {
            case "foo bar":
                for fi12 := 0; fi12 < 10; fi12++ {
                    if (686 > uint(v6)) {
                        fmt.Println(v11, v2, v10)
                    }
                    _ = len(v4)
                    v10 -= 42
                    _ = fi12
                }
            case "\n\t":
                v13 := GenS2{F0: (v5 % ((704) | 1))}
                _ = v13
            default:
                if ("a" != "a") {
                    v14 := GenS3{F0: v5, F1: v3}
                    switch "foo bar" {
                    case "foo bar":
                        if true {
                            for fi15 := 0; fi15 < 5; fi15++ {
                                v16 := GenS4{F0: (v10 >> ((1) & 63)), F1: 42}
                                _ = fi15
                                _ = v16
                            }
                            var v17 int64 = int64(v1)
                            _ = v17
                        } else {
                            if len(v4) > 0 {
                                v4[0] = (v6 >> ((52) & 63))
                            }
                        }
                    case "hello":
                        v18 := GenS5{F0: 44}
                        _ = v18
                    case "\n\t":
                        v10 -= v10
                    default:
                        for fi19 := 0; fi19 < 15; fi19++ {
                            if false {
                                fmt.Println(fi19, v2)
                            }
                            _ = fi19
                        }
                    }
                    _ = v14
                }
            }
            fmt.Println(v3)
            var v20 uint = 800
            _ = v20
        } else {
            v10 += v11
            if len(v4) > 0 {
                v4[0] = 2147483647
            }
        }
    case 127:
        var v21 int16 = -1
        _ = v21
    }
    v22 := GenS3{F0: v5, F1: 65535}
    v6 -= (v6 >> ((1) & 63))
    v23 := GenS5{F0: 325}
    fmt.Println(v5, v11)
    if (32767 < 32768) {
        v4 = append(v4, v6)
    }
    fmt.Println(v10, v11, v2)
    v1 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
