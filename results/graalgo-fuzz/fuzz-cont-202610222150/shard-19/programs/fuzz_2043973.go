package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint64, p1 uint) string {
    v5 := []int{4294967295, -9223372036854775808, 32768, 4294967295}
    for fi6 := 0; fi6 < 15; fi6++ {
        v5 = append(v5, (fi6 + (fi6 % (((fi6 % (((fi6 >> ((4294967295) & 63))) | 1))) | 1))))
        if (!(true && (474 != 518))) {
            v7 := GenS1{F0: 255, F1: -2147483648, F2: 32767}
            switch "" {
            case "hello":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                var v8 int = (fi6 / ((-2147483648) | 1))
                _ = v8
            }
            for fi9 := 0; fi9 < 3; fi9++ {
                p0 = (p0 / ((1) | 1))
                fmt.Println(p0, p1, fi9)
                v10 := GenS1{F0: 42, F1: 1, F2: 703}
                _ = fi9
                _ = v10
            }
            _ = v7
        }
        _ = fi6
    }
    var v11 int16 = -1
    p0 -= (^(p0))
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    return ""
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    v4 := []string{"hello", "foo bar"}
    for fi12 := 0; fi12 < 6; fi12++ {
        switch (("" + "hello") + (("foo bar" + "foo bar") + "")) {
        case "a":
            for fi13 := 0; fi13 < 12; fi13++ {
                v14 := GenS1{F0: 2, F1: -1, F2: 1}
                _ = fi13
                _ = v14
            }
        default:
            for fi15 := 0; fi15 < 3; fi15++ {
                _ = len(v4)
                _ = len(v4)
                v16 := GenS1{F0: 42, F1: -2147483648, F2: 2}
                _ = fi15
                _ = v16
            }
        }
        var v17 float64 = 1e10
        _ = fi12
        _ = v17
    }
    if (1 != 0) {
        if true {
            _ = len(v4)
            delete(v3, "hello")
            var v18 int32 = 1
            _ = v18
        }
    }
    v19 := GenS1{F0: 127, F1: -537, F2: 255}
    v1 = v1
    for fi20 := 0; fi20 < 6; fi20++ {
        var v21 uint = 255
        _ = fi20
        _ = v21
    }
    _ = len(v4)
    if true {
        if (1 != 658) {
            v22 := GenS2{F0: 0, F1: 1, F2: 752}
            v1 = (^(v1))
            v23 := GenS2{F0: 100, F1: 172, F2: 2147483648}
            _ = v22
            _ = v23
        } else {
            var v24 uint32 = 1
            v24 = (v24 + v24)
            _ = v24
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println(v1)
    v1 += v1
    for fi25 := 0; fi25 < 14; fi25++ {
        switch ("" + "foo bar") {
        case "hello":
            var v26 int = (fi25 % ((fi25) | 1))
            _ = v26
        default:
            v1 = 2
        }
        if ((254 != 255) && false) {
            var v27 uint32 = 322
            var v28 uint8 = 127
            _ = v27
            _ = v28
        } else {
            for fi29 := 0; fi29 < 11; fi29++ {
                v4 = append(v4, "foo bar")
                v1 = (^(v1))
                var v30 uint16 = 502
                fmt.Println(v1, v30, v30)
                _ = fi29
                _ = v30
            }
            v1 -= 535
        }
        _ = fi25
    }
    var v31 uint = 706
    var v32 string = "a"
    v31 = 0
    v33 := GenS3{F0: 32767, F1: -1}
    var v34 float32 = 1e10
    v31 -= v31
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
