package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint64 }
func genFn0() uint64 {
    var v6 float64 = (1.5 * (float64((float32(0.0) / 1.5)) + -2.25))
    for fi7 := 0; fi7 < 3; fi7++ {
        if true {
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 int16 = -32768
                if (true || false) {
                    switch "" {
                    case "hello":
                        v10 := GenS1{F0: 128, F1: v9}
                        _ = v10
                    default:
                        switch (fi7 % (((^(fi7))) | 1)) {
                        case 128:
                            v11 := GenS1{F0: 128, F1: (v9 / (((v9 & v9)) | 1))}
                            _ = v11
                        default:
                            fmt.Println(v9)
                        }
                    }
                    if (!false) {
                        if (!((-2147483648 <= 17) || true)) {
                            switch "" {
                            case "foo bar":
                                var v12 int8 = 100
                                _ = v12
                            case "\n\t":
                                if (!true) {
                                    v13 := GenS1{F0: 104, F1: v9}
                                    _ = v13
                                }
                            }
                        } else {
                            v9 += (v9 * (v9 % (((v9 * v9)) | 1)))
                        }
                        v9 += (v9 - 1)
                        v14 := []int{fi7, (fi8 / (((fi8 % ((0) | 1))) | 1)), (fi8 >> ((255) & 63))}
                        _ = v14
                    }
                }
                _ = fi8
                _ = v9
            }
            v15 := []int{(fi7 * (fi7 * 0))}
            _ = v15
        } else {
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 int16 = 0
                v6 = -2.25
                _ = fi16
                _ = v17
            }
            v6 += v6
            v6 = v6
        }
        _ = fi7
    }
    _ = v6
    return 42
}

func main() {
    var v1 int16 = -32768
    v2 := []uint32{2147483648, 0, 4294967295}
    var v3 uint64 = 518
    v4 := []int{603, 229}
    v5 := make(map[int]int)
    v3 = v3
    for fi18 := 0; fi18 < 18; fi18++ {
        v3 = v3
        var v19 float64 = 0.0
        v19 = v19
        _ = fi18
        _ = v19
    }
    if false {
        v3 -= 2
        fmt.Println(v1, v1)
        switch -464 {
        case 2147483647:
            if (0 == v3) {
                var v20 uint32 = 2147483648
                _ = v20
            }
        default:
            v21 := v5[-597]
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi22
    }
    v3 += (v3 ^ (^(v3)))
    var v23 uint32 = 2147483648
    var v24 int32 = 0
    v23 = 2147483647
    switch 32767 {
    case 4294967295:
        for fi25 := 0; fi25 < 3; fi25++ {
            var v26 uint8 = 192
            var v27 uint16 = 821
            _ = fi25
            _ = v26
            _ = v27
        }
    case -328:
        v24 -= v24
    default:
        switch 128 {
        case -9223372036854775808:
            switch (("\n\t" + "a") + "\n\t") {
            case "hello":
                v28 := GenS2{F0: 273}
                _ = v28
            case "\n\t":
                for fi29 := 0; fi29 < 15; fi29++ {
                    v2 = append(v2, (v23 ^ (v23 ^ (v23 % ((v23) | 1)))))
                    _ = fi29
                }
            default:
                var v30 uint16 = 0
                _ = v30
            }
        default:
            v24 += (v24 + (v24 & 2147483647))
        }
    }
    fmt.Println(v23)
    if (!true) {
        fmt.Println(v1, v24, v3)
    } else {
        if (true && true) {
            if false {
                var v31 float64 = 1.5
                _ = v31
            }
        } else {
            fmt.Println(v3, v24, v1)
            delete(v5, 2147483647)
        }
    }
    v1 += 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
