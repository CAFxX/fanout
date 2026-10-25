package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint16 }

func main() {
    v1 := []uint32{65535, 42, 1}
    var v2 bool = (true && true)
    v3 := []uint32{669, 1, 690, 2147483648}
    var v4 int64 = -1
    switch 0 {
    case -9223372036854775808:
        v5 := GenS1{F0: 2147483647, F1: 0}
        _ = v5
    case -2147483648:
        for fi6 := 0; fi6 < 3; fi6++ {
            for fi7 := 0; fi7 < 18; fi7++ {
                switch fi7 {
                case 686:
                    var v8 uint64 = 4294967295
                    _ = v8
                case 128:
                    var v9 float64 = (-2.25 / 3.14159)
                    _ = v9
                case 546:
                    if (0 >= 2147483647) {
                        var v10 int16 = 32767
                        if ((-(fi6)) >= fi6) {
                            v2 = (v2 && (2147483647 <= 552))
                        }
                        _ = v10
                    }
                    var v11 int = -1
                    _ = v11
                default:
                    v3 = append(v3, 2)
                }
                var v12 uint16 = 714
                var v13 float32 = 1e10
                _ = fi7
                _ = v12
                _ = v13
            }
            _ = fi6
        }
    }
    fmt.Println(v2, v2)
    v3 = append(v3, 0)
    fmt.Println(v2, v4)
    if (579 <= 18446744073709551615) {
        if (2 <= 0) {
            if ("" >= "\n\t") {
                fmt.Println(v2)
                if v2 {
                    var v14 uint32 = 2
                    _ = v14
                }
            }
        }
        v2 = v2
        switch -609 {
        case -886:
            v15 := GenS1{F0: -356, F1: 255}
            _ = v15
        default:
            v16 := GenS2{F0: -9223372036854775808, F1: 0}
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := GenS3{F0: 128, F1: 151, F2: 2}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v4 += v4
        _ = fi17
        _ = v18
    }
    if (42 != 0) {
        fmt.Println(v2, v2)
        v2 = (!((!v2) && (127 <= 1)))
    }
    v19 := GenS1{F0: 2147483647, F1: 18446744073709551615}
    var v20 float32 = -2.25
    v21 := GenS1{F0: 42, F1: 955}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
