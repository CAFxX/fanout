package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
func genFn0() int {
    if false {
        var v7 uint = 2
        v7 += (v7 + v7)
        var v8 uint = (v7 << (((v7 ^ v7)) & 63))
        _ = v7
        _ = v8
    }
    var v9 uint32 = 644
    var v10 int8 = 6
    _ = v10
    _ = v9
    return -636
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = -862
    v3 := []int{127, -9223372036854775808, 490, 4294967295}
    var v4 int8 = 0
    var v5 int64 = 9223372036854775807
    var v6 uint32 = 4294967295
    v1[v6] = 32767
    v11 := GenS1{F0: 65535}
    switch "hello" {
    case "foo bar":
        v12 := GenS1{F0: 42}
        _ = v12
    }
    v4 = (v4 << ((5) & 63))
    if (false && (42 <= 42)) {
        fmt.Println(v2, v2, v2)
        v2 = 1000
        if (!false) {
            var v13 uint32 = 931
            _ = v13
        }
    }
    var v14 float32 = ((float32(((float64(1e10) - 1e10) / ((1.5 / 1.5) + (1.5 - (3.14159 * 0.0))))) - (1.5 / 3.14159)) - 3.14159)
    var v15 float32 = 0.0
    if false {
        v16 := GenS1{F0: 532}
        var v17 float64 = ((0.1 + 3.14159) - 3.14159)
        v5 = v5
        _ = v16
        _ = v17
    } else {
        switch "a" {
        case "\n\t":
            fmt.Println(v4)
        }
        switch ("hello" + (("" + "") + "")) {
        case "hello":
            if len(v3) > 0 {
                v3[0] = 32767
            }
        case "":
            var v18 int64 = v5
            _ = v18
        }
    }
    if (2147483647 != 643) {
        var v19 uint8 = 22
        if false {
            v20 := GenS2{F0: 32767}
            switch genFn0() {
            case 0:
                if true {
                    _ = len(v1)
                }
            case -903:
                v19 = (v19 % ((v19) | 1))
            case 640:
                v4 += v4
            default:
                var v21 uint = 1
                _ = v21
            }
            _ = v20
        }
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
