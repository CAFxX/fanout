package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int8) int {
    p0 += (p0 / (((p0 & (p0 * 100))) | 1))
    var v6 uint32 = 4294967295
    if false {
        fmt.Println(v6, p0)
        var v7 bool = false
        _ = v7
    }
    p0 -= p0
    var v8 uint = 4294967295
    _ = p0
    _ = v6
    _ = v8
    return -541
}
func genFn1(p0 int32, p1 int) uint {
    p0 += -25
    v9 := []int{(p1 << ((65535) & 63)), p1, p1, genFn0(127)}
    fmt.Println(p1, p1)
    if true {
        if true {
            var v10 float64 = (0.1 * -2.25)
            switch "" {
            case "foo bar":
                if len(v9) > 0 {
                    v9[0] = genFn0(-42)
                }
            case "hello":
                var v11 uint = 344
                _ = v11
            }
            v12 := GenS1{F0: 65535, F1: 18446744073709551615}
            _ = v10
            _ = v12
        } else {
            v13 := GenS1{F0: 2147483648, F1: 42}
            p1 = p1
            _ = v13
        }
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v9
    return 774
}

func main() {
    v1 := make(map[int]string)
    var v2 int = -1
    v3 := []int{(^(v2)), (v2 & 9223372036854775807), 357}
    var v4 string = "foo bar"
    v5 := []uint32{}
    fmt.Println(v2, v2)
    v2 = genFn0(1)
    v4 = ("a" + ("foo bar" + v4))
    var v14 int8 = 127
    fmt.Println(v4, v14, v2)
    v14 -= (v14 >> ((255) & 63))
    for fi15 := 0; fi15 < 15; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            if (genFn1(-594, (-(v2))) == 207) {
                var v17 int = fi15
                fmt.Println(fi16)
                _ = v17
            }
            _ = fi16
        }
        if (("a" + "\n\t") != "\n\t") {
            if true {
                var v18 float32 = 0.0
                _ = v18
            }
            delete(v1, (^(v2)))
        }
        var v19 uint32 = 0
        _ = fi15
        _ = v19
    }
    v1[(v2 / ((v2) | 1))] = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
