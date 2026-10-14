package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint8, p1 uint, p2 int) uint {
    v4 := GenS1{F0: 0}
    if (18446744073709551615 < 18446744073709551615) {
        var v5 int32 = 42
        var v6 float64 = 3.14159
        fmt.Println(p1)
        _ = v5
        _ = v6
    }
    var v7 uint16 = 264
    p0 += (p0 / ((p0) | 1))
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    _ = v8
    return p1
}
func genFn1(p0 int) int16 {
    v9 := []int{(p0 % ((86) | 1))}
    var v10 uint64 = 725
    if len(v9) > 0 {
        v9[0] = (^(p0))
    }
    _ = p0
    _ = v10
    _ = v9
    return 111
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 890
    v3 := make(map[string]string)
    for fi11 := 0; fi11 < 10; fi11++ {
        if ((51 <= 42) && (42 == 1)) {
            for fi12 := 0; fi12 < 19; fi12++ {
                switch -115 {
                case 128:
                    var v13 uint8 = 0
                    _ = v13
                case 42:
                    var v14 bool = (!false)
                    _ = v14
                case 9223372036854775807:
                    var v15 bool = true
                    fmt.Println(v2, fi11)
                    var v16 uint8 = 2
                    _ = v15
                    _ = v16
                }
                _ = fi12
            }
        } else {
            if true {
                v2 -= v2
            } else {
                v2 -= (v2 >> ((genFn0(22, 477, -420)) & 63))
                var v17 int16 = 1
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(fi18, v2, v2)
            var v19 int64 = 42
            _ = fi18
            _ = v19
        }
        _ = fi11
    }
    var v20 uint = 0
    if ((v20 << (((^(v20))) & 63)) != genFn0(180, 255, 1)) {
        v20 = v20
        v3[""] = (("hello" + "hello") + "hello")
        v20 += v20
    }
    v20 -= (^(v20))
    v21 := GenS2{F0: 0}
    fmt.Println(v2, v20)
    v22 := []int{-2147483648, -552}
    var v23 uint = (^(v20))
    if (!((1 < 377) && true)) {
        fmt.Println(v20, v20, v23)
        v2 = (v2 ^ (v2 >> (((v23 << (((^(v20))) & 63))) & 63)))
    }
    var v24 int32 = 1
    if false {
        v25 := GenS1{F0: 74}
        _ = v25
    } else {
        fmt.Println(v23)
        v20 = uint(v2)
        var v26 int = 32768
        _ = v26
    }
    switch 32767 {
    case -859:
        var v27 uint8 = 42
        _ = v27
    }
    switch 32768 {
    case 111:
        for fi28 := 0; fi28 < 6; fi28++ {
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 int = -9223372036854775808
                var v31 bool = true
                _ = fi29
                _ = v30
                _ = v31
            }
            v32 := GenS3{F0: 1}
            _ = fi28
            _ = v32
        }
    }
    var v33 float32 = float32(((1.5 * (-2.25 + 1e10)) * -2.25))
    var v34 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v34
}
