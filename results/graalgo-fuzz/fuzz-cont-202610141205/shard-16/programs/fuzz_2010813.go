package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]bool)
    var v3 uint8 = 128
    var v4 uint8 = v3
    v5 := []int{-2147483648, 32768}
    v3 = uint8(v1)
    v3 += 154
    v6 := v2["\n\t"]
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: v6, F1: ("" + "a")}
        _ = len(v5)
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 329
    for fi10 := 0; fi10 < 15; fi10++ {
        fmt.Println(v6, v3)
        if (!(!true)) {
            v9 -= uint64(v1)
            v11 := GenS2{F0: (fi10 - fi10), F1: 2}
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 uint16 = 21
                fmt.Println(v1, v6, fi10)
                fmt.Println(v1)
                _ = fi12
                _ = v13
            }
            _ = v11
        } else {
            v4 -= (v4 * (v3 * v3))
            for fi14 := 0; fi14 < 8; fi14++ {
                if len(v5) > 0 {
                    v5[0] = (fi14 * int(v4))
                }
                _ = fi14
            }
            for fi15 := 0; fi15 < 18; fi15++ {
                if false {
                    v16 := GenS3{F0: 770}
                    if v6 {
                        v3 = (v4 >> ((0) & 63))
                    }
                    v17 := v2["hello"]
                    var v18 bool = false
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = fi15
            }
        }
        v2["\n\t"] = true
        _ = fi10
    }
    var v19 int32 = 2147483647
    v6 = ((1 == 889) && false)
    v6 = (!(2 >= 992))
    var v20 int8 = 100
    v21 := GenS2{F0: 359, F1: 65535}
    var v22 int32 = v19
    if (v3 <= (v4 % ((v4) | 1))) {
        v6 = (!(0 <= 824))
        v20 = (-(v20))
        v23 := GenS2{F0: 42, F1: 2}
        _ = v23
    } else {
        var v24 float32 = 3.14159
        var v25 int = -2147483648
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
