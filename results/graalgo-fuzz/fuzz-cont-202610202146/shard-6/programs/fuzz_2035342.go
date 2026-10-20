package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int64 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int8; F1 int8; F2 int8 }
type GenS4 struct { F0 int; F1 uint }
func genFn0(p0 uint32, p1 int32, p2 uint8) string {
    v5 := []int{128, 1}
    switch 2147483647 {
    case 32768:
        p1 -= (p1 - p1)
    case 806:
        fmt.Println(p1, p0)
    case 2147483647:
        v6 := GenS1{F0: -32768, F1: 1, F2: 1}
        _ = v6
    default:
        fmt.Println(p0, p0)
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ("\n\t" + "foo bar")
}
func genFn1(p0 uint, p1 int16) int {
    var v7 int16 = (p1 + p1)
    v8 := GenS2{F0: 255, F1: false}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 739
}

func main() {
    var v1 uint = 4294967295
    var v2 bool = false
    var v3 int64 = 9223372036854775807
    var v4 int32 = 42
    switch "foo bar" {
    case "a":
        switch "" {
        case "\n\t":
            for fi9 := 0; fi9 < 12; fi9++ {
                v2 = ("foo bar" != "")
                _ = fi9
            }
        case "hello":
            fmt.Println(v3, v4, v4)
        default:
            v10 := []int{}
            _ = v10
        }
    }
    if (1e10 >= (float64(float32(1.5)) - 1e10)) {
        if v2 {
            v11 := []int{1, -1, 127}
            _ = v11
        } else {
            v12 := []int{639}
            for fi13 := 0; fi13 < 20; fi13++ {
                v2 = (fi13 < genFn1(v1, 1))
                v3 = (v3 * 1)
                var v14 int8 = -1
                _ = fi13
                _ = v14
            }
            _ = v12
        }
        v15 := GenS3{F0: 127, F1: 127, F2: 0}
        for fi16 := 0; fi16 < 12; fi16++ {
            if true {
                v4 -= (v4 >> ((1) & 63))
                v1 -= v1
            } else {
                v4 += (v4 ^ -2147483648)
                for fi17 := 0; fi17 < 5; fi17++ {
                    switch fi17 {
                    case 32767:
                        switch (fi17 & (-(fi17))) {
                        case -162:
                            switch ("" + "hello") {
                            case "a":
                                v1 = v1
                                var v18 int32 = v4
                                v19 := []int{genFn1(632, 50), 1, (fi16 << (((v1 ^ v1)) & 63))}
                                _ = v18
                                _ = v19
                            }
                        case 145:
                            v1 = (v1 ^ v1)
                        case -910:
                            v20 := []int{}
                            _ = v20
                        default:
                            if (v4 != v4) {
                                v1 = (v1 ^ v1)
                                var v21 string = "hello"
                                if (32767 != 602) {
                                    v4 = 1
                                }
                                _ = v21
                            }
                        }
                    case 9223372036854775807:
                        v1 -= 1
                    case 4294967295:
                        var v22 uint32 = 65535
                        _ = v22
                    default:
                        var v23 string = ("hello" + "foo bar")
                        _ = v23
                    }
                    _ = fi17
                }
            }
            v3 -= 39
            if v2 {
                v24 := []int{(-(fi16)), fi16, 802, (fi16 % ((2147483647) | 1))}
                v3 = (^(v3))
                _ = v24
            } else {
                var v25 uint64 = 189
                if (!false) {
                    var v26 uint16 = 65535
                    for fi27 := 0; fi27 < 17; fi27++ {
                        var v28 string = ((("hello" + "hello") + "\n\t") + "a")
                        var v29 int = fi16
                        var v30 int64 = v3
                        _ = fi27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    var v31 bool = (1 > 149)
                    _ = v26
                    _ = v31
                } else {
                    v3 -= (v3 ^ 1)
                    v32 := GenS1{F0: 32767, F1: 1, F2: (v3 / ((9223372036854775807) | 1))}
                    _ = v32
                }
                v33 := []int{(fi16 % ((fi16) | 1)), 32767, 2147483647}
                _ = v25
                _ = v33
            }
            _ = fi16
        }
        _ = v15
    }
    for fi34 := 0; fi34 < 9; fi34++ {
        var v35 uint = v1
        _ = fi34
        _ = v35
    }
    v36 := []int{409, -909}
    fmt.Println(v3, v1)
    v37 := GenS4{F0: 128, F1: v1}
    v36 = append(v36, -2147483648)
    v38 := GenS1{F0: 0, F1: 32767, F2: -424}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
