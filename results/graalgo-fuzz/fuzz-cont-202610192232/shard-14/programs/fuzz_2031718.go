package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 bool }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int16) int {
    var v5 string = ("" + "")
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(v5, v5)
        var v7 int64 = -1
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v5
    return 9223372036854775807
}

func main() {
    v1 := []int{-814, 2147483647, -210}
    var v2 int64 = -9223372036854775808
    var v3 uint8 = 255
    v4 := make(map[uint32]string)
    v3 += (^(v3))
    v8 := GenS1{F0: -2147483648, F1: 560, F2: (!(42 <= 0))}
    if false {
        v2 += (v2 & (^(v2)))
        for fi9 := 0; fi9 < 7; fi9++ {
            v3 = (v3 + v3)
            _ = fi9
        }
    }
    var v10 uint16 = 42
    switch "a" {
    case "\n\t":
        v11 := GenS2{F0: 1}
        _ = v11
    case "hello":
        if true {
            if (1 == 128) {
                for fi12 := 0; fi12 < 19; fi12++ {
                    var v13 int32 = int32(v10)
                    _ = fi12
                    _ = v13
                }
                v3 = 241
            } else {
                v14 := GenS2{F0: 4294967295}
                v15 := GenS2{F0: 42}
                _ = v14
                _ = v15
            }
            if (!true) {
                var v16 uint32 = 0
                var v17 uint32 = (v16 & 1)
                v18 := GenS2{F0: (v16 * (v16 + v17))}
                _ = v16
                _ = v17
                _ = v18
            }
        }
    }
    if false {
        var v19 int32 = -1
        v19 -= -439
        fmt.Println(v10)
        _ = v19
    }
    if false {
        fmt.Println(v10, v2)
        v2 = v2
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if ((v3 < (v3 - v3)) || (480 > 2)) {
            fmt.Println(v10)
        }
    }
    v3 += v3
    var v20 int16 = 0
    var v21 uint = 18446744073709551615
    if false {
        _ = len(v4)
    } else {
        switch 0 {
        case 166:
            var v22 int16 = 0
            _ = v22
        case 926:
            if false {
                for fi23 := 0; fi23 < 9; fi23++ {
                    var v24 int = genFn0(1)
                    var v25 float32 = 1e10
                    _ = len(v4)
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                for fi26 := 0; fi26 < 20; fi26++ {
                    if (v3 < v3) {
                        v3 += (v3 / ((255) | 1))
                        fmt.Println(v10, v2, v20)
                        var v27 float32 = ((float32(3.14159) * -2.25) * ((1.5 * 1e10) * 0.1))
                        _ = v27
                    }
                    if (v10 < v10) {
                        var v28 uint16 = (v10 * (^(v10)))
                        var v29 float32 = (-2.25 * (3.14159 + (1.5 * 0.1)))
                        _ = v28
                        _ = v29
                    } else {
                        if ((((3.14159 - 0.1) / (3.14159 * 1e10)) * (3.14159 + 1.5)) < 1.5) {
                            v30 := GenS2{F0: 284}
                            for fi31 := 0; fi31 < 18; fi31++ {
                                v10 -= 65535
                                if len(v1) > 0 {
                                    v1[0] = (fi31 % (((fi26 | fi31)) | 1))
                                }
                                _ = fi31
                            }
                            _ = v30
                        } else {
                            switch fi26 {
                            case 128:
                                v21 += v21
                            case 127:
                                if len(v1) > 0 {
                                    v1[0] = fi26
                                }
                            }
                            for fi32 := 0; fi32 < 5; fi32++ {
                                v1 = append(v1, int(v3))
                                _ = fi32
                            }
                        }
                    }
                    _ = fi26
                }
            } else {
                v33 := GenS2{F0: 4294967295}
                var v34 int8 = -92
                var v35 int8 = (^(v34))
                _ = v33
                _ = v34
                _ = v35
            }
        case 9223372036854775807:
            var v36 float32 = 3.14159
            _ = v36
        }
    }
    v10 -= (v10 ^ (^(v10)))
    fmt.Println(v21)
    v10 += (^(v10))
    switch -459 {
    case 4294967295:
        v21 = (v21 >> (((v21 - (v21 + (v21 & (^(v21)))))) & 63))
    case 1:
        v20 = (v20 / ((v20) | 1))
    case 32768:
        if (74 <= 2147483647) {
            fmt.Println(v10)
            v3 += v3
        }
    default:
        fmt.Println(v20)
    }
    var v37 int = 626
    v38 := GenS2{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v8
}
