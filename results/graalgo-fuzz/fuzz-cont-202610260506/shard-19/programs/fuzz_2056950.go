package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int64; F2 int }
type GenS4 struct { F0 int64; F1 uint16; F2 bool }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 127
    var v3 uint64 = 42
    switch -1 {
    case 0:
        v4 := []int{-2147483648, 549, 205}
        _ = v4
    default:
        if false {
            if true {
                v1[42] = (!(153 >= 692))
                for fi5 := 0; fi5 < 6; fi5++ {
                    v3 -= v3
                    _ = fi5
                }
            }
        } else {
            v6 := GenS1{F0: 87}
            var v7 uint64 = (v3 & v3)
            v8 := GenS1{F0: -19}
            _ = v6
            _ = v7
            _ = v8
        }
    }
    switch "" {
    case "":
        fmt.Println(v3, v3)
    case "\n\t":
        switch ("foo bar" + "foo bar") {
        case "":
            v2 = (^(v2))
        case "\n\t":
            var v9 int64 = 42
            _ = v9
        default:
            v2 += (^(v2))
        }
    case "a":
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 float32 = 3.14159
            for fi12 := 0; fi12 < 3; fi12++ {
                v3 = v3
                v2 -= (v2 & v2)
                v13 := GenS1{F0: -128}
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
    default:
        v14 := GenS1{F0: int8(v2)}
        _ = v14
    }
    var v15 bool = true
    v16 := GenS2{F0: (("a" + "a") + "foo bar")}
    v17 := GenS3{F0: v15, F1: 326, F2: 110}
    var v18 int64 = -9223372036854775808
    v19 := GenS1{F0: 42}
    var v20 uint32 = 2147483647
    v18 += v18
    var v21 uint16 = 539
    v3 -= v3
    if (!v15) {
        switch -9223372036854775808 {
        case 2147483647:
            var v22 uint64 = (v3 * 18446744073709551615)
            _ = v22
        case 4294967295:
            v23 := GenS4{F0: v18, F1: (^(v21)), F2: (v2 <= (v2 ^ (^(v2))))}
            _ = v23
        case -449:
            fmt.Println(v2, v21, v3)
        }
        if v15 {
            for fi24 := 0; fi24 < 4; fi24++ {
                for fi25 := 0; fi25 < 13; fi25++ {
                    v3 += 2
                    _ = fi25
                }
                var v26 int = (^(fi24))
                if false {
                    switch 0 {
                    case -740:
                        var v27 uint8 = v2
                        _ = v27
                    default:
                        var v28 float32 = 1.5
                        _ = v28
                    }
                }
                _ = fi24
                _ = v26
            }
            switch 32767 {
            case -9223372036854775808:
                fmt.Println(v21)
            case -842:
                v29 := []int{-552, 1, 2147483647, 4294967295}
                _ = v29
            }
            v30 := []int{998, 32767, 4294967295}
            _ = v30
        } else {
            v15 = ((false && (v20 > uint32(v21))) || v15)
        }
    } else {
        switch ("" + "a") {
        case "a":
            v3 += 2
        case "":
            fmt.Println(v2, v21, v2)
        default:
            v31 := []int{2147483647}
            _ = v31
        }
    }
    var v32 int = 0
    for fi33 := 0; fi33 < 8; fi33++ {
        v21 = (^(v21))
        _ = fi33
    }
    var v34 uint = 889
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v34
}
