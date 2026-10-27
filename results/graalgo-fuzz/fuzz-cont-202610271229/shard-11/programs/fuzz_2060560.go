package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int16 }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]string)
    v3 := make(map[string]uint32)
    v4 := make(map[int]string)
    v5 := []int{}
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: ("hello" + v1), F1: 2147483647}
        _ = fi6
        _ = v7
    }
    var v8 int16 = 471
    v8 = (v8 | v8)
    var v9 float64 = 1e10
    if false {
        switch -194 {
        case 0:
            var v10 int16 = 159
            _ = v10
        case 128:
            v11 := GenS2{F0: 890, F1: ("" + "\n\t")}
            _ = v11
        case -2147483648:
            switch 2147483647 {
            case 2147483647:
                v1 = ""
            case 128:
                if true {
                    v12 := v2["\n\t"]
                    var v13 uint16 = 42
                    _ = v12
                    _ = v13
                }
            case -857:
                var v14 int32 = -118
                _ = v14
            default:
                var v15 int32 = 42
                _ = v15
            }
        }
        var v16 float64 = ((v9 / v9) / ((v9 + (0.0 * (v9 + 0.0))) + 0.1))
        var v17 float64 = 0.1
        _ = v16
        _ = v17
    } else {
        _ = len(v5)
        _ = len(v3)
    }
    var v18 uint64 = 0
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        v20 := v2["a"]
        fmt.Println(v9)
        var v21 string = (v20 + v1)
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 int = 32768
    if (42 <= 0) {
        if true {
            var v23 int8 = 23
            var v24 uint64 = 1
            _ = v23
            _ = v24
        } else {
            _ = len(v5)
            v18 -= 845
        }
        delete(v4, v22)
        v18 = v18
    } else {
        if (false || true) {
            if ((100 > 1) || ((v8 == 1000) || (255 >= 128))) {
                v8 += (v8 % (((v8 << ((42) & 63))) | 1))
                v18 -= (v18 + v18)
                v25 := GenS2{F0: -1, F1: (v1 + "foo bar")}
                _ = v25
            }
            v8 += v8
        } else {
            var v26 uint16 = 514
            v27 := GenS3{F0: (v8 / ((32767) | 1))}
            _ = v26
            _ = v27
        }
        switch "foo bar" {
        case "a":
            switch v1 {
            case "":
                v1 = ((v1 + ("foo bar" + ("\n\t" + "\n\t"))) + ("foo bar" + ("a" + v1)))
            case "hello":
                v8 = -32768
            case "\n\t":
                var v28 int = (v22 | v22)
                _ = v28
            default:
                v9 = v9
            }
        case "\n\t":
            v9 = 1.5
        case "":
            v8 = v8
        }
        v18 -= v18
    }
    v8 += v8
    for fi29 := 0; fi29 < 19; fi29++ {
        v18 = (v18 << ((407) & 63))
        v9 -= (v9 * 0.1)
        fmt.Println(v8, v8)
        _ = fi29
    }
    v18 += v18
    var v30 int32 = int32(v18)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
