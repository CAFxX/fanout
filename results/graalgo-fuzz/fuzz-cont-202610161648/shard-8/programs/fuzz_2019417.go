package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int }
func genFn0() int32 {
    v6 := []int{-1, 1, 127}
    if true {
        if len(v6) > 0 {
            v6[0] = -422
        }
    } else {
        var v7 int8 = 127
        _ = v7
    }
    var v8 int32 = 42
    _ = v6
    _ = v8
    return -77
}

func main() {
    var v1 uint16 = 0
    v2 := []int{-9223372036854775808}
    var v3 int32 = -682
    var v4 uint = 65535
    v5 := []int{}
    v1 = v1
    for fi9 := 0; fi9 < 12; fi9++ {
        if (2 <= 0) {
            v10 := GenS1{F0: (v3 / ((genFn0()) | 1)), F1: 4294967295}
            if (155 < 42) {
                if ((false || false) || false) {
                    v11 := GenS1{F0: int32(v4), F1: 42}
                    var v12 int32 = v3
                    _ = v11
                    _ = v12
                }
                switch 32768 {
                case 1:
                    switch fi9 {
                    case 2147483647:
                        var v13 uint8 = 127
                        _ = v13
                    case 127:
                        for fi14 := 0; fi14 < 7; fi14++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi14
                        }
                    case 388:
                        v15 := GenS1{F0: v3, F1: 540}
                        _ = v15
                    }
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                v4 = (v4 - v4)
                var v16 int8 = -4
                _ = v16
            }
            _ = v10
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v17 := GenS2{F0: (fi9 << (((v4 & 4294967295)) & 63))}
            switch "" {
            case "":
                for fi18 := 0; fi18 < 8; fi18++ {
                    var v19 uint8 = 179
                    v4 += v4
                    _ = fi18
                    _ = v19
                }
            case "\n\t":
                v4 = v4
            }
            _ = v17
        }
        v1 = (v1 + v1)
        _ = fi9
    }
    v1 = v1
    switch "\n\t" {
    case "":
        fmt.Println(v3, v4, v1)
    }
    v4 = v4
    var v20 int32 = -537
    v21 := GenS1{F0: v3, F1: uint32(v1)}
    v22 := GenS1{F0: v20, F1: 65535}
    fmt.Println(v20, v3, v20)
    v1 = (v1 % (((^(v1))) | 1))
    fmt.Println(v20)
    fmt.Println(v4, v3)
    var v23 uint16 = (v1 - v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
