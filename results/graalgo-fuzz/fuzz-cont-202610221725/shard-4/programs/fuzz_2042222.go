package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint; F2 uint64 }
func genFn0(p0 uint8, p1 int) int32 {
    v6 := []int{0, p1, p1}
    var v7 float64 = 1.5
    var v8 uint = 899
    if false {
        if (!false) {
            if false {
                v9 := GenS1{F0: p0}
                p1 -= p1
                _ = v9
            }
            var v10 string = "\n\t"
            fmt.Println(v8)
            _ = v10
        } else {
            var v11 string = "hello"
            _ = v11
        }
        if ((1 > -830) || (942 <= 0)) {
            var v12 string = ""
            v8 = v8
            var v13 int32 = 1
            _ = v12
            _ = v13
        } else {
            var v14 int32 = 1
            switch "hello" {
            case "":
                var v15 uint32 = 2147483648
                _ = v15
            case "a":
                v6 = append(v6, p1)
            }
            var v16 int16 = -592
            _ = v14
            _ = v16
        }
        var v17 uint8 = p0
        _ = v17
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v6 = append(v6, -9223372036854775808)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int = 32767
    var v2 uint8 = 0
    v3 := []uint32{}
    var v4 uint = 2
    var v5 bool = false
    v4 = (v4 / (((v4 - v4)) | 1))
    if len(v3) > 0 {
        v3[0] = 1
    }
    var v18 uint16 = 552
    var v19 uint32 = 962
    v19 += (^(v19))
    v20 := GenS2{F0: v5, F1: 518, F2: 987}
    if v5 {
        var v21 bool = (!((v19 * (v19 >> ((v4) & 63))) >= v19))
        _ = v21
    } else {
        switch (-(v1)) {
        case 96:
            v1 = (v1 % (((v1 ^ (v1 ^ v1))) | 1))
        case 32767:
            var v22 int = v1
            _ = v22
        default:
            v23 := GenS1{F0: (v2 + (v2 >> ((v4) & 63)))}
            _ = v23
        }
    }
    var v24 int32 = 1
    v5 = (int(v4) < v1)
    v2 = 2
    if ("a" < ("a" + "hello")) {
        var v25 float32 = 3.14159
        _ = v25
    } else {
        v3 = append(v3, (^(v19)))
    }
    v19 = (^(v19))
    var v26 bool = false
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
