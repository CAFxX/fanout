package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint8; F2 bool }
type GenS3 struct { F0 int16; F1 uint8; F2 uint }

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 281
    var v3 float64 = 0.0
    v4 := []int{32768, 4294967295}
    var v5 string = "foo bar"
    if ((false || (!true)) && true) {
        if true {
            fmt.Println(v3, v3)
        } else {
            fmt.Println(v5, v5, v3)
            _ = len(v4)
            v6 := GenS1{F0: true}
            _ = v6
        }
    }
    if (!((2147483647 >= 993) || true)) {
        if (0 >= -1) {
            v1[710] = ((((v5 + v5) + (v5 + v5)) + "\n\t") == v5)
        } else {
            var v7 int64 = 9223372036854775807
            _ = v7
        }
        v5 = ""
    } else {
        var v8 string = "hello"
        fmt.Println(v2, v8)
        _ = v8
    }
    v5 = v5
    v9 := GenS2{F0: 1, F1: 1, F2: (true && (!false))}
    v10 := GenS1{F0: (46 == 9223372036854775807)}
    if ((v2 | v2) < (v2 * 317)) {
        v5 = "hello"
    } else {
        v5 = ("" + "hello")
        switch v5 {
        case "a":
            switch -346 {
            case -2147483648:
                var v11 uint32 = 557
                _ = v11
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        case "foo bar":
            v3 = 3.14159
        }
        v2 = v2
    }
    if (!(0 <= -28)) {
        v12 := v1[0]
        fmt.Println(v5, v12)
        _ = v12
    } else {
        if len(v4) > 0 {
            v4[0] = -9223372036854775808
        }
    }
    var v13 uint = 2
    switch -1 {
    case 0:
        switch v5 {
        case "\n\t":
            var v14 int64 = 1
            _ = v14
        case "":
            for fi15 := 0; fi15 < 16; fi15++ {
                v16 := GenS3{F0: -1, F1: 2, F2: v13}
                _ = fi15
                _ = v16
            }
        case "hello":
            if (184 >= 819) {
                v17 := GenS2{F0: 0, F1: 128, F2: ((("\n\t" >= "a") && (42 == -9223372036854775808)) || (2147483648 > 555))}
                v18 := GenS2{F0: -1, F1: 127, F2: true}
                _ = v17
                _ = v18
            } else {
                fmt.Println(v5, v5)
                var v19 float64 = 1e10
                _ = v19
            }
        default:
            v20 := GenS3{F0: 1000, F1: 2, F2: v13}
            _ = v20
        }
    }
    v13 -= 269
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
