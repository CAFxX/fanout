package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int32; F1 bool }
func genFn0(p0 bool, p1 int32) uint64 {
    fmt.Println(p0, p1)
    v4 := []int{839, 32767, -2147483648, 128}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v5 uint64 = 922
    var v6 int8 = 110
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return v5
}

func main() {
    var v1 uint64 = 462
    var v2 int64 = 35
    var v3 uint16 = 32767
    v7 := GenS1{F0: 777, F1: -765}
    v2 = v2
    var v8 bool = (true || true)
    v3 = v3
    var v9 string = ("foo bar" + "")
    for fi10 := 0; fi10 < 8; fi10++ {
        switch fi10 {
        case 925:
            v8 = true
        case -9223372036854775808:
            var v11 int64 = v2
            _ = v11
        case 1:
            var v12 string = (v9 + (v9 + "foo bar"))
            _ = v12
        }
        _ = fi10
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int = fi13
        _ = fi13
        _ = v14
    }
    var v15 float32 = 1e10
    if v8 {
        v16 := []int{0}
        v15 = 1.5
        _ = v16
    }
    v2 += v2
    var v17 bool = ((v8 || (42 != 42)) || v8)
    var v18 uint32 = 1
    if false {
        v18 = v18
        if (((v18 & v18) > 0) && true) {
            v19 := GenS1{F0: 18446744073709551615, F1: -9223372036854775808}
            if (-1 >= -128) {
                var v20 uint8 = 2
                v21 := GenS1{F0: 0, F1: 32768}
                v20 += (v20 + v20)
                _ = v20
                _ = v21
            }
            _ = v19
        } else {
            var v22 int8 = -128
            v23 := []int{130, 309, 1}
            fmt.Println(v17)
            _ = v22
            _ = v23
        }
        v24 := GenS1{F0: 4294967295, F1: 551}
        _ = v24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        if (-694 == 557) {
            switch fi25 {
            case 32767:
                for fi26 := 0; fi26 < 12; fi26++ {
                    v2 += (v2 << ((255) & 63))
                    fmt.Println(fi25, v2)
                    _ = fi26
                }
            case 127:
                v27 := []int{2147483647, (^(fi25)), (^(fi25)), fi25}
                _ = v27
            case -2147483648:
                var v28 string = v9
                _ = v28
            default:
                v29 := GenS1{F0: 42, F1: (fi25 - (^(fi25)))}
                _ = v29
            }
        }
        for fi30 := 0; fi30 < 9; fi30++ {
            v31 := GenS2{F0: 694, F1: false}
            v18 = v18
            _ = fi30
            _ = v31
        }
        for fi32 := 0; fi32 < 4; fi32++ {
            fmt.Println(v8, v18)
            _ = fi32
        }
        _ = fi25
    }
    v33 := []int{-278, -155, -9223372036854775808, 1}
    var v34 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v7
    _ = v8
    _ = v9
}
