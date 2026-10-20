package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 bool; F1 string; F2 uint64 }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS6 struct { F0 uint64; F1 int8; F2 int16 }
func genFn0(p0 int32) int16 {
    var v7 int16 = 1000
    v8 := GenS1{F0: 18446744073709551615}
    if (true && (127 <= 128)) {
        if true {
            v9 := []int{1, 42, 32768}
            if true {
                v10 := GenS2{F0: (p0 << ((42) & 63))}
                var v11 int64 = 42
                fmt.Println(v7, p0)
                _ = v10
                _ = v11
            }
            _ = v9
        } else {
            var v12 uint64 = 42
            fmt.Println(v12)
            v7 = 42
            _ = v12
        }
    } else {
        v13 := []int{-610, -2147483648, 288}
        v14 := GenS2{F0: p0}
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v7
    _ = v8
    return (v7 + (v7 * 1))
}
func genFn1(p0 int32, p1 string) bool {
    v15 := GenS3{F0: -68, F1: 2}
    if ((335 <= 66) && (!false)) {
        var v16 string = "a"
        p1 = v16
        v17 := GenS1{F0: 42}
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v15
    return false
}

func main() {
    var v1 uint8 = 120
    var v2 int64 = 9223372036854775807
    var v3 uint32 = 1
    var v4 uint32 = 202
    var v5 uint = 828
    v6 := make(map[uint32]bool)
    v4 -= (v3 ^ (v3 ^ 909))
    fmt.Println(v3)
    v3 -= (^(v3))
    fmt.Println(v3)
    fmt.Println(v5, v4)
    if false {
        fmt.Println(v3)
        fmt.Println(v2)
    }
    var v18 int8 = 0
    switch 586 {
    case 541:
        delete(v6, (v4 % (((v3 ^ 471)) | 1)))
    default:
        v19 := []int{32767, 9223372036854775807, 2147483647, -354}
        _ = v19
    }
    switch 9223372036854775807 {
    case -630:
        if true {
            v20 := GenS1{F0: (v5 * (^(v5)))}
            var v21 uint = 0
            var v22 uint8 = (v1 ^ 127)
            _ = v20
            _ = v21
            _ = v22
        } else {
            var v23 uint = 2
            _ = v23
        }
    }
    v5 = (v5 & (v5 >> ((525) & 63)))
    switch 153 {
    case 455:
        var v24 float32 = 1.5
        _ = v24
    case 4294967295:
        v4 += 2
    case 127:
        var v25 uint16 = uint16(v4)
        _ = v25
    default:
        v4 += (v4 / ((871) | 1))
    }
    switch "hello" {
    case "a":
        v26 := GenS4{F0: false, F1: "hello", F2: 864}
        _ = v26
    case "hello":
        switch -613 {
        case -287:
            var v27 int = 350
            _ = v27
        }
    case "\n\t":
        switch "\n\t" {
        case "":
            for fi28 := 0; fi28 < 12; fi28++ {
                v29 := []int{}
                fmt.Println(v5, fi28, v3)
                v30 := GenS5{F0: v3, F1: 255, F2: (v1 + (v1 ^ uint8(v3)))}
                _ = fi28
                _ = v29
                _ = v30
            }
        }
    default:
        switch 111 {
        case 509:
            v31 := GenS6{F0: 42, F1: (v18 * v18), F2: genFn0(42)}
            _ = v31
        case -1:
            fmt.Println(v5)
        }
    }
    v3 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
