package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int16 }
type GenS2 struct { F0 uint }
func genFn0(p0 string) string {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint32 = 2147483647
        _ = fi6
        _ = v7
    }
    switch "a" {
    case "":
        v8 := []int{32768, 128, -957}
        _ = v8
    case "hello":
        if false {
            v9 := GenS1{F0: 100, F1: (2147483647 < 1), F2: 437}
            _ = v9
        } else {
            switch 4294967295 {
            case -910:
                p0 = (p0 + ("foo bar" + (p0 + (((p0 + "\n\t") + "foo bar") + p0))))
            default:
                p0 = ""
            }
        }
    case "foo bar":
        for fi10 := 0; fi10 < 7; fi10++ {
            switch "\n\t" {
            case "":
                var v11 int8 = 42
                _ = v11
            case "hello":
                v12 := []int{}
                _ = v12
            default:
                v13 := GenS1{F0: 42, F1: (fi10 >= (^(fi10))), F2: 32767}
                _ = v13
            }
            _ = fi10
        }
    default:
        if true {
            v14 := []int{4294967295, 32767}
            _ = v14
        } else {
            var v15 int64 = -1
            for fi16 := 0; fi16 < 10; fi16++ {
                fmt.Println(fi16, p0, fi16)
                v15 += (-(v15))
                _ = fi16
            }
            fmt.Println(v15, v15)
            _ = v15
        }
    }
    var v17 int32 = -317
    v17 -= v17
    _ = p0
    _ = v17
    return "hello"
}

func main() {
    var v1 int64 = -196
    var v2 int = 0
    var v3 uint16 = 185
    var v4 int32 = 1
    v5 := make(map[int]uint32)
    v2 -= (v2 - (v2 + 387))
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 float64 = 1e10
        _ = fi18
        _ = v19
    }
    var v20 bool = false
    v21 := GenS2{F0: 123}
    v3 = 155
    var v22 int32 = 391
    if (false && (v20 && (0 == 42))) {
        switch -9223372036854775808 {
        case 32768:
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 int64 = (v1 << ((0) & 63))
                fmt.Println(v24, v2)
                _ = fi23
                _ = v24
            }
        }
        if v20 {
            var v25 uint8 = 127
            _ = v25
        }
        if ((v3 > 2) && (-624 <= -521)) {
            if (1 >= 135) {
                fmt.Println(v4)
            }
        } else {
            var v26 uint8 = 92
            v27 := v5[-1]
            _ = v26
            _ = v27
        }
    } else {
        if (!v20) {
            var v28 uint = 0
            _ = v28
        }
        fmt.Println(v1, v4)
        var v29 uint32 = 955
        _ = v29
    }
    v3 -= (v3 ^ (^(v3)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
