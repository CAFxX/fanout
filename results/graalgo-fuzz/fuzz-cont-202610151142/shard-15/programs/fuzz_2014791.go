package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 bool }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []int{42, -9223372036854775808, -9223372036854775808}
    var v2 uint32 = 2
    var v3 uint64 = 2
    v4 := []uint32{(v2 - (^(v2))), v2, (v2 & v2), (v2 + 65535)}
    v5 := make(map[int]string)
    v2 = v2
    for fi6 := 0; fi6 < 3; fi6++ {
        switch fi6 {
        case 9223372036854775807:
            var v7 string = "\n\t"
            _ = v7
        case 32767:
            v8 := GenS1{F0: (((v2 + 4294967295) != (v2 * (v2 - 815))) || ("a" == "")), F1: -1, F2: (18446744073709551615 > 0)}
            _ = v8
        default:
            v3 = uint64(v2)
        }
        _ = fi6
    }
    _ = len(v1)
    if len(v4) > 0 {
        v4[0] = (v2 - v2)
    }
    fmt.Println(v2, v2)
    v2 += v2
    var v9 uint32 = (v2 ^ 0)
    v10 := v5[42]
    if true {
        var v11 float64 = 0.0
        _ = v11
    } else {
        switch -477 {
        case 9223372036854775807:
            v3 = (v3 * (v3 + v3))
        default:
            v12 := GenS1{F0: (230 == 42), F1: 127, F2: true}
            _ = v12
        }
    }
    v9 = v9
    var v13 uint16 = 65535
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v13, v2)
    for fi14 := 0; fi14 < 13; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            v2 = 2147483648
            if (false || false) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v16 float32 = (1.5 / 3.14159)
                fmt.Println(fi14, v10)
                _ = v16
            } else {
                var v17 int16 = -218
                v4 = append(v4, v2)
                _ = v17
            }
            switch v10 {
            case "a":
                _ = len(v4)
            case "\n\t":
                fmt.Println(v9, v13, fi14)
            }
            _ = fi15
        }
        v10 = v10
        _ = fi14
    }
    _ = len(v1)
    if (true && ((v13 | v13) != (v13 << ((1) & 63)))) {
        if true {
            switch -1 {
            case 42:
                switch "\n\t" {
                case "a":
                    fmt.Println(v10, v3)
                case "foo bar":
                    if (false && false) {
                        for fi18 := 0; fi18 < 15; fi18++ {
                            v2 += 2147483647
                            fmt.Println(v2)
                            _ = fi18
                        }
                    } else {
                        v19 := GenS1{F0: (-64 > 128), F1: 42, F2: false}
                        switch (v10 + v10) {
                        case "foo bar":
                            v13 = (^(v13))
                        }
                        _ = v19
                    }
                default:
                    var v20 int32 = 2147483647
                    _ = v20
                }
            case -3:
                delete(v5, -548)
            default:
                v13 -= v13
            }
        } else {
            if true {
                if ((true || (32767 >= 127)) && (1 == -1)) {
                    v2 = (v9 & 2)
                    switch -9223372036854775808 {
                    case 2147483647:
                        v10 = "a"
                    case -9223372036854775808:
                        var v21 int64 = 0
                        _ = v21
                    case 4294967295:
                        v22 := GenS1{F0: (((255 >= 1) && true) && false), F1: -441, F2: (!true)}
                        _ = v22
                    default:
                        if ((v10 + v10) <= v10) {
                            v13 = v13
                        }
                    }
                }
                v13 += 989
                v23 := GenS1{F0: ((!false) || (1 == -2147483648)), F1: 128, F2: false}
                _ = v23
            } else {
                if (v10 < (v10 + v10)) {
                    switch "a" {
                    case "":
                        v10 = v10
                    default:
                        if true {
                            fmt.Println(v3, v2)
                            var v24 int64 = 1
                            for fi25 := 0; fi25 < 8; fi25++ {
                                var v26 int = (fi25 | (^(fi25)))
                                _ = fi25
                                _ = v26
                            }
                            _ = v24
                        } else {
                            switch 127 {
                            case 4294967295:
                                v27 := GenS1{F0: (!false), F1: -9223372036854775808, F2: ((!(1 > 213)) || true)}
                                _ = v27
                            case -90:
                                fmt.Println(v10, v9)
                            case -924:
                                fmt.Println(v3)
                            }
                            if (0.1 < (-2.25 + 3.14159)) {
                                v3 += (v3 * v3)
                            }
                            for fi28 := 0; fi28 < 4; fi28++ {
                                v5[fi28] = v10
                                var v29 string = "\n\t"
                                _ = fi28
                                _ = v29
                            }
                        }
                    }
                } else {
                    v2 = 42
                    var v30 float64 = 0.0
                    _ = v30
                }
                v31 := GenS2{F0: 42}
                _ = v31
            }
            var v32 uint8 = 0
            _ = v32
        }
    } else {
        switch v10 {
        case "":
            v13 = v13
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
