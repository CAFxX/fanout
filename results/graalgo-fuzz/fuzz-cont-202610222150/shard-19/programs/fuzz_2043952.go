package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int8; F1 int8; F2 int }
type GenS3 struct { F0 int16; F1 int32 }

func main() {
    var v1 int16 = 192
    var v2 uint16 = 508
    v3 := []int{}
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 bool = false
        _ = fi4
        _ = v5
    }
    v1 -= v1
    if true {
        if (!(!(!(float64(3.14159) > ((0.0 * 1.5) / 1.5))))) {
            fmt.Println(v2, v2)
            v1 -= (v1 + 42)
        } else {
            switch "\n\t" {
            case "":
                v1 = v1
            case "a":
                for fi6 := 0; fi6 < 16; fi6++ {
                    for fi7 := 0; fi7 < 15; fi7++ {
                        v2 -= v2
                        v1 -= v1
                        _ = fi7
                    }
                    _ = fi6
                }
            case "foo bar":
                v1 -= (v1 % ((v1) | 1))
            }
            for fi8 := 0; fi8 < 9; fi8++ {
                if true {
                    for fi9 := 0; fi9 < 6; fi9++ {
                        v2 = v2
                        v2 -= (^(v2))
                        if true {
                            if (!true) {
                                fmt.Println(fi8, fi8)
                            } else {
                                var v10 int32 = -2147483648
                                v3 = append(v3, (fi8 ^ (fi9 - (fi9 * (fi8 / (((fi8 << ((1) & 63))) | 1))))))
                                v2 = 2
                                _ = v10
                            }
                            v11 := GenS1{F0: 4294967295, F1: (v2 * (v2 & (v2 - (v2 >> ((1) & 63))))), F2: (^(v2))}
                            _ = v11
                        }
                        _ = fi9
                    }
                }
                var v12 float32 = ((-2.25 / 0.1) / 1e10)
                _ = fi8
                _ = v12
            }
            var v13 int = 128
            _ = v13
        }
        v14 := GenS1{F0: 0, F1: v2, F2: 42}
        _ = v14
    }
    var v15 uint = 348
    switch "\n\t" {
    case "":
        if true {
            v1 += int16(v15)
            v2 += v2
            fmt.Println(v1)
        }
    case "\n\t":
        v16 := GenS2{F0: -128, F1: 100, F2: 32767}
        _ = v16
    }
    var v17 int16 = 42
    fmt.Println(v2, v2)
    var v18 uint32 = 672
    v19 := GenS3{F0: (v1 * -721), F1: -1}
    v17 = -32768
    var v20 float32 = (1e10 * 0.1)
    v20 += (v20 + (float32(0.1) - (float32(((1e10 / -2.25) / (-2.25 / ((1e10 * ((0.1 * (1e10 - 3.14159)) * (-2.25 * 1.5))) + 3.14159)))) + v20)))
    var v21 uint64 = 508
    var v22 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
