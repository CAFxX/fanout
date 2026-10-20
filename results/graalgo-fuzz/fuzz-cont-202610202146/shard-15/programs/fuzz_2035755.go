package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 string; F2 int16 }
type GenS4 struct { F0 bool; F1 int32 }
func genFn0(p0 uint32, p1 string) uint64 {
    p1 = p1
    v5 := GenS1{F0: 58, F1: (p0 % ((p0) | 1)), F2: 0}
    _ = p0
    _ = p1
    _ = v5
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 1
    var v3 uint64 = 4294967295
    v4 := []int{}
    v6 := GenS1{F0: 0, F1: 599, F2: 9223372036854775807}
    if ((0.1 < 0.1) && (0 > 42)) {
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 3; fi8++ {
                delete(v1, 2)
                _ = fi8
            }
            var v9 int8 = 127
            var v10 uint32 = 1
            _ = fi7
            _ = v10
            _ = v9
        }
    }
    var v11 int8 = 31
    var v12 uint8 = 255
    if true {
        if true {
            v13 := GenS1{F0: (v12 ^ (v12 & 127)), F1: 2, F2: 175}
            var v14 int32 = -993
            _ = v13
            _ = v14
        } else {
            v12 = (v12 - (v12 * (v12 | (v12 >> (((v2 % ((v2) | 1))) & 63)))))
            switch -959 {
            case 32767:
                var v15 uint32 = 65535
                _ = v15
            }
            v11 = (v11 & 0)
        }
        v1[765] = false
    }
    if (1.5 >= 3.14159) {
        fmt.Println(v2, v11)
    }
    v16 := GenS2{F0: -2147483648}
    v17 := v1[2147483648]
    if v17 {
        if v17 {
            fmt.Println(v11, v2)
            delete(v1, 731)
        } else {
            switch "a" {
            case "hello":
                fmt.Println(v11)
            case "a":
                v18 := v1[65535]
                _ = v18
            default:
                v19 := GenS3{F0: int(v11), F1: "", F2: 116}
                _ = v19
            }
            var v20 int64 = -369
            _ = v20
        }
    } else {
        for fi21 := 0; fi21 < 17; fi21++ {
            var v22 uint8 = (v12 % ((v12) | 1))
            switch "" {
            case "foo bar":
                for fi23 := 0; fi23 < 3; fi23++ {
                    v24 := v1[0]
                    _ = fi23
                    _ = v24
                }
            default:
                var v25 bool = ((fi21 - 541) >= (fi21 * -9223372036854775808))
                _ = v25
            }
            if ("a" >= "foo bar") {
                fmt.Println(v22)
                v26 := v1[704]
                _ = v26
            } else {
                v22 += v22
                var v27 int32 = 0
                _ = v27
            }
            _ = fi21
            _ = v22
        }
        v3 += (v3 & v3)
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        var v29 float64 = (float64(1e10) / (float64((0.1 * 1.5)) * 1.5))
        v30 := GenS4{F0: false, F1: 1}
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println(v3, v3)
    switch 128 {
    case 42:
        fmt.Println(v17)
    case 32768:
        v11 = v11
    case 32767:
        v2 += 2
    }
    v3 += v3
    var v31 float64 = 1e10
    if v17 {
        v17 = (!(32767 < 717))
        v1[42] = v17
        if v17 {
            switch 604 {
            case 652:
                v32 := v1[0]
                _ = v32
            case -728:
                if (v17 && (!true)) {
                    if v17 {
                        v31 = v31
                        if true {
                            var v33 int64 = -758
                            var v34 bool = v17
                            fmt.Println(v12, v34)
                            _ = v33
                            _ = v34
                        }
                        v11 = 42
                    }
                    v35 := GenS1{F0: (v12 - v12), F1: uint32(v12), F2: 9223372036854775807}
                    var v36 uint32 = 690
                    _ = v35
                    _ = v36
                }
            default:
                for fi37 := 0; fi37 < 3; fi37++ {
                    fmt.Println(v11, v3)
                    for fi38 := 0; fi38 < 13; fi38++ {
                        v39 := GenS3{F0: -3, F1: ("\n\t" + (("\n\t" + "a") + "a")), F2: 1000}
                        _ = fi38
                        _ = v39
                    }
                    _ = fi37
                }
            }
            for fi40 := 0; fi40 < 5; fi40++ {
                fmt.Println(v31)
                _ = fi40
            }
        }
    } else {
        if (42 > -1) {
            for fi41 := 0; fi41 < 19; fi41++ {
                delete(v1, 65535)
                v2 -= (^(v2))
                v3 = 18446744073709551615
                _ = fi41
            }
        }
        for fi42 := 0; fi42 < 6; fi42++ {
            v43 := v1[62]
            for fi44 := 0; fi44 < 9; fi44++ {
                v12 -= v12
                fmt.Println(v11, v11, v2)
                v45 := GenS4{F0: false, F1: 941}
                _ = fi44
                _ = v45
            }
            _ = fi42
            _ = v43
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v6
}
