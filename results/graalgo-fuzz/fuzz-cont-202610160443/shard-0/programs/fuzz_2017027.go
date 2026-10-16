package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint8; F1 string; F2 uint64 }
func genFn0(p0 string, p1 int64) int32 {
    p1 = (p1 % (((p1 ^ (p1 & -1))) | 1))
    switch (("a" + "foo bar") + (p0 + "a")) {
    case "":
        v4 := GenS1{F0: 2}
        _ = v4
    default:
        var v5 float64 = float64(0.1)
        _ = v5
    }
    v6 := []int{128, 2147483647, 661, 42}
    switch p0 {
    case "":
        p0 = p0
    default:
        p1 = (^(p1))
    }
    v7 := GenS2{F0: false, F1: 1000, F2: 144}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    v1 := make(map[int]int)
    var v2 float32 = 1.5
    var v3 uint32 = 65535
    var v8 float64 = float64(((0.1 * ((0.0 + v2) - 1.5)) + 0.1))
    switch -875 {
    case 32767:
        fmt.Println(v8, v3, v3)
    case 9223372036854775807:
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 uint = 42
            var v11 int64 = -9223372036854775808
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        v8 += v8
    }
    var v12 int8 = -47
    fmt.Println(v8, v2)
    v8 = v8
    switch -829 {
    case 127:
        for fi13 := 0; fi13 < 19; fi13++ {
            v14 := []int{fi13}
            var v15 int64 = 9223372036854775807
            _ = fi13
            _ = v14
            _ = v15
        }
    default:
        switch -9223372036854775808 {
        case 1:
            for fi16 := 0; fi16 < 5; fi16++ {
                v1[fi16] = (fi16 | 32768)
                _ = fi16
            }
        case 32768:
            if false {
                v3 -= v3
            }
        }
    }
    var v17 uint32 = (v3 / ((v3) | 1))
    v18 := []int{}
    if false {
        if (false && false) {
            var v19 uint8 = 162
            _ = v19
        }
        v20 := GenS2{F0: (128 < 127), F1: -32768, F2: 128}
        v21 := GenS2{F0: (!((float64(0.1) - (1.5 / (1e10 + v8))) < float64(float32((3.14159 - (v8 - v8)))))), F1: 1000, F2: 96}
        _ = v20
        _ = v21
    } else {
        if true {
            v22 := GenS2{F0: (v17 == v17), F1: 32767, F2: 127}
            _ = v22
        }
        var v23 uint64 = 1
        v8 -= v8
        _ = v23
    }
    if (!(!(-601 > -2147483648))) {
        var v24 uint8 = 255
        v24 -= (v24 & v24)
        if true {
            var v25 int16 = 0
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v25, v8)
            default:
                v2 -= (float32(v8) * float32(1.5))
            }
            v12 = (v12 << ((371) & 63))
            _ = v25
        }
        _ = v24
    } else {
        v26 := GenS3{F0: 1, F1: 1}
        switch 9223372036854775807 {
        case 9223372036854775807:
            v3 += (v3 * v3)
        }
        _ = v26
    }
    var v27 uint32 = v3
    fmt.Println(v12)
    var v28 uint32 = (^(v27))
    switch ("\n\t" + ("a" + "\n\t")) {
    case "hello":
        if true {
            var v29 int64 = 9223372036854775807
            _ = v29
        } else {
            var v30 float32 = v2
            _ = v30
        }
    default:
        if false {
            var v31 uint32 = v3
            _ = v31
        } else {
            _ = len(v18)
        }
    }
    v8 = 0.1
    if true {
        switch 42 {
        case -2147483648:
            for fi32 := 0; fi32 < 13; fi32++ {
                if (!false) {
                    var v33 int32 = 42
                    _ = v33
                }
                switch ("a" + ("foo bar" + ("\n\t" + "hello"))) {
                case "foo bar":
                    var v34 uint32 = (v28 / (((^(v28))) | 1))
                    _ = v34
                case "":
                    _ = len(v18)
                case "a":
                    fmt.Println(v28)
                default:
                    for fi35 := 0; fi35 < 4; fi35++ {
                        v17 += (^(v17))
                        fmt.Println(v28, v27, v28)
                        switch (fi35 - int(v17)) {
                        case 32767:
                            var v36 int64 = 1
                            _ = v36
                        }
                        var v37 int64 = -688
                        _ = fi35
                        _ = v37
                    }
                }
                for fi38 := 0; fi38 < 1; fi38++ {
                    v8 = (((v8 / (((v8 / v8) - 3.14159) - float64(0.0))) * v8) - 0.1)
                    _ = fi38
                }
                _ = fi32
            }
        case 9223372036854775807:
            v39 := GenS4{F0: 28, F1: "\n\t", F2: 2}
            _ = v39
        }
        var v40 int64 = 0
        for fi41 := 0; fi41 < 12; fi41++ {
            fmt.Println(v40, v28)
            _ = fi41
        }
        _ = v40
    } else {
        var v42 int16 = -1
        _ = v42
    }
    v43 := GenS1{F0: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v43
    _ = v8
}
