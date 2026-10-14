package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint16 = 42
    var v2 string = "hello"
    v3 := make(map[uint32]string)
    var v4 uint = 619
    if ((v2 < v2) && true) {
        fmt.Println(v1)
        v1 = (v1 - (v1 - v1))
    }
    v5 := GenS1{F0: 0}
    var v6 uint = (v4 & v4)
    fmt.Println(v6, v4)
    v6 = (^(v4))
    v4 = 736
    if (2 <= (v1 >> ((v4) & 63))) {
        v7 := GenS1{F0: 0}
        for fi8 := 0; fi8 < 18; fi8++ {
            fmt.Println(fi8, v6, v4)
            var v9 float64 = 1.5
            v4 += 2
            _ = fi8
            _ = v9
        }
        v2 = "a"
        _ = v7
    }
    v2 = ""
    v10 := v3[42]
    switch (v2 + "\n\t") {
    case "a":
        fmt.Println(v10, v2)
    case "foo bar":
        v11 := v3[2147483647]
        _ = v11
    default:
        v6 = 396
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int16 = int16(v4)
        switch -134 {
        case 679:
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 uint8 = 255
                _ = fi14
                _ = v15
            }
        case 127:
            v10 = v10
        }
        var v16 int16 = (v13 << ((v4) & 63))
        _ = fi12
        _ = v13
        _ = v16
    }
    var v17 bool = false
    for fi18 := 0; fi18 < 3; fi18++ {
        if v17 {
            var v19 int32 = 179
            _ = v19
        }
        var v20 int32 = -1
        _ = fi18
        _ = v20
    }
    v21 := []int{-9223372036854775808, 1}
    if v17 {
        for fi22 := 0; fi22 < 12; fi22++ {
            if v17 {
                for fi23 := 0; fi23 < 18; fi23++ {
                    fmt.Println(v17, v10, fi23)
                    v6 = (v6 / ((1) | 1))
                    v10 = v10
                    _ = fi23
                }
                v10 = "foo bar"
            } else {
                var v24 uint8 = 0
                var v25 int32 = -1
                _ = v24
                _ = v25
            }
            v26 := GenS2{F0: 1000}
            var v27 uint = (v6 - (v6 | (v4 / ((18446744073709551615) | 1))))
            _ = fi22
            _ = v26
            _ = v27
        }
    }
    fmt.Println(v17)
    var v28 uint16 = (v1 / ((476) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
