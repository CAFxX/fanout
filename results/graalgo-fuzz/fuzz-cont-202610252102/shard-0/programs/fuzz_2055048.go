package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int32; F1 int64 }
func genFn0(p0 uint32, p1 string, p2 int8) int8 {
    for fi4 := 0; fi4 < 10; fi4++ {
        for fi5 := 0; fi5 < 18; fi5++ {
            var v6 int = -563
            _ = fi5
            _ = v6
        }
        v7 := GenS1{F0: 18446744073709551615, F1: p2, F2: 91}
        v8 := GenS1{F0: 42, F1: (^(p2)), F2: 2}
        _ = fi4
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: 9, F1: 911}
    p0 += p0
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p2
}
func genFn1(p0 uint32, p1 int32, p2 int32) int {
    v10 := []int{4294967295, -259, 2147483647}
    var v11 int = 4294967295
    var v12 float64 = -2.25
    if (!false) {
        v11 = (v11 - 9223372036854775807)
    } else {
        for fi13 := 0; fi13 < 7; fi13++ {
            p2 = (p1 ^ (p2 * (p1 - (p2 + 1))))
            _ = fi13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return (v11 << ((18446744073709551615) & 63))
}

func main() {
    var v1 int = 463
    v2 := []string{}
    var v3 int8 = -128
    for fi14 := 0; fi14 < 5; fi14++ {
        v1 = genFn1(65535, -1, 42)
        _ = fi14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        v16 := GenS3{F0: 1, F1: -186}
        var v17 int = fi15
        fmt.Println(v17)
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 580, F1: v3, F2: 2}
    if false {
        if true {
            v19 := GenS2{F0: 2, F1: 4294967295}
            fmt.Println(v3)
            _ = v19
        } else {
            v20 := GenS2{F0: 529, F1: 42}
            if (1 <= 1) {
                var v21 int64 = 9223372036854775807
                _ = v21
            } else {
                if (2 > 42) {
                    fmt.Println(v1)
                } else {
                    _ = len(v2)
                    v22 := GenS2{F0: 42, F1: 255}
                    _ = v22
                }
                fmt.Println(v1, v1)
                var v23 uint32 = 4294967295
                _ = v23
            }
            _ = v20
        }
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v24 int32 = 0
        v3 += (v3 ^ 100)
        _ = v24
    }
    v3 -= 1
    v3 -= (^(v3))
    v2 = append(v2, (("foo bar" + "\n\t") + ("" + "foo bar")))
    var v25 float32 = 0.0
    v26 := GenS2{F0: 665, F1: 255}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
