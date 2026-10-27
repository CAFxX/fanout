package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint32 }

func main() {
    var v1 int = 32768
    v2 := make(map[uint32]bool)
    var v3 bool = true
    var v4 string = "foo bar"
    if v3 {
        v4 = v4
        var v5 int = v1
        v6 := GenS1{F0: (v4 + "\n\t"), F1: 1, F2: v4}
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := []int{}
            v3 = (v3 || (!false))
            v1 -= v1
            _ = fi7
            _ = v8
        }
        v3 = (1 < v1)
        v4 = v4
    }
    v4 = ((v4 + v4) + v4)
    var v9 uint = 2
    v10 := GenS2{F0: 170}
    fmt.Println(v4, v1, v1)
    v9 = v9
    if (v4 == v4) {
        v11 := GenS1{F0: v4, F1: 2, F2: v4}
        _ = v11
    } else {
        var v12 uint8 = 126
        var v13 uint8 = (v12 % ((11) | 1))
        v14 := GenS3{F0: 2147483647, F1: 953, F2: 2147483647}
        _ = v12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        if v3 {
            v1 += v1
            var v16 uint = v9
            _ = v16
        } else {
            if ((v1 >> ((915) & 63)) > (fi15 | (fi15 + fi15))) {
                fmt.Println(v4)
                for fi17 := 0; fi17 < 20; fi17++ {
                    v18 := GenS1{F0: v4, F1: 4294967295, F2: "foo bar"}
                    v19 := []int{(fi15 * v1), (v1 >> (((v9 - v9)) & 63)), (v1 + (fi17 << ((18446744073709551615) & 63))), (fi17 >> ((0) & 63))}
                    v1 += fi17
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
        }
        if false {
            var v20 float32 = 0.0
            v3 = (!(((1.5 + float64(3.14159)) / 1e10) <= ((3.14159 + 0.0) - -2.25)))
            _ = v20
        } else {
            var v21 uint32 = 515
            _ = v21
        }
        _ = fi15
    }
    _ = len(v2)
    var v22 string = "foo bar"
    var v23 float32 = (0.0 * 1e10)
    v1 = 274
    var v24 uint32 = 4294967295
    switch v1 {
    case -657:
        v25 := []int{v1}
        _ = v25
    default:
        switch v1 {
        case 2147483647:
            switch (v1 % ((v1) | 1)) {
            case -462:
                var v26 int64 = -1
                _ = v26
            }
        case 32768:
            for fi27 := 0; fi27 < 13; fi27++ {
                var v28 float64 = 1.5
                v29 := GenS2{F0: -1}
                _ = fi27
                _ = v28
                _ = v29
            }
        default:
            var v30 uint8 = 2
            _ = v30
        }
    }
    v24 -= (v24 ^ v24)
    if (-1 <= 1) {
        fmt.Println(v24, v3, v22)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
