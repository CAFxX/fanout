package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 string; F1 int; F2 uint }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 int }
func genFn0(p0 bool) int8 {
    switch 32768 {
    case -9223372036854775808:
        fmt.Println(p0)
    case 1:
        v6 := []int{42, 0, 9223372036854775807}
        _ = v6
    case 32767:
        v7 := GenS1{F0: 42, F1: 32767}
        _ = v7
    default:
        v8 := []int{}
        _ = v8
    }
    if p0 {
        var v9 uint32 = 2
        _ = v9
    }
    var v10 uint32 = 42
    var v11 int32 = 42
    fmt.Println(v11, p0)
    _ = p0
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 int16 = -32768
    var v2 uint32 = 1
    var v3 int32 = -511
    var v4 float32 = (0.0 + (float32(1e10) - 0.1))
    var v5 string = ("foo bar" + (("\n\t" + ("" + ("" + ("a" + "a")))) + "foo bar"))
    v12 := GenS1{F0: -18, F1: v1}
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := []int{}
        _ = fi13
        _ = v14
    }
    v1 -= v1
    v1 -= v1
    var v15 uint32 = v2
    v16 := []int{}
    v17 := GenS1{F0: 23, F1: v1}
    switch 9223372036854775807 {
    case 128:
        if (false || (true || (1 < uint(v15)))) {
            v18 := GenS2{F0: (!true), F1: -2147483648}
            _ = v18
        }
    case -865:
        var v19 int = 0
        _ = v19
    case 269:
        _ = len(v16)
    }
    v5 = v5
    v20 := GenS3{F0: "hello", F1: 2147483647, F2: 65535}
    if (true || true) {
        v21 := GenS4{F0: true, F1: 1}
        var v22 uint32 = 65535
        v1 += -1
        _ = v21
        _ = v22
    } else {
        for fi23 := 0; fi23 < 9; fi23++ {
            for fi24 := 0; fi24 < 20; fi24++ {
                var v25 int64 = 319
                _ = fi24
                _ = v25
            }
            v26 := GenS2{F0: (false && (42 != -1)), F1: -1}
            v3 = (-(v3))
            _ = fi23
            _ = v26
        }
        if (0 != 1) {
            v27 := GenS1{F0: -1, F1: -32768}
            _ = v27
        }
        if ((v5 == v5) || false) {
            v3 = v3
            if true {
                if len(v16) > 0 {
                    v16[0] = 32767
                }
            } else {
                fmt.Println(v1)
                var v28 int = -98
                _ = v28
            }
            v1 = v1
        } else {
            if false {
                v29 := GenS1{F0: 42, F1: (v1 | (v1 - 42))}
                if ((255 < 255) && false) {
                    var v30 uint = 65535
                    _ = v30
                }
                _ = v29
            } else {
                for fi31 := 0; fi31 < 1; fi31++ {
                    v2 -= 687
                    _ = fi31
                }
            }
            if true {
                var v32 float64 = 1.5
                _ = v32
            } else {
                v33 := GenS2{F0: false, F1: (v3 + v3)}
                for fi34 := 0; fi34 < 4; fi34++ {
                    if ((2147483647 < v15) || (!("" != ("" + "a")))) {
                        switch fi34 {
                        case 118:
                            for fi35 := 0; fi35 < 19; fi35++ {
                                if (255 == 2) {
                                    v36 := GenS4{F0: (834 >= 42), F1: 899}
                                    v1 -= int16(fi34)
                                    v37 := GenS5{F0: 42}
                                    v4 = v4
                                    var v38 float64 = -2.25
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                } else {
                                    if len(v16) > 0 {
                                        v16[0] = int(v2)
                                    }
                                }
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                                _ = fi35
                            }
                        default:
                            v4 -= 1.5
                        }
                        for fi39 := 0; fi39 < 19; fi39++ {
                            v40 := GenS2{F0: (2 == 255), F1: (v3 - (v3 << ((42) & 63)))}
                            fmt.Println(v2, v3, fi34)
                            _ = fi39
                            _ = v40
                        }
                    }
                    _ = fi34
                }
                _ = v33
            }
        }
    }
    for fi41 := 0; fi41 < 15; fi41++ {
        v4 = -2.25
        v42 := GenS6{F0: 127, F1: (fi41 * -2147483648)}
        if len(v16) > 0 {
            v16[0] = 32767
        }
        _ = fi41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
