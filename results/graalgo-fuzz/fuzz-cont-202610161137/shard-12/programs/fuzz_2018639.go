package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int64; F1 bool }

func main() {
    var v1 int = 506
    v2 := []int{(v1 % (((v1 >> ((614) & 63))) | 1)), 1}
    v3 := []string{}
    v4 := make(map[uint32]uint32)
    var v5 string = ("\n\t" + ("foo bar" + "\n\t"))
    var v6 float32 = -2.25
    if (!((-2.25 - float64(-2.25)) <= 1.5)) {
        v1 = v1
        v1 = (-(v1))
        if true {
            if len(v2) > 0 {
                v2[0] = (v1 ^ -206)
            }
            fmt.Println(v6)
        }
    }
    fmt.Println(v6)
    if true {
        var v7 int = v1
        var v8 float32 = 3.14159
        _ = v7
        _ = v8
    }
    fmt.Println(v1, v5, v1)
    switch (v1 * 164) {
    case -291:
        if (!(1 >= 960)) {
            var v9 uint = 154
            fmt.Println(v9, v6)
            _ = v9
        } else {
            v10 := GenS1{F0: 4294967295, F1: 42}
            var v11 uint = 42
            delete(v4, 2147483647)
            _ = v10
            _ = v11
        }
    case -1:
        fmt.Println(v5, v6, v1)
    }
    if (-88 == -27) {
        v6 += 0.1
        if false {
            if (2 > 42) {
                var v12 uint = 19
                _ = v12
            }
        } else {
            var v13 uint32 = 42
            v14 := GenS2{F0: v13, F1: 42, F2: 1}
            _ = v13
            _ = v14
        }
    } else {
        if false {
            v5 = ((v5 + v5) + ((v5 + ("foo bar" + v5)) + ("\n\t" + v5)))
            switch ("\n\t" + v5) {
            case "":
                if (372 != 42) {
                    switch 888 {
                    case 307:
                        v15 := GenS3{F0: 2, F1: 42}
                        _ = v15
                    case 538:
                        var v16 uint64 = 42
                        _ = v16
                    }
                }
            }
        }
        delete(v4, 720)
    }
    v17 := GenS3{F0: 1, F1: -147}
    v18 := GenS1{F0: 255, F1: 255}
    var v19 uint32 = 2147483648
    _ = len(v4)
    if ((v1 | (v1 >> ((829) & 63))) < v1) {
        var v20 float64 = 1.5
        var v21 string = (v5 + v5)
        v22 := GenS3{F0: 128, F1: 129}
        _ = v20
        _ = v21
        _ = v22
    }
    v23 := GenS4{F0: 651, F1: false}
    switch ("a" + v5) {
    case "\n\t":
        fmt.Println(v1)
    default:
        v5 = (v5 + (((v5 + v5) + v5) + v5))
    }
    if (32767 != 222) {
        switch (^(v1)) {
        case 128:
            v24 := GenS1{F0: 1, F1: 128}
            _ = v24
        case 254:
            var v25 uint = uint(v19)
            _ = v25
        }
        switch v1 {
        case 279:
            v5 = v5
        case 42:
            if (false || false) {
                var v26 bool = true
                var v27 float64 = (1e10 - (float64(((v6 * (float32(1e10) / 1.0)) * (v6 * (1e10 / (v6 + v6))))) - 0.1))
                v28 := v4[(v19 >> ((633) & 63))]
                _ = v26
                _ = v27
                _ = v28
            }
        }
    }
    var v29 int16 = 902
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
