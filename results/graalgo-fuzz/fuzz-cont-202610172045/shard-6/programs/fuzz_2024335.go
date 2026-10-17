package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint32, p1 uint32) int {
    v7 := GenS1{F0: 65535}
    var v8 int16 = 32767
    p1 -= (p1 - (p0 >> ((255) & 63)))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 int64, p1 int32) string {
    switch 4294967295 {
    case 2147483647:
        if (true || true) {
            for fi9 := 0; fi9 < 14; fi9++ {
                v10 := GenS1{F0: 42}
                _ = fi9
                _ = v10
            }
            switch ("hello" + "") {
            case "a":
                if true {
                    fmt.Println(p1, p0)
                }
            case "\n\t":
                var v11 uint32 = 1
                _ = v11
            default:
                var v12 uint32 = 4294967295
                _ = v12
            }
        }
    case -148:
        v13 := []int{9223372036854775807}
        _ = v13
    }
    var v14 int32 = p1
    var v15 float32 = float32(-2.25)
    var v16 string = "hello"
    v17 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return v16
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]bool)
    var v3 uint16 = 323
    v4 := []uint32{}
    var v5 float64 = 0.1
    v6 := make(map[int]int)
    if ((3.14159 / -2.25) < ((0.0 + 3.14159) + (0.0 * float32(3.14159)))) {
        v3 = v3
    }
    var v18 int16 = -771
    v18 += v18
    var v19 uint8 = 42
    v3 -= v3
    fmt.Println(v3, v3, v19)
    v1 = (v1 << ((18446744073709551615) & 63))
    var v20 uint = 356
    var v21 int64 = 464
    if true {
        var v22 uint16 = 0
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        v1 += 32767
        _ = fi23
    }
    var v24 uint16 = (v3 ^ (v3 >> ((v20) & 63)))
    switch "foo bar" {
    case "\n\t":
        var v25 int16 = (v18 % ((v18) | 1))
        _ = v25
    case "hello":
        switch -343 {
        case -9223372036854775808:
            if (1 != -128) {
                switch genFn0(uint32(v19), 4294967295) {
                case 0:
                    for fi26 := 0; fi26 < 14; fi26++ {
                        switch genFn0(898, 2147483648) {
                        case 127:
                            var v27 int = fi26
                            _ = v27
                        default:
                            switch fi26 {
                            case -372:
                                v1 = v1
                            case 233:
                                v28 := GenS1{F0: v20}
                                _ = v28
                            }
                        }
                        _ = len(v2)
                        _ = fi26
                    }
                case 32767:
                    v29 := GenS2{F0: v19}
                    _ = v29
                default:
                    _ = len(v4)
                }
            } else {
                var v30 int32 = -1
                if true {
                    for fi31 := 0; fi31 < 14; fi31++ {
                        fmt.Println(v5, v20)
                        _ = fi31
                    }
                    if true {
                        if (-1 <= 0) {
                            _ = len(v2)
                            v32 := GenS2{F0: (v19 ^ 1)}
                            _ = v32
                        }
                        var v33 float64 = v5
                        v1 += (v18 >> (((v20 - uint(v24))) & 63))
                        _ = v33
                    } else {
                        var v34 uint64 = 18446744073709551615
                        var v35 int64 = v21
                        _ = v34
                        _ = v35
                    }
                }
                fmt.Println(v30, v21, v30)
                _ = v30
            }
        case 536:
            var v36 string = "\n\t"
            _ = v36
        case 2147483647:
            v4 = append(v4, 842)
        default:
            if len(v4) > 0 {
                v4[0] = 42
            }
        }
    case "foo bar":
        for fi37 := 0; fi37 < 4; fi37++ {
            v18 -= v1
            v1 = v18
            _ = fi37
        }
    default:
        var v38 uint8 = uint8(v3)
        _ = v38
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v20)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
