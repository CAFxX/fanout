package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64 }
func genFn0() int8 {
    var v5 int64 = 9223372036854775807
    if false {
        if true {
            v5 -= (^(v5))
            for fi6 := 0; fi6 < 11; fi6++ {
                v5 = -9223372036854775808
                _ = fi6
            }
            v5 = v5
        } else {
            var v7 float32 = 1e10
            _ = v7
        }
        switch 1 {
        case 2147483647:
            v8 := GenS1{F0: 2, F1: -1}
            _ = v8
        case 127:
            v9 := GenS2{F0: -103, F1: 255}
            _ = v9
        case -951:
            v10 := []int{128, 9223372036854775807}
            _ = v10
        }
        v11 := GenS1{F0: 255, F1: 1000}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 float64 = (0.1 / 1e10)
            if false {
                switch fi12 {
                case 4294967295:
                    var v14 uint = 395
                    _ = v14
                case 929:
                    v5 = v5
                default:
                    var v15 bool = ((v5 * (v5 << ((255) & 63))) > v5)
                    _ = v15
                }
                v13 = ((1.5 / v13) / v13)
                for fi16 := 0; fi16 < 11; fi16++ {
                    v17 := GenS2{F0: -45, F1: 0}
                    v5 -= (v5 << ((42) & 63))
                    v5 += v5
                    _ = fi16
                    _ = v17
                }
            } else {
                v5 = (v5 | v5)
                var v18 uint64 = 2
                _ = v18
            }
            _ = fi12
            _ = v13
        }
        v5 += (-(v5))
        if (!true) {
            var v19 int8 = 127
            var v20 int64 = v5
            var v21 int = 346
            _ = v19
            _ = v20
            _ = v21
        }
    }
    switch 128 {
    case -2147483648:
        if (!(0 <= 50)) {
            v22 := []int{1, 0}
            fmt.Println(v5)
            _ = v22
        } else {
            if (-1 != 127) {
                v5 -= v5
                switch "\n\t" {
                case "":
                    var v23 uint64 = 1
                    _ = v23
                case "a":
                    v24 := GenS2{F0: 1, F1: 252}
                    _ = v24
                case "hello":
                    v25 := GenS3{F0: 1}
                    _ = v25
                }
            } else {
                if (343 != 4294967295) {
                    v26 := []int{}
                    var v27 uint16 = 2
                    _ = v26
                    _ = v27
                }
            }
            v5 = -9223372036854775808
        }
    case 520:
        var v28 float32 = (-2.25 * 1e10)
        _ = v28
    case 4294967295:
        var v29 float64 = float64(-2.25)
        _ = v29
    default:
        v5 += v5
    }
    _ = v5
    return -128
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{9223372036854775807, -9223372036854775808}
    var v3 bool = ((-2147483648 > 32768) && false)
    v4 := make(map[int]string)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 127 {
    case 2147483647:
        switch 32767 {
        case 891:
            delete(v4, 692)
        case 1:
            v30 := GenS4{F0: 42}
            _ = v30
        case 32767:
            var v31 uint16 = 255
            _ = v31
        default:
            v32 := GenS3{F0: 422}
            _ = v32
        }
    case -153:
        v3 = (!v3)
    default:
        v3 = false
    }
    _ = len(v1)
    _ = len(v2)
    var v33 uint32 = 354
    v33 = 2
    if v3 {
        var v34 float32 = (0.1 - 0.0)
        if v3 {
            v35 := GenS4{F0: -196}
            _ = v35
        }
        _ = v34
    }
    v3 = v3
    v3 = (1 <= -1)
    for fi36 := 0; fi36 < 4; fi36++ {
        if true {
            v37 := GenS2{F0: 0, F1: 255}
            var v38 uint = 1
            v38 = (v38 - v38)
            _ = v37
            _ = v38
        }
        fmt.Println(v3, v3)
        _ = fi36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
