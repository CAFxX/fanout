package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int64; F2 int8 }
func genFn0(p0 int8, p1 string) uint {
    v7 := GenS1{F0: 0, F1: 394, F2: 2147483648}
    var v8 uint16 = 32767
    if (1 >= 663) {
        p0 = p0
        v8 = v8
    } else {
        var v9 int16 = 32767
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int32 = -1
    v2 := make(map[int]string)
    var v3 int32 = 2147483647
    var v4 uint16 = 2
    var v5 int8 = -1
    v6 := make(map[int]int)
    v10 := GenS1{F0: -1, F1: 2147483647, F2: 42}
    var v11 int64 = -975
    var v12 int32 = 0
    var v13 int16 = 32767
    var v14 bool = false
    switch 1 {
    case -729:
        v12 = -1
    case 1:
        switch "\n\t" {
        case "\n\t":
            v15 := v2[128]
            _ = v15
        }
    case -126:
        v1 += (v1 + v12)
    default:
        var v16 float32 = 0.1
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        if v14 {
            for fi18 := 0; fi18 < 19; fi18++ {
                var v19 int64 = v11
                var v20 float32 = (-2.25 + 3.14159)
                _ = fi18
                _ = v19
                _ = v20
            }
            var v21 string = "foo bar"
            v22 := GenS2{F0: v4, F1: fi17}
            _ = v21
            _ = v22
        }
        v23 := GenS2{F0: v4, F1: (-(fi17))}
        var v24 int64 = (^(v11))
        _ = fi17
        _ = v23
        _ = v24
    }
    var v25 uint16 = (v4 / (((^(v4))) | 1))
    v4 -= (v25 % ((2) | 1))
    if ((((!v14) && v14) && (!v14)) || (("" + ("a" + "\n\t")) > "foo bar")) {
        var v26 int16 = 1
        _ = v26
    } else {
        v5 -= int8(v3)
        switch "\n\t" {
        case "\n\t":
            var v27 int = 571
            _ = v27
        default:
            var v28 int64 = v11
            _ = v28
        }
    }
    switch 52 {
    case 32768:
        if v14 {
            var v29 uint = 42
            if true {
                if true {
                    v4 += 32767
                    var v30 int64 = int64(v4)
                    _ = v30
                }
            } else {
                if ("a" == ("" + "")) {
                    v13 += v13
                    fmt.Println(v4, v14)
                    v12 -= (v12 * 42)
                }
                v31 := GenS3{F0: 1, F1: v13, F2: 1}
                _ = v31
            }
            v32 := []int{32767, 128}
            _ = v29
            _ = v32
        } else {
            switch "\n\t" {
            case "hello":
                v12 = (-(v1))
            case "\n\t":
                var v33 uint = 1
                _ = v33
            }
            v34 := []int{-9223372036854775808}
            var v35 uint = 0
            _ = v34
            _ = v35
        }
    case -9223372036854775808:
        switch "hello" {
        case "a":
            fmt.Println(v5, v1, v11)
        case "foo bar":
            switch 1 {
            case 127:
                var v36 uint64 = 0
                _ = v36
            case -404:
                for fi37 := 0; fi37 < 12; fi37++ {
                    v38 := GenS4{F0: v25, F1: v11, F2: (v5 >> ((18446744073709551615) & 63))}
                    _ = fi37
                    _ = v38
                }
            default:
                v11 += 42
            }
        default:
            var v39 float32 = (1.5 / 1.5)
            _ = v39
        }
    default:
        var v40 int = -1
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
