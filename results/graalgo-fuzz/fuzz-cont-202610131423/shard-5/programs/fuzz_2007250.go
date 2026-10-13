package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 bool; F1 int32; F2 int8 }
func genFn0(p0 int16, p1 int, p2 int16) uint64 {
    v7 := []int{p1, p1}
    p0 -= (p2 - 32767)
    v8 := GenS1{F0: 1, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 112
}
func genFn1(p0 uint8, p1 bool, p2 bool) uint64 {
    v9 := []int{9223372036854775807}
    for fi10 := 0; fi10 < 6; fi10++ {
        p0 = 85
        p0 += 128
        var v11 int64 = 678
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: (p0 - (^(p0))), F1: -9223372036854775808}
    if p2 {
        var v13 float64 = -2.25
        _ = v13
    } else {
        if len(v9) > 0 {
            v9[0] = 42
        }
        var v14 float32 = 0.1
        _ = v14
    }
    var v15 uint = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v15
    _ = v9
    return 42
}

func main() {
    var v1 float64 = 1.5
    v2 := []uint32{616, 2147483647, 42, 2147483647}
    v3 := make(map[string]uint32)
    var v4 uint16 = 42
    var v5 uint = 4294967295
    var v6 uint32 = 183
    for fi16 := 0; fi16 < 14; fi16++ {
        v6 -= v6
        _ = fi16
    }
    var v17 string = "a"
    if (2 >= (v4 + 385)) {
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 int = (fi18 & 0)
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 int64 = 0
            _ = fi20
            _ = v21
        }
    } else {
        for fi22 := 0; fi22 < 10; fi22++ {
            if true {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v23 int32 = 42
                if false {
                    var v24 float64 = (((v1 - v1) + 0.1) / v1)
                    _ = v24
                }
                _ = v23
            } else {
                v17 = ("a" + "a")
            }
            _ = fi22
        }
    }
    if false {
        v4 += 0
        for fi25 := 0; fi25 < 8; fi25++ {
            _ = len(v3)
            v26 := GenS1{F0: 128, F1: 42}
            _ = fi25
            _ = v26
        }
        var v27 uint = (v5 / ((4294967295) | 1))
        _ = v27
    }
    v6 -= 901
    var v28 uint16 = 255
    v5 = 65535
    for fi29 := 0; fi29 < 17; fi29++ {
        var v30 uint = (v5 | (^(v5)))
        switch 32768 {
        case 4294967295:
            var v31 uint64 = 1
            _ = v31
        case -2147483648:
            v6 = uint32(v4)
        case -426:
            var v32 uint32 = 0
            _ = v32
        default:
            for fi33 := 0; fi33 < 14; fi33++ {
                if (32767 > fi33) {
                    v34 := GenS2{F0: (v6 == (v6 - v6)), F1: 0, F2: 0}
                    v2 = append(v2, (^(v6)))
                    var v35 uint64 = genFn1(128, true, true)
                    _ = v34
                    _ = v35
                } else {
                    fmt.Println(v4, fi33, v5)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                v36 := GenS1{F0: 127, F1: 1}
                var v37 uint32 = 1
                if true {
                    v1 = (((v1 + v1) * -2.25) + ((v1 * float64(-2.25)) * float64(0.1)))
                    v6 -= v37
                }
                _ = fi33
                _ = v36
                _ = v37
            }
        }
        v5 = uint(v4)
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
