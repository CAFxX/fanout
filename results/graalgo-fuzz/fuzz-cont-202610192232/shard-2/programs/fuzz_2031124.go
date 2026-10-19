package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int8; F2 uint32 }

func main() {
    var v1 uint32 = 891
    var v2 uint32 = 2147483648
    var v3 uint8 = 234
    var v4 int = -421
    v4 = 653
    v3 = (v3 >> ((42) & 63))
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := []int{(fi5 % ((v4) | 1)), (fi5 * (fi5 % (((^(fi5))) | 1))), 4294967295, (fi5 | fi5)}
        v7 := GenS1{F0: 0}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 255}
    for fi9 := 0; fi9 < 18; fi9++ {
        if false {
            if (42 >= 558) {
                v10 := GenS1{F0: 130}
                _ = v10
            } else {
                var v11 int64 = 1
                var v12 bool = true
                _ = v11
                _ = v12
            }
        } else {
            v2 = 2
            v4 += (v4 >> ((978) & 63))
        }
        v13 := []int{(^(fi9)), fi9, (fi9 ^ (^(v4))), v4}
        var v14 float32 = 0.0
        _ = fi9
        _ = v13
        _ = v14
    }
    fmt.Println(v1, v2, v1)
    var v15 uint16 = 0
    for fi16 := 0; fi16 < 5; fi16++ {
        v2 = (v2 >> ((65535) & 63))
        var v17 int64 = 0
        fmt.Println(fi16, v2, v4)
        _ = fi16
        _ = v17
    }
    v2 -= (^(v2))
    v4 -= (v4 * 0)
    v18 := []int{-1, 127}
    var v19 uint = 255
    switch (-(v4)) {
    case 127:
        v19 = (v19 >> (((^(v19))) & 63))
    case 673:
        _ = len(v18)
    case 128:
        if true {
            var v20 uint8 = v3
            var v21 int32 = 42
            fmt.Println(v3, v19)
            _ = v20
            _ = v21
        }
    default:
        v3 = uint8(v2)
    }
    v22 := GenS2{F0: -2147483648, F1: 127, F2: 121}
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    var v23 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v8
}
