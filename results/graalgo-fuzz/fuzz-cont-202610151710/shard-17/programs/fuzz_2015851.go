package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 bool }
type GenS2 struct { F0 uint64; F1 string; F2 uint8 }
func genFn0(p0 bool) uint32 {
    v7 := []int{387, 0, -2147483648, 32768}
    for fi8 := 0; fi8 < 13; fi8++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return 2147483648
}
func genFn1() int32 {
    var v9 uint32 = 224
    v10 := []int{32768}
    var v11 int64 = 1
    var v12 uint8 = 127
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return -757
}

func main() {
    var v1 int8 = 1
    v2 := make(map[int]bool)
    var v3 uint32 = 4294967295
    var v4 int32 = 1
    var v5 float64 = 1.5
    var v6 uint8 = 0
    v3 = (v3 + (^(v3)))
    if (9223372036854775807 != 9223372036854775807) {
        v13 := v2[480]
        _ = v13
    }
    v14 := GenS1{F0: 956, F1: ("foo bar" + ""), F2: ("a" <= "hello")}
    var v15 uint32 = (v3 ^ (v3 & (v3 % (((v3 * genFn0(false))) | 1))))
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 int64 = 0
        var v18 string = ""
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 int = 0
    v20 := []int{(v19 >> ((42) & 63)), v19}
    for fi21 := 0; fi21 < 18; fi21++ {
        if false {
            if (!(!true)) {
                var v22 float64 = 0.1
                fmt.Println(v6)
                v2[-2147483648] = false
                _ = v22
            } else {
                v3 += uint32(v6)
            }
        } else {
            v23 := GenS2{F0: 974, F1: ("hello" + "foo bar"), F2: 1}
            _ = v23
        }
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
