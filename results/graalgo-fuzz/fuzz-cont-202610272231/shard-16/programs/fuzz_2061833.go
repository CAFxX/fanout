package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int32; F2 int }
type GenS3 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint8) uint32 {
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 float64 = -2.25
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: p0}
    _ = p0
    _ = v8
    return 42
}
func genFn1() uint16 {
    var v9 uint32 = 4294967295
    v9 = 702
    v9 = 48
    _ = v9
    return 365
}

func main() {
    v1 := []int{-147, -632, 127}
    var v2 bool = (0 > 7)
    v3 := []int{1, -56}
    var v4 int32 = 0
    var v5 uint = 154
    switch 4294967295 {
    case 1:
        fmt.Println(v2, v2)
    case -9223372036854775808:
        if v2 {
            var v10 int64 = 42
            if v2 {
                v5 -= v5
                var v11 float32 = (float32((-2.25 * 0.1)) * 0.0)
                _ = v11
            }
            _ = v10
        }
    default:
        for fi12 := 0; fi12 < 11; fi12++ {
            v13 := GenS1{F0: 68}
            var v14 uint32 = 869
            if v2 {
                v4 = v4
                var v15 uint = 0
                fmt.Println(fi12, v2)
                _ = v15
            }
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    fmt.Println(v4, v4, v4)
    v5 = (v5 + 65535)
    if v2 {
        v4 -= v4
        var v16 int = 495
        switch ("" + ("" + "")) {
        case "":
            v1 = append(v1, 1)
        default:
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        }
        _ = v16
    } else {
        v4 += v4
    }
    var v17 bool = (("a" + "\n\t") < "\n\t")
    switch (("" + (("" + ("\n\t" + "foo bar")) + "a")) + (("hello" + "\n\t") + ("" + "a"))) {
    case "foo bar":
        for fi18 := 0; fi18 < 18; fi18++ {
            switch (("" + ("foo bar" + "hello")) + "hello") {
            case "foo bar":
                if ((v5 & 65535) >= (v5 % ((v5) | 1))) {
                    v17 = v17
                    v5 -= (^(v5))
                    v2 = (!(0 >= -1))
                }
            case "\n\t":
                v17 = (genFn0(255) <= 436)
            default:
                switch (("a" + ("\n\t" + "foo bar")) + "") {
                case "a":
                    if v17 {
                        v17 = (true || (!false))
                        if (32767 != 255) {
                            v19 := GenS1{F0: 127}
                            v3 = append(v3, (fi18 * fi18))
                            v20 := GenS2{F0: 472, F1: v4, F2: (fi18 - (fi18 * fi18))}
                            _ = v19
                            _ = v20
                        }
                    }
                case "hello":
                    fmt.Println(v17)
                case "":
                    for fi21 := 0; fi21 < 7; fi21++ {
                        v22 := GenS2{F0: 42, F1: (v4 / ((v4) | 1)), F2: -412}
                        switch fi21 {
                        case -1:
                            fmt.Println(v2)
                        }
                        _ = fi21
                        _ = v22
                    }
                }
            }
            v17 = v17
            _ = fi18
        }
    }
    var v23 uint = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        v5 -= v23
        var v24 uint32 = 2
        _ = v24
    }
    var v25 string = "\n\t"
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v26 := GenS3{F0: v2, F1: 2}
    fmt.Println(v2, v17)
    var v27 float64 = float64((-2.25 - float32((0.1 * -2.25))))
    if true {
        v2 = (!false)
        v17 = v2
        v5 = 0
    } else {
        if (255 >= 2) {
            var v28 string = v25
            _ = v28
        }
        v17 = (!((100 != -72) && false))
    }
    v4 -= (v4 / (((^(v4))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
