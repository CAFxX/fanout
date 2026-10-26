package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint8 }

func main() {
    v1 := []int{-504, 1, 127, 127}
    var v2 uint8 = 42
    v3 := make(map[int]string)
    var v4 int32 = 0
    switch -9223372036854775808 {
    case 127:
        if (true || (1 != 63)) {
            fmt.Println(v2, v4)
            switch "a" {
            case "":
                v2 -= 84
            }
        } else {
            v5 := v3[128]
            var v6 uint = 1
            _ = v5
            _ = v6
        }
    case -1:
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 bool = true
            v2 += v2
            _ = fi7
            _ = v8
        }
    }
    if false {
        fmt.Println(v2, v2)
        fmt.Println(v2)
        v9 := GenS1{F0: 2147483648, F1: "\n\t"}
        _ = v9
    } else {
        fmt.Println(v2)
    }
    if len(v1) > 0 {
        v1[0] = 42
    }
    v10 := GenS2{F0: 65535, F1: 128}
    if (42 == -1) {
        v4 -= 328
        v2 = (v2 & (v2 + 0))
    } else {
        v1 = append(v1, 42)
        fmt.Println(v2, v2)
    }
    fmt.Println(v2, v2)
    for fi11 := 0; fi11 < 2; fi11++ {
        delete(v3, fi11)
        var v12 int64 = -1
        _ = fi11
        _ = v12
    }
    var v13 int64 = 1
    v2 = v2
    var v14 bool = true
    v2 -= v2
    v15 := GenS1{F0: 317, F1: (((("foo bar" + ((("" + ("" + "\n\t")) + "a") + "\n\t")) + "foo bar") + "hello") + "foo bar")}
    v16 := GenS3{F0: 818, F1: v2}
    v17 := GenS2{F0: 255, F1: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
