package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint64; F2 bool }
type GenS3 struct { F0 uint8; F1 uint; F2 uint }
type GenS4 struct { F0 int32; F1 int16; F2 uint64 }
func genFn0() bool {
    if (!false) {
        var v5 int64 = 278
        var v6 uint = 2
        if (0 != 127) {
            v7 := []int{375, 0, -131, -431}
            _ = len(v7)
            if ((2 != 42) && (!true)) {
                for fi8 := 0; fi8 < 4; fi8++ {
                    v9 := GenS1{F0: 931}
                    _ = fi8
                    _ = v9
                }
                fmt.Println(v5, v6)
            } else {
                var v10 uint64 = 18446744073709551615
                for fi11 := 0; fi11 < 14; fi11++ {
                    v5 = -358
                    _ = fi11
                }
                for fi12 := 0; fi12 < 8; fi12++ {
                    var v13 int32 = 0
                    _ = fi12
                    _ = v13
                }
                _ = v10
            }
            _ = v7
        } else {
            var v14 uint8 = 128
            for fi15 := 0; fi15 < 14; fi15++ {
                v14 -= (v14 * v14)
                v14 += (v14 * (v14 / (((v14 | (v14 | v14))) | 1)))
                _ = fi15
            }
            v16 := []int{-1, -9223372036854775808}
            _ = v14
            _ = v16
        }
        _ = v5
        _ = v6
    } else {
        if true {
            var v17 float64 = (float64(0.0) + (1.5 * 0.1))
            _ = v17
        }
        switch "hello" {
        case "hello":
            if (false || true) {
                var v18 float64 = 1.5
                var v19 bool = ((!(((("a" + "a") + "hello") + "foo bar") == ("a" + ("\n\t" + "foo bar")))) || (-1 < 0))
                var v20 float32 = 0.1
                _ = v18
                _ = v19
                _ = v20
            } else {
                var v21 uint64 = 432
                _ = v21
            }
        case "a":
            for fi22 := 0; fi22 < 19; fi22++ {
                fmt.Println(fi22)
                for fi23 := 0; fi23 < 9; fi23++ {
                    var v24 int64 = 0
                    v25 := []int{(-(fi22)), fi22}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                var v26 int64 = 1
                _ = fi22
                _ = v26
            }
        default:
            var v27 float32 = -2.25
            _ = v27
        }
        var v28 int8 = -8
        _ = v28
    }
    var v29 uint16 = 42
    _ = v29
    return (true || true)
}

func main() {
    var v1 uint64 = 460
    var v2 int8 = -128
    var v3 uint = 18446744073709551615
    v4 := make(map[string]string)
    fmt.Println(v3, v3, v1)
    if true {
        v30 := GenS2{F0: 0, F1: (v1 % ((101) | 1)), F2: false}
        v3 = v3
        switch "\n\t" {
        case "hello":
            fmt.Println(v3)
        case "":
            var v31 int32 = -1
            _ = v31
        default:
            for fi32 := 0; fi32 < 17; fi32++ {
                v33 := []int{-2147483648}
                if true {
                    v33 = append(v33, fi32)
                } else {
                    var v34 int64 = -1
                    _ = v34
                }
                _ = fi32
                _ = v33
            }
        }
        _ = v30
    }
    var v35 bool = false
    var v36 int64 = -992
    fmt.Println(v3)
    if (1 > 0) {
        v37 := GenS3{F0: 2, F1: (^(v3)), F2: 581}
        v38 := []int{32768, 42, -1}
        var v39 bool = v35
        _ = v37
        _ = v38
        _ = v39
    }
    if (32767 > 1) {
        if v35 {
            for fi40 := 0; fi40 < 12; fi40++ {
                if v35 {
                    if ((0.1 != 1.5) && (!("foo bar" >= "\n\t"))) {
                        var v41 float64 = 3.14159
                        _ = v41
                    }
                    var v42 int32 = -1
                    _ = v42
                } else {
                    var v43 float32 = 1e10
                    v2 += 42
                    v44 := []int{(fi40 ^ fi40), (^(fi40)), (fi40 * (fi40 - fi40)), 32767}
                    _ = v43
                    _ = v44
                }
                var v45 uint8 = 103
                v35 = (!(("foo bar" + "") <= "hello"))
                _ = fi40
                _ = v45
            }
            v46 := GenS3{F0: uint8(v1), F1: v3, F2: (v3 / ((430) | 1))}
            fmt.Println(v35, v1, v36)
            _ = v46
        }
        var v47 int16 = 0
        for fi48 := 0; fi48 < 11; fi48++ {
            v3 = (^(v3))
            _ = len(v4)
            v49 := GenS1{F0: (v1 * v1)}
            _ = fi48
            _ = v49
        }
        _ = v47
    }
    var v50 int16 = -32768
    for fi51 := 0; fi51 < 8; fi51++ {
        v52 := GenS1{F0: v1}
        v53 := GenS4{F0: -1, F1: (v50 | (v50 << ((360) & 63))), F2: (v1 * 18446744073709551615)}
        v35 = v35
        _ = fi51
        _ = v52
        _ = v53
    }
    for fi54 := 0; fi54 < 9; fi54++ {
        v55 := GenS4{F0: 42, F1: (v50 & -937), F2: (v1 / ((v1) | 1))}
        v4["hello"] = "a"
        _ = fi54
        _ = v55
    }
    var v56 float64 = (-2.25 + (0.1 - (1e10 - ((1.5 + (((3.14159 / 1.5) / ((-2.25 / 1.0) - 0.0)) - 3.14159)) - float64((1e10 / 3.14159))))))
    switch "\n\t" {
    case "":
        if true {
            var v57 int16 = 242
            var v58 float32 = ((3.14159 / (0.1 / (-2.25 - (0.0 * -2.25)))) + (-2.25 + -2.25))
            _ = v57
            _ = v58
        }
    case "\n\t":
        switch (("hello" + ("\n\t" + "a")) + "") {
        case "":
            v36 += 544
        case "a":
            fmt.Println(v36, v36, v50)
        case "foo bar":
            fmt.Println(v50, v50)
        }
    default:
        v4["hello"] = "foo bar"
    }
    v35 = (!(!v35))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v50:", v50)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v50
    _ = v56
}
