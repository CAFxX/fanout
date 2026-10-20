package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint8 }
func genFn0() uint32 {
    var v4 float64 = 1.5
    v4 = 3.14159
    v4 -= -2.25
    _ = v4
    return 65535
}
func genFn1(p0 int16) int {
    v5 := GenS1{F0: 2, F1: 2147483647, F2: 0}
    var v6 bool = (("hello" != "hello") && false)
    _ = p0
    _ = v5
    _ = v6
    return 697
}

func main() {
    var v1 string = ("\n\t" + ((("a" + "a") + "") + (("\n\t" + "") + "\n\t")))
    var v2 int8 = -128
    var v3 float64 = 1e10
    v7 := GenS1{F0: 20, F1: 28, F2: 18446744073709551615}
    v8 := []int{42, 42, 1}
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            if (!(540 < 1)) {
                if false {
                    _ = len(v8)
                } else {
                    var v11 int = (fi9 | fi10)
                    _ = v11
                }
                var v12 uint16 = 1
                if true {
                    switch v1 {
                    case "":
                        var v13 int32 = -548
                        _ = v13
                    }
                    var v14 int16 = -834
                    _ = v14
                }
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = v12
            }
            _ = fi10
        }
        _ = fi9
    }
    switch genFn1(1) {
    case 436:
        v2 = v2
    case 491:
        var v15 uint = 1
        _ = v15
    default:
        v8 = append(v8, 0)
    }
    switch -9223372036854775808 {
    case -703:
        v3 += v3
    case 32767:
        switch (v1 + "a") {
        case "a":
            for fi16 := 0; fi16 < 15; fi16++ {
                v17 := GenS1{F0: genFn0(), F1: 2147483648, F2: 1}
                _ = fi16
                _ = v17
            }
        case "foo bar":
            v2 = v2
        case "":
            if true {
                var v18 uint64 = 18446744073709551615
                v18 += (v18 >> ((222) & 63))
                switch 128 {
                case -2147483648:
                    v2 = v2
                }
                _ = v18
            }
        default:
            if (false && (32767 <= 277)) {
                v3 -= (0.0 * 1.5)
                v19 := GenS2{F0: 0, F1: 401, F2: 255}
                var v20 float32 = (((-2.25 + (1.5 * -2.25)) * (-2.25 + 1e10)) * (float32(v3) * (3.14159 + float32(float64((0.0 + ((1e10 + (1.5 - float32((float64((((1e10 * (1e10 + 0.0)) * 0.0) + (3.14159 + (1.5 / 3.14159)))) * v3)))) - (1e10 / float32(v3)))))))))
                _ = v19
                _ = v20
            }
        }
    default:
        var v21 int64 = -1
        _ = v21
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 int16 = 32767
        var v24 int32 = int32(v23)
        _ = fi22
        _ = v23
        _ = v24
    }
    switch 4294967295 {
    case -1:
        fmt.Println(v3)
    case 1:
        v25 := GenS2{F0: 255, F1: -1, F2: 42}
        _ = v25
    }
    v3 = v3
    if (9223372036854775807 != 110) {
        if len(v8) > 0 {
            v8[0] = -2147483648
        }
    } else {
        for fi26 := 0; fi26 < 14; fi26++ {
            fmt.Println(v2)
            _ = fi26
        }
    }
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
