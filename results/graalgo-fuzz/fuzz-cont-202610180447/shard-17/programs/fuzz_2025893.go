package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int }
type GenS2 struct { F0 int64; F1 uint16; F2 int }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int16 }

func main() {
    var v1 string = ("" + "a")
    var v2 int8 = -124
    v3 := make(map[int]string)
    var v4 bool = true
    v5 := []int{4294967295, 127, 42}
    v6 := make(map[string]uint32)
    if v4 {
        fmt.Println(v1)
    } else {
        fmt.Println(v2)
    }
    switch 4294967295 {
    case 127:
        v7 := GenS1{F0: 2, F1: v2, F2: 0}
        _ = v7
    default:
        v8 := GenS1{F0: 276, F1: 42, F2: -114}
        _ = v8
    }
    var v9 float32 = 1e10
    fmt.Println(v2, v2, v1)
    switch 4294967295 {
    case 32768:
        v2 = (v2 >> ((18446744073709551615) & 63))
    case 128:
        if false {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            var v10 string = "a"
            v11 := GenS1{F0: 18446744073709551615, F1: 127, F2: 62}
            _ = v10
            _ = v11
        }
    case 32767:
        var v12 int = 32767
        _ = v12
    }
    if (-2.25 < float64(((float32(-2.25) / 1.0) / -2.25))) {
        var v13 float32 = 3.14159
        v13 += -2.25
        var v14 uint16 = 1
        _ = v13
        _ = v14
    }
    switch "a" {
    case "":
        v9 -= 3.14159
    case "hello":
        for fi15 := 0; fi15 < 2; fi15++ {
            switch fi15 {
            case 448:
                if true {
                    v16 := GenS1{F0: 2, F1: (v2 & v2), F2: (fi15 % ((fi15) | 1))}
                    var v17 int8 = v2
                    _ = v16
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 17; fi18++ {
                        if (-1 > 1) {
                            v4 = true
                        } else {
                            var v19 int16 = 1000
                            _ = v19
                        }
                        fmt.Println(fi18)
                        _ = fi18
                    }
                    var v20 bool = false
                    v9 = 3.14159
                    _ = v20
                }
            case 411:
                v1 = v1
            default:
                v21 := GenS2{F0: 42, F1: 2, F2: (fi15 % ((fi15) | 1))}
                _ = v21
            }
            _ = len(v3)
            var v22 int16 = 1
            _ = fi15
            _ = v22
        }
    case "a":
        v2 += 127
    }
    _ = len(v6)
    if (!((^(v2)) > -1)) {
        v23 := GenS3{F0: -2147483648, F1: -9223372036854775808}
        v9 -= v9
        _ = v23
    }
    if false {
        fmt.Println(v1, v1, v9)
        if (206 <= 127) {
            var v24 int = 4294967295
            v24 = (-(v24))
            _ = v24
        } else {
            if len(v5) > 0 {
                v5[0] = 165
            }
        }
    } else {
        v25 := GenS4{F0: 42}
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
