package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint8 }
type GenS2 struct { F0 bool; F1 bool; F2 int64 }
type GenS3 struct { F0 string; F1 uint8; F2 uint32 }

func main() {
    var v1 string = ("hello" + "hello")
    var v2 uint16 = 733
    v3 := []int{-137, -2147483648, 878, -1}
    v4 := GenS1{F0: -9223372036854775808, F1: (v1 + v1), F2: 0}
    var v5 int64 = -9223372036854775808
    switch -755 {
    case 1:
        v6 := GenS1{F0: (v5 / (((v5 ^ v5)) | 1)), F1: v1, F2: uint8(v2)}
        _ = v6
    default:
        if (1 == 0) {
            var v7 int64 = 0
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v7
        }
    }
    v8 := GenS2{F0: (2 != 2147483648), F1: (42 < -1), F2: (v5 / ((v5) | 1))}
    var v9 int32 = 630
    v9 = 42
    if ((false && (!(0 < 255))) && true) {
        v10 := GenS2{F0: false, F1: true, F2: 1}
        if (0 != 1) {
            v11 := GenS2{F0: true, F1: false, F2: (v5 >> ((565) & 63))}
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 int = int(v2)
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        var v14 float32 = 1e10
        _ = v10
        _ = v14
    } else {
        v1 = (v1 + v1)
        for fi15 := 0; fi15 < 6; fi15++ {
            if (-520 < 1) {
                for fi16 := 0; fi16 < 9; fi16++ {
                    if len(v3) > 0 {
                        v3[0] = fi16
                    }
                    switch v1 {
                    case "a":
                        if (!true) {
                            v17 := GenS1{F0: v5, F1: ((((v1 + (v1 + v1)) + (v1 + (v1 + v1))) + (v1 + v1)) + v1), F2: 109}
                            _ = v17
                        }
                    case "\n\t":
                        var v18 uint8 = 127
                        _ = v18
                    default:
                        _ = len(v3)
                    }
                    _ = fi16
                }
                v9 += v9
            }
            _ = fi15
        }
    }
    v5 = (-(v5))
    v19 := GenS3{F0: "\n\t", F1: 0, F2: 916}
    fmt.Println(v1, v1)
    v3 = append(v3, 568)
    var v20 string = v1
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v21 int = 9223372036854775807
    var v22 uint32 = 42
    if true {
        if (127 > 127) {
            v3 = append(v3, (v21 & 128))
        }
        var v23 int16 = -32768
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
