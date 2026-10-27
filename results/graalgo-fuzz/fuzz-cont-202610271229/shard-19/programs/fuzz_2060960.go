package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint16; F2 int32 }
type GenS4 struct { F0 int32; F1 uint }
func genFn0() uint8 {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int32 = -901
        var v9 uint16 = 2
        v8 += v8
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint64 = 1
    _ = v10
    return 244
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    var v4 string = "a"
    var v5 float32 = (1e10 / (-2.25 + (((0.0 * (1.5 - (0.1 / float32(1e10)))) * 1.5) - 1.5)))
    var v6 uint64 = 0
    v5 -= (v5 - (float32(1e10) * v5))
    v5 -= ((1e10 * v5) / (float32(-2.25) - v5))
    v5 = 1.5
    v6 += (v6 % ((v6) | 1))
    v6 += (v6 + 4294967295)
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS1{F0: 42}
        switch v4 {
        case "\n\t":
            fmt.Println(v5, v6)
        case "":
            v13 := GenS2{F0: (fi11 % ((fi11) | 1))}
            _ = v13
        case "a":
            if false {
                if ((-(fi11)) > 1) {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        if (127 > 42) {
                            v15 := GenS3{F0: -1, F1: 65535, F2: -1}
                            switch v4 {
                            case "a":
                                v16 := []int{(fi11 & (-(fi11))), 127, 1, (fi14 ^ (fi11 << ((42) & 63)))}
                                _ = v16
                            case "foo bar":
                                v17 := []int{}
                                _ = v17
                            case "":
                                v18 := v1[0]
                                _ = v18
                            default:
                                var v19 int8 = 100
                                _ = v19
                            }
                            fmt.Println(v4, fi11, v6)
                            var v20 int8 = -2
                            _ = v15
                            _ = v20
                        } else {
                            _ = len(v1)
                        }
                        _ = fi14
                    }
                } else {
                    v4 = v4
                    v21 := v3[2]
                    v1[592] = 2147483647
                    _ = v21
                }
                v4 = ((v4 + v4) + (((v4 + ("" + v4)) + (("hello" + v4) + v4)) + v4))
                v5 += v5
            }
        }
        _ = fi11
        _ = v12
    }
    v6 = v6
    for fi22 := 0; fi22 < 5; fi22++ {
        v6 += (v6 % (((v6 << ((4294967295) & 63))) | 1))
        _ = fi22
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := []int{(fi23 - fi23), 1, fi23, 9223372036854775807}
        switch fi23 {
        case -7:
            var v25 uint16 = 312
            _ = v25
        case 127:
            v26 := GenS4{F0: -883, F1: 697}
            _ = v26
        case 4294967295:
            v6 += v6
        }
        _ = fi23
        _ = v24
    }
    var v27 int16 = 1
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 float32 = 1e10
        _ = fi28
        _ = v29
    }
    v4 = ("hello" + "")
    var v30 uint64 = (v6 % ((4294967295) | 1))
    for fi31 := 0; fi31 < 7; fi31++ {
        for fi32 := 0; fi32 < 4; fi32++ {
            var v33 int8 = 127
            var v34 int16 = v27
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi31
    }
    var v35 int = 9223372036854775807
    if true {
        for fi36 := 0; fi36 < 9; fi36++ {
            v5 += float32(((1e10 - 0.1) - (((0.1 + 1.5) - (0.1 - (-2.25 * (((1.5 * (1.5 / 1.5)) / float64(v5)) / (1.5 + 3.14159))))) / (-2.25 + 1.5))))
            v37 := []int{v35, (v35 ^ (fi36 >> ((1000) & 63))), fi36}
            _ = fi36
            _ = v37
        }
        if true {
            v2[uint32(v30)] = 65535
            var v38 int8 = 1
            _ = v38
        }
    } else {
        for fi39 := 0; fi39 < 20; fi39++ {
            v40 := GenS2{F0: fi39}
            var v41 int16 = (^(v27))
            var v42 uint8 = 1
            _ = fi39
            _ = v40
            _ = v41
            _ = v42
        }
        var v43 uint8 = 155
        var v44 uint64 = v30
        _ = v43
        _ = v44
    }
    v27 += (v27 ^ 32767)
    if false {
        var v45 int16 = v27
        fmt.Println(v45)
        var v46 int = v35
        _ = v45
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
