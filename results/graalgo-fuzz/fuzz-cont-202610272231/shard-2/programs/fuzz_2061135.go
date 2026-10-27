package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int64; F2 int32 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 bool = true
    var v2 int32 = 42
    v3 := []uint32{4294967295, 822}
    var v4 string = ("a" + "")
    var v5 int = -9223372036854775808
    v6 := make(map[int]string)
    if ((!v1) && (0.1 > -2.25)) {
        if (602 < -1) {
            if true {
                var v7 int32 = 42
                v7 -= v2
                _ = v7
            } else {
                var v8 bool = v1
                var v9 int8 = -4
                _ = len(v6)
                _ = v8
                _ = v9
            }
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 bool = (((v2 + 42) != (v2 % ((v2) | 1))) && v1)
                fmt.Println(v5)
                v5 += fi10
                _ = fi10
                _ = v11
            }
        } else {
            var v12 string = (("foo bar" + v4) + v4)
            v13 := GenS1{F0: 1}
            _ = v12
            _ = v13
        }
        var v14 int16 = 1000
        _ = v14
    }
    var v15 int32 = (v2 & 460)
    v16 := GenS2{F0: (v15 ^ v2), F1: -9223372036854775808, F2: v2}
    var v17 float64 = (-2.25 + (-2.25 + (1e10 * (((0.1 - 3.14159) - -2.25) + 3.14159))))
    if false {
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 bool = (42 >= 2147483647)
            v15 += v15
            _ = fi18
            _ = v19
        }
        var v20 uint32 = 745
        _ = v20
    } else {
        var v21 int32 = v2
        v5 -= (v5 + v5)
        _ = v21
    }
    var v22 uint32 = 2147483648
    v2 += v2
    v1 = ((("hello" + v4) >= "") && (0.1 < float32(-2.25)))
    if ((!true) || (!v1)) {
        var v23 uint = 1
        _ = v23
    } else {
        if false {
            var v24 uint8 = 128
            v15 += (v15 + -234)
            _ = v24
        } else {
            v2 = 1
        }
    }
    if false {
        var v25 uint = 65535
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v26 := GenS3{F0: 24}
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        v4 = v4
        _ = fi27
    }
    var v28 uint16 = 65535
    if len(v3) > 0 {
        v3[0] = (^(v22))
    }
    var v29 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
