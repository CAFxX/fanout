package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int; F1 string; F2 uint16 }
type GenS4 struct { F0 int16 }
func genFn0(p0 int32, p1 uint8, p2 int32) uint8 {
    p0 -= p2
    p0 -= -267
    p2 = p0
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 uint64 = 0
    v2 := []int{}
    var v3 float64 = ((3.14159 + (3.14159 / ((0.1 / 1.0) / -2.25))) * float64(float32(3.14159)))
    v4 := make(map[uint32]int)
    var v5 float64 = v3
    var v6 uint32 = 2
    var v7 float64 = v5
    v8 := GenS1{F0: (!(!(!((65535 != 944) || true)))), F1: 169}
    if false {
        if (!false) {
            if len(v2) > 0 {
                v2[0] = -461
            }
            switch "foo bar" {
            case "a":
                fmt.Println(v5)
            case "foo bar":
                var v9 uint32 = uint32(v1)
                _ = v9
            default:
                v3 = -2.25
            }
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            for fi11 := 0; fi11 < 12; fi11++ {
                v6 += (v6 & 65535)
                var v12 int8 = 0
                _ = fi11
                _ = v12
            }
            v6 = (v6 | 43)
            _ = fi10
        }
    } else {
        v1 -= v1
        var v13 float64 = float64(0.1)
        v14 := GenS2{F0: -1, F1: -9223372036854775808}
        _ = v13
        _ = v14
    }
    if true {
        if true {
            v7 += v3
            var v15 int16 = 32767
            v15 = (v15 | (^(v15)))
            _ = v15
        }
        var v16 int16 = 32767
        _ = v16
    } else {
        var v17 string = "foo bar"
        var v18 float64 = ((0.1 + float64(((0.0 * 3.14159) / (1e10 / ((3.14159 * 1e10) + 0.0))))) / 3.14159)
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := GenS1{F0: false, F1: 65535}
        _ = fi19
        _ = v20
    }
    var v21 uint64 = v1
    var v22 uint = 255
    var v23 int16 = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v24 uint8 = 255
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        var v25 int = -316
        for fi26 := 0; fi26 < 19; fi26++ {
            if true {
                var v27 int8 = 1
                v6 = v6
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v27
            } else {
                v4[v6] = (^(v25))
            }
            var v28 int32 = -2147483648
            for fi29 := 0; fi29 < 9; fi29++ {
                if false {
                    var v30 float32 = 0.1
                    _ = v30
                }
                v4[156] = v25
                if (v21 < uint64(v22)) {
                    switch (fi29 - (fi29 ^ (v25 & fi29))) {
                    case -2147483648:
                        if len(v2) > 0 {
                            v2[0] = v25
                        }
                    }
                }
                _ = fi29
            }
            _ = fi26
            _ = v28
        }
        _ = v25
    } else {
        var v31 int64 = 42
        if true {
            v32 := GenS3{F0: -9223372036854775808, F1: "hello", F2: 65535}
            v6 += v6
            _ = v32
        } else {
            for fi33 := 0; fi33 < 4; fi33++ {
                v34 := GenS4{F0: (-(v23))}
                _ = fi33
                _ = v34
            }
        }
        for fi35 := 0; fi35 < 20; fi35++ {
            v23 -= (v23 / ((v23) | 1))
            _ = fi35
        }
        _ = v31
    }
    switch -9223372036854775808 {
    case -303:
        var v36 float32 = (((0.0 * 1e10) * 0.0) / float32(((v7 + -2.25) - (((3.14159 + (float64(float32((1e10 * ((v7 * 3.14159) + (v3 - 1e10))))) - 1.5)) * v7) - (v3 * ((v7 / (float64(3.14159) / v3)) / v7))))))
        _ = v36
    case 32767:
        v2 = append(v2, 127)
    case 32768:
        var v37 string = ((("\n\t" + "hello") + "foo bar") + (("a" + "") + ""))
        _ = v37
    }
    v5 = v7
    v23 += v23
    delete(v4, uint32(v22))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
