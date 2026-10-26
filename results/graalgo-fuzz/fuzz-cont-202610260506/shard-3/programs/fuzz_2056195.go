package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int32 }

func main() {
    var v1 int64 = -1
    var v2 int16 = 0
    v3 := make(map[int]string)
    v4 := []string{"a"}
    v5 := make(map[string]bool)
    var v6 uint32 = 754
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 18446744073709551615, F1: 0, F2: -1}
        v9 := GenS2{F0: "a"}
        _ = fi7
        _ = v8
        _ = v9
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v10 string = ""
    if ((!false) && (0 > -52)) {
        v11 := GenS3{F0: true, F1: 42}
        v12 := GenS2{F0: (v10 + "a")}
        _ = v11
        _ = v12
    } else {
        var v13 uint8 = 156
        _ = v13
    }
    v2 += v2
    v1 -= (v1 % (((v1 >> ((42) & 63))) | 1))
    v6 = 2147483647
    switch "\n\t" {
    case "a":
        if len(v4) > 0 {
            v4[0] = ("hello" + ("a" + v10))
        }
    case "hello":
        switch v10 {
        case "foo bar":
            var v14 uint = 794
            _ = v14
        case "\n\t":
            _ = len(v4)
        case "a":
            var v15 int16 = v2
            _ = v15
        }
    default:
        if (!false) {
            var v16 int8 = 1
            v10 = v10
            _ = len(v3)
            _ = v16
        } else {
            v17 := GenS3{F0: false, F1: 2147483647}
            if len(v4) > 0 {
                v4[0] = v10
            }
            _ = v17
        }
    }
    fmt.Println(v2, v10, v1)
    v6 -= v6
    if (1 != 48) {
        if (18446744073709551615 < 255) {
            v18 := v3[-1]
            v2 = (v2 % ((-295) | 1))
            _ = v18
        } else {
            var v19 int32 = -415
            _ = v19
        }
        v1 = v1
        v2 -= 0
    }
    v20 := GenS2{F0: ("hello" + ((v10 + (v10 + v10)) + v10))}
    var v21 uint64 = 467
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
