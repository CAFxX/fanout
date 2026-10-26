package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int8 }
func genFn0(p0 int64) uint8 {
    if false {
        var v5 int64 = p0
        fmt.Println(v5, v5)
        p0 -= p0
        _ = v5
    }
    fmt.Println(p0)
    if (!(1 <= 42)) {
        for fi6 := 0; fi6 < 7; fi6++ {
            for fi7 := 0; fi7 < 5; fi7++ {
                switch (fi7 - (^(fi6))) {
                case -415:
                    var v8 float64 = float64((((1e10 - (0.1 + (((1.5 / ((0.0 - 0.1) - (-2.25 - (-2.25 * 3.14159)))) - (float32((3.14159 * 0.1)) - (1.5 / 1.0))) + (0.1 * 0.1)))) + 3.14159) / 1e10))
                    _ = v8
                case 2147483647:
                    v9 := []int{32768, (fi7 >> ((65535) & 63)), -684}
                    _ = v9
                default:
                    var v10 uint16 = 255
                    _ = v10
                }
                switch fi6 {
                case 931:
                    v11 := []int{-1}
                    _ = v11
                case -2147483648:
                    var v12 uint16 = 1
                    _ = v12
                }
                switch "" {
                case "a":
                    p0 -= p0
                default:
                    fmt.Println(fi6)
                }
                _ = fi7
            }
            for fi13 := 0; fi13 < 10; fi13++ {
                p0 += (p0 | (p0 & 42))
                var v14 int64 = -421
                fmt.Println(p0, fi13, p0)
                _ = fi13
                _ = v14
            }
            _ = fi6
        }
        fmt.Println(p0)
        v15 := []int{}
        _ = v15
    } else {
        for fi16 := 0; fi16 < 9; fi16++ {
            for fi17 := 0; fi17 < 8; fi17++ {
                switch (("" + "hello") + "foo bar") {
                case "\n\t":
                    v18 := GenS1{F0: 98, F1: 0, F2: 127}
                    _ = v18
                }
                if (p0 <= 9223372036854775807) {
                    v19 := []int{1}
                    fmt.Println(p0, fi16)
                    _ = v19
                } else {
                    v20 := []int{fi16, fi16, (fi17 + (-(fi16))), (fi17 ^ fi17)}
                    _ = v20
                }
                p0 += -1
                _ = fi17
            }
            for fi21 := 0; fi21 < 14; fi21++ {
                p0 += (p0 + (p0 | 42))
                v22 := GenS1{F0: (p0 * p0), F1: 127, F2: 0}
                _ = fi21
                _ = v22
            }
            _ = fi16
        }
        v23 := []int{-2147483648, -52, 32767, 513}
        p0 = (p0 & p0)
        _ = v23
    }
    _ = p0
    return 0
}

func main() {
    var v1 uint32 = 805
    v2 := make(map[uint32]bool)
    var v3 int = -484
    var v4 uint16 = 42
    if (-32768 > 1000) {
        v3 = 2147483647
        v4 = (v4 ^ (v4 % (((v4 >> ((42) & 63))) | 1)))
        for fi24 := 0; fi24 < 13; fi24++ {
            var v25 uint32 = (v1 ^ v1)
            v3 = v3
            v3 = fi24
            _ = fi24
            _ = v25
        }
    }
    v4 += v4
    fmt.Println(v4, v4, v3)
    if false {
        v26 := GenS1{F0: 42, F1: 127, F2: 1}
        if false {
            v27 := GenS1{F0: 1, F1: 127, F2: int8(v1)}
            if false {
                v28 := GenS1{F0: 0, F1: 255, F2: 1}
                _ = v28
            }
            _ = v27
        }
        v1 = (v1 << ((2) & 63))
        _ = v26
    } else {
        v4 = v4
        v4 = (v4 | 1)
    }
    fmt.Println(v1, v1)
    v29 := []int{-2147483648, v3, (v3 << ((65535) & 63))}
    if (3.14159 <= -2.25) {
        var v30 float32 = 0.0
        _ = v30
    } else {
        if true {
            var v31 float32 = float32(1.5)
            for fi32 := 0; fi32 < 13; fi32++ {
                var v33 int16 = 1
                var v34 bool = true
                v35 := GenS1{F0: -9223372036854775808, F1: 2, F2: 1}
                _ = fi32
                _ = v33
                _ = v34
                _ = v35
            }
            var v36 uint32 = 42
            _ = v31
            _ = v36
        } else {
            var v37 int64 = -216
            var v38 uint32 = v1
            if false {
                fmt.Println(v3)
                var v39 float64 = 3.14159
                _ = v39
            } else {
                var v40 bool = ("foo bar" > "foo bar")
                var v41 int8 = 127
                var v42 float32 = float32(0.0)
                _ = v40
                _ = v41
                _ = v42
            }
            _ = v37
            _ = v38
        }
    }
    switch (((("a" + "a") + ("a" + "hello")) + ((("foo bar" + ("foo bar" + "\n\t")) + "foo bar") + "\n\t")) + "hello") {
    case "foo bar":
        for fi43 := 0; fi43 < 12; fi43++ {
            var v44 int32 = 586
            _ = fi43
            _ = v44
        }
    case "":
        fmt.Println(v3, v3, v1)
    }
    v1 = (v1 * v1)
    var v45 float64 = ((3.14159 / -2.25) * (1.5 - -2.25))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v45
}
