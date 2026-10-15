package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int8 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 float32 = 0.1
    var v2 int16 = -1
    var v3 float64 = 3.14159
    v4 := []int{127, 1, 2147483647, 4294967295}
    fmt.Println(v3)
    fmt.Println(v1, v1)
    v5 := GenS1{F0: -209, F1: false, F2: -128}
    var v6 string = "foo bar"
    switch v6 {
    case "foo bar":
        for fi7 := 0; fi7 < 4; fi7++ {
            v2 += v2
            _ = fi7
        }
    case "":
        switch (v6 + (((v6 + "\n\t") + (((v6 + v6) + v6) + ("\n\t" + v6))) + "")) {
        case "":
            v2 += (-(v2))
        case "hello":
            v8 := GenS1{F0: 0, F1: true, F2: 100}
            _ = v8
        case "a":
            v6 = ""
        default:
            if (!(true || true)) {
                fmt.Println(v6, v1)
            } else {
                switch -616 {
                case 32768:
                    if false {
                        if false {
                            v6 = v6
                            v9 := GenS1{F0: 42, F1: (true || false), F2: -17}
                            _ = v9
                        }
                        v3 -= float64(v1)
                        var v10 int16 = v2
                        _ = v10
                    } else {
                        fmt.Println(v3)
                        v3 = (v3 / ((0.1 - (float64((0.0 / v1)) / (v3 / 3.14159))) * v3))
                    }
                case 2147483647:
                    v1 = v1
                case 42:
                    if (32767 < 704) {
                        var v11 bool = (v3 <= (v3 * (float64(float32(3.14159)) + v3)))
                        if v11 {
                            v11 = (4294967295 < 0)
                            var v12 int = 42
                            _ = v12
                        } else {
                            for fi13 := 0; fi13 < 9; fi13++ {
                                var v14 float64 = ((v3 * (0.0 - ((((1.5 - v3) / 1e10) - v3) / (v3 - v3)))) - (v3 / float64(0.1)))
                                _ = fi13
                                _ = v14
                            }
                        }
                        _ = v11
                    } else {
                        switch 128 {
                        case -1:
                            v15 := GenS1{F0: 0, F1: (float64(3.14159) < v3), F2: 1}
                            _ = v15
                        case -325:
                            if (18446744073709551615 >= 325) {
                                var v16 int8 = 1
                                _ = v16
                            }
                        }
                    }
                default:
                    v6 = v6
                }
            }
        }
    case "\n\t":
        var v17 float32 = -2.25
        _ = v17
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println(v6)
    for fi18 := 0; fi18 < 9; fi18++ {
        for fi19 := 0; fi19 < 2; fi19++ {
            if (((-(v2)) >= v2) || true) {
                switch 216 {
                case -2147483648:
                    var v20 float64 = (v3 * v3)
                    _ = v20
                }
                v1 -= -2.25
            }
            fmt.Println(v1)
            _ = fi19
        }
        if false {
            var v21 uint64 = 42
            switch ("a" + "foo bar") {
            case "\n\t":
                v22 := GenS2{F0: ((("\n\t" + v6) + v6) + v6), F1: v2}
                _ = v22
            case "foo bar":
                var v23 bool = false
                _ = v23
            case "":
                for fi24 := 0; fi24 < 13; fi24++ {
                    if true {
                        if false {
                            switch fi18 {
                            case 340:
                                fmt.Println(v2)
                            default:
                                v4 = append(v4, fi18)
                            }
                            fmt.Println(v21, fi24, v2)
                            if ((v6 == ("a" + ("hello" + "\n\t"))) && (v3 != float64((v1 / 1e10)))) {
                                v25 := GenS3{F0: 0}
                                fmt.Println(fi18, fi24)
                                _ = v25
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = -279
                                }
                            }
                        } else {
                            v21 += v21
                        }
                        v1 = float32(v3)
                    }
                    _ = fi24
                }
            default:
                _ = len(v4)
            }
            for fi26 := 0; fi26 < 18; fi26++ {
                fmt.Println(fi18, fi18)
                v2 -= v2
                switch -2147483648 {
                case 128:
                    v1 = v1
                }
                _ = fi26
            }
            _ = v21
        }
        v27 := GenS1{F0: -9223372036854775808, F1: false, F2: 100}
        _ = fi18
        _ = v27
    }
    fmt.Println(v2, v2, v3)
    switch "foo bar" {
    case "foo bar":
        v3 = 1e10
    case "\n\t":
        fmt.Println(v3)
    default:
        fmt.Println(v1, v1, v6)
    }
    v2 = (v2 ^ (v2 / (((-(v2))) | 1)))
    var v28 float32 = (v1 + (((float32(v3) - v1) * 1e10) + (0.0 - ((3.14159 - (float32((v3 / 3.14159)) * 3.14159)) + 0.0))))
    if true {
        fmt.Println(v1, v1, v6)
    } else {
        if (!false) {
            if len(v4) > 0 {
                v4[0] = 32768
            }
            var v29 uint32 = 42
            _ = v29
        }
        fmt.Println(v28, v3, v2)
    }
    fmt.Println(v6, v28, v3)
    v30 := GenS2{F0: (((v6 + ("foo bar" + (v6 + v6))) + (v6 + v6)) + "hello"), F1: (v2 >> ((1) & 63))}
    var v31 float32 = v1
    v32 := GenS2{F0: v6, F1: (v2 - (v2 << ((42) & 63)))}
    for fi33 := 0; fi33 < 7; fi33++ {
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v34 int = (fi33 / ((-403) | 1))
            v35 := GenS3{F0: 1}
            _ = v34
            _ = v35
        }
        switch (fi33 - fi33) {
        case 4294967295:
            var v36 uint = 0
            _ = v36
        case 559:
            v37 := GenS1{F0: 1, F1: (9223372036854775807 == 1), F2: -128}
            _ = v37
        case -9223372036854775808:
            _ = len(v4)
        }
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
