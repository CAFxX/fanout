package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0() uint {
    v4 := GenS1{F0: 42, F1: 2, F2: 255}
    var v5 int64 = -887
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 38
    v3 := make(map[string]bool)
    for fi6 := 0; fi6 < 6; fi6++ {
        v2 = (v2 * (^(v2)))
        _ = fi6
    }
    v2 = (v2 + 2147483647)
    var v7 uint16 = 2
    v2 = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v8 := GenS1{F0: 716, F1: v7, F2: 1}
    _ = len(v1)
    v2 = (v2 * v2)
    switch 128 {
    case 163:
        var v9 int16 = 42
        _ = v9
    case 927:
        var v10 uint64 = 4294967295
        _ = v10
    default:
        v7 += (v7 & v7)
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        switch (fi11 ^ fi11) {
        case 9223372036854775807:
            v12 := GenS1{F0: 2147483647, F1: (^(v7)), F2: 127}
            _ = v12
        case 0:
            for fi13 := 0; fi13 < 7; fi13++ {
                if true {
                    switch ("foo bar" + "\n\t") {
                    case "foo bar":
                        if false {
                            v7 += 65535
                        } else {
                            v1 = append(v1, (^(v2)))
                            v14 := GenS1{F0: int32(v2), F1: 843, F2: 127}
                            switch ("\n\t" + ("a" + "foo bar")) {
                            case "hello":
                                var v15 int32 = 2147483647
                                _ = v15
                            default:
                                var v16 bool = false
                                _ = v16
                            }
                            v2 -= v2
                            _ = v14
                        }
                    default:
                        v1 = append(v1, (v2 + (v2 / (((v2 >> ((18446744073709551615) & 63))) | 1))))
                        switch "hello" {
                        case "a":
                            var v17 uint64 = 773
                            _ = v17
                        case "":
                            fmt.Println(v7, v7)
                        }
                    }
                    v2 -= v2
                    var v18 int32 = -1
                    _ = v18
                } else {
                    var v19 int = fi13
                    if false {
                        v20 := GenS1{F0: 0, F1: 2, F2: 199}
                        v1 = append(v1, (^(v2)))
                        var v21 int32 = -450
                        _ = v20
                        _ = v21
                    }
                    _ = v19
                }
                _ = fi13
            }
        case 879:
            v22 := GenS1{F0: 1, F1: (v7 * v7), F2: 1}
            _ = v22
        }
        _ = fi11
    }
    v2 += uint32(v7)
    v2 -= v2
    v23 := GenS1{F0: 1, F1: 65535, F2: 0}
    if (1.5 <= 0.1) {
        v7 = (^(v7))
        v24 := GenS1{F0: -863, F1: (v7 | 406), F2: 0}
        v2 += v2
        _ = v24
    } else {
        v7 = v7
        var v25 int = 9223372036854775807
        var v26 uint = 18446744073709551615
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: 2147483647, F1: v7, F2: 7}
    v7 = 65535
    v28 := GenS2{F0: (v7 | v7)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v7
    _ = v8
}
