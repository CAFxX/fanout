package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint16 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = -2.25
    v3 := make(map[int]bool)
    var v4 uint32 = 2
    var v5 uint16 = 255
    if (255 != 2) {
        for fi6 := 0; fi6 < 17; fi6++ {
            fmt.Println(fi6)
            v7 := v3[(fi6 / ((fi6) | 1))]
            _ = fi6
            _ = v7
        }
        delete(v3, 9223372036854775807)
        switch 766 {
        case 9223372036854775807:
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := []int{fi8}
                _ = fi8
                _ = v9
            }
        case 0:
            var v10 uint8 = 201
            _ = v10
        default:
            if (-9223372036854775808 > 9223372036854775807) {
                v5 -= v5
                if (!(!(2147483647 > -360))) {
                    v5 = v5
                    var v11 string = ""
                    var v12 float32 = v2
                    _ = v11
                    _ = v12
                } else {
                    if false {
                        v13 := []int{32768}
                        var v14 uint = 4294967295
                        v14 += 42
                        _ = v13
                        _ = v14
                    } else {
                        fmt.Println(v4, v4, v2)
                        var v15 int16 = 965
                        v16 := GenS1{F0: 2, F1: 215, F2: v5}
                        _ = v15
                        _ = v16
                    }
                    v17 := GenS1{F0: 1, F1: 255, F2: 1}
                    switch 700 {
                    case 788:
                        if false {
                            v18 := GenS1{F0: 18446744073709551615, F1: 42, F2: (^(v5))}
                            var v19 int16 = 32767
                            for fi20 := 0; fi20 < 16; fi20++ {
                                var v21 uint32 = 124
                                v4 -= (v4 % ((v4) | 1))
                                _ = fi20
                                _ = v21
                            }
                            _ = v18
                            _ = v19
                        } else {
                            v5 += v5
                        }
                    }
                    _ = v17
                }
            }
        }
    }
    delete(v3, 127)
    v2 = v2
    v4 += (v4 * v4)
    var v22 uint = 0
    fmt.Println(v4, v22, v2)
    switch -9223372036854775808 {
    case 32767:
        v22 -= (v22 >> ((v22) & 63))
    case 42:
        v23 := GenS1{F0: 0, F1: 120, F2: 32767}
        _ = v23
    default:
        v24 := GenS1{F0: 42, F1: 42, F2: 425}
        _ = v24
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        var v26 uint32 = v4
        _ = fi25
        _ = v26
    }
    v22 = 211
    var v27 int = -2147483648
    for fi28 := 0; fi28 < 14; fi28++ {
        v29 := GenS1{F0: 216, F1: 160, F2: 42}
        switch ("" + "\n\t") {
        case "a":
            for fi30 := 0; fi30 < 11; fi30++ {
                v4 = 0
                _ = fi30
            }
        case "":
            var v31 uint64 = 18446744073709551615
            _ = v31
        case "\n\t":
            var v32 uint16 = 42
            _ = v32
        }
        _ = fi28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
