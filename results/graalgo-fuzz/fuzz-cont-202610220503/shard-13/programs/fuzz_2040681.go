package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint64; F1 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = ((-2.25 - 1.5) * 1.5)
    v3 := make(map[uint32]string)
    var v4 float32 = float32(v2)
    var v5 int64 = -56
    switch ((("hello" + ("\n\t" + "\n\t")) + "a") + "foo bar") {
    case "hello":
        for fi6 := 0; fi6 < 1; fi6++ {
            switch fi6 {
            case 1:
                v4 = 0.1
            case 32768:
                var v7 float32 = v4
                _ = v7
            case -805:
                v2 = float64(v4)
            default:
                if false {
                    if (true || true) {
                        v5 = (^(v5))
                        fmt.Println(v4, v5, fi6)
                        if true {
                            if (-35 > int32(v5)) {
                                v2 = v2
                            }
                        }
                    }
                }
            }
            v5 -= (v5 << ((4294967295) & 63))
            var v8 bool = ((fi6 / ((fi6) | 1)) == fi6)
            _ = fi6
            _ = v8
        }
    case "a":
        var v9 float32 = float32(v2)
        _ = v9
    }
    var v10 float64 = v2
    v11 := GenS1{F0: false}
    var v12 float64 = v2
    switch "a" {
    case "foo bar":
        delete(v3, 2147483648)
    case "\n\t":
        switch "a" {
        case "hello":
            v13 := []int{}
            _ = v13
        default:
            var v14 int64 = v5
            _ = v14
        }
    default:
        v2 += (v2 / (v2 - v10))
    }
    if true {
        v15 := GenS1{F0: ("\n\t" >= "hello")}
        var v16 uint16 = 58
        v12 = -2.25
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 uint64 = 2
            var v19 int = (fi17 >> ((2) & 63))
            fmt.Println(v12)
            _ = fi17
            _ = v18
            _ = v19
        }
        switch "hello" {
        case "":
            for fi20 := 0; fi20 < 17; fi20++ {
                v21 := []int{(fi20 & (fi20 * (fi20 << ((487) & 63)))), fi20, fi20}
                switch "hello" {
                case "foo bar":
                    v5 = v5
                case "":
                    v22 := GenS1{F0: (!false)}
                    _ = v22
                default:
                    v2 = v2
                }
                v2 += float64(v4)
                _ = fi20
                _ = v21
            }
        case "\n\t":
            var v23 int = 739
            _ = v23
        case "foo bar":
            if true {
                var v24 float64 = (float64(0.0) * v2)
                var v25 float64 = 0.1
                _ = v24
                _ = v25
            }
        default:
            fmt.Println(v5, v5, v5)
        }
    }
    v12 -= 1e10
    if (!true) {
        fmt.Println(v4)
        var v26 float32 = float32(-2.25)
        var v27 float64 = ((v2 - float64(v26)) * v2)
        _ = v26
        _ = v27
    }
    v28 := GenS1{F0: false}
    v2 = 0.1
    delete(v3, 231)
    v29 := GenS2{F0: v5, F1: "\n\t"}
    v10 -= v10
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 int32 = -950
        for fi32 := 0; fi32 < 12; fi32++ {
            if true {
                v33 := GenS1{F0: false}
                v34 := GenS2{F0: (v5 - v5), F1: ""}
                _ = v33
                _ = v34
            } else {
                v31 = 1
                v10 = v2
            }
            fmt.Println(v31)
            v35 := GenS3{F0: 233, F1: 255}
            _ = fi32
            _ = v35
        }
        _ = fi30
        _ = v31
    }
    if (!false) {
        v36 := []int{9223372036854775807, -677}
        _ = v36
    }
    var v37 uint8 = 28
    v38 := GenS2{F0: (v5 + -9223372036854775808), F1: "a"}
    v4 -= 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
