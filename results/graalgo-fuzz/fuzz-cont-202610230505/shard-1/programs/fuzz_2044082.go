package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int; F1 uint; F2 uint16 }

func main() {
    var v1 bool = true
    v2 := []int{9223372036854775807}
    v3 := make(map[int]uint32)
    v1 = v1
    fmt.Println(v1)
    var v4 float64 = 0.1
    v5 := GenS1{F0: false, F1: 250, F2: 42}
    var v6 uint8 = 42
    for fi7 := 0; fi7 < 6; fi7++ {
        v6 = 61
        var v8 string = "foo bar"
        v9 := GenS1{F0: v1, F1: 2, F2: 487}
        _ = fi7
        _ = v8
        _ = v9
    }
    switch ((("foo bar" + "a") + "") + ("hello" + "foo bar")) {
    case "a":
        var v10 uint8 = v6
        _ = v10
    default:
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 bool = v1
            v13 := GenS1{F0: ((1e10 - (3.14159 - (1.5 + (1e10 - 0.0)))) <= (-2.25 - float32((float64(1e10) * (1e10 * 1.5))))), F1: 65535, F2: 682}
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    v4 = v4
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 string = ("hello" + ("" + "hello"))
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: 9223372036854775807, F1: 510, F2: 42}
    for fi17 := 0; fi17 < 5; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            fmt.Println(fi18)
            if true {
                v1 = (65535 >= 255)
                _ = len(v3)
            } else {
                v6 = 1
            }
            _ = fi18
        }
        for fi19 := 0; fi19 < 7; fi19++ {
            v1 = ((v6 == v6) || v1)
            switch fi19 {
            case -213:
                var v20 uint32 = 662
                _ = v20
            case 9223372036854775807:
                var v21 uint16 = 255
                var v22 float64 = v4
                v6 -= v6
                _ = v21
                _ = v22
            case 737:
                if (((0.1 + 3.14159) / 1e10) <= 1.5) {
                    if v1 {
                        var v23 int32 = 2147483647
                        fmt.Println(v4, fi17, fi19)
                        var v24 int = (fi19 ^ 0)
                        _ = v23
                        _ = v24
                    }
                    v6 += (v6 + 1)
                }
                _ = len(v2)
            default:
                var v25 uint32 = 2147483648
                _ = v25
            }
            var v26 uint = 99
            _ = fi19
            _ = v26
        }
        var v27 uint16 = 263
        _ = fi17
        _ = v27
    }
    if v1 {
        v28 := v3[32767]
        _ = v28
    }
    var v29 int = 322
    fmt.Println(v6, v29, v4)
    fmt.Println(v4, v6, v29)
    v3[v29] = 1
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
