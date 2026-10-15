package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }

func main() {
    v1 := []string{"foo bar", "\n\t", "hello"}
    var v2 int32 = 294
    v3 := make(map[int]uint32)
    v4 := GenS1{F0: 128, F1: 1}
    var v5 string = "foo bar"
    for fi6 := 0; fi6 < 17; fi6++ {
        v5 = "\n\t"
        var v7 int8 = -113
        _ = fi6
        _ = v7
    }
    v5 = v5
    switch -157 {
    case 42:
        fmt.Println(v5)
    case -2147483648:
        var v8 int32 = v2
        _ = v8
    case 1:
        var v9 uint16 = 2
        _ = v9
    default:
        if false {
            var v10 float64 = float64(3.14159)
            for fi11 := 0; fi11 < 10; fi11++ {
                var v12 uint8 = 127
                var v13 int32 = v2
                v13 -= -2147483648
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
        } else {
            _ = len(v1)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    v14 := v3[-2147483648]
    fmt.Println(v5, v14)
    fmt.Println(v2, v5, v14)
    if (!(9223372036854775807 >= 42)) {
        if (true || true) {
            switch 0 {
            case 2147483647:
                if false {
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                    for fi15 := 0; fi15 < 10; fi15++ {
                        v16 := GenS1{F0: 127, F1: -32768}
                        _ = fi15
                        _ = v16
                    }
                }
            case -1:
                v3[-2147483648] = v14
            case 9223372036854775807:
                var v17 float32 = 0.1
                _ = v17
            }
            if (4294967295 >= 4294967295) {
                _ = len(v1)
            } else {
                for fi18 := 0; fi18 < 13; fi18++ {
                    v19 := GenS1{F0: 2, F1: 1000}
                    for fi20 := 0; fi20 < 1; fi20++ {
                        v14 -= v14
                        _ = fi20
                    }
                    _ = fi18
                    _ = v19
                }
            }
            if (v5 <= (v5 + "foo bar")) {
                var v21 int32 = v2
                v3[9223372036854775807] = 2147483647
                if false {
                    delete(v3, 32767)
                    var v22 int16 = 42
                    _ = v22
                }
                _ = v21
            } else {
                v5 = v5
            }
        } else {
            v23 := GenS1{F0: 1, F1: -76}
            _ = v23
        }
        var v24 float64 = float64((((1.5 + 1.5) * 1e10) - 1.5))
        switch -9223372036854775808 {
        case -1:
            var v25 uint8 = 83
            _ = v25
        case 539:
            var v26 int16 = -32768
            _ = v26
        case 0:
            var v27 int = -485
            _ = v27
        }
        _ = v24
    } else {
        var v28 uint8 = 127
        var v29 int32 = (v2 ^ (v2 - v2))
        _ = v28
        _ = v29
    }
    var v30 float32 = (((3.14159 + 1.5) + 3.14159) - ((1e10 / ((1e10 - 0.0) - 0.0)) * 0.1))
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    if len(v1) > 0 {
        v1[0] = (((v5 + (v5 + (v5 + ""))) + ((v5 + v5) + (v5 + "hello"))) + v5)
    }
    v2 = (v2 * v2)
    var v31 int32 = (v2 ^ (v2 - -265))
    v1 = append(v1, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
