package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }

func main() {
    v1 := []int{-9223372036854775808, 32768}
    var v2 uint = 4294967295
    v3 := make(map[uint32]int)
    var v4 uint16 = 482
    var v5 int32 = -2147483648
    var v6 float64 = -2.25
    v2 = 223
    fmt.Println(v2, v5)
    v7 := GenS1{F0: 1, F1: 2}
    if (!true) {
        switch "" {
        case "hello":
            v8 := GenS1{F0: -443, F1: 0}
            _ = v8
        }
    } else {
        v6 += -2.25
    }
    var v9 bool = (!false)
    if false {
        fmt.Println(v2, v4)
    } else {
        v6 = (v6 + (v6 + v6))
        v4 -= (v4 >> ((v2) & 63))
        var v10 float64 = (3.14159 + v6)
        _ = v10
    }
    _ = len(v3)
    var v11 int64 = 0
    var v12 float64 = (0.0 / v6)
    v11 = v11
    switch 4294967295 {
    case 32768:
        v3[0] = 32768
    case -851:
        for fi13 := 0; fi13 < 18; fi13++ {
            if (false && (2147483647 > 1)) {
                v1 = append(v1, 493)
                var v14 int64 = (v11 | (v11 + (v11 & 1)))
                fmt.Println(v2)
                _ = v14
            } else {
                var v15 float64 = float64((0.1 + -2.25))
                if (-241 > 42) {
                    v9 = false
                    if v9 {
                        fmt.Println(v12, v15)
                        delete(v3, 42)
                    } else {
                        v5 = -309
                    }
                    var v16 float64 = (1.5 - float64((-2.25 - 1e10)))
                    _ = v16
                }
                _ = v15
            }
            v17 := GenS1{F0: v11, F1: 2}
            _ = fi13
            _ = v17
        }
    case 1:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v18 uint32 = 2147483648
    v2 = (v2 % (((v2 & v2)) | 1))
    if (452 != uint16(v18)) {
        for fi19 := 0; fi19 < 20; fi19++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v12 += v6
            _ = fi19
        }
        switch -9223372036854775808 {
        case -1:
            v18 += 0
        default:
            v4 += (v4 * 255)
        }
        switch 0 {
        case -1:
            delete(v3, (v18 - v18))
        case -282:
            var v20 float32 = 1.5
            _ = v20
        }
    } else {
        var v21 uint32 = 1
        fmt.Println(v11, v12)
        for fi22 := 0; fi22 < 9; fi22++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi22
        }
        _ = v21
    }
    if len(v1) > 0 {
        v1[0] = 1
    }
    v1 = append(v1, 0)
    var v23 uint64 = 2
    delete(v3, v18)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
