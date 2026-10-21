package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }

func main() {
    var v1 float32 = -2.25
    var v2 string = ("" + "")
    v3 := make(map[int]int)
    var v4 uint16 = 2
    var v5 int64 = 42
    v6 := make(map[string]string)
    v7 := GenS1{F0: -623, F1: -1}
    fmt.Println(v1)
    var v8 uint32 = 668
    if (0.1 < -2.25) {
        var v9 uint32 = (v8 - 812)
        v3[2147483647] = 42
        if false {
            fmt.Println(v4, v1, v4)
            v10 := []int{2147483647}
            v2 = (v2 + (v2 + ((v2 + (v2 + v2)) + "foo bar")))
            _ = v10
        } else {
            switch "foo bar" {
            case "a":
                v8 = v8
            default:
                v1 = v1
            }
            v8 += (v8 & (v9 ^ (^(v9))))
        }
        _ = v9
    }
    switch "foo bar" {
    case "\n\t":
        if true {
            for fi11 := 0; fi11 < 6; fi11++ {
                if (603 == 2) {
                    v2 = (("foo bar" + ((("" + v2) + v2) + ((v2 + v2) + v2))) + v2)
                }
                _ = fi11
            }
        } else {
            var v12 string = "a"
            v8 += 1
            _ = v12
        }
    case "hello":
        v5 = (-(v5))
    case "a":
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := GenS1{F0: -894, F1: 42}
            _ = fi13
            _ = v14
        }
    default:
        v8 = 2147483648
    }
    var v15 int = 970
    if (((((-2.25 - float64(0.1)) / 1e10) + -2.25) - 0.1) >= 0.1) {
        fmt.Println(v8, v4)
    } else {
        for fi16 := 0; fi16 < 2; fi16++ {
            v2 = v2
            for fi17 := 0; fi17 < 5; fi17++ {
                v18 := GenS1{F0: -1, F1: 0}
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
        if false {
            var v19 int16 = 1000
            v8 += (v8 ^ (v8 << ((255) & 63)))
            v20 := []int{(^(v15))}
            _ = v19
            _ = v20
        }
        if false {
            var v21 string = v2
            if (2147483647 < 822) {
                v22 := GenS1{F0: 2147483647, F1: 62}
                _ = v22
            }
            _ = v21
        }
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        v2 = "a"
        _ = fi23
    }
    v24 := []int{(v15 & -1), 192, 128}
    v4 -= (v4 - v4)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
