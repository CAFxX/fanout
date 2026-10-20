package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint; F1 int; F2 uint8 }
type GenS5 struct { F0 int32; F1 string; F2 uint32 }
func genFn0(p0 int8) bool {
    v5 := GenS1{F0: -119, F1: 42}
    var v6 bool = (!(0 <= 4294967295))
    v7 := []int{-2147483648}
    v8 := GenS1{F0: (p0 / ((1) | 1)), F1: 118}
    v9 := GenS2{F0: 18446744073709551615, F1: p0}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (179 > 32767)
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 bool = true
    v3 := []int{-623}
    var v4 uint8 = 127
    if ("foo bar" > "a") {
        if v2 {
            v10 := GenS1{F0: 100, F1: 42}
            var v11 uint = 42
            _ = v10
            _ = v11
        } else {
            var v12 uint32 = 795
            if len(v3) > 0 {
                v3[0] = 42
            }
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 float64 = 0.0
                fmt.Println(v14, fi13)
                var v15 int32 = -886
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        }
        if v2 {
            fmt.Println(v1)
        }
    }
    fmt.Println(v2)
    v16 := GenS1{F0: 0, F1: 165}
    var v17 bool = v2
    for fi18 := 0; fi18 < 3; fi18++ {
        v17 = (1000 < 32767)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi18
    }
    if (128 == (v4 << ((0) & 63))) {
        fmt.Println(v1, v17, v1)
        if v17 {
            var v19 int16 = -248
            _ = v19
        } else {
            if len(v3) > 0 {
                v3[0] = 1
            }
        }
    }
    v2 = v2
    if (-32768 >= 1) {
        var v20 int32 = 42
        _ = v20
    } else {
        fmt.Println(v2, v4, v2)
        if (uint16(v4) > 32767) {
            switch ("a" + ("a" + ("foo bar" + "hello"))) {
            case "a":
                var v21 int64 = v1
                _ = v21
            }
            v22 := GenS3{F0: v1, F1: (v1 / ((v1) | 1))}
            if v17 {
                for fi23 := 0; fi23 < 13; fi23++ {
                    v24 := GenS4{F0: 65535, F1: fi23, F2: (v4 | (v4 << ((4294967295) & 63)))}
                    _ = len(v3)
                    _ = fi23
                    _ = v24
                }
                v1 = (-(v1))
                v25 := GenS5{F0: 82, F1: "", F2: 2147483647}
                _ = v25
            }
            _ = v22
        } else {
            var v26 int8 = 62
            _ = v26
        }
        v4 += v4
    }
    var v27 int32 = -59
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
