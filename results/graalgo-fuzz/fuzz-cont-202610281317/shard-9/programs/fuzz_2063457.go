package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint; F1 bool; F2 uint }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]int)
    var v4 uint32 = 65535
    v5 := make(map[int]uint32)
    var v6 int = 0
    v6 = (-(v6))
    v7 := []int{(v6 * (-(v6))), 2147483647, 431, 32768}
    if false {
        switch 9223372036854775807 {
        case 617:
            v8 := GenS1{F0: -155, F1: "a"}
            _ = v8
        default:
            v6 -= (-(v6))
        }
        v4 += (v4 >> ((914) & 63))
    }
    v7 = append(v7, 127)
    var v9 uint16 = 749
    if (!(!true)) {
        v10 := GenS1{F0: 477, F1: ""}
        var v11 float64 = 1e10
        var v12 uint32 = uint32(v9)
        _ = v10
        _ = v11
        _ = v12
    }
    v6 = v6
    var v13 uint16 = (v9 << ((415) & 63))
    if true {
        switch (v6 / (((v6 % (((^(v6))) | 1))) | 1)) {
        case 32768:
            if len(v7) > 0 {
                v7[0] = (v6 / (((v6 * (v6 % ((-30) | 1)))) | 1))
            }
        case 127:
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 bool = true
                switch ("" + "foo bar") {
                case "":
                    if len(v7) > 0 {
                        v7[0] = (^(fi14))
                    }
                case "foo bar":
                    var v16 uint16 = v9
                    _ = v16
                default:
                    v17 := GenS1{F0: -646, F1: "foo bar"}
                    _ = v17
                }
                var v18 int = (v6 & 290)
                _ = fi14
                _ = v15
                _ = v18
            }
        case 9223372036854775807:
            switch (("hello" + "") + ("hello" + "\n\t")) {
            case "\n\t":
                var v19 uint64 = 880
                _ = v19
            default:
                if ((v6 ^ (v6 + v6)) > 2147483647) {
                    v20 := GenS1{F0: 0, F1: "foo bar"}
                    _ = v20
                }
            }
        default:
            for fi21 := 0; fi21 < 1; fi21++ {
                var v22 uint16 = v13
                _ = fi21
                _ = v22
            }
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        v23 := GenS2{F0: 0, F1: (0.0 >= 1e10), F2: 2}
        _ = v23
    } else {
        var v24 float32 = ((((((3.14159 - 3.14159) - 1.5) - 1.5) / 1e10) - 1.5) * 0.0)
        fmt.Println(v13, v6)
        _ = v24
    }
    if (65535 == 2) {
        for fi25 := 0; fi25 < 10; fi25++ {
            fmt.Println(fi25, v4)
            if (4294967295 < (v4 | (v4 * (v4 + v4)))) {
                var v26 uint = 127
                if false {
                    var v27 uint64 = 678
                    _ = len(v7)
                    v28 := GenS2{F0: v26, F1: ((^(v26)) > (v26 & v26)), F2: v26}
                    _ = v27
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 2; fi29++ {
                        if (!false) {
                            v9 -= v13
                            fmt.Println(v4)
                            v30 := GenS2{F0: (v26 ^ v26), F1: (!true), F2: (^(v26))}
                            _ = v30
                        } else {
                            v31 := GenS1{F0: 0, F1: "\n\t"}
                            _ = v31
                        }
                        _ = fi29
                    }
                }
                var v32 int32 = 42
                _ = v26
                _ = v32
            } else {
                var v33 int = v6
                var v34 uint32 = (v4 >> ((65535) & 63))
                var v35 int8 = 0
                _ = v33
                _ = v34
                _ = v35
            }
            var v36 uint64 = 2
            _ = fi25
            _ = v36
        }
    }
    var v37 uint8 = 44
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
