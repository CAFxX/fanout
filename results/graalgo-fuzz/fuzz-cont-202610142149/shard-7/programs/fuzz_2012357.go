package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint16) int64 {
    v5 := []int{}
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    var v3 int16 = 32767
    var v4 uint = 2
    fmt.Println(v4, v3)
    var v6 uint16 = 65535
    var v7 uint16 = uint16(v4)
    var v8 int16 = (^(v3))
    delete(v2, 2147483648)
    v9 := []int{32767, 46, 42}
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(v6, v7, v7)
        if false {
            var v11 float32 = 1e10
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            var v12 uint8 = 206
            _ = v11
            _ = v12
        }
        _ = fi10
    }
    v6 -= (^(v6))
    for fi13 := 0; fi13 < 13; fi13++ {
        fmt.Println(fi13)
        v6 = v7
        if len(v9) > 0 {
            v9[0] = fi13
        }
        _ = fi13
    }
    var v14 int64 = 42
    var v15 uint8 = 52
    v9 = append(v9, 128)
    switch "foo bar" {
    case "a":
        var v16 bool = true
        _ = v16
    default:
        v17 := GenS1{F0: (v15 / ((2) | 1))}
        _ = v17
    }
    v7 = (v6 ^ (v7 & (v6 << ((42) & 63))))
    switch "" {
    case "":
        if (!(42 <= v15)) {
            switch 9223372036854775807 {
            case 4294967295:
                v14 += (v14 | (v14 - int64(v8)))
            case -155:
                if (false && false) {
                    v15 = v15
                } else {
                    var v18 float64 = -2.25
                    _ = v18
                }
            case 725:
                for fi19 := 0; fi19 < 17; fi19++ {
                    for fi20 := 0; fi20 < 10; fi20++ {
                        var v21 int8 = 100
                        _ = fi20
                        _ = v21
                    }
                    _ = fi19
                }
            }
            fmt.Println(v8)
        }
    case "a":
        v9 = append(v9, -542)
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        switch 32767 {
        case 633:
            v2[4294967295] = "foo bar"
        }
        if len(v9) > 0 {
            v9[0] = (^(fi22))
        }
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
