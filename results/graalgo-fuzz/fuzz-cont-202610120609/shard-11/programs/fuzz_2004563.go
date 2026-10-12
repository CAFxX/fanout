package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint8; F1 uint8; F2 int }

func main() {
    v1 := []int{-164, 32768, 32767, -9223372036854775808}
    var v2 float64 = 1e10
    v3 := []int{32768}
    if (!(475 <= -193)) {
        v3 = append(v3, 0)
        var v4 bool = (!((-2.25 * float32((v2 * (v2 * v2)))) != 0.0))
        _ = v4
    } else {
        if false {
            _ = len(v1)
            v2 -= (v2 + v2)
        } else {
            if (!true) {
                switch 1 {
                case 2147483647:
                    for fi5 := 0; fi5 < 13; fi5++ {
                        switch fi5 {
                        case 32767:
                            v6 := GenS1{F0: 63, F1: 880, F2: -32768}
                            _ = v6
                        case 134:
                            v2 -= ((3.14159 * ((3.14159 - ((v2 - v2) - (-2.25 + (3.14159 + -2.25)))) / (v2 * (0.0 * v2)))) + (v2 * v2))
                        default:
                            v7 := GenS2{F0: -32768}
                            _ = v7
                        }
                        var v8 uint = 0
                        for fi9 := 0; fi9 < 18; fi9++ {
                            var v10 float64 = (v2 - (1e10 + v2))
                            switch ((("\n\t" + "a") + "a") + ("hello" + "foo bar")) {
                            case "hello":
                                v2 -= float64(1.5)
                            case "a":
                                var v11 int = -1
                                _ = v11
                            }
                            fmt.Println(fi5, v10, fi9)
                            _ = fi9
                            _ = v10
                        }
                        _ = fi5
                        _ = v8
                    }
                }
            }
        }
        var v12 float32 = (float32((((((v2 - v2) * float64(0.0)) / 0.1) + 0.0) * v2)) / 0.1)
        var v13 uint16 = 0
        _ = v12
        _ = v13
    }
    switch ("\n\t" + ("" + ("" + ("hello" + "")))) {
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "foo bar":
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 int64 = 1
            var v16 uint = 674
            var v17 uint16 = 255
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
    }
    var v18 bool = ((!true) || false)
    if v18 {
        v19 := GenS2{F0: 32767}
        var v20 uint64 = 2
        fmt.Println(v18, v2, v18)
        _ = v19
        _ = v20
    }
    v18 = v18
    fmt.Println(v2, v2)
    for fi21 := 0; fi21 < 14; fi21++ {
        fmt.Println(fi21, fi21)
        var v22 float64 = 3.14159
        if len(v3) > 0 {
            v3[0] = (fi21 * fi21)
        }
        _ = fi21
        _ = v22
    }
    v18 = ((1e10 / (1e10 / (1.5 / v2))) == (float64(0.1) * v2))
    for fi23 := 0; fi23 < 18; fi23++ {
        for fi24 := 0; fi24 < 11; fi24++ {
            v25 := GenS2{F0: int16(fi23)}
            _ = fi24
            _ = v25
        }
        var v26 bool = (111 >= 645)
        if (62 == 973) {
            var v27 int64 = -740
            v27 -= (v27 >> ((360) & 63))
            _ = v27
        }
        _ = fi23
        _ = v26
    }
    v2 = ((3.14159 + -2.25) - v2)
    var v28 bool = v18
    fmt.Println(v28)
    v18 = (0 == -9223372036854775808)
    var v29 int = -2147483648
    if false {
        v30 := GenS3{F0: v29, F1: 65535}
        v31 := GenS4{F0: 2, F1: 143, F2: v29}
        _ = v30
        _ = v31
    }
    var v32 bool = false
    v29 -= 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
}
