package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
func genFn0(p0 uint, p1 uint64, p2 int8) string {
    p1 = (p1 << ((p0) & 63))
    p0 = (p0 & (^(p0)))
    _ = p0
    _ = p1
    _ = p2
    return "hello"
}
func genFn1() uint {
    if false {
        var v4 bool = (!false)
        _ = v4
    }
    var v5 int8 = 100
    if (2 < 0) {
        v5 = v5
        var v6 int64 = 0
        v7 := []int{}
        _ = v6
        _ = v7
    }
    var v8 int32 = -823
    var v9 string = ""
    _ = v5
    _ = v8
    _ = v9
    return 65535
}

func main() {
    var v1 float64 = (0.0 / 1e10)
    var v2 int16 = -439
    v3 := make(map[string]string)
    v10 := []int{32767, 615}
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(v2, fi11)
        _ = fi11
    }
    switch 862 {
    case -604:
        var v12 int8 = -128
        _ = v12
    case 42:
        if (!true) {
            fmt.Println(v2, v1)
            v1 += (((v1 - float64(-2.25)) / (((1e10 + 1.5) + ((-2.25 + (1.5 + v1)) - v1)) / (-2.25 + v1))) + 1e10)
            var v13 int8 = 78
            _ = v13
        } else {
            var v14 float32 = ((0.0 + (((float32(v1) - 0.1) + 0.1) / 1.0)) - (0.0 + 1e10))
            v15 := GenS1{F0: 0, F1: -9223372036854775808}
            switch 127 {
            case 9223372036854775807:
                for fi16 := 0; fi16 < 2; fi16++ {
                    v17 := GenS1{F0: 4294967295, F1: -971}
                    _ = fi16
                    _ = v17
                }
            case 4294967295:
                var v18 uint32 = 2147483648
                _ = v18
            case -458:
                var v19 uint64 = 0
                _ = v19
            }
            _ = v14
            _ = v15
        }
    }
    var v20 int = 404
    v1 -= 1e10
    switch 128 {
    case 4294967295:
        switch "" {
        case "":
            fmt.Println(v2, v1, v1)
        default:
            v2 -= v2
        }
    }
    fmt.Println(v20)
    switch (("" + ((("" + "foo bar") + "\n\t") + "\n\t")) + "foo bar") {
    case "\n\t":
        v20 = v20
    case "":
        v21 := GenS1{F0: 4294967295, F1: 9223372036854775807}
        _ = v21
    }
    fmt.Println(v1)
    v2 -= (v2 >> ((255) & 63))
    switch "hello" {
    case "hello":
        var v22 uint = 85
        _ = v22
    case "":
        switch "a" {
        case "foo bar":
            v2 -= -32768
        }
    default:
        v1 = float64(float32(-2.25))
    }
    fmt.Println(v1, v20, v2)
    v2 = (v2 ^ v2)
    var v23 int = (v20 - 1)
    v23 = (v20 + -1)
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 bool = (1 < 2147483648)
        _ = fi24
        _ = v25
    }
    var v26 float32 = (1.5 / 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v23
    _ = v26
    _ = v3
}
