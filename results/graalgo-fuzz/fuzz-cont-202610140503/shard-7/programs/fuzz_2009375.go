package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }

func main() {
    var v1 uint64 = 2
    var v2 bool = true
    v3 := make(map[string]string)
    var v4 int = 9223372036854775807
    var v5 float32 = ((0.0 * 1.5) * float32(1.5))
    switch ("a" + "\n\t") {
    case "foo bar":
        var v6 int32 = 1
        _ = v6
    default:
        var v7 int8 = 42
        _ = v7
    }
    v5 += ((-2.25 - 1e10) * v5)
    switch ("\n\t" + ("foo bar" + "")) {
    case "hello":
        for fi8 := 0; fi8 < 20; fi8++ {
            for fi9 := 0; fi9 < 9; fi9++ {
                switch fi8 {
                case -9223372036854775808:
                    var v10 float64 = (((1e10 + 1e10) * 0.0) * (-2.25 + 1e10))
                    _ = v10
                default:
                    var v11 float64 = 0.0
                    _ = v11
                }
                _ = fi9
            }
            v12 := []int{v4}
            v4 = v4
            _ = fi8
            _ = v12
        }
    case "a":
        var v13 uint8 = 0
        _ = v13
    }
    if (1 == 100) {
        delete(v3, ("hello" + "\n\t"))
    }
    v2 = ("hello" != "hello")
    v14 := []int{32767, v4}
    var v15 uint = 18446744073709551615
    if (0 <= 9223372036854775807) {
        fmt.Println(v2, v5)
        switch v4 {
        case 402:
            switch ("" + "\n\t") {
            case "":
                if (v2 || (803 <= 0)) {
                    v1 += v1
                }
            }
        case 42:
            var v16 int16 = 359
            _ = v16
        }
        if v2 {
            for fi17 := 0; fi17 < 19; fi17++ {
                switch (^(fi17)) {
                case 32767:
                    _ = len(v3)
                default:
                    var v18 uint32 = 42
                    _ = v18
                }
                if false {
                    if len(v14) > 0 {
                        v14[0] = v4
                    }
                    if len(v14) > 0 {
                        v14[0] = (-(fi17))
                    }
                } else {
                    for fi19 := 0; fi19 < 11; fi19++ {
                        var v20 float64 = float64(v5)
                        var v21 int = 728
                        fmt.Println(v5, v15, v1)
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                var v22 float32 = v5
                _ = fi17
                _ = v22
            }
            v15 += 42
            v15 = (v15 | v15)
        } else {
            fmt.Println(v2)
        }
    } else {
        v23 := GenS1{F0: (^(v15)), F1: 57}
        _ = v23
    }
    v24 := GenS1{F0: (v15 + (v15 << ((v15) & 63))), F1: 127}
    var v25 uint64 = (^(v1))
    var v26 int32 = -634
    if ((838 > 2) || ((42 != (-(v26))) || (128 == 1))) {
        v5 = 0.0
        v27 := GenS1{F0: (v15 * v15), F1: -89}
        _ = v27
    }
    v3[""] = "foo bar"
    var v28 int16 = 624
    v1 = 716
    if false {
        v25 = (v1 - (v25 | v1))
        switch (v4 >> ((4294967295) & 63)) {
        case -452:
            v15 = 1
        case 128:
            fmt.Println(v26)
        case 127:
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        default:
            fmt.Println(v2)
        }
    }
    v29 := GenS1{F0: (v15 ^ (v15 & 255)), F1: -128}
    fmt.Println(v1, v26)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
