package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint, p1 int64, p2 uint16) string {
    p2 = (^(p2))
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint = (p0 / (((p0 + 625)) | 1))
        var v8 uint = p0
        v8 += (v8 % (((p0 & uint(p2))) | 1))
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int64 = p1
    p0 = (p0 >> ((747) & 63))
    var v10 string = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return (v10 + v10)
}
func genFn1(p0 int32) uint64 {
    fmt.Println(p0)
    var v11 uint16 = 113
    p0 = 1
    _ = p0
    _ = v11
    return 299
}

func main() {
    var v1 string = "hello"
    var v2 int8 = -128
    var v3 uint8 = 1
    var v4 uint32 = 2
    var v5 int32 = -406
    if true {
        v5 -= (v5 & v5)
        var v12 float64 = 1e10
        for fi13 := 0; fi13 < 5; fi13++ {
            fmt.Println(v12)
            for fi14 := 0; fi14 < 10; fi14++ {
                if true {
                    var v15 int16 = 1000
                    _ = v15
                } else {
                    switch (fi13 << ((753) & 63)) {
                    case 127:
                        v16 := []int{}
                        _ = v16
                    default:
                        var v17 uint16 = uint16(v4)
                        _ = v17
                    }
                }
                var v18 bool = (1 > -1)
                _ = fi14
                _ = v18
            }
            _ = fi13
        }
        _ = v12
    }
    v2 -= v2
    v1 = v1
    v19 := GenS1{F0: (v2 + 1), F1: 526}
    v3 = v3
    var v20 uint = 2
    var v21 bool = true
    var v22 uint8 = v3
    v1 = (v1 + "foo bar")
    v23 := []int{127, 4294967295}
    if true {
        v22 += 2
        _ = len(v23)
        v24 := GenS2{F0: v21, F1: 18446744073709551615}
        _ = v24
    }
    var v25 uint32 = (v4 & v4)
    fmt.Println(v3, v21, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
