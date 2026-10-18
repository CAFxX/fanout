package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int32; F2 int16 }
type GenS3 struct { F0 uint8; F1 uint32; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 842
    v3 := make(map[string]string)
    if (!(!(true && false))) {
        v2 -= -1
        for fi4 := 0; fi4 < 19; fi4++ {
            v2 += (v2 - (-(v2)))
            _ = fi4
        }
    }
    switch 537 {
    case 360:
        var v5 string = "foo bar"
        _ = v5
    case -357:
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 float32 = 0.1
            switch fi6 {
            case -160:
                var v8 int8 = -97
                _ = v8
            case 1:
                switch "a" {
                case "hello":
                    v7 -= v7
                case "foo bar":
                    switch "foo bar" {
                    case "\n\t":
                        v9 := GenS1{F0: 1}
                        _ = v9
                    case "hello":
                        var v10 int32 = (v2 / ((1) | 1))
                        _ = v10
                    case "":
                        for fi11 := 0; fi11 < 13; fi11++ {
                            if (!(!false)) {
                                var v12 uint = 18446744073709551615
                                switch (fi6 / ((fi6) | 1)) {
                                case -9223372036854775808:
                                    v2 = (v2 | -869)
                                case 128:
                                    var v13 float64 = 0.1
                                    _ = v13
                                default:
                                    v14 := GenS1{F0: 0}
                                    _ = v14
                                }
                                _ = v12
                            } else {
                                fmt.Println(fi11, fi6)
                                v7 = v7
                                v7 += (0.0 * 0.1)
                            }
                            var v15 string = ""
                            _ = fi11
                            _ = v15
                        }
                    }
                default:
                    switch "" {
                    case "":
                        var v16 string = "a"
                        _ = v16
                    case "foo bar":
                        v2 = v2
                    case "hello":
                        fmt.Println(v2, v7)
                    default:
                        var v17 int32 = (^(v2))
                        _ = v17
                    }
                }
            case 0:
                switch (fi6 << ((65535) & 63)) {
                case 2147483647:
                    switch fi6 {
                    case 127:
                        switch "hello" {
                        case "\n\t":
                            if (42 < 1) {
                                fmt.Println(v2)
                            } else {
                                if (!true) {
                                    v18 := GenS1{F0: 121}
                                    _ = v18
                                }
                                switch "hello" {
                                case "a":
                                    switch "" {
                                    case "a":
                                        v19 := GenS1{F0: -78}
                                        _ = v19
                                    case "":
                                        _ = len(v1)
                                    }
                                }
                                v7 += (v7 - v7)
                            }
                        }
                    case -1:
                        v20 := GenS2{F0: (v2 / ((v2) | 1)), F1: (v2 / ((v2) | 1)), F2: 872}
                        _ = v20
                    default:
                        if (true || ((true || true) && true)) {
                            for fi21 := 0; fi21 < 14; fi21++ {
                                v7 -= 0.1
                                fmt.Println(v2)
                                v7 += (float32(3.14159) * v7)
                                _ = fi21
                            }
                        } else {
                            var v22 uint32 = 2
                            _ = v22
                        }
                    }
                case 485:
                    v7 += (v7 + float32(-2.25))
                case 42:
                    var v23 uint64 = 2
                    _ = v23
                default:
                    for fi24 := 0; fi24 < 16; fi24++ {
                        v25 := v3[""]
                        var v26 string = "hello"
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                }
            }
            v7 -= v7
            _ = fi6
            _ = v7
        }
    case -1:
        for fi27 := 0; fi27 < 4; fi27++ {
            var v28 int8 = -1
            v2 = 2147483647
            switch (("" + "a") + "hello") {
            case "a":
                for fi29 := 0; fi29 < 2; fi29++ {
                    var v30 uint32 = 2
                    if ("hello" == "foo bar") {
                        delete(v1, fi29)
                        if (!(!((true && (704 == 326)) && true))) {
                            switch ("" + "\n\t") {
                            case "foo bar":
                                v30 -= 2147483647
                            }
                            fmt.Println(fi27, v2, fi29)
                            v31 := GenS1{F0: (v28 - v28)}
                            _ = v31
                        }
                        v32 := []int{}
                        _ = v32
                    }
                    _ = fi29
                    _ = v30
                }
            case "\n\t":
                var v33 float32 = 3.14159
                _ = v33
            }
            _ = fi27
            _ = v28
        }
    }
    v34 := GenS3{F0: 127, F1: 0, F2: 9223372036854775807}
    var v35 uint32 = 2147483648
    if true {
        var v36 uint16 = 32767
        v37 := []int{2147483647, 60}
        _ = v36
        _ = v37
    }
    var v38 float64 = (((((((0.0 - 0.0) + 0.1) - 1e10) / 1e10) / float64(((-2.25 / (0.1 + 0.0)) / 1e10))) * 0.0) - 0.1)
    var v39 int8 = -1
    for fi40 := 0; fi40 < 12; fi40++ {
        v41 := v1[fi40]
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        v3["a"] = "hello"
        _ = fi42
    }
    v38 -= (float64((-2.25 + float32(float64((3.14159 - 1.5))))) + 1.5)
    v35 += (v35 & (v35 - v35))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v39
}
