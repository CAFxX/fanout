package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int16 }
func genFn0(p0 int32) int32 {
    var v7 uint8 = 128
    if ((928 >= 801) || false) {
        v8 := GenS1{F0: 22}
        _ = v8
    } else {
        p0 = p0
    }
    _ = p0
    _ = v7
    return (p0 - -2147483648)
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[uint32]int)
    var v3 bool = (false && (!true))
    var v4 int16 = 42
    v5 := make(map[uint32]int)
    var v6 int16 = 1000
    fmt.Println(v4)
    var v9 float64 = 3.14159
    var v10 float32 = 3.14159
    v4 += 31
    v1 += v1
    var v11 bool = v3
    v12 := []int{}
    var v13 uint8 = 79
    var v14 uint32 = v1
    for fi15 := 0; fi15 < 14; fi15++ {
        v16 := GenS2{F0: 42, F1: (v4 % ((v4) | 1))}
        v4 += (v4 << ((2) & 63))
        _ = fi15
        _ = v16
    }
    var v17 int8 = 100
    v3 = v3
    if v3 {
        if ((v14 + v14) >= 2147483648) {
            v1 = 160
            v13 -= 235
            if (v6 <= v4) {
                fmt.Println(v10, v11, v6)
            } else {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            }
        } else {
            var v18 string = ("hello" + "a")
            v19 := GenS2{F0: int64(v14), F1: int16(v17)}
            _ = v18
            _ = v19
        }
    } else {
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 bool = (4294967295 < 18446744073709551615)
            _ = fi20
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        fmt.Println(v14, v9, v1)
        delete(v2, (v14 - (^(v14))))
        switch 648 {
        case -2147483648:
            var v23 string = "foo bar"
            _ = v23
        }
        _ = fi22
    }
    var v24 int16 = (-(v6))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
