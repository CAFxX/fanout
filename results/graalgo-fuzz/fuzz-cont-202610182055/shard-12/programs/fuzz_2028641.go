package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = -2147483648
    var v3 int64 = -647
    v4 := []int{32768, (v2 * (v2 / (((v2 - 2147483647)) | 1)))}
    var v5 int = -84
    var v6 string = "hello"
    var v7 string = ("\n\t" + v6)
    v2 = 929
    if false {
        var v8 int32 = -1
        v4 = append(v4, v5)
        if false {
            fmt.Println(v7, v5)
            v3 += -890
            v7 = v7
        }
        _ = v8
    } else {
        v1[(v2 & (v5 & 32768))] = (v3 == -1)
        if (!(!(!(-27 == -128)))) {
            fmt.Println(v2)
            switch (v2 + (v5 % (((v5 + (v5 - 127))) | 1))) {
            case 410:
                var v9 bool = false
                _ = v9
            case 2147483647:
                v10 := GenS1{F0: 1}
                _ = v10
            default:
                delete(v1, 4294967295)
            }
        }
    }
    _ = len(v4)
    _ = len(v1)
    if (4294967295 > 2) {
        fmt.Println(v5, v2)
    } else {
        if true {
            delete(v1, 2147483647)
            switch v6 {
            case "hello":
                v3 = (v3 - 0)
            default:
                var v11 uint = 42
                _ = v11
            }
            v12 := GenS1{F0: 0}
            _ = v12
        } else {
            v7 = (("\n\t" + "") + v6)
            _ = len(v1)
        }
        if true {
            var v13 int8 = 1
            var v14 float32 = ((0.0 - (0.1 * (((1.5 / 1.5) * float32((((1.5 / (3.14159 + (0.0 / 1.0))) * 0.1) - (((0.0 / -2.25) + 0.1) - 1.5)))) - 1e10))) / 1.0)
            for fi15 := 0; fi15 < 12; fi15++ {
                for fi16 := 0; fi16 < 12; fi16++ {
                    v13 = (^(v13))
                    var v17 uint32 = 925
                    v18 := GenS2{F0: (fi15 & (fi15 + 127))}
                    var v19 uint8 = 91
                    switch 0 {
                    case -745:
                        v20 := GenS3{F0: 965, F1: (v19 / (((v19 & 66)) | 1)), F2: 1}
                        _ = v20
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                _ = fi15
            }
            _ = v13
            _ = v14
        } else {
            v21 := v1[2147483647]
            _ = v21
        }
        if true {
            v7 = v6
        }
    }
    v7 = (v7 + v7)
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 bool = true
        if ((float64(0.1) / 1.5) < 3.14159) {
            v3 -= (v3 << ((18446744073709551615) & 63))
        } else {
            v23 = (66 < 128)
            for fi24 := 0; fi24 < 7; fi24++ {
                fmt.Println(fi24, v7)
                _ = fi24
            }
        }
        _ = fi22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
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
