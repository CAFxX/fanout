package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 184
    var v3 uint32 = 371
    var v4 uint64 = 1
    var v5 int16 = -520
    var v6 int32 = (v2 >> ((42) & 63))
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint = 473
        if (!false) {
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := GenS1{F0: "hello", F1: 18446744073709551615, F2: 0}
                if true {
                    v11 := GenS1{F0: "hello", F1: v4, F2: v5}
                    var v12 int64 = -554
                    var v13 int8 = -1
                    switch fi9 {
                    case 42:
                        v13 += v13
                    case 624:
                        v8 = v8
                    default:
                        switch (fi9 & int(v5)) {
                        case 42:
                            v14 := v1[42]
                            _ = v14
                        case 1:
                            var v15 uint8 = 109
                            _ = v15
                        default:
                            var v16 uint32 = v3
                            _ = v16
                        }
                    }
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    v17 := v1[fi7]
                    fmt.Println(v4, fi7, fi7)
                    fmt.Println(v5, v4, v17)
                    _ = v17
                }
                _ = fi9
                _ = v10
            }
            if (-9223372036854775808 >= 0) {
                switch int(v2) {
                case -1:
                    v18 := []int{(fi7 / ((fi7) | 1)), (fi7 << ((v8) & 63)), fi7}
                    _ = v18
                case 32768:
                    if true {
                        for fi19 := 0; fi19 < 6; fi19++ {
                            fmt.Println(v3, v3)
                            var v20 float32 = 1.5
                            v21 := GenS1{F0: ((("\n\t" + "hello") + "") + "a"), F1: 0, F2: v5}
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        v2 -= 2147483647
                    }
                }
            } else {
                v22 := []int{-1, (fi7 * (fi7 + (fi7 / ((fi7) | 1)))), -424}
                var v23 int = 127
                _ = v22
                _ = v23
            }
        } else {
            v24 := []int{(fi7 - (-(fi7))), fi7, fi7, (fi7 | (fi7 % ((fi7) | 1)))}
            _ = v24
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v2, v2)
    for fi25 := 0; fi25 < 2; fi25++ {
        for fi26 := 0; fi26 < 19; fi26++ {
            v27 := []int{fi26, int(v3)}
            _ = fi26
            _ = v27
        }
        v2 -= (v2 % ((2147483647) | 1))
        _ = fi25
    }
    fmt.Println(v6, v4)
    if false {
        switch -340 {
        case 977:
            fmt.Println(v2, v5, v5)
        default:
            fmt.Println(v5)
        }
        for fi28 := 0; fi28 < 17; fi28++ {
            var v29 int = -555
            _ = fi28
            _ = v29
        }
    } else {
        v2 = (v2 ^ v6)
        v6 = (v6 / ((v6) | 1))
        switch 306 {
        case 128:
            switch 42 {
            case 4294967295:
                fmt.Println(v3, v2)
            default:
                var v30 uint32 = 791
                _ = v30
            }
        case 0:
            v2 -= v6
        case 981:
            switch "\n\t" {
            case "foo bar":
                v31 := []int{42, 128, 728}
                _ = v31
            case "hello":
                switch ("a" + "foo bar") {
                case "\n\t":
                    v2 += (v2 - 2147483647)
                case "hello":
                    var v32 float32 = (0.0 / 1.0)
                    _ = v32
                default:
                    v33 := []int{-2147483648, 1}
                    _ = v33
                }
            case "a":
                for fi34 := 0; fi34 < 7; fi34++ {
                    var v35 int32 = v2
                    var v36 uint16 = 2
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            }
        default:
            v37 := GenS1{F0: "", F1: v4, F2: v5}
            _ = v37
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
