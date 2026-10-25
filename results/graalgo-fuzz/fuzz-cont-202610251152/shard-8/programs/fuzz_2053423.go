package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint32; F1 int64; F2 int8 }
func genFn0(p0 uint16) uint16 {
    v6 := GenS1{F0: 672, F1: 65535, F2: "hello"}
    p0 = p0
    v7 := []int{-351, 2147483647, 32768}
    if (32767 < 784) {
        var v8 bool = false
        _ = v8
    }
    var v9 float64 = ((3.14159 / (0.1 * (1e10 + float64(1.5)))) + 3.14159)
    _ = p0
    _ = v6
    _ = v7
    _ = v9
    return p0
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{-1, 4294967295}
    v3 := make(map[int]string)
    v4 := make(map[uint32]int)
    v5 := []int{2147483647, 127, 127}
    var v10 float32 = float32(((((1e10 * 0.0) / 0.1) / float64((-2.25 / 1.0))) + 1.5))
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 int64 = -9223372036854775808
        v13 := GenS2{F0: 1}
        _ = fi11
        _ = v12
        _ = v13
    }
    if (1 != -342) {
        var v14 int = -9223372036854775808
        v15 := GenS3{F0: (100 == 0), F1: -32768}
        fmt.Println(v14, v10)
        _ = v14
        _ = v15
    }
    v10 = (v10 / 1.0)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v16 int32 = -683
    for fi17 := 0; fi17 < 19; fi17++ {
        v16 -= (-(v16))
        for fi18 := 0; fi18 < 16; fi18++ {
            v16 = 42
            switch fi18 {
            case -414:
                fmt.Println(v16)
            default:
                v19 := GenS3{F0: false, F1: 0}
                _ = v19
            }
            var v20 uint = 892
            _ = fi18
            _ = v20
        }
        fmt.Println(fi17, fi17)
        _ = fi17
    }
    if true {
        var v21 uint = 4294967295
        _ = v21
    }
    if (!true) {
        switch -276 {
        case 32768:
            for fi22 := 0; fi22 < 7; fi22++ {
                fmt.Println(fi22, v10, v10)
                _ = fi22
            }
        case 128:
            if ((false && false) && false) {
                var v23 uint32 = 0
                _ = v23
            }
        }
        if true {
            v16 = v16
        }
    } else {
        v10 = (v10 - 1.5)
        v24 := GenS4{F0: 0, F1: 708, F2: 1}
        for fi25 := 0; fi25 < 9; fi25++ {
            for fi26 := 0; fi26 < 12; fi26++ {
                fmt.Println(v10, fi26, v16)
                _ = fi26
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v27 := v3[fi25]
            _ = fi25
            _ = v27
        }
        _ = v24
    }
    v10 += v10
    for fi28 := 0; fi28 < 3; fi28++ {
        if len(v2) > 0 {
            v2[0] = (-(fi28))
        }
        fmt.Println(fi28, v10)
        _ = fi28
    }
    var v29 int8 = 113
    if (true || (2147483647 == 0)) {
        for fi30 := 0; fi30 < 13; fi30++ {
            switch (fi30 ^ (-(fi30))) {
            case 32768:
                switch fi30 {
                case 32767:
                    delete(v4, 4294967295)
                }
            }
            for fi31 := 0; fi31 < 4; fi31++ {
                var v32 uint64 = 18446744073709551615
                _ = fi31
                _ = v32
            }
            _ = fi30
        }
        if (416 != 401) {
            v10 = -2.25
        } else {
            switch "foo bar" {
            case "a":
                v29 -= (^(v29))
            case "foo bar":
                var v33 uint32 = 248
                _ = v33
            case "":
                var v34 bool = (0 < 32767)
                _ = v34
            }
            v10 += (0.0 + (v10 * -2.25))
            v10 = ((3.14159 + -2.25) * v10)
        }
    } else {
        v29 -= -128
    }
    v35 := GenS5{F0: 0, F1: 9223372036854775807, F2: (^(v29))}
    if (true && false) {
        v36 := v3[-751]
        _ = v36
    }
    var v37 int = 658
    _ = len(v4)
    v10 -= (float32((float64((-2.25 / ((1e10 * 3.14159) + ((0.1 * v10) * 0.0)))) * (0.1 + 1.5))) / (((v10 + v10) / v10) / v10))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v37
    _ = v4
    _ = v5
}
