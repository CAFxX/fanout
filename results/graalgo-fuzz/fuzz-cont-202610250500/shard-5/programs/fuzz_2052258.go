package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint; F1 int32; F2 bool }
type GenS5 struct { F0 uint32; F1 string }
func genFn0(p0 int16, p1 uint16, p2 bool) int {
    var v4 uint8 = 128
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 uint8 = 2
        fmt.Println(v6)
        _ = fi5
        _ = v6
    }
    v4 -= (^(v4))
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 683, F1: 883}
        p1 += (p1 >> ((340) & 63))
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 128
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 518
    var v3 uint32 = 895
    v9 := GenS2{F0: 127}
    fmt.Println(v3, v2)
    var v10 float32 = 3.14159
    v10 = v10
    v1[(("" + "a") + "a")] = (0.1 > float32(-2.25))
    v11 := []int{-58, -9223372036854775808}
    var v12 int = -158
    for fi13 := 0; fi13 < 12; fi13++ {
        switch "\n\t" {
        case "a":
            switch (fi13 * fi13) {
            case 127:
                v12 = v12
            case 9223372036854775807:
                var v14 float64 = 1.5
                _ = v14
            case 4294967295:
                v15 := GenS1{F0: 45, F1: 254}
                _ = v15
            default:
                if (255 != 2) {
                    if false {
                        v16 := GenS1{F0: 4294967295, F1: 4294967295}
                        v17 := GenS2{F0: 76}
                        _ = v16
                        _ = v17
                    }
                    v18 := GenS2{F0: 127}
                    v19 := GenS3{F0: "\n\t", F1: ("hello" + ("\n\t" + "a"))}
                    _ = v18
                    _ = v19
                } else {
                    fmt.Println(v2, fi13)
                    for fi20 := 0; fi20 < 13; fi20++ {
                        v12 = (fi13 | (fi20 >> ((0) & 63)))
                        if (!(65535 > uint16(v3))) {
                            var v21 int64 = -1
                            _ = v21
                        }
                        var v22 int32 = -1
                        v23 := GenS4{F0: 445, F1: -499, F2: (-2.25 >= -2.25)}
                        v3 = 2147483648
                        v24 := GenS5{F0: v3, F1: "foo bar"}
                        _ = fi20
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                }
            }
        }
        switch ("foo bar" + "hello") {
        case "foo bar":
            var v25 bool = true
            _ = v25
        default:
            v2 = -1
        }
        _ = fi13
    }
    delete(v1, "\n\t")
    switch v12 {
    case -2147483648:
        v10 = (v10 * v10)
    case 525:
        if (true && true) {
            fmt.Println(v10, v3)
            for fi26 := 0; fi26 < 17; fi26++ {
                var v27 int32 = 343
                v10 = v10
                _ = fi26
                _ = v27
            }
            if (true || (!(32767 != 560))) {
                switch ("hello" + "\n\t") {
                case "\n\t":
                    var v28 uint16 = 32767
                    _ = v28
                default:
                    v11 = append(v11, v12)
                }
                fmt.Println(v10, v12, v12)
            } else {
                for fi29 := 0; fi29 < 18; fi29++ {
                    v12 -= (v12 ^ (^(v12)))
                    _ = fi29
                }
                v30 := GenS3{F0: (("" + "a") + ("hello" + "hello")), F1: "hello"}
                fmt.Println(v10, v3)
                _ = v30
            }
        }
    default:
        var v31 uint32 = (v3 ^ v3)
        _ = v31
    }
    if len(v11) > 0 {
        v11[0] = (v12 % ((v12) | 1))
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    if (true && true) {
        if (4294967295 == 908) {
            v2 = v2
            var v32 float64 = 0.0
            v32 = v32
            _ = v32
        }
        v12 += (^(v12))
        var v33 int = 32767
        _ = v33
    } else {
        for fi34 := 0; fi34 < 2; fi34++ {
            if (1e10 < (0.1 + (0.0 + 3.14159))) {
                var v35 float32 = v10
                if false {
                    v36 := v1[("\n\t" + ("" + "\n\t"))]
                    v37 := GenS1{F0: 4294967295, F1: 255}
                    v35 = ((0.0 + 0.0) - v10)
                    _ = v36
                    _ = v37
                }
                for fi38 := 0; fi38 < 7; fi38++ {
                    var v39 string = "a"
                    _ = len(v1)
                    var v40 int16 = 964
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                _ = v35
            } else {
                switch "a" {
                case "foo bar":
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                }
                var v41 string = "foo bar"
                var v42 int8 = 127
                _ = v41
                _ = v42
            }
            v10 = float32((((1.5 + 1.5) + ((float64(1.5) - (3.14159 / 1.0)) * float64(v10))) / 1.5))
            v3 += v3
            _ = fi34
        }
        v12 = v12
        var v43 int32 = -870
        _ = v43
    }
    fmt.Println(v2, v2, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
