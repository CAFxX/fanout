package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int64; F2 bool }
type GenS3 struct { F0 int16 }
func genFn0(p0 bool, p1 uint) int64 {
    v5 := []int{}
    _ = len(v5)
    switch "a" {
    case "a":
        p1 = (^(p1))
    case "\n\t":
        for fi6 := 0; fi6 < 10; fi6++ {
            for fi7 := 0; fi7 < 16; fi7++ {
                fmt.Println(fi6, p0, fi7)
                fmt.Println(fi6, p1, p0)
                _ = fi7
            }
            var v8 uint = (^(p1))
            v9 := GenS1{F0: p0}
            _ = fi6
            _ = v8
            _ = v9
        }
    case "foo bar":
        p0 = false
    }
    _ = p0
    _ = p1
    _ = v5
    return -719
}

func main() {
    var v1 uint64 = 807
    var v2 int16 = -32768
    v3 := []int{-1, 128, 0}
    var v4 uint16 = uint16(v1)
    v4 = (^(v4))
    _ = len(v3)
    switch "" {
    case "foo bar":
        var v10 uint32 = 2147483647
        _ = v10
    case "a":
        v1 -= (v1 >> ((472) & 63))
    case "\n\t":
        var v11 uint32 = 2
        _ = v11
    }
    v3 = append(v3, -9223372036854775808)
    if false {
        if false {
            v3 = append(v3, 0)
            if false {
                var v12 uint = 354
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 int16 = v2
                    v14 += v14
                    _ = fi13
                    _ = v14
                }
                _ = v12
            } else {
                var v15 int16 = v2
                var v16 int64 = 463
                v2 = v2
                _ = v15
                _ = v16
            }
        }
        v3 = append(v3, 128)
    }
    if false {
        fmt.Println(v2, v1)
        var v17 uint16 = (v4 + (v4 ^ v4))
        _ = v17
    } else {
        v18 := GenS2{F0: -49, F1: 822, F2: false}
        var v19 int8 = 1
        v19 -= v19
        _ = v18
        _ = v19
    }
    v1 = v1
    var v20 string = "hello"
    var v21 int = -1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        if (-1 != -1) {
            if true {
                if (!(false || (255 >= 128))) {
                    var v23 int8 = 42
                    var v24 int32 = -1
                    var v25 uint8 = 2
                    _ = v23
                    _ = v24
                    _ = v25
                } else {
                    if ((-2.25 + 3.14159) == 3.14159) {
                        v26 := GenS3{F0: 42}
                        v20 = v20
                        _ = v26
                    }
                    var v27 float64 = ((0.0 - -2.25) / ((1.5 / 0.1) + (-2.25 * (float64(0.1) + 0.0))))
                    _ = v27
                }
                if len(v3) > 0 {
                    v3[0] = v21
                }
            }
            if ((float32(3.14159) * 0.1) != 3.14159) {
                var v28 int64 = 42
                _ = v28
            }
        }
        var v29 int32 = 2147483647
        v29 += 142
        _ = fi22
        _ = v29
    }
    var v30 float64 = float64(0.1)
    var v31 int32 = -161
    var v32 uint8 = 127
    var v33 uint8 = uint8(v1)
    var v34 int = (v21 & v21)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
