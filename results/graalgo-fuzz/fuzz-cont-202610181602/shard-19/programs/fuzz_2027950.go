package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 int; F2 uint16 }
type GenS5 struct { F0 uint; F1 int8; F2 int64 }
func genFn0(p0 int, p1 int, p2 int) int16 {
    p0 -= p1
    v7 := []int{}
    for fi8 := 0; fi8 < 4; fi8++ {
        switch p2 {
        case -1:
            if (!(2 > 42)) {
                for fi9 := 0; fi9 < 13; fi9++ {
                    v10 := GenS1{F0: 1, F1: 2}
                    v11 := GenS2{F0: 4294967295, F1: 1000}
                    if false {
                        var v12 uint64 = 1
                        _ = v12
                    }
                    var v13 int32 = 72
                    _ = len(v7)
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v13
                }
                if false {
                    p0 += -949
                }
                p0 -= -1
            } else {
                v14 := GenS3{F0: 2}
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v14
            }
        default:
            if false {
                var v15 uint32 = 2147483648
                fmt.Println(p2, v15)
                _ = v15
            } else {
                p2 = (fi8 << ((666) & 63))
                if (true || (true && (-128 != -21))) {
                    _ = len(v7)
                } else {
                    p1 = (fi8 & fi8)
                    v7 = append(v7, (p2 * (p1 + p0)))
                }
                fmt.Println(p2)
            }
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi8
    }
    var v16 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v7
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 0
    var v3 int64 = 42
    var v4 uint32 = 4294967295
    var v5 float64 = (1.5 / 1.0)
    v6 := make(map[uint32]int)
    var v17 string = (("foo bar" + "foo bar") + ((("hello" + "\n\t") + "hello") + ("foo bar" + "foo bar")))
    v18 := []int{}
    v4 -= (v4 | v4)
    var v19 string = ((("a" + "") + v17) + (v17 + v17))
    switch 0 {
    case -880:
        if false {
            _ = len(v18)
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 int16 = (v2 * -301)
                _ = fi20
                _ = v21
            }
            v22 := GenS1{F0: 4294967295, F1: (v4 & 2147483648)}
            _ = v22
        }
    default:
        fmt.Println(v2, v2, v17)
    }
    switch 128 {
    case 9223372036854775807:
        fmt.Println(v3, v17, v2)
    case 0:
        var v23 int32 = -80
        _ = v23
    case 32767:
        var v24 float64 = ((((v5 * v5) - (((((1.5 - -2.25) + v5) * (float64(0.0) / ((v5 * 3.14159) * v5))) / ((v5 + (v5 / v5)) / 1.0)) + float64(0.1))) + ((v5 + v5) - 0.1)) * 1.5)
        _ = v24
    }
    v3 += 1
    for fi25 := 0; fi25 < 17; fi25++ {
        var v26 uint = 255
        _ = fi25
        _ = v26
    }
    var v27 int8 = 29
    for fi28 := 0; fi28 < 7; fi28++ {
        v29 := v6[677]
        fmt.Println(v27, v17)
        _ = fi28
        _ = v29
    }
    var v30 float32 = 1e10
    switch 303 {
    case -2147483648:
        v31 := GenS1{F0: 65535, F1: 932}
        _ = v31
    default:
        v18 = append(v18, 127)
    }
    switch 0 {
    case 697:
        for fi32 := 0; fi32 < 1; fi32++ {
            switch (fi32 >> ((2) & 63)) {
            case -9223372036854775808:
                v33 := GenS4{F0: (-(v27)), F1: fi32, F2: 0}
                _ = v33
            case 127:
                for fi34 := 0; fi34 < 12; fi34++ {
                    if (!false) {
                        var v35 int8 = v27
                        fmt.Println(v19, v19)
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(v35)
                        case "foo bar":
                            v17 = (v19 + v19)
                        default:
                            v36 := GenS5{F0: 2, F1: v35, F2: (v3 >> ((775) & 63))}
                            _ = v36
                        }
                        _ = v35
                    }
                    delete(v1, (fi34 - fi32))
                    fmt.Println(v3, v3, fi32)
                    _ = fi34
                }
            case 842:
                if (201 < uint8(v4)) {
                    v37 := GenS5{F0: 18446744073709551615, F1: 127, F2: v3}
                    if ((-(v3)) == v3) {
                        switch v19 {
                        case "":
                            fmt.Println(v17)
                        }
                    }
                    var v38 float64 = v5
                    _ = v37
                    _ = v38
                } else {
                    v5 -= v5
                    for fi39 := 0; fi39 < 20; fi39++ {
                        switch 128 {
                        case 127:
                            var v40 bool = (((v17 + "foo bar") == (v17 + v17)) || true)
                            _ = v40
                        case -639:
                            var v41 int = fi39
                            _ = v41
                        default:
                            v5 -= (v5 / ((0.0 + (v5 / 3.14159)) - float64(v30)))
                        }
                        v3 = (v3 * (v3 << ((4294967295) & 63)))
                        _ = fi39
                    }
                }
            }
            switch fi32 {
            case 2147483647:
                var v42 int32 = -134
                _ = v42
            default:
                var v43 uint64 = 1
                _ = v43
            }
            var v44 int8 = 14
            _ = fi32
            _ = v44
        }
    default:
        if false {
            v5 = float64(0.0)
            if true {
                fmt.Println(v4)
                v4 -= v4
            } else {
                v17 = v17
                _ = len(v18)
            }
        }
    }
    var v45 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
