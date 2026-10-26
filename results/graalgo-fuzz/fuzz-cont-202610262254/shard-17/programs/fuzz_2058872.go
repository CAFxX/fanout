package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 int }
func genFn0() uint16 {
    var v5 int64 = -577
    switch "foo bar" {
    case "":
        for fi6 := 0; fi6 < 7; fi6++ {
            v5 += v5
            var v7 int32 = 754
            _ = fi6
            _ = v7
        }
    case "foo bar":
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := GenS1{F0: 4294967295, F1: -91}
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            if false {
                var v12 uint8 = 1
                _ = v12
            }
            _ = fi11
        }
        var v13 bool = true
        _ = fi10
        _ = v13
    }
    switch 0 {
    case 524:
        for fi14 := 0; fi14 < 15; fi14++ {
            for fi15 := 0; fi15 < 14; fi15++ {
                v16 := []int{(^(fi15)), 542, (fi15 % ((fi15) | 1)), fi14}
                _ = fi15
                _ = v16
            }
            v17 := GenS2{F0: 225, F1: "hello"}
            if true {
                v5 = (v5 & (v5 | v5))
            }
            _ = fi14
            _ = v17
        }
    case -614:
        if true {
            fmt.Println(v5)
            for fi18 := 0; fi18 < 14; fi18++ {
                for fi19 := 0; fi19 < 10; fi19++ {
                    v20 := []int{fi18, (fi19 % ((fi19) | 1)), -635, fi19}
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
        }
    }
    _ = v5
    return 761
}
func genFn1(p0 int64, p1 uint32, p2 string) uint {
    v21 := []int{}
    var v22 uint8 = 203
    var v23 float64 = 0.0
    v24 := GenS3{F0: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    return 129
}

func main() {
    var v1 bool = (!true)
    v2 := make(map[uint32]int)
    v3 := []int{128, 42, 4294967295, -9223372036854775808}
    var v4 uint8 = 42
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 uint16 = 2
        v2[1] = fi25
        _ = fi25
        _ = v26
    }
    if ((303 != 4294967295) || v1) {
        v27 := GenS2{F0: 255, F1: "\n\t"}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v27
    }
    var v28 float32 = ((0.0 + 0.1) * (((float32(0.1) - (1e10 / (1.5 * (((3.14159 - (((float32(-2.25) - ((0.1 + (float32((0.1 - float64(0.0))) * (3.14159 * 1e10))) + float32(1e10))) - 1e10) + 0.1)) + (float32(0.0) / 1e10)) / 0.1)))) - (((1e10 + float32((1e10 / 1.0))) * 3.14159) + 0.0)) / (0.0 + 1.5)))
    var v29 uint = 2
    for fi30 := 0; fi30 < 18; fi30++ {
        var v31 int8 = 102
        if true {
            if ((((3.14159 - 0.0) - (((1.5 * (0.1 / 1e10)) * 1.5) / (float64(3.14159) - 0.1))) * float64(v28)) <= (0.0 / (((1.5 + 1e10) + 0.1) * float64(v28)))) {
                fmt.Println(v29, fi30, v1)
                v31 += 127
                v4 = v4
            } else {
                v1 = v1
                if false {
                    v32 := v2[2147483647]
                    _ = v32
                }
            }
            for fi33 := 0; fi33 < 17; fi33++ {
                v4 = v4
                _ = fi33
            }
        } else {
            v29 -= (^(v29))
            v34 := GenS4{F0: (2147483647 <= 2147483647)}
            _ = len(v2)
            _ = v34
        }
        if (32767 < genFn0()) {
            var v35 float32 = 0.1
            var v36 string = (((("" + "a") + "foo bar") + "foo bar") + "foo bar")
            _ = v35
            _ = v36
        }
        _ = fi30
        _ = v31
    }
    v4 = v4
    var v37 uint32 = 147
    switch 128 {
    case 128:
        switch "hello" {
        case "a":
            v38 := v2[v37]
            _ = v38
        case "foo bar":
            fmt.Println(v28, v28, v29)
        case "":
            _ = len(v2)
        }
    case -2147483648:
        var v39 int = 4294967295
        _ = v39
    default:
        v40 := GenS2{F0: 501, F1: ""}
        _ = v40
    }
    switch ("hello" + "a") {
    case "a":
        if len(v3) > 0 {
            v3[0] = int(v4)
        }
    }
    v41 := GenS5{F0: "hello", F1: 479}
    if false {
        v28 += (float32(1e10) + 1.5)
    }
    if v1 {
        var v42 float64 = 1e10
        _ = v42
    } else {
        for fi43 := 0; fi43 < 8; fi43++ {
            v29 += 18446744073709551615
            if (-553 < 1) {
                if (!(genFn0() >= 32767)) {
                    var v44 uint16 = 849
                    var v45 uint32 = (v37 << ((genFn1(42, uint32(v4), "a")) & 63))
                    _ = v44
                    _ = v45
                } else {
                    for fi46 := 0; fi46 < 13; fi46++ {
                        var v47 uint16 = uint16(v37)
                        _ = fi46
                        _ = v47
                    }
                    _ = len(v3)
                    fmt.Println(v37)
                }
                if ((v29 + v29) != (v29 << ((genFn1(-18, v37, (("a" + ("foo bar" + "a")) + (((("" + "") + ("foo bar" + "")) + ("foo bar" + "hello")) + "foo bar")))) & 63))) {
                    if v1 {
                        fmt.Println(v28, v1)
                        v4 += (v4 | uint8(v29))
                    } else {
                        var v48 int16 = 1000
                        _ = v48
                    }
                }
                var v49 uint64 = 2
                _ = v49
            }
            _ = fi43
        }
        var v50 float64 = -2.25
        fmt.Println(v28, v50)
        _ = v50
    }
    var v51 float64 = 3.14159
    fmt.Println(v37, v1, v51)
    fmt.Println(v51, v29)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v51
}
