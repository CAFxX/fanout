package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int8 }
type GenS2 struct { F0 int8; F1 uint }
func genFn0(p0 string, p1 int64) uint32 {
    v4 := GenS1{F0: 0, F1: "hello", F2: 42}
    var v5 int8 = 86
    var v6 int16 = -32768
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 581
}

func main() {
    v1 := make(map[int]int)
    var v2 int8 = 0
    v3 := make(map[string]int)
    v1[-2147483648] = 9223372036854775807
    v7 := GenS1{F0: 624, F1: "\n\t", F2: (^(v2))}
    if false {
        var v8 string = ("hello" + ("a" + "\n\t"))
        _ = v8
    } else {
        switch ("hello" + "foo bar") {
        case "a":
            var v9 float32 = ((float32(-2.25) * (0.0 + -2.25)) / 1.5)
            _ = v9
        case "foo bar":
            v1[-2147483648] = 0
        default:
            var v10 int32 = 1
            _ = v10
        }
        var v11 int = 32767
        fmt.Println(v11, v2)
        _ = v11
    }
    v12 := v1[854]
    fmt.Println(v12)
    v2 += v2
    v13 := GenS1{F0: 127, F1: "a", F2: v2}
    for fi14 := 0; fi14 < 2; fi14++ {
        if false {
            switch (("\n\t" + "foo bar") + "\n\t") {
            case "":
                fmt.Println(v12, fi14)
            case "\n\t":
                fmt.Println(fi14, fi14)
            case "hello":
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 float32 = 0.1
                    switch (-(fi15)) {
                    case 32768:
                        _ = len(v1)
                    case 53:
                        v17 := GenS1{F0: fi14, F1: ("foo bar" + ""), F2: -118}
                        _ = v17
                    default:
                        var v18 uint16 = 0
                        if false {
                            _ = len(v1)
                        }
                        var v19 int8 = (v2 >> ((1) & 63))
                        fmt.Println(v19, v18)
                        var v20 uint = 4294967295
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    fmt.Println(v12, fi14)
                    v2 -= v2
                    _ = fi15
                    _ = v16
                }
            default:
                var v21 uint8 = 36
                _ = v21
            }
        }
        fmt.Println(v12)
        if ((true && (!false)) && false) {
            fmt.Println(v2, fi14)
            v12 = 795
            v22 := GenS1{F0: 2147483647, F1: (("\n\t" + "hello") + "foo bar"), F2: v2}
            _ = v22
        } else {
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 uint32 = 2147483648
                v12 += v12
                v25 := GenS1{F0: (-(v12)), F1: "", F2: v2}
                _ = fi23
                _ = v24
                _ = v25
            }
            v26 := GenS1{F0: (fi14 + v12), F1: "hello", F2: -124}
            _ = v26
        }
        _ = fi14
    }
    var v27 float32 = 1.5
    v12 = (v12 >> ((1) & 63))
    v2 += (v2 & (v2 - v2))
    if (("\n\t" + "\n\t") == "a") {
        v28 := v3["hello"]
        _ = v28
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 int = (v12 / ((fi29) | 1))
        _ = fi29
        _ = v30
    }
    v31 := GenS1{F0: 42, F1: "foo bar", F2: (v2 ^ -1)}
    for fi32 := 0; fi32 < 13; fi32++ {
        for fi33 := 0; fi33 < 18; fi33++ {
            if true {
                v12 = 1
                v34 := GenS2{F0: 59, F1: 18446744073709551615}
                _ = v34
            } else {
                var v35 uint = 207
                var v36 int16 = -690
                switch (v12 ^ (fi33 % (((fi33 | (-(fi33)))) | 1))) {
                case 32768:
                    if (!true) {
                        var v37 float64 = 1.5
                        v38 := []int{128}
                        _ = v37
                        _ = v38
                    }
                }
                _ = v35
                _ = v36
            }
            var v39 int = (fi33 % ((-607) | 1))
            _ = fi33
            _ = v39
        }
        fmt.Println(fi32)
        v2 = v2
        _ = fi32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v7
}
