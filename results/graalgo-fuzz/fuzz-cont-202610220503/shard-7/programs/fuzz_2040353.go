package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 int8; F2 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 3.14159
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    switch 42 {
    case 32768:
        v2 += v2
    default:
        for fi5 := 0; fi5 < 12; fi5++ {
            var v6 float64 = 1.5
            for fi7 := 0; fi7 < 3; fi7++ {
                var v8 uint16 = 0
                _ = fi7
                _ = v8
            }
            _ = fi5
            _ = v6
        }
    }
    v9 := GenS1{F0: -9223372036854775808, F1: -128}
    v10 := GenS2{F0: 65535}
    fmt.Println(v2)
    v11 := []int{}
    var v12 int8 = 86
    var v13 bool = ((711 == 1000) && (1 == -2147483648))
    fmt.Println(v13, v12, v2)
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 uint16 = 65535
        _ = fi14
        _ = v15
    }
    v13 = v13
    if (42 <= -2147483648) {
        v13 = (9223372036854775807 < 9223372036854775807)
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 int64 = 42
            switch "hello" {
            case "":
                v13 = (false || (2147483648 <= 2147483648))
            case "\n\t":
                var v18 float32 = (v2 / ((v2 * (v2 * v2)) / v2))
                _ = v18
            case "foo bar":
                for fi19 := 0; fi19 < 19; fi19++ {
                    var v20 uint32 = 65535
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    v20 += 2147483648
                    _ = fi19
                    _ = v20
                }
            }
            _ = fi16
            _ = v17
        }
    } else {
        v21 := GenS1{F0: -9223372036854775808, F1: (-(v12))}
        v13 = (!(((-1 != -1) || v13) && (0.0 <= -2.25)))
        v2 = v2
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        for fi23 := 0; fi23 < 1; fi23++ {
            v13 = v13
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = fi23
        }
        _ = fi22
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        v25 := GenS2{F0: 2}
        var v26 int32 = 522
        var v27 uint = 255
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    v28 := GenS3{F0: (v13 && (2 > 0)), F1: -128, F2: -288}
    if (!true) {
        v12 += v12
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        var v30 int64 = -9223372036854775808
        v12 = v12
        fmt.Println(v13)
        _ = fi29
        _ = v30
    }
    v2 += -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v9
}
