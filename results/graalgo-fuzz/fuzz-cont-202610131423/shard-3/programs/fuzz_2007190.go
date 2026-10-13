package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int }
func genFn0(p0 int64) int8 {
    for fi5 := 0; fi5 < 19; fi5++ {
        fmt.Println(p0)
        _ = fi5
    }
    p0 = (^(p0))
    if (true || true) {
        v6 := GenS1{F0: 4294967295, F1: 42}
        _ = v6
    }
    p0 = p0
    _ = p0
    return 124
}
func genFn1(p0 int32, p1 uint8, p2 uint8) uint {
    if (4294967295 > 2147483647) {
        v7 := GenS1{F0: 660, F1: 65535}
        v8 := []int{-1, -576, 2147483647}
        _ = v7
        _ = v8
    }
    p2 += p2
    v9 := []int{42, 0, 0, 1}
    for fi10 := 0; fi10 < 2; fi10++ {
        if (!((3.14159 / (0.0 + -2.25)) > 0.0)) {
            p0 = p0
            fmt.Println(p0, fi10, p1)
        }
        var v11 int32 = (p0 | -340)
        fmt.Println(p2)
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 65535
}

func main() {
    var v1 int = 128
    var v2 int = v1
    v3 := []uint32{672}
    var v4 bool = (65535 == 183)
    if false {
        if len(v3) > 0 {
            v3[0] = 937
        }
    } else {
        for fi12 := 0; fi12 < 7; fi12++ {
            switch 128 {
            case 2147483647:
                var v13 float32 = float32(float64(0.1))
                _ = v13
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            if (0 <= 0) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            v2 = (v2 * 32768)
            _ = fi12
        }
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
        _ = fi14
    }
    var v15 uint16 = 922
    for fi16 := 0; fi16 < 8; fi16++ {
        if (377 < 0) {
            v2 = v2
        }
        fmt.Println(v4)
        _ = fi16
    }
    var v17 string = "\n\t"
    fmt.Println(v1)
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 uint32 = 1
        v4 = false
        var v20 uint32 = v19
        _ = fi18
        _ = v19
        _ = v20
    }
    if v4 {
        if v4 {
            var v21 uint16 = v15
            for fi22 := 0; fi22 < 14; fi22++ {
                fmt.Println(v1, v17, v15)
                var v23 int64 = 322
                _ = fi22
                _ = v23
            }
            _ = v21
        } else {
            fmt.Println(v2, v4)
            _ = len(v3)
            var v24 uint = 2
            _ = v24
        }
        fmt.Println(v15)
    }
    var v25 int8 = 0
    v17 = v17
    v3 = append(v3, 2147483648)
    fmt.Println(v1)
    var v26 int32 = -1
    for fi27 := 0; fi27 < 5; fi27++ {
        switch (((v17 + v17) + "hello") + (v17 + v17)) {
        case "a":
            v26 = v26
        case "foo bar":
            fmt.Println(fi27)
        case "":
            if ((2147483648 < 2) && (0 < 2)) {
                var v28 int = 4294967295
                _ = v28
            }
        default:
            if (-379 < -32768) {
                var v29 int = 414
                if (!(!true)) {
                    for fi30 := 0; fi30 < 11; fi30++ {
                        v17 = (v17 + v17)
                        if false {
                            v25 += (v25 + (v25 - v25))
                        }
                        _ = fi30
                    }
                    switch (v1 % (((-(v29))) | 1)) {
                    case -2147483648:
                        if v4 {
                            v31 := GenS2{F0: -2147483648}
                            _ = v31
                        }
                    default:
                        v4 = v4
                    }
                } else {
                    switch (v29 & (v2 / ((v2) | 1))) {
                    case 971:
                        if (v4 && (v4 && (!true))) {
                            var v32 int = v1
                            _ = v32
                        } else {
                            v4 = ("hello" < v17)
                            if len(v3) > 0 {
                                v3[0] = 1
                            }
                            fmt.Println(v4, v26, v15)
                        }
                    default:
                        if v4 {
                            if false {
                                for fi33 := 0; fi33 < 6; fi33++ {
                                    var v34 uint = 893
                                    v4 = (v17 != v17)
                                    v34 = (v34 | genFn1(965, 1, 1))
                                    _ = fi33
                                    _ = v34
                                }
                                v2 += (fi27 & int(v25))
                            }
                            v1 -= (v2 << ((1) & 63))
                        } else {
                            v35 := GenS3{F0: -1, F1: (v1 | (v1 >> ((65535) & 63)))}
                            v17 = "\n\t"
                            _ = v35
                        }
                    }
                    _ = len(v3)
                }
                _ = v29
            }
        }
        _ = fi27
    }
    v26 -= 0
    v36 := GenS3{F0: v25, F1: (v2 << ((558) & 63))}
    var v37 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
