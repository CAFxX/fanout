package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 uint64 = 127
    var v2 int8 = 100
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]int)
    v5 := GenS1{F0: 0}
    _ = len(v4)
    v3 += v3
    if (0.0 <= 1e10) {
        if (676 < 848) {
            var v6 int = -514
            _ = v6
        } else {
            if (!(true && (!false))) {
                var v7 int32 = -1
                _ = v7
            }
        }
        switch 32767 {
        case 128:
            v1 += (^(v3))
        }
        v8 := GenS1{F0: 128}
        _ = v8
    } else {
        if ((1.5 <= 1e10) && false) {
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 string = ("hello" + ("a" + ("hello" + "foo bar")))
                v11 := v4[349]
                _ = fi9
                _ = v10
                _ = v11
            }
        } else {
            if (9223372036854775807 >= int64(v2)) {
                v12 := GenS2{F0: 0, F1: false}
                v2 = 1
                var v13 int8 = (v2 | v2)
                _ = v12
                _ = v13
            }
            var v14 int16 = -32768
            v15 := GenS3{F0: 2147483647, F1: 65535, F2: 2}
            _ = v14
            _ = v15
        }
        switch -9223372036854775808 {
        case -608:
            v2 += 1
        case -612:
            var v16 int8 = (v2 >> ((727) & 63))
            _ = v16
        default:
            switch "" {
            case "foo bar":
                for fi17 := 0; fi17 < 8; fi17++ {
                    if (false && (100 > (v2 - 0))) {
                        for fi18 := 0; fi18 < 3; fi18++ {
                            v19 := []int{fi17, fi17, 9223372036854775807, -2147483648}
                            if (0 != 1) {
                                var v20 uint32 = 65535
                                v1 = v1
                                var v21 float64 = 0.1
                                _ = v20
                                _ = v21
                            }
                            v2 += (^(v2))
                            if true {
                                v1 += v1
                            } else {
                                var v22 uint16 = 2
                                v2 = 42
                                _ = v22
                            }
                            switch "hello" {
                            case "":
                                v2 += v2
                            }
                            v19 = append(v19, (fi17 * fi17))
                            v3 = 18446744073709551615
                            _ = fi18
                            _ = v19
                        }
                    }
                    _ = fi17
                }
            case "\n\t":
                fmt.Println(v3, v3, v1)
            case "a":
                v23 := GenS4{F0: 67}
                _ = v23
            default:
                switch 969 {
                case 2147483647:
                    v1 += 83
                case -943:
                    v3 -= 0
                case -9223372036854775808:
                    for fi24 := 0; fi24 < 17; fi24++ {
                        v3 = v1
                        switch "hello" {
                        case "hello":
                            if (!false) {
                                delete(v4, 1)
                                var v25 uint = 743
                                v1 = 86
                                _ = v25
                            }
                        case "\n\t":
                            v26 := GenS4{F0: 213}
                            _ = v26
                        default:
                            v1 = (v3 << ((2) & 63))
                        }
                        _ = fi24
                    }
                default:
                    fmt.Println(v3, v2, v3)
                }
            }
        }
        v27 := GenS1{F0: 255}
        _ = v27
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        v29 := GenS4{F0: 294}
        _ = fi28
        _ = v29
    }
    fmt.Println(v3, v1, v1)
    v30 := GenS3{F0: 1, F1: 555, F2: 42}
    v1 += (^(v1))
    v31 := v4[2147483647]
    var v32 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
