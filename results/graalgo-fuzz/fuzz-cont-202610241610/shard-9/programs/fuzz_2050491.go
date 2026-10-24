package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 bool; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint64, p1 uint, p2 int64) int64 {
    var v7 int16 = 620
    var v8 float64 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint16) string {
    v9 := GenS1{F0: true, F1: 1}
    if (1e10 < ((-2.25 * 1.5) / (1e10 / ((3.14159 - ((1.5 + -2.25) * (1.5 + (-2.25 - (1.5 - (1.5 / 1.0)))))) / ((0.1 - 0.0) - 1e10))))) {
        var v10 float32 = (1.5 / 3.14159)
        _ = v10
    } else {
        var v11 int = 42
        p0 = (^(p0))
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        switch fi12 {
        case 32768:
            switch "hello" {
            case "a":
                v13 := GenS2{F0: 1, F1: "a"}
                _ = v13
            case "":
                v14 := []int{(-(fi12)), (fi12 * (-(fi12))), fi12}
                _ = v14
            default:
                for fi15 := 0; fi15 < 3; fi15++ {
                    v16 := []int{32768, fi15}
                    if len(v16) > 0 {
                        v16[0] = fi15
                    }
                    if len(v16) > 0 {
                        v16[0] = (fi12 << ((0) & 63))
                    }
                    _ = fi15
                    _ = v16
                }
            }
        case 1:
            switch (fi12 - fi12) {
            case -9223372036854775808:
                v17 := []int{(^(fi12)), (fi12 ^ fi12)}
                _ = v17
            }
        default:
            for fi18 := 0; fi18 < 19; fi18++ {
                p0 = p0
                _ = fi18
            }
        }
        p0 -= (p0 * 609)
        _ = fi12
    }
    v19 := []int{-421, -805, -1}
    var v20 uint8 = 255
    _ = p0
    _ = v19
    _ = v20
    _ = v9
    return ("a" + ("hello" + ("hello" + ("\n\t" + ("hello" + "foo bar")))))
}

func main() {
    var v1 int8 = 100
    v2 := make(map[int]bool)
    v3 := []uint32{38, 324, 2147483647}
    var v4 int64 = 42
    var v5 float64 = -2.25
    var v6 int8 = v1
    var v21 uint16 = 37
    var v22 uint32 = 1
    v23 := GenS2{F0: 255, F1: ("a" + "")}
    var v24 float32 = float32(v5)
    _ = len(v3)
    for fi25 := 0; fi25 < 9; fi25++ {
        v5 += 1.5
        v26 := GenS3{F0: false, F1: v22, F2: 1}
        v27 := GenS2{F0: 545, F1: "foo bar"}
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        for fi29 := 0; fi29 < 5; fi29++ {
            v30 := GenS4{F0: v21}
            v31 := GenS4{F0: 0}
            switch (fi28 << ((71) & 63)) {
            case 9223372036854775807:
                v32 := GenS2{F0: 439, F1: "\n\t"}
                _ = v32
            default:
                var v33 int32 = 1
                _ = v33
            }
            _ = fi29
            _ = v30
            _ = v31
        }
        v3 = append(v3, v22)
        if (true && false) {
            v22 += 0
            v34 := GenS2{F0: 0, F1: ("\n\t" + ("foo bar" + ("" + "")))}
            v35 := GenS4{F0: v21}
            _ = v34
            _ = v35
        }
        _ = fi28
    }
    v36 := GenS4{F0: v21}
    var v37 float64 = (v5 + 0.1)
    var v38 int16 = -909
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
