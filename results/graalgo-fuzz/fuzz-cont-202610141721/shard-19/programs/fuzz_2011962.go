package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
func genFn0() int8 {
    var v4 uint = 65535
    var v5 int64 = 661
    v5 += (v5 - v5)
    v6 := []int{}
    v6 = append(v6, -706)
    _ = v4
    _ = v5
    _ = v6
    return -128
}

func main() {
    var v1 int32 = 257
    var v2 int32 = v1
    v3 := make(map[int]int)
    fmt.Println(v2, v1)
    fmt.Println(v1, v2)
    switch -1 {
    case -2147483648:
        switch (("a" + "\n\t") + "\n\t") {
        case "a":
            var v7 uint = 774
            _ = v7
        default:
            v2 -= v2
        }
    case 32768:
        v8 := []int{}
        _ = v8
    case -1:
        fmt.Println(v2, v2)
    }
    v9 := GenS1{F0: 1, F1: 42}
    v1 -= (v1 >> ((18446744073709551615) & 63))
    if ((!false) && (0 > -1)) {
        fmt.Println(v2)
        switch "foo bar" {
        case "\n\t":
            if (42 >= -50) {
                var v10 uint64 = 1
                _ = v10
            }
        case "":
            if false {
                fmt.Println(v1)
            } else {
                fmt.Println(v2, v1)
                v1 -= v2
                v2 = (v1 << ((42) & 63))
            }
        case "foo bar":
            var v11 int32 = 532
            _ = v11
        default:
            var v12 bool = false
            _ = v12
        }
        var v13 float64 = (1.5 * -2.25)
        _ = v13
    } else {
        v14 := v3[127]
        fmt.Println(v2, v2)
        switch ("\n\t" + "hello") {
        case "":
            v15 := []int{2147483647, (v14 ^ (v14 << ((4294967295) & 63))), (v14 % (((v14 << ((777) & 63))) | 1))}
            _ = v15
        }
        _ = v14
    }
    switch "" {
    case "foo bar":
        if (42 != 42) {
            v2 += 42
            var v16 uint = 65535
            if (!((v16 << (((v16 & v16)) & 63)) <= v16)) {
                switch -9223372036854775808 {
                case 159:
                    var v17 int8 = -1
                    _ = v17
                }
                fmt.Println(v2, v2, v16)
                v2 += v1
            }
            _ = v16
        }
    case "hello":
        var v18 bool = ((1000 >= -32768) && ((0.1 == 0.1) && false))
        _ = v18
    }
    if (2 <= 65535) {
        fmt.Println(v1, v1)
    }
    v2 += -519
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v9
}
