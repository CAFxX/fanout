package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
func genFn0(p0 uint64, p1 uint16) uint {
    v5 := []int{}
    var v6 int8 = 1
    for fi7 := 0; fi7 < 11; fi7++ {
        v6 -= (v6 / ((v6) | 1))
        _ = fi7
    }
    v6 = (v6 % ((v6) | 1))
    var v8 int64 = 0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v8
    return 21
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    var v3 int64 = 280
    v4 := []uint32{747, 978, 0}
    v3 = (^(v3))
    fmt.Println(v3)
    v2[2147483647] = (("\n\t" + "hello") + "")
    if len(v4) > 0 {
        v4[0] = 42
    }
    v4 = append(v4, 42)
    v9 := GenS1{F0: 1, F1: -1}
    if true {
        var v10 int32 = -2147483648
        _ = v10
    } else {
        var v11 int8 = -1
        v3 = (^(v3))
        for fi12 := 0; fi12 < 10; fi12++ {
            v3 -= 1
            fmt.Println(fi12)
            _ = fi12
        }
        _ = v11
    }
    if (!false) {
        fmt.Println(v3)
    } else {
        for fi13 := 0; fi13 < 8; fi13++ {
            v1[(fi13 % ((-2147483648) | 1))] = 423
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi14 := 0; fi14 < 4; fi14++ {
                if (true && false) {
                    var v15 uint = 65535
                    delete(v1, -1)
                    _ = v15
                }
                _ = fi14
            }
            _ = fi13
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            v3 -= v3
            _ = fi16
        }
    }
    var v17 int = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        v3 -= -1
        v19 := GenS1{F0: 508, F1: 2147483647}
        var v20 uint16 = 0
        _ = fi18
        _ = v19
        _ = v20
    }
    v3 = v3
    var v21 uint = 123
    var v22 uint64 = 765
    var v23 bool = false
    fmt.Println(v23, v23, v22)
    var v24 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
