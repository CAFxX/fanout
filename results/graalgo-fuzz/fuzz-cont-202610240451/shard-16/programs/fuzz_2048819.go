package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int16; F1 int8; F2 uint }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint16) int8 {
    fmt.Println(p0)
    v4 := GenS1{F0: 2, F1: 0}
    v5 := GenS2{F0: -32768, F1: -51, F2: 156}
    _ = p0
    _ = v4
    _ = v5
    return -1
}
func genFn1() int32 {
    var v6 int = -1
    switch "foo bar" {
    case "hello":
        if false {
            var v7 bool = (true && (v6 == 861))
            _ = v7
        } else {
            v6 -= 108
        }
    default:
        v6 = v6
    }
    v8 := GenS3{F0: "a", F1: 18446744073709551615}
    fmt.Println(v6)
    _ = v6
    _ = v8
    return 171
}

func main() {
    v1 := []int{-2147483648}
    var v2 uint8 = 127
    var v3 uint8 = (v2 >> ((879) & 63))
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 int = (fi9 >> ((81) & 63))
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 uint64 = 0
            v2 = 2
            v10 -= (v10 % ((127) | 1))
            _ = fi11
            _ = v12
        }
        if len(v1) > 0 {
            v1[0] = fi9
        }
        _ = fi9
        _ = v10
    }
    var v13 int64 = -9223372036854775808
    if true {
        var v14 uint = 986
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 bool = (1 <= -20)
        var v17 uint64 = 0
        _ = fi15
        _ = v16
        _ = v17
    }
    v1 = append(v1, -1)
    v13 -= v13
    var v18 uint = 553
    for fi19 := 0; fi19 < 14; fi19++ {
        if (4294967295 != 1) {
            var v20 int16 = 0
            v2 = v2
            _ = v20
        } else {
            fmt.Println(v18)
        }
        _ = fi19
    }
    var v21 uint16 = 32767
    switch 0 {
    case 918:
        var v22 int32 = -208
        _ = v22
    case 891:
        v1 = append(v1, 961)
    case 2147483647:
        v21 = (^(v21))
    default:
        v21 += (v21 + v21)
    }
    fmt.Println(v3)
    if true {
        var v23 uint64 = 1
        v18 = v18
        _ = v23
    }
    fmt.Println(v3, v13, v18)
    switch 979 {
    case 42:
        v24 := GenS4{F0: v21}
        _ = v24
    case -9223372036854775808:
        v25 := GenS4{F0: (v21 % ((uint16(v3)) | 1))}
        _ = v25
    case 4294967295:
        if (-875 != -708) {
            switch (("\n\t" + "foo bar") + "hello") {
            case "foo bar":
                var v26 int64 = (v13 % ((1) | 1))
                _ = v26
            }
            if len(v1) > 0 {
                v1[0] = 128
            }
            if len(v1) > 0 {
                v1[0] = 459
            }
        }
    default:
        var v27 int8 = 5
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v21
    _ = v3
}
