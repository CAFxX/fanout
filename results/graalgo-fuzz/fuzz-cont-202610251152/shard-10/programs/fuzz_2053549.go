package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int8 }

func main() {
    var v1 uint32 = 0
    var v2 string = "\n\t"
    v3 := []int{1, 9223372036854775807, 128}
    var v4 int64 = 390
    v5 := make(map[uint32]uint32)
    if false {
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := GenS1{F0: 65535}
            v2 = v2
            _ = fi6
            _ = v7
        }
        v1 += v1
        if true {
            if ((1 < int32(v1)) && false) {
                v4 = v4
                v8 := v5[v1]
                v9 := GenS2{F0: 250, F1: 100}
                _ = v8
                _ = v9
            } else {
                if (3.14159 >= float32(0.0)) {
                    if ((127 < 2) && (!false)) {
                        v1 += v1
                        var v10 uint32 = (v1 % ((v1) | 1))
                        _ = v10
                    }
                    fmt.Println(v2)
                    v4 -= 0
                }
            }
        } else {
            var v11 uint32 = v1
            _ = v11
        }
    }
    switch ((v2 + (v2 + "foo bar")) + "a") {
    case "":
        for fi12 := 0; fi12 < 4; fi12++ {
            fmt.Println(v4)
            var v13 int8 = -128
            var v14 int64 = v4
            _ = fi12
            _ = v13
            _ = v14
        }
    case "a":
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 uint64 = 16
            var v17 uint32 = (v1 - 1)
            fmt.Println(v17, v1)
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    var v18 int = 32767
    fmt.Println(v1, v4, v2)
    for fi19 := 0; fi19 < 10; fi19++ {
        fmt.Println(v1)
        _ = fi19
    }
    var v20 uint8 = 165
    var v21 uint32 = v1
    var v22 int32 = -2147483648
    var v23 uint32 = (^(v1))
    v1 = v1
    v18 = v18
    v21 -= (v23 * (^(v23)))
    var v24 string = "\n\t"
    if (-55 < -1) {
        var v25 int8 = -62
        for fi26 := 0; fi26 < 15; fi26++ {
            v4 = v4
            var v27 int64 = v4
            fmt.Println(v4, v22)
            _ = fi26
            _ = v27
        }
        _ = v25
    } else {
        var v28 float32 = 1.5
        v5[uint32(v20)] = 2147483648
        var v29 float32 = (v28 + ((float32((-2.25 * 0.1)) / v28) / v28))
        _ = v28
        _ = v29
    }
    var v30 float64 = 0.1
    for fi31 := 0; fi31 < 6; fi31++ {
        _ = len(v5)
        var v32 uint = 0
        _ = fi31
        _ = v32
    }
    v33 := v5[v23]
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
}
