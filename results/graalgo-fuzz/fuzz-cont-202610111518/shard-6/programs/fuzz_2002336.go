package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint }

func main() {
    v1 := []uint32{0, 2147483647}
    v2 := make(map[int]int)
    var v3 int32 = 731
    var v4 int = 1
    v5 := []uint32{1, 65535, 675, 42}
    v3 += (v3 | -2147483648)
    switch 869 {
    case 42:
        switch (v4 << ((892) & 63)) {
        case 32767:
            v3 = (v3 % (((v3 ^ v3)) | 1))
        case 1:
            _ = len(v5)
        case 42:
            if len(v5) > 0 {
                v5[0] = 4294967295
            }
        default:
            v3 = (v3 * 426)
        }
    default:
        if true {
            if false {
                if (!(2 < 738)) {
                    var v6 bool = (-959 <= -9223372036854775808)
                    if false {
                        var v7 float32 = -2.25
                        _ = v7
                    } else {
                        fmt.Println(v3, v3)
                        for fi8 := 0; fi8 < 16; fi8++ {
                            v2[(-(fi8))] = v4
                            if (v6 && ((65535 <= 405) || v6)) {
                                if (!(502 != 695)) {
                                    var v9 string = "hello"
                                    var v10 float64 = 3.14159
                                    for fi11 := 0; fi11 < 13; fi11++ {
                                        var v12 int8 = 5
                                        var v13 float32 = 1e10
                                        var v14 int16 = -1
                                        _ = fi11
                                        _ = v12
                                        _ = v13
                                        _ = v14
                                    }
                                    _ = v10
                                    _ = v9
                                }
                                v3 = v3
                            }
                            switch "" {
                            case "foo bar":
                                if v6 {
                                    for fi15 := 0; fi15 < 7; fi15++ {
                                        fmt.Println(fi15)
                                        switch ((("\n\t" + "\n\t") + ((("" + "foo bar") + "hello") + "a")) + "\n\t") {
                                        case "\n\t":
                                            var v16 float64 = 3.14159
                                            _ = v16
                                        case "a":
                                            v4 = (fi15 << ((4294967295) & 63))
                                        case "foo bar":
                                            var v17 int16 = -343
                                            _ = v17
                                        default:
                                            fmt.Println(fi8, fi15, v4)
                                        }
                                        _ = fi15
                                    }
                                    v18 := GenS1{F0: 0}
                                    for fi19 := 0; fi19 < 20; fi19++ {
                                        var v20 int = (fi8 / (((-(v4))) | 1))
                                        _ = fi19
                                        _ = v20
                                    }
                                    _ = v18
                                }
                            case "a":
                                var v21 string = ""
                                _ = v21
                            case "\n\t":
                                v3 -= (v3 & (v3 / (((v3 & (v3 >> ((1) & 63)))) | 1)))
                            default:
                                v22 := GenS1{F0: 9223372036854775807}
                                _ = v22
                            }
                            _ = fi8
                        }
                    }
                    var v23 int = v4
                    _ = v23
                    _ = v6
                }
                v24 := GenS1{F0: -1}
                _ = v24
            }
        }
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        var v26 uint32 = 2147483647
        _ = fi25
        _ = v26
    }
    v3 = (v3 ^ (^(v3)))
    v3 = -2147483648
    switch "a" {
    case "a":
        var v27 uint64 = 2
        _ = v27
    case "":
        v4 = (v4 + v4)
    }
    var v28 float32 = -2.25
    v28 += v28
    var v29 int16 = -32768
    fmt.Println(v3, v4)
    if false {
        var v30 string = "foo bar"
        var v31 int = v4
        _ = v30
        _ = v31
    }
    v28 += (v28 - 1e10)
    if (!true) {
        var v32 int8 = -29
        if (0 != 42) {
            switch (v4 % (((-(v4))) | 1)) {
            case 9223372036854775807:
                if false {
                    v33 := GenS2{F0: -70, F1: 255, F2: 65535}
                    if false {
                        for fi34 := 0; fi34 < 13; fi34++ {
                            v35 := GenS1{F0: 9223372036854775807}
                            _ = fi34
                            _ = v35
                        }
                        v4 = 42
                        var v36 float64 = 0.1
                        _ = v36
                    } else {
                        var v37 int16 = v29
                        v2[(v4 << ((1) & 63))] = v4
                        _ = v37
                    }
                    _ = v33
                }
            default:
                var v38 int8 = (v32 / (((^(v32))) | 1))
                _ = v38
            }
        } else {
            v39 := GenS2{F0: (v32 / ((6) | 1)), F1: 0, F2: 477}
            _ = v39
        }
        v3 = v3
        _ = v32
    }
    var v40 int64 = -294
    v41 := GenS2{F0: -128, F1: 32767, F2: 4294967295}
    var v42 uint32 = 4294967295
    var v43 uint32 = (v42 % ((v42) | 1))
    v29 = v29
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
