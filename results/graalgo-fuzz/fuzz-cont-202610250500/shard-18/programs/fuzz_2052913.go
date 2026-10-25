package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int64 }
func genFn0(p0 uint, p1 int) uint {
    var v7 uint16 = 255
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint64 = 116
        v10 := GenS1{F0: 0, F1: 9223372036854775807, F2: 1}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 0
}
func genFn1(p0 int, p1 uint64, p2 int) string {
    var v11 bool = false
    fmt.Println(p0, p2, p0)
    var v12 int16 = -150
    if v11 {
        var v13 uint = 65535
        _ = v13
    } else {
        if (50 < 100) {
            var v14 uint = 4294967295
            v14 = genFn0(genFn0((v14 ^ (^(v14))), p2), p2)
            _ = v14
        }
        v15 := []int{(p2 % (((p0 - 42)) | 1)), p2, p2, (p0 - p0)}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return "a"
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int32 = 42
    v3 := []string{"", "hello", "a", "hello"}
    var v4 uint = 1
    v5 := make(map[string]string)
    var v6 uint16 = 2
    switch "" {
    case "hello":
        var v16 uint16 = (v6 >> (((v4 & (v4 >> ((v4) & 63)))) & 63))
        _ = v16
    default:
        var v17 string = ("foo bar" + "foo bar")
        _ = v17
    }
    v18 := v5["foo bar"]
    v5["hello"] = (v18 + (v18 + (v18 + (v18 + "a"))))
    var v19 int64 = 0
    var v20 uint16 = (v6 | v6)
    var v21 int8 = 127
    v18 = (v18 + "foo bar")
    v6 = v20
    v18 = "a"
    var v22 string = v18
    var v23 uint64 = (v1 * 0)
    v24 := GenS1{F0: (v23 >> ((genFn0(v4, 4294967295)) & 63)), F1: v19, F2: v19}
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
