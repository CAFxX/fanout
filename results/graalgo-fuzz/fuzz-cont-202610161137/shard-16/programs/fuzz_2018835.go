package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }

func main() {
    var v1 float32 = (3.14159 - (1.5 * 1e10))
    var v2 uint64 = 42
    var v3 int8 = -128
    v4 := make(map[uint32]int)
    if false {
        var v5 string = ""
        v6 := GenS1{F0: 1, F1: 472}
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := []int{fi7, (fi7 / ((145) | 1))}
        v2 = v2
        var v9 uint16 = 65535
        _ = fi7
        _ = v8
        _ = v9
    }
    v1 += 0.1
    fmt.Println(v2)
    _ = len(v4)
    v1 -= v1
    switch 9223372036854775807 {
    case 1:
        v10 := v4[2]
        _ = v10
    }
    v2 -= (v2 * (^(v2)))
    if (4294967295 > -915) {
        for fi11 := 0; fi11 < 12; fi11++ {
            switch fi11 {
            case 4294967295:
                if false {
                    v2 = (v2 * v2)
                    for fi12 := 0; fi12 < 10; fi12++ {
                        v1 = (0.0 - ((v1 / -2.25) + 1e10))
                        switch ("" + ("foo bar" + "foo bar")) {
                        case "":
                            var v13 int64 = -741
                            _ = v13
                        case "foo bar":
                            var v14 uint32 = 65535
                            _ = v14
                        case "\n\t":
                            var v15 uint8 = uint8(v2)
                            _ = v15
                        default:
                            v3 = v3
                        }
                        _ = fi12
                    }
                }
            case -1:
                delete(v4, 2147483648)
            case 9223372036854775807:
                v16 := GenS1{F0: 29, F1: 411}
                _ = v16
            default:
                v17 := GenS1{F0: 42, F1: int16(v3)}
                _ = v17
            }
            v3 -= (v3 >> ((65535) & 63))
            _ = fi11
        }
        var v18 uint16 = 32767
        if false {
            v3 += -128
            for fi19 := 0; fi19 < 14; fi19++ {
                if (!(true && (true && true))) {
                    var v20 uint8 = 106
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 16; fi21++ {
                        v1 = float32((1.5 * 0.0))
                        fmt.Println(fi19, v2)
                        _ = fi21
                    }
                    v3 += v3
                }
                var v22 int16 = 32767
                _ = fi19
                _ = v22
            }
        }
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
