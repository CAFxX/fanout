package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 int }
func genFn0(p0 uint8, p1 uint64, p2 uint64) int {
    fmt.Println(p1, p1, p1)
    var v5 bool = false
    v6 := []int{}
    switch 32767 {
    case -9223372036854775808:
        fmt.Println(p2, v5)
    default:
        var v7 int8 = 42
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 uint32 = 2147483647
    var v2 int64 = -9223372036854775808
    var v3 float32 = 1.5
    var v4 int64 = (v2 % (((-(v2))) | 1))
    var v8 float32 = v3
    var v9 int32 = 42
    var v10 uint16 = uint16(v1)
    v3 = float32(1.5)
    var v11 uint64 = 0
    v12 := GenS1{F0: 100}
    v13 := []int{2147483647, 9223372036854775807, genFn0(42, v11, v11), 127}
    v3 = v3
    var v14 int16 = 645
    fmt.Println(v3, v14)
    var v15 uint32 = (v1 << ((340) & 63))
    v16 := GenS2{F0: 1}
    for fi17 := 0; fi17 < 9; fi17++ {
        v13 = append(v13, (fi17 | (fi17 % (((fi17 | fi17)) | 1))))
        if ((1.5 == 0.0) || true) {
            var v18 int8 = 127
            if ((2 <= 7) && (false && true)) {
                switch 127 {
                case -1:
                    v1 += 118
                case 76:
                    for fi19 := 0; fi19 < 16; fi19++ {
                        v13 = append(v13, (fi17 ^ genFn0(75, 108, 2)))
                        _ = fi19
                    }
                case 485:
                    switch "foo bar" {
                    case "a":
                        switch int(v10) {
                        case -465:
                            if true {
                                v1 += v15
                                for fi20 := 0; fi20 < 4; fi20++ {
                                    var v21 uint32 = (^(v1))
                                    _ = fi20
                                    _ = v21
                                }
                                v18 -= v18
                            } else {
                                v11 += 42
                            }
                        default:
                            var v22 int32 = 0
                            _ = v22
                        }
                    case "":
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    case "\n\t":
                        var v23 uint64 = v11
                        _ = v23
                    default:
                        if false {
                            v24 := GenS3{F0: (v18 % ((v18) | 1)), F1: fi17}
                            if ((fi17 + genFn0(127, (v11 + (v11 % (((^(v11))) | 1))), (v11 >> ((667) & 63)))) < (fi17 & 9223372036854775807)) {
                                v8 = v8
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                var v25 uint = 0
                                _ = v25
                            }
                            _ = v24
                        }
                    }
                }
            } else {
                var v26 int16 = v14
                fmt.Println(v26, v8)
                _ = v26
            }
            _ = v18
        }
        fmt.Println(v9, v3, v10)
        _ = fi17
    }
    v27 := GenS3{F0: 100, F1: -9223372036854775808}
    v28 := GenS2{F0: 65}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
