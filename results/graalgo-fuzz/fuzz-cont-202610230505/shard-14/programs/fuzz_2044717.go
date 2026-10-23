package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint16, p1 uint16) int8 {
    v4 := []int{}
    v5 := GenS1{F0: -906, F1: 4294967295}
    var v6 float64 = 0.0
    if false {
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 float32 = 3.14159
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        p1 = (p0 % ((p1) | 1))
        v6 -= (0.0 / 0.1)
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return -128
}

func main() {
    var v1 uint32 = 360
    v2 := []string{}
    v3 := []int{}
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            v1 += 553
            _ = fi11
        }
        _ = fi10
    }
    v1 = v1
    switch ("hello" + "a") {
    case "foo bar":
        if false {
            v1 += v1
            v12 := GenS2{F0: 1}
            _ = v12
        }
    default:
        if (!false) {
            v1 -= v1
            if ((v1 << ((4294967295) & 63)) >= v1) {
                fmt.Println(v1)
                var v13 int32 = 1
                _ = v13
            }
            v1 -= (v1 % ((v1) | 1))
        }
    }
    v14 := GenS3{F0: (-431 >= 631), F1: (v1 / ((2147483647) | 1)), F2: 0}
    if true {
        v1 = (v1 << ((42) & 63))
        fmt.Println(v1)
    } else {
        if true {
            for fi15 := 0; fi15 < 18; fi15++ {
                if (30 <= 1) {
                    var v16 uint8 = 127
                    for fi17 := 0; fi17 < 11; fi17++ {
                        switch ("foo bar" + ("a" + "\n\t")) {
                        case "a":
                            if false {
                                fmt.Println(v1)
                                v18 := GenS2{F0: 0}
                                _ = v18
                            }
                        default:
                            var v19 uint64 = 0
                            _ = v19
                        }
                        _ = fi17
                    }
                    _ = v16
                } else {
                    v1 = v1
                    v20 := GenS4{F0: 840}
                    _ = v20
                }
                fmt.Println(v1)
                _ = fi15
            }
            var v21 int8 = -128
            _ = v21
        }
        var v22 float32 = 0.1
        _ = v22
    }
    var v23 uint32 = v1
    var v24 float32 = ((3.14159 * (-2.25 + (0.0 / 1.0))) / (3.14159 * 0.1))
    var v25 bool = false
    v26 := GenS1{F0: 9223372036854775807, F1: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
