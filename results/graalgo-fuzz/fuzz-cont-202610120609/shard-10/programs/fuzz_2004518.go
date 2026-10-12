package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 bool }
type GenS2 struct { F0 bool; F1 int32; F2 string }
func genFn0(p0 int16, p1 uint16) uint16 {
    v6 := []int{42, -534, 378}
    p0 -= p0
    _ = p0
    _ = p1
    _ = v6
    return p1
}
func genFn1(p0 int64) bool {
    switch 9223372036854775807 {
    case -120:
        for fi7 := 0; fi7 < 17; fi7++ {
            switch fi7 {
            case 340:
                var v8 bool = (100 != 42)
                _ = v8
            case 0:
                var v9 uint8 = 127
                _ = v9
            case 256:
                var v10 float64 = 1e10
                _ = v10
            }
            switch "a" {
            case "\n\t":
                switch fi7 {
                case -928:
                    p0 = 42
                case 4294967295:
                    var v11 uint32 = 65535
                    _ = v11
                }
            case "a":
                var v12 float32 = (float32(((1e10 + 0.1) * -2.25)) + (float32(float64(1e10)) - ((-2.25 * 3.14159) + 0.1)))
                _ = v12
            }
            if true {
                fmt.Println(fi7, fi7)
                fmt.Println(p0, p0)
            }
            _ = fi7
        }
    case 0:
        var v13 uint64 = 2
        _ = v13
    case -9223372036854775808:
        var v14 int8 = -57
        _ = v14
    default:
        fmt.Println(p0)
    }
    fmt.Println(p0)
    p0 = 0
    v15 := []int{221, 42, 0}
    fmt.Println(p0)
    _ = p0
    _ = v15
    return false
}

func main() {
    var v1 int64 = 0
    var v2 string = "a"
    v3 := make(map[uint32]int)
    v4 := []int{-2147483648, -946}
    var v5 uint8 = 128
    v3[0] = 4294967295
    _ = len(v4)
    v16 := v3[1]
    for fi17 := 0; fi17 < 3; fi17++ {
        for fi18 := 0; fi18 < 15; fi18++ {
            if true {
                var v19 uint32 = 2147483648
                _ = v19
            } else {
                var v20 string = (v2 + v2)
                fmt.Println(v2, v16)
                if (!false) {
                    v2 = v2
                    if (4294967295 != 500) {
                        switch v16 {
                        case 166:
                            fmt.Println(v16, v20, v16)
                            v21 := GenS1{F0: 519, F1: (^(fi18)), F2: true}
                            v3[714] = 1
                            v16 = fi18
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v21
                        case -744:
                            v4 = append(v4, (^(v16)))
                        }
                        var v22 uint16 = 255
                        _ = v22
                    }
                    _ = len(v3)
                    var v23 uint16 = 1
                    _ = v23
                }
                v24 := GenS2{F0: (18446744073709551615 != 4294967295), F1: 1, F2: "foo bar"}
                _ = v20
                _ = v24
            }
            _ = fi18
        }
        var v25 uint = 418
        var v26 int8 = 42
        _ = fi17
        _ = v25
        _ = v26
    }
    if (v16 > v16) {
        for fi27 := 0; fi27 < 14; fi27++ {
            v1 += -1
            v28 := GenS2{F0: true, F1: -2147483648, F2: v2}
            switch v16 {
            case 0:
                v2 = v2
            case 9223372036854775807:
                v29 := v3[4294967295]
                _ = v29
            case 735:
                v16 += -575
            }
            _ = fi27
            _ = v28
        }
        v30 := GenS2{F0: true, F1: 0, F2: v2}
        v16 += int(v5)
        _ = v30
    }
    for fi31 := 0; fi31 < 7; fi31++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v1 = (v1 >> ((715) & 63))
        _ = fi31
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        var v33 bool = false
        v16 = v16
        _ = fi32
        _ = v33
    }
    fmt.Println(v5, v16)
    v5 = 1
    _ = len(v4)
    if (!false) {
        for fi34 := 0; fi34 < 3; fi34++ {
            _ = len(v4)
            if false {
                switch (fi34 - -2147483648) {
                case 877:
                    v16 -= (v16 % ((912) | 1))
                case -101:
                    v16 += (v16 | v16)
                }
                v16 += (-(v16))
            }
            for fi35 := 0; fi35 < 10; fi35++ {
                v16 = v16
                _ = fi35
            }
            _ = fi34
        }
    } else {
        var v36 int32 = 2147483647
        _ = v36
    }
    v37 := GenS1{F0: 32767, F1: 42, F2: true}
    v16 = 128
    var v38 bool = (546 <= 0)
    switch 9223372036854775807 {
    case 1:
        if (177 > 495) {
            var v39 string = "a"
            fmt.Println(v39, v39)
            _ = v39
        }
    case 127:
        var v40 int8 = 33
        _ = v40
    default:
        switch v2 {
        case "a":
            var v41 uint32 = 2147483648
            _ = v41
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
