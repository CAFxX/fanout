package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int }

func main() {
    var v1 uint = 4294967295
    var v2 bool = false
    var v3 bool = v2
    if (false && (1 <= 20)) {
        v3 = true
        for fi4 := 0; fi4 < 13; fi4++ {
            v1 += v1
            _ = fi4
        }
    } else {
        v5 := []int{32767, 949, 370, -9223372036854775808}
        if v2 {
            if true {
                _ = len(v5)
            }
        } else {
            for fi6 := 0; fi6 < 2; fi6++ {
                switch "\n\t" {
                case "a":
                    v3 = (25 < 32767)
                default:
                    v7 := GenS1{F0: 42}
                    _ = v7
                }
                switch (fi6 + 42) {
                case 128:
                    fmt.Println(fi6, v1)
                case 696:
                    if (100 >= 100) {
                        fmt.Println(v3, v3, fi6)
                        v1 = 1
                        var v8 uint = (v1 / (((v1 / ((v1) | 1))) | 1))
                        _ = v8
                    } else {
                        for fi9 := 0; fi9 < 18; fi9++ {
                            v10 := GenS1{F0: 18446744073709551615}
                            _ = fi9
                            _ = v10
                        }
                    }
                case 82:
                    fmt.Println(v3, v3)
                }
                _ = fi6
            }
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 uint32 = 2
                switch "" {
                case "":
                    v13 := GenS1{F0: 582}
                    _ = v13
                case "foo bar":
                    v12 = (v12 ^ v12)
                default:
                    v2 = (v12 > v12)
                }
                _ = fi11
                _ = v12
            }
            _ = len(v5)
        }
        v14 := GenS2{F0: -119, F1: 141}
        _ = v14
        _ = v5
    }
    if v3 {
        var v15 int64 = 9223372036854775807
        var v16 int32 = -2147483648
        _ = v15
        _ = v16
    } else {
        v17 := []int{-899}
        _ = v17
    }
    switch 128 {
    case 42:
        for fi18 := 0; fi18 < 1; fi18++ {
            v3 = (803 < 0)
            _ = fi18
        }
    case 54:
        v1 -= v1
    }
    v2 = false
    v2 = v3
    if (!v2) {
        v19 := GenS2{F0: 127, F1: 52}
        _ = v19
    }
    v20 := []int{128, -1, 127}
    if v2 {
        var v21 int32 = -569
        _ = v21
    }
    v3 = (v2 && false)
    v22 := GenS3{F0: 1, F1: 32768}
    if v2 {
        if len(v20) > 0 {
            v20[0] = 32767
        }
        var v23 float32 = 1.5
        _ = v23
    } else {
        v2 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
}
