package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint16, p1 uint8, p2 int64) uint16 {
    v7 := []int{0, 1, 32767, -964}
    fmt.Println(p1)
    v8 := GenS1{F0: 42, F1: 0, F2: 989}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 int8 = -1
    var v2 uint64 = 2
    var v3 int = 32768
    v4 := []string{"hello"}
    var v5 float64 = -2.25
    var v6 int8 = v1
    var v9 bool = (255 <= 2)
    var v10 int64 = 0
    var v11 uint = 745
    var v12 uint16 = 42
    var v13 uint16 = v12
    v13 = v12
    v10 -= (v10 / (((v10 | -1)) | 1))
    for fi14 := 0; fi14 < 16; fi14++ {
        v5 = v5
        switch 1 {
        case 128:
            switch "\n\t" {
            case "":
                var v15 uint8 = 255
                _ = v15
            case "hello":
                v16 := GenS2{F0: 2147483647}
                _ = v16
            case "foo bar":
                v6 = (v1 << (((^(v11))) & 63))
            }
        case 1:
            switch ("a" + "a") {
            case "a":
                if (3.14159 == (0.0 + float32(-2.25))) {
                    v10 = v10
                    var v17 uint32 = 702
                    _ = v17
                }
            case "\n\t":
                if true {
                    v18 := GenS3{F0: 2147483647, F1: (v1 * 127), F2: 128}
                    _ = v18
                }
            }
        }
        _ = fi14
    }
    fmt.Println(v12, v5)
    var v19 int32 = 0
    v4 = append(v4, "")
    v5 = 1.5
    v19 -= v19
    switch v3 {
    case 32768:
        fmt.Println(v6, v6)
    case 846:
        v20 := GenS2{F0: 581}
        _ = v20
    default:
        v5 += v5
    }
    v11 -= (v11 & (v11 - v11))
    if false {
        switch "foo bar" {
        case "\n\t":
            if false {
                v21 := GenS4{F0: 430}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v21
            } else {
                var v22 int32 = (v19 * v19)
                if true {
                    _ = len(v4)
                } else {
                    if (((!v9) || false) && (!false)) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if (1 <= -391) {
                            for fi23 := 0; fi23 < 18; fi23++ {
                                for fi24 := 0; fi24 < 5; fi24++ {
                                    if (v1 >= v1) {
                                        v3 = v3
                                        var v25 int8 = (v1 % ((-1) | 1))
                                        _ = v25
                                    }
                                    _ = len(v4)
                                    _ = fi24
                                }
                                _ = fi23
                            }
                            if v9 {
                                fmt.Println(v9, v6)
                            }
                        }
                    } else {
                        var v26 uint = 628
                        _ = len(v4)
                        _ = v26
                    }
                    v1 += v6
                }
                _ = v22
            }
        case "foo bar":
            var v27 bool = v9
            _ = v27
        }
        var v28 int16 = 1000
        _ = v28
    } else {
        v9 = ("foo bar" >= "")
        v29 := GenS3{F0: uint32(v11), F1: v1, F2: 255}
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
