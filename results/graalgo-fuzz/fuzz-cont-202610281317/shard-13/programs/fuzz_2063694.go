package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
func genFn0(p0 int8, p1 uint) int8 {
    v5 := []int{int(p1), 2147483647, 0}
    if false {
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS1{F0: -9223372036854775808, F1: 32767}
            if false {
                fmt.Println(fi6)
            }
            if (-2147483648 > -1) {
                p1 = p1
                if false {
                    v8 := GenS1{F0: 0, F1: 231}
                    v9 := GenS1{F0: 70, F1: 255}
                    var v10 string = ""
                    _ = v10
                    _ = v8
                    _ = v9
                }
                if ((!true) || false) {
                    p1 += (p1 % (((p1 * (p1 % ((2) | 1)))) | 1))
                }
            } else {
                fmt.Println(fi6, p1)
            }
            _ = fi6
            _ = v7
        }
    }
    v11 := GenS1{F0: 9223372036854775807, F1: 108}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    return (p0 | 0)
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[string]uint32)
    var v3 float32 = -2.25
    v4 := make(map[int]bool)
    var v12 uint8 = 63
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 string = ((("hello" + "a") + "foo bar") + "")
        v15 := []int{(fi13 | (fi13 % (((fi13 ^ (fi13 ^ fi13))) | 1))), (fi13 - (fi13 % ((fi13) | 1))), (fi13 | (^(fi13))), fi13}
        _ = len(v2)
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 9223372036854775807, F1: 65535}
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(fi18)
            if true {
                var v19 int32 = 453
                v1 = v1
                _ = v19
            }
            var v20 uint8 = (^(v12))
            _ = fi18
            _ = v20
        }
        _ = fi17
    }
    fmt.Println(v1, v12)
    v21 := GenS1{F0: 9223372036854775807, F1: 457}
    if false {
        if (3.14159 <= 1e10) {
            v3 = v3
        }
    } else {
        switch 881 {
        case 1:
            switch 929 {
            case 42:
                v3 = (0.1 - v3)
            }
        case 4294967295:
            for fi22 := 0; fi22 < 17; fi22++ {
                switch fi22 {
                case 127:
                    for fi23 := 0; fi23 < 5; fi23++ {
                        fmt.Println(v12, v12, fi23)
                        v1 = v1
                        _ = fi23
                    }
                case 959:
                    for fi24 := 0; fi24 < 15; fi24++ {
                        v12 += (^(v12))
                        _ = fi24
                    }
                }
                v25 := GenS1{F0: -9223372036854775808, F1: 395}
                v3 += 1.5
                _ = fi22
                _ = v25
            }
        case -803:
            v1 = v1
        default:
            v3 -= ((0.1 * -2.25) / 3.14159)
        }
        var v26 int8 = 0
        _ = v26
    }
    if (!((float64(v3) == 1.5) && (false && (9223372036854775807 == 0)))) {
        for fi27 := 0; fi27 < 14; fi27++ {
            v28 := []int{}
            var v29 int32 = 0
            switch int(v12) {
            case 504:
                v29 = v29
            default:
                fmt.Println(v29, v12)
            }
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
