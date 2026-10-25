package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 uint }
func genFn0(p0 bool, p1 uint8) string {
    v5 := GenS1{F0: 2147483647, F1: 430, F2: 2147483647}
    switch 4294967295 {
    case 53:
        var v6 uint32 = 2147483648
        _ = v6
    case 1:
        if (p0 || (255 <= 797)) {
            p1 += 127
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return ("a" + "hello")
}
func genFn1(p0 uint64, p1 int16) int8 {
    switch "hello" {
    case "a":
        if true {
            v7 := []int{2147483647}
            fmt.Println(p0, p1)
            _ = v7
        } else {
            if false {
                p0 -= (^(p0))
            } else {
                var v8 uint = 65535
                _ = v8
            }
            v9 := []int{127, 970, 0, 2147483647}
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = v9
        }
    case "foo bar":
        v10 := GenS1{F0: 42, F1: 0, F2: 32767}
        _ = v10
    case "hello":
        if true {
            p0 = 2
        }
    }
    v11 := []int{-1}
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := GenS2{F0: 372, F1: -564}
        var v14 float64 = 1.5
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v11
    return 127
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    var v4 int16 = 32767
    if ((751 > 2147483648) && false) {
        if (v4 >= 42) {
            v15 := GenS3{F0: 2147483647}
            for fi16 := 0; fi16 < 19; fi16++ {
                v4 += (v4 & -32768)
                var v17 int16 = v4
                _ = fi16
                _ = v17
            }
            _ = v15
        } else {
            switch 670 {
            case 32767:
                v4 = v4
            case 508:
                for fi18 := 0; fi18 < 7; fi18++ {
                    var v19 uint16 = 255
                    _ = fi18
                    _ = v19
                }
            }
            var v20 uint64 = 526
            v20 -= 4294967295
            _ = v20
        }
        v21 := v3["a"]
        v4 = (v4 >> ((65535) & 63))
        _ = v21
    } else {
        var v22 int16 = (v4 >> ((0) & 63))
        _ = v22
    }
    if false {
        v4 -= v4
    } else {
        v23 := GenS1{F0: -228, F1: 2147483647, F2: 128}
        fmt.Println(v4)
        _ = v23
    }
    v24 := GenS4{F0: 1}
    if false {
        _ = len(v2)
    }
    fmt.Println(v4)
    var v25 int = -194
    v25 += 4294967295
    v26 := GenS2{F0: 9223372036854775807, F1: -2147483648}
    var v27 int = v25
    fmt.Println(v27)
    v28 := GenS5{F0: 836, F1: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
