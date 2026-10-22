package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint64; F1 int; F2 int64 }
func genFn0(p0 uint8) int8 {
    p0 = (p0 ^ (p0 << ((458) & 63)))
    p0 = (p0 & p0)
    v6 := []int{-53, -1, 237}
    for fi7 := 0; fi7 < 11; fi7++ {
        switch 0 {
        case 2147483647:
            var v8 uint64 = 4294967295
            _ = v8
        default:
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 int8 = -15
                var v11 uint32 = 1
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 float32 = -2.25
            _ = fi12
            _ = v13
        }
        if len(v6) > 0 {
            v6[0] = 128
        }
        _ = fi7
    }
    _ = p0
    _ = v6
    return 100
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 64
    v3 := []int{}
    var v4 float32 = 1.5
    v5 := []uint32{4294967295, 5, 0}
    v2 -= (v2 % (((^(v2))) | 1))
    if (false && true) {
        for fi14 := 0; fi14 < 8; fi14++ {
            fmt.Println(fi14, fi14, v2)
            var v15 uint32 = 581
            if (1.5 != v4) {
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 uint64 = 104
                    v2 = (^(v2))
                    _ = fi16
                    _ = v17
                }
            } else {
                var v18 uint8 = v2
                fmt.Println(v2, v4)
                for fi19 := 0; fi19 < 11; fi19++ {
                    if (v4 == 1.5) {
                        v2 = uint8(v15)
                        var v20 uint = 65535
                        var v21 uint16 = 543
                        _ = v20
                        _ = v21
                    } else {
                        v4 = ((0.0 / v4) / (1.5 + 1.5))
                        fmt.Println(v4, v2)
                    }
                    v22 := GenS1{F0: "\n\t", F1: true}
                    v15 = v15
                    _ = fi19
                    _ = v22
                }
                _ = v18
            }
            _ = fi14
            _ = v15
        }
        var v23 int64 = 0
        v4 -= v4
        _ = v23
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        _ = len(v5)
        var v25 float64 = (0.0 / 3.14159)
        for fi26 := 0; fi26 < 18; fi26++ {
            if ("foo bar" == "hello") {
                var v27 uint32 = 2147483647
                v27 = v27
                if len(v5) > 0 {
                    v5[0] = 42
                }
                _ = v27
            }
            fmt.Println(v25, v4)
            if (1 <= 237) {
                v25 += v25
                var v28 int8 = 45
                _ = v28
            }
            _ = fi26
        }
        _ = fi24
        _ = v25
    }
    v29 := GenS1{F0: "\n\t", F1: (!(!false))}
    for fi30 := 0; fi30 < 5; fi30++ {
        v2 += (v2 / ((248) | 1))
        _ = fi30
    }
    var v31 int16 = 32767
    if true {
        if false {
            v32 := GenS2{F0: 489, F1: 0, F2: 1}
            _ = v32
        } else {
            v31 = (v31 * -224)
            var v33 string = "a"
            var v34 bool = (0 < 128)
            _ = v33
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        v36 := GenS2{F0: 629, F1: (fi35 << ((0) & 63)), F2: -150}
        _ = fi35
        _ = v36
    }
    fmt.Println(v31, v2)
    var v37 int32 = -593
    for fi38 := 0; fi38 < 2; fi38++ {
        v39 := GenS1{F0: "hello", F1: false}
        var v40 string = "hello"
        fmt.Println(fi38, v4, v31)
        _ = fi38
        _ = v39
        _ = v40
    }
    v41 := GenS1{F0: "foo bar", F1: (true || true)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v37
    _ = v4
    _ = v41
    _ = v5
}
