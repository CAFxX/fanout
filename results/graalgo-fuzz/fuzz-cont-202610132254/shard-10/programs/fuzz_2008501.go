package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 string = "a"
    var v2 uint32 = 4294967295
    var v3 float64 = 3.14159
    v4 := make(map[int]bool)
    var v5 float32 = 0.0
    var v6 int64 = -9223372036854775808
    var v7 uint32 = 1
    v3 -= v3
    var v8 int64 = 1
    if (0 >= 2) {
        v9 := GenS1{F0: ("" + v1), F1: 435}
        if true {
            v4[127] = (false || (!false))
            for fi10 := 0; fi10 < 20; fi10++ {
                for fi11 := 0; fi11 < 8; fi11++ {
                    var v12 bool = (!(true && true))
                    if (2147483647 >= 42) {
                        if v12 {
                            v7 = v7
                            var v13 uint64 = 360
                            _ = v13
                        } else {
                            switch v1 {
                            case "foo bar":
                                fmt.Println(fi11)
                            }
                            var v14 int32 = -2147483648
                            var v15 int64 = -1
                            _ = v14
                            _ = v15
                        }
                        var v16 int16 = -752
                        _ = v16
                    }
                    _ = fi11
                    _ = v12
                }
                fmt.Println(v6, v3)
                v17 := GenS2{F0: 2147483647}
                _ = fi10
                _ = v17
            }
            var v18 float32 = (1e10 * -2.25)
            _ = v18
        }
        _ = v9
    }
    var v19 int16 = 32767
    v20 := []int{113, 9223372036854775807, -2147483648}
    v1 = v1
    v20 = append(v20, -9223372036854775808)
    var v21 string = v1
    _ = len(v20)
    var v22 float32 = v5
    var v23 int16 = v19
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 float32 = ((v5 * v22) * (1e10 - v5))
        _ = fi24
        _ = v25
    }
    if (65535 < 0) {
        var v26 int16 = -975
        for fi27 := 0; fi27 < 6; fi27++ {
            if false {
                _ = len(v20)
            } else {
                var v28 int16 = 42
                if false {
                    v5 = v5
                } else {
                    _ = len(v4)
                    v8 = (v6 >> ((349) & 63))
                }
                _ = v28
            }
            fmt.Println(v3, v6, v21)
            var v29 uint64 = 0
            _ = fi27
            _ = v29
        }
        v1 = (("a" + v21) + ("hello" + (v1 + (("" + v21) + v1))))
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
