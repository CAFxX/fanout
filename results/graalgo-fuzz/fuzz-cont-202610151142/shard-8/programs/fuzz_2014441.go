package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int }
type GenS2 struct { F0 int8; F1 int16 }
func genFn0(p0 bool, p1 uint64) uint8 {
    p1 -= (p1 & 4294967295)
    p1 = (p1 >> ((0) & 63))
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 float32 = float32((1.5 / 1e10))
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 int64 = 0
            v8 = 1.5
            switch (fi7 % ((fi9) | 1)) {
            case 78:
                var v11 int8 = 118
                _ = v11
            case -1:
                v8 = v8
            case 42:
                v8 = 3.14159
            }
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 bool, p1 uint8, p2 int) uint16 {
    v12 := []int{p2, (p2 / ((p2) | 1)), p2, p2}
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 uint16 = 897
        if ((2147483647 < 175) || (!p0)) {
            for fi15 := 0; fi15 < 14; fi15++ {
                if true {
                    var v16 string = "\n\t"
                    _ = v16
                }
                if (!true) {
                    v17 := GenS1{F0: 923, F1: v14, F2: 127}
                    fmt.Println(v14)
                    _ = v17
                }
                var v18 bool = (!false)
                _ = fi15
                _ = v18
            }
        }
        var v19 int32 = 0
        _ = fi13
        _ = v14
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 701
    var v3 uint16 = 32767
    v4 := []int{}
    v5 := []int{}
    var v6 int16 = 42
    var v20 uint64 = 42
    v2 = v2
    v3 = (v3 / (((v3 % (((v3 - v3)) | 1))) | 1))
    for fi21 := 0; fi21 < 14; fi21++ {
        v22 := GenS2{F0: 42, F1: (v6 << ((0) & 63))}
        _ = fi21
        _ = v22
    }
    var v23 float32 = (0.0 - 1e10)
    if false {
        v3 = v3
        fmt.Println(v3, v20)
    } else {
        var v24 uint32 = 2147483648
        v25 := GenS2{F0: -1, F1: v6}
        var v26 float64 = -2.25
        _ = v24
        _ = v25
        _ = v26
    }
    var v27 int64 = 1
    var v28 float32 = v23
    var v29 uint8 = 1
    switch ((("hello" + "\n\t") + (("\n\t" + "foo bar") + "\n\t")) + "a") {
    case "a":
        var v30 int8 = 78
        _ = v30
    }
    fmt.Println(v20, v27, v20)
    var v31 float32 = v28
    fmt.Println(v6)
    var v32 bool = (v29 > (^(v29)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
