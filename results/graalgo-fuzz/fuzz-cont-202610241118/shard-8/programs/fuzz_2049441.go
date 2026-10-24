package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int64 }

func main() {
    var v1 uint64 = 2
    var v2 uint8 = 75
    v3 := make(map[uint32]uint32)
    var v4 uint8 = v2
    fmt.Println(v2, v2)
    if true {
        for fi5 := 0; fi5 < 2; fi5++ {
            switch fi5 {
            case -2147483648:
                switch "" {
                case "a":
                    for fi6 := 0; fi6 < 4; fi6++ {
                        var v7 float32 = -2.25
                        _ = fi6
                        _ = v7
                    }
                case "foo bar":
                    if false {
                        for fi8 := 0; fi8 < 14; fi8++ {
                            switch ("a" + "") {
                            case "foo bar":
                                if (2 <= 2) {
                                    v9 := v3[941]
                                    var v10 uint8 = v4
                                    v9 = (v9 + (v9 % ((v9) | 1)))
                                    _ = v10
                                    _ = v9
                                }
                            case "hello":
                                if (2 == 42) {
                                    v11 := []int{(-(fi8))}
                                    _ = v11
                                }
                            default:
                                v1 -= v1
                                var v12 uint64 = (^(v1))
                                _ = v12
                            }
                            v1 = (^(v1))
                            v2 -= v2
                            _ = fi8
                        }
                        v4 = v4
                    }
                case "\n\t":
                    var v13 int8 = -82
                    _ = v13
                }
            case 1:
                v14 := GenS1{F0: 2, F1: -32768, F2: -26}
                _ = v14
            default:
                v4 -= (v2 << ((4294967295) & 63))
            }
            fmt.Println(v1, v1)
            fmt.Println(v4, v2)
            _ = fi5
        }
        var v15 string = "hello"
        _ = v15
    } else {
        switch 42 {
        case -2147483648:
            v16 := []int{-9223372036854775808, -2147483648}
            _ = v16
        default:
            switch 2147483647 {
            case -601:
                fmt.Println(v2, v1, v1)
            case -116:
                var v17 uint = 18446744073709551615
                _ = v17
            default:
                for fi18 := 0; fi18 < 6; fi18++ {
                    v4 = (v4 << ((1) & 63))
                    _ = fi18
                }
            }
        }
        v19 := GenS1{F0: 1, F1: 500, F2: -1}
        _ = v19
    }
    v3[65535] = 775
    var v20 int64 = 0
    fmt.Println(v2)
    var v21 int32 = -1
    var v22 int = 32768
    var v23 uint64 = (v1 * 42)
    if (((^(v2)) == (^(v4))) && (false && ((v22 << ((1) & 63)) != (v22 >> ((555) & 63))))) {
        var v24 int64 = 1
        for fi25 := 0; fi25 < 19; fi25++ {
            v26 := GenS1{F0: 1, F1: -30, F2: 739}
            if true {
                if (0 < 1) {
                    v4 = v4
                    v27 := GenS2{F0: 1, F1: 2147483648, F2: v20}
                    fmt.Println(v22, fi25, v1)
                    _ = v27
                }
                if (uint32(v23) == 2) {
                    v1 += (v1 % ((v23) | 1))
                    var v28 int32 = -1
                    var v29 float32 = 1e10
                    _ = v28
                    _ = v29
                }
            } else {
                fmt.Println(v1)
                var v30 int64 = int64(v2)
                _ = v30
            }
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    v31 := GenS2{F0: 2, F1: 65535, F2: (v20 - 0)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v4
}
