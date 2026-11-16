package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int32 }
func genFn0(p0 int64, p1 int) int {
    p0 -= p0
    v4 := []int{}
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = v4
    return (p1 >> ((255) & 63))
}

func main() {
    v1 := []uint32{2147483648, 572}
    v2 := []uint32{2, 42, 1, 2}
    v3 := make(map[uint32]uint32)
    var v5 string = ("\n\t" + "foo bar")
    if false {
        v5 = ((v5 + "") + "")
    } else {
        var v6 bool = (!(1 <= 381))
        v7 := v3[1]
        v5 = (v5 + "hello")
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: 68, F1: 32767, F2: 2147483647}
        v10 := GenS1{F0: -1, F1: -1, F2: 1}
        v5 = "hello"
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        if (false && (589 > 267)) {
            if true {
                for fi12 := 0; fi12 < 16; fi12++ {
                    v3[42] = 42
                    switch ((("\n\t" + "\n\t") + "a") + ("foo bar" + (v5 + "a"))) {
                    case "a":
                        v13 := GenS1{F0: -1, F1: 0, F2: 467}
                        _ = v13
                    case "foo bar":
                        var v14 int8 = 1
                        _ = v14
                    case "\n\t":
                        v5 = v5
                    default:
                        v5 = "\n\t"
                    }
                    v15 := GenS1{F0: 0, F1: 1000, F2: 2147483647}
                    _ = fi12
                    _ = v15
                }
                fmt.Println(fi11, fi11)
            }
            var v16 uint = 255
            v17 := GenS1{F0: -1, F1: 1, F2: 42}
            _ = v16
            _ = v17
        }
        var v18 uint8 = 127
        _ = fi11
        _ = v18
    }
    switch genFn0(-9223372036854775808, 1) {
    case 725:
        var v19 float64 = (-2.25 - 1e10)
        _ = v19
    }
    v20 := GenS1{F0: -100, F1: 638, F2: -982}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        v3[65535] = 305
        for fi22 := 0; fi22 < 4; fi22++ {
            if (false || false) {
                v5 = v5
                v23 := GenS1{F0: -1, F1: 1000, F2: -1}
                v5 = "foo bar"
                _ = v23
            }
            v5 = ""
            v24 := v3[2147483648]
            var v25 int = (fi21 / ((fi21) | 1))
            v3[v24] = v24
            v25 = -687
            var v26 int32 = -299
            _ = fi22
            _ = v24
            _ = v25
            _ = v26
        }
        var v27 bool = (("hello" > "") || true)
        _ = fi21
        _ = v27
    }
    if true {
        v5 = v5
        fmt.Println(v5)
    }
    _ = len(v3)
    var v28 int16 = 1000
    v29 := GenS1{F0: -1, F1: v28, F2: -2147483648}
    if false {
        if (0 >= 1) {
            var v30 string = v5
            fmt.Println(v30)
            _ = v30
        }
        v31 := GenS1{F0: 127, F1: 1000, F2: -172}
        _ = len(v1)
        _ = v31
    } else {
        var v32 uint32 = 806
        var v33 uint = 1
        var v34 int = 32767
        _ = v32
        _ = v33
        _ = v34
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v35 float64 = float64(3.14159)
    v36 := GenS1{F0: 42, F1: -485, F2: -161}
    for fi37 := 0; fi37 < 4; fi37++ {
        if (-682 != genFn0(291, (-(fi37)))) {
            for fi38 := 0; fi38 < 18; fi38++ {
                var v39 uint8 = 212
                _ = fi38
                _ = v39
            }
            fmt.Println(v28, v5)
            if len(v1) > 0 {
                v1[0] = 1
            }
        }
        _ = fi37
    }
    var v40 bool = (float32(-2.25) != ((1.5 / (float32((((v35 * v35) / v35) - -2.25)) + (float32(v35) * (-2.25 - (0.1 - 3.14159))))) * 0.0))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v40
    _ = v5
}
