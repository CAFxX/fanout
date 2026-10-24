package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int64; F1 uint32 }

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]int)
    var v3 uint = 611
    v4 := make(map[int]uint32)
    _ = len(v2)
    v1 -= (v1 ^ 18446744073709551615)
    v2[32768] = -698
    var v5 int16 = 0
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{-2147483648, 4294967295, -2147483648}
        fmt.Println(v1)
        var v8 string = ("a" + "hello")
        _ = fi6
        _ = v7
        _ = v8
    }
    if (((3.14159 / 3.14159) > 3.14159) && (255 != 9)) {
        v9 := []int{}
        v5 += -1
        var v10 bool = false
        _ = v10
        _ = v9
    } else {
        var v11 int = -1
        var v12 float64 = 3.14159
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        _ = len(v2)
        for fi14 := 0; fi14 < 9; fi14++ {
            if (!true) {
                var v15 uint32 = 2147483647
                v1 += (v1 & v1)
                _ = v15
            } else {
                var v16 int64 = 0
                _ = v16
            }
            if true {
                v3 = 711
                var v17 int = 4294967295
                fmt.Println(v17, v17)
                _ = v17
            }
            _ = fi14
        }
        _ = fi13
    }
    if false {
        var v18 int = 32768
        v19 := GenS1{F0: 1000, F1: true}
        _ = v18
        _ = v19
    } else {
        v20 := GenS2{F0: v5, F1: v5}
        _ = v20
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 int16 = (v5 / (((v5 + (-(v5)))) | 1))
        var v23 uint32 = 1
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v3 = v3
        _ = fi24
    }
    v3 = v3
    var v25 uint8 = 42
    if ((-1 == 0) || ((!(!false)) && (624 >= 0))) {
        switch "a" {
        case "foo bar":
            var v26 uint16 = 2
            _ = v26
        default:
            v27 := []int{32767, -702, 0}
            _ = v27
        }
    }
    v28 := GenS3{F0: 1, F1: 42}
    _ = len(v4)
    fmt.Println(v3, v1, v1)
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
