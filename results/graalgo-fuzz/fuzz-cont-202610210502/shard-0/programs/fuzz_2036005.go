package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int16 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 float32 = 3.14159
    var v2 int16 = 947
    v3 := make(map[int]int)
    var v4 int16 = v2
    v5 := make(map[int]string)
    v6 := []uint32{153, 0, 2147483648}
    for fi7 := 0; fi7 < 5; fi7++ {
        v2 -= -1
        var v8 int32 = -2147483648
        v9 := GenS1{F0: 656, F1: (^(v8)), F2: 1}
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v2)
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS1{F0: 1, F1: 2147483647, F2: (v2 + 1000)}
        v1 = (v1 * v1)
        if (900 < 42) {
            v12 := GenS2{F0: 42, F1: (v4 ^ v4), F2: 699}
            var v13 uint = 18446744073709551615
            fmt.Println(v4, v2, v4)
            _ = v12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    if false {
        if ((v4 & (^(v4))) == (v4 << ((42) & 63))) {
            for fi14 := 0; fi14 < 2; fi14++ {
                v1 += v1
                v1 = 0.0
                for fi15 := 0; fi15 < 15; fi15++ {
                    v16 := GenS1{F0: 65535, F1: 999, F2: v2}
                    _ = fi15
                    _ = v16
                }
                _ = fi14
            }
            var v17 uint = 4294967295
            _ = v17
        } else {
            v2 = -479
            v2 -= (v4 & v2)
        }
    } else {
        fmt.Println(v2, v2)
        v18 := GenS1{F0: 65535, F1: -106, F2: (v4 % (((v4 | v4)) | 1))}
        if (-2.25 >= 0.0) {
            var v19 uint = 2
            _ = v19
        }
        _ = v18
    }
    var v20 uint = 42
    v21 := GenS2{F0: 2147483647, F1: v4, F2: 2}
    if (!((42 > 567) && (!true))) {
        switch 2147483647 {
        case 32767:
            for fi22 := 0; fi22 < 12; fi22++ {
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(fi22)
                case "hello":
                    var v23 uint32 = 436
                    _ = v23
                case "foo bar":
                    for fi24 := 0; fi24 < 10; fi24++ {
                        v25 := GenS3{F0: 42}
                        var v26 int64 = 9223372036854775807
                        v27 := GenS4{F0: 255}
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                }
                v1 += v1
                var v28 uint16 = 253
                _ = fi22
                _ = v28
            }
        case -457:
            _ = len(v3)
        default:
            _ = len(v5)
        }
    } else {
        fmt.Println(v2)
    }
    var v29 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
