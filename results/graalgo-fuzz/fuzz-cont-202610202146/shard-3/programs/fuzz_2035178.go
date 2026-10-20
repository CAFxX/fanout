package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 string; F2 int16 }
func genFn0(p0 int, p1 int32) int32 {
    v6 := []int{1, (p0 + 867)}
    var v7 bool = (!(true && false))
    p1 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p1 << ((255) & 63))
}
func genFn1(p0 int8, p1 int32, p2 int64) uint64 {
    v8 := GenS1{F0: (p2 | (p2 & (p2 << ((18446744073709551615) & 63)))), F1: 215, F2: 2147483647}
    if true {
        switch ("foo bar" + "\n\t") {
        case "a":
            for fi9 := 0; fi9 < 1; fi9++ {
                p2 += 855
                switch fi9 {
                case 127:
                    p2 -= (p2 << ((255) & 63))
                }
                _ = fi9
            }
        case "\n\t":
            p1 = p1
        }
        switch "foo bar" {
        case "a":
            switch 32768 {
            case 32767:
                for fi10 := 0; fi10 < 18; fi10++ {
                    var v11 float64 = float64(0.1)
                    _ = fi10
                    _ = v11
                }
            }
        }
        p2 = (^(p2))
    }
    v12 := []int{-673, -9223372036854775808, 127, 4294967295}
    var v13 int = -2147483648
    if true {
        if true {
            fmt.Println(v13, p0, p0)
            v14 := GenS2{F0: -360}
            _ = v14
        }
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        fmt.Println(p2, v13, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 int64 = 42
    var v2 int32 = 0
    var v3 uint64 = 0
    v4 := make(map[string]uint32)
    var v5 int64 = 9223372036854775807
    v15 := []int{-9223372036854775808}
    v15 = append(v15, -2147483648)
    if true {
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 uint = 696
            _ = fi16
            _ = v17
        }
        v5 = (v5 << ((2) & 63))
        v2 -= 925
    } else {
        v3 = v3
        _ = len(v15)
    }
    v2 = -1
    if true {
        v18 := GenS3{F0: -2147483648}
        switch "" {
        case "a":
            v5 += v5
        case "foo bar":
            fmt.Println(v1)
        default:
            if len(v15) > 0 {
                v15[0] = -93
            }
        }
        _ = v18
    } else {
        v19 := GenS4{F0: v5, F1: ("" + ""), F2: -1}
        if true {
            switch "foo bar" {
            case "a":
                v20 := GenS4{F0: (v5 % ((v5) | 1)), F1: ("hello" + "foo bar"), F2: 601}
                _ = v20
            }
        }
        _ = v19
    }
    switch 0 {
    case 42:
        var v21 uint = 2
        _ = v21
    case 1:
        fmt.Println(v3)
    case 0:
        for fi22 := 0; fi22 < 19; fi22++ {
            v23 := GenS1{F0: v5, F1: 0, F2: (fi22 / (((fi22 >> ((4294967295) & 63))) | 1))}
            _ = fi22
            _ = v23
        }
    }
    v2 += (v2 % (((v2 >> ((506) & 63))) | 1))
    v4["hello"] = 2147483647
    if (-2.25 != 0.0) {
        v24 := GenS1{F0: v1, F1: -3, F2: -180}
        _ = v24
    }
    fmt.Println(v2)
    v2 -= (-(v2))
    v3 = (v3 / (((v3 << ((2) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
