package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint32; F1 int8; F2 int64 }
type GenS3 struct { F0 int64; F1 int32 }

func main() {
    var v1 uint = 255
    v2 := make(map[int]bool)
    var v3 int8 = -1
    var v4 uint64 = 246
    var v5 float32 = 0.1
    if ((("a" + (("hello" + "a") + ("a" + "hello"))) == "\n\t") && (!(v1 <= (v1 * 0)))) {
        if true {
            _ = len(v2)
        }
        for fi6 := 0; fi6 < 4; fi6++ {
            switch fi6 {
            case 32767:
                v3 = v3
            case 127:
                var v7 bool = (-1 > -9223372036854775808)
                _ = v7
            case -2147483648:
                v4 -= (v4 % (((^(v4))) | 1))
            default:
                v4 -= (v4 ^ (v4 + 16))
            }
            v8 := GenS1{F0: (!false), F1: -515}
            _ = fi6
            _ = v8
        }
        if ((455 <= 2147483648) && (((true || true) && ((v1 / ((477) | 1)) > (v1 - v1))) || true)) {
            _ = len(v2)
        }
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := GenS1{F0: ((fi9 / ((fi9) | 1)) >= fi9), F1: 991}
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: true, F1: -1}
        _ = v11
    }
    delete(v2, -2147483648)
    fmt.Println(v4, v3, v1)
    v12 := GenS1{F0: false, F1: -9223372036854775808}
    v4 = 645
    if false {
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 uint32 = 78
            v15 := []int{fi13, -9223372036854775808, fi13}
            var v16 float32 = v5
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 int64 = -1
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: (("" + "foo bar") > "hello"), F1: -354}
    var v20 int64 = -9223372036854775808
    v21 := []int{127, -9223372036854775808}
    for fi22 := 0; fi22 < 11; fi22++ {
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := GenS2{F0: 65535, F1: v3, F2: (v20 | (^(v20)))}
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    var v25 uint32 = 162
    v25 = v25
    v4 = 0
    var v26 int64 = -9223372036854775808
    if (v26 != v26) {
        for fi27 := 0; fi27 < 16; fi27++ {
            fmt.Println(v1, v4)
            _ = fi27
        }
        v28 := v2[-2147483648]
        v29 := GenS3{F0: 1, F1: -2147483648}
        _ = v28
        _ = v29
    } else {
        v30 := GenS3{F0: v26, F1: 2147483647}
        v4 -= v4
        var v31 uint16 = 65535
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
