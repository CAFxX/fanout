package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int = -128
    var v3 float64 = 1e10
    var v4 uint64 = 641
    if false {
        v5 := GenS1{F0: v2}
        if true {
            v6 := []int{}
            _ = len(v6)
            _ = v6
        }
        var v7 float32 = (-2.25 * (1.5 + -2.25))
        _ = v5
        _ = v7
    } else {
        v8 := []int{378, v2, v2}
        v9 := GenS1{F0: v2}
        _ = v8
        _ = v9
    }
    var v10 uint = 42
    v11 := []int{(-(v2)), (v2 - v2), v2}
    switch ("" + "hello") {
    case "":
        for fi12 := 0; fi12 < 11; fi12++ {
            switch fi12 {
            case 42:
                var v13 int32 = 437
                _ = v13
            default:
                fmt.Println(fi12, v1)
            }
            _ = fi12
        }
    case "a":
        var v14 int64 = (v1 * 1)
        _ = v14
    case "\n\t":
        v15 := GenS1{F0: (v2 % ((v2) | 1))}
        _ = v15
    default:
        var v16 float32 = -2.25
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 uint32 = 42
        switch (fi17 ^ fi17) {
        case 0:
            if (!(false && false)) {
                switch v2 {
                case -155:
                    v10 += v10
                case -9223372036854775808:
                    for fi19 := 0; fi19 < 18; fi19++ {
                        if len(v11) > 0 {
                            v11[0] = (fi17 / ((fi19) | 1))
                        }
                        switch fi17 {
                        case 127:
                            var v20 int8 = 127
                            _ = v20
                        default:
                            v3 = float64(-2.25)
                        }
                        if (!((0.1 + 1.5) != 1.5)) {
                            var v21 int8 = 100
                            _ = v21
                        }
                        _ = fi19
                    }
                case -1:
                    v4 = uint64(v18)
                }
            } else {
                v18 = (v18 / ((v18) | 1))
                v1 = v1
                fmt.Println(v4)
            }
        }
        var v22 float64 = float64(-2.25)
        _ = fi17
        _ = v18
        _ = v22
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        switch v2 {
        case -1:
            switch 2147483647 {
            case -1:
                switch v2 {
                case 127:
                    v3 = v3
                default:
                    v4 += (v4 >> ((v10) & 63))
                }
            case 1:
                v4 -= 18446744073709551615
            default:
                var v24 uint = 677
                _ = v24
            }
        case -604:
            _ = len(v11)
        case 100:
            var v25 uint8 = 127
            _ = v25
        }
        var v26 uint16 = 112
        _ = fi23
        _ = v26
    }
    v4 += (^(v4))
    fmt.Println(v10, v4)
    for fi27 := 0; fi27 < 16; fi27++ {
        v4 = (^(v4))
        _ = fi27
    }
    var v28 uint32 = uint32(v10)
    if (!((^(v10)) != (^(v10)))) {
        v29 := GenS1{F0: (v2 ^ int(v4))}
        var v30 uint = (v10 ^ v10)
        _ = v29
        _ = v30
    } else {
        fmt.Println(v28)
        v10 += (v10 & (v10 | (v10 - (v10 - 42))))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
