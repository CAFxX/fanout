package main

import (
    "fmt"
)

func genFn0(p0 uint32) uint8 {
    var v7 float64 = 1e10
    fmt.Println(v7, p0)
    var v8 uint8 = 42
    _ = p0
    _ = v7
    _ = v8
    return (^(v8))
}

func main() {
    var v1 uint32 = 433
    v2 := []uint32{}
    var v3 int64 = 85
    v4 := make(map[int]int)
    var v5 int64 = -1
    var v6 uint16 = 0
    fmt.Println(v6, v3, v3)
    fmt.Println(v6, v5)
    var v9 int16 = 1
    var v10 bool = true
    v9 -= (v9 ^ (v9 % ((v9) | 1)))
    v1 = 350
    fmt.Println(v5, v5)
    v3 = v3
    switch 925 {
    case 127:
        if len(v2) > 0 {
            v2[0] = (v1 % ((2147483648) | 1))
        }
    case 0:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -432:
        var v11 int64 = (v3 >> ((163) & 63))
        _ = v11
    }
    fmt.Println(v10, v3)
    for fi12 := 0; fi12 < 2; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 uint16 = v6
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 uint = 65535
            v9 = v9
            var v17 int = fi15
            _ = fi15
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            if (v6 == 895) {
                v9 += (v9 + (-(v9)))
            } else {
                switch 42 {
                case 42:
                    var v19 int16 = (^(v9))
                    _ = v19
                case 9223372036854775807:
                    _ = len(v2)
                case 182:
                    var v20 uint8 = 0
                    _ = v20
                default:
                    var v21 float64 = 1e10
                    _ = v21
                }
                delete(v4, 0)
            }
            fmt.Println(v6, v1, v10)
            _ = fi18
        }
        _ = fi12
    }
    var v22 float64 = 0.1
    var v23 uint8 = 1
    var v24 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
