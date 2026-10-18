package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v4 int64 = 1
    v5 := make(map[string]uint32)
    var v6 uint32 = 916
    v6 += v6
    var v7 uint8 = 0
    var v8 bool = (0.0 < -2.25)
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 string = "hello"
        if ((65535 < 32767) && (v7 == 2)) {
            v6 = v6
            var v11 uint16 = 0
            if ((!true) && (0.0 == float32(0.1))) {
                v4 = (^(v4))
            }
            _ = v11
        } else {
            switch (^(fi9)) {
            case -2147483648:
                v12 := GenS1{F0: 18446744073709551615}
                _ = v12
            default:
                v10 = v10
            }
        }
        _ = fi9
        _ = v10
    }
    if (float64(((0.0 + -2.25) * (0.0 + 1.5))) != (1e10 / 1e10)) {
        var v13 int64 = 9223372036854775807
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        v6 -= 927
        _ = fi14
    }
    fmt.Println(v6, v4, v7)
    var v15 int = -1
    v6 = v6
    var v16 uint16 = 255
    fmt.Println(v6, v8)
    for fi17 := 0; fi17 < 4; fi17++ {
        v15 = 42
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 uint64 = 845
            var v20 int = (fi17 * fi18)
            v7 -= 174
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 uint64 = 759
        _ = fi17
        _ = v21
    }
    switch v15 {
    case 9223372036854775807:
        if true {
            switch (v15 | (v15 & v15)) {
            case -9223372036854775808:
                v22 := GenS1{F0: 469}
                _ = v22
            default:
                for fi23 := 0; fi23 < 15; fi23++ {
                    for fi24 := 0; fi24 < 5; fi24++ {
                        var v25 int16 = 1000
                        fmt.Println(v8, v15, v4)
                        _ = fi24
                        _ = v25
                    }
                    fmt.Println(v8)
                    _ = fi23
                }
            }
        } else {
            if (0 < 120) {
                for fi26 := 0; fi26 < 16; fi26++ {
                    v8 = false
                    _ = fi26
                }
                v27 := GenS1{F0: 2}
                _ = v27
            }
            if (1.5 != -2.25) {
                switch "" {
                case "hello":
                    v15 = v15
                default:
                    for fi28 := 0; fi28 < 3; fi28++ {
                        v29 := GenS1{F0: 18446744073709551615}
                        _ = fi28
                        _ = v29
                    }
                }
            }
        }
    case 1:
        var v30 int16 = -32768
        _ = v30
    default:
        v6 -= 65535
    }
    fmt.Println(v4, v6, v16)
    v16 += v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
