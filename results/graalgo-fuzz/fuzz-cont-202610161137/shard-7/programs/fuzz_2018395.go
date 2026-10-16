package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 int }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[uint32]bool)
    var v3 float32 = 1.5
    v4 := make(map[int]int)
    var v5 int64 = 374
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint32 = 42
        _ = fi6
        _ = v7
    }
    var v8 int32 = -282
    if (747 < 255) {
        switch 2147483647 {
        case 32767:
            switch -9223372036854775808 {
            case -2147483648:
                v5 -= v5
            default:
                switch "\n\t" {
                case "":
                    for fi9 := 0; fi9 < 2; fi9++ {
                        for fi10 := 0; fi10 < 12; fi10++ {
                            v11 := GenS1{F0: 1, F1: "a"}
                            v3 -= v3
                            if true {
                                v3 = (v3 + v3)
                                v12 := GenS1{F0: 4294967295, F1: "a"}
                                v1 = 9223372036854775807
                                _ = v12
                            } else {
                                v13 := []int{fi10, -311}
                                if false {
                                    fmt.Println(fi9)
                                    v14 := GenS2{F0: 1000, F1: 542}
                                    _ = v14
                                }
                                _ = v13
                            }
                            _ = fi10
                            _ = v11
                        }
                        v5 = 1
                        _ = fi9
                    }
                case "hello":
                    if (1 <= 0) {
                        v15 := []int{-9223372036854775808}
                        _ = v15
                    } else {
                        delete(v4, -2147483648)
                        v3 = ((3.14159 * ((v3 / v3) + v3)) + 0.0)
                    }
                default:
                    _ = len(v2)
                }
            }
        case 127:
            v16 := []int{584, 0}
            _ = v16
        case -2147483648:
            fmt.Println(v3)
        default:
            var v17 int = -159
            _ = v17
        }
    }
    fmt.Println(v3)
    var v18 int = -2147483648
    var v19 string = "hello"
    for fi20 := 0; fi20 < 11; fi20++ {
        switch (v19 + ("" + "hello")) {
        case "foo bar":
            switch v18 {
            case 32768:
                var v21 int = 128
                _ = v21
            case 586:
                v22 := GenS3{F0: 886}
                _ = v22
            case -2147483648:
                for fi23 := 0; fi23 < 15; fi23++ {
                    var v24 int32 = int32(v1)
                    _ = fi23
                    _ = v24
                }
            default:
                for fi25 := 0; fi25 < 15; fi25++ {
                    v26 := GenS4{F0: v19, F1: 42}
                    _ = len(v2)
                    var v27 float64 = float64(float32(0.0))
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
        case "\n\t":
            var v28 int16 = -32768
            _ = v28
        case "hello":
            v29 := []int{}
            _ = v29
        }
        v8 = (v8 << ((1) & 63))
        _ = fi20
    }
    v30 := []int{128}
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v8
}
