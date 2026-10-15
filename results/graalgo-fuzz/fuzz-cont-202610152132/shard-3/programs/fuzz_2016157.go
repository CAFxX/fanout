package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 bool = true
    var v2 int64 = -502
    var v3 int8 = -128
    var v4 string = ("foo bar" + "a")
    if ((v3 | (v3 % ((1) | 1))) <= (v3 & (v3 << ((65535) & 63)))) {
        v5 := []int{}
        fmt.Println(v1, v1, v3)
        _ = v5
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint16 = 62
            v9 := GenS1{F0: ("" + (("a" + v4) + ((v4 + v4) + v4)))}
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v10 := []int{841, 32767, 308, 375}
    switch 2147483647 {
    case -169:
        for fi11 := 0; fi11 < 7; fi11++ {
            if v1 {
                switch v4 {
                case "foo bar":
                    v12 := GenS2{F0: v2, F1: -264}
                    _ = v12
                case "a":
                    var v13 int16 = 663
                    _ = v13
                default:
                    v14 := GenS1{F0: v4}
                    _ = v14
                }
                v15 := GenS3{F0: 0}
                _ = v15
            } else {
                var v16 bool = v1
                for fi17 := 0; fi17 < 10; fi17++ {
                    var v18 float64 = 1.5
                    _ = len(v10)
                    var v19 string = v4
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = v16
            }
            _ = fi11
        }
    }
    var v20 float64 = 1.5
    switch 4294967295 {
    case 4294967295:
        var v21 int64 = (v2 >> ((4294967295) & 63))
        _ = v21
    }
    var v22 int16 = -1
    v22 += v22
    if false {
        v1 = true
        var v23 bool = (!(4294967295 >= 65535))
        _ = v23
    } else {
        for fi24 := 0; fi24 < 17; fi24++ {
            v2 = (v2 & v2)
            v25 := GenS2{F0: 1, F1: 663}
            _ = fi24
            _ = v25
        }
        if (v22 <= v22) {
            _ = len(v10)
        }
    }
    var v26 uint16 = 255
    var v27 string = v4
    var v28 int64 = v2
    var v29 int8 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
