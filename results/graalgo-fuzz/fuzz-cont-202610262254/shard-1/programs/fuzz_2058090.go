package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int16; F1 int16; F2 int64 }
type GenS3 struct { F0 string; F1 bool; F2 int }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 1e10
    var v3 int16 = 231
    var v4 int = 32767
    v2 += float64(-2.25)
    v4 += v4
    switch -429 {
    case -2147483648:
        v4 -= 32768
    case 0:
        v5 := []int{v4}
        _ = v5
    }
    v3 -= (v3 & (-(v3)))
    var v6 int8 = -128
    v2 -= float64(1e10)
    v7 := GenS1{F0: 65535, F1: 1}
    var v8 uint16 = 145
    if false {
        var v9 int16 = (v3 * v3)
        v4 = v4
        v2 = float64(3.14159)
        _ = v9
    } else {
        var v10 uint64 = 445
        fmt.Println(v8, v6)
        if true {
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := GenS1{F0: 862, F1: 2147483647}
                var v13 bool = false
                v13 = ((v2 - (0.1 * ((0.0 * v2) - 3.14159))) == (0.1 - (v2 - v2)))
                _ = fi11
                _ = v12
                _ = v13
            }
            fmt.Println(v3)
            if true {
                switch "" {
                case "foo bar":
                    var v14 int32 = -1
                    _ = v14
                default:
                    v15 := GenS2{F0: (-(v3)), F1: (v3 / ((v3) | 1)), F2: 9223372036854775807}
                    _ = v15
                }
                var v16 uint = 4294967295
                v2 = ((v2 * float64(0.0)) - v2)
                _ = v16
            } else {
                var v17 string = "a"
                if false {
                    fmt.Println(v10, v4, v6)
                    var v18 uint16 = v8
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        var v20 bool = false
                        v10 -= (v10 ^ v10)
                        for fi21 := 0; fi21 < 12; fi21++ {
                            switch v17 {
                            case "foo bar":
                                switch v4 {
                                case 32768:
                                    var v22 uint8 = 42
                                    _ = v22
                                case -9223372036854775808:
                                    v1[(fi19 % ((-360) | 1))] = "a"
                                }
                            default:
                                fmt.Println(v4)
                            }
                            switch v17 {
                            case "":
                                var v23 uint = 515
                                _ = v23
                            case "\n\t":
                                v6 = -11
                            case "a":
                                var v24 int = v4
                                _ = v24
                            }
                            fmt.Println(fi19, v6, v8)
                            _ = fi21
                        }
                        _ = fi19
                        _ = v20
                    }
                    v25 := GenS3{F0: "\n\t", F1: (328 < 4294967295), F2: v4}
                    _ = v25
                }
                _ = v17
            }
        } else {
            var v26 int32 = 2147483647
            _ = v26
        }
        _ = v10
    }
    v8 -= (v8 % (((v8 + 65535)) | 1))
    fmt.Println(v8)
    v3 = 829
    v6 = (v6 << ((42) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
