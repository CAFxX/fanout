package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint32; F1 int8; F2 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int64; F1 uint; F2 int16 }
type GenS6 struct { F0 int16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 string = "a"
    var v4 float32 = (-2.25 * 0.1)
    v4 = 3.14159
    if true {
        v3 = (v3 + "hello")
    } else {
        fmt.Println(v4)
    }
    v5 := GenS1{F0: 264, F1: 586}
    v6 := GenS2{F0: 0, F1: -124, F2: -1}
    var v7 int = -688
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: -9223372036854775808, F1: -981}
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 42
    v11 := GenS1{F0: 42, F1: v7}
    if false {
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS1{F0: 1, F1: 127}
            v4 -= 1e10
            var v14 float64 = 1e10
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := GenS3{F0: (v7 % ((42) | 1))}
        _ = v15
    }
    fmt.Println(v3)
    var v16 uint64 = 0
    v17 := GenS1{F0: -2, F1: int(v10)}
    switch v7 {
    case 32768:
        var v18 int64 = 617
        _ = v18
    case 811:
        for fi19 := 0; fi19 < 10; fi19++ {
            var v20 int16 = 42
            v4 = (v4 + v4)
            _ = fi19
            _ = v20
        }
    case 412:
        v7 += 127
    default:
        var v21 string = "hello"
        _ = v21
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        if true {
            for fi23 := 0; fi23 < 7; fi23++ {
                v10 -= (v10 % ((113) | 1))
                var v24 uint = 2
                var v25 int = v7
                _ = fi23
                _ = v24
                _ = v25
            }
            var v26 bool = false
            if v26 {
                switch "a" {
                case "\n\t":
                    var v27 string = v3
                    _ = v27
                case "a":
                    fmt.Println(v3, v3)
                default:
                    v3 = "foo bar"
                }
            } else {
                var v28 int64 = -1
                v29 := GenS4{F0: -2147483648, F1: (fi22 - (-(fi22)))}
                _ = v28
                _ = v29
            }
            _ = v26
        }
        for fi30 := 0; fi30 < 13; fi30++ {
            switch 32767 {
            case 32767:
                v16 -= (^(v16))
            case 32768:
                v31 := GenS5{F0: 1, F1: 757, F2: 1000}
                _ = v31
            case -9223372036854775808:
                var v32 int64 = -152
                _ = v32
            }
            fmt.Println(fi30, v4, v10)
            if (-665 >= 0) {
                var v33 int = v7
                _ = v33
            }
            _ = fi30
        }
        switch v3 {
        case "":
            for fi34 := 0; fi34 < 5; fi34++ {
                var v35 uint16 = 1
                v36 := GenS5{F0: 1, F1: 255, F2: 1000}
                v37 := GenS6{F0: -32768}
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
            }
        }
        _ = fi22
    }
    v38 := GenS6{F0: 1000}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
