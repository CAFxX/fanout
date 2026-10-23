package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 bool; F1 int }
func genFn0(p0 uint, p1 int64, p2 int) uint64 {
    var v7 uint64 = 725
    var v8 uint32 = 42
    var v9 uint16 = 2
    var v10 int8 = 0
    fmt.Println(v10, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (v7 * v7)
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    v3 := []uint32{4294967295, 586, 65535, 2147483647}
    v4 := []string{((((((("\n\t" + "a") + "") + "a") + "\n\t") + "hello") + "a") + "\n\t")}
    var v5 float64 = 1.5
    var v6 int64 = 1
    if true {
        fmt.Println(v5)
        fmt.Println(v5, v5)
    } else {
        if (false || false) {
            fmt.Println(v6)
            v11 := GenS1{F0: v6}
            _ = v11
        } else {
            var v12 float64 = float64(1e10)
            v13 := GenS1{F0: v6}
            _ = v12
            _ = v13
        }
        v6 += 1
        v6 = v6
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        v6 = v6
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 uint16 = 255
            if (4294967295 != 311) {
                if ((42 != -1) || false) {
                    v1[539] = fi15
                    v6 += v6
                    v3 = append(v3, 183)
                    v17 := GenS1{F0: 343}
                    _ = v17
                }
            }
            fmt.Println(v16, fi14, fi15)
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v18 := GenS2{F0: 1, F1: 4294967295}
    v19 := GenS1{F0: (-(v6))}
    v20 := v2[-1]
    v5 += v5
    for fi21 := 0; fi21 < 7; fi21++ {
        for fi22 := 0; fi22 < 5; fi22++ {
            v5 -= v5
            _ = fi22
        }
        _ = fi21
    }
    v6 -= 9223372036854775807
    if (false || (0 != 2147483648)) {
        if true {
            v23 := GenS3{F0: 42}
            for fi24 := 0; fi24 < 6; fi24++ {
                v6 += (v6 << ((0) & 63))
                var v25 int64 = v6
                _ = fi24
                _ = v25
            }
            _ = v23
        }
    }
    v26 := GenS2{F0: 255, F1: 1}
    var v27 int = v20
    v28 := v2[v27]
    for fi29 := 0; fi29 < 2; fi29++ {
        if (v5 < (v5 - 0.1)) {
            v27 -= v20
            v27 += (-(v28))
            v30 := GenS4{F0: 365, F1: 1}
            _ = v30
        } else {
            if true {
                if false {
                    for fi31 := 0; fi31 < 20; fi31++ {
                        var v32 int = (-(fi31))
                        var v33 int16 = -32768
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    switch (("" + "\n\t") + ("a" + ("hello" + "foo bar"))) {
                    case "":
                        v34 := GenS3{F0: 255}
                        _ = v34
                    case "\n\t":
                        var v35 float64 = 1e10
                        _ = v35
                    default:
                        v36 := GenS2{F0: 889, F1: 95}
                        _ = v36
                    }
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        for fi37 := 0; fi37 < 19; fi37++ {
                            var v38 float32 = 1e10
                            v39 := GenS5{F0: true, F1: 32767}
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                    case "hello":
                        v20 += fi29
                    default:
                        var v40 int = v28
                        _ = v40
                    }
                    for fi41 := 0; fi41 < 8; fi41++ {
                        var v42 uint = 1
                        v6 += 9223372036854775807
                        _ = fi41
                        _ = v42
                    }
                }
            }
        }
        _ = fi29
    }
    v1[42] = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
