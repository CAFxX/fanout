package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 int32; F2 bool }
type GenS5 struct { F0 uint; F1 bool }
func genFn0(p0 uint32, p1 uint64) uint8 {
    switch "foo bar" {
    case "hello":
        if (65535 != 32767) {
            fmt.Println(p1, p1)
            v4 := GenS1{F0: 100, F1: (p1 - (p1 | (^(p1)))), F2: p1}
            _ = v4
        }
    case "a":
        v5 := GenS1{F0: 127, F1: 18446744073709551615, F2: 0}
        _ = v5
    case "":
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := GenS1{F0: 127, F1: 758, F2: 193}
            _ = fi6
            _ = v7
        }
    default:
        for fi8 := 0; fi8 < 1; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                v10 := []int{fi8, (fi8 >> ((255) & 63))}
                var v11 bool = true
                var v12 uint8 = 1
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            var v13 float64 = 1e10
            p1 += p1
            _ = fi8
            _ = v13
        }
    }
    v14 := []int{4294967295, 127, 228}
    p0 = 65535
    for fi15 := 0; fi15 < 16; fi15++ {
        if (!(3.14159 < float64(float32((0.1 + float64((0.1 / -2.25))))))) {
            for fi16 := 0; fi16 < 20; fi16++ {
                p1 += (^(p1))
                fmt.Println(p0, p1)
                var v17 uint = 447
                _ = fi16
                _ = v17
            }
            v18 := GenS2{F0: (!(0 != 867))}
            if ((("" + (("" + "\n\t") + (("" + "\n\t") + "\n\t"))) >= ("\n\t" + ("\n\t" + ""))) || false) {
                var v19 int = fi15
                for fi20 := 0; fi20 < 15; fi20++ {
                    _ = len(v14)
                    fmt.Println(p0)
                    _ = fi20
                }
                _ = v19
            } else {
                for fi21 := 0; fi21 < 2; fi21++ {
                    p1 = (p1 / (((p1 >> ((255) & 63))) | 1))
                    if false {
                        var v22 string = "hello"
                        if ((false || (!false)) || false) {
                            v23 := GenS3{F0: (p1 - p1)}
                            var v24 uint16 = 65535
                            _ = v23
                            _ = v24
                        }
                        _ = v22
                    }
                    _ = fi21
                }
                var v25 int32 = -2147483648
                _ = v25
            }
            _ = v18
        }
        v26 := GenS1{F0: 100, F1: p1, F2: p1}
        _ = fi15
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v14
    return 130
}

func main() {
    var v1 int32 = 1
    var v2 int32 = (v1 >> ((18446744073709551615) & 63))
    v3 := []int{799, -827}
    fmt.Println(v1, v2)
    switch "a" {
    case "a":
        v2 = (^(v2))
    case "":
        v2 += (v1 % (((v2 >> ((747) & 63))) | 1))
    }
    fmt.Println(v1)
    _ = len(v3)
    var v27 uint = 65535
    if true {
        for fi28 := 0; fi28 < 11; fi28++ {
            var v29 int64 = 0
            _ = fi28
            _ = v29
        }
    } else {
        fmt.Println(v27, v2, v2)
        v2 = int32(v27)
    }
    v2 -= v1
    switch "a" {
    case "hello":
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    case "a":
        v30 := GenS4{F0: 127, F1: (^(v2)), F2: ((-(v2)) < (v1 ^ v1))}
        _ = v30
    default:
        v27 = (v27 | 65535)
    }
    v27 += 581
    if (!false) {
        v31 := GenS1{F0: 100, F1: 2, F2: 2}
        var v32 float64 = 0.1
        _ = v31
        _ = v32
    }
    v33 := GenS5{F0: v27, F1: (int8(v2) > 1)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
}
