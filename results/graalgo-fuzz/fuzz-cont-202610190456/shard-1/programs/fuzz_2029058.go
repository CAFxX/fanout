package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 bool }

func main() {
    v1 := make(map[int]int)
    var v2 bool = (false && ((1 > 255) && true))
    var v3 float64 = (1e10 * (1e10 * -2.25))
    var v4 uint64 = 4294967295
    var v5 int8 = 0
    v6 := GenS1{F0: 2147483648, F1: 2, F2: false}
    v7 := GenS1{F0: 2147483647, F1: 1, F2: false}
    v1[777] = 32768
    if (32767 > 819) {
        v2 = v2
        var v8 int64 = 1
        var v9 uint8 = 140
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 1, F1: 65535, F2: (!(42 < 1))}
    if v2 {
        v11 := GenS1{F0: 164, F1: 0, F2: (18446744073709551615 == (^(v4)))}
        _ = v11
    } else {
        if (638 >= v4) {
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 float32 = 0.1
                var v14 int64 = 42
                _ = fi12
                _ = v13
                _ = v14
            }
            if v2 {
                if (0 >= 1) {
                    if (false || v2) {
                        v2 = v2
                        v4 -= (v4 >> ((2) & 63))
                        v3 += float64((3.14159 / (0.1 / (1.5 - -2.25))))
                    } else {
                        v15 := []int{0}
                        _ = v15
                    }
                    var v16 uint32 = 1
                    v17 := []int{32768, 32767}
                    _ = v16
                    _ = v17
                }
            }
            var v18 string = ((("hello" + "") + "") + "")
            _ = v18
        } else {
            var v19 int = 2147483647
            switch (v19 >> ((255) & 63)) {
            case -1:
                v5 = -102
            case 9223372036854775807:
                var v20 int32 = int32(v5)
                _ = v20
            }
            _ = v19
        }
        fmt.Println(v2, v2, v2)
    }
    v21 := GenS1{F0: 65535, F1: 588, F2: true}
    for fi22 := 0; fi22 < 17; fi22++ {
        fmt.Println(fi22, v3, fi22)
        _ = fi22
    }
    v4 -= (v4 >> ((296) & 63))
    v2 = ((v3 / (float64(-2.25) / v3)) != (((v3 - 0.0) * ((v3 - v3) * v3)) * v3))
    var v23 float32 = 1.5
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 uint = 4294967295
        _ = fi24
        _ = v25
    }
    v23 -= 1.5
    if true {
        var v26 int64 = 1
        switch 32767 {
        case 2147483647:
            _ = len(v1)
        case 0:
            var v27 string = "foo bar"
            _ = v27
        default:
            fmt.Println(v23, v23, v2)
        }
        for fi28 := 0; fi28 < 13; fi28++ {
            v2 = (!(766 < 2))
            switch (fi28 ^ fi28) {
            case 2147483647:
                if (-2147483648 <= (fi28 * fi28)) {
                    var v29 int64 = -9223372036854775808
                    v30 := []int{}
                    _ = v29
                    _ = v30
                } else {
                    var v31 float64 = (v3 + 1e10)
                    _ = v31
                }
            case -67:
                v2 = ("foo bar" > "hello")
            case 32768:
                v2 = (207 > 0)
            default:
                var v32 int64 = v26
                _ = v32
            }
            _ = fi28
        }
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
