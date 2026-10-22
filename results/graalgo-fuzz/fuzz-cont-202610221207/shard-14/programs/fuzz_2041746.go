package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
func genFn0(p0 uint32, p1 int32) uint8 {
    if (!(((!(!(32767 >= 0))) && false) || true)) {
        v5 := GenS1{F0: 100, F1: (false && true), F2: 32767}
        for fi6 := 0; fi6 < 19; fi6++ {
            v7 := GenS1{F0: -86, F1: true, F2: uint16(p0)}
            v8 := GenS2{F0: 128}
            _ = fi6
            _ = v7
            _ = v8
        }
        switch "a" {
        case "foo bar":
            switch -223 {
            case 1:
                p1 -= p1
            }
        }
        _ = v5
    }
    v9 := GenS3{F0: 1}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v9
    return 178
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = -1
    var v3 float64 = -2.25
    v4 := make(map[int]int)
    v4[int(v2)] = 2147483647
    if false {
        if (1.5 == float32(v3)) {
            v2 += v2
            fmt.Println(v2, v3)
        }
        if true {
            var v10 float32 = (1.5 - 1e10)
            switch "" {
            case "foo bar":
                v3 = v3
            }
            _ = v10
        } else {
            v3 = 0.1
        }
        v2 = v2
    } else {
        if true {
            fmt.Println(v3, v3)
        } else {
            v3 += v3
            v3 = 1.5
        }
    }
    var v11 bool = false
    var v12 bool = ((2147483647 == 2) || (802 > 1))
    v3 = 1e10
    var v13 int32 = 753
    switch 4294967295 {
    case -351:
        v13 -= v13
    default:
        v4[32768] = 32768
    }
    switch "a" {
    case "":
        switch 127 {
        case -1:
            for fi14 := 0; fi14 < 6; fi14++ {
                v15 := GenS4{F0: 42}
                if v12 {
                    switch (fi14 + (fi14 % (((-(fi14))) | 1))) {
                    case 1:
                        var v16 float64 = (-2.25 * ((v3 - 0.1) / ((v3 - ((v3 - v3) - (v3 / v3))) / v3)))
                        _ = v16
                    case 32768:
                        var v17 float32 = (float32(v3) + 3.14159)
                        _ = v17
                    case -2147483648:
                        fmt.Println(v12, v3, v3)
                    }
                }
                _ = fi14
                _ = v15
            }
        case 32767:
            fmt.Println(v13, v11)
        }
    default:
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 int64 = (^(v2))
            _ = fi18
            _ = v19
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
