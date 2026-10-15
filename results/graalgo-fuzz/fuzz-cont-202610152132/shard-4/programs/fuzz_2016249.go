package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int8; F1 uint; F2 string }
type GenS3 struct { F0 int; F1 int8; F2 int32 }
type GenS4 struct { F0 int8; F1 int8; F2 uint }

func main() {
    var v1 float64 = 3.14159
    var v2 string = "hello"
    v3 := []int{-1, -13, -2147483648}
    v4 := make(map[string]string)
    v2 = v2
    var v5 int8 = -123
    switch v2 {
    case "hello":
        var v6 uint8 = 127
        _ = v6
    case "foo bar":
        for fi7 := 0; fi7 < 18; fi7++ {
            v5 = (v5 ^ 127)
            v5 = -128
            _ = fi7
        }
    case "\n\t":
        v8 := GenS1{F0: v2, F1: 9223372036854775807}
        _ = v8
    }
    delete(v4, (v2 + v2))
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 uint64 = 42
        v2 = "hello"
        _ = fi9
        _ = v10
    }
    v1 = float64(1.5)
    if (-909 > 434) {
        var v11 int = -678
        v12 := GenS1{F0: "a", F1: 9223372036854775807}
        _ = v11
        _ = v12
    } else {
        v5 -= -1
        v13 := GenS1{F0: v2, F1: 42}
        v5 = (v5 + (v5 >> ((65535) & 63)))
        _ = v13
    }
    v4[("hello" + v2)] = v2
    for fi14 := 0; fi14 < 16; fi14++ {
        if (213 > 2147483647) {
            v2 = (v2 + "hello")
            for fi15 := 0; fi15 < 10; fi15++ {
                if true {
                    fmt.Println(v1, fi14, fi15)
                    if ((v5 / ((v5) | 1)) <= v5) {
                        v16 := GenS1{F0: (v2 + v2), F1: -9223372036854775808}
                        switch fi15 {
                        case 127:
                            v1 += (v1 / v1)
                        case -1:
                            var v17 int = fi14
                            _ = v17
                        default:
                            v1 -= v1
                        }
                        _ = v16
                    }
                } else {
                    v18 := GenS1{F0: v2, F1: 0}
                    var v19 int = fi15
                    switch (v2 + (v2 + (v2 + "\n\t"))) {
                    case "a":
                        fmt.Println(v5, v5)
                    case "\n\t":
                        var v20 float32 = 0.1
                        _ = v20
                    }
                    var v21 uint32 = 126
                    _ = v18
                    _ = v19
                    _ = v21
                }
                _ = fi15
            }
        }
        v2 = "a"
        _ = fi14
    }
    var v22 uint32 = 2147483647
    if (1 == 383) {
        switch -9223372036854775808 {
        case 9223372036854775807:
            for fi23 := 0; fi23 < 11; fi23++ {
                for fi24 := 0; fi24 < 13; fi24++ {
                    v25 := GenS2{F0: 124, F1: 515, F2: ""}
                    v2 = (v2 + "")
                    _ = fi24
                    _ = v25
                }
                v5 += -1
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi23
            }
        case 32768:
            v26 := GenS3{F0: 921, F1: -60, F2: -270}
            _ = v26
        }
        v22 += (v22 - v22)
    } else {
        if (681 > 65535) {
            v1 -= (v1 + float64((1.5 + 1e10)))
        } else {
            var v27 int8 = 100
            switch -292 {
            case 250:
                var v28 int8 = v5
                _ = v28
            case 128:
                switch 4294967295 {
                case 42:
                    v2 = v2
                case 4294967295:
                    var v29 uint8 = 128
                    _ = v29
                default:
                    v22 -= (^(v22))
                }
            default:
                switch "a" {
                case "a":
                    v30 := GenS4{F0: (v5 ^ (v27 ^ (v27 * (v27 << ((0) & 63))))), F1: (v27 * v27), F2: 42}
                    _ = v30
                case "foo bar":
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        var v31 uint8 = 104
                        _ = v31
                    default:
                        if true {
                            v32 := GenS4{F0: v27, F1: 31, F2: 42}
                            v27 += (v27 >> ((580) & 63))
                            fmt.Println(v27)
                            _ = v32
                        } else {
                            if true {
                                v27 += (v27 % (((v27 >> ((18446744073709551615) & 63))) | 1))
                                if (2 > 42) {
                                    _ = len(v4)
                                    for fi33 := 0; fi33 < 11; fi33++ {
                                        var v34 int64 = 0
                                        fmt.Println(v1, v2)
                                        v4[((("\n\t" + (v2 + v2)) + (v2 + v2)) + v2)] = ("foo bar" + "foo bar")
                                        _ = fi33
                                        _ = v34
                                    }
                                    v5 = -85
                                }
                            }
                            var v35 float64 = v1
                            _ = v35
                        }
                    }
                case "hello":
                    if (-9223372036854775808 != -9223372036854775808) {
                        if true {
                            var v36 uint64 = 2
                            v36 += v36
                            var v37 float64 = (v1 / (1.5 / (float64(1.5) - (v1 / v1))))
                            _ = v36
                            _ = v37
                        } else {
                            var v38 uint32 = v22
                            _ = v38
                        }
                        if ((^(v5)) >= 100) {
                            v39 := v4["hello"]
                            v4[v39] = v2
                            delete(v4, "")
                            _ = v39
                        }
                    }
                default:
                    var v40 float32 = float32(0.0)
                    _ = v40
                }
            }
            v4[v2] = v2
            _ = v27
        }
        for fi41 := 0; fi41 < 3; fi41++ {
            v1 = float64(0.1)
            _ = fi41
        }
        for fi42 := 0; fi42 < 8; fi42++ {
            fmt.Println(fi42, v1)
            switch "" {
            case "":
                v43 := GenS3{F0: fi42, F1: (v5 ^ 127), F2: 735}
                _ = v43
            case "\n\t":
                if ((fi42 >= (fi42 / ((fi42) | 1))) || true) {
                    v44 := GenS3{F0: (-(fi42)), F1: -1, F2: 963}
                    var v45 uint = 65535
                    v5 += v5
                    _ = v44
                    _ = v45
                } else {
                    for fi46 := 0; fi46 < 12; fi46++ {
                        v47 := GenS3{F0: (fi46 / ((-1) | 1)), F1: 1, F2: -1}
                        v22 = v22
                        v5 = v5
                        _ = fi46
                        _ = v47
                    }
                    if (-1 <= v5) {
                        if false {
                            var v48 uint8 = 128
                            switch (v2 + ("\n\t" + ("" + "a"))) {
                            case "foo bar":
                                fmt.Println(fi42, fi42)
                            case "\n\t":
                                var v49 int8 = (v5 * 0)
                                _ = v49
                            }
                            var v50 bool = false
                            _ = v48
                            _ = v50
                        }
                        for fi51 := 0; fi51 < 4; fi51++ {
                            var v52 int64 = 19
                            v52 -= v52
                            _ = fi51
                            _ = v52
                        }
                    }
                }
            }
            v5 = v5
            _ = fi42
        }
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    v22 = v22
    var v53 uint8 = 42
    v53 += v53
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v53
}
