package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 bool = (42 != 1)
    v2 := make(map[int]bool)
    var v3 uint16 = 32767
    switch 2147483647 {
    case 42:
        var v4 int64 = 1
        _ = v4
    }
    delete(v2, -321)
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: 42, F1: 639}
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 18446744073709551615, F1: 895}
    v8 := v2[9223372036854775807]
    var v9 uint64 = 776
    if v8 {
        v2[-9223372036854775808] = false
    } else {
        fmt.Println(v3, v9, v1)
        switch -395 {
        case 32768:
            v10 := GenS1{F0: (v9 - v9), F1: 0}
            _ = v10
        case 2147483647:
            var v11 uint64 = (^(v9))
            _ = v11
        }
        if v1 {
            switch 127 {
            case 127:
                fmt.Println(v8, v9)
            case -347:
                if ((v9 & (v9 & 4294967295)) == (^(v9))) {
                    v8 = (true || (-996 > 1000))
                    v1 = true
                    switch 4294967295 {
                    case 128:
                        switch 1 {
                        case 9223372036854775807:
                            var v12 bool = ((1e10 / ((1.5 + ((0.0 * (((((3.14159 + 1.5) / 1e10) - (float32(-2.25) + ((3.14159 + 3.14159) / float32(1.5)))) / ((0.1 + 1e10) + 0.0)) - -2.25)) / -2.25)) / 3.14159)) != 0.0)
                            _ = v12
                        case 526:
                            if (!("a" != "foo bar")) {
                                var v13 uint16 = (v3 / (((v3 / ((32767) | 1))) | 1))
                                v14 := GenS2{F0: 0, F1: v8, F2: 160}
                                v15 := []int{-1, 127, 309, 1}
                                _ = v13
                                _ = v14
                                _ = v15
                            } else {
                                v16 := GenS3{F0: 391}
                                _ = v16
                            }
                        case 32767:
                            v17 := GenS3{F0: 1}
                            _ = v17
                        }
                    case 453:
                        v1 = (v3 == v3)
                    default:
                        for fi18 := 0; fi18 < 1; fi18++ {
                            v19 := []int{fi18, (-(fi18))}
                            _ = fi18
                            _ = v19
                        }
                    }
                } else {
                    v8 = (uint32(v9) >= 2)
                    v8 = false
                }
            default:
                v20 := []int{9223372036854775807, 32768}
                _ = v20
            }
        }
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        v21 := []int{4294967295, 0}
        _ = v21
    default:
        v1 = v1
    }
    v22 := GenS2{F0: 65535, F1: v8, F2: v3}
    for fi23 := 0; fi23 < 7; fi23++ {
        v24 := GenS2{F0: 1, F1: (v3 == (^(v3))), F2: v3}
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        v9 = v9
        _ = fi25
    }
    if (v9 == v9) {
        if ("hello" <= "a") {
            fmt.Println(v3)
            v26 := []int{9223372036854775807}
            var v27 int = 953
            _ = v26
            _ = v27
        } else {
            v3 += 356
            var v28 uint32 = 314
            _ = v28
        }
    }
    switch 189 {
    case -9223372036854775808:
        var v29 int64 = 489
        _ = v29
    case 32767:
        v30 := GenS2{F0: 42, F1: v8, F2: (v3 >> ((255) & 63))}
        _ = v30
    case 128:
        fmt.Println(v9, v8, v9)
    }
    var v31 int64 = 1
    for fi32 := 0; fi32 < 14; fi32++ {
        _ = len(v2)
        _ = fi32
    }
    delete(v2, 1)
    v31 -= v31
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v7
    _ = v8
    _ = v9
}
