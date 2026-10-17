package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 string; F1 uint32; F2 int16 }
func genFn0(p0 uint32) uint {
    fmt.Println(p0)
    v4 := []int{0, 4294967295}
    _ = p0
    _ = v4
    return 253
}

func main() {
    var v1 float32 = float32(0.1)
    v2 := make(map[uint32]int)
    var v3 int8 = 42
    var v5 bool = true
    v6 := GenS1{F0: 0, F1: 1}
    v5 = v5
    if true {
        if (true && ((!((v1 - v1) == 1.5)) || (!(!(65535 < 742))))) {
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 int = fi7
                var v9 float64 = (3.14159 - (float64(1e10) - 0.1))
                v9 += float64(float32((3.14159 - ((v9 * (((v9 - float64(-2.25)) * float64(0.0)) / (v9 + v9))) / v9))))
                _ = fi7
                _ = v8
                _ = v9
            }
            v2[2147483648] = 459
        } else {
            v5 = (!v5)
        }
        fmt.Println(v3, v3)
    } else {
        var v10 int32 = -67
        var v11 int = 32768
        _ = v10
        _ = v11
    }
    v3 += (v3 << ((42) & 63))
    for fi12 := 0; fi12 < 7; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            switch (fi12 + fi12) {
            case 127:
                switch ("a" + "hello") {
                case "\n\t":
                    if v5 {
                        if v5 {
                            if true {
                                var v14 uint8 = 2
                                var v15 int64 = -300
                                v16 := []int{-9223372036854775808, (fi13 ^ (-(fi13))), fi12}
                                _ = v14
                                _ = v15
                                _ = v16
                            } else {
                                v17 := []int{fi12, fi13, fi12, -355}
                                _ = v17
                            }
                            var v18 string = ""
                            _ = v18
                        } else {
                            fmt.Println(v3, fi12)
                        }
                        switch fi12 {
                        case -9223372036854775808:
                            v19 := v2[1]
                            fmt.Println(v1)
                            if v5 {
                                var v20 float64 = -2.25
                                _ = v20
                            } else {
                                fmt.Println(v19, v1)
                                v3 = (v3 | v3)
                            }
                            _ = v19
                        case 128:
                            switch "\n\t" {
                            case "a":
                                var v21 int = (fi13 / ((fi12) | 1))
                                _ = v21
                            case "foo bar":
                                var v22 int64 = -73
                                _ = v22
                            }
                        case 9223372036854775807:
                            v23 := GenS1{F0: 65535, F1: 150}
                            _ = v23
                        default:
                            var v24 float64 = 0.1
                            _ = v24
                        }
                        v25 := []int{}
                        _ = v25
                    } else {
                        v26 := GenS2{F0: "foo bar", F1: 2147483647, F2: 1000}
                        var v27 int32 = 460
                        if true {
                            v28 := []int{-128}
                            v29 := GenS2{F0: "foo bar", F1: 4294967295, F2: 1000}
                            _ = v28
                            _ = v29
                        } else {
                            v5 = true
                        }
                        _ = v26
                        _ = v27
                    }
                case "foo bar":
                    var v30 int8 = (v3 / ((v3) | 1))
                    _ = v30
                case "a":
                    v31 := GenS2{F0: "a", F1: 4294967295, F2: -513}
                    _ = v31
                default:
                    if false {
                        v3 += (v3 | -32)
                        delete(v2, 2147483648)
                        var v32 int16 = 1
                        _ = v32
                    }
                }
            case 399:
                var v33 int16 = 42
                _ = v33
            }
            _ = fi13
        }
        _ = fi12
    }
    v5 = ((1.5 * 0.0) == 1e10)
    if (100 == v3) {
        var v34 int64 = -1
        switch -495 {
        case 9223372036854775807:
            var v35 int32 = 42
            _ = v35
        case 4294967295:
            fmt.Println(v5, v3, v1)
        }
        for fi36 := 0; fi36 < 15; fi36++ {
            v1 -= ((v1 / (v1 * float32((1e10 / 3.14159)))) + 0.0)
            _ = fi36
        }
        _ = v34
    }
    var v37 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v5
    _ = v6
}
