package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 bool; F1 uint; F2 int16 }
func genFn0(p0 uint8) uint16 {
    fmt.Println(p0)
    var v5 float64 = (3.14159 * (0.0 - ((float64((3.14159 + 1e10)) / -2.25) * (0.0 + 1e10))))
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint16 = 1
        if (float64(-2.25) <= 0.0) {
            v8 := []int{128, 128, (fi6 + 9223372036854775807), fi6}
            _ = v8
        } else {
            v9 := GenS1{F0: 2147483647}
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v5
    return 284
}

func main() {
    var v1 int8 = 127
    var v2 uint32 = 4294967295
    var v3 float32 = -2.25
    var v4 int = 32767
    var v10 int16 = -382
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 12; fi12++ {
            v3 += (((1.5 + v3) - v3) + 1e10)
            _ = fi12
        }
        _ = fi11
    }
    v3 = ((3.14159 - v3) - 0.0)
    v13 := []int{(v4 + int(v2))}
    if (v3 <= (3.14159 + (((-2.25 - v3) - (v3 / 1.5)) - (((v3 * float32(1.5)) * 0.1) + v3)))) {
        var v14 uint64 = 126
        if (true && (!(!(255 != 308)))) {
            switch (v4 >> ((516) & 63)) {
            case 604:
                v1 = (v1 + v1)
            case 4294967295:
                var v15 uint = 817
                _ = v15
            default:
                for fi16 := 0; fi16 < 5; fi16++ {
                    switch "a" {
                    case "":
                        var v17 int8 = v1
                        _ = v17
                    case "a":
                        v18 := GenS1{F0: 1}
                        _ = v18
                    default:
                        switch "hello" {
                        case "a":
                            v13 = append(v13, (v4 >> ((113) & 63)))
                        }
                    }
                    _ = fi16
                }
            }
            v2 = v2
        } else {
            if true {
                for fi19 := 0; fi19 < 2; fi19++ {
                    if true {
                        for fi20 := 0; fi20 < 5; fi20++ {
                            v14 += (v14 * 4294967295)
                            _ = fi20
                        }
                    }
                    v21 := GenS2{F0: 1}
                    v1 = v1
                    _ = fi19
                    _ = v21
                }
                var v22 uint8 = 127
                for fi23 := 0; fi23 < 1; fi23++ {
                    v1 -= (v1 * (v1 ^ int8(v4)))
                    v1 = -1
                    _ = fi23
                }
                _ = v22
            }
        }
        if (((v10 - (v10 >> ((59) & 63))) >= v10) || ((false && (!(false || false))) || false)) {
            v10 += 42
            if (1 != 2) {
                if len(v13) > 0 {
                    v13[0] = (v4 >> ((2) & 63))
                }
            } else {
                for fi24 := 0; fi24 < 6; fi24++ {
                    v14 -= (v14 ^ 1)
                    if (((v3 - v3) - v3) != v3) {
                        v10 = v10
                        var v25 string = "\n\t"
                        _ = v25
                    }
                    _ = fi24
                }
            }
            if len(v13) > 0 {
                v13[0] = v4
            }
        }
        _ = v14
    } else {
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            v10 = v10
        }
    }
    _ = len(v13)
    fmt.Println(v1)
    for fi26 := 0; fi26 < 12; fi26++ {
        v27 := GenS1{F0: 2147483647}
        v28 := GenS3{F0: 630, F1: 762}
        _ = fi26
        _ = v27
        _ = v28
    }
    v3 = v3
    for fi29 := 0; fi29 < 7; fi29++ {
        var v30 uint16 = 255
        var v31 bool = true
        var v32 uint64 = uint64(v2)
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    v33 := GenS4{F0: (2 != (v2 - v2)), F1: 4294967295, F2: v10}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
