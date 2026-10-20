package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 string = ("" + "a")
    var v2 int16 = 298
    v3 := make(map[string]string)
    var v4 int16 = (-(v2))
    var v5 float64 = 1.5
    var v6 int32 = -990
    if (!false) {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 int = (fi7 % (((fi7 + fi7)) | 1))
            v9 := GenS1{F0: 54, F1: v1, F2: 0}
            _ = fi7
            _ = v8
            _ = v9
        }
        if true {
            var v10 bool = true
            var v11 bool = (v6 > v6)
            var v12 bool = (-1 >= -320)
            _ = v10
            _ = v11
            _ = v12
        }
    } else {
        v13 := GenS1{F0: 1, F1: v1, F2: 2}
        delete(v3, v1)
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        switch 128 {
        case 127:
            for fi15 := 0; fi15 < 11; fi15++ {
                v4 += v2
                v16 := GenS2{F0: 0}
                _ = fi15
                _ = v16
            }
        case 869:
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 uint32 = 4294967295
                _ = fi17
                _ = v18
            }
        case -528:
            var v19 uint32 = 2
            _ = v19
        }
        v20 := GenS1{F0: 42, F1: v1, F2: 255}
        v6 += 1
        _ = fi14
        _ = v20
    }
    v2 = 42
    v21 := []int{127}
    if true {
        v22 := GenS2{F0: v4}
        _ = len(v3)
        _ = v22
    }
    var v23 uint32 = 259
    v24 := GenS2{F0: v4}
    var v25 uint8 = 1
    v6 -= (v6 & (v6 >> ((2) & 63)))
    var v26 int64 = 1
    fmt.Println(v23)
    if true {
        var v27 uint64 = 18446744073709551615
        v28 := GenS3{F0: (v6 * -2147483648)}
        _ = v27
        _ = v28
    } else {
        v4 = v2
        fmt.Println(v6, v1, v23)
        if true {
            switch 2147483647 {
            case 32768:
                fmt.Println(v6)
            case -800:
                if ((v26 + v26) <= v26) {
                    var v29 uint64 = 4294967295
                    var v30 uint8 = v25
                    _ = v29
                    _ = v30
                }
            default:
                fmt.Println(v23)
            }
            if ((^(v6)) != -2147483648) {
                fmt.Println(v4, v23)
            }
            fmt.Println(v25, v25, v2)
        }
    }
    var v31 float64 = 1.5
    fmt.Println(v31, v31)
    if (((true || true) && (!false)) && (!true)) {
        var v32 string = "hello"
        _ = v32
    } else {
        var v33 float64 = v5
        v23 = v23
        var v34 float32 = -2.25
        _ = v33
        _ = v34
    }
    var v35 uint32 = 42
    v26 += v26
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
