package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }

func main() {
    var v1 int = -2147483648
    v2 := make(map[int]int)
    var v3 int32 = -75
    var v4 int = 2147483647
    var v5 int = (v1 * v4)
    v1 += v4
    if (false || (-933 > -363)) {
        if (v3 > 0) {
            var v6 float64 = 1e10
            if (127 == 0) {
                if (v6 >= float64(((-2.25 + ((1e10 * 0.1) - 1e10)) * (float32(0.0) * 1.5)))) {
                    if (false || false) {
                        v1 += (v4 ^ (v4 * v5))
                    }
                    if (32767 > 255) {
                        delete(v2, v4)
                        v2[v4] = (-(v5))
                        var v7 uint8 = 1
                        _ = v7
                    }
                    fmt.Println(v6)
                }
                for fi8 := 0; fi8 < 17; fi8++ {
                    var v9 uint32 = 4294967295
                    v9 = (v9 + (v9 << ((2) & 63)))
                    _ = fi8
                    _ = v9
                }
            } else {
                for fi10 := 0; fi10 < 4; fi10++ {
                    v11 := GenS1{F0: "a", F1: -1}
                    var v12 int16 = 1
                    v3 += (v3 >> ((1) & 63))
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                if (true && (-545 >= -1)) {
                    v3 += v3
                    v4 = (v5 << ((184) & 63))
                }
            }
            _ = v6
        }
        var v13 float32 = (0.1 / -2.25)
        _ = v13
    } else {
        var v14 uint8 = 0
        v15 := []int{(v4 * (v4 * -802)), 1, -1, -9223372036854775808}
        _ = v14
        _ = v15
    }
    if false {
        var v16 int8 = -65
        if (291 <= 735) {
            for fi17 := 0; fi17 < 13; fi17++ {
                for fi18 := 0; fi18 < 8; fi18++ {
                    if (true && true) {
                        v19 := GenS1{F0: ("foo bar" + "a"), F1: 32767}
                        _ = v19
                    }
                    fmt.Println(v3, fi18, v1)
                    _ = fi18
                }
                for fi20 := 0; fi20 < 13; fi20++ {
                    v21 := []int{1, fi17}
                    if (!false) {
                        v4 = 835
                    }
                    _ = fi20
                    _ = v21
                }
                _ = fi17
            }
        }
        fmt.Println(v1)
        _ = v16
    } else {
        v5 = (v4 << ((255) & 63))
    }
    var v22 int64 = 670
    var v23 bool = ("" == "hello")
    for fi24 := 0; fi24 < 18; fi24++ {
        fmt.Println(v3, v4, v4)
        _ = fi24
    }
    v3 = v3
    v1 = (v1 | (v4 % ((v1) | 1)))
    for fi25 := 0; fi25 < 6; fi25++ {
        fmt.Println(v23, v4, v23)
        var v26 string = ""
        delete(v2, (fi25 - fi25))
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
