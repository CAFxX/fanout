package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint64 }

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 42
    var v3 uint8 = 2
    var v4 uint32 = 4294967295
    var v5 string = "a"
    var v6 uint = uint(v4)
    v6 -= v6
    delete(v1, 4294967295)
    var v7 float32 = (0.0 / ((-2.25 / 1.5) - (((-2.25 / 1.5) + 0.1) / 0.1)))
    if ((v4 > v4) && ((^(v4)) != (v4 + v4))) {
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 uint32 = v4
            fmt.Println(v2, v3, v5)
            _ = fi8
            _ = v9
        }
    }
    v10 := GenS1{F0: ((!(((v3 % (((v3 / ((0) | 1))) | 1)) > v3) || false)) && false), F1: false, F2: 1}
    var v11 int16 = 0
    v12 := v1[832]
    _ = len(v1)
    var v13 int = 4294967295
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := v1[(^(v12))]
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 int64 = 527
            if ((true && ((1.5 + float64(v7)) != (float64(1e10) - 1.5))) && (!(42 > 618))) {
                var v18 bool = (v11 == v11)
                _ = v18
            } else {
                var v19 uint32 = 42
                var v20 uint = uint(v4)
                _ = v19
                _ = v20
            }
            fmt.Println(v2, v13)
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    v4 = v4
    v21 := GenS1{F0: (true || true), F1: (v4 <= 2147483648), F2: uint64(v4)}
    var v22 int32 = 573
    if false {
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 string = v5
            if (false && true) {
                v6 += 42
                for fi25 := 0; fi25 < 19; fi25++ {
                    v13 += 42
                    _ = fi25
                }
            } else {
                fmt.Println(v3, v13, v7)
            }
            _ = fi23
            _ = v24
        }
        var v26 uint16 = (v2 | (v2 | v2))
        for fi27 := 0; fi27 < 2; fi27++ {
            v28 := GenS1{F0: ("a" != v5), F1: (!false), F2: 1}
            fmt.Println(v26)
            _ = fi27
            _ = v28
        }
        _ = v26
    } else {
        v6 = v6
    }
    var v29 uint = 255
    v4 -= v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
