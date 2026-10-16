package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint16 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint32 = 822
    v2 := make(map[string]uint32)
    v3 := []int{9223372036854775807, 0}
    var v4 uint = 0
    v5 := GenS1{F0: 42, F1: "", F2: 689}
    v1 = v1
    switch ("hello" + "hello") {
    case "foo bar":
        if ((14 <= -24) || true) {
            fmt.Println(v1, v1)
            v4 += v4
        } else {
            var v6 uint = v4
            v1 -= v1
            _ = v6
        }
    case "a":
        v7 := GenS1{F0: 1, F1: "", F2: 2}
        _ = v7
    case "":
        _ = len(v2)
    default:
        v8 := v2["a"]
        _ = v8
    }
    v9 := GenS1{F0: 255, F1: (("hello" + "foo bar") + ("foo bar" + "a")), F2: 32767}
    switch 209 {
    case 9223372036854775807:
        switch 42 {
        case -633:
            if true {
                if (255 != 197) {
                    v4 = v4
                } else {
                    var v10 int8 = 42
                    _ = v10
                }
                var v11 uint8 = 152
                _ = v11
            } else {
                var v12 int16 = -1
                v1 -= (^(v1))
                var v13 uint32 = v1
                _ = v12
                _ = v13
            }
        case 1:
            delete(v2, "")
        case -9223372036854775808:
            v1 = (v1 % ((2) | 1))
        default:
            switch 1 {
            case 0:
                v14 := GenS2{F0: 980}
                _ = v14
            default:
                if (18446744073709551615 <= 1) {
                    v4 += (v4 * 18446744073709551615)
                }
            }
        }
    case 127:
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 uint = 4294967295
            v17 := v2[""]
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    v1 += (v1 >> ((v4) & 63))
    v1 = v1
    for fi18 := 0; fi18 < 9; fi18++ {
        v4 = (v4 & 42)
        _ = fi18
    }
    v4 -= 1
    if true {
        fmt.Println(v1)
        v2["a"] = (^(v1))
        var v19 float64 = 1.5
        _ = v19
    } else {
        var v20 uint64 = 2
        _ = v20
    }
    if false {
        var v21 uint = (v4 >> ((65535) & 63))
        v2["foo bar"] = (v1 & v1)
        var v22 bool = ((v1 & (^(v1))) != v1)
        _ = v21
        _ = v22
    } else {
        if (true && true) {
            var v23 uint32 = 2147483647
            _ = v23
        }
        for fi24 := 0; fi24 < 17; fi24++ {
            if len(v3) > 0 {
                v3[0] = 82
            }
            _ = fi24
        }
        if (4294967295 != 18446744073709551615) {
            v3 = append(v3, -1)
        }
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        for fi26 := 0; fi26 < 4; fi26++ {
            v27 := GenS2{F0: 891}
            fmt.Println(fi25, v4)
            _ = fi26
            _ = v27
        }
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 float32 = 1.5
            _ = fi28
            _ = v29
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
