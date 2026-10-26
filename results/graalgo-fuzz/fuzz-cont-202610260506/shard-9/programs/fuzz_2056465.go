package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint }

func main() {
    v1 := []int{32768, 32768, 128}
    v2 := make(map[int]string)
    var v3 uint16 = 23
    v4 := []int{}
    v5 := []string{""}
    v6 := []string{("a" + ("a" + ((("\n\t" + "\n\t") + "\n\t") + "hello")))}
    if (9223372036854775807 == -9223372036854775808) {
        fmt.Println(v3)
        if (true || (812 < 4294967295)) {
            if false {
                fmt.Println(v3)
            } else {
                v3 += 0
            }
        }
        v3 = 32767
    } else {
        v3 = v3
    }
    v3 -= (v3 * v3)
    for fi7 := 0; fi7 < 10; fi7++ {
        if (127 != -1) {
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 int32 = 681
                if ((1e10 >= 3.14159) && true) {
                    var v10 uint8 = 128
                    v9 = v9
                    switch "" {
                    case "\n\t":
                        var v11 uint16 = v3
                        _ = v11
                    case "foo bar":
                        _ = len(v5)
                    case "a":
                        fmt.Println(v3)
                    }
                    _ = v10
                }
                var v12 int64 = 9223372036854775807
                _ = fi8
                _ = v12
                _ = v9
            }
            v2[fi7] = "foo bar"
            var v13 int64 = -9223372036854775808
            _ = v13
        }
        var v14 uint32 = 2
        _ = fi7
        _ = v14
    }
    if true {
        var v15 string = ("" + ("" + ""))
        _ = v15
    } else {
        fmt.Println(v3)
    }
    var v16 int64 = 9223372036854775807
    v16 -= v16
    v16 = 1
    _ = len(v2)
    v3 -= v3
    switch -1 {
    case 565:
        v17 := GenS1{F0: 37}
        _ = v17
    case 538:
        if true {
            v16 = (v16 * 1)
            for fi18 := 0; fi18 < 9; fi18++ {
                switch -2147483648 {
                case -9223372036854775808:
                    _ = len(v6)
                case 2147483647:
                    _ = len(v1)
                case 0:
                    v16 += (v16 - v16)
                }
                _ = fi18
            }
            var v19 float64 = 1.5
            _ = v19
        }
    case 128:
        switch "hello" {
        case "a":
            fmt.Println(v3)
        case "":
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 uint8 = 255
                v16 = 0
                var v22 int = fi20
                _ = fi20
                _ = v21
                _ = v22
            }
        case "foo bar":
            _ = len(v2)
        default:
            var v23 float32 = (((1e10 + float32(((1e10 + float64(0.1)) + 3.14159))) - (0.1 + 3.14159)) * ((3.14159 * (-2.25 + (0.1 - 0.0))) * ((float32(1.5) / 1e10) + 3.14159)))
            _ = v23
        }
    }
    if false {
        v16 -= 0
        switch "\n\t" {
        case "foo bar":
            var v24 bool = (-2147483648 > 2147483647)
            _ = v24
        case "\n\t":
            v25 := GenS1{F0: 42}
            _ = v25
        default:
            if true {
                fmt.Println(v3, v16)
            } else {
                var v26 uint32 = 2
                for fi27 := 0; fi27 < 17; fi27++ {
                    v28 := GenS1{F0: -128}
                    for fi29 := 0; fi29 < 14; fi29++ {
                        v30 := GenS1{F0: -69}
                        _ = fi29
                        _ = v30
                    }
                    v16 -= v16
                    _ = fi27
                    _ = v28
                }
                _ = v26
            }
        }
        v16 -= (v16 + 934)
    } else {
        v31 := GenS2{F0: (v3 ^ (v3 ^ 255)), F1: 1, F2: uint(v3)}
        _ = len(v1)
        _ = v31
    }
    v16 += (v16 >> ((0) & 63))
    v16 -= v16
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
