package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint8; F1 int }
func genFn0(p0 int8) int8 {
    fmt.Println(p0)
    v5 := GenS1{F0: 42, F1: 4294967295, F2: false}
    var v6 int = -1
    _ = p0
    _ = v5
    _ = v6
    return (-(p0))
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 65535
    var v3 uint = 2
    v4 := make(map[int]uint32)
    var v7 uint8 = 1
    var v8 float64 = 0.1
    v9 := GenS2{F0: "a"}
    switch 9223372036854775807 {
    case -1:
        fmt.Println(v8, v2, v8)
    case -100:
        v3 -= (v3 << ((42) & 63))
    case -2147483648:
        if (1.5 <= 3.14159) {
            var v10 int64 = 0
            fmt.Println(v10, v3)
            v11 := GenS3{F0: (v7 | v7), F1: ""}
            _ = v10
            _ = v11
        }
    default:
        var v12 uint32 = (v2 / ((2147483647) | 1))
        _ = v12
    }
    if true {
        var v13 uint32 = uint32(v3)
        var v14 uint64 = 1
        var v15 bool = (!true)
        _ = v13
        _ = v14
        _ = v15
    }
    delete(v4, 630)
    var v16 int8 = 100
    v7 = (v7 | v7)
    if true {
        var v17 string = ("hello" + "hello")
        var v18 float32 = 1.5
        var v19 uint32 = 4294967295
        _ = v17
        _ = v18
        _ = v19
    }
    fmt.Println(v2, v16)
    v20 := GenS4{F0: (^(v7)), F1: -2147483648}
    v8 += (((1.5 * ((v8 - (v8 * (v8 * v8))) * (v8 / (3.14159 + v8)))) + (0.1 * (v8 / -2.25))) - 0.1)
    v21 := []int{966, 646, int(v3), 9223372036854775807}
    switch ((("a" + "hello") + (("foo bar" + "a") + "a")) + "") {
    case "\n\t":
        if false {
            var v22 int8 = -1
            switch -977 {
            case 9223372036854775807:
                var v23 uint8 = 127
                _ = v23
            case 42:
                for fi24 := 0; fi24 < 4; fi24++ {
                    var v25 int = 127
                    _ = fi24
                    _ = v25
                }
            }
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
            _ = v22
        }
    case "hello":
        v2 -= 42
    case "a":
        if (!false) {
            for fi26 := 0; fi26 < 14; fi26++ {
                if len(v21) > 0 {
                    v21[0] = 0
                }
                _ = fi26
            }
            fmt.Println(v16, v2, v7)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
