package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int; F2 uint64 }
func genFn0(p0 uint8) uint {
    var v7 float32 = 1.5
    var v8 float64 = 1.5
    _ = p0
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 string) bool {
    p0 = p0
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 float64 = (0.1 * 3.14159)
        _ = fi9
        _ = v10
    }
    p0 = (((p0 + p0) + (p0 + p0)) + (("a" + p0) + ""))
    _ = p0
    return false
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]bool)
    var v3 string = "hello"
    var v4 uint16 = 489
    var v5 string = v3
    var v6 int8 = 127
    for fi11 := 0; fi11 < 1; fi11++ {
        v6 += (v6 << ((454) & 63))
        if false {
            v3 = v5
        } else {
            _ = len(v2)
            v5 = v3
        }
        _ = fi11
    }
    var v12 int32 = -2147483648
    if true {
        v12 += (v12 & -2147483648)
        if ((v4 % (((v4 | (v4 | (v4 & v4)))) | 1)) == v4) {
            for fi13 := 0; fi13 < 15; fi13++ {
                var v14 uint = 1
                _ = fi13
                _ = v14
            }
            switch v5 {
            case "foo bar":
                var v15 int64 = 42
                _ = v15
            case "\n\t":
                var v16 int64 = 0
                _ = v16
            default:
                var v17 int = 478
                _ = v17
            }
        }
    }
    fmt.Println(v3)
    var v18 int64 = -9223372036854775808
    if false {
        var v19 int16 = -1
        fmt.Println(v18, v4)
        _ = v19
    }
    if false {
        fmt.Println(v1, v18)
        switch v3 {
        case "foo bar":
            v20 := GenS1{F0: genFn0(0), F1: -32768}
            _ = v20
        case "a":
            var v21 int = 2147483647
            _ = v21
        default:
            v22 := v2[585]
            _ = v22
        }
    }
    switch v5 {
    case "foo bar":
        for fi23 := 0; fi23 < 8; fi23++ {
            v3 = "a"
            fmt.Println(v3, v12)
            var v24 int32 = (^(v12))
            _ = fi23
            _ = v24
        }
    case "hello":
        v3 = v3
    default:
        var v25 int8 = v6
        _ = v25
    }
    switch ("a" + ("\n\t" + "hello")) {
    case "foo bar":
        fmt.Println(v18)
    case "hello":
        v6 -= v6
    default:
        if ((v6 >> ((2) & 63)) == 100) {
            var v26 int16 = 42
            _ = v26
        }
    }
    fmt.Println(v6, v12, v1)
    var v27 uint = 227
    _ = len(v2)
    var v28 bool = (!(-170 <= -1))
    if false {
        v29 := GenS2{F0: ("\n\t" + ("\n\t" + ""))}
        for fi30 := 0; fi30 < 5; fi30++ {
            fmt.Println(v18, v1)
            if (((v5 + ("hello" + v5)) > "") && v28) {
                var v31 uint16 = v4
                v1 += ((((v1 * v1) / 0.1) * 1.5) + -2.25)
                _ = v31
            }
            var v32 float64 = 3.14159
            _ = fi30
            _ = v32
        }
        fmt.Println(v27, v6)
        _ = v29
    }
    if v28 {
        v33 := v2[4294967295]
        if v33 {
            switch "foo bar" {
            case "":
                for fi34 := 0; fi34 < 12; fi34++ {
                    v27 = v27
                    v35 := GenS3{F0: 42, F1: fi34, F2: 1}
                    v3 = v5
                    _ = fi34
                    _ = v35
                }
            case "a":
                v36 := GenS1{F0: v27, F1: -32768}
                _ = v36
            }
            v27 -= (v27 - v27)
            var v37 uint32 = 811
            _ = v37
        } else {
            switch -1 {
            case 16:
                v28 = (v28 || false)
            case 2147483647:
                v38 := []int{-1, 4294967295, -2147483648, -820}
                _ = v38
            }
        }
        for fi39 := 0; fi39 < 3; fi39++ {
            v40 := []int{(fi39 % ((-2147483648) | 1))}
            _ = fi39
            _ = v40
        }
        _ = v33
    } else {
        fmt.Println(v18)
    }
    var v41 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
