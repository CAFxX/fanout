package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 bool }
type GenS2 struct { F0 uint16; F1 int; F2 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int64; F2 int16 }
func genFn0(p0 int8, p1 uint) int16 {
    fmt.Println(p1, p0)
    for fi5 := 0; fi5 < 20; fi5++ {
        p1 += (p1 ^ (p1 | p1))
        var v6 uint64 = 1
        var v7 uint16 = 255
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    return 1000
}
func genFn1(p0 int) bool {
    p0 += -2147483648
    if false {
        switch p0 {
        case 32768:
            v8 := GenS1{F0: 0, F1: -28, F2: false}
            _ = v8
        case 827:
            p0 += p0
        case -9223372036854775808:
            var v9 float32 = -2.25
            _ = v9
        }
    }
    v10 := GenS1{F0: 877, F1: 42, F2: (true || true)}
    if true {
        v11 := GenS2{F0: 453, F1: 42, F2: 255}
        if false {
            fmt.Println(p0)
            for fi12 := 0; fi12 < 7; fi12++ {
                fmt.Println(fi12, fi12)
                _ = fi12
            }
        } else {
            if (false && true) {
                v13 := GenS2{F0: 637, F1: (p0 & 32768), F2: 0}
                _ = v13
            } else {
                var v14 uint64 = 18446744073709551615
                if true {
                    v15 := []int{p0, (p0 | 128), p0, 2147483647}
                    _ = v15
                }
                p0 += p0
                _ = v14
            }
            if (883 < 0) {
                v16 := GenS1{F0: 2, F1: 782, F2: true}
                p0 += -9223372036854775808
                var v17 uint32 = 661
                _ = v16
                _ = v17
            } else {
                var v18 float64 = -2.25
                var v19 uint32 = 0
                _ = v18
                _ = v19
            }
            var v20 float64 = 1.5
            _ = v20
        }
        v21 := GenS2{F0: 874, F1: p0, F2: 91}
        _ = v11
        _ = v21
    }
    _ = p0
    _ = v10
    return true
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[string]int)
    var v3 uint = 2
    var v4 int32 = 0
    v22 := []int{-2147483648, 817}
    if (uint8(v1) >= 128) {
        for fi23 := 0; fi23 < 19; fi23++ {
            var v24 int16 = -32768
            for fi25 := 0; fi25 < 5; fi25++ {
                var v26 uint64 = 2
                var v27 uint8 = 127
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = fi23
            _ = v24
        }
    }
    var v28 uint32 = (v1 * 2)
    fmt.Println(v4, v3)
    var v29 int32 = (v4 % ((0) | 1))
    fmt.Println(v4)
    v30 := GenS2{F0: 255, F1: 42, F2: 0}
    var v31 float64 = 0.0
    v32 := GenS3{F0: -2147483648}
    v33 := GenS4{F0: 127, F1: -1, F2: -729}
    if true {
        fmt.Println(v29, v4)
        var v34 int64 = 9223372036854775807
        var v35 float64 = (v31 * (v31 - (v31 * 1e10)))
        _ = v34
        _ = v35
    } else {
        switch 32768 {
        case 42:
            if ((280 <= (v29 + (v4 >> (((^(v3))) & 63)))) || ((true && (true || false)) || (false && true))) {
                var v36 string = ("foo bar" + "hello")
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                switch 42 {
                case 434:
                    var v37 float64 = (1.5 / v31)
                    _ = v37
                case 2147483647:
                    v29 -= v29
                case -622:
                    v3 -= (v3 % (((^(v3))) | 1))
                }
                _ = v36
            } else {
                switch 127 {
                case -1:
                    v28 -= (v1 | 0)
                }
                if len(v22) > 0 {
                    v22[0] = -685
                }
            }
        case -9223372036854775808:
            v31 = 1.5
        case 1:
            for fi38 := 0; fi38 < 12; fi38++ {
                for fi39 := 0; fi39 < 10; fi39++ {
                    var v40 int32 = (v29 & v4)
                    _ = fi39
                    _ = v40
                }
                var v41 int8 = 111
                _ = fi38
                _ = v41
            }
        default:
            v1 -= v28
        }
        if (!(true || (((v29 ^ v29) != v4) && (!((-128 < -1) && (!(42 >= 1000))))))) {
            v3 = (^(v3))
            v4 -= (v4 % (((v29 * v29)) | 1))
        }
        switch 62 {
        case 127:
            _ = len(v22)
        case -9223372036854775808:
            v42 := GenS2{F0: 42, F1: -953, F2: 575}
            _ = v42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
