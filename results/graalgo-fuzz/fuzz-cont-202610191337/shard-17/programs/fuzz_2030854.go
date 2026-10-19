package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    v3 := make(map[int]bool)
    var v4 uint32 = 2
    v5 := make(map[string]string)
    fmt.Println(v4)
    var v6 uint8 = 127
    v7 := []int{-2147483648}
    v8 := GenS1{F0: 42, F1: 149, F2: 454}
    fmt.Println(v4, v6)
    if (614 > 42) {
        v9 := GenS1{F0: 42, F1: 222, F2: 1000}
        switch 32767 {
        case 699:
            for fi10 := 0; fi10 < 3; fi10++ {
                if true {
                    v4 = (v4 - v4)
                    v6 -= (v6 & v6)
                    v6 -= (v6 / ((127) | 1))
                } else {
                    v7 = append(v7, fi10)
                }
                _ = fi10
            }
        case 128:
            v4 = (v4 * (^(v4)))
        default:
            var v11 uint64 = 250
            _ = v11
        }
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS1{F0: -103, F1: 0, F2: 1}
            v14 := v2[fi12]
            if v14 {
                if (!(32767 > 65535)) {
                    if (-147 > 177) {
                        v14 = ((0.1 != 1.5) && (!v14))
                    }
                    v15 := v5[("foo bar" + "foo bar")]
                    v2[-1] = v14
                    _ = v15
                }
            } else {
                fmt.Println(v6, v14, v6)
                if true {
                    if (!v14) {
                        v16 := GenS1{F0: 42, F1: 384, F2: 42}
                        _ = v16
                    }
                } else {
                    v14 = v14
                    for fi17 := 0; fi17 < 18; fi17++ {
                        fmt.Println(v6)
                        _ = fi17
                    }
                }
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v9
    }
    var v18 uint = 197
    var v19 uint64 = uint64(v18)
    for fi20 := 0; fi20 < 20; fi20++ {
        delete(v3, 2147483647)
        var v21 uint = (v18 ^ (v18 % ((uint(v19)) | 1)))
        if ((^(v21)) < v18) {
            for fi22 := 0; fi22 < 3; fi22++ {
                var v23 int64 = -1
                v7 = append(v7, (fi20 & (fi20 % (((fi22 + 42)) | 1))))
                _ = fi22
                _ = v23
            }
        }
        _ = fi20
        _ = v21
    }
    switch 32768 {
    case 9223372036854775807:
        v4 += (v4 ^ v4)
    }
    _ = len(v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
