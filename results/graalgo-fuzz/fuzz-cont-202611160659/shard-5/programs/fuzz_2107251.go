package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int8 }

func main() {
    v1 := []uint32{292}
    var v2 uint16 = 1
    var v3 float32 = 0.1
    var v4 int32 = -1
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            v3 -= ((float32(0.0) - ((1e10 + 0.1) + -2.25)) - float32((3.14159 - 0.0)))
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
            _ = fi6
        }
        for fi7 := 0; fi7 < 1; fi7++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v8 int16 = 798
            switch 941 {
            case -158:
                var v9 float64 = 3.14159
                _ = v9
            case 4294967295:
                if len(v1) > 0 {
                    v1[0] = 0
                }
            default:
                var v10 int32 = v4
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        switch 260 {
        case 127:
            if (168 == 796) {
                v3 += v3
                if true {
                    switch "foo bar" {
                    case "":
                        if (false || (4294967295 < 42)) {
                            var v11 int16 = -1
                            for fi12 := 0; fi12 < 5; fi12++ {
                                v4 -= (-(v4))
                                _ = fi12
                            }
                            _ = v11
                        }
                    case "\n\t":
                        var v13 uint = 328
                        _ = v13
                    case "a":
                        fmt.Println(v4, v4)
                    default:
                        var v14 float64 = float64(0.1)
                        _ = v14
                    }
                    var v15 uint = 2
                    v2 += v2
                    _ = v15
                }
            }
        case -487:
            fmt.Println(v2)
        }
        _ = fi5
    }
    v3 = v3
    if true {
        for fi16 := 0; fi16 < 4; fi16++ {
            for fi17 := 0; fi17 < 19; fi17++ {
                v4 = (^(v4))
                v2 = v2
                _ = fi17
            }
            v2 = v2
            _ = fi16
        }
        switch 128 {
        case 1:
            v18 := GenS1{F0: "", F1: 0, F2: int8(v2)}
            _ = v18
        case 32768:
            v19 := GenS1{F0: ("foo bar" + "\n\t"), F1: 186, F2: 0}
            _ = v19
        }
        switch "\n\t" {
        case "":
            for fi20 := 0; fi20 < 3; fi20++ {
                var v21 int16 = -32768
                _ = fi20
                _ = v21
            }
        default:
            for fi22 := 0; fi22 < 19; fi22++ {
                if len(v1) > 0 {
                    v1[0] = 733
                }
                _ = fi22
            }
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v23 int64 = 0
    v2 += (v2 << ((18446744073709551615) & 63))
    switch 4294967295 {
    case 4294967295:
        var v24 int32 = v4
        _ = v24
    }
    v3 -= v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        if len(v1) > 0 {
            v1[0] = 2
        }
        for fi25 := 0; fi25 < 5; fi25++ {
            switch 4294967295 {
            case 1:
                var v26 uint = 42
                _ = v26
            default:
                if (4294967295 > 17) {
                    v2 = v2
                    if (true || true) {
                        if false {
                            var v27 bool = true
                            v27 = v27
                            _ = v27
                        }
                    } else {
                        v23 += v23
                        fmt.Println(v2, v3)
                    }
                    if (2147483648 != 2147483647) {
                        v1 = append(v1, 42)
                    }
                }
            }
            _ = fi25
        }
    }
    v28 := GenS1{F0: "", F1: v4, F2: 100}
    v23 = (v23 | (v23 + v23))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v4
}
