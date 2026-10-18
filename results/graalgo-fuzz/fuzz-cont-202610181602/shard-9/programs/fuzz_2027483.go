package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int16; F2 int32 }
type GenS3 struct { F0 uint64; F1 int64; F2 bool }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 string }
func genFn0() uint {
    for fi4 := 0; fi4 < 15; fi4++ {
        var v5 int16 = -190
        fmt.Println(v5)
        v6 := GenS1{F0: fi4}
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := []int{}
    if (((!true) && false) || true) {
        _ = len(v7)
        for fi8 := 0; fi8 < 4; fi8++ {
            fmt.Println(fi8)
            _ = fi8
        }
    }
    var v9 int32 = -1
    v10 := GenS2{F0: "a", F1: 381, F2: v9}
    _ = v10
    _ = v7
    _ = v9
    return 105
}

func main() {
    var v1 bool = (2 > 127)
    var v2 int64 = 9223372036854775807
    var v3 bool = true
    switch 9223372036854775807 {
    case -9223372036854775808:
        var v11 uint32 = 4294967295
        _ = v11
    case 151:
        if (-9223372036854775808 >= -2147483648) {
            if v1 {
                v12 := []int{897, 0}
                for fi13 := 0; fi13 < 12; fi13++ {
                    v1 = (true && (109 != 0))
                    _ = fi13
                }
                _ = v12
            }
            fmt.Println(v2, v2)
        } else {
            var v14 int = int(v2)
            if ((v2 >> ((genFn0()) & 63)) > (v2 << ((255) & 63))) {
                if (-916 <= 2147483647) {
                    v15 := GenS1{F0: -9223372036854775808}
                    _ = v15
                }
            }
            v1 = (-918 < -840)
            _ = v14
        }
    }
    var v16 int16 = -1
    var v17 int8 = 1
    switch "foo bar" {
    case "\n\t":
        v18 := []int{4294967295, -499, 617, -530}
        _ = v18
    default:
        for fi19 := 0; fi19 < 19; fi19++ {
            v20 := GenS3{F0: 64, F1: (v2 ^ 211), F2: ((((1.5 - 1.5) - 0.1) / 0.1) > (-2.25 + 0.0))}
            _ = fi19
            _ = v20
        }
    }
    fmt.Println(v3)
    if (0 != -2147483648) {
        v21 := []int{666, 128, 127, 175}
        fmt.Println(v17, v2, v1)
        v22 := GenS4{F0: (-(v16)), F1: 1}
        _ = v21
        _ = v22
    } else {
        v1 = false
    }
    if (!(9223372036854775807 > -422)) {
        var v23 int64 = v2
        _ = v23
    } else {
        for fi24 := 0; fi24 < 5; fi24++ {
            v25 := []int{fi24}
            _ = fi24
            _ = v25
        }
        v26 := GenS5{F0: ("foo bar" + "foo bar")}
        if true {
            for fi27 := 0; fi27 < 4; fi27++ {
                var v28 uint64 = 355
                _ = fi27
                _ = v28
            }
        }
        _ = v26
    }
    v2 = (v2 / ((-1) | 1))
    switch 2147483647 {
    case -277:
        v29 := []int{9223372036854775807, -2147483648, 128, 42}
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
