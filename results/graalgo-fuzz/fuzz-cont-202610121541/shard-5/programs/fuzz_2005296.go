package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
func genFn0() uint16 {
    var v6 int64 = 233
    var v7 int16 = 1000
    if false {
        v6 = 680
        if false {
            var v8 int64 = v6
            _ = v8
        } else {
            for fi9 := 0; fi9 < 18; fi9++ {
                for fi10 := 0; fi10 < 10; fi10++ {
                    v6 = 42
                    _ = fi10
                }
                _ = fi9
            }
            v11 := GenS1{F0: 1, F1: -9223372036854775808}
            _ = v11
        }
    }
    _ = v6
    _ = v7
    return 673
}
func genFn1() uint32 {
    var v12 uint32 = 65535
    v12 = 1
    _ = v12
    return v12
}

func main() {
    v1 := []int{42}
    var v2 uint = 65535
    var v3 int8 = 100
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]bool)
    var v13 uint64 = 1
    v3 += 0
    v14 := GenS1{F0: -2147483648, F1: -450}
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 uint8 = uint8(v13)
        _ = fi15
        _ = v16
    }
    var v17 string = "a"
    delete(v4, 42)
    if false {
        fmt.Println(v3)
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
    }
    var v18 string = (("\n\t" + (v17 + v17)) + (v17 + "\n\t"))
    v2 += (v2 - (v2 | uint(v13)))
    fmt.Println(v13, v13, v2)
    if false {
        if (-1 != -959) {
            v19 := GenS1{F0: 32768, F1: -1}
            switch "a" {
            case "\n\t":
                v13 += (v13 % ((v13) | 1))
            }
            v20 := GenS1{F0: 32768, F1: 0}
            _ = v19
            _ = v20
        } else {
            if false {
                if (!(false || (!true))) {
                    if (42 <= 682) {
                        for fi21 := 0; fi21 < 19; fi21++ {
                            var v22 uint32 = 471
                            var v23 float32 = (((-2.25 * (1.5 * 1e10)) * (1e10 - 1e10)) - (-2.25 + (float32((1.5 / 1e10)) / ((1e10 / 1.0) / float32(-2.25)))))
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                    } else {
                        v18 = "\n\t"
                        var v24 bool = (3.14159 < 3.14159)
                        v2 -= v2
                        _ = v24
                    }
                    var v25 uint = v2
                    _ = v25
                }
                var v26 int = 28
                _ = v26
            }
            _ = len(v4)
            v18 = (v18 + (v18 + v18))
        }
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 uint = 255
            _ = fi27
            _ = v28
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
