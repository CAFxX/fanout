package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 bool }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]int)
    var v3 int8 = 42
    var v4 uint8 = 42
    v5 := v2[9223372036854775807]
    switch "\n\t" {
    case "":
        var v6 float64 = 0.1
        _ = v6
    case "foo bar":
        v7 := v2[v5]
        _ = v7
    }
    if false {
        var v8 uint64 = 0
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint32 = 65535
            var v11 uint16 = 255
            v8 += 42
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 uint64 = 2
        _ = v12
        _ = v8
    } else {
        v13 := []int{(v5 >> ((796) & 63)), v5, (v5 << ((255) & 63))}
        _ = v13
    }
    if (float64(float32(0.1)) <= 3.14159) {
        switch (v5 >> ((0) & 63)) {
        case 906:
            v5 = v5
        case 42:
            v14 := v2[v5]
            _ = v14
        default:
            var v15 int32 = 2147483647
            _ = v15
        }
    } else {
        if (417 < 780) {
            switch v5 {
            case 0:
                switch v5 {
                case 9223372036854775807:
                    for fi16 := 0; fi16 < 8; fi16++ {
                        var v17 int16 = 1
                        _ = fi16
                        _ = v17
                    }
                case 884:
                    v18 := []int{(v5 + int(v3)), v5, -724, 9223372036854775807}
                    _ = v18
                case 42:
                    switch 32768 {
                    case 1:
                        var v19 uint8 = v4
                        _ = v19
                    default:
                        var v20 int16 = 793
                        _ = v20
                    }
                default:
                    switch -2147483648 {
                    case 746:
                        switch "" {
                        case "":
                            v21 := GenS1{F0: 137, F1: v4, F2: ((true && (int8(v4) > 1)) && false)}
                            _ = v21
                        default:
                            for fi22 := 0; fi22 < 20; fi22++ {
                                _ = len(v1)
                                _ = fi22
                            }
                        }
                    }
                }
            case 1:
                var v23 uint = 1
                _ = v23
            default:
                v24 := GenS1{F0: 2, F1: 13, F2: false}
                _ = v24
            }
            switch "\n\t" {
            case "\n\t":
                var v25 int = v5
                _ = v25
            }
        } else {
            var v26 bool = false
            switch v5 {
            case 0:
                v3 -= v3
            }
            var v27 uint8 = (v4 % ((0) | 1))
            _ = v26
            _ = v27
        }
    }
    v28 := GenS1{F0: 18446744073709551615, F1: 1, F2: (-149 != 0)}
    if (true && true) {
        fmt.Println(v4, v5, v3)
        v29 := []int{(v5 | 391), v5, (v5 % ((-1) | 1))}
        _ = v29
    }
    fmt.Println(v5, v3, v5)
    v30 := []int{32768}
    var v31 int = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
