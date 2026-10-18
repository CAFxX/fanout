package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 bool }
func genFn0(p0 bool, p1 uint8, p2 uint32) int64 {
    p0 = p0
    p1 -= 25
    p1 = p1
    v6 := GenS1{F0: 255, F1: p0}
    v7 := GenS2{F0: 0, F1: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -436
}
func genFn1() uint64 {
    var v8 bool = (255 == 1)
    v9 := []int{-207, 2147483647}
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 int64 = -9223372036854775808
        _ = fi10
        _ = v11
    }
    v12 := GenS3{F0: -114, F1: -855}
    var v13 int32 = 2147483647
    _ = v12
    _ = v13
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 int64 = -1
    v2 := make(map[int]bool)
    var v3 int8 = -128
    v4 := make(map[uint32]bool)
    var v5 uint16 = 312
    for fi14 := 0; fi14 < 16; fi14++ {
        v3 = 0
        v15 := GenS4{F0: false}
        switch "" {
        case "\n\t":
            if (float64(0.0) > 3.14159) {
                var v16 bool = ((^(v3)) == 42)
                var v17 int = fi14
                v18 := []int{}
                _ = v16
                _ = v17
                _ = v18
            } else {
                v19 := []int{(fi14 / ((9223372036854775807) | 1))}
                _ = v19
            }
        case "a":
            switch "a" {
            case "hello":
                if false {
                    var v20 int8 = (v3 ^ (v3 * (v3 + v3)))
                    v2[(fi14 & (fi14 - -9223372036854775808))] = (float32(1.5) < 1e10)
                    _ = v20
                }
            }
        case "":
            for fi21 := 0; fi21 < 3; fi21++ {
                var v22 int16 = 1
                _ = fi21
                _ = v22
            }
        default:
            if false {
                var v23 uint16 = (v5 & (v5 & (v5 | 32767)))
                switch (fi14 / ((fi14) | 1)) {
                case 2147483647:
                    if true {
                        var v24 uint64 = 42
                        if true {
                            if false {
                                _ = len(v4)
                                v25 := []int{}
                                _ = v25
                            }
                        } else {
                            var v26 float32 = 0.0
                            var v27 uint32 = 474
                            _ = v26
                            _ = v27
                        }
                        _ = v24
                    } else {
                        var v28 int = (fi14 | (fi14 / (((fi14 << ((91) & 63))) | 1)))
                        _ = v28
                    }
                case 32768:
                    v5 += v23
                case 4294967295:
                    switch -479 {
                    case 32768:
                        v29 := v2[410]
                        _ = v29
                    default:
                        v3 = 42
                    }
                }
                if (!(!false)) {
                    if ((!false) && true) {
                        v30 := []int{128, 800}
                        _ = v30
                    }
                } else {
                    if true {
                        v31 := GenS2{F0: v3, F1: 0}
                        _ = v31
                    }
                }
                _ = v23
            }
        }
        _ = fi14
        _ = v15
    }
    var v32 uint32 = 4294967295
    v32 = (^(v32))
    for fi33 := 0; fi33 < 17; fi33++ {
        v5 -= 589
        var v34 int16 = -32768
        v35 := []int{fi33, (fi33 * fi33), (fi33 & (^(fi33)))}
        _ = fi33
        _ = v34
        _ = v35
    }
    if true {
        v1 += 704
        v36 := []int{-9223372036854775808, 9223372036854775807}
        _ = v36
    } else {
        var v37 int8 = (v3 & (v3 / (((v3 - 127)) | 1)))
        _ = v37
    }
    v3 += v3
    v1 = 0
    var v38 int8 = (^(v3))
    for fi39 := 0; fi39 < 7; fi39++ {
        for fi40 := 0; fi40 < 18; fi40++ {
            if false {
                var v41 bool = ((v1 ^ v1) != v1)
                var v42 uint64 = 0
                _ = v41
                _ = v42
            }
            v5 -= (v5 & (v5 - v5))
            _ = fi40
        }
        if (!(!true)) {
            v43 := []int{(fi39 - fi39), (fi39 % (((fi39 * -9223372036854775808)) | 1))}
            _ = v43
        } else {
            if true {
                for fi44 := 0; fi44 < 19; fi44++ {
                    v45 := GenS1{F0: v5, F1: false}
                    v38 -= (v38 | (v3 ^ v38))
                    _ = fi44
                    _ = v45
                }
                var v46 uint8 = 186
                _ = v46
            }
            var v47 int64 = v1
            if true {
                v3 = (v38 + (^(v38)))
            } else {
                v5 = (v5 << ((2) & 63))
                var v48 string = "\n\t"
                _ = v48
            }
            _ = v47
        }
        var v49 uint16 = v5
        _ = fi39
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v4
    _ = v5
}
