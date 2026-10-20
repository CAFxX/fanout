package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint64 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint32, p1 uint) uint64 {
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 float32 = (-2.25 / 1.0)
        v8 -= v8
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: true, F1: -1, F2: 18446744073709551615}
    var v10 float64 = ((1.5 - 0.1) + 0.0)
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 int64 = -600
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 142
}
func genFn1(p0 uint8) uint32 {
    p0 -= p0
    var v14 float32 = (1e10 * float32(-2.25))
    var v15 int64 = 9223372036854775807
    _ = p0
    _ = v14
    _ = v15
    return uint32(p0)
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    var v3 uint16 = 2
    v4 := make(map[string]string)
    var v5 uint64 = 767
    v6 := make(map[string]bool)
    for fi16 := 0; fi16 < 13; fi16++ {
        v17 := GenS2{F0: 1000, F1: 32767}
        var v18 int8 = -1
        for fi19 := 0; fi19 < 16; fi19++ {
            v20 := v4["a"]
            _ = fi19
            _ = v20
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    var v21 uint16 = (v3 ^ 42)
    v5 = 0
    v4["hello"] = "\n\t"
    v21 = (^(v3))
    v22 := []int{4294967295, 128}
    if true {
        for fi23 := 0; fi23 < 4; fi23++ {
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 int32 = 42
                if ((fi23 * 0) != -9223372036854775808) {
                    if ((v25 | 1) > (v25 ^ v25)) {
                        _ = len(v2)
                    }
                    var v26 uint16 = 0
                    fmt.Println(v25, v5)
                    if true {
                        switch ("" + ("\n\t" + "")) {
                        case "\n\t":
                            if false {
                                v27 := GenS2{F0: 1000, F1: 0}
                                _ = v27
                            }
                        }
                    } else {
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                    }
                    if len(v22) > 0 {
                        v22[0] = fi23
                    }
                    _ = v26
                }
                v5 -= 0
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
        v5 += 42
    } else {
        for fi28 := 0; fi28 < 11; fi28++ {
            v29 := GenS1{F0: (float32(0.0) < 1e10), F1: (fi28 + -640), F2: (v5 - genFn0(0, 42))}
            v21 -= (v21 / ((v21) | 1))
            _ = fi28
            _ = v29
        }
        var v30 int64 = -796
        v31 := GenS1{F0: (-128 <= -128), F1: 127, F2: (^(v5))}
        _ = v30
        _ = v31
    }
    if false {
        v32 := GenS3{F0: 42}
        _ = v32
    } else {
        var v33 uint16 = 2
        var v34 int8 = 92
        _ = v33
        _ = v34
    }
    v35 := GenS4{F0: 42}
    v22 = append(v22, -134)
    var v36 uint16 = (v3 / ((283) | 1))
    _ = len(v22)
    fmt.Println(v21)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
