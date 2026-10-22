package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 string; F2 string }
type GenS6 struct { F0 int32; F1 int16; F2 bool }
func genFn0(p0 int64, p1 uint64) uint8 {
    v4 := []int{42, -667, 547, 869}
    var v5 bool = (479 == 42)
    if (!v5) {
        if v5 {
            if (1 < -128) {
                if len(v4) > 0 {
                    v4[0] = 9223372036854775807
                }
                if (true || (42 >= 386)) {
                    v6 := GenS1{F0: -1, F1: 127}
                    fmt.Println(p1, p0, p1)
                    _ = v6
                } else {
                    v5 = true
                    v7 := GenS2{F0: 56}
                    if (("\n\t" < "\n\t") || (0.1 <= 0.1)) {
                        v5 = v5
                        var v8 string = "a"
                        _ = v8
                    }
                    _ = v7
                }
            } else {
                var v9 float32 = (((float32(((float64(3.14159) - 1.5) + 1e10)) * (-2.25 * -2.25)) + (float32(0.1) / 0.1)) + 1e10)
                _ = v9
            }
            var v10 int64 = p0
            var v11 string = "\n\t"
            _ = v10
            _ = v11
        }
        _ = len(v4)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        var v12 int32 = 42
        _ = v12
    }
    _ = len(v4)
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 127
}

func main() {
    var v1 uint32 = 137
    v2 := make(map[string]uint32)
    var v3 uint8 = 255
    for fi13 := 0; fi13 < 20; fi13++ {
        v1 = v1
        fmt.Println(v3)
        for fi14 := 0; fi14 < 6; fi14++ {
            if (!(uint(v3) < uint(v1))) {
                if true {
                    if true {
                        v15 := GenS2{F0: 930}
                        fmt.Println(fi14)
                        fmt.Println(fi13, v3, fi13)
                        _ = v15
                    }
                    var v16 int8 = 100
                    _ = v16
                }
                if true {
                    var v17 int32 = 2147483647
                    if true {
                        if false {
                            var v18 uint16 = 2
                            v19 := []int{(fi14 >> ((636) & 63)), fi14, (fi13 >> ((255) & 63))}
                            v20 := GenS1{F0: 32767, F1: (fi14 / (((fi13 & (fi13 - fi14))) | 1))}
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        v21 := GenS2{F0: 316}
                        fmt.Println(v17)
                        _ = v21
                    } else {
                        if true {
                            v22 := []int{-1}
                            _ = v22
                        } else {
                            v1 -= v1
                            v17 = v17
                        }
                        v17 = -1
                    }
                    _ = v17
                }
                v3 -= genFn0(-1, 0)
            } else {
                var v23 float32 = 3.14159
                _ = v23
            }
            _ = fi14
        }
        _ = fi13
    }
    v24 := []int{int(v3), 42, 2147483647}
    v1 = (v1 * 370)
    fmt.Println(v1)
    if len(v24) > 0 {
        v24[0] = 32768
    }
    v25 := v2["foo bar"]
    v2[(("hello" + ("" + "a")) + "hello")] = (v25 & 1)
    v26 := GenS3{F0: v3, F1: 1000, F2: v3}
    v27 := GenS4{F0: 127}
    for fi28 := 0; fi28 < 7; fi28++ {
        if false {
            v2[("\n\t" + "foo bar")] = v1
            v2[("hello" + "hello")] = v25
            switch (fi28 | (fi28 / ((-1) | 1))) {
            case 775:
                _ = len(v24)
            default:
                fmt.Println(v3, fi28)
            }
        } else {
            v1 += v1
        }
        for fi29 := 0; fi29 < 6; fi29++ {
            if false {
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
                v1 = 0
            }
            _ = fi29
        }
        v30 := GenS5{F0: 0, F1: "foo bar", F2: "hello"}
        _ = fi28
        _ = v30
    }
    v3 -= v3
    if true {
        var v31 int16 = 32767
        var v32 bool = true
        v25 = v25
        _ = v31
        _ = v32
    } else {
        for fi33 := 0; fi33 < 5; fi33++ {
            if (!false) {
                var v34 uint32 = (v25 % ((65535) | 1))
                fmt.Println(fi33, v25)
                v2["foo bar"] = (^(v25))
                _ = v34
            }
            var v35 string = (("a" + "foo bar") + "a")
            v25 -= (v1 | 2)
            _ = fi33
            _ = v35
        }
        if len(v24) > 0 {
            v24[0] = 750
        }
        var v36 int16 = 1000
        _ = v36
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        for fi38 := 0; fi38 < 16; fi38++ {
            var v39 uint32 = v25
            var v40 int64 = 9223372036854775807
            var v41 float64 = 0.0
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
        }
        v42 := GenS6{F0: -18, F1: 42, F2: (true || ("a" == "a"))}
        _ = fi37
        _ = v42
    }
    switch 81 {
    case -9223372036854775808:
        var v43 float64 = 3.14159
        _ = v43
    default:
        v3 = (v3 * (v3 | 91))
    }
    fmt.Println(v3, v25)
    var v44 int32 = 123
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v44
}
