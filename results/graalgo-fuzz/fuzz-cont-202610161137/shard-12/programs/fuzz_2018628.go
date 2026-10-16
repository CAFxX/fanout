package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint16; F1 uint; F2 uint32 }
func genFn0(p0 uint8) int16 {
    switch "\n\t" {
    case "\n\t":
        var v4 float64 = (3.14159 * 1e10)
        _ = v4
    case "foo bar":
        p0 -= (p0 * 127)
    }
    v5 := []int{}
    _ = p0
    _ = v5
    return -32768
}
func genFn1(p0 uint, p1 uint8) int {
    p1 += p1
    var v6 uint = 130
    var v7 uint64 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 637
}

func main() {
    var v1 uint16 = 96
    v2 := make(map[uint32]bool)
    v3 := []int{}
    if true {
        v1 -= v1
        v8 := GenS1{F0: 75, F1: 23, F2: 127}
        _ = v8
    }
    if len(v3) > 0 {
        v3[0] = 872
    }
    var v9 int16 = -32768
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(fi10)
        _ = fi10
    }
    if false {
        if true {
            v11 := GenS1{F0: -64, F1: (-(v9)), F2: 76}
            if false {
                v3 = append(v3, -121)
            } else {
                for fi12 := 0; fi12 < 11; fi12++ {
                    fmt.Println(v9, v9)
                    var v13 int8 = 100
                    _ = fi12
                    _ = v13
                }
            }
            switch 42 {
            case 42:
                for fi14 := 0; fi14 < 11; fi14++ {
                    _ = len(v2)
                    v9 = -983
                    _ = fi14
                }
            case 2147483647:
                switch "" {
                case "hello":
                    if (3.14159 < 3.14159) {
                        v2[2147483647] = (1.5 > -2.25)
                        v1 += 42
                        for fi15 := 0; fi15 < 13; fi15++ {
                            v9 += v9
                            _ = fi15
                        }
                    }
                }
            case 349:
                fmt.Println(v1, v9)
            default:
                var v16 uint16 = (v1 / ((855) | 1))
                _ = v16
            }
            _ = v11
        }
        var v17 uint32 = 4294967295
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS2{F0: 255, F1: v9}
            v20 := GenS3{F0: 1, F1: 1}
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = v17
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        v22 := GenS3{F0: -571, F1: 167}
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        if ((fi23 / (((fi23 >> ((65535) & 63))) | 1)) >= (fi23 + 128)) {
            v2[2] = false
            fmt.Println(v9, fi23)
            v24 := GenS1{F0: -107, F1: v9, F2: 42}
            _ = v24
        } else {
            v1 -= 255
            if false {
                var v25 int64 = 9223372036854775807
                _ = v25
            }
            fmt.Println(v9)
        }
        fmt.Println(v9)
        _ = fi23
    }
    var v26 float64 = 1e10
    var v27 int = -1
    var v28 int8 = 100
    if true {
        v27 += v27
        v29 := GenS2{F0: 123, F1: (v9 << ((18446744073709551615) & 63))}
        switch 2147483647 {
        case 1:
            v30 := GenS4{F0: true, F1: 65535, F2: (v1 & 255)}
            _ = v30
        default:
            switch (^(v27)) {
            case 53:
                var v31 uint32 = 2147483648
                _ = v31
            case -9223372036854775808:
                v32 := GenS2{F0: 18446744073709551615, F1: 32767}
                _ = v32
            }
        }
        _ = v29
    }
    v28 += (v28 & (^(v28)))
    for fi33 := 0; fi33 < 5; fi33++ {
        var v34 string = ""
        _ = fi33
        _ = v34
    }
    v35 := GenS5{F0: (^(v1)), F1: 65535, F2: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v9
}
