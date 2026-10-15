package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int16; F2 bool }

func main() {
    var v1 uint8 = 255
    var v2 int64 = -9223372036854775808
    var v3 string = ""
    var v4 bool = false
    var v5 float64 = (float64((0.0 + 0.0)) - (1e10 / (3.14159 + (0.1 - (float64((0.1 * (0.1 * (0.0 / 1e10)))) + -2.25)))))
    var v6 int16 = 32767
    v7 := GenS1{F0: -74}
    var v8 int8 = 0
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 uint32 = 4294967295
        switch 42 {
        case 2147483647:
            v11 := []int{-1, 42, fi9, int(v8)}
            _ = v11
        case 629:
            fmt.Println(v10)
        }
        if true {
            var v12 uint16 = 0
            _ = v12
        } else {
            if false {
                for fi13 := 0; fi13 < 16; fi13++ {
                    v2 = v2
                    v14 := []int{42, fi9}
                    _ = fi13
                    _ = v14
                }
            } else {
                var v15 uint64 = 602
                if true {
                    var v16 int64 = int64(fi9)
                    _ = v16
                } else {
                    v2 -= -1
                    v5 = v5
                    v17 := []int{fi9, -253, (fi9 % ((991) | 1)), fi9}
                    _ = v17
                }
                _ = v15
            }
            v18 := GenS2{F0: (v1 / ((v1) | 1)), F1: (v6 - (v6 ^ int16(v8))), F2: false}
            if v4 {
                var v19 int8 = v8
                var v20 uint32 = v10
                _ = v19
                _ = v20
            } else {
                v21 := GenS2{F0: v1, F1: (v6 - v6), F2: v4}
                _ = v21
            }
            _ = v18
        }
        _ = fi9
        _ = v10
    }
    v2 += (v2 % ((0) | 1))
    var v22 int = -2147483648
    v4 = true
    for fi23 := 0; fi23 < 13; fi23++ {
        v22 += (^(fi23))
        _ = fi23
    }
    var v24 float32 = 1.5
    v3 = "a"
    v25 := []int{}
    v25 = append(v25, (v22 * v22))
    if len(v25) > 0 {
        v25[0] = -1
    }
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    if (!v4) {
        v1 = 255
        for fi26 := 0; fi26 < 11; fi26++ {
            v22 = v22
            fmt.Println(v2)
            _ = fi26
        }
        var v27 int16 = (v6 % (((^(v6))) | 1))
        _ = v27
    }
    var v28 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
