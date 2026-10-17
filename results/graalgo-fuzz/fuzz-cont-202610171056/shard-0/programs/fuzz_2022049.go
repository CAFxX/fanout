package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint16) int64 {
    p0 -= p0
    switch "" {
    case "":
        var v6 uint8 = 42
        _ = v6
    case "a":
        switch ("foo bar" + "") {
        case "foo bar":
            v7 := []int{-277}
            _ = v7
        }
    default:
        p0 = (p0 / (((p0 * 467)) | 1))
    }
    v8 := GenS1{F0: 1, F1: -247}
    p0 -= p0
    v9 := GenS2{F0: (("foo bar" + "foo bar") + "hello"), F1: -218}
    _ = p0
    _ = v8
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 uint = 4294967295
    var v2 int8 = 127
    var v3 uint8 = 255
    var v4 int8 = (v2 - (v2 ^ -79))
    var v5 uint = v1
    v3 = v3
    var v10 bool = false
    v11 := GenS2{F0: ("hello" + "hello"), F1: 128}
    var v12 uint8 = 39
    if ((v2 + 13) > (v4 ^ (v2 >> ((v5) & 63)))) {
        if true {
            v2 = v4
        }
    }
    switch 127 {
    case -2147483648:
        v3 += (v3 ^ (v12 % ((v3) | 1)))
    case 2147483647:
        var v13 int32 = -1
        _ = v13
    case 128:
        v14 := []int{0, 0, 42, -2147483648}
        _ = v14
    }
    var v15 int8 = v2
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 float64 = (1.5 * 1e10)
        _ = fi16
        _ = v17
    }
    var v18 uint8 = 42
    v19 := GenS2{F0: "foo bar", F1: -563}
    if (!v10) {
        v20 := []int{-67}
        v18 += v3
        switch 0 {
        case 4294967295:
            if len(v20) > 0 {
                v20[0] = 0
            }
        case 0:
            v21 := GenS2{F0: "hello", F1: 410}
            _ = v21
        case 212:
            v22 := GenS3{F0: 153}
            _ = v22
        }
        _ = v20
    } else {
        v5 = (v1 >> (((^(v1))) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
