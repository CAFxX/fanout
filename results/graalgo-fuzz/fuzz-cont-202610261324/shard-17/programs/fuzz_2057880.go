package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint32 }

func main() {
    var v1 float64 = (float64((-2.25 + 3.14159)) / 1.0)
    v2 := make(map[int]int)
    var v3 float64 = 1e10
    var v4 int16 = 5
    var v5 uint16 = 65535
    v6 := make(map[int]int)
    v7 := GenS1{F0: 9223372036854775807, F1: 42, F2: 498}
    v6[32768] = 32767
    fmt.Println(v5, v4, v4)
    switch ("\n\t" + "") {
    case "a":
        var v8 int32 = 1
        _ = v8
    case "":
        switch 127 {
        case 0:
            for fi9 := 0; fi9 < 17; fi9++ {
                fmt.Println(v3, fi9)
                for fi10 := 0; fi10 < 12; fi10++ {
                    var v11 uint16 = 0
                    v11 = v11
                    var v12 int64 = 0
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                switch ("foo bar" + ("\n\t" + (("foo bar" + "foo bar") + "hello"))) {
                case "foo bar":
                    var v13 int16 = v4
                    _ = v13
                default:
                    var v14 uint16 = 1
                    _ = v14
                }
                _ = fi9
            }
        }
    }
    if false {
        v15 := GenS1{F0: 0, F1: 9223372036854775807, F2: 2147483647}
        v16 := []int{127, 0, -2147483648}
        for fi17 := 0; fi17 < 18; fi17++ {
            v1 -= (3.14159 * float64((0.0 - -2.25)))
            switch 454 {
            case -159:
                v4 = -778
            default:
                for fi18 := 0; fi18 < 17; fi18++ {
                    var v19 int8 = 0
                    v19 -= v19
                    v6[fi17] = fi17
                    _ = fi18
                    _ = v19
                }
            }
            var v20 uint = 375
            _ = fi17
            _ = v20
        }
        _ = v15
        _ = v16
    }
    fmt.Println(v4, v3)
    for fi21 := 0; fi21 < 18; fi21++ {
        for fi22 := 0; fi22 < 1; fi22++ {
            if true {
                if (false || (false && (!false))) {
                    switch 1 {
                    case 0:
                        var v23 uint8 = 1
                        _ = v23
                    case 157:
                        v1 -= (v1 * ((v1 * v1) + v3))
                    }
                }
            }
            v24 := GenS1{F0: -647, F1: -1, F2: 42}
            _ = fi22
            _ = v24
        }
        v25 := GenS1{F0: -9223372036854775808, F1: -903, F2: 2147483648}
        if true {
            if (true || true) {
                var v26 int64 = 1
                _ = v26
            } else {
                for fi27 := 0; fi27 < 1; fi27++ {
                    v28 := GenS1{F0: -9223372036854775808, F1: 9223372036854775807, F2: 1}
                    _ = fi27
                    _ = v28
                }
            }
        }
        _ = fi21
        _ = v25
    }
    if (0 <= 0) {
        fmt.Println(v5, v5, v5)
        for fi29 := 0; fi29 < 12; fi29++ {
            v3 += float64(1.5)
            var v30 int16 = (v4 + 0)
            _ = fi29
            _ = v30
        }
    }
    var v31 int16 = v4
    v32 := []int{-1, 1, 419, 127}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
