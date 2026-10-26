package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint16; F2 int64 }

func main() {
    var v1 uint32 = 654
    var v2 uint64 = 114
    var v3 uint8 = 1
    var v4 int16 = -32768
    var v5 float64 = 0.1
    v6 := GenS1{F0: (^(v1)), F1: v1}
    if ((!true) || ((v3 % ((v3) | 1)) < (v3 / ((uint8(v1)) | 1)))) {
        v1 -= v1
        v5 += ((v5 - float64((1e10 + 0.1))) * v5)
        v5 = float64(1e10)
    } else {
        switch ("foo bar" + ("" + ("foo bar" + "a"))) {
        case "hello":
            switch 9223372036854775807 {
            case 127:
                var v7 uint = 2
                _ = v7
            case 32768:
                var v8 uint64 = uint64(v1)
                _ = v8
            default:
                v9 := []int{-137, 630, 32767}
                _ = v9
            }
        default:
            v10 := GenS2{F0: int16(v2), F1: uint16(v3), F2: -9223372036854775808}
            _ = v10
        }
        for fi11 := 0; fi11 < 12; fi11++ {
            if ((^(v2)) > (v2 >> ((147) & 63))) {
                var v12 string = ("" + "foo bar")
                var v13 float32 = (-2.25 / ((0.1 + -2.25) + 0.0))
                _ = v12
                _ = v13
            } else {
                var v14 uint8 = v3
                var v15 int16 = (v4 | v4)
                if (65535 > 4294967295) {
                    v16 := GenS2{F0: 1, F1: 32767, F2: int64(v14)}
                    var v17 int32 = 119
                    _ = v16
                    _ = v17
                }
                _ = v14
                _ = v15
            }
            for fi18 := 0; fi18 < 4; fi18++ {
                var v19 uint = 255
                switch "a" {
                case "a":
                    if ((!false) && false) {
                        if (false || (false || (false && (((!false) && false) || (289 != v19))))) {
                            fmt.Println(v5)
                            fmt.Println(v5, v2)
                            v20 := []int{fi11, 42, 128}
                            _ = v20
                        }
                    }
                    v4 = v4
                    v4 -= (v4 & v4)
                case "\n\t":
                    switch fi18 {
                    case 127:
                        fmt.Println(v2, fi11, fi11)
                    case 679:
                        var v21 int32 = 1
                        _ = v21
                    case -526:
                        if (!(0 > 1)) {
                            fmt.Println(v2)
                            switch 766 {
                            case 9223372036854775807:
                                v1 += v1
                            default:
                                var v22 string = ("hello" + "a")
                                _ = v22
                            }
                        }
                    }
                }
                _ = fi18
                _ = v19
            }
            for fi23 := 0; fi23 < 2; fi23++ {
                var v24 uint8 = v3
                _ = fi23
                _ = v24
            }
            _ = fi11
        }
        for fi25 := 0; fi25 < 8; fi25++ {
            var v26 int = -9223372036854775808
            for fi27 := 0; fi27 < 10; fi27++ {
                v5 = ((3.14159 - 0.0) * ((-2.25 + (float64(0.1) / (v5 + v5))) - 3.14159))
                fmt.Println(v2, v26, fi25)
                _ = fi27
            }
            v28 := GenS2{F0: v4, F1: 2, F2: 42}
            _ = fi25
            _ = v26
            _ = v28
        }
    }
    var v29 uint32 = (v1 + (v1 % (((^(v1))) | 1)))
    v3 += 0
    var v30 uint = 18446744073709551615
    var v31 float32 = (0.0 + (1e10 / 1.5))
    var v32 uint64 = v2
    for fi33 := 0; fi33 < 1; fi33++ {
        for fi34 := 0; fi34 < 20; fi34++ {
            var v35 uint8 = v3
            var v36 int64 = 42
            v30 = 1
            _ = fi34
            _ = v35
            _ = v36
        }
        var v37 int8 = -128
        _ = fi33
        _ = v37
    }
    v5 = v5
    fmt.Println(v29)
    v32 = 1
    for fi38 := 0; fi38 < 19; fi38++ {
        var v39 int8 = 42
        v40 := GenS2{F0: -850, F1: 42, F2: 9223372036854775807}
        _ = fi38
        _ = v39
        _ = v40
    }
    var v41 uint32 = (^(v29))
    v4 -= -847
    v2 -= (^(v32))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
