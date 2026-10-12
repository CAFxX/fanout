package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int32; F1 int }

func main() {
    var v1 int16 = -32768
    v2 := []uint32{2, 42, 4294967295, 42}
    var v3 uint8 = 127
    v3 = v3
    for fi4 := 0; fi4 < 17; fi4++ {
        switch fi4 {
        case 1:
            if (42 == 42) {
                v3 += v3
            } else {
                fmt.Println(fi4, v3)
            }
        case -402:
            var v5 uint8 = (v3 & v3)
            _ = v5
        default:
            for fi6 := 0; fi6 < 8; fi6++ {
                var v7 int32 = -590
                var v8 int64 = 240
                v3 = v3
                v9 := GenS1{F0: v3, F1: 0}
                v3 = 2
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
        }
        _ = fi4
    }
    switch "a" {
    case "\n\t":
        var v10 uint = 1
        _ = v10
    default:
        if false {
            _ = len(v2)
            _ = len(v2)
            switch "" {
            case "a":
                if ((483 >= 42) && (2147483648 >= 701)) {
                    var v11 uint16 = 42
                    switch "" {
                    case "foo bar":
                        v12 := GenS1{F0: (^(v3)), F1: 4294967295}
                        _ = v12
                    case "a":
                        var v13 int32 = 0
                        _ = v13
                    case "":
                        v3 += v3
                    default:
                        for fi14 := 0; fi14 < 5; fi14++ {
                            v15 := GenS2{F0: 124, F1: 0}
                            for fi16 := 0; fi16 < 5; fi16++ {
                                v1 = 42
                                v17 := GenS2{F0: 58, F1: 471}
                                if false {
                                    var v18 int32 = 2147483647
                                    _ = v18
                                }
                                _ = fi16
                                _ = v17
                            }
                            _ = fi14
                            _ = v15
                        }
                    }
                    v19 := GenS2{F0: 255, F1: 42}
                    _ = v11
                    _ = v19
                } else {
                    var v20 uint8 = 42
                    v1 -= v1
                    switch -172 {
                    case 668:
                        var v21 float64 = 1e10
                        _ = v21
                    }
                    _ = v20
                }
            case "hello":
                switch 0 {
                case 2147483647:
                    if len(v2) > 0 {
                        v2[0] = 601
                    }
                default:
                    v22 := GenS1{F0: v3, F1: 640}
                    _ = v22
                }
            }
        }
    }
    fmt.Println(v1)
    if false {
        fmt.Println(v1, v3)
        var v23 bool = true
        _ = v23
    }
    v24 := GenS3{F0: 698, F1: -9223372036854775808}
    fmt.Println(v1, v3)
    v1 = (-(v1))
    switch 32768 {
    case 678:
        v3 = (^(v3))
    }
    v2 = append(v2, 407)
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
}
