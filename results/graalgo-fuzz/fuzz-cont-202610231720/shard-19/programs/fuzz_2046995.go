package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint }
func genFn0(p0 int16) string {
    fmt.Println(p0)
    switch "\n\t" {
    case "":
        p0 = (p0 << ((2) & 63))
    case "foo bar":
        if ((1 <= 0) && (4294967295 >= 2)) {
            v7 := GenS1{F0: 466, F1: -72}
            if (127 <= 128) {
                var v8 float64 = 3.14159
                _ = v8
            } else {
                v9 := []int{}
                v9 = append(v9, -1)
                _ = v9
            }
            p0 = p0
            _ = v7
        }
    case "hello":
        fmt.Println(p0)
    }
    p0 = p0
    v10 := GenS1{F0: 1, F1: 0}
    _ = p0
    _ = v10
    return ""
}

func main() {
    var v1 int = 32768
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    v5 := []int{}
    var v6 int32 = -934
    fmt.Println(v1, v6, v6)
    v11 := GenS2{F0: 652}
    var v12 string = "hello"
    switch (v12 + ("hello" + "a")) {
    case "foo bar":
        v3["foo bar"] = "\n\t"
    case "":
        v13 := GenS2{F0: 300}
        _ = v13
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    switch "" {
    case "\n\t":
        _ = len(v3)
    case "foo bar":
        var v14 uint64 = 277
        _ = v14
    case "hello":
        var v15 uint32 = 650
        _ = v15
    default:
        var v16 uint32 = 65535
        _ = v16
    }
    switch (v1 << ((774) & 63)) {
    case 42:
        for fi17 := 0; fi17 < 18; fi17++ {
            v6 = (v6 ^ (v6 % ((v6) | 1)))
            for fi18 := 0; fi18 < 1; fi18++ {
                v19 := GenS2{F0: 18446744073709551615}
                switch (v12 + ((v12 + "foo bar") + v12)) {
                case "\n\t":
                    var v20 string = ("hello" + "\n\t")
                    _ = v20
                case "hello":
                    _ = len(v4)
                case "foo bar":
                    v1 = v1
                }
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    default:
        var v21 uint16 = 42
        _ = v21
    }
    v12 = "hello"
    var v22 uint32 = 799
    for fi23 := 0; fi23 < 13; fi23++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v24 bool = true
        _ = fi23
        _ = v24
    }
    var v25 float64 = (float64(float32(1e10)) - (((1.5 + float64(-2.25)) * 3.14159) - 1e10))
    v2 -= 0
    v6 += (v6 << ((2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
