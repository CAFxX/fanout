package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 string }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint8; F1 int8; F2 bool }

func main() {
    v1 := []int{}
    var v2 int64 = 0
    var v3 int8 = 7
    v4 := make(map[string]bool)
    v5 := make(map[string]string)
    v6 := []int{-9223372036854775808, 128, 1}
    var v7 int32 = 0
    for fi8 := 0; fi8 < 14; fi8++ {
        switch "\n\t" {
        case "":
            var v9 uint8 = 12
            _ = v9
        case "hello":
            var v10 uint = 754
            _ = v10
        case "foo bar":
            fmt.Println(v7, v2)
        default:
            for fi11 := 0; fi11 < 7; fi11++ {
                fmt.Println(fi11, fi8, v7)
                var v12 int32 = (v7 & 42)
                fmt.Println(v12, v12)
                _ = fi11
                _ = v12
            }
        }
        _ = fi8
    }
    var v13 float32 = (0.0 * -2.25)
    for fi14 := 0; fi14 < 10; fi14++ {
        switch -9223372036854775808 {
        case -2147483648:
            if (!true) {
                switch 0 {
                case -666:
                    for fi15 := 0; fi15 < 12; fi15++ {
                        v16 := GenS1{F0: v3}
                        _ = fi15
                        _ = v16
                    }
                case 1:
                    v3 -= (v3 * 100)
                default:
                    v17 := GenS1{F0: v3}
                    _ = v17
                }
            }
        case 32767:
            v18 := GenS1{F0: (v3 / ((127) | 1))}
            _ = v18
        }
        var v19 uint64 = 1
        v13 = 3.14159
        _ = fi14
        _ = v19
    }
    v20 := GenS2{F0: 2, F1: 128, F2: ("hello" + "a")}
    for fi21 := 0; fi21 < 5; fi21++ {
        for fi22 := 0; fi22 < 19; fi22++ {
            var v23 int = (fi21 * fi21)
            var v24 uint = 65535
            v23 = (fi21 + fi22)
            _ = fi22
            _ = v23
            _ = v24
        }
        var v25 int16 = int16(v3)
        for fi26 := 0; fi26 < 9; fi26++ {
            fmt.Println(fi26, v13)
            switch (fi21 >> ((261) & 63)) {
            case 133:
                if ((fi21 % ((fi26) | 1)) != (fi26 / (((fi21 | (fi21 ^ fi26))) | 1))) {
                    var v27 int64 = 0
                    v28 := v4["\n\t"]
                    switch "" {
                    case "":
                        v13 = ((0.0 / v13) - (v13 + 3.14159))
                        if ((fi26 % ((fi21) | 1)) >= -1) {
                            switch (("a" + "\n\t") + ("foo bar" + ((("\n\t" + "foo bar") + "hello") + ""))) {
                            case "":
                                v25 = 868
                            case "a":
                                if ((v27 + -9223372036854775808) >= (v27 - 42)) {
                                    var v29 string = ""
                                    if len(v1) > 0 {
                                        v1[0] = (fi26 / (((fi21 % (((fi26 / (((fi26 << ((18446744073709551615) & 63))) | 1))) | 1))) | 1))
                                    }
                                    v30 := GenS1{F0: 1}
                                    _ = v29
                                    _ = v30
                                }
                            }
                            v31 := GenS3{F0: 255, F1: ((("foo bar" + "\n\t") + ("a" + "hello")) + "\n\t")}
                            _ = v31
                        } else {
                            v4[("\n\t" + ("hello" + "foo bar"))] = ((float64(v13) + 1e10) == float64(0.1))
                        }
                        v32 := v4[""]
                        _ = v32
                    case "\n\t":
                        v28 = false
                        v33 := GenS4{F0: 42, F1: (-(v3)), F2: false}
                        _ = v33
                    }
                    delete(v4, ("foo bar" + "hello"))
                    var v34 float32 = v13
                    _ = v27
                    _ = v28
                    _ = v34
                }
            default:
                delete(v4, "a")
            }
            _ = fi26
        }
        _ = fi21
        _ = v25
    }
    fmt.Println(v7)
    if true {
        fmt.Println(v2, v3)
    }
    var v35 float64 = ((1.5 * 0.0) * 1.5)
    var v36 uint16 = 21
    if len(v6) > 0 {
        v6[0] = 857
    }
    var v37 uint32 = 450
    v7 -= v7
    v4[""] = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
