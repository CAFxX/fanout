package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint }
func genFn0(p0 int16) uint64 {
    v7 := GenS1{F0: 127}
    var v8 int32 = 42
    var v9 int64 = 1
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 768
}

func main() {
    var v1 uint16 = 32767
    var v2 uint8 = uint8(v1)
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]int)
    var v5 int = 391
    var v6 int64 = -9223372036854775808
    var v10 float64 = ((1e10 + 1e10) - (0.1 + 0.1))
    if true {
        v11 := GenS2{F0: v5, F1: 806}
        fmt.Println(v2)
        _ = v11
    }
    v12 := GenS1{F0: (^(v5))}
    var v13 uint32 = 42
    if false {
        v10 = v10
        v14 := []int{(^(v5))}
        var v15 uint16 = (v1 & (v1 >> ((255) & 63)))
        _ = v14
        _ = v15
    } else {
        var v16 int64 = (v6 * 938)
        switch (v5 ^ v5) {
        case 127:
            v17 := GenS1{F0: v5}
            _ = v17
        case 1:
            _ = len(v3)
        case -401:
            switch "\n\t" {
            case "":
                var v18 int8 = -1
                _ = v18
            case "hello":
                if true {
                    switch "" {
                    case "":
                        v19 := []int{int(v1)}
                        _ = v19
                    case "a":
                        switch ("\n\t" + ("a" + "")) {
                        case "\n\t":
                            v20 := []int{v5, int(v1), (v5 % (((v5 & v5)) | 1)), 773}
                            _ = v20
                        case "":
                            v21 := []int{v5, (v5 % ((v5) | 1)), v5}
                            _ = v21
                        case "hello":
                            if ((!true) && (94 >= 0)) {
                                fmt.Println(v16, v13, v1)
                                if false {
                                    var v22 float32 = 3.14159
                                    v10 += 3.14159
                                    _ = v22
                                } else {
                                    var v23 int = (v5 >> ((1) & 63))
                                    for fi24 := 0; fi24 < 7; fi24++ {
                                        fmt.Println(v16)
                                        _ = fi24
                                    }
                                    v25 := GenS2{F0: 0, F1: 255}
                                    _ = v23
                                    _ = v25
                                }
                            }
                        }
                    }
                } else {
                    v13 = 2147483648
                    var v26 uint64 = 81
                    _ = v26
                }
            default:
                v13 = v13
            }
        }
        _ = v16
    }
    switch 9223372036854775807 {
    case 404:
        fmt.Println(v10, v13)
    case 1:
        var v27 int8 = -54
        _ = v27
    case 127:
        var v28 uint8 = (v2 << ((42) & 63))
        _ = v28
    default:
        v29 := GenS2{F0: (v5 / (((v5 ^ v5)) | 1)), F1: uint(v1)}
        _ = v29
    }
    v13 += (v13 << ((350) & 63))
    var v30 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
