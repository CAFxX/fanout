package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint }

func main() {
    v1 := []int{81, 32768, 4294967295}
    var v2 string = "a"
    var v3 string = v2
    var v4 uint16 = 1
    var v5 uint32 = 2
    if len(v1) > 0 {
        v1[0] = int(v5)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 := GenS1{F0: -157}
    switch -9223372036854775808 {
    case -169:
        v5 = (v5 / (((v5 + v5)) | 1))
    case 107:
        var v7 bool = false
        _ = v7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: -1}
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS1{F0: -793}
            _ = fi10
            _ = v11
        }
        switch fi8 {
        case 2147483647:
            if (-1 < -547) {
                if (false || (!(255 > 127))) {
                    var v12 string = v3
                    switch 645 {
                    case -9223372036854775808:
                        var v13 float64 = (1.5 * -2.25)
                        _ = v13
                    default:
                        for fi14 := 0; fi14 < 3; fi14++ {
                            if (9223372036854775807 < 42) {
                                switch fi8 {
                                case 42:
                                    switch fi8 {
                                    case 127:
                                        var v15 int = (^(fi8))
                                        _ = v15
                                    }
                                case -9223372036854775808:
                                    var v16 int = (fi8 % ((int(v5)) | 1))
                                    _ = v16
                                case 799:
                                    switch v2 {
                                    case "hello":
                                        var v17 bool = true
                                        _ = v17
                                    case "\n\t":
                                        if len(v1) > 0 {
                                            v1[0] = (fi14 + (fi14 >> ((255) & 63)))
                                        }
                                    }
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                if (!(-1 == -1)) {
                                    v4 = v4
                                    v3 = ("hello" + "foo bar")
                                } else {
                                    switch "\n\t" {
                                    case "":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    default:
                                        fmt.Println(fi14, fi14, v4)
                                        if false {
                                            v4 = 727
                                        }
                                        switch (fi14 - (fi14 << ((1) & 63))) {
                                        case -303:
                                            var v18 float64 = (0.1 + ((((-2.25 * (float64(3.14159) / 1.5)) / (3.14159 / 1e10)) * (float64(-2.25) * 1e10)) - (-2.25 / 0.1)))
                                            _ = v18
                                        case 575:
                                            switch "hello" {
                                            case "foo bar":
                                                fmt.Println(v3, fi8)
                                            case "hello":
                                                v19 := GenS1{F0: 9223372036854775807}
                                                _ = v19
                                            case "":
                                                fmt.Println(v4, fi8)
                                            }
                                        default:
                                            var v20 uint = 54
                                            _ = v20
                                        }
                                    }
                                    var v21 uint32 = (^(v5))
                                    v4 = v4
                                    _ = v21
                                }
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = (fi8 - fi14)
                                }
                                v22 := GenS1{F0: -986}
                                v1 = append(v1, (fi14 + 42))
                                _ = v22
                            }
                            _ = fi14
                        }
                    }
                    var v23 int = (fi8 >> ((150) & 63))
                    _ = v12
                    _ = v23
                }
                var v24 int16 = 1000
                v1 = append(v1, fi8)
                _ = v24
            }
        case -9223372036854775808:
            fmt.Println(v4, v5, fi8)
        case -533:
            for fi25 := 0; fi25 < 5; fi25++ {
                var v26 int8 = 0
                var v27 bool = ((fi8 * fi8) < fi25)
                var v28 int = 4294967295
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
        }
        _ = fi8
        _ = v9
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v29 := GenS2{F0: -151, F1: 255}
    var v30 int64 = 592
    var v31 string = v2
    v31 = v3
    v31 = v3
    v30 = v30
    for fi32 := 0; fi32 < 15; fi32++ {
        for fi33 := 0; fi33 < 14; fi33++ {
            v34 := GenS1{F0: 9223372036854775807}
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    switch ("a" + v2) {
    case "hello":
        v30 -= v30
    case "a":
        v5 += v5
    case "foo bar":
        fmt.Println(v3, v3, v5)
    default:
        var v35 int32 = 42
        _ = v35
    }
    var v36 float32 = (0.1 + -2.25)
    var v37 float32 = v36
    var v38 uint16 = (v4 % (((v4 << ((uint(v4)) & 63))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
