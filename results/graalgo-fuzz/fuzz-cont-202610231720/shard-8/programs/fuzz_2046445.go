package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := []int{32768, 272, 128}
    v2 := make(map[uint32]bool)
    var v3 uint = 0
    var v4 uint64 = 823
    v5 := []uint32{0, 2}
    switch "hello" {
    case "":
        fmt.Println(v3, v3)
    default:
        var v6 string = "\n\t"
        _ = v6
    }
    if ("a" < (((("\n\t" + "hello") + "a") + "foo bar") + (("foo bar" + "\n\t") + (("hello" + "") + ("foo bar" + "\n\t"))))) {
        if (!((false || (2 >= v3)) && true)) {
            switch 4294967295 {
            case 0:
                var v7 int64 = 0
                _ = v7
            case 127:
                fmt.Println(v4)
            case 540:
                v8 := GenS1{F0: 2147483647}
                _ = v8
            }
            var v9 uint64 = 1
            _ = v9
        }
        if (true && (-2147483648 >= 21)) {
            var v10 uint64 = v4
            v3 -= (v3 ^ (v3 >> ((v3) & 63)))
            _ = v10
        } else {
            if len(v1) > 0 {
                v1[0] = 128
            }
            v4 += (v4 | v4)
        }
    }
    var v11 float64 = 0.1
    switch 127 {
    case 42:
        v11 -= 1e10
    case 935:
        var v12 float32 = float32((float64(0.0) / 1e10))
        _ = v12
    case 9223372036854775807:
        var v13 uint64 = (v4 >> (((v3 & (v3 >> ((v3) & 63)))) & 63))
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 uint32 = 1
        _ = fi14
        _ = v15
    }
    var v16 string = ""
    var v17 float64 = 3.14159
    v1 = append(v1, 1)
    if (-790 >= 32768) {
        var v18 float64 = 0.0
        var v19 uint = (v3 & (^(v3)))
        v2[42] = true
        _ = v18
        _ = v19
    }
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
