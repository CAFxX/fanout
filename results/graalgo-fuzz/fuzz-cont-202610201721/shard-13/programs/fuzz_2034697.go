package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 bool; F1 int; F2 uint }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint64, p1 uint8) int64 {
    if (true || (float32(((-2.25 - 1.5) / 1.0)) > 0.0)) {
        if false {
            p1 += 42
            var v6 string = (("" + ("hello" + "a")) + "\n\t")
            for fi7 := 0; fi7 < 19; fi7++ {
                v8 := GenS1{F0: 2, F1: (fi7 + fi7)}
                _ = fi7
                _ = v8
            }
            _ = v6
        }
    } else {
        fmt.Println(p0, p0)
        var v9 uint = 214
        _ = v9
    }
    switch 128 {
    case 42:
        p1 += 128
    case 4294967295:
        v10 := GenS1{F0: 0, F1: 128}
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        p1 -= (p1 % (((p1 << ((42) & 63))) | 1))
        p0 += (p0 / ((18446744073709551615) | 1))
        _ = fi11
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 566
}

func main() {
    var v1 string = ("" + ("foo bar" + "foo bar"))
    v2 := make(map[uint32]bool)
    v3 := []string{(v1 + ("" + v1))}
    var v4 uint = 150
    var v5 int8 = 100
    var v12 float64 = 0.1
    var v13 uint64 = 3
    fmt.Println(v13, v13, v13)
    _ = len(v3)
    var v14 float64 = 1e10
    if false {
        v12 += float64(0.0)
    } else {
        v4 = v4
        v15 := GenS2{F0: true, F1: 32768, F2: 18446744073709551615}
        var v16 int64 = -9223372036854775808
        _ = v15
        _ = v16
    }
    v5 += v5
    v17 := GenS2{F0: false, F1: 9223372036854775807, F2: v4}
    v18 := GenS1{F0: (v4 - 18446744073709551615), F1: -183}
    v2[65535] = true
    switch v1 {
    case "foo bar":
        for fi19 := 0; fi19 < 18; fi19++ {
            if false {
                v2[0] = true
                v13 += (v13 % (((v13 % (((v13 | (v13 - (^(v13))))) | 1))) | 1))
                v12 += (v14 + float64(1e10))
            }
            fmt.Println(v13, v13, fi19)
            v20 := GenS3{F0: 0}
            _ = fi19
            _ = v20
        }
    case "hello":
        v21 := v2[42]
        _ = v21
    default:
        v4 = (v4 / ((v4) | 1))
    }
    fmt.Println(v13)
    v13 = v13
    v22 := GenS1{F0: v4, F1: 128}
    if (42 <= 178) {
        v5 = 1
    }
    if (!(-48 != -2147483648)) {
        v23 := GenS2{F0: true, F1: 1, F2: (v4 ^ v4)}
        switch 293 {
        case -9223372036854775808:
            var v24 int = 9223372036854775807
            _ = v24
        default:
            v13 = v13
        }
        _ = v23
    }
    v1 = "hello"
    if false {
        v25 := GenS3{F0: -204}
        fmt.Println(v4, v14)
        _ = v25
    } else {
        for fi26 := 0; fi26 < 7; fi26++ {
            v27 := GenS3{F0: -1}
            fmt.Println(v1, v13, v4)
            _ = fi26
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
