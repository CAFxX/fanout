package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint }

func main() {
    var v1 bool = (1 <= 18446744073709551615)
    var v2 uint = 2
    v3 := make(map[string]string)
    v4 := make(map[int]uint32)
    v5 := make(map[int]uint32)
    v6 := []int{0, 4294967295, 925}
    if v1 {
        var v7 int16 = 32767
        _ = v7
    } else {
        if v1 {
            v2 += 0
            for fi8 := 0; fi8 < 17; fi8++ {
                v1 = (714 <= 42)
                _ = len(v6)
                _ = fi8
            }
            fmt.Println(v1, v1)
        } else {
            if true {
                fmt.Println(v1, v2)
                v9 := GenS1{F0: 0, F1: 973, F2: (v2 + (v2 | (v2 << ((v2) & 63))))}
                _ = v9
            }
        }
        v2 -= v2
        var v10 int = 157
        _ = v10
    }
    var v11 int = 20
    v1 = true
    v11 = (v11 << ((v2) & 63))
    v11 = -9223372036854775808
    var v12 int32 = -2147483648
    switch ("" + "foo bar") {
    case "a":
        delete(v5, 32767)
    case "":
        var v13 uint16 = 0
        _ = v13
    default:
        v12 += (v12 * v12)
    }
    var v14 string = "foo bar"
    if (132 < 32767) {
        fmt.Println(v1, v12)
        var v15 uint64 = 42
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 string = (v14 + ((("hello" + v14) + v14) + v14))
            for fi18 := 0; fi18 < 15; fi18++ {
                var v19 int = fi18
                v20 := GenS1{F0: (^(v15)), F1: 255, F2: 255}
                _ = fi18
                _ = v19
                _ = v20
            }
            var v21 int8 = 1
            _ = fi16
            _ = v17
            _ = v21
        }
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
