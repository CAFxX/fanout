package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int8; F2 uint64 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int64, p1 int) uint16 {
    fmt.Println(p0, p1)
    p1 = p1
    if (66 > 42) {
        p0 += p0
        v6 := GenS1{F0: -2147483648}
        _ = v6
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    var v4 uint = 2
    var v5 uint64 = 1
    v7 := GenS2{F0: -9223372036854775808, F1: 1, F2: 2}
    v4 -= v4
    var v8 uint64 = (v5 / ((853) | 1))
    v8 = (v8 >> ((v4) & 63))
    switch 32768 {
    case 993:
        var v9 int64 = 9223372036854775807
        _ = v9
    default:
        if true {
            fmt.Println(v4, v4)
        } else {
            v10 := []int{127, 317, 0, -2147483648}
            _ = v10
        }
    }
    var v11 int16 = 1
    switch "" {
    case "a":
        if false {
            var v12 uint32 = 2
            _ = v12
        } else {
            var v13 uint = v4
            _ = v13
        }
    }
    switch 128 {
    case -2147483648:
        fmt.Println(v11, v5)
    }
    v5 = uint64(v4)
    if false {
        v8 -= v5
        v8 -= v8
    } else {
        for fi14 := 0; fi14 < 13; fi14++ {
            v11 -= 0
            var v15 float64 = 0.1
            _ = fi14
            _ = v15
        }
        v11 = 42
    }
    if false {
        var v16 int64 = 487
        for fi17 := 0; fi17 < 3; fi17++ {
            v11 = (^(v11))
            _ = fi17
        }
        _ = v16
    } else {
        v18 := GenS2{F0: -338, F1: -113, F2: (v5 << ((520) & 63))}
        v19 := []int{514, 127}
        v5 = v5
        _ = v18
        _ = v19
    }
    v20 := GenS3{F0: 0, F1: 42}
    if true {
        var v21 bool = true
        _ = v21
    } else {
        switch 127 {
        case 4294967295:
            if false {
                v22 := GenS4{F0: 0, F1: 2147483647}
                _ = v22
            } else {
                fmt.Println(v8)
                v5 = v8
                fmt.Println(v4, v5)
            }
        case -9223372036854775808:
            switch 42 {
            case 42:
                _ = len(v2)
            case -387:
                var v23 bool = (0 <= 1)
                _ = v23
            case 1:
                if (false && false) {
                    var v24 bool = (!(27 != v4))
                    if ("\n\t" <= "a") {
                        v25 := []int{-2147483648, -1, 32767, 2147483647}
                        var v26 float64 = ((0.1 - 1e10) / ((-2.25 + 3.14159) * -2.25))
                        fmt.Println(v26)
                        _ = v25
                        _ = v26
                    } else {
                        v27 := GenS1{F0: -452}
                        fmt.Println(v4)
                        _ = v27
                    }
                    _ = v24
                }
            }
        default:
            v28 := GenS2{F0: -1, F1: -116, F2: 4294967295}
            _ = v28
        }
        v11 = v11
    }
    var v29 int = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
