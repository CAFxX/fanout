package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 string; F2 bool }
type GenS3 struct { F0 int64 }
func genFn0() bool {
    var v7 uint32 = 0
    var v8 int = 32768
    _ = v7
    _ = v8
    return true
}
func genFn1(p0 uint16) int64 {
    var v9 uint8 = 255
    var v10 bool = (32768 < 32767)
    for fi11 := 0; fi11 < 6; fi11++ {
        switch "\n\t" {
        case "\n\t":
            var v12 int8 = -128
            _ = v12
        default:
            v13 := []int{(fi11 + 148), (fi11 >> ((255) & 63)), (fi11 + 9223372036854775807), 785}
            _ = v13
        }
        v9 += 81
        _ = fi11
    }
    _ = p0
    _ = v10
    _ = v9
    return 570
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -9223372036854775808
    v3 := make(map[int]bool)
    var v4 float64 = float64(float32((-2.25 * 0.0)))
    v5 := make(map[int]string)
    var v6 bool = (!(180 >= 313))
    v4 = v4
    if (0 < 979) {
        v14 := GenS1{F0: 1}
        _ = v14
    }
    var v15 uint8 = 203
    v16 := []int{970, -9223372036854775808, -9223372036854775808}
    var v17 int16 = -1
    for fi18 := 0; fi18 < 12; fi18++ {
        v15 += (v15 >> ((18446744073709551615) & 63))
        var v19 float64 = 0.0
        if false {
            var v20 float32 = 1.5
            v3[32768] = v6
            _ = v20
        } else {
            fmt.Println(v19)
            switch int(v17) {
            case 979:
                v21 := GenS1{F0: 46}
                _ = v21
            }
            fmt.Println(v2)
        }
        _ = fi18
        _ = v19
    }
    v1[2] = v6
    fmt.Println(v15, v15, v17)
    if (208 < 1) {
        v22 := GenS2{F0: 1, F1: "\n\t", F2: v6}
        _ = v22
    } else {
        if (((v17 >> ((939) & 63)) <= v17) && false) {
            _ = len(v1)
            _ = len(v16)
            var v23 int16 = 1000
            _ = v23
        } else {
            v4 = float64(((0.0 - 0.1) / 1.5))
            if (-998 <= 262) {
                fmt.Println(v4, v2)
            }
            for fi24 := 0; fi24 < 8; fi24++ {
                v17 += v17
                fmt.Println(v15, fi24, v4)
                _ = fi24
            }
        }
        v16 = append(v16, 943)
    }
    v25 := GenS3{F0: -1}
    if (147 > v2) {
        for fi26 := 0; fi26 < 10; fi26++ {
            v4 = v4
            _ = fi26
        }
        for fi27 := 0; fi27 < 16; fi27++ {
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 int32 = -2147483648
                var v30 int = (fi28 >> ((844) & 63))
                _ = fi28
                _ = v29
                _ = v30
            }
            switch "hello" {
            case "hello":
                switch (fi27 ^ (fi27 + (fi27 << ((0) & 63)))) {
                case 32767:
                    for fi31 := 0; fi31 < 11; fi31++ {
                        fmt.Println(v17, v15, v15)
                        v3[fi27] = false
                        _ = fi31
                    }
                case 9223372036854775807:
                    var v32 int64 = 9223372036854775807
                    _ = v32
                default:
                    v2 += (v2 << ((uint(v15)) & 63))
                }
            default:
                var v33 uint64 = uint64(v15)
                _ = v33
            }
            _ = fi27
        }
    } else {
        var v34 bool = v6
        v4 = float64((0.1 - ((0.1 * 1.5) * float32((v4 + v4)))))
        _ = v34
    }
    fmt.Println(v6, v17, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
