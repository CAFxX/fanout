package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int64 }
func genFn0(p0 uint, p1 int32, p2 bool) int16 {
    p2 = (p2 || p2)
    v4 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}

func main() {
    v1 := []uint32{2147483647}
    v2 := []uint32{2147483647, 4294967295, 0, 42}
    v3 := make(map[string]string)
    v5 := GenS1{F0: "", F1: 1000, F2: 263}
    v2 = append(v2, 887)
    switch 783 {
    case -2147483648:
        switch -1 {
        case 1:
            var v6 int8 = 0
            _ = v6
        case -9223372036854775808:
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := GenS1{F0: "", F1: 1, F2: 42}
                _ = fi7
                _ = v8
            }
        case 2147483647:
            var v9 uint = 65535
            _ = v9
        default:
            for fi10 := 0; fi10 < 12; fi10++ {
                var v11 uint16 = 776
                for fi12 := 0; fi12 < 5; fi12++ {
                    var v13 int8 = -1
                    var v14 int32 = 42
                    v14 += 0
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = fi10
                _ = v11
            }
        }
    }
    var v15 int32 = 554
    for fi16 := 0; fi16 < 9; fi16++ {
        switch 32767 {
        case 42:
            var v17 uint16 = 42
            _ = v17
        case 0:
            v15 = v15
        default:
            v18 := v3["\n\t"]
            _ = v18
        }
        v19 := GenS1{F0: ("hello" + "a"), F1: genFn0(255, (v15 % ((v15) | 1)), (false && (0.0 < ((0.1 - 1e10) - ((-2.25 - float32((1.5 * 3.14159))) / ((3.14159 + -2.25) - 0.1)))))), F2: -580}
        if (!(2 == 42)) {
            var v20 int16 = genFn0(249, v15, false)
            for fi21 := 0; fi21 < 2; fi21++ {
                v15 += v15
                _ = fi21
            }
            v20 += 127
            _ = v20
        } else {
            var v22 int64 = -1
            if ((-128 >= 125) || (!true)) {
                v22 = v22
                var v23 int64 = (v22 | v22)
                fmt.Println(fi16, fi16, fi16)
                _ = v23
            }
            if (true && (v22 != (^(v22)))) {
                fmt.Println(fi16)
            }
            _ = v22
        }
        _ = fi16
        _ = v19
    }
    v15 = (v15 % (((v15 * v15)) | 1))
    v15 = 2147483647
    switch -180 {
    case 0:
        v15 += v15
    case -1:
        for fi24 := 0; fi24 < 20; fi24++ {
            switch (-(fi24)) {
            case -9223372036854775808:
                var v25 int32 = v15
                _ = v25
            }
            _ = len(v1)
            _ = fi24
        }
    case 127:
        v15 -= (v15 % ((v15) | 1))
    default:
        v15 -= (^(v15))
    }
    v15 += v15
    var v26 uint = 18446744073709551615
    for fi27 := 0; fi27 < 20; fi27++ {
        var v28 int16 = genFn0(474, -1, (705 == 0))
        v15 += 1
        var v29 int8 = -67
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v5
}
