package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint; F2 int }
func genFn0(p0 string, p1 int8, p2 int16) bool {
    p1 += 127
    v7 := []int{}
    for fi8 := 0; fi8 < 7; fi8++ {
        p2 = (p2 % ((110) | 1))
        p1 += (p1 >> ((65535) & 63))
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := GenS1{F0: 18446744073709551615, F1: 0}
            switch "foo bar" {
            case "foo bar":
                var v11 float64 = 1.5
                _ = v11
            }
            var v12 float32 = 3.14159
            _ = fi9
            _ = v10
            _ = v12
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p0 > "a")
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{1, 991, 830}
    var v3 float32 = 1e10
    var v4 int64 = -295
    var v5 float64 = 0.0
    var v6 float32 = -2.25
    fmt.Println(v4, v4, v6)
    if (v5 != (v5 * (v5 * float64(float32(1e10))))) {
        _ = len(v1)
    } else {
        v13 := GenS1{F0: 849, F1: 305}
        var v14 bool = true
        _ = v13
        _ = v14
    }
    switch 0 {
    case 546:
        v15 := GenS1{F0: 548, F1: v4}
        _ = v15
    case 0:
        fmt.Println(v6, v3, v4)
    default:
        var v16 uint8 = 2
        _ = v16
    }
    _ = len(v2)
    for fi17 := 0; fi17 < 18; fi17++ {
        switch "hello" {
        case "a":
            var v18 string = "hello"
            _ = v18
        }
        for fi19 := 0; fi19 < 6; fi19++ {
            switch (("hello" + "\n\t") + "foo bar") {
            case "":
                v20 := GenS1{F0: 1, F1: (-(v4))}
                _ = v20
            default:
                switch "hello" {
                case "\n\t":
                    if true {
                        if len(v2) > 0 {
                            v2[0] = fi17
                        }
                        v4 += v4
                        v3 -= float32(v5)
                    }
                case "foo bar":
                    var v21 int64 = v4
                    _ = v21
                }
            }
            var v22 uint8 = 42
            var v23 uint32 = 0
            _ = fi19
            _ = v22
            _ = v23
        }
        _ = fi17
    }
    if false {
        v2 = append(v2, 1)
        var v24 int8 = 46
        _ = v24
    } else {
        var v25 int16 = -32768
        _ = v25
    }
    if ((0 < -1) || ((32767 <= 42) && (42 != 127))) {
        if false {
            if false {
                v4 = v4
            }
            v3 = v6
        } else {
            v5 += (v5 + v5)
        }
    } else {
        fmt.Println(v6)
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        for fi27 := 0; fi27 < 15; fi27++ {
            var v28 int32 = -1
            var v29 int16 = 42
            v30 := GenS2{F0: 1}
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
        _ = fi26
    }
    v3 = 3.14159
    for fi31 := 0; fi31 < 5; fi31++ {
        var v32 string = ("" + "\n\t")
        _ = fi31
        _ = v32
    }
    v4 -= 42
    v3 -= ((v6 + float32(-2.25)) - v6)
    if ((-32768 < 1000) && (!false)) {
        var v33 int32 = -356
        _ = v33
    } else {
        v4 -= v4
        switch -964 {
        case 32768:
            v6 = v3
        case 753:
            switch "\n\t" {
            case "\n\t":
                switch 4294967295 {
                case 4294967295:
                    if false {
                        v34 := GenS2{F0: 687}
                        if (((1 >= 0) && (2 <= 42)) && (true || false)) {
                            for fi35 := 0; fi35 < 2; fi35++ {
                                fmt.Println(v6)
                                switch ("hello" + "foo bar") {
                                case "a":
                                    var v36 int = -435
                                    _ = v36
                                case "\n\t":
                                    v6 += (0.1 * v6)
                                case "":
                                    for fi37 := 0; fi37 < 4; fi37++ {
                                        v38 := GenS1{F0: 4294967295, F1: v4}
                                        if (-749 > fi35) {
                                            v2 = append(v2, (fi35 & (-(fi37))))
                                        } else {
                                            if true {
                                                v6 += ((float32(3.14159) * v3) - 3.14159)
                                                fmt.Println(fi35, v4)
                                                var v39 float64 = (0.1 - float64((v6 / float32(1.5))))
                                                _ = v39
                                            }
                                        }
                                        _ = fi37
                                        _ = v38
                                    }
                                }
                                if (-128 >= -1) {
                                    var v40 uint64 = 63
                                    var v41 uint8 = 42
                                    _ = v40
                                    _ = v41
                                }
                                _ = fi35
                            }
                        }
                        v5 += -2.25
                        _ = v34
                    }
                }
            case "a":
                for fi42 := 0; fi42 < 10; fi42++ {
                    var v43 float64 = float64((float32(0.1) * (float32(1.5) * 1e10)))
                    _ = len(v2)
                    _ = fi42
                    _ = v43
                }
            default:
                v44 := GenS1{F0: 42, F1: (v4 >> ((455) & 63))}
                _ = v44
            }
        default:
            _ = len(v2)
        }
    }
    var v45 int32 = 1
    v46 := GenS3{F0: v45, F1: 1, F2: 9223372036854775807}
    var v47 float64 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
