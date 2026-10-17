package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int32; F2 uint16 }

func main() {
    var v1 int16 = -1
    var v2 uint = 42
    var v3 uint16 = 377
    v4 := make(map[uint32]int)
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 int16 = (v1 % (((v1 ^ v1)) | 1))
        _ = fi5
        _ = v6
    }
    v1 -= (v1 << ((v2) & 63))
    switch "hello" {
    case "":
        for fi7 := 0; fi7 < 11; fi7++ {
            if (3.14159 < (1.5 + (1.5 - (float64(0.0) - 1.5)))) {
                v8 := GenS1{F0: (v2 * (v2 * (v2 / ((v2) | 1)))), F1: 132, F2: 42}
                v9 := []int{(-(fi7)), fi7, (fi7 | fi7)}
                var v10 int16 = v1
                _ = v10
                _ = v8
                _ = v9
            }
            if (504 >= v2) {
                fmt.Println(fi7)
            }
            switch "\n\t" {
            case "hello":
                v11 := []int{fi7, (fi7 | (fi7 << ((v2) & 63))), fi7}
                _ = v11
            case "\n\t":
                var v12 uint = (v2 + (v2 - v2))
                _ = v12
            case "a":
                if (!(true && (false && (18446744073709551615 < 539)))) {
                    v13 := []int{(^(fi7)), 32768, (fi7 ^ fi7)}
                    if false {
                        var v14 string = "\n\t"
                        fmt.Println(v14, v2, v2)
                        var v15 float64 = 3.14159
                        _ = v14
                        _ = v15
                    }
                    if false {
                        for fi16 := 0; fi16 < 12; fi16++ {
                            v13 = append(v13, (^(fi7)))
                            fmt.Println(fi7)
                            var v17 uint32 = 2
                            _ = fi16
                            _ = v17
                        }
                    }
                    _ = v13
                } else {
                    for fi18 := 0; fi18 < 5; fi18++ {
                        if false {
                            switch ("a" + (("a" + "") + "")) {
                            case "hello":
                                v19 := GenS2{F0: 42, F1: -2147483648, F2: 1}
                                _ = v19
                            case "":
                                var v20 int64 = 1
                                _ = v20
                            }
                        } else {
                            if (!(!(!(true && (true || (127 != 42)))))) {
                                var v21 int = fi7
                                _ = v21
                            } else {
                                v22 := GenS1{F0: (v2 + (v2 ^ 42)), F1: v3, F2: 224}
                                v23 := []int{(fi7 / ((int(v2)) | 1)), (fi18 | (^(fi18)))}
                                v2 = v2
                                _ = v22
                                _ = v23
                            }
                            v24 := []int{fi18, 916}
                            var v25 uint = (v2 >> ((v2) & 63))
                            v3 = 0
                            _ = v24
                            _ = v25
                        }
                        _ = fi18
                    }
                }
            }
            _ = fi7
        }
    default:
        for fi26 := 0; fi26 < 11; fi26++ {
            v27 := []int{127}
            _ = fi26
            _ = v27
        }
    }
    v1 = (v1 * v1)
    var v28 bool = (!false)
    v29 := GenS1{F0: 2, F1: 1, F2: 65535}
    v30 := GenS2{F0: 128, F1: 1, F2: 824}
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
