package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 bool; F1 bool; F2 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = -197
    var v3 uint8 = 255
    var v4 float64 = 1e10
    var v5 bool = true
    var v6 bool = (0.0 == ((1e10 / 1.0) * 1e10))
    for fi7 := 0; fi7 < 1; fi7++ {
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            var v8 int8 = 1
            _ = v8
        case "":
            if (!(!false)) {
                v9 := GenS1{F0: -716}
                v4 = (v4 / v4)
                _ = v9
            }
        case "a":
            v1[-372] = ""
        default:
            var v10 int64 = 0
            _ = v10
        }
        v6 = (-657 <= -9223372036854775808)
        var v11 float64 = ((((-2.25 * v4) - float64(float32(((1.5 * ((1e10 / v4) + v4)) / v4)))) - 1e10) + v4)
        _ = fi7
        _ = v11
    }
    if (0 > 42) {
        var v12 bool = (v5 || (3.14159 < float32(-2.25)))
        var v13 int16 = 0
        v14 := []int{0, 1, 42}
        _ = v12
        _ = v13
        _ = v14
    }
    v4 = (float64(1e10) - float64(1e10))
    if (!v6) {
        if v6 {
            var v15 int16 = -1
            v16 := GenS1{F0: v15}
            fmt.Println(v2)
            _ = v15
            _ = v16
        }
        switch 127 {
        case 0:
            v17 := GenS1{F0: 1}
            _ = v17
        case 42:
            if (("hello" + "\n\t") < "\n\t") {
                v4 += -2.25
                v3 = 2
                var v18 int8 = 99
                _ = v18
            }
        case 648:
            v3 = 255
        }
    }
    var v19 uint = 1
    fmt.Println(v5)
    v6 = ((v19 << ((42) & 63)) > 4294967295)
    var v20 int32 = v2
    fmt.Println(v20, v19, v4)
    if true {
        v21 := GenS1{F0: 0}
        v22 := GenS1{F0: -1}
        var v23 uint32 = 4294967295
        _ = v21
        _ = v22
        _ = v23
    }
    if v6 {
        for fi24 := 0; fi24 < 10; fi24++ {
            _ = len(v1)
            _ = fi24
        }
        if v6 {
            var v25 uint16 = 1
            _ = v25
        }
    } else {
        v4 = v4
    }
    v26 := []int{-961, 788, 1, 9223372036854775807}
    switch "\n\t" {
    case "":
        fmt.Println(v4)
    default:
        v27 := GenS2{F0: -128, F1: 935}
        _ = v27
    }
    if v5 {
        v28 := GenS3{F0: (!false), F1: v6, F2: -32768}
        fmt.Println(v6, v20, v4)
        v19 = v19
        _ = v28
    } else {
        if v6 {
            if false {
                v5 = (!(v6 && ((((((v4 + -2.25) - (3.14159 + v4)) * ((((v4 * (float64((0.0 * (-2.25 * ((0.1 * 0.1) / (((-2.25 / -2.25) + 0.0) - ((0.1 - -2.25) + 0.1)))))) * 0.1)) / (0.0 - 3.14159)) - 0.1) + 0.1)) >= (0.0 - 1e10)) || ((1e10 - ((3.14159 / -2.25) * v4)) == v4)) || (((v20 / ((v2) | 1)) == int32(v19)) || (0.1 >= 1e10)))))
                for fi29 := 0; fi29 < 2; fi29++ {
                    delete(v1, 127)
                    var v30 bool = v6
                    _ = fi29
                    _ = v30
                }
            }
        }
        _ = len(v1)
        if ((-(v20)) <= v2) {
            v20 = v2
            v5 = (v6 || v5)
        }
    }
    v6 = (!(v5 && (0 > 42)))
    var v31 string = (("" + "") + ("foo bar" + ""))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
