package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint64 }
func genFn0() int {
    for fi6 := 0; fi6 < 8; fi6++ {
        if (true && false) {
            if (!(false || ("\n\t" > ""))) {
                v7 := []int{128, fi6, (fi6 / ((fi6) | 1))}
                _ = v7
            }
            if true {
                fmt.Println(fi6)
            }
        } else {
            var v8 uint16 = 1
            _ = v8
        }
        var v9 float32 = 1e10
        _ = fi6
        _ = v9
    }
    if (221 < 255) {
        var v10 uint8 = 127
        v10 = v10
        _ = v10
    }
    v11 := GenS1{F0: 42, F1: 1000}
    _ = v11
    return 32767
}
func genFn1() uint {
    var v12 int16 = 0
    fmt.Println(v12)
    v13 := GenS1{F0: 59, F1: v12}
    switch "" {
    case "hello":
        for fi14 := 0; fi14 < 14; fi14++ {
            for fi15 := 0; fi15 < 5; fi15++ {
                v16 := GenS1{F0: 42, F1: (v12 / ((-1) | 1))}
                if (0 < 1) {
                    if true {
                        v17 := []int{fi14, (fi14 & fi14), (fi14 - fi15)}
                        v12 = v12
                        var v18 float64 = 0.1
                        _ = v17
                        _ = v18
                    }
                } else {
                    if (-2.25 < -2.25) {
                        v19 := GenS2{F0: 255, F1: 18446744073709551615}
                        _ = v19
                    }
                    v20 := []int{fi15, fi14, (-(fi15))}
                    _ = v20
                }
                v21 := GenS3{F0: 1, F1: -115, F2: 2}
                v12 -= v12
                var v22 uint32 = 2147483647
                var v23 string = ("foo bar" + ("a" + "a"))
                var v24 float64 = 1e10
                v25 := []int{9223372036854775807, (fi15 << ((0) & 63)), fi14, fi14}
                _ = fi15
                _ = v16
                _ = v21
                _ = v22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi14
        }
    case "a":
        v26 := []int{}
        _ = v26
    }
    _ = v12
    _ = v13
    return 707
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{4294967295, 0}
    v3 := make(map[int]bool)
    v4 := make(map[string]bool)
    v5 := make(map[uint32]uint32)
    v27 := GenS2{F0: 2, F1: 0}
    switch 32767 {
    case -742:
        if (false || (("\n\t" >= ("a" + "\n\t")) || ((!false) || (1 < 4294967295)))) {
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
            for fi28 := 0; fi28 < 4; fi28++ {
                var v29 float64 = 0.0
                for fi30 := 0; fi30 < 19; fi30++ {
                    v29 = v29
                    _ = fi30
                }
                v31 := GenS3{F0: 2, F1: 1, F2: 4294967295}
                _ = fi28
                _ = v29
                _ = v31
            }
            var v32 int64 = -350
            _ = v32
        } else {
            delete(v5, 165)
            var v33 float64 = ((0.0 / 1e10) - 1e10)
            _ = v33
        }
    case 2147483647:
        var v34 int = 1
        _ = v34
    case -9223372036854775808:
        v35 := GenS2{F0: 255, F1: 1}
        _ = v35
    }
    var v36 int8 = -1
    if false {
        var v37 int64 = 1
        _ = v37
    }
    v36 += 1
    fmt.Println(v36)
    var v38 int = 0
    v38 = (-(v38))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v36
    _ = v38
    _ = v4
    _ = v5
}
